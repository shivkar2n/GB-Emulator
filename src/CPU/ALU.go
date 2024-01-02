// Opcodes for 8-bit ALU instructions

package CPU

import "fmt"

func (cpu *CPU) ADCAR8(reg string) (string, int, int, bool) { // reg[A] = reg[A] + C + reg[]
	opcode := fmt.Sprintf("ADC A,%s", reg)
	regVal := cpu.Reg.Read(reg)
	val := int(int8(cpu.Reg.Read("A") + regVal + cpu.GetFlag("C")))
	halfcarry := cpu.Reg.Read("A")&0xF + regVal&0xF + cpu.GetFlag("C")
	carry := cpu.Reg.Read("A")&0xFF + regVal&0xFF + cpu.GetFlag("C")
	cpu.Reg.Write(val, "A")
	cpu.SetFlagADD(halfcarry, carry, val)
	return opcode, 1, 4, true
}

func (cpu *CPU) ADCAHL() (string, int, int, bool) { // reg[A] = reg[A] + C + M[reg[HL]]
	opcode := fmt.Sprintf("ADC A,(HL)")
	val := int(int8(cpu.Reg.Read("A") + cpu.GetHLVal() + cpu.GetFlag("C")))
	halfcarry := cpu.Reg.Read("A")&0xF + cpu.GetHLVal()&0xF + cpu.GetFlag("C")
	carry := cpu.Reg.Read("A")&0xFF + cpu.GetHLVal()&0xFF + cpu.GetFlag("C")
	cpu.Reg.Write(val, "A")
	cpu.SetFlagADD(halfcarry, carry, val)
	return opcode, 1, 8, true
}

func (cpu *CPU) ADCAN8() (string, int, int, bool) { // reg[A] = reg[A] + n8 + C
	opcode := fmt.Sprintf("ADC A,$%x", cpu.Mem.Read(cpu.Reg.Read("PC")+1))
	addr := cpu.Mem.Read(cpu.Reg.Read("PC") + 1)
	val := int(int8(cpu.Reg.Read("A") + addr + cpu.GetFlag("C")))
	halfcarry := cpu.Reg.Read("A")&0xF + addr&0xF + cpu.GetFlag("C")
	carry := cpu.Reg.Read("A")&0xFF + addr&0xFF + cpu.GetFlag("C")
	cpu.Reg.Write(val, "A")
	cpu.SetFlagADD(halfcarry, carry, val)
	return opcode, 2, 8, true
}

func (cpu *CPU) ADDAR8(reg string) (string, int, int, bool) { // reg[A] = reg[A] + reg[]
	opcode := fmt.Sprintf("ADD A,%s", reg)
	regVal := cpu.Reg.Read(reg)
	regAVal := cpu.Reg.Read("A")
	val := int(int8(regAVal + regVal))
	halfcarry := regAVal&0xF + regVal&0xF
	carry := regAVal&0xFF + regVal&0xFF
	cpu.Reg.Write(val, "A")
	cpu.SetFlagADD(halfcarry, carry, val)
	return opcode, 1, 4, true
}

func (cpu *CPU) ADDAHL() (string, int, int, bool) { // reg[A] = reg[A] + M[reg[HL]]
	opcode := fmt.Sprintf("ADD A,(HL)")
	val := int(int8(cpu.Reg.Read("A") + cpu.GetHLVal()))
	halfcarry := cpu.Reg.Read("A")&0xF + cpu.GetHLVal()&0xF
	carry := cpu.Reg.Read("A")&0xFF + cpu.GetHLVal()&0xFF
	cpu.Reg.Write(val, "A")
	cpu.SetFlagADD(halfcarry, carry, val)
	return opcode, 1, 8, true
}

func (cpu *CPU) ADDAN8() (string, int, int, bool) { // reg[A] = reg[A] + n8
	op := cpu.Mem.Read(cpu.Reg.Read("PC") + 1)
	opcode := fmt.Sprintf("ADD A,$%x", uint8(op))
	val := int(int8(cpu.Reg.Read("A") + op))
	halfcarry := cpu.Reg.Read("A")&0xF + op&0xF
	carry := cpu.Reg.Read("A")&0xFF + op&0xFF
	cpu.Reg.Write(val, "A")
	cpu.SetFlagADD(halfcarry, carry, val)
	return opcode, 2, 8, true
}

func (cpu *CPU) ANDAR8(reg string) (string, int, int, bool) { // reg[A] = reg[A] & reg[]
	opcode := fmt.Sprintf("AND A,%s", reg)
	regVal := cpu.Reg.Read(reg)
	val := cpu.Reg.Read("A") & regVal
	cpu.Reg.Write(val, "A")
	cpu.SetFlagAND(val)
	return opcode, 1, 4, true
}

func (cpu *CPU) ANDAHL() (string, int, int, bool) { // reg[A] = reg[A] & M[HL]
	opcode := fmt.Sprintf("AND A,(HL)")
	val := cpu.Reg.Read("A") & cpu.GetHLVal()
	cpu.Reg.Write(val, "A")
	cpu.SetFlagAND(val)
	return opcode, 1, 8, true
}

func (cpu *CPU) ANDAN8() (string, int, int, bool) { // reg[A] = reg[A] & n8
	op := cpu.Mem.Read(cpu.Reg.Read("PC") + 1)
	opcode := fmt.Sprintf("AND A,$%x", uint8(op))
	val := cpu.Reg.Read("A") & op
	cpu.Reg.Write(val, "A")
	cpu.SetFlagAND(val)
	return opcode, 2, 8, true
}

func (cpu *CPU) CPAR8(reg string) (string, int, int, bool) { // res = reg[A] - reg[]
	opcode := fmt.Sprintf("CP A,%s", reg)
	regVal := cpu.Reg.Read(reg)
	halfBorrow := cpu.Reg.Read("A")&0x0F - regVal&0x0F
	borrow := cpu.Reg.Read("A")&0xFF - regVal&0xFF
	val := int(int8(cpu.Reg.Read("A") - regVal))
	cpu.SetFlagCP(halfBorrow, borrow, val)
	return opcode, 1, 4, true
}

func (cpu *CPU) CPAHL() (string, int, int, bool) { // res = reg[A] - M[HL]
	opcode := fmt.Sprintf("CP A,(HL)")
	halfBorrow := cpu.Reg.Read("A")&0x0F - cpu.GetHLVal()&0x0F
	borrow := cpu.Reg.Read("A")&0xFF - cpu.GetHLVal()&0xFF
	val := int(int8(cpu.Reg.Read("A") - cpu.GetHLVal()))
	cpu.SetFlagCP(halfBorrow, borrow, val)
	return opcode, 1, 8, true
}

func (cpu *CPU) CPAN8() (string, int, int, bool) { // res = reg[A] - n8
	op := cpu.Mem.Read(cpu.Reg.Read("PC") + 1)
	opcode := fmt.Sprintf("CP $%x", uint8(op))
	halfBorrow := cpu.Reg.Read("A")&0x0F - op&0x0F
	borrow := cpu.Reg.Read("A")&0xFF - op&0xFF
	val := int(int8(cpu.Reg.Read("A") - op))
	cpu.SetFlagCP(halfBorrow, borrow, val)
	return opcode, 2, 8, true
}

func (cpu *CPU) DAA() (string, int, int, bool) { // res = reg[A] - n8
	opcode := fmt.Sprintf("DAA")
	if cpu.GetFlag("N") == 0 {
		if cpu.GetFlag("C") == 1 || cpu.Reg.Read("A") > 0x99 {
			cpu.Reg.Write(int(int8(cpu.Reg.Read("A")+0x60)), "A")
			cpu.SetFlag("C")
		}
		if cpu.GetFlag("H") == 1 || cpu.Reg.Read("A")&0x0F > 0x09 {
			cpu.Reg.Write(int(int8(cpu.Reg.Read("A")+0x6)), "A")
		}
	} else {
		if cpu.GetFlag("C") == 1 {
			cpu.Reg.Write(int(int8(cpu.Reg.Read("A")-0x60)), "A")
		}
		if cpu.GetFlag("H") == 1 {
			cpu.Reg.Write(int(int8(cpu.Reg.Read("A")-0x6)), "A")
		}
	}
	if cpu.Reg.Read("A") == 0 {
		cpu.SetFlag("Z")
	} else {
		cpu.ResetFlag("Z")
	}
	cpu.ResetFlag("H")
	return opcode, 1, 4, true
}

func (cpu *CPU) DECR8(reg string) (string, int, int, bool) { // reg[] = reg[] - 1
	opcode := fmt.Sprintf("DEC %s", reg)
	val := int(int8(cpu.Reg.Read(reg)) - 1)
	halfBorrow := cpu.Reg.Read(reg)&0x0F - 1
	cpu.SetFlagDEC(halfBorrow, val)
	cpu.Reg.Write(val, reg)
	return opcode, 1, 4, true
}

func (cpu *CPU) DECHL() (string, int, int, bool) { // M[HL] = M[HL] - 1
	opcode := fmt.Sprintf("DEC (HL)")
	val := int(int8(cpu.GetHLVal() - 1))
	halfBorrow := cpu.GetHLVal()&0x0F - 1
	cpu.SetHLVal(val)
	cpu.SetFlagDEC(halfBorrow, val)
	return opcode, 1, 12, true
}

func (cpu *CPU) INCR8(reg string) (string, int, int, bool) { // reg[] = reg[] + 1
	opcode := fmt.Sprintf("INC %s", reg)
	val := int(int8(cpu.Reg.Read(reg) + 1))
	halfcarry := cpu.Reg.Read(reg)&0x0F + 1
	cpu.Reg.Write(val, reg)
	cpu.SetFlagINC(halfcarry, val)
	return opcode, 1, 4, true
}

func (cpu *CPU) INCRHL() (string, int, int, bool) { // M[HL] = M[HL] + 1
	opcode := fmt.Sprintf("INC (HL)")
	val := int(int8(cpu.GetHLVal() + 1))
	halfcarry := cpu.GetHLVal()&0x0F + 1
	cpu.SetHLVal(val)
	cpu.SetFlagINC(halfcarry, val)
	return opcode, 1, 12, true
}

func (cpu *CPU) ORAR8(reg string) (string, int, int, bool) { // reg[A] = reg[A] | reg[]
	opcode := fmt.Sprintf("OR A,%s", reg)
	val := cpu.Reg.Read("A") | cpu.Reg.Read(reg)
	cpu.Reg.Write(val, "A")
	cpu.SetFlagOR(val)
	return opcode, 1, 4, true
}

func (cpu *CPU) ORAHL() (string, int, int, bool) { // reg[A] = M[HL] | reg[A]
	opcode := fmt.Sprintf("OR A,(HL)")
	val := cpu.Reg.Read("A") | cpu.GetHLVal()
	cpu.Reg.Write(val, "A")
	cpu.SetFlagOR(val)
	return opcode, 1, 8, true
}

func (cpu *CPU) ORAN8() (string, int, int, bool) { // reg[A] = reg[A] | n8
	op := cpu.Mem.Read(cpu.Reg.Read("PC") + 1)
	opcode := fmt.Sprintf("OR A,$%x", int8(op))
	val := cpu.Reg.Read("A") | op
	cpu.Reg.Write(val, "A")
	cpu.SetFlagOR(val)
	return opcode, 2, 8, true
}

func (cpu *CPU) SUBAR8(reg string) (string, int, int, bool) { // reg[A] = reg[A] - reg[]
	opcode := fmt.Sprintf("SUB A,%s", reg)
	regVal := cpu.Reg.Read(reg)
	val := cpu.Reg.Read("A") - regVal
	halfborrow := cpu.Reg.Read("A")&0xF - regVal&0xF
	borrow := cpu.Reg.Read("A") - regVal
	cpu.Reg.Write(val, "A")
	cpu.SetFlagSUB(halfborrow, borrow, val)
	return opcode, 1, 4, true
}

func (cpu *CPU) SUBAHL() (string, int, int, bool) { // reg[A] = reg[A] - M[HL]
	opcode := fmt.Sprintf("SUB A,(HL)")
	val := cpu.Reg.Read("A") - cpu.GetHLVal()
	halfborrow := cpu.Reg.Read("A")&0xF - cpu.GetHLVal()&0xF
	borrow := cpu.Reg.Read("A") - cpu.GetHLVal()
	cpu.Reg.Write(val, "A")
	cpu.SetFlagSUB(halfborrow, borrow, val)
	return opcode, 1, 8, true
}

func (cpu *CPU) SUBAN8() (string, int, int, bool) { // reg[A] = reg[A] - n8
	opcode := fmt.Sprintf("SUB A,$%x", cpu.Mem.Read(cpu.Reg.Read("PC")+1))
	addr := cpu.Mem.Read(cpu.Reg.Read("PC") + 1)
	val := cpu.Reg.Read("A") - addr
	halfborrow := cpu.Reg.Read("A")&0xF - addr&0xF
	borrow := cpu.Reg.Read("A") - addr
	cpu.Reg.Write(val, "A")
	cpu.SetFlagSUB(halfborrow, borrow, val)
	return opcode, 2, 8, true
}

func (cpu *CPU) SBCAR8(reg string) (string, int, int, bool) { // reg[A] = reg[A] - reg[] - C
	opcode := fmt.Sprintf("SBC A,B")
	regVal := cpu.Reg.Read(reg)
	val := int(int8(cpu.Reg.Read("A") - regVal - cpu.GetFlag("C")))
	halfborrow := cpu.Reg.Read("A")&0xF - regVal&0xF - cpu.GetFlag("C")
	borrow := cpu.Reg.Read("A")&0xFF - regVal&0xFF - cpu.GetFlag("C")
	cpu.Reg.Write(val, "A")
	cpu.SetFlagSUB(halfborrow, borrow, val)
	return opcode, 1, 4, true
}

func (cpu *CPU) SBCAHL() (string, int, int, bool) { // reg[A] = reg[A] - M[HL] - C
	opcode := fmt.Sprintf("SBC A,(HL)")
	val := int(int8(cpu.Reg.Read("A") - cpu.GetHLVal() - cpu.GetFlag("C")))
	halfborrow := cpu.Reg.Read("A")&0xF - cpu.GetHLVal()&0xF - cpu.GetFlag("C")
	borrow := cpu.Reg.Read("A") - cpu.GetHLVal() - cpu.GetFlag("C")
	cpu.Reg.Write(val, "A")
	cpu.SetFlagSUB(halfborrow, borrow, val)
	return opcode, 1, 8, true
}

func (cpu *CPU) SBCAN8() (string, int, int, bool) { // reg[A] = reg[A] - n8 - C
	op := cpu.Mem.Read(cpu.Reg.Read("PC") + 1)
	opcode := fmt.Sprintf("SBC A,$%x", uint8(op))
	val := int(int8(cpu.Reg.Read("A") - op - cpu.GetFlag("C")))
	halfborrow := cpu.Reg.Read("A")&0xF - op&0xF - cpu.GetFlag("C")
	borrow := cpu.Reg.Read("A")&0xFF - op&0xFF - cpu.GetFlag("C")
	cpu.Reg.Write(val, "A")
	cpu.SetFlagSUB(halfborrow, borrow, val)
	return opcode, 2, 8, true
}

func (cpu *CPU) XORAR8(reg string) (string, int, int, bool) { // reg[A] = reg[A] ^ reg[]
	opcode := fmt.Sprintf("XOR A,%s", reg)
	val := cpu.Reg.Read("A") ^ cpu.Reg.Read(reg)
	cpu.Reg.Write(val, "A")
	cpu.SetFlagOR(val)
	return opcode, 1, 4, true
}

func (cpu *CPU) XORAHL() (string, int, int, bool) { // reg[A] = reg[A] ^ M[HL]
	opcode := fmt.Sprintf("XOR A,(HL)")
	val := cpu.Reg.Read("A") ^ cpu.GetHLVal()
	cpu.Reg.Write(val, "A")
	cpu.SetFlagOR(val)
	return opcode, 1, 8, true
}

func (cpu *CPU) XORAN8() (string, int, int, bool) { // reg[A] = reg[A] ^ op
	opcode := fmt.Sprintf("XOR $%x", uint8(cpu.Mem.Read(cpu.Reg.Read("PC")+1)))
	op := cpu.Mem.Read(cpu.Reg.Read("PC") + 1)
	val := cpu.Reg.Read("A") ^ op
	cpu.Reg.Write(val, "A")
	cpu.SetFlagOR(val)
	return opcode, 2, 8, true
}

// Opcodes for 16-bit ALU instructions
// half carry -> carry from 11th bit
// carry -> carry from 15th bit

func (cpu *CPU) ADDHLR16(reg string) (string, int, int, bool) { // reg[HL] = reg[r16] + reg[HL]
	opcode := fmt.Sprintf("ADD HL, %s", reg)
	val := int(cpu.Reg.Read("HL") + cpu.Reg.Read(reg))
	var tpcarry int
	if (cpu.Reg.Read("HL")&0xFF+cpu.Reg.Read(reg)&0xFF)&0x100 == 0x100 {
		tpcarry = 1
	} else {
		tpcarry = 0
	}
	halfcarry := ((cpu.Reg.Read("HL")&0x0F00)>>8 + (cpu.Reg.Read(reg)&0x0F00)>>8 + tpcarry) >> 4
	carry := val >> 16
	cpu.Reg.Write(int(uint16(val)), "HL")
	cpu.SetFlagADD16(halfcarry, carry)
	return opcode, 1, 8, true
}

func (cpu *CPU) DECR16(reg string) (string, int, int, bool) { // reg[r16] = reg[r16] - 1
	opcode := fmt.Sprintf("DEC %s", reg)
	cpu.Reg.Write(int(int16(cpu.Reg.Read(reg)-1)), reg)
	return opcode, 1, 8, true
}

func (cpu *CPU) INCR16(reg string) (string, int, int, bool) { // reg[r16] = reg[r16] + 1
	opcode := fmt.Sprintf("INC %s", reg)
	cpu.Reg.Write(int(int16(cpu.Reg.Read(reg)+1)), reg)
	return opcode, 1, 8, true
}
