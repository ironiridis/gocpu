package gocpu_mem

import "github.com/ironiridis/gocpu/lib"

func Make(capacity uint) (p gocpu.Peripheral) {
	p.State = &gocpu.PeripheralState{
		KV: make(map[string]interface{}),
	}
	p.State.KV["a"] = uint(0)
	p.State.KV["d"] = make([]uint, capacity, capacity)
	p.Calls = []gocpu.PeripheralCall{
		{
			Symbol: "memaddr",
			Fn: func(p *gocpu.Peripheral, b gocpu.Bus) {
				p.State.Mutex.Lock()
				p.State.KV["a"] = b.Read()
				p.State.Mutex.Unlock()
			},
		},
		{
			Symbol: "memput",
			Fn: func(p *gocpu.Peripheral, b gocpu.Bus) {
				p.State.Mutex.Lock()
				p.State.KV["d"].([]uint)[p.State.KV["a"].(uint)] = b.Read()
				p.State.Mutex.Unlock()
			},
		},
		{
			Symbol: "memget",
			Fn: func(p *gocpu.Peripheral, b gocpu.Bus) {
				p.State.Mutex.Lock()
				b.Write(p.State.KV["d"].([]uint)[p.State.KV["a"].(uint)])
				p.State.Mutex.Unlock()
			},
		},
	}

	return
}
