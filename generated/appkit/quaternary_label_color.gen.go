
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [quaternaryLabelColor] class.
var quaternaryLabelColorClass _quaternaryLabelColorClass

func init() {
	quaternaryLabelColorClass = _quaternaryLabelColorClass{objc.GetClass("quaternaryLabelColor")}
}

type _quaternaryLabelColorClass struct {
	objc.Class
}

// An interface definition for the [quaternaryLabelColor] class.
type IquaternaryLabelColor interface {
	ID() objc.ID
}

type quaternaryLabelColor struct {
	id objc.ID
}

func quaternaryLabelColorFrom(ptr unsafe.Pointer) quaternaryLabelColor {
	return quaternaryLabelColor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (q_ quaternaryLabelColor) ID() objc.ID {
	return q_.id
}

// Alloc allocates a new instance without initialization.
func (qc _quaternaryLabelColorClass) Alloc() quaternaryLabelColor {
	rv := objc.Send[quaternaryLabelColor](objc.ID(qc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (qc _quaternaryLabelColorClass) New() quaternaryLabelColor {
	rv := objc.Send[quaternaryLabelColor](objc.ID(qc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewquaternaryLabelColor creates and returns a new initialized instance.
func NewquaternaryLabelColor() quaternaryLabelColor {
	return quaternaryLabelColorClass.New()
}

// Init initializes the instance.
func (q_ quaternaryLabelColor) Init() quaternaryLabelColor {
	rv := objc.Send[quaternaryLabelColor](q_.ID(), selInit)
	return rv
}
