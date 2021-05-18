package main

import "sync"

type memoryStore struct {
	Mu    sync.Mutex
	Store map[string]string
}

func (m memoryStore) Write(data secretData) {
	m.Mu.Lock()
	defer m.Mu.Unlock()

	m.Store[data.Id] = data.Secret

}

func (m memoryStore) Read(id string) string {
	m.Mu.Lock()
	defer m.Mu.Unlock()

	data := m.Store[id]
	delete(m.Store, id)

	return data
}
