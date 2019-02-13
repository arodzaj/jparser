package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	log "github.com/Sirupsen/logrus"
)

type File struct {
	Root Node
	Name string
	Date time.Time
}

func (f *File) String() string {
	return fmt.Sprintf("<name:%s, date:%s>", f.Name, f.Date)
}

func (f *File) Init(body, name string) error {
	b := []byte(body)
	buffer := map[string]interface{}{}
	if err := json.Unmarshal(b, &buffer); err != nil {
		log.Error(err)
		return errors.New("Unable to parse file")
	}

	f.Root = Parse(buffer)
	f.Date = time.Now()
	f.Name = name

	return nil
}
