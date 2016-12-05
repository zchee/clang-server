package main_test

import (
	"os/exec"
	"testing"
)

func TestGoClangDump(t *testing.T) {
	for _, fname := range []string{
		"../../testdata/basicparsing.c",
	} {
		cmd := exec.Command("go-clang-dump", "-fname", fname)
		err := cmd.Run()
		if err != nil {
			t.Fatalf("error running go-clang-dump on %q: %v\n", fname, err)
		}
	}
}
