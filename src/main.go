package main

import (
	"github.com/shivkar2n/GB-Emulator/CPU"
	"github.com/shivkar2n/GB-Emulator/Display"
	"github.com/shivkar2n/GB-Emulator/MMU"
)

func main() {
	display := Display.Init()
	mmu := MMU.Init()
	cpu := CPU.Init(mmu)

	defer display.Free()

	var exec bool
	var length, time int
	// var opcode string
	exec = true

	for 1 == 1 {
		if exec {
			for cpu.TotalT < cpu.ClkRate/cpu.FrameRate {
				// cpu.StateInfo()
				_, length, time, exec = cpu.ExecuteOpcode()
				cpu.IncrementTimer(time)
				cpu.IncrementCounter(length)
				// println(opcode)
				if !exec {
					break
				}
				cpu.LogSerialIO()
				time, exec = cpu.InterruptHandler(exec)
				cpu.IncrementTimer(time)
			}
			cpu.TotalT = cpu.TotalT % (cpu.ClkRate / cpu.FrameRate)

		} else {
			// cpu.StateInfo()
			cpu.LogSerialIO()
			cpu.IncrementTimer(4)
			time, exec = cpu.InterruptHandler(exec)
			cpu.IncrementTimer(time)

		}
		display.RenderFrame(cpu)
	}

}
