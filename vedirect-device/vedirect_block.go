package vedirect

import (
	"encoding/json"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

type VedirectBlock struct {
	Label map[string]VedirectField `json:"label"`
}

type VedirectField struct {
	Unit        string `json:"unit"`
	Description string `json:"description"`
	Value       any    `json:"value"`
}

func newVedirectBlock(text string) *VedirectBlock {
	block := new(VedirectBlock)
	vedirectRefData := vedirectRefData()
	fields := strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n")
	block.Label = make(map[string]VedirectField)
	for _, raw_field := range fields {
		raw_fields := strings.Fields(raw_field)
		if len(raw_fields) > 0 {
			block.Label[raw_fields[0]] = VedirectField{
				Value:       raw_fields[1],
				Unit:        vedirectRefData[raw_fields[0]][0],
				Description: vedirectRefData[raw_fields[0]][1],
			}
		}
	}

	log.WithField("VEDirect Labels", block.GetLabels()).Debug("Found labels in block")
	return block
}

func (veblock VedirectBlock) GetLabels() []string {
	field_labels := make([]string, len(veblock.Label))
	i := 0
	for k := range veblock.Label {
		field_labels[i] = k
		i++
	}
	return field_labels
}

func (veblock VedirectBlock) ToJson() []byte {
	json, err := json.MarshalIndent(veblock, "", "    ")
	if err != nil {
		log.WithFields(log.Fields{"veblock": veblock}).Panicf("Failed to marshal JSON from input: %s", err)
	}
	return json
}

func (vefield VedirectField) AsFloat() float64 {
	if vefield.Value.(string) == "ON" {
		return float64(1)
	}
	if vefield.Value.(string) == "OFF" {
		return float64(0)
	}
	f, err := strconv.ParseFloat(vefield.Value.(string), 64)
	if err != nil {
		log.Fatal(err)
	}
	return f
}
