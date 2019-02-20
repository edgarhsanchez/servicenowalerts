package settings

import (
	"encoding/gob"
	"os"
)

// Settings represent all the service settings used for performing queries and opening a web page
type Settings struct {
	URL        string
	SvcNowUser string
	SvcNowPass string
	Query1     string
}

// Save saves the settings in the settings.dat file
func (s *Settings) Save() {
	file, err := os.Create("settings.dat")
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(s)
	}
	file.Close()
}

// Open retrieves the settings from the settings.dat file
func (s *Settings) Open() {
	file, err := os.Open("settings.dat")
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(s)
	}
	file.Close()
}
