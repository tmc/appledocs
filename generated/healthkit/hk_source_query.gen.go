// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKSourceQuery */


/* debug [class_header]: Header for HKSourceQuery */
// The class instance for the [HKSourceQuery] class.
var (
	HKSourceQueryClass     _HKSourceQueryClass
	HKSourceQueryClassOnce sync.Once
)

func getHKSourceQueryClass() _HKSourceQueryClass {
	HKSourceQueryClassOnce.Do(func() {
		HKSourceQueryClass = _HKSourceQueryClass{objc.GetClass("HKSourceQuery")}
	})
	return HKSourceQueryClass
}

type _HKSourceQueryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKSourceQuery */
// An interface definition for the [HKSourceQuery] class.
type IHKSourceQuery interface {
	IHKQuery
	
/* debug [class_interface_properties]: Properties for HKSourceQuery */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKSourceQuery */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKSourceQuery */
// Alloc allocates a new instance without initialization.
func (hc _HKSourceQueryClass) Alloc() HKSourceQuery {
	rv := objc.Send[HKSourceQuery](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKSourceQueryClass) New() HKSourceQuery {
	rv := objc.Send[HKSourceQuery](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKSourceQuery) Init() HKSourceQuery {
	rv := objc.Send[HKSourceQuery](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKSourceQuery) Autorelease() HKSourceQuery {
	rv := objc.Send[HKSourceQuery](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKSourceQuery creates a new HKSourceQuery instance.
func NewHKSourceQuery() HKSourceQuery {
	return getHKSourceQueryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKSourceQuery */
// A query that returns a list of sources, such as apps and devices, that have saved matching queries to the HealthKit store.
//
// Source queries return a list of sources that have saved samples matching the specified sample types. Sources can be apps or devices (like Apple Watch or Bluetooth heart-rate monitors). Source queries are immutable: Their properties are set when they are first created, and they can’t change.


// A query that returns a list of sources, such as apps and devices, that have saved matching queries to the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSourceQuery
type HKSourceQuery struct {
	HKQuery
}

// HKSourceQueryFrom constructs a [HKSourceQuery] from an unsafe.Pointer.
//
// A query that returns a list of sources, such as apps and devices, that have saved matching queries to the HealthKit store.
func HKSourceQueryFrom(ptr unsafe.Pointer) HKSourceQuery {
	return HKSourceQuery{
		HKQuery: HKQueryFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKSourceQuery */

// Instantiates and returns a source query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSourceQuery/init(sampleType:samplePredicate:completionHandler:)
func NewHKSourceQueryWithSampleTypeSamplePredicateCompletionHandler(sampleType IHKSampleType, objectPredicate foundation.Predicate, completionHandler unsafe.Pointer) HKSourceQuery {
	instance := getHKSourceQueryClass().Alloc()
	rv := objc.Send[HKSourceQuery](instance.ID, objc.Sel("initWithSampleType:samplePredicate:completionHandler:"), sampleType, objectPredicate, completionHandler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKSourceQueryWithSampleTypeSamplePredicateCompletionHandler */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKSourceQuery */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKSourceQuery */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKSourceQuery */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKSourceQuery */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKSourceQuery */


