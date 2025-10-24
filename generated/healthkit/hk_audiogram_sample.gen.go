// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKAudiogramSample */


/* debug [class_header]: Header for HKAudiogramSample */
// The class instance for the [HKAudiogramSample] class.
var (
	HKAudiogramSampleClass     _HKAudiogramSampleClass
	HKAudiogramSampleClassOnce sync.Once
)

func getHKAudiogramSampleClass() _HKAudiogramSampleClass {
	HKAudiogramSampleClassOnce.Do(func() {
		HKAudiogramSampleClass = _HKAudiogramSampleClass{objc.GetClass("HKAudiogramSample")}
	})
	return HKAudiogramSampleClass
}

type _HKAudiogramSampleClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKAudiogramSample */
// An interface definition for the [HKAudiogramSample] class.
type IHKAudiogramSample interface {
	IHKSample
	
/* debug [class_interface_properties]: Properties for HKAudiogramSample */
	// properties:
	SensitivityPoints() []HKAudiogramSensitivityPoint
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKAudiogramSample */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKAudiogramSample */
// Alloc allocates a new instance without initialization.
func (hc _HKAudiogramSampleClass) Alloc() HKAudiogramSample {
	rv := objc.Send[HKAudiogramSample](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKAudiogramSampleClass) New() HKAudiogramSample {
	rv := objc.Send[HKAudiogramSample](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKAudiogramSample) Init() HKAudiogramSample {
	rv := objc.Send[HKAudiogramSample](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKAudiogramSample) Autorelease() HKAudiogramSample {
	rv := objc.Send[HKAudiogramSample](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKAudiogramSample creates a new HKAudiogramSample instance.
func NewHKAudiogramSample() HKAudiogramSample {
	return getHKAudiogramSampleClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKAudiogramSample */
// A sample that stores an audiogram.
//
// This sample stores the results from a hearing test. The sample stores the audiogram data as an array of sensitivity points.


// A sample that stores an audiogram.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramSample
type HKAudiogramSample struct {
	HKSample
}

// HKAudiogramSampleFrom constructs a [HKAudiogramSample] from an unsafe.Pointer.
//
// A sample that stores an audiogram.
func HKAudiogramSampleFrom(ptr unsafe.Pointer) HKAudiogramSample {
	return HKAudiogramSample{
		HKSample: HKSampleFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKAudiogramSample */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramSample/init(sensitivityPoints:start:end:device:metadata:)
func NewHKAudiogramSampleWithSensitivityPointsStartDateEndDateDeviceMetadata(sensitivityPoints []HKAudiogramSensitivityPoint, startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, device IHKDevice, metadata foundation.IDictionary) HKAudiogramSample {
	rv := objc.Send[HKAudiogramSample](objc.ID(getHKAudiogramSampleClass().class), objc.Sel("audiogramSampleWithSensitivityPoints:startDate:endDate:device:metadata:"), sensitivityPoints, startDate, endDate, device, metadata)
	return rv
}/* debug [class_init_methods/constructor]: NewHKAudiogramSampleWithSensitivityPointsStartDateEndDateDeviceMetadata */


// Creates a new audiogram sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramSample/init(sensitivityPoints:start:end:metadata:)
func NewHKAudiogramSampleWithSensitivityPointsStartDateEndDateMetadata(sensitivityPoints []HKAudiogramSensitivityPoint, startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, metadata foundation.IDictionary) HKAudiogramSample {
	rv := objc.Send[HKAudiogramSample](objc.ID(getHKAudiogramSampleClass().class), objc.Sel("audiogramSampleWithSensitivityPoints:startDate:endDate:metadata:"), sensitivityPoints, startDate, endDate, metadata)
	return rv
}/* debug [class_init_methods/constructor]: NewHKAudiogramSampleWithSensitivityPointsStartDateEndDateMetadata */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKAudiogramSample */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramSample/init(sensitivityPoints:start:end:device:metadata:)
func (hc _HKAudiogramSampleClass) AudiogramSampleWithSensitivityPointsStartDateEndDateDeviceMetadata(sensitivityPoints []HKAudiogramSensitivityPoint, startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, device IHKDevice, metadata foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("audiogramSampleWithSensitivityPoints:startDate:endDate:device:metadata:"), sensitivityPoints, startDate, endDate, device, metadata)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AudiogramSampleWithSensitivityPointsStartDateEndDateDeviceMetadata) */


// Creates a new audiogram sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramSample/init(sensitivityPoints:start:end:metadata:)
func (hc _HKAudiogramSampleClass) AudiogramSampleWithSensitivityPointsStartDateEndDateMetadata(sensitivityPoints []HKAudiogramSensitivityPoint, startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, metadata foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("audiogramSampleWithSensitivityPoints:startDate:endDate:metadata:"), sensitivityPoints, startDate, endDate, metadata)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AudiogramSampleWithSensitivityPointsStartDateEndDateMetadata) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKAudiogramSample */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKAudiogramSample */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKAudiogramSample */

// An array of sensitivity point objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramSample/sensitivityPoints
func (h_ HKAudiogramSample) SensitivityPoints() []HKAudiogramSensitivityPoint {
	rv := objc.Send[[]HKAudiogramSensitivityPoint](h_.ID, objc.Sel("sensitivityPoints"))
	return rv
}/* debug [instance_properties/getter]: sensitivityPoints */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKAudiogramSample */


