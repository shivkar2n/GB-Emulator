package PPU

import (
	"sort"

	"github.com/shivkar2n/GB-Emulator/pkg/CPU"
	"github.com/shivkar2n/GB-Emulator/pkg/MMU"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	WINDOW_WIDTH       = 800
	WINDOW_HEIGHT      = 720
	NO_PIXEL_WIDTH     = 160
	NO_PIXEL_HEIGHT    = 144
	NO_TILES_ROW       = 32
	NO_TILES_COL       = 32
	OBJECT_SIZE        = 8
	TALL_OBJECT_SIZE   = 16
	OBJECT_BUFFER_SIZE = 10
	BG_QUEUE_SIZE      = 8
	LCD_CTRL           = 0xFF40
	WX                 = 0xFF4B
	WY                 = 0xFF4A
	SCX                = 0xFF43
	SCY                = 0xFF42
	BGP                = 0xFF47
	OBP0               = 0xFF48
	OBP1               = 0xFF49
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

type PPU struct {
	Window   *sdl.Window
	Renderer *sdl.Renderer
	Pixels   [NO_PIXEL_HEIGHT][NO_PIXEL_WIDTH]*sdl.Rect

	CPU *CPU.CPU
	MMU *MMU.MMU

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

func Init(CPU *CPU.CPU, MMU *MMU.MMU) *PPU {
	sdl.Init(sdl.INIT_EVERYTHING)

	Window, _ := sdl.CreateWindow("GB-Emulator", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, WINDOW_WIDTH, WINDOW_HEIGHT, sdl.WINDOW_RESIZABLE)
	Renderer, _ := sdl.CreateRenderer(Window, -1, sdl.RENDERER_SOFTWARE)
	var Pixels [NO_PIXEL_HEIGHT][NO_PIXEL_WIDTH]*sdl.Rect

	for i := 0; i < NO_PIXEL_WIDTH; i++ {
		for j := 0; j < NO_PIXEL_HEIGHT; j++ {
			Pixels[j][i] = &sdl.Rect{int32(i * WINDOW_HEIGHT / NO_PIXEL_HEIGHT), int32(j * WINDOW_WIDTH / NO_PIXEL_WIDTH), WINDOW_WIDTH / NO_PIXEL_WIDTH, WINDOW_HEIGHT / NO_PIXEL_HEIGHT}
		}
	}

	return &PPU{
		Window:   Window,
		Renderer: Renderer,
		Pixels:   Pixels,
		CPU:      CPU,
		MMU:      MMU,
	}
}

func (PPU *PPU) ResetScanline() {
	PPU.ObjectBuffer = []Object{}
	PPU.FetcherX = 0
	PPU.FetcherY = 0
	PPU.WindowPositionX = 0
	PPU.BackgroundPositionX = 0
	PPU.CurrentPositionX = 0
	PPU.WindowMode = false
	PPU.IsBackgroundTile = true
}

func (PPU *PPU) FetchBackgroundTile() int {
	var tileAddr, tileIndex int
	wX := PPU.MMU.Read(WX)
	wY := PPU.MMU.Read(WY)
	PPU.IsBackgroundTile = !(PPU.WindowMode && (PPU.MMU.Read(LCD_CTRL)>>5)&1 == 1 && 0 <= wX && wX <= 166 && 0 <= wY && wY <= 143)

	if !PPU.IsBackgroundTile {
		PPU.FetcherX = PPU.WindowPositionX & 0x1F
		PPU.FetcherY = PPU.WindowPositionY & 0xFF

	} else {
		scX := int(PPU.MMU.Read(SCX))
		scY := int(PPU.MMU.Read(SCY))
		PPU.FetcherX = (scX/8 + PPU.BackgroundPositionX) & 0x1F
		PPU.FetcherY = (scY + PPU.BackgroundPositionY) & 0xFF
	}

	if ((PPU.MMU.Read(LCD_CTRL)>>6)&1 == 1 && !PPU.IsBackgroundTile) || ((PPU.MMU.Read(LCD_CTRL)>>3)&1 == 1 && PPU.IsBackgroundTile) {
		tileIndex = PPU.MMU.Read(0x9C00 + (NO_TILES_ROW*(PPU.FetcherY>>3)+PPU.FetcherX)&0x3FF)
	} else {
		tileIndex = PPU.MMU.Read(0x9800 + (NO_TILES_ROW*(PPU.FetcherY>>3)+PPU.FetcherX)&0x3FF)
	}

	if (PPU.MMU.Read(LCD_CTRL)>>4)&1 == 1 {
		tileAddr = 0x8000 + 16*int(uint8(tileIndex)) + 2*(PPU.FetcherY%8)
	} else {
		tileAddr = 0x9000 + 16*int(int8(tileIndex)) + 2*(PPU.FetcherY%8)
	}
	return tileAddr
}

func (PPU *PPU) BackgroundColor(id int) int {
	return (PPU.MMU.Read(BGP) >> (id * 2)) & 0x03
}

func (PPU *PPU) ObjectColor(id int, objectFlag int) int {
	if (objectFlag>>4)&1 == 1 {
		return (PPU.MMU.Read(OBP1) >> (id * 2)) & 0x03
	}
	return (PPU.MMU.Read(OBP0) >> (id * 2)) & 0x03
}

func (PPU *PPU) LoadObjectBuffer() {
	for addr := 0xFE00; addr < 0xFEA0; addr += 4 {
		spriteX := PPU.MMU.Read(addr + 1)
		spriteY := PPU.MMU.Read(addr)
		tileIndex := PPU.MMU.Read(addr + 2)
		flags := PPU.MMU.Read(addr + 3)
		objSize := OBJECT_SIZE
		if (PPU.MMU.Read(LCD_CTRL)>>2)&1 == 1 {
			objSize = TALL_OBJECT_SIZE
		}

		if spriteX > 0 && PPU.BackgroundPositionY+16 >= spriteY && PPU.BackgroundPositionY+16 < spriteY+objSize && len(PPU.ObjectBuffer) < OBJECT_BUFFER_SIZE {
			PPU.ObjectBuffer = append(PPU.ObjectBuffer, Object{spriteX, spriteY, tileIndex, flags})
		}
	}

	sort.SliceStable(PPU.ObjectBuffer, func(i, j int) bool {
		return PPU.ObjectBuffer[i].PositionX >= PPU.ObjectBuffer[j].PositionX
	})
}


func (PPU *PPU) IncrementCounter() {
	if !PPU.IsBackgroundTile { // Increment once when Window Pixels pushed to LCD for scanline
		PPU.WindowPositionY++
	}
	PPU.BackgroundPositionY++
	PPU.MMU.Write(PPU.BackgroundPositionY, 0xFF44)
	if PPU.BackgroundPositionY-PPU.ObjectPositionY >= 16 {
		PPU.ObjectPositionY = -1
	}
}

func (PPU *PPU) PushBackgroundPixelRow(tileAddr int) []int {
	wX := PPU.MMU.Read(WX)
	wY := PPU.MMU.Read(WY)
	var backgroundQueue []int

	for i := 0; len(backgroundQueue) < BG_QUEUE_SIZE; i++ {
		if (PPU.MMU.Read(LCD_CTRL))&1 == 0 { // Disable background
			backgroundQueue = append(backgroundQueue, 0)

		} else if PPU.IsBackgroundTile && (PPU.MMU.Read(LCD_CTRL)>>5)&1 == 1 && PPU.BackgroundPositionY >= wY && PPU.CurrentPositionX+i >= (wX-7) { // Window fetching
			PPU.WindowMode = true
			backgroundQueue = nil
			break

		} else { // Push pixel to BGFIFO
			backgroundQueue = append(backgroundQueue, int(2*((PPU.MMU.Read(tileAddr+1)>>(7-i))&1)+((PPU.MMU.Read(tileAddr)>>(7-i))&1)))
		}
	}

	if len(backgroundQueue) != 0 {
		if PPU.IsBackgroundTile {
			PPU.BackgroundPositionX++
		} else {
			PPU.WindowPositionX++
		}
	}
	return backgroundQueue
}

func (PPU *PPU) PushObjectPixel() [2]int {
	objectPixel := [2]int{0, 0} // Push transparent pixel with lowest priority

	for _, object := range PPU.ObjectBuffer {
		tileIndex := object.TileIndex
		tileAddr := 0

		if object.PositionX <= PPU.CurrentPositionX+8 {
			val := PPU.CurrentPositionX - object.PositionX + 8

			if (PPU.MMU.Read(LCD_CTRL)>>2)&1 == 1 || PPU.ObjectPositionY != -1 { // Tall Object Mode
				if PPU.ObjectPositionY == -1 {
					PPU.ObjectPositionY = PPU.BackgroundPositionY
				}
				tileIndex &= 0xFE
				if (PPU.BackgroundPositionY-PPU.ObjectPositionY < 8 && (object.Flags>>6)&1 == 1) || (PPU.BackgroundPositionY-PPU.ObjectPositionY >= 8 && (object.Flags>>6)&1 == 0) { // Vertically flip Object
					tileIndex += 1
				}
			}

			if (object.Flags>>6)&1 == 1 { // Vertically flip object
				tileAddr = 0x8000 + 16*int(uint8(tileIndex)) + 2*(7-(PPU.BackgroundPositionY%8))
			} else {
				tileAddr = 0x8000 + 16*int(uint8(tileIndex)) + 2*(PPU.BackgroundPositionY%8)
			}

			if 0 <= val && val < 8 {
				if (object.Flags>>5)&1 == 1 { // Horizontally flip object
					objectPixel = [2]int{(2*((PPU.MMU.Read(tileAddr+1)>>val)&1) + ((PPU.MMU.Read(tileAddr) >> val) & 1)), object.Flags}

				} else {
					objectPixel = [2]int{(2*((PPU.MMU.Read(tileAddr+1)>>(7-val))&1) + ((PPU.MMU.Read(tileAddr) >> (7 - val)) & 1)), object.Flags}
				}
			}
		}
	}
	return objectPixel
}

func (PPU *PPU) PushPixelToLCD(backgroundQueue []int) {
	for len(backgroundQueue) != 0 && PPU.CurrentPositionX < NO_PIXEL_WIDTH { // Push pixel to LCD
		colorBG := palette[PPU.BackgroundColor(backgroundQueue[0])]
		objectQueue := PPU.PushObjectPixel()

		if len(objectQueue) != 0 { // Mixing of object and background Pixels
			colorOBJ := palette[PPU.ObjectColor(objectQueue[0], objectQueue[1])]

			if objectQueue[0] == 0 { // If object pixel is transparent
				PPU.Renderer.SetDrawColor(colorBG[0], colorBG[1], colorBG[2], colorBG[3])

			} else if (objectQueue[1]>>7)&1 == 1 && backgroundQueue[0] != 0 {
				PPU.Renderer.SetDrawColor(colorBG[0], colorBG[1], colorBG[2], colorBG[3])

			} else if (PPU.MMU.Read(LCD_CTRL)>>1)&1 == 1 { // Check if objects can be PPUed
				PPU.Renderer.SetDrawColor(colorOBJ[0], colorOBJ[1], colorOBJ[2], colorOBJ[3])
			}

		} else {
			PPU.Renderer.SetDrawColor(colorBG[0], colorBG[1], colorBG[2], colorBG[3])
		}

		PPU.Renderer.DrawRect(PPU.Pixels[PPU.BackgroundPositionY][PPU.CurrentPositionX])
		PPU.Renderer.FillRect(PPU.Pixels[PPU.BackgroundPositionY][PPU.CurrentPositionX])
		backgroundQueue = backgroundQueue[1:]
		PPU.CurrentPositionX++ // Increment Internal-X-Position-Counter
	}
}

func (PPU *PPU) DrawScanline() {
	for PPU.CurrentPositionX < NO_PIXEL_WIDTH {
		tileAddr := PPU.FetchBackgroundTile()
		backgroundQueue := PPU.PushBackgroundPixelRow(tileAddr)
		PPU.PushPixelToLCD(backgroundQueue)
	}
}

func (PPU *PPU) Free() {
	PPU.Window = nil
	PPU.Renderer = nil
	for i := 0; i < NO_PIXEL_HEIGHT; i++ {
		for j := 0; j < NO_PIXEL_WIDTH; j++ {
			PPU.Pixels[i][j] = nil
		}
	}
}
