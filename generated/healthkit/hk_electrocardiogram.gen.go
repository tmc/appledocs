// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HKElectrocardiogram] class.
var (
	HKElectrocardiogramClass     _HKElectrocardiogramClass
	HKElectrocardiogramClassOnce sync.Once
)

func getHKElectrocardiogramClass() _HKElectrocardiogramClass {
	HKElectrocardiogramClassOnce.Do(func() {
		HKElectrocardiogramClass = _HKElectrocardiogramClass{objc.GetClass("HKElectrocardiogram")}
	})
	return HKElectrocardiogramClass
}

type _HKElectrocardiogramClass struct {
	class objc.Class
}

// An interface definition for the [HKElectrocardiogram] class.
type IHKElectrocardiogram interface {
	IHKSample
}

// A sample for electrocardiogram data.
//
// An is a collection of voltage values representing waveforms from one or more leads. The sample provides high-level details about the ECG reading, such as the sampling frequency or classification. HealthKit provides read-only access to electrocardiogram (ECG) data saved by Apple Watch. You can query for samples using an . After retrieving an sample, you can access the voltage measurements associated with the sample use an query.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKElectrocardiogram
type HKElectrocardiogram struct {
	HKSample
}

// HKElectrocardiogramFrom constructs a [HKElectrocardiogram] from an unsafe.Pointer.
//
// A sample for electrocardiogram data.
func HKElectrocardiogramFrom(ptr unsafe.Pointer) HKElectrocardiogram {
	return HKElectrocardiogram{
		HKSample: HKSampleFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HKElectrocardiogramClass) Alloc() HKElectrocardiogram {
	rv := objc.Send[HKElectrocardiogram](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKElectrocardiogramClass) New() HKElectrocardiogram {
	rv := objc.Send[HKElectrocardiogram](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKElectrocardiogram) Init() HKElectrocardiogram {
	rv := objc.Send[HKElectrocardiogram](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKElectrocardiogram) Autorelease() HKElectrocardiogram {
	rv := objc.Send[HKElectrocardiogram](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKElectrocardiogram creates a new HKElectrocardiogram instance.
func NewHKElectrocardiogram() HKElectrocardiogram {
	return getHKElectrocardiogramClass().New()
}


// The user’s average heart rate during the ECG.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKElectrocardiogram/averageHeartRate
func (h_ HKElectrocardiogram) AverageHeartRate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("averageHeartRate"))
	return rv
}

// The number of voltage measurements associated with this sample.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKElectrocardiogram/numberOfVoltageMeasurements
func (h_ HKElectrocardiogram) NumberOfVoltageMeasurements() int {
	rv := objc.Send[int](h_.ID, objc.Sel("numberOfVoltageMeasurements"))
	return rv
}



