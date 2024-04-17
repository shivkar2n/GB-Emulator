package GB

import (
	"fmt"

	"github.com/shivkar2n/GB-Emulator/pkg/CPU"
	"github.com/shivkar2n/GB-Emulator/pkg/MMU"
	"github.com/shivkar2n/GB-Emulator/pkg/PPU"
)

type GB struct {
	SleepMode       bool
	Length, TCycles int
	CPU             *CPU.CPU
	MMU             *MMU.MMU
	PPU             *PPU.PPU
}

func Init() *GB {
	cpu := CPU.Init()
	mmu := MMU.Init()
	ppu := PPU.Init(cpu, mmu)

	return &GB{
		SleepMode: true,
		Length:    0,
		TCycles:   0,
		CPU:       cpu,
		MMU:       mmu,
		PPU:       ppu,
	}
}

func (GB *GB) GetHLVal() int { // Get value pointed by HL register
	return GB.MMU.Read(GB.CPU.Reg.Read("HL"))
}

func (GB *GB) SetHLVal(val int) { // Set value pointed by HL register
	GB.MMU.Write(val, GB.CPU.Reg.Read("HL"))
}

func (GB *GB) GetPCVal() int { // Get value pointed by PC register
	return GB.MMU.Read(GB.CPU.Reg.Read("PC"))
}

func (GB *GB) SetPCVal(val int) { // Set value pointed by PC register
	GB.MMU.Write(val, GB.CPU.Reg.Read("PC"))
}

func (GB *GB) LogSerialIO() {
	if GB.MMU.Read(0xFF02) == 0x81 {
		fmt.Printf("%c", GB.MMU.Read(0xFF01))
		GB.MMU.Write(0x00, 0xFF02)
	}
}

func (GB *GB) StateInfo() { // Get info about state of Console
	a := GB.CPU.Reg.Read("A")
	f := GB.CPU.Reg.Read("F")
	b := GB.CPU.Reg.Read("B")
	c := GB.CPU.Reg.Read("C")
	d := GB.CPU.Reg.Read("D")
	e := GB.CPU.Reg.Read("E")
	h := GB.CPU.Reg.Read("H")
	l := GB.CPU.Reg.Read("L")
	sp := GB.CPU.Reg.Read("SP")
	pc := GB.CPU.Reg.Read("PC")
	pcMem := GB.MMU.Read(GB.CPU.Reg.Read("PC"))
	pcMem1 := GB.MMU.Read(GB.CPU.Reg.Read("PC") + 1)
	pcMem2 := GB.MMU.Read(GB.CPU.Reg.Read("PC") + 2)
	pcMem3 := GB.MMU.Read(GB.CPU.Reg.Read("PC") + 3)

	// fmt.Printf("A: %02X F: %02X B: %02X C: %02X D: %02X E: %02X H: %02X L: %02X SP: %04X PC: 00:%04X (%02X %02X %02X %02X)\n", a, f, b, c, d, e, h, l, sp, pc, pcMem, pcMem1, pcMem2, pcMem3)
	fmt.Printf("A:%02X F:%02X B:%02X C:%02X D:%02X E:%02X H:%02X L:%02X SP:%04X PC:%04X PCMEM:%02X,%02X,%02X,%02X\n", a, f, b, c, d, e, h, l, sp, pc, pcMem, pcMem1, pcMem2, pcMem3)
}

func (GB *GB) Run() {
	for 1 == 1 {
		if GB.SleepMode {
			for GB.CPU.TotalT < GB.CPU.ClkRate/GB.CPU.FrameRate {
				// CPU.StateInfo()
				_, GB.Length, GB.TCycles, GB.SleepMode = GB.ExecuteOpcode()
				GB.IncrementTimer()
				GB.CPU.IncrementCounter(GB.Length)
				// println(opcode)
				if !GB.SleepMode {
					break
				}
				GB.LogSerialIO()
				GB.TCycles, GB.SleepMode = GB.InterruptHandler(GB.SleepMode)
				GB.IncrementTimer()
			}
			GB.CPU.TotalT = GB.CPU.TotalT % (GB.CPU.ClkRate / GB.CPU.FrameRate)

		} else { // Stall CPU until it recieves interrupt (Sleep mode)
			// CPU.StateInfo()
			GB.LogSerialIO()
			GB.TCycles = 4
			GB.IncrementTimer()
			GB.TCycles, GB.SleepMode = GB.InterruptHandler(GB.SleepMode)
			GB.IncrementTimer()

		}
		GB.RenderFrame()
	}

}
