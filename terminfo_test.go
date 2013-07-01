package terminfo

import (
	"testing"
)

func TestInfocmp(t *testing.T) {
	out, err := infocmp("", false)
	t.Log(string(out))
	if err != nil {
		t.Fatal(err)
	}
}

func TestInfocmpTermcap(t *testing.T) {
	out, err := infocmp("", true)
	t.Log(string(out))
	if err != nil {
		t.Fatal(err)
	}
}

func TestGet(t *testing.T) {
	ti, err := Get(false)
	t.Log(ti)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetTermcap(t *testing.T) {
	ti, err := Get(true)
	t.Log(ti)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTerm(t *testing.T) {
	ti, err := Term("xterm", false)
	t.Log(ti)
	if err != nil {
		t.Fatal(err)
	}
}
