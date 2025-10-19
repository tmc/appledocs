// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Color] class.
var (
	colorClass     _ColorClass
	colorClassOnce sync.Once
)

func getColorClass() _ColorClass {
	colorClassOnce.Do(func() {
		colorClass = _ColorClass{objc.GetClass("NSColor")}
	})
	return colorClass
}

type _ColorClass struct {
	class objc.Class
}

// An interface definition for the [Color] class.
type IColor interface {
	objectivec.IObject
	ColorUsingColorSpaceName(name unsafe.Pointer) unsafe.Pointer
}

// An object that stores color data and sometimes opacity (alpha value). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor

type Color struct {
	objectivec.Object
}

// ColorFrom constructs a [Color] from an unsafe.Pointer.
//
// An object that stores color data and sometimes opacity (alpha value).
func ColorFrom(ptr unsafe.Pointer) Color {
	return Color{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (cc _ColorClass) Alloc() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (cc _ColorClass) New() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ Color) Init() Color {
	rv := objc.Send[Color](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ Color) Autorelease() Color {
	rv := objc.Send[Color](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewColor creates a new Color instance.
func NewColor() Color {
	return getColorClass().New()
}


// Creates a new color object whose color is the same as the receiver’s, except that the new color object is in the specified color space. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/usingColorSpaceName(_:)
func (c_ Color) ColorUsingColorSpaceName(name unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("colorUsingColorSpaceName:"), name)
	return rv
}


