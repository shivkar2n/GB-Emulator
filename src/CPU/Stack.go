package CPU

import "fmt"

func (cpu *CPU) ADDHLSP() (string, int, int, bool) { // reg[HL] = reg[HL] + reg[SP]
	return cpu.ADDHLR16("SP")
}

func (cpu *CPU) ADDSPE8() (string, int, int, bool) { // reg[SP] = reg[SP] + e8
	op := int(int8(cpu.Mem.Read(cpu.Reg.Read("PC") + 1)))
	opcode := fmt.Sprintf("ADD SP,%x", op)
	val := int(int16(cpu.Reg.Read("SP") + op))
	halfcarry := (cpu.Reg.Read("SP") & 0xFF & 0xF) + (op & 0xF)
	carry := cpu.Reg.Read("SP")&0xFF&0xFF + op&0xFF
	cpu.SetFlagADD(halfcarry, carry, 1)
	cpu.Reg.Write(val, "SP")
	return opcode, 2, 16, true
}

func (cpu *CPU) LDN16SP() (string, int, int, bool) {
	opcode := fmt.Sprintf("LD (%x), SP", cpu.Mem.Read(cpu.Reg.Read("PC")+0)+cpu.Mem.Read(cpu.Reg.Read("PC")+2)<<8)
	addr := cpu.Mem.Read(cpu.Reg.Read("PC")+1) + (cpu.Mem.Read(cpu.Reg.Read("PC")+2) << 8)
	cpu.Mem.Write(cpu.Reg.Read("SP")&0xFF, int(uint16(addr))) // M[n16] = reg[SP]
	cpu.Mem.Write(cpu.Reg.Read("SP")>>8, int(uint16(addr+1))) // M[n16+1] = reg[SP] >> 8
	return opcode, 3, 20, true
}

func (cpu *CPU) LDHLSPE8() (string, int, int, bool) { // reg[HL] = reg[SP] + e8
	op := int(int8(cpu.Mem.Read(cpu.Reg.Read("PC") + 1)))
	opcode := fmt.Sprintf("LD HL,SP+%x", op)
	val := int(int16(cpu.Reg.Read("SP") + op))
	halfcarry := cpu.Reg.Read("SP")&0xFF&0xF + op&0xF
	carry := cpu.Reg.Read("SP")&0xFF&0xFF + op&0xFF
	cpu.Reg.Write(val, "HL")
	cpu.SetFlagADD(halfcarry, carry, 1)
	return opcode, 2, 12, true
}

func (cpu *CPU) LDSPHL() (string, int, int, bool) { // reg[SP] = reg[HL]
	opcode := fmt.Sprintf("LD SP,HL")
	cpu.Reg.Write(cpu.Reg.Read("HL"), "SP")
	return opcode, 1, 8, true
}

func (cpu *CPU) POPFA() (string, int, int, bool) { // POP reg[AF] from Stack
	opcode := fmt.Sprintf("POP AF")
	cpu.Reg.Write(cpu.Mem.Read(cpu.Reg.Read("SP"))&0xF0, "F")
	cpu.Reg.Write(int(int16(cpu.Reg.Read("SP")+1)), "SP")
	cpu.Reg.Write(cpu.Mem.Read(cpu.Reg.Read("SP")), "A")
	cpu.Reg.Write(int(int16(cpu.Reg.Read("SP")+1)), "SP")
	return opcode, 1, 12, true
}

func (cpu *CPU) POPR16(reg string) (string, int, int, bool) { // POP reg[] from Stack
	opcode := fmt.Sprintf("POP %s", reg)
	x := cpu.Mem.Read(cpu.Reg.Read("SP"))
	cpu.Reg.Write(cpu.Reg.Read("SP")+1, "SP")
	y := cpu.Mem.Read(cpu.Reg.Read("SP")) << 8
	cpu.Reg.Write(cpu.Reg.Read("SP")+1, "SP")
	cpu.Reg.Write(int(int16(x+y)), reg)
	return opcode, 1, 12, true
}

func (cpu *CPU) PUSHFA() (string, int, int, bool) { // PUSH reg[AF] from Stack
	opcode := fmt.Sprintf("PUSH AF")
	cpu.Reg.Write(int(int16(cpu.Reg.Read("SP")-1)), "SP")
	cpu.Mem.Write(cpu.Reg.Read("A"), cpu.Reg.Read("SP"))
	cpu.Reg.Write(int(int16(cpu.Reg.Read("SP")-1)), "SP")
	cpu.Mem.Write((cpu.GetFlag("Z")<<7)|(cpu.GetFlag("N")<<6)|(cpu.GetFlag("H")<<5)|(cpu.GetFlag("C")<<4), cpu.Reg.Read("SP"))
	return opcode, 1, 16, true
}

func (cpu *CPU) PUSHR16(reg string) (string, int, int, bool) { // PUSH reg[] from Stack
	opcode := fmt.Sprintf("PUSH %s", reg)
	cpu.Reg.Write(int(int16(cpu.Reg.Read("SP")-1)), "SP")
	cpu.Mem.Write((cpu.Reg.Read(reg)&0xFF00)>>8, cpu.Reg.Read("SP"))
	cpu.Reg.Write(int(int16(cpu.Reg.Read("SP")-1)), "SP")
	cpu.Mem.Write(cpu.Reg.Read(reg)&0xFF, cpu.Reg.Read("SP"))
	return opcode, 1, 16, true
}

func (cpu *CPU) PUSHN16(op int) { // PUSH N16 from Stack (NOT AN OPCODE)
	cpu.Reg.Write(int(int16(cpu.Reg.Read("SP")-1)), "SP")
	cpu.Mem.Write((op&0xFF00)>>8, cpu.Reg.Read("SP"))
	cpu.Reg.Write(int(int16(cpu.Reg.Read("SP")-1)), "SP")
	cpu.Mem.Write(op&0xFF, cpu.Reg.Read("SP"))
}
