// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Set] class.
var setClass = _SetClass{objc.GetClass("NSSet")}

type _SetClass struct {
	class objc.Class
}

// An interface definition for the [Set] class.
type ISet interface {
	objectivec.IObject
}

// A static, unordered collection of unique objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet

type Set struct {
	objectivec.Object
}

// SetFrom constructs a [Set] from an unsafe.Pointer.
//
// A static, unordered collection of unique objects.
func SetFrom(ptr unsafe.Pointer) Set {
	return Set{objectivec.Object{objc.ID(ptr)}}
}



