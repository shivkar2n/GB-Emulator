package CPU

func (cpu *CPU) PrefixTable() (string, int, int, bool) {
	opCode := uint8(cpu.Mem.Read(cpu.Reg.Read("PC") + 1))

	switch opCode {
	case 0x00:
		return cpu.RLCR8("B")

	case 0x01:
		return cpu.RLCR8("C")

	case 0x02:
		return cpu.RLCR8("D")

	case 0x03:
		return cpu.RLCR8("E")

	case 0x04:
		return cpu.RLCR8("H")

	case 0x05:
		return cpu.RLCR8("L")

	case 0x06:
		return cpu.RLCHL()

	case 0x07:
		return cpu.RLCR8("A")

	case 0x08:
		return cpu.RRCR8("B")

	case 0x09:
		return cpu.RRCR8("C")

	case 0x0A:
		return cpu.RRCR8("D")

	case 0x0B:
		return cpu.RRCR8("E")

	case 0x0C:
		return cpu.RRCR8("H")

	case 0x0D:
		return cpu.RRCR8("L")

	case 0x0E:
		return cpu.RRCHL()

	case 0x0F:
		return cpu.RRCR8("A")

	case 0x10:
		return cpu.RLR8("B")

	case 0x11:
		return cpu.RLR8("C")

	case 0x12:
		return cpu.RLR8("D")

	case 0x13:
		return cpu.RLR8("E")

	case 0x14:
		return cpu.RLR8("H")

	case 0x15:
		return cpu.RLR8("L")

	case 0x16:
		return cpu.RLHL()

	case 0x17:
		return cpu.RLR8("A")

	case 0x18:
		return cpu.RRR8("B")

	case 0x19:
		return cpu.RRR8("C")

	case 0x1A:
		return cpu.RRR8("D")

	case 0x1B:
		return cpu.RRR8("E")

	case 0x1C:
		return cpu.RRR8("H")

	case 0x1D:
		return cpu.RRR8("L")

	case 0x1E:
		return cpu.RRHL()

	case 0x1F:
		return cpu.RRR8("A")

	case 0x20:
		return cpu.SLAR8("B")

	case 0x21:
		return cpu.SLAR8("C")

	case 0x22:
		return cpu.SLAR8("D")

	case 0x23:
		return cpu.SLAR8("E")

	case 0x24:
		return cpu.SLAR8("H")

	case 0x25:
		return cpu.SLAR8("L")

	case 0x26:
		return cpu.SLAHL()

	case 0x27:
		return cpu.SLAR8("A")

	case 0x28:
		return cpu.SRAR8("B")

	case 0x29:
		return cpu.SRAR8("C")

	case 0x2A:
		return cpu.SRAR8("D")

	case 0x2B:
		return cpu.SRAR8("E")

	case 0x2C:
		return cpu.SRAR8("H")

	case 0x2D:
		return cpu.SRAR8("L")

	case 0x2E:
		return cpu.SRAHL()

	case 0x2F:
		return cpu.SRAR8("A")

	case 0x30:
		return cpu.SWAPR8("B")

	case 0x31:
		return cpu.SWAPR8("C")

	case 0x32:
		return cpu.SWAPR8("D")

	case 0x33:
		return cpu.SWAPR8("E")

	case 0x34:
		return cpu.SWAPR8("H")

	case 0x35:
		return cpu.SWAPR8("L")

	case 0x36:
		return cpu.SWAPHL()

	case 0x37:
		return cpu.SWAPR8("A")

	case 0x38:
		return cpu.SRLR8("B")

	case 0x39:
		return cpu.SRLR8("C")

	case 0x3A:
		return cpu.SRLR8("D")

	case 0x3B:
		return cpu.SRLR8("E")

	case 0x3C:
		return cpu.SRLR8("H")

	case 0x3D:
		return cpu.SRLR8("L")

	case 0x3E:
		return cpu.SRLHL()

	case 0x3F:
		return cpu.SRLR8("A")

	case 0x40:
		return cpu.BITU3R8(0, "B")

	case 0x41:
		return cpu.BITU3R8(0, "C")

	case 0x42:
		return cpu.BITU3R8(0, "D")

	case 0x43:
		return cpu.BITU3R8(0, "E")

	case 0x44:
		return cpu.BITU3R8(0, "H")

	case 0x45:
		return cpu.BITU3R8(0, "L")

	case 0x46:
		return cpu.BITU3HL(0)

	case 0x47:
		return cpu.BITU3R8(0, "A")

	case 0x48:
		return cpu.BITU3R8(1, "B")

	case 0x49:
		return cpu.BITU3R8(1, "C")

	case 0x4A:
		return cpu.BITU3R8(1, "D")

	case 0x4B:
		return cpu.BITU3R8(1, "E")

	case 0x4C:
		return cpu.BITU3R8(1, "H")

	case 0x4D:
		return cpu.BITU3R8(1, "L")

	case 0x4E:
		return cpu.BITU3HL(1)

	case 0x4F:
		return cpu.BITU3R8(1, "A")

	case 0x50:
		return cpu.BITU3R8(2, "B")

	case 0x51:
		return cpu.BITU3R8(2, "C")

	case 0x52:
		return cpu.BITU3R8(2, "D")

	case 0x53:
		return cpu.BITU3R8(2, "E")

	case 0x54:
		return cpu.BITU3R8(2, "H")

	case 0x55:
		return cpu.BITU3R8(2, "L")

	case 0x56:
		return cpu.BITU3HL(2)

	case 0x57:
		return cpu.BITU3R8(2, "A")

	case 0x58:
		return cpu.BITU3R8(3, "B")

	case 0x59:
		return cpu.BITU3R8(3, "C")

	case 0x5A:
		return cpu.BITU3R8(3, "D")

	case 0x5B:
		return cpu.BITU3R8(3, "E")

	case 0x5C:
		return cpu.BITU3R8(3, "H")

	case 0x5D:
		return cpu.BITU3R8(3, "L")

	case 0x5E:
		return cpu.BITU3HL(3)

	case 0x5F:
		return cpu.BITU3R8(3, "A")

	case 0x60:
		return cpu.BITU3R8(4, "B")

	case 0x61:
		return cpu.BITU3R8(4, "C")

	case 0x62:
		return cpu.BITU3R8(4, "D")

	case 0x63:
		return cpu.BITU3R8(4, "E")

	case 0x64:
		return cpu.BITU3R8(4, "H")

	case 0x65:
		return cpu.BITU3R8(4, "L")

	case 0x66:
		return cpu.BITU3HL(4)

	case 0x67:
		return cpu.BITU3R8(4, "A")

	case 0x68:
		return cpu.BITU3R8(5, "B")

	case 0x69:
		return cpu.BITU3R8(5, "C")

	case 0x6A:
		return cpu.BITU3R8(5, "D")

	case 0x6B:
		return cpu.BITU3R8(5, "E")

	case 0x6C:
		return cpu.BITU3R8(5, "H")

	case 0x6D:
		return cpu.BITU3R8(5, "L")

	case 0x6E:
		return cpu.BITU3HL(5)

	case 0x6F:
		return cpu.BITU3R8(5, "A")

	case 0x70:
		return cpu.BITU3R8(6, "B")

	case 0x71:
		return cpu.BITU3R8(6, "C")

	case 0x72:
		return cpu.BITU3R8(6, "D")

	case 0x73:
		return cpu.BITU3R8(6, "E")

	case 0x74:
		return cpu.BITU3R8(6, "H")

	case 0x75:
		return cpu.BITU3R8(6, "L")

	case 0x76:
		return cpu.BITU3HL(6)

	case 0x77:
		return cpu.BITU3R8(6, "A")

	case 0x78:
		return cpu.BITU3R8(7, "B")

	case 0x79:
		return cpu.BITU3R8(7, "C")

	case 0x7A:
		return cpu.BITU3R8(7, "D")

	case 0x7B:
		return cpu.BITU3R8(7, "E")

	case 0x7C:
		return cpu.BITU3R8(7, "H")

	case 0x7D:
		return cpu.BITU3R8(7, "L")

	case 0x7E:
		return cpu.BITU3HL(7)

	case 0x7F:
		return cpu.BITU3R8(7, "A")

	case 0x80:
		return cpu.RESU3R8(0, "B")

	case 0x81:
		return cpu.RESU3R8(0, "C")

	case 0x82:
		return cpu.RESU3R8(0, "D")

	case 0x83:
		return cpu.RESU3R8(0, "E")

	case 0x84:
		return cpu.RESU3R8(0, "H")

	case 0x85:
		return cpu.RESU3R8(0, "L")

	case 0x86:
		return cpu.RESU3HL(0)

	case 0x87:
		return cpu.RESU3R8(0, "A")

	case 0x88:
		return cpu.RESU3R8(1, "B")

	case 0x89:
		return cpu.RESU3R8(1, "C")

	case 0x8A:
		return cpu.RESU3R8(1, "D")

	case 0x8B:
		return cpu.RESU3R8(1, "E")

	case 0x8C:
		return cpu.RESU3R8(1, "H")

	case 0x8D:
		return cpu.RESU3R8(1, "L")

	case 0x8E:
		return cpu.RESU3HL(1)

	case 0x8F:
		return cpu.RESU3R8(1, "A")

	case 0x90:
		return cpu.RESU3R8(2, "B")

	case 0x91:
		return cpu.RESU3R8(2, "C")

	case 0x92:
		return cpu.RESU3R8(2, "D")

	case 0x93:
		return cpu.RESU3R8(2, "E")

	case 0x94:
		return cpu.RESU3R8(2, "H")

	case 0x95:
		return cpu.RESU3R8(2, "L")

	case 0x96:
		return cpu.RESU3HL(2)

	case 0x97:
		return cpu.RESU3R8(2, "A")

	case 0x98:
		return cpu.RESU3R8(3, "B")

	case 0x99:
		return cpu.RESU3R8(3, "C")

	case 0x9A:
		return cpu.RESU3R8(3, "D")

	case 0x9B:
		return cpu.RESU3R8(3, "E")

	case 0x9C:
		return cpu.RESU3R8(3, "H")

	case 0x9D:
		return cpu.RESU3R8(3, "L")

	case 0x9E:
		return cpu.RESU3HL(3)

	case 0x9F:
		return cpu.RESU3R8(3, "A")

	case 0xA0:
		return cpu.RESU3R8(4, "B")

	case 0xA1:
		return cpu.RESU3R8(4, "C")

	case 0xA2:
		return cpu.RESU3R8(4, "D")

	case 0xA3:
		return cpu.RESU3R8(4, "E")

	case 0xA4:
		return cpu.RESU3R8(4, "H")

	case 0xA5:
		return cpu.RESU3R8(4, "L")

	case 0xA6:
		return cpu.RESU3HL(4)

	case 0xA7:
		return cpu.RESU3R8(4, "A")

	case 0xA8:
		return cpu.RESU3R8(5, "B")

	case 0xA9:
		return cpu.RESU3R8(5, "C")

	case 0xAA:
		return cpu.RESU3R8(5, "D")

	case 0xAB:
		return cpu.RESU3R8(5, "E")

	case 0xAC:
		return cpu.RESU3R8(5, "H")

	case 0xAD:
		return cpu.RESU3R8(5, "L")

	case 0xAE:
		return cpu.RESU3HL(5)

	case 0xAF:
		return cpu.RESU3R8(5, "A")

	case 0xB0:
		return cpu.RESU3R8(6, "B")

	case 0xB1:
		return cpu.RESU3R8(6, "C")

	case 0xB2:
		return cpu.RESU3R8(6, "D")

	case 0xB3:
		return cpu.RESU3R8(6, "E")

	case 0xB4:
		return cpu.RESU3R8(6, "H")

	case 0xB5:
		return cpu.RESU3R8(6, "L")

	case 0xB6:
		return cpu.RESU3HL(6)

	case 0xB7:
		return cpu.RESU3R8(6, "A")

	case 0xB8:
		return cpu.RESU3R8(7, "B")

	case 0xB9:
		return cpu.RESU3R8(7, "C")

	case 0xBA:
		return cpu.RESU3R8(7, "D")

	case 0xBB:
		return cpu.RESU3R8(7, "E")

	case 0xBC:
		return cpu.RESU3R8(7, "H")

	case 0xBD:
		return cpu.RESU3R8(7, "L")

	case 0xBE:
		return cpu.RESU3HL(7)

	case 0xBF:
		return cpu.RESU3R8(7, "A")

	case 0xC0:
		return cpu.SETU3R8(0, "B")

	case 0xC1:
		return cpu.SETU3R8(0, "C")

	case 0xC2:
		return cpu.SETU3R8(0, "D")

	case 0xC3:
		return cpu.SETU3R8(0, "E")

	case 0xC4:
		return cpu.SETU3R8(0, "H")

	case 0xC5:
		return cpu.SETU3R8(0, "L")

	case 0xC6:
		return cpu.SETU3HL(0)

	case 0xC7:
		return cpu.SETU3R8(0, "A")

	case 0xC8:
		return cpu.SETU3R8(1, "B")

	case 0xC9:
		return cpu.SETU3R8(1, "C")

	case 0xCA:
		return cpu.SETU3R8(1, "D")

	case 0xCB:
		return cpu.SETU3R8(1, "E")

	case 0xCC:
		return cpu.SETU3R8(1, "H")

	case 0xCD:
		return cpu.SETU3R8(1, "L")

	case 0xCE:
		return cpu.SETU3HL(1)

	case 0xCF:
		return cpu.SETU3R8(1, "A")

	case 0xD0:
		return cpu.SETU3R8(2, "B")

	case 0xD1:
		return cpu.SETU3R8(2, "C")

	case 0xD2:
		return cpu.SETU3R8(2, "D")

	case 0xD3:
		return cpu.SETU3R8(2, "E")

	case 0xD4:
		return cpu.SETU3R8(2, "H")

	case 0xD5:
		return cpu.SETU3R8(2, "L")

	case 0xD6:
		return cpu.SETU3HL(2)

	case 0xD7:
		return cpu.SETU3R8(2, "A")

	case 0xD8:
		return cpu.SETU3R8(3, "B")

	case 0xD9:
		return cpu.SETU3R8(3, "C")

	case 0xDA:
		return cpu.SETU3R8(3, "D")

	case 0xDB:
		return cpu.SETU3R8(3, "E")

	case 0xDC:
		return cpu.SETU3R8(3, "H")

	case 0xDD:
		return cpu.SETU3R8(3, "L")

	case 0xDE:
		return cpu.SETU3HL(3)

	case 0xDF:
		return cpu.SETU3R8(3, "A")

	case 0xE0:
		return cpu.SETU3R8(4, "B")

	case 0xE1:
		return cpu.SETU3R8(4, "C")

	case 0xE2:
		return cpu.SETU3R8(4, "D")

	case 0xE3:
		return cpu.SETU3R8(4, "E")

	case 0xE4:
		return cpu.SETU3R8(4, "H")

	case 0xE5:
		return cpu.SETU3R8(4, "L")

	case 0xE6:
		return cpu.SETU3HL(4)

	case 0xE7:
		return cpu.SETU3R8(4, "A")

	case 0xE8:
		return cpu.SETU3R8(5, "B")

	case 0xE9:
		return cpu.SETU3R8(5, "C")

	case 0xEA:
		return cpu.SETU3R8(5, "D")

	case 0xEB:
		return cpu.SETU3R8(5, "E")

	case 0xEC:
		return cpu.SETU3R8(5, "H")

	case 0xED:
		return cpu.SETU3R8(5, "L")

	case 0xEE:
		return cpu.SETU3HL(5)

	case 0xEF:
		return cpu.SETU3R8(5, "A")

	case 0xF0:
		return cpu.SETU3R8(6, "B")

	case 0xF1:
		return cpu.SETU3R8(6, "C")

	case 0xF2:
		return cpu.SETU3R8(6, "D")

	case 0xF3:
		return cpu.SETU3R8(6, "E")

	case 0xF4:
		return cpu.SETU3R8(6, "H")

	case 0xF5:
		return cpu.SETU3R8(6, "L")

	case 0xF6:
		return cpu.SETU3HL(6)

	case 0xF7:
		return cpu.SETU3R8(6, "A")

	case 0xF8:
		return cpu.SETU3R8(7, "B")

	case 0xF9:
		return cpu.SETU3R8(7, "C")

	case 0xFA:
		return cpu.SETU3R8(7, "D")

	case 0xFB:
		return cpu.SETU3R8(7, "E")

	case 0xFC:
		return cpu.SETU3R8(7, "H")

	case 0xFD:
		return cpu.SETU3R8(7, "L")

	case 0xFE:
		return cpu.SETU3HL(7)

	case 0xFF:
		return cpu.SETU3R8(7, "A")

	}
	return "", 0, 0, false
}
