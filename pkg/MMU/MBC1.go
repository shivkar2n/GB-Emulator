package MMU

func (m *MMU) Read(addr int) int {
	return int(m.Ram[addr])
}

func (m *MMU) Write(val int, addr int) {
	m.Ram[addr] = byte(val)
}

// func (m *MMU) Read(addr int) int {
// 	if addr >= 0x0000 && addr <= 0x3FFF {
// 			return int(m.Rom[addr])

// 	} else if addr >= 0x4000 && addr <= 0x7FFF {
// 		if len(m.Rom) < 1024*1024 {
// 			m.highBank = m.romBank & m.bitMask
// 		}
// 		return int(m.Rom[0x4000*m.highBank+addr-0x4000])

// 	} else if addr >= 0x8000 && addr <= 0x9FFF {
// 		return int(m.Ram[addr])

// 	} else if addr >= 0xA000 && addr <= 0xBFFF {
// 		if m.enableExtRam {
// 			m.highBank = m.romBank & m.bitMask
// 			return int(m.Rom[0x4000*m.highBank+addr-0x4000])
// 		}
// 		return 0xFF

// 	} else if addr >= 0xC000 && addr <= 0xDFFF {
// 		return int(m.Ram[addr])

// 	} else if addr >= 0xE000 && addr <= 0xFE9F {
// 		return int(m.Ram[addr])

// 	} else if addr >= 0xFEA0 && addr <= 0xFEFF {
// 		return 0 

// 	} else if addr >= 0xFF00 && addr <= 0xFFFF {
// 		return int(m.Ram[addr])
// 	} 
 
// 	return 0
// }

// func (m *MMU) Write(val int, addr int) {
// 	if addr >= 0x0000 && addr <= 0x1FFF {
// 		if val&0xF == 0xA {
// 			m.enableExtRam = true
// 		} else {
// 			m.enableExtRam = false
// 		}
// 		m.Ram[addr] = byte(val)

// 	} else if addr >= 0x2000 && addr <= 0x3FFF {
// 		if val == 0 {
// 			m.romBank = 1
// 		} else {
// 			m.romBank = val & m.bitMask
// 		}
// 		m.Ram[addr] = byte(val)

// 	} else if addr >= 0x4000 && addr <= 0x5FFF {
// 		m.ramBank = m.ramBank | (val & 0x3)
// 		m.Ram[addr] = byte(val)

// 	} else if addr >= 0x6000 && addr <= 0x7FFF {
// 		m.mode = (val & 1) == 1
// 		m.Ram[addr] = byte(val)

// 	} else if addr >= 0x8000 && addr <= 0x9FFF {
// 		m.Ram[addr] = byte(val)

// 	} else if addr >= 0xA000 && addr <= 0xBFFF {
// 		if m.enableExtRam {
// 		}

// 	} else if addr >= 0xC000 && addr <= 0xDFFF {
// 		m.Ram[addr] = byte(val)

// 	} else if addr >= 0xFF00 && addr <= 0xFFFF {
// 		m.Ram[addr] = byte(val)
// 	}

// }
