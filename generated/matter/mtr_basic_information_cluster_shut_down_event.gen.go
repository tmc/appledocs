// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRBasicInformationClusterShutDownEvent] class.
var (
	MTRBasicInformationClusterShutDownEventClass     _MTRBasicInformationClusterShutDownEventClass
	MTRBasicInformationClusterShutDownEventClassOnce sync.Once
)

func getMTRBasicInformationClusterShutDownEventClass() _MTRBasicInformationClusterShutDownEventClass {
	MTRBasicInformationClusterShutDownEventClassOnce.Do(func() {
		MTRBasicInformationClusterShutDownEventClass = _MTRBasicInformationClusterShutDownEventClass{objc.GetClass("MTRBasicInformationClusterShutDownEvent")}
	})
	return MTRBasicInformationClusterShutDownEventClass
}

type _MTRBasicInformationClusterShutDownEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRBasicInformationClusterShutDownEvent] class.
type IMTRBasicInformationClusterShutDownEvent interface {
	objectivec.IObject
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBasicInformationClusterShutDownEvent
type MTRBasicInformationClusterShutDownEvent struct {
	objectivec.Object
}

// MTRBasicInformationClusterShutDownEventFrom constructs a [MTRBasicInformationClusterShutDownEvent] from an unsafe.Pointer.
func MTRBasicInformationClusterShutDownEventFrom(ptr unsafe.Pointer) MTRBasicInformationClusterShutDownEvent {
	return MTRBasicInformationClusterShutDownEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBasicInformationClusterShutDownEventClass) Alloc() MTRBasicInformationClusterShutDownEvent {
	rv := objc.Send[MTRBasicInformationClusterShutDownEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBasicInformationClusterShutDownEventClass) New() MTRBasicInformationClusterShutDownEvent {
	rv := objc.Send[MTRBasicInformationClusterShutDownEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBasicInformationClusterShutDownEvent) Init() MTRBasicInformationClusterShutDownEvent {
	rv := objc.Send[MTRBasicInformationClusterShutDownEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBasicInformationClusterShutDownEvent) Autorelease() MTRBasicInformationClusterShutDownEvent {
	rv := objc.Send[MTRBasicInformationClusterShutDownEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBasicInformationClusterShutDownEvent creates a new MTRBasicInformationClusterShutDownEvent instance.
func NewMTRBasicInformationClusterShutDownEvent() MTRBasicInformationClusterShutDownEvent {
	return getMTRBasicInformationClusterShutDownEventClass().New()
}
