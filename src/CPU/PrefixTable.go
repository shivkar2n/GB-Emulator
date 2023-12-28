package CPU

import (
	"fmt"
)

func (s *CPU) PrefixTable() {
	opCode := uint8(s.Mem.Read(s.GetReg16Val("PC") + 1))

	switch opCode {
	case 0x00:
		s.RLCR8("B")
		s.CurrOpcode = fmt.Sprintf("RLC B")

	case 0x01:
		s.RLCR8("C")
		s.CurrOpcode = fmt.Sprintf("RLC C")

	case 0x02:
		s.RLCR8("D")
		s.CurrOpcode = fmt.Sprintf("RLC D")

	case 0x03:
		s.RLCR8("E")
		s.CurrOpcode = fmt.Sprintf("RLC E")

	case 0x04:
		s.RLCR8("H")
		s.CurrOpcode = fmt.Sprintf("RLC H")

	case 0x05:
		s.RLCR8("L")
		s.CurrOpcode = fmt.Sprintf("RLC L")

	case 0x06:
		s.RLCHL()
		s.CurrOpcode = fmt.Sprintf("RLC (HL)")

	case 0x07:
		s.RLCR8("A")
		s.CurrOpcode = fmt.Sprintf("RLC A")

	case 0x08:
		s.RRCR8("B")
		s.CurrOpcode = fmt.Sprintf("RRC B")

	case 0x09:
		s.RRCR8("C")
		s.CurrOpcode = fmt.Sprintf("RRC C")

	case 0x0A:
		s.RRCR8("D")
		s.CurrOpcode = fmt.Sprintf("RRC D")

	case 0x0B:
		s.RRCR8("E")
		s.CurrOpcode = fmt.Sprintf("RRC E")

	case 0x0C:
		s.RRCR8("H")
		s.CurrOpcode = fmt.Sprintf("RRC H")

	case 0x0D:
		s.RRCR8("L")
		s.CurrOpcode = fmt.Sprintf("RRC L")

	case 0x0E:
		s.RRCHL()
		s.CurrOpcode = fmt.Sprintf("RRC (HL)")

	case 0x0F:
		s.RRCR8("A")
		s.CurrOpcode = fmt.Sprintf("RRC A")

	case 0x10:
		s.RLR8("B")
		s.CurrOpcode = fmt.Sprintf("RL B")

	case 0x11:
		s.RLR8("C")
		s.CurrOpcode = fmt.Sprintf("RL C")

	case 0x12:
		s.RLR8("D")
		s.CurrOpcode = fmt.Sprintf("RL D")

	case 0x13:
		s.RLR8("E")
		s.CurrOpcode = fmt.Sprintf("RL E")

	case 0x14:
		s.RLR8("H")
		s.CurrOpcode = fmt.Sprintf("RL H")

	case 0x15:
		s.RLR8("L")
		s.CurrOpcode = fmt.Sprintf("RL L")

	case 0x16:
		s.RLHL()
		s.CurrOpcode = fmt.Sprintf("RL (HL)")

	case 0x17:
		s.RLR8("A")
		s.CurrOpcode = fmt.Sprintf("RL A")

	case 0x18:
		s.RRR8("B")
		s.CurrOpcode = fmt.Sprintf("RR B")

	case 0x19:
		s.RRR8("C")
		s.CurrOpcode = fmt.Sprintf("RR C")

	case 0x1A:
		s.RRR8("D")
		s.CurrOpcode = fmt.Sprintf("RR D")

	case 0x1B:
		s.RRR8("E")
		s.CurrOpcode = fmt.Sprintf("RR E")

	case 0x1C:
		s.RRR8("H")
		s.CurrOpcode = fmt.Sprintf("RR H")

	case 0x1D:
		s.RRR8("L")
		s.CurrOpcode = fmt.Sprintf("RR L")

	case 0x1E:
		s.RRHL()
		s.CurrOpcode = fmt.Sprintf("RR (HL)")

	case 0x1F:
		s.RRR8("A")
		s.CurrOpcode = fmt.Sprintf("RR A")

	case 0x20:
		s.SLAR8("B")
		s.CurrOpcode = fmt.Sprintf("SLA B")

	case 0x21:
		s.SLAR8("C")
		s.CurrOpcode = fmt.Sprintf("SLA C")

	case 0x22:
		s.SLAR8("D")
		s.CurrOpcode = fmt.Sprintf("SLA D")

	case 0x23:
		s.SLAR8("E")
		s.CurrOpcode = fmt.Sprintf("SLA E")

	case 0x24:
		s.SLAR8("H")
		s.CurrOpcode = fmt.Sprintf("SLA H")

	case 0x25:
		s.SLAR8("L")
		s.CurrOpcode = fmt.Sprintf("SLA L")

	case 0x26:
		s.SLAHL()
		s.CurrOpcode = fmt.Sprintf("SLA (HL)")

	case 0x27:
		s.SLAR8("A")
		s.CurrOpcode = fmt.Sprintf("SLA A")

	case 0x28:
		s.SRAR8("B")
		s.CurrOpcode = fmt.Sprintf("SRA B")

	case 0x29:
		s.SRAR8("C")
		s.CurrOpcode = fmt.Sprintf("SRA C")

	case 0x2A:
		s.SRAR8("D")
		s.CurrOpcode = fmt.Sprintf("SRA D")

	case 0x2B:
		s.SRAR8("E")
		s.CurrOpcode = fmt.Sprintf("SRA E")

	case 0x2C:
		s.SRAR8("H")
		s.CurrOpcode = fmt.Sprintf("SRA H")

	case 0x2D:
		s.SRAR8("L")
		s.CurrOpcode = fmt.Sprintf("SRA L")

	case 0x2E:
		s.SRAHL()
		s.CurrOpcode = fmt.Sprintf("SRA (HL)")

	case 0x2F:
		s.SRAR8("A")
		s.CurrOpcode = fmt.Sprintf("SRA A")

	case 0x30:
		s.SWAPR8("B")
		s.CurrOpcode = fmt.Sprintf("SWAP B")

	case 0x31:
		s.SWAPR8("C")
		s.CurrOpcode = fmt.Sprintf("SWAP C")

	case 0x32:
		s.SWAPR8("D")
		s.CurrOpcode = fmt.Sprintf("SWAP D")

	case 0x33:
		s.SWAPR8("E")
		s.CurrOpcode = fmt.Sprintf("SWAP E")

	case 0x34:
		s.SWAPR8("H")
		s.CurrOpcode = fmt.Sprintf("SWAP H")

	case 0x35:
		s.SWAPR8("L")
		s.CurrOpcode = fmt.Sprintf("SWAP L")

	case 0x36:
		s.SWAPHL()
		s.CurrOpcode = fmt.Sprintf("SWAP (HL)")

	case 0x37:
		s.SWAPR8("A")
		s.CurrOpcode = fmt.Sprintf("SWAP A")

	case 0x38:
		s.SRLR8("B")
		s.CurrOpcode = fmt.Sprintf("SRA B")

	case 0x39:
		s.SRLR8("C")
		s.CurrOpcode = fmt.Sprintf("SRA C")

	case 0x3A:
		s.SRLR8("D")
		s.CurrOpcode = fmt.Sprintf("SRA D")

	case 0x3B:
		s.SRLR8("E")
		s.CurrOpcode = fmt.Sprintf("SRA E")

	case 0x3C:
		s.SRLR8("H")
		s.CurrOpcode = fmt.Sprintf("SRA H")

	case 0x3D:
		s.SRLR8("L")
		s.CurrOpcode = fmt.Sprintf("SRA L")

	case 0x3E:
		s.SRLHL()
		s.CurrOpcode = fmt.Sprintf("SRA (HL)")

	case 0x3F:
		s.SRLR8("A")
		s.CurrOpcode = fmt.Sprintf("SRA A")

	case 0x40:
		s.BITU3R8(0, "B")
		s.CurrOpcode = fmt.Sprintf("BIT 0,B")

	case 0x41:
		s.BITU3R8(0, "C")
		s.CurrOpcode = fmt.Sprintf("BIT 0,C")

	case 0x42:
		s.BITU3R8(0, "D")
		s.CurrOpcode = fmt.Sprintf("BIT 0,D")

	case 0x43:
		s.BITU3R8(0, "E")
		s.CurrOpcode = fmt.Sprintf("BIT 0,E")

	case 0x44:
		s.BITU3R8(0, "H")
		s.CurrOpcode = fmt.Sprintf("BIT 0,H")

	case 0x45:
		s.BITU3R8(0, "L")
		s.CurrOpcode = fmt.Sprintf("BIT 0,L")

	case 0x46:
		s.BITU3HL(0)
		s.CurrOpcode = fmt.Sprintf("BIT 0,(HL)")

	case 0x47:
		s.BITU3R8(0, "A")
		s.CurrOpcode = fmt.Sprintf("BIT 0,A")

	case 0x48:
		s.BITU3R8(1, "B")
		s.CurrOpcode = fmt.Sprintf("BIT 1,B")

	case 0x49:
		s.BITU3R8(1, "C")
		s.CurrOpcode = fmt.Sprintf("BIT 1,C")

	case 0x4A:
		s.BITU3R8(1, "D")
		s.CurrOpcode = fmt.Sprintf("BIT 1,D")

	case 0x4B:
		s.BITU3R8(1, "E")
		s.CurrOpcode = fmt.Sprintf("BIT 1,E")

	case 0x4C:
		s.BITU3R8(1, "H")
		s.CurrOpcode = fmt.Sprintf("BIT 1,H")

	case 0x4D:
		s.BITU3R8(1, "L")
		s.CurrOpcode = fmt.Sprintf("BIT 1,L")

	case 0x4E:
		s.BITU3HL(1)
		s.CurrOpcode = fmt.Sprintf("BIT 1,(HL)")

	case 0x4F:
		s.BITU3R8(1, "A")
		s.CurrOpcode = fmt.Sprintf("BIT 1,A")

	case 0x50:
		s.BITU3R8(2, "B")
		s.CurrOpcode = fmt.Sprintf("BIT 2,B")

	case 0x51:
		s.BITU3R8(2, "C")
		s.CurrOpcode = fmt.Sprintf("BIT 2,C")

	case 0x52:
		s.BITU3R8(2, "D")
		s.CurrOpcode = fmt.Sprintf("BIT 2,D")

	case 0x53:
		s.BITU3R8(2, "E")
		s.CurrOpcode = fmt.Sprintf("BIT 2,E")

	case 0x54:
		s.BITU3R8(2, "H")
		s.CurrOpcode = fmt.Sprintf("BIT 2,H")

	case 0x55:
		s.BITU3R8(2, "L")
		s.CurrOpcode = fmt.Sprintf("BIT 2,L")

	case 0x56:
		s.BITU3HL(2)
		s.CurrOpcode = fmt.Sprintf("BIT 2,(HL)")

	case 0x57:
		s.BITU3R8(2, "A")
		s.CurrOpcode = fmt.Sprintf("BIT 2,A")

	case 0x58:
		s.BITU3R8(3, "B")
		s.CurrOpcode = fmt.Sprintf("BIT 3,B")

	case 0x59:
		s.BITU3R8(3, "C")
		s.CurrOpcode = fmt.Sprintf("BIT 3,C")

	case 0x5A:
		s.BITU3R8(3, "D")
		s.CurrOpcode = fmt.Sprintf("BIT 3,D")

	case 0x5B:
		s.BITU3R8(3, "E")
		s.CurrOpcode = fmt.Sprintf("BIT 3,E")

	case 0x5C:
		s.BITU3R8(3, "H")
		s.CurrOpcode = fmt.Sprintf("BIT 3,H")

	case 0x5D:
		s.BITU3R8(3, "L")
		s.CurrOpcode = fmt.Sprintf("BIT 3,L")

	case 0x5E:
		s.BITU3HL(3)
		s.CurrOpcode = fmt.Sprintf("BIT 3,(HL)")

	case 0x5F:
		s.BITU3R8(3, "A")
		s.CurrOpcode = fmt.Sprintf("BIT 3,A")

	case 0x60:
		s.BITU3R8(4, "B")
		s.CurrOpcode = fmt.Sprintf("BIT 4,B")

	case 0x61:
		s.BITU3R8(4, "C")
		s.CurrOpcode = fmt.Sprintf("BIT 4,C")

	case 0x62:
		s.BITU3R8(4, "D")
		s.CurrOpcode = fmt.Sprintf("BIT 4,D")

	case 0x63:
		s.BITU3R8(4, "E")
		s.CurrOpcode = fmt.Sprintf("BIT 4,E")

	case 0x64:
		s.BITU3R8(4, "H")
		s.CurrOpcode = fmt.Sprintf("BIT 4,H")

	case 0x65:
		s.BITU3R8(4, "L")
		s.CurrOpcode = fmt.Sprintf("BIT 4,L")

	case 0x66:
		s.BITU3HL(4)
		s.CurrOpcode = fmt.Sprintf("BIT 4,(HL)")

	case 0x67:
		s.BITU3R8(4, "A")
		s.CurrOpcode = fmt.Sprintf("BIT 4,A")

	case 0x68:
		s.BITU3R8(5, "B")
		s.CurrOpcode = fmt.Sprintf("BIT 5,B")

	case 0x69:
		s.BITU3R8(5, "C")
		s.CurrOpcode = fmt.Sprintf("BIT 5,C")

	case 0x6A:
		s.BITU3R8(5, "D")
		s.CurrOpcode = fmt.Sprintf("BIT 5,D")

	case 0x6B:
		s.BITU3R8(5, "E")
		s.CurrOpcode = fmt.Sprintf("BIT 5,E")

	case 0x6C:
		s.BITU3R8(5, "H")
		s.CurrOpcode = fmt.Sprintf("BIT 5,H")

	case 0x6D:
		s.BITU3R8(5, "L")
		s.CurrOpcode = fmt.Sprintf("BIT 5,L")

	case 0x6E:
		s.BITU3HL(5)
		s.CurrOpcode = fmt.Sprintf("BIT 5,(HL)")

	case 0x6F:
		s.BITU3R8(5, "A")
		s.CurrOpcode = fmt.Sprintf("BIT 5,A")

	case 0x70:
		s.BITU3R8(6, "B")
		s.CurrOpcode = fmt.Sprintf("BIT 6,B")

	case 0x71:
		s.BITU3R8(6, "C")
		s.CurrOpcode = fmt.Sprintf("BIT 6,C")

	case 0x72:
		s.BITU3R8(6, "D")
		s.CurrOpcode = fmt.Sprintf("BIT 6,D")

	case 0x73:
		s.BITU3R8(6, "E")
		s.CurrOpcode = fmt.Sprintf("BIT 6,E")

	case 0x74:
		s.BITU3R8(6, "H")
		s.CurrOpcode = fmt.Sprintf("BIT 6,H")

	case 0x75:
		s.BITU3R8(6, "L")
		s.CurrOpcode = fmt.Sprintf("BIT 6,L")

	case 0x76:
		s.BITU3HL(6)
		s.CurrOpcode = fmt.Sprintf("BIT 6,(HL)")

	case 0x77:
		s.BITU3R8(6, "A")
		s.CurrOpcode = fmt.Sprintf("BIT 6,A")

	case 0x78:
		s.BITU3R8(7, "B")
		s.CurrOpcode = fmt.Sprintf("BIT 7,B")

	case 0x79:
		s.BITU3R8(7, "C")
		s.CurrOpcode = fmt.Sprintf("BIT 7,C")

	case 0x7A:
		s.BITU3R8(7, "D")
		s.CurrOpcode = fmt.Sprintf("BIT 7,D")

	case 0x7B:
		s.BITU3R8(7, "E")
		s.CurrOpcode = fmt.Sprintf("BIT 7,E")

	case 0x7C:
		s.BITU3R8(7, "H")
		s.CurrOpcode = fmt.Sprintf("BIT 7,H")

	case 0x7D:
		s.BITU3R8(7, "L")
		s.CurrOpcode = fmt.Sprintf("BIT 7,L")

	case 0x7E:
		s.BITU3HL(7)
		s.CurrOpcode = fmt.Sprintf("BIT 7,(HL)")

	case 0x7F:
		s.BITU3R8(7, "A")
		s.CurrOpcode = fmt.Sprintf("BIT 7,A")

	case 0x80:
		s.RESU3R8(0, "B")
		s.CurrOpcode = fmt.Sprintf("RES 0,B")

	case 0x81:
		s.RESU3R8(0, "C")
		s.CurrOpcode = fmt.Sprintf("RES 0,C")

	case 0x82:
		s.RESU3R8(0, "D")
		s.CurrOpcode = fmt.Sprintf("RES 0,D")

	case 0x83:
		s.RESU3R8(0, "E")
		s.CurrOpcode = fmt.Sprintf("RES 0,E")

	case 0x84:
		s.RESU3R8(0, "H")
		s.CurrOpcode = fmt.Sprintf("RES 0,H")

	case 0x85:
		s.RESU3R8(0, "L")
		s.CurrOpcode = fmt.Sprintf("RES 0,L")

	case 0x86:
		s.RESU3HL(0)
		s.CurrOpcode = fmt.Sprintf("RES 0,(HL)")

	case 0x87:
		s.RESU3R8(0, "A")
		s.CurrOpcode = fmt.Sprintf("RES 0,A")

	case 0x88:
		s.RESU3R8(1, "B")
		s.CurrOpcode = fmt.Sprintf("RES 1,B")

	case 0x89:
		s.RESU3R8(1, "C")
		s.CurrOpcode = fmt.Sprintf("RES 1,C")

	case 0x8A:
		s.RESU3R8(1, "D")
		s.CurrOpcode = fmt.Sprintf("RES 1,D")

	case 0x8B:
		s.RESU3R8(1, "E")
		s.CurrOpcode = fmt.Sprintf("RES 1,E")

	case 0x8C:
		s.RESU3R8(1, "H")
		s.CurrOpcode = fmt.Sprintf("RES 1,H")

	case 0x8D:
		s.RESU3R8(1, "L")
		s.CurrOpcode = fmt.Sprintf("RES 1,L")

	case 0x8E:
		s.RESU3HL(1)
		s.CurrOpcode = fmt.Sprintf("RES 1,(HL)")

	case 0x8F:
		s.RESU3R8(1, "A")
		s.CurrOpcode = fmt.Sprintf("RES 1,A")

	case 0x90:
		s.RESU3R8(2, "B")
		s.CurrOpcode = fmt.Sprintf("RES 2,B")

	case 0x91:
		s.RESU3R8(2, "C")
		s.CurrOpcode = fmt.Sprintf("RES 2,C")

	case 0x92:
		s.RESU3R8(2, "D")
		s.CurrOpcode = fmt.Sprintf("RES 2,D")

	case 0x93:
		s.RESU3R8(2, "E")
		s.CurrOpcode = fmt.Sprintf("RES 2,E")

	case 0x94:
		s.RESU3R8(2, "H")
		s.CurrOpcode = fmt.Sprintf("RES 2,H")

	case 0x95:
		s.RESU3R8(2, "L")
		s.CurrOpcode = fmt.Sprintf("RES 2,L")

	case 0x96:
		s.RESU3HL(2)
		s.CurrOpcode = fmt.Sprintf("RES 2,(HL)")

	case 0x97:
		s.RESU3R8(2, "A")
		s.CurrOpcode = fmt.Sprintf("RES 2,A")

	case 0x98:
		s.RESU3R8(3, "B")
		s.CurrOpcode = fmt.Sprintf("RES 3,B")

	case 0x99:
		s.RESU3R8(3, "C")
		s.CurrOpcode = fmt.Sprintf("RES 3,C")

	case 0x9A:
		s.RESU3R8(3, "D")
		s.CurrOpcode = fmt.Sprintf("RES 3,D")

	case 0x9B:
		s.RESU3R8(3, "E")
		s.CurrOpcode = fmt.Sprintf("RES 3,E")

	case 0x9C:
		s.RESU3R8(3, "H")
		s.CurrOpcode = fmt.Sprintf("RES 3,H")

	case 0x9D:
		s.RESU3R8(3, "L")
		s.CurrOpcode = fmt.Sprintf("RES 3,L")

	case 0x9E:
		s.RESU3HL(3)
		s.CurrOpcode = fmt.Sprintf("RES 3,(HL)")

	case 0x9F:
		s.RESU3R8(3, "A")
		s.CurrOpcode = fmt.Sprintf("RES 3,A")

	case 0xA0:
		s.RESU3R8(4, "B")
		s.CurrOpcode = fmt.Sprintf("RES 4,B")

	case 0xA1:
		s.RESU3R8(4, "C")
		s.CurrOpcode = fmt.Sprintf("RES 4,C")

	case 0xA2:
		s.RESU3R8(4, "D")
		s.CurrOpcode = fmt.Sprintf("RES 4,D")

	case 0xA3:
		s.RESU3R8(4, "E")
		s.CurrOpcode = fmt.Sprintf("RES 4,E")

	case 0xA4:
		s.RESU3R8(4, "H")
		s.CurrOpcode = fmt.Sprintf("RES 4,H")

	case 0xA5:
		s.RESU3R8(4, "L")
		s.CurrOpcode = fmt.Sprintf("RES 4,L")

	case 0xA6:
		s.RESU3HL(4)
		s.CurrOpcode = fmt.Sprintf("RES A,(HL)")

	case 0xA7:
		s.RESU3R8(4, "A")
		s.CurrOpcode = fmt.Sprintf("RES 4,A")

	case 0xA8:
		s.RESU3R8(5, "B")
		s.CurrOpcode = fmt.Sprintf("RES 5,B")

	case 0xA9:
		s.RESU3R8(5, "C")
		s.CurrOpcode = fmt.Sprintf("RES 5,C")

	case 0xAA:
		s.RESU3R8(5, "D")
		s.CurrOpcode = fmt.Sprintf("RES 5,D")

	case 0xAB:
		s.RESU3R8(5, "E")
		s.CurrOpcode = fmt.Sprintf("RES 5,E")

	case 0xAC:
		s.RESU3R8(5, "H")
		s.CurrOpcode = fmt.Sprintf("RES 5,H")

	case 0xAD:
		s.RESU3R8(5, "L")
		s.CurrOpcode = fmt.Sprintf("RES 5,L")

	case 0xAE:
		s.RESU3HL(5)
		s.CurrOpcode = fmt.Sprintf("RES 5,(HL)")

	case 0xAF:
		s.RESU3R8(5, "A")
		s.CurrOpcode = fmt.Sprintf("RES 5,A")

	case 0xB0:
		s.RESU3R8(6, "B")
		s.CurrOpcode = fmt.Sprintf("RES 6,B")

	case 0xB1:
		s.RESU3R8(6, "C")
		s.CurrOpcode = fmt.Sprintf("RES 6,C")

	case 0xB2:
		s.RESU3R8(6, "D")
		s.CurrOpcode = fmt.Sprintf("RES 6,D")

	case 0xB3:
		s.RESU3R8(6, "E")
		s.CurrOpcode = fmt.Sprintf("RES 6,E")

	case 0xB4:
		s.RESU3R8(6, "H")
		s.CurrOpcode = fmt.Sprintf("RES 6,H")

	case 0xB5:
		s.RESU3R8(6, "L")
		s.CurrOpcode = fmt.Sprintf("RES 6,L")

	case 0xB6:
		s.RESU3HL(6)
		s.CurrOpcode = fmt.Sprintf("RES 6,(HL)")

	case 0xB7:
		s.RESU3R8(6, "A")
		s.CurrOpcode = fmt.Sprintf("RES 6,A")

	case 0xB8:
		s.RESU3R8(7, "B")
		s.CurrOpcode = fmt.Sprintf("RES 7,B")

	case 0xB9:
		s.RESU3R8(7, "C")
		s.CurrOpcode = fmt.Sprintf("RES 7,C")

	case 0xBA:
		s.RESU3R8(7, "D")
		s.CurrOpcode = fmt.Sprintf("RES 7,D")

	case 0xBB:
		s.RESU3R8(7, "E")
		s.CurrOpcode = fmt.Sprintf("RES 7,E")

	case 0xBC:
		s.RESU3R8(7, "H")
		s.CurrOpcode = fmt.Sprintf("RES 7,H")

	case 0xBD:
		s.RESU3R8(7, "L")
		s.CurrOpcode = fmt.Sprintf("RES 7,L")

	case 0xBE:
		s.RESU3HL(7)
		s.CurrOpcode = fmt.Sprintf("RES 7,(HL)")

	case 0xBF:
		s.RESU3R8(7, "A")
		s.CurrOpcode = fmt.Sprintf("RES 7,A")

	case 0xC0:
		s.SETU3R8(0, "B")
		s.CurrOpcode = fmt.Sprintf("SET 0,B")

	case 0xC1:
		s.SETU3R8(0, "C")
		s.CurrOpcode = fmt.Sprintf("SET 0,C")

	case 0xC2:
		s.SETU3R8(0, "D")
		s.CurrOpcode = fmt.Sprintf("SET 0,D")

	case 0xC3:
		s.SETU3R8(0, "E")
		s.CurrOpcode = fmt.Sprintf("SET 0,E")

	case 0xC4:
		s.SETU3R8(0, "H")
		s.CurrOpcode = fmt.Sprintf("SET 0,H")

	case 0xC5:
		s.SETU3R8(0, "L")
		s.CurrOpcode = fmt.Sprintf("SET 0,L")

	case 0xC6:
		s.SETU3HL(0)
		s.CurrOpcode = fmt.Sprintf("SET 0,(HL)")

	case 0xC7:
		s.SETU3R8(0, "A")
		s.CurrOpcode = fmt.Sprintf("SET 0,A")

	case 0xC8:
		s.SETU3R8(1, "B")
		s.CurrOpcode = fmt.Sprintf("SET 1,B")

	case 0xC9:
		s.SETU3R8(1, "C")
		s.CurrOpcode = fmt.Sprintf("SET 1,C")

	case 0xCA:
		s.SETU3R8(1, "D")
		s.CurrOpcode = fmt.Sprintf("SET 1,D")

	case 0xCB:
		s.SETU3R8(1, "E")
		s.CurrOpcode = fmt.Sprintf("SET 1,E")

	case 0xCC:
		s.SETU3R8(1, "H")
		s.CurrOpcode = fmt.Sprintf("SET 1,H")

	case 0xCD:
		s.SETU3R8(1, "L")
		s.CurrOpcode = fmt.Sprintf("SET 1,L")

	case 0xCE:
		s.SETU3HL(1)
		s.CurrOpcode = fmt.Sprintf("SET 1,(HL)")

	case 0xCF:
		s.SETU3R8(1, "A")
		s.CurrOpcode = fmt.Sprintf("SET 1,A")

	case 0xD0:
		s.SETU3R8(2, "B")
		s.CurrOpcode = fmt.Sprintf("SET 2,B")

	case 0xD1:
		s.SETU3R8(2, "C")
		s.CurrOpcode = fmt.Sprintf("SET 2,C")

	case 0xD2:
		s.SETU3R8(2, "D")
		s.CurrOpcode = fmt.Sprintf("SET 2,D")

	case 0xD3:
		s.SETU3R8(2, "E")
		s.CurrOpcode = fmt.Sprintf("SET 2,E")

	case 0xD4:
		s.SETU3R8(2, "H")
		s.CurrOpcode = fmt.Sprintf("SET 2,H")

	case 0xD5:
		s.SETU3R8(2, "L")
		s.CurrOpcode = fmt.Sprintf("SET 2,L")

	case 0xD6:
		s.SETU3HL(2)
		s.CurrOpcode = fmt.Sprintf("SET 2,(HL)")

	case 0xD7:
		s.SETU3R8(2, "A")
		s.CurrOpcode = fmt.Sprintf("SET 2,A")

	case 0xD8:
		s.SETU3R8(3, "B")
		s.CurrOpcode = fmt.Sprintf("SET 3,B")

	case 0xD9:
		s.SETU3R8(3, "C")
		s.CurrOpcode = fmt.Sprintf("SET 3,C")

	case 0xDA:
		s.SETU3R8(3, "D")
		s.CurrOpcode = fmt.Sprintf("SET 3,D")

	case 0xDB:
		s.SETU3R8(3, "E")
		s.CurrOpcode = fmt.Sprintf("SET 3,E")

	case 0xDC:
		s.SETU3R8(3, "H")
		s.CurrOpcode = fmt.Sprintf("SET 3,H")

	case 0xDD:
		s.SETU3R8(3, "L")
		s.CurrOpcode = fmt.Sprintf("SET 3,L")

	case 0xDE:
		s.SETU3HL(3)
		s.CurrOpcode = fmt.Sprintf("SET 3,(HL)")

	case 0xDF:
		s.SETU3R8(3, "A")
		s.CurrOpcode = fmt.Sprintf("SET 3,A")

	case 0xE0:
		s.SETU3R8(4, "B")
		s.CurrOpcode = fmt.Sprintf("SET 4,B")

	case 0xE1:
		s.SETU3R8(4, "C")
		s.CurrOpcode = fmt.Sprintf("SET 4,C")

	case 0xE2:
		s.SETU3R8(4, "D")
		s.CurrOpcode = fmt.Sprintf("SET 4,D")

	case 0xE3:
		s.SETU3R8(4, "E")
		s.CurrOpcode = fmt.Sprintf("SET 4,E")

	case 0xE4:
		s.SETU3R8(4, "H")
		s.CurrOpcode = fmt.Sprintf("SET 4,H")

	case 0xE5:
		s.SETU3R8(4, "L")
		s.CurrOpcode = fmt.Sprintf("SET 4,L")

	case 0xE6:
		s.SETU3HL(4)
		s.CurrOpcode = fmt.Sprintf("SET 4,(HL)")

	case 0xE7:
		s.SETU3R8(4, "A")
		s.CurrOpcode = fmt.Sprintf("SET 4,A")

	case 0xE8:
		s.SETU3R8(5, "B")
		s.CurrOpcode = fmt.Sprintf("SET 5,B")

	case 0xE9:
		s.SETU3R8(5, "C")
		s.CurrOpcode = fmt.Sprintf("SET 5,C")

	case 0xEA:
		s.SETU3R8(5, "D")
		s.CurrOpcode = fmt.Sprintf("SET 5,D")

	case 0xEB:
		s.SETU3R8(5, "E")
		s.CurrOpcode = fmt.Sprintf("SET 5,E")

	case 0xEC:
		s.SETU3R8(5, "H")
		s.CurrOpcode = fmt.Sprintf("SET 5,H")

	case 0xED:
		s.SETU3R8(5, "L")
		s.CurrOpcode = fmt.Sprintf("SET 5,L")

	case 0xEE:
		s.SETU3HL(5)
		s.CurrOpcode = fmt.Sprintf("SET 5,(HL)")

	case 0xEF:
		s.SETU3R8(5, "A")
		s.CurrOpcode = fmt.Sprintf("SET 5,A")

	case 0xF0:
		s.SETU3R8(6, "B")
		s.CurrOpcode = fmt.Sprintf("SET 6,B")

	case 0xF1:
		s.SETU3R8(6, "C")
		s.CurrOpcode = fmt.Sprintf("SET 6,C")

	case 0xF2:
		s.SETU3R8(6, "D")
		s.CurrOpcode = fmt.Sprintf("SET 6,D")

	case 0xF3:
		s.SETU3R8(6, "E")
		s.CurrOpcode = fmt.Sprintf("SET 6,E")

	case 0xF4:
		s.SETU3R8(6, "H")
		s.CurrOpcode = fmt.Sprintf("SET 6,H")

	case 0xF5:
		s.SETU3R8(6, "L")
		s.CurrOpcode = fmt.Sprintf("SET 6,L")

	case 0xF6:
		s.SETU3HL(6)
		s.CurrOpcode = fmt.Sprintf("SET 6,(HL)")

	case 0xF7:
		s.SETU3R8(6, "A")
		s.CurrOpcode = fmt.Sprintf("SET 6,A")

	case 0xF8:
		s.SETU3R8(7, "B")
		s.CurrOpcode = fmt.Sprintf("SET 7,B")

	case 0xF9:
		s.SETU3R8(7, "C")
		s.CurrOpcode = fmt.Sprintf("SET 7,C")

	case 0xFA:
		s.SETU3R8(7, "D")
		s.CurrOpcode = fmt.Sprintf("SET 7,D")

	case 0xFB:
		s.SETU3R8(7, "E")
		s.CurrOpcode = fmt.Sprintf("SET 7,E")

	case 0xFC:
		s.SETU3R8(7, "H")
		s.CurrOpcode = fmt.Sprintf("SET 7,H")

	case 0xFD:
		s.SETU3R8(7, "L")
		s.CurrOpcode = fmt.Sprintf("SET 7,L")

	case 0xFE:
		s.SETU3HL(7)
		s.CurrOpcode = fmt.Sprintf("SET 7,(HL)")

	case 0xFF:
		s.SETU3R8(7, "A")
		s.CurrOpcode = fmt.Sprintf("SET 7,A")

	}
}
