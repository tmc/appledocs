// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRBridgedDeviceBasicInformationClusterActiveChangedEvent] class.
var (
	MTRBridgedDeviceBasicInformationClusterActiveChangedEventClass     _MTRBridgedDeviceBasicInformationClusterActiveChangedEventClass
	MTRBridgedDeviceBasicInformationClusterActiveChangedEventClassOnce sync.Once
)

func getMTRBridgedDeviceBasicInformationClusterActiveChangedEventClass() _MTRBridgedDeviceBasicInformationClusterActiveChangedEventClass {
	MTRBridgedDeviceBasicInformationClusterActiveChangedEventClassOnce.Do(func() {
		MTRBridgedDeviceBasicInformationClusterActiveChangedEventClass = _MTRBridgedDeviceBasicInformationClusterActiveChangedEventClass{objc.GetClass("MTRBridgedDeviceBasicInformationClusterActiveChangedEvent")}
	})
	return MTRBridgedDeviceBasicInformationClusterActiveChangedEventClass
}

type _MTRBridgedDeviceBasicInformationClusterActiveChangedEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRBridgedDeviceBasicInformationClusterActiveChangedEvent] class.
type IMTRBridgedDeviceBasicInformationClusterActiveChangedEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBridgedDeviceBasicInformationClusterActiveChangedEvent
type MTRBridgedDeviceBasicInformationClusterActiveChangedEvent struct {
	objectivec.Object
}

// MTRBridgedDeviceBasicInformationClusterActiveChangedEventFrom constructs a [MTRBridgedDeviceBasicInformationClusterActiveChangedEvent] from an unsafe.Pointer.
func MTRBridgedDeviceBasicInformationClusterActiveChangedEventFrom(ptr unsafe.Pointer) MTRBridgedDeviceBasicInformationClusterActiveChangedEvent {
	return MTRBridgedDeviceBasicInformationClusterActiveChangedEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBridgedDeviceBasicInformationClusterActiveChangedEventClass) Alloc() MTRBridgedDeviceBasicInformationClusterActiveChangedEvent {
	rv := objc.Send[MTRBridgedDeviceBasicInformationClusterActiveChangedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBridgedDeviceBasicInformationClusterActiveChangedEventClass) New() MTRBridgedDeviceBasicInformationClusterActiveChangedEvent {
	rv := objc.Send[MTRBridgedDeviceBasicInformationClusterActiveChangedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBridgedDeviceBasicInformationClusterActiveChangedEvent) Init() MTRBridgedDeviceBasicInformationClusterActiveChangedEvent {
	rv := objc.Send[MTRBridgedDeviceBasicInformationClusterActiveChangedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBridgedDeviceBasicInformationClusterActiveChangedEvent) Autorelease() MTRBridgedDeviceBasicInformationClusterActiveChangedEvent {
	rv := objc.Send[MTRBridgedDeviceBasicInformationClusterActiveChangedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBridgedDeviceBasicInformationClusterActiveChangedEvent creates a new MTRBridgedDeviceBasicInformationClusterActiveChangedEvent instance.
func NewMTRBridgedDeviceBasicInformationClusterActiveChangedEvent() MTRBridgedDeviceBasicInformationClusterActiveChangedEvent {
	return getMTRBridgedDeviceBasicInformationClusterActiveChangedEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBridgedDeviceBasicInformationClusterActiveChangedEvent/promisedActiveDuration
func (m_ MTRBridgedDeviceBasicInformationClusterActiveChangedEvent) PromisedActiveDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("promisedActiveDuration"))
	return rv
}


// SetPromisedActiveDuration sets the value of the promisedActiveDuration property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBridgedDeviceBasicInformationClusterActiveChangedEvent/promisedActiveDuration
func (m_ MTRBridgedDeviceBasicInformationClusterActiveChangedEvent) SetPromisedActiveDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPromisedActiveDuration:"), value)
}


