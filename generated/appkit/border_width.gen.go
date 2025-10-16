
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [borderWidth] class.
var borderWidthClass _borderWidthClass

func init() {
	borderWidthClass = _borderWidthClass{objc.GetClass("borderWidth")}
}

type _borderWidthClass struct {
	objc.Class
}

// An interface definition for the [borderWidth] class.
type IborderWidth interface {
	ID() objc.ID
}

type borderWidth struct {
	id objc.ID
}

func borderWidthFrom(ptr unsafe.Pointer) borderWidth {
	return borderWidth{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (b_ borderWidth) ID() objc.ID {
	return b_.id
}

// Alloc allocates a new instance without initialization.
func (bc _borderWidthClass) Alloc() borderWidth {
	rv := objc.Send[borderWidth](objc.ID(bc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (bc _borderWidthClass) New() borderWidth {
	rv := objc.Send[borderWidth](objc.ID(bc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewborderWidth creates and returns a new initialized instance.
func NewborderWidth() borderWidth {
	return borderWidthClass.New()
}

// Init initializes the instance.
func (b_ borderWidth) Init() borderWidth {
	rv := objc.Send[borderWidth](b_.ID(), selInit)
	return rv
}
