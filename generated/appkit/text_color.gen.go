
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [textColor] class.
var textColorClass _textColorClass

func init() {
	textColorClass = _textColorClass{objc.GetClass("textColor")}
}

type _textColorClass struct {
	objc.Class
}

// An interface definition for the [textColor] class.
type ItextColor interface {
	ID() objc.ID
}

type textColor struct {
	id objc.ID
}

func textColorFrom(ptr unsafe.Pointer) textColor {
	return textColor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ textColor) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _textColorClass) Alloc() textColor {
	rv := objc.Send[textColor](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _textColorClass) New() textColor {
	rv := objc.Send[textColor](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewtextColor creates and returns a new initialized instance.
func NewtextColor() textColor {
	return textColorClass.New()
}

// Init initializes the instance.
func (t_ textColor) Init() textColor {
	rv := objc.Send[textColor](t_.ID(), selInit)
	return rv
}
