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

func (cpu *CPU) IncrementTimer(t int) {
	cpu.TotalT += t

	for i := 0; i < t; i++ {
		cpu.Sysclk = (cpu.Sysclk + 1) & 0xFFFF
		cpu.Mem.Write(cpu.Sysclk&0xFF, 0xFF04)
		timerEnable := (cpu.Mem.Read(0xFF07) >> 2) & 1
		a := cpu.Mem.Read(0xFF07) & 0x03
		b := (cpu.Mem.Read(0xFF04) >> bitPos[a]) & 1
		currLevel := b & timerEnable
		if cpu.Level == 1 && currLevel == 0 {
			if cpu.Mem.Read(0xFF05) == 0xFF {
				cpu.Mem.Write(cpu.Mem.Read(0xFF06), 0xFF05)
				cpu.SetIFBit(2)
			} else {
				cpu.Mem.Write(cpu.Mem.Read(0xFF05)+1, 0xFF05)
			}
		}
		cpu.Level = int(currLevel)
	}
}
