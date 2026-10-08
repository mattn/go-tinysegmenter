package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/mattn/go-tinysegmenter"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return fmt.Sprintf("%v", *i)
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

type token struct {
	text   string
	spaced bool // preceded by whitespace
}

func isDigits(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return s != ""
}

func isPoint(s string) bool {
	return s == "." || s == "．"
}

// joinDecimals joins numbers and the decimal points between them into a
// single word. Tokens separated by whitespace are never joined.
func joinDecimals(tokens []token) []string {
	var words []string
	inNumber := false
	for i := 0; i < len(tokens); i++ {
		t := tokens[i]
		if inNumber && !t.spaced && isPoint(t.text) &&
			i+1 < len(tokens) && !tokens[i+1].spaced && isDigits(tokens[i+1].text) {
			words[len(words)-1] += t.text + tokens[i+1].text
			i++
			continue
		}
		words = append(words, t.text)
		inNumber = isDigits(t.text)
	}
	return words
}

// segmentLine segments a line, dropping whitespace-only segments and
// joining decimal numbers.
func segmentLine(seg *tinysegmenter.TinySegmenter, line string) []string {
	var tokens []token
	spaced := false
	for _, s := range seg.Segment(line) {
		if text := strings.TrimSpace(s); text == "" {
			spaced = true
		} else {
			tokens = append(tokens, token{text: text, spaced: spaced})
			spaced = false
		}
	}
	return joinDecimals(tokens)
}

func doReader(seg *tinysegmenter.TinySegmenter, in io.Reader, out io.Writer) error {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		fmt.Fprintln(out, strings.Join(segmentLine(seg, scanner.Text()), " "))
	}
	return scanner.Err()
}

func doFile(seg *tinysegmenter.TinySegmenter, name string) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := doReader(seg, f, os.Stdout); err != nil {
		return err
	}
	return nil
}

func main() {
	var preserveTokens bool
	var preserveList arrayFlags
	flag.BoolVar(&preserveTokens, "p", false, "Preserve tokens in the input")
	flag.Var(&preserveList, "w", "List of tokens to preserve")
	flag.Parse()

	seg := tinysegmenter.New()
	seg.SetPreserveTokens(preserveTokens)
	seg.SetPreserveList(preserveList)

	if flag.NArg() == 0 {
		if err := doReader(seg, os.Stdin, os.Stdout); err != nil {
			log.Fatal(err)
		}
	} else {
		for _, name := range flag.Args() {
			if err := doFile(seg, name); err != nil {
				log.Fatal(err)
			}
		}
	}
}
