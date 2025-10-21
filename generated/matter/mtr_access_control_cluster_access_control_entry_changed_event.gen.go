// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRAccessControlClusterAccessControlEntryChangedEvent] class.
var (
	MTRAccessControlClusterAccessControlEntryChangedEventClass     _MTRAccessControlClusterAccessControlEntryChangedEventClass
	MTRAccessControlClusterAccessControlEntryChangedEventClassOnce sync.Once
)

func getMTRAccessControlClusterAccessControlEntryChangedEventClass() _MTRAccessControlClusterAccessControlEntryChangedEventClass {
	MTRAccessControlClusterAccessControlEntryChangedEventClassOnce.Do(func() {
		MTRAccessControlClusterAccessControlEntryChangedEventClass = _MTRAccessControlClusterAccessControlEntryChangedEventClass{objc.GetClass("MTRAccessControlClusterAccessControlEntryChangedEvent")}
	})
	return MTRAccessControlClusterAccessControlEntryChangedEventClass
}

type _MTRAccessControlClusterAccessControlEntryChangedEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRAccessControlClusterAccessControlEntryChangedEvent] class.
type IMTRAccessControlClusterAccessControlEntryChangedEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterAccessControlEntryChangedEvent
type MTRAccessControlClusterAccessControlEntryChangedEvent struct {
	objectivec.Object
}

// MTRAccessControlClusterAccessControlEntryChangedEventFrom constructs a [MTRAccessControlClusterAccessControlEntryChangedEvent] from an unsafe.Pointer.
func MTRAccessControlClusterAccessControlEntryChangedEventFrom(ptr unsafe.Pointer) MTRAccessControlClusterAccessControlEntryChangedEvent {
	return MTRAccessControlClusterAccessControlEntryChangedEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRAccessControlClusterAccessControlEntryChangedEventClass) Alloc() MTRAccessControlClusterAccessControlEntryChangedEvent {
	rv := objc.Send[MTRAccessControlClusterAccessControlEntryChangedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRAccessControlClusterAccessControlEntryChangedEventClass) New() MTRAccessControlClusterAccessControlEntryChangedEvent {
	rv := objc.Send[MTRAccessControlClusterAccessControlEntryChangedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAccessControlClusterAccessControlEntryChangedEvent) Init() MTRAccessControlClusterAccessControlEntryChangedEvent {
	rv := objc.Send[MTRAccessControlClusterAccessControlEntryChangedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAccessControlClusterAccessControlEntryChangedEvent) Autorelease() MTRAccessControlClusterAccessControlEntryChangedEvent {
	rv := objc.Send[MTRAccessControlClusterAccessControlEntryChangedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAccessControlClusterAccessControlEntryChangedEvent creates a new MTRAccessControlClusterAccessControlEntryChangedEvent instance.
func NewMTRAccessControlClusterAccessControlEntryChangedEvent() MTRAccessControlClusterAccessControlEntryChangedEvent {
	return getMTRAccessControlClusterAccessControlEntryChangedEventClass().New()
}




