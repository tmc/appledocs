// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HKAnchoredObjectQuery] class.
var (
	HKAnchoredObjectQueryClass     _HKAnchoredObjectQueryClass
	HKAnchoredObjectQueryClassOnce sync.Once
)

func getHKAnchoredObjectQueryClass() _HKAnchoredObjectQueryClass {
	HKAnchoredObjectQueryClassOnce.Do(func() {
		HKAnchoredObjectQueryClass = _HKAnchoredObjectQueryClass{objc.GetClass("HKAnchoredObjectQuery")}
	})
	return HKAnchoredObjectQueryClass
}

type _HKAnchoredObjectQueryClass struct {
	class objc.Class
}

// An interface definition for the [HKAnchoredObjectQuery] class.
type IHKAnchoredObjectQuery interface {
	IHKQuery
	// properties:
	UpdateHandler() unsafe.Pointer
	SetUpdateHandler(value unsafe.Pointer)
	HKObjectQueryNoLimit() int
	// methods:
}

// A query that returns changes to the HealthKit store, including a snapshot of new changes and continuous monitoring as a long-running query.
//
// Anchored object queries provide an easy way to search for new data in the HealthKit store. An returns an anchor value that corresponds to the last sample or deleted object received by that query. Subsequent queries can use this anchor to restrict their results to only newer saved or deleted objects. Anchored object queries are mostly immutable. You can assign the query’s property after instantiating the object, but you must set all other properties when you instantiate the object. You can’t change them.


// A query that returns changes to the HealthKit store, including a snapshot of new changes and continuous monitoring as a long-running query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAnchoredObjectQuery
type HKAnchoredObjectQuery struct {
	HKQuery
}

// HKAnchoredObjectQueryFrom constructs a [HKAnchoredObjectQuery] from an unsafe.Pointer.
//
// A query that returns changes to the HealthKit store, including a snapshot of new changes and continuous monitoring as a long-running query.
func HKAnchoredObjectQueryFrom(ptr unsafe.Pointer) HKAnchoredObjectQuery {
	return HKAnchoredObjectQuery{
		HKQuery: HKQueryFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HKAnchoredObjectQueryClass) Alloc() HKAnchoredObjectQuery {
	rv := objc.Send[HKAnchoredObjectQuery](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKAnchoredObjectQueryClass) New() HKAnchoredObjectQuery {
	rv := objc.Send[HKAnchoredObjectQuery](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKAnchoredObjectQuery) Init() HKAnchoredObjectQuery {
	rv := objc.Send[HKAnchoredObjectQuery](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKAnchoredObjectQuery) Autorelease() HKAnchoredObjectQuery {
	rv := objc.Send[HKAnchoredObjectQuery](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKAnchoredObjectQuery creates a new HKAnchoredObjectQuery instance.
func NewHKAnchoredObjectQuery() HKAnchoredObjectQuery {
	return getHKAnchoredObjectQueryClass().New()
}



// Handler for monitoring updates to the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkanchoredobjectquery/updatehandler
func (h_ HKAnchoredObjectQuery) UpdateHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("updateHandler"))
	return rv
}


// Handler for monitoring updates to the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkanchoredobjectquery/updatehandler
func (h_ HKAnchoredObjectQuery) SetUpdateHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setUpdateHandler:"), value)
}


// A value indicating that the query returns all the matching samples in the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkobjectquerynolimit
func (h_ HKAnchoredObjectQuery) HKObjectQueryNoLimit() int {
	rv := objc.Send[int](h_.ID, objc.Sel("HKObjectQueryNoLimit"))
	return rv
}



