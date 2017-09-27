package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
		id++
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
