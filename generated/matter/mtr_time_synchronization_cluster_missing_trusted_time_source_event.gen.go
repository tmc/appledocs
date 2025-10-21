// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent] class.
var (
	MTRTimeSynchronizationClusterMissingTrustedTimeSourceEventClass     _MTRTimeSynchronizationClusterMissingTrustedTimeSourceEventClass
	MTRTimeSynchronizationClusterMissingTrustedTimeSourceEventClassOnce sync.Once
)

func getMTRTimeSynchronizationClusterMissingTrustedTimeSourceEventClass() _MTRTimeSynchronizationClusterMissingTrustedTimeSourceEventClass {
	MTRTimeSynchronizationClusterMissingTrustedTimeSourceEventClassOnce.Do(func() {
		MTRTimeSynchronizationClusterMissingTrustedTimeSourceEventClass = _MTRTimeSynchronizationClusterMissingTrustedTimeSourceEventClass{objc.GetClass("MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent")}
	})
	return MTRTimeSynchronizationClusterMissingTrustedTimeSourceEventClass
}

type _MTRTimeSynchronizationClusterMissingTrustedTimeSourceEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent] class.
type IMTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent
type MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent struct {
	objectivec.Object
}

// MTRTimeSynchronizationClusterMissingTrustedTimeSourceEventFrom constructs a [MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent] from an unsafe.Pointer.
func MTRTimeSynchronizationClusterMissingTrustedTimeSourceEventFrom(ptr unsafe.Pointer) MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent {
	return MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTimeSynchronizationClusterMissingTrustedTimeSourceEventClass) Alloc() MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent {
	rv := objc.Send[MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTimeSynchronizationClusterMissingTrustedTimeSourceEventClass) New() MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent {
	rv := objc.Send[MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent) Init() MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent {
	rv := objc.Send[MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent) Autorelease() MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent {
	rv := objc.Send[MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent creates a new MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent instance.
func NewMTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent() MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent {
	return getMTRTimeSynchronizationClusterMissingTrustedTimeSourceEventClass().New()
}




