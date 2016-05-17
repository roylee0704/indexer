package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {

	tf := harvest(strings.NewReader("hello hello world"))
	fmt.Printf("%v\n", tf)

	fd, _ := os.Open("fixtures/short.txt")
	tf = harvest(fd)
	fmt.Printf("%v\n", tf)

	//fmt.Printf("%d\n", len(tf))

}
