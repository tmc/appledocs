// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [numBytes] class.
var numBytesClass objc.Class

func init() {
	numBytesClass = objc.GetClass("numBytes")
}

type numBytes struct {
	objc.ID
}

func numBytesFrom(ptr unsafe.Pointer) numBytes {
	return numBytes{
		ID: objc.ID(ptr),
	}
}




