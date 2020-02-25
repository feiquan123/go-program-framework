package utils

import (
	"testing"
)

func TestIsFileExists(t *testing.T) {
	filename := "/var/log/go-program-famework.log"
	exist, err := IsFileExists(filename)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(exist)
	}
}

func TestFileSplitAsDate(t *testing.T) {
	filename := "/var/log/go-program-famework.log"
	err := FileSplitAsDate(filename, TimeFormat5, true)
	if err != nil {
		t.Log(err)
	}
}
