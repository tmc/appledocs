// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [eccVersion] class.
var eccVersionClass = _eccVersionClass{objc.GetClass("eccVersion")}

type _eccVersionClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/eccVersion-c.ivar

type eccVersion struct {
	objectivec.Object
}

// eccVersionFrom constructs a [eccVersion] from an unsafe.Pointer.
func eccVersionFrom(ptr unsafe.Pointer) eccVersion {
	return eccVersion{objectivec.Object{objc.ID(ptr)}}
}



