package resources

import (
	. "testing"
)

func TestPackageBundle(t *T) {
	cp, err := OpenCurrentPackage()
	if err != nil {
		t.Fatalf("OpenCurrentPacakge(): %v", err)
	}

	t.Log("CurrentPackage:", cp)
	t.Log("cp.List():", cp.(Lister).List())
	if fs, err := cp.(Searcher).Glob("*.go"); err != nil {
		t.Fatal("Glob(*.go):", err)
	} else {
		t.Log("Glob(*.go):", fs)
	}
}
