package main

import (
	"bufio"
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os"
)

var (
	dictFile = flag.String("d", "/usr/share/dict/words",
		"File from which to read the dictionary")
	length = flag.Int("n", 4, "Desired number of words in the password")
)

func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}

func main() {
	// Setup
	flag.Parse()
	file, err := os.Open(*dictFile)
	check(err)
	defer file.Close()
	dict := bufio.NewReader(file)

	// Count the number of words in the file.
	var nwords int64
	for r, _, err := dict.ReadRune(); err == nil; {
		if r == '\n' {
			nwords++
		}
		r, _, err = dict.ReadRune()
	}

	// Make some storage for our password
	result := make([]string, *length)

	for i := 0; i < *length; i++ {
		// Seek back to the beginning of the file to generate a new word.
		_, err := file.Seek(0, 0)
		check(err)

		// Make a new buffer.Reader; we can't seek on one of these as
		// best as I can tell.
		dict = bufio.NewReader(file)

		// Generate a random index into the file...
		jthword, err := rand.Int(rand.Reader, big.NewInt(nwords))
		check(err)

		// ...and go find it.
		for j := int64(0); j < jthword.Int64(); {
			r, _, err := dict.ReadRune()
			check(err)
			if r == '\n' {
				j++
			}
		}

		// Okay, we've found our word. now let's read it in and add it to
		// the password.
		result[i], err = dict.ReadString('\n')
		check(err)
	}

	// We're done! print it out.
	for _, v := range result {
		fmt.Print(v)
	}
}
