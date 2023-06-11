package CPU

func (s *CPU) BITU3R8(offset int, reg string) {
	val := (s.GetReg8Val(reg) >> offset) & 0x01
	s.SetFlagBIT(val)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+2)
}

func (s *CPU) BITU3HL(offset int) { // if reg[] >> u3 == 0 then setZeroFlag
	val := s.GetHLVal() >> offset
	s.SetFlagBIT(val)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+2)
}

func (s *CPU) RESU3R8(offset int, reg string) { // Set bit u3 in reg[] to 0
	k := 0
	for i := 0; i < 8; i++ {
		k = k << 1
		if 7-i != offset {
			k += 1
		}
	}
	s.SetReg8Val(reg, s.GetReg8Val(reg)&k)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+2)
}

func (s *CPU) RESU3HL(offset int) { // Set bit u3 in reg[HL] to 0
	k := 0
	for i := 0; i < 8; i++ {
		k = k << 1
		if 7-i != offset {
			k += 1
		}
	}
	s.SetHLVal(s.GetHLVal() & k)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+2)
}

func (s *CPU) SETU3R8(offset int, reg string) { // Set bit u3 in reg[] to 1
	k := 1
	for i := 0; i < offset; i++ {
		k = k << 1
	}
	s.SetReg8Val(reg, s.GetReg8Val(reg)|k)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+2)
}

func (s *CPU) SETU3HL(offset int) { // Set bit u3 in M[reg[HL]] to 1
	k := 1
	for i := 0; i < offset; i++ {
		k = k << 1
	}
	s.SetHLVal(s.GetHLVal() | k)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+2)
}

func (s *CPU) SWAPR8(reg string) { // Swap upper, lower 4 bits in reg[]
	low := (s.GetReg8Val(reg) & 0xF)
	high := (s.GetReg8Val(reg) & 0xF0) >> 4
	val := low<<4 | high
	s.SetReg8Val(reg, val)
	s.SetFlagSWAP(val)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+2)
}

func (s *CPU) SWAPHL() { // Swap upper, lower 4 bits in mem[reg[HL]]
	low := s.GetHLVal() & 0xF
	high := (s.GetHLVal() & 0xF0) >> 4
	val := low<<4 | high
	s.Mem[s.GetReg16Val("HL")] = byte(val)
	s.SetFlagSWAP(val)
	s.SetReg16Val("PC", s.GetReg16Val("PC")+2)
}