// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MeasurementFormatter] class.
var MeasurementFormatterClass = _MeasurementFormatterClass{objc.GetClass("NSMeasurementFormatter")}

type _MeasurementFormatterClass struct {
	class objc.Class
}

type MeasurementFormatter struct {
	objc.ID
}

func MeasurementFormatterFrom(ptr unsafe.Pointer) MeasurementFormatter {
	return MeasurementFormatter{
		ID: objc.ID(ptr),
	}
}




