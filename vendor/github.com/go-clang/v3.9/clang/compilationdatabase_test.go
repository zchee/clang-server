package clang

import (
	"testing"
)

func TestCompilationDatabaseError(t *testing.T) {
	err, _ := FromDirectory("../testdata-not-there")
	if err != CompilationDatabase_CanNotLoadDatabase {
		t.Fatalf("expected %v", CompilationDatabase_CanNotLoadDatabase)
	}
}

func TestCompilationDatabase(t *testing.T) {
	err, db := FromDirectory("../testdata")
	if err != CompilationDatabase_NoError {
		t.Fatalf("error loading compilation database: %v", err)
	}
	defer db.Dispose()

	table := []struct {
		directory string
		args      []string
	}{
		{
			directory: "/home/user/llvm/build",
			args: []string{
				"/usr/bin/clang++",
				"-Irelative",
				//FIXME: bug in clang ?
				//`-DSOMEDEF="With spaces, quotes and \-es.`,
				"-DSOMEDEF=With spaces, quotes and -es.",
				"-c",
				"-o",
				"file.o",
				"file.cc",
			},
		},
		{
			directory: "@TESTDIR@",
			args:      []string{"g++", "-c", "-DMYMACRO=a", "subdir/a.cpp"},
		},
	}

	cmds := db.AllCompileCommands()
	if int(cmds.Size()) != len(table) {
		t.Errorf("expected #cmds=%d. got=%d", len(table), cmds.Size())
	}

	for i := 0; i < int(cmds.Size()); i++ {
		cmd := cmds.Command(uint32(i))
		if cmd.Directory() != table[i].directory {
			t.Errorf("expected dir=%q. got=%q", table[i].directory, cmd.Directory())
		}

		nargs := int(cmd.NumArgs())
		if nargs != len(table[i].args) {
			t.Errorf("expected #args=%d. got=%d", len(table[i].args), nargs)
		}
		if nargs > len(table[i].args) {
			nargs = len(table[i].args)
		}
		for j := 0; j < nargs; j++ {
			arg := cmd.Arg(uint32(j))
			if arg != table[i].args[j] {
				t.Errorf("expected arg[%d]=%q. got=%q", j, table[i].args[j], arg)
			}
		}
	}
}
