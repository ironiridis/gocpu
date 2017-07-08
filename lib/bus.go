package gocpu

import "time"
import "errors"

type BusUInt32 struct {
	v chan uint32
	expire *time.Timer
	expireTime time.Duration
}

type BusUInt8 struct {
	v chan uint8
	expire *time.Timer
	expireTime time.Duration
}

func NewBusUInt32(expireTime time.Duration) (b BusUInt32) {
	b.v = make(chan uint32)
	b.expireTime = expireTime
	b.expire = time.NewTimer(expireTime)
	b.expire.Stop() // we actually don't require a timeout immediately
}

func (b BusUInt32) Read() (v uint) {
	b.expire.Reset(b.expireTime)
	select {
	case v = <-b.v:
		b.expire.Stop()
		return
	case <-b.expire.C:
		panic(errors.New("Bus read timeout"))
	}
}

func (b BusUInt32) Write(v uint) {
	b.expire.Reset(b.expireTime)
	select {
	case b.v <- v:
		b.expire.Stop()
		return
	case <-b.expire.C:
		panic(errors.New("Bus write timeout"))
	}
}

