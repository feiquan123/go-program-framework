package utils

import (
	"testing"
	"time"
)

func TestTimeStamp2String(t *testing.T) {
	str := TimeStamp2String(TimeFormat3, time.Now().Unix())
	t.Log(str)

	str = TimeStamp2String(TimeFormat5, time.Now().Unix())
	t.Log(str)

	str = TimeStamp2String(TimeFormat5, time.Now().AddDate(0, 0, -1).Unix())
	t.Log(str)
}
