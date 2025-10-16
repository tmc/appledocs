
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [bezelColor] class.
var bezelColorClass _bezelColorClass

func init() {
	bezelColorClass = _bezelColorClass{objc.GetClass("bezelColor")}
}

type _bezelColorClass struct {
	objc.Class
}

// An interface definition for the [bezelColor] class.
type IbezelColor interface {
	ID() objc.ID
}

type bezelColor struct {
	id objc.ID
}

func bezelColorFrom(ptr unsafe.Pointer) bezelColor {
	return bezelColor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (b_ bezelColor) ID() objc.ID {
	return b_.id
}

// Alloc allocates a new instance without initialization.
func (bc _bezelColorClass) Alloc() bezelColor {
	rv := objc.Send[bezelColor](objc.ID(bc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (bc _bezelColorClass) New() bezelColor {
	rv := objc.Send[bezelColor](objc.ID(bc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewbezelColor creates and returns a new initialized instance.
func NewbezelColor() bezelColor {
	return bezelColorClass.New()
}

// Init initializes the instance.
func (b_ bezelColor) Init() bezelColor {
	rv := objc.Send[bezelColor](b_.ID(), selInit)
	return rv
}
