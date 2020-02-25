package utils

import "testing"

func TestGetLocalIPv4(t *testing.T) {
	ip := GetLocalIPv4()
	t.Log("ip:",ip)
}

func TestGetLocalRandomPort(t *testing.T) {
	port := GetLocalRandomPort()
	t.Log(port)
}
