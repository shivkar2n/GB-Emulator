package CPU

import (
	"log"
)

func (s *CPU) ExecuteOpcode(stateLog *log.Logger, log *log.Logger) {
	s.StateInfo(stateLog) // Log CPU State info

	switch s.Mem.Read(s.GetReg16Val("PC")) {
	case 0x00:
		log.Printf("NOP\n")
		s.NOP()

	case 0x01:
		log.Printf("LD BC,$%x \n", uint16(int(s.Mem.Read(s.GetReg16Val("PC")+1))+(int(s.Mem.Read(s.GetReg16Val("PC")+2))<<8)))
		s.LDR16N16("BC")

	case 0x02:
		log.Printf("LD (BC), A\n")
		s.LDR16A("BC")

	case 0x03:
		log.Printf("INC BC\n")
		s.INCR16("BC")

	case 0x04:
		log.Printf("INC B\n")
		s.INCR8("B")

	case 0x05:
		log.Printf("DEC B\n")
		s.DECR8("B")

	case 0x06:
		log.Printf("LD B,$%x\n", uint8(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.LDR8N8("B")

	case 0x07:
		log.Printf("RLCA\n")
		s.RLCA()

	case 0x08:
		log.Printf("LD (%x), SP\n", int(s.Mem.Read(s.GetReg16Val("PC")+1))+(int(s.Mem.Read(s.GetReg16Val("PC")+2))<<8))
		s.LDN16SP()

	case 0x09:
		log.Printf("ADD HL, BC\n")
		s.ADDHLR16("BC")

	case 0x0A:
		log.Printf("LD A, (BC)\n")
		s.LDAR16("BC")

	case 0x0B:
		log.Printf("DEC BC\n")
		s.DECR16("BC")

	case 0x0C:
		log.Printf("INC C\n")
		s.INCR8("C")

	case 0x0D:
		log.Printf("DEC C\n")
		s.DECR8("C")

	case 0x0E:
		log.Printf("LD C, $%x\n", int(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.LDR8N8("C")

	case 0x0F:
		log.Printf("RRCA \n")
		s.RRCA()

	case 0x10:
		log.Printf("STOP $%x\n", int8(s.Mem.Read(s.GetReg16Val("PC")+1)))

	case 0x11:
		log.Printf("LD DE,$%x\n", uint16(int(s.Mem.Read(s.GetReg16Val("PC")+1))+int(s.Mem.Read(s.GetReg16Val("PC")+2)<<8)))
		s.LDR16N16("DE")

	case 0x12:
		log.Printf("LD (DE),A\n")
		s.LDR16A("DE")

	case 0x13:
		log.Printf("INC (DE)\n")
		s.INCR16("DE")

	case 0x14:
		log.Printf("INC D\n")
		s.INCR8("D")

	case 0x15:
		log.Printf("DEC D\n")
		s.DECR8("D")

	case 0x16:
		log.Printf("LD D,$%x\n", uint8(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.LDR8N8("D")

	case 0x17:
		log.Printf("RLA\n")
		s.RLA()

	case 0x18:
		log.Printf("JR $%x\n", int8(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.JRN16()

	case 0x19:
		log.Printf("ADD HL, DE\n")
		s.ADDHLR16("DE")

	case 0x1A:
		log.Printf("LD A, (DE)\n")
		s.LDAR16("DE")

	case 0x1B:
		log.Printf("DEC DE\n")
		s.DECR16("DE")

	case 0x1C:
		log.Printf("INC E\n")
		s.INCR8("E")

	case 0x1D:
		log.Printf("DEC E\n")
		s.DECR8("E")

	case 0x1E:
		log.Printf("LD E,$%x\n", uint8(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.LDR8N8("E")

	case 0x1F:
		log.Printf("RRA\n")
		s.RRA()

	case 0x20:
		log.Printf("JR NZ,$%x\n", int8(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.JRCCN16("NZ")

	case 0x21:
		log.Printf("LD HL,$%x\n", uint16(int(s.Mem.Read(s.GetReg16Val("PC")+1))+(int(s.Mem.Read(s.GetReg16Val("PC")+2))<<8)))
		s.LDR16N16("HL")

	case 0x22:
		log.Printf("LD (HL+),A\n")
		s.LDHLIA()

	case 0x23:
		log.Printf("INC HL\n")
		s.INCR16("HL")

	case 0x24:
		log.Printf("INC H\n")
		s.INCR8("H")

	case 0x25:
		log.Printf("DEC H\n")
		s.DECR8("H")

	case 0x26:
		log.Printf("LD H,$%x\n", uint8(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.LDR8N8("H")

	case 0x27:
		log.Printf("DAA\n")
		s.DAA()

	case 0x28:
		log.Printf("JR Z,$%x\n", int8(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.JRCCN16("Z")

	case 0x29:
		log.Printf("ADD HL,HL\n")
		s.ADDHLR16("HL")

	case 0x2A:
		log.Printf("LD A,(HL+)\n")
		s.LDAHLI()

	case 0x2B:
		log.Printf("DEC HL\n")
		s.DECR16("HL")

	case 0x2C:
		log.Printf("INC L\n")
		s.INCR8("L")

	case 0x2D:
		log.Printf("DEC L\n")
		s.DECR8("L")

	case 0x2E:
		log.Printf("LD L,$%x\n", uint8(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.LDR8N8("L")

	case 0x2F:
		log.Printf("CPL\n")
		s.CPL()

	case 0x30:
		log.Printf("JR NC,$%x\n", int8(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.JRCCN16("NC")

	case 0x31:
		log.Printf("LD SP $%x\n", uint16(int(s.Mem.Read(s.GetReg16Val("PC")+1))+(int(s.Mem.Read(s.GetReg16Val("PC")+2))<<8)))
		s.LDR16N16("SP")

	case 0x32:
		log.Printf("LD (HL-),A\n")
		s.LDHLDA()

	case 0x33:
		log.Printf("INC SP\n")
		s.INCR16("SP")

	case 0x34:
		log.Printf("INC (HL)\n")
		s.INCRHL()

	case 0x35:
		log.Printf("DEC (HL)\n")
		s.DECHL()

	case 0x36:
		log.Printf("LD (HL),$%x\n", uint8(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.LDHLN8()

	case 0x37:
		log.Printf("SCF\n")
		s.SCF()

	case 0x38:
		log.Printf("JR C,$%x\n", int8(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.JRCCN16("C")

	case 0x39:
		log.Printf("ADD HL,SP\n")
		s.ADDHLSP()

	case 0x3A:
		log.Printf("LD A,(HL-)\n")
		s.LDAHLD()

	case 0x3B:
		log.Printf("DEC SP\n")
		s.DECR16("SP")

	case 0x3C:
		log.Printf("INC A\n")
		s.INCR8("A")

	case 0x3D:
		log.Printf("DEC A\n")
		s.DECR8("A")

	case 0x3E:
		log.Printf("LD A,$%x\n", uint8(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.LDR8N8("A")

	case 0x3F:
		log.Printf("CCF\n")
		s.CCF()

	case 0x40:
		log.Printf("LD B,B\n")
		s.LDR8R8("B", "B")

	case 0x41:
		log.Printf("LD B,C\n")
		s.LDR8R8("B", "C")

	case 0x42:
		log.Printf("LD B,D\n")
		s.LDR8R8("B", "D")

	case 0x43:
		log.Printf("LD B,E\n")
		s.LDR8R8("B", "E")

	case 0x44:
		log.Printf("LD B,H\n")
		s.LDR8R8("B", "H")

	case 0x45:
		log.Printf("LD B,L\n")
		s.LDR8R8("B", "L")

	case 0x46:
		log.Printf("LD B,(HL)\n")
		s.LDR8HL("B")

	case 0x47:
		log.Printf("LD B,A\n")
		s.LDR8R8("B", "A")

	case 0x48:
		log.Printf("LD C,B\n")
		s.LDR8R8("C", "B")

	case 0x49:
		log.Printf("LD C,C\n")
		s.LDR8R8("C", "C")

	case 0x4A:
		log.Printf("LD C,D\n")
		s.LDR8R8("C", "D")

	case 0x4B:
		log.Printf("LD C,E\n")
		s.LDR8R8("C", "E")

	case 0x4C:
		log.Printf("LD C,H\n")
		s.LDR8R8("C", "H")

	case 0x4D:
		log.Printf("LD C,L\n")
		s.LDR8R8("C", "L")

	case 0x4E:
		log.Printf("LD C,(HL)\n")
		s.LDR8HL("C")

	case 0x4F:
		log.Printf("LD C,A\n")
		s.LDR8R8("C", "A")

	case 0x50:
		log.Printf("LD D,B\n")
		s.LDR8R8("D", "B")

	case 0x51:
		log.Printf("LD D,C\n")
		s.LDR8R8("D", "C")

	case 0x52:
		log.Printf("LD D,D\n")
		s.LDR8R8("D", "D")

	case 0x53:
		log.Printf("LD D,E\n")
		s.LDR8R8("D", "E")

	case 0x54:
		log.Printf("LD D,H\n")
		s.LDR8R8("D", "H")

	case 0x55:
		log.Printf("LD D,L\n")
		s.LDR8R8("D", "L")

	case 0x56:
		log.Printf("LD D,(HL)\n")
		s.LDR8HL("D")

	case 0x57:
		log.Printf("LD D,A\n")
		s.LDR8R8("D", "A")

	case 0x58:
		log.Printf("LD E,B\n")
		s.LDR8R8("E", "B")

	case 0x59:
		log.Printf("LD E,C\n")
		s.LDR8R8("E", "C")

	case 0x5A:
		log.Printf("LD E,D\n")
		s.LDR8R8("E", "D")

	case 0x5B:
		log.Printf("LD E,E\n")
		s.LDR8R8("E", "E")

	case 0x5C:
		log.Printf("LD E,H\n")
		s.LDR8R8("E", "H")

	case 0x5D:
		log.Printf("LD E,L\n")
		s.LDR8R8("E", "L")

	case 0x5E:
		log.Printf("LD E,(HL)\n")
		s.LDR8HL("E")

	case 0x5F:
		log.Printf("LD E,A\n")
		s.LDR8R8("E", "A")

	case 0x60:
		log.Printf("LD H,B\n")
		s.LDR8R8("H", "B")

	case 0x61:
		log.Printf("LD H,C\n")
		s.LDR8R8("H", "C")

	case 0x62:
		log.Printf("LD H,D\n")
		s.LDR8R8("H", "D")

	case 0x63:
		log.Printf("LD H,E\n")
		s.LDR8R8("H", "E")

	case 0x64:
		log.Printf("LD H,H\n")
		s.LDR8R8("H", "H")

	case 0x65:
		log.Printf("LD H,L\n")
		s.LDR8R8("H", "L")

	case 0x66:
		log.Printf("LD H,(HL)\n")
		s.LDR8HL("H")

	case 0x67:
		log.Printf("LD H,A\n")
		s.LDR8R8("H", "A")

	case 0x68:
		log.Printf("LD L,B\n")
		s.LDR8R8("L", "B")

	case 0x69:
		log.Printf("LD L,C\n")
		s.LDR8R8("L", "C")

	case 0x6A:
		log.Printf("LD L,D\n")
		s.LDR8R8("L", "D")

	case 0x6B:
		log.Printf("LD L,E\n")
		s.LDR8R8("L", "E")

	case 0x6C:
		log.Printf("LD L,H\n")
		s.LDR8R8("L", "H")

	case 0x6D:
		log.Printf("LD L,L\n")
		s.LDR8R8("L", "L")

	case 0x6E:
		log.Printf("LD L,(HL)\n")
		s.LDR8HL("L")

	case 0x6F:
		log.Printf("LD L,A\n")
		s.LDR8R8("L", "A")

	case 0x70:
		log.Printf("LD (HL),B\n")
		s.LDHLR8("B")

	case 0x71:
		log.Printf("LD (HL),C\n")
		s.LDHLR8("C")

	case 0x72:
		log.Printf("LD (HL),D\n")
		s.LDHLR8("D")

	case 0x73:
		log.Printf("LD (HL),E\n")
		s.LDHLR8("E")

	case 0x74:
		log.Printf("LD (HL),H\n")
		s.LDHLR8("H")

	case 0x75:
		log.Printf("LD (HL),L\n")
		s.LDHLR8("L")

	case 0x76:
		log.Printf("HALT\n")
		s.HALT()

	case 0x77:
		log.Printf("LD (HL),A\n")
		s.LDHLR8("A")

	case 0x78:
		log.Printf("LD A,B\n")
		s.LDR8R8("A", "B")

	case 0x79:
		log.Printf("LD A,C\n")
		s.LDR8R8("A", "C")

	case 0x7A:
		log.Printf("LD A,D\n")
		s.LDR8R8("A", "D")

	case 0x7B:
		log.Printf("LD A,E\n")
		s.LDR8R8("A", "E")

	case 0x7C:
		log.Printf("LD A,H\n")
		s.LDR8R8("A", "H")

	case 0x7D:
		log.Printf("LD A,L\n")
		s.LDR8R8("A", "L")

	case 0x7E:
		log.Printf("LD A,(HL)\n")
		s.LDR8HL("A")

	case 0x7F:
		log.Printf("LD A,A\n")
		s.LDR8R8("A", "A")

	case 0x80:
		log.Printf("ADD A,B\n")
		s.ADDAR8("B")

	case 0x81:
		log.Printf("ADD A,C\n")
		s.ADDAR8("C")

	case 0x82:
		log.Printf("ADD A,D\n")
		s.ADDAR8("D")

	case 0x83:
		s.ADDAR8("E")
		log.Printf("ADD A,E\n")

	case 0x84:
		log.Printf("ADD A,H\n")
		s.ADDAR8("H")

	case 0x85:
		log.Printf("ADD A,L\n")
		s.ADDAR8("L")

	case 0x86:
		log.Printf("ADD A,(HL)\n")
		s.ADDAHL()

	case 0x87:
		log.Printf("ADD A,A\n")
		s.ADDAR8("A")

	case 0x88:
		log.Printf("ADC A,B\n")
		s.ADCAR8("B")

	case 0x89:
		log.Printf("ADC A,C\n")
		s.ADCAR8("C")

	case 0x8A:
		log.Printf("ADC A,D\n")
		s.ADCAR8("D")

	case 0x8B:
		log.Printf("ADC A,E\n")
		s.ADCAR8("E")

	case 0x8C:
		log.Printf("ADC A,H\n")
		s.ADCAR8("H")

	case 0x8D:
		log.Printf("ADC A,L\n")
		s.ADCAR8("L")

	case 0x8E:
		log.Printf("ADC A,(HL)\n")
		s.ADCAHL()

	case 0x8F:
		log.Printf("ADC A,A\n")
		s.ADCAR8("A")

	case 0x90:
		log.Printf("SUB A,B\n")
		s.SUBAR8("B")

	case 0x91:
		log.Printf("SUB A,C\n")
		s.SUBAR8("C")

	case 0x92:
		log.Printf("SUB A,D\n")
		s.SUBAR8("D")

	case 0x93:
		log.Printf("SUB A,E\n")
		s.SUBAR8("E")

	case 0x94:
		log.Printf("SUB A,H\n")
		s.SUBAR8("H")

	case 0x95:
		log.Printf("SUB A,L\n")
		s.SUBAR8("L")

	case 0x96:
		log.Printf("SUB A,(HL)\n")
		s.SUBAHL()

	case 0x97:
		log.Printf("SUB A,A\n")
		s.SUBAR8("A")

	case 0x98:
		log.Printf("SBC A,B\n")
		s.SBCAR8("B")

	case 0x99:
		log.Printf("SBC A,C\n")
		s.SBCAR8("C")

	case 0x9A:
		log.Printf("SBC A,D\n")
		s.SBCAR8("D")

	case 0x9B:
		log.Printf("SBC A,E\n")
		s.SBCAR8("E")

	case 0x9C:
		log.Printf("SBC A,H\n")
		s.SBCAR8("H")

	case 0x9D:
		log.Printf("SBC A,L\n")
		s.SBCAR8("L")

	case 0x9E:
		log.Printf("SBC A,(HL)\n")
		s.SBCAHL()

	case 0x9F:
		log.Printf("SBC A,A\n")
		s.SBCAR8("A")

	case 0xA0:
		log.Printf("AND A,B\n")
		s.ANDAR8("B")

	case 0xA1:
		log.Printf("AND A,C\n")
		s.ANDAR8("C")

	case 0xA2:
		log.Printf("AND A,D\n")
		s.ANDAR8("D")

	case 0xA3:
		log.Printf("AND A,E\n")
		s.ANDAR8("E")

	case 0xA4:
		log.Printf("AND A,H\n")
		s.ANDAR8("H")

	case 0xA5:
		log.Printf("AND A,L\n")
		s.ANDAR8("L")

	case 0xA6:
		log.Printf("AND A,(HL)\n")
		s.ANDAHL()

	case 0xA7:
		log.Printf("AND A,A\n")
		s.ANDAR8("A")

	case 0xA8:
		log.Printf("XOR A,B\n")
		s.XORAR8("B")

	case 0xA9:
		log.Printf("XOR A,C\n")
		s.XORAR8("C")

	case 0xAA:
		log.Printf("XOR A,D\n")
		s.XORAR8("D")

	case 0xAB:
		log.Printf("XOR A,E\n")
		s.XORAR8("E")

	case 0xAC:
		log.Printf("XOR A,H\n")
		s.XORAR8("H")

	case 0xAD:
		log.Printf("XOR A,L\n")
		s.XORAR8("L")

	case 0xAE:
		log.Printf("XOR A,(HL)\n")
		s.XORAHL()

	case 0xAF:
		log.Printf("XOR A,A\n")
		s.XORAR8("A")

	case 0xB0:
		log.Printf("OR A,B\n")
		s.ORAR8("B")

	case 0xB1:
		log.Printf("OR A,C\n")
		s.ORAR8("C")

	case 0xB2:
		log.Printf("OR A,D\n")
		s.ORAR8("D")

	case 0xB3:
		log.Printf("OR A,E\n")
		s.ORAR8("E")

	case 0xB4:
		log.Printf("OR A,H\n")
		s.ORAR8("H")

	case 0xB5:
		log.Printf("OR A,L\n")
		s.ORAR8("L")

	case 0xB6:
		log.Printf("OR A,(HL)\n")
		s.ORAHL()

	case 0xB7:
		log.Printf("OR A,A\n")
		s.ORAR8("A")

	case 0xB8:
		log.Printf("CP A,B\n")
		s.CPAR8("B")

	case 0xB9:
		log.Printf("CP A,C\n")
		s.CPAR8("C")

	case 0xBA:
		log.Printf("CP A,D\n")
		s.CPAR8("D")

	case 0xBB:
		log.Printf("CP A,E\n")
		s.CPAR8("E")

	case 0xBC:
		log.Printf("CP A,H\n")
		s.CPAR8("H")

	case 0xBD:
		log.Printf("CP A,L\n")
		s.CPAR8("L")

	case 0xBE:
		log.Printf("CP A,(HL)\n")
		s.CPAHL()

	case 0xBF:
		log.Printf("CP A,A\n")
		s.CPAR8("A")

	case 0xC0:
		log.Printf("RET NZ\n")
		s.RETCC("NZ")

	case 0xC1:
		log.Printf("POP BC\n")
		s.POPR16("BC")

	case 0xC2:
		addr := int(s.Mem.Read(s.GetReg16Val("PC")+1)) + (int(s.Mem.Read(s.GetReg16Val("PC")+2)) << 8)
		log.Printf("JP NZ,$%x \n", addr)
		s.JPCCN16("NZ", addr)

	case 0xC3:
		addr := int(s.Mem.Read(s.GetReg16Val("PC")+1)) + (int(s.Mem.Read(s.GetReg16Val("PC")+2)) << 8)
		log.Printf("JP $%x \n", addr)
		s.JPN16(addr)

	case 0xC4:
		addr := int(uint16(s.Mem.Read(s.GetReg16Val("PC")+1)) + (uint16(s.Mem.Read(s.GetReg16Val("PC")+2)) << 8))
		log.Printf("CALL NZ,$%x \n", addr)
		s.CALLCCN16("NZ", addr)

	case 0xC5:
		log.Printf("PUSH BC\n")
		s.PUSHR16("BC")

	case 0xC6:
		log.Printf("ADD A,$%x\n", uint8(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.ADDAN8()

	case 0xC7:
		log.Printf("RST 00H\n")
		s.RST("00H")

	case 0xC8:
		log.Printf("RET Z\n")
		s.RETCC("Z")

	case 0xC9:
		log.Printf("RET\n")
		s.RET()

	case 0xCA:
		addr := int(s.Mem.Read(s.GetReg16Val("PC")+1)) + (int(s.Mem.Read(s.GetReg16Val("PC")+2)) << 8)
		s.JPCCN16("Z", addr)
		log.Printf("JP Z,$%x\n", addr)

	case 0xCB:
		s.PrefixTable(log)

	case 0xCC:
		addr := int(s.Mem.Read(s.GetReg16Val("PC")+1)) + (int(s.Mem.Read(s.GetReg16Val("PC")+2)) << 8)
		log.Printf("CALL Z,$%x\n", addr)
		s.CALLCCN16("Z", addr)

	case 0xCD:
		addr := int(s.Mem.Read(s.GetReg16Val("PC")+1)) + (int(s.Mem.Read(s.GetReg16Val("PC")+2)) << 8)
		log.Printf("CALL $%x\n", addr)
		s.CALLN16(addr)

	case 0xCE:
		log.Printf("ADC A,$%x\n", int(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.ADCAN8()

	case 0xCF:
		log.Printf("RST 08H\n")
		s.RST("08H")

	case 0xD0:
		log.Printf("RET NC\n")
		s.RETCC("NC")

	case 0xD1:
		log.Printf("POP DE\n")
		s.POPR16("DE")

	case 0xD2:
		addr := int(s.Mem.Read(s.GetReg16Val("PC")+1)) + (int(s.Mem.Read(s.GetReg16Val("PC")+2)) << 8)
		log.Printf("JP NC,$%x \n", addr)
		s.JPCCN16("NC", addr)

	case 0xD4:
		addr := int(s.Mem.Read(s.GetReg16Val("PC")+1)) + (int(s.Mem.Read(s.GetReg16Val("PC")+2)) << 8)
		log.Printf("CALL NC,$%x \n", addr)
		s.CALLCCN16("NC", addr)

	case 0xD5:
		log.Printf("PUSH DE\n")
		s.PUSHR16("DE")

	case 0xD6:
		log.Printf("SUB A,$%x\n", int(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.SUBAN8()

	case 0xD7:
		log.Printf("RST 10H\n")
		s.RST("10H")

	case 0xD8:
		log.Printf("RET C\n")
		s.RETCC("C")

	case 0xD9:
		log.Printf("RETI\n")
		s.RETI()

	case 0xDA:
		addr := int(s.Mem.Read(s.GetReg16Val("PC")+1)) + (int(s.Mem.Read(s.GetReg16Val("PC")+2)) << 8)
		log.Printf("JP C,$%x\n", addr)
		s.JPCCN16("C", addr)

	case 0xDC:
		addr := int(s.Mem.Read(s.GetReg16Val("PC")+1)) + (int(s.Mem.Read(s.GetReg16Val("PC")+2)) << 8)
		log.Printf("CALL C,$%x\n", addr)
		s.CALLCCN16("C", addr)

	case 0xDE:
		log.Printf("SBC A,$%x\n", uint8(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.SBCAN8()

	case 0xDF:
		log.Printf("RST 18H\n")
		s.RST("18H")

	case 0xE0:
		addr := int(uint16(0xFF00 + int(s.Mem.Read(s.GetReg16Val("PC")+1))))
		log.Printf("LD $(%x),A\n", addr)
		s.LDN16A(addr)
		s.SetReg16Val("PC", s.GetReg16Val("PC")-1)

	case 0xE1:
		log.Printf("POP HL\n")
		s.POPR16("HL")

	case 0xE2:
		log.Printf("LD ($FF00+C),A\n")
		s.LDN16A(int(uint16(0xFF00 + int(s.GetReg8Val("C")))))
		s.SetReg16Val("PC", s.GetReg16Val("PC")-2)

	case 0xE5:
		log.Printf("PUSH HL\n")
		s.PUSHR16("HL")

	case 0xE6:
		log.Printf("AND A,$%x\n", uint8(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.ANDAN8()

	case 0xE7:
		s.RST("20H")
		log.Printf("RST 20H\n")

	case 0xE8:
		s.ADDSPE8()
		log.Printf("ADD SP,%x\n", int8(s.Mem.Read(s.GetReg16Val("PC")+1)))

	case 0xE9:
		log.Printf("JP HL\n")
		s.JPHL()

	case 0xEA:
		addr := int(uint16(int(s.Mem.Read(s.GetReg16Val("PC")+1)) + (int(s.Mem.Read(s.GetReg16Val("PC")+2)) << 8)))
		log.Printf("LD ($%x),A\n", addr)
		s.LDN16A(addr)

	case 0xEE:
		log.Printf("XOR $%x\n", uint8(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.XORAN8()

	case 0xEF:
		log.Printf("RST 28H\n")
		s.RST("28H")

	case 0xF0:
		addr := int(uint16(0xFF00 + int(uint8(s.Mem.Read(s.GetReg16Val("PC")+1)))))
		log.Printf("LD A,($%x)\n", addr)
		s.LDAN16(addr)

	case 0xF1:
		log.Printf("POP AF\n")
		s.POPFA()

	case 0xF2:
		log.Printf("LD A,(0xFF00+C)\n")
		s.LDAN16(int(uint16(0xFF00 + s.GetReg8Val("C"))))
		s.SetReg16Val("PC", s.GetReg16Val("PC")-1)

	case 0xF3:
		log.Printf("DI\n")
		s.DI()

	case 0xF4:

	case 0xF5:
		log.Printf("PUSH AF\n")
		s.PUSHFA()

	case 0xF6:
		log.Printf("OR A,$%x\n", int8(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.ORAN8()

	case 0xF7:
		log.Printf("RST 30H\n")
		s.RST("30H")

	case 0xF8:
		log.Printf("LD HL,SP+%x\n", int8(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.LDHLSPE8()

	case 0xF9:
		log.Printf("LD SP,HL\n")
		s.LDSPHL()

	case 0xFA:
		addr := int(uint16(int(s.Mem.Read(s.GetReg16Val("PC")+1)) + int(s.Mem.Read(s.GetReg16Val("PC")+2))<<8))
		log.Printf("LD A,($%x)\n", addr)
		s.LDAN16(addr)
		s.SetReg16Val("PC", s.GetReg16Val("PC")+1)

	case 0xFB:
		log.Printf("EI\n")
		s.EI()

	case 0xFE:
		log.Printf("CP $%x\n", int8(s.Mem.Read(s.GetReg16Val("PC")+1)))
		s.CPAN8()

	case 0xFF:
		log.Printf("RST 38H\n")
		s.RST("38H")

	}
}
