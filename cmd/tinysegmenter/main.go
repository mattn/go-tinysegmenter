package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mattn/go-tinysegmenter"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	seg := tinysegmenter.New()
	for scanner.Scan() {
		fmt.Println(strings.Join(seg.Segment(scanner.Text()), " "))
	}
}
