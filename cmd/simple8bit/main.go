package main

import "github.com/ironiridis/gocpu/lib"
import "time"

func main() {
	c := gocpu.NewCore(gocpu.NewBusUInt8(time.Second * 1))
	
	p := gocpu.Peripheral{
		State: &gocpu.PeripheralState{},
		Calls: make([]gocpu.PeripheralCall, 1),
	}
	
	p.State.KV = make(map[string]interface{})
	p.Init = nil // no initialization needed
	
	err := c.RegisterPeripheral(&p)
	if err != nil { panic(err) }
	panic("eof")
}
