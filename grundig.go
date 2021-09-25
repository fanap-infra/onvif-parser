package onvif

import "strings"

const GRUNDIG = "Grundig"

type Grundig struct {
}

func (p Grundig) CheckMotion(data []byte) bool {
	return strings.Contains(string(data), "MotionAlarm")
}
