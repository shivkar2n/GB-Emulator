package GB

var bitPos = map[int]int{
	0: 9,
	1: 3,
	2: 5,
	3: 7,
}

func (GB *GB) IncrementTimer(cycles int) {
	GB.CPU.TotalCycles += cycles

	for i := 0; i < cycles; i++ {
		GB.CPU.Sysclk = (GB.CPU.Sysclk + 1) & 0xFFFF
		GB.MMU.Write(GB.CPU.Sysclk&0xFF, DIV)
		timerEnable := (GB.MMU.Read(TAC) >> 2) & 1
		a := GB.MMU.Read(TAC) & 0x03
		b := (GB.MMU.Read(DIV) >> bitPos[a]) & 1
		currLevel := b & timerEnable
		if GB.CPU.Level == 1 && currLevel == 0 {
			if GB.MMU.Read(TIMA) == 0xFF {
				GB.MMU.Write(GB.MMU.Read(TMA), TIMA)
				GB.SetIFBit(2)
			} else {
				GB.MMU.Write(GB.MMU.Read(TIMA)+1, TIMA)
			}
		}
		GB.CPU.Level = int(currLevel)
	}
}
