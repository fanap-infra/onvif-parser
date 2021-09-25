package onvif

import "strings"

const DAHUA = "General"

type Dahua struct {
}

func (p Dahua) CheckMotion(data []byte) bool {
	return strings.Contains(string(data), "MotionAlarm")
}
