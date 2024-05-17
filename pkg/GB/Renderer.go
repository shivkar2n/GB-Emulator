package GB

import "github.com/veandco/go-sdl2/sdl"

const (
	LCD_STAT          = 0xFF41
	LYC               = 0xFF45
	NO_SCANLINES      = 154
	NO_REAL_SCANLINES = 144
)

const (
	HBLANK  = 0
	VBLANK  = 1
	OAM     = 2
	DRAWING = 3
)

func (GB *GB) StatInterrupt() {
	var condition bool
	condition = false

	if GB.MMU.Read(LCD_STAT)&0x3 == HBLANK && GB.MMU.Read(LCD_STAT)>>3&1 == 1 {
		condition = true

	} else if GB.MMU.Read(LCD_STAT)&0x3 == VBLANK && GB.MMU.Read(LCD_STAT)>>4&1 == 1 {
		condition = true

	} else if GB.MMU.Read(LCD_STAT)&0x3 == OAM && GB.MMU.Read(LCD_STAT)>>5&1 == 1 {
		condition = true

	} else if GB.MMU.Read(LCD_STAT)&0x3 == DRAWING {
		condition = true

	} else if (GB.MMU.Read(LCD_STAT)>>6)&1 == 1 && int(GB.MMU.Read(LYC)) == GB.PPU.BackgroundPositionY {
		condition = true
	}

	if condition {
		GB.SetIFBit(1)
	}
}

func (GB *GB) RenderScanline() {

	if GB.PPU.BackgroundPositionY < NO_REAL_SCANLINES {
		GB.PPU.ResetScanline()
		GB.PPU.LoadObjectBuffer()
		GB.PPU.DrawScanline()
		GB.PPU.IncrementCounter()
		GB.StatInterrupt()

		// VBLANK MODE
	} else if NO_REAL_SCANLINES <= GB.PPU.BackgroundPositionY && GB.PPU.BackgroundPositionY < NO_SCANLINES {
		GB.PPU.IncrementCounter()
		GB.SetIFBit(0) // Set V-Blank interrupt

	} else {
		// Reset PPU Values
		GB.MMU.Write(0xFF44, 0)
		GB.PPU.BackgroundPositionY = 0
		GB.PPU.WindowPositionY = 0
		GB.PPU.ObjectPositionY = -1

		// Render frame once all scanlines are drawn
		GB.PPU.Frame++
		GB.PPU.Renderer.Present()
	}
}

func (GB *GB) FrameDelay() {
	// Delay between frame
	GB.PPU.Frame = GB.PPU.Frame % FRAME_RATE
	sdl.Delay(0)
}
