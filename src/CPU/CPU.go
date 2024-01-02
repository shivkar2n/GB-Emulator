package CPU

import (
	"fmt"

	"github.com/shivkar2n/GB-Emulator/MMU"
)

type CPU struct {
	Reg                   Register
	Mem                   MMU.MMU
	Sysclk, TotalT, Level int
	IME                   bool
	ClkRate, FrameRate    int
}

type Register struct {
	FA, CB, ED, LH, PC, SP [2]byte
}

func InitRegister() Register {
	return Register{
		FA: [2]byte{byte(0xB0), byte(0x01)},
		CB: [2]byte{byte(0x13), byte(0x00)},
		LH: [2]byte{byte(0x4D), byte(0x01)},
		SP: [2]byte{byte(0xFE), byte(0xFF)},
		PC: [2]byte{byte(0x00), byte(0x01)},
		ED: [2]byte{byte(0xD8), byte(0x00)},
	}
}

func (r *Register) Read(name string) int {
	if name == "A" {
		return int(r.FA[1])

	} else if name == "F" {
		return int(r.FA[0])

	} else if name == "B" {
		return int(r.CB[1])

	} else if name == "C" {
		return int(r.CB[0])

	} else if name == "D" {
		return int(r.ED[1])

	} else if name == "E" {
		return int(r.ED[0])

	} else if name == "H" {
		return int(r.LH[1])

	} else if name == "L" {
		return int(r.LH[0])

	} else if name == "AF" {
		return 256*int(r.FA[1]) + int(r.FA[0])

	} else if name == "BC" {
		return 256*int(r.CB[1]) + int(r.CB[0])

	} else if name == "DE" {
		return 256*int(r.ED[1]) + int(r.ED[0])

	} else if name == "HL" {
		return 256*int(r.LH[1]) + int(r.LH[0])

	} else if name == "SP" {
		return 256*int(r.SP[1]) + int(r.SP[0])

	} else if name == "PC" {
		return 256*int(r.PC[1]) + int(r.PC[0])
	}
	return 0
}

func (r *Register) Write(val int, name string) {
	if name == "A" {
		r.FA[1] = byte(val)

	} else if name == "F" {
		r.FA[0] = byte(val)

	} else if name == "B" {
		r.CB[1] = byte(val)

	} else if name == "C" {
		r.CB[0] = byte(val)

	} else if name == "D" {
		r.ED[1] = byte(val)

	} else if name == "E" {
		r.ED[0] = byte(val)

	} else if name == "H" {
		r.LH[1] = byte(val)

	} else if name == "L" {
		r.LH[0] = byte(val)

	} else if name == "AF" {
		r.FA[0] = byte(val & 0x00FF)
		r.FA[1] = byte((val & 0xFF00) >> 8)

	} else if name == "BC" {
		r.CB[0] = byte(val & 0x00FF)
		r.CB[1] = byte((val & 0xFF00) >> 8)

	} else if name == "DE" {
		r.ED[0] = byte(val & 0x00FF)
		r.ED[1] = byte((val & 0xFF00) >> 8)

	} else if name == "HL" {
		r.LH[0] = byte(val & 0x00FF)
		r.LH[1] = byte((val & 0xFF00) >> 8)

	} else if name == "SP" {
		r.SP[0] = byte(val & 0x00FF)
		r.SP[1] = byte((val & 0xFF00) >> 8)

	} else if name == "PC" {
		r.PC[0] = byte(val & 0x00FF)
		r.PC[1] = byte((val & 0xFF00) >> 8)
	}
}

func Init(mem MMU.MMU) *CPU {
	c := CPU{
		Reg:       InitRegister(),
		Mem:       mem,
		IME:       false,
		TotalT:    0,
		Sysclk:    0xABCC,
		Level:     0,
		ClkRate:   4194304,
		FrameRate: 60,
	}
	return &c
}

func (cpu *CPU) GetFlag(flagType string) int { // Get value of flag
	if flagType == "Z" {
		return cpu.Reg.Read("F") & 0x80 >> 7
	} else if flagType == "N" {
		return cpu.Reg.Read("F") & 0x40 >> 6
	} else if flagType == "H" {
		return cpu.Reg.Read("F") & 0x20 >> 5
	} else if flagType == "C" {
		return cpu.Reg.Read("F") & 0x10 >> 4
	}
	return -1
}

func (cpu *CPU) ResetFlag(flagType string) { // Check if flag bit set to 1, then change to 0
	if flagType == "Z" && (cpu.Reg.Read("F"))&0x80>>7 == 1 {
		val := cpu.Reg.Read("F") & 0x7F
		cpu.Reg.Write(val, "F")

	} else if flagType == "N" && (cpu.Reg.Read("F")&0x40)>>6 == 1 {
		val := cpu.Reg.Read("F") & 0xBF
		cpu.Reg.Write(val, "F")

	} else if flagType == "H" && (cpu.Reg.Read("F")&0x20)>>5 == 1 {
		val := cpu.Reg.Read("F") & 0xDF
		cpu.Reg.Write(val, "F")

	} else if flagType == "C" && (cpu.Reg.Read("F")&0x10)>>4 == 1 {
		val := cpu.Reg.Read("F") & 0xEF
		cpu.Reg.Write(val, "F")
	}
}

func (cpu *CPU) SetFlag(flagType string) { // Check if flag bit set to 0, then change to 1
	if flagType == "Z" && (cpu.Reg.Read("F")&0x80)>>7 == 0 {
		val := cpu.Reg.Read("F") ^ 0x80
		cpu.Reg.Write(val, "F")

	} else if flagType == "N" && (cpu.Reg.Read("F")&0x40)>>6 == 0 {
		val := cpu.Reg.Read("F") ^ 0x40
		cpu.Reg.Write(val, "F")

	} else if flagType == "H" && (cpu.Reg.Read("F")&0x20)>>5 == 0 {
		val := cpu.Reg.Read("F") ^ 0x20
		cpu.Reg.Write(val, "F")

	} else if flagType == "C" && (cpu.Reg.Read("F")&0x10)>>4 == 0 {
		val := cpu.Reg.Read("F") ^ 0x10
		cpu.Reg.Write(val, "F")
	}
}

func (cpu *CPU) StateInfo() { // Get info about state of CPU
	a := cpu.Reg.Read("A")
	f := cpu.Reg.Read("F")
	b := cpu.Reg.Read("B")
	c := cpu.Reg.Read("C")
	d := cpu.Reg.Read("D")
	e := cpu.Reg.Read("E")
	h := cpu.Reg.Read("H")
	l := cpu.Reg.Read("L")
	sp := cpu.Reg.Read("SP")
	pc := cpu.Reg.Read("PC")
	pcMem := cpu.Mem.Read(cpu.Reg.Read("PC"))
	pcMem1 := cpu.Mem.Read(cpu.Reg.Read("PC") + 1)
	pcMem2 := cpu.Mem.Read(cpu.Reg.Read("PC") + 2)
	pcMem3 := cpu.Mem.Read(cpu.Reg.Read("PC") + 3)

	// fmt.Printf("A: %02X F: %02X B: %02X C: %02X D: %02X E: %02X H: %02X L: %02X SP: %04X PC: 00:%04X (%02X %02X %02X %02X)\n", a, f, b, c, d, e, h, l, sp, pc, pcMem, pcMem1, pcMem2, pcMem3)
	fmt.Printf("A:%02X F:%02X B:%02X C:%02X D:%02X E:%02X H:%02X L:%02X SP:%04X PC:%04X PCMEM:%02X,%02X,%02X,%02X\n", a, f, b, c, d, e, h, l, sp, pc, pcMem, pcMem1, pcMem2, pcMem3)
}

func (cpu *CPU) GetHLVal() int { // Get value pointed by HL register
	return cpu.Mem.Read(cpu.Reg.Read("HL"))
}

func (cpu *CPU) SetHLVal(val int) { // Set value pointed by HL register
	cpu.Mem.Write(val, cpu.Reg.Read("HL"))
}

func (cpu *CPU) GetPCVal() int { // Get value pointed by PC register
	return cpu.Mem.Read(cpu.Reg.Read("PC"))
}

func (cpu *CPU) SetPCVal(val int) { // Set value pointed by PC register
	cpu.Mem.Write(val, cpu.Reg.Read("PC"))
}

func (cpu *CPU) checkCC(cc string) bool { // Check condition
	if cc == "Z" {
		return cpu.GetFlag("Z") == 1
	} else if cc == "NZ" {
		return cpu.GetFlag("Z") == 0
	} else if cc == "C" {
		return cpu.GetFlag("C") == 1
	} else if cc == "NC" {
		return cpu.GetFlag("C") == 0
	}
	return false
}

func (cpu *CPU) LogSerialIO() {
	if cpu.Mem.Read(0xFF02) == 0x81 {
		fmt.Printf("%c", cpu.Mem.Read(0xFF01))
		cpu.Mem.Write(0x00, 0xFF02)
	}
}

func (cpu *CPU) IncrCounter(val int) {
	cpu.Reg.Write(cpu.Reg.Read("PC")+val, "PC")
}
