// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [Event] class.
var (
	EventClass     _EventClass
	EventClassOnce sync.Once
)

func getEventClass() _EventClass {
	EventClassOnce.Do(func() {
		EventClass = _EventClass{objc.GetClass("NSEvent")}
	})
	return EventClass
}

type _EventClass struct {
	class objc.Class
}

// An interface definition for the [Event] class.
type IEvent interface {
	objectivec.IObject
	LocationInNode(node unsafe.Pointer) coregraphics.CGPoint
}

// An object that contains information about an input action, such as a mouse click or a key press.
//
// AppKit reports events that occur in a window to the app that created the window. Events include mouse clicks, key presses, and other types of input to the system. An object contains pertinent information about each event, such as the event type and when the event occurred. The event type defines what other information is available in the event object. For example, a keyboard event contains information about the pressed keys. Although you can create objects directly, you typically don’t. The system generates them automatically in response to input from the mouse, keyboard, trackpad, or other peripherals such as connected tablets. It enqueues those events in its event queue, and dequeues them when it’s ready to process them. The system delivers events to the most relevant object, which might be the first responder or the object where the event occurred. For example, the system delivers mouse-click events to the view that contains the event location. To handle events, add support to your app’s objects. You can also use gesture recognizers to handle some events for you and execute your app’s code at appropriate times. For more information, see the reference. You can also monitor the events your app receives and modify or cancel some events as needed. Install a local monitor using the method to detect specific types of events and take action when your app receives them. Install a global monitor using the method to monitor events systemwide, although without the ability to modify them.
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


// Installs an event monitor that receives copies of events the system posts to this app prior to their dispatch.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/addLocalMonitorForEvents(matching:handler:)
func (ec _EventClass) AddLocalMonitorForEventsMatchingMaskHandler(mask unsafe.Pointer, block unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(ec.class), objc.Sel("addLocalMonitorForEventsMatchingMask:handler:"), mask, block)
	return rv
}

// Reports the current mouse position in screen coordinates.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/mouseLocation
func (ec _EventClass) MouseLocation() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](objc.ID(ec.class), objc.Sel("mouseLocation"))
	return rv
}
// Returns the location of the receiver in the coordinate system of the given node.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/location(in:)
func (e_ Event) LocationInNode(node unsafe.Pointer) coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](e_.ID, objc.Sel("locationInNode:"), node)
	return rv
}

// The Core Graphics event object corresponding to this event.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/cgEvent
func (e_ Event) CGEvent() coregraphics.CGEventRef {
	rv := objc.Send[coregraphics.CGEventRef](e_.ID, objc.Sel("CGEvent"))
	return rv
}

// An opaque Carbon type associated with this event.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/eventRef
func (e_ Event) EventRef() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("eventRef"))
	return rv
}

// The event location in the base coordinate system of the associated window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/locationInWindow
func (e_ Event) LocationInWindow() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](e_.ID, objc.Sel("locationInWindow"))
	return rv
}

// An integer bit field that indicates the pressed modifier keys.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/modifierFlags-swift.property
func (e_ Event) ModifierFlags() EventModifierFlags {
	rv := objc.Send[EventModifierFlags](e_.ID, objc.Sel("modifierFlags"))
	return rv
}

// Reports the current mouse position in screen coordinates.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/mouseLocation
func (e_ Event) MouseLocation() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](e_.ID, objc.Sel("mouseLocation"))
	return rv
}

// The time when the event occurred in seconds since system startup.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/timestamp
func (e_ Event) Timestamp() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("timestamp"))
	return rv
}

// The window object associated with the event.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/window
func (e_ Event) Window() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("window"))
	return rv
}

// The identifier for the window device associated with the event.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/windowNumber
func (e_ Event) WindowNumber() int {
	rv := objc.Send[int](e_.ID, objc.Sel("windowNumber"))
	return rv
}



