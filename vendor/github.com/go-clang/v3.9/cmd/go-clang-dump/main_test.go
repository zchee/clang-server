package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoClangDump(t *testing.T) {
	for _, fname := range []string{
		"../../testdata/basicparsing.c",
	} {
		assert.Equal(t, 0, cmd([]string{"-fname", fname}))
	}
}
