package clang

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiagnostics(t *testing.T) {
	idx := NewIndex(0, 0)
	defer idx.Dispose()

	tu := idx.ParseTranslationUnit("cursor.c", nil, nil, 0)
	assert.True(t, tu.IsValid())
	defer tu.Dispose()

	diags := tu.Diagnostics()
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
		t.Log(d)
		t.Log(d.Severity(), d.Spelling())
		t.Log(d.FormatDiagnostic(uint32(Diagnostic_DisplayCategoryName | Diagnostic_DisplaySourceLocation)))
	}
	assert.True(t, ok)
}
