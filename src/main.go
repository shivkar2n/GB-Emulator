package main

import (
	CPU "github.com/shivkar2n/GB-Emulator/CPU"
	Disasembler "github.com/shivkar2n/GB-Emulator/Disassembler"
)

func main() {
	// State of CPU after executing BOOT ROM
	c := CPU.InitCPU()
	c.FA = [2]byte{byte(0xB0), byte(0x01)}
	c.CB = [2]byte{byte(0x13), byte(0x00)}
	c.LH = [2]byte{byte(0x4D), byte(0x01)}
	c.SP = [2]byte{byte(0xFE), byte(0xFF)}
	c.PC = [2]byte{byte(0x00), byte(0x01)}
	c.ED = [2]byte{byte(0xD8), byte(0x00)}
	c.Mem[0xFF44] = byte(0x90)

	Disasembler.OpCodeTable(c)
}
