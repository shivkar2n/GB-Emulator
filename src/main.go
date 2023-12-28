package main

import (
	"github.com/shivkar2n/GB-Emulator/CPU"
	"github.com/shivkar2n/GB-Emulator/Display"
	"github.com/shivkar2n/GB-Emulator/MMU"
)

func main() {
	d := Display.Init()
	m := MMU.Init()
	c := CPU.Init(m)

	defer d.Free()

	// var exec bool
	// exec = false

	for 1 == 1 {
		if !c.StopExec {
			for c.TotalT < c.ClkRate/c.FrameRate {
				// c.StateInfo()
				c.ExecuteOpcode()
				c.IncrTimer()
				if c.StopExec {
					break
				}
				c.LogSerialIO()
				c.InterruptHandler()
				c.IncrTimer()
			}
			c.TotalT = c.TotalT % (c.ClkRate / c.FrameRate)

		} else {
			// c.StateInfo()
			c.LogSerialIO()
			c.T = 4
			c.IncrTimer()
			c.InterruptHandler()
			c.IncrTimer()

		}
		d.RenderFrame(c)
	}

}
