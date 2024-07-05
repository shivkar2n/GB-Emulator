package MMU

func (MMU *MMU) Read(addr int) int {
	switch MMU.Ram[CARTRIDGE_TYPE] {
	case MBC0:
		{
			if addr == SERIAL { // Stub for serial COMS
				return int(0xFF)
			}
			return int(MMU.Ram[addr])
		}

	case MBC1:
		{
			if addr >= 0x0000 && addr <= 0x3FFF {
				return int(MMU.Rom[addr])

			} else if addr >= 0x4000 && addr <= 0x7FFF {
				if len(MMU.Rom) < 1024*1024 {
					MMU.highBank = MMU.romBank
				}
				return int(MMU.Rom[0x4000*MMU.highBank+addr-0x4000])

			} else if addr >= 0x8000 && addr <= 0x9FFF {
				return int(MMU.Ram[addr])

			} else if addr >= 0xA000 && addr <= 0xBFFF {
				if MMU.enableExtRam && MMU.Rom[0x147] == 0x02 {
					MMU.highBank = MMU.romBank
					return int(MMU.Rom[0x4000*MMU.highBank+addr-0x4000])
				}
				return 0xFF

			} else if addr >= 0xC000 && addr <= 0xFFFF {
				if addr == 0xFF4D {
					return 0xFF
				}
				return int(MMU.Ram[addr])
			}
		}
	}
	return 0
}

func (MMU *MMU) Write(val int, addr int) {
	switch MMU.Ram[CARTRIDGE_TYPE] {
	case MBC0:
		{
			if addr == JOYP {
				if val == 0x30 {
					MMU.Ram[addr] = byte(0xFF)

				} else if val == 0x20 { // Select dpad
					MMU.Ram[addr] = (MMU.Ram[addr] & 0xF0) | MMU.DpadState

				} else if val == 0x10 { // Select buttons
					MMU.Ram[addr] = (MMU.Ram[addr] & 0xF0) | MMU.ButtonState
				} else {
					MMU.Ram[addr] = (MMU.Ram[addr] & 0xF0) | MMU.ButtonState
					MMU.Ram[addr] = (MMU.Ram[addr] & 0xF0) | MMU.DpadState

				}

			} else if addr == DMA { // OAM DMA Transfer
				start := (val << 8)
				for oamAddr, srcAddr := 0xFE00, start; oamAddr < 0xFEA0; oamAddr, srcAddr = oamAddr+1, srcAddr+1 {
					MMU.Ram[oamAddr] = MMU.Ram[srcAddr]
				}

			} else if 0x0000 <= addr && addr <= 0x3FFF {
			} else {
				MMU.Ram[addr] = byte(val)
			}

		}
	case MBC1:
		{
			if addr >= 0x0000 && addr <= 0x1FFF {
				if val&0xF == 0xA {
					MMU.enableExtRam = true
				} else {
					MMU.enableExtRam = false
				}
				MMU.Ram[addr] = byte(val)

			} else if addr >= 0x2000 && addr <= 0x3FFF {
				if val == 0 {
					MMU.romBank = 1
				} else {
					MMU.romBank = val
				}
				MMU.Ram[addr] = byte(val)

			} else if addr >= 0x4000 && addr <= 0x5FFF {
				MMU.ramBank = val & 0x3
				MMU.Ram[addr] = byte(val)

			} else if addr >= 0x6000 && addr <= 0x7FFF {
				MMU.mode = (val & 1) == 1
				MMU.Ram[addr] = byte(val)

			} else if addr >= 0x8000 && addr <= 0x9FFF {
				MMU.Ram[addr] = byte(val)

			} else if addr >= 0xA000 && addr <= 0xBFFF {
				if MMU.enableExtRam {
					MMU.Ram[0xA000+addr] = byte(val)
				}

			} else if addr >= 0xC000 && addr <= 0xFFFF {
				MMU.Ram[addr] = byte(val)
			}
		}

	}
}
