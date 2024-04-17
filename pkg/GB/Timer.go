package GB

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

func (GB *GB) IncrementTimer() {
	GB.CPU.TotalT += GB.TCycles

	for i := 0; i < GB.TCycles; i++ {
		GB.CPU.Sysclk = (GB.CPU.Sysclk + 1) & 0xFFFF
		GB.MMU.Write(GB.CPU.Sysclk&0xFF, 0xFF04)
		timerEnable := (GB.MMU.Read(0xFF07) >> 2) & 1
		a := GB.MMU.Read(0xFF07) & 0x03
		b := (GB.MMU.Read(0xFF04) >> bitPos[a]) & 1
		currLevel := b & timerEnable
		if GB.CPU.Level == 1 && currLevel == 0 {
			if GB.MMU.Read(0xFF05) == 0xFF {
				GB.MMU.Write(GB.MMU.Read(0xFF06), 0xFF05)
				GB.SetIFBit(2)
			} else {
				GB.MMU.Write(GB.MMU.Read(0xFF05)+1, 0xFF05)
			}
		}
		GB.CPU.Level = int(currLevel)
	}
}
