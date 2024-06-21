package GB

// import "fmt"

const (
	IE = 0xFFFF
	IF = 0xFF0F
)

func (GB *GB) GetIEBit(pos int) int { // Get interrupt enable
	return (GB.MMU.Read(IE) >> pos) & 1
}

func (GB *GB) GetIFBit(pos int) int { // Get interrupt flag
	return (GB.MMU.Read(IF) >> pos) & 1
}

func (GB *GB) SetIEBit(op int) { // Set interrupt enable bit
	GB.MMU.Write(GB.MMU.Read(IE)|1<<op, IE)
}

func (GB *GB) ResetIEBit(offset int) { // Reset interrupt enable bit
	k := 0
	for i := 0; i < 8; i++ {
		k = k << 1
		if 7-i != offset {
			k += 1
		}
	}
	GB.MMU.Write(GB.MMU.Read(IE)&k, IE)
}

func (GB *GB) SetIFBit(op int) { // Set interrupt flag bit
	GB.MMU.Write(GB.MMU.Read(IF)|1<<op, IF)
}

func (GB *GB) ResetIFBit(offset int) { // Reset interrupt flag bit
	k := 0
	for i := 0; i < 8; i++ {
		k = k << 1
		if 7-i != offset {
			k += 1
		}
	}
	GB.MMU.Write(GB.MMU.Read(IF)&k, IF)
}

func (GB *GB) InterruptHandler() {
	if GB.CPU.IME {
		if GB.GetIEBit(0) == 1 && GB.GetIFBit(0) == 1 { // VBlank
			GB.CPU.IME = false
			// fmt.Printf("V-Blank Interrupt!\t")
			GB.PUSHN16(GB.CPU.Reg.Read("PC"))
			GB.JPN16(0x40)
			GB.ResetIFBit(0)
			GB.IncrementTimer(20)
			GB.CPU.Awake = true
			GB.CPU.InterruptHit = true
		}
		if GB.GetIEBit(1) == 1 && GB.GetIFBit(1) == 1 { // LCD STAT
			GB.CPU.IME = false
			// fmt.Printf("LCD Interrupt!\n")
			GB.PUSHN16(GB.CPU.Reg.Read("PC"))
			GB.JPN16(0x48)
			GB.ResetIFBit(1)
			GB.IncrementTimer(20)
			GB.CPU.Awake = true
			GB.CPU.InterruptHit = true
		}
		if GB.GetIEBit(2) == 1 && GB.GetIFBit(2) == 1 { // Timer
			GB.CPU.IME = false
			// fmt.Printf("Timer Interrupt!\n")
			GB.PUSHN16(GB.CPU.Reg.Read("PC"))
			GB.JPN16(0x50)
			GB.IncrementTimer(20)
			GB.ResetIFBit(2)
			GB.CPU.Awake = true
			GB.CPU.InterruptHit = true
		}
		if GB.GetIEBit(3) == 1 && GB.GetIFBit(3) == 1 { // Serial
			GB.CPU.IME = false
			// fmt.Printf("Serial Interrupt!\t")
			GB.PUSHN16(GB.CPU.Reg.Read("PC"))
			GB.JPN16(0x58)
			GB.ResetIFBit(3)
			GB.IncrementTimer(20)
			GB.CPU.Awake = true
			GB.CPU.InterruptHit = true

		}
		if GB.GetIEBit(4) == 1 && GB.GetIFBit(4) == 1 { // Joypad
			GB.CPU.IME = false
			// fmt.Printf("Joypad Interrupt!\t")
			GB.PUSHN16(GB.CPU.Reg.Read("PC"))
			GB.JPN16(0x60)
			GB.ResetIFBit(4)
			GB.IncrementTimer(20)
			GB.CPU.Awake = true
			GB.CPU.InterruptHit = true
		}
	} else if !GB.CPU.IME && !GB.CPU.Awake && (GB.MMU.Read(IF)&GB.MMU.Read(IE)) != 0 {
		GB.CPU.Awake = true
	}
}
