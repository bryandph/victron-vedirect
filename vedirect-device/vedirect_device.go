package vedirect

import (
	"bufio"
	"bytes"

	log "github.com/sirupsen/logrus"

	"go.bug.st/serial"
)

type VEDirectDevice struct {
	serialport serial.Port
	bufscanner bufio.Scanner
}

func NewVEDirectDevice(portname string) (*VEDirectDevice, error) {
	device := new(VEDirectDevice)

	// initialize serial port
	mode := &serial.Mode{
		BaudRate: 19200,
		Parity:   serial.NoParity,
		StopBits: serial.OneStopBit,
		DataBits: 8,
	}
	port, err := serial.Open(portname, mode)
	if err != nil {
		return &VEDirectDevice{}, err
	}
	device.serialport = port

	// initialize scanner with custom frame splitter
	device.bufscanner = *bufio.NewScanner(port)
	device.bufscanner.Split(victronBlockSplitter)

	return device, nil
}

// victronBlockSplitter handles a buffer and returns checksum ensured block of VEDirect TEXT data.
func victronBlockSplitter(data []byte, atEOF bool) (advance int, token []byte, err error) {
	for i := 0; i < len(data)-7; i++ {
		if bytes.Equal(data[i:i+3], []byte("sum")) {
			checksum := func(arr []byte) int {
				sum := 0
				for idx := 0; idx < len(arr); idx++ {
					sum += int(arr[idx])
				}
				return (sum % 256)
			}(data[:i+7])
			if checksum == 0 {
				log.Debugln("Checksum validated on block")
				return i + 7, data[:i-5], nil
			} else {
				log.Infoln("Checksum failed on block")
				return i + 7, nil, nil
			}
		}
	}
	if !atEOF {
		return 0, nil, nil
	}
	return 0, data, bufio.ErrFinalToken
}

func (vedev VEDirectDevice) GetBlock() VedirectBlock {
	vedev.bufscanner.Scan()
	text := vedev.bufscanner.Text()
	return *newVedirectBlock(text)
}
