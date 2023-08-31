package vedirect

import (
	"bufio"
	"os"
	"testing"
)

func TestVEDirectParsing(t *testing.T) {
	f, err := os.Open("sample.hex")
	if err != nil {
		t.Fatal(err)
	}
	bufscanner := bufio.NewScanner(f)
	bufscanner.Split(victronBlockSplitter)
	bufscanner.Scan()
	veblock := newVedirectBlock(bufscanner.Text())
	if veblock.Label["FW"].Value != "159" || err != nil {
		t.Fatalf("Failed to parse FW version 159 from sample serial output: %s, %s", err, veblock.Label["FW"].Value)
	}
}
