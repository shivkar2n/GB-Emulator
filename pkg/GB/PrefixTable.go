package GB

func (GB *GB) PrefixTable() {
	GB.IncrementTimer(4)
	opCode := uint8(GB.MMU.Read(GB.CPU.Reg.Read("PC") + 1))

	switch opCode {
	case 0x00:
		GB.CPU.Curr.Opcode = func() int { return GB.RLCR8("B") }

	case 0x01:
		GB.CPU.Curr.Opcode = func() int { return GB.RLCR8("C") }

	case 0x02:
		GB.CPU.Curr.Opcode = func() int { return GB.RLCR8("D") }

	case 0x03:
		GB.CPU.Curr.Opcode = func() int { return GB.RLCR8("E") }

	case 0x04:
		GB.CPU.Curr.Opcode = func() int { return GB.RLCR8("H") }

	case 0x05:
		GB.CPU.Curr.Opcode = func() int { return GB.RLCR8("L") }

	case 0x06:
		GB.CPU.Curr.Opcode = func() int { return GB.RLCHL() }

	case 0x07:
		GB.CPU.Curr.Opcode = func() int { return GB.RLCR8("A") }

	case 0x08:
		GB.CPU.Curr.Opcode = func() int { return GB.RRCR8("B") }

	case 0x09:
		GB.CPU.Curr.Opcode = func() int { return GB.RRCR8("C") }

	case 0x0A:
		GB.CPU.Curr.Opcode = func() int { return GB.RRCR8("D") }

	case 0x0B:
		GB.CPU.Curr.Opcode = func() int { return GB.RRCR8("E") }

	case 0x0C:
		GB.CPU.Curr.Opcode = func() int { return GB.RRCR8("H") }

	case 0x0D:
		GB.CPU.Curr.Opcode = func() int { return GB.RRCR8("L") }

	case 0x0E:
		GB.CPU.Curr.Opcode = func() int { return GB.RRCHL() }

	case 0x0F:
		GB.CPU.Curr.Opcode = func() int { return GB.RRCR8("A") }

	case 0x10:
		GB.CPU.Curr.Opcode = func() int { return GB.RLR8("B") }

	case 0x11:
		GB.CPU.Curr.Opcode = func() int { return GB.RLR8("C") }

	case 0x12:
		GB.CPU.Curr.Opcode = func() int { return GB.RLR8("D") }

	case 0x13:
		GB.CPU.Curr.Opcode = func() int { return GB.RLR8("E") }

	case 0x14:
		GB.CPU.Curr.Opcode = func() int { return GB.RLR8("H") }

	case 0x15:
		GB.CPU.Curr.Opcode = func() int { return GB.RLR8("L") }

	case 0x16:
		GB.CPU.Curr.Opcode = func() int { return GB.RLHL() }

	case 0x17:
		GB.CPU.Curr.Opcode = func() int { return GB.RLR8("A") }

	case 0x18:
		GB.CPU.Curr.Opcode = func() int { return GB.RRR8("B") }

	case 0x19:
		GB.CPU.Curr.Opcode = func() int { return GB.RRR8("C") }

	case 0x1A:
		GB.CPU.Curr.Opcode = func() int { return GB.RRR8("D") }

	case 0x1B:
		GB.CPU.Curr.Opcode = func() int { return GB.RRR8("E") }

	case 0x1C:
		GB.CPU.Curr.Opcode = func() int { return GB.RRR8("H") }

	case 0x1D:
		GB.CPU.Curr.Opcode = func() int { return GB.RRR8("L") }

	case 0x1E:
		GB.CPU.Curr.Opcode = func() int { return GB.RRHL() }

	case 0x1F:
		GB.CPU.Curr.Opcode = func() int { return GB.RRR8("A") }

	case 0x20:
		GB.CPU.Curr.Opcode = func() int { return GB.SLAR8("B") }

	case 0x21:
		GB.CPU.Curr.Opcode = func() int { return GB.SLAR8("C") }

	case 0x22:
		GB.CPU.Curr.Opcode = func() int { return GB.SLAR8("D") }

	case 0x23:
		GB.CPU.Curr.Opcode = func() int { return GB.SLAR8("E") }

	case 0x24:
		GB.CPU.Curr.Opcode = func() int { return GB.SLAR8("H") }

	case 0x25:
		GB.CPU.Curr.Opcode = func() int { return GB.SLAR8("L") }

	case 0x26:
		GB.CPU.Curr.Opcode = func() int { return GB.SLAHL() }

	case 0x27:
		GB.CPU.Curr.Opcode = func() int { return GB.SLAR8("A") }

	case 0x28:
		GB.CPU.Curr.Opcode = func() int { return GB.SRAR8("B") }

	case 0x29:
		GB.CPU.Curr.Opcode = func() int { return GB.SRAR8("C") }

	case 0x2A:
		GB.CPU.Curr.Opcode = func() int { return GB.SRAR8("D") }

	case 0x2B:
		GB.CPU.Curr.Opcode = func() int { return GB.SRAR8("E") }

	case 0x2C:
		GB.CPU.Curr.Opcode = func() int { return GB.SRAR8("H") }

	case 0x2D:
		GB.CPU.Curr.Opcode = func() int { return GB.SRAR8("L") }

	case 0x2E:
		GB.CPU.Curr.Opcode = func() int { return GB.SRAHL() }

	case 0x2F:
		GB.CPU.Curr.Opcode = func() int { return GB.SRAR8("A") }

	case 0x30:
		GB.CPU.Curr.Opcode = func() int { return GB.SWAPR8("B") }

	case 0x31:
		GB.CPU.Curr.Opcode = func() int { return GB.SWAPR8("C") }

	case 0x32:
		GB.CPU.Curr.Opcode = func() int { return GB.SWAPR8("D") }

	case 0x33:
		GB.CPU.Curr.Opcode = func() int { return GB.SWAPR8("E") }

	case 0x34:
		GB.CPU.Curr.Opcode = func() int { return GB.SWAPR8("H") }

	case 0x35:
		GB.CPU.Curr.Opcode = func() int { return GB.SWAPR8("L") }

	case 0x36:
		GB.CPU.Curr.Opcode = func() int { return GB.SWAPHL() }

	case 0x37:
		GB.CPU.Curr.Opcode = func() int { return GB.SWAPR8("A") }

	case 0x38:
		GB.CPU.Curr.Opcode = func() int { return GB.SRLR8("B") }

	case 0x39:
		GB.CPU.Curr.Opcode = func() int { return GB.SRLR8("C") }

	case 0x3A:
		GB.CPU.Curr.Opcode = func() int { return GB.SRLR8("D") }

	case 0x3B:
		GB.CPU.Curr.Opcode = func() int { return GB.SRLR8("E") }

	case 0x3C:
		GB.CPU.Curr.Opcode = func() int { return GB.SRLR8("H") }

	case 0x3D:
		GB.CPU.Curr.Opcode = func() int { return GB.SRLR8("L") }

	case 0x3E:
		GB.CPU.Curr.Opcode = func() int { return GB.SRLHL() }

	case 0x3F:
		GB.CPU.Curr.Opcode = func() int { return GB.SRLR8("A") }

	case 0x40:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(0, "B") }

	case 0x41:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(0, "C") }

	case 0x42:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(0, "D") }

	case 0x43:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(0, "E") }

	case 0x44:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(0, "H") }

	case 0x45:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(0, "L") }

	case 0x46:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3HL(0) }

	case 0x47:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(0, "A") }

	case 0x48:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(1, "B") }

	case 0x49:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(1, "C") }

	case 0x4A:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(1, "D") }

	case 0x4B:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(1, "E") }

	case 0x4C:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(1, "H") }

	case 0x4D:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(1, "L") }

	case 0x4E:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3HL(1) }

	case 0x4F:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(1, "A") }

	case 0x50:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(2, "B") }

	case 0x51:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(2, "C") }

	case 0x52:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(2, "D") }

	case 0x53:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(2, "E") }

	case 0x54:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(2, "H") }

	case 0x55:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(2, "L") }

	case 0x56:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3HL(2) }

	case 0x57:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(2, "A") }

	case 0x58:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(3, "B") }

	case 0x59:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(3, "C") }

	case 0x5A:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(3, "D") }

	case 0x5B:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(3, "E") }

	case 0x5C:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(3, "H") }

	case 0x5D:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(3, "L") }

	case 0x5E:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3HL(3) }

	case 0x5F:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(3, "A") }

	case 0x60:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(4, "B") }

	case 0x61:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(4, "C") }

	case 0x62:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(4, "D") }

	case 0x63:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(4, "E") }

	case 0x64:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(4, "H") }

	case 0x65:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(4, "L") }

	case 0x66:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3HL(4) }

	case 0x67:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(4, "A") }

	case 0x68:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(5, "B") }

	case 0x69:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(5, "C") }

	case 0x6A:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(5, "D") }

	case 0x6B:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(5, "E") }

	case 0x6C:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(5, "H") }

	case 0x6D:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(5, "L") }

	case 0x6E:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3HL(5) }

	case 0x6F:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(5, "A") }

	case 0x70:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(6, "B") }

	case 0x71:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(6, "C") }

	case 0x72:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(6, "D") }

	case 0x73:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(6, "E") }

	case 0x74:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(6, "H") }

	case 0x75:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(6, "L") }

	case 0x76:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3HL(6) }

	case 0x77:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(6, "A") }

	case 0x78:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(7, "B") }

	case 0x79:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(7, "C") }

	case 0x7A:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(7, "D") }

	case 0x7B:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(7, "E") }

	case 0x7C:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(7, "H") }

	case 0x7D:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(7, "L") }

	case 0x7E:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3HL(7) }

	case 0x7F:
		GB.CPU.Curr.Opcode = func() int { return GB.BITU3R8(7, "A") }

	case 0x80:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(0, "B") }

	case 0x81:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(0, "C") }

	case 0x82:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(0, "D") }

	case 0x83:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(0, "E") }

	case 0x84:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(0, "H") }

	case 0x85:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(0, "L") }

	case 0x86:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3HL(0) }

	case 0x87:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(0, "A") }

	case 0x88:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(1, "B") }

	case 0x89:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(1, "C") }

	case 0x8A:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(1, "D") }

	case 0x8B:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(1, "E") }

	case 0x8C:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(1, "H") }

	case 0x8D:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(1, "L") }

	case 0x8E:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3HL(1) }

	case 0x8F:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(1, "A") }

	case 0x90:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(2, "B") }

	case 0x91:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(2, "C") }

	case 0x92:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(2, "D") }

	case 0x93:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(2, "E") }

	case 0x94:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(2, "H") }

	case 0x95:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(2, "L") }

	case 0x96:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3HL(2) }

	case 0x97:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(2, "A") }

	case 0x98:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(3, "B") }

	case 0x99:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(3, "C") }

	case 0x9A:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(3, "D") }

	case 0x9B:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(3, "E") }

	case 0x9C:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(3, "H") }

	case 0x9D:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(3, "L") }

	case 0x9E:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3HL(3) }

	case 0x9F:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(3, "A") }

	case 0xA0:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(4, "B") }

	case 0xA1:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(4, "C") }

	case 0xA2:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(4, "D") }

	case 0xA3:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(4, "E") }

	case 0xA4:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(4, "H") }

	case 0xA5:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(4, "L") }

	case 0xA6:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3HL(4) }

	case 0xA7:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(4, "A") }

	case 0xA8:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(5, "B") }

	case 0xA9:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(5, "C") }

	case 0xAA:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(5, "D") }

	case 0xAB:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(5, "E") }

	case 0xAC:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(5, "H") }

	case 0xAD:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(5, "L") }

	case 0xAE:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3HL(5) }

	case 0xAF:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(5, "A") }

	case 0xB0:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(6, "B") }

	case 0xB1:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(6, "C") }

	case 0xB2:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(6, "D") }

	case 0xB3:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(6, "E") }

	case 0xB4:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(6, "H") }

	case 0xB5:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(6, "L") }

	case 0xB6:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3HL(6) }

	case 0xB7:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(6, "A") }

	case 0xB8:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(7, "B") }

	case 0xB9:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(7, "C") }

	case 0xBA:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(7, "D") }

	case 0xBB:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(7, "E") }

	case 0xBC:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(7, "H") }

	case 0xBD:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(7, "L") }

	case 0xBE:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3HL(7) }

	case 0xBF:
		GB.CPU.Curr.Opcode = func() int { return GB.RESU3R8(7, "A") }

	case 0xC0:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(0, "B") }

	case 0xC1:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(0, "C") }

	case 0xC2:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(0, "D") }

	case 0xC3:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(0, "E") }

	case 0xC4:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(0, "H") }

	case 0xC5:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(0, "L") }

	case 0xC6:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3HL(0) }

	case 0xC7:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(0, "A") }

	case 0xC8:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(1, "B") }

	case 0xC9:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(1, "C") }

	case 0xCA:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(1, "D") }

	case 0xCB:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(1, "E") }

	case 0xCC:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(1, "H") }

	case 0xCD:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(1, "L") }

	case 0xCE:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3HL(1) }

	case 0xCF:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(1, "A") }

	case 0xD0:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(2, "B") }

	case 0xD1:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(2, "C") }

	case 0xD2:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(2, "D") }

	case 0xD3:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(2, "E") }

	case 0xD4:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(2, "H") }

	case 0xD5:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(2, "L") }

	case 0xD6:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3HL(2) }

	case 0xD7:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(2, "A") }

	case 0xD8:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(3, "B") }

	case 0xD9:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(3, "C") }

	case 0xDA:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(3, "D") }

	case 0xDB:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(3, "E") }

	case 0xDC:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(3, "H") }

	case 0xDD:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(3, "L") }

	case 0xDE:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3HL(3) }

	case 0xDF:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(3, "A") }

	case 0xE0:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(4, "B") }

	case 0xE1:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(4, "C") }

	case 0xE2:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(4, "D") }

	case 0xE3:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(4, "E") }

	case 0xE4:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(4, "H") }

	case 0xE5:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(4, "L") }

	case 0xE6:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3HL(4) }

	case 0xE7:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(4, "A") }

	case 0xE8:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(5, "B") }

	case 0xE9:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(5, "C") }

	case 0xEA:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(5, "D") }

	case 0xEB:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(5, "E") }

	case 0xEC:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(5, "H") }

	case 0xED:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(5, "L") }

	case 0xEE:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3HL(5) }

	case 0xEF:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(5, "A") }

	case 0xF0:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(6, "B") }

	case 0xF1:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(6, "C") }

	case 0xF2:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(6, "D") }

	case 0xF3:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(6, "E") }

	case 0xF4:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(6, "H") }

	case 0xF5:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(6, "L") }

	case 0xF6:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3HL(6) }

	case 0xF7:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(6, "A") }

	case 0xF8:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(7, "B") }

	case 0xF9:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(7, "C") }

	case 0xFA:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(7, "D") }

	case 0xFB:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(7, "E") }

	case 0xFC:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(7, "H") }

	case 0xFD:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(7, "L") }

	case 0xFE:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3HL(7) }

	case 0xFF:
		GB.CPU.Curr.Opcode = func() int { return GB.SETU3R8(7, "A") }

	}
}
