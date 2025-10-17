// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Shadow] class.
var shadowClass = _ShadowClass{objc.GetClass("NSShadow")}

type _ShadowClass struct {
	class objc.Class
}

// An object you use to specify attributes to create and style a drop shadow during drawing operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSShadow

type Shadow struct {
	objectivec.Object
}

// ShadowFrom constructs a [Shadow] from an unsafe.Pointer.
//
// An object you use to specify attributes to create and style a drop shadow during drawing operations.
func ShadowFrom(ptr unsafe.Pointer) Shadow {
	return Shadow{objectivec.Object{objc.ID(ptr)}}
}



