package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var host = flag.String("host", "127.0.0.1:3000", "host:port")

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
	if err == io.EOF {
		return
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
		return
	}
}
