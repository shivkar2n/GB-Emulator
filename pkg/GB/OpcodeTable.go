package GB

import "fmt"

func (GB *GB) ExecuteOpcode() (string, int, int, bool) {

	switch GB.MMU.Read(GB.CPU.Reg.Read("PC")) {

	case 0x00:
		return GB.NOP()

	case 0x01:
		return GB.LDR16N16("BC")

	case 0x02:
		return GB.LDR16A("BC")

	case 0x03:
		return GB.INCR16("BC")

	case 0x04:
		return GB.INCR8("B")

	case 0x05:
		return GB.DECR8("B")

	case 0x06:
		return GB.LDR8N8("B")

	case 0x07:
		return GB.RLCA()

	case 0x08:
		return GB.LDN16SP()

	case 0x09:
		return GB.ADDHLR16("BC")

	case 0x0A:
		return GB.LDAR16("BC")

	case 0x0B:
		return GB.DECR16("BC")

	case 0x0C:
		return GB.INCR8("C")

	case 0x0D:
		return GB.DECR8("C")

	case 0x0E:
		return GB.LDR8N8("C")

	case 0x0F:
		return GB.RRCA()

	case 0x10:
		opcode := fmt.Sprintf("STOP $%x", int8(GB.MMU.Read(GB.CPU.Reg.Read("PC")+1)))
		return opcode, 0, 0, false

	case 0x11:
		return GB.LDR16N16("DE")

	case 0x12:
		return GB.LDR16A("DE")

	case 0x13:
		return GB.INCR16("DE")

	case 0x14:
		return GB.INCR8("D")

	case 0x15:
		return GB.DECR8("D")

	case 0x16:
		return GB.LDR8N8("D")

	case 0x17:
		return GB.RLA()

	case 0x18:
		return GB.JRN16()

	case 0x19:
		return GB.ADDHLR16("DE")

	case 0x1A:
		return GB.LDAR16("DE")

	case 0x1B:
		return GB.DECR16("DE")

	case 0x1C:
		return GB.INCR8("E")

	case 0x1D:
		return GB.DECR8("E")

	case 0x1E:
		return GB.LDR8N8("E")

	case 0x1F:
		return GB.RRA()

	case 0x20:
		return GB.JRCCN16("NZ")

	case 0x21:
		return GB.LDR16N16("HL")

	case 0x22:
		return GB.LDHLIA()

	case 0x23:
		return GB.INCR16("HL")

	case 0x24:
		return GB.INCR8("H")

	case 0x25:
		return GB.DECR8("H")

	case 0x26:
		return GB.LDR8N8("H")

	case 0x27:
		return GB.DAA()

	case 0x28:
		return GB.JRCCN16("Z")

	case 0x29:
		return GB.ADDHLR16("HL")

	case 0x2A:
		return GB.LDAHLI()

	case 0x2B:
		return GB.DECR16("HL")

	case 0x2C:
		return GB.INCR8("L")

	case 0x2D:
		return GB.DECR8("L")

	case 0x2E:
		return GB.LDR8N8("L")

	case 0x2F:
		return GB.CPL()

	case 0x30:
		return GB.JRCCN16("NC")

	case 0x31:
		return GB.LDR16N16("SP")

	case 0x32:
		return GB.LDHLDA()

	case 0x33:
		return GB.INCR16("SP")

	case 0x34:
		return GB.INCRHL()

	case 0x35:
		return GB.DECHL()

	case 0x36:
		return GB.LDHLN8()

	case 0x37:
		return GB.SCF()

	case 0x38:
		return GB.JRCCN16("C")

	case 0x39:
		return GB.ADDHLSP()

	case 0x3A:
		return GB.LDAHLD()

	case 0x3B:
		return GB.DECR16("SP")

	case 0x3C:
		return GB.INCR8("A")

	case 0x3D:
		return GB.DECR8("A")

	case 0x3E:
		return GB.LDR8N8("A")

	case 0x3F:
		return GB.CCF()

	case 0x40:
		return GB.LDR8R8("B", "B")

	case 0x41:
		return GB.LDR8R8("B", "C")

	case 0x42:
		return GB.LDR8R8("B", "D")

	case 0x43:
		return GB.LDR8R8("B", "E")

	case 0x44:
		return GB.LDR8R8("B", "H")

	case 0x45:
		return GB.LDR8R8("B", "L")

	case 0x46:
		return GB.LDR8HL("B")

	case 0x47:
		return GB.LDR8R8("B", "A")

	case 0x48:
		return GB.LDR8R8("C", "B")

	case 0x49:
		return GB.LDR8R8("C", "C")

	case 0x4A:
		return GB.LDR8R8("C", "D")

	case 0x4B:
		return GB.LDR8R8("C", "E")

	case 0x4C:
		return GB.LDR8R8("C", "H")

	case 0x4D:
		return GB.LDR8R8("C", "L")

	case 0x4E:
		return GB.LDR8HL("C")

	case 0x4F:
		return GB.LDR8R8("C", "A")

	case 0x50:
		return GB.LDR8R8("D", "B")

	case 0x51:
		return GB.LDR8R8("D", "C")

	case 0x52:
		return GB.LDR8R8("D", "D")

	case 0x53:
		return GB.LDR8R8("D", "E")

	case 0x54:
		return GB.LDR8R8("D", "H")

	case 0x55:
		return GB.LDR8R8("D", "L")

	case 0x56:
		return GB.LDR8HL("D")

	case 0x57:
		return GB.LDR8R8("D", "A")

	case 0x58:
		return GB.LDR8R8("E", "B")

	case 0x59:
		return GB.LDR8R8("E", "C")

	case 0x5A:
		return GB.LDR8R8("E", "D")

	case 0x5B:
		return GB.LDR8R8("E", "E")

	case 0x5C:
		return GB.LDR8R8("E", "H")

	case 0x5D:
		return GB.LDR8R8("E", "L")

	case 0x5E:
		return GB.LDR8HL("E")

	case 0x5F:
		return GB.LDR8R8("E", "A")

	case 0x60:
		return GB.LDR8R8("H", "B")

	case 0x61:
		return GB.LDR8R8("H", "C")

	case 0x62:
		return GB.LDR8R8("H", "D")

	case 0x63:
		return GB.LDR8R8("H", "E")

	case 0x64:
		return GB.LDR8R8("H", "H")

	case 0x65:
		return GB.LDR8R8("H", "L")

	case 0x66:
		return GB.LDR8HL("H")

	case 0x67:
		return GB.LDR8R8("H", "A")

	case 0x68:
		return GB.LDR8R8("L", "B")

	case 0x69:
		return GB.LDR8R8("L", "C")

	case 0x6A:
		return GB.LDR8R8("L", "D")

	case 0x6B:
		return GB.LDR8R8("L", "E")

	case 0x6C:
		return GB.LDR8R8("L", "H")

	case 0x6D:
		return GB.LDR8R8("L", "L")

	case 0x6E:
		return GB.LDR8HL("L")

	case 0x6F:
		return GB.LDR8R8("L", "A")

	case 0x70:
		return GB.LDHLR8("B")

	case 0x71:
		return GB.LDHLR8("C")

	case 0x72:
		return GB.LDHLR8("D")

	case 0x73:
		return GB.LDHLR8("E")

	case 0x74:
		return GB.LDHLR8("H")

	case 0x75:
		return GB.LDHLR8("L")

	case 0x76:
		return GB.HALT()

	case 0x77:
		return GB.LDHLR8("A")

	case 0x78:
		return GB.LDR8R8("A", "B")

	case 0x79:
		return GB.LDR8R8("A", "C")

	case 0x7A:
		return GB.LDR8R8("A", "D")

	case 0x7B:
		return GB.LDR8R8("A", "E")

	case 0x7C:
		return GB.LDR8R8("A", "H")

	case 0x7D:
		return GB.LDR8R8("A", "L")

	case 0x7E:
		return GB.LDR8HL("A")

	case 0x7F:
		return GB.LDR8R8("A", "A")

	case 0x80:
		return GB.ADDAR8("B")

	case 0x81:
		return GB.ADDAR8("C")

	case 0x82:
		return GB.ADDAR8("D")

	case 0x83:
		return GB.ADDAR8("E")

	case 0x84:
		return GB.ADDAR8("H")

	case 0x85:
		return GB.ADDAR8("L")

	case 0x86:
		return GB.ADDAHL()

	case 0x87:
		return GB.ADDAR8("A")

	case 0x88:
		return GB.ADCAR8("B")

	case 0x89:
		return GB.ADCAR8("C")

	case 0x8A:
		return GB.ADCAR8("D")

	case 0x8B:
		return GB.ADCAR8("E")

	case 0x8C:
		return GB.ADCAR8("H")

	case 0x8D:
		return GB.ADCAR8("L")

	case 0x8E:
		return GB.ADCAHL()

	case 0x8F:
		return GB.ADCAR8("A")

	case 0x90:
		return GB.SUBAR8("B")

	case 0x91:
		return GB.SUBAR8("C")

	case 0x92:
		return GB.SUBAR8("D")

	case 0x93:
		return GB.SUBAR8("E")

	case 0x94:
		return GB.SUBAR8("H")

	case 0x95:
		return GB.SUBAR8("L")

	case 0x96:
		return GB.SUBAHL()

	case 0x97:
		return GB.SUBAR8("A")

	case 0x98:
		return GB.SBCAR8("B")

	case 0x99:
		return GB.SBCAR8("C")

	case 0x9A:
		return GB.SBCAR8("D")

	case 0x9B:
		return GB.SBCAR8("E")

	case 0x9C:
		return GB.SBCAR8("H")

	case 0x9D:
		return GB.SBCAR8("L")

	case 0x9E:
		return GB.SBCAHL()

	case 0x9F:
		return GB.SBCAR8("A")

	case 0xA0:
		return GB.ANDAR8("B")

	case 0xA1:
		return GB.ANDAR8("C")

	case 0xA2:
		return GB.ANDAR8("D")

	case 0xA3:
		return GB.ANDAR8("E")

	case 0xA4:
		return GB.ANDAR8("H")

	case 0xA5:
		return GB.ANDAR8("L")

	case 0xA6:
		return GB.ANDAHL()

	case 0xA7:
		return GB.ANDAR8("A")

	case 0xA8:
		return GB.XORAR8("B")

	case 0xA9:
		return GB.XORAR8("C")

	case 0xAA:
		return GB.XORAR8("D")

	case 0xAB:
		return GB.XORAR8("E")

	case 0xAC:
		return GB.XORAR8("H")

	case 0xAD:
		return GB.XORAR8("L")

	case 0xAE:
		return GB.XORAHL()

	case 0xAF:
		return GB.XORAR8("A")

	case 0xB0:
		return GB.ORAR8("B")

	case 0xB1:
		return GB.ORAR8("C")

	case 0xB2:
		return GB.ORAR8("D")

	case 0xB3:
		return GB.ORAR8("E")

	case 0xB4:
		return GB.ORAR8("H")

	case 0xB5:
		return GB.ORAR8("L")

	case 0xB6:
		return GB.ORAHL()

	case 0xB7:
		return GB.ORAR8("A")

	case 0xB8:
		return GB.CPAR8("B")

	case 0xB9:
		return GB.CPAR8("C")

	case 0xBA:
		return GB.CPAR8("D")

	case 0xBB:
		return GB.CPAR8("E")

	case 0xBC:
		return GB.CPAR8("H")

	case 0xBD:
		return GB.CPAR8("L")

	case 0xBE:
		return GB.CPAHL()

	case 0xBF:
		return GB.CPAR8("A")

	case 0xC0:
		return GB.RETCC("NZ")

	case 0xC1:
		return GB.POPR16("BC")

	case 0xC2:
		return GB.JPCCN16("NZ")

	case 0xC3:
		addr := GB.MMU.Read(GB.CPU.Reg.Read("PC")+1) + GB.MMU.Read(GB.CPU.Reg.Read("PC")+2)<<8
		return GB.JPN16(addr)

	case 0xC4:
		return GB.CALLCCN16("NZ")

	case 0xC5:
		return GB.PUSHR16("BC")

	case 0xC6:
		return GB.ADDAN8()

	case 0xC7:
		return GB.RST("00H")

	case 0xC8:
		return GB.RETCC("Z")

	case 0xC9:
		return GB.RET()

	case 0xCA:
		return GB.JPCCN16("Z")

	case 0xCB:
		return GB.PrefixTable()

	case 0xCC:
		return GB.CALLCCN16("Z")

	case 0xCD:
		return GB.CALLN16()

	case 0xCE:
		return GB.ADCAN8()

	case 0xCF:
		return GB.RST("08H")

	case 0xD0:
		return GB.RETCC("NC")

	case 0xD1:
		return GB.POPR16("DE")

	case 0xD2:
		return GB.JPCCN16("NC")

	case 0xD4:
		return GB.CALLCCN16("NC")

	case 0xD5:
		return GB.PUSHR16("DE")

	case 0xD6:
		return GB.SUBAN8()

	case 0xD7:
		return GB.RST("10H")

	case 0xD8:
		return GB.RETCC("C")

	case 0xD9:
		return GB.RETI()

	case 0xDA:
		return GB.JPCCN16("C")

	case 0xDC:
		return GB.CALLCCN16("C")

	case 0xDE:
		return GB.SBCAN8()

	case 0xDF:
		return GB.RST("18H")

	case 0xE0:
		return GB.LDHU8A()

	case 0xE1:
		return GB.POPR16("HL")

	case 0xE2:
		return GB.LDHCA()

	case 0xE5:
		return GB.PUSHR16("HL")

	case 0xE6:
		return GB.ANDAN8()

	case 0xE7:
		return GB.RST("20H")

	case 0xE8:
		return GB.ADDSPE8()

	case 0xE9:
		return GB.JPHL()

	case 0xEA:
		return GB.LDN16A()

	case 0xEE:
		return GB.XORAN8()

	case 0xEF:
		return GB.RST("28H")

	case 0xF0:
		return GB.LDHAU8()

	case 0xF1:
		return GB.POPFA()

	case 0xF2:
		return GB.LDHAC()

	case 0xF3:
		return GB.DI()

	case 0xF4:

	case 0xF5:
		return GB.PUSHFA()

	case 0xF6:
		return GB.ORAN8()

	case 0xF7:
		return GB.RST("30H")

	case 0xF8:
		return GB.LDHLSPE8()

	case 0xF9:
		return GB.LDSPHL()

	case 0xFA:
		return GB.LDAN16()

	case 0xFB:
		return GB.EI()

	case 0xFE:
		return GB.CPAN8()

	case 0xFF:
		return GB.RST("38H")

	}
	return "", 0, 0, false
}
