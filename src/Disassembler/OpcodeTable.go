package Disassembler

import (
	"os"

	CPU "github.com/shivkar2n/GB-Emulator/CPU"
)

func OpCodeTable(c *CPU.CPU) {
	f, _ := os.OpenFile("message.gb", os.O_RDWR|os.O_CREATE, 0644)
	// Load rom into memory
	c.LoadROM()

	// Switch to Stop execution
	var stopExec bool
	stopExec = false

	// Main execution loop
	for c.GetReg16Val("PC") < 0xFFFF && !stopExec {
		// For handling interrupt service requests
		c.InterruptHandler()

		// For incrementing timer
		c.INCRDIV()

		c.DebugInfo()

		var b []byte
		b = append([]byte{c.Mem[0xFF01]}, b...)
		f.Write(b)

		// if c.GetReg16Val("PC") == 0xC2BB {
		// 	stopExec = true
		// 	// c.MemInfo()
		// 	println(c.GetIFVal())
		// }

		switch uint8(c.Mem[c.GetReg16Val("PC")] & 0xF0) {
		case 0x00:
			switch c.Mem[c.GetReg16Val("PC")] {
			case 0x00:
				// fmt.Prinf("NOP\n")
				c.NOP()

			case 0x01:
				c.LDR16N16("BC")
				// fmt.Prinf("LD BC,$%x \n", addr)

			case 0x02:
				c.LDR16A("BC")
				// fmt.Prinf("LD BC, A\n")

			case 0x03:
				c.INCR16("BC")
				// fmt.Prinf("INC BC\n")

			case 0x04:
				c.INCR8("B")
				// fmt.Prinf("INC B\n")

			case 0x05:
				c.DECR8("B")
				// fmt.Prinf("DEC B\n")

			case 0x06:
				addr := int(c.Mem[c.GetReg16Val("PC")+1])
				c.LDR8N8("B", addr)
				// fmt.Prinf("LD B,$%x\n", addr)

			case 0x07:
				c.RLCA()
				// fmt.Prinf("RLCA\n")

			case 0x08:
				addr := int(c.Mem[c.GetReg16Val("PC")+1]) + (int(c.Mem[c.GetReg16Val("PC")+2]) << 8)
				c.LDN16SP(addr)
				// fmt.Prinf("LD (%x), SP\n", addr)

			case 0x09:
				c.ADDHLR16("BC")
				// fmt.Prinf("ADD HL, BC\n")

			case 0x0A:
				c.LDAR16("BC")
				// fmt.Prinf("LD A, (BC)\n")

			case 0x0B:
				c.DECR16("BC")
				// fmt.Prinf("DEC BC\n")

			case 0x0C:
				c.INCR8("C")
				// fmt.Prinf("INC C\n")

			case 0x0D:
				c.DECR8("C")
				// fmt.Prinf("DEC C\n")

			case 0x0E:
				addr := int(c.Mem[c.GetReg16Val("PC")+1])
				c.LDR8N8("C", addr)
				// fmt.Prinf("LD C, $%x\n", addr)

			case 0x0F:
				c.RRCA()
				// fmt.Prinf("RRCA \n")

			}

		case 0x10:
			switch c.Mem[c.GetReg16Val("PC")] {
			case 0x10:
				// addr := c.Mem[c.GetReg16Val("PC")+1]
				// fmt.Prinf("STOP $%x\n", addr)

			case 0x11:
				c.LDR16N16("DE")
				// fmt.Prinf("LD DE,$%x\n", addr)

			case 0x12:
				c.LDR16A("DE")
				// fmt.Prinf("LD (DE),A\n")

			case 0x13:
				c.INCR16("DE")
				// fmt.Prinf("INC (DE)\n")

			case 0x14:
				c.INCR8("D")
				// fmt.Prinf("INC D\n")

			case 0x15:
				c.DECR8("D")
				// fmt.Prinf("DEC D\n")

			case 0x16:
				addr := int(c.Mem[c.GetReg16Val("PC")+1])
				c.LDR8N8("D", addr)
				// fmt.Prinf("LD D,$%x\n", addr)

			case 0x17:
				c.RLA()
				// fmt.Prinf("RLA\n")

			case 0x18:
				// addr := int(int8(c.Mem[c.GetReg16Val("PC")+1]))
				c.JRN16()
				// fmt.Prinf("JR r%x\n", addr)

			case 0x19:
				c.ADDHLR16("DE")
				// fmt.Prinf("ADD HL, DE\n")

			case 0x1A:
				c.LDAR16("DE")
				// fmt.Prinf("LD A, (DE)\n")

			case 0x1B:
				c.DECR16("DE")
				// fmt.Prinf("DEC DE\n")

			case 0x1C:
				c.INCR8("E")
				// fmt.Prinf("INC E\n")

			case 0x1D:
				c.DECR8("E")
				// fmt.Prinf("DEC E\n")

			case 0x1E:
				addr := int(c.Mem[c.GetReg16Val("PC")+1])
				c.LDR8N8("E", addr)
				// fmt.Prinf("LD E,$%x\n", addr)

			case 0x1F:
				c.RRA()
				// fmt.Prinf("RRA\n")

			}

		case 0x20:
			switch c.Mem[c.GetReg16Val("PC")] {
			case 0x20:
				// addr := int(c.Mem[c.GetReg16Val("PC")+1])
				// fmt.Prinf("JR NZ,%x\n", addr)
				c.JRCCN16("NZ")

			case 0x21:
				c.LDR16N16("HL")
				// fmt.Prinf("LD HL,$%x\n", addr)

			case 0x22:
				c.LDHLIA()
				// fmt.Prinf("LD (HL+),A\n")

			case 0x23:
				c.INCR16("HL")
				// fmt.Prinf("INC HL\n")

			case 0x24:
				c.INCR8("H")
				// fmt.Prinf("INC H\n")

			case 0x25:
				c.DECR8("H")
				// fmt.Prinf("DEC H\n")

			case 0x26:
				addr := int(c.Mem[c.GetReg16Val("PC")+1])
				c.LDR8N8("H", addr)
				// fmt.Prinf("LD H,$%x\n", addr)

			case 0x27:
				// fmt.Prinf("DAA\n")

			case 0x28:
				// addr := int(c.Mem[c.GetReg16Val("PC")+1])
				// fmt.Prinf("JR Z,r%x\n", addr)
				c.JRCCN16("Z")

			case 0x29:
				c.ADDHLR16("HL")
				// fmt.Prinf("ADD HL,HL\n")

			case 0x2A:
				c.LDAHLI()
				// fmt.Prinf("LD A,(HL+)\n")

			case 0x2B:
				c.DECR16("HL")
				// fmt.Prinf("DEC HL\n")

			case 0x2C:
				c.INCR8("L")
				// fmt.Prinf("INC L\n")

			case 0x2D:
				c.DECR8("L")
				// fmt.Prinf("DEC L\n")

			case 0x2E:
				addr := int(c.Mem[c.GetReg16Val("PC")+1])
				c.LDR8N8("L", addr)
				// fmt.Prinf("LD L,$%x\n", addr)

			case 0x2F:
				c.CPL()
				// fmt.Prinf("CPL\n")

			}

		case 0x30:
			switch c.Mem[c.GetReg16Val("PC")] {
			case 0x30:
				// addr := int(c.Mem[c.GetReg16Val("PC")+1])
				// fmt.Prinf("JR NC,r%x\n", addr)
				c.JRCCN16("NC")

			case 0x31:
				c.LDR16N16("SP")
				// fmt.Prinf("LD SP $%x\n", addr)

			case 0x32:
				c.LDHLDA()
				// fmt.Prinf("LD (HL-),A\n")

			case 0x33:
				c.INCR16("SP")
				// fmt.Prinf("INC SP\n")

			case 0x34:
				c.INCRHL()
				// fmt.Prinf("INC (HL)\n")

			case 0x35:
				c.DECHL()
				// fmt.Prinf("DEC (HL)\n")

			case 0x36:
				addr := int(c.Mem[c.GetReg16Val("PC")+1])
				c.LDHLN8(addr)
				// fmt.Prinf("LD (HL),$%x\n", addr)

			case 0x37:
				c.SCF()
				// fmt.Prinf("SCF\n")

			case 0x38:
				// fmt.Prinf("JR C,r%x\n", int(c.Mem[c.GetReg16Val("PC")+1]))
				c.JRCCN16("C")

			case 0x39:
				c.ADDHLSP()
				// fmt.Prinf("ADD HL,SP\n")

			case 0x3A:
				c.LDAHLD()
				// fmt.Prinf("LD A,(HL-)\n")

			case 0x3B:
				c.DECR16("SP")
				// fmt.Prinf("DEC SP\n")

			case 0x3C:
				c.INCR8("A")
				// fmt.Prinf("INC A\n")

			case 0x3D:
				c.DECR8("A")
				// fmt.Prinf("DEC A\n")

			case 0x3E:
				addr := int(c.Mem[c.GetReg16Val("PC")+1])
				c.LDR8N8("A", addr)
				// fmt.Prinf("LD A,$%x\n", addr)

			case 0x3F:
				c.CCF()
				// fmt.Prinf("CCF\n")

			}

		case 0x40:
			switch c.Mem[c.GetReg16Val("PC")] {
			case 0x40:
				c.LDR8R8("B", "B")
				// fmt.Prinf("LD B,B\n")

			case 0x41:
				c.LDR8R8("B", "C")
				// fmt.Prinf("LD B,C\n")

			case 0x42:
				c.LDR8R8("B", "D")
				// fmt.Prinf("LD B,D\n")

			case 0x43:
				c.LDR8R8("B", "E")
				// fmt.Prinf("LD B,E\n")

			case 0x44:
				c.LDR8R8("B", "H")
				// fmt.Prinf("LD B,H\n")

			case 0x45:
				c.LDR8R8("B", "L")
				// fmt.Prinf("LD B,L\n")

			case 0x46:
				c.LDR8HL("B")
				// fmt.Prinf("LD B,(HL)\n")

			case 0x47:
				c.LDR8R8("B", "A")
				// fmt.Prinf("LD B,A\n")

			case 0x48:
				c.LDR8R8("C", "B")
				// fmt.Prinf("LD C,B\n")

			case 0x49:
				c.LDR8R8("C", "C")
				// fmt.Prinf("LD C,C\n")

			case 0x4A:
				c.LDR8R8("C", "D")
				// fmt.Prinf("LD C,D\n")

			case 0x4B:
				c.LDR8R8("C", "E")
				// fmt.Prinf("LD C,E\n")

			case 0x4C:
				c.LDR8R8("C", "H")
				// fmt.Prinf("LD C,H\n")

			case 0x4D:
				c.LDR8R8("C", "L")
				// fmt.Prinf("LD C,L\n")

			case 0x4E:
				c.LDR8HL("C")
				// fmt.Prinf("LD C,(HL)\n")

			case 0x4F:
				c.LDR8R8("C", "A")
				// fmt.Prinf("LD C,A\n")

			}

		case 0x50:
			switch c.Mem[c.GetReg16Val("PC")] {
			case 0x50:
				c.LDR8R8("D", "B")
				// fmt.Prinf("LD D,B\n")

			case 0x51:
				c.LDR8R8("D", "C")
				// fmt.Prinf("LD D,C\n")

			case 0x52:
				c.LDR8R8("D", "D")
				// fmt.Prinf("LD D,D\n")

			case 0x53:
				c.LDR8R8("D", "E")
				// fmt.Prinf("LD D,E\n")

			case 0x54:
				c.LDR8R8("D", "H")
				// fmt.Prinf("LD D,H\n")

			case 0x55:
				c.LDR8R8("D", "L")
				// fmt.Prinf("LD D,L\n")

			case 0x56:
				c.LDR8HL("D")
				// fmt.Prinf("LD D,(HL)\n")

			case 0x57:
				c.LDR8R8("D", "A")
				// fmt.Prinf("LD D,A\n")

			case 0x58:
				c.LDR8R8("E", "B")
				// fmt.Prinf("LD E,B\n")

			case 0x59:
				c.LDR8R8("E", "C")
				// fmt.Prinf("LD E,C\n")

			case 0x5A:
				c.LDR8R8("E", "D")
				// fmt.Prinf("LD E,D\n")

			case 0x5B:
				c.LDR8R8("E", "E")
				// fmt.Prinf("LD E,E\n")

			case 0x5C:
				c.LDR8R8("E", "H")
				// fmt.Prinf("LD E,H\n")

			case 0x5D:
				c.LDR8R8("E", "L")
				// fmt.Prinf("LD E,L\n")

			case 0x5E:
				c.LDR8HL("E")
				// fmt.Prinf("LD E,(HL)\n")

			case 0x5F:
				c.LDR8R8("E", "A")
				// fmt.Prinf("LD E,A\n")

			}

		case 0x60:
			switch c.Mem[c.GetReg16Val("PC")] {
			case 0x60:
				c.LDR8R8("H", "B")
				// fmt.Prinf("LD H,B\n")

			case 0x61:
				c.LDR8R8("H", "C")
				// fmt.Prinf("LD H,C\n")

			case 0x62:
				c.LDR8R8("H", "D")
				// fmt.Prinf("LD H,D\n")

			case 0x63:
				c.LDR8R8("H", "E")
				// fmt.Prinf("LD H,E\n")

			case 0x64:
				c.LDR8R8("H", "H")
				// fmt.Prinf("LD H,H\n")

			case 0x65:
				c.LDR8R8("H", "L")
				// fmt.Prinf("LD H,L\n")

			case 0x66:
				c.LDR8HL("H")
				// fmt.Prinf("LD H,(HL)\n")

			case 0x67:
				c.LDR8R8("H", "A")
				// fmt.Prinf("LD H,A\n")

			case 0x68:
				c.LDR8R8("L", "B")
				// fmt.Prinf("LD L,B\n")

			case 0x69:
				c.LDR8R8("L", "C")
				// fmt.Prinf("LD L,C\n")

			case 0x6A:
				c.LDR8R8("L", "D")
				// fmt.Prinf("LD L,D\n")

			case 0x6B:
				c.LDR8R8("L", "E")
				// fmt.Prinf("LD L,E\n")

			case 0x6C:
				c.LDR8R8("L", "H")
				// fmt.Prinf("LD L,H\n")

			case 0x6D:
				c.LDR8R8("L", "L")
				// fmt.Prinf("LD L,L\n")

			case 0x6E:
				c.LDR8HL("L")
				// fmt.Prinf("LD L,(HL)\n")

			case 0x6F:
				c.LDR8R8("L", "A")
				// fmt.Prinf("LD L,A\n")

			}

		case 0x70:
			switch c.Mem[c.GetReg16Val("PC")] {
			case 0x70:
				c.LDHLR8("B")
				// fmt.Prinf("LD (HL),B\n")

			case 0x71:
				c.LDHLR8("C")
				// fmt.Prinf("LD (HL),C\n")

			case 0x72:
				c.LDHLR8("D")
				// fmt.Prinf("LD (HL),D\n")

			case 0x73:
				c.LDHLR8("E")
				// fmt.Prinf("LD (HL),E\n")

			case 0x74:
				c.LDHLR8("H")
				// fmt.Prinf("LD (HL),H\n")

			case 0x75:
				c.LDHLR8("L")
				// fmt.Prinf("LD (HL),L\n")

			case 0x76:
				c.HALT()
				// fmt.Prinf("HALT\n")

			case 0x77:
				c.LDHLR8("A")
				// fmt.Prinf("LD (HL),A\n")

			case 0x78:
				c.LDR8R8("A", "B")
				// fmt.Prinf("LD A,B\n")

			case 0x79:
				c.LDR8R8("A", "C")
				// fmt.Prinf("LD A,C\n")

			case 0x7A:
				c.LDR8R8("A", "D")
				// fmt.Prinf("LD A,D\n")

			case 0x7B:
				c.LDR8R8("A", "E")
				// fmt.Prinf("LD A,E\n")

			case 0x7C:
				c.LDR8R8("A", "H")
				// fmt.Prinf("LD A,H\n")

			case 0x7D:
				c.LDR8R8("A", "L")
				// fmt.Prinf("LD A,L\n")

			case 0x7E:
				c.LDR8HL("A")
				// fmt.Prinf("LD A,(HL)\n")

			case 0x7F:
				c.LDR8R8("A", "A")
				// fmt.Prinf("LD A,A\n")

			}

		case 0x80:
			switch c.Mem[c.GetReg16Val("PC")] {
			case 0x80:
				c.ADDAR8("B")
				// fmt.Prinf("ADD A,B\n")

			case 0x81:
				c.ADDAR8("C")
				// fmt.Prinf("ADD A,C\n")

			case 0x82:
				c.ADDAR8("D")
				// fmt.Prinf("ADD A,D\n")

			case 0x83:
				c.ADDAR8("E")
				// fmt.Prinf("ADD A,E\n")

			case 0x84:
				c.ADDAR8("H")
				// fmt.Prinf("ADD A,H\n")

			case 0x85:
				c.ADDAR8("L")
				// fmt.Prinf("ADD A,L\n")

			case 0x86:
				c.ADDAHL()
				// fmt.Prinf("ADD A,(HL)\n")

			case 0x87:
				c.ADDAR8("A")
				// fmt.Prinf("ADD A,A\n")

			case 0x88:
				c.ADCAR8("B")
				// fmt.Prinf("ADC A,B\n")

			case 0x89:
				c.ADCAR8("C")
				// fmt.Prinf("ADC A,C\n")

			case 0x8A:
				c.ADCAR8("D")
				// fmt.Prinf("ADC A,D\n")

			case 0x8B:
				c.ADCAR8("E")
				// fmt.Prinf("ADC A,E\n")

			case 0x8C:
				c.ADCAR8("H")
				// fmt.Prinf("ADC A,H\n")

			case 0x8D:
				c.ADCAR8("L")
				// fmt.Prinf("ADC A,L\n")

			case 0x8E:
				c.ADCAHL()
				// fmt.Prinf("ADC A,(HL)\n")

			case 0x8F:
				c.ADCAR8("A")
				// fmt.Prinf("ADC A,A\n")

			}

		case 0x90:
			switch c.Mem[c.GetReg16Val("PC")] {
			case 0x90:
				c.SUBAR8("B")
				// fmt.Prinf("SUB A,B\n")

			case 0x91:
				c.SUBAR8("C")
				// fmt.Prinf("SUB A,C\n")

			case 0x92:
				c.SUBAR8("D")
				// fmt.Prinf("SUB A,D\n")

			case 0x93:
				c.SUBAR8("E")
				// fmt.Prinf("SUB A,E\n")

			case 0x94:
				c.SUBAR8("H")
				// fmt.Prinf("SUB A,H\n")

			case 0x95:
				c.SUBAR8("L")
				// fmt.Prinf("SUB A,L\n")

			case 0x96:
				c.SUBAHL()
				// fmt.Prinf("SUB A,(HL)\n")

			case 0x97:
				c.SUBAR8("A")
				// fmt.Prinf("SUB A,A\n")

			case 0x98:
				c.SBCAR8("B")
				// fmt.Prinf("SBC A,B\n")

			case 0x99:
				c.SBCAR8("C")
				// fmt.Prinf("SBC A,C\n")

			case 0x9A:
				c.SBCAR8("D")
				// fmt.Prinf("SBC A,D\n")

			case 0x9B:
				c.SBCAR8("E")
				// fmt.Prinf("SBC A,E\n")

			case 0x9C:
				c.SBCAR8("H")
				// fmt.Prinf("SBC A,H\n")

			case 0x9D:
				c.SBCAR8("L")
				// fmt.Prinf("SBC A,L\n")

			case 0x9E:
				c.SBCAHL()
				// fmt.Prinf("SBC A,(HL)\n")

			case 0x9F:
				c.SBCAR8("A")
				// fmt.Prinf("SBC A,A\n")

			}

		case 0xA0:
			switch c.Mem[c.GetReg16Val("PC")] {
			case 0xA0:
				c.ANDAR8("B")
				// fmt.Prinf("AND A,B\n")

			case 0xA1:
				c.ANDAR8("C")
				// fmt.Prinf("AND A,C\n")

			case 0xA2:
				c.ANDAR8("D")
				// fmt.Prinf("AND A,D\n")

			case 0xA3:
				c.ANDAR8("E")
				// fmt.Prinf("AND A,E\n")

			case 0xA4:
				c.ANDAR8("H")
				// fmt.Prinf("AND A,H\n")

			case 0xA5:
				c.ANDAR8("L")
				// fmt.Prinf("AND A,L\n")

			case 0xA6:
				c.ANDAHL()
				// fmt.Prinf("AND A,(HL)\n")

			case 0xA7:
				c.ANDAR8("A")
				// fmt.Prinf("AND A,A\n")

			case 0xA8:
				c.XORAR8("B")
				// fmt.Prinf("XOR A,B\n")

			case 0xA9:
				c.XORAR8("C")
				// fmt.Prinf("XOR A,C\n")

			case 0xAA:
				c.XORAR8("D")
				// fmt.Prinf("XOR A,D\n")

			case 0xAB:
				c.XORAR8("E")
				// fmt.Prinf("XOR A,E\n")

			case 0xAC:
				c.XORAR8("H")
				// fmt.Prinf("XOR A,H\n")

			case 0xAD:
				c.XORAR8("L")
				// fmt.Prinf("XOR A,L\n")

			case 0xAE:
				c.XORAHL()
				// fmt.Prinf("XOR A,(HL)\n")

			case 0xAF:
				c.XORAR8("A")
				// fmt.Prinf("XOR A,A\n")

			}

		case 0xB0:
			switch c.Mem[c.GetReg16Val("PC")] {
			case 0xB0:
				c.ORAR8("B")
				// fmt.Prinf("OR A,B\n")

			case 0xB1:
				c.ORAR8("C")
				// fmt.Prinf("OR A,C\n")

			case 0xB2:
				c.ORAR8("D")
				// fmt.Prinf("OR A,D\n")

			case 0xB3:
				c.ORAR8("E")
				// fmt.Prinf("OR A,E\n")

			case 0xB4:
				c.ORAR8("H")
				// fmt.Prinf("OR A,H\n")

			case 0xB5:
				c.ORAR8("L")
				// fmt.Prinf("OR A,L\n")

			case 0xB6:
				c.ORAHL()
				// fmt.Prinf("OR A,(HL)\n")

			case 0xB7:
				c.ORAR8("A")
				// fmt.Prinf("OR A,A\n")

			case 0xB8:
				c.CPAR8("B")
				// fmt.Prinf("CP A,B\n")

			case 0xB9:
				c.CPAR8("C")
				// fmt.Prinf("CP A,C\n")

			case 0xBA:
				c.CPAR8("D")
				// fmt.Prinf("CP A,D\n")

			case 0xBB:
				c.CPAR8("E")
				// fmt.Prinf("CP A,E\n")

			case 0xBC:
				c.CPAR8("H")
				// fmt.Prinf("CP A,H\n")

			case 0xBD:
				c.CPAR8("L")
				// fmt.Prinf("CP A,L\n")

			case 0xBE:
				c.CPAHL()
				// fmt.Prinf("CP A,(HL)\n")

			case 0xBF:
				c.CPAR8("A")
				// fmt.Prinf("CP A,A\n")

			}

		case 0xC0:
			switch c.Mem[c.GetReg16Val("PC")] {
			case 0xC0:
				c.RETCC("NZ")
				// fmt.Prinf("RET NZ\n")

			case 0xC1:
				c.POPR16("BC")
				// fmt.Prinf("POP BC\n")

			case 0xC2:
				addr := int(c.Mem[c.GetReg16Val("PC")+1]) + (int(c.Mem[c.GetReg16Val("PC")+2]) << 8)
				c.JPCCN16("NZ", addr)
				// fmt.Prinf("JP NZ,$%x \n", addr)

			case 0xC3:
				addr := int(c.Mem[c.GetReg16Val("PC")+1]) + (int(c.Mem[c.GetReg16Val("PC")+2]) << 8)
				c.JPN16(addr)
				// fmt.Prinf("JP $%x \n", addr)

			case 0xC4:
				addr := int(uint16(c.Mem[c.GetReg16Val("PC")+1]) + (uint16(c.Mem[c.GetReg16Val("PC")+2]) << 8))
				c.CALLCCN16("NZ", addr)
				// fmt.Prinf("CALL NZ,$%x \n", addr)

			case 0xC5:
				c.PUSHR16("BC")
				// fmt.Prinf("PUSH BC\n")

			case 0xC6:
				addr := int(c.Mem[c.GetReg16Val("PC")+1])
				c.ADDAN8(addr)
				// fmt.Prinf("ADD A,$%x\n", addr)

			case 0xC7:
				c.RST("00H")
				// fmt.Prinf("RST 00H\n")

			case 0xC8:
				c.RETCC("Z")
				// fmt.Prinf("RET Z\n")

			case 0xC9:
				c.RET()
				// fmt.Prinf("RET\n")

			case 0xCA:
				addr := int(c.Mem[c.GetReg16Val("PC")+1]) + (int(c.Mem[c.GetReg16Val("PC")+2]) << 8)
				// fmt.Prinf("JP Z,$%x\n", addr)
				c.JPCCN16("Z", addr)

			case 0xCB:
				PrefixTable(c)

			case 0xCC:
				addr := int(c.Mem[c.GetReg16Val("PC")+1]) + (int(c.Mem[c.GetReg16Val("PC")+2]) << 8)
				c.CALLCCN16("Z", addr)
				// fmt.Prinf("CALL Z,$%x\n", addr)

			case 0xCD:
				addr := int(c.Mem[c.GetReg16Val("PC")+1]) + (int(c.Mem[c.GetReg16Val("PC")+2]) << 8)
				c.CALLN16(addr)
				// fmt.Prinf("CALL $%x\n", addr)

			case 0xCE:
				addr := int(c.Mem[c.GetReg16Val("PC")+1])
				c.ADCAN8(addr)
				// fmt.Prinf("ADC A,$%x\n", addr)

			case 0xCF:
				c.RST("08H")
				// fmt.Prinf("RST 08H\n")

			}

		case 0xD0:
			switch c.Mem[c.GetReg16Val("PC")] {
			case 0xD0:
				c.RETCC("NC")
				// fmt.Prinf("RET NC\n")

			case 0xD1:
				c.POPR16("DE")
				// fmt.Prinf("POP DE\n")

			case 0xD2:
				addr := int(c.Mem[c.GetReg16Val("PC")+1]) + (int(c.Mem[c.GetReg16Val("PC")+2]) << 8)
				c.JPCCN16("NC", addr)
				// fmt.Prinf("JP NC,$%x \n", addr)

			case 0xD4:
				addr := int(c.Mem[c.GetReg16Val("PC")+1]) + (int(c.Mem[c.GetReg16Val("PC")+2]) << 8)
				c.CALLCCN16("NC", addr)
				// fmt.Prinf("CALL NC,$%x \n", addr)

			case 0xD5:
				c.PUSHR16("DE")
				// fmt.Prinf("PUSH DE\n")

			case 0xD6:
				addr := int(c.Mem[c.GetReg16Val("PC")+1])
				c.SUBAN8(addr)
				// fmt.Prinf("SUB A,$%x\n", addr)

			case 0xD7:
				c.RST("10H")
				// fmt.Prinf("RST 10H\n")

			case 0xD8:
				c.RETCC("C")
				// fmt.Prinf("RET C\n")

			case 0xD9:
				c.RETI()
				// fmt.Prinf("RETI\n")

			case 0xDA:
				addr := int(c.Mem[c.GetReg16Val("PC")+1]) + (int(c.Mem[c.GetReg16Val("PC")+2]) << 8)
				c.JPCCN16("C", addr)
				// fmt.Prinf("JP C,$%x\n", addr)

			case 0xDC:
				addr := int(c.Mem[c.GetReg16Val("PC")+1]) + (int(c.Mem[c.GetReg16Val("PC")+2]) << 8)
				c.CALLCCN16("C", addr)
				// fmt.Prinf("CALL C,$%x\n", addr)

			case 0xDE:
				addr := int(int8(c.Mem[c.GetReg16Val("PC")+1]))
				c.SBCAN8(addr)
				// fmt.Prinf("SBC A,$%x\n", addr)

			case 0xDF:
				c.RST("18H")
				// fmt.Prinf("RST 18H\n")

			default:
				c.SetReg16Val("PC", c.GetReg16Val("PC")+1)

			}

		case 0xE0:
			switch c.Mem[c.GetReg16Val("PC")] {
			case 0xE0:
				addr := int(uint16(0xFF00 + int(c.Mem[c.GetReg16Val("PC")+1])))
				c.LDN16A(addr)
				c.SetReg16Val("PC", c.GetReg16Val("PC")-1)
				// fmt.Prinf("LD ($FFOO+$%x),A\n", addr)

			case 0xE1:
				c.POPR16("HL")
				// fmt.Prinf("POP HL\n")

			case 0xE2:
				c.LDN16A(int(uint16(0xFF00 + int(c.GetReg8Val("C")))))
				c.SetReg16Val("PC", c.GetReg16Val("PC")-2)
				// fmt.Prinf("LD ($FF00+C),A\n")

			case 0xE5:
				c.PUSHR16("HL")
				// fmt.Prinf("PUSH HL\n")

			case 0xE6:
				addr := int(c.Mem[c.GetReg16Val("PC")+1])
				c.ANDAN8(addr)
				// fmt.Prinf("AND,$%x\n", addr)

			case 0xE7:
				c.RST("20H")
				// fmt.Prinf("RST 20H\n")

			case 0xE8:
				addr := int(int8(int(c.Mem[c.GetReg16Val("PC")+1])))
				c.ADDSPE8(addr)
				// fmt.Prinf("ADD SP,r%x\n", addr)

			case 0xE9:
				c.JPHL()
				// fmt.Prinf("JP HL\n")

			case 0xEA:
				addr := int(uint16(int(c.Mem[c.GetReg16Val("PC")+1]) + (int(c.Mem[c.GetReg16Val("PC")+2]) << 8)))
				c.LDN16A(addr)
				// fmt.Prinf("LD ($%x),A\n", addr)

			case 0xEE:
				addr := int(c.Mem[c.GetReg16Val("PC")+1])
				c.XORAN8(addr)
				// fmt.Prinf("XOR $%x\n", addr)

			case 0xEF:
				c.RST("28H")
				// fmt.Prinf("RST 28H\n")

			default:
				c.SetReg16Val("PC", c.GetReg16Val("PC")+1)

			}

		case 0xF0:
			switch c.Mem[c.GetReg16Val("PC")] {
			case 0xF0:
				addr := int(uint16(0xFF00 + int(uint8(c.Mem[c.GetReg16Val("PC")+1]))))
				c.LDAN16(addr)
				// fmt.Prinf("LD A,($FF00+%x)\n", addr)

			case 0xF1:
				c.POPFA()
				// fmt.Prinf("POP AF\n")

			case 0xF2:
				c.LDAN16(int(uint16(0xFF00 + c.GetReg8Val("C"))))
				c.SetReg16Val("PC", c.GetReg16Val("PC")-1)

				// fmt.Prinf("LD A,(C)\n")

			case 0xF3:
				c.DI()
				// fmt.Prinf("DI\n")

			case 0xF4:

			case 0xF5:
				c.PUSHAF()
				// fmt.Prinf("PUSH AF\n")

			case 0xF6:
				addr := int(c.Mem[c.GetReg16Val("PC")+1])
				c.ORAN8(addr)
				// fmt.Prinf("OR A,$%x\n", addr)

			case 0xF7:
				c.RST("30H")
				// fmt.Prinf("RST 30H\n")

			case 0xF8:
				addr := int(int8(c.Mem[c.GetReg16Val("PC")+1]))
				c.LDHLSPE8(addr)
				// fmt.Prinf("LD HL,SP+E8%x\n", addr)

			case 0xF9:
				c.LDSPHL()
				// fmt.Prinf("LD SP,HL\n")

			case 0xFA:
				addr := int(uint16(int(c.Mem[c.GetReg16Val("PC")+1]) + int(c.Mem[c.GetReg16Val("PC")+2])<<8))
				// if c.GetReg16Val("PC") == 0xC7FE {
				// 	fmt.Printf("%x\n", addr)
				// 	fmt.Printf("%d\n", c.Mem[addr])
				// }
				c.LDAN16(addr)
				c.SetReg16Val("PC", c.GetReg16Val("PC")+1)
				// fmt.Prinf("LD A,($%x)\n", addr)

			case 0xFB:
				c.EI()
				// fmt.Prinf("EI\n")

			case 0xFE:
				addr := int(c.Mem[c.GetReg16Val("PC")+1])
				c.CPAN8(addr)
				// fmt.Prinf("CP $%x\n", addr)

			case 0xFF:
				c.RST("38H")
				// fmt.Prinf("RST 38H\n")

			default:
				c.SetReg16Val("PC", c.GetReg16Val("PC")+1)

			}
		default:
			break
		}
	}

}
