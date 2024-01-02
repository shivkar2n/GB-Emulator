// Opcodes for LOAD instructions
// n8 - 8 bit int
// u8 - 8 bit int unsigned
// n16 - 16 bit const
// e8 - 8 bit offset
// u3 - 3 bit unsigned

package CPU

import "fmt"

func (cpu *CPU) LDR8R8(reg1 string, reg2 string) (string, int, int, bool) { // reg[reg1] = reg[reg2]
	opcode := fmt.Sprintf("LD %s,%s", reg1, reg2)
	cpu.Reg.Write(cpu.Reg.Read(reg2), reg1)
	return opcode, 1, 4, true
}

func (cpu *CPU) LDR8N8(reg string) (string, int, int, bool) { // reg[] = n8
	opcode := fmt.Sprintf("LD %s,$%x", reg, uint8(cpu.Mem.Read(cpu.Reg.Read("PC")+1)))
	addr := cpu.Mem.Read(cpu.Reg.Read("PC") + 1)
	cpu.Reg.Write(addr, reg)
	return opcode, 2, 8, true
}

func (cpu *CPU) LDR16N16(reg string) (string, int, int, bool) { // reg16[reg] = n16
	opcode := fmt.Sprintf("LD %s,$%x ", reg, uint16(cpu.Mem.Read(cpu.Reg.Read("PC")+1)+cpu.Mem.Read(cpu.Reg.Read("PC")+2)<<8))
	addr := int(uint16(cpu.Mem.Read(cpu.Reg.Read("PC")+1) + (cpu.Mem.Read(cpu.Reg.Read("PC")+2) << 8)))
	cpu.Reg.Write(addr, reg)
	return opcode, 3, 12, true
}

func (cpu *CPU) LDHLR8(reg string) (string, int, int, bool) { // M[HL] = reg[]
	opcode := fmt.Sprintf("LD (HL),%s", reg)
	cpu.SetHLVal(cpu.Reg.Read(reg))
	return opcode, 1, 8, true
}

func (cpu *CPU) LDHLN8() (string, int, int, bool) { // M[HL] = n8
	opcode := fmt.Sprintf("LD (HL),$%x", uint8(cpu.Mem.Read(cpu.Reg.Read("PC")+1)))
	op := cpu.Mem.Read(cpu.Reg.Read("PC") + 1)
	cpu.SetHLVal(op)
	return opcode, 2, 12, true
}

func (cpu *CPU) LDR8HL(reg string) (string, int, int, bool) { // reg[] = M[HL]
	opcode := fmt.Sprintf("LD %s,(HL)", reg)
	cpu.Reg.Write(cpu.GetHLVal(), reg)
	return opcode, 1, 12, true
}

func (cpu *CPU) LDR16A(reg string) (string, int, int, bool) { // M[reg] = reg[A]
	opcode := fmt.Sprintf("LD (%s), A", reg)
	cpu.Mem.Write(cpu.Reg.Read("A"), cpu.Reg.Read(reg))
	return opcode, 1, 8, true
}

func (cpu *CPU) LDHCA() (string, int, int, bool) { // M[$FF00+C] = reg[A]
	opcode := fmt.Sprintf("LD ($FF00+C),A")
	cpu.Mem.Write(cpu.Reg.Read("A"), int(uint16(0xFF00+int(cpu.Reg.Read("C")))))
	return opcode, 1, 8, true
}

func (cpu *CPU) LDN16A() (string, int, int, bool) { // M[n16] = reg[A]
	addr := int(uint16(cpu.Mem.Read(cpu.Reg.Read("PC")+1) + cpu.Mem.Read(cpu.Reg.Read("PC")+2)<<8))
	opcode := fmt.Sprintf("LD ($%x),A", addr)
	cpu.Mem.Write(cpu.Reg.Read("A"), addr)
	return opcode, 3, 16, true
}

func (cpu *CPU) LDHU8A() (string, int, int, bool) { // M[$FF00+n8] = reg[A]
	addr := int(uint16(0xFF00 + cpu.Mem.Read(cpu.Reg.Read("PC")+1)))
	opcode := fmt.Sprintf("LD $(%x),A", addr)
	cpu.Mem.Write(cpu.Reg.Read("A"), addr)
	return opcode, 2, 12, true
}

func (cpu *CPU) LDAR16(reg string) (string, int, int, bool) { // reg[A] = M[reg[]]
	opcode := fmt.Sprintf("LD A, (%s)", reg)
	cpu.Reg.Write(cpu.Mem.Read(cpu.Reg.Read(reg)), "A")
	return opcode, 1, 8, true
}

func (cpu *CPU) LDHAU8() (string, int, int, bool) { // reg[A] = M[$FF00+u8]
	addr := int(uint16(0xFF00 + int(uint8(cpu.Mem.Read(cpu.Reg.Read("PC")+1)))))
	opcode := fmt.Sprintf("LD A,($%x)", addr)
	cpu.Reg.Write(cpu.Mem.Read(addr), "A")
	return opcode, 2, 12, true
}

func (cpu *CPU) LDAN16() (string, int, int, bool) { // reg[A] = M[u16]
	addr := int(uint16(cpu.Mem.Read(cpu.Reg.Read("PC")+1) + (cpu.Mem.Read(cpu.Reg.Read("PC")+2) << 8)))
	opcode := fmt.Sprintf("LD A,($%x)", addr)
	cpu.Reg.Write(cpu.Mem.Read(addr), "A")
	return opcode, 3, 16, true
}

func (cpu *CPU) LDHAC() (string, int, int, bool) { // reg[A] = M[$FF00+reg[C]]
	opcode := fmt.Sprintf("LD A,(0xFF00+C)")
	cpu.Reg.Write(cpu.Mem.Read(int(uint16(0xFF00+cpu.Reg.Read("C")))), "A")
	return opcode, 1, 8, true
}

func (cpu *CPU) LDHLIA() (string, int, int, bool) { // M[HL] = reg[A], HL++
	opcode := fmt.Sprintf("LD (HL+),A")
	cpu.SetHLVal(cpu.Reg.Read("A"))
	cpu.Reg.Write(int(uint16(cpu.Reg.Read("HL")+1)), "HL")
	return opcode, 1, 8, true
}

func (cpu *CPU) LDHLDA() (string, int, int, bool) { // M[HL] = reg[A], HL--
	opcode := fmt.Sprintf("LD (HL-),A")
	cpu.SetHLVal(cpu.Reg.Read("A"))
	cpu.Reg.Write(int(uint16(cpu.Reg.Read("HL")-1)), "HL")
	return opcode, 1, 8, true
}

func (cpu *CPU) LDAHLI() (string, int, int, bool) { // reg[A] = M[HL], HL++
	opcode := fmt.Sprintf("LD A,(HL+)")
	cpu.Reg.Write(cpu.GetHLVal(), "A")
	cpu.Reg.Write(int(uint16(cpu.Reg.Read("HL")+1)), "HL")
	return opcode, 1, 8, true
}

func (cpu *CPU) LDAHLD() (string, int, int, bool) { // reg[A] = M[HL], HL--
	opcode := fmt.Sprintf("LD A,(HL-)")
	cpu.Reg.Write(cpu.GetHLVal(), "A")
	cpu.Reg.Write(int(uint16(cpu.Reg.Read("HL")-1)), "HL")
	return opcode, 1, 8, true
}
