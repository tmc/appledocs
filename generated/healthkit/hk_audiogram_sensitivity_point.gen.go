// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKAudiogramSensitivityPoint */


/* debug [class_header]: Header for HKAudiogramSensitivityPoint */
// The class instance for the [HKAudiogramSensitivityPoint] class.
var (
	HKAudiogramSensitivityPointClass     _HKAudiogramSensitivityPointClass
	HKAudiogramSensitivityPointClassOnce sync.Once
)

func getHKAudiogramSensitivityPointClass() _HKAudiogramSensitivityPointClass {
	HKAudiogramSensitivityPointClassOnce.Do(func() {
		HKAudiogramSensitivityPointClass = _HKAudiogramSensitivityPointClass{objc.GetClass("HKAudiogramSensitivityPoint")}
	})
	return HKAudiogramSensitivityPointClass
}

type _HKAudiogramSensitivityPointClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKAudiogramSensitivityPoint */
// An interface definition for the [HKAudiogramSensitivityPoint] class.
type IHKAudiogramSensitivityPoint interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKAudiogramSensitivityPoint */
	// properties:
	Frequency() IHKQuantity
	LeftEarSensitivity() IHKQuantity
	RightEarSensitivity() IHKQuantity
	Tests() []HKAudiogramSensitivityTest
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKAudiogramSensitivityPoint */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKAudiogramSensitivityPoint */
// Alloc allocates a new instance without initialization.
func (hc _HKAudiogramSensitivityPointClass) Alloc() HKAudiogramSensitivityPoint {
	rv := objc.Send[HKAudiogramSensitivityPoint](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKAudiogramSensitivityPointClass) New() HKAudiogramSensitivityPoint {
	rv := objc.Send[HKAudiogramSensitivityPoint](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKAudiogramSensitivityPoint) Init() HKAudiogramSensitivityPoint {
	rv := objc.Send[HKAudiogramSensitivityPoint](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKAudiogramSensitivityPoint) Autorelease() HKAudiogramSensitivityPoint {
	rv := objc.Send[HKAudiogramSensitivityPoint](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKAudiogramSensitivityPoint creates a new HKAudiogramSensitivityPoint instance.
func NewHKAudiogramSensitivityPoint() HKAudiogramSensitivityPoint {
	return getHKAudiogramSensitivityPointClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKAudiogramSensitivityPoint */
// A hearing sensitivity reading associated with a hearing test.


// A hearing sensitivity reading associated with a hearing test.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramSensitivityPoint
type HKAudiogramSensitivityPoint struct {
	objectivec.Object
}

// HKAudiogramSensitivityPointFrom constructs a [HKAudiogramSensitivityPoint] from an unsafe.Pointer.
//
// A hearing sensitivity reading associated with a hearing test.
func HKAudiogramSensitivityPointFrom(ptr unsafe.Pointer) HKAudiogramSensitivityPoint {
	return HKAudiogramSensitivityPoint{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKAudiogramSensitivityPoint */

// Creates a new sensitivity point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramSensitivityPoint/init(frequency:leftEarSensitivity:rightEarSensitivity:)
func NewHKAudiogramSensitivityPointWithFrequencyLeftEarSensitivityRightEarSensitivityError(frequency IHKQuantity, leftEarSensitivity IHKQuantity, rightEarSensitivity IHKQuantity, error_ objectivec.IObject) HKAudiogramSensitivityPoint {
	rv := objc.Send[HKAudiogramSensitivityPoint](objc.ID(getHKAudiogramSensitivityPointClass().class), objc.Sel("sensitivityPointWithFrequency:leftEarSensitivity:rightEarSensitivity:error:"), frequency, leftEarSensitivity, rightEarSensitivity, error_)
	return rv
}/* debug [class_init_methods/constructor]: NewHKAudiogramSensitivityPointWithFrequencyLeftEarSensitivityRightEarSensitivityError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramSensitivityPoint/init(frequency:tests:)
func NewHKAudiogramSensitivityPointWithFrequencyTestsError(frequency IHKQuantity, tests []HKAudiogramSensitivityTest, errorOut objectivec.IObject) HKAudiogramSensitivityPoint {
	rv := objc.Send[HKAudiogramSensitivityPoint](objc.ID(getHKAudiogramSensitivityPointClass().class), objc.Sel("sensitivityPointWithFrequency:tests:error:"), frequency, tests, errorOut)
	return rv
}/* debug [class_init_methods/constructor]: NewHKAudiogramSensitivityPointWithFrequencyTestsError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKAudiogramSensitivityPoint */

// Creates a new sensitivity point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramSensitivityPoint/init(frequency:leftEarSensitivity:rightEarSensitivity:)
func (hc _HKAudiogramSensitivityPointClass) SensitivityPointWithFrequencyLeftEarSensitivityRightEarSensitivityError(frequency IHKQuantity, leftEarSensitivity IHKQuantity, rightEarSensitivity IHKQuantity, error_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("sensitivityPointWithFrequency:leftEarSensitivity:rightEarSensitivity:error:"), frequency, leftEarSensitivity, rightEarSensitivity, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SensitivityPointWithFrequencyLeftEarSensitivityRightEarSensitivityError) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramSensitivityPoint/init(frequency:tests:)
func (hc _HKAudiogramSensitivityPointClass) SensitivityPointWithFrequencyTestsError(frequency IHKQuantity, tests []HKAudiogramSensitivityTest, errorOut objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("sensitivityPointWithFrequency:tests:error:"), frequency, tests, errorOut)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SensitivityPointWithFrequencyTestsError) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKAudiogramSensitivityPoint */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKAudiogramSensitivityPoint */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKAudiogramSensitivityPoint */

// The frequency tested in the hearing test.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramSensitivityPoint/frequency
func (h_ HKAudiogramSensitivityPoint) Frequency() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("frequency"))
	return rv
}/* debug [instance_properties/getter]: frequency */


// The sensitivity of the left ear.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramSensitivityPoint/leftEarSensitivity
func (h_ HKAudiogramSensitivityPoint) LeftEarSensitivity() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("leftEarSensitivity"))
	return rv
}/* debug [instance_properties/getter]: leftEarSensitivity */


// The sensitivity of the right ear.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramSensitivityPoint/rightEarSensitivity
func (h_ HKAudiogramSensitivityPoint) RightEarSensitivity() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("rightEarSensitivity"))
	return rv
}/* debug [instance_properties/getter]: rightEarSensitivity */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramSensitivityPoint/tests
func (h_ HKAudiogramSensitivityPoint) Tests() []HKAudiogramSensitivityTest {
	rv := objc.Send[[]HKAudiogramSensitivityTest](h_.ID, objc.Sel("tests"))
	return rv
}/* debug [instance_properties/getter]: tests */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKAudiogramSensitivityPoint */


