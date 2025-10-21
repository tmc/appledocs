// Code generated from Apple documentation for AVRouting. DO NOT EDIT.

package avrouting

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CustomDeviceRoute] class.
var (
	CustomDeviceRouteClass     _CustomDeviceRouteClass
	CustomDeviceRouteClassOnce sync.Once
)

func getCustomDeviceRouteClass() _CustomDeviceRouteClass {
	CustomDeviceRouteClassOnce.Do(func() {
		CustomDeviceRouteClass = _CustomDeviceRouteClass{objc.GetClass("AVCustomDeviceRoute")}
	})
	return CustomDeviceRouteClass
}

type _CustomDeviceRouteClass struct {
	class objc.Class
}

// An interface definition for the [CustomDeviceRoute] class.
type ICustomDeviceRoute interface {
	objectivec.IObject
}

// An object that represents a custom device route.
//
// Use the value of a route’s or property to establish a connection to a device. Typically, only one of the properties provides a valid value, depending on the type of device. In certain cases, both properties may provide valid values, in which case your app determines which one to use.
//
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVCustomDeviceRoute
type CustomDeviceRoute struct {
	objectivec.Object
}

// CustomDeviceRouteFrom constructs a [CustomDeviceRoute] from an unsafe.Pointer.
//
// An object that represents a custom device route.
func CustomDeviceRouteFrom(ptr unsafe.Pointer) CustomDeviceRoute {
	return CustomDeviceRoute{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CustomDeviceRouteClass) Alloc() CustomDeviceRoute {
	rv := objc.Send[CustomDeviceRoute](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CustomDeviceRouteClass) New() CustomDeviceRoute {
	rv := objc.Send[CustomDeviceRoute](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CustomDeviceRoute) Init() CustomDeviceRoute {
	rv := objc.Send[CustomDeviceRoute](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CustomDeviceRoute) Autorelease() CustomDeviceRoute {
	rv := objc.Send[CustomDeviceRoute](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCustomDeviceRoute creates a new CustomDeviceRoute instance.
func NewCustomDeviceRoute() CustomDeviceRoute {
	return getCustomDeviceRouteClass().New()
}


// An identifier to use to establish a connection to a Bluetooth device.
//
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVCustomDeviceRoute/bluetoothIdentifier
func (c_ CustomDeviceRoute) BluetoothIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("bluetoothIdentifier"))
	return rv
}

// A local or remote endpoint to connect to.
//
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVCustomDeviceRoute/networkEndpoint
func (c_ CustomDeviceRoute) NetworkEndpoint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("networkEndpoint"))
	return rv
}

// A reason for an event, such as a user request to activate or deactivate a route.
//
// [Full Topic]: https://developer.apple.com/documentation/avrouting/avcustomroutingevent/reason
func (c_ CustomDeviceRoute) Reason() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("reason"))
	return rv
}


// SetReason sets the value of the reason property.
// A reason for an event, such as a user request to activate or deactivate a route.

//
// [Full Topic]: https://developer.apple.com/documentation/avrouting/avcustomroutingevent/reason
func (c_ CustomDeviceRoute) SetReason(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setReason:"), value)
}

// A route for the event.
//
// [Full Topic]: https://developer.apple.com/documentation/avrouting/avcustomroutingevent/route
func (c_ CustomDeviceRoute) Route() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("route"))
	return rv
}


// SetRoute sets the value of the route property.
// A route for the event.

//
// [Full Topic]: https://developer.apple.com/documentation/avrouting/avcustomroutingevent/route
func (c_ CustomDeviceRoute) SetRoute(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRoute:"), value)
}



