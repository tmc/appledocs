// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKAudiogramSensitivityTest */


/* debug [class_header]: Header for HKAudiogramSensitivityTest */
// The class instance for the [HKAudiogramSensitivityTest] class.
var (
	HKAudiogramSensitivityTestClass     _HKAudiogramSensitivityTestClass
	HKAudiogramSensitivityTestClassOnce sync.Once
)

func getHKAudiogramSensitivityTestClass() _HKAudiogramSensitivityTestClass {
	HKAudiogramSensitivityTestClassOnce.Do(func() {
		HKAudiogramSensitivityTestClass = _HKAudiogramSensitivityTestClass{objc.GetClass("HKAudiogramSensitivityTest")}
	})
	return HKAudiogramSensitivityTestClass
}

type _HKAudiogramSensitivityTestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKAudiogramSensitivityTest */
// An interface definition for the [HKAudiogramSensitivityTest] class.
type IHKAudiogramSensitivityTest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKAudiogramSensitivityTest */
	// properties:
	ClampingRange() IHKAudiogramSensitivityPointClampingRange
	Masked() bool
	Sensitivity() IHKQuantity
	Side() HKAudiogramSensitivityTestSide
	Type() HKAudiogramConductionType
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKAudiogramSensitivityTest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKAudiogramSensitivityTest */
// Alloc allocates a new instance without initialization.
func (hc _HKAudiogramSensitivityTestClass) Alloc() HKAudiogramSensitivityTest {
	rv := objc.Send[HKAudiogramSensitivityTest](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKAudiogramSensitivityTestClass) New() HKAudiogramSensitivityTest {
	rv := objc.Send[HKAudiogramSensitivityTest](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKAudiogramSensitivityTest) Init() HKAudiogramSensitivityTest {
	rv := objc.Send[HKAudiogramSensitivityTest](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKAudiogramSensitivityTest) Autorelease() HKAudiogramSensitivityTest {
	rv := objc.Send[HKAudiogramSensitivityTest](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKAudiogramSensitivityTest creates a new HKAudiogramSensitivityTest instance.
func NewHKAudiogramSensitivityTest() HKAudiogramSensitivityTest {
	return getHKAudiogramSensitivityTestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKAudiogramSensitivityTest */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramSensitivityTest
type HKAudiogramSensitivityTest struct {
	objectivec.Object
}

// HKAudiogramSensitivityTestFrom constructs a [HKAudiogramSensitivityTest] from an unsafe.Pointer.
func HKAudiogramSensitivityTestFrom(ptr unsafe.Pointer) HKAudiogramSensitivityTest {
	return HKAudiogramSensitivityTest{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKAudiogramSensitivityTest */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramSensitivityTest/init(sensitivity:type:masked:side:clampingRange:)
func NewHKAudiogramSensitivityTestWithSensitivityTypeMaskedSideClampingRangeError(sensitivity IHKQuantity, type_ HKAudiogramConductionType, masked bool, side HKAudiogramSensitivityTestSide, clampingRange IHKAudiogramSensitivityPointClampingRange, errorOut objectivec.IObject) HKAudiogramSensitivityTest {
	instance := getHKAudiogramSensitivityTestClass().Alloc()
	rv := objc.Send[HKAudiogramSensitivityTest](instance.ID, objc.Sel("initWithSensitivity:type:masked:side:clampingRange:error:"), sensitivity, type_, masked, side, clampingRange, errorOut)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKAudiogramSensitivityTestWithSensitivityTypeMaskedSideClampingRangeError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKAudiogramSensitivityTest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKAudiogramSensitivityTest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKAudiogramSensitivityTest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKAudiogramSensitivityTest */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramSensitivityTest/clampingRange
func (h_ HKAudiogramSensitivityTest) ClampingRange() IHKAudiogramSensitivityPointClampingRange {
	rv := objc.Send[HKAudiogramSensitivityPointClampingRange](h_.ID, objc.Sel("clampingRange"))
	return rv
}/* debug [instance_properties/getter]: clampingRange */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramSensitivityTest/masked
func (h_ HKAudiogramSensitivityTest) Masked() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("masked"))
	return rv
}/* debug [instance_properties/getter]: masked */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramSensitivityTest/sensitivity
func (h_ HKAudiogramSensitivityTest) Sensitivity() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("sensitivity"))
	return rv
}/* debug [instance_properties/getter]: sensitivity */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramSensitivityTest/side
func (h_ HKAudiogramSensitivityTest) Side() HKAudiogramSensitivityTestSide {
	rv := objc.Send[HKAudiogramSensitivityTestSide](h_.ID, objc.Sel("side"))
	return rv
}/* debug [instance_properties/getter]: side */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramSensitivityTest/type
func (h_ HKAudiogramSensitivityTest) Type() HKAudiogramConductionType {
	rv := objc.Send[HKAudiogramConductionType](h_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKAudiogramSensitivityTest */


