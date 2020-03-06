package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// IsFileExists : file is exists.
func IsFileExists(filename string) (bool, error) {
	_, err := os.Stat(filename)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, err
	}
	return false, err
}

// NewFile : create new file
func NewFile(filename string, head string) (file *os.File) {
	filename, err := filepath.Abs(filename)
	if err != nil {
		log.Fatalln(err)
	}

	if exist, _ := IsFileExists(filename); !exist {
		// create fold
		if err = os.MkdirAll(filepath.Dir(filename), os.ModePerm); err != nil {
			log.Fatalln("can not create fold,", filepath.Dir(filename))
		}
		// create file
		file, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModeAppend)
		if err != nil {
			log.Fatalln("failed to open file,", err)
		}
		// modifyfile authority
		if err := os.Chmod(filename, 0666); err != nil {
			log.Fatalln("can not modify file authority to 0666,", err)
		}
	} else {
		file, err = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
		if err != nil {
			log.Fatalln("failed to open file,", err)
		}
	}

	if head != "" {
		file.WriteString(head)
	}
	return file
}

// FileSplitAsDate spilt file of project, rename current file.
func FileSplitAsDate(filename, timeformat string, isBeforeDay bool) (err error) {
	filename, err = filepath.Abs(filename)
	if err != nil {
		log.Fatalln(err)
	}
	if _, err := IsFileExists(filename); err != nil {
		return err
	}

	var date int64
	if isBeforeDay {
		date = time.Now().AddDate(0, 0, -1).Unix()
	} else {
		date = time.Now().Unix()
	}

	root := filepath.Dir(filename)
	base := filepath.Base(filename)
	arr := strings.Split(base, ".")
	var newFilename string
	if len(arr) > 1 {
		newFilename = strings.Join(arr[:len(arr)-1], ",") + "_" + TimeStamp2String(timeformat, date) + "." + arr[len(arr)-1:][0]
	} else {
		newFilename = arr[0] + "_" + TimeStamp2String(timeformat, date)
	}
	newFilename = root + "/" + newFilename

	// mv
	command := fmt.Sprintf("cat %s > %s && cat /dev/null > %s", filename, newFilename, filename)
	out, err := Run(command)
	fmt.Println(out)
	return err
}
