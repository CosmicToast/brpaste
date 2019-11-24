package storage

const (
	// Collision is a fatal collision error, only triggered when it fails
	Collision = Error("collision detected")
	// Unhealthy is a fatal error when the backend ceases to be healthy
	Unhealthy = Error("backend unhealthy")
)

// Error is a sentinel error type for storage engines
type Error string

func (e Error) Error() string { return string(e) }

// CHR - Create, Health, Read
type CHR interface {
	Create(key, value string, checkcollision bool) error
	Healthy() bool
	Read(key string) (string, error)
}
