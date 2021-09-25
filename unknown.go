package onvif

import (
	"strings"
)

type Unknown struct {
}

func (p Unknown) CheckMotion(data []byte) bool {
	if strings.Contains(string(data), "MotionAlarm") {
		return true
	}
	return false
}
