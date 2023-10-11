package CPU

import (
	"log"
)

func (s *CPU) PrefixTable(log *log.Logger) {
	opCode := uint8(s.Mem[s.GetReg16Val("PC")+1])

	switch opCode {
	case 0x00:
		s.RLCR8("B")
		log.Printf("RLC B\n")

	case 0x01:
		s.RLCR8("C")
		log.Printf("RLC C\n")

	case 0x02:
		s.RLCR8("D")
		log.Printf("RLC D\n")

	case 0x03:
		s.RLCR8("E")
		log.Printf("RLC E\n")

	case 0x04:
		s.RLCR8("H")
		log.Printf("RLC H\n")

	case 0x05:
		s.RLCR8("L")
		log.Printf("RLC L\n")

	case 0x06:
		s.RLCHL()
		log.Printf("RLC (HL)\n")

	case 0x07:
		s.RLCR8("A")
		log.Printf("RLC A\n")

	case 0x08:
		s.RRCR8("B")
		log.Printf("RRC B\n")

	case 0x09:
		s.RRCR8("C")
		log.Printf("RRC C\n")

	case 0x0A:
		s.RRCR8("D")
		log.Printf("RRC D\n")

	case 0x0B:
		s.RRCR8("E")
		log.Printf("RRC E\n")

	case 0x0C:
		s.RRCR8("H")
		log.Printf("RRC H\n")

	case 0x0D:
		s.RRCR8("L")
		log.Printf("RRC L\n")

	case 0x0E:
		s.RRCHL()
		log.Printf("RRC (HL)\n")

	case 0x0F:
		s.RRCR8("A")
		log.Printf("RRC A\n")

	case 0x10:
		s.RLR8("B")
		log.Printf("RL B\n")

	case 0x11:
		s.RLR8("C")
		log.Printf("RL C\n")

	case 0x12:
		s.RLR8("D")
		log.Printf("RL D\n")

	case 0x13:
		s.RLR8("E")
		log.Printf("RL E\n")

	case 0x14:
		s.RLR8("H")
		log.Printf("RL H\n")

	case 0x15:
		s.RLR8("L")
		log.Printf("RL L\n")

	case 0x16:
		s.RLHL()
		log.Printf("RL (HL)\n")

	case 0x17:
		s.RLR8("A")
		log.Printf("RL A\n")

	case 0x18:
		s.RRR8("B")
		log.Printf("RR B\n")

	case 0x19:
		s.RRR8("C")
		log.Printf("RR C\n")

	case 0x1A:
		s.RRR8("D")
		log.Printf("RR D\n")

	case 0x1B:
		s.RRR8("E")
		log.Printf("RR E\n")

	case 0x1C:
		s.RRR8("H")
		log.Printf("RR H\n")

	case 0x1D:
		s.RRR8("L")
		log.Printf("RR L\n")

	case 0x1E:
		s.RRHL()
		log.Printf("RR (HL)\n")

	case 0x1F:
		s.RRR8("A")
		log.Printf("RR A\n")

	case 0x20:
		s.SLAR8("B")
		log.Printf("SLA B\n")

	case 0x21:
		s.SLAR8("C")
		log.Printf("SLA C\n")

	case 0x22:
		s.SLAR8("D")
		log.Printf("SLA D\n")

	case 0x23:
		s.SLAR8("E")
		log.Printf("SLA E\n")

	case 0x24:
		s.SLAR8("H")
		log.Printf("SLA H\n")

	case 0x25:
		s.SLAR8("L")
		log.Printf("SLA L\n")

	case 0x26:
		s.SLAHL()
		log.Printf("SLA (HL)\n")

	case 0x27:
		s.SLAR8("A")
		log.Printf("SLA A\n")

	case 0x28:
		s.SRAR8("B")
		log.Printf("SRA B\n")

	case 0x29:
		s.SRAR8("C")
		log.Printf("SRA C\n")

	case 0x2A:
		s.SRAR8("D")
		log.Printf("SRA D\n")

	case 0x2B:
		s.SRAR8("E")
		log.Printf("SRA E\n")

	case 0x2C:
		s.SRAR8("H")
		log.Printf("SRA H\n")

	case 0x2D:
		s.SRAR8("L")
		log.Printf("SRA L\n")

	case 0x2E:
		s.SRAHL()
		log.Printf("SRA (HL)\n")

	case 0x2F:
		s.SRAR8("A")
		log.Printf("SRA A\n")

	case 0x30:
		s.SWAPR8("B")
		log.Printf("SWAP B\n")

	case 0x31:
		s.SWAPR8("C")
		log.Printf("SWAP C\n")

	case 0x32:
		s.SWAPR8("D")
		log.Printf("SWAP D\n")

	case 0x33:
		s.SWAPR8("E")
		log.Printf("SWAP E\n")

	case 0x34:
		s.SWAPR8("H")
		log.Printf("SWAP H\n")

	case 0x35:
		s.SWAPR8("L")
		log.Printf("SWAP L\n")

	case 0x36:
		s.SWAPHL()
		log.Printf("SWAP (HL)\n")

	case 0x37:
		s.SWAPR8("A")
		log.Printf("SWAP A\n")

	case 0x38:
		s.SRLR8("B")
		log.Printf("SRA B\n")

	case 0x39:
		s.SRLR8("C")
		log.Printf("SRA C\n")

	case 0x3A:
		s.SRLR8("D")
		log.Printf("SRA D\n")

	case 0x3B:
		s.SRLR8("E")
		log.Printf("SRA E\n")

	case 0x3C:
		s.SRLR8("H")
		log.Printf("SRA H\n")

	case 0x3D:
		s.SRLR8("L")
		log.Printf("SRA L\n")

	case 0x3E:
		s.SRLHL()
		log.Printf("SRA (HL)\n")

	case 0x3F:
		s.SRLR8("A")
		log.Printf("SRA A\n")

	case 0x40:
		s.BITU3R8(0, "B")
		log.Printf("BIT 0,B\n")

	case 0x41:
		s.BITU3R8(0, "C")
		log.Printf("BIT 0,C\n")

	case 0x42:
		s.BITU3R8(0, "D")
		log.Printf("BIT 0,D\n")

	case 0x43:
		s.BITU3R8(0, "E")
		log.Printf("BIT 0,E\n")

	case 0x44:
		s.BITU3R8(0, "H")
		log.Printf("BIT 0,H\n")

	case 0x45:
		s.BITU3R8(0, "L")
		log.Printf("BIT 0,L\n")

	case 0x46:
		s.BITU3HL(0)
		log.Printf("BIT 0,(HL)\n")

	case 0x47:
		s.BITU3R8(0, "A")
		log.Printf("BIT 0,A\n")

	case 0x48:
		s.BITU3R8(1, "B")
		log.Printf("BIT 1,B\n")

	case 0x49:
		s.BITU3R8(1, "C")
		log.Printf("BIT 1,C\n")

	case 0x4A:
		s.BITU3R8(1, "D")
		log.Printf("BIT 1,D\n")

	case 0x4B:
		s.BITU3R8(1, "E")
		log.Printf("BIT 1,E\n")

	case 0x4C:
		s.BITU3R8(1, "H")
		log.Printf("BIT 1,H\n")

	case 0x4D:
		s.BITU3R8(1, "L")
		log.Printf("BIT 1,L\n")

	case 0x4E:
		s.BITU3HL(1)
		log.Printf("BIT 1,(HL)\n")

	case 0x4F:
		s.BITU3R8(1, "A")
		log.Printf("BIT 1,A\n")

	case 0x50:
		s.BITU3R8(2, "B")
		log.Printf("BIT 2,B\n")

	case 0x51:
		s.BITU3R8(2, "C")
		log.Printf("BIT 2,C\n")

	case 0x52:
		s.BITU3R8(2, "D")
		log.Printf("BIT 2,D\n")

	case 0x53:
		s.BITU3R8(2, "E")
		log.Printf("BIT 2,E\n")

	case 0x54:
		s.BITU3R8(2, "H")
		log.Printf("BIT 2,H\n")

	case 0x55:
		s.BITU3R8(2, "L")
		log.Printf("BIT 2,L\n")

	case 0x56:
		s.BITU3HL(2)
		log.Printf("BIT 2,(HL)\n")

	case 0x57:
		s.BITU3R8(2, "A")
		log.Printf("BIT 2,A\n")

	case 0x58:
		s.BITU3R8(3, "B")
		log.Printf("BIT 3,B\n")

	case 0x59:
		s.BITU3R8(3, "C")
		log.Printf("BIT 3,C\n")

	case 0x5A:
		s.BITU3R8(3, "D")
		log.Printf("BIT 3,D\n")

	case 0x5B:
		s.BITU3R8(3, "E")
		log.Printf("BIT 3,E\n")

	case 0x5C:
		s.BITU3R8(3, "H")
		log.Printf("BIT 3,H\n")

	case 0x5D:
		s.BITU3R8(3, "L")
		log.Printf("BIT 3,L\n")

	case 0x5E:
		s.BITU3HL(3)
		log.Printf("BIT 3,(HL)\n")

	case 0x5F:
		s.BITU3R8(3, "A")
		log.Printf("BIT 3,A\n")

	case 0x60:
		s.BITU3R8(4, "B")
		log.Printf("BIT 4,B\n")

	case 0x61:
		s.BITU3R8(4, "C")
		log.Printf("BIT 4,C\n")

	case 0x62:
		s.BITU3R8(4, "D")
		log.Printf("BIT 4,D\n")

	case 0x63:
		s.BITU3R8(4, "E")
		log.Printf("BIT 4,E\n")

	case 0x64:
		s.BITU3R8(4, "H")
		log.Printf("BIT 4,H\n")

	case 0x65:
		s.BITU3R8(4, "L")
		log.Printf("BIT 4,L\n")

	case 0x66:
		s.BITU3HL(4)
		log.Printf("BIT 4,(HL)\n")

	case 0x67:
		s.BITU3R8(4, "A")
		log.Printf("BIT 4,A\n")

	case 0x68:
		s.BITU3R8(5, "B")
		log.Printf("BIT 5,B\n")

	case 0x69:
		s.BITU3R8(5, "C")
		log.Printf("BIT 5,C\n")

	case 0x6A:
		s.BITU3R8(5, "D")
		log.Printf("BIT 5,D\n")

	case 0x6B:
		s.BITU3R8(5, "E")
		log.Printf("BIT 5,E\n")

	case 0x6C:
		s.BITU3R8(5, "H")
		log.Printf("BIT 5,H\n")

	case 0x6D:
		s.BITU3R8(5, "L")
		log.Printf("BIT 5,L\n")

	case 0x6E:
		s.BITU3HL(5)
		log.Printf("BIT 5,(HL)\n")

	case 0x6F:
		s.BITU3R8(5, "A")
		log.Printf("BIT 5,A\n")

	case 0x70:
		s.BITU3R8(6, "B")
		log.Printf("BIT 6,B\n")

	case 0x71:
		s.BITU3R8(6, "C")
		log.Printf("BIT 6,C\n")

	case 0x72:
		s.BITU3R8(6, "D")
		log.Printf("BIT 6,D\n")

	case 0x73:
		s.BITU3R8(6, "E")
		log.Printf("BIT 6,E\n")

	case 0x74:
		s.BITU3R8(6, "H")
		log.Printf("BIT 6,H\n")

	case 0x75:
		s.BITU3R8(6, "L")
		log.Printf("BIT 6,L\n")

	case 0x76:
		s.BITU3HL(6)
		log.Printf("BIT 6,(HL)\n")

	case 0x77:
		s.BITU3R8(6, "A")
		log.Printf("BIT 6,A\n")

	case 0x78:
		s.BITU3R8(7, "B")
		log.Printf("BIT 7,B\n")

	case 0x79:
		s.BITU3R8(7, "C")
		log.Printf("BIT 7,C\n")

	case 0x7A:
		s.BITU3R8(7, "D")
		log.Printf("BIT 7,D\n")

	case 0x7B:
		s.BITU3R8(7, "E")
		log.Printf("BIT 7,E\n")

	case 0x7C:
		s.BITU3R8(7, "H")
		log.Printf("BIT 7,H\n")

	case 0x7D:
		s.BITU3R8(7, "L")
		log.Printf("BIT 7,L\n")

	case 0x7E:
		s.BITU3HL(7)
		log.Printf("BIT 7,(HL)\n")

	case 0x7F:
		s.BITU3R8(7, "A")
		log.Printf("BIT 7,A\n")

	case 0x80:
		s.RESU3R8(0, "B")
		log.Printf("RES 0,B\n")

	case 0x81:
		s.RESU3R8(0, "C")
		log.Printf("RES 0,C\n")

	case 0x82:
		s.RESU3R8(0, "D")
		log.Printf("RES 0,D\n")

	case 0x83:
		s.RESU3R8(0, "E")
		log.Printf("RES 0,E\n")

	case 0x84:
		s.RESU3R8(0, "H")
		log.Printf("RES 0,H\n")

	case 0x85:
		s.RESU3R8(0, "L")
		log.Printf("RES 0,L\n")

	case 0x86:
		s.RESU3HL(0)
		log.Printf("RES 0,(HL)\n")

	case 0x87:
		s.RESU3R8(0, "A")
		log.Printf("RES 0,A\n")

	case 0x88:
		s.RESU3R8(1, "B")
		log.Printf("RES 1,B\n")

	case 0x89:
		s.RESU3R8(1, "C")
		log.Printf("RES 1,C\n")

	case 0x8A:
		s.RESU3R8(1, "D")
		log.Printf("RES 1,D\n")

	case 0x8B:
		s.RESU3R8(1, "E")
		log.Printf("RES 1,E\n")

	case 0x8C:
		s.RESU3R8(1, "H")
		log.Printf("RES 1,H\n")

	case 0x8D:
		s.RESU3R8(1, "L")
		log.Printf("RES 1,L\n")

	case 0x8E:
		s.RESU3HL(1)
		log.Printf("RES 1,(HL)\n")

	case 0x8F:
		s.RESU3R8(1, "A")
		log.Printf("RES 1,A\n")

	case 0x90:
		s.RESU3R8(2, "B")
		log.Printf("RES 2,B\n")

	case 0x91:
		s.RESU3R8(2, "C")
		log.Printf("RES 2,C\n")

	case 0x92:
		s.RESU3R8(2, "D")
		log.Printf("RES 2,D\n")

	case 0x93:
		s.RESU3R8(2, "E")
		log.Printf("RES 2,E\n")

	case 0x94:
		s.RESU3R8(2, "H")
		log.Printf("RES 2,H\n")

	case 0x95:
		s.RESU3R8(2, "L")
		log.Printf("RES 2,L\n")

	case 0x96:
		s.RESU3HL(2)
		log.Printf("RES 2,(HL)\n")

	case 0x97:
		s.RESU3R8(2, "A")
		log.Printf("RES 2,A\n")

	case 0x98:
		s.RESU3R8(3, "B")
		log.Printf("RES 3,B\n")

	case 0x99:
		s.RESU3R8(3, "C")
		log.Printf("RES 3,C\n")

	case 0x9A:
		s.RESU3R8(3, "D")
		log.Printf("RES 3,D\n")

	case 0x9B:
		s.RESU3R8(3, "E")
		log.Printf("RES 3,E\n")

	case 0x9C:
		s.RESU3R8(3, "H")
		log.Printf("RES 3,H\n")

	case 0x9D:
		s.RESU3R8(3, "L")
		log.Printf("RES 3,L\n")

	case 0x9E:
		s.RESU3HL(3)
		log.Printf("RES 3,(HL)\n")

	case 0x9F:
		s.RESU3R8(3, "A")
		log.Printf("RES 3,A\n")

	case 0xA0:
		s.RESU3R8(4, "B")
		log.Printf("RES 4,B\n")

	case 0xA1:
		s.RESU3R8(4, "C")
		log.Printf("RES 4,C\n")

	case 0xA2:
		s.RESU3R8(4, "D")
		log.Printf("RES 4,D\n")

	case 0xA3:
		s.RESU3R8(4, "E")
		log.Printf("RES 4,E\n")

	case 0xA4:
		s.RESU3R8(4, "H")
		log.Printf("RES 4,H\n")

	case 0xA5:
		s.RESU3R8(4, "L")
		log.Printf("RES 4,L\n")

	case 0xA6:
		s.RESU3HL(4)
		log.Printf("RES A,(HL)\n")

	case 0xA7:
		s.RESU3R8(4, "A")
		log.Printf("RES 4,A\n")

	case 0xA8:
		s.RESU3R8(5, "B")
		log.Printf("RES 5,B\n")

	case 0xA9:
		s.RESU3R8(5, "C")
		log.Printf("RES 5,C\n")

	case 0xAA:
		s.RESU3R8(5, "D")
		log.Printf("RES 5,D\n")

	case 0xAB:
		s.RESU3R8(5, "E")
		log.Printf("RES 5,E\n")

	case 0xAC:
		s.RESU3R8(5, "H")
		log.Printf("RES 5,H\n")

	case 0xAD:
		s.RESU3R8(5, "L")
		log.Printf("RES 5,L\n")

	case 0xAE:
		s.RESU3HL(5)
		log.Printf("RES 5,(HL)\n")

	case 0xAF:
		s.RESU3R8(5, "A")
		log.Printf("RES 5,A\n")

	case 0xB0:
		s.RESU3R8(6, "B")
		log.Printf("RES 6,B\n")

	case 0xB1:
		s.RESU3R8(6, "C")
		log.Printf("RES 6,C\n")

	case 0xB2:
		s.RESU3R8(6, "D")
		log.Printf("RES 6,D\n")

	case 0xB3:
		s.RESU3R8(6, "E")
		log.Printf("RES 6,E\n")

	case 0xB4:
		s.RESU3R8(6, "H")
		log.Printf("RES 6,H\n")

	case 0xB5:
		s.RESU3R8(6, "L")
		log.Printf("RES 6,L\n")

	case 0xB6:
		s.RESU3HL(6)
		log.Printf("RES 6,(HL)\n")

	case 0xB7:
		s.RESU3R8(6, "A")
		log.Printf("RES 6,A\n")

	case 0xB8:
		s.RESU3R8(7, "B")
		log.Printf("RES 7,B\n")

	case 0xB9:
		s.RESU3R8(7, "C")
		log.Printf("RES 7,C\n")

	case 0xBA:
		s.RESU3R8(7, "D")
		log.Printf("RES 7,D\n")

	case 0xBB:
		s.RESU3R8(7, "E")
		log.Printf("RES 7,E\n")

	case 0xBC:
		s.RESU3R8(7, "H")
		log.Printf("RES 7,H\n")

	case 0xBD:
		s.RESU3R8(7, "L")
		log.Printf("RES 7,L\n")

	case 0xBE:
		s.RESU3HL(7)
		log.Printf("RES 7,(HL)\n")

	case 0xBF:
		s.RESU3R8(7, "A")
		log.Printf("RES 7,A\n")

	case 0xC0:
		s.SETU3R8(0, "B")
		log.Printf("SET 0,B\n")

	case 0xC1:
		s.SETU3R8(0, "C")
		log.Printf("SET 0,C\n")

	case 0xC2:
		s.SETU3R8(0, "D")
		log.Printf("SET 0,D\n")

	case 0xC3:
		s.SETU3R8(0, "E")
		log.Printf("SET 0,E\n")

	case 0xC4:
		s.SETU3R8(0, "H")
		log.Printf("SET 0,H\n")

	case 0xC5:
		s.SETU3R8(0, "L")
		log.Printf("SET 0,L\n")

	case 0xC6:
		s.SETU3HL(0)
		log.Printf("SET 0,(HL)\n")

	case 0xC7:
		s.SETU3R8(0, "A")
		log.Printf("SET 0,A\n")

	case 0xC8:
		s.SETU3R8(1, "B")
		log.Printf("SET 1,B\n")

	case 0xC9:
		s.SETU3R8(1, "C")
		log.Printf("SET 1,C\n")

	case 0xCA:
		s.SETU3R8(1, "D")
		log.Printf("SET 1,D\n")

	case 0xCB:
		s.SETU3R8(1, "E")
		log.Printf("SET 1,E\n")

	case 0xCC:
		s.SETU3R8(1, "H")
		log.Printf("SET 1,H\n")

	case 0xCD:
		s.SETU3R8(1, "L")
		log.Printf("SET 1,L\n")

	case 0xCE:
		s.SETU3HL(1)
		log.Printf("SET 1,(HL)\n")

	case 0xCF:
		s.SETU3R8(1, "A")
		log.Printf("SET 1,A\n")

	case 0xD0:
		s.SETU3R8(2, "B")
		log.Printf("SET 2,B\n")

	case 0xD1:
		s.SETU3R8(2, "C")
		log.Printf("SET 2,C\n")

	case 0xD2:
		s.SETU3R8(2, "D")
		log.Printf("SET 2,D\n")

	case 0xD3:
		s.SETU3R8(2, "E")
		log.Printf("SET 2,E\n")

	case 0xD4:
		s.SETU3R8(2, "H")
		log.Printf("SET 2,H\n")

	case 0xD5:
		s.SETU3R8(2, "L")
		log.Printf("SET 2,L\n")

	case 0xD6:
		s.SETU3HL(2)
		log.Printf("SET 2,(HL)\n")

	case 0xD7:
		s.SETU3R8(2, "A")
		log.Printf("SET 2,A\n")

	case 0xD8:
		s.SETU3R8(3, "B")
		log.Printf("SET 3,B\n")

	case 0xD9:
		s.SETU3R8(3, "C")
		log.Printf("SET 3,C\n")

	case 0xDA:
		s.SETU3R8(3, "D")
		log.Printf("SET 3,D\n")

	case 0xDB:
		s.SETU3R8(3, "E")
		log.Printf("SET 3,E\n")

	case 0xDC:
		s.SETU3R8(3, "H")
		log.Printf("SET 3,H\n")

	case 0xDD:
		s.SETU3R8(3, "L")
		log.Printf("SET 3,L\n")

	case 0xDE:
		s.SETU3HL(3)
		log.Printf("SET 3,(HL)\n")

	case 0xDF:
		s.SETU3R8(3, "A")
		log.Printf("SET 3,A\n")

	case 0xE0:
		s.SETU3R8(4, "B")
		log.Printf("SET 4,B\n")

	case 0xE1:
		s.SETU3R8(4, "C")
		log.Printf("SET 4,C\n")

	case 0xE2:
		s.SETU3R8(4, "D")
		log.Printf("SET 4,D\n")

	case 0xE3:
		s.SETU3R8(4, "E")
		log.Printf("SET 4,E\n")

	case 0xE4:
		s.SETU3R8(4, "H")
		log.Printf("SET 4,H\n")

	case 0xE5:
		s.SETU3R8(4, "L")
		log.Printf("SET 4,L\n")

	case 0xE6:
		s.SETU3HL(4)
		log.Printf("SET 4,(HL)\n")

	case 0xE7:
		s.SETU3R8(4, "A")
		log.Printf("SET 4,A\n")

	case 0xE8:
		s.SETU3R8(5, "B")
		log.Printf("SET 5,B\n")

	case 0xE9:
		s.SETU3R8(5, "C")
		log.Printf("SET 5,C\n")

	case 0xEA:
		s.SETU3R8(5, "D")
		log.Printf("SET 5,D\n")

	case 0xEB:
		s.SETU3R8(5, "E")
		log.Printf("SET 5,E\n")

	case 0xEC:
		s.SETU3R8(5, "H")
		log.Printf("SET 5,H\n")

	case 0xED:
		s.SETU3R8(5, "L")
		log.Printf("SET 5,L\n")

	case 0xEE:
		s.SETU3HL(5)
		log.Printf("SET 5,(HL)\n")

	case 0xEF:
		s.SETU3R8(5, "A")
		log.Printf("SET 5,A\n")

	case 0xF0:
		s.SETU3R8(6, "B")
		log.Printf("SET 6,B\n")

	case 0xF1:
		s.SETU3R8(6, "C")
		log.Printf("SET 6,C\n")

	case 0xF2:
		s.SETU3R8(6, "D")
		log.Printf("SET 6,D\n")

	case 0xF3:
		s.SETU3R8(6, "E")
		log.Printf("SET 6,E\n")

	case 0xF4:
		s.SETU3R8(6, "H")
		log.Printf("SET 6,H\n")

	case 0xF5:
		s.SETU3R8(6, "L")
		log.Printf("SET 6,L\n")

	case 0xF6:
		s.SETU3HL(6)
		log.Printf("SET 6,(HL)\n")

	case 0xF7:
		s.SETU3R8(6, "A")
		log.Printf("SET 6,A\n")

	case 0xF8:
		s.SETU3R8(7, "B")
		log.Printf("SET 7,B\n")

	case 0xF9:
		s.SETU3R8(7, "C")
		log.Printf("SET 7,C\n")

	case 0xFA:
		s.SETU3R8(7, "D")
		log.Printf("SET 7,D\n")

	case 0xFB:
		s.SETU3R8(7, "E")
		log.Printf("SET 7,E\n")

	case 0xFC:
		s.SETU3R8(7, "H")
		log.Printf("SET 7,H\n")

	case 0xFD:
		s.SETU3R8(7, "L")
		log.Printf("SET 7,L\n")

	case 0xFE:
		s.SETU3HL(7)
		log.Printf("SET 7,(HL)\n")

	case 0xFF:
		s.SETU3R8(7, "A")
		log.Printf("SET 7,A\n")

	}
}
