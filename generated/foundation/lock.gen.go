// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Lock] class.
var lockClass = _LockClass{objc.GetClass("NSLock")}

type _LockClass struct {
	class objc.Class
}

// An interface definition for the [Lock] class.
type ILock interface {
	objectivec.IObject
}

// An object that coordinates the operation of multiple threads of execution within the same application. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLock

type Lock struct {
	objectivec.Object
}

// LockFrom constructs a [Lock] from an unsafe.Pointer.
//
// An object that coordinates the operation of multiple threads of execution within the same application.
func LockFrom(ptr unsafe.Pointer) Lock {
	return Lock{objectivec.Object{objc.ID(ptr)}}
}



