package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func køyrKøyr(argumenter ...string) {
	// ta utgangspunkt i full kontroll her, som om me var i main()

	var feil error

	// kontroller fyrste argument (må vere ei .gå-fil)
	fil := argumenter[0]
	if !strings.HasSuffix(fil, ".gå") {
		fmt.Fprintln(os.Stderr, "fyrste argument må ver ei .gå-fil")
		os.Exit(1)
	}
	filstoda, feil := os.Stat(fil)
	if feil != nil {
		fmt.Fprintln(os.Stderr, feil)
		os.Exit(1)
	}
	if !filstoda.Mode().IsRegular() {
		fmt.Fprintln(os.Stderr, "inga vanleg fil:", fil)
		os.Exit(1)
	}

	// me kan fjerna fila frå `argumenter` nå
	argumenter = argumenter[1:]

	// åpne .gå-fil
	inn, feil := os.Open(fil)
	if feil != nil {
		fmt.Fprintln(os.Stderr, feil)
		os.Exit(1)
	}

	// lag mellombels katalog
	mellombelsKatalog, feil := ioutil.TempDir("", "gå-køyr")
	if feil != nil {
		fmt.Fprintln(os.Stderr, feil)
		os.Exit(1)
	}
	defer os.RemoveAll(mellombelsKatalog)

	// lag .go-fil
	nyFilnavn := strings.TrimSuffix(filepath.Base(fil), ".gå") + ".go"
	nyFil := filepath.Join(mellombelsKatalog, nyFilnavn)
	ut, feil := os.Create(nyFil)
	if feil != nil {
		fmt.Fprintln(os.Stderr, feil)
		os.Exit(1)
	}
	defer ut.Close()

	// omset
	feil = Omset(inn, ut)
	if feil != nil {
		fmt.Fprintln(os.Stderr, feil)
		os.Exit(1)
	}

	// køyr `go run <fil> <resten av argumenter...>`
	kommando := exec.Command("go", append([]string{"run", ut.Name()}, argumenter...)...)
	kommando.Stdin = os.Stdin
	kommando.Stdout = os.Stdout
	kommando.Stderr = os.Stderr
	feil = kommando.Run()
	if feil != nil {
		fmt.Fprintln(os.Stderr, feil)
		os.Exit(1)
	}
}
