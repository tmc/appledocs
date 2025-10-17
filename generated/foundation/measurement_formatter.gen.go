// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MeasurementFormatter] class.
var measurementFormatterClass = _MeasurementFormatterClass{objc.GetClass("NSMeasurementFormatter")}

type _MeasurementFormatterClass struct {
	class objc.Class
}

// A formatter that provides localized representations of units and measurements. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/MeasurementFormatter

type MeasurementFormatter struct {
	Formatter
}

// MeasurementFormatterFrom constructs a [MeasurementFormatter] from an unsafe.Pointer.
//
// A formatter that provides localized representations of units and measurements.
func MeasurementFormatterFrom(ptr unsafe.Pointer) MeasurementFormatter {
	return MeasurementFormatter{
		Formatter: FormatterFrom(ptr),
	}
}



