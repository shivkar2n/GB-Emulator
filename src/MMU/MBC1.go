package MMU

func (m *MMU) Read(addr int) byte {
	return m.Ram[addr]
}

func (m *MMU) Write(val int, addr int) {
	m.Ram[addr] = byte(val)
}

// func (m *MMU) Read(addr int) byte {
// 	if addr >= 0x0000 && addr <= 0x3FFF {
// 		if !m.mode {
// 			return m.Rom[addr]
// 		} else {
// 			return m.Rom[0x4000*m.zeroBank+addr]
// 		}

// 	} else if addr >= 0x4000 && addr <= 0x7FFF {
// 		if len(m.Rom) < 1024*1024 {
// 			m.highBank = m.romBank & m.bitMask
// 		}
// 		return m.Rom[0x4000*m.highBank+addr-0x4000]

// 	} else if addr >= 0xA000 && addr <= 0xBFFF {
// 		if m.enableExtRam {
// 			m.highBank = m.romBank & m.bitMask
// 			return m.Rom[0x4000*m.highBank+addr-0x4000]
// 		}
// 		return 0xFF

// 	} else if addr >= 0xC000 && addr <= 0xDFFF {
// 		return m.Ram[addr]

// 	} else if addr >= 0xFF00 && addr <= 0xFFFF {
// 		return m.Ram[addr]
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

// 	} else if addr >= 0x2000 && addr <= 0x3FFF {
// 		if val == 0 {
// 			m.romBank = 1
// 		} else {
// 			m.romBank = val & m.bitMask
// 		}

// 	} else if addr >= 0x4000 && addr <= 0x5FFF {
// 		m.ramBank = m.ramBank | (val & 0x3)

// 	} else if addr >= 0x6000 && addr <= 0x7FFF {
// 		m.mode = (val & 1) == 1

// 	} else if addr >= 0xA000 && addr <= 0xBFFF {
// 		if m.enableExtRam {
// 		}

// 	} else if addr >= 0xC000 && addr <= 0xDFFF {
// 		m.Ram[addr] = byte(val)

// 	} else if addr >= 0xFF00 && addr <= 0xFFFF {
// 		if addr == 0xFF04 {
// 			m.Ram[addr] = byte(0)
// 		} else {
// 			m.Ram[addr] = byte(val)
// 		}
// 	}

// }
