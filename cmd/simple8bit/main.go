package main

import "github.com/ironiridis/gocpu/lib"
import "github.com/ironiridis/gocpu/per/io"
import "github.com/ironiridis/gocpu/per/mem"
import "time"
import "os"

func main() {
	var err error
	c := gocpu.NewCore(gocpu.NewBusUInt8(time.Second * 1))
	per_io := gocpu_io.Make(nil, os.Stdout)
	per_mem := gocpu_mem.Make(65536)
	err = c.RegisterPeripheral(&per_mem)
	if err != nil {
		panic(err)
	}
	err = c.RegisterPeripheral(&per_io)
	if err != nil {
		panic(err)
	}
	panic("eof")
}
