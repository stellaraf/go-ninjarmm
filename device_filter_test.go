package ninjarmm_test

import (
	"testing"
	"time"

	"github.com/stellaraf/go-ninjarmm"
	"github.com/stretchr/testify/assert"
)

func Test_DeviceFilter(t *testing.T) {
	t.Run("org eq", func(t *testing.T) {
		t.Parallel()
		df := ninjarmm.NewDeviceFilter()
		result := df.Org(ninjarmm.EQ, 1).String()
		assert.Equal(t, "org eq 1", result)
	})
	t.Run("org neq", func(t *testing.T) {
		t.Parallel()
		df := ninjarmm.NewDeviceFilter()
		result := df.Org(ninjarmm.NEQ, 1).String()
		assert.Equal(t, "org neq 1", result)
	})
	t.Run("org in", func(t *testing.T) {
		t.Parallel()
		df := ninjarmm.NewDeviceFilter()
		result := df.Org(ninjarmm.IN, 1, 2, 3).String()
		assert.Equal(t, "org in (1,2,3)", result)
	})
	t.Run("org notin", func(t *testing.T) {
		t.Parallel()
		df := ninjarmm.NewDeviceFilter()
		result := df.Org(ninjarmm.NOTIN, 1, 2, 3).String()
		assert.Equal(t, "org notin (1,2,3)", result)
	})
	t.Run("org and loc", func(t *testing.T) {
		t.Parallel()
		df := ninjarmm.NewDeviceFilter()
		result := df.Org(ninjarmm.EQ, 1).Loc(ninjarmm.EQ, 2).String()
		assert.Equal(t, "org eq 1 AND loc eq 2", result)
	})
	t.Run("status", func(t *testing.T) {
		t.Parallel()
		df := ninjarmm.NewDeviceFilter()
		result := df.Status(ninjarmm.EQ, ninjarmm.APPROVED).String()
		assert.Equal(t, "status eq APPROVED", result)
	})
	t.Run("status wrong op", func(t *testing.T) {
		t.Parallel()
		df := ninjarmm.NewDeviceFilter()
		assert.Panics(t, func() {
			df.Status(ninjarmm.IN, ninjarmm.APPROVED)
		})
	})
	t.Run("offline", func(t *testing.T) {
		t.Parallel()
		df := ninjarmm.NewDeviceFilter()
		result := df.Offline().String()
		assert.Equal(t, "offline", result)
	})
	t.Run("online", func(t *testing.T) {
		t.Parallel()
		df := ninjarmm.NewDeviceFilter()
		result := df.Online().String()
		assert.Equal(t, "online", result)
	})
	t.Run("org and online", func(t *testing.T) {
		t.Parallel()
		df := ninjarmm.NewDeviceFilter()
		result := df.Org(ninjarmm.EQ, 1).Online().String()
		assert.Equal(t, "org eq 1 AND online", result)
	})
	t.Run("class and role", func(t *testing.T) {
		t.Parallel()
		df := ninjarmm.NewDeviceFilter()
		result := df.Class(ninjarmm.IN, ninjarmm.NodeClass_MAC, ninjarmm.NodeClass_WINDOWS_SERVER).Role(ninjarmm.EQ, 2).String()
		assert.Equal(t, "class in (MAC,WINDOWS_SERVER) AND role eq 2", result)
	})
	t.Run("id not in", func(t *testing.T) {
		t.Parallel()
		df := ninjarmm.NewDeviceFilter()
		result := df.ID(ninjarmm.NOTIN, 1, 2, 3).String()
		assert.Equal(t, "id notin (1,2,3)", result)
	})
	t.Run("created", func(t *testing.T) {
		t.Parallel()
		df := ninjarmm.NewDeviceFilter()
		before := time.Date(2024, 05, 25, 0, 0, 0, 0, time.UTC)
		result := df.Created(ninjarmm.BEFORE, before).String()
		assert.Equal(t, "created before 2024-05-25", result)
	})
	t.Run("org and group", func(t *testing.T) {
		t.Parallel()
		df := ninjarmm.NewDeviceFilter()
		result := df.Org(ninjarmm.EQ, 1).Group(2).String()
		assert.Equal(t, "org eq 1 AND group 2", result)
	})
	t.Run("time wrong op", func(t *testing.T) {
		t.Parallel()
		assert.Panics(t, func() {
			df := ninjarmm.NewDeviceFilter()
			df.Created(ninjarmm.IN, time.Now())
		})
	})
	t.Run("encode", func(t *testing.T) {
		t.Parallel()
		df := ninjarmm.NewDeviceFilter()
		result := df.Org(ninjarmm.EQ, 1).Encode()
		assert.Equal(t, "org%20eq%201", result)
	})
}
