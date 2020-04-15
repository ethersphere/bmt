// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command generate_legacy generates bmt hashes of sequential byte inputs
// for every possible length of legacy bmt hasher
package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gitlab.com/nolash/go-mockbytes"
)

func main() {

	// create output directory, fail if it already exists or error creating
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: generate-hashes <output_directory>\n")
		os.Exit(1)
	}
	outputDir := os.Args[1]
	err := os.Mkdir(outputDir, 0755)
	if err == os.ErrExist {
		fmt.Fprintf(os.Stderr, "Directory %s already exists\n", outputDir)
		os.Exit(1)
	} else if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	// create sequence generator and outputs
	var i int
	g := mockbytes.New(0, mockbytes.MockTypeStandard).WithModulus(255)
	for i = 0; i < 4096; i++ {
		s := fmt.Sprintf("processing %d...", i)
		fmt.Fprintf(os.Stderr, "%-64s\r", s)
		filename := fmt.Sprintf(".data/%d.bin", i)
		b, err := g.SequentialBytes(i)
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
		}
		err = ioutil.WriteFile(filename, b, 0644)
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
		}
		err = ioutil.WriteFile(filename, b, 0644)
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
		}
	}

	// Be kind and give feedback to user
	fmt.Printf("%-64s\n", "Done. Data is in .data. Enjoy!")
}
