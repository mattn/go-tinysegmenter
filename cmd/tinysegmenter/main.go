package main

import (
	"bufio"
	"flag"
	"fmt"
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

func main() {
	var preserveTokens bool
	var preserveList arrayFlags
	flag.BoolVar(&preserveTokens, "p", false, "Preserve tokens in the input")
	flag.Var(&preserveList, "w", "List of tokens to preserve")
	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)
	seg := tinysegmenter.New()
	seg.SetPreserveTokens(preserveTokens)
	seg.SetPreserveList(preserveList)
	for scanner.Scan() {
		fmt.Println(strings.Join(seg.Segment(scanner.Text()), " "))
	}
}
