// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FileAccessIntent] class.
var fileAccessIntentClass = _FileAccessIntentClass{objc.GetClass("NSFileAccessIntent")}

type _FileAccessIntentClass struct {
	class objc.Class
}

// An interface definition for the [FileAccessIntent] class.
type IFileAccessIntent interface {
	objectivec.IObject
}

// The details of a coordinated-read or coordinated-write operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileAccessIntent

type FileAccessIntent struct {
	objectivec.Object
}

// FileAccessIntentFrom constructs a [FileAccessIntent] from an unsafe.Pointer.
//
// The details of a coordinated-read or coordinated-write operation.
func FileAccessIntentFrom(ptr unsafe.Pointer) FileAccessIntent {
	return FileAccessIntent{objectivec.Object{objc.ID(ptr)}}
}



