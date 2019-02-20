package settings

import (
	"testing"
)

func TestSaveOpenSettings(t *testing.T) {
	//arrange
	settings := Settings{"http://test.com", "username", "password", "query"}
	//act
	settings.Save()
	settings2 := Settings{}
	settings2.Open()
	//assert

	if settings2.URL != "http://test.com" {
		t.Error("Did not save URL")
	}

	if settings2.SvcNowUser != "username" {
		t.Error("Did not save username")
	}

	if settings2.SvcNowPass != "password" {
		t.Error("Did not save password")
	}

	if settings2.Query1 != "query" {
		t.Error("Did not save query")
	}

}
