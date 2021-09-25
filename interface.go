package onvif

type Parser interface {
	CheckMotion(data []byte) bool
}
