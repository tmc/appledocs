// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [HKSourceQuery] class.
type IHKSourceQuery interface {
	IHKQuery
}

// A query that returns a list of sources, such as apps and devices, that have saved matching queries to the HealthKit store.
//
// Source queries return a list of sources that have saved samples matching the specified sample types. Sources can be apps or devices (like Apple Watch or Bluetooth heart-rate monitors). Source queries are immutable: Their properties are set when they are first created, and they can’t change.
//
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

// Alloc allocates a new instance without initialization.
func (hc _HKSourceQueryClass) Alloc() HKSourceQuery {
	rv := objc.Send[HKSourceQuery](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




