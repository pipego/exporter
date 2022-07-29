package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHost(t *testing.T) {
	h := host()
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
