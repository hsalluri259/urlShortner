package store

import "sync"

type URLStore struct {
	urls map[string]string
	mu   sync.RWMutex
}

func NewURLStore() *URLStore {
	return &URLStore{urls: make(map[string]string)}
}

// Get the long URL using the key
func (s *URLStore) Get(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.urls[key]
}

// Store long URL as value for generated short url as key
func (s *URLStore) Set(key, url string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	// check for key existence
	if _, present := s.urls[key]; present {
		return false
	}

	s.urls[key] = url
	return true
}

func (s *URLStore) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.urls)
}

func (s *URLStore) Put(url string) string {
	for {
		//  generate the short URL
		key := genKey(s.Count())
		if ok := s.Set(key, url); ok {
			return key
		}
	}
	// shouldn't get here
	return ""
}
