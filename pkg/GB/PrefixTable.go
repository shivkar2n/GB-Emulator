package GB

func (GB *GB) PrefixTable() (string, int, int, bool) {
	opCode := uint8(GB.MMU.Read(GB.CPU.Reg.Read("PC") + 1))

	switch opCode {
	case 0x00:
		return GB.RLCR8("B")

	case 0x01:
		return GB.RLCR8("C")

	case 0x02:
		return GB.RLCR8("D")

	case 0x03:
		return GB.RLCR8("E")

	case 0x04:
		return GB.RLCR8("H")

	case 0x05:
		return GB.RLCR8("L")

	case 0x06:
		return GB.RLCHL()

	case 0x07:
		return GB.RLCR8("A")

	case 0x08:
		return GB.RRCR8("B")

	case 0x09:
		return GB.RRCR8("C")

	case 0x0A:
		return GB.RRCR8("D")

	case 0x0B:
		return GB.RRCR8("E")

	case 0x0C:
		return GB.RRCR8("H")

	case 0x0D:
		return GB.RRCR8("L")

	case 0x0E:
		return GB.RRCHL()

	case 0x0F:
		return GB.RRCR8("A")

	case 0x10:
		return GB.RLR8("B")

	case 0x11:
		return GB.RLR8("C")

	case 0x12:
		return GB.RLR8("D")

	case 0x13:
		return GB.RLR8("E")

	case 0x14:
		return GB.RLR8("H")

	case 0x15:
		return GB.RLR8("L")

	case 0x16:
		return GB.RLHL()

	case 0x17:
		return GB.RLR8("A")

	case 0x18:
		return GB.RRR8("B")

	case 0x19:
		return GB.RRR8("C")

	case 0x1A:
		return GB.RRR8("D")

	case 0x1B:
		return GB.RRR8("E")

	case 0x1C:
		return GB.RRR8("H")

	case 0x1D:
		return GB.RRR8("L")

	case 0x1E:
		return GB.RRHL()

	case 0x1F:
		return GB.RRR8("A")

	case 0x20:
		return GB.SLAR8("B")

	case 0x21:
		return GB.SLAR8("C")

	case 0x22:
		return GB.SLAR8("D")

	case 0x23:
		return GB.SLAR8("E")

	case 0x24:
		return GB.SLAR8("H")

	case 0x25:
		return GB.SLAR8("L")

	case 0x26:
		return GB.SLAHL()

	case 0x27:
		return GB.SLAR8("A")

	case 0x28:
		return GB.SRAR8("B")

	case 0x29:
		return GB.SRAR8("C")

	case 0x2A:
		return GB.SRAR8("D")

	case 0x2B:
		return GB.SRAR8("E")

	case 0x2C:
		return GB.SRAR8("H")

	case 0x2D:
		return GB.SRAR8("L")

	case 0x2E:
		return GB.SRAHL()

	case 0x2F:
		return GB.SRAR8("A")

	case 0x30:
		return GB.SWAPR8("B")

	case 0x31:
		return GB.SWAPR8("C")

	case 0x32:
		return GB.SWAPR8("D")

	case 0x33:
		return GB.SWAPR8("E")

	case 0x34:
		return GB.SWAPR8("H")

	case 0x35:
		return GB.SWAPR8("L")

	case 0x36:
		return GB.SWAPHL()

	case 0x37:
		return GB.SWAPR8("A")

	case 0x38:
		return GB.SRLR8("B")

	case 0x39:
		return GB.SRLR8("C")

	case 0x3A:
		return GB.SRLR8("D")

	case 0x3B:
		return GB.SRLR8("E")

	case 0x3C:
		return GB.SRLR8("H")

	case 0x3D:
		return GB.SRLR8("L")

	case 0x3E:
		return GB.SRLHL()

	case 0x3F:
		return GB.SRLR8("A")

	case 0x40:
		return GB.BITU3R8(0, "B")

	case 0x41:
		return GB.BITU3R8(0, "C")

	case 0x42:
		return GB.BITU3R8(0, "D")

	case 0x43:
		return GB.BITU3R8(0, "E")

	case 0x44:
		return GB.BITU3R8(0, "H")

	case 0x45:
		return GB.BITU3R8(0, "L")

	case 0x46:
		return GB.BITU3HL(0)

	case 0x47:
		return GB.BITU3R8(0, "A")

	case 0x48:
		return GB.BITU3R8(1, "B")

	case 0x49:
		return GB.BITU3R8(1, "C")

	case 0x4A:
		return GB.BITU3R8(1, "D")

	case 0x4B:
		return GB.BITU3R8(1, "E")

	case 0x4C:
		return GB.BITU3R8(1, "H")

	case 0x4D:
		return GB.BITU3R8(1, "L")

	case 0x4E:
		return GB.BITU3HL(1)

	case 0x4F:
		return GB.BITU3R8(1, "A")

	case 0x50:
		return GB.BITU3R8(2, "B")

	case 0x51:
		return GB.BITU3R8(2, "C")

	case 0x52:
		return GB.BITU3R8(2, "D")

	case 0x53:
		return GB.BITU3R8(2, "E")

	case 0x54:
		return GB.BITU3R8(2, "H")

	case 0x55:
		return GB.BITU3R8(2, "L")

	case 0x56:
		return GB.BITU3HL(2)

	case 0x57:
		return GB.BITU3R8(2, "A")

	case 0x58:
		return GB.BITU3R8(3, "B")

	case 0x59:
		return GB.BITU3R8(3, "C")

	case 0x5A:
		return GB.BITU3R8(3, "D")

	case 0x5B:
		return GB.BITU3R8(3, "E")

	case 0x5C:
		return GB.BITU3R8(3, "H")

	case 0x5D:
		return GB.BITU3R8(3, "L")

	case 0x5E:
		return GB.BITU3HL(3)

	case 0x5F:
		return GB.BITU3R8(3, "A")

	case 0x60:
		return GB.BITU3R8(4, "B")

	case 0x61:
		return GB.BITU3R8(4, "C")

	case 0x62:
		return GB.BITU3R8(4, "D")

	case 0x63:
		return GB.BITU3R8(4, "E")

	case 0x64:
		return GB.BITU3R8(4, "H")

	case 0x65:
		return GB.BITU3R8(4, "L")

	case 0x66:
		return GB.BITU3HL(4)

	case 0x67:
		return GB.BITU3R8(4, "A")

	case 0x68:
		return GB.BITU3R8(5, "B")

	case 0x69:
		return GB.BITU3R8(5, "C")

	case 0x6A:
		return GB.BITU3R8(5, "D")

	case 0x6B:
		return GB.BITU3R8(5, "E")

	case 0x6C:
		return GB.BITU3R8(5, "H")

	case 0x6D:
		return GB.BITU3R8(5, "L")

	case 0x6E:
		return GB.BITU3HL(5)

	case 0x6F:
		return GB.BITU3R8(5, "A")

	case 0x70:
		return GB.BITU3R8(6, "B")

	case 0x71:
		return GB.BITU3R8(6, "C")

	case 0x72:
		return GB.BITU3R8(6, "D")

	case 0x73:
		return GB.BITU3R8(6, "E")

	case 0x74:
		return GB.BITU3R8(6, "H")

	case 0x75:
		return GB.BITU3R8(6, "L")

	case 0x76:
		return GB.BITU3HL(6)

	case 0x77:
		return GB.BITU3R8(6, "A")

	case 0x78:
		return GB.BITU3R8(7, "B")

	case 0x79:
		return GB.BITU3R8(7, "C")

	case 0x7A:
		return GB.BITU3R8(7, "D")

	case 0x7B:
		return GB.BITU3R8(7, "E")

	case 0x7C:
		return GB.BITU3R8(7, "H")

	case 0x7D:
		return GB.BITU3R8(7, "L")

	case 0x7E:
		return GB.BITU3HL(7)

	case 0x7F:
		return GB.BITU3R8(7, "A")

	case 0x80:
		return GB.RESU3R8(0, "B")

	case 0x81:
		return GB.RESU3R8(0, "C")

	case 0x82:
		return GB.RESU3R8(0, "D")

	case 0x83:
		return GB.RESU3R8(0, "E")

	case 0x84:
		return GB.RESU3R8(0, "H")

	case 0x85:
		return GB.RESU3R8(0, "L")

	case 0x86:
		return GB.RESU3HL(0)

	case 0x87:
		return GB.RESU3R8(0, "A")

	case 0x88:
		return GB.RESU3R8(1, "B")

	case 0x89:
		return GB.RESU3R8(1, "C")

	case 0x8A:
		return GB.RESU3R8(1, "D")

	case 0x8B:
		return GB.RESU3R8(1, "E")

	case 0x8C:
		return GB.RESU3R8(1, "H")

	case 0x8D:
		return GB.RESU3R8(1, "L")

	case 0x8E:
		return GB.RESU3HL(1)

	case 0x8F:
		return GB.RESU3R8(1, "A")

	case 0x90:
		return GB.RESU3R8(2, "B")

	case 0x91:
		return GB.RESU3R8(2, "C")

	case 0x92:
		return GB.RESU3R8(2, "D")

	case 0x93:
		return GB.RESU3R8(2, "E")

	case 0x94:
		return GB.RESU3R8(2, "H")

	case 0x95:
		return GB.RESU3R8(2, "L")

	case 0x96:
		return GB.RESU3HL(2)

	case 0x97:
		return GB.RESU3R8(2, "A")

	case 0x98:
		return GB.RESU3R8(3, "B")

	case 0x99:
		return GB.RESU3R8(3, "C")

	case 0x9A:
		return GB.RESU3R8(3, "D")

	case 0x9B:
		return GB.RESU3R8(3, "E")

	case 0x9C:
		return GB.RESU3R8(3, "H")

	case 0x9D:
		return GB.RESU3R8(3, "L")

	case 0x9E:
		return GB.RESU3HL(3)

	case 0x9F:
		return GB.RESU3R8(3, "A")

	case 0xA0:
		return GB.RESU3R8(4, "B")

	case 0xA1:
		return GB.RESU3R8(4, "C")

	case 0xA2:
		return GB.RESU3R8(4, "D")

	case 0xA3:
		return GB.RESU3R8(4, "E")

	case 0xA4:
		return GB.RESU3R8(4, "H")

	case 0xA5:
		return GB.RESU3R8(4, "L")

	case 0xA6:
		return GB.RESU3HL(4)

	case 0xA7:
		return GB.RESU3R8(4, "A")

	case 0xA8:
		return GB.RESU3R8(5, "B")

	case 0xA9:
		return GB.RESU3R8(5, "C")

	case 0xAA:
		return GB.RESU3R8(5, "D")

	case 0xAB:
		return GB.RESU3R8(5, "E")

	case 0xAC:
		return GB.RESU3R8(5, "H")

	case 0xAD:
		return GB.RESU3R8(5, "L")

	case 0xAE:
		return GB.RESU3HL(5)

	case 0xAF:
		return GB.RESU3R8(5, "A")

	case 0xB0:
		return GB.RESU3R8(6, "B")

	case 0xB1:
		return GB.RESU3R8(6, "C")

	case 0xB2:
		return GB.RESU3R8(6, "D")

	case 0xB3:
		return GB.RESU3R8(6, "E")

	case 0xB4:
		return GB.RESU3R8(6, "H")

	case 0xB5:
		return GB.RESU3R8(6, "L")

	case 0xB6:
		return GB.RESU3HL(6)

	case 0xB7:
		return GB.RESU3R8(6, "A")

	case 0xB8:
		return GB.RESU3R8(7, "B")

	case 0xB9:
		return GB.RESU3R8(7, "C")

	case 0xBA:
		return GB.RESU3R8(7, "D")

	case 0xBB:
		return GB.RESU3R8(7, "E")

	case 0xBC:
		return GB.RESU3R8(7, "H")

	case 0xBD:
		return GB.RESU3R8(7, "L")

	case 0xBE:
		return GB.RESU3HL(7)

	case 0xBF:
		return GB.RESU3R8(7, "A")

	case 0xC0:
		return GB.SETU3R8(0, "B")

	case 0xC1:
		return GB.SETU3R8(0, "C")

	case 0xC2:
		return GB.SETU3R8(0, "D")

	case 0xC3:
		return GB.SETU3R8(0, "E")

	case 0xC4:
		return GB.SETU3R8(0, "H")

	case 0xC5:
		return GB.SETU3R8(0, "L")

	case 0xC6:
		return GB.SETU3HL(0)

	case 0xC7:
		return GB.SETU3R8(0, "A")

	case 0xC8:
		return GB.SETU3R8(1, "B")

	case 0xC9:
		return GB.SETU3R8(1, "C")

	case 0xCA:
		return GB.SETU3R8(1, "D")

	case 0xCB:
		return GB.SETU3R8(1, "E")

	case 0xCC:
		return GB.SETU3R8(1, "H")

	case 0xCD:
		return GB.SETU3R8(1, "L")

	case 0xCE:
		return GB.SETU3HL(1)

	case 0xCF:
		return GB.SETU3R8(1, "A")

	case 0xD0:
		return GB.SETU3R8(2, "B")

	case 0xD1:
		return GB.SETU3R8(2, "C")

	case 0xD2:
		return GB.SETU3R8(2, "D")

	case 0xD3:
		return GB.SETU3R8(2, "E")

	case 0xD4:
		return GB.SETU3R8(2, "H")

	case 0xD5:
		return GB.SETU3R8(2, "L")

	case 0xD6:
		return GB.SETU3HL(2)

	case 0xD7:
		return GB.SETU3R8(2, "A")

	case 0xD8:
		return GB.SETU3R8(3, "B")

	case 0xD9:
		return GB.SETU3R8(3, "C")

	case 0xDA:
		return GB.SETU3R8(3, "D")

	case 0xDB:
		return GB.SETU3R8(3, "E")

	case 0xDC:
		return GB.SETU3R8(3, "H")

	case 0xDD:
		return GB.SETU3R8(3, "L")

	case 0xDE:
		return GB.SETU3HL(3)

	case 0xDF:
		return GB.SETU3R8(3, "A")

	case 0xE0:
		return GB.SETU3R8(4, "B")

	case 0xE1:
		return GB.SETU3R8(4, "C")

	case 0xE2:
		return GB.SETU3R8(4, "D")

	case 0xE3:
		return GB.SETU3R8(4, "E")

	case 0xE4:
		return GB.SETU3R8(4, "H")

	case 0xE5:
		return GB.SETU3R8(4, "L")

	case 0xE6:
		return GB.SETU3HL(4)

	case 0xE7:
		return GB.SETU3R8(4, "A")

	case 0xE8:
		return GB.SETU3R8(5, "B")

	case 0xE9:
		return GB.SETU3R8(5, "C")

	case 0xEA:
		return GB.SETU3R8(5, "D")

	case 0xEB:
		return GB.SETU3R8(5, "E")

	case 0xEC:
		return GB.SETU3R8(5, "H")

	case 0xED:
		return GB.SETU3R8(5, "L")

	case 0xEE:
		return GB.SETU3HL(5)

	case 0xEF:
		return GB.SETU3R8(5, "A")

	case 0xF0:
		return GB.SETU3R8(6, "B")

	case 0xF1:
		return GB.SETU3R8(6, "C")

	case 0xF2:
		return GB.SETU3R8(6, "D")

	case 0xF3:
		return GB.SETU3R8(6, "E")

	case 0xF4:
		return GB.SETU3R8(6, "H")

	case 0xF5:
		return GB.SETU3R8(6, "L")

	case 0xF6:
		return GB.SETU3HL(6)

	case 0xF7:
		return GB.SETU3R8(6, "A")

	case 0xF8:
		return GB.SETU3R8(7, "B")

	case 0xF9:
		return GB.SETU3R8(7, "C")

	case 0xFA:
		return GB.SETU3R8(7, "D")

	case 0xFB:
		return GB.SETU3R8(7, "E")

	case 0xFC:
		return GB.SETU3R8(7, "H")

	case 0xFD:
		return GB.SETU3R8(7, "L")

	case 0xFE:
		return GB.SETU3HL(7)

	case 0xFF:
		return GB.SETU3R8(7, "A")

	}
	return "", 0, 0, false
}
