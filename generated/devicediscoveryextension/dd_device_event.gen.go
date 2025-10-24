// Code generated from Apple documentation for DeviceDiscoveryExtension. DO NOT EDIT.

package devicediscoveryextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class DDDeviceEvent */


/* debug [class_header]: Header for DDDeviceEvent */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DDDeviceEvent */
// An interface definition for the [DDDeviceEvent] class.
type IDDDeviceEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for DDDeviceEvent */
	// properties:
	Device() IDDDevice
	EventType() DDEventType
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DDDeviceEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DDDeviceEvent */
// Alloc allocates a new instance without initialization.
func (dc _DDDeviceEventClass) Alloc() DDDeviceEvent {
	rv := objc.Send[DDDeviceEvent](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DDDeviceEvent */
// An object that provides a device or communicates its change in status.
//
// The extension creates and configures an instance of this class to represent a moment of interest in the device discovery life cycle. The event’s ( ) describes a particular status. For example, when the extension discovers a device of interest, it instantiates an instance of this class with the type . Then, the extension provides the discovered device to the system using for eventual display in the route picker view ( ).


// An object that provides a device or communicates its change in status.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DDDeviceEvent */

// Creates an event object that conveys status for a discovered device of interest.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDeviceEvent/init(eventType:device:)
func NewDDDeviceEventWithEventTypeDevice(type_ DDEventType, device IDDDevice) DDDeviceEvent {
	instance := getDDDeviceEventClass().Alloc()
	rv := objc.Send[DDDeviceEvent](instance.ID, objc.Sel("initWithEventType:device:"), type_, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDDDeviceEventWithEventTypeDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DDDeviceEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DDDeviceEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DDDeviceEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DDDeviceEvent */

// An object that describes a third-party media receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDeviceEvent/device
func (d_ DDDeviceEvent) Device() IDDDevice {
	rv := objc.Send[DDDevice](d_.ID, objc.Sel("device"))
	return rv
}/* debug [instance_properties/getter]: device */


// A type for the event that describes the discovery status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDeviceEvent/eventType-swift.property
func (d_ DDDeviceEvent) EventType() DDEventType {
	rv := objc.Send[DDEventType](d_.ID, objc.Sel("eventType"))
	return rv
}/* debug [instance_properties/getter]: eventType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DDDeviceEvent */


