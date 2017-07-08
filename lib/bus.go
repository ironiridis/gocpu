package gocpu

import "time"
import "errors"

// Bus represents a channel (with a particular width) for data to flow between Peripherals
type Bus interface {
	Read() (uint)
	Write(uint)
}

// BusUInt32 is a 32-bit Bus
type BusUInt32 struct {
	v chan uint32
	expire *time.Timer
	expireTime time.Duration
}

// BusUInt8 is an 8-bit Bus
type BusUInt8 struct {
	v chan uint8
	expire *time.Timer
	expireTime time.Duration
}

// NewBusUInt32 initializes a 32-bit Bus
func NewBusUInt32(expireTime time.Duration) (b BusUInt32) {
	b.v = make(chan uint32)
	b.expireTime = expireTime
	b.expire = time.NewTimer(expireTime)
	b.expire.Stop() // we actually don't require a timeout immediately
	return
}

func (b BusUInt32) Read() (v uint) {
	b.expire.Reset(b.expireTime)
	select {
	case n := <-b.v:
		b.expire.Stop()
		v = uint(n)
	case <-b.expire.C:
		panic(errors.New("Bus read timeout"))
	}
	return
}

func (b BusUInt32) Write(v uint) {
	b.expire.Reset(b.expireTime)
	select {
	case b.v <- uint32(v):
		b.expire.Stop()
	case <-b.expire.C:
		panic(errors.New("Bus write timeout"))
	}
}

// NewBusUInt8 initializes an 8-bit Bus
func NewBusUInt8(expireTime time.Duration) (b BusUInt8) {
	b.v = make(chan uint8)
	b.expireTime = expireTime
	b.expire = time.NewTimer(expireTime)
	b.expire.Stop() // we actually don't require a timeout immediately
	return
}

func (b BusUInt8) Read() (v uint) {
	b.expire.Reset(b.expireTime)
	select {
	case n := <-b.v:
		b.expire.Stop()
		v = uint(n)
	case <-b.expire.C:
		panic(errors.New("Bus read timeout"))
	}
	return
}

func (b BusUInt8) Write(v uint) {
	b.expire.Reset(b.expireTime)
	select {
	case b.v <- uint8(v):
		b.expire.Stop()
	case <-b.expire.C:
		panic(errors.New("Bus write timeout"))
	}
}

