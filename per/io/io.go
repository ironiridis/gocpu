package gocpu_io

import "io"
import "github.com/ironiridis/gocpu/lib"

// This is a gross hack, but it is enough to test with
func Make(r io.Reader, w io.Writer) (p *gocpu.Peripheral) {
	p.State = &gocpu.PeripheralState{
	}
	p.Calls = []gocpu.PeripheralCall{{
		Symbol: "putc",
		Fn: func(p *gocpu.Peripheral, b gocpu.Bus) {
			n, err := w.Write([]byte{byte(b.Read())})
			if err != nil { panic(err) }
			if n != 1 { panic("putc didn't write 1 byte") }
		}}}
	
	return
}
