package MMU

import "os"

const (
	JOYP           = 0xFF00
	SERIAL         = 0xFF01
	DMA            = 0xFF46
	IE             = 0xFFFF
	IF             = 0xFF0F
	CARTRIDGE_TYPE = 0x0147
)

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
	ButtonState  byte
	DpadState    byte
}

func Init() *MMU { // Load ROM
	rom, _ := os.ReadFile(os.Args[1])
	mask := 0
	for i := 0; i < len(rom)/32768; i++ {
		mask = mask << 1
		mask = mask | 1
	}

	m := &MMU{
		Rom:         rom,
		ramBank:     0,
		romBank:     1,
		mode:        false,
		zeroBank:    0,
		highBank:    0,
		bitMask:     mask,
		ButtonState: 0xF,
		DpadState:   0xF,
	}
	copy(m.Ram[:], rom)
	m.Ram[JOYP] = byte(0xFF)

	// Load Nintendo Boot logo
	// logo := []byte{0xCE, 0xED, 0x66, 0x66, 0xCC, 0x0D, 0x00, 0x0B, 0x03, 0x73, 0x00, 0x83, 0x00, 0x0C, 0x00, 0x0D,
	// 	0x00, 0x08, 0x11, 0x1F, 0x88, 0x89, 0x00, 0x0E, 0xDC, 0xCC, 0x6E, 0xE6, 0xDD, 0xDD, 0xD9, 0x99,
	// 	0xBB, 0xBB, 0x67, 0x63, 0x6E, 0x0E, 0xEC, 0xCC, 0xDD, 0xDC, 0x99, 0x9F, 0xBB, 0xB9, 0x33, 0x3E}

	// for i := 0x0104; i <= 0x0133; i++ {
	// 	m.Ram[i] = logo[i-0x0104]
	// }

	return m
}
