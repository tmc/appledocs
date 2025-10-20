// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var EnergyFormatterClass _EnergyFormatterClass

func init() {
	EnergyFormatterClass = _EnergyFormatterClass{objc.GetClass("NSEnergyFormatter")}
}

type _EnergyFormatterClass struct {
	class objc.Class
}

type EnergyFormatter struct {
	objc.ID
}

func EnergyFormatterFrom(ptr unsafe.Pointer) EnergyFormatter {
	return EnergyFormatter{
		ID: objc.ID(ptr),
	}
}




