package CPU

import (
	"fmt"
)

func (cpu *CPU) ExecuteOpcode() (string, int, int, bool) {

	switch cpu.Mem.Read(cpu.Reg.Read("PC")) {

	case 0x00:
		return cpu.NOP()

	case 0x01:
		return cpu.LDR16N16("BC")

	case 0x02:
		return cpu.LDR16A("BC")

	case 0x03:
		return cpu.INCR16("BC")

	case 0x04:
		return cpu.INCR8("B")

	case 0x05:
		return cpu.DECR8("B")

	case 0x06:
		return cpu.LDR8N8("B")

	case 0x07:
		return cpu.RLCA()

	case 0x08:
		return cpu.LDN16SP()

	case 0x09:
		return cpu.ADDHLR16("BC")

	case 0x0A:
		return cpu.LDAR16("BC")

	case 0x0B:
		return cpu.DECR16("BC")

	case 0x0C:
		return cpu.INCR8("C")

	case 0x0D:
		return cpu.DECR8("C")

	case 0x0E:
		return cpu.LDR8N8("C")

	case 0x0F:
		return cpu.RRCA()

	case 0x10:
		opcode := fmt.Sprintf("STOP $%x", int8(cpu.Mem.Read(cpu.Reg.Read("PC")+1)))
		return opcode, 0, 0, false

	case 0x11:
		return cpu.LDR16N16("DE")

	case 0x12:
		return cpu.LDR16A("DE")

	case 0x13:
		return cpu.INCR16("DE")

	case 0x14:
		return cpu.INCR8("D")

	case 0x15:
		return cpu.DECR8("D")

	case 0x16:
		return cpu.LDR8N8("D")

	case 0x17:
		return cpu.RLA()

	case 0x18:
		return cpu.JRN16()

	case 0x19:
		return cpu.ADDHLR16("DE")

	case 0x1A:
		return cpu.LDAR16("DE")

	case 0x1B:
		return cpu.DECR16("DE")

	case 0x1C:
		return cpu.INCR8("E")

	case 0x1D:
		return cpu.DECR8("E")

	case 0x1E:
		return cpu.LDR8N8("E")

	case 0x1F:
		return cpu.RRA()

	case 0x20:
		return cpu.JRCCN16("NZ")

	case 0x21:
		return cpu.LDR16N16("HL")

	case 0x22:
		return cpu.LDHLIA()

	case 0x23:
		return cpu.INCR16("HL")

	case 0x24:
		return cpu.INCR8("H")

	case 0x25:
		return cpu.DECR8("H")

	case 0x26:
		return cpu.LDR8N8("H")

	case 0x27:
		return cpu.DAA()

	case 0x28:
		return cpu.JRCCN16("Z")

	case 0x29:
		return cpu.ADDHLR16("HL")

	case 0x2A:
		return cpu.LDAHLI()

	case 0x2B:
		return cpu.DECR16("HL")

	case 0x2C:
		return cpu.INCR8("L")

	case 0x2D:
		return cpu.DECR8("L")

	case 0x2E:
		return cpu.LDR8N8("L")

	case 0x2F:
		return cpu.CPL()

	case 0x30:
		return cpu.JRCCN16("NC")

	case 0x31:
		return cpu.LDR16N16("SP")

	case 0x32:
		return cpu.LDHLDA()

	case 0x33:
		return cpu.INCR16("SP")

	case 0x34:
		return cpu.INCRHL()

	case 0x35:
		return cpu.DECHL()

	case 0x36:
		return cpu.LDHLN8()

	case 0x37:
		return cpu.SCF()

	case 0x38:
		return cpu.JRCCN16("C")

	case 0x39:
		return cpu.ADDHLSP()

	case 0x3A:
		return cpu.LDAHLD()

	case 0x3B:
		return cpu.DECR16("SP")

	case 0x3C:
		return cpu.INCR8("A")

	case 0x3D:
		return cpu.DECR8("A")

	case 0x3E:
		return cpu.LDR8N8("A")

	case 0x3F:
		return cpu.CCF()

	case 0x40:
		return cpu.LDR8R8("B", "B")

	case 0x41:
		return cpu.LDR8R8("B", "C")

	case 0x42:
		return cpu.LDR8R8("B", "D")

	case 0x43:
		return cpu.LDR8R8("B", "E")

	case 0x44:
		return cpu.LDR8R8("B", "H")

	case 0x45:
		return cpu.LDR8R8("B", "L")

	case 0x46:
		return cpu.LDR8HL("B")

	case 0x47:
		return cpu.LDR8R8("B", "A")

	case 0x48:
		return cpu.LDR8R8("C", "B")

	case 0x49:
		return cpu.LDR8R8("C", "C")

	case 0x4A:
		return cpu.LDR8R8("C", "D")

	case 0x4B:
		return cpu.LDR8R8("C", "E")

	case 0x4C:
		return cpu.LDR8R8("C", "H")

	case 0x4D:
		return cpu.LDR8R8("C", "L")

	case 0x4E:
		return cpu.LDR8HL("C")

	case 0x4F:
		return cpu.LDR8R8("C", "A")

	case 0x50:
		return cpu.LDR8R8("D", "B")

	case 0x51:
		return cpu.LDR8R8("D", "C")

	case 0x52:
		return cpu.LDR8R8("D", "D")

	case 0x53:
		return cpu.LDR8R8("D", "E")

	case 0x54:
		return cpu.LDR8R8("D", "H")

	case 0x55:
		return cpu.LDR8R8("D", "L")

	case 0x56:
		return cpu.LDR8HL("D")

	case 0x57:
		return cpu.LDR8R8("D", "A")

	case 0x58:
		return cpu.LDR8R8("E", "B")

	case 0x59:
		return cpu.LDR8R8("E", "C")

	case 0x5A:
		return cpu.LDR8R8("E", "D")

	case 0x5B:
		return cpu.LDR8R8("E", "E")

	case 0x5C:
		return cpu.LDR8R8("E", "H")

	case 0x5D:
		return cpu.LDR8R8("E", "L")

	case 0x5E:
		return cpu.LDR8HL("E")

	case 0x5F:
		return cpu.LDR8R8("E", "A")

	case 0x60:
		return cpu.LDR8R8("H", "B")

	case 0x61:
		return cpu.LDR8R8("H", "C")

	case 0x62:
		return cpu.LDR8R8("H", "D")

	case 0x63:
		return cpu.LDR8R8("H", "E")

	case 0x64:
		return cpu.LDR8R8("H", "H")

	case 0x65:
		return cpu.LDR8R8("H", "L")

	case 0x66:
		return cpu.LDR8HL("H")

	case 0x67:
		return cpu.LDR8R8("H", "A")

	case 0x68:
		return cpu.LDR8R8("L", "B")

	case 0x69:
		return cpu.LDR8R8("L", "C")

	case 0x6A:
		return cpu.LDR8R8("L", "D")

	case 0x6B:
		return cpu.LDR8R8("L", "E")

	case 0x6C:
		return cpu.LDR8R8("L", "H")

	case 0x6D:
		return cpu.LDR8R8("L", "L")

	case 0x6E:
		return cpu.LDR8HL("L")

	case 0x6F:
		return cpu.LDR8R8("L", "A")

	case 0x70:
		return cpu.LDHLR8("B")

	case 0x71:
		return cpu.LDHLR8("C")

	case 0x72:
		return cpu.LDHLR8("D")

	case 0x73:
		return cpu.LDHLR8("E")

	case 0x74:
		return cpu.LDHLR8("H")

	case 0x75:
		return cpu.LDHLR8("L")

	case 0x76:
		return cpu.HALT()

	case 0x77:
		return cpu.LDHLR8("A")

	case 0x78:
		return cpu.LDR8R8("A", "B")

	case 0x79:
		return cpu.LDR8R8("A", "C")

	case 0x7A:
		return cpu.LDR8R8("A", "D")

	case 0x7B:
		return cpu.LDR8R8("A", "E")

	case 0x7C:
		return cpu.LDR8R8("A", "H")

	case 0x7D:
		return cpu.LDR8R8("A", "L")

	case 0x7E:
		return cpu.LDR8HL("A")

	case 0x7F:
		return cpu.LDR8R8("A", "A")

	case 0x80:
		return cpu.ADDAR8("B")

	case 0x81:
		return cpu.ADDAR8("C")

	case 0x82:
		return cpu.ADDAR8("D")

	case 0x83:
		return cpu.ADDAR8("E")

	case 0x84:
		return cpu.ADDAR8("H")

	case 0x85:
		return cpu.ADDAR8("L")

	case 0x86:
		return cpu.ADDAHL()

	case 0x87:
		return cpu.ADDAR8("A")

	case 0x88:
		return cpu.ADCAR8("B")

	case 0x89:
		return cpu.ADCAR8("C")

	case 0x8A:
		return cpu.ADCAR8("D")

	case 0x8B:
		return cpu.ADCAR8("E")

	case 0x8C:
		return cpu.ADCAR8("H")

	case 0x8D:
		return cpu.ADCAR8("L")

	case 0x8E:
		return cpu.ADCAHL()

	case 0x8F:
		return cpu.ADCAR8("A")

	case 0x90:
		return cpu.SUBAR8("B")

	case 0x91:
		return cpu.SUBAR8("C")

	case 0x92:
		return cpu.SUBAR8("D")

	case 0x93:
		return cpu.SUBAR8("E")

	case 0x94:
		return cpu.SUBAR8("H")

	case 0x95:
		return cpu.SUBAR8("L")

	case 0x96:
		return cpu.SUBAHL()

	case 0x97:
		return cpu.SUBAR8("A")

	case 0x98:
		return cpu.SBCAR8("B")

	case 0x99:
		return cpu.SBCAR8("C")

	case 0x9A:
		return cpu.SBCAR8("D")

	case 0x9B:
		return cpu.SBCAR8("E")

	case 0x9C:
		return cpu.SBCAR8("H")

	case 0x9D:
		return cpu.SBCAR8("L")

	case 0x9E:
		return cpu.SBCAHL()

	case 0x9F:
		return cpu.SBCAR8("A")

	case 0xA0:
		return cpu.ANDAR8("B")

	case 0xA1:
		return cpu.ANDAR8("C")

	case 0xA2:
		return cpu.ANDAR8("D")

	case 0xA3:
		return cpu.ANDAR8("E")

	case 0xA4:
		return cpu.ANDAR8("H")

	case 0xA5:
		return cpu.ANDAR8("L")

	case 0xA6:
		return cpu.ANDAHL()

	case 0xA7:
		return cpu.ANDAR8("A")

	case 0xA8:
		return cpu.XORAR8("B")

	case 0xA9:
		return cpu.XORAR8("C")

	case 0xAA:
		return cpu.XORAR8("D")

	case 0xAB:
		return cpu.XORAR8("E")

	case 0xAC:
		return cpu.XORAR8("H")

	case 0xAD:
		return cpu.XORAR8("L")

	case 0xAE:
		return cpu.XORAHL()

	case 0xAF:
		return cpu.XORAR8("A")

	case 0xB0:
		return cpu.ORAR8("B")

	case 0xB1:
		return cpu.ORAR8("C")

	case 0xB2:
		return cpu.ORAR8("D")

	case 0xB3:
		return cpu.ORAR8("E")

	case 0xB4:
		return cpu.ORAR8("H")

	case 0xB5:
		return cpu.ORAR8("L")

	case 0xB6:
		return cpu.ORAHL()

	case 0xB7:
		return cpu.ORAR8("A")

	case 0xB8:
		return cpu.CPAR8("B")

	case 0xB9:
		return cpu.CPAR8("C")

	case 0xBA:
		return cpu.CPAR8("D")

	case 0xBB:
		return cpu.CPAR8("E")

	case 0xBC:
		return cpu.CPAR8("H")

	case 0xBD:
		return cpu.CPAR8("L")

	case 0xBE:
		return cpu.CPAHL()

	case 0xBF:
		return cpu.CPAR8("A")

	case 0xC0:
		return cpu.RETCC("NZ")

	case 0xC1:
		return cpu.POPR16("BC")

	case 0xC2:
		return cpu.JPCCN16("NZ")

	case 0xC3:
		addr := cpu.Mem.Read(cpu.Reg.Read("PC")+1) + cpu.Mem.Read(cpu.Reg.Read("PC")+2)<<8
		return cpu.JPN16(addr)

	case 0xC4:
		return cpu.CALLCCN16("NZ")

	case 0xC5:
		return cpu.PUSHR16("BC")

	case 0xC6:
		return cpu.ADDAN8()

	case 0xC7:
		return cpu.RST("00H")

	case 0xC8:
		return cpu.RETCC("Z")

	case 0xC9:
		return cpu.RET()

	case 0xCA:
		return cpu.JPCCN16("Z")

	case 0xCB:
		return cpu.PrefixTable()

	case 0xCC:
		return cpu.CALLCCN16("Z")

	case 0xCD:
		return cpu.CALLN16()

	case 0xCE:
		return cpu.ADCAN8()

	case 0xCF:
		return cpu.RST("08H")

	case 0xD0:
		return cpu.RETCC("NC")

	case 0xD1:
		return cpu.POPR16("DE")

	case 0xD2:
		return cpu.JPCCN16("NC")

	case 0xD4:
		return cpu.CALLCCN16("NC")

	case 0xD5:
		return cpu.PUSHR16("DE")

	case 0xD6:
		return cpu.SUBAN8()

	case 0xD7:
		return cpu.RST("10H")

	case 0xD8:
		return cpu.RETCC("C")

	case 0xD9:
		return cpu.RETI()

	case 0xDA:
		return cpu.JPCCN16("C")

	case 0xDC:
		return cpu.CALLCCN16("C")

	case 0xDE:
		return cpu.SBCAN8()

	case 0xDF:
		return cpu.RST("18H")

	case 0xE0:
		return cpu.LDHU8A()

	case 0xE1:
		return cpu.POPR16("HL")

	case 0xE2:
		return cpu.LDHCA()

	case 0xE5:
		return cpu.PUSHR16("HL")

	case 0xE6:
		return cpu.ANDAN8()

	case 0xE7:
		return cpu.RST("20H")

	case 0xE8:
		return cpu.ADDSPE8()

	case 0xE9:
		return cpu.JPHL()

	case 0xEA:
		return cpu.LDN16A()

	case 0xEE:
		return cpu.XORAN8()

	case 0xEF:
		return cpu.RST("28H")

	case 0xF0:
		return cpu.LDHAU8()

	case 0xF1:
		return cpu.POPFA()

	case 0xF2:
		return cpu.LDHAC()

	case 0xF3:
		return cpu.DI()

	case 0xF4:

	case 0xF5:
		return cpu.PUSHFA()

	case 0xF6:
		return cpu.ORAN8()

	case 0xF7:
		return cpu.RST("30H")

	case 0xF8:
		return cpu.LDHLSPE8()

	case 0xF9:
		return cpu.LDSPHL()

	case 0xFA:
		return cpu.LDAN16()

	case 0xFB:
		return cpu.EI()

	case 0xFE:
		return cpu.CPAN8()

	case 0xFF:
		return cpu.RST("38H")

	}
	return "", 0, 0, false
}
