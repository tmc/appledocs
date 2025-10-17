// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PressureConfiguration] class.
var pressureConfigurationClass = _PressureConfigurationClass{objc.GetClass("NSPressureConfiguration")}

type _PressureConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [PressureConfiguration] class.
type IPressureConfiguration interface {
	objectivec.IObject
}

// An encapsulation of the behavior and progression of a Force Touch trackpad as it responds to specific events. [Full Topic]
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



