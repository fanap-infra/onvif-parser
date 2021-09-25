package onvif

import "strings"

const BOSCH = "BOSCH"

type Bosch struct {
}

func (p Bosch) CheckMotion(data []byte) bool {
	return strings.Contains(string(data), "MotionAlarm")
}
