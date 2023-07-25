package Disassembler

import (
	CPU "github.com/shivkar2n/GB-Emulator/CPU"
	Logger "github.com/shivkar2n/GB-Emulator/Logger"
)

func OpCodeTable(c *CPU.CPU) {
	// Load rom into memory
	c.LoadROM()
	var stopSwitch bool
	stopSwitch = false

	// Logging
	log := Logger.DebugLogger()
	stateLog := Logger.StateLogger()

	// Main execution loop
	for c.GetReg16Val("PC") < 0xFFFF && !stopSwitch {
		// For handling interrupt service requests
		c.InterruptHandler()

		// For incrementing timer
		c.INCRTIMER()

		// Get CPU State info
		c.StateInfo(stateLog)

		// if c.GetReg16Val("PC") == 0xC321 {
		// 	c.StopExec = true
		// 	stopSwitch = true
		// 	// c.MemInfo()
		// 	// println(c.GetIFVal())
		// 	// println(c.GetIEVal())
		// }

		if !c.StopExec {
			switch c.Mem[c.GetReg16Val("PC")] {
			case 0x00:
				log.Printf("NOP\n")
				c.NOP()

			case 0x01:
				log.Printf("LD BC,$%x \n", uint16(int(c.Mem[c.GetReg16Val("PC")+1])+(int(c.Mem[c.GetReg16Val("PC")+2])<<8)))
				c.LDR16N16("BC")

			case 0x02:
				log.Printf("LD (BC), A\n")
				c.LDR16A("BC")

			case 0x03:
				log.Printf("INC BC\n")
				c.INCR16("BC")

			case 0x04:
				log.Printf("INC B\n")
				c.INCR8("B")

			case 0x05:
				log.Printf("DEC B\n")
				c.DECR8("B")

			case 0x06:
				log.Printf("LD B,$%x\n", uint8(c.Mem[c.GetReg16Val("PC")+1]))
				c.LDR8N8("B")

			case 0x07:
				log.Printf("RLCA\n")
				c.RLCA()

			case 0x08:
				log.Printf("LD (%x), SP\n", int(c.Mem[c.GetReg16Val("PC")+1])+(int(c.Mem[c.GetReg16Val("PC")+2])<<8))
				c.LDN16SP()

			case 0x09:
				log.Printf("ADD HL, BC\n")
				c.ADDHLR16("BC")

			case 0x0A:
				log.Printf("LD A, (BC)\n")
				c.LDAR16("BC")

			case 0x0B:
				log.Printf("DEC BC\n")
				c.DECR16("BC")

			case 0x0C:
				log.Printf("INC C\n")
				c.INCR8("C")

			case 0x0D:
				log.Printf("DEC C\n")
				c.DECR8("C")

			case 0x0E:
				log.Printf("LD C, $%x\n", int(c.Mem[c.GetReg16Val("PC")+1]))
				c.LDR8N8("C")

			case 0x0F:
				log.Printf("RRCA \n")
				c.RRCA()

			case 0x10:
				log.Printf("STOP $%x\n", int8(c.Mem[c.GetReg16Val("PC")+1]))

			case 0x11:
				log.Printf("LD DE,$%x\n", uint16(int(c.Mem[c.GetReg16Val("PC")+1])+(int(c.Mem[c.GetReg16Val("PC")+2])<<8)))
				c.LDR16N16("DE")

			case 0x12:
				log.Printf("LD (DE),A\n")
				c.LDR16A("DE")

			case 0x13:
				log.Printf("INC (DE)\n")
				c.INCR16("DE")

			case 0x14:
				log.Printf("INC D\n")
				c.INCR8("D")

			case 0x15:
				log.Printf("DEC D\n")
				c.DECR8("D")

			case 0x16:
				log.Printf("LD D,$%x\n", uint8(c.Mem[c.GetReg16Val("PC")+1]))
				c.LDR8N8("D")

			case 0x17:
				log.Printf("RLA\n")
				c.RLA()

			case 0x18:
				log.Printf("JR $%x\n", int8(c.Mem[c.GetReg16Val("PC")+1]))
				c.JRN16()

			case 0x19:
				log.Printf("ADD HL, DE\n")
				c.ADDHLR16("DE")

			case 0x1A:
				log.Printf("LD A, (DE)\n")
				c.LDAR16("DE")

			case 0x1B:
				log.Printf("DEC DE\n")
				c.DECR16("DE")

			case 0x1C:
				log.Printf("INC E\n")
				c.INCR8("E")

			case 0x1D:
				log.Printf("DEC E\n")
				c.DECR8("E")

			case 0x1E:
				log.Printf("LD E,$%x\n", uint8(c.Mem[c.GetReg16Val("PC")+1]))
				c.LDR8N8("E")

			case 0x1F:
				log.Printf("RRA\n")
				c.RRA()

			case 0x20:
				log.Printf("JR NZ,$%x\n", int8(c.Mem[c.GetReg16Val("PC")+1]))
				c.JRCCN16("NZ")

			case 0x21:
				log.Printf("LD HL,$%x\n", uint16(int(c.Mem[c.GetReg16Val("PC")+1])+(int(c.Mem[c.GetReg16Val("PC")+2])<<8)))
				c.LDR16N16("HL")

			case 0x22:
				log.Printf("LD (HL+),A\n")
				c.LDHLIA()

			case 0x23:
				log.Printf("INC HL\n")
				c.INCR16("HL")

			case 0x24:
				log.Printf("INC H\n")
				c.INCR8("H")

			case 0x25:
				log.Printf("DEC H\n")
				c.DECR8("H")

			case 0x26:
				log.Printf("LD H,$%x\n", uint8(c.Mem[c.GetReg16Val("PC")+1]))
				c.LDR8N8("H")

			case 0x27:
				log.Printf("DAA\n")
				c.DAA()

			case 0x28:
				log.Printf("JR Z,$%x\n", int8(c.Mem[c.GetReg16Val("PC")+1]))
				c.JRCCN16("Z")

			case 0x29:
				log.Printf("ADD HL,HL\n")
				c.ADDHLR16("HL")

			case 0x2A:
				log.Printf("LD A,(HL+)\n")
				c.LDAHLI()

			case 0x2B:
				log.Printf("DEC HL\n")
				c.DECR16("HL")

			case 0x2C:
				log.Printf("INC L\n")
				c.INCR8("L")

			case 0x2D:
				log.Printf("DEC L\n")
				c.DECR8("L")

			case 0x2E:
				log.Printf("LD L,$%x\n", uint8(c.Mem[c.GetReg16Val("PC")+1]))
				c.LDR8N8("L")

			case 0x2F:
				log.Printf("CPL\n")
				c.CPL()

			case 0x30:
				log.Printf("JR NC,$%x\n", int8(c.Mem[c.GetReg16Val("PC")+1]))
				c.JRCCN16("NC")

			case 0x31:
				log.Printf("LD SP $%x\n", uint16(int(c.Mem[c.GetReg16Val("PC")+1])+(int(c.Mem[c.GetReg16Val("PC")+2])<<8)))
				c.LDR16N16("SP")

			case 0x32:
				log.Printf("LD (HL-),A\n")
				c.LDHLDA()

			case 0x33:
				log.Printf("INC SP\n")
				c.INCR16("SP")

			case 0x34:
				log.Printf("INC (HL)\n")
				c.INCRHL()

			case 0x35:
				log.Printf("DEC (HL)\n")
				c.DECHL()

			case 0x36:
				log.Printf("LD (HL),$%x\n", uint8(c.Mem[c.GetReg16Val("PC")+1]))
				c.LDHLN8()

			case 0x37:
				log.Printf("SCF\n")
				c.SCF()

			case 0x38:
				log.Printf("JR C,$%x\n", int8(c.Mem[c.GetReg16Val("PC")+1]))
				c.JRCCN16("C")

			case 0x39:
				log.Printf("ADD HL,SP\n")
				c.ADDHLSP()

			case 0x3A:
				log.Printf("LD A,(HL-)\n")
				c.LDAHLD()

			case 0x3B:
				log.Printf("DEC SP\n")
				c.DECR16("SP")

			case 0x3C:
				log.Printf("INC A\n")
				c.INCR8("A")

			case 0x3D:
				log.Printf("DEC A\n")
				c.DECR8("A")

			case 0x3E:
				log.Printf("LD A,$%x\n", uint8(c.Mem[c.GetReg16Val("PC")+1]))
				c.LDR8N8("A")

			case 0x3F:
				log.Printf("CCF\n")
				c.CCF()

			case 0x40:
				log.Printf("LD B,B\n")
				c.LDR8R8("B", "B")

			case 0x41:
				log.Printf("LD B,C\n")
				c.LDR8R8("B", "C")

			case 0x42:
				log.Printf("LD B,D\n")
				c.LDR8R8("B", "D")

			case 0x43:
				log.Printf("LD B,E\n")
				c.LDR8R8("B", "E")

			case 0x44:
				log.Printf("LD B,H\n")
				c.LDR8R8("B", "H")

			case 0x45:
				log.Printf("LD B,L\n")
				c.LDR8R8("B", "L")

			case 0x46:
				log.Printf("LD B,(HL)\n")
				c.LDR8HL("B")

			case 0x47:
				log.Printf("LD B,A\n")
				c.LDR8R8("B", "A")

			case 0x48:
				log.Printf("LD C,B\n")
				c.LDR8R8("C", "B")

			case 0x49:
				log.Printf("LD C,C\n")
				c.LDR8R8("C", "C")

			case 0x4A:
				log.Printf("LD C,D\n")
				c.LDR8R8("C", "D")

			case 0x4B:
				log.Printf("LD C,E\n")
				c.LDR8R8("C", "E")

			case 0x4C:
				log.Printf("LD C,H\n")
				c.LDR8R8("C", "H")

			case 0x4D:
				log.Printf("LD C,L\n")
				c.LDR8R8("C", "L")

			case 0x4E:
				log.Printf("LD C,(HL)\n")
				c.LDR8HL("C")

			case 0x4F:
				log.Printf("LD C,A\n")
				c.LDR8R8("C", "A")

			case 0x50:
				log.Printf("LD D,B\n")
				c.LDR8R8("D", "B")

			case 0x51:
				log.Printf("LD D,C\n")
				c.LDR8R8("D", "C")

			case 0x52:
				log.Printf("LD D,D\n")
				c.LDR8R8("D", "D")

			case 0x53:
				log.Printf("LD D,E\n")
				c.LDR8R8("D", "E")

			case 0x54:
				log.Printf("LD D,H\n")
				c.LDR8R8("D", "H")

			case 0x55:
				log.Printf("LD D,L\n")
				c.LDR8R8("D", "L")

			case 0x56:
				log.Printf("LD D,(HL)\n")
				c.LDR8HL("D")

			case 0x57:
				log.Printf("LD D,A\n")
				c.LDR8R8("D", "A")

			case 0x58:
				log.Printf("LD E,B\n")
				c.LDR8R8("E", "B")

			case 0x59:
				log.Printf("LD E,C\n")
				c.LDR8R8("E", "C")

			case 0x5A:
				log.Printf("LD E,D\n")
				c.LDR8R8("E", "D")

			case 0x5B:
				log.Printf("LD E,E\n")
				c.LDR8R8("E", "E")

			case 0x5C:
				log.Printf("LD E,H\n")
				c.LDR8R8("E", "H")

			case 0x5D:
				log.Printf("LD E,L\n")
				c.LDR8R8("E", "L")

			case 0x5E:
				log.Printf("LD E,(HL)\n")
				c.LDR8HL("E")

			case 0x5F:
				log.Printf("LD E,A\n")
				c.LDR8R8("E", "A")

			case 0x60:
				log.Printf("LD H,B\n")
				c.LDR8R8("H", "B")

			case 0x61:
				log.Printf("LD H,C\n")
				c.LDR8R8("H", "C")

			case 0x62:
				log.Printf("LD H,D\n")
				c.LDR8R8("H", "D")

			case 0x63:
				log.Printf("LD H,E\n")
				c.LDR8R8("H", "E")

			case 0x64:
				log.Printf("LD H,H\n")
				c.LDR8R8("H", "H")

			case 0x65:
				log.Printf("LD H,L\n")
				c.LDR8R8("H", "L")

			case 0x66:
				log.Printf("LD H,(HL)\n")
				c.LDR8HL("H")

			case 0x67:
				log.Printf("LD H,A\n")
				c.LDR8R8("H", "A")

			case 0x68:
				log.Printf("LD L,B\n")
				c.LDR8R8("L", "B")

			case 0x69:
				log.Printf("LD L,C\n")
				c.LDR8R8("L", "C")

			case 0x6A:
				log.Printf("LD L,D\n")
				c.LDR8R8("L", "D")

			case 0x6B:
				log.Printf("LD L,E\n")
				c.LDR8R8("L", "E")

			case 0x6C:
				log.Printf("LD L,H\n")
				c.LDR8R8("L", "H")

			case 0x6D:
				log.Printf("LD L,L\n")
				c.LDR8R8("L", "L")

			case 0x6E:
				log.Printf("LD L,(HL)\n")
				c.LDR8HL("L")

			case 0x6F:
				log.Printf("LD L,A\n")
				c.LDR8R8("L", "A")

			case 0x70:
				log.Printf("LD (HL),B\n")
				c.LDHLR8("B")

			case 0x71:
				log.Printf("LD (HL),C\n")
				c.LDHLR8("C")

			case 0x72:
				log.Printf("LD (HL),D\n")
				c.LDHLR8("D")

			case 0x73:
				log.Printf("LD (HL),E\n")
				c.LDHLR8("E")

			case 0x74:
				log.Printf("LD (HL),H\n")
				c.LDHLR8("H")

			case 0x75:
				log.Printf("LD (HL),L\n")
				c.LDHLR8("L")

			case 0x76:
				log.Printf("HALT\n")
				c.HALT()

			case 0x77:
				log.Printf("LD (HL),A\n")
				c.LDHLR8("A")

			case 0x78:
				log.Printf("LD A,B\n")
				c.LDR8R8("A", "B")

			case 0x79:
				log.Printf("LD A,C\n")
				c.LDR8R8("A", "C")

			case 0x7A:
				log.Printf("LD A,D\n")
				c.LDR8R8("A", "D")

			case 0x7B:
				log.Printf("LD A,E\n")
				c.LDR8R8("A", "E")

			case 0x7C:
				log.Printf("LD A,H\n")
				c.LDR8R8("A", "H")

			case 0x7D:
				log.Printf("LD A,L\n")
				c.LDR8R8("A", "L")

			case 0x7E:
				log.Printf("LD A,(HL)\n")
				c.LDR8HL("A")

			case 0x7F:
				log.Printf("LD A,A\n")
				c.LDR8R8("A", "A")

			case 0x80:
				log.Printf("ADD A,B\n")
				c.ADDAR8("B")

			case 0x81:
				log.Printf("ADD A,C\n")
				c.ADDAR8("C")

			case 0x82:
				log.Printf("ADD A,D\n")
				c.ADDAR8("D")

			case 0x83:
				c.ADDAR8("E")
				log.Printf("ADD A,E\n")

			case 0x84:
				log.Printf("ADD A,H\n")
				c.ADDAR8("H")

			case 0x85:
				log.Printf("ADD A,L\n")
				c.ADDAR8("L")

			case 0x86:
				log.Printf("ADD A,(HL)\n")
				c.ADDAHL()

			case 0x87:
				log.Printf("ADD A,A\n")
				c.ADDAR8("A")

			case 0x88:
				log.Printf("ADC A,B\n")
				c.ADCAR8("B")

			case 0x89:
				log.Printf("ADC A,C\n")
				c.ADCAR8("C")

			case 0x8A:
				log.Printf("ADC A,D\n")
				c.ADCAR8("D")

			case 0x8B:
				log.Printf("ADC A,E\n")
				c.ADCAR8("E")

			case 0x8C:
				log.Printf("ADC A,H\n")
				c.ADCAR8("H")

			case 0x8D:
				log.Printf("ADC A,L\n")
				c.ADCAR8("L")

			case 0x8E:
				log.Printf("ADC A,(HL)\n")
				c.ADCAHL()

			case 0x8F:
				log.Printf("ADC A,A\n")
				c.ADCAR8("A")

			case 0x90:
				log.Printf("SUB A,B\n")
				c.SUBAR8("B")

			case 0x91:
				log.Printf("SUB A,C\n")
				c.SUBAR8("C")

			case 0x92:
				log.Printf("SUB A,D\n")
				c.SUBAR8("D")

			case 0x93:
				log.Printf("SUB A,E\n")
				c.SUBAR8("E")

			case 0x94:
				log.Printf("SUB A,H\n")
				c.SUBAR8("H")

			case 0x95:
				log.Printf("SUB A,L\n")
				c.SUBAR8("L")

			case 0x96:
				log.Printf("SUB A,(HL)\n")
				c.SUBAHL()

			case 0x97:
				log.Printf("SUB A,A\n")
				c.SUBAR8("A")

			case 0x98:
				log.Printf("SBC A,B\n")
				c.SBCAR8("B")

			case 0x99:
				log.Printf("SBC A,C\n")
				c.SBCAR8("C")

			case 0x9A:
				log.Printf("SBC A,D\n")
				c.SBCAR8("D")

			case 0x9B:
				log.Printf("SBC A,E\n")
				c.SBCAR8("E")

			case 0x9C:
				log.Printf("SBC A,H\n")
				c.SBCAR8("H")

			case 0x9D:
				log.Printf("SBC A,L\n")
				c.SBCAR8("L")

			case 0x9E:
				log.Printf("SBC A,(HL)\n")
				c.SBCAHL()

			case 0x9F:
				log.Printf("SBC A,A\n")
				c.SBCAR8("A")

			case 0xA0:
				log.Printf("AND A,B\n")
				c.ANDAR8("B")

			case 0xA1:
				log.Printf("AND A,C\n")
				c.ANDAR8("C")

			case 0xA2:
				log.Printf("AND A,D\n")
				c.ANDAR8("D")

			case 0xA3:
				log.Printf("AND A,E\n")
				c.ANDAR8("E")

			case 0xA4:
				log.Printf("AND A,H\n")
				c.ANDAR8("H")

			case 0xA5:
				log.Printf("AND A,L\n")
				c.ANDAR8("L")

			case 0xA6:
				log.Printf("AND A,(HL)\n")
				c.ANDAHL()

			case 0xA7:
				log.Printf("AND A,A\n")
				c.ANDAR8("A")

			case 0xA8:
				log.Printf("XOR A,B\n")
				c.XORAR8("B")

			case 0xA9:
				log.Printf("XOR A,C\n")
				c.XORAR8("C")

			case 0xAA:
				log.Printf("XOR A,D\n")
				c.XORAR8("D")

			case 0xAB:
				log.Printf("XOR A,E\n")
				c.XORAR8("E")

			case 0xAC:
				log.Printf("XOR A,H\n")
				c.XORAR8("H")

			case 0xAD:
				log.Printf("XOR A,L\n")
				c.XORAR8("L")

			case 0xAE:
				log.Printf("XOR A,(HL)\n")
				c.XORAHL()

			case 0xAF:
				log.Printf("XOR A,A\n")
				c.XORAR8("A")

			case 0xB0:
				log.Printf("OR A,B\n")
				c.ORAR8("B")

			case 0xB1:
				log.Printf("OR A,C\n")
				c.ORAR8("C")

			case 0xB2:
				log.Printf("OR A,D\n")
				c.ORAR8("D")

			case 0xB3:
				log.Printf("OR A,E\n")
				c.ORAR8("E")

			case 0xB4:
				log.Printf("OR A,H\n")
				c.ORAR8("H")

			case 0xB5:
				log.Printf("OR A,L\n")
				c.ORAR8("L")

			case 0xB6:
				log.Printf("OR A,(HL)\n")
				c.ORAHL()

			case 0xB7:
				log.Printf("OR A,A\n")
				c.ORAR8("A")

			case 0xB8:
				log.Printf("CP A,B\n")
				c.CPAR8("B")

			case 0xB9:
				log.Printf("CP A,C\n")
				c.CPAR8("C")

			case 0xBA:
				log.Printf("CP A,D\n")
				c.CPAR8("D")

			case 0xBB:
				log.Printf("CP A,E\n")
				c.CPAR8("E")

			case 0xBC:
				log.Printf("CP A,H\n")
				c.CPAR8("H")

			case 0xBD:
				log.Printf("CP A,L\n")
				c.CPAR8("L")

			case 0xBE:
				log.Printf("CP A,(HL)\n")
				c.CPAHL()

			case 0xBF:
				log.Printf("CP A,A\n")
				c.CPAR8("A")

			case 0xC0:
				log.Printf("RET NZ\n")
				c.RETCC("NZ")

			case 0xC1:
				log.Printf("POP BC\n")
				c.POPR16("BC")

			case 0xC2:
				addr := int(c.Mem[c.GetReg16Val("PC")+1]) + (int(c.Mem[c.GetReg16Val("PC")+2]) << 8)
				log.Printf("JP NZ,$%x \n", addr)
				c.JPCCN16("NZ", addr)

			case 0xC3:
				addr := int(c.Mem[c.GetReg16Val("PC")+1]) + (int(c.Mem[c.GetReg16Val("PC")+2]) << 8)
				log.Printf("JP $%x \n", addr)
				c.JPN16(addr)

			case 0xC4:
				addr := int(uint16(c.Mem[c.GetReg16Val("PC")+1]) + (uint16(c.Mem[c.GetReg16Val("PC")+2]) << 8))
				log.Printf("CALL NZ,$%x \n", addr)
				c.CALLCCN16("NZ", addr)

			case 0xC5:
				log.Printf("PUSH BC\n")
				c.PUSHR16("BC")

			case 0xC6:
				log.Printf("ADD A,$%x\n", uint8(c.Mem[c.GetReg16Val("PC")+1]))
				c.ADDAN8()

			case 0xC7:
				log.Printf("RST 00H\n")
				c.RST("00H")

			case 0xC8:
				log.Printf("RET Z\n")
				c.RETCC("Z")

			case 0xC9:
				log.Printf("RET\n")
				c.RET()

			case 0xCA:
				addr := int(c.Mem[c.GetReg16Val("PC")+1]) + (int(c.Mem[c.GetReg16Val("PC")+2]) << 8)
				c.JPCCN16("Z", addr)
				log.Printf("JP Z,$%x\n", addr)

			case 0xCB:
				PrefixTable(c, log)

			case 0xCC:
				addr := int(c.Mem[c.GetReg16Val("PC")+1]) + (int(c.Mem[c.GetReg16Val("PC")+2]) << 8)
				log.Printf("CALL Z,$%x\n", addr)
				c.CALLCCN16("Z", addr)

			case 0xCD:
				addr := int(c.Mem[c.GetReg16Val("PC")+1]) + (int(c.Mem[c.GetReg16Val("PC")+2]) << 8)
				log.Printf("CALL $%x\n", addr)
				c.CALLN16(addr)

			case 0xCE:
				log.Printf("ADC A,$%x\n", int(c.Mem[c.GetReg16Val("PC")+1]))
				c.ADCAN8()

			case 0xCF:
				log.Printf("RST 08H\n")
				c.RST("08H")

			case 0xD0:
				log.Printf("RET NC\n")
				c.RETCC("NC")

			case 0xD1:
				log.Printf("POP DE\n")
				c.POPR16("DE")

			case 0xD2:
				addr := int(c.Mem[c.GetReg16Val("PC")+1]) + (int(c.Mem[c.GetReg16Val("PC")+2]) << 8)
				log.Printf("JP NC,$%x \n", addr)
				c.JPCCN16("NC", addr)

			case 0xD4:
				addr := int(c.Mem[c.GetReg16Val("PC")+1]) + (int(c.Mem[c.GetReg16Val("PC")+2]) << 8)
				log.Printf("CALL NC,$%x \n", addr)
				c.CALLCCN16("NC", addr)

			case 0xD5:
				log.Printf("PUSH DE\n")
				c.PUSHR16("DE")

			case 0xD6:
				log.Printf("SUB A,$%x\n", int(c.Mem[c.GetReg16Val("PC")+1]))
				c.SUBAN8()

			case 0xD7:
				log.Printf("RST 10H\n")
				c.RST("10H")

			case 0xD8:
				log.Printf("RET C\n")
				c.RETCC("C")

			case 0xD9:
				log.Printf("RETI\n")
				c.RETI()

			case 0xDA:
				addr := int(c.Mem[c.GetReg16Val("PC")+1]) + (int(c.Mem[c.GetReg16Val("PC")+2]) << 8)
				log.Printf("JP C,$%x\n", addr)
				c.JPCCN16("C", addr)

			case 0xDC:
				addr := int(c.Mem[c.GetReg16Val("PC")+1]) + (int(c.Mem[c.GetReg16Val("PC")+2]) << 8)
				log.Printf("CALL C,$%x\n", addr)
				c.CALLCCN16("C", addr)

			case 0xDE:
				log.Printf("SBC A,$%x\n", uint8(c.Mem[c.GetReg16Val("PC")+1]))
				c.SBCAN8()

			case 0xDF:
				log.Printf("RST 18H\n")
				c.RST("18H")

			case 0xE0:
				addr := int(uint16(0xFF00 + int(c.Mem[c.GetReg16Val("PC")+1])))
				log.Printf("LD $(%x),A\n", addr)
				c.LDN16A(addr)
				c.SetReg16Val("PC", c.GetReg16Val("PC")-1)

			case 0xE1:
				log.Printf("POP HL\n")
				c.POPR16("HL")

			case 0xE2:
				log.Printf("LD ($FF00+C),A\n")
				c.LDN16A(int(uint16(0xFF00 + int(c.GetReg8Val("C")))))
				c.SetReg16Val("PC", c.GetReg16Val("PC")-2)

			case 0xE5:
				log.Printf("PUSH HL\n")
				c.PUSHR16("HL")

			case 0xE6:
				log.Printf("AND A,$%x\n", uint8(c.Mem[c.GetReg16Val("PC")+1]))
				c.ANDAN8()

			case 0xE7:
				c.RST("20H")
				log.Printf("RST 20H\n")

			case 0xE8:
				c.ADDSPE8()
				log.Printf("ADD SP,%x\n", int8(c.Mem[c.GetReg16Val("PC")+1]))

			case 0xE9:
				log.Printf("JP HL\n")
				c.JPHL()

			case 0xEA:
				addr := int(uint16(int(c.Mem[c.GetReg16Val("PC")+1]) + (int(c.Mem[c.GetReg16Val("PC")+2]) << 8)))
				log.Printf("LD ($%x),A\n", addr)
				c.LDN16A(addr)

			case 0xEE:
				log.Printf("XOR $%x\n", uint8(c.Mem[c.GetReg16Val("PC")+1]))
				c.XORAN8()

			case 0xEF:
				log.Printf("RST 28H\n")
				c.RST("28H")

			case 0xF0:
				addr := int(uint16(0xFF00 + int(uint8(c.Mem[c.GetReg16Val("PC")+1]))))
				log.Printf("LD A,(%x)\n", addr)
				c.LDAN16(addr)

			case 0xF1:
				log.Printf("POP AF\n")
				c.POPFA()

			case 0xF2:
				log.Printf("LD A,(0xFF00+C)\n")
				c.LDAN16(int(uint16(0xFF00 + c.GetReg8Val("C"))))
				c.SetReg16Val("PC", c.GetReg16Val("PC")-1)

			case 0xF3:
				log.Printf("DI\n")
				c.DI()

			case 0xF4:

			case 0xF5:
				log.Printf("PUSH AF\n")
				c.PUSHFA()

			case 0xF6:
				log.Printf("OR A,$%x\n", int8(c.Mem[c.GetReg16Val("PC")+1]))
				c.ORAN8()

			case 0xF7:
				log.Printf("RST 30H\n")
				c.RST("30H")

			case 0xF8:
				log.Printf("LD HL,SP+%x\n", int8(c.Mem[c.GetReg16Val("PC")+1]))
				c.LDHLSPE8()

			case 0xF9:
				log.Printf("LD SP,HL\n")
				c.LDSPHL()

			case 0xFA:
				addr := int(uint16(int(c.Mem[c.GetReg16Val("PC")+1]) + int(c.Mem[c.GetReg16Val("PC")+2])<<8))
				log.Printf("LD A,($%x)\n", addr)
				c.LDAN16(addr)
				c.SetReg16Val("PC", c.GetReg16Val("PC")+1)

			case 0xFB:
				log.Printf("EI\n")
				c.EI()

			case 0xFE:
				log.Printf("CP $%x\n", int8(c.Mem[c.GetReg16Val("PC")+1]))
				c.CPAN8()

			case 0xFF:
				log.Printf("RST 38H\n")
				c.RST("38H")

			}
		}
	}

}
