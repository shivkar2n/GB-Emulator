package CPU

import "fmt"

func (cpu *CPU) RLR8(reg string) (string, int, int, bool) { // Rotate bits reg[] left through Flag C
	opcode := fmt.Sprintf("RL %s", reg)
	carry := cpu.Reg.Read(reg) >> 7
	val := int(int8(cpu.Reg.Read(reg)<<1) | int8(cpu.GetFlag("C")))
	cpu.Reg.Write(val, reg)
	cpu.SetFlagRL(val, carry)
	return opcode, 2, 8, true
}

func (cpu *CPU) RLHL() (string, int, int, bool) { // Rotate bits M[reg[HL]] left through Flag C
	opcode := fmt.Sprintf("RL (HL)")
	carry := cpu.GetHLVal() >> 7
	val := int(int8((cpu.GetHLVal() << 1) | cpu.GetFlag("C")))
	cpu.SetHLVal(val)
	cpu.SetFlagRL(val, carry)
	return opcode, 2, 16, true
}

func (cpu *CPU) RLA() (string, int, int, bool) { // Rotate bits reg[A] left through Flag C
	opcode := fmt.Sprintf("RLA")
	carry := cpu.Reg.Read("A") >> 7
	val := int(int8((cpu.Reg.Read("A") << 1) | cpu.GetFlag("C")))
	cpu.Reg.Write(val, "A")
	cpu.SetFlagRL(1, carry)
	return opcode, 1, 4, true
}

func (cpu *CPU) RLCR8(reg string) (string, int, int, bool) { // Rotate bits reg[A] left through Flag c
	opcode := fmt.Sprintf("RLC %s", reg)
	carry := cpu.Reg.Read(reg) >> 7
	val := int(int8((cpu.Reg.Read(reg) << 1) | carry))
	cpu.Reg.Write(val, reg)
	cpu.SetFlagRL(val, carry)
	return opcode, 2, 8, true
}

func (cpu *CPU) RLCHL() (string, int, int, bool) { // Rotate bits M[reg[HL]] left through flag C
	opcode := fmt.Sprintf("RLC (HL)")
	carry := cpu.GetHLVal() >> 7
	val := int(int8((cpu.GetHLVal() << 1) | carry))
	cpu.SetHLVal(val)
	cpu.SetFlagRL(val, carry)
	return opcode, 2, 16, true
}

func (cpu *CPU) RLCA() (string, int, int, bool) { // Rotate bits reg[A] left
	opcode := fmt.Sprintf("RLCA")
	carry := cpu.Reg.Read("A") >> 7
	val := int(int8((cpu.Reg.Read("A") << 1) | carry))
	cpu.Reg.Write(val, "A")
	cpu.SetFlagRL(1, carry)
	return opcode, 1, 4, true
}

func (cpu *CPU) RRR8(reg string) (string, int, int, bool) { // Rotate bits reg[] right through flag C
	opcode := fmt.Sprintf("RR %s", reg)
	carry := cpu.Reg.Read(reg) & 0x01
	val := int(int8((cpu.Reg.Read(reg) >> 1) | (cpu.GetFlag("C") << 7)))
	cpu.Reg.Write(val, reg)
	cpu.SetFlagRL(val, carry)
	return opcode, 2, 8, true
}

func (cpu *CPU) RRHL() (string, int, int, bool) { // Rotate bits M[reg[HL]] right through flag C
	opcode := fmt.Sprintf("RR (HL)")
	carry := cpu.GetHLVal() & 0x01
	val := int(int8((cpu.GetHLVal() >> 1) | (cpu.GetFlag("C") << 7)))
	cpu.SetHLVal(val)
	cpu.SetFlagRL(val, carry)
	return opcode, 2, 16, true
}

func (cpu *CPU) RRA() (string, int, int, bool) { // Rotate bits reg[A] right through flag C
	opcode := fmt.Sprintf("RRA")
	carry := cpu.Reg.Read("A") & 0x01
	val := (cpu.Reg.Read("A") >> 1) | (cpu.GetFlag("C") << 7)
	cpu.Reg.Write(val, "A")
	cpu.SetFlagRL(1, carry)
	return opcode, 1, 4, true
}

func (cpu *CPU) RRCR8(reg string) (string, int, int, bool) { // Rotate reg[] right
	opcode := fmt.Sprintf("RRC %s", reg)
	carry := cpu.Reg.Read(reg) & 0x01
	val := (cpu.Reg.Read(reg) >> 1) | (carry << 7)
	cpu.Reg.Write(val, reg)
	cpu.SetFlagRL(val, carry)
	return opcode, 2, 8, true
}

func (cpu *CPU) RRCHL() (string, int, int, bool) { // Rotate M[reg[HL]] right
	opcode := fmt.Sprintf("RRC (HL)")
	carry := cpu.GetHLVal() & 0x01
	val := (cpu.GetHLVal() >> 1) | (carry << 7)
	cpu.SetHLVal(val)
	cpu.SetFlagRL(val, carry)
	return opcode, 2, 16, true
}

func (cpu *CPU) RRCA() (string, int, int, bool) { // Rotate reg[A] right
	opcode := fmt.Sprintf("RRCA ")
	carry := cpu.Reg.Read("A") & 0x01
	val := (cpu.Reg.Read("A") >> 1) | (carry << 7)
	cpu.Reg.Write(val, "A")
	cpu.SetFlagRL(1, carry)
	return opcode, 1, 4, true
}

func (cpu *CPU) SLAR8(reg string) (string, int, int, bool) { // reg[] = reg[] << 1
	opcode := fmt.Sprintf("SLA B")
	carry := cpu.Reg.Read(reg) >> 7
	val := int(int8(cpu.Reg.Read(reg) << 1))
	cpu.Reg.Write(val, reg)
	cpu.SetFlagRL(val, carry)
	return opcode, 2, 8, true
}

func (cpu *CPU) SLAHL() (string, int, int, bool) { // M[reg[HL]] = M[reg[HL]] << 1
	opcode := fmt.Sprintf("SLA (HL)")
	carry := cpu.GetHLVal() >> 7
	val := int(int8(cpu.GetHLVal() << 1))
	cpu.SetHLVal(val)
	cpu.SetFlagRL(val, carry)
	return opcode, 2, 16, true
}

func (cpu *CPU) SRAR8(reg string) (string, int, int, bool) { // reg[] = reg[] >> 1, bit 7 unchanged
	opcode := fmt.Sprintf("SRA %s", reg)
	carry := cpu.Reg.Read(reg) & 0x01
	val := (cpu.Reg.Read(reg) >> 1) | cpu.Reg.Read(reg)&0x80
	cpu.Reg.Write(val, reg)
	cpu.SetFlagRL(val, carry)
	return opcode, 2, 8, true
}

func (cpu *CPU) SRAHL() (string, int, int, bool) { // M[reg[HL]] = M[reg[HL]] >> 1, bit 7 unchanged
	opcode := fmt.Sprintf("SRA (HL)")
	carry := cpu.GetHLVal() & 0x01
	val := (cpu.GetHLVal() >> 1) | cpu.GetHLVal()&0x80
	cpu.SetHLVal(val)
	cpu.SetFlagRL(val, carry)
	return opcode, 2, 16, true
}

func (cpu *CPU) SRLR8(reg string) (string, int, int, bool) { // reg[] = reg[] >> 1
	opcode := fmt.Sprintf("SRL %s", reg)
	carry := cpu.Reg.Read(reg) & 0x01
	val := cpu.Reg.Read(reg) >> 1
	cpu.Reg.Write(val, reg)
	cpu.SetFlagRL(val, carry)
	return opcode, 2, 8, true
}

func (cpu *CPU) SRLHL() (string, int, int, bool) { // M[reg[HL]] = M[reg[HL]] >> 1
	opcode := fmt.Sprintf("SRL (HL)")
	carry := cpu.GetHLVal() & 0x01
	val := cpu.GetHLVal() >> 1
	cpu.SetHLVal(val)
	cpu.SetFlagRL(val, carry)
	return opcode, 2, 16, true
}
