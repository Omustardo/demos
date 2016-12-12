// bytecoder provides functions for converting standard arrays into byte arrays.
// Based on the Bytes function from golang.org\x\mobile\exp\f32\f32.go
package bytecoder

import (
	"encoding/binary"
	"fmt"
	"math"
)

// getByteOrder returns false for LittleEndian and true for BigEndian. It panics if the provided parameter isn't one of the two.
func getByteOrder(byteOrder binary.ByteOrder) bool {
	le := false
	switch byteOrder {
	case binary.BigEndian:
	case binary.LittleEndian:
		le = true
	default:
		panic(fmt.Sprintf("invalid byte order %v", byteOrder))
	}
	return le
}

// Bytes returns the byte representation of float32 values in the given byte
// order. byteOrder must be either binary.BigEndian or binary.LittleEndian.
func Float32(byteOrder binary.ByteOrder, values ...float32) []byte {
	le := getByteOrder(byteOrder)
	width := 4
	b := make([]byte, width*len(values))
	for i, v := range values {
		u := math.Float32bits(v)
		if le {
			b[width*i+0] = byte(u >> 0)
			b[width*i+1] = byte(u >> 8)
			b[width*i+2] = byte(u >> 16)
			b[width*i+3] = byte(u >> 24)
		} else {
			b[width*i+0] = byte(u >> 24)
			b[width*i+1] = byte(u >> 16)
			b[width*i+2] = byte(u >> 8)
			b[width*i+3] = byte(u >> 0)
		}
	}
	return b
}
