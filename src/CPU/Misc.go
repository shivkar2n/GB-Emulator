package CPU

import "fmt"

func (cpu *CPU) EI() (string, int, int, bool) {
	opcode := fmt.Sprintf("EI")
	cpu.IME = true
	return opcode, 1, 4, true
}

func (cpu *CPU) DI() (string, int, int, bool) {
	opcode := fmt.Sprintf("DI")
	cpu.IME = false
	return opcode, 1, 4, true
}

func (cpu *CPU) CPL() (string, int, int, bool) { // Reg[A] = ~Reg[A]
	opcode := fmt.Sprintf("CPL")
	cpu.Reg.Write(^cpu.Reg.Read("A"), "A")
	cpu.SetFlag("N")
	cpu.SetFlag("H")
	return opcode, 1, 4, true
}

func (cpu *CPU) SCF() (string, int, int, bool) { // Flag[C] = true
	opcode := fmt.Sprintf("SCF")
	cpu.ResetFlag("N")
	cpu.ResetFlag("H")
	cpu.SetFlag("C")
	return opcode, 1, 4, true
}

func (cpu *CPU) CCF() (string, int, int, bool) { // Flag[C] = ~Flag[C]
	opcode := fmt.Sprintf("CCF")
	if cpu.GetFlag("C") == 1 {
		cpu.ResetFlag("C")
	} else {
		cpu.SetFlag("C")
	}
	cpu.ResetFlag("N")
	cpu.ResetFlag("H")
	return opcode, 1, 4, true
}

func (cpu *CPU) NOP() (string, int, int, bool) {
	opcode := fmt.Sprintf("NOP")
	return opcode, 1, 4, true
}

func (cpu *CPU) HALT() (string, int, int, bool) {
	opcode := fmt.Sprintf("HALT")
	return opcode, 1, 4, false
}
