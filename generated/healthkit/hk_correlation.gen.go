// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HKCorrelation] class.
var (
	HKCorrelationClass     _HKCorrelationClass
	HKCorrelationClassOnce sync.Once
)

func getHKCorrelationClass() _HKCorrelationClass {
	HKCorrelationClassOnce.Do(func() {
		HKCorrelationClass = _HKCorrelationClass{objc.GetClass("HKCorrelation")}
	})
	return HKCorrelationClass
}

type _HKCorrelationClass struct {
	class objc.Class
}

// An interface definition for the [HKCorrelation] class.
type IHKCorrelation interface {
	IHKSample
	// properties:
	CorrelationType() IHKCorrelationType
	SetCorrelationType(value IHKCorrelationType)
	Objects() IHKSample
	SetObjects(value IHKSample)
	HKMetadataKeyFoodType() string
	HKPredicateKeyPathCorrelation() string
	// methods:
}

// A sample that groups multiple related samples into a single entry.
//
// HealthKit uses correlations to represent both blood pressure and food. Blood pressure correlations always include two quantity samples, representing the systolic and diastolic values. Food correlations can contain a wide range of dietary information about the food, including information about the fat, protein, carbohydrates, energy, and vitamins consumed. In general, a food correlation should include at least a sample. You can also add nutritional quantity samples for any other items you want to track. Use the key to indicate the food’s name. The class is a concrete subclass of the class. Correlations are immutable: You set the correlation’s properties when the object is first created, and they cannot change.


// A sample that groups multiple related samples into a single entry.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCorrelation
type HKCorrelation struct {
	HKSample
}

// HKCorrelationFrom constructs a [HKCorrelation] from an unsafe.Pointer.
//
// A sample that groups multiple related samples into a single entry.
func HKCorrelationFrom(ptr unsafe.Pointer) HKCorrelation {
	return HKCorrelation{
		HKSample: HKSampleFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HKCorrelationClass) Alloc() HKCorrelation {
	rv := objc.Send[HKCorrelation](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKCorrelationClass) New() HKCorrelation {
	rv := objc.Send[HKCorrelation](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKCorrelation) Init() HKCorrelation {
	rv := objc.Send[HKCorrelation](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKCorrelation) Autorelease() HKCorrelation {
	rv := objc.Send[HKCorrelation](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKCorrelation creates a new HKCorrelation instance.
func NewHKCorrelation() HKCorrelation {
	return getHKCorrelationClass().New()
}



// The type for this correlation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkcorrelation/correlationtype
func (h_ HKCorrelation) CorrelationType() IHKCorrelationType {
	rv := objc.Send[HKCorrelationType](h_.ID, objc.Sel("correlationType"))
	return rv
}


// The type for this correlation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkcorrelation/correlationtype
func (h_ HKCorrelation) SetCorrelationType(value IHKCorrelationType) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setCorrelationType:"), value)
}


// The set of sample objects that make up the correlation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkcorrelation/objects
func (h_ HKCorrelation) Objects() IHKSample {
	rv := objc.Send[HKSample](h_.ID, objc.Sel("objects"))
	return rv
}


// The set of sample objects that make up the correlation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkcorrelation/objects
func (h_ HKCorrelation) SetObjects(value IHKSample) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setObjects:"), value)
}


// The type of food that the HealthKit object represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkmetadatakeyfoodtype
func (h_ HKCorrelation) HKMetadataKeyFoodType() string {
	rv := objc.Send[string](h_.ID, objc.Sel("HKMetadataKeyFoodType"))
	return rv
}


// The key path for accessing the object’s correlation inside a predicate format string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathcorrelation
func (h_ HKCorrelation) HKPredicateKeyPathCorrelation() string {
	rv := objc.Send[string](h_.ID, objc.Sel("HKPredicateKeyPathCorrelation"))
	return rv
}



