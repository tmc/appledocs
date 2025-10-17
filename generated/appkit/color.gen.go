
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Color] class.
var ColorClass _ColorClass

func init() {
	ColorClass = _ColorClass{objc.GetClass("NSColor")}
}

type _ColorClass struct {
	objc.Class
}

// An interface definition for the [Color] class.
type IColor interface {
	ID() objc.ID
	ColorUsingColorSpaceName(name unsafe.Pointer) unsafe.Pointer
}

type Color struct {
	id objc.ID
}

func ColorFrom(ptr unsafe.Pointer) Color {
	return Color{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ Color) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _ColorClass) Alloc() Color {
	rv := objc.Send[Color](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _ColorClass) New() Color {
	rv := objc.Send[Color](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewColor creates and returns a new initialized instance.
func NewColor() Color {
	return ColorClass.New()
}

// Init initializes the instance.
func (c_ Color) Init() Color {
	rv := objc.Send[Color](c_.ID(), selInit)
	return rv
}
// Creates a new color object whose color is the same as the receiver’s, except that the new color object is in the specified color space. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSColor/usingColorSpaceName(_:)
func (c_ Color) ColorUsingColorSpaceName(name unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("colorUsingColorSpaceName:"), name)
	return rv
}
