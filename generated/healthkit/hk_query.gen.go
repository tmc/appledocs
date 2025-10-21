// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

// An abstract class for all the query classes in HealthKit.
//
// The class is the basis for all the query objects that retrieve data from the HealthKit store. The class is an abstract class. You should never instantiate it directly. Instead, you always work with one of its concrete subclasses.
//
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


// Returns a predicate for a specific FHIR resource.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForClinicalRecords(from:fhirResourceType:identifier:)
func (hc _HKQueryClass) PredicateForClinicalRecordsFromSourceFHIRResourceTypeIdentifier(source unsafe.Pointer, resourceType unsafe.Pointer, identifier string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("predicateForClinicalRecordsFromSource:FHIRResourceType:identifier:"), source, resourceType, objc.String(identifier))
	return rv
}

// Returns a predicate for a specific FHIR type.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForClinicalRecords(withFHIRResourceType:)
func (hc _HKQueryClass) PredicateForClinicalRecordsWithFHIRResourceType(resourceType unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("predicateForClinicalRecordsWithFHIRResourceType:"), resourceType)
	return rv
}

// Returns a predicate that matches any objects that have been associated with the provided workout.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForObjects(from:)-5irg9
func (hc _HKQueryClass) PredicateForObjectsFromWorkout(workout unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("predicateForObjectsFromWorkout:"), workout)
	return rv
}

// The type of objects being queried.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkquery/objecttype
func (h_ HKQuery) ObjectType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("objectType"))
	return rv
}


// SetObjectType sets the value of the objectType property.
// The type of objects being queried.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkquery/objecttype
func (h_ HKQuery) SetObjectType(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setObjectType:"), value)
}

// The key path for accessing the object’s metadata dictionary inside a predicate format string.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathmetadata
func (h_ HKQuery) HKPredicateKeyPathMetadata() string {
	rv := objc.Send[string](h_.ID, objc.Sel("HKPredicateKeyPathMetadata"))
	return rv
}

// The key path for accessing the object’s UUID inside a predicate format string.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathuuid
func (h_ HKQuery) HKPredicateKeyPathUUID() string {
	rv := objc.Send[string](h_.ID, objc.Sel("HKPredicateKeyPathUUID"))
	return rv
}

// The type of objects being queried.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkquery/sampletype
func (h_ HKQuery) SampleType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("sampleType"))
	return rv
}


// SetSampleType sets the value of the sampleType property.
// The type of objects being queried.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkquery/sampletype
func (h_ HKQuery) SetSampleType(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setSampleType:"), value)
}

// A predicate used to filter the objects returned from the HealthKit store.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkquery/predicate
func (h_ HKQuery) Predicate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("predicate"))
	return rv
}


// SetPredicate sets the value of the predicate property.
// A predicate used to filter the objects returned from the HealthKit store.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkquery/predicate
func (h_ HKQuery) SetPredicate(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setPredicate:"), value)
}



