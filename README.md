# go-tinysegmenter

Go port of TinySegmenter 0.2. A lightweight library for segmenting Japanese text into words.

This implementation includes the NN__ feature addition from the original, which prevents consecutive numbers from being split.

## Installation

```bash
go get github.com/mattn/go-tinysegmenter
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/mattn/go-tinysegmenter"
)

func main() {
    ts := tinysegmenter.New()
    result := ts.Segment("私の名前は中野です")
    fmt.Println(result) // [私 の 名前 は 中野 です]
}
```

## License

Modified BSD License (same as original TinySegmenter)

## Original Implementation

- [TinySegmenter](http://chasen.org/~taku/software/TinySegmenter/) by Taku Kudo

## Author

Yasuhiro Matsumoto (a.k.a mattn)
