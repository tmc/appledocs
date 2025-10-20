// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var RecursiveLockClass _RecursiveLockClass

func init() {
	RecursiveLockClass = _RecursiveLockClass{objc.GetClass("NSRecursiveLock")}
}

type _RecursiveLockClass struct {
	class objc.Class
}

type RecursiveLock struct {
	objc.ID
}

func RecursiveLockFrom(ptr unsafe.Pointer) RecursiveLock {
	return RecursiveLock{
		ID: objc.ID(ptr),
	}
}




