package cloud

import (
	"log"
	"sync"
)

type MockCloudProvider struct {
	mu    sync.Mutex
	nodes []string
}

func NewMockCloudProvider() *MockCloudProvider {
	return &MockCloudProvider{
		nodes: []string{"node-1"},
	}
}

func (m *MockCloudProvider) AddNode(name string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.nodes = append(m.nodes, name)
	log.Printf("[CLOUD] Added %s (total=%d)\n", name, len(m.nodes))
	return nil
}

func (m *MockCloudProvider) RemoveNode(name string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, n := range m.nodes {
		if n == name {
			m.nodes = append(m.nodes[:i], m.nodes[i+1:]...)
			log.Printf("[CLOUD] Removed %s (total=%d)\n", name, len(m.nodes))
			return nil
		}
	}
	log.Fatalf("node %s not found", name)
	return nil

}

func (m *MockCloudProvider) ListNodes() ([]string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	return append([]string{}, m.nodes...), nil
}
