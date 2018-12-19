package store

import (
	"errors"
	"sync"

	log "github.com/Sirupsen/logrus"
)

// Store - main store type
type Store struct {
	sync.Mutex

	files []File
	cur   int // cursor
}

func New() *Store {
	s := new(Store)

	return s
}

func (s *Store) GetNext() (string, error) {

	return s.Get(s.cur)
}

func (s *Store) GetCur() (string, error) {
	return s.Get(s.cur)
}

func (s *Store) Get(id int) (string, error) {
	if id < 0 || id > s.cur {
		return "", errors.New("Index out of range")
	}

	return s.files[id].String(), nil
}

func (s *Store) Add(body string) int {
	f := File{}

	if err := f.Init(body, "json"); err != nil {
		log.Panic("Unable to parse file")
	}

	var id int

	s.Lock()
	s.files = append(s.files, f)
	id = len(s.files) - 1
	s.Unlock()

	return id
}
