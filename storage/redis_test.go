// +build redis

package storage

import (
	"testing"

	"github.com/go-redis/redis/v7"
	"github.com/stretchr/testify/assert"
)

func TestRedis(t *testing.T) {
	assert := assert.New(t)
	const (
		k  = "key"
		v1 = "val1"
		v2 = "val2"
	)

	// when using tag "redis", we expect all of this to be as-is
	var (
		o, _ = redis.ParseURL("redis://localhost:6379")
		c    = redis.NewClient(o)
		s    = (*Redis)(c)
	)

	assert.True(s.Healthy())

	err := s.Create(k, v1, false)
	assert.Nil(err)

	err = s.Create(k, v2, false)
	assert.Nil(err)

	err = s.Create(k, v1, true)
	assert.Equal(Collision, err)

	v, err := s.Read(k)
	assert.Nil(err)
	assert.Equal(v2, v)

	_, err = s.Read(v1)
	assert.NotNil(err)
	assert.Equal(NotFound, err)
}
