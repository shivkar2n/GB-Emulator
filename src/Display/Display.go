package Display

import (
	"sort"

	"github.com/shivkar2n/GB-Emulator/CPU"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	WINDOW_WIDTH    = 800
	WINDOW_HEIGHT   = 720
	NO_PIXEL_WIDTH  = 160
	NO_PIXEL_HEIGHT = 144
	NO_TILES_ROW    = 32
	NO_TILES_COL    = 32
)

var palette = [4][4]uint8{
	{202, 220, 159, sdl.ALPHA_OPAQUE},
	{139, 172, 15, sdl.ALPHA_OPAQUE},
	{48, 98, 48, sdl.ALPHA_OPAQUE},
	{15, 56, 15, sdl.ALPHA_OPAQUE},
}

type Object struct {
	PositionX int
	PositionY int
	TileIndex int
	Flags     int
}

type Display struct {
	window   *sdl.Window
	renderer *sdl.Renderer
	pixels   [NO_PIXEL_HEIGHT][NO_PIXEL_WIDTH]*sdl.Rect

	ObjectBuffer        []Object
	FetcherX            int
	FetcherY            int
	ObjectPositionY     int
	WindowPositionX     int
	WindowPositionY     int // Window Line Counter
	BackgroundPositionX int
	BackgroundPositionY int // ly
	CurrentPositionX    int // Internal-X-Pos-Counter
	WindowMode          bool
	IsBackgroundTile    bool
}

func Init() *Display {
	sdl.Init(sdl.INIT_EVERYTHING)

	window, _ := sdl.CreateWindow("GB-Emulator", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, WINDOW_WIDTH, WINDOW_HEIGHT, sdl.WINDOW_RESIZABLE)
	renderer, _ := sdl.CreateRenderer(window, -1, sdl.RENDERER_SOFTWARE)
	var pixels [NO_PIXEL_HEIGHT][NO_PIXEL_WIDTH]*sdl.Rect

	for i := 0; i < NO_PIXEL_WIDTH; i++ {
		for j := 0; j < NO_PIXEL_HEIGHT; j++ {
			pixels[j][i] = &sdl.Rect{int32(i * WINDOW_HEIGHT / NO_PIXEL_HEIGHT), int32(j * WINDOW_WIDTH / NO_PIXEL_WIDTH), WINDOW_WIDTH / NO_PIXEL_WIDTH, WINDOW_HEIGHT / NO_PIXEL_HEIGHT}
		}
	}

	return &Display{
		window:   window,
		renderer: renderer,
		pixels:   pixels,
	}
}

func (display *Display) ResetScanline() {
	display.ObjectBuffer = []Object{}
	display.FetcherX = 0
	display.FetcherY = 0
	display.WindowPositionX = 0
	display.BackgroundPositionX = 0
	display.CurrentPositionX = 0
	display.WindowMode = false
	display.IsBackgroundTile = true
}

func (display *Display) FetchBackgroundTile(cpu *CPU.CPU) int {
	var tileAddr, tileIndex int
	wX := cpu.Mem.Read(0xFF4B)
	wY := cpu.Mem.Read(0xFF4A)
	display.IsBackgroundTile = !(display.WindowMode && (cpu.Mem.Read(0xFF40)>>5)&1 == 1 && 0 <= wX && wX <= 166 && 0 <= wY && wY <= 143)

	if !display.IsBackgroundTile {
		display.FetcherX = display.WindowPositionX & 0x1F
		display.FetcherY = display.WindowPositionY & 0xFF

	} else {
		scX := int(cpu.Mem.Read(0xFF43))
		scY := int(cpu.Mem.Read(0xFF42))
		display.FetcherX = (scX/8 + display.BackgroundPositionX) & 0x1F
		display.FetcherY = (scY + display.BackgroundPositionY) & 0xFF
	}

	if ((cpu.Mem.Read(0xFF40)>>6)&1 == 1 && !display.IsBackgroundTile) || ((cpu.Mem.Read(0xFF40)>>3)&1 == 1 && display.IsBackgroundTile) {
		tileIndex = cpu.Mem.Read(0x9C00 + (NO_TILES_ROW*(display.FetcherY>>3)+display.FetcherX)&0x3FF)
	} else {
		tileIndex = cpu.Mem.Read(0x9800 + (NO_TILES_ROW*(display.FetcherY>>3)+display.FetcherX)&0x3FF)
	}

	if (cpu.Mem.Read(0xFF40)>>4)&1 == 1 {
		tileAddr = 0x8000 + 16*int(uint8(tileIndex)) + 2*(display.FetcherY%8)
	} else {
		tileAddr = 0x9000 + 16*int(int8(tileIndex)) + 2*(display.FetcherY%8)
	}
	return tileAddr
}

func (display *Display) BackgroundColor(cpu *CPU.CPU, id int) int {
	return (cpu.Mem.Read(0xFF47) >> (id * 2)) & 0x03
}

func (display *Display) ObjectColor(cpu *CPU.CPU, id int, objectFlag int) int {
	if (objectFlag>>4)&1 == 1 {
		return (cpu.Mem.Read(0xFF49) >> (id * 2)) & 0x03
	}
	return (cpu.Mem.Read(0xFF48) >> (id * 2)) & 0x03
}

func (display *Display) LoadObjectBuffer(cpu *CPU.CPU) {
	for addr := 0xFE00; addr < 0xFEA0; addr += 4 {
		spriteX := cpu.Mem.Read(addr + 1)
		spriteY := cpu.Mem.Read(addr)
		tileIndex := cpu.Mem.Read(addr + 2)
		flags := cpu.Mem.Read(addr + 3)
		objSize := 8
		if (cpu.Mem.Read(0xFF40)>>2)&1 == 1 {
			objSize = 16
		}

		if spriteX > 0 && display.BackgroundPositionY+16 >= spriteY && display.BackgroundPositionY+16 < spriteY+objSize && len(display.ObjectBuffer) < 10 {
			display.ObjectBuffer = append(display.ObjectBuffer, Object{spriteX, spriteY, tileIndex, flags})
		}
	}

	sort.SliceStable(display.ObjectBuffer, func(i, j int) bool {
		return display.ObjectBuffer[i].PositionX >= display.ObjectBuffer[j].PositionX
	})
}

func (display *Display) StatInterrupt(cpu *CPU.CPU) {
	if (cpu.Mem.Read(0xFF41)>>6)&1 == 1 && int(cpu.Mem.Read(0xFF45)) == display.BackgroundPositionY {
		cpu.SetIFBit(1)
		cpu.InterruptHandler(true)
		for !cpu.IME {
			_, length, time, _ := cpu.ExecuteOpcode()
			cpu.IncrementTimer(time)
			cpu.IncrementCounter(length)
		}
	}
}

func (display *Display) IncrementCounter(cpu *CPU.CPU) {
	if !display.IsBackgroundTile { // Increment once when window pixels pushed to LCD for scanline
		display.WindowPositionY++
	}
	display.BackgroundPositionY++
	cpu.Mem.Write(display.BackgroundPositionY, 0xFF44)
	if display.BackgroundPositionY-display.ObjectPositionY >= 16 {
		display.ObjectPositionY = -1
	}
}

func (display *Display) PushBackgroundPixelRow(cpu *CPU.CPU, tileAddr int) []int {
	wX := cpu.Mem.Read(0xFF4B)
	wY := cpu.Mem.Read(0xFF4A)
	var backgroundQueue []int

	for i := 0; len(backgroundQueue) < 8; i++ {
		if (cpu.Mem.Read(0xFF40))&1 == 0 { // Disable background
			backgroundQueue = append(backgroundQueue, 0)

		} else if display.IsBackgroundTile && (cpu.Mem.Read(0xFF40)>>5)&1 == 1 && display.BackgroundPositionY >= wY && display.CurrentPositionX+i >= (wX-7) { // Window fetching
			display.WindowMode = true
			backgroundQueue = nil
			break

		} else { // Push pixel to BGFIFO
			backgroundQueue = append(backgroundQueue, int(2*((cpu.Mem.Read(tileAddr+1)>>(7-i))&1)+((cpu.Mem.Read(tileAddr)>>(7-i))&1)))
		}
	}

	if len(backgroundQueue) != 0 {
		if display.IsBackgroundTile {
			display.BackgroundPositionX++
		} else {
			display.WindowPositionX++
		}
	}
	return backgroundQueue
}

func (display *Display) PushObjectPixel(cpu *CPU.CPU) [2]int {
	objectPixel := [2]int{0, 0} // Push transparent pixel with lowest priority

	for _, object := range display.ObjectBuffer {
		tileIndex := object.TileIndex
		tileAddr := 0

		if object.PositionX <= display.CurrentPositionX+8 {
			val := display.CurrentPositionX - object.PositionX + 8

			if (cpu.Mem.Read(0xFF40)>>2)&1 == 1 || display.ObjectPositionY != -1 { // Tall Object Mode
				if display.ObjectPositionY == -1 {
					display.ObjectPositionY = display.BackgroundPositionY
				}
				tileIndex &= 0xFE
				if (display.BackgroundPositionY-display.ObjectPositionY < 8 && (object.Flags>>6)&1 == 1) || (display.BackgroundPositionY-display.ObjectPositionY >= 8 && (object.Flags>>6)&1 == 0) { // Vertically flip Object
					tileIndex += 1
				}
			}

			if (object.Flags>>6)&1 == 1 { // Vertically flip object
				tileAddr = 0x8000 + 16*int(uint8(tileIndex)) + 2*(7-(display.BackgroundPositionY%8))
			} else {
				tileAddr = 0x8000 + 16*int(uint8(tileIndex)) + 2*(display.BackgroundPositionY%8)
			}

			if 0 <= val && val < 8 {
				if (object.Flags>>5)&1 == 1 { // Horizontally flip object
					objectPixel = [2]int{(2*((cpu.Mem.Read(tileAddr+1)>>val)&1) + ((cpu.Mem.Read(tileAddr) >> val) & 1)), object.Flags}

				} else {
					objectPixel = [2]int{(2*((cpu.Mem.Read(tileAddr+1)>>(7-val))&1) + ((cpu.Mem.Read(tileAddr) >> (7 - val)) & 1)), object.Flags}
				}
			}
		}
	}
	return objectPixel
}

func (display *Display) DrawScanline(cpu *CPU.CPU) {

	for display.CurrentPositionX < NO_PIXEL_WIDTH {
		tileAddr := display.FetchBackgroundTile(cpu)
		backgroundQueue := display.PushBackgroundPixelRow(cpu, tileAddr)

		for len(backgroundQueue) != 0 && display.CurrentPositionX < NO_PIXEL_WIDTH { // Push pixel to LCD
			colorBG := palette[display.BackgroundColor(cpu, backgroundQueue[0])]
			objectQueue := display.PushObjectPixel(cpu)

			if len(objectQueue) != 0 { // Mixing of object and background pixels
				colorOBJ := palette[display.ObjectColor(cpu, objectQueue[0], objectQueue[1])]

				if objectQueue[0] == 0 { // If object pixel is transparent
					display.renderer.SetDrawColor(colorBG[0], colorBG[1], colorBG[2], colorBG[3])

				} else if (objectQueue[1]>>7)&1 == 1 && backgroundQueue[0] != 0 {
					display.renderer.SetDrawColor(colorBG[0], colorBG[1], colorBG[2], colorBG[3])

				} else if (cpu.Mem.Read(0xFF40)>>1)&1 == 1 { // Check if objects can be displayed
					display.renderer.SetDrawColor(colorOBJ[0], colorOBJ[1], colorOBJ[2], colorOBJ[3])
				}

			} else {
				display.renderer.SetDrawColor(colorBG[0], colorBG[1], colorBG[2], colorBG[3])
			}

			display.renderer.DrawRect(display.pixels[display.BackgroundPositionY][display.CurrentPositionX])
			display.renderer.FillRect(display.pixels[display.BackgroundPositionY][display.CurrentPositionX])
			backgroundQueue = backgroundQueue[1:]
			display.CurrentPositionX++ // Increment Internal-X-Position-Counter
		}
	}
}

func (display *Display) RenderFrame(cpu *CPU.CPU) {
	display.renderer.Clear()
	display.BackgroundPositionY = 0
	display.WindowPositionY = 0
	display.ObjectPositionY = -1

	for display.BackgroundPositionY < NO_PIXEL_HEIGHT {
		display.ResetScanline()
		display.LoadObjectBuffer(cpu)
		display.DrawScanline(cpu)
		display.IncrementCounter(cpu)
		display.StatInterrupt(cpu)

	}
	sdl.Delay(0)
	display.renderer.Present()
}

func (display *Display) Free() {
	display.window = nil
	display.renderer = nil
	for i := 0; i < NO_PIXEL_HEIGHT; i++ {
		for j := 0; j < NO_PIXEL_WIDTH; j++ {
			display.pixels[i][j] = nil
		}
	}
}
