package gocpu

import "sync"

// A Peripheral on the CPU can be invoked by CPU microcode to invoke some behavior. Memory, registers,
// arithmetic, and others are all peripherals for the CPU to use.
type Peripheral struct {
	State *PeripheralState
	Calls []PeripheralCall
	Init  func(*Peripheral) error
}

// PeripheralState is precisely what it sounds like.
type PeripheralState struct {
	Mutex sync.Mutex
	KV    map[string]interface{}
}

// A PeripheralCall represents some behavior that is invoked by the CPU microcode through a control
// bit. Peripherals cannot (directly) interact; all data must flow through the Bus.
type PeripheralCall struct {
	Symbol string
	Fn     func(*Peripheral, Bus)
}
