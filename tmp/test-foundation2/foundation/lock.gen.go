// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var lockClass _LockClass

func init() {
	lockClass = _LockClass{objc.GetClass("NSLock")}
}

type _LockClass struct {
	class objc.Class
}

type Lock struct {
	objc.ID
}

func LockFrom(ptr unsafe.Pointer) Lock {
	return Lock{
		ID: objc.ID(ptr),
	}
}




