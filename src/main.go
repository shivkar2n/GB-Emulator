package main

import (
	CPU "github.com/shivkar2n/GB-Emulator/CPU"
	Logger "github.com/shivkar2n/GB-Emulator/Logger"
)

func main() {
	log := Logger.DebugLogger()
	stateLog := Logger.StateLogger()

	c := CPU.InitCPU()
	c.LoadROM()

	// var exec bool
	// exec = false

	for 1 == 1 {

		for c.TotalM < c.ClkRate/c.FrameRate {
			c.ExecuteOpcode(stateLog, log)
			c.IncrTimer() // Increment timer registers
			c.LogSerialIO()
			c.T = 0
			c.M = 0
			c.InterruptHandler(stateLog)
			c.IncrTimer() // Increment timer registers
		}
		// RenderFrame()
		c.TotalM = c.TotalM % (c.ClkRate / c.FrameRate)
	}

}
