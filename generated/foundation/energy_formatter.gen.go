// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [EnergyFormatter] class.
var EnergyFormatterClass objc.Class

func init() {
	EnergyFormatterClass = objc.GetClass("NSEnergyFormatter")
}

type EnergyFormatter struct {
	objc.ID
}

func EnergyFormatterFrom(ptr unsafe.Pointer) EnergyFormatter {
	return EnergyFormatter{
		ID: objc.ID(ptr),
	}
}



