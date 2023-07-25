// Opcodes for LOAD instructions
// n8 - 8 bit int
// n16 - 16 bit const
// e8 - 8 bit offset
// u3 - 3 bit unsigned

package CPU

func (s *CPU) LDR8R8(reg1 string, reg2 string) { // reg[reg1] = reg[reg2]
	s.SetReg8Val(reg1, s.GetReg8Val(reg2))
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) LDR8N8(reg string) { // reg[] = n8
	addr := int(s.Mem[s.GetReg16Val("PC")+1])
	s.SetReg8Val(reg, addr)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+2)
}

func (s *CPU) LDR16N16(reg string) { // reg16[reg] = n16
	addr := int(uint16(int(s.Mem[s.GetReg16Val("PC")+1]) + (int(s.Mem[s.GetReg16Val("PC")+2]) << 8)))
	s.SetReg16Val(reg, addr)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+3)
}

func (s *CPU) LDHLR8(reg string) { // M[HL] = reg[]
	s.SetHLVal(s.GetReg8Val(reg))
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) LDHLN8() { // M[HL] = n8
	op := int(s.Mem[s.GetReg16Val("PC")+1])
	s.SetHLVal(op)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+2)
}

func (s *CPU) LDR8HL(reg string) { // reg[] = M[HL]
	s.SetReg8Val(reg, s.GetHLVal())
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) LDR16A(reg string) { // M[reg] = reg[A]
	s.Mem[s.GetReg16Val(reg)] = byte(s.GetReg8Val("A"))
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) LDN16A(op int) { // M[n16] = reg[A]
	s.Mem[op] = byte(s.GetReg8Val("A"))
	s.SetReg16Val("PC", s.GetReg16Val("PC")+3)
}

func (s *CPU) LDAR16(reg string) { // reg[A] = M[reg[]]
	s.SetReg8Val("A", int(s.Mem[s.GetReg16Val(reg)]))
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) LDAN16(op int) { // regA = M[n16]
	s.SetReg8Val("A", int(s.Mem[op]))
	s.SetReg16Val("PC", s.GetReg16Val("PC")+2)
}

func (s *CPU) LDHLIA() { // M[HL] = reg[A], HL++
	s.SetHLVal(s.GetReg8Val("A"))
	s.SetReg16Val("HL", int(uint16(s.GetReg16Val("HL")+1)))
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) LDHLDA() { // M[HL] = reg[A], HL--
	s.SetHLVal(s.GetReg8Val("A"))
	s.SetReg16Val("HL", int(uint16(s.GetReg16Val("HL")-1)))
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) LDAHLI() { // reg[A] = M[HL], HL++
	s.SetReg8Val("A", s.GetHLVal())
	s.SetReg16Val("HL", int(uint16(s.GetReg16Val("HL")+1)))
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) LDAHLD() { // reg[A] = M[HL], HL--
	s.SetReg8Val("A", s.GetHLVal())
	s.SetReg16Val("HL", int(uint16(s.GetReg16Val("HL")-1)))
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}
