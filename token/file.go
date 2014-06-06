// Copyright (c) 2014, Rob Thornton
// All rights reserved.
// This source code is governed by a Simplied BSD-License. Please see the
// LICENSE included in this distribution for a copy of the full license
// or, if one is not included, you may also find a copy at
// http://opensource.org/licenses/BSD-2-Clause

package token

type File struct {
	base  int
	name  string
	lines []int
	size  int
}

func NewFile(name string, base, size int) *File {
	return &File{
		base:  base,
		name:  name,
		lines: make([]int, 0, 16),
		size:  size,
	}
}

func (f *File) AddLine(offset int) {
	if offset >= f.base-1 && offset < f.base+f.size {
		f.lines = append(f.lines, offset)
	}
}

func (f *File) Base() int {
	return f.base
}

func (f *File) Pos(offset int) Pos {
	if offset < 0 || offset >= f.size {
		panic("illegal file offset")
	}
	return Pos(f.base + offset)
}

func (f *File) Position(p Pos) Position {
	col, row := int(p), 1

	for i, nl := range f.lines {
		if p > f.Pos(nl) {
			col, row = int(p-f.Pos(nl)), i+1
		}
	}

	return Position{Filename: f.name, Col: col, Row: row}
}

func (f *File) Size() int {
	return f.size
}
