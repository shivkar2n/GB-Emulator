package CPU

type CPU struct {
	Reg                   Register
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

func Init() *CPU {
	cpu := CPU{
		Reg:       InitRegister(),
		IME:       false,
		TotalT:    0,
		Sysclk:    0xABCC,
		Level:     0,
		ClkRate:   4194304,
		FrameRate: 60,
	}
	return &cpu
}

func (CPU *CPU) GetFlag(flagType string) int { // Get value of flag
	if flagType == "Z" {
		return CPU.Reg.Read("F") & 0x80 >> 7
	} else if flagType == "N" {
		return CPU.Reg.Read("F") & 0x40 >> 6
	} else if flagType == "H" {
		return CPU.Reg.Read("F") & 0x20 >> 5
	} else if flagType == "C" {
		return CPU.Reg.Read("F") & 0x10 >> 4
	}
	return -1
}

func (CPU *CPU) ResetFlag(flagType string) { // Check if flag bit set to 1, then change to 0
	if flagType == "Z" && (CPU.Reg.Read("F"))&0x80>>7 == 1 {
		val := CPU.Reg.Read("F") & 0x7F
		CPU.Reg.Write(val, "F")

	} else if flagType == "N" && (CPU.Reg.Read("F")&0x40)>>6 == 1 {
		val := CPU.Reg.Read("F") & 0xBF
		CPU.Reg.Write(val, "F")

	} else if flagType == "H" && (CPU.Reg.Read("F")&0x20)>>5 == 1 {
		val := CPU.Reg.Read("F") & 0xDF
		CPU.Reg.Write(val, "F")

	} else if flagType == "C" && (CPU.Reg.Read("F")&0x10)>>4 == 1 {
		val := CPU.Reg.Read("F") & 0xEF
		CPU.Reg.Write(val, "F")
	}
}

func (CPU *CPU) SetFlag(flagType string) { // Check if flag bit set to 0, then change to 1
	if flagType == "Z" && (CPU.Reg.Read("F")&0x80)>>7 == 0 {
		val := CPU.Reg.Read("F") ^ 0x80
		CPU.Reg.Write(val, "F")

	} else if flagType == "N" && (CPU.Reg.Read("F")&0x40)>>6 == 0 {
		val := CPU.Reg.Read("F") ^ 0x40
		CPU.Reg.Write(val, "F")

	} else if flagType == "H" && (CPU.Reg.Read("F")&0x20)>>5 == 0 {
		val := CPU.Reg.Read("F") ^ 0x20
		CPU.Reg.Write(val, "F")

	} else if flagType == "C" && (CPU.Reg.Read("F")&0x10)>>4 == 0 {
		val := CPU.Reg.Read("F") ^ 0x10
		CPU.Reg.Write(val, "F")
	}
}

func (CPU *CPU) CheckCC(cc string) bool { // Check condition
	if cc == "Z" {
		return CPU.GetFlag("Z") == 1
	} else if cc == "NZ" {
		return CPU.GetFlag("Z") == 0
	} else if cc == "C" {
		return CPU.GetFlag("C") == 1
	} else if cc == "NC" {
		return CPU.GetFlag("C") == 0
	}
	return false
}

func (CPU *CPU) IncrementCounter(val int) {
	CPU.Reg.Write(CPU.Reg.Read("PC")+val, "PC")
}
