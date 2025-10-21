// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRAccessControlClusterAccessControlExtensionChangedEvent] class.
var (
	MTRAccessControlClusterAccessControlExtensionChangedEventClass     _MTRAccessControlClusterAccessControlExtensionChangedEventClass
	MTRAccessControlClusterAccessControlExtensionChangedEventClassOnce sync.Once
)

func getMTRAccessControlClusterAccessControlExtensionChangedEventClass() _MTRAccessControlClusterAccessControlExtensionChangedEventClass {
	MTRAccessControlClusterAccessControlExtensionChangedEventClassOnce.Do(func() {
		MTRAccessControlClusterAccessControlExtensionChangedEventClass = _MTRAccessControlClusterAccessControlExtensionChangedEventClass{objc.GetClass("MTRAccessControlClusterAccessControlExtensionChangedEvent")}
	})
	return MTRAccessControlClusterAccessControlExtensionChangedEventClass
}

type _MTRAccessControlClusterAccessControlExtensionChangedEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRAccessControlClusterAccessControlExtensionChangedEvent] class.
type IMTRAccessControlClusterAccessControlExtensionChangedEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterAccessControlExtensionChangedEvent
type MTRAccessControlClusterAccessControlExtensionChangedEvent struct {
	objectivec.Object
}

// MTRAccessControlClusterAccessControlExtensionChangedEventFrom constructs a [MTRAccessControlClusterAccessControlExtensionChangedEvent] from an unsafe.Pointer.
func MTRAccessControlClusterAccessControlExtensionChangedEventFrom(ptr unsafe.Pointer) MTRAccessControlClusterAccessControlExtensionChangedEvent {
	return MTRAccessControlClusterAccessControlExtensionChangedEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRAccessControlClusterAccessControlExtensionChangedEventClass) Alloc() MTRAccessControlClusterAccessControlExtensionChangedEvent {
	rv := objc.Send[MTRAccessControlClusterAccessControlExtensionChangedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRAccessControlClusterAccessControlExtensionChangedEventClass) New() MTRAccessControlClusterAccessControlExtensionChangedEvent {
	rv := objc.Send[MTRAccessControlClusterAccessControlExtensionChangedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAccessControlClusterAccessControlExtensionChangedEvent) Init() MTRAccessControlClusterAccessControlExtensionChangedEvent {
	rv := objc.Send[MTRAccessControlClusterAccessControlExtensionChangedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAccessControlClusterAccessControlExtensionChangedEvent) Autorelease() MTRAccessControlClusterAccessControlExtensionChangedEvent {
	rv := objc.Send[MTRAccessControlClusterAccessControlExtensionChangedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAccessControlClusterAccessControlExtensionChangedEvent creates a new MTRAccessControlClusterAccessControlExtensionChangedEvent instance.
func NewMTRAccessControlClusterAccessControlExtensionChangedEvent() MTRAccessControlClusterAccessControlExtensionChangedEvent {
	return getMTRAccessControlClusterAccessControlExtensionChangedEventClass().New()
}




