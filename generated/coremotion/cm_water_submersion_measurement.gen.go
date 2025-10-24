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



