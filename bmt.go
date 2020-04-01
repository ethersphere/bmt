package bmt

import (
	"hash"
)

// BMT provides the necessary extension of the hash interface to add the length-prefix of the BMT hash
//
// Any implementation should make it possible to generate a BMT hash using the hash.Hash interface only. However, the limitation will be that the Span of the BMT hash always must be limited to the amount of bytes actually written.
type BMTHash interface {
	hash.Hash

	// SetSpan sets the length prefix of BMT hash.
	SetSpan(int64) error

	// Maximum number of BlockSize() units of data that this BMT hasher will process.
	MaxSections() int

	// WriteSection writes to a specific section of the data to be hashed.
	WriteSection(idx int, data []byte) error
}
