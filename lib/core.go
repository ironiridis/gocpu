package gocpu

import "errors"

// ErrorStub is returned because the author is lazy
var ErrorStub = errors.New("Not implemented")

// ErrorPeripheralNoCalls is returned when a Peripheral doesn't have any calls to invoke
var ErrorPeripheralNoCalls = errors.New("Peripheral doesn't define any calls")

// ErrorPeripheralTooMany is returned when the Core doesn't have sufficient control bits
var ErrorPeripheralTooMany = errors.New("Not enough available control bits")

// Core represents the execution environment for a CPU including Peripherals, symbols, and state
type Core struct {
	p []*Peripheral
	cb [32]PeripheralCall
	cbCount int
	cbMap map[string]uint
	b Bus
}

// NewCore returns an instance of Core.
func NewCore(b Bus) (c Core) {
	c.p = make([]*Peripheral, 0, 8)
	c.cbMap = make(map[string]uint)
	c.b = b
	return
}

// RegisterPeripheral adds a Peripheral to Core, reads its symbols and turns them into control
// bits, runs the Init routine (if any)
func (c *Core) RegisterPeripheral(p *Peripheral) (error) {
	if len(p.Calls) == 0 {
		return ErrorPeripheralNoCalls
	}

	if c.cbCount + len(p.Calls) > cap(c.cb) {
		return ErrorPeripheralTooMany
	}

	if p.Init != nil {
		err := p.Init(p)
		if err != nil { return err }
		p.Init = nil // don't run again later
	}

	c.p = append(c.p, p)

	// incomplete implementation...
	return ErrorStub
}
