// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [numBytes] class.
var numBytesClass = _numBytesClass{objc.GetClass("numBytes")}

type _numBytesClass struct {
	class objc.Class
}

type numBytes struct {
	objc.ID
}

func numBytesFrom(ptr unsafe.Pointer) numBytes {
	return numBytes{
		ID: objc.ID(ptr),
	}
}




