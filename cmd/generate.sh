#!/bin/bash

# Copyright 2020 The Swarm Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# This script generates bmt hashes for all possible input lengths

# prepare working directory
wd=$(realpath $(dirname ${BASH_SOURCE[0]}))
pushd $wd > /dev/null
mkdir -p .data
go build -o ./_bmt ./main_legacy.go

# generate data to be hashed
data=''
>&2 echo -ne 		"preparing data...\r"
for i in {0..4095}; do
	n=`printf %02x $((i%255))`
	data="${data}\x${n}"
done
echo -ne ${data} > .data/src.bin

# hash data and write to data directory
#for i in {1..4096}; do
for i in {1..4096}; do
	>&2 echo -ne 	"processing ${i}                                          \r"
	dd if=.data/src.bin bs=1 count=${i} 2> /dev/null | ./_bmt > .data/${i}.bin
done

>&2 echo -ne 		"hashing zero-length input                                \r"
echo -n | ./_bmt > .data/empty.bin
>&2 echo -ne 		"hashing 'foo'                                            \r"
echo -n "foo" | ./_bmt > .data/foo.bin
>&2 echo -ne 		"hashing 'foo\\\n'                                        \r"
echo "foo" | ./_bmt > .data/foo_lf.bin
>&2 echo -ne 		"hashing 'foo\\\r\\\n'                                    \r"
echo -e "foo\r" | ./_bmt > .data/foo_crlf.bin

# clean up
unlink .data/src.bin
unlink ./_bmt
echo 			"Done. Data is in ${wd}/.data. Enjoy!                       "
popd > /dev/null
