
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [backgroundColors] class.
var backgroundColorsClass _backgroundColorsClass

func init() {
	backgroundColorsClass = _backgroundColorsClass{objc.GetClass("backgroundColors")}
}

type _backgroundColorsClass struct {
	objc.Class
}

// An interface definition for the [backgroundColors] class.
type IbackgroundColors interface {
	ID() objc.ID
}

type backgroundColors struct {
	id objc.ID
}

func backgroundColorsFrom(ptr unsafe.Pointer) backgroundColors {
	return backgroundColors{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (b_ backgroundColors) ID() objc.ID {
	return b_.id
}

// Alloc allocates a new instance without initialization.
func (bc _backgroundColorsClass) Alloc() backgroundColors {
	rv := objc.Send[backgroundColors](objc.ID(bc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (bc _backgroundColorsClass) New() backgroundColors {
	rv := objc.Send[backgroundColors](objc.ID(bc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewbackgroundColors creates and returns a new initialized instance.
func NewbackgroundColors() backgroundColors {
	return backgroundColorsClass.New()
}

// Init initializes the instance.
func (b_ backgroundColors) Init() backgroundColors {
	rv := objc.Send[backgroundColors](b_.ID(), selInit)
	return rv
}
