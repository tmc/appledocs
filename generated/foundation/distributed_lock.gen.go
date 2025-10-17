// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DistributedLock] class.
var distributedLockClass = _DistributedLockClass{objc.GetClass("NSDistributedLock")}

type _DistributedLockClass struct {
	class objc.Class
}

// An interface definition for the [DistributedLock] class.
type IDistributedLock interface {
	objectivec.IObject
}

// A lock that multiple applications on multiple hosts can use to restrict access to some shared resource, such as a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDistributedLock

type DistributedLock struct {
	objectivec.Object
}

// DistributedLockFrom constructs a [DistributedLock] from an unsafe.Pointer.
//
// A lock that multiple applications on multiple hosts can use to restrict access to some shared resource, such as a file.
func DistributedLockFrom(ptr unsafe.Pointer) DistributedLock {
	return DistributedLock{objectivec.Object{objc.ID(ptr)}}
}



