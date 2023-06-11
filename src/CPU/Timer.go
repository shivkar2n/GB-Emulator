package CPU

func (s *CPU) INCRDIV() {
	s.Mem[0xFF04] = byte(int8(s.Mem[0xFF04] + 1))
}

func (s *CPU) INCRTIMA() { // Rate of INC according to TAC Register
	if int(s.Mem[0xFF05])+1 == 0x100 {
		s.Mem[0xFF05] = s.Mem[0xFF06]
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
