package main

import (
	"os"

	"labs.lesiw.io/ops/goapp"
	"lesiw.io/ops"
)

type Ops struct{ goapp.Ops }

func main() {
	goapp.Name = "http2https"
	if len(os.Args) < 2 {
		os.Args = append(os.Args, "build")
	}
	ops.Handle(Ops{})
}