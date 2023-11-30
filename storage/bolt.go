package storage

import bolt "go.etcd.io/bbolt"

var _ CHR = &Bolt{}

// Redis storage engine
type Bolt bolt.DB

var bname = []byte("brpaste")

func OpenBolt(db *bolt.DB) (*Bolt, error) {
	err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bname)
		return err
	})
	return (*Bolt)(db), err
}

// Create an entry in redis
func (db *Bolt) Create(key, value string, checkcollision bool) error {
	err := (*bolt.DB)(db).Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bname)
		if b == nil {
			return Unhealthy
		}
		k := []byte(key)
		if checkcollision {
			if b.Get(k) != nil {
				return Collision
			}
		}
		return b.Put(k, []byte(value))
	})
	if err != nil {
		return err
	}
	return nil
}

func (db *Bolt) Read(key string) (string, error) {
	var out string
	err := (*bolt.DB)(db).View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bname)
		if b == nil {
			return Unhealthy
		}
		v := b.Get([]byte(key))
		if v == nil {
			return NotFound
		}
		out = string(v)
		return nil
	})
	if err != nil {
		return "", err
	}
	return out, nil
}

// Healthy TODO: destub ?
func (db *Bolt) Healthy() bool {
	return db != nil
}
