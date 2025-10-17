// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [EnergyFormatter] class.
var energyFormatterClass = _EnergyFormatterClass{objc.GetClass("NSEnergyFormatter")}

type _EnergyFormatterClass struct {
	class objc.Class
}

// A formatter that provides localized descriptions of energy values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/EnergyFormatter

type EnergyFormatter struct {
	Formatter
}

// EnergyFormatterFrom constructs a [EnergyFormatter] from an unsafe.Pointer.
//
// A formatter that provides localized descriptions of energy values.
func EnergyFormatterFrom(ptr unsafe.Pointer) EnergyFormatter {
	return EnergyFormatter{
		Formatter: FormatterFrom(ptr),
	}
}



