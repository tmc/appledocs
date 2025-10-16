
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [tintColor] class.
var tintColorClass _tintColorClass

func init() {
	tintColorClass = _tintColorClass{objc.GetClass("tintColor")}
}

type _tintColorClass struct {
	objc.Class
}

// An interface definition for the [tintColor] class.
type ItintColor interface {
	ID() objc.ID
}

type tintColor struct {
	id objc.ID
}

func tintColorFrom(ptr unsafe.Pointer) tintColor {
	return tintColor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ tintColor) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _tintColorClass) Alloc() tintColor {
	rv := objc.Send[tintColor](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _tintColorClass) New() tintColor {
	rv := objc.Send[tintColor](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewtintColor creates and returns a new initialized instance.
func NewtintColor() tintColor {
	return tintColorClass.New()
}

// Init initializes the instance.
func (t_ tintColor) Init() tintColor {
	rv := objc.Send[tintColor](t_.ID(), selInit)
	return rv
}
