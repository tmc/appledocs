// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	

	// properties:
	AbsoluteX() int
	AbsoluteY() int
	AbsoluteZ() int
	AssociatedEventsMask() EventMask
	ButtonMask() EventButtonMask
	ButtonNumber() int
	CGEvent() EventRef /* not a class type */
	Characters() foundation.foundation.INSString
	CharactersIgnoringModifiers() foundation.foundation.INSString
	ClickCount() int
	Data1() int
	Data2() int
	DeltaX() float64
	DeltaY() float64
	DeltaZ() float64
	EventNumber() int
	EventRef() objectivec.IObject
	HasPreciseScrollingDeltas() bool
	ARepeat() bool
	DirectionInvertedFromDevice() bool
	KeyCode() objectivec.IObject
	LocationInWindow() corefoundation.CGPoint
	Magnification() float64
	ModifierFlags() EventModifierFlags
	MomentumPhase() EventPhase
	Phase() EventPhase
	Pressure() float32
	PressureBehavior() PressureBehavior
	Rotation() float32
	ScrollingDeltaX() float64
	ScrollingDeltaY() float64
	Stage() int
	StageTransition() float64
	Subtype() EventSubtype
	TangentialPressure() float32
	Tilt() corefoundation.CGPoint
	Timestamp() float64
	TrackingArea() ITrackingArea
	TrackingNumber() int
	Type() EventType
	UserData() objectivec.IObject
	VendorDefined() objc.ID
	Window() IWindow
	WindowNumber() int
	CapabilityMask() int
	SetCapabilityMask(value int)
	Context() IGraphicsContext
	SetContext(value IGraphicsContext)
	DeviceID() int
	SetDeviceID(value int)
	IsARepeat() bool
	SetIsARepeat(value bool)
	IsDirectionInvertedFromDevice() bool
	SetIsDirectionInvertedFromDevice(value bool)
	IsEnteringProximity() bool
	SetIsEnteringProximity(value bool)
	PointingDeviceID() int
	SetPointingDeviceID(value int)
	PointingDeviceSerialNumber() int
	SetPointingDeviceSerialNumber(value int)
	PointingDeviceType() objectivec.IObject
	SetPointingDeviceType(value objectivec.IObject)
	SpecialKey() objectivec.IObject
	SetSpecialKey(value objectivec.IObject)
	SystemTabletID() int
	SetSystemTabletID(value int)
	TabletID() int
	SetTabletID(value int)
	UniqueID() uint64
	SetUniqueID(value uint64)
	VendorID() int
	SetVendorID(value int)
	VendorPointingDeviceType() int
	SetVendorPointingDeviceType(value int)


	

	// methods:
	AllTouches() unsafe.Pointer
	CharactersByApplyingModifiers(modifiers EventModifierFlags) foundation.String
	CoalescedTouchesForTouch(touch ITouch) []Touch
	TouchesForView(view IView) unsafe.Pointer
	TouchesMatchingPhaseInView(phase TouchPhase, view IView) unsafe.Pointer


}





// Alloc allocates a new instance without initialization.
func (ec _EventClass) Alloc() Event {
	rv := objc.Send[Event](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// An object that contains information about an input action, such as a mouse click or a key press.
//
// AppKit reports events that occur in a window to the app that created the window. Events include mouse clicks, key presses, and other types of input to the system. An object contains pertinent information about each event, such as the event type and when the event occurred. The event type defines what other information is available in the event object. For example, a keyboard event contains information about the pressed keys. Although you can create objects directly, you typically don’t. The system generates them automatically in response to input from the mouse, keyboard, trackpad, or other peripherals such as connected tablets. It enqueues those events in its event queue, and dequeues them when it’s ready to process them. The system delivers events to the most relevant object, which might be the first responder or the object where the event occurred. For example, the system delivers mouse-click events to the view that contains the event location. To handle events, add support to your app’s objects. You can also use gesture recognizers to handle some events for you and execute your app’s code at appropriate times. For more information, see the reference. You can also monitor the events your app receives and modify or cancel some events as needed. Install a local monitor using the method to detect specific types of events and take action when your app receives them. Install a global monitor using the method to monitor events systemwide, although without the ability to modify them.


// An object that contains information about an input action, such as a mouse click or a key press.
//
// [Full Topic]
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






// Creates and returns an event object for a Core Graphics event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/init(cgEvent:)
func NewEventWithCGEvent(cgEvent EventRef /* not a class type */) Event {
	rv := objc.Send[Event](objc.ID(getEventClass().class), objc.Sel("eventWithCGEvent:"), cgEvent)
	return rv
}


// Creates and returns a new event object for a Carbon event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/init(eventRef:)
func NewEventWithEventRef(eventRef objectivec.IObject) Event {
	rv := objc.Send[Event](objc.ID(getEventClass().class), objc.Sel("eventWithEventRef:"), eventRef)
	return rv
}







// Installs an event monitor that receives copies of events the system posts to other applications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/addGlobalMonitorForEvents(matching:handler:)
func (ec _EventClass) AddGlobalMonitorForEventsMatchingMaskHandler(mask EventMask, block unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(ec.class), objc.Sel("addGlobalMonitorForEventsMatchingMask:handler:"), mask, block)
	return rv
}


// Installs an event monitor that receives copies of events the system posts to this app prior to their dispatch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/addLocalMonitorForEvents(matching:handler:)
func (ec _EventClass) AddLocalMonitorForEventsMatchingMaskHandler(mask EventMask, block unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(ec.class), objc.Sel("addLocalMonitorForEventsMatchingMask:handler:"), mask, block)
	return rv
}


// Creates and returns a new event object that describes a tracking-rectangle or cursor-update event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/enterExitEvent(with:location:modifierFlags:timestamp:windowNumber:context:eventNumber:trackingNumber:userData:)
func (ec _EventClass) EnterExitEventWithTypeLocationModifierFlagsTimestampWindowNumberContextEventNumberTrackingNumberUserData(type_ EventType, location corefoundation.CGPoint, flags EventModifierFlags, time float64, wNum int, unusedPassNil IGraphicsContext, eNum int, tNum int, data objectivec.IObject) IEvent {
	rv := objc.Send[Event](objc.ID(ec.class), objc.Sel("enterExitEventWithType:location:modifierFlags:timestamp:windowNumber:context:eventNumber:trackingNumber:userData:"), type_, location, flags, time, wNum, unusedPassNil, eNum, tNum, data)
	return rv
}


// Creates and returns an event object for a Core Graphics event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/init(cgEvent:)
func (ec _EventClass) EventWithCGEvent(cgEvent EventRef /* not a class type */) IEvent {
	rv := objc.Send[Event](objc.ID(ec.class), objc.Sel("eventWithCGEvent:"), cgEvent)
	return rv
}


// Creates and returns a new event object for a Carbon event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/init(eventRef:)
func (ec _EventClass) EventWithEventRef(eventRef objectivec.IObject) IEvent {
	rv := objc.Send[Event](objc.ID(ec.class), objc.Sel("eventWithEventRef:"), eventRef)
	return rv
}


// Creates and returns a new event object that describes a key event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/keyEvent(with:location:modifierFlags:timestamp:windowNumber:context:characters:charactersIgnoringModifiers:isARepeat:keyCode:)
func (ec _EventClass) KeyEventWithTypeLocationModifierFlagsTimestampWindowNumberContextCharactersCharactersIgnoringModifiersIsARepeatKeyCode(type_ EventType, location corefoundation.CGPoint, flags EventModifierFlags, time float64, wNum int, unusedPassNil IGraphicsContext, keys foundation.foundation.INSString, ukeys foundation.foundation.INSString, flag bool, code objectivec.IObject) IEvent {
	rv := objc.Send[Event](objc.ID(ec.class), objc.Sel("keyEventWithType:location:modifierFlags:timestamp:windowNumber:context:characters:charactersIgnoringModifiers:isARepeat:keyCode:"), type_, location, flags, time, wNum, unusedPassNil, keys, ukeys, flag, code)
	return rv
}


// Creates and returns a new event object that describes a mouse-down, -up, -moved, or -dragged event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/mouseEvent(with:location:modifierFlags:timestamp:windowNumber:context:eventNumber:clickCount:pressure:)
func (ec _EventClass) MouseEventWithTypeLocationModifierFlagsTimestampWindowNumberContextEventNumberClickCountPressure(type_ EventType, location corefoundation.CGPoint, flags EventModifierFlags, time float64, wNum int, unusedPassNil IGraphicsContext, eNum int, cNum int, pressure float32) IEvent {
	rv := objc.Send[Event](objc.ID(ec.class), objc.Sel("mouseEventWithType:location:modifierFlags:timestamp:windowNumber:context:eventNumber:clickCount:pressure:"), type_, location, flags, time, wNum, unusedPassNil, eNum, cNum, pressure)
	return rv
}


// Creates and returns a new event object that describes a custom event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/otherEvent(with:location:modifierFlags:timestamp:windowNumber:context:subtype:data1:data2:)
func (ec _EventClass) OtherEventWithTypeLocationModifierFlagsTimestampWindowNumberContextSubtypeData1Data2(type_ EventType, location corefoundation.CGPoint, flags EventModifierFlags, time float64, wNum int, unusedPassNil IGraphicsContext, subtype objectivec.IObject, d1 int, d2 int) IEvent {
	rv := objc.Send[Event](objc.ID(ec.class), objc.Sel("otherEventWithType:location:modifierFlags:timestamp:windowNumber:context:subtype:data1:data2:"), type_, location, flags, time, wNum, unusedPassNil, subtype, d1, d2)
	return rv
}


// Removes the specified event monitor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/removeMonitor(_:)
func (ec _EventClass) RemoveMonitor(eventMonitor objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(ec.class), objc.Sel("removeMonitor:"), eventMonitor)
}







// The maximum number of seconds in which a second mouse click must occur for an event to be a double-click event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/doubleClickInterval
func (ec _EventClass) DoubleClickInterval() float64 {
	rv := objc.Send[float64](objc.ID(ec.class), objc.Sel("doubleClickInterval"))
	return rv
}

// A Boolean value that indicates whether the system coalesces mouse movement events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/isMouseCoalescingEnabled
func (ec _EventClass) MouseCoalescingEnabled() bool {
	rv := objc.Send[bool](objc.ID(ec.class), objc.Sel("mouseCoalescingEnabled"))
	return rv
}

// The number of seconds someone must hold down a key before the first key repeat event occurs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/keyRepeatDelay
func (ec _EventClass) KeyRepeatDelay() float64 {
	rv := objc.Send[float64](objc.ID(ec.class), objc.Sel("keyRepeatDelay"))
	return rv
}

// The number of seconds someone must hold down a key to generate key-repeat events after the initial delay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/keyRepeatInterval
func (ec _EventClass) KeyRepeatInterval() float64 {
	rv := objc.Send[float64](objc.ID(ec.class), objc.Sel("keyRepeatInterval"))
	return rv
}

// Reports the current mouse position in screen coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/mouseLocation
func (ec _EventClass) MouseLocation() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](objc.ID(ec.class), objc.Sel("mouseLocation"))
	return rv
}

// The indices of the currently pressed mouse buttons.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/pressedMouseButtons
func (ec _EventClass) PressedMouseButtons() uint {
	rv := objc.Send[uint](objc.ID(ec.class), objc.Sel("pressedMouseButtons"))
	return rv
}






// Returns all touch objects associated with the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/allTouches()
func (e_ Event) AllTouches() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("allTouches"))
	return rv
}


// Returns the new characters that result if you apply the specified modifier keys to the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/characters(byApplyingModifiers:)
func (e_ Event) CharactersByApplyingModifiers(modifiers EventModifierFlags) foundation.String {
	rv := objc.Send[foundation.String](e_.ID, objc.Sel("charactersByApplyingModifiers:"), modifiers)
	return rv
}


// Returns all of the touch objects associated with the specified main touch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/coalescedTouches(for:)
func (e_ Event) CoalescedTouchesForTouch(touch ITouch) []Touch {
	rv := objc.Send[[]Touch](e_.ID, objc.Sel("coalescedTouchesForTouch:"), touch)
	return rv
}


// Returns the touch objects from the event that belong to the specified view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/touches(for:)
func (e_ Event) TouchesForView(view IView) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("touchesForView:"), view)
	return rv
}


// Returns the touch objects associated with the specified phase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/touches(matching:in:)
func (e_ Event) TouchesMatchingPhaseInView(phase TouchPhase, view IView) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("touchesMatchingPhase:inView:"), phase, view)
	return rv
}







// The absolute x coordinate of a pointing device on its tablet at full tablet resolution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/absoluteX
func (e_ Event) AbsoluteX() int {
	rv := objc.Send[int](e_.ID, objc.Sel("absoluteX"))
	return rv
}


// The absolute y coordinate of a pointing device on its tablet at full tablet resolution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/absoluteY
func (e_ Event) AbsoluteY() int {
	rv := objc.Send[int](e_.ID, objc.Sel("absoluteY"))
	return rv
}


// The absolute z coordinate of pointing device on its tablet at full tablet resolution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/absoluteZ
func (e_ Event) AbsoluteZ() int {
	rv := objc.Send[int](e_.ID, objc.Sel("absoluteZ"))
	return rv
}


// The associated events mask of a mouse event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/associatedEventsMask
func (e_ Event) AssociatedEventsMask() EventMask {
	rv := objc.Send[EventMask](e_.ID, objc.Sel("associatedEventsMask"))
	return rv
}


// A bit mask identifying the buttons pressed for a tablet event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/buttonMask-swift.property
func (e_ Event) ButtonMask() EventButtonMask {
	rv := objc.Send[EventButtonMask](e_.ID, objc.Sel("buttonMask"))
	return rv
}


// The button number for a mouse event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/buttonNumber
func (e_ Event) ButtonNumber() int {
	rv := objc.Send[int](e_.ID, objc.Sel("buttonNumber"))
	return rv
}


// The Core Graphics event object corresponding to this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/cgEvent
func (e_ Event) CGEvent() EventRef /* not a class type */ {
	rv := objc.Send[EventRef](e_.ID, objc.Sel("CGEvent"))
	return rv
}


// The characters associated with a key-up or key-down event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/characters
func (e_ Event) Characters() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("characters"))
	return rv
}


// The characters generated by a key event as if no modifier key (except for Shift) applies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/charactersIgnoringModifiers
func (e_ Event) CharactersIgnoringModifiers() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("charactersIgnoringModifiers"))
	return rv
}


// The number of mouse clicks associated with a mouse-down or mouse-up event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/clickCount
func (e_ Event) ClickCount() int {
	rv := objc.Send[int](e_.ID, objc.Sel("clickCount"))
	return rv
}


// Additional data associated with this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/data1
func (e_ Event) Data1() int {
	rv := objc.Send[int](e_.ID, objc.Sel("data1"))
	return rv
}


// Additional data associated with this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/data2
func (e_ Event) Data2() int {
	rv := objc.Send[int](e_.ID, objc.Sel("data2"))
	return rv
}


// The x-coordinate change for scroll wheel, mouse-move, mouse-drag, and swipe events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/deltaX
func (e_ Event) DeltaX() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("deltaX"))
	return rv
}


// The y-coordinate change for scroll wheel, mouse-move, mouse-drag, and swipe events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/deltaY
func (e_ Event) DeltaY() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("deltaY"))
	return rv
}


// The z-coordinate change for a scroll wheel, mouse-move, or mouse-drag event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/deltaZ
func (e_ Event) DeltaZ() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("deltaZ"))
	return rv
}


// The maximum number of seconds in which a second mouse click must occur for an event to be a double-click event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/doubleClickInterval
func (e_ Event) DoubleClickInterval() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("doubleClickInterval"))
	return rv
}


// The counter value of the latest mouse or tracking-rectangle event object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/eventNumber
func (e_ Event) EventNumber() int {
	rv := objc.Send[int](e_.ID, objc.Sel("eventNumber"))
	return rv
}


// An opaque Carbon type associated with this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/eventRef
func (e_ Event) EventRef() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](e_.ID, objc.Sel("eventRef"))
	return rv
}


// A Boolean value that indicates whether precise scrolling deltas are available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/hasPreciseScrollingDeltas
func (e_ Event) HasPreciseScrollingDeltas() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("hasPreciseScrollingDeltas"))
	return rv
}


// A Boolean value that indicates whether the key event is a repeat.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/isARepeat
func (e_ Event) ARepeat() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("ARepeat"))
	return rv
}


// A Boolean value that indicates whether the user has changed the device inversion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/isDirectionInvertedFromDevice
func (e_ Event) DirectionInvertedFromDevice() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("directionInvertedFromDevice"))
	return rv
}


// A Boolean value that indicates whether the system coalesces mouse movement events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/isMouseCoalescingEnabled
func (e_ Event) MouseCoalescingEnabled() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("mouseCoalescingEnabled"))
	return rv
}


// A Boolean value that indicates whether the system coalesces mouse movement events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/isMouseCoalescingEnabled
func (e_ Event) SetMouseCoalescingEnabled(value bool) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setMouseCoalescingEnabled:"), value)
}


// The virtual code for the key associated with the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/keyCode
func (e_ Event) KeyCode() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](e_.ID, objc.Sel("keyCode"))
	return rv
}


// The number of seconds someone must hold down a key before the first key repeat event occurs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/keyRepeatDelay
func (e_ Event) KeyRepeatDelay() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("keyRepeatDelay"))
	return rv
}


// The number of seconds someone must hold down a key to generate key-repeat events after the initial delay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/keyRepeatInterval
func (e_ Event) KeyRepeatInterval() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("keyRepeatInterval"))
	return rv
}


// The event location in the base coordinate system of the associated window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/locationInWindow
func (e_ Event) LocationInWindow() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](e_.ID, objc.Sel("locationInWindow"))
	return rv
}


// The amount of change to add to a magnification gesture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/magnification
func (e_ Event) Magnification() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("magnification"))
	return rv
}


// An integer bit field that indicates the pressed modifier keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/modifierFlags-swift.property
func (e_ Event) ModifierFlags() EventModifierFlags {
	rv := objc.Send[EventModifierFlags](e_.ID, objc.Sel("modifierFlags"))
	return rv
}


// The momentum phase for a scroll or flick gesture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/momentumPhase
func (e_ Event) MomentumPhase() EventPhase {
	rv := objc.Send[EventPhase](e_.ID, objc.Sel("momentumPhase"))
	return rv
}


// Reports the current mouse position in screen coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/mouseLocation
func (e_ Event) MouseLocation() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](e_.ID, objc.Sel("mouseLocation"))
	return rv
}


// The phase of a gesture event, such as a magnify, scroll, or pressure change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/phase-swift.property
func (e_ Event) Phase() EventPhase {
	rv := objc.Send[EventPhase](e_.ID, objc.Sel("phase"))
	return rv
}


// The indices of the currently pressed mouse buttons.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/pressedMouseButtons
func (e_ Event) PressedMouseButtons() uint {
	rv := objc.Send[uint](e_.ID, objc.Sel("pressedMouseButtons"))
	return rv
}


// A normalized value that indicates the degree of pressure applied to an appropriate input device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/pressure
func (e_ Event) Pressure() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("pressure"))
	return rv
}


// The behavior and progression for a pressure event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/pressureBehavior-swift.property
func (e_ Event) PressureBehavior() PressureBehavior {
	rv := objc.Send[PressureBehavior](e_.ID, objc.Sel("pressureBehavior"))
	return rv
}


// The rotation in degrees of the tablet pointing device associated with this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/rotation
func (e_ Event) Rotation() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("rotation"))
	return rv
}


// The scroll wheel’s horizontal delta.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/scrollingDeltaX
func (e_ Event) ScrollingDeltaX() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("scrollingDeltaX"))
	return rv
}


// The scroll wheel’s vertical delta.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/scrollingDeltaY
func (e_ Event) ScrollingDeltaY() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("scrollingDeltaY"))
	return rv
}


// A value that indicates the stage of a pressure gesture event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/stage
func (e_ Event) Stage() int {
	rv := objc.Send[int](e_.ID, objc.Sel("stage"))
	return rv
}


// The transition value for the stage of a pressure gesture event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/stageTransition
func (e_ Event) StageTransition() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("stageTransition"))
	return rv
}


// The event’s subtype.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/subtype
func (e_ Event) Subtype() EventSubtype {
	rv := objc.Send[EventSubtype](e_.ID, objc.Sel("subtype"))
	return rv
}


// The tangential pressure on the device that generated this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/tangentialPressure
func (e_ Event) TangentialPressure() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("tangentialPressure"))
	return rv
}


// The scaled tilt values of the pointing device that generated this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/tilt
func (e_ Event) Tilt() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](e_.ID, objc.Sel("tilt"))
	return rv
}


// The time when the event occurred in seconds since system startup.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/timestamp
func (e_ Event) Timestamp() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("timestamp"))
	return rv
}


// The tracking area for the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/trackingArea
func (e_ Event) TrackingArea() ITrackingArea {
	rv := objc.Send[TrackingArea](e_.ID, objc.Sel("trackingArea"))
	return rv
}


// The identifier of a mouse-tracking event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/trackingNumber
func (e_ Event) TrackingNumber() int {
	rv := objc.Send[int](e_.ID, objc.Sel("trackingNumber"))
	return rv
}


// The event’s type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/type
func (e_ Event) Type() EventType {
	rv := objc.Send[EventType](e_.ID, objc.Sel("type"))
	return rv
}


// The data associated with a mouse-tracking event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/userData
func (e_ Event) UserData() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](e_.ID, objc.Sel("userData"))
	return rv
}


// An array of three vendor-defined number objects associated with a pointing-type event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/vendorDefined
func (e_ Event) VendorDefined() objc.ID {
	rv := objc.Send[objc.ID](e_.ID, objc.Sel("vendorDefined"))
	return rv
}


// The window object associated with the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/window
func (e_ Event) Window() IWindow {
	rv := objc.Send[Window](e_.ID, objc.Sel("window"))
	return rv
}


// The identifier for the window device associated with the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/windowNumber
func (e_ Event) WindowNumber() int {
	rv := objc.Send[int](e_.ID, objc.Sel("windowNumber"))
	return rv
}


// A mask that indicates the capabilities of the tablet device that generated this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/capabilitymask
func (e_ Event) CapabilityMask() int {
	rv := objc.Send[int](e_.ID, objc.Sel("capabilityMask"))
	return rv
}


// A mask that indicates the capabilities of the tablet device that generated this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/capabilitymask
func (e_ Event) SetCapabilityMask(value int) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setCapabilityMask:"), value)
}


// The display graphics context for this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/context
func (e_ Event) Context() IGraphicsContext {
	rv := objc.Send[GraphicsContext](e_.ID, objc.Sel("context"))
	return rv
}


// The display graphics context for this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/context
func (e_ Event) SetContext(value IGraphicsContext) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setContext:"), value)
}


// A special identifier the system matches against tablet-pointer and tablet-proximity events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/deviceid
func (e_ Event) DeviceID() int {
	rv := objc.Send[int](e_.ID, objc.Sel("deviceID"))
	return rv
}


// A special identifier the system matches against tablet-pointer and tablet-proximity events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/deviceid
func (e_ Event) SetDeviceID(value int) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setDeviceID:"), value)
}


// A Boolean value that indicates whether the key event is a repeat.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/isarepeat
func (e_ Event) IsARepeat() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("isARepeat"))
	return rv
}


// A Boolean value that indicates whether the key event is a repeat.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/isarepeat
func (e_ Event) SetIsARepeat(value bool) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setIsARepeat:"), value)
}


// A Boolean value that indicates whether the user has changed the device inversion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/isdirectioninvertedfromdevice
func (e_ Event) IsDirectionInvertedFromDevice() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("isDirectionInvertedFromDevice"))
	return rv
}


// A Boolean value that indicates whether the user has changed the device inversion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/isdirectioninvertedfromdevice
func (e_ Event) SetIsDirectionInvertedFromDevice(value bool) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setIsDirectionInvertedFromDevice:"), value)
}


// A Boolean value that indicates whether a pointing device is entering or leaving the proximity of its tablet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/isenteringproximity
func (e_ Event) IsEnteringProximity() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("isEnteringProximity"))
	return rv
}


// A Boolean value that indicates whether a pointing device is entering or leaving the proximity of its tablet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/isenteringproximity
func (e_ Event) SetIsEnteringProximity(value bool) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setIsEnteringProximity:"), value)
}


// The index of the pointing device currently in proximity with the tablet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/pointingdeviceid
func (e_ Event) PointingDeviceID() int {
	rv := objc.Send[int](e_.ID, objc.Sel("pointingDeviceID"))
	return rv
}


// The index of the pointing device currently in proximity with the tablet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/pointingdeviceid
func (e_ Event) SetPointingDeviceID(value int) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setPointingDeviceID:"), value)
}


// The vendor-assigned serial number of a pointing device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/pointingdeviceserialnumber
func (e_ Event) PointingDeviceSerialNumber() int {
	rv := objc.Send[int](e_.ID, objc.Sel("pointingDeviceSerialNumber"))
	return rv
}


// The vendor-assigned serial number of a pointing device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/pointingdeviceserialnumber
func (e_ Event) SetPointingDeviceSerialNumber(value int) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setPointingDeviceSerialNumber:"), value)
}


// The kind of pointing device associated with this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/pointingdevicetype-swift.property
func (e_ Event) PointingDeviceType() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](e_.ID, objc.Sel("pointingDeviceType"))
	return rv
}


// The kind of pointing device associated with this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/pointingdevicetype-swift.property
func (e_ Event) SetPointingDeviceType(value objectivec.IObject) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setPointingDeviceType:"), value)
}


// The code associated with a function key or other special key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/specialkey-swift.property
func (e_ Event) SpecialKey() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](e_.ID, objc.Sel("specialKey"))
	return rv
}


// The code associated with a function key or other special key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/specialkey-swift.property
func (e_ Event) SetSpecialKey(value objectivec.IObject) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setSpecialKey:"), value)
}


// The index of the tablet device connected to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/systemtabletid
func (e_ Event) SystemTabletID() int {
	rv := objc.Send[int](e_.ID, objc.Sel("systemTabletID"))
	return rv
}


// The index of the tablet device connected to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/systemtabletid
func (e_ Event) SetSystemTabletID(value int) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setSystemTabletID:"), value)
}


// The USB model identifier of the tablet device associated with this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/tabletid
func (e_ Event) TabletID() int {
	rv := objc.Send[int](e_.ID, objc.Sel("tabletID"))
	return rv
}


// The USB model identifier of the tablet device associated with this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/tabletid
func (e_ Event) SetTabletID(value int) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setTabletID:"), value)
}


// The unique identifier of the pointing device that generated this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/uniqueid
func (e_ Event) UniqueID() uint64 {
	rv := objc.Send[uint64](e_.ID, objc.Sel("uniqueID"))
	return rv
}


// The unique identifier of the pointing device that generated this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/uniqueid
func (e_ Event) SetUniqueID(value uint64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setUniqueID:"), value)
}


// The vendor identifier of the tablet associated with the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/vendorid
func (e_ Event) VendorID() int {
	rv := objc.Send[int](e_.ID, objc.Sel("vendorID"))
	return rv
}


// The vendor identifier of the tablet associated with the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/vendorid
func (e_ Event) SetVendorID(value int) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setVendorID:"), value)
}


// A coded bit field whose set bits indicate the type of pointing device (within a vendor selection) associated with the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/vendorpointingdevicetype
func (e_ Event) VendorPointingDeviceType() int {
	rv := objc.Send[int](e_.ID, objc.Sel("vendorPointingDeviceType"))
	return rv
}


// A coded bit field whose set bits indicate the type of pointing device (within a vendor selection) associated with the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/vendorpointingdevicetype
func (e_ Event) SetVendorPointingDeviceType(value int) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setVendorPointingDeviceType:"), value)
}







