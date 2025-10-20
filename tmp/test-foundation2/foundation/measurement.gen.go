// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var measurementClass _MeasurementClass

func init() {
	measurementClass = _MeasurementClass{objc.GetClass("NSMeasurement")}
}

type _MeasurementClass struct {
	class objc.Class
}

type Measurement struct {
	objc.ID
}

func MeasurementFrom(ptr unsafe.Pointer) Measurement {
	return Measurement{
		ID: objc.ID(ptr),
	}
}




