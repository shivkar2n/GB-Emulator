package Display

import (
	"github.com/shivkar2n/GB-Emulator/CPU"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	WINDOW_WIDTH    = 800
	WINDOW_HEIGHT   = 800
	NO_PIXEL_WIDTH  = 256
	NO_PIXEL_HEIGHT = 256
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

	for j := 0; j < NO_PIXEL_HEIGHT; j++ {
		for i := 0; i < NO_PIXEL_WIDTH; i++ {
			pixels[i][j] = &sdl.Rect{int32(i * WINDOW_WIDTH / NO_PIXEL_WIDTH), int32(j * WINDOW_HEIGHT / NO_PIXEL_HEIGHT), WINDOW_WIDTH / NO_PIXEL_WIDTH, WINDOW_HEIGHT / NO_PIXEL_HEIGHT}
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

	tileIndices := [NO_TILES_COL][NO_TILES_ROW]byte{}
	mapOffset := 0x9800

	for j := 0; j < NO_TILES_COL; j++ {
		for i := 0; i < NO_TILES_ROW; i++ {
			addr := int(uint16(mapOffset + j*NO_TILES_ROW + i))
			tileIndices[i][j] = c.Mem.Read(addr)
		}
	}

	for j := 0; j < NO_TILES_COL; j++ {
		for i := 0; i < NO_TILES_ROW; i++ {
			var base int
			if (c.Mem.Read(0xFF40)>>4)&1 == 1 {
				base = int(uint16(0x8000 + 16*int(uint8(tileIndices[i][j]))))
			} else {
				base = int(uint16(0x9000 + 16*int(int8(tileIndices[i][j]))))
			}

			pixelData := [8][8]int{}

			for k := 0; k < 16; k += 2 {
				for l := 0; l < 8; l += 1 {
					pixelData[k/2][l] = 2*((int(c.Mem.Read(base+k+1))>>l)&1) + (int(c.Mem.Read(base+k))>>l)&1
				}
			}

			for l := 0; l < 8; l++ {
				for k := 0; k < 8; k++ {
					color := palette[pixelData[k][l]]
					d.renderer.SetDrawColor(color[0], color[1], color[2], color[3])
					d.renderer.DrawRect(d.pixels[8*i+7-l][8*j+k])
					d.renderer.FillRect(d.pixels[8*i+7-l][8*j+k])
				}
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
