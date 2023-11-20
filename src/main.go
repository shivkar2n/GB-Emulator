package main

import (
	"github.com/shivkar2n/GB-Emulator/CPU"
	"github.com/shivkar2n/GB-Emulator/Display"
	"github.com/shivkar2n/GB-Emulator/Logger"
	"github.com/shivkar2n/GB-Emulator/MMU"
)

func main() {
	log := Logger.DebugLogger()
	stateLog := Logger.StateLogger()

	d := Display.Init()
	m := MMU.Init()
	c := CPU.Init(m)

	defer d.Free()

	// var exec bool
	// exec = false

	for 1 == 1 {
		for c.TotalM < c.ClkRate/c.FrameRate {

			c.ExecuteOpcode(stateLog, log)
			c.IncrTimer()
			c.LogSerialIO()
			c.T = 0
			c.M = 0
			c.InterruptHandler()
		}
		d.RenderFrame(c)
		c.TotalM = c.TotalM % (c.ClkRate / c.FrameRate)
	}

}
