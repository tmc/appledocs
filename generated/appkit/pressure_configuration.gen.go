
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PressureConfiguration] class.
var PressureConfigurationClass _PressureConfigurationClass

func init() {
	PressureConfigurationClass = _PressureConfigurationClass{objc.GetClass("NSPressureConfiguration")}
}

type _PressureConfigurationClass struct {
	objc.Class
}

// An interface definition for the [PressureConfiguration] class.
type IPressureConfiguration interface {
	ID() objc.ID
}

type PressureConfiguration struct {
	id objc.ID
}

func PressureConfigurationFrom(ptr unsafe.Pointer) PressureConfiguration {
	return PressureConfiguration{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ PressureConfiguration) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _PressureConfigurationClass) Alloc() PressureConfiguration {
	rv := objc.Send[PressureConfiguration](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _PressureConfigurationClass) New() PressureConfiguration {
	rv := objc.Send[PressureConfiguration](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewPressureConfiguration creates and returns a new initialized instance.
func NewPressureConfiguration() PressureConfiguration {
	return PressureConfigurationClass.New()
}

// Init initializes the instance.
func (p_ PressureConfiguration) Init() PressureConfiguration {
	rv := objc.Send[PressureConfiguration](p_.ID(), selInit)
	return rv
}
