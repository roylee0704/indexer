package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// A plain hit consists of a capitalization bit, font size,
// and 12 bits of word position in a document
type hit struct {
	wordPos int // 12bits
}

// termVector represents term->hit-list / term -> (freq, offset, position)
type termVector map[string][]hit

// implements stringer()
func (tf termVector) String() string {

	var s string
	for term, hits := range tf {
		s += fmt.Sprintf("[%s]: {size: %d, hits: %v}\n", term, len(hits), hits)
	}
	return s
}

// harvest consumes data from reader (line by line), and
// splits out termVector -- (term - []hits)
func harvest(r io.Reader) termVector {

	freq := make(termVector)

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	var token string
	var c = 1

	for scanner.Scan() {
		token = strings.ToLower(scanner.Text())
		if _, ok := freq[token]; !ok {
			//	freq[token] = newHitPack()
		}

		//freq[token].add(hit{pos: c})

		c++
	}
	return freq
}
