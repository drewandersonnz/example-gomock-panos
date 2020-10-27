package services

import "fmt"

type pingServiceInterface interface {
	PingService() (string, error)
}

// PingServiceStruct  blah blah
type PingServiceStruct struct{}

// PingServiceVar blahb lah
var PingServiceVar pingServiceInterface = PingServiceStruct{}

// PingService returns pong
func (service PingServiceStruct) PingService() (string, error) {
	fmt.Println("Connecting to an external Database")
	return "pong", nil
}
