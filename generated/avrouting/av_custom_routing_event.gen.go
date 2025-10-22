// Code generated from Apple documentation for AVRouting. DO NOT EDIT.

package avrouting

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CustomRoutingEvent] class.
var (
	CustomRoutingEventClass     _CustomRoutingEventClass
	CustomRoutingEventClassOnce sync.Once
)

func getCustomRoutingEventClass() _CustomRoutingEventClass {
	CustomRoutingEventClassOnce.Do(func() {
		CustomRoutingEventClass = _CustomRoutingEventClass{objc.GetClass("AVCustomRoutingEvent")}
	})
	return CustomRoutingEventClass
}

type _CustomRoutingEventClass struct {
	class objc.Class
}

// An interface definition for the [CustomRoutingEvent] class.
type ICustomRoutingEvent interface {
	objectivec.IObject
	Reason() CustomRoutingEventReason
	Route() AVCustomDeviceRoute
}

// An object that represents an event that occurs on a route.
//
// Depending on the route’s reason, apps establish or tear down a connection to a specified route.
//
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVCustomRoutingEvent
type CustomRoutingEvent struct {
	objectivec.Object
}

// CustomRoutingEventFrom constructs a [CustomRoutingEvent] from an unsafe.Pointer.
//
// An object that represents an event that occurs on a route.
func CustomRoutingEventFrom(ptr unsafe.Pointer) CustomRoutingEvent {
	return CustomRoutingEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CustomRoutingEventClass) Alloc() CustomRoutingEvent {
	rv := objc.Send[CustomRoutingEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CustomRoutingEventClass) New() CustomRoutingEvent {
	rv := objc.Send[CustomRoutingEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CustomRoutingEvent) Init() CustomRoutingEvent {
	rv := objc.Send[CustomRoutingEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CustomRoutingEvent) Autorelease() CustomRoutingEvent {
	rv := objc.Send[CustomRoutingEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCustomRoutingEvent creates a new CustomRoutingEvent instance.
func NewCustomRoutingEvent() CustomRoutingEvent {
	return getCustomRoutingEventClass().New()
}


// A reason for an event, such as a user request to activate or deactivate a route.
//
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVCustomRoutingEvent/reason
func (c_ CustomRoutingEvent) Reason() CustomRoutingEventReason {
	rv := objc.Send[CustomRoutingEventReason](c_.ID, objc.Sel("reason"))
	return rv
}

// A route for the event.
//
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVCustomRoutingEvent/route
func (c_ CustomRoutingEvent) Route() AVCustomDeviceRoute {
	rv := objc.Send[AVCustomDeviceRoute](c_.ID, objc.Sel("route"))
	return rv
}



