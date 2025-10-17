// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ColorList] class.
var colorListClass = _ColorListClass{objc.GetClass("NSColorList")}

type _ColorListClass struct {
	class objc.Class
}

// An interface definition for the [ColorList] class.
type IColorList interface {
	objectivec.IObject
}

// An ordered list of color objects, identified by keys. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorList

type ColorList struct {
	objectivec.Object
}

// ColorListFrom constructs a [ColorList] from an unsafe.Pointer.
//
// An ordered list of color objects, identified by keys.
func ColorListFrom(ptr unsafe.Pointer) ColorList {
	return ColorList{objectivec.Object{objc.ID(ptr)}}
}



