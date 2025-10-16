
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [effectiveAppearance] class.
var effectiveAppearanceClass _effectiveAppearanceClass

func init() {
	effectiveAppearanceClass = _effectiveAppearanceClass{objc.GetClass("effectiveAppearance")}
}

type _effectiveAppearanceClass struct {
	objc.Class
}

// An interface definition for the [effectiveAppearance] class.
type IeffectiveAppearance interface {
	ID() objc.ID
}

type effectiveAppearance struct {
	id objc.ID
}

func effectiveAppearanceFrom(ptr unsafe.Pointer) effectiveAppearance {
	return effectiveAppearance{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (e_ effectiveAppearance) ID() objc.ID {
	return e_.id
}

// Alloc allocates a new instance without initialization.
func (ec _effectiveAppearanceClass) Alloc() effectiveAppearance {
	rv := objc.Send[effectiveAppearance](objc.ID(ec.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ec _effectiveAppearanceClass) New() effectiveAppearance {
	rv := objc.Send[effectiveAppearance](objc.ID(ec.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NeweffectiveAppearance creates and returns a new initialized instance.
func NeweffectiveAppearance() effectiveAppearance {
	return effectiveAppearanceClass.New()
}

// Init initializes the instance.
func (e_ effectiveAppearance) Init() effectiveAppearance {
	rv := objc.Send[effectiveAppearance](e_.ID(), selInit)
	return rv
}
