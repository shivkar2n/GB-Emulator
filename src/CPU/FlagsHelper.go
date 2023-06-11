// Flag helpers for corresponding opcode type
// Set flag for given opcode instruction

package CPU

func (s *CPU) SetFlagADD(halfcarry int, carry int, val int) {
	s.ResetFlag("N")
	if halfcarry&0x10 == 0x10 {
		s.SetFlag("H")
	}
	if carry&0x100 == 0x100 {
		s.SetFlag("C")
	}
	if val == 0 {
		s.SetFlag("Z")
	}
}

func (s *CPU) SetFlagAND(val int) {
	s.ResetFlag("N")
	s.ResetFlag("C")
	s.SetFlag("H")
	if val == 0 {
		s.SetFlag("Z")
	}
}

func (s *CPU) SetFlagOR(val int) { // Applies for XOR and OR opcodes
	if val == 0 {
		s.SetFlag("Z")
	}
	s.ResetFlag("N")
	s.ResetFlag("H")
	s.ResetFlag("C")
}

func (s *CPU) SetFlagSUB(halfborrow bool, borrow bool, val int) { // Applies for SUB opcodes
	s.SetFlag("N")
	if halfborrow {
		s.SetFlag("H")
	}
	if borrow {
		s.SetFlag("C")
	}
	if val == 0 {
		s.SetFlag("Z")
	}
}

func (s *CPU) SetFlagBIT(val int) { // Applies for BIT opcodes
	if val == 0 {
		s.SetFlag("Z")
	}
	s.ResetFlag("N")
	s.SetFlag("H")
}

func (s *CPU) SetFlagSWAP(val int) { // Applies for BIT opcodes
	if val == 0 {
		s.SetFlag("Z")
	}
	s.ResetFlag("N")
	s.ResetFlag("H")
	s.ResetFlag("C")
}

func (s *CPU) SetFlagRL(val int, carry int) { // Applies for RL,RR,SL,SR opcodes
	if val == 0 {
		s.SetFlag("Z")
	}
	s.SetFlag("N")
	s.SetFlag("H")
	if carry == 0 {
		s.ResetFlag("C")
	} else {
		s.SetFlag("C")
	}
}
