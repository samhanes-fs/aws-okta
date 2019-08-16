package lib

import (
	"errors"
	"fmt"
	"time"

	u2fhost "github.com/marshallbrekka/go-u2fhost"
)

func promptU2f(request *u2fhost.AuthenticateRequest, timeout int) (response *u2fhost.AuthenticateResponse, err error) {
	allDevices := u2fhost.Devices()
	openDevices := []u2fhost.Device{}
	for _, device := range allDevices {
		fmt.printf("U2f device %s", device)
		err := device.Open()
		if err == nil {
			openDevices = append(openDevices, device)
			defer device.Close()
		}
	}
	if len(openDevices) == 0 {
		return nil, errors.New("no accessible U2F devices")
	}

	prompted := false
	timeoutC := time.After(time.Second * time.Duration(timeout))
	interval := time.NewTicker(time.Millisecond * 250)
	defer interval.Stop()
	for {
		select {
		case <-timeoutC:
			return nil, fmt.Errorf("Failed to get U2F response after %d seconds", timeout)
		case <-interval.C:
			for _, device := range openDevices {
				response, err = device.Authenticate(request)
				if err == nil {
					return response, nil
				} else if _, ok := err.(*u2fhost.TestOfUserPresenceRequiredError); ok {
					if !prompted {
						fmt.Println("Touch the flashing U2F device to authenticate...")
						fmt.Println()
					}
					prompted = true
				} else {
					return nil, err
				}
			}
		}
	}
}
