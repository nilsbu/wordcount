package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"unicode"
)

func main() {
	dat, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	printStats(string(dat))
}

func splitWords(s string) []string {
	f := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	return strings.FieldsFunc(s, f)
}

func countUnique(ss []string, n []int) []int {
	words := make(map[string]bool)
	counts := []int{}

	nc := 0
	for i, s := range ss {
		lower := strings.ToLower(s)
		words[lower] = true

		if i+1 == n[nc] {
			counts = append(counts, len(words))
			nc++
		}
	}
	return counts
}

func printStats(s string) {
	ss := splitWords(s)
	n := []int{10000, 25000, 50000, 100000, len(ss)}
	sort.IntSlice.Sort(n)
	counts := countUnique(ss, n)

	fmt.Println("word count:", len(ss))
	for i, count := range counts {
		fmt.Println("after", n[i], "word:", count)
	}
}
