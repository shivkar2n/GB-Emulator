package MMU

import (
	"github.com/veandco/go-sdl2/sdl"
)

func (MMU *MMU) PressDpad(direction string) {
	if direction == "RIGHT" {
		MMU.DpadState &= 0xE

	} else if direction == "LEFT" {
		MMU.DpadState &= 0xD

	} else if direction == "DOWN" {
		MMU.DpadState &= 0x7

	} else if direction == "UP" {
		MMU.DpadState &= 0xB
	}
	// MMU.SetIFBit(4)
}

func (MMU *MMU) PressButton(button string) {
	if button == "A" {
		MMU.ButtonState &= 0xE

	} else if button == "B" {
		MMU.ButtonState &= 0xD

	} else if button == "START" {
		MMU.ButtonState &= 0x7

	} else if button == "SELECT" {
		MMU.ButtonState &= 0xB
	}
	// MMU.SetIFBit(4)
}

func (MMU *MMU) ReleaseDpad(direction string) {
	if direction == "RIGHT" {
		MMU.DpadState |= 0x1

	} else if direction == "LEFT" {
		MMU.DpadState |= 0x2

	} else if direction == "DOWN" {
		MMU.DpadState |= 0x8

	} else if direction == "UP" {
		MMU.DpadState |= 0x4
	}
}

func (MMU *MMU) ReleaseButton(button string) {
	if button == "A" {
		MMU.ButtonState |= 0x1

	} else if button == "B" {
		MMU.ButtonState |= 0x2

	} else if button == "START" {
		MMU.ButtonState |= 0x8

	} else if button == "SELECT" {
		MMU.ButtonState |= 0x4
	}
}

func (MMU *MMU) PollJoyPadPress() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case sdl.KeyboardEvent:
			keyCode := t.Keysym.Sym
			if t.State == sdl.RELEASED {
				switch sdl.GetKeyName(keyCode) {
				case "Z":
					MMU.ReleaseButton("A")
					break

				case "X":
					MMU.ReleaseButton("B")
					break

				case "A":
					MMU.ReleaseButton("SELECT")
					break

				case "S":
					MMU.ReleaseButton("START")
					break

				case "Left":
					MMU.ReleaseDpad("LEFT")
					break

				case "Right":
					MMU.ReleaseDpad("RIGHT")
					break

				case "Up":
					MMU.ReleaseDpad("UP")
					break

				case "Down":
					MMU.ReleaseDpad("DOWN")
					break
				}

			} else if t.State == sdl.PRESSED {
				switch sdl.GetKeyName(keyCode) {
				case "Z":
					MMU.PressButton("A")
					break

				case "X":
					MMU.PressButton("B")
					break

				case "A":
					MMU.PressButton("SELECT")
					break

				case "S":
					MMU.PressButton("START")
					break

				case "Left":
					MMU.PressDpad("LEFT")
					break

				case "Right":
					MMU.PressDpad("RIGHT")
					break

				case "Up":
					MMU.PressDpad("UP")
					break

				case "Down":
					MMU.PressDpad("DOWN")
					break

				}
			}
		}
	}
}
