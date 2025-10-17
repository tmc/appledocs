// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PressureConfiguration] class.
var PressureConfigurationClass objc.Class

func init() {
	PressureConfigurationClass = objc.GetClass("NSPressureConfiguration")
}

type PressureConfiguration struct {
	objc.ID
}

func PressureConfigurationFrom(ptr unsafe.Pointer) PressureConfiguration {
	return PressureConfiguration{
		ID: objc.ID(ptr),
	}
}




