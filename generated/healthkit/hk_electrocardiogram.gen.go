// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class HKElectrocardiogram */


/* debug [class_header]: Header for HKElectrocardiogram */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKElectrocardiogram */
// An interface definition for the [HKElectrocardiogram] class.
type IHKElectrocardiogram interface {
	IHKSample
	
/* debug [class_interface_properties]: Properties for HKElectrocardiogram */
	// properties:
	AverageHeartRate() IHKQuantity
	Classification() HKElectrocardiogramClassification
	NumberOfVoltageMeasurements() int
	SamplingFrequency() IHKQuantity
	SymptomsStatus() HKElectrocardiogramSymptomsStatus
	HKMetadataKeyAppleECGAlgorithmVersion() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathAverageHeartRate() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathECGClassification() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathECGSymptomsStatus() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKElectrocardiogram */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKElectrocardiogram */
// Alloc allocates a new instance without initialization.
func (hc _HKElectrocardiogramClass) Alloc() HKElectrocardiogram {
	rv := objc.Send[HKElectrocardiogram](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKElectrocardiogram */
// A sample for electrocardiogram data.
//
// An is a collection of voltage values representing waveforms from one or more leads. The sample provides high-level details about the ECG reading, such as the sampling frequency or classification. HealthKit provides read-only access to electrocardiogram (ECG) data saved by Apple Watch. You can query for samples using an . After retrieving an sample, you can access the voltage measurements associated with the sample use an query.


// A sample for electrocardiogram data.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKElectrocardiogram *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKElectrocardiogram */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKElectrocardiogram */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKElectrocardiogram */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKElectrocardiogram */

// The user’s average heart rate during the ECG.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKElectrocardiogram/averageHeartRate
func (h_ HKElectrocardiogram) AverageHeartRate() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("averageHeartRate"))
	return rv
}/* debug [instance_properties/getter]: averageHeartRate */


// The ECG’s classification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKElectrocardiogram/classification-swift.property
func (h_ HKElectrocardiogram) Classification() HKElectrocardiogramClassification {
	rv := objc.Send[HKElectrocardiogramClassification](h_.ID, objc.Sel("classification"))
	return rv
}/* debug [instance_properties/getter]: classification */


// The number of voltage measurements associated with this sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKElectrocardiogram/numberOfVoltageMeasurements
func (h_ HKElectrocardiogram) NumberOfVoltageMeasurements() int {
	rv := objc.Send[int](h_.ID, objc.Sel("numberOfVoltageMeasurements"))
	return rv
}/* debug [instance_properties/getter]: numberOfVoltageMeasurements */


// The frequency at which the Apple Watch sampled the voltage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKElectrocardiogram/samplingFrequency
func (h_ HKElectrocardiogram) SamplingFrequency() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("samplingFrequency"))
	return rv
}/* debug [instance_properties/getter]: samplingFrequency */


// A value that indicates whether the user entered a symptom when they recorded the ECG.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKElectrocardiogram/symptomsStatus-swift.property
func (h_ HKElectrocardiogram) SymptomsStatus() HKElectrocardiogramSymptomsStatus {
	rv := objc.Send[HKElectrocardiogramSymptomsStatus](h_.ID, objc.Sel("symptomsStatus"))
	return rv
}/* debug [instance_properties/getter]: symptomsStatus */


// A key for metadata indicating the version number of the algorithm Apple Watch uses to generate an ECG reading.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkmetadatakeyappleecgalgorithmversion
func (h_ HKElectrocardiogram) HKMetadataKeyAppleECGAlgorithmVersion() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKMetadataKeyAppleECGAlgorithmVersion"))
	return rv
}/* debug [instance_properties/getter]: HKMetadataKeyAppleECGAlgorithmVersion */


// The key path for the sample’s average heart rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathaverageheartrate
func (h_ HKElectrocardiogram) HKPredicateKeyPathAverageHeartRate() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathAverageHeartRate"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathAverageHeartRate */


// The key path for the sample’s classification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathecgclassification
func (h_ HKElectrocardiogram) HKPredicateKeyPathECGClassification() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathECGClassification"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathECGClassification */


// The key path for the sample’s symptom status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathecgsymptomsstatus
func (h_ HKElectrocardiogram) HKPredicateKeyPathECGSymptomsStatus() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathECGSymptomsStatus"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathECGSymptomsStatus */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKElectrocardiogram */



