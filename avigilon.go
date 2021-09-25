package onvif

import "strings"

const AVIGILON = "Avigilon"

type Avigilon struct {
}

func (p Avigilon) CheckMotion(data []byte) bool {
	return strings.Contains(string(data), `Name="MotionActive" Value="1"`)
}
