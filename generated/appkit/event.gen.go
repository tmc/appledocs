// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Event] class.
var eventClass = _EventClass{objc.GetClass("NSEvent")}

type _EventClass struct {
	class objc.Class
}

// An object that contains information about an input action, such as a mouse click or a key press. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent

type Event struct {
	objectivec.Object
}

// EventFrom constructs a [Event] from an unsafe.Pointer.
//
// An object that contains information about an input action, such as a mouse click or a key press.
func EventFrom(ptr unsafe.Pointer) Event {
	return Event{objectivec.Object{objc.ID(ptr)}}
}



