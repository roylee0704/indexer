package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"unicode/utf8"
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

// harvest consumes data from reader using typical bufio.ScanWords and
// splits out termVector -- (term - []hits)
func harvest(r io.Reader) termVector {

	tv := make(termVector)

	scanner := bufio.NewScanner(r)
	scanner.Split(ScanTerms)

	var token string
	var c = 1

	for scanner.Scan() {
		token = strings.ToLower(scanner.Text())
		tv[token] = append(tv[token], hit{wordPos: c})
		c++
	}
	return tv
}

// isControlBreak checks if a rune is a control breaker.
// - isSpace
// - isPunctuation
func isControlBreak(r rune) bool {
	return isSpace(r) || isPunctuation(r) //space automatically be control breaker.
}

func isSignificant(r rune) bool {
	return false
}

// ScanTerms is a split function for a Scanner that returns each
// space-separated word of text, with surrounding space deleted. The definition
// of space is set by unicode.IsSpace.
//
// On top of that, it might also return punctuation-seperated word of text.
// Definition of punctuation is set by significant & insignificant.
//
// cases (simplified): return (advance, token, err)
// 1. "ABC " - spliter found!: return indexOf(last-term-char), token.
//
// if spliter not found:
// 2a. "ABC" - !atEOF: request for more data.
// 2b. "ABC" - atEOF: return token, err = finalToken
//
// advance: number of characters(if not ascii, + width) that you have visited.
func ScanTerms(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// ignore spaces
	var start int
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if !isControlBreak(r) {
			break
		}
	}

	// seek for control-break.
	for i := start; i < len(data); {
		r, width := utf8.DecodeRune(data[i:])
		if isControlBreak(r) {
			return i + width, data[start:i], nil
		}
		i += width
	}

	// no control-break found, check for EOF.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	return start, nil, nil
}

// isPunctuation serves as control-break checker, atm while insignificant and
// significant hasn't been implemented, temporarily, all punctuations are
// considered control-break.
func isPunctuation(r rune) bool {

	if r == '\'' || r == '-' {
		return false
	}
	if r >= '!' && r <= '/' ||
		r >= ':' && r <= '@' ||
		r >= '[' && r <= '`' ||
		r >= '{' && r <= '~' {
		return true
	}

	return false
}

// isSpace reports whether the character is a Unicode white space character.
// We avoid dependency on the unicode package, but check validity of the implementation
// in the tests.
func isSpace(r rune) bool {
	if r <= '\u00FF' {
		// Obvious ASCII ones: \t through \r plus space. Plus two Latin-1 oddballs.
		switch r {
		case ' ', '\t', '\n', '\v', '\f', '\r':
			return true
		case '\u0085', '\u00A0':
			return true
		}
		return false
	}
	// High-valued ones.
	if '\u2000' <= r && r <= '\u200a' {
		return true
	}
	switch r {
	case '\u1680', '\u2028', '\u2029', '\u202f', '\u205f', '\u3000':
		return true
	}
	return false
}
