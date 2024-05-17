package GB

import (
	// "fmt"

	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	JOYP = 0xFF00
)

func (GB *GB) PressDpad(direction string) {
	GB.MMU.Write(GB.MMU.Read(JOYP)&0x2F, JOYP)
	if direction == "RIGHT" {
		GB.MMU.Write(GB.MMU.Read(JOYP)&0xFE, JOYP)

	} else if direction == "LEFT" {
		GB.MMU.Write(GB.MMU.Read(JOYP)&0xFD, JOYP)

	} else if direction == "DOWN" {
		GB.MMU.Write(GB.MMU.Read(JOYP)&0xF7, JOYP)

	} else if direction == "UP" {
		GB.MMU.Write(GB.MMU.Read(JOYP)&0xFB, JOYP)
	}
	GB.SetIFBit(4)
}

func (GB *GB) PressButton(button string) {
	GB.MMU.Write(GB.MMU.Read(JOYP)&0x1F, JOYP)
	if button == "A" {
		GB.MMU.Write(GB.MMU.Read(JOYP)&0xFE, JOYP)

	} else if button == "B" {
		GB.MMU.Write(GB.MMU.Read(JOYP)&0xFD, JOYP)

	} else if button == "START" {
		GB.MMU.Write(GB.MMU.Read(JOYP)&0xF7, JOYP)

	} else if button == "SELECT" {
		GB.MMU.Write(GB.MMU.Read(JOYP)&0xFB, JOYP)
	}
	GB.SetIFBit(4)
}

func (GB *GB) ReleaseDpad(direction string) {
	GB.MMU.Write(GB.MMU.Read(JOYP)|0x1F, JOYP)
	if direction == "RIGHT" {
		GB.MMU.Write(GB.MMU.Read(JOYP)|0x01, JOYP)

	} else if direction == "LEFT" {
		GB.MMU.Write(GB.MMU.Read(JOYP)|0x02, JOYP)

	} else if direction == "DOWN" {
		GB.MMU.Write(GB.MMU.Read(JOYP)|0x08, JOYP)

	} else if direction == "UP" {
		GB.MMU.Write(GB.MMU.Read(JOYP)|0x04, JOYP)
	}
}

func (GB *GB) ReleaseButton(button string) {
	GB.MMU.Write(GB.MMU.Read(JOYP)|0x2F, JOYP)
	if button == "A" {
		GB.MMU.Write(GB.MMU.Read(JOYP)|0x01, JOYP)

	} else if button == "B" {
		GB.MMU.Write(GB.MMU.Read(JOYP)|0x02, JOYP)

	} else if button == "START" {
		GB.MMU.Write(GB.MMU.Read(JOYP)|0x08, JOYP)

	} else if button == "SELECT" {
		GB.MMU.Write(GB.MMU.Read(JOYP)|0x04, JOYP)
	}
}

func (GB *GB) PollJoyPadPress() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case sdl.KeyboardEvent:
			keyCode := t.Keysym.Sym
			if t.State == sdl.RELEASED {
				// fmt.Printf("BEFORE RELEASE EVENT:%02X %02X %02X\n", GB.MMU.Read(JOYP), GB.MMU.Read(IF), GB.MMU.Read(IE))
				switch sdl.GetKeyName(keyCode) {
				case "Z":
					GB.ReleaseButton("A")
					break

				case "X":
					GB.ReleaseButton("B")
					break

				case "A":
					GB.ReleaseButton("SELECT")
					break

				case "S":
					GB.ReleaseButton("START")
					break

				case "Left":
					GB.ReleaseDpad("LEFT")
					break

				case "Right":
					GB.ReleaseDpad("RIGHT")
					break

				case "Up":
					GB.ReleaseDpad("UP")
					break

				case "Down":
					GB.ReleaseDpad("DOWN")
					break
				}
				// fmt.Printf("AFTER RELEASE EVENT:%02X %02X\n", GB.MMU.Read(JOYP), GB.MMU.Read(IF))

			} else if t.State == sdl.PRESSED {
				fmt.Printf("BEFORE PRESS EVENT:%02X %02X %02X\n", GB.MMU.Read(JOYP), GB.MMU.Read(IF), GB.MMU.Read(IE))
				switch sdl.GetKeyName(keyCode) {
				case "Z":
					GB.PressButton("A")
					break

				case "X":
					GB.PressButton("B")
					break

				case "A":
					GB.PressButton("SELECT")
					break

				case "S":
					GB.PressButton("START")
					break

				case "Left":
					GB.PressDpad("LEFT")
					break

				case "Right":
					GB.PressDpad("RIGHT")
					break

				case "Up":
					GB.PressDpad("UP")
					break

				case "Down":
					GB.PressDpad("DOWN")
					break

				}
				fmt.Printf("AFTER PRESS EVENT: %02X %02X %02X\n", GB.MMU.Read(JOYP), GB.MMU.Read(IF), GB.MMU.Read(IE))
			}
		}
	}
}
