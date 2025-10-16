
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [tertiaryLabelColor] class.
var tertiaryLabelColorClass _tertiaryLabelColorClass

func init() {
	tertiaryLabelColorClass = _tertiaryLabelColorClass{objc.GetClass("tertiaryLabelColor")}
}

type _tertiaryLabelColorClass struct {
	objc.Class
}

// An interface definition for the [tertiaryLabelColor] class.
type ItertiaryLabelColor interface {
	ID() objc.ID
}

type tertiaryLabelColor struct {
	id objc.ID
}

func tertiaryLabelColorFrom(ptr unsafe.Pointer) tertiaryLabelColor {
	return tertiaryLabelColor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ tertiaryLabelColor) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _tertiaryLabelColorClass) Alloc() tertiaryLabelColor {
	rv := objc.Send[tertiaryLabelColor](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _tertiaryLabelColorClass) New() tertiaryLabelColor {
	rv := objc.Send[tertiaryLabelColor](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewtertiaryLabelColor creates and returns a new initialized instance.
func NewtertiaryLabelColor() tertiaryLabelColor {
	return tertiaryLabelColorClass.New()
}

// Init initializes the instance.
func (t_ tertiaryLabelColor) Init() tertiaryLabelColor {
	rv := objc.Send[tertiaryLabelColor](t_.ID(), selInit)
	return rv
}
