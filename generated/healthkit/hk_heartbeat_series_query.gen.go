// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HKHeartbeatSeriesQuery] class.
var (
	HKHeartbeatSeriesQueryClass     _HKHeartbeatSeriesQueryClass
	HKHeartbeatSeriesQueryClassOnce sync.Once
)

func getHKHeartbeatSeriesQueryClass() _HKHeartbeatSeriesQueryClass {
	HKHeartbeatSeriesQueryClassOnce.Do(func() {
		HKHeartbeatSeriesQueryClass = _HKHeartbeatSeriesQueryClass{objc.GetClass("HKHeartbeatSeriesQuery")}
	})
	return HKHeartbeatSeriesQueryClass
}

type _HKHeartbeatSeriesQueryClass struct {
	class objc.Class
}

// An interface definition for the [HKHeartbeatSeriesQuery] class.
type IHKHeartbeatSeriesQuery interface {
	IHKQuery
	// properties:
	// methods:
}

// A query that returns the heartbeat data contained in a heartbeat series sample.


// A query that returns the heartbeat data contained in a heartbeat series sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHeartbeatSeriesQuery
type HKHeartbeatSeriesQuery struct {
	HKQuery
}

// HKHeartbeatSeriesQueryFrom constructs a [HKHeartbeatSeriesQuery] from an unsafe.Pointer.
//
// A query that returns the heartbeat data contained in a heartbeat series sample.
func HKHeartbeatSeriesQueryFrom(ptr unsafe.Pointer) HKHeartbeatSeriesQuery {
	return HKHeartbeatSeriesQuery{
		HKQuery: HKQueryFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HKHeartbeatSeriesQueryClass) Alloc() HKHeartbeatSeriesQuery {
	rv := objc.Send[HKHeartbeatSeriesQuery](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKHeartbeatSeriesQueryClass) New() HKHeartbeatSeriesQuery {
	rv := objc.Send[HKHeartbeatSeriesQuery](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKHeartbeatSeriesQuery) Init() HKHeartbeatSeriesQuery {
	rv := objc.Send[HKHeartbeatSeriesQuery](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKHeartbeatSeriesQuery) Autorelease() HKHeartbeatSeriesQuery {
	rv := objc.Send[HKHeartbeatSeriesQuery](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKHeartbeatSeriesQuery creates a new HKHeartbeatSeriesQuery instance.
func NewHKHeartbeatSeriesQuery() HKHeartbeatSeriesQuery {
	return getHKHeartbeatSeriesQueryClass().New()
}




