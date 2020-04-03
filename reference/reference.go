// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reference

import (
	"hash"
)

// RefHasher is the non-optimized easy-to-read reference implementation of BMT.
type RefHasher struct {
	maxDataLength int       // c * hashSize, where c = 2 ^ ceil(log2(count)), where count = ceil(length / hashSize)
	sectionLength int       // 2 * hashSize
	hasher        hash.Hash // base hash func (Keccak256 SHA3)
}

// NewRefHasher returns a new RefHasher.
func NewRefHasher(h hash.Hash, count int) *RefHasher {
	hashsize := h.Size()
	c := 2
	for ; c < count; c *= 2 {
	}
	return &RefHasher{
		sectionLength: 2 * hashsize,
		maxDataLength: c * hashsize,
		hasher:        h,
	}
}

// Hash returns the BMT hash of the byte slice.
func (rh *RefHasher) Hash(data []byte) []byte {
	// if data is shorter than the base length (maxDataLength), we provide padding with zeros
	d := make([]byte, rh.maxDataLength)
	length := len(data)
	if length > rh.maxDataLength {
		length = rh.maxDataLength
	}
	copy(d, data[:length])
	return rh.hash(d, rh.maxDataLength)
}

// data has length maxDataLength = segmentSize * 2^k
// hash calls itself recursively on both halves of the given slice
// concatenates the results, and returns the hash of that
// if the length of d is 2 * segmentSize then just returns the hash of that section
func (rh *RefHasher) hash(data []byte, length int) []byte {
	var section []byte
	if length == rh.sectionLength {
		// section contains two data segments (d)
		section = data
	} else {
		// section contains hashes of left and right BMT subtreea
		// to be calculated by calling hash recursively on left and right half of d
		length /= 2
		section = append(rh.hash(data[:length], length), rh.hash(data[length:], length)...)
	}
	rh.hasher.Reset()
	rh.hasher.Write(section)
	return rh.hasher.Sum(nil)
}
