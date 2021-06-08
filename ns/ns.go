package main

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
	"time"
)

// Debug implements log.Printf
var Debug = func(string, ...interface{}) {}

// a rament is a stem a slice of []ReadDirFS
// It must implement, at minimum, fs.ReadDirFS.
// We require ReadDirFS so that we need not support
// open directory files where not needed.
// The entries are sorted so that the longest stem is first.
// These are only entries, no files.
type rament struct {
	Stem string
	// Rep *rament TODO: replacement mount
	FS []*rament
}

var _ fs.ReadDirFS = &rament{}

var _ fs.DirEntry = &rament{}

func NewRAMFS() *rament {
	return &rament{Stem: "/"}
}

func (r *rament) Open(n string) (fs.File, error) {
	return nil, fmt.Errorf("not yet")
}

func (r *rament) ReadDir(n string) ([]fs.DirEntry, error) {
	var ents []fs.DirEntry
	fs := r.FS
	var ret []*rament

consumed:
	for len(n) > 0 {
		Debug("Find %v in %q", n, fs)

		for n != "" && fs != nil {
			Debug("n is now %q, fs is %q", n, fs)
			for _, e := range fs {
				// Really need a filepath version of this I guess.
				Debug("Check %q with name %q %v", n, e.Stem, strings.HasPrefix(n, e.Stem))
				if strings.HasPrefix(n, e.Stem) {
					ret = e.FS
					fs = e.FS
					n = strings.TrimPrefix(n, e.Stem)
					if len(n) == 0 {
						break consumed
					}

					break
				}
			}
		}
		if ret == nil {
			return nil, os.ErrNotExist
		}
	}

	for _, e := range fs {
		ents = append(ents, e)
	}
	return ents, nil
}

func (r *rament) Info() (fs.FileInfo, error) {
	return nil, fmt.Errorf("not yet")
}

func (r *rament) Name() string {
	return r.Stem
}

func (r *rament) Size() int64 {
	return int64(len(r.FS))
}

func (r *rament) Mode() fs.FileMode {
	return fs.FileMode(0666)
}

func (r *rament) Type() fs.FileMode {
	return fs.FileMode(0666)
}

func (r *rament) ModTime() time.Time {
	return time.Now()
}

func (r *rament) IsDir() bool {
	return true
}

func (r *rament) Sys() interface{} {
	return nil
}
