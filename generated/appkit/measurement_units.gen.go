
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [measurementUnits] class.
var measurementUnitsClass _measurementUnitsClass

func init() {
	measurementUnitsClass = _measurementUnitsClass{objc.GetClass("measurementUnits")}
}

type _measurementUnitsClass struct {
	objc.Class
}

// An interface definition for the [measurementUnits] class.
type ImeasurementUnits interface {
	ID() objc.ID
}

type measurementUnits struct {
	id objc.ID
}

func measurementUnitsFrom(ptr unsafe.Pointer) measurementUnits {
	return measurementUnits{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ measurementUnits) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _measurementUnitsClass) Alloc() measurementUnits {
	rv := objc.Send[measurementUnits](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _measurementUnitsClass) New() measurementUnits {
	rv := objc.Send[measurementUnits](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewmeasurementUnits creates and returns a new initialized instance.
func NewmeasurementUnits() measurementUnits {
	return measurementUnitsClass.New()
}

// Init initializes the instance.
func (m_ measurementUnits) Init() measurementUnits {
	rv := objc.Send[measurementUnits](m_.ID(), selInit)
	return rv
}
