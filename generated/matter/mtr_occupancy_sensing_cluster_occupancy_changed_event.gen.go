// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROccupancySensingClusterOccupancyChangedEvent] class.
var (
	MTROccupancySensingClusterOccupancyChangedEventClass     _MTROccupancySensingClusterOccupancyChangedEventClass
	MTROccupancySensingClusterOccupancyChangedEventClassOnce sync.Once
)

func getMTROccupancySensingClusterOccupancyChangedEventClass() _MTROccupancySensingClusterOccupancyChangedEventClass {
	MTROccupancySensingClusterOccupancyChangedEventClassOnce.Do(func() {
		MTROccupancySensingClusterOccupancyChangedEventClass = _MTROccupancySensingClusterOccupancyChangedEventClass{objc.GetClass("MTROccupancySensingClusterOccupancyChangedEvent")}
	})
	return MTROccupancySensingClusterOccupancyChangedEventClass
}

type _MTROccupancySensingClusterOccupancyChangedEventClass struct {
	class objc.Class
}

// An interface definition for the [MTROccupancySensingClusterOccupancyChangedEvent] class.
type IMTROccupancySensingClusterOccupancyChangedEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROccupancySensingClusterOccupancyChangedEvent
type MTROccupancySensingClusterOccupancyChangedEvent struct {
	objectivec.Object
}

// MTROccupancySensingClusterOccupancyChangedEventFrom constructs a [MTROccupancySensingClusterOccupancyChangedEvent] from an unsafe.Pointer.
func MTROccupancySensingClusterOccupancyChangedEventFrom(ptr unsafe.Pointer) MTROccupancySensingClusterOccupancyChangedEvent {
	return MTROccupancySensingClusterOccupancyChangedEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROccupancySensingClusterOccupancyChangedEventClass) Alloc() MTROccupancySensingClusterOccupancyChangedEvent {
	rv := objc.Send[MTROccupancySensingClusterOccupancyChangedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROccupancySensingClusterOccupancyChangedEventClass) New() MTROccupancySensingClusterOccupancyChangedEvent {
	rv := objc.Send[MTROccupancySensingClusterOccupancyChangedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROccupancySensingClusterOccupancyChangedEvent) Init() MTROccupancySensingClusterOccupancyChangedEvent {
	rv := objc.Send[MTROccupancySensingClusterOccupancyChangedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROccupancySensingClusterOccupancyChangedEvent) Autorelease() MTROccupancySensingClusterOccupancyChangedEvent {
	rv := objc.Send[MTROccupancySensingClusterOccupancyChangedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROccupancySensingClusterOccupancyChangedEvent creates a new MTROccupancySensingClusterOccupancyChangedEvent instance.
func NewMTROccupancySensingClusterOccupancyChangedEvent() MTROccupancySensingClusterOccupancyChangedEvent {
	return getMTROccupancySensingClusterOccupancyChangedEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROccupancySensingClusterOccupancyChangedEvent/occupancy
func (m_ MTROccupancySensingClusterOccupancyChangedEvent) Occupancy() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("occupancy"))
	return rv
}


// SetOccupancy sets the value of the occupancy property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROccupancySensingClusterOccupancyChangedEvent/occupancy
func (m_ MTROccupancySensingClusterOccupancyChangedEvent) SetOccupancy(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOccupancy:"), value)
}



