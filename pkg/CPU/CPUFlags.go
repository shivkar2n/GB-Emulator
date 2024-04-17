// Flag helpers for corresponding opcode type
// Set flag for given opcode instruction

package CPU

func (CPU *CPU) SetFlagADD(halfcarry int, carry int, val int) {
	CPU.ResetFlag("N")
	if halfcarry&0x10 == 0x10 {
		CPU.SetFlag("H")
	} else {
		CPU.ResetFlag("H")
	}
	if carry&0x100 == 0x100 {
		CPU.SetFlag("C")
	} else {
		CPU.ResetFlag("C")
	}
	if val == 0 {
		CPU.SetFlag("Z")
	} else {
		CPU.ResetFlag("Z")
	}
}

func (CPU *CPU) SetFlagAND(val int) {
	if val == 0 {
		CPU.SetFlag("Z")
	} else {
		CPU.ResetFlag("Z")
	}
	CPU.ResetFlag("N")
	CPU.SetFlag("H")
	CPU.ResetFlag("C")
}

func (CPU *CPU) SetFlagOR(val int) { // Applies for XOR and OR opcodes
	if val == 0 {
		CPU.SetFlag("Z")
	} else {
		CPU.ResetFlag("Z")
	}
	CPU.ResetFlag("N")
	CPU.ResetFlag("H")
	CPU.ResetFlag("C")
}

func (CPU *CPU) SetFlagSUB(halfborrow int, borrow int, val int) { // Applies for SUB opcodes
	CPU.SetFlag("N")
	if halfborrow < 0 {
		CPU.SetFlag("H")
	} else {
		CPU.ResetFlag("H")
	}
	if borrow < 0 {
		CPU.SetFlag("C")
	} else {
		CPU.ResetFlag("C")
	}
	if val == 0 {
		CPU.SetFlag("Z")
	} else {
		CPU.ResetFlag("Z")
	}
}

func (CPU *CPU) SetFlagBIT(val int) { // Applies for BIT opcodes
	if val == 0 {
		CPU.SetFlag("Z")
	} else {
		CPU.ResetFlag("Z")
	}
	CPU.ResetFlag("N")
	CPU.SetFlag("H")
}

func (CPU *CPU) SetFlagSWAP(val int) { // Applies for BIT opcodes
	if val == 0 {
		CPU.SetFlag("Z")
	} else {
		CPU.ResetFlag("Z")
	}
	CPU.ResetFlag("N")
	CPU.ResetFlag("H")
	CPU.ResetFlag("C")
}

func (CPU *CPU) SetFlagRL(val int, carry int) { // Applies for RL,RR,SL,SR opcodes
	if val == 0 {
		CPU.SetFlag("Z")
	} else {
		CPU.ResetFlag("Z")
	}
	CPU.ResetFlag("N")
	CPU.ResetFlag("H")
	if carry == 0 {
		CPU.ResetFlag("C")
	} else {
		CPU.SetFlag("C")
	}
}

func (CPU *CPU) SetFlagCP(halfBorrow int, borrow int, val int) {
	CPU.SetFlag("N")
	if halfBorrow < 0 {
		CPU.SetFlag("H")
	} else {
		CPU.ResetFlag("H")
	}
	if borrow < 0 {
		CPU.SetFlag("C")
	} else {
		CPU.ResetFlag("C")
	}
	if val == 0 {
		CPU.SetFlag("Z")
	} else {
		CPU.ResetFlag("Z")
	}
}

func (CPU *CPU) SetFlagDEC(halfBorrow int, val int) {
	CPU.SetFlag("N")
	if halfBorrow < 0 {
		CPU.SetFlag("H")
	} else {
		CPU.ResetFlag("H")
	}
	if val == 0 {
		CPU.SetFlag("Z")
	} else {
		CPU.ResetFlag("Z")
	}
}

func (CPU *CPU) SetFlagINC(halfCarry int, val int) {
	CPU.ResetFlag("N")
	if halfCarry&0x10 == 0x10 {
		CPU.SetFlag("H")
	} else {
		CPU.ResetFlag("H")
	}
	if val == 0 {
		CPU.SetFlag("Z")
	} else {
		CPU.ResetFlag("Z")
	}
}

func (CPU *CPU) SetFlagADD16(halfcarry int, carry int) {
	CPU.ResetFlag("N")
	if halfcarry == 1 {
		CPU.SetFlag("H")
	} else {
		CPU.ResetFlag("H")
	}
	if carry == 1 {
		CPU.SetFlag("C")
	} else {
		CPU.ResetFlag("C")
	}
}
