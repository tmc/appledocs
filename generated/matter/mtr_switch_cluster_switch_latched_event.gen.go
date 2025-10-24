// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRSwitchClusterSwitchLatchedEvent] class.
var (
	MTRSwitchClusterSwitchLatchedEventClass     _MTRSwitchClusterSwitchLatchedEventClass
	MTRSwitchClusterSwitchLatchedEventClassOnce sync.Once
)

func getMTRSwitchClusterSwitchLatchedEventClass() _MTRSwitchClusterSwitchLatchedEventClass {
	MTRSwitchClusterSwitchLatchedEventClassOnce.Do(func() {
		MTRSwitchClusterSwitchLatchedEventClass = _MTRSwitchClusterSwitchLatchedEventClass{objc.GetClass("MTRSwitchClusterSwitchLatchedEvent")}
	})
	return MTRSwitchClusterSwitchLatchedEventClass
}

type _MTRSwitchClusterSwitchLatchedEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRSwitchClusterSwitchLatchedEvent] class.
type IMTRSwitchClusterSwitchLatchedEvent interface {
	objectivec.IObject
	// properties:
	NewPosition() objc.IObject /* cross-framework: NSNumber */
	SetNewPosition(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSwitchClusterSwitchLatchedEvent
type MTRSwitchClusterSwitchLatchedEvent struct {
	objectivec.Object
}

// MTRSwitchClusterSwitchLatchedEventFrom constructs a [MTRSwitchClusterSwitchLatchedEvent] from an unsafe.Pointer.
func MTRSwitchClusterSwitchLatchedEventFrom(ptr unsafe.Pointer) MTRSwitchClusterSwitchLatchedEvent {
	return MTRSwitchClusterSwitchLatchedEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRSwitchClusterSwitchLatchedEventClass) Alloc() MTRSwitchClusterSwitchLatchedEvent {
	rv := objc.Send[MTRSwitchClusterSwitchLatchedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRSwitchClusterSwitchLatchedEventClass) New() MTRSwitchClusterSwitchLatchedEvent {
	rv := objc.Send[MTRSwitchClusterSwitchLatchedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRSwitchClusterSwitchLatchedEvent) Init() MTRSwitchClusterSwitchLatchedEvent {
	rv := objc.Send[MTRSwitchClusterSwitchLatchedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRSwitchClusterSwitchLatchedEvent) Autorelease() MTRSwitchClusterSwitchLatchedEvent {
	rv := objc.Send[MTRSwitchClusterSwitchLatchedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRSwitchClusterSwitchLatchedEvent creates a new MTRSwitchClusterSwitchLatchedEvent instance.
func NewMTRSwitchClusterSwitchLatchedEvent() MTRSwitchClusterSwitchLatchedEvent {
	return getMTRSwitchClusterSwitchLatchedEventClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrswitchclusterswitchlatchedevent/newposition
func (m_ MTRSwitchClusterSwitchLatchedEvent) NewPosition() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("newPosition"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrswitchclusterswitchlatchedevent/newposition
func (m_ MTRSwitchClusterSwitchLatchedEvent) SetNewPosition(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNewPosition:"), value)
}



