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
// Alloc allocates a new instance without initialization.
func (cc _ColorListClass) Alloc() ColorList {
	rv := objc.Send[ColorList](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (cc _ColorListClass) New() ColorList {
	rv := objc.Send[ColorList](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ColorList) Init() ColorList {
	rv := objc.Send[ColorList](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ColorList) Autorelease() ColorList {
	rv := objc.Send[ColorList](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewColorList creates a new ColorList instance.
func NewColorList() ColorList {
	return colorListClass.New()
}




