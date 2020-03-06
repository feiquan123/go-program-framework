package utils

import (
	"testing"
)

func TestIsFileExists(t *testing.T) {
	filename := "/tmp/go-program-famework.log"
	exist, err := IsFileExists(filename)
	if err != nil {
		t.Error(err)
	}
	t.Log(exist)
}

func TestNewFile(t *testing.T) {
	filename := "/tmp/go-program-famework.log"
	wFile := NewFile(filename, "id,source\n")
	wFile.WriteString("0,dida\n")
	exist, err := IsFileExists(filename)
	if err != nil {
		t.Error(err)
	}
	t.Log(exist)
}

// func TestFileSplitAsDate(t *testing.T) {
// 	filename := "/tmp/go-program-famework.log"
// 	err := FileSplitAsDate(filename, TimeFormat5, true)
// 	if err != nil {
// 		t.Log(err)
// 	}
// }
