// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent] class.
var (
	MTROTASoftwareUpdateRequestorClusterVersionAppliedEventClass     _MTROTASoftwareUpdateRequestorClusterVersionAppliedEventClass
	MTROTASoftwareUpdateRequestorClusterVersionAppliedEventClassOnce sync.Once
)

func getMTROTASoftwareUpdateRequestorClusterVersionAppliedEventClass() _MTROTASoftwareUpdateRequestorClusterVersionAppliedEventClass {
	MTROTASoftwareUpdateRequestorClusterVersionAppliedEventClassOnce.Do(func() {
		MTROTASoftwareUpdateRequestorClusterVersionAppliedEventClass = _MTROTASoftwareUpdateRequestorClusterVersionAppliedEventClass{objc.GetClass("MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent")}
	})
	return MTROTASoftwareUpdateRequestorClusterVersionAppliedEventClass
}

type _MTROTASoftwareUpdateRequestorClusterVersionAppliedEventClass struct {
	class objc.Class
}

// An interface definition for the [MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent] class.
type IMTROTASoftwareUpdateRequestorClusterVersionAppliedEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent-94prr
type MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent struct {
	objectivec.Object
}

// MTROTASoftwareUpdateRequestorClusterVersionAppliedEventFrom constructs a [MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent] from an unsafe.Pointer.
func MTROTASoftwareUpdateRequestorClusterVersionAppliedEventFrom(ptr unsafe.Pointer) MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent {
	return MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROTASoftwareUpdateRequestorClusterVersionAppliedEventClass) Alloc() MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent {
	rv := objc.Send[MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROTASoftwareUpdateRequestorClusterVersionAppliedEventClass) New() MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent {
	rv := objc.Send[MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent) Init() MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent {
	rv := objc.Send[MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent) Autorelease() MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent {
	rv := objc.Send[MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROTASoftwareUpdateRequestorClusterVersionAppliedEvent creates a new MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent instance.
func NewMTROTASoftwareUpdateRequestorClusterVersionAppliedEvent() MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent {
	return getMTROTASoftwareUpdateRequestorClusterVersionAppliedEventClass().New()
}




