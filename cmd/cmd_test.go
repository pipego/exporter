package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pipego/exporter/config"
)

func TestHost(t *testing.T) {
	h := _host()
	assert.NotEqual(t, "", h)
}

func TestMilliCPU(t *testing.T) {
	alloc, request := milliCPU()

	assert.NotEqual(t, -1, alloc)
	assert.NotEqual(t, -1, request)
}

func TestMemory(t *testing.T) {
	alloc, request := memory()

	assert.NotEqual(t, -1, alloc)
	assert.NotEqual(t, -1, request)
}

func TestStorage(t *testing.T) {
	alloc, request := storage()

	assert.NotEqual(t, -1, alloc)
	assert.NotEqual(t, -1, request)
}

func TestStats(t *testing.T) {
	c := config.New()

	c.AllocatableResource.MilliCPU, c.RequestedResource.MilliCPU = milliCPU()
	c.AllocatableResource.Memory, c.RequestedResource.Memory = memory()
	c.AllocatableResource.Storage, c.RequestedResource.Storage = storage()

	cpu, os, _memory, _storage := stats(c.AllocatableResource, c.RequestedResource)

	assert.NotEqual(t, nil, cpu)
	assert.NotEqual(t, "", os)
	assert.NotEqual(t, nil, _memory)
	assert.NotEqual(t, nil, _storage)
}
