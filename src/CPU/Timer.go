package CPU

// CLOCK_SPEED=4194304Hz
// FF04 -> DIV
// FF05 -> TIMA
// FF06 -> TMA
// FF07 -> TAC

func (s *CPU) IncrTimer() {
	s.TotalM += s.M
	s.TotalT += s.T
	s.TIMA += s.M

	if s.M >= 4 {
		s.Sysclk += 1
		if s.Sysclk >= 16 {
			s.Mem[0xFF04] = byte(s.Mem[0xFF04] + 1)
			s.Sysclk = s.Sysclk % 16
		}
	}

	if s.TIMA > s.GetTACFreq() {
		s.TIMA = s.TIMA % s.GetTACFreq()
		val := s.Mem[0xFF05] + 1
		s.Mem[0xFF05] = byte(val)
		if val > 0xFF {
			s.Mem[0xFF05] = s.Mem[0xFF06]
			s.SetIFBit(2)
		}
	}
}

func (s *CPU) GetTACFreq() int { // Get Clock Cycle rate for given (f)
	switch s.Mem[0xFF07] & 0x02 {
	case 0:
		return s.ClkRate / 4096

	case 1:
		return s.ClkRate / 262144

	case 2:
		return s.ClkRate / 65536

	case 3:
		return s.ClkRate / 16384

	default:
		return -1
	}
}
