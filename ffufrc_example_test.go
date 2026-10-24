package main

import (
	"os"
	"testing"

	"github.com/ffuf/ffuf/v2/pkg/ffuf"
	"github.com/pelletier/go-toml"
)

// TestFfufrcExampleKeysAreRead decodes ffufrc.example in strict mode, which
// fails on any key that matches no ConfigOptions field. go-toml otherwise
// ignores such keys silently, so a misspelled key in the example would just do
// nothing.
func TestFfufrcExampleKeysAreRead(t *testing.T) {
	f, err := os.Open("ffufrc.example")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	if err := toml.NewDecoder(f).Strict(true).Decode(ffuf.NewConfigOptions()); err != nil {
		t.Errorf("ffufrc.example: %s", err)
	}
}
