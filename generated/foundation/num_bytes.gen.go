// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [numBytes] class.
var numBytesClass = _numBytesClass{objc.GetClass("numBytes")}

type _numBytesClass struct {
	class objc.Class
}

// An interface definition for the [numBytes] class.
type InumBytes interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSimpleCString/numBytes

type numBytes struct {
	objectivec.Object
}

// numBytesFrom constructs a [numBytes] from an unsafe.Pointer.
func numBytesFrom(ptr unsafe.Pointer) numBytes {
	return numBytes{objectivec.Object{objc.ID(ptr)}}
}



