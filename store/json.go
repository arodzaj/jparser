package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	log "github.com/Sirupsen/logrus"
)

type File struct {
	Ext  string
	Root Node
	Name string
	Date time.Time
}

func (f *File) String() string {
	return ""
}

func (f *File) Init(body, ext string) error {
	if ext != "json" {
		return errors.New("Not supported file type")
	}

	f.Ext = ext

	b := []byte(body)
	buffer := map[string]interface{}{}
	if err := json.Unmarshal(b, &buffer); err != nil {
		log.Error(err)
		return errors.New("Unable to parse file")
	}

	f.Root = Parse(buffer)

	fmt.Print(f.Root.String())

	return nil
}
