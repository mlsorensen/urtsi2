package main

import (
	. "github.com/mlsorensen/urtsi2/pkg/serial"
	"time"
)

func main() {
	session := RTSSession{SerialPort: "/dev/ttyUSB2"}

	// using as basic Somfy buttons
	err := session.Send(CommandClose)
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 10)

	// using as theater screen control where "open" and "close" correlate to aspect ratio.
	//The difference is merely semantic.
	err = session.SetAspect(Aspect_24)
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 10)

	err = session.SetAspect(Aspect_4x3)
	if err != nil {
		panic(err)
	}
}
