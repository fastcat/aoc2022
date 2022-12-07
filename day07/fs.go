package day07

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/fastcat/aoc2022/i"
	"golang.org/x/exp/slices"
)

type entry interface {
	name() string
	isDir() bool
	size() int64
}

type dir struct {
	_name   string
	_size   int64
	entries []entry
}

func (*dir) isDir() bool    { return true }
func (d *dir) name() string { return d._name }
func (d *dir) size() int64 {
	if d._size <= 0 {
		d._size = 0
		for _, e := range d.entries {
			d._size += e.size()
		}
	}
	return d._size
}

func (d *dir) add(e entry) {
	// TODO: check unique names
	d.entries = append(d.entries, e)
	// clear size cache
	d._size = -1
}

type file struct {
	_name string
	_size int64
}

func (f *file) name() string { return f._name }
func (*file) isDir() bool    { return false }
func (f *file) size() int64  { return f._size }

func newDir(name string) *dir {
	return &dir{_name: name}
}
func newFile(name string, size int64) *file {
	return &file{name, size}
}

func parse(in string) *dir {
	dirStack := []*dir{newDir("/")}
	cwd := dirStack[0]
	lines := i.Split[rune](i.Runes(in), []rune{'\n'})
	inLs := false
	i.For(lines, func(l []rune, idx int) {
		ls := string(l)
		if idx == 0 {
			if ls != "$ cd /" {
				panic(errors.New("first cmd must be cd /"))
			}
			return // continue
		} else if ls == "$ cd /" {
			panic(errors.New("can't cd / after first cmd"))
		}
		words := strings.Split(ls, " ")
		switch words[0] {
		case "$":
			inLs = false
			if len(words) < 2 {
				panic(errors.New("missing command"))
			}
			switch words[1] {
			case "cd":
				if len(words) != 3 {
					panic(errors.New("bad cd command"))
				}
				if words[2] == ".." {
					if len(dirStack) == 1 {
						panic(errors.New("can't cd outside root"))
					}
					dirStack = dirStack[:len(dirStack)-1]
					cwd = dirStack[len(dirStack)-1]
				} else {
					eidx := slices.IndexFunc(cwd.entries, func(e entry) bool { return e.name() == words[2] })
					if eidx < 0 {
						panic(fmt.Errorf("dir not found %s", words[2]))
					}
					e := cwd.entries[eidx]
					if e, ok := e.(*dir); !ok || !e.isDir() {
						panic(fmt.Errorf("not a dir: %s", words[2]))
					} else {
						dirStack = append(dirStack, e)
						cwd = e
					}
				}
			case "ls":
				if len(words) != 2 {
					panic(errors.New("ls doesn't accept args"))
				}
				inLs = true
			default:
				panic(fmt.Errorf("invalid command %s", words[1]))
			}
		case "dir":
			if len(words) != 2 {
				panic(errors.New("bad dir line"))
			}
			if !inLs {
				panic(errors.New("data outside ls"))
			}
			d := newDir(words[1])
			cwd.add(d)
		default:
			if len(words) != 2 {
				panic(errors.New("bad file line"))
			}
			size, err := strconv.ParseInt(words[0], 10, 64)
			if err != nil {
				panic(err)
			}
			f := newFile(words[1], size)
			cwd.add(f)
		}
	})
	return dirStack[0]
}

func (d *dir) recursiveFiles() i.Iterator[*file] {
	idx := -1
	var it i.Iterator[*file]
	return i.Func(func() (*file, bool) {
		for {
			if it == nil {
				// move to the next entry, if it's a file yield it
				idx++
				if idx >= len(d.entries) {
					return nil, true
				}
				e := d.entries[idx]
				if f, ok := e.(*file); ok {
					return f, false
				} else if d, ok := e.(*dir); ok {
					it = d.recursiveFiles()
				} else {
					panic(errors.New("bad entry type"))
				}
			}
			f, done := it.Next()
			if done {
				// move to the next entry
				it = nil
				continue
			}
			return f, false
		}
	})
}

func (d *dir) recursiveDirs() i.Iterator[*dir] {
	idx := -2
	var it i.Iterator[*dir]
	return i.Func(func() (*dir, bool) {
		if idx == -2 {
			idx++
			return d, false
		}
		for {
			if it == nil {
				// move to the next entry, if it's a file yield it
				idx++
				if idx >= len(d.entries) {
					return nil, true
				}
				e := d.entries[idx]
				if _, ok := e.(*file); ok {
					continue
				} else if d, ok := e.(*dir); ok {
					it = d.recursiveDirs()
				} else {
					panic(errors.New("bad entry type"))
				}
			}
			d, done := it.Next()
			if done {
				// move to the next entry
				it = nil
				continue
			}
			return d, false
		}
	})
}
