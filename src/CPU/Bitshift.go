package CPU

func (s *CPU) RLR8(reg string) { // Rotate bits reg[] left through Flag C
	carry := s.GetReg8Val(reg) >> 7
	val := int(int8(s.GetReg8Val(reg)<<1) | int8(s.GetFlag("C")))
	s.SetReg8Val(reg, val)
	s.SetFlagRL(val, carry)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+2)
}

func (s *CPU) RLHL() { // Rotate bits M[reg[HL]] left through Flag C
	carry := s.GetHLVal() >> 7
	val := int(int8((s.GetHLVal() << 1) | s.GetFlag("C")))
	s.SetHLVal(val)
	s.SetFlagRL(val, carry)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+2)
}

func (s *CPU) RLA() { // Rotate bits reg[A] left through Flag C
	carry := s.GetReg8Val("A") >> 7
	val := (s.GetReg8Val("A") << 1) | s.GetFlag("C")
	s.SetReg8Val("A", val)
	s.SetFlagRL(1, carry)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) RLCR8(reg string) { // Rotate bits reg[A] left through Flag c
	carry := s.GetReg8Val(reg) >> 7
	val := (s.GetReg8Val(reg) << 1) | carry
	s.SetReg8Val(reg, val)
	s.SetFlagRL(val, carry)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+2)
}

func (s *CPU) RLCHL() { // Rotate bits M[reg[HL]] left through flag C
	carry := s.GetHLVal() >> 7
	val := (s.GetHLVal() << 1) | carry
	s.SetHLVal(val)
	s.SetFlagRL(val, carry)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+2)
}

func (s *CPU) RLCA() { // Rotate bits reg[A] left
	carry := s.GetReg8Val("A") >> 7
	val := (s.GetReg8Val("A") << 1) | carry
	s.SetReg8Val("A", val)
	s.SetFlagRL(1, carry)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) RRR8(reg string) { // Rotate bits reg[] right through flag C
	carry := s.GetReg8Val(reg) & 0x01
	val := (s.GetReg8Val(reg) >> 1) | (s.GetFlag("C") << 7)
	s.SetReg8Val(reg, val)
	s.SetFlagRL(val, carry)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+2)
}

func (s *CPU) RRHL() { // Rotate bits M[reg[HL]] right through flag C
	carry := s.GetHLVal() & 0x01
	val := (s.GetHLVal() >> 1) | (s.GetFlag("C") << 7)
	s.SetHLVal(val)
	s.SetFlagRL(val, carry)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+2)
}

func (s *CPU) RRA() { // Rotate bits reg[A] right through flag C
	carry := s.GetReg8Val("A") & 0x01
	val := (s.GetReg8Val("A") >> 1) | (s.GetFlag("C") << 7)
	s.SetReg8Val("A", val)
	s.SetFlagRL(1, carry)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) RRCR8(reg string) { // Rotate reg[] right
	carry := s.GetReg8Val(reg) & 0x01
	val := (s.GetReg8Val(reg) >> 1) | (carry << 7)
	s.SetReg8Val(reg, val)
	s.SetFlagRL(val, carry)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+2)
}

func (s *CPU) RRCHL() { // Rotate M[reg[HL]] right
	carry := s.GetHLVal() & 0x01
	val := (s.GetHLVal() >> 1) | (carry << 7)
	s.SetHLVal(val)
	s.SetFlagRL(val, carry)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+2)
}

func (s *CPU) RRCA() { // Rotate reg[A] right
	carry := s.GetReg8Val("A") & 0x01
	val := (s.GetReg8Val("A") >> 1) | (carry << 7)
	s.SetReg8Val("A", val)
	s.SetFlagRL(1, carry)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) SLAR8(reg string) { // reg[] = reg[] << 1
	carry := s.GetReg8Val(reg) >> 7
	val := int(int8(s.GetReg8Val(reg) << 1))
	s.SetReg8Val(reg, val)
	s.SetFlagRL(val, carry)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+2)
}

func (s *CPU) SLAHL() { // M[reg[HL]] = M[reg[HL]] << 1
	carry := s.GetHLVal() >> 7
	val := int(int8(s.GetHLVal() << 1))
	s.SetHLVal(val)
	s.SetFlagRL(val, carry)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+2)
}

func (s *CPU) SRAR8(reg string) { // reg[] = reg[] >> 1, bit 7 unchanged
	carry := s.GetReg8Val(reg) & 0x01
	val := (s.GetReg8Val(reg) >> 1) | s.GetReg8Val(reg)&0xA0
	s.SetReg8Val(reg, val)
	s.SetFlagRL(val, carry)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+2)
}

func (s *CPU) SRAHL() { // M[reg[HL]] = M[reg[HL]] >> 1, bit 7 unchanged
	carry := s.GetHLVal() & 0x01
	val := (s.GetHLVal() >> 1) | s.GetHLVal()&0xA0
	s.SetHLVal(val)
	s.SetFlagRL(val, carry)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+2)
}

func (s *CPU) SRLR8(reg string) { // reg[] = reg[] >> 1
	carry := s.GetReg8Val(reg) & 0x01
	val := s.GetReg8Val(reg) >> 1
	s.SetReg8Val(reg, val)
	s.SetFlagRL(val, carry)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+2)
}

func (s *CPU) SRLHL() { // M[reg[HL]] = M[reg[HL]] >> 1
	carry := s.GetHLVal() & 0x01
	val := s.GetHLVal() >> 1
	s.SetHLVal(val)
	s.SetFlagRL(val, carry)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+2)
}
