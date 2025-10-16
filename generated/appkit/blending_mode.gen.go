
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [blendingMode] class.
var blendingModeClass _blendingModeClass

func init() {
	blendingModeClass = _blendingModeClass{objc.GetClass("blendingMode")}
}

type _blendingModeClass struct {
	objc.Class
}

// An interface definition for the [blendingMode] class.
type IblendingMode interface {
	ID() objc.ID
}

type blendingMode struct {
	id objc.ID
}

func blendingModeFrom(ptr unsafe.Pointer) blendingMode {
	return blendingMode{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (b_ blendingMode) ID() objc.ID {
	return b_.id
}

// Alloc allocates a new instance without initialization.
func (bc _blendingModeClass) Alloc() blendingMode {
	rv := objc.Send[blendingMode](objc.ID(bc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (bc _blendingModeClass) New() blendingMode {
	rv := objc.Send[blendingMode](objc.ID(bc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewblendingMode creates and returns a new initialized instance.
func NewblendingMode() blendingMode {
	return blendingModeClass.New()
}

// Init initializes the instance.
func (b_ blendingMode) Init() blendingMode {
	rv := objc.Send[blendingMode](b_.ID(), selInit)
	return rv
}
