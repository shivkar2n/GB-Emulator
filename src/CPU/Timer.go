package CPU

// CLOCK_SPEED=4194304Hz
// FF04 -> DIV
// FF05 -> TIMA
// FF06 -> TMA
// FF07 -> TAC

var bitPos = map[int]int{
	0: 9,
	1: 3,
	2: 5,
	3: 7,
}

func (s *CPU) IncrTimer() {
	s.TotalM += s.M
	s.TotalT += s.T

	for i := 0; i < s.T; i++ {
		s.Sysclk = (s.Sysclk + 1) & 0xFFFF
		s.Mem.Write(s.Sysclk&0xFF, 0xFF04)
		timerEnable := (int(s.Mem.Read(0xFF07)) >> 2) & 1
		a := int(s.Mem.Read(0xFF07)) & 0x03
		b := (s.Mem.Read(0xFF04) >> bitPos[a]) & 1
		currLevel := b & byte(timerEnable)
		if s.Level == 1 && currLevel == 0 {
			if int(s.Mem.Read(0xFF05)) == 0xFF {
				s.Mem.Write(int(s.Mem.Read(0xFF06)), 0xFF05)
				s.SetIFBit(2)
			} else {
				s.Mem.Write(int(s.Mem.Read(0xFF05))+1, 0xFF05)
			}
		}
		s.Level = int(currLevel)
	}

	s.T = 0
	s.M = 0
}
