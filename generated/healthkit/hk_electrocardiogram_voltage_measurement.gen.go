// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HKElectrocardiogramVoltageMeasurement] class.
var (
	HKElectrocardiogramVoltageMeasurementClass     _HKElectrocardiogramVoltageMeasurementClass
	HKElectrocardiogramVoltageMeasurementClassOnce sync.Once
)

func getHKElectrocardiogramVoltageMeasurementClass() _HKElectrocardiogramVoltageMeasurementClass {
	HKElectrocardiogramVoltageMeasurementClassOnce.Do(func() {
		HKElectrocardiogramVoltageMeasurementClass = _HKElectrocardiogramVoltageMeasurementClass{objc.GetClass("HKElectrocardiogramVoltageMeasurement")}
	})
	return HKElectrocardiogramVoltageMeasurementClass
}

type _HKElectrocardiogramVoltageMeasurementClass struct {
	class objc.Class
}

// An interface definition for the [HKElectrocardiogramVoltageMeasurement] class.
type IHKElectrocardiogramVoltageMeasurement interface {
	objectivec.IObject
	TimeSinceSampleStart() unsafe.Pointer
	SetTimeSinceSampleStart(value unsafe.Pointer)
}

// The voltage for all leads at a single point in time.


// The voltage for all leads at a single point in time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKElectrocardiogram/VoltageMeasurement

type HKElectrocardiogramVoltageMeasurement struct {
	objectivec.Object
}

// HKElectrocardiogramVoltageMeasurementFrom constructs a [HKElectrocardiogramVoltageMeasurement] from an unsafe.Pointer.
//
// The voltage for all leads at a single point in time.
func HKElectrocardiogramVoltageMeasurementFrom(ptr unsafe.Pointer) HKElectrocardiogramVoltageMeasurement {
	return HKElectrocardiogramVoltageMeasurement{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKElectrocardiogramVoltageMeasurementClass) Alloc() HKElectrocardiogramVoltageMeasurement {
	rv := objc.Send[HKElectrocardiogramVoltageMeasurement](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKElectrocardiogramVoltageMeasurementClass) New() HKElectrocardiogramVoltageMeasurement {
	rv := objc.Send[HKElectrocardiogramVoltageMeasurement](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKElectrocardiogramVoltageMeasurement) Init() HKElectrocardiogramVoltageMeasurement {
	rv := objc.Send[HKElectrocardiogramVoltageMeasurement](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKElectrocardiogramVoltageMeasurement) Autorelease() HKElectrocardiogramVoltageMeasurement {
	rv := objc.Send[HKElectrocardiogramVoltageMeasurement](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKElectrocardiogramVoltageMeasurement creates a new HKElectrocardiogramVoltageMeasurement instance.
func NewHKElectrocardiogramVoltageMeasurement() HKElectrocardiogramVoltageMeasurement {
	return getHKElectrocardiogramVoltageMeasurementClass().New()
}



// The time of the measurement relative to the sample’s start time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkelectrocardiogram/voltagemeasurement/timesincesamplestart

func (h_ HKElectrocardiogramVoltageMeasurement) TimeSinceSampleStart() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("timeSinceSampleStart"))
	return rv
}


// The time of the measurement relative to the sample’s start time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkelectrocardiogram/voltagemeasurement/timesincesamplestart

func (h_ HKElectrocardiogramVoltageMeasurement) SetTimeSinceSampleStart(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setTimeSinceSampleStart:"), value)
}



