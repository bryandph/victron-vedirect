# Victron VEDirect CLI

This project implements a CLI to get status information from a [Victron VE Direct](https://www.victronenergy.com/upload/documents/VE.Direct-Protocol-3.33.pdf) device over a serial interface.

#### Usage

```
NAME:
   vedirect-status - Get serial output from an attached Victron VEDirect device and output in JSON format

USAGE:
   vedirect-status [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --serialport COM3, -s COM3  serial port device to use eg /dev/ttyUSB0, COM3
   --verbosity, -v             (default: false)
   --help, -h                  show help
```

Example Output ([BlueSolar MPPT 75/10](https://www.victronenergy.com/solar-charge-controllers/mppt7510))
```
{
    "label": {
        "CS": {
            "unit": "n/a",
            "description": "State of operation",
            "value": "0"
        },
        "ERR": {
            "unit": "n/a",
            "description": "Error code",
            "value": "0"
        },
        "FW": {
            "unit": "n/a",
            "description": "Firmware version (16 bit)",
            "value": "161"
        },
        "H19": {
            "unit": "0.01kWh",
            "description": "Yield total (user resettable counter)",
            "value": "11"
        },
        "H20": {
            "unit": "0.01kWh",
            "description": "Yield today",
            "value": "0"
        },
        "H21": {
            "unit": "W",
            "description": "Maximum power today",
            "value": "0"
        },
        "H22": {
            "unit": "0.01kWh",
            "description": "Yield yesterday",
            "value": "4"
        },
        "H23": {
            "unit": "W",
            "description": "Maximum power yesterday",
            "value": "19"
        },
        "HSDS": {
            "unit": "n/a",
            "description": "Day sequence number (0..364)",
            "value": "3"
        },
        "I": {
            "unit": "mA",
            "description": "Main or channel 1 battery current",
            "value": "-310"
        },
        "IL": {
            "unit": "mA",
            "description": "Load current",
            "value": "300"
        },
        "LOAD": {
            "unit": "n/a",
            "description": "Load output state (ON/OFF) ",
            "value": "ON"
        },
        "MPPT": {
            "unit": "n/a",
            "description": "Tracker operation mode",
            "value": "0"
        },
        "PID": {
            "unit": "n/a",
            "description": "Product ID",
            "value": "0xA04C"
        },
        "PPV": {
            "unit": "W",
            "description": "Panel power",
            "value": "0"
        },
        "SER#": {
            "unit": "n/a",
            "description": "Serial number",
            "value": "HQ2151JJCZR"
        },
        "V": {
            "unit": "mV",
            "description": "Main or channel 1 (battery) voltage",
            "value": "13830"
        },
        "VPV": {
            "unit": "mV",
            "description": "Panel voltage",
            "value": "80"
        }
    }
}
```