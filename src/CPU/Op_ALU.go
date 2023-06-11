// Opcodes for 8-bit ALU instructions

package CPU

func (s *CPU) ADCAR8(reg string) { // reg[A] = reg[A] + C + reg[]
	regVal := s.GetReg8Val(reg)
	val := s.GetReg8Val("A") + regVal + int(s.GetFlag("C"))
	halfcarry := (s.GetReg8Val("A") & 0xF) + regVal + (int(s.GetFlag("C")) & 0xF)
	carry := (s.GetReg8Val("A") & 0xFF) + regVal + (int(s.GetFlag("C")) & 0xFF)
	s.SetReg8Val("A", val)
	s.SetFlagADD(halfcarry, carry, val)
}

func (s *CPU) ADCAHL() { // reg[A] = reg[A] + C + M[HL]
	val := int(s.AF[1]) + s.GetHLVal() + int(s.GetFlag("C"))
	halfcarry := (int(s.AF[1]) & 0xF) + s.GetHLVal() + (int(s.GetFlag("C")) & 0xF)
	carry := (int(s.AF[1]) & 0xFF) + s.GetHLVal() + (int(s.GetFlag("C")) & 0xFF)
	s.SetReg8Val("A", val)
	s.SetFlagADD(halfcarry, carry, val)
}

func (s *CPU) ADCAN8(op int) { // reg[A] = reg[A] + C + n8
	val := s.GetReg8Val("A") + op + int(s.GetFlag("C"))
	halfcarry := (s.GetReg8Val("A") & 0xF) + op + (int(s.GetFlag("C")) & 0xF)
	carry := (s.GetReg8Val("A") & 0xFF) + op + (int(s.GetFlag("C")) & 0xFF)
	s.SetReg8Val("A", val)
	s.SetFlagADD(halfcarry, carry, val)
}

func (s *CPU) ADDAR8(reg string) { // reg[A] = reg[A] + reg[]
	regVal := s.GetReg8Val(reg)
	val := s.GetReg8Val("A") + regVal
	halfcarry := (s.GetReg8Val("A") & 0xF) + regVal
	carry := (s.GetReg8Val("A") & 0xFF) + regVal
	s.SetReg8Val("A", val)
	s.SetFlagADD(halfcarry, carry, val)
}

func (s *CPU) ADDAHL() { // reg[A] = reg[A] + M[HL]
	val := int(s.AF[1]) + s.GetHLVal()
	halfcarry := (int(s.AF[1]) & 0xF) + val
	carry := (s.GetReg8Val("A") & 0xFF) + val
	s.SetReg8Val("A", val)
	s.SetFlagADD(halfcarry, carry, val)
}
func (s *CPU) ADDAN8(op int) { // reg[A] = reg[A] + n8
	val := s.GetReg8Val("A") + op
	halfcarry := (s.GetReg8Val("A") & 0xF) + op
	carry := (s.GetReg8Val("A") & 0xFF)
	s.SetReg8Val("A", val)
	s.SetFlagADD(halfcarry, carry, val)
}

func (s *CPU) ANDAR8(reg string) { // reg[A] = reg[A] & reg[]
	regVal := s.GetReg8Val(reg)
	val := s.GetReg8Val("A") & regVal
	s.SetReg8Val("A", val)
	s.SetFlagAND(val)
}

func (s *CPU) ANDAHL() { // reg[A] = reg[A] & M[HL]
	val := int(s.AF[1]) & s.GetHLVal()
	s.SetReg8Val("A", val)
	s.SetFlagAND(val)
}

func (s *CPU) ANDAN8(op int) { // reg[A] = reg[A] & n8
	val := s.GetReg8Val("A") & op
	s.SetReg8Val("A", val)
	s.SetFlagAND(val)
}

func (s *CPU) CPAR8(reg string) { // res = reg[A] - reg[]
	regVal := s.GetReg8Val(reg)
	halfBorrow := ((int(s.AF[1]) & 0x0F) - (regVal & 0x0F)) < 0
	borrow := int(s.AF[1])-regVal < 0
	zero := int(s.AF[1])-regVal == 0
	s.SetFlag("N")
	if halfBorrow {
		s.SetFlag("H")
	}
	if borrow {
		s.SetFlag("C")
	}
	if zero {
		s.SetFlag("Z")
	}
}

func (s *CPU) CPAHL() { // res = reg[A] - M[HL]
	val := s.GetReg8Val("A") - s.GetHLVal()
	halfBorrow := ((s.GetReg8Val("A") & 0x0F) - (s.GetHLVal() & 0x0F)) < 0
	borrow := val < 0
	zero := val == 0
	s.SetFlag("N")
	if halfBorrow {
		s.SetFlag("H")
	}
	if borrow {
		s.SetFlag("C")
	}
	if zero {
		s.SetFlag("Z")
	}
}

func (s *CPU) CPAN8(op int) { // res = reg[A] - n8
	val := s.GetReg8Val("A") - op
	halfBorrow := ((s.GetReg8Val("A") & 0x0F) - (op & 0x0F)) < 0
	borrow := val < 0
	zero := val == 0
	s.SetFlag("N")
	if halfBorrow {
		s.SetFlag("H")
	}
	if borrow {
		s.SetFlag("C")
	}
	if zero {
		s.SetFlag("Z")
	}
}

func (s *CPU) DECR8(reg string) { // reg[] = reg[] - 1
	val := s.GetReg8Val(reg) - 0x01
	halfBorrow := ((s.GetReg8Val("A") & 0x0F) - 0x01) < 0
	zero := val == 0
	s.SetReg8Val(reg, val)
	s.SetFlag("N")
	if halfBorrow {
		s.SetFlag("H")
	}
	if zero {
		s.SetFlag("Z")
	}
}

func (s *CPU) DECHL() { // M[HL] = M[HL] - 1
	val := s.GetHLVal() - 0x01
	halfBorrow := ((s.GetHLVal() & 0x0F) - 0x01) < 0
	s.SetHLVal(val)
	zero := val == 0
	s.SetFlag("N")
	if halfBorrow {
		s.SetFlag("H")
	}
	if zero {
		s.SetFlag("Z")
	}
}

func (s *CPU) INCR8(reg string) { // reg[] = reg[] + 1
	val := s.GetReg8Val(reg) + 0x01
	halfcarry := ((s.GetReg8Val("A") & 0x0F) + 0x01)
	zero := val == 0
	s.SetReg8Val(reg, val)
	s.ResetFlag("N")
	if halfcarry&0x10 == 0x10 {
		s.SetFlag("H")
	}
	if zero {
		s.SetFlag("Z")
	}
}

func (s *CPU) INCRHL() { // M[HL] = M[HL] + 1
	val := s.GetHLVal() + 0x01
	halfcarry := (s.GetHLVal() & 0x0F) + 0x01
	zero := val == 0
	s.SetHLVal(val)
	s.ResetFlag("N")
	if halfcarry&0x10 == 0x10 {
		s.SetFlag("H")
	}
	if zero {
		s.SetFlag("Z")
	}
}

func (s *CPU) ORAR8(reg string) { // reg[A] = reg[A] | reg[]
	regVal := s.GetReg8Val(reg)
	val := s.GetReg8Val("A") | regVal
	s.SetReg8Val("A", val)
	s.SetFlagOR(val)
}

func (s *CPU) ORAHL() { // M[HL] = M[HL] | reg[A]
	val := s.GetReg8Val("A") | s.GetHLVal()
	s.SetReg8Val("A", val)
	s.SetFlagOR(val)
}

func (s *CPU) ORAN8(op int) { // reg[A] = reg[A] | n8
	val := s.GetReg8Val("A") | op
	s.SetReg8Val("A", val)
	s.SetFlagOR(val)
}

func (s *CPU) SBCAR8(reg string) { // reg[A] = reg[A] - reg[] - C
	regVal := s.GetReg8Val(reg)
	val := s.GetReg8Val("A") - regVal - int(s.GetFlag("C"))
	halfborrow := (s.GetReg8Val("A")&0xF)-(regVal&0xF)-(int(s.GetFlag("C"))&0xF) < 0
	borrow := s.GetReg8Val("A")-regVal-int(s.GetFlag("C")) < 0
	s.SetReg8Val("A", val)
	s.SetFlagSUB(halfborrow, borrow, val)
}

func (s *CPU) SBCAHL() { // reg[A] = reg[A] - M[HL] - C
	val := s.GetReg8Val("A") - s.GetHLVal() - int(s.GetFlag("C"))
	halfborrow := (s.GetReg8Val("A")&0xF)-(s.GetHLVal()&0xF)-(int(s.GetFlag("C"))&0xF) < 0
	borrow := s.GetReg8Val("A")-s.GetHLVal()-int(s.GetFlag("C")) < 0
	s.SetReg8Val("A", val)
	s.SetFlagSUB(halfborrow, borrow, val)
}

func (s *CPU) SBCAN8(op int) { // reg[A] = reg[A] - n8 - C
	val := s.GetReg8Val("A") - op - int(s.GetFlag("C"))
	halfborrow := (s.GetReg8Val("A")&0xF)-(op&0xF)-(int(s.GetFlag("C"))&0xF) < 0
	borrow := s.GetReg8Val("A")-op-int(s.GetFlag("C")) < 0
	s.SetReg8Val("A", val)
	s.SetFlagSUB(halfborrow, borrow, val)
}

func (s *CPU) XORAR8(reg string) { // reg[A] = reg[A] ^ reg[]
	regVal := s.GetReg8Val(reg)
	val := s.GetReg8Val("A") ^ regVal
	s.SetReg8Val("A", val)
	s.SetFlagOR(val)
}

func (s *CPU) XORAHL() { // reg[A] = reg[A] ^ M[HL]
	val := s.GetReg8Val("A") ^ s.GetHLVal()
	s.SetReg8Val("A", val)
	s.SetFlagOR(val)
}

func (s *CPU) XORAN8(op int) { // reg[A] = reg[A] ^ op
	val := s.GetReg8Val("A") ^ op
	s.SetReg8Val("A", val)
	s.SetFlagOR(val)
}

// Opcodes for 16-bit ALU instructions

func (s *CPU) ADDHLR16(reg string) { // reg[HL] = reg[r16] + reg[HL]
	val := s.GetReg16Val("HL") + s.GetReg16Val(reg)
	s.SetReg16Val("HL", val)
	var tpcarry int
	if (int(s.HL[0])+s.GetReg16Val(reg)&0xFF)&0x100 == 0x100 {
		tpcarry = 1
	}
	halfcarry := (int(s.HL[1]) & 0xF) + (s.GetReg16Val(reg) & 0xF00) + tpcarry
	carry := (int(s.HL[1]) + s.GetReg16Val(reg)&0xFF) & 0x100
	s.SetFlagADD(halfcarry, carry, val)
}

func (s *CPU) DECR16(reg string) { // reg[r16] = reg[r16] - 1
	s.SetReg16Val(reg, s.GetReg16Val(reg)-1)
}

func (s *CPU) INCR16(reg string) { // reg[r16] = reg[r16] + 1
	s.SetReg16Val(reg, s.GetReg16Val(reg)+1)
}
