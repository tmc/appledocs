
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Event] class.
var EventClass _EventClass

func init() {
	EventClass = _EventClass{objc.GetClass("NSEvent")}
}

type _EventClass struct {
	objc.Class
}

// An interface definition for the [Event] class.
type IEvent interface {
	ID() objc.ID
}

type Event struct {
	id objc.ID
}

func EventFrom(ptr unsafe.Pointer) Event {
	return Event{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (e_ Event) ID() objc.ID {
	return e_.id
}

// Alloc allocates a new instance without initialization.
func (ec _EventClass) Alloc() Event {
	rv := objc.Send[Event](objc.ID(ec.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ec _EventClass) New() Event {
	rv := objc.Send[Event](objc.ID(ec.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewEvent creates and returns a new initialized instance.
func NewEvent() Event {
	return EventClass.New()
}

// Init initializes the instance.
func (e_ Event) Init() Event {
	rv := objc.Send[Event](e_.ID(), selInit)
	return rv
}
// The Core Graphics event object corresponding to this event. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSEvent/cgEvent
func (e_ Event) CGEvent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID(), objc.RegisterName("CGEvent"))
	return rv
}
// An opaque Carbon type associated with this event. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSEvent/eventRef
func (e_ Event) EventRef() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID(), objc.RegisterName("eventRef"))
	return rv
}
// The event location in the base coordinate system of the associated window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSEvent/locationInWindow
func (e_ Event) LocationInWindow() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID(), objc.RegisterName("locationInWindow"))
	return rv
}
// An integer bit field that indicates the pressed modifier keys. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSEvent/modifierFlags-swift.property
func (e_ Event) ModifierFlags() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID(), objc.RegisterName("modifierFlags"))
	return rv
}
// The time when the event occurred in seconds since system startup. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSEvent/timestamp
func (e_ Event) Timestamp() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID(), objc.RegisterName("timestamp"))
	return rv
}
// The window object associated with the event. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSEvent/window
func (e_ Event) Window() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID(), objc.RegisterName("window"))
	return rv
}
// The identifier for the window device associated with the event. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSEvent/windowNumber
func (e_ Event) WindowNumber() int {
	rv := objc.Send[int](e_.ID(), objc.RegisterName("windowNumber"))
	return rv
}
