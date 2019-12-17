package storage

var _ CHR = &Memory{}

// Memory is an in-memory non-persistent storage type
// using it in production is not recommended
type Memory struct {
	store map[string]string
}

// Create will store a value in a key in memory
func (r *Memory) Create(key, value string, checkcollision bool) error {
	if !r.Healthy() {
		return Unhealthy
	}
	if checkcollision {
		if _, ok := r.store[key]; ok {
			return Collision
		}
	}
	r.store[key] = value
	return nil
}

// Read will read a key from memory, if it's there
func (r *Memory) Read(key string) (string, error) {
	if !r.Healthy() {
		return "", Unhealthy
	}
	if val, ok := r.store[key]; ok {
		return val, nil
	}
	return "", Error("value not found")
}

// Healthy checks if the memory storage is initialized
func (r *Memory) Healthy() bool {
	return r.store != nil
}

// NewMemory initializes a memory backend for use
func NewMemory() CHR {
	m := Memory{
		store: make(map[string]string),
	}
	return &m
}
