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
	)

	var (
		// setup
		err  = Error("testing error")
		dumC = Dummy{collide: true, found: true, healthy: true, err: nil}
		dumE = Dummy{collide: false, found: true, healthy: true, err: err}
		dumF = Dummy{collide: false, found: false, healthy: true, err: nil}
		dumH = Dummy{collide: false, found: true, healthy: false, err: nil}
		dumN = Dummy{collide: false, found: true, healthy: true, err: nil}

		// ---- TESTS
		// Health
		hc = dumC.Healthy()
		he = dumE.Healthy()
		hf = dumF.Healthy()
		hh = dumH.Healthy()
		hn = dumN.Healthy()

		// Write - no collisions
		wc = dumC.Create(k, v, false)
		we = dumE.Create(k, v, false)
		wf = dumF.Create(k, v, false)
		wh = dumH.Create(k, v, false)
		wn = dumN.Create(k, v, false)

		// Write - collisions
		cc = dumC.Create(k, v, true)
		ce = dumE.Create(k, v, true)
		cf = dumF.Create(k, v, true)
		ch = dumH.Create(k, v, true)
		cn = dumN.Create(k, v, true)

		// Read
		rc, erc = dumC.Read(k)
		re, ere = dumE.Read(k)
		rf, erf = dumF.Read(k)
		rh, erh = dumH.Read(k)
		rn, ern = dumN.Read(k)
	)
	assert.Equal("testing error", err.Error())

	assert.True(hc)
	assert.True(he)
	assert.True(hf)
	assert.False(hh)
	assert.True(hn)

	assert.Nil(wc)
	assert.NotNil(we)
	assert.Nil(wf)
	assert.NotNil(wh)
	assert.Nil(wn)

	assert.Equal(err, we)
	assert.Equal(Unhealthy, wh)

	assert.NotNil(cc)
	assert.NotNil(ce)
	assert.Nil(cf)
	assert.NotNil(ch)
	assert.Nil(cn)

	assert.Equal(Collision, cc)
	assert.Equal(err, ce)
	assert.Equal(Unhealthy, ch)

	assert.Nil(erc)
	assert.Nil(ere)
	assert.NotNil(erf)
	assert.NotNil(erh)
	assert.Nil(ern)

	assert.Equal(Unhealthy, erh)
	assert.Equal(NotFound, erf)

	assert.Equal(k, rc)
	assert.Equal(k, re)
	assert.Equal(k, rf)
	assert.Equal(k, rh)
	assert.Equal(k, rn)
}
