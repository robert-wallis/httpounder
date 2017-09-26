package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

var host = flag.String("host", "127.0.0.1:3000", "host:port")

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

func sendHttp(method string, location string) (res string, err error) {
	message := fmt.Sprintf("%s %s\r\n\r\n", method, location)
	return sendText(message)
}

func replayLine(id int, line string) error {
	fields := strings.Fields(line)
	if len(fields) < 2 {
		return fmt.Errorf("line should be \"METHOD location\" \"%s\"", line)
	}
	fmt.Printf("#%d: %s => ", id, line)
	res, err := sendHttp(fields[0], fields[1])
	if err != nil {
		fmt.Println("")
		return err
	}
	fmt.Println(res)
	return nil
}

func replayFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	id := 1
	for scanner.Scan() {
		if err := replayLine(id, scanner.Text()); err != nil {
			return err
		}
		id ++
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Please specify log file name.")
		flag.PrintDefaults()
		os.Exit(-100)
		return
	}
	err := replayFile(args[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
		return
	}
	os.Exit(0)
}
