package clang

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompletion(t *testing.T) {
	idx := NewIndex(0, 0)
	defer idx.Dispose()

	tu := idx.ParseTranslationUnit("cursor.c", nil, nil, 0)
	assert.True(t, tu.IsValid())
	defer tu.Dispose()

	res := tu.CodeCompleteAt("cursor.c", 5, 18, nil, 0)
	assert.NotNil(t, res)
	defer res.Dispose()

	if n := len(res.Results()); n < 10 {
		t.Errorf("Expected more results than %d", n)
	}

	t.Logf("%+v", res)
	for _, r := range res.Results() {
		t.Logf("%+v", r)

		cs := r.CompletionString()

		for i := uint32(0); i < cs.NumChunks(); i++ {
			t.Logf("\t%s %s", cs.ChunkKind(i), cs.ChunkText(i))
		}
	}

	diags := res.Diagnostics()
	defer func() {
		for _, d := range diags {
			d.Dispose()
		}
	}()

	ok := false
	for _, d := range diags {
		if strings.Contains(d.Spelling(), "_cgo_export.h") {
			ok = true
		}
		t.Log(d.Severity(), d.Spelling())
	}
	assert.True(t, ok)
}
