// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [MeasurementFormatter] class.
var MeasurementFormatterClass objc.Class

func init() {
	MeasurementFormatterClass = objc.GetClass("NSMeasurementFormatter")
}

type MeasurementFormatter struct {
	objc.ID
}

func MeasurementFormatterFrom(ptr unsafe.Pointer) MeasurementFormatter {
	return MeasurementFormatter{
		ID: objc.ID(ptr),
	}
}



