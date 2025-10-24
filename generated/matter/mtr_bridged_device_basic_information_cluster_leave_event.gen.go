// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRBridgedDeviceBasicInformationClusterLeaveEvent] class.
var (
	MTRBridgedDeviceBasicInformationClusterLeaveEventClass     _MTRBridgedDeviceBasicInformationClusterLeaveEventClass
	MTRBridgedDeviceBasicInformationClusterLeaveEventClassOnce sync.Once
)

func getMTRBridgedDeviceBasicInformationClusterLeaveEventClass() _MTRBridgedDeviceBasicInformationClusterLeaveEventClass {
	MTRBridgedDeviceBasicInformationClusterLeaveEventClassOnce.Do(func() {
		MTRBridgedDeviceBasicInformationClusterLeaveEventClass = _MTRBridgedDeviceBasicInformationClusterLeaveEventClass{objc.GetClass("MTRBridgedDeviceBasicInformationClusterLeaveEvent")}
	})
	return MTRBridgedDeviceBasicInformationClusterLeaveEventClass
}

type _MTRBridgedDeviceBasicInformationClusterLeaveEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRBridgedDeviceBasicInformationClusterLeaveEvent] class.
type IMTRBridgedDeviceBasicInformationClusterLeaveEvent interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBridgedDeviceBasicInformationClusterLeaveEvent
type MTRBridgedDeviceBasicInformationClusterLeaveEvent struct {
	objectivec.Object
}

// MTRBridgedDeviceBasicInformationClusterLeaveEventFrom constructs a [MTRBridgedDeviceBasicInformationClusterLeaveEvent] from an unsafe.Pointer.
func MTRBridgedDeviceBasicInformationClusterLeaveEventFrom(ptr unsafe.Pointer) MTRBridgedDeviceBasicInformationClusterLeaveEvent {
	return MTRBridgedDeviceBasicInformationClusterLeaveEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBridgedDeviceBasicInformationClusterLeaveEventClass) Alloc() MTRBridgedDeviceBasicInformationClusterLeaveEvent {
	rv := objc.Send[MTRBridgedDeviceBasicInformationClusterLeaveEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBridgedDeviceBasicInformationClusterLeaveEventClass) New() MTRBridgedDeviceBasicInformationClusterLeaveEvent {
	rv := objc.Send[MTRBridgedDeviceBasicInformationClusterLeaveEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBridgedDeviceBasicInformationClusterLeaveEvent) Init() MTRBridgedDeviceBasicInformationClusterLeaveEvent {
	rv := objc.Send[MTRBridgedDeviceBasicInformationClusterLeaveEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBridgedDeviceBasicInformationClusterLeaveEvent) Autorelease() MTRBridgedDeviceBasicInformationClusterLeaveEvent {
	rv := objc.Send[MTRBridgedDeviceBasicInformationClusterLeaveEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBridgedDeviceBasicInformationClusterLeaveEvent creates a new MTRBridgedDeviceBasicInformationClusterLeaveEvent instance.
func NewMTRBridgedDeviceBasicInformationClusterLeaveEvent() MTRBridgedDeviceBasicInformationClusterLeaveEvent {
	return getMTRBridgedDeviceBasicInformationClusterLeaveEventClass().New()
}




