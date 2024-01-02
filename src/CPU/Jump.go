package CPU

import "fmt"

func (cpu *CPU) JPN16(op int) (string, int, int, bool) { // reg[PC] = n16
	opcode := fmt.Sprintf("JP $%x ", op)
	cpu.Reg.Write(op, "PC")
	return opcode, 0, 16, true
}

func (cpu *CPU) JPCCN16(cc string) (string, int, int, bool) { // reg[PC] = n16, if cc == true
	addr := cpu.Mem.Read(cpu.Reg.Read("PC")+1) + cpu.Mem.Read(cpu.Reg.Read("PC")+2)<<8
	opcode := fmt.Sprintf("JP %s,$%x ", cc, addr)
	if cpu.checkCC(cc) {
		cpu.JPN16(addr)
		return opcode, 0, 16, true
	}
	return opcode, 3, 12, true

}

func (cpu *CPU) JPHL() (string, int, int, bool) { // reg[PC] = reg[HL]
	opcode := fmt.Sprintf("JP HL")
	cpu.Reg.Write(cpu.Reg.Read("HL"), "PC")
	return opcode, 0, 4, true
}

func (cpu *CPU) JRN16() (string, int, int, bool) { // reg[PC] = reg[PC] + n16
	opcode := fmt.Sprintf("JR $%x", int8(cpu.Mem.Read(cpu.Reg.Read("PC")+1)))
	addr := int(int8(cpu.Mem.Read(cpu.Reg.Read("PC") + 1)))
	cpu.JPN16(cpu.Reg.Read("PC") + addr + 2)
	return opcode, 0, 12, true
}

func (cpu *CPU) JRCCN16(cc string) (string, int, int, bool) { // reg[PC] = reg[PC] + n16, if cc == true
	opcode := fmt.Sprintf("JR %s,$%x", cc, uint8(cpu.Mem.Read(cpu.Reg.Read("PC")+1)))
	if cpu.checkCC(cc) {
		cpu.JRN16()
		return opcode, 0, 12, true
	}
	return opcode, 2, 8, true
}

func (cpu *CPU) RET() (string, int, int, bool) { //
	opcode := fmt.Sprintf("RET")
	cpu.POPR16("PC")
	return opcode, 0, 16, true
}

func (cpu *CPU) RETCC(cc string) (string, int, int, bool) { //
	opcode := fmt.Sprintf("RET %s", cc)
	if cpu.checkCC(cc) {
		cpu.RET()
		return opcode, 0, 20, true
	}
	return opcode, 1, 8, true
}

func (cpu *CPU) RETI() (string, int, int, bool) { // Equivalent to:
	opcode := fmt.Sprintf("RETI")
	cpu.IME = true // EI
	cpu.RET()      // RET
	return opcode, 0, 16, true
}

func (cpu *CPU) CALLN16() (string, int, int, bool) { // Equivalent to following
	addr := cpu.Mem.Read(cpu.Reg.Read("PC")+1) + cpu.Mem.Read(cpu.Reg.Read("PC")+2)<<8
	opcode := fmt.Sprintf("CALL $%x", addr)
	cpu.PUSHN16(cpu.Reg.Read("PC") + 3) // PUSH M[reg[PC+3]]
	cpu.JPN16(addr)                     // JP N16
	return opcode, 0, 24, true
}

func (cpu *CPU) CALLN8(op int) (string, int, int, bool) { // Equivalent to following
	cpu.PUSHN16(cpu.Reg.Read("PC")) // PUSH M[reg[PC]]
	cpu.JPN16(op)                   // JP N16
	return "", 0, 0, true
}

func (cpu *CPU) CALLCCN16(cc string) (string, int, int, bool) { // CALLN16 if cc == true
	addr := cpu.Mem.Read(cpu.Reg.Read("PC")+1) + cpu.Mem.Read(cpu.Reg.Read("PC")+2)<<8
	opcode := fmt.Sprintf("CALL %s,$%x", cc, addr)
	if cpu.checkCC(cc) {
		cpu.CALLN16()
		return opcode, 0, 24, true
	}
	return opcode, 3, 12, true
}

func (cpu *CPU) RST(vec string) (string, int, int, bool) {
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
	cpu.PUSHN16(cpu.Reg.Read("PC") + 1)
	cpu.JPN16(vecAddr)
	return opcode, 0, 16, true
}
