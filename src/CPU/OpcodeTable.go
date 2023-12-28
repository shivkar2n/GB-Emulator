package CPU

import (
	"fmt"
)

func (s *CPU) ExecuteOpcode() {

	switch s.Mem.Read(s.GetReg16Val("PC")) {
	case 0x00:
		s.CurrOpcode = fmt.Sprintf("NOP")
		s.NOP()

	case 0x01:
		s.CurrOpcode = fmt.Sprintf("LD BC,$%x ", uint16(int(s.Mem.Read(s.GetReg16Val("PC")+1))+(int(s.Mem.Read(s.GetReg16Val("PC")+2))<<8)))
		s.LDR16N16("BC")

	case 0x02:
		s.CurrOpcode = fmt.Sprintf("LD (BC), A")
		s.LDR16A("BC")

	case 0x03:
		s.CurrOpcode = fmt.Sprintf("INC BC")
		s.INCR16("BC")

	case 0x04:
		s.CurrOpcode = fmt.Sprintf("INC B")
		s.INCR8("B")

	case 0x05:
		s.CurrOpcode = fmt.Sprintf("DEC B")
		s.DECR8("B")

	case 0x06:
		s.CurrOpcode = fmt.Sprintf("LD B,$%x", uint8(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.LDR8N8("B")

	case 0x07:
		s.CurrOpcode = fmt.Sprintf("RLCA")
		s.RLCA()

	case 0x08:
		s.CurrOpcode = fmt.Sprintf("LD (%x), SP", int(s.Mem.Read(s.GetReg16Val("PC")+1))+(int(s.Mem.Read(s.GetReg16Val("PC")+2))<<8))
		s.LDN16SP()

	case 0x09:
		s.CurrOpcode = fmt.Sprintf("ADD HL, BC")
		s.ADDHLR16("BC")

	case 0x0A:
		s.CurrOpcode = fmt.Sprintf("LD A, (BC)")
		s.LDAR16("BC")

	case 0x0B:
		s.CurrOpcode = fmt.Sprintf("DEC BC")
		s.DECR16("BC")

	case 0x0C:
		s.CurrOpcode = fmt.Sprintf("INC C")
		s.INCR8("C")

	case 0x0D:
		s.CurrOpcode = fmt.Sprintf("DEC C")
		s.DECR8("C")

	case 0x0E:
		s.CurrOpcode = fmt.Sprintf("LD C, $%x", int(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.LDR8N8("C")

	case 0x0F:
		s.CurrOpcode = fmt.Sprintf("RRCA ")
		s.RRCA()

	case 0x10:
		s.CurrOpcode = fmt.Sprintf("STOP $%x", int8(s.Mem.Read(s.GetReg16Val("PC")+1)))

	case 0x11:
		s.CurrOpcode = fmt.Sprintf("LD DE,$%x", uint16(int(s.Mem.Read(s.GetReg16Val("PC")+1))+int(s.Mem.Read(s.GetReg16Val("PC")+2)<<8)))
		s.LDR16N16("DE")

	case 0x12:
		s.CurrOpcode = fmt.Sprintf("LD (DE),A")
		s.LDR16A("DE")

	case 0x13:
		s.CurrOpcode = fmt.Sprintf("INC (DE)")
		s.INCR16("DE")

	case 0x14:
		s.CurrOpcode = fmt.Sprintf("INC D")
		s.INCR8("D")

	case 0x15:
		s.CurrOpcode = fmt.Sprintf("DEC D")
		s.DECR8("D")

	case 0x16:
		s.CurrOpcode = fmt.Sprintf("LD D,$%x", uint8(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.LDR8N8("D")

	case 0x17:
		s.CurrOpcode = fmt.Sprintf("RLA")
		s.RLA()

	case 0x18:
		s.CurrOpcode = fmt.Sprintf("JR $%x", int8(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.JRN16()

	case 0x19:
		s.CurrOpcode = fmt.Sprintf("ADD HL, DE")
		s.ADDHLR16("DE")

	case 0x1A:
		s.CurrOpcode = fmt.Sprintf("LD A, (DE)")
		s.LDAR16("DE")

	case 0x1B:
		s.CurrOpcode = fmt.Sprintf("DEC DE")
		s.DECR16("DE")

	case 0x1C:
		s.CurrOpcode = fmt.Sprintf("INC E")
		s.INCR8("E")

	case 0x1D:
		s.CurrOpcode = fmt.Sprintf("DEC E")
		s.DECR8("E")

	case 0x1E:
		s.CurrOpcode = fmt.Sprintf("LD E,$%x", uint8(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.LDR8N8("E")

	case 0x1F:
		s.CurrOpcode = fmt.Sprintf("RRA")
		s.RRA()

	case 0x20:
		s.CurrOpcode = fmt.Sprintf("JR NZ,$%x", uint8(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.JRCCN16("NZ")

	case 0x21:
		s.CurrOpcode = fmt.Sprintf("LD HL,$%x", uint16(int(s.Mem.Read(s.GetReg16Val("PC")+1))+(int(s.Mem.Read(s.GetReg16Val("PC")+2))<<8)))
		s.LDR16N16("HL")

	case 0x22:
		s.CurrOpcode = fmt.Sprintf("LD (HL+),A")
		s.LDHLIA()

	case 0x23:
		s.CurrOpcode = fmt.Sprintf("INC HL")
		s.INCR16("HL")

	case 0x24:
		s.CurrOpcode = fmt.Sprintf("INC H")
		s.INCR8("H")

	case 0x25:
		s.CurrOpcode = fmt.Sprintf("DEC H")
		s.DECR8("H")

	case 0x26:
		s.CurrOpcode = fmt.Sprintf("LD H,$%x", uint8(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.LDR8N8("H")

	case 0x27:
		s.CurrOpcode = fmt.Sprintf("DAA")
		s.DAA()

	case 0x28:
		s.CurrOpcode = fmt.Sprintf("JR Z,$%x", int8(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.JRCCN16("Z")

	case 0x29:
		s.CurrOpcode = fmt.Sprintf("ADD HL,HL")
		s.ADDHLR16("HL")

	case 0x2A:
		s.CurrOpcode = fmt.Sprintf("LD A,(HL+)")
		s.LDAHLI()

	case 0x2B:
		s.CurrOpcode = fmt.Sprintf("DEC HL")
		s.DECR16("HL")

	case 0x2C:
		s.CurrOpcode = fmt.Sprintf("INC L")
		s.INCR8("L")

	case 0x2D:
		s.CurrOpcode = fmt.Sprintf("DEC L")
		s.DECR8("L")

	case 0x2E:
		s.CurrOpcode = fmt.Sprintf("LD L,$%x", uint8(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.LDR8N8("L")

	case 0x2F:
		s.CurrOpcode = fmt.Sprintf("CPL")
		s.CPL()

	case 0x30:
		s.CurrOpcode = fmt.Sprintf("JR NC,$%x", int8(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.JRCCN16("NC")

	case 0x31:
		s.CurrOpcode = fmt.Sprintf("LD SP $%x", uint16(int(s.Mem.Read(s.GetReg16Val("PC")+1))+(int(s.Mem.Read(s.GetReg16Val("PC")+2))<<8)))
		s.LDR16N16("SP")

	case 0x32:
		s.CurrOpcode = fmt.Sprintf("LD (HL-),A")
		s.LDHLDA()

	case 0x33:
		s.CurrOpcode = fmt.Sprintf("INC SP")
		s.INCR16("SP")

	case 0x34:
		s.CurrOpcode = fmt.Sprintf("INC (HL)")
		s.INCRHL()

	case 0x35:
		s.CurrOpcode = fmt.Sprintf("DEC (HL)")
		s.DECHL()

	case 0x36:
		s.CurrOpcode = fmt.Sprintf("LD (HL),$%x", uint8(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.LDHLN8()

	case 0x37:
		s.CurrOpcode = fmt.Sprintf("SCF")
		s.SCF()

	case 0x38:
		s.CurrOpcode = fmt.Sprintf("JR C,$%x", int8(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.JRCCN16("C")

	case 0x39:
		s.CurrOpcode = fmt.Sprintf("ADD HL,SP")
		s.ADDHLSP()

	case 0x3A:
		s.CurrOpcode = fmt.Sprintf("LD A,(HL-)")
		s.LDAHLD()

	case 0x3B:
		s.CurrOpcode = fmt.Sprintf("DEC SP")
		s.DECR16("SP")

	case 0x3C:
		s.CurrOpcode = fmt.Sprintf("INC A")
		s.INCR8("A")

	case 0x3D:
		s.CurrOpcode = fmt.Sprintf("DEC A")
		s.DECR8("A")

	case 0x3E:
		s.CurrOpcode = fmt.Sprintf("LD A,$%x", uint8(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.LDR8N8("A")

	case 0x3F:
		s.CurrOpcode = fmt.Sprintf("CCF")
		s.CCF()

	case 0x40:
		s.CurrOpcode = fmt.Sprintf("LD B,B")
		s.LDR8R8("B", "B")

	case 0x41:
		s.CurrOpcode = fmt.Sprintf("LD B,C")
		s.LDR8R8("B", "C")

	case 0x42:
		s.CurrOpcode = fmt.Sprintf("LD B,D")
		s.LDR8R8("B", "D")

	case 0x43:
		s.CurrOpcode = fmt.Sprintf("LD B,E")
		s.LDR8R8("B", "E")

	case 0x44:
		s.CurrOpcode = fmt.Sprintf("LD B,H")
		s.LDR8R8("B", "H")

	case 0x45:
		s.CurrOpcode = fmt.Sprintf("LD B,L")
		s.LDR8R8("B", "L")

	case 0x46:
		s.CurrOpcode = fmt.Sprintf("LD B,(HL)")
		s.LDR8HL("B")

	case 0x47:
		s.CurrOpcode = fmt.Sprintf("LD B,A")
		s.LDR8R8("B", "A")

	case 0x48:
		s.CurrOpcode = fmt.Sprintf("LD C,B")
		s.LDR8R8("C", "B")

	case 0x49:
		s.CurrOpcode = fmt.Sprintf("LD C,C")
		s.LDR8R8("C", "C")

	case 0x4A:
		s.CurrOpcode = fmt.Sprintf("LD C,D")
		s.LDR8R8("C", "D")

	case 0x4B:
		s.CurrOpcode = fmt.Sprintf("LD C,E")
		s.LDR8R8("C", "E")

	case 0x4C:
		s.CurrOpcode = fmt.Sprintf("LD C,H")
		s.LDR8R8("C", "H")

	case 0x4D:
		s.CurrOpcode = fmt.Sprintf("LD C,L")
		s.LDR8R8("C", "L")

	case 0x4E:
		s.CurrOpcode = fmt.Sprintf("LD C,(HL)")
		s.LDR8HL("C")

	case 0x4F:
		s.CurrOpcode = fmt.Sprintf("LD C,A")
		s.LDR8R8("C", "A")

	case 0x50:
		s.CurrOpcode = fmt.Sprintf("LD D,B")
		s.LDR8R8("D", "B")

	case 0x51:
		s.CurrOpcode = fmt.Sprintf("LD D,C")
		s.LDR8R8("D", "C")

	case 0x52:
		s.CurrOpcode = fmt.Sprintf("LD D,D")
		s.LDR8R8("D", "D")

	case 0x53:
		s.CurrOpcode = fmt.Sprintf("LD D,E")
		s.LDR8R8("D", "E")

	case 0x54:
		s.CurrOpcode = fmt.Sprintf("LD D,H")
		s.LDR8R8("D", "H")

	case 0x55:
		s.CurrOpcode = fmt.Sprintf("LD D,L")
		s.LDR8R8("D", "L")

	case 0x56:
		s.CurrOpcode = fmt.Sprintf("LD D,(HL)")
		s.LDR8HL("D")

	case 0x57:
		s.CurrOpcode = fmt.Sprintf("LD D,A")
		s.LDR8R8("D", "A")

	case 0x58:
		s.CurrOpcode = fmt.Sprintf("LD E,B")
		s.LDR8R8("E", "B")

	case 0x59:
		s.CurrOpcode = fmt.Sprintf("LD E,C")
		s.LDR8R8("E", "C")

	case 0x5A:
		s.CurrOpcode = fmt.Sprintf("LD E,D")
		s.LDR8R8("E", "D")

	case 0x5B:
		s.CurrOpcode = fmt.Sprintf("LD E,E")
		s.LDR8R8("E", "E")

	case 0x5C:
		s.CurrOpcode = fmt.Sprintf("LD E,H")
		s.LDR8R8("E", "H")

	case 0x5D:
		s.CurrOpcode = fmt.Sprintf("LD E,L")
		s.LDR8R8("E", "L")

	case 0x5E:
		s.CurrOpcode = fmt.Sprintf("LD E,(HL)")
		s.LDR8HL("E")

	case 0x5F:
		s.CurrOpcode = fmt.Sprintf("LD E,A")
		s.LDR8R8("E", "A")

	case 0x60:
		s.CurrOpcode = fmt.Sprintf("LD H,B")
		s.LDR8R8("H", "B")

	case 0x61:
		s.CurrOpcode = fmt.Sprintf("LD H,C")
		s.LDR8R8("H", "C")

	case 0x62:
		s.CurrOpcode = fmt.Sprintf("LD H,D")
		s.LDR8R8("H", "D")

	case 0x63:
		s.CurrOpcode = fmt.Sprintf("LD H,E")
		s.LDR8R8("H", "E")

	case 0x64:
		s.CurrOpcode = fmt.Sprintf("LD H,H")
		s.LDR8R8("H", "H")

	case 0x65:
		s.CurrOpcode = fmt.Sprintf("LD H,L")
		s.LDR8R8("H", "L")

	case 0x66:
		s.CurrOpcode = fmt.Sprintf("LD H,(HL)")
		s.LDR8HL("H")

	case 0x67:
		s.CurrOpcode = fmt.Sprintf("LD H,A")
		s.LDR8R8("H", "A")

	case 0x68:
		s.CurrOpcode = fmt.Sprintf("LD L,B")
		s.LDR8R8("L", "B")

	case 0x69:
		s.CurrOpcode = fmt.Sprintf("LD L,C")
		s.LDR8R8("L", "C")

	case 0x6A:
		s.CurrOpcode = fmt.Sprintf("LD L,D")
		s.LDR8R8("L", "D")

	case 0x6B:
		s.CurrOpcode = fmt.Sprintf("LD L,E")
		s.LDR8R8("L", "E")

	case 0x6C:
		s.CurrOpcode = fmt.Sprintf("LD L,H")
		s.LDR8R8("L", "H")

	case 0x6D:
		s.CurrOpcode = fmt.Sprintf("LD L,L")
		s.LDR8R8("L", "L")

	case 0x6E:
		s.CurrOpcode = fmt.Sprintf("LD L,(HL)")
		s.LDR8HL("L")

	case 0x6F:
		s.CurrOpcode = fmt.Sprintf("LD L,A")
		s.LDR8R8("L", "A")

	case 0x70:
		s.CurrOpcode = fmt.Sprintf("LD (HL),B")
		s.LDHLR8("B")

	case 0x71:
		s.CurrOpcode = fmt.Sprintf("LD (HL),C")
		s.LDHLR8("C")

	case 0x72:
		s.CurrOpcode = fmt.Sprintf("LD (HL),D")
		s.LDHLR8("D")

	case 0x73:
		s.CurrOpcode = fmt.Sprintf("LD (HL),E")
		s.LDHLR8("E")

	case 0x74:
		s.CurrOpcode = fmt.Sprintf("LD (HL),H")
		s.LDHLR8("H")

	case 0x75:
		s.CurrOpcode = fmt.Sprintf("LD (HL),L")
		s.LDHLR8("L")

	case 0x76:
		s.CurrOpcode = fmt.Sprintf("HALT")
		s.HALT()

	case 0x77:
		s.CurrOpcode = fmt.Sprintf("LD (HL),A")
		s.LDHLR8("A")

	case 0x78:
		s.CurrOpcode = fmt.Sprintf("LD A,B")
		s.LDR8R8("A", "B")

	case 0x79:
		s.CurrOpcode = fmt.Sprintf("LD A,C")
		s.LDR8R8("A", "C")

	case 0x7A:
		s.CurrOpcode = fmt.Sprintf("LD A,D")
		s.LDR8R8("A", "D")

	case 0x7B:
		s.CurrOpcode = fmt.Sprintf("LD A,E")
		s.LDR8R8("A", "E")

	case 0x7C:
		s.CurrOpcode = fmt.Sprintf("LD A,H")
		s.LDR8R8("A", "H")

	case 0x7D:
		s.CurrOpcode = fmt.Sprintf("LD A,L")
		s.LDR8R8("A", "L")

	case 0x7E:
		s.CurrOpcode = fmt.Sprintf("LD A,(HL)")
		s.LDR8HL("A")

	case 0x7F:
		s.CurrOpcode = fmt.Sprintf("LD A,A")
		s.LDR8R8("A", "A")

	case 0x80:
		s.CurrOpcode = fmt.Sprintf("ADD A,B")
		s.ADDAR8("B")

	case 0x81:
		s.CurrOpcode = fmt.Sprintf("ADD A,C")
		s.ADDAR8("C")

	case 0x82:
		s.CurrOpcode = fmt.Sprintf("ADD A,D")
		s.ADDAR8("D")

	case 0x83:
		s.ADDAR8("E")
		s.CurrOpcode = fmt.Sprintf("ADD A,E")

	case 0x84:
		s.CurrOpcode = fmt.Sprintf("ADD A,H")
		s.ADDAR8("H")

	case 0x85:
		s.CurrOpcode = fmt.Sprintf("ADD A,L")
		s.ADDAR8("L")

	case 0x86:
		s.CurrOpcode = fmt.Sprintf("ADD A,(HL)")
		s.ADDAHL()

	case 0x87:
		s.CurrOpcode = fmt.Sprintf("ADD A,A")
		s.ADDAR8("A")

	case 0x88:
		s.CurrOpcode = fmt.Sprintf("ADC A,B")
		s.ADCAR8("B")

	case 0x89:
		s.CurrOpcode = fmt.Sprintf("ADC A,C")
		s.ADCAR8("C")

	case 0x8A:
		s.CurrOpcode = fmt.Sprintf("ADC A,D")
		s.ADCAR8("D")

	case 0x8B:
		s.CurrOpcode = fmt.Sprintf("ADC A,E")
		s.ADCAR8("E")

	case 0x8C:
		s.CurrOpcode = fmt.Sprintf("ADC A,H")
		s.ADCAR8("H")

	case 0x8D:
		s.CurrOpcode = fmt.Sprintf("ADC A,L")
		s.ADCAR8("L")

	case 0x8E:
		s.CurrOpcode = fmt.Sprintf("ADC A,(HL)")
		s.ADCAHL()

	case 0x8F:
		s.CurrOpcode = fmt.Sprintf("ADC A,A")
		s.ADCAR8("A")

	case 0x90:
		s.CurrOpcode = fmt.Sprintf("SUB A,B")
		s.SUBAR8("B")

	case 0x91:
		s.CurrOpcode = fmt.Sprintf("SUB A,C")
		s.SUBAR8("C")

	case 0x92:
		s.CurrOpcode = fmt.Sprintf("SUB A,D")
		s.SUBAR8("D")

	case 0x93:
		s.CurrOpcode = fmt.Sprintf("SUB A,E")
		s.SUBAR8("E")

	case 0x94:
		s.CurrOpcode = fmt.Sprintf("SUB A,H")
		s.SUBAR8("H")

	case 0x95:
		s.CurrOpcode = fmt.Sprintf("SUB A,L")
		s.SUBAR8("L")

	case 0x96:
		s.CurrOpcode = fmt.Sprintf("SUB A,(HL)")
		s.SUBAHL()

	case 0x97:
		s.CurrOpcode = fmt.Sprintf("SUB A,A")
		s.SUBAR8("A")

	case 0x98:
		s.CurrOpcode = fmt.Sprintf("SBC A,B")
		s.SBCAR8("B")

	case 0x99:
		s.CurrOpcode = fmt.Sprintf("SBC A,C")
		s.SBCAR8("C")

	case 0x9A:
		s.CurrOpcode = fmt.Sprintf("SBC A,D")
		s.SBCAR8("D")

	case 0x9B:
		s.CurrOpcode = fmt.Sprintf("SBC A,E")
		s.SBCAR8("E")

	case 0x9C:
		s.CurrOpcode = fmt.Sprintf("SBC A,H")
		s.SBCAR8("H")

	case 0x9D:
		s.CurrOpcode = fmt.Sprintf("SBC A,L")
		s.SBCAR8("L")

	case 0x9E:
		s.CurrOpcode = fmt.Sprintf("SBC A,(HL)")
		s.SBCAHL()

	case 0x9F:
		s.CurrOpcode = fmt.Sprintf("SBC A,A")
		s.SBCAR8("A")

	case 0xA0:
		s.CurrOpcode = fmt.Sprintf("AND A,B")
		s.ANDAR8("B")

	case 0xA1:
		s.CurrOpcode = fmt.Sprintf("AND A,C")
		s.ANDAR8("C")

	case 0xA2:
		s.CurrOpcode = fmt.Sprintf("AND A,D")
		s.ANDAR8("D")

	case 0xA3:
		s.CurrOpcode = fmt.Sprintf("AND A,E")
		s.ANDAR8("E")

	case 0xA4:
		s.CurrOpcode = fmt.Sprintf("AND A,H")
		s.ANDAR8("H")

	case 0xA5:
		s.CurrOpcode = fmt.Sprintf("AND A,L")
		s.ANDAR8("L")

	case 0xA6:
		s.CurrOpcode = fmt.Sprintf("AND A,(HL)")
		s.ANDAHL()

	case 0xA7:
		s.CurrOpcode = fmt.Sprintf("AND A,A")
		s.ANDAR8("A")

	case 0xA8:
		s.CurrOpcode = fmt.Sprintf("XOR A,B")
		s.XORAR8("B")

	case 0xA9:
		s.CurrOpcode = fmt.Sprintf("XOR A,C")
		s.XORAR8("C")

	case 0xAA:
		s.CurrOpcode = fmt.Sprintf("XOR A,D")
		s.XORAR8("D")

	case 0xAB:
		s.CurrOpcode = fmt.Sprintf("XOR A,E")
		s.XORAR8("E")

	case 0xAC:
		s.CurrOpcode = fmt.Sprintf("XOR A,H")
		s.XORAR8("H")

	case 0xAD:
		s.CurrOpcode = fmt.Sprintf("XOR A,L")
		s.XORAR8("L")

	case 0xAE:
		s.CurrOpcode = fmt.Sprintf("XOR A,(HL)")
		s.XORAHL()

	case 0xAF:
		s.CurrOpcode = fmt.Sprintf("XOR A,A")
		s.XORAR8("A")

	case 0xB0:
		s.CurrOpcode = fmt.Sprintf("OR A,B")
		s.ORAR8("B")

	case 0xB1:
		s.CurrOpcode = fmt.Sprintf("OR A,C")
		s.ORAR8("C")

	case 0xB2:
		s.CurrOpcode = fmt.Sprintf("OR A,D")
		s.ORAR8("D")

	case 0xB3:
		s.CurrOpcode = fmt.Sprintf("OR A,E")
		s.ORAR8("E")

	case 0xB4:
		s.CurrOpcode = fmt.Sprintf("OR A,H")
		s.ORAR8("H")

	case 0xB5:
		s.CurrOpcode = fmt.Sprintf("OR A,L")
		s.ORAR8("L")

	case 0xB6:
		s.CurrOpcode = fmt.Sprintf("OR A,(HL)")
		s.ORAHL()

	case 0xB7:
		s.CurrOpcode = fmt.Sprintf("OR A,A")
		s.ORAR8("A")

	case 0xB8:
		s.CurrOpcode = fmt.Sprintf("CP A,B")
		s.CPAR8("B")

	case 0xB9:
		s.CurrOpcode = fmt.Sprintf("CP A,C")
		s.CPAR8("C")

	case 0xBA:
		s.CurrOpcode = fmt.Sprintf("CP A,D")
		s.CPAR8("D")

	case 0xBB:
		s.CurrOpcode = fmt.Sprintf("CP A,E")
		s.CPAR8("E")

	case 0xBC:
		s.CurrOpcode = fmt.Sprintf("CP A,H")
		s.CPAR8("H")

	case 0xBD:
		s.CurrOpcode = fmt.Sprintf("CP A,L")
		s.CPAR8("L")

	case 0xBE:
		s.CurrOpcode = fmt.Sprintf("CP A,(HL)")
		s.CPAHL()

	case 0xBF:
		s.CurrOpcode = fmt.Sprintf("CP A,A")
		s.CPAR8("A")

	case 0xC0:
		s.CurrOpcode = fmt.Sprintf("RET NZ")
		s.RETCC("NZ")

	case 0xC1:
		s.CurrOpcode = fmt.Sprintf("POP BC")
		s.POPR16("BC")

	case 0xC2:
		addr := int(s.Mem.Read(s.GetReg16Val("PC")+1)) + (int(s.Mem.Read(s.GetReg16Val("PC")+2)) << 8)
		s.CurrOpcode = fmt.Sprintf("JP NZ,$%x ", addr)
		s.JPCCN16("NZ", addr)

	case 0xC3:
		addr := int(s.Mem.Read(s.GetReg16Val("PC")+1)) + (int(s.Mem.Read(s.GetReg16Val("PC")+2)) << 8)
		s.CurrOpcode = fmt.Sprintf("JP $%x ", addr)
		s.JPN16(addr)

	case 0xC4:
		addr := int(uint16(s.Mem.Read(s.GetReg16Val("PC")+1)) + (uint16(s.Mem.Read(s.GetReg16Val("PC")+2)) << 8))
		s.CurrOpcode = fmt.Sprintf("CALL NZ,$%x ", addr)
		s.CALLCCN16("NZ", addr)

	case 0xC5:
		s.CurrOpcode = fmt.Sprintf("PUSH BC")
		s.PUSHR16("BC")

	case 0xC6:
		s.CurrOpcode = fmt.Sprintf("ADD A,$%x", uint8(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.ADDAN8()

	case 0xC7:
		s.CurrOpcode = fmt.Sprintf("RST 00H")
		s.RST("00H")

	case 0xC8:
		s.CurrOpcode = fmt.Sprintf("RET Z")
		s.RETCC("Z")

	case 0xC9:
		s.CurrOpcode = fmt.Sprintf("RET")
		s.RET()

	case 0xCA:
		addr := int(s.Mem.Read(s.GetReg16Val("PC")+1)) + (int(s.Mem.Read(s.GetReg16Val("PC")+2)) << 8)
		s.JPCCN16("Z", addr)
		s.CurrOpcode = fmt.Sprintf("JP Z,$%x", addr)

	case 0xCB:
		s.PrefixTable()

	case 0xCC:
		addr := int(s.Mem.Read(s.GetReg16Val("PC")+1)) + (int(s.Mem.Read(s.GetReg16Val("PC")+2)) << 8)
		s.CurrOpcode = fmt.Sprintf("CALL Z,$%x", addr)
		s.CALLCCN16("Z", addr)

	case 0xCD:
		addr := int(s.Mem.Read(s.GetReg16Val("PC")+1)) + (int(s.Mem.Read(s.GetReg16Val("PC")+2)) << 8)
		s.CurrOpcode = fmt.Sprintf("CALL $%x", addr)
		s.CALLN16(addr)

	case 0xCE:
		s.CurrOpcode = fmt.Sprintf("ADC A,$%x", int(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.ADCAN8()

	case 0xCF:
		s.CurrOpcode = fmt.Sprintf("RST 08H")
		s.RST("08H")

	case 0xD0:
		s.CurrOpcode = fmt.Sprintf("RET NC")
		s.RETCC("NC")

	case 0xD1:
		s.CurrOpcode = fmt.Sprintf("POP DE")
		s.POPR16("DE")

	case 0xD2:
		addr := int(s.Mem.Read(s.GetReg16Val("PC")+1)) + (int(s.Mem.Read(s.GetReg16Val("PC")+2)) << 8)
		s.CurrOpcode = fmt.Sprintf("JP NC,$%x ", addr)
		s.JPCCN16("NC", addr)

	case 0xD4:
		addr := int(s.Mem.Read(s.GetReg16Val("PC")+1)) + (int(s.Mem.Read(s.GetReg16Val("PC")+2)) << 8)
		s.CurrOpcode = fmt.Sprintf("CALL NC,$%x ", addr)
		s.CALLCCN16("NC", addr)

	case 0xD5:
		s.CurrOpcode = fmt.Sprintf("PUSH DE")
		s.PUSHR16("DE")

	case 0xD6:
		s.CurrOpcode = fmt.Sprintf("SUB A,$%x", int(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.SUBAN8()

	case 0xD7:
		s.CurrOpcode = fmt.Sprintf("RST 10H")
		s.RST("10H")

	case 0xD8:
		s.CurrOpcode = fmt.Sprintf("RET C")
		s.RETCC("C")

	case 0xD9:
		s.CurrOpcode = fmt.Sprintf("RETI")
		s.RETI()

	case 0xDA:
		addr := int(s.Mem.Read(s.GetReg16Val("PC")+1)) + (int(s.Mem.Read(s.GetReg16Val("PC")+2)) << 8)
		s.CurrOpcode = fmt.Sprintf("JP C,$%x", addr)
		s.JPCCN16("C", addr)

	case 0xDC:
		addr := int(s.Mem.Read(s.GetReg16Val("PC")+1)) + (int(s.Mem.Read(s.GetReg16Val("PC")+2)) << 8)
		s.CurrOpcode = fmt.Sprintf("CALL C,$%x", addr)
		s.CALLCCN16("C", addr)

	case 0xDE:
		s.CurrOpcode = fmt.Sprintf("SBC A,$%x", uint8(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.SBCAN8()

	case 0xDF:
		s.CurrOpcode = fmt.Sprintf("RST 18H")
		s.RST("18H")

	case 0xE0:
		addr := int(uint16(0xFF00 + int(s.Mem.Read(s.GetReg16Val("PC")+1))))
		s.CurrOpcode = fmt.Sprintf("LD $(%x),A", addr)
		s.LDN16A(addr)
		s.SetReg16Val("PC", s.GetReg16Val("PC")-1)

	case 0xE1:
		s.CurrOpcode = fmt.Sprintf("POP HL")
		s.POPR16("HL")

	case 0xE2:
		s.CurrOpcode = fmt.Sprintf("LD ($FF00+C),A")
		s.LDN16A(int(uint16(0xFF00 + int(s.GetReg8Val("C")))))
		s.SetReg16Val("PC", s.GetReg16Val("PC")-2)

	case 0xE5:
		s.CurrOpcode = fmt.Sprintf("PUSH HL")
		s.PUSHR16("HL")

	case 0xE6:
		s.CurrOpcode = fmt.Sprintf("AND A,$%x", uint8(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.ANDAN8()

	case 0xE7:
		s.RST("20H")
		s.CurrOpcode = fmt.Sprintf("RST 20H")

	case 0xE8:
		s.ADDSPE8()
		s.CurrOpcode = fmt.Sprintf("ADD SP,%x", int8(s.Mem.Read(s.GetReg16Val("PC")+1)))

	case 0xE9:
		s.CurrOpcode = fmt.Sprintf("JP HL")
		s.JPHL()

	case 0xEA:
		addr := int(uint16(int(s.Mem.Read(s.GetReg16Val("PC")+1)) + (int(s.Mem.Read(s.GetReg16Val("PC")+2)) << 8)))
		s.CurrOpcode = fmt.Sprintf("LD ($%x),A", addr)
		s.LDN16A(addr)

	case 0xEE:
		s.CurrOpcode = fmt.Sprintf("XOR $%x", uint8(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.XORAN8()

	case 0xEF:
		s.CurrOpcode = fmt.Sprintf("RST 28H")
		s.RST("28H")

	case 0xF0:
		addr := int(uint16(0xFF00 + int(uint8(s.Mem.Read(s.GetReg16Val("PC")+1)))))
		s.CurrOpcode = fmt.Sprintf("LD A,($%x)", addr)
		s.LDAN16(addr)

	case 0xF1:
		s.CurrOpcode = fmt.Sprintf("POP AF")
		s.POPFA()

	case 0xF2:
		s.CurrOpcode = fmt.Sprintf("LD A,(0xFF00+C)")
		s.LDAN16(int(uint16(0xFF00 + s.GetReg8Val("C"))))
		s.SetReg16Val("PC", s.GetReg16Val("PC")-1)

	case 0xF3:
		s.CurrOpcode = fmt.Sprintf("DI")
		s.DI()

	case 0xF4:

	case 0xF5:
		s.CurrOpcode = fmt.Sprintf("PUSH AF")
		s.PUSHFA()

	case 0xF6:
		s.CurrOpcode = fmt.Sprintf("OR A,$%x", int8(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.ORAN8()

	case 0xF7:
		s.CurrOpcode = fmt.Sprintf("RST 30H")
		s.RST("30H")

	case 0xF8:
		s.CurrOpcode = fmt.Sprintf("LD HL,SP+%x", int8(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.LDHLSPE8()

	case 0xF9:
		s.CurrOpcode = fmt.Sprintf("LD SP,HL")
		s.LDSPHL()

	case 0xFA:
		addr := int(uint16(int(s.Mem.Read(s.GetReg16Val("PC")+1)) + int(s.Mem.Read(s.GetReg16Val("PC")+2))<<8))
		s.CurrOpcode = fmt.Sprintf("LD A,($%x)", addr)
		s.LDAN16(addr)
		s.SetReg16Val("PC", s.GetReg16Val("PC")+1)

	case 0xFB:
		s.CurrOpcode = fmt.Sprintf("EI")
		s.EI()

	case 0xFE:
		s.CurrOpcode = fmt.Sprintf("CP $%x", uint8(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.CPAN8()

	case 0xFF:
		s.CurrOpcode = fmt.Sprintf("RST 38H")
		s.RST("38H")

	}
}
