package gocpu

type PeripheralCall func()
func (c *Core) RegisterControlBit(f *PeripheralCall) error
