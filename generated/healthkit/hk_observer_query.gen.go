// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKObserverQuery */


/* debug [class_header]: Header for HKObserverQuery */
// The class instance for the [HKObserverQuery] class.
var (
	HKObserverQueryClass     _HKObserverQueryClass
	HKObserverQueryClassOnce sync.Once
)

func getHKObserverQueryClass() _HKObserverQueryClass {
	HKObserverQueryClassOnce.Do(func() {
		HKObserverQueryClass = _HKObserverQueryClass{objc.GetClass("HKObserverQuery")}
	})
	return HKObserverQueryClass
}

type _HKObserverQueryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKObserverQuery */
// An interface definition for the [HKObserverQuery] class.
type IHKObserverQuery interface {
	IHKQuery
	
/* debug [class_interface_properties]: Properties for HKObserverQuery */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKObserverQuery */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKObserverQuery */
// Alloc allocates a new instance without initialization.
func (hc _HKObserverQueryClass) Alloc() HKObserverQuery {
	rv := objc.Send[HKObserverQuery](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKObserverQueryClass) New() HKObserverQuery {
	rv := objc.Send[HKObserverQuery](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKObserverQuery) Init() HKObserverQuery {
	rv := objc.Send[HKObserverQuery](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKObserverQuery) Autorelease() HKObserverQuery {
	rv := objc.Send[HKObserverQuery](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKObserverQuery creates a new HKObserverQuery instance.
func NewHKObserverQuery() HKObserverQuery {
	return getHKObserverQueryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKObserverQuery */
// A long-running query that monitors the HealthKit store and updates your app when the HealthKit store saves or deletes a matching sample.
//
// Observer queries set up a long-running task on a background queue. This task watches the HealthKit store, and alerts you when the store saves or removes matching data. Your app uses observer queries to respond to changes made by other apps and devices. Observer queries are immutable: You set their properties when you first create them, and you can’t change them.


// A long-running query that monitors the HealthKit store and updates your app when the HealthKit store saves or deletes a matching sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObserverQuery
type HKObserverQuery struct {
	HKQuery
}

// HKObserverQueryFrom constructs a [HKObserverQuery] from an unsafe.Pointer.
//
// A long-running query that monitors the HealthKit store and updates your app when the HealthKit store saves or deletes a matching sample.
func HKObserverQueryFrom(ptr unsafe.Pointer) HKObserverQuery {
	return HKObserverQuery{
		HKQuery: HKQueryFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKObserverQuery */

// Creates a query that monitors the HealthKit store and responds to any changes matching any of the query descriptors you provided.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObserverQuery/init(queryDescriptors:updateHandler:)
func NewHKObserverQueryWithQueryDescriptorsUpdateHandler(queryDescriptors []HKQueryDescriptor, updateHandler unsafe.Pointer) HKObserverQuery {
	instance := getHKObserverQueryClass().Alloc()
	rv := objc.Send[HKObserverQuery](instance.ID, objc.Sel("initWithQueryDescriptors:updateHandler:"), queryDescriptors, updateHandler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKObserverQueryWithQueryDescriptorsUpdateHandler */


// Instantiates and returns a query that monitors the HealthKit store and responds to changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObserverQuery/init(sampleType:predicate:updateHandler:)
func NewHKObserverQueryWithSampleTypePredicateUpdateHandler(sampleType IHKSampleType, predicate foundation.Predicate, updateHandler unsafe.Pointer) HKObserverQuery {
	instance := getHKObserverQueryClass().Alloc()
	rv := objc.Send[HKObserverQuery](instance.ID, objc.Sel("initWithSampleType:predicate:updateHandler:"), sampleType, predicate, updateHandler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKObserverQueryWithSampleTypePredicateUpdateHandler */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKObserverQuery */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKObserverQuery */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKObserverQuery */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKObserverQuery */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKObserverQuery */


