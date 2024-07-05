package GB

import (
	"fmt"

	"github.com/shivkar2n/GB-Emulator/pkg/CPU"
	"github.com/shivkar2n/GB-Emulator/pkg/MMU"
	"github.com/shivkar2n/GB-Emulator/pkg/PPU"
)

const (
	JOYP           = 0xFF00
	SERIAL         = 0xFF01
	SC             = 0xFF02
	DIV            = 0xFF04
	TIMA           = 0xFF05
	TMA            = 0xFF06
	TAC            = 0xFF07
	FRAME_RATE     = 60
	CPU_CLOCK_RATE = 4194304
)

type GB struct {
	CPU *CPU.CPU
	MMU *MMU.MMU
	PPU *PPU.PPU
}

func Init() *GB {
	cpu := CPU.Init()
	mmu := MMU.Init()
	ppu := PPU.Init(cpu, mmu)

	return &GB{
		CPU: cpu,
		MMU: mmu,
		PPU: ppu,
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
	if GB.MMU.Read(SC) == 0x81 {
		fmt.Printf("%c", GB.MMU.Read(SERIAL))
		GB.MMU.Write(0x00, SC)
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
	// ly := GB.MMU.Read(0xFF44)
	// iE := GB.MMU.Read(IE)
	// iF := GB.MMU.Read(IF)
	// cdn := GB.MMU.Read(0xFFA6)

	// fmt.Printf("A:%02X B:%02X C:%02X D:%02X E:%02X H:%02X L:%02X SP:%04X PC:%04X PCMEM:%02X,%02X,%02X,%02X LY:%02X IF:%02X IE:%02X %s\n", a, b, c, d, e, h, l, sp, pc, pcMem, pcMem1, pcMem2, pcMem3, ly, iE, iF, opcode) fmt.Printf("A:%02X F:%02X B:%02X C:%02X D:%02X E:%02X H:%02X L:%02X SP:%04X PC:%04X PCMEM:%02X,%02X,%02X,%02X\n", a, f, b, c, d, e, h, l, sp, pc, pcMem, pcMem1, pcMem2, pcMem3)
	fmt.Printf("A:%02X F:%02X B:%02X C:%02X D:%02X E:%02X H:%02X L:%02X SP:%04X PC:%04X PCMEM:%02X,%02X,%02X,%02X\n", a, f, b, c, d, e, h, l, sp, pc, pcMem, pcMem1, pcMem2, pcMem3)
}

func (GB *GB) Run() {
	for 1 == 1 {
		if GB.PPU.Frame < FRAME_RATE {
			// opcode := ""
			for GB.CPU.TotalCycles < CPU_CLOCK_RATE/(FRAME_RATE*NO_SCANLINES) {
				if GB.CPU.Awake {
					GB.Fetch()
					GB.InterruptHandler()
					if GB.CPU.InterruptHit { // Reset Fetch-Execute cycle on interrupt
						GB.CPU.InterruptHit = false
						continue
					}
					GB.Execute()
					// GB.StateInfo()

				} else { // Stall CPU until it recieves interrupt (Sleep mode)
					GB.IncrementTimer(4)
					GB.InterruptHandler()
				}
			}
			GB.CPU.TotalCycles = GB.CPU.TotalCycles % (CPU_CLOCK_RATE / (FRAME_RATE * NO_SCANLINES))
			GB.LogSerialIO()
			GB.RenderScanline()

		} else {
			GB.FrameDelay()
		}
	}

}
