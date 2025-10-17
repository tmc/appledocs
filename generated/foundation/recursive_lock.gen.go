// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [RecursiveLock] class.
var RecursiveLockClass objc.Class

func init() {
	RecursiveLockClass = objc.GetClass("NSRecursiveLock")
}

type RecursiveLock struct {
	objc.ID
}

func RecursiveLockFrom(ptr unsafe.Pointer) RecursiveLock {
	return RecursiveLock{
		ID: objc.ID(ptr),
	}
}



