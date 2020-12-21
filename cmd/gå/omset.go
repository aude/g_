package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Omsetjing struct {
	fastOmgrep *regexp.Regexp
	ombyte     []byte
}

func NyOmsetjing(anglisk, nynorsk string) *Omsetjing {
	omsetjing := new(Omsetjing)
	omsetjing.fastOmgrep = regexp.MustCompile(`([^\p{L}\d]|^)` + anglisk + `([^\p{L}\d]|$)`)
	omsetjing.ombyte = []byte("${1}" + nynorsk + "${2}")
	return omsetjing
}

var Omsetjingar = []*Omsetjing{
	NyOmsetjing("hop", "package"),
	NyOmsetjing("høve", "func"),
	NyOmsetjing("lånord", "import"),
	NyOmsetjing("skildre", "var"),
	NyOmsetjing("fast", "const"),
	NyOmsetjing("slag", "type"),
	NyOmsetjing("att", "return"),
	NyOmsetjing("medan", "for"),
	NyOmsetjing("famn", "range"),
	NyOmsetjing("viss", "if"),
	NyOmsetjing("elles", "else"),
	NyOmsetjing("ymse", "switch"),
	NyOmsetjing("lei", "case"),
	NyOmsetjing("korkje", "default"),
	NyOmsetjing("dett", "fallthrough"),
	NyOmsetjing("kaffipause", "break"),
	NyOmsetjing("fortset", "continue"),
	NyOmsetjing("seinare", "defer"),
	NyOmsetjing("skriv", "print"),
	NyOmsetjing("skrivlnj", "println"),
	NyOmsetjing("attåt", "append"),
	NyOmsetjing("tal", "len"),
	NyOmsetjing("ny", "new"),
	NyOmsetjing("lag", "make"),
	NyOmsetjing("døy", "panic"),
	NyOmsetjing("gå", "go"),
	NyOmsetjing("heiltal", "int"),
	NyOmsetjing("flyttal", "float"),
	NyOmsetjing("tekst", "string"),
	NyOmsetjing("bete", "byte"),
	NyOmsetjing("teikn", "rune"),
	NyOmsetjing("kanskje", "bool"),
	NyOmsetjing("sant", "true"),
	NyOmsetjing("usant", "false"),
	NyOmsetjing("null", "nil"),
	NyOmsetjing("kart", "map"),
	NyOmsetjing("drag", "chan"),
	NyOmsetjing("bokmål", "error"),
}

func Omset(inn io.Reader, ut io.Writer) error {
	søkjar := bufio.NewScanner(inn)
	for søkjar.Scan() {
		linje := søkjar.Bytes()
		for _, omsetjing := range Omsetjingar {
			linje = omsetjing.fastOmgrep.ReplaceAll(linje, omsetjing.ombyte)
		}
		_, feil := ut.Write(append(linje, '\n'))
		if feil != nil {
			return fmt.Errorf("skrivefeil: %q", feil)
		}
	}
	if feil := søkjar.Err(); feil != nil {
		return fmt.Errorf("lesefeil: %q", feil)
	}
	return nil
}

func køyrOmset(argumenter ...string) {
	// rekna med full kontroll her, som om me var i main()

	// berre hald aksept for katalogar og filer
	katalogar := []string{}
	filer := []string{}
	for _, argument := range argumenter {
		stoda, feil := os.Stat(argument)
		if feil == nil && stoda.Mode().IsDir() {
			katalogar = append(katalogar, argument)
		} else if feil == nil && stoda.Mode().IsRegular() {
			filer = append(filer, argument)
		} else {
			fmt.Fprintf(os.Stderr, "%q er korkje katalog eller fil\n", argument)
			os.Exit(1)
		}
	}

	// viss ingen argumenter, gå til pipemodus
	if len(katalogar) == 0 && len(filer) == 0 {
		feil := Omset(os.Stdin, os.Stdout)
		if feil != nil {
			fmt.Fprintln(os.Stderr, feil)
			os.Exit(1)
		}
	} else {
		// omset .gå-filer til .go-filer

		// legg filer frå directories til `filer`
		for _, katalog := range katalogar {
			handtak, feil := os.Open(katalog)
			if feil != nil {
				fmt.Fprintln(os.Stderr, feil)
				os.Exit(1)
			}
			defer handtak.Close()

			filinfo, feil := handtak.Readdir(-1)
			if feil != nil {
				fmt.Fprintln(os.Stderr, feil)
				os.Exit(1)
			}
			for _, fil := range filinfo {
				filer = append(filer, filepath.Join(katalog, fil.Name()))
			}
		}

		// omset tilførte .gå-filer
		for _, fil := range filer {
			if !strings.HasSuffix(fil, ".gå") {
				continue
			}

			var feil error

			inn, feil := os.Open(fil)
			if feil != nil {
				fmt.Fprintln(os.Stderr, feil)
				fmt.Fprintln(os.Stderr, "høpp øve...")
				continue
			}
			defer inn.Close()

			nyFil := strings.TrimSuffix(fil, ".gå") + ".go"
			ut, feil := os.Create(nyFil)
			if feil != nil {
				fmt.Fprintln(os.Stderr, feil)
				fmt.Fprintln(os.Stderr, "høpp øve...")
				continue
			}
			defer ut.Close()

			feil = Omset(inn, ut)
			if feil != nil {
				fmt.Fprintln(os.Stderr, feil)
				continue
			}
		}
	}
}
