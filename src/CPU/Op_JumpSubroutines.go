package CPU

func (s *CPU) JPN16(op int) { // reg[PC] = n16
	s.SetReg16Val("PC", op)
}

func (s *CPU) JPCCN16(cc string, op int) { // reg[PC] = n16, if cc == true
	if s.checkCC(cc) {
		s.JPN16(op)
	}
}

func (s *CPU) JPNHL() { // reg[PC] = reg[HL]
	s.SetReg16Val("PC", s.GetHLVal())
}

func (s *CPU) JRN16(op int) { // reg[PC] = reg[PC] + n16
	s.JPN16(s.GetPCVal() + op)
}

func (s *CPU) JRCCN16(cc string, op int) { // reg[PC] = reg[PC] + n16, if cc == true
	s.JPCCN16(cc, s.GetPCVal()+op)
}

func (s *CPU) RET() { //
	s.POPR16("PC")
}

func (s *CPU) RETCC(cc string) { //
	if s.checkCC(cc) {
		s.POPR16("PC")
	}
}

// RETI
// RST VEC
