// Flag helpers for corresponding opcode type
// Set flag for given opcode instruction

package CPU

func (s *CPU) SetFlagADD(halfcarry int, carry int, val int) {
	s.ResetFlag("N")
	if halfcarry&0x10 == 0x10 {
		s.SetFlag("H")
	} else {
		s.ResetFlag("H")
	}
	if carry&0x100 == 0x100 {
		s.SetFlag("C")
	} else {
		s.ResetFlag("C")
	}
	if val == 0 {
		s.SetFlag("Z")
	} else {
		s.ResetFlag("Z")
	}
}

func (s *CPU) SetFlagAND(val int) {
	if val == 0 {
		s.SetFlag("Z")
	} else {
		s.ResetFlag("Z")
	}
	s.ResetFlag("N")
	s.SetFlag("H")
	s.ResetFlag("C")
}

func (s *CPU) SetFlagOR(val int) { // Applies for XOR and OR opcodes
	if val == 0 {
		s.SetFlag("Z")
	} else {
		s.ResetFlag("Z")
	}
	s.ResetFlag("N")
	s.ResetFlag("H")
	s.ResetFlag("C")
}

func (s *CPU) SetFlagSUB(halfborrow int, borrow int, val int) { // Applies for SUB opcodes
	s.SetFlag("N")
	if halfborrow < 0 {
		s.SetFlag("H")
	} else {
		s.ResetFlag("H")
	}
	if borrow < 0 {
		s.SetFlag("C")
	} else {
		s.ResetFlag("C")
	}
	if val == 0 {
		s.SetFlag("Z")
	} else {
		s.ResetFlag("Z")
	}
}

func (s *CPU) SetFlagBIT(val int) { // Applies for BIT opcodes
	if val == 0 {
		s.SetFlag("Z")
	} else {
		s.ResetFlag("Z")
	}
	s.ResetFlag("N")
	s.SetFlag("H")
}

func (s *CPU) SetFlagSWAP(val int) { // Applies for BIT opcodes
	if val == 0 {
		s.SetFlag("Z")
	} else {
		s.ResetFlag("Z")
	}
	s.ResetFlag("N")
	s.ResetFlag("H")
	s.ResetFlag("C")
}

func (s *CPU) SetFlagRL(val int, carry int) { // Applies for RL,RR,SL,SR opcodes
	if val == 0 {
		s.SetFlag("Z")
	} else {
		s.ResetFlag("Z")
	}
	s.ResetFlag("N")
	s.ResetFlag("H")
	if carry == 0 {
		s.ResetFlag("C")
	} else {
		s.SetFlag("C")
	}
}

func (s *CPU) SetFlagCP(halfBorrow int, borrow int, val int) {
	s.SetFlag("N")
	if halfBorrow < 0 {
		s.SetFlag("H")
	} else {
		s.ResetFlag("H")
	}
	if borrow < 0 {
		s.SetFlag("C")
	} else {
		s.ResetFlag("C")
	}
	if val == 0 {
		s.SetFlag("Z")
	} else {
		s.ResetFlag("Z")
	}
}

func (s *CPU) SetFlagDEC(halfBorrow int, val int) {
	s.SetFlag("N")
	if halfBorrow < 0 {
		s.SetFlag("H")
	} else {
		s.ResetFlag("H")
	}
	if val == 0 {
		s.SetFlag("Z")
	} else {
		s.ResetFlag("Z")
	}
}

func (s *CPU) SetFlagINC(halfCarry int, val int) {
	s.ResetFlag("N")
	if halfCarry&0x10 == 0x10 {
		s.SetFlag("H")
	} else {
		s.ResetFlag("H")
	}
	if val == 0 {
		s.SetFlag("Z")
	} else {
		s.ResetFlag("Z")
	}
}

func (s *CPU) SetFlagADD16(halfcarry int, carry int) {
	s.ResetFlag("N")
	if halfcarry == 1 {
		s.SetFlag("H")
	} else {
		s.ResetFlag("H")
	}
	if carry == 1 {
		s.SetFlag("C")
	} else {
		s.ResetFlag("C")
	}
}
