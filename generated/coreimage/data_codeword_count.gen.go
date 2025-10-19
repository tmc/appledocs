// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [dataCodewordCount] class.
var dataCodewordCountClass = _dataCodewordCountClass{objc.GetClass("dataCodewordCount")}

type _dataCodewordCountClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIAztecCodeDescriptor/dataCodewordCount-c.ivar

type dataCodewordCount struct {
	objectivec.Object
}

// dataCodewordCountFrom constructs a [dataCodewordCount] from an unsafe.Pointer.
func dataCodewordCountFrom(ptr unsafe.Pointer) dataCodewordCount {
	return dataCodewordCount{objectivec.Object{objc.ID(ptr)}}
}



