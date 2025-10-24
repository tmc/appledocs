// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKVisionPrism */


/* debug [class_header]: Header for HKVisionPrism */
// The class instance for the [HKVisionPrism] class.
var (
	HKVisionPrismClass     _HKVisionPrismClass
	HKVisionPrismClassOnce sync.Once
)

func getHKVisionPrismClass() _HKVisionPrismClass {
	HKVisionPrismClassOnce.Do(func() {
		HKVisionPrismClass = _HKVisionPrismClass{objc.GetClass("HKVisionPrism")}
	})
	return HKVisionPrismClass
}

type _HKVisionPrismClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKVisionPrism */
// An interface definition for the [HKVisionPrism] class.
type IHKVisionPrism interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKVisionPrism */
	// properties:
	Amount() IHKQuantity
	Angle() IHKQuantity
	Eye() HKVisionEye
	HorizontalAmount() IHKQuantity
	HorizontalBase() HKPrismBase
	VerticalAmount() IHKQuantity
	VerticalBase() HKPrismBase
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKVisionPrism */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKVisionPrism */
// Alloc allocates a new instance without initialization.
func (hc _HKVisionPrismClass) Alloc() HKVisionPrism {
	rv := objc.Send[HKVisionPrism](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKVisionPrismClass) New() HKVisionPrism {
	rv := objc.Send[HKVisionPrism](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKVisionPrism) Init() HKVisionPrism {
	rv := objc.Send[HKVisionPrism](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKVisionPrism) Autorelease() HKVisionPrism {
	rv := objc.Send[HKVisionPrism](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKVisionPrism creates a new HKVisionPrism instance.
func NewHKVisionPrism() HKVisionPrism {
	return getHKVisionPrismClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKVisionPrism */
// Prescription data for eye alignment.
//
// To include prism information in a glasses prescription, start by creating an object. Then, pass this value to the ’s initializer. Finally, create the glasses prescription and save it to the HealthKit store.


// Prescription data for eye alignment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVisionPrism
type HKVisionPrism struct {
	objectivec.Object
}

// HKVisionPrismFrom constructs a [HKVisionPrism] from an unsafe.Pointer.
//
// Prescription data for eye alignment.
func HKVisionPrismFrom(ptr unsafe.Pointer) HKVisionPrism {
	return HKVisionPrism{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKVisionPrism */

// Creates a new vision prism object, using a single quantity and an alignment angle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVisionPrism/init(amount:angle:eye:)
func NewHKVisionPrismWithAmountAngleEye(amount IHKQuantity, angle IHKQuantity, eye HKVisionEye) HKVisionPrism {
	instance := getHKVisionPrismClass().Alloc()
	rv := objc.Send[HKVisionPrism](instance.ID, objc.Sel("initWithAmount:angle:eye:"), amount, angle, eye)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKVisionPrismWithAmountAngleEye */


// Creates a new vision prism object that separates the correction strength into horizontal and vertical components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVisionPrism/init(verticalAmount:verticalBase:horizontalAmount:horizontalBase:eye:)
func NewHKVisionPrismWithVerticalAmountVerticalBaseHorizontalAmountHorizontalBaseEye(verticalAmount IHKQuantity, verticalBase HKPrismBase, horizontalAmount IHKQuantity, horizontalBase HKPrismBase, eye HKVisionEye) HKVisionPrism {
	instance := getHKVisionPrismClass().Alloc()
	rv := objc.Send[HKVisionPrism](instance.ID, objc.Sel("initWithVerticalAmount:verticalBase:horizontalAmount:horizontalBase:eye:"), verticalAmount, verticalBase, horizontalAmount, horizontalBase, eye)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKVisionPrismWithVerticalAmountVerticalBaseHorizontalAmountHorizontalBaseEye */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKVisionPrism */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKVisionPrism */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKVisionPrism */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKVisionPrism */

// The strength of the correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVisionPrism/amount
func (h_ HKVisionPrism) Amount() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("amount"))
	return rv
}/* debug [instance_properties/getter]: amount */


// The orientation of the adjustment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVisionPrism/angle
func (h_ HKVisionPrism) Angle() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("angle"))
	return rv
}/* debug [instance_properties/getter]: angle */


// A value indicating which eye the correction applies to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVisionPrism/eye
func (h_ HKVisionPrism) Eye() HKVisionEye {
	rv := objc.Send[HKVisionEye](h_.ID, objc.Sel("eye"))
	return rv
}/* debug [instance_properties/getter]: eye */


// The strength of the horizontal correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVisionPrism/horizontalAmount
func (h_ HKVisionPrism) HorizontalAmount() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("horizontalAmount"))
	return rv
}/* debug [instance_properties/getter]: horizontalAmount */


// The orientation of the horizontal portion of the correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVisionPrism/horizontalBase
func (h_ HKVisionPrism) HorizontalBase() HKPrismBase {
	rv := objc.Send[HKPrismBase](h_.ID, objc.Sel("horizontalBase"))
	return rv
}/* debug [instance_properties/getter]: horizontalBase */


// The strength of the vertical correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVisionPrism/verticalAmount
func (h_ HKVisionPrism) VerticalAmount() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("verticalAmount"))
	return rv
}/* debug [instance_properties/getter]: verticalAmount */


// The orientation of the vertical portion of the correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVisionPrism/verticalBase
func (h_ HKVisionPrism) VerticalBase() HKPrismBase {
	rv := objc.Send[HKPrismBase](h_.ID, objc.Sel("verticalBase"))
	return rv
}/* debug [instance_properties/getter]: verticalBase */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKVisionPrism */


