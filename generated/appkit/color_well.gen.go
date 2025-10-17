
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ColorWell] class.
var ColorWellClass _ColorWellClass

func init() {
	ColorWellClass = _ColorWellClass{objc.GetClass("NSColorWell")}
}

type _ColorWellClass struct {
	objc.Class
}

// An interface definition for the [ColorWell] class.
type IColorWell interface {
	ID() objc.ID
}

type ColorWell struct {
	id objc.ID
}

func ColorWellFrom(ptr unsafe.Pointer) ColorWell {
	return ColorWell{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ ColorWell) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _ColorWellClass) Alloc() ColorWell {
	rv := objc.Send[ColorWell](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _ColorWellClass) New() ColorWell {
	rv := objc.Send[ColorWell](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewColorWell creates and returns a new initialized instance.
func NewColorWell() ColorWell {
	return ColorWellClass.New()
}

// Init initializes the instance.
func (c_ ColorWell) Init() ColorWell {
	rv := objc.Send[ColorWell](c_.ID(), selInit)
	return rv
}
// A Boolean value that determines whether the color well has a border. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSColorWell/isBordered
func (c_ ColorWell) Bordered() bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("bordered"))
	return rv
}
// SetBordered sets the value of the bordered property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSColorWell/isBordered
func (c_ ColorWell) SetBordered(value bool) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setBordered:"), value)
}
// The target object that defines the action you want to perform when someone interacts with the color well. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSColorWell/pulldownTarget
func (c_ ColorWell) PulldownTarget() objc.ID {
	rv := objc.Send[objc.ID](c_.ID(), objc.RegisterName("pulldownTarget"))
	return rv
}
// SetPulldownTarget sets the value of the pulldownTarget property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSColorWell/pulldownTarget
func (c_ ColorWell) SetPulldownTarget(value objc.ID) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setPulldownTarget:"), value)
}
