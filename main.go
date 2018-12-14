package main

import (
	"flag"
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/arodzaj/jparser/store"
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

// variables
var (
	space *store.Store
)

func run(js string) error {
	space = store.New()

	if js != "" {
		space.Add(js)
	}

	if *watchArg {
		go func() {

		}()
	}

	for {
	}

	return nil
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
