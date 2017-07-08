package main

import "github.com/ironiridis/gocpu/lib"
import "github.com/ironiridis/gocpu/per/io"
import "time"
import "os"

func main() {
	c := gocpu.NewCore(gocpu.NewBusUInt8(time.Second * 1))
	per_io := gocpu_io.Make(nil, os.Stdout)
	err := c.RegisterPeripheral(&per_io)
	if err != nil { panic(err) }
	panic("eof")
}
