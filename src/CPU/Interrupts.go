package CPU

func (s *CPU) EI() {
	s.SetIMEval()
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) DI() {
	s.ResetIMEval()
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) SetIMEval() {
	s.IME = true
}

func (s *CPU) ResetIMEval() {
	s.IME = false
}

func (s *CPU) GetIMEVal() bool { // Get interrupt master enable flag
	return s.IME
}

func (s *CPU) GetIEVal() int { // Get interrupt enable
	return int(s.Mem[0xFFFF])
}

func (s *CPU) SetIEVal(op int) { // Set interrupt enable
	s.Mem[0xFFFF] = byte(op)
}

func (s *CPU) GetIFVal() int { // Get interrupt flag
	return int(s.Mem[0xFF0F])
}

func (s *CPU) SetIEBit(op int) { // Set interrupt enable bit
	s.Mem[0xFFFF] = byte(s.Mem[0xFFFF] | 1<<op)
}

func (s *CPU) ResetIEBit(offset int) { // Reset interrupt enable bit
	k := 0
	for i := 0; i < 8; i++ {
		k = k << 1
		if 7-i != offset {
			k += 1
		}
	}
	s.Mem[0xFFFF] = byte(int(s.Mem[0xFFFF]) & k)
}

func (s *CPU) SetIFBit(op int) { // Set interrupt flag bit
	s.Mem[0xFF0F] = byte(s.Mem[0xFF0F] | 1<<op)
}

func (s *CPU) ResetIFBit(offset int) { // Reset interrupt flag bit
	k := 0
	for i := 0; i < 8; i++ {
		k = k << 1
		if 7-i != offset {
			k += 1
		}
	}
	s.Mem[0xFF0F] = byte(int(s.Mem[0xFF0F]) & k)
}

func (s *CPU) InterruptHandler() { // Interrupt handler
	if s.GetIMEVal() {
		if s.GetIEVal()>>0&1 == 1 { // VBlank
			s.CALLN8(0x40)
			s.ResetIFBit(0)

		} else if s.GetIEVal()>>1&1 == 1 { // LCD STAT
			s.CALLN8(0x48)
			s.ResetIFBit(1)

		} else if s.GetIEVal()>>2&1 == 1 { // Timer
			s.CALLN8(0x50)
			s.ResetIFBit(2)

		} else if s.GetIEVal()>>3&1 == 1  { // Serial
			s.CALLN8(0x58)
			s.ResetIFBit(3)

		} else if s.GetIEVal()>>4&1 == 1 { // Joypad
			s.CALLN8(0x60)
			s.ResetIFBit(4)

		}
		s.RETI()
		s.ResetIMEval()
	}
}
