package clang

func (ccr *CodeCompleteResults) Diagnostics() []Diagnostic { // TODO this can be generated https://github.com/go-clang/gen/issues/47
	s := make([]Diagnostic, ccr.NumDiagnostics())

	for i := range s {
		s[i] = ccr.Diagnostic(uint32(i))
	}

	return s
}
