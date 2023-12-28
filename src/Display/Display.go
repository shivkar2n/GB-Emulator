package Display

import (
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

type Display struct {
	window   *sdl.Window
	renderer *sdl.Renderer
	pixels   [NO_PIXEL_HEIGHT][NO_PIXEL_WIDTH]*sdl.Rect
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

func (d *Display) RenderFrame(c *CPU.CPU) {
	d.renderer.Clear()

	for ly := 0; ly < NO_PIXEL_HEIGHT; {

		// STAT interrupt if ly==lyc
		if (c.Mem.Read(0xFF41)>>6)&1 == 1 && int(c.Mem.Read(0xFF45)) == ly {
			c.SetIFBit(1)
			c.InterruptHandler()
		}

		// Load object buffer (OAM Scan - Mode 2)
		objBuffer := [][4]int{}
		for obj := 0xFE00; obj < 0xFEA0; obj += 4 {
			spriteX := int(c.Mem.Read(obj + 1))
			spriteY := int(c.Mem.Read(obj))
			tileIndex := int(c.Mem.Read(obj + 2))
			flags := int(c.Mem.Read(obj + 3))

			if int(spriteX-8) > 0 && ly+16 >= spriteY && ly+16 < spriteY+8 && len(objBuffer) < 10 {
				objBuffer = append(objBuffer, [4]int{spriteX, spriteY, tileIndex, flags})
			}
		}

		bgFIFO := []int{}
		fetcherX := 0
		fetcherY := 0
		tileAddr := 0
		tileIndex := 0

		scX := int(c.Mem.Read(0xFF43))
		scY := int(c.Mem.Read(0xFF42))

		// Draw on a scanline (Drawing - Mode 3)
		ixpc := 0
		posX := 0
		for x := 0; x < NO_PIXEL_WIDTH; {

			// BG Tile Fetcher
			fetcherX = (scX/8 + ixpc) & 0x1F
			fetcherY = (ly + scY) & 0xFF

			if (c.Mem.Read(0xFF40)>>3)&1 == 1 {
				tileIndex = int(c.Mem.Read(0x9C00 + NO_TILES_ROW*(fetcherY>>3) + fetcherX))
			} else {
				tileIndex = int(c.Mem.Read(0x9800 + NO_TILES_ROW*(fetcherY>>3) + fetcherX))
			}

			if (c.Mem.Read(0xFF40)>>4)&1 == 1 {
				tileAddr = 0x8000 + 16*int(uint8(tileIndex)) + 2*((ly+scY)%8)
			} else {
				tileAddr = 0x9000 + 16*int(int8(tileIndex)) + 2*((ly+scY)%8)
			}

			for i := scX; i < 8 && len(bgFIFO) < 8; i++ {
				// Disable background
				if (c.Mem.Read(0xFF40))&1 == 0 {
					bgFIFO = append(bgFIFO, 0)
				} else {
					bgFIFO = append(bgFIFO, int(2*((c.Mem.Read(tileAddr+1)>>(7-i))&1)+((c.Mem.Read(tileAddr)>>(7-i))&1)))
				}
				posX++
			}

			// Load pixel into LCD
			if len(bgFIFO) != 0 {
				ixpc++
			}
			for len(bgFIFO) != 0 {
				colorBG := palette[bgFIFO[0]]
				bgFIFO = bgFIFO[1:]
				d.renderer.SetDrawColor(colorBG[0], colorBG[1], colorBG[2], colorBG[3])

				d.renderer.DrawRect(d.pixels[ly][x])
				d.renderer.FillRect(d.pixels[ly][x])
				x++
			}
		}
		ly++
		c.Mem.Write(ly, 0xFF44)
	}
	sdl.Delay(0)
	d.renderer.Present()
}

func (d *Display) Free() {
	d.window = nil
	d.renderer = nil
	for i := 0; i < NO_PIXEL_HEIGHT; i++ {
		for j := 0; j < NO_PIXEL_WIDTH; j++ {
			d.pixels[i][j] = nil
		}
	}
}
