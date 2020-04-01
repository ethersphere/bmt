package bmt

import (
	"hash"
)

// BMT provides the necessary extension of the hash interface to add the length-prefix of the BMT hash
type BMT interface {
	hash.Hash
	SetLength(int)
}
