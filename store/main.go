package store

import (
	"errors"
	"sync"
)

// Store - main store type
type Store struct {
	sync.Mutex

	jsons []json
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

	return s.jsons[id].String(), nil
}

func (s *Store) Add(js string) int {
	j := parseJS(js)
	var id int

	s.Lock()
	s.jsons = append(s.jsons, j)
	id = len(s.jsons) - 1
	s.Unlock()

	return id
}
