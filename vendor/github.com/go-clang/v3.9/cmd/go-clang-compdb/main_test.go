package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoClangCompDB(t *testing.T) {
	for _, path := range []string{
		"../../testdata",
	} {
		assert.Equal(t, 0, cmd([]string{path}))
	}
}
