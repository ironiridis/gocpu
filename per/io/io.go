package gocpu_io

import "io"
import "fmt"
import "github.com/ironiridis/gocpu/lib"

// This is a gross hack, but it is enough to test with
func Make(r io.Reader, w io.Writer) (p gocpu.Peripheral) {
	p.State = &gocpu.PeripheralState{}
	p.Calls = []gocpu.PeripheralCall{
		{
			Symbol: "putc",
			Fn: func(p *gocpu.Peripheral, b gocpu.Bus) {
				n, err := w.Write([]byte{byte(b.Read())})
				if err != nil {
					panic(err)
				}
				if n != 1 {
					panic("putc didn't write 1 byte")
				}
			},
		}, {
			Symbol: "putfmt02x",
			Fn: func(p *gocpu.Peripheral, b gocpu.Bus) {
				n, err := fmt.Fprintf(w, "%02x", b.Read())
				if err != nil {
					panic(err)
				}
				if n < 2 {
					panic("putfmt02x didn't write at least 2 bytes")
				}
			},
		}, {
			Symbol: "putfmtu",
			Fn: func(p *gocpu.Peripheral, b gocpu.Bus) {
				n, err := fmt.Fprintf(w, "%u", b.Read())
				if err != nil {
					panic(err)
				}
				if n < 1 {
					panic("putfmtu didn't write at least 1 byte")
				}
			},
		},
	}

	return
}
