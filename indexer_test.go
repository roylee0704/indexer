package main

import (
	"fmt"
	"os"
	"testing"
)

func TestSplit(t *testing.T) {

	//tf := harvest(strings.NewReader("hello hello world"))
	//	fmt.Printf("%v\n", tf)

	fd, _ := os.Open("fixtures/short.txt")
	tf := harvest(fd)
	fmt.Printf("%v\n", tf)

}

// TestSkipLeadingSpace tests: given a sentence, trim all spaces,
func TestSkipLeadingSpace(t *testing.T) {
	// fmt.Println(ScanTerms([]byte("   A "), false))
	// fmt.Println(ScanTerms([]byte("   A"), false))
	// fmt.Println(ScanTerms([]byte("   A"), true))

}

func TestIsPunctuation(t *testing.T) {

}
