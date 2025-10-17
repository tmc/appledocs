// Code generated from Apple documentation for ObjectiveC. DO NOT EDIT.

package objectivec

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [isa] class.
var isaClass = _isaClass{objc.GetClass("isa")}

type _isaClass struct {
	class objc.Class
}

// An interface definition for the [isa] class.
type Iisa interface {
	IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/isa

type isa struct {
	Object
}

// isaFrom constructs a [isa] from an unsafe.Pointer.
func isaFrom(ptr unsafe.Pointer) isa {
	return isa{Object{objc.ID(ptr)}}
}



