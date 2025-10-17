// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [RecursiveLock] class.
var recursiveLockClass = _RecursiveLockClass{objc.GetClass("NSRecursiveLock")}

type _RecursiveLockClass struct {
	class objc.Class
}

// An interface definition for the [RecursiveLock] class.
type IRecursiveLock interface {
	objectivec.IObject
}

// A lock that may be acquired multiple times by the same thread without causing a deadlock. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRecursiveLock

type RecursiveLock struct {
	objectivec.Object
}

// RecursiveLockFrom constructs a [RecursiveLock] from an unsafe.Pointer.
//
// A lock that may be acquired multiple times by the same thread without causing a deadlock.
func RecursiveLockFrom(ptr unsafe.Pointer) RecursiveLock {
	return RecursiveLock{objectivec.Object{objc.ID(ptr)}}
}



