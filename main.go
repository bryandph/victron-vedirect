package main

import (
	"fmt"
	"os"

	vedirect_device "github.com/bryandph/victron-vedirect/vedirect-device"

	"go.bug.st/serial/enumerator"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func getFTDIPort() string {
	ports, err := enumerator.GetDetailedPortsList()
	if (err != nil) || (len(ports) == 0) {
		log.Fatal("No serial ports found!")
	}
	var portname string
	for _, port := range ports {
		if port.IsUSB {
			if port.VID == "0403" {
				portname = port.Name
			}
		}
	}
	if portname != "" {
		log.WithFields(log.Fields{"serialport": portname}).Debug("Using found FTDI device")
	}
	return portname
}

func getVEDirectStatus(portname string) string {
	vedevice, err := vedirect_device.NewVEDirectDevice(portname)
	if err != nil {
		log.Fatal(err)
	}
	return string(vedevice.GetBlock().ToJson())
}

func init() {
	log.SetLevel(log.WarnLevel)
}

func main() {

	app := &cli.App{
		Name:  "vedirect-status",
		Usage: "Get serial output from an attached Victron VEDirect device and output in JSON format",
		Action: func(cCtx *cli.Context) error {
			if cCtx.Bool("verbosity") {
				log.SetLevel(log.DebugLevel)
			}
			if cCtx.String("serialport") == "" {
				fmt.Println(getVEDirectStatus(getFTDIPort()))
			} else {
				fmt.Println(getVEDirectStatus(cCtx.String("serialport")))
			}
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "serialport",
				Aliases: []string{"s"},
				Value:   "",
				Usage:   "serial port device to use eg /dev/ttyUSB0, `COM3`",
			},
			&cli.BoolFlag{
				Name:    "verbosity",
				Aliases: []string{"v"},
				Value:   false,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
