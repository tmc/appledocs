// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var conditionLockClass _ConditionLockClass

func init() {
	conditionLockClass = _ConditionLockClass{objc.GetClass("NSConditionLock")}
}

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




