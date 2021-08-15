package serial

import (
	"github.com/tarm/serial"
)

type Aspect string

const (
	DefaultBaud = 9600

	Aspect_24   Aspect = "0101U\n"                   // 2.4:1
	Aspect_4x3  Aspect = "0101D\n"                   // 4:3
	Aspect_16x9 Aspect = "0101S\n"                   // 16:9
	Aspect_185  Aspect = "0101S;W9;0101U;W1;0101S\n" // 1.85:1
	Aspect_22   Aspect = "0101U;W9;0101D;W2;0101S\n" // 2.2:1
	Aspect_143  Aspect = "0101D;W9;0101U;W1;0101S\n" // 1.43:1
)

type RTSSession struct {
	SerialPort string
	openPort   *serial.Port
}

func (r *RTSSession) NewSession() error {
	c := &serial.Config{Name: r.SerialPort, Baud: DefaultBaud}
	s, err := serial.OpenPort(c)
	if err != nil {
		return err
	}
	r.openPort = s
	return nil
}

func (r *RTSSession) Send(command string) error {
	if r.openPort == nil {
		err := r.NewSession()
		if err != nil {
			return err
		}
	}

	_, err := r.openPort.Write([]byte(command))
	if err != nil {
		return err
	}

	return nil
}

func (r *RTSSession) SetAspect(aspect Aspect) error {
	return r.Send(string(aspect))
}
