package CPU

import "fmt"

func (cpu *CPU) BITU3R8(offset int, reg string) (string, int, int, bool) { // if reg[] >> u3 == 0 then setZeroFlag
	opcode := fmt.Sprintf("BIT %d,%s", offset, reg)
	val := (cpu.Reg.Read(reg) >> offset) & 0x01
	cpu.SetFlagBIT(val)
	return opcode, 2, 8, true
}

func (cpu *CPU) BITU3HL(offset int) (string, int, int, bool) { // if M[reg[HL]] >> u3 == 0 then setZeroFlag
	opcode := fmt.Sprintf("BIT %d,(HL)", offset)
	val := (cpu.GetHLVal() >> offset) & 0x01
	cpu.SetFlagBIT(val)
	return opcode, 2, 12, true
}

func (cpu *CPU) RESU3R8(offset int, reg string) (string, int, int, bool) { // Set bit u3 in reg[] to 0
	opcode := fmt.Sprintf("RES %d,%s", offset, reg)
	k := 0
	for i := 0; i < 8; i++ {
		k = k << 1
		if 7-i != offset {
			k += 1
		}
	}
	cpu.Reg.Write(cpu.Reg.Read(reg)&k, reg)
	return opcode, 2, 8, true
}

func (cpu *CPU) RESU3HL(offset int) (string, int, int, bool) { // Set bit u3 in reg[HL] to 0
	opcode := fmt.Sprintf("RES %d,(HL)", offset)
	k := 0
	for i := 0; i < 8; i++ {
		k = k << 1
		if 7-i != offset {
			k += 1
		}
	}
	cpu.SetHLVal(cpu.GetHLVal() & k)
	return opcode, 2, 16, true
}

func (cpu *CPU) SETU3R8(offset int, reg string) (string, int, int, bool) { // Set bit u3 in reg[] to 1
	opcode := fmt.Sprintf("SET %d,%s", offset, reg)
	k := 1
	for i := 0; i < offset; i++ {
		k = k << 1
	}
	cpu.Reg.Write(cpu.Reg.Read(reg)|k, reg)
	return opcode, 2, 8, true
}

func (cpu *CPU) SETU3HL(offset int) (string, int, int, bool) { // Set bit u3 in M[reg[HL]] to 1
	opcode := fmt.Sprintf("SET %d,(HL)", offset)
	k := 1
	for i := 0; i < offset; i++ {
		k = k << 1
	}
	cpu.SetHLVal(cpu.GetHLVal() | k)
	return opcode, 2, 16, true
}

func (cpu *CPU) SWAPR8(reg string) (string, int, int, bool) { // Swap upper, lower 4 bits in reg[]
	opcode := fmt.Sprintf("SWAP B")
	low := (cpu.Reg.Read(reg) & 0xF)
	high := (cpu.Reg.Read(reg) & 0xF0) >> 4
	val := low<<4 | high
	cpu.Reg.Write(val, reg)
	cpu.SetFlagSWAP(val)
	return opcode, 2, 8, true
}

func (cpu *CPU) SWAPHL() (string, int, int, bool) { // Swap upper, lower 4 bits in mem[reg[HL]]
	opcode := fmt.Sprintf("SWAP (HL)")
	low := cpu.GetHLVal() & 0xF
	high := (cpu.GetHLVal() & 0xF0) >> 4
	val := low<<4 | high
	cpu.SetHLVal(val)
	cpu.SetFlagSWAP(val)
	return opcode, 2, 16, true
}
