package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemory(t *testing.T) {
	assert := assert.New(t)
	const (
		k  = "key"
		v1 = "val1"
		v2 = "val2"
	)

	m := NewMemory()

	err := m.Create(k, v1, false)
	assert.Nil(err)

	err = m.Create(k, v2, false)
	assert.Nil(err)

	err = m.Create(k, v1, true)
	assert.Equal(Collision, err)

	v, err := m.Read(k)
	assert.Nil(err)
	assert.Equal(v2, v)

	_, err = m.Read(v1)
	assert.NotNil(err)
	assert.Equal(NotFound, err)
}

func TestUnhealthyMemory(t *testing.T) {
	assert := assert.New(t)
	m := &Memory{}

	var (
		e1    = m.Create("", "", false)
		e2    = m.Create("", "", true)
		_, e3 = m.Read("")
	)

	assert.NotNil(e1)
	assert.NotNil(e2)
	assert.NotNil(e3)

	assert.Equal(Unhealthy, e1)
	assert.Equal(Unhealthy, e2)
	assert.Equal(Unhealthy, e3)

	assert.False(m.Healthy())
}
