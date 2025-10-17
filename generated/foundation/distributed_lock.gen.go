// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [DistributedLock] class.
var DistributedLockClass objc.Class

func init() {
	DistributedLockClass = objc.GetClass("NSDistributedLock")
}

type DistributedLock struct {
	objc.ID
}

func DistributedLockFrom(ptr unsafe.Pointer) DistributedLock {
	return DistributedLock{
		ID: objc.ID(ptr),
	}
}



