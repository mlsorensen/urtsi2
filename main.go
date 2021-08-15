package main

import (
	. "github.com/mlsorensen/urtsi2/pkg/serial"
	"time"
)

func main() {
	session := RTSSession{SerialPort: "/dev/ttyUSB2"}

	err := session.SetAspect(Aspect_24)
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 10)

	err = session.SetAspect(Aspect_16x9)
	if err != nil {
		panic(err)
	}
}
