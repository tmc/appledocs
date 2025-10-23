// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HKQuery] class.
var (
	HKQueryClass     _HKQueryClass
	HKQueryClassOnce sync.Once
)

func getHKQueryClass() _HKQueryClass {
	HKQueryClassOnce.Do(func() {
		HKQueryClass = _HKQueryClass{objc.GetClass("HKQuery")}
	})
	return HKQueryClass
}

type _HKQueryClass struct {
	class objc.Class
}

// An interface definition for the [HKQuery] class.
type IHKQuery interface {
	objectivec.IObject
	// properties:
	HKPredicateKeyPathMetadata() string
	HKPredicateKeyPathUUID() string
	ObjectType() IHKObjectType
	SetObjectType(value IHKObjectType)
	Predicate() foundation.Predicate
	SetPredicate(value foundation.Predicate)
	SampleType() IHKSampleType
	SetSampleType(value IHKSampleType)
	// methods:
}

// An abstract class for all the query classes in HealthKit.
//
// The class is the basis for all the query objects that retrieve data from the HealthKit store. The class is an abstract class. You should never instantiate it directly. Instead, you always work with one of its concrete subclasses.


// An abstract class for all the query classes in HealthKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery
type HKQuery struct {
	objectivec.Object
}

// HKQueryFrom constructs a [HKQuery] from an unsafe.Pointer.
//
// An abstract class for all the query classes in HealthKit.
func HKQueryFrom(ptr unsafe.Pointer) HKQuery {
	return HKQuery{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKQueryClass) Alloc() HKQuery {
	rv := objc.Send[HKQuery](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKQueryClass) New() HKQuery {
	rv := objc.Send[HKQuery](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKQuery) Init() HKQuery {
	rv := objc.Send[HKQuery](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKQuery) Autorelease() HKQuery {
	rv := objc.Send[HKQuery](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKQuery creates a new HKQuery instance.
func NewHKQuery() HKQuery {
	return getHKQueryClass().New()
}



// The key path for accessing the object’s metadata dictionary inside a predicate format string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathmetadata
func (h_ HKQuery) HKPredicateKeyPathMetadata() string {
	rv := objc.Send[string](h_.ID, objc.Sel("HKPredicateKeyPathMetadata"))
	return rv
}


// The key path for accessing the object’s UUID inside a predicate format string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathuuid
func (h_ HKQuery) HKPredicateKeyPathUUID() string {
	rv := objc.Send[string](h_.ID, objc.Sel("HKPredicateKeyPathUUID"))
	return rv
}


// The type of objects being queried.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkquery/objecttype
func (h_ HKQuery) ObjectType() IHKObjectType {
	rv := objc.Send[HKObjectType](h_.ID, objc.Sel("objectType"))
	return rv
}


// The type of objects being queried.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkquery/objecttype
func (h_ HKQuery) SetObjectType(value IHKObjectType) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setObjectType:"), value)
}


// A predicate used to filter the objects returned from the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkquery/predicate
func (h_ HKQuery) Predicate() foundation.Predicate {
	rv := objc.Send[foundation.Predicate](h_.ID, objc.Sel("predicate"))
	return rv
}


// A predicate used to filter the objects returned from the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkquery/predicate
func (h_ HKQuery) SetPredicate(value foundation.Predicate) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setPredicate:"), value)
}


// The type of objects being queried.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkquery/sampletype
func (h_ HKQuery) SampleType() IHKSampleType {
	rv := objc.Send[HKSampleType](h_.ID, objc.Sel("sampleType"))
	return rv
}


// The type of objects being queried.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkquery/sampletype
func (h_ HKQuery) SetSampleType(value IHKSampleType) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setSampleType:"), value)
}



