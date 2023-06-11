// Opcodes for LOAD instructions

package CPU

func (s *CPU) LDR8R8(reg1 string, reg2 string) { // reg[reg1] = reg[reg2]
	s.SetReg8Val(reg1, s.GetReg8Val(reg2))
}

func (s *CPU) LDR8N8(reg string, op int) { // reg[] = n8
	s.SetReg8Val(reg, op)
}

func (s *CPU) LDR16R16(reg1 string, reg2 string) { // reg16[reg1] = reg16[reg2]
	s.SetReg16Val(reg1, s.GetReg16Val(reg2))
}

func (s *CPU) LDHLR8(reg string) { // M[HL] = reg[]
	s.SetHLVal(s.GetReg16Val(reg))
}

func (s *CPU) LDHLN8(op int) { // M[HL] = n8
	s.SetHLVal(op)
}

func (s *CPU) LDR8HL(reg string) { // M[HL] = reg[]
	s.SetReg8Val(reg, s.GetHLVal())
}

func (s *CPU) LDR16A(reg string) { // M[reg] = reg[A]
	s.Mem[s.GetReg16Val(reg)] = byte(s.GetReg8Val("A"))
}

func (s *CPU) LDN16A(op int) { // M[n16] = reg[A]
	s.Mem[op] = byte(s.GetReg8Val("A"))
}

func (s *CPU) LDHN16A(op int) { // M[n16] = reg[A] provided n16 b/w 0xFF00 and 0xFFFF
	if 0xFFFF <= op && op <= 0xFFFF {
		s.Mem[op] = byte(s.GetReg8Val("A"))
	}
}

func (s *CPU) LDHCA(C int) { // M[0xFF00+C] = reg[A]
	s.Mem[0xFF00+C] = byte(s.GetReg8Val("A"))
}

func (s *CPU) LDAR16(reg string) { // reg[A] = M[reg[]]
	s.SetReg8Val("A", int(s.Mem[s.GetReg16Val(reg)]))
}

func (s *CPU) LDAN16(op int) { // regA = M[n16]
	s.SetReg8Val("A", int(s.Mem[op]))
}

func (s *CPU) LDHAC() { // reg[A] = M[0xFF00+reg[C]]
	s.SetReg8Val("A", int(s.Mem[0xFF00+s.GetReg8Val("C")]))
}

func (s *CPU) LDHLIA() { // M[HL] = reg[A], HL++
	s.Mem[s.GetHLVal()] = byte(s.GetReg8Val("A"))
	s.SetReg16Val("HL", s.GetReg16Val("HL")+1)
}

func (s *CPU) LDHLDA() { // M[HL] = reg[A], HL--
	s.Mem[s.GetHLVal()] = byte(s.GetReg8Val("A"))
	s.SetReg16Val("HL", s.GetReg16Val("HL")-1)
}

func (s *CPU) LDAHLI() { // reg[A] = M[HL], HL++
	s.SetReg8Val("A", s.GetHLVal())
	s.SetReg16Val("HL", s.GetReg16Val("HL")+1)
}

func (s *CPU) LDAHLD() { // reg[A] = M[HL], HL--
	s.SetReg8Val("A", s.GetHLVal())
	s.SetReg16Val("HL", s.GetReg16Val("HL")-1)
}
