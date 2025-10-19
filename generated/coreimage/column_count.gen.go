// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [columnCount] class.
var columnCountClass = _columnCountClass{objc.GetClass("columnCount")}

type _columnCountClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIPDF417CodeDescriptor/columnCount-c.ivar

type columnCount struct {
	objectivec.Object
}

// columnCountFrom constructs a [columnCount] from an unsafe.Pointer.
func columnCountFrom(ptr unsafe.Pointer) columnCount {
	return columnCount{objectivec.Object{objc.ID(ptr)}}
}



