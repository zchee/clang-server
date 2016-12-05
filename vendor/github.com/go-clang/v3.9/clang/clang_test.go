package clang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicParsing(t *testing.T) {
	idx := NewIndex(0, 1)
	defer idx.Dispose()

	tu := idx.ParseTranslationUnit("../testdata/basicparsing.c", nil, nil, 0)
	assert.True(t, tu.IsValid())
	defer tu.Dispose()

	found := 0

	tu.TranslationUnitCursor().Visit(func(cursor, parent Cursor) ChildVisitResult {
		if cursor.IsNull() {
			return ChildVisit_Continue
		}

		switch cursor.Kind() {
		case Cursor_FunctionDecl:
			assert.Equal(t, "foo", cursor.Spelling())

			found++
		case Cursor_ParmDecl:
			assert.Equal(t, "bar", cursor.Spelling())

			found++
		}

		return ChildVisit_Recurse
	})

	assert.Equal(t, 2, found, "Did not find all nodes")
}

func TestReparse(t *testing.T) {
	us := []UnsavedFile{
		NewUnsavedFile("hello.cpp", "int world();"),
	}

	idx := NewIndex(0, 0)
	defer idx.Dispose()

	tu := idx.ParseTranslationUnit("hello.cpp", nil, us, 0)
	assert.True(t, tu.IsValid())
	defer tu.Dispose()

	ok := false
	tu.TranslationUnitCursor().Visit(func(cursor, parent Cursor) ChildVisitResult {
		if cursor.Spelling() == "world" {
			ok = true

			return ChildVisit_Break
		}

		return ChildVisit_Continue
	})
	if !ok {
		t.Error("Expected to find 'world', but didn't")
	}

	us[0] = NewUnsavedFile("hello.cpp", "int world2();")
	tu.ReparseTranslationUnit(us, 0)

	ok = false
	tu.TranslationUnitCursor().Visit(func(cursor, parent Cursor) ChildVisitResult {
		if s := cursor.Spelling(); s == "world2" {
			ok = true

			return ChildVisit_Break
		} else if s == "world" {
			t.Errorf("'world' should no longer be part of the translationunit, but it still is")
		}

		return ChildVisit_Continue
	})
	if !ok {
		t.Error("Expected to find 'world2', but didn't")
	}
}
