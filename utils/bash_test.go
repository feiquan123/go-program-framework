package utils

import (
	"testing"
)

func TestRun(t *testing.T) {
	commad := "iwconfig"
	stdout, err := Run(commad)
	if err != nil {
		t.Error(err)
	}
	t.Log(stdout)
}
