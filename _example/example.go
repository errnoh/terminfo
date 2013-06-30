package main

import (
	"github.com/errnoh/terminfo"
	"bytes"
	"fmt"
	"log"
)


func main() {
	ti, err := terminfo.Get()
	if err != nil {
		log.Fatal(err)
	}

	b := new(bytes.Buffer)

	b.Write(ti.String["clear"])
	b.Write(ti.String["bold"])
	b.WriteString("Bold Crispy Bacon")
	b.Write(ti.String["sgr0"])
	fmt.Println(b.String())
}
