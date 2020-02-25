package utils

import (
	"os/exec"
	"bytes"
	"strings"
	"fmt"
)

// Run : run bash command
func Run(command string)(stdout string,err error){
	cmd := exec.Command("sh")
	var out, stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	cmd.Stdin = strings.NewReader(command)
	err = cmd.Run()
	if err != nil{
		fmt.Println(fmt.Sprint(err) + ":" + stderr.String())
	}
	return out.String(),err
}
