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

	for ly := 0; ly < NO_PIXEL_HEIGHT; ly++ {

		// Load object buffer (OAM Scan)
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

		objFIFO := []int{}
		bgFIFO := []int{}
		fetcherX := 0
		fetcherY := 0

		// Draw on a scanline
		i := 0
		for posX := 0; posX < NO_PIXEL_WIDTH; i++ {

			// BG Tile Fetcher
			scX := int(c.Mem.Read(0xFF43))
			scY := int(c.Mem.Read(0xFF42))
			println(scX, scY)
			fetcherX = (scX/8 + i) & 0x1F
			fetcherY = (ly + scY) & 0xFF

			var tileAddr, tileIndex int
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

			for i := scX % 8; len(bgFIFO) < 8; i++ {
				bgFIFO = append(bgFIFO, int(2*((c.Mem.Read(tileAddr+1)>>(7-i))&1)+((c.Mem.Read(tileAddr)>>(7-i))&1)))
			}

			bgToObj := 0
			// xFlip := 0
			yFlip := 0

			// // Fetch OBJ pixel
			if len(objBuffer) != 0 {
				tileAddr := 0x8000 + 16*int(uint8(objBuffer[0][2])) + 2*(fetcherY%8)
				bgToObj = int((objBuffer[0][3] >> 7) & 1)
				yFlip = int((objBuffer[0][3] >> 6) & 1)
				objBuffer = objBuffer[1:]

				// 	// Load pixel into OBJ FIFO
				if yFlip == 0 {
					for i := 0; len(objFIFO) < 8; i++ {
						objFIFO = append(objFIFO, int(2*((c.Mem.Read(tileAddr+1)>>(7-i))&1)+((c.Mem.Read(tileAddr)>>(7-i))&1)))
					}
				} else {
					for i := 0; len(objFIFO) < 8; i++ {
						objFIFO = append(objFIFO, int(2*((c.Mem.Read(tileAddr+1)>>i)&1)+((c.Mem.Read(tileAddr)>>i)&1)))
					}
				}
			}

			// Load pixel into LCD
			for len(bgFIFO) != 0 {

				colorBG := palette[bgFIFO[0]]
				colorOBJ := [4]uint8{}
				bgFIFO = bgFIFO[1:]

				// Disable background
				if (c.Mem.Read(0xFF40))&1 == 0 {
					colorBG = palette[0]
				}

				if len(objFIFO) != 0 {
					colorOBJ = palette[objFIFO[0]]
					objFIFO = objFIFO[1:]
				}

				if len(objFIFO) != 0 {
					if objFIFO[0] == 0 || (bgFIFO[0] != 0 && bgToObj == 1) {
						d.renderer.SetDrawColor(colorBG[0], colorBG[1], colorBG[2], colorBG[3])
					} else {
						d.renderer.SetDrawColor(colorOBJ[0], colorOBJ[1], colorOBJ[2], colorOBJ[3])
					}
				} else {
					d.renderer.SetDrawColor(colorBG[0], colorBG[1], colorBG[2], colorBG[3])
				}

				d.renderer.DrawRect(d.pixels[ly][posX])
				d.renderer.FillRect(d.pixels[ly][posX])
				posX++
			}
		}
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
