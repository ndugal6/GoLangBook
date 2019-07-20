//Dup3, unlike 1 & 2, read the ent re input into memory in one big gulp, split it into lines all at once, then process the lines. 
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	lineCounts := make(map[string]int)
	wordCounts := make(map[string]int)
	for _, filename:= range os.Args[1:] {
		data, err := ioutil.ReadFile(filename) //ReadFile returns a byte slice that must be converted into a string so it can be split by strings.Split
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			lineCounts[line]++
			for _, word := range strings.Split(string(line), " ") {
				wordCounts[word]++
			}
		}
	}
	for line, n := range lineCounts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
	for word, n := range wordCounts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, word)
		}
	}
}