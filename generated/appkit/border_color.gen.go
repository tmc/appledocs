
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [borderColor] class.
var borderColorClass _borderColorClass

func init() {
	borderColorClass = _borderColorClass{objc.GetClass("borderColor")}
}

type _borderColorClass struct {
	objc.Class
}

// An interface definition for the [borderColor] class.
type IborderColor interface {
	ID() objc.ID
}

type borderColor struct {
	id objc.ID
}

func borderColorFrom(ptr unsafe.Pointer) borderColor {
	return borderColor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (b_ borderColor) ID() objc.ID {
	return b_.id
}

// Alloc allocates a new instance without initialization.
func (bc _borderColorClass) Alloc() borderColor {
	rv := objc.Send[borderColor](objc.ID(bc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (bc _borderColorClass) New() borderColor {
	rv := objc.Send[borderColor](objc.ID(bc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewborderColor creates and returns a new initialized instance.
func NewborderColor() borderColor {
	return borderColorClass.New()
}

// Init initializes the instance.
func (b_ borderColor) Init() borderColor {
	rv := objc.Send[borderColor](b_.ID(), selInit)
	return rv
}
