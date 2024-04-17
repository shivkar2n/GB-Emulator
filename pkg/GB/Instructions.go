package GB

import "fmt"

// OPCODES FOR 8-BIT ALU INSTRUCTIONS
func (GB *GB) ADCAR8(reg string) (string, int, int, bool) { // reg[A] = reg[A] + C + reg[]
	opcode := fmt.Sprintf("ADC A,%s", reg)
	regVal := GB.CPU.Reg.Read(reg)
	val := int(int8(GB.CPU.Reg.Read("A") + regVal + GB.CPU.GetFlag("C")))
	halfcarry := GB.CPU.Reg.Read("A")&0xF + regVal&0xF + GB.CPU.GetFlag("C")
	carry := GB.CPU.Reg.Read("A")&0xFF + regVal&0xFF + GB.CPU.GetFlag("C")
	GB.CPU.Reg.Write(val, "A")
	GB.CPU.SetFlagADD(halfcarry, carry, val)
	return opcode, 1, 4, true
}

func (GB *GB) ADCAHL() (string, int, int, bool) { // reg[A] = reg[A] + C + M[reg[HL]]
	opcode := fmt.Sprintf("ADC A,(HL)")
	val := int(int8(GB.CPU.Reg.Read("A") + GB.GetHLVal() + GB.CPU.GetFlag("C")))
	halfcarry := GB.CPU.Reg.Read("A")&0xF + GB.GetHLVal()&0xF + GB.CPU.GetFlag("C")
	carry := GB.CPU.Reg.Read("A")&0xFF + GB.GetHLVal()&0xFF + GB.CPU.GetFlag("C")
	GB.CPU.Reg.Write(val, "A")
	GB.CPU.SetFlagADD(halfcarry, carry, val)
	return opcode, 1, 8, true
}

func (GB *GB) ADCAN8() (string, int, int, bool) { // reg[A] = reg[A] + n8 + C
	opcode := fmt.Sprintf("ADC A,$%x", GB.MMU.Read(GB.CPU.Reg.Read("PC")+1))
	addr := GB.MMU.Read(GB.CPU.Reg.Read("PC") + 1)
	val := int(int8(GB.CPU.Reg.Read("A") + addr + GB.CPU.GetFlag("C")))
	halfcarry := GB.CPU.Reg.Read("A")&0xF + addr&0xF + GB.CPU.GetFlag("C")
	carry := GB.CPU.Reg.Read("A")&0xFF + addr&0xFF + GB.CPU.GetFlag("C")
	GB.CPU.Reg.Write(val, "A")
	GB.CPU.SetFlagADD(halfcarry, carry, val)
	return opcode, 2, 8, true
}

func (GB *GB) ADDAR8(reg string) (string, int, int, bool) { // reg[A] = reg[A] + reg[]
	opcode := fmt.Sprintf("ADD A,%s", reg)
	regVal := GB.CPU.Reg.Read(reg)
	regAVal := GB.CPU.Reg.Read("A")
	val := int(int8(regAVal + regVal))
	halfcarry := regAVal&0xF + regVal&0xF
	carry := regAVal&0xFF + regVal&0xFF
	GB.CPU.Reg.Write(val, "A")
	GB.CPU.SetFlagADD(halfcarry, carry, val)
	return opcode, 1, 4, true
}

func (GB *GB) ADDAHL() (string, int, int, bool) { // reg[A] = reg[A] + M[reg[HL]]
	opcode := fmt.Sprintf("ADD A,(HL)")
	val := int(int8(GB.CPU.Reg.Read("A") + GB.GetHLVal()))
	halfcarry := GB.CPU.Reg.Read("A")&0xF + GB.GetHLVal()&0xF
	carry := GB.CPU.Reg.Read("A")&0xFF + GB.GetHLVal()&0xFF
	GB.CPU.Reg.Write(val, "A")
	GB.CPU.SetFlagADD(halfcarry, carry, val)
	return opcode, 1, 8, true
}

func (GB *GB) ADDAN8() (string, int, int, bool) { // reg[A] = reg[A] + n8
	op := GB.MMU.Read(GB.CPU.Reg.Read("PC") + 1)
	opcode := fmt.Sprintf("ADD A,$%x", uint8(op))
	val := int(int8(GB.CPU.Reg.Read("A") + op))
	halfcarry := GB.CPU.Reg.Read("A")&0xF + op&0xF
	carry := GB.CPU.Reg.Read("A")&0xFF + op&0xFF
	GB.CPU.Reg.Write(val, "A")
	GB.CPU.SetFlagADD(halfcarry, carry, val)
	return opcode, 2, 8, true
}

func (GB *GB) ANDAR8(reg string) (string, int, int, bool) { // reg[A] = reg[A] & reg[]
	opcode := fmt.Sprintf("AND A,%s", reg)
	regVal := GB.CPU.Reg.Read(reg)
	val := GB.CPU.Reg.Read("A") & regVal
	GB.CPU.Reg.Write(val, "A")
	GB.CPU.SetFlagAND(val)
	return opcode, 1, 4, true
}

func (GB *GB) ANDAHL() (string, int, int, bool) { // reg[A] = reg[A] & M[HL]
	opcode := fmt.Sprintf("AND A,(HL)")
	val := GB.CPU.Reg.Read("A") & GB.GetHLVal()
	GB.CPU.Reg.Write(val, "A")
	GB.CPU.SetFlagAND(val)
	return opcode, 1, 8, true
}

func (GB *GB) ANDAN8() (string, int, int, bool) { // reg[A] = reg[A] & n8
	op := GB.MMU.Read(GB.CPU.Reg.Read("PC") + 1)
	opcode := fmt.Sprintf("AND A,$%x", uint8(op))
	val := GB.CPU.Reg.Read("A") & op
	GB.CPU.Reg.Write(val, "A")
	GB.CPU.SetFlagAND(val)
	return opcode, 2, 8, true
}

func (GB *GB) CPAR8(reg string) (string, int, int, bool) { // res = reg[A] - reg[]
	opcode := fmt.Sprintf("CP A,%s", reg)
	regVal := GB.CPU.Reg.Read(reg)
	halfBorrow := GB.CPU.Reg.Read("A")&0x0F - regVal&0x0F
	borrow := GB.CPU.Reg.Read("A")&0xFF - regVal&0xFF
	val := int(int8(GB.CPU.Reg.Read("A") - regVal))
	GB.CPU.SetFlagCP(halfBorrow, borrow, val)
	return opcode, 1, 4, true
}

func (GB *GB) CPAHL() (string, int, int, bool) { // res = reg[A] - M[HL]
	opcode := fmt.Sprintf("CP A,(HL)")
	halfBorrow := GB.CPU.Reg.Read("A")&0x0F - GB.GetHLVal()&0x0F
	borrow := GB.CPU.Reg.Read("A")&0xFF - GB.GetHLVal()&0xFF
	val := int(int8(GB.CPU.Reg.Read("A") - GB.GetHLVal()))
	GB.CPU.SetFlagCP(halfBorrow, borrow, val)
	return opcode, 1, 8, true
}

func (GB *GB) CPAN8() (string, int, int, bool) { // res = reg[A] - n8
	op := GB.MMU.Read(GB.CPU.Reg.Read("PC") + 1)
	opcode := fmt.Sprintf("CP $%x", uint8(op))
	halfBorrow := GB.CPU.Reg.Read("A")&0x0F - op&0x0F
	borrow := GB.CPU.Reg.Read("A")&0xFF - op&0xFF
	val := int(int8(GB.CPU.Reg.Read("A") - op))
	GB.CPU.SetFlagCP(halfBorrow, borrow, val)
	return opcode, 2, 8, true
}

func (GB *GB) DAA() (string, int, int, bool) { // res = reg[A] - n8
	opcode := fmt.Sprintf("DAA")
	if GB.CPU.GetFlag("N") == 0 {
		if GB.CPU.GetFlag("C") == 1 || GB.CPU.Reg.Read("A") > 0x99 {
			GB.CPU.Reg.Write(int(int8(GB.CPU.Reg.Read("A")+0x60)), "A")
			GB.CPU.SetFlag("C")
		}
		if GB.CPU.GetFlag("H") == 1 || GB.CPU.Reg.Read("A")&0x0F > 0x09 {
			GB.CPU.Reg.Write(int(int8(GB.CPU.Reg.Read("A")+0x6)), "A")
		}
	} else {
		if GB.CPU.GetFlag("C") == 1 {
			GB.CPU.Reg.Write(int(int8(GB.CPU.Reg.Read("A")-0x60)), "A")
		}
		if GB.CPU.GetFlag("H") == 1 {
			GB.CPU.Reg.Write(int(int8(GB.CPU.Reg.Read("A")-0x6)), "A")
		}
	}
	if GB.CPU.Reg.Read("A") == 0 {
		GB.CPU.SetFlag("Z")
	} else {
		GB.CPU.ResetFlag("Z")
	}
	GB.CPU.ResetFlag("H")
	return opcode, 1, 4, true
}

func (GB *GB) DECR8(reg string) (string, int, int, bool) { // reg[] = reg[] - 1
	opcode := fmt.Sprintf("DEC %s", reg)
	val := int(int8(GB.CPU.Reg.Read(reg)) - 1)
	halfBorrow := GB.CPU.Reg.Read(reg)&0x0F - 1
	GB.CPU.SetFlagDEC(halfBorrow, val)
	GB.CPU.Reg.Write(val, reg)
	return opcode, 1, 4, true
}

func (GB *GB) DECHL() (string, int, int, bool) { // M[HL] = M[HL] - 1
	opcode := fmt.Sprintf("DEC (HL)")
	val := int(int8(GB.GetHLVal() - 1))
	halfBorrow := GB.GetHLVal()&0x0F - 1
	GB.SetHLVal(val)
	GB.CPU.SetFlagDEC(halfBorrow, val)
	return opcode, 1, 12, true
}

func (GB *GB) INCR8(reg string) (string, int, int, bool) { // reg[] = reg[] + 1
	opcode := fmt.Sprintf("INC %s", reg)
	val := int(int8(GB.CPU.Reg.Read(reg) + 1))
	halfcarry := GB.CPU.Reg.Read(reg)&0x0F + 1
	GB.CPU.Reg.Write(val, reg)
	GB.CPU.SetFlagINC(halfcarry, val)
	return opcode, 1, 4, true
}

func (GB *GB) INCRHL() (string, int, int, bool) { // M[HL] = M[HL] + 1
	opcode := fmt.Sprintf("INC (HL)")
	val := int(int8(GB.GetHLVal() + 1))
	halfcarry := GB.GetHLVal()&0x0F + 1
	GB.SetHLVal(val)
	GB.CPU.SetFlagINC(halfcarry, val)
	return opcode, 1, 12, true
}

func (GB *GB) ORAR8(reg string) (string, int, int, bool) { // reg[A] = reg[A] | reg[]
	opcode := fmt.Sprintf("OR A,%s", reg)
	val := GB.CPU.Reg.Read("A") | GB.CPU.Reg.Read(reg)
	GB.CPU.Reg.Write(val, "A")
	GB.CPU.SetFlagOR(val)
	return opcode, 1, 4, true
}

func (GB *GB) ORAHL() (string, int, int, bool) { // reg[A] = M[HL] | reg[A]
	opcode := fmt.Sprintf("OR A,(HL)")
	val := GB.CPU.Reg.Read("A") | GB.GetHLVal()
	GB.CPU.Reg.Write(val, "A")
	GB.CPU.SetFlagOR(val)
	return opcode, 1, 8, true
}

func (GB *GB) ORAN8() (string, int, int, bool) { // reg[A] = reg[A] | n8
	op := GB.MMU.Read(GB.CPU.Reg.Read("PC") + 1)
	opcode := fmt.Sprintf("OR A,$%x", int8(op))
	val := GB.CPU.Reg.Read("A") | op
	GB.CPU.Reg.Write(val, "A")
	GB.CPU.SetFlagOR(val)
	return opcode, 2, 8, true
}

func (GB *GB) SUBAR8(reg string) (string, int, int, bool) { // reg[A] = reg[A] - reg[]
	opcode := fmt.Sprintf("SUB A,%s", reg)
	regVal := GB.CPU.Reg.Read(reg)
	val := GB.CPU.Reg.Read("A") - regVal
	halfborrow := GB.CPU.Reg.Read("A")&0xF - regVal&0xF
	borrow := GB.CPU.Reg.Read("A") - regVal
	GB.CPU.Reg.Write(val, "A")
	GB.CPU.SetFlagSUB(halfborrow, borrow, val)
	return opcode, 1, 4, true
}

func (GB *GB) SUBAHL() (string, int, int, bool) { // reg[A] = reg[A] - M[HL]
	opcode := fmt.Sprintf("SUB A,(HL)")
	val := GB.CPU.Reg.Read("A") - GB.GetHLVal()
	halfborrow := GB.CPU.Reg.Read("A")&0xF - GB.GetHLVal()&0xF
	borrow := GB.CPU.Reg.Read("A") - GB.GetHLVal()
	GB.CPU.Reg.Write(val, "A")
	GB.CPU.SetFlagSUB(halfborrow, borrow, val)
	return opcode, 1, 8, true
}

func (GB *GB) SUBAN8() (string, int, int, bool) { // reg[A] = reg[A] - n8
	opcode := fmt.Sprintf("SUB A,$%x", GB.MMU.Read(GB.CPU.Reg.Read("PC")+1))
	addr := GB.MMU.Read(GB.CPU.Reg.Read("PC") + 1)
	val := GB.CPU.Reg.Read("A") - addr
	halfborrow := GB.CPU.Reg.Read("A")&0xF - addr&0xF
	borrow := GB.CPU.Reg.Read("A") - addr
	GB.CPU.Reg.Write(val, "A")
	GB.CPU.SetFlagSUB(halfborrow, borrow, val)
	return opcode, 2, 8, true
}

func (GB *GB) SBCAR8(reg string) (string, int, int, bool) { // reg[A] = reg[A] - reg[] - C
	opcode := fmt.Sprintf("SBC A,B")
	regVal := GB.CPU.Reg.Read(reg)
	val := int(int8(GB.CPU.Reg.Read("A") - regVal - GB.CPU.GetFlag("C")))
	halfborrow := GB.CPU.Reg.Read("A")&0xF - regVal&0xF - GB.CPU.GetFlag("C")
	borrow := GB.CPU.Reg.Read("A")&0xFF - regVal&0xFF - GB.CPU.GetFlag("C")
	GB.CPU.Reg.Write(val, "A")
	GB.CPU.SetFlagSUB(halfborrow, borrow, val)
	return opcode, 1, 4, true
}

func (GB *GB) SBCAHL() (string, int, int, bool) { // reg[A] = reg[A] - M[HL] - C
	opcode := fmt.Sprintf("SBC A,(HL)")
	val := int(int8(GB.CPU.Reg.Read("A") - GB.GetHLVal() - GB.CPU.GetFlag("C")))
	halfborrow := GB.CPU.Reg.Read("A")&0xF - GB.GetHLVal()&0xF - GB.CPU.GetFlag("C")
	borrow := GB.CPU.Reg.Read("A") - GB.GetHLVal() - GB.CPU.GetFlag("C")
	GB.CPU.Reg.Write(val, "A")
	GB.CPU.SetFlagSUB(halfborrow, borrow, val)
	return opcode, 1, 8, true
}

func (GB *GB) SBCAN8() (string, int, int, bool) { // reg[A] = reg[A] - n8 - C
	op := GB.MMU.Read(GB.CPU.Reg.Read("PC") + 1)
	opcode := fmt.Sprintf("SBC A,$%x", uint8(op))
	val := int(int8(GB.CPU.Reg.Read("A") - op - GB.CPU.GetFlag("C")))
	halfborrow := GB.CPU.Reg.Read("A")&0xF - op&0xF - GB.CPU.GetFlag("C")
	borrow := GB.CPU.Reg.Read("A")&0xFF - op&0xFF - GB.CPU.GetFlag("C")
	GB.CPU.Reg.Write(val, "A")
	GB.CPU.SetFlagSUB(halfborrow, borrow, val)
	return opcode, 2, 8, true
}

func (GB *GB) XORAR8(reg string) (string, int, int, bool) { // reg[A] = reg[A] ^ reg[]
	opcode := fmt.Sprintf("XOR A,%s", reg)
	val := GB.CPU.Reg.Read("A") ^ GB.CPU.Reg.Read(reg)
	GB.CPU.Reg.Write(val, "A")
	GB.CPU.SetFlagOR(val)
	return opcode, 1, 4, true
}

func (GB *GB) XORAHL() (string, int, int, bool) { // reg[A] = reg[A] ^ M[HL]
	opcode := fmt.Sprintf("XOR A,(HL)")
	val := GB.CPU.Reg.Read("A") ^ GB.GetHLVal()
	GB.CPU.Reg.Write(val, "A")
	GB.CPU.SetFlagOR(val)
	return opcode, 1, 8, true
}

func (GB *GB) XORAN8() (string, int, int, bool) { // reg[A] = reg[A] ^ op
	opcode := fmt.Sprintf("XOR $%x", uint8(GB.MMU.Read(GB.CPU.Reg.Read("PC")+1)))
	op := GB.MMU.Read(GB.CPU.Reg.Read("PC") + 1)
	val := GB.CPU.Reg.Read("A") ^ op
	GB.CPU.Reg.Write(val, "A")
	GB.CPU.SetFlagOR(val)
	return opcode, 2, 8, true
}

// Opcodes for 16-bit ALU instructions
// half carry -> carry from 11th bit
// carry -> carry from 15th bit

func (GB *GB) ADDHLR16(reg string) (string, int, int, bool) { // reg[HL] = reg[r16] + reg[HL]
	opcode := fmt.Sprintf("ADD HL, %s", reg)
	val := int(GB.CPU.Reg.Read("HL") + GB.CPU.Reg.Read(reg))
	var tpcarry int
	if (GB.CPU.Reg.Read("HL")&0xFF+GB.CPU.Reg.Read(reg)&0xFF)&0x100 == 0x100 {
		tpcarry = 1
	} else {
		tpcarry = 0
	}
	halfcarry := ((GB.CPU.Reg.Read("HL")&0x0F00)>>8 + (GB.CPU.Reg.Read(reg)&0x0F00)>>8 + tpcarry) >> 4
	carry := val >> 16
	GB.CPU.Reg.Write(int(uint16(val)), "HL")
	GB.CPU.SetFlagADD16(halfcarry, carry)
	return opcode, 1, 8, true
}

func (GB *GB) DECR16(reg string) (string, int, int, bool) { // reg[r16] = reg[r16] - 1
	opcode := fmt.Sprintf("DEC %s", reg)
	GB.CPU.Reg.Write(int(int16(GB.CPU.Reg.Read(reg)-1)), reg)
	return opcode, 1, 8, true
}

func (GB *GB) INCR16(reg string) (string, int, int, bool) { // reg[r16] = reg[r16] + 1
	opcode := fmt.Sprintf("INC %s", reg)
	GB.CPU.Reg.Write(int(int16(GB.CPU.Reg.Read(reg)+1)), reg)
	return opcode, 1, 8, true
}

// OPCODES FOR BITSHIFT INSTRUCTIONS
func (GB *GB) RLR8(reg string) (string, int, int, bool) { // Rotate bits reg[] left through Flag C
	opcode := fmt.Sprintf("RL %s", reg)
	carry := GB.CPU.Reg.Read(reg) >> 7
	val := int(int8(GB.CPU.Reg.Read(reg)<<1) | int8(GB.CPU.GetFlag("C")))
	GB.CPU.Reg.Write(val, reg)
	GB.CPU.SetFlagRL(val, carry)
	return opcode, 2, 8, true
}

func (GB *GB) RLHL() (string, int, int, bool) { // Rotate bits M[reg[HL]] left through Flag C
	opcode := fmt.Sprintf("RL (HL)")
	carry := GB.GetHLVal() >> 7
	val := int(int8((GB.GetHLVal() << 1) | GB.CPU.GetFlag("C")))
	GB.SetHLVal(val)
	GB.CPU.SetFlagRL(val, carry)
	return opcode, 2, 16, true
}

func (GB *GB) RLA() (string, int, int, bool) { // Rotate bits reg[A] left through Flag C
	opcode := fmt.Sprintf("RLA")
	carry := GB.CPU.Reg.Read("A") >> 7
	val := int(int8((GB.CPU.Reg.Read("A") << 1) | GB.CPU.GetFlag("C")))
	GB.CPU.Reg.Write(val, "A")
	GB.CPU.SetFlagRL(1, carry)
	return opcode, 1, 4, true
}

func (GB *GB) RLCR8(reg string) (string, int, int, bool) { // Rotate bits reg[A] left through Flag c
	opcode := fmt.Sprintf("RLC %s", reg)
	carry := GB.CPU.Reg.Read(reg) >> 7
	val := int(int8((GB.CPU.Reg.Read(reg) << 1) | carry))
	GB.CPU.Reg.Write(val, reg)
	GB.CPU.SetFlagRL(val, carry)
	return opcode, 2, 8, true
}

func (GB *GB) RLCHL() (string, int, int, bool) { // Rotate bits M[reg[HL]] left through flag C
	opcode := fmt.Sprintf("RLC (HL)")
	carry := GB.GetHLVal() >> 7
	val := int(int8((GB.GetHLVal() << 1) | carry))
	GB.SetHLVal(val)
	GB.CPU.SetFlagRL(val, carry)
	return opcode, 2, 16, true
}

func (GB *GB) RLCA() (string, int, int, bool) { // Rotate bits reg[A] left
	opcode := fmt.Sprintf("RLCA")
	carry := GB.CPU.Reg.Read("A") >> 7
	val := int(int8((GB.CPU.Reg.Read("A") << 1) | carry))
	GB.CPU.Reg.Write(val, "A")
	GB.CPU.SetFlagRL(1, carry)
	return opcode, 1, 4, true
}

func (GB *GB) RRR8(reg string) (string, int, int, bool) { // Rotate bits reg[] right through flag C
	opcode := fmt.Sprintf("RR %s", reg)
	carry := GB.CPU.Reg.Read(reg) & 0x01
	val := int(int8((GB.CPU.Reg.Read(reg) >> 1) | (GB.CPU.GetFlag("C") << 7)))
	GB.CPU.Reg.Write(val, reg)
	GB.CPU.SetFlagRL(val, carry)
	return opcode, 2, 8, true
}

func (GB *GB) RRHL() (string, int, int, bool) { // Rotate bits M[reg[HL]] right through flag C
	opcode := fmt.Sprintf("RR (HL)")
	carry := GB.GetHLVal() & 0x01
	val := int(int8((GB.GetHLVal() >> 1) | (GB.CPU.GetFlag("C") << 7)))
	GB.SetHLVal(val)
	GB.CPU.SetFlagRL(val, carry)
	return opcode, 2, 16, true
}

func (GB *GB) RRA() (string, int, int, bool) { // Rotate bits reg[A] right through flag C
	opcode := fmt.Sprintf("RRA")
	carry := GB.CPU.Reg.Read("A") & 0x01
	val := (GB.CPU.Reg.Read("A") >> 1) | (GB.CPU.GetFlag("C") << 7)
	GB.CPU.Reg.Write(val, "A")
	GB.CPU.SetFlagRL(1, carry)
	return opcode, 1, 4, true
}

func (GB *GB) RRCR8(reg string) (string, int, int, bool) { // Rotate reg[] right
	opcode := fmt.Sprintf("RRC %s", reg)
	carry := GB.CPU.Reg.Read(reg) & 0x01
	val := (GB.CPU.Reg.Read(reg) >> 1) | (carry << 7)
	GB.CPU.Reg.Write(val, reg)
	GB.CPU.SetFlagRL(val, carry)
	return opcode, 2, 8, true
}

func (GB *GB) RRCHL() (string, int, int, bool) { // Rotate M[reg[HL]] right
	opcode := fmt.Sprintf("RRC (HL)")
	carry := GB.GetHLVal() & 0x01
	val := (GB.GetHLVal() >> 1) | (carry << 7)
	GB.SetHLVal(val)
	GB.CPU.SetFlagRL(val, carry)
	return opcode, 2, 16, true
}

func (GB *GB) RRCA() (string, int, int, bool) { // Rotate reg[A] right
	opcode := fmt.Sprintf("RRCA ")
	carry := GB.CPU.Reg.Read("A") & 0x01
	val := (GB.CPU.Reg.Read("A") >> 1) | (carry << 7)
	GB.CPU.Reg.Write(val, "A")
	GB.CPU.SetFlagRL(1, carry)
	return opcode, 1, 4, true
}

func (GB *GB) SLAR8(reg string) (string, int, int, bool) { // reg[] = reg[] << 1
	opcode := fmt.Sprintf("SLA B")
	carry := GB.CPU.Reg.Read(reg) >> 7
	val := int(int8(GB.CPU.Reg.Read(reg) << 1))
	GB.CPU.Reg.Write(val, reg)
	GB.CPU.SetFlagRL(val, carry)
	return opcode, 2, 8, true
}

func (GB *GB) SLAHL() (string, int, int, bool) { // M[reg[HL]] = M[reg[HL]] << 1
	opcode := fmt.Sprintf("SLA (HL)")
	carry := GB.GetHLVal() >> 7
	val := int(int8(GB.GetHLVal() << 1))
	GB.SetHLVal(val)
	GB.CPU.SetFlagRL(val, carry)
	return opcode, 2, 16, true
}

func (GB *GB) SRAR8(reg string) (string, int, int, bool) { // reg[] = reg[] >> 1, bit 7 unchanged
	opcode := fmt.Sprintf("SRA %s", reg)
	carry := GB.CPU.Reg.Read(reg) & 0x01
	val := (GB.CPU.Reg.Read(reg) >> 1) | GB.CPU.Reg.Read(reg)&0x80
	GB.CPU.Reg.Write(val, reg)
	GB.CPU.SetFlagRL(val, carry)
	return opcode, 2, 8, true
}

func (GB *GB) SRAHL() (string, int, int, bool) { // M[reg[HL]] = M[reg[HL]] >> 1, bit 7 unchanged
	opcode := fmt.Sprintf("SRA (HL)")
	carry := GB.GetHLVal() & 0x01
	val := (GB.GetHLVal() >> 1) | GB.GetHLVal()&0x80
	GB.SetHLVal(val)
	GB.CPU.SetFlagRL(val, carry)
	return opcode, 2, 16, true
}

func (GB *GB) SRLR8(reg string) (string, int, int, bool) { // reg[] = reg[] >> 1
	opcode := fmt.Sprintf("SRL %s", reg)
	carry := GB.CPU.Reg.Read(reg) & 0x01
	val := GB.CPU.Reg.Read(reg) >> 1
	GB.CPU.Reg.Write(val, reg)
	GB.CPU.SetFlagRL(val, carry)
	return opcode, 2, 8, true
}

func (GB *GB) SRLHL() (string, int, int, bool) { // M[reg[HL]] = M[reg[HL]] >> 1
	opcode := fmt.Sprintf("SRL (HL)")
	carry := GB.GetHLVal() & 0x01
	val := GB.GetHLVal() >> 1
	GB.SetHLVal(val)
	GB.CPU.SetFlagRL(val, carry)
	return opcode, 2, 16, true
}

// OPCODES FOR BITWISE INSTRUCTIONS
func (GB *GB) BITU3R8(offset int, reg string) (string, int, int, bool) { // if reg[] >> u3 == 0 then setZeroFlag
	opcode := fmt.Sprintf("BIT %d,%s", offset, reg)
	val := (GB.CPU.Reg.Read(reg) >> offset) & 0x01
	GB.CPU.SetFlagBIT(val)
	return opcode, 2, 8, true
}

func (GB *GB) BITU3HL(offset int) (string, int, int, bool) { // if M[reg[HL]] >> u3 == 0 then setZeroFlag
	opcode := fmt.Sprintf("BIT %d,(HL)", offset)
	val := (GB.GetHLVal() >> offset) & 0x01
	GB.CPU.SetFlagBIT(val)
	return opcode, 2, 12, true
}

func (GB *GB) RESU3R8(offset int, reg string) (string, int, int, bool) { // Set bit u3 in reg[] to 0
	opcode := fmt.Sprintf("RES %d,%s", offset, reg)
	k := 0
	for i := 0; i < 8; i++ {
		k = k << 1
		if 7-i != offset {
			k += 1
		}
	}
	GB.CPU.Reg.Write(GB.CPU.Reg.Read(reg)&k, reg)
	return opcode, 2, 8, true
}

func (GB *GB) RESU3HL(offset int) (string, int, int, bool) { // Set bit u3 in reg[HL] to 0
	opcode := fmt.Sprintf("RES %d,(HL)", offset)
	k := 0
	for i := 0; i < 8; i++ {
		k = k << 1
		if 7-i != offset {
			k += 1
		}
	}
	GB.SetHLVal(GB.GetHLVal() & k)
	return opcode, 2, 16, true
}

func (GB *GB) SETU3R8(offset int, reg string) (string, int, int, bool) { // Set bit u3 in reg[] to 1
	opcode := fmt.Sprintf("SET %d,%s", offset, reg)
	k := 1
	for i := 0; i < offset; i++ {
		k = k << 1
	}
	GB.CPU.Reg.Write(GB.CPU.Reg.Read(reg)|k, reg)
	return opcode, 2, 8, true
}

func (GB *GB) SETU3HL(offset int) (string, int, int, bool) { // Set bit u3 in M[reg[HL]] to 1
	opcode := fmt.Sprintf("SET %d,(HL)", offset)
	k := 1
	for i := 0; i < offset; i++ {
		k = k << 1
	}
	GB.SetHLVal(GB.GetHLVal() | k)
	return opcode, 2, 16, true
}

func (GB *GB) SWAPR8(reg string) (string, int, int, bool) { // Swap upper, lower 4 bits in reg[]
	opcode := fmt.Sprintf("SWAP B")
	low := (GB.CPU.Reg.Read(reg) & 0xF)
	high := (GB.CPU.Reg.Read(reg) & 0xF0) >> 4
	val := low<<4 | high
	GB.CPU.Reg.Write(val, reg)
	GB.CPU.SetFlagSWAP(val)
	return opcode, 2, 8, true
}

func (GB *GB) SWAPHL() (string, int, int, bool) { // Swap upper, lower 4 bits in mem[reg[HL]]
	opcode := fmt.Sprintf("SWAP (HL)")
	low := GB.GetHLVal() & 0xF
	high := (GB.GetHLVal() & 0xF0) >> 4
	val := low<<4 | high
	GB.SetHLVal(val)
	GB.CPU.SetFlagSWAP(val)
	return opcode, 2, 16, true
}

// OPCODES FOR LOAD INSTRUCTIONS
// n8 - 8 bit int
// u8 - 8 bit int unsigned
// n16 - 16 bit const
// e8 - 8 bit offset
// u3 - 3 bit unsigned
func (GB *GB) LDR8R8(reg1 string, reg2 string) (string, int, int, bool) { // reg[reg1] = reg[reg2]
	opcode := fmt.Sprintf("LD %s,%s", reg1, reg2)
	GB.CPU.Reg.Write(GB.CPU.Reg.Read(reg2), reg1)
	return opcode, 1, 4, true
}

func (GB *GB) LDR8N8(reg string) (string, int, int, bool) { // reg[] = n8
	opcode := fmt.Sprintf("LD %s,$%x", reg, uint8(GB.MMU.Read(GB.CPU.Reg.Read("PC")+1)))
	addr := GB.MMU.Read(GB.CPU.Reg.Read("PC") + 1)
	GB.CPU.Reg.Write(addr, reg)
	return opcode, 2, 8, true
}

func (GB *GB) LDR16N16(reg string) (string, int, int, bool) { // reg16[reg] = n16
	opcode := fmt.Sprintf("LD %s,$%x ", reg, uint16(GB.MMU.Read(GB.CPU.Reg.Read("PC")+1)+GB.MMU.Read(GB.CPU.Reg.Read("PC")+2)<<8))
	addr := int(uint16(GB.MMU.Read(GB.CPU.Reg.Read("PC")+1) + (GB.MMU.Read(GB.CPU.Reg.Read("PC")+2) << 8)))
	GB.CPU.Reg.Write(addr, reg)
	return opcode, 3, 12, true
}

func (GB *GB) LDHLR8(reg string) (string, int, int, bool) { // M[HL] = reg[]
	opcode := fmt.Sprintf("LD (HL),%s", reg)
	GB.SetHLVal(GB.CPU.Reg.Read(reg))
	return opcode, 1, 8, true
}

func (GB *GB) LDHLN8() (string, int, int, bool) { // M[HL] = n8
	opcode := fmt.Sprintf("LD (HL),$%x", uint8(GB.MMU.Read(GB.CPU.Reg.Read("PC")+1)))
	op := GB.MMU.Read(GB.CPU.Reg.Read("PC") + 1)
	GB.SetHLVal(op)
	return opcode, 2, 12, true
}

func (GB *GB) LDR8HL(reg string) (string, int, int, bool) { // reg[] = M[HL]
	opcode := fmt.Sprintf("LD %s,(HL)", reg)
	GB.CPU.Reg.Write(GB.GetHLVal(), reg)
	return opcode, 1, 12, true
}

func (GB *GB) LDR16A(reg string) (string, int, int, bool) { // M[reg] = reg[A]
	opcode := fmt.Sprintf("LD (%s), A", reg)
	GB.MMU.Write(GB.CPU.Reg.Read("A"), GB.CPU.Reg.Read(reg))
	return opcode, 1, 8, true
}

func (GB *GB) LDHCA() (string, int, int, bool) { // M[$FF00+C] = reg[A]
	opcode := fmt.Sprintf("LD ($FF00+C),A")
	GB.MMU.Write(GB.CPU.Reg.Read("A"), int(uint16(0xFF00+int(GB.CPU.Reg.Read("C")))))
	return opcode, 1, 8, true
}

func (GB *GB) LDN16A() (string, int, int, bool) { // M[n16] = reg[A]
	addr := int(uint16(GB.MMU.Read(GB.CPU.Reg.Read("PC")+1) + GB.MMU.Read(GB.CPU.Reg.Read("PC")+2)<<8))
	opcode := fmt.Sprintf("LD ($%x),A", addr)
	GB.MMU.Write(GB.CPU.Reg.Read("A"), addr)
	return opcode, 3, 16, true
}

func (GB *GB) LDHU8A() (string, int, int, bool) { // M[$FF00+n8] = reg[A]
	addr := int(uint16(0xFF00 + GB.MMU.Read(GB.CPU.Reg.Read("PC")+1)))
	opcode := fmt.Sprintf("LD $(%x),A", addr)
	GB.MMU.Write(GB.CPU.Reg.Read("A"), addr)
	return opcode, 2, 12, true
}

func (GB *GB) LDAR16(reg string) (string, int, int, bool) { // reg[A] = M[reg[]]
	opcode := fmt.Sprintf("LD A, (%s)", reg)
	GB.CPU.Reg.Write(GB.MMU.Read(GB.CPU.Reg.Read(reg)), "A")
	return opcode, 1, 8, true
}

func (GB *GB) LDHAU8() (string, int, int, bool) { // reg[A] = M[$FF00+u8]
	addr := int(uint16(0xFF00 + int(uint8(GB.MMU.Read(GB.CPU.Reg.Read("PC")+1)))))
	opcode := fmt.Sprintf("LD A,($%x)", addr)
	GB.CPU.Reg.Write(GB.MMU.Read(addr), "A")
	return opcode, 2, 12, true
}

func (GB *GB) LDAN16() (string, int, int, bool) { // reg[A] = M[u16]
	addr := int(uint16(GB.MMU.Read(GB.CPU.Reg.Read("PC")+1) + (GB.MMU.Read(GB.CPU.Reg.Read("PC")+2) << 8)))
	opcode := fmt.Sprintf("LD A,($%x)", addr)
	GB.CPU.Reg.Write(GB.MMU.Read(addr), "A")
	return opcode, 3, 16, true
}

func (GB *GB) LDHAC() (string, int, int, bool) { // reg[A] = M[$FF00+reg[C]]
	opcode := fmt.Sprintf("LD A,(0xFF00+C)")
	GB.CPU.Reg.Write(GB.MMU.Read(int(uint16(0xFF00+GB.CPU.Reg.Read("C")))), "A")
	return opcode, 1, 8, true
}

func (GB *GB) LDHLIA() (string, int, int, bool) { // M[HL] = reg[A], HL++
	opcode := fmt.Sprintf("LD (HL+),A")
	GB.SetHLVal(GB.CPU.Reg.Read("A"))
	GB.CPU.Reg.Write(int(uint16(GB.CPU.Reg.Read("HL")+1)), "HL")
	return opcode, 1, 8, true
}

func (GB *GB) LDHLDA() (string, int, int, bool) { // M[HL] = reg[A], HL--
	opcode := fmt.Sprintf("LD (HL-),A")
	GB.SetHLVal(GB.CPU.Reg.Read("A"))
	GB.CPU.Reg.Write(int(uint16(GB.CPU.Reg.Read("HL")-1)), "HL")
	return opcode, 1, 8, true
}

func (GB *GB) LDAHLI() (string, int, int, bool) { // reg[A] = M[HL], HL++
	opcode := fmt.Sprintf("LD A,(HL+)")
	GB.CPU.Reg.Write(GB.GetHLVal(), "A")
	GB.CPU.Reg.Write(int(uint16(GB.CPU.Reg.Read("HL")+1)), "HL")
	return opcode, 1, 8, true
}

func (GB *GB) LDAHLD() (string, int, int, bool) { // reg[A] = M[HL], HL--
	opcode := fmt.Sprintf("LD A,(HL-)")
	GB.CPU.Reg.Write(GB.GetHLVal(), "A")
	GB.CPU.Reg.Write(int(uint16(GB.CPU.Reg.Read("HL")-1)), "HL")
	return opcode, 1, 8, true
}

// OPCODES JUMP INSTRUCTIONS
func (GB *GB) JPN16(op int) (string, int, int, bool) { // reg[PC] = n16
	opcode := fmt.Sprintf("JP $%x ", op)
	GB.CPU.Reg.Write(op, "PC")
	return opcode, 0, 16, true
}

func (GB *GB) JPCCN16(cc string) (string, int, int, bool) { // reg[PC] = n16, if cc == true
	addr := GB.MMU.Read(GB.CPU.Reg.Read("PC")+1) + GB.MMU.Read(GB.CPU.Reg.Read("PC")+2)<<8
	opcode := fmt.Sprintf("JP %s,$%x ", cc, addr)
	if GB.CPU.CheckCC(cc) {
		GB.JPN16(addr)
		return opcode, 0, 16, true
	}
	return opcode, 3, 12, true

}

func (GB *GB) JPHL() (string, int, int, bool) { // reg[PC] = reg[HL]
	opcode := fmt.Sprintf("JP HL")
	GB.CPU.Reg.Write(GB.CPU.Reg.Read("HL"), "PC")
	return opcode, 0, 4, true
}

func (GB *GB) JRN16() (string, int, int, bool) { // reg[PC] = reg[PC] + n16
	opcode := fmt.Sprintf("JR $%x", int8(GB.MMU.Read(GB.CPU.Reg.Read("PC")+1)))
	addr := int(int8(GB.MMU.Read(GB.CPU.Reg.Read("PC") + 1)))
	GB.JPN16(GB.CPU.Reg.Read("PC") + addr + 2)
	return opcode, 0, 12, true
}

func (GB *GB) JRCCN16(cc string) (string, int, int, bool) { // reg[PC] = reg[PC] + n16, if cc == true
	opcode := fmt.Sprintf("JR %s,$%x", cc, uint8(GB.MMU.Read(GB.CPU.Reg.Read("PC")+1)))
	if GB.CPU.CheckCC(cc) {
		GB.JRN16()
		return opcode, 0, 12, true
	}
	return opcode, 2, 8, true
}

func (GB *GB) RET() (string, int, int, bool) { //
	opcode := fmt.Sprintf("RET")
	GB.POPR16("PC")
	return opcode, 0, 16, true
}

func (GB *GB) RETCC(cc string) (string, int, int, bool) { //
	opcode := fmt.Sprintf("RET %s", cc)
	if GB.CPU.CheckCC(cc) {
		GB.RET()
		return opcode, 0, 20, true
	}
	return opcode, 1, 8, true
}

func (GB *GB) RETI() (string, int, int, bool) { // Equivalent to:
	opcode := fmt.Sprintf("RETI")
	GB.CPU.IME = true // EI
	GB.RET()          // RET
	return opcode, 0, 16, true
}

func (GB *GB) CALLN16() (string, int, int, bool) { // Equivalent to following
	addr := GB.MMU.Read(GB.CPU.Reg.Read("PC")+1) + GB.MMU.Read(GB.CPU.Reg.Read("PC")+2)<<8
	opcode := fmt.Sprintf("CALL $%x", addr)
	GB.PUSHN16(GB.CPU.Reg.Read("PC") + 3) // PUSH M[reg[PC+3]]
	GB.JPN16(addr)                        // JP N16
	return opcode, 0, 24, true
}

func (GB *GB) CALLN8(op int) (string, int, int, bool) { // Equivalent to following
	GB.PUSHN16(GB.CPU.Reg.Read("PC")) // PUSH M[reg[PC]]
	GB.JPN16(op)                      // JP N16
	return "", 0, 0, true
}

func (GB *GB) CALLCCN16(cc string) (string, int, int, bool) { // CALLN16 if cc == true
	addr := GB.MMU.Read(GB.CPU.Reg.Read("PC")+1) + GB.MMU.Read(GB.CPU.Reg.Read("PC")+2)<<8
	opcode := fmt.Sprintf("CALL %s,$%x", cc, addr)
	if GB.CPU.CheckCC(cc) {
		GB.CALLN16()
		return opcode, 0, 24, true
	}
	return opcode, 3, 12, true
}

func (GB *GB) RST(vec string) (string, int, int, bool) {
	opcode := fmt.Sprintf("RST %s", vec)
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
	GB.PUSHN16(GB.CPU.Reg.Read("PC") + 1)
	GB.JPN16(vecAddr)
	return opcode, 0, 16, true
}

// OPCODES FOR STACK INSTRUCTIONS
func (GB *GB) ADDHLSP() (string, int, int, bool) { // reg[HL] = reg[HL] + reg[SP]
	return GB.ADDHLR16("SP")
}

func (GB *GB) ADDSPE8() (string, int, int, bool) { // reg[SP] = reg[SP] + e8
	op := int(int8(GB.MMU.Read(GB.CPU.Reg.Read("PC") + 1)))
	opcode := fmt.Sprintf("ADD SP,%x", op)
	val := int(int16(GB.CPU.Reg.Read("SP") + op))
	halfcarry := (GB.CPU.Reg.Read("SP") & 0xFF & 0xF) + (op & 0xF)
	carry := GB.CPU.Reg.Read("SP")&0xFF&0xFF + op&0xFF
	GB.CPU.SetFlagADD(halfcarry, carry, 1)
	GB.CPU.Reg.Write(val, "SP")
	return opcode, 2, 16, true
}

func (GB *GB) LDN16SP() (string, int, int, bool) {
	opcode := fmt.Sprintf("LD (%x), SP", GB.MMU.Read(GB.CPU.Reg.Read("PC")+0)+GB.MMU.Read(GB.CPU.Reg.Read("PC")+2)<<8)
	addr := GB.MMU.Read(GB.CPU.Reg.Read("PC")+1) + (GB.MMU.Read(GB.CPU.Reg.Read("PC")+2) << 8)
	GB.MMU.Write(GB.CPU.Reg.Read("SP")&0xFF, int(uint16(addr))) // M[n16] = reg[SP]
	GB.MMU.Write(GB.CPU.Reg.Read("SP")>>8, int(uint16(addr+1))) // M[n16+1] = reg[SP] >> 8
	return opcode, 3, 20, true
}

func (GB *GB) LDHLSPE8() (string, int, int, bool) { // reg[HL] = reg[SP] + e8
	op := int(int8(GB.MMU.Read(GB.CPU.Reg.Read("PC") + 1)))
	opcode := fmt.Sprintf("LD HL,SP+%x", op)
	val := int(int16(GB.CPU.Reg.Read("SP") + op))
	halfcarry := GB.CPU.Reg.Read("SP")&0xFF&0xF + op&0xF
	carry := GB.CPU.Reg.Read("SP")&0xFF&0xFF + op&0xFF
	GB.CPU.Reg.Write(val, "HL")
	GB.CPU.SetFlagADD(halfcarry, carry, 1)
	return opcode, 2, 12, true
}

func (GB *GB) LDSPHL() (string, int, int, bool) { // reg[SP] = reg[HL]
	opcode := fmt.Sprintf("LD SP,HL")
	GB.CPU.Reg.Write(GB.CPU.Reg.Read("HL"), "SP")
	return opcode, 1, 8, true
}

func (GB *GB) POPFA() (string, int, int, bool) { // POP reg[AF] from Stack
	opcode := fmt.Sprintf("POP AF")
	GB.CPU.Reg.Write(GB.MMU.Read(GB.CPU.Reg.Read("SP"))&0xF0, "F")
	GB.CPU.Reg.Write(int(int16(GB.CPU.Reg.Read("SP")+1)), "SP")
	GB.CPU.Reg.Write(GB.MMU.Read(GB.CPU.Reg.Read("SP")), "A")
	GB.CPU.Reg.Write(int(int16(GB.CPU.Reg.Read("SP")+1)), "SP")
	return opcode, 1, 12, true
}

func (GB *GB) POPR16(reg string) (string, int, int, bool) { // POP reg[] from Stack
	opcode := fmt.Sprintf("POP %s", reg)
	x := GB.MMU.Read(GB.CPU.Reg.Read("SP"))
	GB.CPU.Reg.Write(GB.CPU.Reg.Read("SP")+1, "SP")
	y := GB.MMU.Read(GB.CPU.Reg.Read("SP")) << 8
	GB.CPU.Reg.Write(GB.CPU.Reg.Read("SP")+1, "SP")
	GB.CPU.Reg.Write(int(int16(x+y)), reg)
	return opcode, 1, 12, true
}

func (GB *GB) PUSHFA() (string, int, int, bool) { // PUSH reg[AF] from Stack
	opcode := fmt.Sprintf("PUSH AF")
	GB.CPU.Reg.Write(int(int16(GB.CPU.Reg.Read("SP")-1)), "SP")
	GB.MMU.Write(GB.CPU.Reg.Read("A"), GB.CPU.Reg.Read("SP"))
	GB.CPU.Reg.Write(int(int16(GB.CPU.Reg.Read("SP")-1)), "SP")
	GB.MMU.Write((GB.CPU.GetFlag("Z")<<7)|(GB.CPU.GetFlag("N")<<6)|(GB.CPU.GetFlag("H")<<5)|(GB.CPU.GetFlag("C")<<4), GB.CPU.Reg.Read("SP"))
	return opcode, 1, 16, true
}

func (GB *GB) PUSHR16(reg string) (string, int, int, bool) { // PUSH reg[] from Stack
	opcode := fmt.Sprintf("PUSH %s", reg)
	GB.CPU.Reg.Write(int(int16(GB.CPU.Reg.Read("SP")-1)), "SP")
	GB.MMU.Write((GB.CPU.Reg.Read(reg)&0xFF00)>>8, GB.CPU.Reg.Read("SP"))
	GB.CPU.Reg.Write(int(int16(GB.CPU.Reg.Read("SP")-1)), "SP")
	GB.MMU.Write(GB.CPU.Reg.Read(reg)&0xFF, GB.CPU.Reg.Read("SP"))
	return opcode, 1, 16, true
}

func (GB *GB) PUSHN16(op int) { // PUSH N16 from Stack (NOT AN OPCODE)
	GB.CPU.Reg.Write(int(int16(GB.CPU.Reg.Read("SP")-1)), "SP")
	GB.MMU.Write((op&0xFF00)>>8, GB.CPU.Reg.Read("SP"))
	GB.CPU.Reg.Write(int(int16(GB.CPU.Reg.Read("SP")-1)), "SP")
	GB.MMU.Write(op&0xFF, GB.CPU.Reg.Read("SP"))
}

// OPCODES FOR MISC INSTRUCTIONS
func (GB *GB) EI() (string, int, int, bool) {
	opcode := fmt.Sprintf("EI")
	GB.CPU.IME = true
	return opcode, 1, 4, true
}

func (GB *GB) DI() (string, int, int, bool) {
	opcode := fmt.Sprintf("DI")
	GB.CPU.IME = false
	return opcode, 1, 4, true
}

func (GB *GB) CPL() (string, int, int, bool) { // Reg[A] = ~Reg[A]
	opcode := fmt.Sprintf("CPL")
	GB.CPU.Reg.Write(GB.CPU.Reg.Read("A"), "A")
	GB.CPU.SetFlag("N")
	GB.CPU.SetFlag("H")
	return opcode, 1, 4, true
}

func (GB *GB) SCF() (string, int, int, bool) { // Flag[C] = true
	opcode := fmt.Sprintf("SCF")
	GB.CPU.ResetFlag("N")
	GB.CPU.ResetFlag("H")
	GB.CPU.SetFlag("C")
	return opcode, 1, 4, true
}

func (GB *GB) CCF() (string, int, int, bool) { // Flag[C] = ~Flag[C]
	opcode := fmt.Sprintf("CCF")
	if GB.CPU.GetFlag("C") == 1 {
		GB.CPU.ResetFlag("C")
	} else {
		GB.CPU.SetFlag("C")
	}
	GB.CPU.ResetFlag("N")
	GB.CPU.ResetFlag("H")
	return opcode, 1, 4, true
}

func (GB *GB) NOP() (string, int, int, bool) {
	opcode := fmt.Sprintf("NOP")
	return opcode, 1, 4, true
}

func (GB *GB) HALT() (string, int, int, bool) {
	opcode := fmt.Sprintf("HALT")
	return opcode, 1, 4, false
}
