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
func (h_ HKElectrocardiogram) AverageHeartRate() HKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("averageHeartRate"))
	return rv
}

// The number of voltage measurements associated with this sample.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKElectrocardiogram/numberOfVoltageMeasurements
func (h_ HKElectrocardiogram) NumberOfVoltageMeasurements() int {
	rv := objc.Send[int](h_.ID, objc.Sel("numberOfVoltageMeasurements"))
	return rv
}

// The ECG’s classification.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkelectrocardiogram/classification-swift.property
func (h_ HKElectrocardiogram) Classification() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("classification"))
	return rv
}


// SetClassification sets the value of the classification property.
// The ECG’s classification.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkelectrocardiogram/classification-swift.property
func (h_ HKElectrocardiogram) SetClassification(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setClassification:"), value)
}

// The frequency at which the Apple Watch sampled the voltage.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkelectrocardiogram/samplingfrequency
func (h_ HKElectrocardiogram) SamplingFrequency() HKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("samplingFrequency"))
	return rv
}


// SetSamplingFrequency sets the value of the samplingFrequency property.
// The frequency at which the Apple Watch sampled the voltage.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkelectrocardiogram/samplingfrequency
func (h_ HKElectrocardiogram) SetSamplingFrequency(value IHKQuantity) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setSamplingFrequency:"), value)
}

// A value that indicates whether the user entered a symptom when they recorded the ECG.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkelectrocardiogram/symptomsstatus-swift.property
func (h_ HKElectrocardiogram) SymptomsStatus() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("symptomsStatus"))
	return rv
}


// SetSymptomsStatus sets the value of the symptomsStatus property.
// A value that indicates whether the user entered a symptom when they recorded the ECG.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkelectrocardiogram/symptomsstatus-swift.property
func (h_ HKElectrocardiogram) SetSymptomsStatus(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setSymptomsStatus:"), value)
}

// A key for metadata indicating the version number of the algorithm Apple Watch uses to generate an ECG reading.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkmetadatakeyappleecgalgorithmversion
func (h_ HKElectrocardiogram) HKMetadataKeyAppleECGAlgorithmVersion() appkit.string {
	rv := objc.Send[appkit.string](h_.ID, objc.Sel("HKMetadataKeyAppleECGAlgorithmVersion"))
	return rv
}

// The key path for the sample’s average heart rate.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathaverageheartrate
func (h_ HKElectrocardiogram) HKPredicateKeyPathAverageHeartRate() appkit.string {
	rv := objc.Send[appkit.string](h_.ID, objc.Sel("HKPredicateKeyPathAverageHeartRate"))
	return rv
}

// The key path for the sample’s classification.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathecgclassification
func (h_ HKElectrocardiogram) HKPredicateKeyPathECGClassification() appkit.string {
	rv := objc.Send[appkit.string](h_.ID, objc.Sel("HKPredicateKeyPathECGClassification"))
	return rv
}

// The key path for the sample’s symptom status.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathecgsymptomsstatus
func (h_ HKElectrocardiogram) HKPredicateKeyPathECGSymptomsStatus() appkit.string {
	rv := objc.Send[appkit.string](h_.ID, objc.Sel("HKPredicateKeyPathECGSymptomsStatus"))
	return rv
}



