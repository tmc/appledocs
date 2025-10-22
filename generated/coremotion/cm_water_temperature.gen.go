// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [WaterTemperature] class.
var (
	WaterTemperatureClass     _WaterTemperatureClass
	WaterTemperatureClassOnce sync.Once
)

func getWaterTemperatureClass() _WaterTemperatureClass {
	WaterTemperatureClassOnce.Do(func() {
		WaterTemperatureClass = _WaterTemperatureClass{objc.GetClass("CMWaterTemperature")}
	})
	return WaterTemperatureClass
}

type _WaterTemperatureClass struct {
	class objc.Class
}

// An interface definition for the [WaterTemperature] class.
type IWaterTemperature interface {
	objectivec.IObject
	Date() foundation.NSDate
	Temperature() unsafe.Pointer
	TemperatureUncertainty() unsafe.Pointer
}

// An update that contains data about the water temperature.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterTemperature
type WaterTemperature struct {
	objectivec.Object
}

// WaterTemperatureFrom constructs a [WaterTemperature] from an unsafe.Pointer.
//
// An update that contains data about the water temperature.
func WaterTemperatureFrom(ptr unsafe.Pointer) WaterTemperature {
	return WaterTemperature{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (wc _WaterTemperatureClass) Alloc() WaterTemperature {
	rv := objc.Send[WaterTemperature](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (wc _WaterTemperatureClass) New() WaterTemperature {
	rv := objc.Send[WaterTemperature](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WaterTemperature) Init() WaterTemperature {
	rv := objc.Send[WaterTemperature](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WaterTemperature) Autorelease() WaterTemperature {
	rv := objc.Send[WaterTemperature](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWaterTemperature creates a new WaterTemperature instance.
func NewWaterTemperature() WaterTemperature {
	return getWaterTemperatureClass().New()
}


// The time and date when the system recorded the measurements.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterTemperature/date
func (w_ WaterTemperature) Date() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](w_.ID, objc.Sel("date"))
	return rv
}

// The water temperature.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterTemperature/temperature
func (w_ WaterTemperature) Temperature() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("temperature"))
	return rv
}

// The amount of uncertainty in the measurement of the water temperature.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterTemperature/temperatureUncertainty
func (w_ WaterTemperature) TemperatureUncertainty() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("temperatureUncertainty"))
	return rv
}




