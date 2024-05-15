package GB

import (
	"fmt"

	"github.com/shivkar2n/GB-Emulator/pkg/CPU"
	"github.com/shivkar2n/GB-Emulator/pkg/MMU"
	"github.com/shivkar2n/GB-Emulator/pkg/PPU"
)

type GB struct {
	AwakeMode       bool
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
		AwakeMode: true,
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

func (GB *GB) StateInfo(opcode string) { // Get info about state of Console
	a := GB.CPU.Reg.Read("A")
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
	ly := GB.MMU.Read(0xFF44)
	cdn := GB.MMU.Read(0xFFA6)

	// fmt.Printf("A: %02X F: %02X B: %02X C: %02X D: %02X E: %02X H: %02X L: %02X SP: %04X PC: 00:%04X (%02X %02X %02X %02X)\n", a, f, b, c, d, e, h, l, sp, pc, pcMem, pcMem1, pcMem2, pcMem3)
	fmt.Printf("A:%02X B:%02X C:%02X D:%02X E:%02X H:%02X L:%02X SP:%04X PC:%04X PCMEM:%02X,%02X,%02X,%02X LY:%02X COUNTDOWN: %02X %s\n", a, b, c, d, e, h, l, sp, pc, pcMem, pcMem1, pcMem2, pcMem3, ly, cdn, opcode)
	// fmt.Printf("A:%02X B:%02X C:%02X D:%02X E:%02X H:%02X L:%02X SP:%04X, PC:%04X PCMEM:%02X,%02X,%02X,%02X %s", a, b, c, d, e, h, l, sp, pc, pcMem, pcMem1, pcMem2, pcMem3, opcode)
	// fmt.Printf(" FLAGS: Z:%d H:%d C:%d N:%d\n", GB.CPU.GetFlag("Z"), GB.CPU.GetFlag("H"), GB.CPU.GetFlag("C"), GB.CPU.GetFlag("N"))
	// fmt.Printf("%d %s\n", ly, opcode)
}

func (GB *GB) Run() {
	for 1 == 1 {
		if GB.PPU.Frame < 60 {
			if GB.AwakeMode {
				for GB.CPU.TotalT < GB.CPU.ClkRate/(GB.PPU.FrameRate*NO_REAL_SCANLINES) {
					// var opcode string
					_, GB.Length, GB.TCycles, GB.AwakeMode = GB.ExecuteOpcode()
					// GB.StateInfo(opcode)
					GB.IncrementTimer()
					GB.CPU.IncrementCounter(GB.Length)
					if !GB.AwakeMode {
						break
					}
					// GB.LogSerialIO()
					GB.TCycles, GB.AwakeMode = GB.InterruptHandler(GB.AwakeMode)
					GB.IncrementTimer()
				}
				GB.CPU.TotalT = GB.CPU.TotalT % (GB.CPU.ClkRate / (GB.PPU.FrameRate * NO_REAL_SCANLINES))

			} else { // Stall CPU until it recieves interrupt (Sleep mode)
				// GB.StateInfo("SLEEP MODE!!")
				// GB.LogSerialIO()
				GB.TCycles = 4
				GB.IncrementTimer()
				GB.TCycles, GB.AwakeMode = GB.InterruptHandler(GB.AwakeMode)
				GB.IncrementTimer()

			}
			GB.RenderScanline()

		} else {
			GB.RenderFrame()
		}
	}

}
