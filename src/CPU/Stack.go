package CPU

func (s *CPU) ADDHLSP() { // reg[HL] = reg[HL] + reg[SP]
	s.ADDHLR16("SP")
}

func (s *CPU) ADDSPE8() { // reg[SP] = reg[SP] + e8
	op := int(int8(int(s.Mem[s.GetReg16Val("PC")+1])))
	val := int(int16(s.GetReg16Val("SP") + op))
	halfcarry := (int(s.SP[0]) & 0xF) + (op & 0xF)
	carry := int(s.SP[0])&0xFF + op&0xFF
	s.SetFlagADD(halfcarry, carry, 1)
	s.SetReg16Val("SP", val)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+2)
}

func (s *CPU) DECSP() { // reg[SP] = reg[SP] - 1
	s.DECR16("SP")
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) INCSP() { // reg[SP] = reg[SP] + 1
	s.INCR16("SP")
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) LDSPN16(op int) { // reg[SP] = n8
	s.SetReg16Val("SP", op)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+3)
}

func (s *CPU) LDN16SP() {
	addr := int(s.Mem[s.GetReg16Val("PC")+1]) + (int(s.Mem[s.GetReg16Val("PC")+2]) << 8)
	s.Mem[uint16(addr)] = byte(s.GetReg16Val("SP") & 0xFF) // M[n16] = reg[SP]
	s.Mem[uint16(addr+1)] = byte(s.GetReg16Val("SP") >> 8) // M[n16+1] = reg[SP] >> 8
	s.SetReg16Val("PC", s.GetReg16Val("PC")+3)
}

func (s *CPU) LDHLSPE8() { // reg[HL] = reg[SP] + e8
	op := int(int8(s.Mem[s.GetReg16Val("PC")+1]))
	val := int(int16(s.GetReg16Val("SP") + op))
	halfcarry := int(s.SP[0])&0xF + op&0xF
	carry := int(s.SP[0])&0xFF + op&0xFF
	s.SetReg16Val("HL", val)
	s.SetFlagADD(halfcarry, carry, 1)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+2)
}

func (s *CPU) LDSPHL() { // reg[SP] = reg[HL]
	s.SetReg16Val("SP", s.GetReg16Val("HL"))
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) POPFA() { // POP reg[AF] from Stack
	s.SetReg8Val("F", int(s.Mem[s.GetReg16Val("SP")])&0xF0)
	s.SetReg16Val("SP", int(int16(s.GetReg16Val("SP")+1)))
	s.SetReg8Val("A", int(s.Mem[s.GetReg16Val("SP")]))
	s.SetReg16Val("SP", int(int16(s.GetReg16Val("SP")+1)))
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) POPR16(reg string) { // POP reg[] from Stack
	x := int(s.Mem[s.GetReg16Val("SP")])
	s.SetReg16Val("SP", s.GetReg16Val("SP")+1)
	y := int(s.Mem[s.GetReg16Val("SP")]) << 8
	s.SetReg16Val("SP", s.GetReg16Val("SP")+1)
	s.SetReg16Val(reg, int(int16(x+y)))
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) PUSHFA() { // PUSH reg[AF] from Stack
	s.SetReg16Val("SP", int(int16(s.GetReg16Val("SP")-1)))
	s.Mem[s.GetReg16Val("SP")] = byte(s.GetReg8Val("A"))
	s.SetReg16Val("SP", int(int16(s.GetReg16Val("SP")-1)))
	s.Mem[s.GetReg16Val("SP")] = byte((s.GetFlag("Z") << 7) | (s.GetFlag("N") << 6) | (s.GetFlag("H") << 5) | (s.GetFlag("C") << 4))
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) PUSHR16(reg string) { // PUSH reg[] from Stack
	s.SetReg16Val("SP", int(int16(s.GetReg16Val("SP")-1)))
	s.Mem[s.GetReg16Val("SP")] = byte((s.GetReg16Val(reg) & 0xFF00) >> 8)
	s.SetReg16Val("SP", int(int16(s.GetReg16Val("SP")-1)))
	s.Mem[s.GetReg16Val("SP")] = byte(s.GetReg16Val(reg) & 0xFF)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) PUSHN16(op int) { // PUSH N16 from Stack
	s.SetReg16Val("SP", int(int16(s.GetReg16Val("SP")-1)))
	s.Mem[s.GetReg16Val("SP")] = byte((op & 0xFF00) >> 8)
	s.SetReg16Val("SP", int(int16(s.GetReg16Val("SP")-1)))
	s.Mem[s.GetReg16Val("SP")] = byte(op & 0xFF)
}
