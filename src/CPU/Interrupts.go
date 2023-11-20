package CPU

// import "log"

func (s *CPU) EI() {
	s.IME = true
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) DI() {
	s.IME = false
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) GetIEBit(pos int) int { // Get interrupt enable
	return (int(s.Mem.Read(0xFFFF)) >> pos) & 1
}

func (s *CPU) GetIFBit(pos int) int { // Get interrupt flag
	return (int(s.Mem.Read(0xFF0F)) >> pos) & 1
}

func (s *CPU) SetIEBit(op int) { // Set interrupt enable bit
	s.Mem.Write(int(s.Mem.Read(0xFFFF))|1<<op, 0xFFFF)
}

func (s *CPU) ResetIEBit(offset int) { // Reset interrupt enable bit
	k := 0
	for i := 0; i < 8; i++ {
		k = k << 1
		if 7-i != offset {
			k += 1
		}
	}
	s.Mem.Write(int(s.Mem.Read(0xFFFF))&k, 0xFFFF)
}

func (s *CPU) SetIFBit(op int) { // Set interrupt flag bit
	s.Mem.Write(int(s.Mem.Read(0xFF0F))|1<<op, 0xFF0F)
}

func (s *CPU) ResetIFBit(offset int) { // Reset interrupt flag bit
	k := 0
	for i := 0; i < 8; i++ {
		k = k << 1
		if 7-i != offset {
			k += 1
		}
	}
	s.Mem.Write(int(s.Mem.Read(0xFF0F))&k, 0xFF0F)
}

func (s *CPU) InterruptHandler() {
	// log.Printf("IME:%t IF:%X IE:%X\n", s.IME, s.Mem.Read(0xFF0F), s.Mem.Read(0xFFFF))
	if s.IME {
		if s.GetIEBit(0) == 1 && s.GetIFBit(0) == 1 { // VBlank
			s.IME = false
			// log.Printf("V-Blank Interrupt!\t")
			s.CALLN8(0x40)
			s.ResetIFBit(0)
			s.SetClockTime(4, 1)

		} else if s.GetIEBit(1) == 1 && s.GetIFBit(1) == 1 { // LCD STAT
			s.IME = false
			// log.Printf("LCD Interrupt!\t")
			s.CALLN8(0x48)
			s.ResetIFBit(1)
			s.SetClockTime(4, 1)

		} else if s.GetIEBit(2) == 1 && s.GetIFBit(2) == 1 { // Timer
			s.IME = false
			// log.Printf("Timer Interrupt!\t")
			s.CALLN8(0x50)
			s.ResetIFBit(2)
			s.SetClockTime(4, 1)

		} else if s.GetIEBit(3) == 1 && s.GetIFBit(3) == 1 { // Serial
			s.IME = false
			// log.Printf("Serial Interrupt!\t")
			s.CALLN8(0x58)
			s.ResetIFBit(3)
			s.SetClockTime(4, 1)

		} else if s.GetIEBit(4) == 1 && s.GetIFBit(4) == 1 { // Joypad
			s.IME = false
			// log.Printf("Joypad Interrupt!\t")
			s.CALLN8(0x60)
			s.ResetIFBit(4)
			s.SetClockTime(4, 1)
		}
	}
}
