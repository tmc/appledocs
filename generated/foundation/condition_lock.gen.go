// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ConditionLock] class.
var ConditionLockClass objc.Class

func init() {
	ConditionLockClass = objc.GetClass("NSConditionLock")
}

type ConditionLock struct {
	objc.ID
}

func ConditionLockFrom(ptr unsafe.Pointer) ConditionLock {
	return ConditionLock{
		ID: objc.ID(ptr),
	}
}



