// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ConditionLock] class.
var ConditionLockClass = _ConditionLockClass{objc.GetClass("NSConditionLock")}

type _ConditionLockClass struct {
	class objc.Class
}

type ConditionLock struct {
	objc.ID
}

func ConditionLockFrom(ptr unsafe.Pointer) ConditionLock {
	return ConditionLock{
		ID: objc.ID(ptr),
	}
}




