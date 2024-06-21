package GB

func (GB *GB) Execute() {
	cycles := GB.CPU.Curr.Opcode()
	GB.IncrementTimer(cycles)
}

func (GB *GB) Fetch() {
	GB.CPU.IncrementCounter()
	GB.IncrementTimer(4)

	switch GB.MMU.Read(GB.CPU.Reg.Read("PC")) {

	case 0x00:
		GB.CPU.Curr.Opcode = func() int { return GB.NOP() }

	case 0x01:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR16N16("BC") }

	case 0x02:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR16A("BC") }

	case 0x03:
		GB.CPU.Curr.Opcode = func() int { return GB.INCR16("BC") }

	case 0x04:
		GB.CPU.Curr.Opcode = func() int { return GB.INCR8("B") }

	case 0x05:
		GB.CPU.Curr.Opcode = func() int { return GB.DECR8("B") }

	case 0x06:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8N8("B") }

	case 0x07:
		GB.CPU.Curr.Opcode = func() int { return GB.RLCA() }

	case 0x08:
		GB.CPU.Curr.Opcode = func() int { return GB.LDN16SP() }

	case 0x09:
		GB.CPU.Curr.Opcode = func() int { return GB.ADDHLR16("BC") }

	case 0x0A:
		GB.CPU.Curr.Opcode = func() int { return GB.LDAR16("BC") }

	case 0x0B:
		GB.CPU.Curr.Opcode = func() int { return GB.DECR16("BC") }

	case 0x0C:
		GB.CPU.Curr.Opcode = func() int { return GB.INCR8("C") }

	case 0x0D:
		GB.CPU.Curr.Opcode = func() int { return GB.DECR8("C") }

	case 0x0E:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8N8("C") }

	case 0x0F:
		GB.CPU.Curr.Opcode = func() int { return GB.RRCA() }

	case 0x10:
		// opcode := fmt.Sprintf("STOP $%x", int8(GB.MMU.Read(GB.CPU.Reg.Read("PC")+1)))

	case 0x11:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR16N16("DE") }

	case 0x12:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR16A("DE") }

	case 0x13:
		GB.CPU.Curr.Opcode = func() int { return GB.INCR16("DE") }

	case 0x14:
		GB.CPU.Curr.Opcode = func() int { return GB.INCR8("D") }

	case 0x15:
		GB.CPU.Curr.Opcode = func() int { return GB.DECR8("D") }

	case 0x16:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8N8("D") }

	case 0x17:
		GB.CPU.Curr.Opcode = func() int { return GB.RLA() }

	case 0x18:
		GB.CPU.Curr.Opcode = func() int { return GB.JRN16() }

	case 0x19:
		GB.CPU.Curr.Opcode = func() int { return GB.ADDHLR16("DE") }

	case 0x1A:
		GB.CPU.Curr.Opcode = func() int { return GB.LDAR16("DE") }

	case 0x1B:
		GB.CPU.Curr.Opcode = func() int { return GB.DECR16("DE") }

	case 0x1C:
		GB.CPU.Curr.Opcode = func() int { return GB.INCR8("E") }

	case 0x1D:
		GB.CPU.Curr.Opcode = func() int { return GB.DECR8("E") }

	case 0x1E:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8N8("E") }

	case 0x1F:
		GB.CPU.Curr.Opcode = func() int { return GB.RRA() }

	case 0x20:
		GB.CPU.Curr.Opcode = func() int { return GB.JRCCN16("NZ") }

	case 0x21:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR16N16("HL") }

	case 0x22:
		GB.CPU.Curr.Opcode = func() int { return GB.LDHLIA() }

	case 0x23:
		GB.CPU.Curr.Opcode = func() int { return GB.INCR16("HL") }

	case 0x24:
		GB.CPU.Curr.Opcode = func() int { return GB.INCR8("H") }

	case 0x25:
		GB.CPU.Curr.Opcode = func() int { return GB.DECR8("H") }

	case 0x26:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8N8("H") }

	case 0x27:
		GB.CPU.Curr.Opcode = func() int { return GB.DAA() }

	case 0x28:
		GB.CPU.Curr.Opcode = func() int { return GB.JRCCN16("Z") }

	case 0x29:
		GB.CPU.Curr.Opcode = func() int { return GB.ADDHLR16("HL") }

	case 0x2A:
		GB.CPU.Curr.Opcode = func() int { return GB.LDAHLI() }

	case 0x2B:
		GB.CPU.Curr.Opcode = func() int { return GB.DECR16("HL") }

	case 0x2C:
		GB.CPU.Curr.Opcode = func() int { return GB.INCR8("L") }

	case 0x2D:
		GB.CPU.Curr.Opcode = func() int { return GB.DECR8("L") }

	case 0x2E:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8N8("L") }

	case 0x2F:
		GB.CPU.Curr.Opcode = func() int { return GB.CPL() }

	case 0x30:
		GB.CPU.Curr.Opcode = func() int { return GB.JRCCN16("NC") }

	case 0x31:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR16N16("SP") }

	case 0x32:
		GB.CPU.Curr.Opcode = func() int { return GB.LDHLDA() }

	case 0x33:
		GB.CPU.Curr.Opcode = func() int { return GB.INCR16("SP") }

	case 0x34:
		GB.CPU.Curr.Opcode = func() int { return GB.INCRHL() }

	case 0x35:
		GB.CPU.Curr.Opcode = func() int { return GB.DECHL() }

	case 0x36:
		GB.CPU.Curr.Opcode = func() int { return GB.LDHLN8() }

	case 0x37:
		GB.CPU.Curr.Opcode = func() int { return GB.SCF() }

	case 0x38:
		GB.CPU.Curr.Opcode = func() int { return GB.JRCCN16("C") }

	case 0x39:
		GB.CPU.Curr.Opcode = func() int { return GB.ADDHLSP() }

	case 0x3A:
		GB.CPU.Curr.Opcode = func() int { return GB.LDAHLD() }

	case 0x3B:
		GB.CPU.Curr.Opcode = func() int { return GB.DECR16("SP") }

	case 0x3C:
		GB.CPU.Curr.Opcode = func() int { return GB.INCR8("A") }

	case 0x3D:
		GB.CPU.Curr.Opcode = func() int { return GB.DECR8("A") }

	case 0x3E:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8N8("A") }

	case 0x3F:
		GB.CPU.Curr.Opcode = func() int { return GB.CCF() }

	case 0x40:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("B", "B") }

	case 0x41:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("B", "C") }

	case 0x42:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("B", "D") }

	case 0x43:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("B", "E") }

	case 0x44:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("B", "H") }

	case 0x45:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("B", "L") }

	case 0x46:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8HL("B") }

	case 0x47:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("B", "A") }

	case 0x48:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("C", "B") }

	case 0x49:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("C", "C") }

	case 0x4A:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("C", "D") }

	case 0x4B:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("C", "E") }

	case 0x4C:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("C", "H") }

	case 0x4D:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("C", "L") }

	case 0x4E:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8HL("C") }

	case 0x4F:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("C", "A") }

	case 0x50:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("D", "B") }

	case 0x51:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("D", "C") }

	case 0x52:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("D", "D") }

	case 0x53:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("D", "E") }

	case 0x54:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("D", "H") }

	case 0x55:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("D", "L") }

	case 0x56:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8HL("D") }

	case 0x57:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("D", "A") }

	case 0x58:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("E", "B") }

	case 0x59:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("E", "C") }

	case 0x5A:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("E", "D") }

	case 0x5B:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("E", "E") }

	case 0x5C:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("E", "H") }

	case 0x5D:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("E", "L") }

	case 0x5E:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8HL("E") }

	case 0x5F:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("E", "A") }

	case 0x60:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("H", "B") }

	case 0x61:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("H", "C") }

	case 0x62:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("H", "D") }

	case 0x63:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("H", "E") }

	case 0x64:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("H", "H") }

	case 0x65:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("H", "L") }

	case 0x66:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8HL("H") }

	case 0x67:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("H", "A") }

	case 0x68:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("L", "B") }

	case 0x69:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("L", "C") }

	case 0x6A:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("L", "D") }

	case 0x6B:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("L", "E") }

	case 0x6C:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("L", "H") }

	case 0x6D:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("L", "L") }

	case 0x6E:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8HL("L") }

	case 0x6F:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("L", "A") }

	case 0x70:
		GB.CPU.Curr.Opcode = func() int { return GB.LDHLR8("B") }

	case 0x71:
		GB.CPU.Curr.Opcode = func() int { return GB.LDHLR8("C") }

	case 0x72:
		GB.CPU.Curr.Opcode = func() int { return GB.LDHLR8("D") }

	case 0x73:
		GB.CPU.Curr.Opcode = func() int { return GB.LDHLR8("E") }

	case 0x74:
		GB.CPU.Curr.Opcode = func() int { return GB.LDHLR8("H") }

	case 0x75:
		GB.CPU.Curr.Opcode = func() int { return GB.LDHLR8("L") }

	case 0x76:
		GB.CPU.Curr.Opcode = func() int { return GB.HALT() }

	case 0x77:
		GB.CPU.Curr.Opcode = func() int { return GB.LDHLR8("A") }

	case 0x78:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("A", "B") }

	case 0x79:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("A", "C") }

	case 0x7A:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("A", "D") }

	case 0x7B:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("A", "E") }

	case 0x7C:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("A", "H") }

	case 0x7D:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("A", "L") }

	case 0x7E:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8HL("A") }

	case 0x7F:
		GB.CPU.Curr.Opcode = func() int { return GB.LDR8R8("A", "A") }

	case 0x80:
		GB.CPU.Curr.Opcode = func() int { return GB.ADDAR8("B") }

	case 0x81:
		GB.CPU.Curr.Opcode = func() int { return GB.ADDAR8("C") }

	case 0x82:
		GB.CPU.Curr.Opcode = func() int { return GB.ADDAR8("D") }

	case 0x83:
		GB.CPU.Curr.Opcode = func() int { return GB.ADDAR8("E") }

	case 0x84:
		GB.CPU.Curr.Opcode = func() int { return GB.ADDAR8("H") }

	case 0x85:
		GB.CPU.Curr.Opcode = func() int { return GB.ADDAR8("L") }

	case 0x86:
		GB.CPU.Curr.Opcode = func() int { return GB.ADDAHL() }

	case 0x87:
		GB.CPU.Curr.Opcode = func() int { return GB.ADDAR8("A") }

	case 0x88:
		GB.CPU.Curr.Opcode = func() int { return GB.ADCAR8("B") }

	case 0x89:
		GB.CPU.Curr.Opcode = func() int { return GB.ADCAR8("C") }

	case 0x8A:
		GB.CPU.Curr.Opcode = func() int { return GB.ADCAR8("D") }

	case 0x8B:
		GB.CPU.Curr.Opcode = func() int { return GB.ADCAR8("E") }

	case 0x8C:
		GB.CPU.Curr.Opcode = func() int { return GB.ADCAR8("H") }

	case 0x8D:
		GB.CPU.Curr.Opcode = func() int { return GB.ADCAR8("L") }

	case 0x8E:
		GB.CPU.Curr.Opcode = func() int { return GB.ADCAHL() }

	case 0x8F:
		GB.CPU.Curr.Opcode = func() int { return GB.ADCAR8("A") }

	case 0x90:
		GB.CPU.Curr.Opcode = func() int { return GB.SUBAR8("B") }

	case 0x91:
		GB.CPU.Curr.Opcode = func() int { return GB.SUBAR8("C") }

	case 0x92:
		GB.CPU.Curr.Opcode = func() int { return GB.SUBAR8("D") }

	case 0x93:
		GB.CPU.Curr.Opcode = func() int { return GB.SUBAR8("E") }

	case 0x94:
		GB.CPU.Curr.Opcode = func() int { return GB.SUBAR8("H") }

	case 0x95:
		GB.CPU.Curr.Opcode = func() int { return GB.SUBAR8("L") }

	case 0x96:
		GB.CPU.Curr.Opcode = func() int { return GB.SUBAHL() }

	case 0x97:
		GB.CPU.Curr.Opcode = func() int { return GB.SUBAR8("A") }

	case 0x98:
		GB.CPU.Curr.Opcode = func() int { return GB.SBCAR8("B") }

	case 0x99:
		GB.CPU.Curr.Opcode = func() int { return GB.SBCAR8("C") }

	case 0x9A:
		GB.CPU.Curr.Opcode = func() int { return GB.SBCAR8("D") }

	case 0x9B:
		GB.CPU.Curr.Opcode = func() int { return GB.SBCAR8("E") }

	case 0x9C:
		GB.CPU.Curr.Opcode = func() int { return GB.SBCAR8("H") }

	case 0x9D:
		GB.CPU.Curr.Opcode = func() int { return GB.SBCAR8("L") }

	case 0x9E:
		GB.CPU.Curr.Opcode = func() int { return GB.SBCAHL() }

	case 0x9F:
		GB.CPU.Curr.Opcode = func() int { return GB.SBCAR8("A") }

	case 0xA0:
		GB.CPU.Curr.Opcode = func() int { return GB.ANDAR8("B") }

	case 0xA1:
		GB.CPU.Curr.Opcode = func() int { return GB.ANDAR8("C") }

	case 0xA2:
		GB.CPU.Curr.Opcode = func() int { return GB.ANDAR8("D") }

	case 0xA3:
		GB.CPU.Curr.Opcode = func() int { return GB.ANDAR8("E") }

	case 0xA4:
		GB.CPU.Curr.Opcode = func() int { return GB.ANDAR8("H") }

	case 0xA5:
		GB.CPU.Curr.Opcode = func() int { return GB.ANDAR8("L") }

	case 0xA6:
		GB.CPU.Curr.Opcode = func() int { return GB.ANDAHL() }

	case 0xA7:
		GB.CPU.Curr.Opcode = func() int { return GB.ANDAR8("A") }

	case 0xA8:
		GB.CPU.Curr.Opcode = func() int { return GB.XORAR8("B") }

	case 0xA9:
		GB.CPU.Curr.Opcode = func() int { return GB.XORAR8("C") }

	case 0xAA:
		GB.CPU.Curr.Opcode = func() int { return GB.XORAR8("D") }

	case 0xAB:
		GB.CPU.Curr.Opcode = func() int { return GB.XORAR8("E") }

	case 0xAC:
		GB.CPU.Curr.Opcode = func() int { return GB.XORAR8("H") }

	case 0xAD:
		GB.CPU.Curr.Opcode = func() int { return GB.XORAR8("L") }

	case 0xAE:
		GB.CPU.Curr.Opcode = func() int { return GB.XORAHL() }

	case 0xAF:
		GB.CPU.Curr.Opcode = func() int { return GB.XORAR8("A") }

	case 0xB0:
		GB.CPU.Curr.Opcode = func() int { return GB.ORAR8("B") }

	case 0xB1:
		GB.CPU.Curr.Opcode = func() int { return GB.ORAR8("C") }

	case 0xB2:
		GB.CPU.Curr.Opcode = func() int { return GB.ORAR8("D") }

	case 0xB3:
		GB.CPU.Curr.Opcode = func() int { return GB.ORAR8("E") }

	case 0xB4:
		GB.CPU.Curr.Opcode = func() int { return GB.ORAR8("H") }

	case 0xB5:
		GB.CPU.Curr.Opcode = func() int { return GB.ORAR8("L") }

	case 0xB6:
		GB.CPU.Curr.Opcode = func() int { return GB.ORAHL() }

	case 0xB7:
		GB.CPU.Curr.Opcode = func() int { return GB.ORAR8("A") }

	case 0xB8:
		GB.CPU.Curr.Opcode = func() int { return GB.CPAR8("B") }

	case 0xB9:
		GB.CPU.Curr.Opcode = func() int { return GB.CPAR8("C") }

	case 0xBA:
		GB.CPU.Curr.Opcode = func() int { return GB.CPAR8("D") }

	case 0xBB:
		GB.CPU.Curr.Opcode = func() int { return GB.CPAR8("E") }

	case 0xBC:
		GB.CPU.Curr.Opcode = func() int { return GB.CPAR8("H") }

	case 0xBD:
		GB.CPU.Curr.Opcode = func() int { return GB.CPAR8("L") }

	case 0xBE:
		GB.CPU.Curr.Opcode = func() int { return GB.CPAHL() }

	case 0xBF:
		GB.CPU.Curr.Opcode = func() int { return GB.CPAR8("A") }

	case 0xC0:
		GB.CPU.Curr.Opcode = func() int { return GB.RETCC("NZ") }

	case 0xC1:
		GB.CPU.Curr.Opcode = func() int { return GB.POPR16("BC") }

	case 0xC2:
		GB.CPU.Curr.Opcode = func() int { return GB.JPCCN16("NZ") }

	case 0xC3:
		addr := GB.MMU.Read(GB.CPU.Reg.Read("PC")+1) + GB.MMU.Read(GB.CPU.Reg.Read("PC")+2)<<8
		GB.CPU.Curr.Opcode = func() int { return GB.JPN16(addr) }

	case 0xC4:
		GB.CPU.Curr.Opcode = func() int { return GB.CALLCCN16("NZ") }

	case 0xC5:
		GB.CPU.Curr.Opcode = func() int { return GB.PUSHR16("BC") }

	case 0xC6:
		GB.CPU.Curr.Opcode = func() int { return GB.ADDAN8() }

	case 0xC7:
		GB.CPU.Curr.Opcode = func() int { return GB.RST("00H") }

	case 0xC8:
		GB.CPU.Curr.Opcode = func() int { return GB.RETCC("Z") }

	case 0xC9:
		GB.CPU.Curr.Opcode = func() int { return GB.RET() }

	case 0xCA:
		GB.CPU.Curr.Opcode = func() int { return GB.JPCCN16("Z") }

	case 0xCB:
		GB.PrefixTable()

	case 0xCC:
		GB.CPU.Curr.Opcode = func() int { return GB.CALLCCN16("Z") }

	case 0xCD:
		GB.CPU.Curr.Opcode = func() int { return GB.CALLN16() }

	case 0xCE:
		GB.CPU.Curr.Opcode = func() int { return GB.ADCAN8() }

	case 0xCF:
		GB.CPU.Curr.Opcode = func() int { return GB.RST("08H") }

	case 0xD0:
		GB.CPU.Curr.Opcode = func() int { return GB.RETCC("NC") }

	case 0xD1:
		GB.CPU.Curr.Opcode = func() int { return GB.POPR16("DE") }

	case 0xD2:
		GB.CPU.Curr.Opcode = func() int { return GB.JPCCN16("NC") }

	case 0xD4:
		GB.CPU.Curr.Opcode = func() int { return GB.CALLCCN16("NC") }

	case 0xD5:
		GB.CPU.Curr.Opcode = func() int { return GB.PUSHR16("DE") }

	case 0xD6:
		GB.CPU.Curr.Opcode = func() int { return GB.SUBAN8() }

	case 0xD7:
		GB.CPU.Curr.Opcode = func() int { return GB.RST("10H") }

	case 0xD8:
		GB.CPU.Curr.Opcode = func() int { return GB.RETCC("C") }

	case 0xD9:
		GB.CPU.Curr.Opcode = func() int { return GB.RETI() }

	case 0xDA:
		GB.CPU.Curr.Opcode = func() int { return GB.JPCCN16("C") }

	case 0xDC:
		GB.CPU.Curr.Opcode = func() int { return GB.CALLCCN16("C") }

	case 0xDE:
		GB.CPU.Curr.Opcode = func() int { return GB.SBCAN8() }

	case 0xDF:
		GB.CPU.Curr.Opcode = func() int { return GB.RST("18H") }

	case 0xE0:
		GB.CPU.Curr.Opcode = func() int { return GB.LDHU8A() }

	case 0xE1:
		GB.CPU.Curr.Opcode = func() int { return GB.POPR16("HL") }

	case 0xE2:
		GB.CPU.Curr.Opcode = func() int { return GB.LDHCA() }

	case 0xE5:
		GB.CPU.Curr.Opcode = func() int { return GB.PUSHR16("HL") }

	case 0xE6:
		GB.CPU.Curr.Opcode = func() int { return GB.ANDAN8() }

	case 0xE7:
		GB.CPU.Curr.Opcode = func() int { return GB.RST("20H") }

	case 0xE8:
		GB.CPU.Curr.Opcode = func() int { return GB.ADDSPE8() }

	case 0xE9:
		GB.CPU.Curr.Opcode = func() int { return GB.JPHL() }

	case 0xEA:
		GB.CPU.Curr.Opcode = func() int { return GB.LDN16A() }

	case 0xEE:
		GB.CPU.Curr.Opcode = func() int { return GB.XORAN8() }

	case 0xEF:
		GB.CPU.Curr.Opcode = func() int { return GB.RST("28H") }

	case 0xF0:
		GB.CPU.Curr.Opcode = func() int { return GB.LDHAU8() }

	case 0xF1:
		GB.CPU.Curr.Opcode = func() int { return GB.POPFA() }

	case 0xF2:
		GB.CPU.Curr.Opcode = func() int { return GB.LDHAC() }

	case 0xF3:
		GB.CPU.Curr.Opcode = func() int { return GB.DI() }

	case 0xF4:

	case 0xF5:
		GB.CPU.Curr.Opcode = func() int { return GB.PUSHFA() }

	case 0xF6:
		GB.CPU.Curr.Opcode = func() int { return GB.ORAN8() }

	case 0xF7:
		GB.CPU.Curr.Opcode = func() int { return GB.RST("30H") }

	case 0xF8:
		GB.CPU.Curr.Opcode = func() int { return GB.LDHLSPE8() }

	case 0xF9:
		GB.CPU.Curr.Opcode = func() int { return GB.LDSPHL() }

	case 0xFA:
		GB.CPU.Curr.Opcode = func() int { return GB.LDAN16() }

	case 0xFB:
		GB.CPU.Curr.Opcode = func() int { return GB.EI() }

	case 0xFE:
		GB.CPU.Curr.Opcode = func() int { return GB.CPAN8() }

	case 0xFF:
		GB.CPU.Curr.Opcode = func() int { return GB.RST("38H") }

	}
}
