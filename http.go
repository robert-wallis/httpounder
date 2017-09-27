package main

import "fmt"

func sendHttp(method string, location string) (res string, err error) {
	message := fmt.Sprintf("%s %s\r\n\r\n", method, location)
	return sendText(message)
}
