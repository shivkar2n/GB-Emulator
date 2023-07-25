package Disassembler

import (
	"log"

	CPU "github.com/shivkar2n/GB-Emulator/CPU"
)

func PrefixTable(c *CPU.CPU, log *log.Logger) {
	opCode := uint8(c.Mem[c.GetReg16Val("PC")+1])

	switch opCode & 0xF0 {
	case 0x00:
		switch opCode {
		case 0x00:
			c.RLCR8("B")
			log.Printf("RLC B\n")

		case 0x01:
			c.RLCR8("C")
			log.Printf("RLC C\n")

		case 0x02:
			c.RLCR8("D")
			log.Printf("RLC D\n")

		case 0x03:
			c.RLCR8("E")
			log.Printf("RLC E\n")

		case 0x04:
			c.RLCR8("H")
			log.Printf("RLC H\n")

		case 0x05:
			c.RLCR8("L")
			log.Printf("RLC L\n")

		case 0x06:
			c.RLCHL()
			log.Printf("RLC (HL)\n")

		case 0x07:
			c.RLCR8("A")
			log.Printf("RLC A\n")

		case 0x08:
			c.RRCR8("B")
			log.Printf("RRC B\n")

		case 0x09:
			c.RRCR8("C")
			log.Printf("RRC C\n")

		case 0x0A:
			c.RRCR8("D")
			log.Printf("RRC D\n")

		case 0x0B:
			c.RRCR8("E")
			log.Printf("RRC E\n")

		case 0x0C:
			c.RRCR8("H")
			log.Printf("RRC H\n")

		case 0x0D:
			c.RRCR8("L")
			log.Printf("RRC L\n")

		case 0x0E:
			c.RRCHL()
			log.Printf("RRC (HL)\n")

		case 0x0F:
			c.RRCR8("A")
			log.Printf("RRC A\n")
		}

	case 0x10:
		switch opCode {
		case 0x10:
			c.RLR8("B")
			log.Printf("RL B\n")

		case 0x11:
			c.RLR8("C")
			log.Printf("RL C\n")

		case 0x12:
			c.RLR8("D")
			log.Printf("RL D\n")

		case 0x13:
			c.RLR8("E")
			log.Printf("RL E\n")

		case 0x14:
			c.RLR8("H")
			log.Printf("RL H\n")

		case 0x15:
			c.RLR8("L")
			log.Printf("RL L\n")

		case 0x16:
			c.RLHL()
			log.Printf("RL (HL)\n")

		case 0x17:
			c.RLR8("A")
			log.Printf("RL A\n")

		case 0x18:
			c.RRR8("B")
			log.Printf("RR B\n")

		case 0x19:
			c.RRR8("C")
			log.Printf("RR C\n")

		case 0x1A:
			c.RRR8("D")
			log.Printf("RR D\n")

		case 0x1B:
			c.RRR8("E")
			log.Printf("RR E\n")

		case 0x1C:
			c.RRR8("H")
			log.Printf("RR H\n")

		case 0x1D:
			c.RRR8("L")
			log.Printf("RR L\n")

		case 0x1E:
			c.RRHL()
			log.Printf("RR (HL)\n")

		case 0x1F:
			c.RRR8("A")
			log.Printf("RR A\n")
		}

	case 0x20:
		switch opCode {
		case 0x20:
			c.SLAR8("B")
			log.Printf("SLA B\n")

		case 0x21:
			c.SLAR8("C")
			log.Printf("SLA C\n")

		case 0x22:
			c.SLAR8("D")
			log.Printf("SLA D\n")

		case 0x23:
			c.SLAR8("E")
			log.Printf("SLA E\n")

		case 0x24:
			c.SLAR8("H")
			log.Printf("SLA H\n")

		case 0x25:
			c.SLAR8("L")
			log.Printf("SLA L\n")

		case 0x26:
			c.SLAHL()
			log.Printf("SLA (HL)\n")

		case 0x27:
			c.SLAR8("A")
			log.Printf("SLA A\n")

		case 0x28:
			c.SRAR8("B")
			log.Printf("SRA B\n")

		case 0x29:
			c.SRAR8("C")
			log.Printf("SRA C\n")

		case 0x2A:
			c.SRAR8("D")
			log.Printf("SRA D\n")

		case 0x2B:
			c.SRAR8("E")
			log.Printf("SRA E\n")

		case 0x2C:
			c.SRAR8("H")
			log.Printf("SRA H\n")

		case 0x2D:
			c.SRAR8("L")
			log.Printf("SRA L\n")

		case 0x2E:
			c.SRAHL()
			log.Printf("SRA (HL)\n")

		case 0x2F:
			c.SRAR8("A")
			log.Printf("SRA A\n")
		}

	case 0x30:
		switch opCode {
		case 0x30:
			c.SWAPR8("B")
			log.Printf("SWAP B\n")

		case 0x31:
			c.SWAPR8("C")
			log.Printf("SWAP C\n")

		case 0x32:
			c.SWAPR8("D")
			log.Printf("SWAP D\n")

		case 0x33:
			c.SWAPR8("E")
			log.Printf("SWAP E\n")

		case 0x34:
			c.SWAPR8("H")
			log.Printf("SWAP H\n")

		case 0x35:
			c.SWAPR8("L")
			log.Printf("SWAP L\n")

		case 0x36:
			c.SWAPHL()
			log.Printf("SWAP (HL)\n")

		case 0x37:
			c.SWAPR8("A")
			log.Printf("SWAP A\n")

		case 0x38:
			c.SRLR8("B")
			log.Printf("SRA B\n")

		case 0x39:
			c.SRLR8("C")
			log.Printf("SRA C\n")

		case 0x3A:
			c.SRLR8("D")
			log.Printf("SRA D\n")

		case 0x3B:
			c.SRLR8("E")
			log.Printf("SRA E\n")

		case 0x3C:
			c.SRLR8("H")
			log.Printf("SRA H\n")

		case 0x3D:
			c.SRLR8("L")
			log.Printf("SRA L\n")

		case 0x3E:
			c.SRLHL()
			log.Printf("SRA (HL)\n")

		case 0x3F:
			c.SRLR8("A")
			log.Printf("SRA A\n")
		}

	case 0x40:
		switch opCode {
		case 0x40:
			c.BITU3R8(0, "B")
			log.Printf("BIT 0,B\n")

		case 0x41:
			c.BITU3R8(0, "C")
			log.Printf("BIT 0,C\n")

		case 0x42:
			c.BITU3R8(0, "D")
			log.Printf("BIT 0,D\n")

		case 0x43:
			c.BITU3R8(0, "E")
			log.Printf("BIT 0,E\n")

		case 0x44:
			c.BITU3R8(0, "H")
			log.Printf("BIT 0,H\n")

		case 0x45:
			c.BITU3R8(0, "L")
			log.Printf("BIT 0,L\n")

		case 0x46:
			c.BITU3HL(0)
			log.Printf("BIT 0,(HL)\n")

		case 0x47:
			c.BITU3R8(0, "A")
			log.Printf("BIT 0,A\n")

		case 0x48:
			c.BITU3R8(1, "B")
			log.Printf("BIT 1,B\n")

		case 0x49:
			c.BITU3R8(1, "C")
			log.Printf("BIT 1,C\n")

		case 0x4A:
			c.BITU3R8(1, "D")
			log.Printf("BIT 1,D\n")

		case 0x4B:
			c.BITU3R8(1, "E")
			log.Printf("BIT 1,E\n")

		case 0x4C:
			c.BITU3R8(1, "H")
			log.Printf("BIT 1,H\n")

		case 0x4D:
			c.BITU3R8(1, "L")
			log.Printf("BIT 1,L\n")

		case 0x4E:
			c.BITU3HL(1)
			log.Printf("BIT 1,(HL)\n")

		case 0x4F:
			c.BITU3R8(1, "A")
			log.Printf("BIT 1,A\n")

		}

	case 0x50:
		switch opCode {
		case 0x50:
			c.BITU3R8(2, "B")
			log.Printf("BIT 2,B\n")

		case 0x51:
			c.BITU3R8(2, "C")
			log.Printf("BIT 2,C\n")

		case 0x52:
			c.BITU3R8(2, "D")
			log.Printf("BIT 2,D\n")

		case 0x53:
			c.BITU3R8(2, "E")
			log.Printf("BIT 2,E\n")

		case 0x54:
			c.BITU3R8(2, "H")
			log.Printf("BIT 2,H\n")

		case 0x55:
			c.BITU3R8(2, "L")
			log.Printf("BIT 2,L\n")

		case 0x56:
			c.BITU3HL(2)
			log.Printf("BIT 2,(HL)\n")

		case 0x57:
			c.BITU3R8(2, "A")
			log.Printf("BIT 2,A\n")

		case 0x58:
			c.BITU3R8(3, "B")
			log.Printf("BIT 3,B\n")

		case 0x59:
			c.BITU3R8(3, "C")
			log.Printf("BIT 3,C\n")

		case 0x5A:
			c.BITU3R8(3, "D")
			log.Printf("BIT 3,D\n")

		case 0x5B:
			c.BITU3R8(3, "E")
			log.Printf("BIT 3,E\n")

		case 0x5C:
			c.BITU3R8(3, "H")
			log.Printf("BIT 3,H\n")

		case 0x5D:
			c.BITU3R8(3, "L")
			log.Printf("BIT 3,L\n")

		case 0x5E:
			c.BITU3HL(3)
			log.Printf("BIT 3,(HL)\n")

		case 0x5F:
			c.BITU3R8(3, "A")
			log.Printf("BIT 3,A\n")

		}

	case 0x60:
		switch opCode {
		case 0x60:
			c.BITU3R8(4, "B")
			log.Printf("BIT 4,B\n")

		case 0x61:
			c.BITU3R8(4, "C")
			log.Printf("BIT 4,C\n")

		case 0x62:
			c.BITU3R8(4, "D")
			log.Printf("BIT 4,D\n")

		case 0x63:
			c.BITU3R8(4, "E")
			log.Printf("BIT 4,E\n")

		case 0x64:
			c.BITU3R8(4, "H")
			log.Printf("BIT 4,H\n")

		case 0x65:
			c.BITU3R8(4, "L")
			log.Printf("BIT 4,L\n")

		case 0x66:
			c.BITU3HL(4)
			log.Printf("BIT 4,(HL)\n")

		case 0x67:
			c.BITU3R8(4, "A")
			log.Printf("BIT 4,A\n")

		case 0x68:
			c.BITU3R8(5, "B")
			log.Printf("BIT 5,B\n")

		case 0x69:
			c.BITU3R8(5, "C")
			log.Printf("BIT 5,C\n")

		case 0x6A:
			c.BITU3R8(5, "D")
			log.Printf("BIT 5,D\n")

		case 0x6B:
			c.BITU3R8(5, "E")
			log.Printf("BIT 5,E\n")

		case 0x6C:
			c.BITU3R8(5, "H")
			log.Printf("BIT 5,H\n")

		case 0x6D:
			c.BITU3R8(5, "L")
			log.Printf("BIT 5,L\n")

		case 0x6E:
			c.BITU3HL(5)
			log.Printf("BIT 5,(HL)\n")

		case 0x6F:
			c.BITU3R8(5, "A")
			log.Printf("BIT 5,A\n")

		}

	case 0x70:
		switch opCode {
		case 0x70:
			c.BITU3R8(6, "B")
			log.Printf("BIT 6,B\n")

		case 0x71:
			c.BITU3R8(6, "C")
			log.Printf("BIT 6,C\n")

		case 0x72:
			c.BITU3R8(6, "D")
			log.Printf("BIT 6,D\n")

		case 0x73:
			c.BITU3R8(6, "E")
			log.Printf("BIT 6,E\n")

		case 0x74:
			c.BITU3R8(6, "H")
			log.Printf("BIT 6,H\n")

		case 0x75:
			c.BITU3R8(6, "L")
			log.Printf("BIT 6,L\n")

		case 0x76:
			c.BITU3HL(6)
			log.Printf("BIT 6,(HL)\n")

		case 0x77:
			c.BITU3R8(6, "A")
			log.Printf("BIT 6,A\n")

		case 0x78:
			c.BITU3R8(7, "B")
			log.Printf("BIT 7,B\n")

		case 0x79:
			c.BITU3R8(7, "C")
			log.Printf("BIT 7,C\n")

		case 0x7A:
			c.BITU3R8(7, "D")
			log.Printf("BIT 7,D\n")

		case 0x7B:
			c.BITU3R8(7, "E")
			log.Printf("BIT 7,E\n")

		case 0x7C:
			c.BITU3R8(7, "H")
			log.Printf("BIT 7,H\n")

		case 0x7D:
			c.BITU3R8(7, "L")
			log.Printf("BIT 7,L\n")

		case 0x7E:
			c.BITU3HL(7)
			log.Printf("BIT 7,(HL)\n")

		case 0x7F:
			c.BITU3R8(7, "A")
			log.Printf("BIT 7,A\n")

		}

	case 0x80:
		switch opCode {
		case 0x80:
			c.RESU3R8(0, "B")
			log.Printf("RES 0,B\n")

		case 0x81:
			c.RESU3R8(0, "C")
			log.Printf("RES 0,C\n")

		case 0x82:
			c.RESU3R8(0, "D")
			log.Printf("RES 0,D\n")

		case 0x83:
			c.RESU3R8(0, "E")
			log.Printf("RES 0,E\n")

		case 0x84:
			c.RESU3R8(0, "H")
			log.Printf("RES 0,H\n")

		case 0x85:
			c.RESU3R8(0, "L")
			log.Printf("RES 0,L\n")

		case 0x86:
			c.RESU3HL(0)
			log.Printf("RES 0,(HL)\n")

		case 0x87:
			c.RESU3R8(0, "A")
			log.Printf("RES 0,A\n")

		case 0x88:
			c.RESU3R8(1, "B")
			log.Printf("RES 1,B\n")

		case 0x89:
			c.RESU3R8(1, "C")
			log.Printf("RES 1,C\n")

		case 0x8A:
			c.RESU3R8(1, "D")
			log.Printf("RES 1,D\n")

		case 0x8B:
			c.RESU3R8(1, "E")
			log.Printf("RES 1,E\n")

		case 0x8C:
			c.RESU3R8(1, "H")
			log.Printf("RES 1,H\n")

		case 0x8D:
			c.RESU3R8(1, "L")
			log.Printf("RES 1,L\n")

		case 0x8E:
			c.RESU3HL(1)
			log.Printf("RES 1,(HL)\n")

		case 0x8F:
			c.RESU3R8(1, "A")
			log.Printf("RES 1,A\n")

		}

	case 0x90:
		switch opCode {
		case 0x90:
			c.RESU3R8(2, "B")
			log.Printf("RES 2,B\n")

		case 0x91:
			c.RESU3R8(2, "C")
			log.Printf("RES 2,C\n")

		case 0x92:
			c.RESU3R8(2, "D")
			log.Printf("RES 2,D\n")

		case 0x93:
			c.RESU3R8(2, "E")
			log.Printf("RES 2,E\n")

		case 0x94:
			c.RESU3R8(2, "H")
			log.Printf("RES 2,H\n")

		case 0x95:
			c.RESU3R8(2, "L")
			log.Printf("RES 2,L\n")

		case 0x96:
			c.RESU3HL(2)
			log.Printf("RES 2,(HL)\n")

		case 0x97:
			c.RESU3R8(2, "A")
			log.Printf("RES 2,A\n")

		case 0x98:
			c.RESU3R8(3, "B")
			log.Printf("RES 3,B\n")

		case 0x99:
			c.RESU3R8(3, "C")
			log.Printf("RES 3,C\n")

		case 0x9A:
			c.RESU3R8(3, "D")
			log.Printf("RES 3,D\n")

		case 0x9B:
			c.RESU3R8(3, "E")
			log.Printf("RES 3,E\n")

		case 0x9C:
			c.RESU3R8(3, "H")
			log.Printf("RES 3,H\n")

		case 0x9D:
			c.RESU3R8(3, "L")
			log.Printf("RES 3,L\n")

		case 0x9E:
			c.RESU3HL(3)
			log.Printf("RES 3,(HL)\n")

		case 0x9F:
			c.RESU3R8(3, "A")
			log.Printf("RES 3,A\n")

		}

	case 0xA0:
		switch opCode {
		case 0xA0:
			c.RESU3R8(4, "B")
			log.Printf("RES 4,B\n")

		case 0xA1:
			c.RESU3R8(4, "C")
			log.Printf("RES 4,C\n")

		case 0xA2:
			c.RESU3R8(4, "D")
			log.Printf("RES 4,D\n")

		case 0xA3:
			c.RESU3R8(4, "E")
			log.Printf("RES 4,E\n")

		case 0xA4:
			c.RESU3R8(4, "H")
			log.Printf("RES 4,H\n")

		case 0xA5:
			c.RESU3R8(4, "L")
			log.Printf("RES 4,L\n")

		case 0xA6:
			c.RESU3HL(4)
			log.Printf("RES A,(HL)\n")

		case 0xA7:
			c.RESU3R8(4, "A")
			log.Printf("RES 4,A\n")

		case 0xA8:
			c.RESU3R8(5, "B")
			log.Printf("RES 5,B\n")

		case 0xA9:
			c.RESU3R8(5, "C")
			log.Printf("RES 5,C\n")

		case 0xAA:
			c.RESU3R8(5, "D")
			log.Printf("RES 5,D\n")

		case 0xAB:
			c.RESU3R8(5, "E")
			log.Printf("RES 5,E\n")

		case 0xAC:
			c.RESU3R8(5, "H")
			log.Printf("RES 5,H\n")

		case 0xAD:
			c.RESU3R8(5, "L")
			log.Printf("RES 5,L\n")

		case 0xAE:
			c.RESU3HL(5)
			log.Printf("RES 5,(HL)\n")

		case 0xAF:
			c.RESU3R8(5, "A")
			log.Printf("RES 5,A\n")

		}

	case 0xB0:
		switch opCode {
		case 0xB0:
			c.RESU3R8(6, "B")
			log.Printf("RES 6,B\n")

		case 0xB1:
			c.RESU3R8(6, "C")
			log.Printf("RES 6,C\n")

		case 0xB2:
			c.RESU3R8(6, "D")
			log.Printf("RES 6,D\n")

		case 0xB3:
			c.RESU3R8(6, "E")
			log.Printf("RES 6,E\n")

		case 0xB4:
			c.RESU3R8(6, "H")
			log.Printf("RES 6,H\n")

		case 0xB5:
			c.RESU3R8(6, "L")
			log.Printf("RES 6,L\n")

		case 0xB6:
			c.RESU3HL(6)
			log.Printf("RES 6,(HL)\n")

		case 0xB7:
			c.RESU3R8(6, "A")
			log.Printf("RES 6,A\n")

		case 0xB8:
			c.RESU3R8(7, "B")
			log.Printf("RES 7,B\n")

		case 0xB9:
			c.RESU3R8(7, "C")
			log.Printf("RES 7,C\n")

		case 0xBA:
			c.RESU3R8(7, "D")
			log.Printf("RES 7,D\n")

		case 0xBB:
			c.RESU3R8(7, "E")
			log.Printf("RES 7,E\n")

		case 0xBC:
			c.RESU3R8(7, "H")
			log.Printf("RES 7,H\n")

		case 0xBD:
			c.RESU3R8(7, "L")
			log.Printf("RES 7,L\n")

		case 0xBE:
			c.RESU3HL(7)
			log.Printf("RES 7,(HL)\n")

		case 0xBF:
			c.RESU3R8(7, "A")
			log.Printf("RES 7,A\n")

		}

	case 0xC0:
		switch opCode {
		case 0xC0:
			c.SETU3R8(0, "B")
			log.Printf("SET 0,B\n")

		case 0xC1:
			c.SETU3R8(0, "C")
			log.Printf("SET 0,C\n")

		case 0xC2:
			c.SETU3R8(0, "D")
			log.Printf("SET 0,D\n")

		case 0xC3:
			c.SETU3R8(0, "E")
			log.Printf("SET 0,E\n")

		case 0xC4:
			c.SETU3R8(0, "H")
			log.Printf("SET 0,H\n")

		case 0xC5:
			c.SETU3R8(0, "L")
			log.Printf("SET 0,L\n")

		case 0xC6:
			c.SETU3HL(0)
			log.Printf("SET 0,(HL)\n")

		case 0xC7:
			c.SETU3R8(0, "A")
			log.Printf("SET 0,A\n")

		case 0xC8:
			c.SETU3R8(1, "B")
			log.Printf("SET 1,B\n")

		case 0xC9:
			c.SETU3R8(1, "C")
			log.Printf("SET 1,C\n")

		case 0xCA:
			c.SETU3R8(1, "D")
			log.Printf("SET 1,D\n")

		case 0xCB:
			c.SETU3R8(1, "E")
			log.Printf("SET 1,E\n")

		case 0xCC:
			c.SETU3R8(1, "H")
			log.Printf("SET 1,H\n")

		case 0xCD:
			c.SETU3R8(1, "L")
			log.Printf("SET 1,L\n")

		case 0xCE:
			c.SETU3HL(1)
			log.Printf("SET 1,(HL)\n")

		case 0xCF:
			c.SETU3R8(1, "A")
			log.Printf("SET 1,A\n")

		}

	case 0xD0:
		switch opCode {
		case 0xD0:
			c.SETU3R8(2, "B")
			log.Printf("SET 2,B\n")

		case 0xD1:
			c.SETU3R8(2, "C")
			log.Printf("SET 2,C\n")

		case 0xD2:
			c.SETU3R8(2, "D")
			log.Printf("SET 2,D\n")

		case 0xD3:
			c.SETU3R8(2, "E")
			log.Printf("SET 2,E\n")

		case 0xD4:
			c.SETU3R8(2, "H")
			log.Printf("SET 2,H\n")

		case 0xD5:
			c.SETU3R8(2, "L")
			log.Printf("SET 2,L\n")

		case 0xD6:
			c.SETU3HL(2)
			log.Printf("SET 2,(HL)\n")

		case 0xD7:
			c.SETU3R8(2, "A")
			log.Printf("SET 2,A\n")

		case 0xD8:
			c.SETU3R8(3, "B")
			log.Printf("SET 3,B\n")

		case 0xD9:
			c.SETU3R8(3, "C")
			log.Printf("SET 3,C\n")

		case 0xDA:
			c.SETU3R8(3, "D")
			log.Printf("SET 3,D\n")

		case 0xDB:
			c.SETU3R8(3, "E")
			log.Printf("SET 3,E\n")

		case 0xDC:
			c.SETU3R8(3, "H")
			log.Printf("SET 3,H\n")

		case 0xDD:
			c.SETU3R8(3, "L")
			log.Printf("SET 3,L\n")

		case 0xDE:
			c.SETU3HL(3)
			log.Printf("SET 3,(HL)\n")

		case 0xDF:
			c.SETU3R8(3, "A")
			log.Printf("SET 3,A\n")

		}

	case 0xE0:
		switch opCode {
		case 0xE0:
			c.SETU3R8(4, "B")
			log.Printf("SET 4,B\n")

		case 0xE1:
			c.SETU3R8(4, "C")
			log.Printf("SET 4,C\n")

		case 0xE2:
			c.SETU3R8(4, "D")
			log.Printf("SET 4,D\n")

		case 0xE3:
			c.SETU3R8(4, "E")
			log.Printf("SET 4,E\n")

		case 0xE4:
			c.SETU3R8(4, "H")
			log.Printf("SET 4,H\n")

		case 0xE5:
			c.SETU3R8(4, "L")
			log.Printf("SET 4,L\n")

		case 0xE6:
			c.SETU3HL(4)
			log.Printf("SET 4,(HL)\n")

		case 0xE7:
			c.SETU3R8(4, "A")
			log.Printf("SET 4,A\n")

		case 0xE8:
			c.SETU3R8(5, "B")
			log.Printf("SET 5,B\n")

		case 0xE9:
			c.SETU3R8(5, "C")
			log.Printf("SET 5,C\n")

		case 0xEA:
			c.SETU3R8(5, "D")
			log.Printf("SET 5,D\n")

		case 0xEB:
			c.SETU3R8(5, "E")
			log.Printf("SET 5,E\n")

		case 0xEC:
			c.SETU3R8(5, "H")
			log.Printf("SET 5,H\n")

		case 0xED:
			c.SETU3R8(5, "L")
			log.Printf("SET 5,L\n")

		case 0xEE:
			c.SETU3HL(5)
			log.Printf("SET 5,(HL)\n")

		case 0xEF:
			c.SETU3R8(5, "A")
			log.Printf("SET 5,A\n")

		}

	case 0xF0:
		switch opCode {
		case 0xF0:
			c.SETU3R8(6, "B")
			log.Printf("SET 6,B\n")

		case 0xF1:
			c.SETU3R8(6, "C")
			log.Printf("SET 6,C\n")

		case 0xF2:
			c.SETU3R8(6, "D")
			log.Printf("SET 6,D\n")

		case 0xF3:
			c.SETU3R8(6, "E")
			log.Printf("SET 6,E\n")

		case 0xF4:
			c.SETU3R8(6, "H")
			log.Printf("SET 6,H\n")

		case 0xF5:
			c.SETU3R8(6, "L")
			log.Printf("SET 6,L\n")

		case 0xF6:
			c.SETU3HL(6)
			log.Printf("SET 6,(HL)\n")

		case 0xF7:
			c.SETU3R8(6, "A")
			log.Printf("SET 6,A\n")

		case 0xF8:
			c.SETU3R8(7, "B")
			log.Printf("SET 7,B\n")

		case 0xF9:
			c.SETU3R8(7, "C")
			log.Printf("SET 7,C\n")

		case 0xFA:
			c.SETU3R8(7, "D")
			log.Printf("SET 7,D\n")

		case 0xFB:
			c.SETU3R8(7, "E")
			log.Printf("SET 7,E\n")

		case 0xFC:
			c.SETU3R8(7, "H")
			log.Printf("SET 7,H\n")

		case 0xFD:
			c.SETU3R8(7, "L")
			log.Printf("SET 7,L\n")

		case 0xFE:
			c.SETU3HL(7)
			log.Printf("SET 7,(HL)\n")

		case 0xFF:
			c.SETU3R8(7, "A")
			log.Printf("SET 7,A\n")

		}
	}
}
