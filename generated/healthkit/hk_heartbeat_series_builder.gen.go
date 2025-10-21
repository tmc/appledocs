// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HKHeartbeatSeriesBuilder] class.
var (
	HKHeartbeatSeriesBuilderClass     _HKHeartbeatSeriesBuilderClass
	HKHeartbeatSeriesBuilderClassOnce sync.Once
)

func getHKHeartbeatSeriesBuilderClass() _HKHeartbeatSeriesBuilderClass {
	HKHeartbeatSeriesBuilderClassOnce.Do(func() {
		HKHeartbeatSeriesBuilderClass = _HKHeartbeatSeriesBuilderClass{objc.GetClass("HKHeartbeatSeriesBuilder")}
	})
	return HKHeartbeatSeriesBuilderClass
}

type _HKHeartbeatSeriesBuilderClass struct {
	class objc.Class
}

// An interface definition for the [HKHeartbeatSeriesBuilder] class.
type IHKHeartbeatSeriesBuilder interface {
	IHKSeriesBuilder
}

// A builder object for incrementally building a heartbeat series.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHeartbeatSeriesBuilder
type HKHeartbeatSeriesBuilder struct {
	HKSeriesBuilder
}

// HKHeartbeatSeriesBuilderFrom constructs a [HKHeartbeatSeriesBuilder] from an unsafe.Pointer.
//
// A builder object for incrementally building a heartbeat series.
func HKHeartbeatSeriesBuilderFrom(ptr unsafe.Pointer) HKHeartbeatSeriesBuilder {
	return HKHeartbeatSeriesBuilder{
		HKSeriesBuilder: HKSeriesBuilderFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HKHeartbeatSeriesBuilderClass) Alloc() HKHeartbeatSeriesBuilder {
	rv := objc.Send[HKHeartbeatSeriesBuilder](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKHeartbeatSeriesBuilderClass) New() HKHeartbeatSeriesBuilder {
	rv := objc.Send[HKHeartbeatSeriesBuilder](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKHeartbeatSeriesBuilder) Init() HKHeartbeatSeriesBuilder {
	rv := objc.Send[HKHeartbeatSeriesBuilder](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKHeartbeatSeriesBuilder) Autorelease() HKHeartbeatSeriesBuilder {
	rv := objc.Send[HKHeartbeatSeriesBuilder](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKHeartbeatSeriesBuilder creates a new HKHeartbeatSeriesBuilder instance.
func NewHKHeartbeatSeriesBuilder() HKHeartbeatSeriesBuilder {
	return getHKHeartbeatSeriesBuilderClass().New()
}




