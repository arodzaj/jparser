package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"os"
)

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
