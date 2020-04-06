#!/bin/bash

# prepare working directory
wd=$(realpath $(dirname ${BASH_SOURCE[0]}))
pushd $wd > /dev/null
mkdir -p .data
go build -o ./_bmt ./main.go

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
for i in {1..2}; do
	>&2 echo -ne 	"processing ${i}                                          \r"
	dd if=.data/src.bin bs=1 count=${i} 2> /dev/null | ./_bmt > .data/${i}.bin
done
echo -n "foo" | ./_bmt > .data/foo.bin
echo "foo" | ./_bmt > .data/foo_lf.bin
echo -e "foo\r"  ./_bmt > .data/foo_crlf.bin

# clean up
unlink .data/src.bin
unlink ./_bmt
echo 			"Done. Data is in ${wd}/.data. Enjoy!                       "
popd > /dev/null
