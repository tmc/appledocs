//go:build darwin && ios

// Code generated from Apple documentation for AVRouting. DO NOT EDIT.

package avrouting

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CustomRoutingEvent


// iOS-only properties

// A reason for an event, such as a user request to activate or deactivate a route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVCustomRoutingEvent/reason
func (c_ CustomRoutingEvent) Reason() CustomRoutingEventReason {
	rv := objc.Send[CustomRoutingEventReason](c_.ID, objc.Sel("reason"))
	return rv
}

// A route for the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVCustomRoutingEvent/route
func (c_ CustomRoutingEvent) Route() IAVCustomDeviceRoute {
	rv := objc.Send[CustomDeviceRoute](c_.ID, objc.Sel("route"))
	return rv
}





