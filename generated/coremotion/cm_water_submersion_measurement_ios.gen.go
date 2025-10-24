//go:build darwin && ios

// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for WaterSubmersionMeasurement


// iOS-only properties

// The time and date when the system recorded the measurements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionMeasurement/date
func (w_ WaterSubmersionMeasurement) Date() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](w_.ID, objc.Sel("date"))
	return rv
}

// The depth under water.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionMeasurement/depth
func (w_ WaterSubmersionMeasurement) Depth() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("depth"))
	return rv
}

// The water pressure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionMeasurement/pressure
func (w_ WaterSubmersionMeasurement) Pressure() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("pressure"))
	return rv
}

// The depth state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionMeasurement/submersionState
func (w_ WaterSubmersionMeasurement) SubmersionState() WaterSubmersionDepthState {
	rv := objc.Send[WaterSubmersionDepthState](w_.ID, objc.Sel("submersionState"))
	return rv
}

// The surface air pressure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionMeasurement/surfacePressure
func (w_ WaterSubmersionMeasurement) SurfacePressure() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("surfacePressure"))
	return rv
}





