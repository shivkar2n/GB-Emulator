package Dissasembler

import (
	"fmt"
	"os"
)

func OpCodeTable() {
	fileName := os.Args[1]
	rom, _ := os.ReadFile(fileName)

	for PC := 0; PC < len(rom); {
		switch uint8(rom[PC] & 0xF0) {

		case 0x00:
			switch rom[PC] {
			case 0x00:
				fmt.Printf("NOP\n")

			case 0x01:
				addr := uint16(rom[PC+1]) + (uint16(rom[PC+2]) << 8)
				fmt.Printf("LD BC,$%x \n", addr)
				PC += 3

			case 0x02:
				fmt.Printf("LD BC, A\n")
				PC += 1

			case 0x03:
				fmt.Printf("INC BC\n")
				PC += 1

			case 0x04:
				fmt.Printf("INC B\n")
				PC += 1

			case 0x05:
				fmt.Printf("DEC B\n")
				PC += 1

			case 0x06:
				addr := rom[PC+1]
				fmt.Printf("LD B,$%x\n", addr)
				PC += 2

			case 0x07:
				fmt.Printf("RCLA")
				PC += 1

			case 0x08:
				addr := uint16(rom[PC+1]) + (uint16(rom[PC+2]) << 8)
				fmt.Printf("LD (%x), SP\n", addr)
				PC += 3

			case 0x09:
				fmt.Printf("ADD HL, BC\n")
				PC += 1

			case 0x0A:
				fmt.Printf("LD A, (BC)\n")
				PC += 1

			case 0x0B:
				fmt.Printf("DEC BC\n")
				PC += 1

			case 0x0C:
				fmt.Printf("INC C\n")
				PC += 1

			case 0x0D:
				fmt.Printf("DEC C\n")
				PC += 1

			case 0x0E:
				addr := rom[PC+1]
				fmt.Printf("LD C, $%x\n", addr)
				PC += 2

			case 0x0F:
				fmt.Printf("RRCA \n")
				PC += 1

			}

		case 0x10:
			switch rom[PC] {
			case 0x10:
				addr := rom[PC+1]
				fmt.Printf("STOP $%x\n", addr)
				PC += 2

			case 0x11:
				addr := uint16(rom[PC+1]) + (uint16(rom[PC+2]) << 8)
				fmt.Printf("LD DE,$%x\n", addr)
				PC += 3

			case 0x12:
				fmt.Printf("LD (DE),A\n")
				PC += 1

			case 0x13:
				fmt.Printf("INC (DE)\n")
				PC += 1

			case 0x14:
				fmt.Printf("INC D\n")
				PC += 1

			case 0x15:
				fmt.Printf("DEC D\n")
				PC += 1

			case 0x16:
				addr := rom[PC+1]
				fmt.Printf("LD D,$%x\n", addr)
				PC += 2

			case 0x17:
				fmt.Printf("RLA\n")
				PC += 1

			case 0x18:
				addr := int8(rom[PC+1])
				fmt.Printf("JR r%x\n", addr)
				PC += 2

			case 0x19:
				fmt.Printf("ADD HL, DE \n")
				PC += 1

			case 0x1A:
				fmt.Printf("LD A, (DE)\n")
				PC += 1

			case 0x1B:
				fmt.Printf("DEC DE\n")
				PC += 1

			case 0x1C:
				fmt.Printf("INC E\n")
				PC += 1

			case 0x1D:
				fmt.Printf("DEC E\n")
				PC += 1

			case 0x1E:
				addr := rom[PC+1]
				fmt.Printf("LD E,$%x\n", addr)
				PC += 2

			case 0x1F:
				fmt.Printf("RRA\n")
				PC += 1
			}

		case 0x20:
			switch rom[PC] {
			case 0x20:
				addr := rom[PC+1]
				fmt.Printf("JR NZ,r%x\n", addr)
				PC += 2

			case 0x21:
				addr := uint16(rom[PC+1]) + (uint16(rom[PC+2]) << 8)
				fmt.Printf("LD HL,$%x\n", addr)
				PC += 3

			case 0x22:
				fmt.Printf("LD (HL+),A\n")
				PC += 1

			case 0x23:
				fmt.Printf("INC HL\n")
				PC += 1

			case 0x24:
				fmt.Printf("INC H\n")
				PC += 1

			case 0x25:
				fmt.Printf("DEC H\n")
				PC += 1

			case 0x26:
				addr := rom[PC+1]
				fmt.Printf("LD H,$%x\n", addr)
				PC += 2

			case 0x27:
				fmt.Printf("DAA\n")
				PC += 1

			case 0x28:
				addr := int8(rom[PC+1])
				fmt.Printf("JR Z,r%x\n", addr)
				PC += 2

			case 0x29:
				fmt.Printf("ADD HL,HL\n")
				PC += 1

			case 0x2A:
				fmt.Printf("LD A,(HL+)\n")
				PC += 1

			case 0x2B:
				fmt.Printf("DEC HL\n")
				PC += 1

			case 0x2C:
				fmt.Printf("INC L\n")
				PC += 1

			case 0x2D:
				fmt.Printf("DEC L\n")
				PC += 1

			case 0x2E:
				addr := rom[PC+1]
				fmt.Printf("LD L,$%x\n", addr)
				PC += 2

			case 0x2F:
				fmt.Printf("CPL\n")
				PC += 1
			}

		case 0x30:
			switch rom[PC] {
			case 0x30:
				addr := rom[PC+1]
				fmt.Printf("JR NC,r%x\n", addr)
				PC += 2

			case 0x31:
				addr := uint16(rom[PC+1]) + (uint16(rom[PC+2]) << 8)
				fmt.Printf("LD SP $%x\n", addr)
				PC += 3

			case 0x32:
				fmt.Printf("LD (HL-),A\n")
				PC += 1

			case 0x33:
				fmt.Printf("INC SP\n")
				PC += 1

			case 0x34:
				fmt.Printf("INC (HL)\n")
				PC += 1

			case 0x35:
				fmt.Printf("DEC (HL)\n")
				PC += 1

			case 0x36:
				addr := rom[PC+1]
				fmt.Printf("LD (HL),$%x\n", addr)
				PC += 2

			case 0x37:
				fmt.Printf("SCF\n")
				PC += 1

			case 0x38:
				addr := int8(rom[PC+1])
				fmt.Printf("JR C,r%x\n", addr)
				PC += 2

			case 0x39:
				fmt.Printf("ADD HL,SP\n")
				PC += 1

			case 0x3A:
				fmt.Printf("LD A,(HL-)\n")
				PC += 1

			case 0x3B:
				fmt.Printf("DEC SP\n")
				PC += 1

			case 0x3C:
				fmt.Printf("INC A\n")
				PC += 1

			case 0x3D:
				fmt.Printf("DEC A\n")
				PC += 1

			case 0x3E:
				addr := rom[PC+1]
				fmt.Printf("LD A,$%x\n", addr)
				PC += 2

			case 0x3F:
				fmt.Printf("CCF\n")
				PC += 1
			}

		case 0x40:
			if (rom[PC] & 0x0F >> 3) < 1 {
				fmt.Printf("LD B,%s\n", reg[rom[PC]&0x0F])
			} else {
				fmt.Printf("LD C,%s\n", reg[(rom[PC]&0x0F)&^0x08])
			}
			PC += 1

		case 0x50:
			if (rom[PC] & 0x0F >> 3) < 1 {
				fmt.Printf("LD D,%s\n", reg[rom[PC]&0x0F])
			} else {
				fmt.Printf("LD E,%s\n", reg[(rom[PC]&0x0F)&^0x08])
			}
			PC += 1

		case 0x60:
			if (rom[PC] & 0x0F >> 3) < 1 {
				fmt.Printf("LD H,%s\n", reg[rom[PC]&0x0F])
			} else {
				fmt.Printf("LD L,%s\n", reg[(rom[PC]&0x0F)&^0x08])
			}
			PC += 1

		case 0x70:
			switch rom[PC] {
			case 0x70:
				fmt.Printf("LD (HL),B\n")
				PC += 1

			case 0x71:
				fmt.Printf("LD (HL),C\n")
				PC += 1

			case 0x72:
				fmt.Printf("LD (HL),D\n")
				PC += 1

			case 0x73:
				fmt.Printf("LD (HL),E\n")
				PC += 1

			case 0x74:
				fmt.Printf("LD (HL),H\n")
				PC += 1

			case 0x75:
				fmt.Printf("LD (HL),L\n")
				PC += 1

			case 0x76:
				fmt.Printf("HALT\n")
				PC += 1

			case 0x77:
				fmt.Printf("LD (HL),A\n")
				PC += 1

			case 0x78:
				fmt.Printf("LD A,B\n")
				PC += 1

			case 0x79:
				fmt.Printf("LD A,C\n")
				PC += 1

			case 0x7A:
				fmt.Printf("LD A,D\n")
				PC += 1

			case 0x7B:
				fmt.Printf("LD A,E\n")
				PC += 1

			case 0x7C:
				fmt.Printf("LD A,H\n")
				PC += 1

			case 0x7D:
				fmt.Printf("LD A,L\n")
				PC += 1

			case 0x7E:
				fmt.Printf("LD A,(HL)\n")
				PC += 1

			case 0x7F:
				fmt.Printf("LD A,A\n")
				PC += 1
			}

		case 0x80:
			if (rom[PC] & 0x0F >> 3) < 1 {
				fmt.Printf("ADD A,%s\n", reg[rom[PC]&0x0F])
			} else {
				fmt.Printf("ADC A,%s\n", reg[(rom[PC]&0x0F)&^0x08])
			}
			PC += 1

		case 0x90:
			if (rom[PC] & 0x0F >> 3) < 1 {
				fmt.Printf("SUB %s\n", reg[rom[PC]&0x0F])
			} else {
				fmt.Printf("SUBC %s\n", reg[(rom[PC]&0x0F)&^0x08])
			}
			PC += 1

		case 0xA0:
			if (rom[PC] & 0x0F >> 3) < 1 {
				fmt.Printf("AND %s\n", reg[rom[PC]&0x0F])
			} else {
				fmt.Printf("XOR %s\n", reg[(rom[PC]&0x0F)&^0x08])
			}
			PC += 1

		case 0xB0:
			if (rom[PC] & 0x0F >> 3) < 1 {
				fmt.Printf("OR %s\n", reg[rom[PC]&0x0F])
			} else {
				fmt.Printf("CP %s\n", reg[(rom[PC]&0x0F)&^0x08])
			}
			PC += 1

		case 0xC0:
			switch rom[PC] {
			case 0xC0:
				fmt.Printf("RET NZ\n")
				PC += 1

			case 0xC1:
				fmt.Printf("POP BC\n")
				PC += 1

			case 0xC2:
				addr := uint16(rom[PC+1]) + (uint16(rom[PC+2]) << 8)
				fmt.Printf("JP NZ,$%x \n", addr)
				PC += 3

			case 0xC3:
				addr := uint16(rom[PC+1]) + (uint16(rom[PC+2]) << 8)
				fmt.Printf("JP $%x \n", addr)
				PC += 3

			case 0xC4:
				addr := uint16(rom[PC+1]) + (uint16(rom[PC+2]) << 8)
				fmt.Printf("CALL NZ,$%x \n", addr)
				PC += 3

			case 0xC5:
				fmt.Printf("PUSH BC\n")
				PC += 1

			case 0xC6:
				addr := rom[PC+1]
				fmt.Printf("ADD A,$%x\n", addr)
				PC += 2

			case 0xC7:
				fmt.Printf("RST 00H\n")
				PC += 1

			case 0xC8:
				fmt.Printf("RET Z\n")
				PC += 1

			case 0xC9:
				fmt.Printf("RET\n")
				PC += 1

			case 0xCA:
				addr := uint16(rom[PC+1]) + (uint16(rom[PC+2]) << 8)
				fmt.Printf("JP Z,$%x\n", addr)
				PC += 3

			case 0xCB:
				PrefixTable(&PC, rom)
				PC += 2

			case 0xCC:
				addr := uint16(rom[PC+1]) + (uint16(rom[PC+2]) << 8)
				fmt.Printf("CALL Z,$%x\n", addr)
				PC += 3

			case 0xCD:
				addr := uint16(rom[PC+1]) + (uint16(rom[PC+2]) << 8)
				fmt.Printf("CALL $%x\n", addr)
				PC += 3

			case 0xCE:
				addr := rom[PC+1]
				fmt.Printf("ADC A,$%x\n", addr)
				PC += 2

			case 0xCF:
				fmt.Printf("RST 08H\n")
				PC += 1

			}

		case 0xD0:
			switch rom[PC] {
			case 0xD0:
				fmt.Printf("RET NC\n")
				PC += 1

			case 0xD1:
				fmt.Printf("POP DE\n")
				PC += 1

			case 0xD2:
				addr := uint16(rom[PC+1]) + (uint16(rom[PC+2]) << 8)
				fmt.Printf("JP NC,$%x \n", addr)
				PC += 3

			case 0xD3:
				PC += 1

			case 0xD4:
				addr := uint16(rom[PC+1]) + (uint16(rom[PC+2]) << 8)
				fmt.Printf("CALL NC,$%x \n", addr)
				PC += 3

			case 0xD5:
				fmt.Printf("PUSH DE\n")
				PC += 1

			case 0xD6:
				addr := rom[PC+1]
				fmt.Printf("SUB,$%x\n", addr)
				PC += 2

			case 0xD7:
				fmt.Printf("RST 10H\n")
				PC += 1

			case 0xD8:
				fmt.Printf("RET C\n")
				PC += 1

			case 0xD9:
				fmt.Printf("RETI\n")
				PC += 1

			case 0xDA:
				addr := uint16(rom[PC+1]) + (uint16(rom[PC+2]) << 8)
				fmt.Printf("JP C,$%x\n", addr)
				PC += 3

			case 0xDB:
				PC += 1

			case 0xDC:
				addr := uint16(rom[PC+1]) + (uint16(rom[PC+2]) << 8)
				fmt.Printf("CALL C,$%x\n", addr)
				PC += 3

			case 0xDD:
				PC += 1

			case 0xDE:
				addr := rom[PC+1]
				fmt.Printf("SBC A,$%x\n", addr)
				PC += 2

			case 0xDF:
				fmt.Printf("RST 18H\n")
				PC += 1

			}

		case 0xE0:
			switch rom[PC] {
			case 0xE0:
				addr := rom[PC+1]
				fmt.Printf("LD ($FFOO+$%x),A\n", addr)
				PC += 2

			case 0xE1:
				fmt.Printf("POP HL\n")
				PC += 1

			case 0xE2:
				fmt.Printf("LD ($FF00+C),A\n")
				PC += 1

			case 0xE3:
				PC += 1

			case 0xE4:
				PC += 1

			case 0xE5:
				fmt.Printf("PUSH HL\n")
				PC += 1

			case 0xE6:
				addr := rom[PC+1]
				fmt.Printf("AND,$%x\n", addr)
				PC += 2

			case 0xE7:
				fmt.Printf("RST 20H\n")
				PC += 1

			case 0xE8:
				addr := rom[PC+1]
				fmt.Printf("ADD SP,r%x\n", addr)
				PC += 2

			case 0xE9:
				fmt.Printf("JP HL\n")
				PC += 1

			case 0xEA:
				addr := uint16(rom[PC+1]) + (uint16(rom[PC+2]) << 8)
				fmt.Printf("LD ($%x),A\n", addr)
				PC += 3

			case 0xEB:
				PC += 1

			case 0xEC:
				PC += 1

			case 0xED:
				PC += 1

			case 0xEE:
				addr := rom[PC+1]
				fmt.Printf("XOR $%x\n", addr)
				PC += 2

			case 0xEF:
				fmt.Printf("RST 28H\n")
				PC += 1
			}

		case 0xF0:
			switch rom[PC] {
			case 0xF0:
				addr := rom[PC+1]
				fmt.Printf("LD A,($FF00+%x)\n", addr)
				PC += 2

			case 0xF1:
				fmt.Printf("POP AF\n")
				PC += 1

			case 0xF2:
				fmt.Printf("LD A,(C)\n")
				PC += 1

			case 0xF3:
				fmt.Printf("DI\n")
				PC += 1

			case 0xF4:
				PC += 1

			case 0xF5:
				fmt.Printf("PUSH AF\n")
				PC += 1

			case 0xF6:
				addr := rom[PC+1]
				fmt.Printf("OR,$%x\n", addr)
				PC += 2

			case 0xF7:
				fmt.Printf("RST 30H\n")
				PC += 1

			case 0xF8:
				addr := rom[PC+1]
				fmt.Printf("LD HL,SP+r%x\n", addr)
				PC += 2

			case 0xF9:
				fmt.Printf("LD SP,HL\n")
				PC += 1

			case 0xFA:
				addr := uint16(rom[PC+1]) + (uint16(rom[PC+2]) << 8)
				fmt.Printf("LD A,($%x)\n", addr)
				PC += 3

			case 0xFB:
				fmt.Printf("EI\n")
				PC += 1

			case 0xFC:
				PC += 1

			case 0xFD:
				PC += 1

			case 0xFE:
				addr := rom[PC+1]
				fmt.Printf("CP $%x\n", addr)
				PC += 2

			case 0xFF:
				fmt.Printf("RST 38H\n")
				PC += 1
			}
		default:
			break
		}
	}

}
