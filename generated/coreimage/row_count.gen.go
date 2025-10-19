// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [rowCount] class.
var rowCountClass = _rowCountClass{objc.GetClass("rowCount")}

type _rowCountClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIPDF417CodeDescriptor/rowCount-c.ivar

type rowCount struct {
	objectivec.Object
}

// rowCountFrom constructs a [rowCount] from an unsafe.Pointer.
func rowCountFrom(ptr unsafe.Pointer) rowCount {
	return rowCount{objectivec.Object{objc.ID(ptr)}}
}



