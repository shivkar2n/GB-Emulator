// Flag helpers for corresponding opcode type
// Set flag for given opcode instruction

package CPU

func (cpu *CPU) SetFlagADD(halfcarry int, carry int, val int) {
	cpu.ResetFlag("N")
	if halfcarry&0x10 == 0x10 {
		cpu.SetFlag("H")
	} else {
		cpu.ResetFlag("H")
	}
	if carry&0x100 == 0x100 {
		cpu.SetFlag("C")
	} else {
		cpu.ResetFlag("C")
	}
	if val == 0 {
		cpu.SetFlag("Z")
	} else {
		cpu.ResetFlag("Z")
	}
}

func (cpu *CPU) SetFlagAND(val int) {
	if val == 0 {
		cpu.SetFlag("Z")
	} else {
		cpu.ResetFlag("Z")
	}
	cpu.ResetFlag("N")
	cpu.SetFlag("H")
	cpu.ResetFlag("C")
}

func (cpu *CPU) SetFlagOR(val int) { // Applies for XOR and OR opcodes
	if val == 0 {
		cpu.SetFlag("Z")
	} else {
		cpu.ResetFlag("Z")
	}
	cpu.ResetFlag("N")
	cpu.ResetFlag("H")
	cpu.ResetFlag("C")
}

func (cpu *CPU) SetFlagSUB(halfborrow int, borrow int, val int) { // Applies for SUB opcodes
	cpu.SetFlag("N")
	if halfborrow < 0 {
		cpu.SetFlag("H")
	} else {
		cpu.ResetFlag("H")
	}
	if borrow < 0 {
		cpu.SetFlag("C")
	} else {
		cpu.ResetFlag("C")
	}
	if val == 0 {
		cpu.SetFlag("Z")
	} else {
		cpu.ResetFlag("Z")
	}
}

func (cpu *CPU) SetFlagBIT(val int) { // Applies for BIT opcodes
	if val == 0 {
		cpu.SetFlag("Z")
	} else {
		cpu.ResetFlag("Z")
	}
	cpu.ResetFlag("N")
	cpu.SetFlag("H")
}

func (cpu *CPU) SetFlagSWAP(val int) { // Applies for BIT opcodes
	if val == 0 {
		cpu.SetFlag("Z")
	} else {
		cpu.ResetFlag("Z")
	}
	cpu.ResetFlag("N")
	cpu.ResetFlag("H")
	cpu.ResetFlag("C")
}

func (cpu *CPU) SetFlagRL(val int, carry int) { // Applies for RL,RR,SL,SR opcodes
	if val == 0 {
		cpu.SetFlag("Z")
	} else {
		cpu.ResetFlag("Z")
	}
	cpu.ResetFlag("N")
	cpu.ResetFlag("H")
	if carry == 0 {
		cpu.ResetFlag("C")
	} else {
		cpu.SetFlag("C")
	}
}

func (cpu *CPU) SetFlagCP(halfBorrow int, borrow int, val int) {
	cpu.SetFlag("N")
	if halfBorrow < 0 {
		cpu.SetFlag("H")
	} else {
		cpu.ResetFlag("H")
	}
	if borrow < 0 {
		cpu.SetFlag("C")
	} else {
		cpu.ResetFlag("C")
	}
	if val == 0 {
		cpu.SetFlag("Z")
	} else {
		cpu.ResetFlag("Z")
	}
}

func (cpu *CPU) SetFlagDEC(halfBorrow int, val int) {
	cpu.SetFlag("N")
	if halfBorrow < 0 {
		cpu.SetFlag("H")
	} else {
		cpu.ResetFlag("H")
	}
	if val == 0 {
		cpu.SetFlag("Z")
	} else {
		cpu.ResetFlag("Z")
	}
}

func (cpu *CPU) SetFlagINC(halfCarry int, val int) {
	cpu.ResetFlag("N")
	if halfCarry&0x10 == 0x10 {
		cpu.SetFlag("H")
	} else {
		cpu.ResetFlag("H")
	}
	if val == 0 {
		cpu.SetFlag("Z")
	} else {
		cpu.ResetFlag("Z")
	}
}

func (cpu *CPU) SetFlagADD16(halfcarry int, carry int) {
	cpu.ResetFlag("N")
	if halfcarry == 1 {
		cpu.SetFlag("H")
	} else {
		cpu.ResetFlag("H")
	}
	if carry == 1 {
		cpu.SetFlag("C")
	} else {
		cpu.ResetFlag("C")
	}
}
