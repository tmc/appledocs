// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRBasicInformationClusterLeaveEvent] class.
var (
	MTRBasicInformationClusterLeaveEventClass     _MTRBasicInformationClusterLeaveEventClass
	MTRBasicInformationClusterLeaveEventClassOnce sync.Once
)

func getMTRBasicInformationClusterLeaveEventClass() _MTRBasicInformationClusterLeaveEventClass {
	MTRBasicInformationClusterLeaveEventClassOnce.Do(func() {
		MTRBasicInformationClusterLeaveEventClass = _MTRBasicInformationClusterLeaveEventClass{objc.GetClass("MTRBasicInformationClusterLeaveEvent")}
	})
	return MTRBasicInformationClusterLeaveEventClass
}

type _MTRBasicInformationClusterLeaveEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRBasicInformationClusterLeaveEvent] class.
type IMTRBasicInformationClusterLeaveEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBasicInformationClusterLeaveEvent
type MTRBasicInformationClusterLeaveEvent struct {
	objectivec.Object
}

// MTRBasicInformationClusterLeaveEventFrom constructs a [MTRBasicInformationClusterLeaveEvent] from an unsafe.Pointer.
func MTRBasicInformationClusterLeaveEventFrom(ptr unsafe.Pointer) MTRBasicInformationClusterLeaveEvent {
	return MTRBasicInformationClusterLeaveEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBasicInformationClusterLeaveEventClass) Alloc() MTRBasicInformationClusterLeaveEvent {
	rv := objc.Send[MTRBasicInformationClusterLeaveEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBasicInformationClusterLeaveEventClass) New() MTRBasicInformationClusterLeaveEvent {
	rv := objc.Send[MTRBasicInformationClusterLeaveEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBasicInformationClusterLeaveEvent) Init() MTRBasicInformationClusterLeaveEvent {
	rv := objc.Send[MTRBasicInformationClusterLeaveEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBasicInformationClusterLeaveEvent) Autorelease() MTRBasicInformationClusterLeaveEvent {
	rv := objc.Send[MTRBasicInformationClusterLeaveEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBasicInformationClusterLeaveEvent creates a new MTRBasicInformationClusterLeaveEvent instance.
func NewMTRBasicInformationClusterLeaveEvent() MTRBasicInformationClusterLeaveEvent {
	return getMTRBasicInformationClusterLeaveEventClass().New()
}




