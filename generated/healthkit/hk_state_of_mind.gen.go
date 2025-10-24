// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKStateOfMind */


/* debug [class_header]: Header for HKStateOfMind */
// The class instance for the [HKStateOfMind] class.
var (
	HKStateOfMindClass     _HKStateOfMindClass
	HKStateOfMindClassOnce sync.Once
)

func getHKStateOfMindClass() _HKStateOfMindClass {
	HKStateOfMindClassOnce.Do(func() {
		HKStateOfMindClass = _HKStateOfMindClass{objc.GetClass("HKStateOfMind")}
	})
	return HKStateOfMindClass
}

type _HKStateOfMindClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKStateOfMind */
// An interface definition for the [HKStateOfMind] class.
type IHKStateOfMind interface {
	IHKSample
	
/* debug [class_interface_properties]: Properties for HKStateOfMind */
	// properties:
	Associations() []foundation.Number
	Kind() HKStateOfMindKind
	Labels() []foundation.Number
	Valence() float64
	ValenceClassification() HKStateOfMindValenceClassification
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKStateOfMind */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKStateOfMind */
// Alloc allocates a new instance without initialization.
func (hc _HKStateOfMindClass) Alloc() HKStateOfMind {
	rv := objc.Send[HKStateOfMind](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKStateOfMindClass) New() HKStateOfMind {
	rv := objc.Send[HKStateOfMind](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKStateOfMind) Init() HKStateOfMind {
	rv := objc.Send[HKStateOfMind](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKStateOfMind) Autorelease() HKStateOfMind {
	rv := objc.Send[HKStateOfMind](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKStateOfMind creates a new HKStateOfMind instance.
func NewHKStateOfMind() HKStateOfMind {
	return getHKStateOfMindClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKStateOfMind */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind
type HKStateOfMind struct {
	HKSample
}

// HKStateOfMindFrom constructs a [HKStateOfMind] from an unsafe.Pointer.
func HKStateOfMindFrom(ptr unsafe.Pointer) HKStateOfMind {
	return HKStateOfMind{
		HKSample: HKSampleFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKStateOfMind *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKStateOfMind */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/stateOfMindWithDate:kind:valence:labels:associations:
func (hc _HKStateOfMindClass) StateOfMindWithDateKindValenceLabelsAssociations(date objc.IObject /* cross-framework: NSDate */, kind HKStateOfMindKind, valence float64, labels []foundation.Number, associations []foundation.Number) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("stateOfMindWithDate:kind:valence:labels:associations:"), date, kind, valence, labels, associations)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=StateOfMindWithDateKindValenceLabelsAssociations) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/stateOfMindWithDate:kind:valence:labels:associations:metadata:
func (hc _HKStateOfMindClass) StateOfMindWithDateKindValenceLabelsAssociationsMetadata(date objc.IObject /* cross-framework: NSDate */, kind HKStateOfMindKind, valence float64, labels []foundation.Number, associations []foundation.Number, metadata foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("stateOfMindWithDate:kind:valence:labels:associations:metadata:"), date, kind, valence, labels, associations, metadata)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=StateOfMindWithDateKindValenceLabelsAssociationsMetadata) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKStateOfMind */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKStateOfMind */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKStateOfMind */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/associations-5vfw8
func (h_ HKStateOfMind) Associations() []foundation.Number {
	rv := objc.Send[[]foundation.Number](h_.ID, objc.Sel("associations"))
	return rv
}/* debug [instance_properties/getter]: associations */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/kind-swift.property
func (h_ HKStateOfMind) Kind() HKStateOfMindKind {
	rv := objc.Send[HKStateOfMindKind](h_.ID, objc.Sel("kind"))
	return rv
}/* debug [instance_properties/getter]: kind */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/labels-11jl3
func (h_ HKStateOfMind) Labels() []foundation.Number {
	rv := objc.Send[[]foundation.Number](h_.ID, objc.Sel("labels"))
	return rv
}/* debug [instance_properties/getter]: labels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/valence
func (h_ HKStateOfMind) Valence() float64 {
	rv := objc.Send[float64](h_.ID, objc.Sel("valence"))
	return rv
}/* debug [instance_properties/getter]: valence */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/valenceClassification-swift.property
func (h_ HKStateOfMind) ValenceClassification() HKStateOfMindValenceClassification {
	rv := objc.Send[HKStateOfMindValenceClassification](h_.ID, objc.Sel("valenceClassification"))
	return rv
}/* debug [instance_properties/getter]: valenceClassification */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKStateOfMind */



