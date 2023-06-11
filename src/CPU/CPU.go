package CPU

import (
	"fmt"
	"os"
)

type CPU struct {
	FA, CB, ED, LH, PC, SP [2]byte
	Mem                    [65536]byte
	IME                    bool
}

func InitCPU() *CPU {
	c := CPU{
		FA:  [2]byte{},
		CB:  [2]byte{},
		ED:  [2]byte{},
		LH:  [2]byte{},
		PC:  [2]byte{},
		SP:  [2]byte{},
		Mem: [65536]byte{},
		IME: false,
	}
	return &c
}

func (s *CPU) GetFlag(flagType string) int { // Get value of flag
	if flagType == "Z" {
		return int(s.FA[0]) & 0x80 >> 7
	} else if flagType == "N" {
		return int(s.FA[0]) & 0x40 >> 6
	} else if flagType == "H" {
		return int(s.FA[0]) & 0x20 >> 5
	} else if flagType == "C" {
		return int(s.FA[0]) & 0x10 >> 4
	}
	return -1
}

func (s *CPU) ResetFlag(flagType string) { // Check if flag bit set to 1, then change to 0
	if flagType == "Z" && (uint8(s.FA[0])&0x80)>>7 == 1 {
		s.FA[0] = byte(s.FA[0] & 0x7F)
	} else if flagType == "N" && (uint8(s.FA[0])&0x40)>>6 == 1 {
		s.FA[0] = byte(s.FA[0] & 0xBF)
	} else if flagType == "H" && (uint8(s.FA[0])&0x20)>>5 == 1 {
		s.FA[0] = byte(s.FA[0] & 0xDF)
	} else if flagType == "C" && (uint8(s.FA[0])&0x10)>>4 == 1 {
		s.FA[0] = byte(s.FA[0] & 0xEF)
	}
}

func (s *CPU) SetFlag(flagType string) { // Check if flag bit set to 0, then change to 1
	if flagType == "Z" && (s.FA[0]&0x80)>>7 == 0 {
		s.FA[0] = byte(s.FA[0] ^ 0x80)
	} else if flagType == "N" && (s.FA[0]&0x40)>>6 == 0 {
		s.FA[0] = byte(s.FA[0] ^ 0x40)
	} else if flagType == "H" && (s.FA[0]&0x20)>>5 == 0 {
		s.FA[0] = byte(s.FA[0] ^ 0x20)
	} else if flagType == "C" && (s.FA[0]&0x10)>>4 == 0 {
		s.FA[0] = byte(s.FA[0] ^ 0x10)
	}
}

// func (s *CPU) DebugInfo() { // Get info about state of CPU
// 	fmt.Printf("Registers AF: %X\n", s.GetReg16Val("AF"))
// 	fmt.Printf("Registers BC: %X\n", s.GetReg16Val("BC"))
// 	fmt.Printf("Registers DE: %X\n", s.GetReg16Val("DE"))
// 	fmt.Printf("Registers HL: %X\n", s.GetReg16Val("HL"))
// 	fmt.Printf("M[HL] : %X\n", s.Mem[s.GetReg16Val("HL")])
// 	fmt.Printf("Registers A: %X\n", s.GetReg8Val("A"))
// 	fmt.Printf("Registers B: %X\n", s.GetReg8Val("B"))
// 	fmt.Printf("Registers C: %X\n", s.GetReg8Val("C"))
// 	fmt.Printf("Registers D: %X\n", s.GetReg8Val("D"))
// 	fmt.Printf("Registers E: %X\n", s.GetReg8Val("E"))
// 	fmt.Printf("Registers F: %X\n", s.GetReg8Val("F"))
// 	fmt.Printf("Registers H: %X\n", s.GetReg8Val("H"))
// 	fmt.Printf("Registers L: %X\n", s.GetReg8Val("L"))
// 	fmt.Printf("Program Counter: %X\n", s.GetReg16Val("PC"))
// 	fmt.Printf("Stack Pointer: %X\n", s.GetReg16Val("SP"))
// 	fmt.Printf("Flag Z: %d\n", s.GetFlag("Z"))
// 	fmt.Printf("Flag N: %d\n", s.GetFlag("N"))
// 	fmt.Printf("Flag H: %d\n", s.GetFlag("H"))
// 	fmt.Printf("Flag C: %d\n\n", s.GetFlag("C"))
// }

func (s *CPU) MemInfo() {
	for i := 0; i < 0xFFFF; i += 16 {
		fmt.Printf("%04X: %02X %02X %02X %02X %02X %02X %02X %02X", i, s.Mem[i], s.Mem[i+1], s.Mem[i+2], s.Mem[i+3], s.Mem[i+4], s.Mem[i+5], s.Mem[i+6], s.Mem[i+7])
		fmt.Printf(" %02X %02X %02X %02X %02X %02X %02X %02X\n", s.Mem[i+8], s.Mem[i+9], s.Mem[i+10], s.Mem[i+11], s.Mem[i+12], s.Mem[i+13], s.Mem[i+14], s.Mem[i+15])
	}
}

func (s *CPU) DebugInfo() { // Get info about state of CPU
	a := s.GetReg8Val("A")
	f := s.GetReg8Val("F")
	b := s.GetReg8Val("B")
	c := s.GetReg8Val("C")
	d := s.GetReg8Val("D")
	e := s.GetReg8Val("E")
	h := s.GetReg8Val("H")
	l := s.GetReg8Val("L")
	sp := s.GetReg16Val("SP")
	pc := s.GetReg16Val("PC")
	pcMem := s.Mem[s.GetReg16Val("PC")]
	pcMem1 := s.Mem[s.GetReg16Val("PC")+1]
	pcMem2 := s.Mem[s.GetReg16Val("PC")+2]
	pcMem3 := s.Mem[s.GetReg16Val("PC")+3]
	// fmt.Printf("A: %02X F: %02X B: %02X C: %02X D: %02X E: %02X H: %02X L: %02X SP: %04X PC: 00:%04X (%02X %02X %02X %02X)\n", a, f, b, c, d, e, h, l, sp, pc, pcMem, pcMem1, pcMem2, pcMem3)
	fmt.Printf("A:%02X F:%02X B:%02X C:%02X D:%02X E:%02X H:%02X L:%02X SP:%04X PC:%04X PCMEM:%02X,%02X,%02X,%02X\n", a, f, b, c, d, e, h, l, sp, pc, pcMem, pcMem1, pcMem2, pcMem3)
}

func (s *CPU) GetReg8Val(name string) int { // Get value of 8-bit register
	if name == "A" {
		return int(s.FA[1])
	} else if name == "F" {
		return int(s.FA[0])
	} else if name == "B" {
		return int(s.CB[1])
	} else if name == "C" {
		return int(s.CB[0])
	} else if name == "D" {
		return int(s.ED[1])
	} else if name == "E" {
		return int(s.ED[0])
	} else if name == "H" {
		return int(s.LH[1])
	} else if name == "L" {
		return int(s.LH[0])
	}
	return 0
}

func (s *CPU) SetReg8Val(name string, val int) { // Set value of 8-bit register
	if name == "A" {
		s.FA[1] = byte(val)
	} else if name == "F" {
		s.FA[0] = byte(val)
	} else if name == "B" {
		s.CB[1] = byte(val)
	} else if name == "C" {
		s.CB[0] = byte(val)
	} else if name == "D" {
		s.ED[1] = byte(val)
	} else if name == "E" {
		s.ED[0] = byte(val)
	} else if name == "H" {
		s.LH[1] = byte(val)
	} else if name == "L" {
		s.LH[0] = byte(val)
	}
}

func (s *CPU) GetHLVal() int { // Get value pointed by HL register
	addr := 256*int(s.LH[1]) + int(s.LH[0])
	return int(s.Mem[addr])
}

func (s *CPU) SetHLVal(val int) { // Set value pointed by HL register
	addr := 256*int(s.LH[1]) + int(s.LH[0])
	s.Mem[addr] = byte(val)
}

func (s *CPU) GetPCVal() int { // Get value pointed by PC register
	addr := 256*int(s.PC[1]) + int(s.PC[0])
	return int(s.Mem[addr])
}

func (s *CPU) SetPCVal(val int) { // Set value pointed by PC register
	addr := 256*int(s.PC[1]) + int(s.PC[0])
	s.Mem[addr] = byte(val)
}

func (s *CPU) GetReg16Val(name string) int { // Get value of 16-bit register
	if name == "AF" {
		return 256*int(s.FA[1]) + int(s.FA[0])
	} else if name == "BC" {
		return 256*int(s.CB[1]) + int(s.CB[0])
	} else if name == "DE" {
		return 256*int(s.ED[1]) + int(s.ED[0])
	} else if name == "HL" {
		return 256*int(s.LH[1]) + int(s.LH[0])
	} else if name == "SP" {
		return 256*int(s.SP[1]) + int(s.SP[0])
	} else if name == "PC" {
		return 256*int(s.PC[1]) + int(s.PC[0])
	}

	return -1
}

func (s *CPU) SetReg16Val(name string, val int) { // Set value of 16-bit register
	if name == "AF" {
		s.FA[0] = byte(val & 0x00FF)
		s.FA[1] = byte((val & 0xFF00) >> 8)
	} else if name == "BC" {
		s.CB[0] = byte(val & 0x00FF)
		s.CB[1] = byte((val & 0xFF00) >> 8)
	} else if name == "DE" {
		s.ED[0] = byte(val & 0x00FF)
		s.ED[1] = byte((val & 0xFF00) >> 8)
	} else if name == "HL" {
		s.LH[0] = byte(val & 0x00FF)
		s.LH[1] = byte((val & 0xFF00) >> 8)
	} else if name == "SP" {
		s.SP[0] = byte(val & 0x00FF)
		s.SP[1] = byte((val & 0xFF00) >> 8)
	} else if name == "PC" {
		s.PC[0] = byte(val & 0x00FF)
		s.PC[1] = byte((val & 0xFF00) >> 8)
	}
}

func (s *CPU) checkCC(cc string) bool { // Check condition
	if cc == "Z" {
		return s.GetFlag("Z") == 1
	} else if cc == "NZ" {
		return s.GetFlag("Z") == 0
	} else if cc == "C" {
		return s.GetFlag("C") == 1
	} else if cc == "NC" {
		return s.GetFlag("C") == 0
	}
	return false
}

func (s *CPU) LoadROM() { // Load ROM into memory
	fileName := os.Args[1]
	rom, _ := os.ReadFile(fileName)
	copy(s.Mem[:], rom)
}