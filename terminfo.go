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

// Get returns Terminfo of current terminal
func Get() (*Terminfo, error) {
	out, err := infocmp("")
	if err != nil {
		return nil, err
	}

	b := bytes.NewBuffer(out)
	return parse(b)

}

// Term returns Terminfo of terminal 'term'.
func Term(term string) (*Terminfo, error) {
	out, err := infocmp(term)
	if err != nil {
		return nil, err
	}

	b := bytes.NewBuffer(out)
	return parse(b)
}

type Terminfo struct {
	Description string
	String      map[string][]byte
	Numeric     map[string]int
	Boolean     map[string]bool
}

func infocmp(s string) ([]byte, error) {
	if s != "" {
		return exec.Command("infocmp", "-1", s).Output()
	}
	return exec.Command("infocmp", "-1").Output()
}

func parse(b *bytes.Buffer) (ti *Terminfo, err error) {
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
		line = bytes.TrimSuffix(line, []byte{','})
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
