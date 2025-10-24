//go:build darwin && ios

// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for GCEventInteraction


// iOS-only properties

// The types of game controller events that should be delivered through the Game Controller framework.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCEventInteraction/handledEventTypes
func (g_ GCEventInteraction) HandledEventTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("handledEventTypes"))
	return rv
}
func (g_ GCEventInteraction) SetHandledEventTypes(value unsafe.Pointer) {
	g_.ID.Send(objc.RegisterName("setHandledEventTypes:"), value)
}

// A Boolean value that determines whether events are delivered exclusively through the Game Controller framework.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCEventInteraction/receivesEventsInView
func (g_ GCEventInteraction) ReceivesEventsInView() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("receivesEventsInView"))
	return rv
}
func (g_ GCEventInteraction) SetReceivesEventsInView(value bool) {
	g_.ID.Send(objc.RegisterName("setReceivesEventsInView:"), value)
}




