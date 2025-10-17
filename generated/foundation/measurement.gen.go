// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Measurement] class.
var MeasurementClass objc.Class

func init() {
	MeasurementClass = objc.GetClass("NSMeasurement")
}

type Measurement struct {
	objc.ID
}

func MeasurementFrom(ptr unsafe.Pointer) Measurement {
	return Measurement{
		ID: objc.ID(ptr),
	}
}



