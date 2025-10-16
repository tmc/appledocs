
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [backgroundTintColor] class.
var backgroundTintColorClass _backgroundTintColorClass

func init() {
	backgroundTintColorClass = _backgroundTintColorClass{objc.GetClass("backgroundTintColor")}
}

type _backgroundTintColorClass struct {
	objc.Class
}

// An interface definition for the [backgroundTintColor] class.
type IbackgroundTintColor interface {
	ID() objc.ID
}

type backgroundTintColor struct {
	id objc.ID
}

func backgroundTintColorFrom(ptr unsafe.Pointer) backgroundTintColor {
	return backgroundTintColor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (b_ backgroundTintColor) ID() objc.ID {
	return b_.id
}

// Alloc allocates a new instance without initialization.
func (bc _backgroundTintColorClass) Alloc() backgroundTintColor {
	rv := objc.Send[backgroundTintColor](objc.ID(bc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (bc _backgroundTintColorClass) New() backgroundTintColor {
	rv := objc.Send[backgroundTintColor](objc.ID(bc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewbackgroundTintColor creates and returns a new initialized instance.
func NewbackgroundTintColor() backgroundTintColor {
	return backgroundTintColorClass.New()
}

// Init initializes the instance.
func (b_ backgroundTintColor) Init() backgroundTintColor {
	rv := objc.Send[backgroundTintColor](b_.ID(), selInit)
	return rv
}
