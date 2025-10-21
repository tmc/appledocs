// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [HKObserverQuery] class.
type IHKObserverQuery interface {
	IHKQuery
}

// A long-running query that monitors the HealthKit store and updates your app when the HealthKit store saves or deletes a matching sample.
//
// Observer queries set up a long-running task on a background queue. This task watches the HealthKit store, and alerts you when the store saves or removes matching data. Your app uses observer queries to respond to changes made by other apps and devices. Observer queries are immutable: You set their properties when you first create them, and you can’t change them.
//
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

// Alloc allocates a new instance without initialization.
func (hc _HKObserverQueryClass) Alloc() HKObserverQuery {
	rv := objc.Send[HKObserverQuery](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Creates a query that monitors the HealthKit store and responds to any changes matching any of the query descriptors you provided.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObserverQuery/init(queryDescriptors:updateHandler:)
func NewHKObserverQueryWithQueryDescriptorsUpdateHandler(queryDescriptors unsafe.Pointer, updateHandler unsafe.Pointer) HKObserverQuery {
	instance := getHKObserverQueryClass().Alloc()
	rv := objc.Send[HKObserverQuery](instance.ID, objc.Sel("initWithQueryDescriptors:updateHandler:"), queryDescriptors, updateHandler)
	rv.Autorelease()
	return rv
}

// Instantiates and returns a query that monitors the HealthKit store and responds to changes.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObserverQuery/init(sampleType:predicate:updateHandler:)
func NewHKObserverQueryWithSampleTypePredicateUpdateHandler(sampleType unsafe.Pointer, predicate unsafe.Pointer, updateHandler unsafe.Pointer) HKObserverQuery {
	instance := getHKObserverQueryClass().Alloc()
	rv := objc.Send[HKObserverQuery](instance.ID, objc.Sel("initWithSampleType:predicate:updateHandler:"), sampleType, predicate, updateHandler)
	rv.Autorelease()
	return rv
}



