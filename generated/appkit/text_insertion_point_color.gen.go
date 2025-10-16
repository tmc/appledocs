
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [textInsertionPointColor] class.
var textInsertionPointColorClass _textInsertionPointColorClass

func init() {
	textInsertionPointColorClass = _textInsertionPointColorClass{objc.GetClass("textInsertionPointColor")}
}

type _textInsertionPointColorClass struct {
	objc.Class
}

// An interface definition for the [textInsertionPointColor] class.
type ItextInsertionPointColor interface {
	ID() objc.ID
}

type textInsertionPointColor struct {
	id objc.ID
}

func textInsertionPointColorFrom(ptr unsafe.Pointer) textInsertionPointColor {
	return textInsertionPointColor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ textInsertionPointColor) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _textInsertionPointColorClass) Alloc() textInsertionPointColor {
	rv := objc.Send[textInsertionPointColor](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _textInsertionPointColorClass) New() textInsertionPointColor {
	rv := objc.Send[textInsertionPointColor](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewtextInsertionPointColor creates and returns a new initialized instance.
func NewtextInsertionPointColor() textInsertionPointColor {
	return textInsertionPointColorClass.New()
}

// Init initializes the instance.
func (t_ textInsertionPointColor) Init() textInsertionPointColor {
	rv := objc.Send[textInsertionPointColor](t_.ID(), selInit)
	return rv
}
