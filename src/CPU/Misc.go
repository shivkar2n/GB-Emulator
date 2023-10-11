package CPU

func (s *CPU) CPL() { // Reg[A] = ~Reg[A]
	s.SetReg8Val("A", ^s.GetReg8Val("A"))
	s.SetFlag("N")
	s.SetFlag("H")
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) SCF() { // Flag[C] = true
	s.ResetFlag("N")
	s.ResetFlag("H")
	s.SetFlag("C")
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) CCF() { // Flag[C] = ~Flag[C]
	if s.GetFlag("C") == 1 {
		s.ResetFlag("C")
	} else {
		s.SetFlag("C")
	}
	s.ResetFlag("N")
	s.ResetFlag("H")
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) NOP() {
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
	s.SetClockTime(4, 1)
}

func (s *CPU) HALT() {
	s.StopExec = true
	s.SetClockTime(4, 1)
}
