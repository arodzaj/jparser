package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
)

const usageMessage = `usage: jsparser [-s][--file FILE][--watch][--version][--help] JSON_STRING

general options:
   -s          send json to the server
   --file      parse given file
   --watch     run server of json monitor
   --version   print script version
   --help      print script help`

func usage() {
	fmt.Fprintf(os.Stderr, usageMessage)
	os.Exit(2)
}

// input arguments variables
var (
	sendArg  = flag.Bool("s", false, "")
	watchArg = flag.Bool("watch", false, "")
	fileArg  = flag.String("file", "", "")
)

func run(js string) error {

	return nil
}

func getJSON() (string, error) {
	stat, _ := os.Stdin.Stat()

	if len(flag.Args()) == 1 {
		// JSON passed as argument
		return flag.Args()[0], nil

	} else if *fileArg != "" {
		// JSON passed as file path
		return *fileArg, nil

	} else if (stat.Mode() & os.ModeCharDevice) == 0 {
		// JSON passed via pipeline
		scanner := bufio.NewScanner(os.Stdin)
		var buffer bytes.Buffer
		for scanner.Scan() {
			buffer.WriteString(scanner.Text())
		}

		return buffer.String(), nil
	} else if *watchArg {
		return "", nil
	}

	return "", errors.New("Wrong parameters")

}

func main() {
	log.Info("Starting script")

	flag.Usage = usage
	flag.Parse()

	js, err := getJSON()
	if err != nil {
		usage()
	}

	if err := run(js); err != nil {
		log.Panic("...")
	}

}
