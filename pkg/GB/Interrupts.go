package GB

func (GB *GB) GetIEBit(pos int) int { // Get interrupt enable
	return (GB.MMU.Read(0xFFFF) >> pos) & 1
}

func (GB *GB) GetIFBit(pos int) int { // Get interrupt flag
	return (GB.MMU.Read(0xFF0F) >> pos) & 1
}

func (GB *GB) SetIEBit(op int) { // Set interrupt enable bit
	GB.MMU.Write(GB.MMU.Read(0xFFFF)|1<<op, 0xFFFF)
}

func (GB *GB) ResetIEBit(offset int) { // Reset interrupt enable bit
	k := 0
	for i := 0; i < 8; i++ {
		k = k << 1
		if 7-i != offset {
			k += 1
		}
	}
	GB.MMU.Write(GB.MMU.Read(0xFFFF)&k, 0xFFFF)
}

func (GB *GB) SetIFBit(op int) { // Set interrupt flag bit
	GB.MMU.Write(GB.MMU.Read(0xFF0F)|1<<op, 0xFF0F)
}

func (GB *GB) ResetIFBit(offset int) { // Reset interrupt flag bit
	k := 0
	for i := 0; i < 8; i++ {
		k = k << 1
		if 7-i != offset {
			k += 1
		}
	}
	GB.MMU.Write(GB.MMU.Read(0xFF0F)&k, 0xFF0F)
}

func (GB *GB) InterruptHandler(exec bool) (int, bool) {
	if GB.CPU.IME {
		if GB.GetIEBit(0) == 1 && GB.GetIFBit(0) == 1 { // VBlank
			GB.CPU.IME = false
			// fmt.Printf("V-Blank Interrupt!\t")
			GB.CALLN8(0x40)
			GB.ResetIFBit(0)
			return 20, true
		}
		if GB.GetIEBit(1) == 1 && GB.GetIFBit(1) == 1 { // LCD STAT
			GB.CPU.IME = false
			// fmt.Printf("LCD Interrupt!\n")
			GB.CALLN8(0x48)
			GB.ResetIFBit(1)
			return 20, true
		}
		if GB.GetIEBit(2) == 1 && GB.GetIFBit(2) == 1 { // Timer
			GB.CPU.IME = false
			// fmt.Printf("Timer Interrupt!\n")
			GB.CALLN8(0x50)
			GB.ResetIFBit(2)
			return 20, true
		}
		if GB.GetIEBit(3) == 1 && GB.GetIFBit(3) == 1 { // Serial
			GB.CPU.IME = false
			// fmt.Printf("Serial Interrupt!\t")
			GB.CALLN8(0x58)
			GB.ResetIFBit(3)
			return 20, true

		}
		if GB.GetIEBit(4) == 1 && GB.GetIFBit(4) == 1 { // Joypad
			GB.CPU.IME = false
			// fmt.Printf("Joypad Interrupt!\t")
			GB.CALLN8(0x60)
			GB.ResetIFBit(4)
			return 20, true
		}
	} else if !GB.CPU.IME && !exec && (GB.MMU.Read(0xFF0F)&GB.MMU.Read(0xFFFF)) != 0 {
		return 0, true
	}
	return 0, exec
}
