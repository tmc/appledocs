// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var LockClass _LockClass

func init() {
	LockClass = _LockClass{objc.GetClass("NSLock")}
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




