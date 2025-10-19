// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Event] class.
var (
	eventClass     _EventClass
	eventClassOnce sync.Once
)

func getEventClass() _EventClass {
	eventClassOnce.Do(func() {
		eventClass = _EventClass{objc.GetClass("NSEvent")}
	})
	return eventClass
}

type _EventClass struct {
	class objc.Class
}

// An interface definition for the [Event] class.
type IEvent interface {
	objectivec.IObject
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

// Alloc allocates a new instance without initialization.
func (ec _EventClass) Alloc() Event {
	rv := objc.Send[Event](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _EventClass) New() Event {
	rv := objc.Send[Event](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ Event) Init() Event {
	rv := objc.Send[Event](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ Event) Autorelease() Event {
	rv := objc.Send[Event](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEvent creates a new Event instance.
func NewEvent() Event {
	return getEventClass().New()
}




