package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func sendText(message string) (res string, err error) {
	conn, err := net.Dial("tcp", *host)
	if err != nil {
		return "", err
	}
	defer conn.Close()
	fmt.Fprintf(conn, message)
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(status), nil
}
