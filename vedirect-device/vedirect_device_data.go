package vedirect

// vedirectRefData gets all known VEDirect fields and returns a map of labels to an array of units and descriptions. CSV is sourced from https://www.victronenergy.com/upload/documents/VE.Direct-Protocol-3.33.pdf
func vedirectRefData() map[string][]string {
	var data = map[string][]string{
		"V":        {"mV", "Main or channel 1 (battery) voltage"},
		"V2":       {"mV", "Channel 2 (battery) voltage"},
		"V3":       {"mV", "Channel 3 (battery) voltage"},
		"VS":       {"mV", "Auxiliary (starter) voltage"},
		"VM":       {"mV", "Mid-point voltage of the battery bank"},
		"DM":       {"%", "Mid-point deviation of the battery bank"},
		"VPV":      {"mV", "Panel voltage"},
		"PPV":      {"W", "Panel power"},
		"I":        {"mA", "Main or channel 1 battery current"},
		"I2":       {"mA", "Channel 2 battery current"},
		"I3":       {"mA", "Channel 3 battery current"},
		"IL":       {"mA", "Load current"},
		"LOAD":     {"n/a", "Load output state (ON/OFF) "},
		"T":        {"°C", "Battery temperature"},
		"P":        {"W", "Instantaneous power"},
		"CE":       {"mAh", "Consumed Amp Hours"},
		"SOC":      {"%", "State-of-charge"},
		"TTG":      {"Minutes", "Time-to-go"},
		"Alarm":    {"n/a", "Alarm condition active"},
		"Relay":    {"n/a", "Relay state"},
		"AR":       {"n/a", "Alarm reason"},
		"OR":       {"n/a", "Off reason"},
		"H1":       {"mAh", "Depth of the deepest discharge"},
		"H2":       {"mAh", "Depth of the last discharge"},
		"H3":       {"mAh", "Depth of the average discharge"},
		"H4":       {"Number", "of charge cycles"},
		"H5":       {"Number", "of full discharges"},
		"H6":       {"mAh", "Cumulative Amp Hours drawn"},
		"H7":       {"mV", "Minimum main (battery) voltage"},
		"H8":       {"mV", "Maximum main (battery) voltage"},
		"H9":       {"Seconds", "Number of seconds since last full charge"},
		"H10":      {"n/a", "Number of automatic synchronizations"},
		"H11":      {"n/a", "Number of low main voltage alarms"},
		"H12":      {"n/a", "Number of high main voltage alarms"},
		"H13":      {"n/a", "Number of low auxiliary voltage alarms"},
		"H14":      {"n/a", "Number of high auxiliary voltage alarms"},
		"H15":      {"mV3", "Minimum auxiliary (battery) voltage"},
		"H16":      {"mV3", "Maximum auxiliary (battery) voltage"},
		"H17":      {"0.01kWh", "Amount of discharged energy (BMV)"},
		"H18":      {"0.01kWh", "Amount of charged energy (BMV)"},
		"H19":      {"0.01kWh", "Yield total (user resettable counter)"},
		"H20":      {"0.01kWh", "Yield today"},
		"H21":      {"W", "Maximum power today"},
		"H22":      {"0.01kWh", "Yield yesterday"},
		"H23":      {"W", "Maximum power yesterday"},
		"ERR":      {"n/a", "Error code"},
		"CS":       {"n/a", "State of operation"},
		"BMV":      {"n/a", "Model description (deprecated)"},
		"FW":       {"n/a", "Firmware version (16 bit)"},
		"FWE":      {"n/a", "Firmware version (24 bit)"},
		"PID":      {"n/a", "Product ID"},
		"SER#":     {"n/a", "Serial number"},
		"HSDS":     {"n/a", "Day sequence number (0..364)"},
		"MODE":     {"n/a", "Device mode"},
		"AC_OUT_V": {"0.01V", "AC output voltage"},
		"AC_OUT_I": {"0.1A", "AC output current"},
		"AC_OUT_S": {"VA", "AC output apparent power"},
		"WARN":     {"n/a", "Warning reason"},
		"MPPT":     {"n/a", "Tracker operation mode"},
		"MON":      {"n/a", "DC monitor mode"},
		"DC_IN_V":  {"0.01V", "DC input voltage"},
		"DC_IN_I":  {"0.1A", "DC input current"},
		"DC_IN_P":  {"1W", "DC input power"},
	}
	return data
}