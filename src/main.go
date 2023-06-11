package main

import (
	// "github.com/shivkar2n/GB-Emulator/CPU"
	Dissasembler "github.com/shivkar2n/GB-Emulator/Disassembler"
)

func main() {
	Dissasembler.OpCodeTable()
	// c := CPU.InitCPU()
	// c.AF = [2]byte{byte(0x00), byte(0xAF)}
	// c.BC = [2]byte{byte(0x1E), byte(0x21)}
	// c.HL = [2]byte{byte(0x00), byte(0x00)}
	// c.SP = [2]byte{byte(0x13), byte(0x10)}
	// c.Mem[0x00] = byte(0xAF)
	// c.DebugInfo()
	// c.SLAR8("A")
	// c.DebugInfo()
	// fmt.Printf("%X", c.Mem[0x00])

}
