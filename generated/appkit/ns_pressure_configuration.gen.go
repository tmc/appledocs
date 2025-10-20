// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PressureConfiguration] class.
var (
	PressureConfigurationClass     _PressureConfigurationClass
	PressureConfigurationClassOnce sync.Once
)

func getPressureConfigurationClass() _PressureConfigurationClass {
	PressureConfigurationClassOnce.Do(func() {
		PressureConfigurationClass = _PressureConfigurationClass{objc.GetClass("NSPressureConfiguration")}
	})
	return PressureConfigurationClass
}

type _PressureConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [PressureConfiguration] class.
type IPressureConfiguration interface {
	objectivec.IObject
	Set()
}

// An encapsulation of the behavior and progression of a Force Touch trackpad as it responds to specific events.
//
// Use an object to configure the behavior and progression of a Force Touch trackpad when it responds to a mouse drag or pressure event sequence. Pressure configurations are assigned to views ( ) and gesture recognizers ( ).
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPressureConfiguration
type PressureConfiguration struct {
	objectivec.Object
}

// PressureConfigurationFrom constructs a [PressureConfiguration] from an unsafe.Pointer.
//
// An encapsulation of the behavior and progression of a Force Touch trackpad as it responds to specific events.
func PressureConfigurationFrom(ptr unsafe.Pointer) PressureConfiguration {
	return PressureConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PressureConfigurationClass) Alloc() PressureConfiguration {
	rv := objc.Send[PressureConfiguration](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PressureConfigurationClass) New() PressureConfiguration {
	rv := objc.Send[PressureConfiguration](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PressureConfiguration) Init() PressureConfiguration {
	rv := objc.Send[PressureConfiguration](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PressureConfiguration) Autorelease() PressureConfiguration {
	rv := objc.Send[PressureConfiguration](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPressureConfiguration creates a new PressureConfiguration instance.
func NewPressureConfiguration() PressureConfiguration {
	return getPressureConfigurationClass().New()
}


// Changes the pressure configuration of the trackpad to the initialized pressure configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPressureConfiguration/set()
func (p_ PressureConfiguration) Set() {
	objc.Send[objc.ID](p_.ID, objc.Sel("set"))
}



