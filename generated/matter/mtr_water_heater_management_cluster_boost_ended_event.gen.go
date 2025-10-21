// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRWaterHeaterManagementClusterBoostEndedEvent] class.
var (
	MTRWaterHeaterManagementClusterBoostEndedEventClass     _MTRWaterHeaterManagementClusterBoostEndedEventClass
	MTRWaterHeaterManagementClusterBoostEndedEventClassOnce sync.Once
)

func getMTRWaterHeaterManagementClusterBoostEndedEventClass() _MTRWaterHeaterManagementClusterBoostEndedEventClass {
	MTRWaterHeaterManagementClusterBoostEndedEventClassOnce.Do(func() {
		MTRWaterHeaterManagementClusterBoostEndedEventClass = _MTRWaterHeaterManagementClusterBoostEndedEventClass{objc.GetClass("MTRWaterHeaterManagementClusterBoostEndedEvent")}
	})
	return MTRWaterHeaterManagementClusterBoostEndedEventClass
}

type _MTRWaterHeaterManagementClusterBoostEndedEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRWaterHeaterManagementClusterBoostEndedEvent] class.
type IMTRWaterHeaterManagementClusterBoostEndedEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterBoostEndedEvent
type MTRWaterHeaterManagementClusterBoostEndedEvent struct {
	objectivec.Object
}

// MTRWaterHeaterManagementClusterBoostEndedEventFrom constructs a [MTRWaterHeaterManagementClusterBoostEndedEvent] from an unsafe.Pointer.
func MTRWaterHeaterManagementClusterBoostEndedEventFrom(ptr unsafe.Pointer) MTRWaterHeaterManagementClusterBoostEndedEvent {
	return MTRWaterHeaterManagementClusterBoostEndedEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRWaterHeaterManagementClusterBoostEndedEventClass) Alloc() MTRWaterHeaterManagementClusterBoostEndedEvent {
	rv := objc.Send[MTRWaterHeaterManagementClusterBoostEndedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRWaterHeaterManagementClusterBoostEndedEventClass) New() MTRWaterHeaterManagementClusterBoostEndedEvent {
	rv := objc.Send[MTRWaterHeaterManagementClusterBoostEndedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWaterHeaterManagementClusterBoostEndedEvent) Init() MTRWaterHeaterManagementClusterBoostEndedEvent {
	rv := objc.Send[MTRWaterHeaterManagementClusterBoostEndedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWaterHeaterManagementClusterBoostEndedEvent) Autorelease() MTRWaterHeaterManagementClusterBoostEndedEvent {
	rv := objc.Send[MTRWaterHeaterManagementClusterBoostEndedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWaterHeaterManagementClusterBoostEndedEvent creates a new MTRWaterHeaterManagementClusterBoostEndedEvent instance.
func NewMTRWaterHeaterManagementClusterBoostEndedEvent() MTRWaterHeaterManagementClusterBoostEndedEvent {
	return getMTRWaterHeaterManagementClusterBoostEndedEventClass().New()
}




