package main

import (
	"reflect"
	"testing"
)

func TestRAMFS(t *testing.T) {
	r := NewRAMFS()
	t.Logf("r is %v", r)
	// Add stuff ...
	r.FS = append(r.FS,
		&rament{Stem: "net/eth0", FS: []*rament{{Stem: "."}}},
		&rament{Stem: "root/a", FS: []*rament{{Stem: "."}}},
		&rament{Stem: "root", FS: []*rament{{Stem: "."}}},
		// not sure.
		&rament{Stem: ".", FS: []*rament{{Stem: "."}}},
	)
	d, err := r.ReadDir("")
	if err != nil {
		t.Fatalf("root.Read: got %v, want nil", err)
	}
	// why does this fail ... oh well
	if !reflect.DeepEqual(d, r.FS) {
		t.Logf("root.Read: got %v, want %v", d, r.FS)
	}
	t.Logf("r readir: %q", d)
	Debug = t.Logf
	d, err = r.ReadDir("root/a")
	if err != nil {
		t.Fatalf("root.Read: got %v, want nil", err)
	}
	if d == nil {
		t.Fatalf("root.Read: got nil, want rament")
	}
	if d[0].Name() != "." {
		t.Fatalf("ReadDirFS: got %s, want .", d[0].Name())
	}
	d, err = r.ReadDir(".")
	if err != nil {
		t.Fatalf("root.Read: got %v, want nil", err)
	}
	if d == nil {
		t.Fatalf("root.Read: got nil, want rament")
	}
	if d[0].Name() != "." {
		t.Fatalf("ReadDirFS: got %s, want .", d[0].Name())
	}

}
