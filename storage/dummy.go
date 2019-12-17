package storage

// Dummy is a dummy storage device that acts predictably for tests
type Dummy struct {
	collide, healthy bool
	err              error
}

// Create is a noop
// if the backend was created unhealthy, it will error out with the unhealthy error
// if the backend was set as colliding, and collision detection is enabled, it will error out with a collision
// if the backend was set as erroring, it will error with that error
func (d *Dummy) Create(key, value string, checkcollision bool) error {
	if !d.Healthy() {
		return Unhealthy
	}
	if checkcollision && d.collide {
		return Collision
	}
	if d.err != nil {
		return d.err
	}
	return nil
}

// Read will return the key that is sent to it, or error out if the backend was created unhealthy
// Read echoes the key given to it back
// if the backend was set to unhealthy, it also returns the Unhealthy error
func (d *Dummy) Read(key string) (string, error) {
	var err error
	if !d.Healthy() {
		err = Unhealthy
	}
	return key, err
}

// Healthy checks whether or not the backend was created healthy
func (d *Dummy) Healthy() bool { return d.healthy }
