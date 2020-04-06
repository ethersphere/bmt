// Command hash executes the BMT hash algorithm on the given data and writes the binary result to standard output
//
// Up to 4096 bytes will be read
//
// If a filename is given as argument, it reads data from the file. Otherwise it reads data from standard input.
package main

import (
	"fmt"
	"os"

	"github.com/ethersphere/bmt/reference"
	"golang.org/x/crypto/sha3"
)

func main() {
	var data [4096]byte
	var err error
	var infile *os.File

	if len(os.Args) > 1 {
		infile, err = os.Open(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(1)
		}
	} else {
		infile = os.Stdin
	}
	var c int
	c, err = infile.Read(data[:])

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		infile.Close()
		os.Exit(1)
	}
	infile.Close()

	hash := sha3.NewLegacyKeccak256()
	bmtHash := reference.NewRefHasher(hash, 128)
	binSum := bmtHash.Hash(data[:c])
	_, err = os.Stdout.Write(binSum)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
