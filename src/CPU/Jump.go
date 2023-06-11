package CPU

func (s *CPU) JPN16(op int) { // reg[PC] = n16
	s.SetReg16Val("PC", op)
}

func (s *CPU) JPCCN16(cc string, op int) { // reg[PC] = n16, if cc == true
	if s.checkCC(cc) {
		s.JPN16(op)
	} else {
		s.SetReg16Val("PC", s.GetReg16Val("PC")+3)
	}
}

func (s *CPU) JPHL() { // reg[PC] = reg[HL]
	s.SetReg16Val("PC", s.GetReg16Val("HL"))
}

func (s *CPU) JRN16() { // reg[PC] = reg[PC] + n16
	addr := int(int8(s.Mem[s.GetReg16Val("PC")+1]))
	s.JPN16(s.GetReg16Val("PC") + addr + 2)
}

func (s *CPU) JRCCN16(cc string) { // reg[PC] = reg[PC] + n16, if cc == true
	if s.checkCC(cc) {
		s.JRN16()
	} else {
		s.SetReg16Val("PC", s.GetReg16Val("PC")+2)
	}
}

func (s *CPU) RET() { //
	s.POPR16("PC")
	s.SetReg16Val("PC", s.GetReg16Val("PC")-1)
}

func (s *CPU) RETCC(cc string) { //
	if s.checkCC(cc) {
		s.RET()
	} else {
		s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
	}
}

func (s *CPU) CALLN16(op int) { // Equivalent to following
	s.PUSHN16(s.GetReg16Val("PC") + 3) // PUSH M[reg[PC+3]]
	s.JPN16(op)                        // JP N16
}

func (s *CPU) CALLN8(op int) { // Equivalent to following
	s.PUSHN16(s.GetReg16Val("PC")) // PUSH M[reg[PC]]
	s.JPN16(op)                    // JP N16
}

func (s *CPU) CALLCCN16(cc string, op int) { // CALLN16 if cc == true
	if s.checkCC(cc) {
		s.CALLN16(op)
	} else {
		s.SetReg16Val("PC", s.GetReg16Val("PC")+3)
	}
}

func (s *CPU) RETI() { // Equivalent to:
	s.IME = true // EI
	s.RET()      // RET
}

func (s *CPU) RST(vec string) {
	var vecAddr int
	if vec == "00H" {
		vecAddr = 0x0000
	} else if vec == "08H" {
		vecAddr = 0x0008
	} else if vec == "10H" {
		vecAddr = 0x0010
	} else if vec == "18H" {
		vecAddr = 0x0018
	} else if vec == "20H" {
		vecAddr = 0x0020
	} else if vec == "28H" {
		vecAddr = 0x0028
	} else if vec == "30H" {
		vecAddr = 0x0030
	} else if vec == "38H" {
		vecAddr = 0x0038
	}
	s.PUSHN16(s.GetReg16Val("PC") + 1)
	s.JPN16(vecAddr)
}
