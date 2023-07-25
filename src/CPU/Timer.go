package CPU

func (s *CPU) INCRDIV() { // Increment DIV register
	s.Mem[0xFF04] = byte(int8(s.Mem[0xFF04] + 1))
}

func (s *CPU) INCRTIMA() { // Rate of INC according to TAC Register
	if int(s.Mem[0xFF05])+1 == 0x100 { // If overflow set to TMA value
		s.Mem[0xFF05] = s.Mem[0xFF06]
		s.SetIFBit(2)
	} else {
		s.Mem[0xFF05] = byte(int8(s.Mem[0xFF05] + 1))
	}
}

func (s *CPU) INCRTMA() {
	if int(s.Mem[0xFF06])+1 == 0x100 {

	} else {
		s.Mem[0xFF06] = byte(int8(s.Mem[0xFF06] + 1))
	}
}

func (s *CPU) GetDIVBit(op int) int {
	return (int(s.Mem[0xFF04]) >> op) & 0x01
}

func (s *CPU) GetTACPos() int {
	return int(s.Mem[0xFF07]) & 0x03
}

func (s *CPU) GetTEVal() int { // Get timer-enable bit
	return (int(s.Mem[0xFF07]) & 0x04) >> 2 & 0x01
}

func (s *CPU) INCRTIMER() {
	s.INCRDIV()
	curCycle := s.GetDIVBit(s.GetTACPos()) & s.GetTEVal()
	if s.PrevCycle == 1 && curCycle == 0 {
		s.INCRTIMA()
	}
	s.PrevCycle = curCycle
}
