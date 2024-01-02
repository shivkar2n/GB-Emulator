package CPU

func (cpu *CPU) GetIEBit(pos int) int { // Get interrupt enable
	return (cpu.Mem.Read(0xFFFF) >> pos) & 1
}

func (cpu *CPU) GetIFBit(pos int) int { // Get interrupt flag
	return (cpu.Mem.Read(0xFF0F) >> pos) & 1
}

func (cpu *CPU) SetIEBit(op int) { // Set interrupt enable bit
	cpu.Mem.Write(cpu.Mem.Read(0xFFFF)|1<<op, 0xFFFF)
}

func (cpu *CPU) ResetIEBit(offset int) { // Reset interrupt enable bit
	k := 0
	for i := 0; i < 8; i++ {
		k = k << 1
		if 7-i != offset {
			k += 1
		}
	}
	cpu.Mem.Write(cpu.Mem.Read(0xFFFF)&k, 0xFFFF)
}

func (cpu *CPU) SetIFBit(op int) { // Set interrupt flag bit
	cpu.Mem.Write(cpu.Mem.Read(0xFF0F)|1<<op, 0xFF0F)
}

func (cpu *CPU) ResetIFBit(offset int) { // Reset interrupt flag bit
	k := 0
	for i := 0; i < 8; i++ {
		k = k << 1
		if 7-i != offset {
			k += 1
		}
	}
	cpu.Mem.Write(cpu.Mem.Read(0xFF0F)&k, 0xFF0F)
}

func (cpu *CPU) InterruptHandler(exec bool) (int, bool) {
	if cpu.IME {
		if cpu.GetIEBit(0) == 1 && cpu.GetIFBit(0) == 1 { // VBlank
			cpu.IME = false
			// fmt.Printf("V-Blank Interrupt!\t")
			cpu.CALLN8(0x40)
			cpu.ResetIFBit(0)
			return 20, true
		}
		if cpu.GetIEBit(1) == 1 && cpu.GetIFBit(1) == 1 { // LCD STAT
			cpu.IME = false
			// fmt.Printf("LCD Interrupt!\n")
			cpu.CALLN8(0x48)
			cpu.ResetIFBit(1)
			return 20, true
		}
		if cpu.GetIEBit(2) == 1 && cpu.GetIFBit(2) == 1 { // Timer
			cpu.IME = false
			// fmt.Printf("Timer Interrupt!\n")
			cpu.CALLN8(0x50)
			cpu.ResetIFBit(2)
			return 20, true
		}
		if cpu.GetIEBit(3) == 1 && cpu.GetIFBit(3) == 1 { // Serial
			cpu.IME = false
			// fmt.Printf("Serial Interrupt!\t")
			cpu.CALLN8(0x58)
			cpu.ResetIFBit(3)
			return 20, true

		}
		if cpu.GetIEBit(4) == 1 && cpu.GetIFBit(4) == 1 { // Joypad
			cpu.IME = false
			// fmt.Printf("Joypad Interrupt!\t")
			cpu.CALLN8(0x60)
			cpu.ResetIFBit(4)
			return 20, true
		}
	} else if !cpu.IME && !exec && (cpu.Mem.Read(0xFF0F)&cpu.Mem.Read(0xFFFF)) != 0 {
		return 0, true
	}
	return 0, exec
}
