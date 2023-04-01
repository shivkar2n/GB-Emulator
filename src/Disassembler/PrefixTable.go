package Dissasembler

import (
	"fmt"
)

var reg = [8]string{"B", "C", "D", "E", "H", "L", "(HL)", "A"}

func PrefixTable(PC *int, rom []byte) {
	opCode := uint8(rom[*PC+1])

	switch opCode & 0xF0 {
	case 0x00:
		if (opCode & 0x0F >> 3) < 1 {
			fmt.Printf("RLC %s\n", reg[opCode&0x0F])
		} else {
			fmt.Printf("RRC %s\n", reg[(opCode&0x0F)&^0x08])
		}

	case 0x10:
		if (opCode & 0x0F >> 3) < 1 {
			fmt.Printf("RL %s\n", reg[opCode&0x0F])
		} else {
			fmt.Printf("RR %s\n", reg[(opCode&0x0F)&^0x08])
		}

	case 0x20:
		if (opCode & 0x0F >> 3) < 1 {
			fmt.Printf("SLA %s\n", reg[opCode&0x0F])
		} else {
			fmt.Printf("SRA %s\n", reg[(opCode&0x0F)&^0x08])
		}

	case 0x30:
		if (opCode & 0x0F >> 3) < 1 {
			fmt.Printf("SWAP %s\n", reg[opCode&0x0F])
		} else {
			fmt.Printf("SRL %s\n", reg[(opCode&0x0F)&^0x08])
		}

	case 0x40:
		if (opCode & 0x0F >> 3) < 1 {
			fmt.Printf("BIT 0,%s\n", reg[opCode&0x0F])
		} else {
			fmt.Printf("BIT 1,%s\n", reg[(opCode&0x0F)&^0x08])
		}

	case 0x50:
		if (opCode & 0x0F >> 3) < 1 {
			fmt.Printf("BIT 2,%s\n", reg[opCode&0x0F])
		} else {
			fmt.Printf("BIT 3,%s\n", reg[(opCode&0x0F)&^0x08])
		}

	case 0x60:
		if (opCode & 0x0F >> 3) < 1 {
			fmt.Printf("BIT 4,%s\n", reg[opCode&0x0F])
		} else {
			fmt.Printf("BIT 5,%s\n", reg[(opCode&0x0F)&^0x08])
		}

	case 0x70:
		if (opCode & 0x0F >> 3) < 1 {
			fmt.Printf("BIT 6,%s\n", reg[opCode&0x0F])
		} else {
			fmt.Printf("BIT 7,%s\n", reg[(opCode&0x0F)&^0x08])
		}

	case 0x80:
		if (opCode & 0x0F >> 3) < 1 {
			fmt.Printf("RES 0,%s\n", reg[opCode&0x0F])
		} else {
			fmt.Printf("RES 1,%s\n", reg[(opCode&0x0F)&^0x08])
		}

	case 0x90:
		if (opCode & 0x0F >> 3) < 1 {
			fmt.Printf("RES 2,%s\n", reg[opCode&0x0F])
		} else {
			fmt.Printf("RES 3,%s\n", reg[(opCode&0x0F)&^0x08])
		}

	case 0xA0:
		if (opCode & 0x0F >> 3) < 1 {
			fmt.Printf("RES 4,%s\n", reg[opCode&0x0F])
		} else {
			fmt.Printf("RES 5,%s\n", reg[(opCode&0x0F)&^0x08])
		}

	case 0xB0:
		if (opCode & 0x0F >> 3) < 1 {
			fmt.Printf("RES 6,%s\n", reg[opCode&0x0F])
		} else {
			fmt.Printf("RES 7,%s\n", reg[(opCode&0x0F)&^0x08])
		}

	case 0xC0:
		if (opCode & 0x0F >> 3) < 1 {
			fmt.Printf("SET 0,%s\n", reg[opCode&0x0F])
		} else {
			fmt.Printf("SET 1,%s\n", reg[(opCode&0x0F)&^0x08])
		}

	case 0xD0:
		if (opCode & 0x0F >> 3) < 1 {
			fmt.Printf("SET 2,%s\n", reg[opCode&0x0F])
		} else {
			fmt.Printf("SET 3,%s\n", reg[(opCode&0x0F)&^0x08])
		}

	case 0xE0:
		if (opCode & 0x0F >> 3) < 1 {
			fmt.Printf("SET 4,%s\n", reg[opCode&0x0F])
		} else {
			fmt.Printf("SET 5,%s\n", reg[(opCode&0x0F)&^0x08])
		}

	case 0xF0:
		if (opCode & 0x0F >> 3) < 1 {
			fmt.Printf("SET 6,%s\n", reg[opCode&0x0F])
		} else {
			fmt.Printf("SET 7,%s\n", reg[(opCode&0x0F)&^0x08])
		}
	}
}
