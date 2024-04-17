package GB

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	LCD_STAT        = 0xFF41
	LYC             = 0xFF45
	NO_PIXEL_HEIGHT = 144
)

func (GB *GB) StatInterrupt() {
	if (GB.MMU.Read(LCD_STAT)>>6)&1 == 1 && int(GB.MMU.Read(LYC)) == GB.PPU.BackgroundPositionY {
		GB.SetIFBit(1)
		GB.InterruptHandler(true)
		for !GB.CPU.IME {
			_, GB.Length, GB.TCycles, _ = GB.ExecuteOpcode()
			GB.IncrementTimer()
			GB.CPU.IncrementCounter(GB.Length)
		}
	}
}
func (GB *GB) RenderFrame() {
	GB.PPU.Renderer.Clear()
	// Initialize PPU values
	GB.PPU.BackgroundPositionY = 0
	GB.PPU.WindowPositionY = 0
	GB.PPU.ObjectPositionY = -1

	// Draw scanlines
	for GB.PPU.BackgroundPositionY < NO_PIXEL_HEIGHT {
		GB.PPU.ResetScanline()
		GB.PPU.LoadObjectBuffer()
		GB.PPU.DrawScanline()
		GB.PPU.IncrementCounter()
		GB.StatInterrupt()

	}
	// Render drawn frame
	sdl.Delay(0)
	GB.PPU.Renderer.Present()
}
