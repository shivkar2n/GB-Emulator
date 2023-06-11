package CPU

import (
	"fmt"
)

type CPU struct {
	AF, BC, DE, HL, PC, SP [2]byte
	Mem                    [65536]byte
}

func InitCPU() *CPU {
	c := CPU{
		AF:  [2]byte{},
		BC:  [2]byte{},
		DE:  [2]byte{},
		HL:  [2]byte{},
		PC:  [2]byte{},
		SP:  [2]byte{},
		Mem: [65536]byte{},
	}
	return &c
}

func (s *CPU) GetFlag(flagType string) int { // Get value of flag
	if flagType == "Z" {
		return int(s.AF[0]) & 0x80 >> 7
	} else if flagType == "N" {
		return int(s.AF[0]) & 0x40 >> 6
	} else if flagType == "H" {
		return int(s.AF[0]) & 0x20 >> 5
	} else if flagType == "C" {
		return int(s.AF[0]) & 0x10 >> 4
	}
	return -1
}

func (s *CPU) ResetFlag(flagType string) { // Check if flag bit set to 1, then change to 0
	if flagType == "Z" && (uint8(s.AF[0])&0x80)>>7 == 1 {
		s.AF[0] = byte(s.AF[0] & 0x7F)
	} else if flagType == "N" && (uint8(s.AF[0])&0x40)>>6 == 1 {
		s.AF[0] = byte(s.AF[0] & 0xBF)
	} else if flagType == "H" && (uint8(s.AF[0])&0x20)>>5 == 1 {
		s.AF[0] = byte(s.AF[0] & 0xDF)
	} else if flagType == "C" && (uint8(s.AF[0])&0x10)>>4 == 1 {
		s.AF[0] = byte(s.AF[0] & 0xEF)
	}
}

func (s *CPU) SetFlag(flagType string) { // Check if flag bit set to 0, then change to 1
	if flagType == "Z" && (s.AF[0]&0x80)>>7 == 0 {
		s.AF[0] = byte(s.AF[0] ^ 0x80)
	} else if flagType == "N" && (s.AF[0]&0x40)>>6 == 0 {
		s.AF[0] = byte(s.AF[0] ^ 0x40)
	} else if flagType == "H" && (s.AF[0]&0x20)>>5 == 0 {
		s.AF[0] = byte(s.AF[0] ^ 0x20)
	} else if flagType == "C" && (s.AF[0]&0x10)>>4 == 0 {
		s.AF[0] = byte(s.AF[0] ^ 0x10)
	}
}

func (s *CPU) DebugInfo() { // Get info about state of CPU
	fmt.Printf("Registers AF: %X\n", s.GetReg16Val("AF"))
	fmt.Printf("Registers BC: %X\n", s.GetReg16Val("BC"))
	fmt.Printf("Registers DE: %X\n", s.GetReg16Val("DE"))
	fmt.Printf("Registers HL: %X\n", s.GetReg16Val("HL"))
	fmt.Printf("Registers A: %X\n", s.GetReg8Val("A"))
	fmt.Printf("Registers B: %X\n", s.GetReg8Val("B"))
	fmt.Printf("Registers C: %X\n", s.GetReg8Val("C"))
	fmt.Printf("Registers D: %X\n", s.GetReg8Val("D"))
	fmt.Printf("Registers E: %X\n", s.GetReg8Val("E"))
	fmt.Printf("Registers F: %X\n", s.GetReg8Val("F"))
	fmt.Printf("Program Counter: %X\n", s.PC)
	fmt.Printf("Stack Pointer: %X\n", s.GetReg16Val("SP"))
	fmt.Printf("Flag Z: %d\n", s.GetFlag("Z"))
	fmt.Printf("Flag N: %d\n", s.GetFlag("N"))
	fmt.Printf("Flag H: %d\n", s.GetFlag("H"))
	fmt.Printf("Flag C: %d\n\n", s.GetFlag("C"))
}

func (s *CPU) GetReg8Val(name string) int { // Get value of 8-bit register
	if name == "A" {
		return int(s.AF[1])
	} else if name == "B" {
		return int(s.BC[1])
	} else if name == "F" {
		return int(s.AF[0])
	} else if name == "C" {
		return int(s.BC[0])
	} else if name == "D" {
		return int(s.DE[1])
	} else if name == "E" {
		return int(s.DE[0])
	} else if name == "H" {
		return int(s.HL[1])
	} else if name == "L" {
		return int(s.HL[0])
	}
	return 0
}

func (s *CPU) SetReg8Val(name string, val int) { // Set value of 8-bit register
	if name == "A" {
		s.AF[1] = byte(val)
	} else if name == "B" {
		s.BC[1] = byte(val)
	} else if name == "C" {
		s.BC[0] = byte(val)
	} else if name == "D" {
		s.DE[1] = byte(val)
	} else if name == "E" {
		s.DE[0] = byte(val)
	} else if name == "H" {
		s.HL[1] = byte(val)
	} else if name == "L" {
		s.HL[0] = byte(val)
	}
}

func (s *CPU) GetHLVal() int { // Get value pointed by HL register
	addr := 256*int(s.HL[1]) + int(s.HL[0])
	return int(s.Mem[addr])
}

func (s *CPU) SetHLVal(val int) { // Set value pointed by HL register
	addr := 256*int(s.HL[1]) + int(s.HL[0])
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
		return 256*int(s.AF[1]) + int(s.AF[0])
	} else if name == "BC" {
		return 256*int(s.BC[1]) + int(s.BC[0])
	} else if name == "DE" {
		return 256*int(s.DE[1]) + int(s.DE[0])
	} else if name == "HL" {
		return 256*int(s.HL[1]) + int(s.HL[0])
	} else if name == "SP" {
		return 256*int(s.SP[1]) + int(s.SP[0])
	} else if name == "PC" {
		return 256*int(s.PC[1]) + int(s.PC[0])
	}

	return -1
}

func (s *CPU) SetReg16Val(name string, val int) { // Set value of 16-bit register
	if name == "AF" {
		s.AF[0] = byte(val & 0x00FF)
		s.AF[1] = byte((val & 0xFF00) >> 8)
	} else if name == "BC" {
		s.BC[0] = byte(val & 0x00FF)
		s.BC[1] = byte((val & 0xFF00) >> 8)
	} else if name == "DE" {
		s.DE[0] = byte(val & 0x00FF)
		s.DE[1] = byte((val & 0xFF00) >> 8)
	} else if name == "HL" {
		s.HL[0] = byte(val & 0x00FF)
		s.HL[1] = byte((val & 0xFF00) >> 8)
	} else if name == "SP" {
		s.SP[0] = byte(val & 0x00FF)
		s.SP[1] = byte((val & 0xFF00) >> 8)
	} else if name == "PC" {
		s.PC[0] = byte(val & 0x00FF)
		s.PC[1] = byte((val & 0xFF00) >> 8)
	}
}

func (s *CPU) checkCC(cc string) bool {
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
