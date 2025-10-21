// Code generated from Apple documentation for DeviceDiscoveryExtension. DO NOT EDIT.

package devicediscoveryextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [DDDeviceEvent] class.
var (
	DDDeviceEventClass     _DDDeviceEventClass
	DDDeviceEventClassOnce sync.Once
)

func getDDDeviceEventClass() _DDDeviceEventClass {
	DDDeviceEventClassOnce.Do(func() {
		DDDeviceEventClass = _DDDeviceEventClass{objc.GetClass("DDDeviceEvent")}
	})
	return DDDeviceEventClass
}

type _DDDeviceEventClass struct {
	class objc.Class
}

// An interface definition for the [DDDeviceEvent] class.
type IDDDeviceEvent interface {
	objectivec.IObject
}

// An object that provides a device or communicates its change in status.
//
// The extension creates and configures an instance of this class to represent a moment of interest in the device discovery life cycle. The event’s ( ) describes a particular status. For example, when the extension discovers a device of interest, it instantiates an instance of this class with the type . Then, the extension provides the discovered device to the system using for eventual display in the route picker view ( ).
//
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDeviceEvent
type DDDeviceEvent struct {
	objectivec.Object
}

// DDDeviceEventFrom constructs a [DDDeviceEvent] from an unsafe.Pointer.
//
// An object that provides a device or communicates its change in status.
func DDDeviceEventFrom(ptr unsafe.Pointer) DDDeviceEvent {
	return DDDeviceEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DDDeviceEventClass) Alloc() DDDeviceEvent {
	rv := objc.Send[DDDeviceEvent](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DDDeviceEventClass) New() DDDeviceEvent {
	rv := objc.Send[DDDeviceEvent](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DDDeviceEvent) Init() DDDeviceEvent {
	rv := objc.Send[DDDeviceEvent](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DDDeviceEvent) Autorelease() DDDeviceEvent {
	rv := objc.Send[DDDeviceEvent](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDDDeviceEvent creates a new DDDeviceEvent instance.
func NewDDDeviceEvent() DDDeviceEvent {
	return getDDDeviceEventClass().New()
}




// Creates an event object that conveys status for a discovered device of interest.
//
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDeviceEvent/init(eventType:device:)
func NewDDDeviceEventWithEventTypeDevice(type_ unsafe.Pointer, device unsafe.Pointer) DDDeviceEvent {
	instance := getDDDeviceEventClass().Alloc()
	rv := objc.Send[DDDeviceEvent](instance.ID, objc.Sel("initWithEventType:device:"), type_, device)
	rv.Autorelease()
	return rv
}


// An object that describes a third-party media receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDeviceEvent/device
func (d_ DDDeviceEvent) Device() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("device"))
	return rv
}

// A type for the event that describes the discovery status.
//
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDeviceEvent/eventType-swift.property
func (d_ DDDeviceEvent) EventType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("eventType"))
	return rv
}


