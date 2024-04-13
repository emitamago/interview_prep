package main

import (
	"encoding/gob"
	"io"
	"log"
	"os"
	"sync"
)

const saveQueueLength = 1000

type URLStore struct {
	urls map[string]string // from short to long urls
	mu   sync.RWMutex      // provide read and write lock for urls maps
	save chan record
}

type record struct {
	Key string
	URL string
}

func (s *URLStore) saveLoop(filename string) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("URLStore: ")
	}
	defer f.Close()
	e := gob.NewEncoder(f)
	for {
		r := s.save // taking a record from the channel and encoding it
		if err := e.Encode(r); err != nil {
			log.Println("URLStore:", err)
		}
	}
}

func (s *URLStore) load(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		log.Println("Error opening URLStore:", err)
		return err
	}
	defer f.Close()
	d := gob.NewDecoder(f)
	for err == nil {
		var r record
		if err = d.Decode(&r); err == nil {
			s.Set(r.Key, r.URL)
		}
	}
	if err == io.EOF {
		return nil
	}
	log.Println("Error decoding URLStore:", err) // map hasn't been read correctly
	return err

}

// Get ---> get long url from existing map with read lock
func (s *URLStore) Get(key string) string {
	s.mu.RLock() // lock read on map
	defer s.mu.RUnlock()
	// url := s.urls[key] ---> because of defer, we do not need to store url into var
	return s.urls[key]
}

// Set --> Set new url map if that does not exist already. Return bool if new map has been set
func (s *URLStore) Set(key, url string) bool {
	s.mu.Lock() // lock write/read on map
	defer s.mu.Unlock()
	_, present := s.urls[key]
	if present {
		// s.mu.Unlock() --> can remove this because of defer
		return false // did not set new url, key is present already
	}
	s.urls[key] = url
	// s.mu.Unlock() --> can remove bacause we are defering unlock on line 22
	return true // set new map

}

// NewURLStore --> initialize new url
func NewURLStore(filename string) *URLStore {
	s := &URLStore{
		urls: make(map[string]string),
		save: make(chan record, saveQueueLength),
	}
	if err := s.load(filename); err != nil {
		log.Fatal("Error loading URLStore: ", err)
	}
	go s.saveLoop(filename)
	return s
}

// Count --> return how many url present in URLStore
func (s *URLStore) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.urls)
}

// Put -- > Take long URL, generate short key, and set new url using the generated key
func (s *URLStore) Put(url string) string {
	for {
		key := genKey(s.Count())
		if s.Set(key, url) {
			s.save <- record{key, url}
			return key
		}
	} // this for loop will retry Set() until it succeed
	// Should not get here
	panic("should not get here")
}
