// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var numBytesClass _numBytesClass

func init() {
	numBytesClass = _numBytesClass{objc.GetClass("numBytes")}
}

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




