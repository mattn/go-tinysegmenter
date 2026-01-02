package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

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

func doReader(seg *tinysegmenter.TinySegmenter, in io.Reader) error {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		fmt.Println(strings.Join(seg.Segment(scanner.Text()), " "))
	}
	return scanner.Err()
}

func doFile(seg *tinysegmenter.TinySegmenter, name string) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := doReader(seg, f); err != nil {
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
		if err := doReader(seg, os.Stdin); err != nil {
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
