//go:build darwin && ios

// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for WaterTemperature


// iOS-only properties

// The time and date when the system recorded the measurements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterTemperature/date
func (w_ WaterTemperature) Date() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](w_.ID, objc.Sel("date"))
	return rv
}

// The water temperature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterTemperature/temperature
func (w_ WaterTemperature) Temperature() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("temperature"))
	return rv
}

// The amount of uncertainty in the measurement of the water temperature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterTemperature/temperatureUncertainty
func (w_ WaterTemperature) TemperatureUncertainty() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("temperatureUncertainty"))
	return rv
}





