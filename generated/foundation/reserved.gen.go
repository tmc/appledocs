// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [reserved] class.
var reservedClass = _reservedClass{objc.GetClass("reserved")}

type _reservedClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPredicate/reserved

type reserved struct {
	objectivec.Object
}

// reservedFrom constructs a [reserved] from an unsafe.Pointer.
func reservedFrom(ptr unsafe.Pointer) reserved {
	return reserved{objectivec.Object{objc.ID(ptr)}}
}



