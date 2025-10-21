// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBasicClusterReachableChangedEvent] class.
var (
	MTRBasicClusterReachableChangedEventClass     _MTRBasicClusterReachableChangedEventClass
	MTRBasicClusterReachableChangedEventClassOnce sync.Once
)

func getMTRBasicClusterReachableChangedEventClass() _MTRBasicClusterReachableChangedEventClass {
	MTRBasicClusterReachableChangedEventClassOnce.Do(func() {
		MTRBasicClusterReachableChangedEventClass = _MTRBasicClusterReachableChangedEventClass{objc.GetClass("MTRBasicClusterReachableChangedEvent")}
	})
	return MTRBasicClusterReachableChangedEventClass
}

type _MTRBasicClusterReachableChangedEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRBasicClusterReachableChangedEvent] class.
type IMTRBasicClusterReachableChangedEvent interface {
	IMTRBasicInformationClusterReachableChangedEvent
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBasicClusterReachableChangedEvent
type MTRBasicClusterReachableChangedEvent struct {
	MTRBasicInformationClusterReachableChangedEvent
}

// MTRBasicClusterReachableChangedEventFrom constructs a [MTRBasicClusterReachableChangedEvent] from an unsafe.Pointer.
func MTRBasicClusterReachableChangedEventFrom(ptr unsafe.Pointer) MTRBasicClusterReachableChangedEvent {
	return MTRBasicClusterReachableChangedEvent{
		MTRBasicInformationClusterReachableChangedEvent: MTRBasicInformationClusterReachableChangedEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBasicClusterReachableChangedEventClass) Alloc() MTRBasicClusterReachableChangedEvent {
	rv := objc.Send[MTRBasicClusterReachableChangedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBasicClusterReachableChangedEventClass) New() MTRBasicClusterReachableChangedEvent {
	rv := objc.Send[MTRBasicClusterReachableChangedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBasicClusterReachableChangedEvent) Init() MTRBasicClusterReachableChangedEvent {
	rv := objc.Send[MTRBasicClusterReachableChangedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBasicClusterReachableChangedEvent) Autorelease() MTRBasicClusterReachableChangedEvent {
	rv := objc.Send[MTRBasicClusterReachableChangedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBasicClusterReachableChangedEvent creates a new MTRBasicClusterReachableChangedEvent instance.
func NewMTRBasicClusterReachableChangedEvent() MTRBasicClusterReachableChangedEvent {
	return getMTRBasicClusterReachableChangedEventClass().New()
}




