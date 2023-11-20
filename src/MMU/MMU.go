package MMU

import "os"

type MMU struct {
	Ram          [65536]byte
	Rom          []byte
	romBank      int
	ramBank      int
	mode         bool
	enableExtRam bool
	zeroBank     int
	highBank     int
	bitMask      int
}

func Init() MMU { // Load ROM
	rom, _ := os.ReadFile(os.Args[1])
	mask := 0
	for i := 0; i < len(rom)/32*1024; i++ {
		mask = mask << 1
		mask = mask | 1
	}
	m := MMU{
		Rom:      rom,
		ramBank:  0,
		romBank:  1,
		mode:     false,
		zeroBank: 0,
		highBank: 0,
		bitMask:  mask,
	}
	for k, v := range rom { // If NO MBC
		m.Ram[k] = v
	}
	return m
}
