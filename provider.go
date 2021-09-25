package onvif

import (
	"fmt"
)

func NewParser(Manufacturer string, Model string) Parser {
	switch {
	case Manufacturer == AVIGILON:
		return Avigilon{}
	case Manufacturer == BOSCH:
		return Bosch{}
	case Manufacturer == GRUNDIG:
		return Grundig{}
	case Manufacturer == DAHUA:
		return Dahua{}
	default:
		fmt.Printf("%v is unknown, model: %v", Manufacturer, Model)
		return Unknown{}
	}
}
