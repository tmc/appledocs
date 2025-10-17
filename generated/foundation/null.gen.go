// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Null] class.
var nullClass = _NullClass{objc.GetClass("NSNull")}

type _NullClass struct {
	class objc.Class
}

// A singleton object used to represent null values in collection objects that don’t allow values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNull

type Null struct {
	objectivec.Object
}

// NullFrom constructs a [Null] from an unsafe.Pointer.
//
// A singleton object used to represent null values in collection objects that don’t allow values.
func NullFrom(ptr unsafe.Pointer) Null {
	return Null{objectivec.Object{objc.ID(ptr)}}
}



