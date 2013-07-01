// Copyright 2013 errnoh. All rights reserved.
// Use of this source code is governed by
// MIT License that can be found in the LICENSE file.

// Wrapper around infocmp command, returns terminfo data.
package terminfo

// Linux Terminfo and Termcap Commands:
// http://comptechdoc.org/os/linux/howlinuxworks/linux_hltermcommands.html

import (
	"bytes"
	"io"
	"os/exec"
	"strconv"
)

// Get returns Terminfo of current terminal. Set termcap if you want to use termcap format instead of terminfo (default).
func Get(termcap bool) (*Terminfo, error) {
	return Term("", termcap)
}

// Term returns Terminfo of terminal term. Set termcap if you want to use termcap format instead of terminfo (default).
func Term(term string, termcap bool) (*Terminfo, error) {
	out, err := infocmp(term, termcap)
	if err != nil {
		return nil, err
	}

	b := bytes.NewBuffer(out)
	return parse(b, termcap)
}

type Terminfo struct {
	Description string
	String      map[string][]byte
	Numeric     map[string]int
	Boolean     map[string]bool
}

func infocmp(s string, termcap bool) ([]byte, error) {
	args := []string{"-1"}

	if s != "" {
		args = append(args, s)
	}

	if termcap {
		args = append(args, "-C")
	}

	return exec.Command("infocmp", args...).Output()
}

func parse(b *bytes.Buffer, termcap bool) (ti *Terminfo, err error) {
	var (
		line []byte
		pair [][]byte
	)

	ti = &Terminfo{
		String:  make(map[string][]byte),
		Numeric: make(map[string]int),
		Boolean: make(map[string]bool),
	}

	for {
		line, err = b.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		} else if len(line) <= 2 {
			continue
		} else if line[0] == '#' {
			continue
		}
		line = bytes.TrimSpace(line)
		if termcap {
			line = bytes.TrimSuffix(line, []byte(":\\"))
			line = bytes.TrimPrefix(line, []byte{':'})
		} else {
			line = bytes.TrimSuffix(line, []byte{','})
		}
		if ti.Description == "" {
			ti.Description = string(line)
			continue
		}

		pair = bytes.Split(line, []byte{'='})
		switch len(pair) {
		case 1:
			pair = bytes.Split(line, []byte{'#'})
			switch len(pair) {
			case 1:
				ti.Boolean[string(line)] = true
			case 2:
				v, err := strconv.Atoi(string(pair[1]))
				if err != nil {
					ti.Boolean[string(line)] = true
				}
				ti.Numeric[string(pair[0])] = v
			}
		case 2:
			pair[1] = bytes.Replace(pair[1], []byte(`\E`), []byte{'\033'}, -1)
			pair[1] = bytes.Replace(pair[1], []byte(`\n`), []byte{'\n'}, -1)
			pair[1] = bytes.Replace(pair[1], []byte(`\r`), []byte{'\r'}, -1)
			pair[1] = bytes.Replace(pair[1], []byte(`\t`), []byte{'\t'}, -1)
			pair[1] = bytes.Replace(pair[1], []byte(`\b`), []byte{'\b'}, -1)
			ti.String[string(pair[0])] = pair[1]
		default:
			continue
		}
	}

	return
}
