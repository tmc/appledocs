// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTROperationalStateClusterOperationalErrorEvent] class.
var (
	MTROperationalStateClusterOperationalErrorEventClass     _MTROperationalStateClusterOperationalErrorEventClass
	MTROperationalStateClusterOperationalErrorEventClassOnce sync.Once
)

func getMTROperationalStateClusterOperationalErrorEventClass() _MTROperationalStateClusterOperationalErrorEventClass {
	MTROperationalStateClusterOperationalErrorEventClassOnce.Do(func() {
		MTROperationalStateClusterOperationalErrorEventClass = _MTROperationalStateClusterOperationalErrorEventClass{objc.GetClass("MTROperationalStateClusterOperationalErrorEvent")}
	})
	return MTROperationalStateClusterOperationalErrorEventClass
}

type _MTROperationalStateClusterOperationalErrorEventClass struct {
	class objc.Class
}

// An interface definition for the [MTROperationalStateClusterOperationalErrorEvent] class.
type IMTROperationalStateClusterOperationalErrorEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalStateClusterOperationalErrorEvent
type MTROperationalStateClusterOperationalErrorEvent struct {
	objectivec.Object
}

// MTROperationalStateClusterOperationalErrorEventFrom constructs a [MTROperationalStateClusterOperationalErrorEvent] from an unsafe.Pointer.
func MTROperationalStateClusterOperationalErrorEventFrom(ptr unsafe.Pointer) MTROperationalStateClusterOperationalErrorEvent {
	return MTROperationalStateClusterOperationalErrorEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROperationalStateClusterOperationalErrorEventClass) Alloc() MTROperationalStateClusterOperationalErrorEvent {
	rv := objc.Send[MTROperationalStateClusterOperationalErrorEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROperationalStateClusterOperationalErrorEventClass) New() MTROperationalStateClusterOperationalErrorEvent {
	rv := objc.Send[MTROperationalStateClusterOperationalErrorEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalStateClusterOperationalErrorEvent) Init() MTROperationalStateClusterOperationalErrorEvent {
	rv := objc.Send[MTROperationalStateClusterOperationalErrorEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalStateClusterOperationalErrorEvent) Autorelease() MTROperationalStateClusterOperationalErrorEvent {
	rv := objc.Send[MTROperationalStateClusterOperationalErrorEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalStateClusterOperationalErrorEvent creates a new MTROperationalStateClusterOperationalErrorEvent instance.
func NewMTROperationalStateClusterOperationalErrorEvent() MTROperationalStateClusterOperationalErrorEvent {
	return getMTROperationalStateClusterOperationalErrorEventClass().New()
}




