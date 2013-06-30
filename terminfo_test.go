package terminfo

import (
	"testing"
)

func TestInfocmp(t *testing.T) {
	out, err := infocmp("")
	t.Log(string(out))
	if err != nil {
		t.Fatal(err)
	}
}

func TestGet(t *testing.T) {
	ti, err := Get()
	t.Log(ti)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTerm(t *testing.T) {
	ti, err := Term("xterm")
	t.Log(ti)
	if err != nil {
		t.Fatal(err)
	}
}
