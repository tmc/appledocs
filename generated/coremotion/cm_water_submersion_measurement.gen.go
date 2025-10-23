// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [WaterSubmersionMeasurement] class.
var (
	WaterSubmersionMeasurementClass     _WaterSubmersionMeasurementClass
	WaterSubmersionMeasurementClassOnce sync.Once
)

func getWaterSubmersionMeasurementClass() _WaterSubmersionMeasurementClass {
	WaterSubmersionMeasurementClassOnce.Do(func() {
		WaterSubmersionMeasurementClass = _WaterSubmersionMeasurementClass{objc.GetClass("CMWaterSubmersionMeasurement")}
	})
	return WaterSubmersionMeasurementClass
}

type _WaterSubmersionMeasurementClass struct {
	class objc.Class
}

// An interface definition for the [WaterSubmersionMeasurement] class.
type IWaterSubmersionMeasurement interface {
	objectivec.IObject
	// properties:
	Date() foundation.objc.IObject /* cross-framework: NSDate */
	Depth() unsafe.Pointer
	Pressure() unsafe.Pointer
	SubmersionState() WaterSubmersionDepthState
	SurfacePressure() unsafe.Pointer
	// methods:
}

// An update that contains data about the pressure and depth.


// An update that contains data about the pressure and depth.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionMeasurement
type WaterSubmersionMeasurement struct {
	objectivec.Object
}

// WaterSubmersionMeasurementFrom constructs a [WaterSubmersionMeasurement] from an unsafe.Pointer.
//
// An update that contains data about the pressure and depth.
func WaterSubmersionMeasurementFrom(ptr unsafe.Pointer) WaterSubmersionMeasurement {
	return WaterSubmersionMeasurement{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (wc _WaterSubmersionMeasurementClass) Alloc() WaterSubmersionMeasurement {
	rv := objc.Send[WaterSubmersionMeasurement](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (wc _WaterSubmersionMeasurementClass) New() WaterSubmersionMeasurement {
	rv := objc.Send[WaterSubmersionMeasurement](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WaterSubmersionMeasurement) Init() WaterSubmersionMeasurement {
	rv := objc.Send[WaterSubmersionMeasurement](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WaterSubmersionMeasurement) Autorelease() WaterSubmersionMeasurement {
	rv := objc.Send[WaterSubmersionMeasurement](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWaterSubmersionMeasurement creates a new WaterSubmersionMeasurement instance.
func NewWaterSubmersionMeasurement() WaterSubmersionMeasurement {
	return getWaterSubmersionMeasurementClass().New()
}



// The time and date when the system recorded the measurements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionMeasurement/date
func (w_ WaterSubmersionMeasurement) Date() foundation.objc.IObject /* cross-framework: NSDate */ {
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



