// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var distributedLockClass _DistributedLockClass

func init() {
	distributedLockClass = _DistributedLockClass{objc.GetClass("NSDistributedLock")}
}

type _DistributedLockClass struct {
	class objc.Class
}

type DistributedLock struct {
	objc.ID
}

func DistributedLockFrom(ptr unsafe.Pointer) DistributedLock {
	return DistributedLock{
		ID: objc.ID(ptr),
	}
}




