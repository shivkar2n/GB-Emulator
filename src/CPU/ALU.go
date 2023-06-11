// Opcodes for 8-bit ALU instructions

package CPU

func (s *CPU) ADCAR8(reg string) { // reg[A] = reg[A] + C + reg[]
	regVal := int(int8(s.GetReg8Val(reg)))
	val := int(int8(s.GetReg8Val("A") + regVal + s.GetFlag("C")))
	halfcarry := s.GetReg8Val("A")&0xF + regVal&0xF + s.GetFlag("C")
	carry := s.GetReg8Val("A")&0xFF + regVal&0xFF + s.GetFlag("C")
	s.SetReg8Val("A", val)
	s.SetFlagADD(halfcarry, carry, val)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) ADCAHL() { // reg[A] = reg[A] + C + M[HL]
	val := int(int8(s.GetReg8Val("A") + s.GetHLVal() + s.GetFlag("C")))
	halfcarry := s.GetReg8Val("A")&0xF + s.GetHLVal()&0xF + s.GetFlag("C")
	carry := s.GetReg8Val("A")&0xFF + s.GetHLVal()&0xFF + s.GetFlag("C")
	s.SetReg8Val("A", val)
	s.SetFlagADD(halfcarry, carry, val)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) ADCAN8(op int) { // reg[A] = reg[A] + n8 + C
	val := int(int8(s.GetReg8Val("A") + op + s.GetFlag("C")))
	halfcarry := s.GetReg8Val("A")&0xF + op&0xF + s.GetFlag("C")
	carry := s.GetReg8Val("A")&0xFF + op&0xFF + s.GetFlag("C")
	s.SetReg8Val("A", val)
	s.SetFlagADD(halfcarry, carry, val)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+2)
}

func (s *CPU) ADDAR8(reg string) { // reg[A] = reg[A] + reg[]
	regVal := int(int8(s.GetReg8Val(reg)))
	regAVal := s.GetReg8Val("A")
	val := int(int8(regAVal + regVal))
	halfcarry := regAVal&0xF + regVal&0xF
	carry := regAVal&0xFF + regVal&0xFF
	s.SetReg8Val("A", val)
	s.SetFlagADD(halfcarry, carry, val)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) ADDAHL() { // reg[A] = reg[A] + M[HL]
	val := int(int8(s.GetReg8Val("A") + s.GetHLVal()))
	halfcarry := s.GetReg8Val("A")&0xF + s.GetHLVal()&0xF
	carry := s.GetReg8Val("A")&0xFF + s.GetHLVal()&0xFF
	s.SetReg8Val("A", val)
	s.SetFlagADD(halfcarry, carry, val)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) ADDAN8(op int) { // reg[A] = reg[A] + n8
	val := int(int8(s.GetReg8Val("A") + op))
	halfcarry := s.GetReg8Val("A")&0xF + op&0xF
	carry := s.GetReg8Val("A")&0xFF + op&0xFF
	s.SetReg8Val("A", val)
	s.SetFlagADD(halfcarry, carry, val)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+2)
}

func (s *CPU) ANDAR8(reg string) { // reg[A] = reg[A] & reg[]
	regVal := s.GetReg8Val(reg)
	val := s.GetReg8Val("A") & regVal
	s.SetReg8Val("A", val)
	s.SetFlagAND(val)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) ANDAHL() { // reg[A] = reg[A] & M[HL]
	val := s.GetReg8Val("A") & s.GetHLVal()
	s.SetReg8Val("A", val)
	s.SetFlagAND(val)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) ANDAN8(op int) { // reg[A] = reg[A] & n8
	val := s.GetReg8Val("A") & op
	s.SetReg8Val("A", val)
	s.SetFlagAND(val)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+2)
}

func (s *CPU) CPAR8(reg string) { // res = reg[A] - reg[]
	regVal := s.GetReg8Val(reg)
	halfBorrow := s.GetReg8Val("A")&0x0F - regVal&0x0F
	borrow := s.GetReg8Val("A")&0xFF - regVal&0xFF
	val := s.GetReg8Val("A") - regVal
	s.SetFlagCP(halfBorrow, borrow, val)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) CPAHL() { // res = reg[A] - M[HL]
	halfBorrow := s.GetReg8Val("A")&0x0F - s.GetHLVal()&0x0F
	borrow := s.GetReg8Val("A")&0xFF - s.GetHLVal()&0xFF
	val := int(int8(s.GetReg8Val("A") - s.GetHLVal()))
	s.SetFlagCP(halfBorrow, borrow, val)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) CPAN8(op int) { // res = reg[A] - n8
	halfBorrow := s.GetReg8Val("A")&0x0F - op&0x0F
	borrow := s.GetReg8Val("A")&0xFF - op&0xFF
	val := s.GetReg8Val("A") - op
	s.SetFlagCP(halfBorrow, borrow, val)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+2)
}

func (s *CPU) DAA() { // res = reg[A] - n8
	if s.GetFlag("N") == 0 {
		if s.GetFlag("C") == 1 || s.GetReg8Val("A") > 0x99 {
			s.SetReg8Val("A", s.GetReg8Val("A")+0x60)
			s.SetFlag("C")
		}
		if s.GetFlag("H") == 1 || s.GetReg8Val("A")&0x0F > 0x09 {
			s.SetReg8Val("A", s.GetReg8Val("A")+0x6)
		}
	} else {
		if s.GetFlag("C") == 1 {
			s.SetReg8Val("A", s.GetReg8Val("A")-0x60)
		}
		if s.GetFlag("H") == 1 {
			s.SetReg8Val("A", s.GetReg8Val("A")-0x6)
		}
	}
	if s.GetReg8Val("A") == 0 {
		s.SetFlag("Z")
	} else {
		s.ResetFlag("Z")
	}
	s.ResetFlag("H")
}

func (s *CPU) DECR8(reg string) { // reg[] = reg[] - 1
	val := int(int8(s.GetReg8Val(reg)) - 1)
	halfBorrow := s.GetReg8Val(reg)&0x0F - 1
	s.SetFlagDEC(halfBorrow, val)
	s.SetReg8Val(reg, val)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) DECHL() { // M[HL] = M[HL] - 1
	val := int(int8(s.GetHLVal() - 1))
	halfBorrow := s.GetHLVal()&0x0F - 1
	s.SetHLVal(val)
	s.SetFlagDEC(halfBorrow, val)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) INCR8(reg string) { // reg[] = reg[] + 1
	val := int(int8(s.GetReg8Val(reg) + 1))
	halfcarry := s.GetReg8Val(reg)&0x0F + 1
	s.SetReg8Val(reg, val)
	s.SetFlagINC(halfcarry, val)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) INCRHL() { // M[HL] = M[HL] + 1
	val := int(int8(s.GetHLVal() + 1))
	halfcarry := s.GetHLVal()&0x0F + 1
	s.SetHLVal(val)
	s.SetFlagINC(halfcarry, val)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) ORAR8(reg string) { // reg[A] = reg[A] | reg[]
	regVal := s.GetReg8Val(reg)
	val := s.GetReg8Val("A") | regVal
	s.SetReg8Val("A", val)
	s.SetFlagOR(val)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) ORAHL() { // reg[A] = M[HL] | reg[A]
	val := s.GetReg8Val("A") | s.GetHLVal()
	s.SetReg8Val("A", val)
	s.SetFlagOR(val)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) ORAN8(op int) { // reg[A] = reg[A] | n8
	val := s.GetReg8Val("A") | op
	s.SetReg8Val("A", val)
	s.SetFlagOR(val)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+2)
}

func (s *CPU) SUBAR8(reg string) { // reg[A] = reg[A] - reg[]
	regVal := s.GetReg8Val(reg)
	val := s.GetReg8Val("A") - regVal
	halfborrow := s.GetReg8Val("A")&0xF - regVal&0xF
	borrow := s.GetReg8Val("A") - regVal
	s.SetReg8Val("A", val)
	s.SetFlagSUB(halfborrow, borrow, val)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) SUBAHL() { // reg[A] = reg[A] - M[HL]
	val := s.GetReg8Val("A") - s.GetHLVal()
	halfborrow := s.GetReg8Val("A")&0xF - s.GetHLVal()&0xF
	borrow := s.GetReg8Val("A") - s.GetHLVal()
	s.SetReg8Val("A", val)
	s.SetFlagSUB(halfborrow, borrow, val)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) SUBAN8(op int) { // reg[A] = reg[A] - n8
	val := s.GetReg8Val("A") - op
	halfborrow := s.GetReg8Val("A")&0xF - op&0xF
	borrow := s.GetReg8Val("A") - op
	s.SetReg8Val("A", val)
	s.SetFlagSUB(halfborrow, borrow, val)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+2)
}

func (s *CPU) SBCAR8(reg string) { // reg[A] = reg[A] - reg[] - C
	regVal := int(int8(s.GetReg8Val(reg)))
	val := int(int8(s.GetReg8Val("A") - regVal - s.GetFlag("C")))
	halfborrow := s.GetReg8Val("A")&0xF - regVal&0xF - s.GetFlag("C")
	borrow := s.GetReg8Val("A")&0xFF - regVal&0xFF - s.GetFlag("C")
	s.SetReg8Val("A", val)
	s.SetFlagSUB(halfborrow, borrow, val)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) SBCAHL() { // reg[A] = reg[A] - M[HL] - C
	val := int(int8(s.GetReg8Val("A") - s.GetHLVal() - int(s.GetFlag("C"))))
	halfborrow := s.GetReg8Val("A")&0xF - s.GetHLVal()&0xF - s.GetFlag("C")
	borrow := s.GetReg8Val("A") - s.GetHLVal() - s.GetFlag("C")
	s.SetReg8Val("A", val)
	s.SetFlagSUB(halfborrow, borrow, val)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) SBCAN8(op int) { // reg[A] = reg[A] - n8 - C
	val := int(int8(s.GetReg8Val("A") - op - s.GetFlag("C")))
	halfborrow := s.GetReg8Val("A")&0xF - op&0xF - s.GetFlag("C")
	borrow := s.GetReg8Val("A")&0xFF - op&0xFF - s.GetFlag("C")
	s.SetReg8Val("A", val)
	s.SetFlagSUB(halfborrow, borrow, val)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+2)
}

func (s *CPU) XORAR8(reg string) { // reg[A] = reg[A] ^ reg[]
	regVal := s.GetReg8Val(reg)
	val := s.GetReg8Val("A") ^ regVal
	s.SetReg8Val("A", val)
	s.SetFlagOR(val)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) XORAHL() { // reg[A] = reg[A] ^ M[HL]
	val := s.GetReg8Val("A") ^ s.GetHLVal()
	s.SetReg8Val("A", val)
	s.SetFlagOR(val)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) XORAN8(op int) { // reg[A] = reg[A] ^ op
	val := s.GetReg8Val("A") ^ op
	s.SetReg8Val("A", val)
	s.SetFlagOR(val)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+2)
}

// Opcodes for 16-bit ALU instructions
// half carry -> carry from 11th bit
// carry -> carry from 15th bit

func (s *CPU) ADDHLR16(reg string) { // reg[HL] = reg[r16] + reg[HL]
	val := int(uint16(s.GetReg16Val("HL") + s.GetReg16Val(reg)))
	var tpcarry int
	if (s.GetReg16Val("HL")&0xFF+s.GetReg16Val(reg)&0xFF)&0x100 == 0x100 {
		tpcarry = 1
	} else {
		tpcarry = 0
	}
	halfcarry := ((s.GetReg16Val("HL")&0x0F00)>>8 + (s.GetReg16Val(reg)&0x0F00)>>8 + tpcarry) >> 4
	carry := val >> 16
	s.SetReg16Val("HL", val)
	s.SetFlagADD16(halfcarry, carry)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) DECR16(reg string) { // reg[r16] = reg[r16] - 1
	s.SetReg16Val(reg, int(int16(s.GetReg16Val(reg)-1)))
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}

func (s *CPU) INCR16(reg string) { // reg[r16] = reg[r16] + 1
	s.SetReg16Val(reg, int(int16(s.GetReg16Val(reg)+1)))
	s.SetReg16Val("PC", s.GetReg16Val("PC")+1)
}
