// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Lock] class.
var LockClass objc.Class

func init() {
	LockClass = objc.GetClass("NSLock")
}

type Lock struct {
	objc.ID
}

func LockFrom(ptr unsafe.Pointer) Lock {
	return Lock{
		ID: objc.ID(ptr),
	}
}




