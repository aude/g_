package main

import (
	"fmt"
	"os"
)

func Hjelpmelding() string {
	return fmt.Sprintf(`Go på nynorsk.

Høve:

	gå omset [katalog...] [fil...]
	gå køyr <fil> [argument...]

Til dømes:

	$ gå omset .
	$ gå omset /src/a /src/b
	$ gå omset framifrå.gå samskap.gå oppslag.gå
	$ gå køyr main.gå
	$ gå køyr main.gå -h
	$ gå køyr main.gå input.txt > output.txt
`)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, Hjelpmelding())
		os.Exit(1)
	}

	switch os.Args[1] {
	case "hjelp", "stønad", "naud", "-h", "--help", "--usage":
		fmt.Println(Hjelpmelding())
	case "omset":
		køyrOmset(os.Args[2:]...)
	case "køyr":
		køyrKøyr(os.Args[2:]...)
	default:
		fmt.Fprintln(os.Stderr, Hjelpmelding())
		os.Exit(1)
	}
}
