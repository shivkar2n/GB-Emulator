package CPU

func (s *CPU) ADDHLSP() { // reg[HL] = reg[HL] + reg[SP]
	s.ADDHLR16("SP")
}

func (s *CPU) ADDSPE8(op int) { // reg[SP] = reg[SP] + e8
	val := s.GetReg16Val("SP") + op
	s.SetReg16Val("SP", val)
	halfcarry := (int(s.SP[0]) & 0xF) + (op & 0xF)
	carry := (int(s.SP[0]) + op)
	s.SetFlagADD(halfcarry, carry, val)
}

func (s *CPU) DECSP() { // reg[SP] = reg[SP] - 1
	s.DECR16("SP")
}

func (s *CPU) INCSP() { // reg[SP] = reg[SP] + 1
	s.INCR16("SP")
}

func (s *CPU) LDSPN16(op int) { // reg[SP] = n8
	s.SetReg16Val("SP", op)
}

func (s *CPU) LDN16SP(op int) { // Store reg SP & 0xFF and SP >> 8 in M[n16+1]
	s.Mem[s.GetReg16Val("SP")&0xFF] = byte(op)
	s.Mem[op+1] = byte(s.GetReg16Val("SP") >> 8)
}

func (s *CPU) LDHLSPE8(op int) { // reg[HL] = reg[SP] + e8
	val := int(int8(s.GetReg16Val("SP")) + int8(op))
	halfcarry := int(s.SP[0])&0xF + op
	carry := int(s.SP[0]) + op
	s.SetReg16Val("HL", val)
	s.SetFlagADD(halfcarry, carry, 0)
}

func (s *CPU) LDSPHL() { // reg[SP] = reg[HL]
	s.SetReg16Val("SP", s.GetReg16Val("HL"))
}

func (s *CPU) POPAF() { // POP reg[AF] from Stack
	s.SetReg8Val("A", int(s.Mem[s.GetReg16Val("SP")]))
	s.INCSP()
	s.SetReg8Val("F", int(s.Mem[s.GetReg16Val("SP")]))
	s.INCSP()
}

func (s *CPU) POPR16(reg string) { // POP reg[] from Stack
	x := int(s.Mem[s.GetReg16Val("SP")])
	s.INCSP()
	y := 256 * int(s.Mem[s.GetReg16Val("SP")])
	s.INCSP()
	s.SetReg16Val(reg, x+y)
}

func (s *CPU) PUSHAF() { // PUSH reg[AF] from Stack
	s.DECSP()
	s.Mem[s.GetReg16Val("SP")] = byte(s.GetReg8Val("A"))
	s.DECSP()
	s.Mem[s.GetReg16Val("SP")] = byte((s.GetFlag("Z") << 7) | (s.GetFlag("N") << 6) | (s.GetFlag("H") << 5) | (s.GetFlag("C") << 4))
}

func (s *CPU) PUSHR16(reg string) { // PUSH reg[] from Stack
	s.DECSP()
	s.Mem[s.GetReg16Val("SP")] = byte((s.GetReg16Val(reg) & 0xFF00) >> 8)
	s.DECSP()
	s.Mem[s.GetReg16Val("SP")] = byte(s.GetReg16Val(reg) & 0xFF)
}
