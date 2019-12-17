package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDummy(t *testing.T) {
	assert := assert.New(t)

	const (
		k = "key"
		v = "val"
		e = "key"
	)

	var (
		// setup
		err  = Error("testing error")
		dumC = Dummy{collide: true, healthy: true, err: nil}
		dumE = Dummy{collide: false, healthy: true, err: err}
		dumH = Dummy{collide: false, healthy: false, err: nil}
		dumN = Dummy{collide: false, healthy: true, err: nil}

		// ---- TESTS
		// Health
		hc = dumC.Healthy()
		he = dumE.Healthy()
		hh = dumH.Healthy()
		hn = dumN.Healthy()

		// Write - no collisions
		wc = dumC.Create(k, v, false)
		we = dumE.Create(k, v, false)
		wh = dumH.Create(k, v, false)
		wn = dumN.Create(k, v, false)

		// Write - collisions
		cc = dumC.Create(k, v, true)
		ce = dumE.Create(k, v, true)
		ch = dumH.Create(k, v, true)
		cn = dumN.Create(k, v, true)

		// Read
		rc, erc = dumC.Read(k)
		re, ere = dumE.Read(k)
		rh, erh = dumH.Read(k)
		rn, ern = dumN.Read(k)
	)
	assert.Equal("testing error", err.Error())

	assert.True(hc)
	assert.True(he)
	assert.False(hh)
	assert.True(hn)

	assert.Nil(wc)
	assert.NotNil(we)
	assert.NotNil(wh)
	assert.Nil(wn)

	assert.Equal(err, we)
	assert.Equal(Unhealthy, wh)

	assert.NotNil(cc)
	assert.NotNil(ce)
	assert.NotNil(ch)
	assert.Nil(cn)

	assert.Equal(Collision, cc)
	assert.Equal(err, ce)
	assert.Equal(Unhealthy, ch)

	assert.Nil(erc)
	assert.Nil(ere)
	assert.NotNil(erh)
	assert.Nil(ern)

	assert.Equal(Unhealthy, erh)

	assert.Equal(e, rc)
	assert.Equal(e, re)
	assert.Equal(e, rh)
	assert.Equal(e, rn)
}
