// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	AbsoluteX() int /* primitive/slice/pointer. */
	AbsoluteY() int /* primitive/slice/pointer. */
	AbsoluteZ() int /* primitive/slice/pointer. */
	AssociatedEventsMask() EventMask
	ButtonMask() EventButtonMask
	ButtonNumber() int /* primitive/slice/pointer. */
	CapabilityMask() uint /* primitive/slice/pointer. */
	CGEvent() EventRef /* not a class type */
	Characters() objc.IObject /* cross-framework: NSString */
	CharactersIgnoringModifiers() objc.IObject /* cross-framework: NSString */
	ClickCount() int /* primitive/slice/pointer. */
	Context() IGraphicsContext
	Data1() int /* primitive/slice/pointer. */
	Data2() int /* primitive/slice/pointer. */
	DeltaX() float64 /* primitive/slice/pointer. */
	DeltaY() float64 /* primitive/slice/pointer. */
	DeltaZ() float64 /* primitive/slice/pointer. */
	DeviceID() uint /* primitive/slice/pointer. */
	EventNumber() int /* primitive/slice/pointer. */
	EventRef() unsafe.Pointer
	HasPreciseScrollingDeltas() bool /* primitive/slice/pointer. */
	ARepeat() bool /* primitive/slice/pointer. */
	DirectionInvertedFromDevice() bool /* primitive/slice/pointer. */
	EnteringProximity() bool /* primitive/slice/pointer. */
	KeyCode() unsafe.Pointer
	LocationInWindow() objc.IObject /* cross-framework: Point */
	Magnification() float64 /* primitive/slice/pointer. */
	ModifierFlags() EventModifierFlags
	MomentumPhase() EventPhase
	Phase() EventPhase
	PointingDeviceID() uint /* primitive/slice/pointer. */
	PointingDeviceSerialNumber() uint /* primitive/slice/pointer. */
	PointingDeviceType() PointingDeviceType
	Pressure() float32 /* primitive/slice/pointer. */
	PressureBehavior() PressureBehavior
	Rotation() float32 /* primitive/slice/pointer. */
	ScrollingDeltaX() float64 /* primitive/slice/pointer. */
	ScrollingDeltaY() float64 /* primitive/slice/pointer. */
	Stage() int /* primitive/slice/pointer. */
	StageTransition() float64 /* primitive/slice/pointer. */
	Subtype() EventSubtype
	SystemTabletID() uint /* primitive/slice/pointer. */
	TabletID() uint /* primitive/slice/pointer. */
	TangentialPressure() float32 /* primitive/slice/pointer. */
	Tilt() objc.IObject /* cross-framework: Point */
	Timestamp() TimeInterval /* not a class type */
	TrackingArea() ITrackingArea
	TrackingNumber() int /* primitive/slice/pointer. */
	Type() EventType
	UniqueID() uint64 /* primitive/slice/pointer. */
	UserData() unsafe.Pointer
	VendorDefined() objc.ID
	VendorID() uint /* primitive/slice/pointer. */
	VendorPointingDeviceType() uint /* primitive/slice/pointer. */
	Window() IWindow
	WindowNumber() int /* primitive/slice/pointer. */
	IsARepeat() bool /* primitive/slice/pointer. */
	SetIsARepeat(value bool /* primitive/slice/pointer. */)
	IsDirectionInvertedFromDevice() bool /* primitive/slice/pointer. */
	SetIsDirectionInvertedFromDevice(value bool /* primitive/slice/pointer. */)
	IsEnteringProximity() bool /* primitive/slice/pointer. */
	SetIsEnteringProximity(value bool /* primitive/slice/pointer. */)
	SpecialKey() unsafe.Pointer
	SetSpecialKey(value unsafe.Pointer)
	// methods:
	AllTouches() unsafe.Pointer
	CharactersByApplyingModifiers(modifiers EventModifierFlags) objc.IObject /* cross-framework: String */
	CoalescedTouchesForTouch(touch ITouch) []Touch /* primitive/slice/pointer. */
	LocationInNode(node Node /* not a class type */) objc.IObject /* cross-framework: Point */
	TouchesForView(view IView) unsafe.Pointer
	TouchesMatchingPhaseInView(phase TouchPhase, view IView) unsafe.Pointer
	TrackSwipeEventWithOptionsDampenAmountThresholdMinMaxUsingHandler(options EventSwipeTrackingOptions, minDampenThreshold float64 /* primitive/slice/pointer. */, maxDampenThreshold float64 /* primitive/slice/pointer. */, trackingHandler unsafe.Pointer)
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
func NewEventWithEventRef(eventRef unsafe.Pointer) Event {
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
func (ec _EventClass) AddLocalMonitorForEventsMatchingMaskHandler(mask EventMask, block Event  * (^)( NSEvent  *  event /* not a class type */) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(ec.class), objc.Sel("addLocalMonitorForEventsMatchingMask:handler:"), mask, block)
	return rv
}


// Creates and returns a new event object that describes a tracking-rectangle or cursor-update event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/enterExitEvent(with:location:modifierFlags:timestamp:windowNumber:context:eventNumber:trackingNumber:userData:)
func (ec _EventClass) EnterExitEventWithTypeLocationModifierFlagsTimestampWindowNumberContextEventNumberTrackingNumberUserData(type_ EventType, location objc.IObject /* cross-framework Point */, flags EventModifierFlags, time TimeInterval /* not a class type */, wNum int /* primitive/slice/pointer. */, unusedPassNil IGraphicsContext, eNum int /* primitive/slice/pointer. */, tNum int /* primitive/slice/pointer. */, data unsafe.Pointer) IEvent {
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
func (ec _EventClass) EventWithEventRef(eventRef unsafe.Pointer) IEvent {
	rv := objc.Send[Event](objc.ID(ec.class), objc.Sel("eventWithEventRef:"), eventRef)
	return rv
}


// Creates and returns a new event object that describes a key event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/keyEvent(with:location:modifierFlags:timestamp:windowNumber:context:characters:charactersIgnoringModifiers:isARepeat:keyCode:)
func (ec _EventClass) KeyEventWithTypeLocationModifierFlagsTimestampWindowNumberContextCharactersCharactersIgnoringModifiersIsARepeatKeyCode(type_ EventType, location objc.IObject /* cross-framework Point */, flags EventModifierFlags, time TimeInterval /* not a class type */, wNum int /* primitive/slice/pointer. */, unusedPassNil IGraphicsContext, keys objc.IObject /* cross-framework NSString */, ukeys objc.IObject /* cross-framework NSString */, flag bool /* primitive/slice/pointer. */, code unsafe.Pointer) IEvent {
	rv := objc.Send[Event](objc.ID(ec.class), objc.Sel("keyEventWithType:location:modifierFlags:timestamp:windowNumber:context:characters:charactersIgnoringModifiers:isARepeat:keyCode:"), type_, location, flags, time, wNum, unusedPassNil, keys, ukeys, flag, code)
	return rv
}


// Creates and returns a new event object that describes a mouse-down, -up, -moved, or -dragged event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/mouseEvent(with:location:modifierFlags:timestamp:windowNumber:context:eventNumber:clickCount:pressure:)
func (ec _EventClass) MouseEventWithTypeLocationModifierFlagsTimestampWindowNumberContextEventNumberClickCountPressure(type_ EventType, location objc.IObject /* cross-framework Point */, flags EventModifierFlags, time TimeInterval /* not a class type */, wNum int /* primitive/slice/pointer. */, unusedPassNil IGraphicsContext, eNum int /* primitive/slice/pointer. */, cNum int /* primitive/slice/pointer. */, pressure float32 /* primitive/slice/pointer. */) IEvent {
	rv := objc.Send[Event](objc.ID(ec.class), objc.Sel("mouseEventWithType:location:modifierFlags:timestamp:windowNumber:context:eventNumber:clickCount:pressure:"), type_, location, flags, time, wNum, unusedPassNil, eNum, cNum, pressure)
	return rv
}


// Creates and returns a new event object that describes a custom event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/otherEvent(with:location:modifierFlags:timestamp:windowNumber:context:subtype:data1:data2:)
func (ec _EventClass) OtherEventWithTypeLocationModifierFlagsTimestampWindowNumberContextSubtypeData1Data2(type_ EventType, location objc.IObject /* cross-framework Point */, flags EventModifierFlags, time TimeInterval /* not a class type */, wNum int /* primitive/slice/pointer. */, unusedPassNil IGraphicsContext, subtype unsafe.Pointer, d1 int /* primitive/slice/pointer. */, d2 int /* primitive/slice/pointer. */) IEvent {
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


// Begins generating periodic events for the current thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/startPeriodicEvents(afterDelay:withPeriod:)
func (ec _EventClass) StartPeriodicEventsAfterDelayWithPeriod(delay TimeInterval /* not a class type */, period TimeInterval /* not a class type */) {
	objc.Send[objc.ID](objc.ID(ec.class), objc.Sel("startPeriodicEventsAfterDelay:withPeriod:"), delay, period)
}


// Stops generating periodic events for the current thread and discards any periodic events remaining in the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/stopPeriodicEvents()
func (ec _EventClass) StopPeriodicEvents() {
	objc.Send[objc.ID](objc.ID(ec.class), objc.Sel("stopPeriodicEvents"))
}


// The maximum number of seconds in which a second mouse click must occur for an event to be a double-click event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/doubleClickInterval
func (ec _EventClass) DoubleClickInterval() TimeInterval /* not a class type */ {
	rv := objc.Send[TimeInterval](objc.ID(ec.class), objc.Sel("doubleClickInterval"))
	return rv
}

// A Boolean value that indicates whether the system coalesces mouse movement events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/isMouseCoalescingEnabled
func (ec _EventClass) MouseCoalescingEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](objc.ID(ec.class), objc.Sel("mouseCoalescingEnabled"))
	return rv
}

// A Boolean value that indicates whether to track fluid swipe gestures using scroll events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/isSwipeTrackingFromScrollEventsEnabled
func (ec _EventClass) SwipeTrackingFromScrollEventsEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](objc.ID(ec.class), objc.Sel("swipeTrackingFromScrollEventsEnabled"))
	return rv
}

// The number of seconds someone must hold down a key before the first key repeat event occurs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/keyRepeatDelay
func (ec _EventClass) KeyRepeatDelay() TimeInterval /* not a class type */ {
	rv := objc.Send[TimeInterval](objc.ID(ec.class), objc.Sel("keyRepeatDelay"))
	return rv
}

// The number of seconds someone must hold down a key to generate key-repeat events after the initial delay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/keyRepeatInterval
func (ec _EventClass) KeyRepeatInterval() TimeInterval /* not a class type */ {
	rv := objc.Send[TimeInterval](objc.ID(ec.class), objc.Sel("keyRepeatInterval"))
	return rv
}

// Reports the current mouse position in screen coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/mouseLocation
func (ec _EventClass) MouseLocation() objc.IObject /* cross-framework: Point */ {
	rv := objc.Send[Point](objc.ID(ec.class), objc.Sel("mouseLocation"))
	return rv
}

// The indices of the currently pressed mouse buttons.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/pressedMouseButtons
func (ec _EventClass) PressedMouseButtons() uint /* primitive/slice/pointer. */ {
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
func (e_ Event) CharactersByApplyingModifiers(modifiers EventModifierFlags) objc.IObject /* cross-framework: String */ {
	rv := objc.Send[String](e_.ID, objc.Sel("charactersByApplyingModifiers:"), modifiers)
	return rv
}


// Returns all of the touch objects associated with the specified main touch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/coalescedTouches(for:)
func (e_ Event) CoalescedTouchesForTouch(touch ITouch) []Touch /* primitive/slice/pointer. */ {
	rv := objc.Send[[]Touch](e_.ID, objc.Sel("coalescedTouchesForTouch:"), touch)
	return rv
}


// Returns the location of the receiver in the coordinate system of the given node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/location(in:)
func (e_ Event) LocationInNode(node Node /* not a class type */) objc.IObject /* cross-framework: Point */ {
	rv := objc.Send[Point](e_.ID, objc.Sel("locationInNode:"), node)
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


// Allows tracking and user interface feedback of scroll wheel events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/trackSwipeEvent(options:dampenAmountThresholdMin:max:usingHandler:)
func (e_ Event) TrackSwipeEventWithOptionsDampenAmountThresholdMinMaxUsingHandler(options EventSwipeTrackingOptions, minDampenThreshold float64 /* primitive/slice/pointer. */, maxDampenThreshold float64 /* primitive/slice/pointer. */, trackingHandler unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("trackSwipeEventWithOptions:dampenAmountThresholdMin:max:usingHandler:"), options, minDampenThreshold, maxDampenThreshold, trackingHandler)
}


// The absolute x coordinate of a pointing device on its tablet at full tablet resolution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/absoluteX
func (e_ Event) AbsoluteX() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](e_.ID, objc.Sel("absoluteX"))
	return rv
}


// The absolute y coordinate of a pointing device on its tablet at full tablet resolution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/absoluteY
func (e_ Event) AbsoluteY() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](e_.ID, objc.Sel("absoluteY"))
	return rv
}


// The absolute z coordinate of pointing device on its tablet at full tablet resolution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/absoluteZ
func (e_ Event) AbsoluteZ() int /* primitive/slice/pointer. */ {
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
func (e_ Event) ButtonNumber() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](e_.ID, objc.Sel("buttonNumber"))
	return rv
}


// A mask that indicates the capabilities of the tablet device that generated this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/capabilityMask
func (e_ Event) CapabilityMask() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](e_.ID, objc.Sel("capabilityMask"))
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
func (e_ Event) Characters() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("characters"))
	return rv
}


// The characters generated by a key event as if no modifier key (except for Shift) applies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/charactersIgnoringModifiers
func (e_ Event) CharactersIgnoringModifiers() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("charactersIgnoringModifiers"))
	return rv
}


// The number of mouse clicks associated with a mouse-down or mouse-up event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/clickCount
func (e_ Event) ClickCount() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](e_.ID, objc.Sel("clickCount"))
	return rv
}


// The display graphics context for this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/context
func (e_ Event) Context() IGraphicsContext {
	rv := objc.Send[GraphicsContext](e_.ID, objc.Sel("context"))
	return rv
}


// Additional data associated with this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/data1
func (e_ Event) Data1() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](e_.ID, objc.Sel("data1"))
	return rv
}


// Additional data associated with this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/data2
func (e_ Event) Data2() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](e_.ID, objc.Sel("data2"))
	return rv
}


// The x-coordinate change for scroll wheel, mouse-move, mouse-drag, and swipe events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/deltaX
func (e_ Event) DeltaX() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](e_.ID, objc.Sel("deltaX"))
	return rv
}


// The y-coordinate change for scroll wheel, mouse-move, mouse-drag, and swipe events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/deltaY
func (e_ Event) DeltaY() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](e_.ID, objc.Sel("deltaY"))
	return rv
}


// The z-coordinate change for a scroll wheel, mouse-move, or mouse-drag event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/deltaZ
func (e_ Event) DeltaZ() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](e_.ID, objc.Sel("deltaZ"))
	return rv
}


// A special identifier the system matches against tablet-pointer and tablet-proximity events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/deviceID
func (e_ Event) DeviceID() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](e_.ID, objc.Sel("deviceID"))
	return rv
}


// The maximum number of seconds in which a second mouse click must occur for an event to be a double-click event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/doubleClickInterval
func (e_ Event) DoubleClickInterval() TimeInterval /* not a class type */ {
	rv := objc.Send[TimeInterval](e_.ID, objc.Sel("doubleClickInterval"))
	return rv
}


// The counter value of the latest mouse or tracking-rectangle event object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/eventNumber
func (e_ Event) EventNumber() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](e_.ID, objc.Sel("eventNumber"))
	return rv
}


// An opaque Carbon type associated with this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/eventRef
func (e_ Event) EventRef() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("eventRef"))
	return rv
}


// A Boolean value that indicates whether precise scrolling deltas are available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/hasPreciseScrollingDeltas
func (e_ Event) HasPreciseScrollingDeltas() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](e_.ID, objc.Sel("hasPreciseScrollingDeltas"))
	return rv
}


// A Boolean value that indicates whether the key event is a repeat.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/isARepeat
func (e_ Event) ARepeat() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](e_.ID, objc.Sel("ARepeat"))
	return rv
}


// A Boolean value that indicates whether the user has changed the device inversion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/isDirectionInvertedFromDevice
func (e_ Event) DirectionInvertedFromDevice() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](e_.ID, objc.Sel("directionInvertedFromDevice"))
	return rv
}


// A Boolean value that indicates whether a pointing device is entering or leaving the proximity of its tablet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/isEnteringProximity
func (e_ Event) EnteringProximity() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](e_.ID, objc.Sel("enteringProximity"))
	return rv
}


// A Boolean value that indicates whether the system coalesces mouse movement events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/isMouseCoalescingEnabled
func (e_ Event) MouseCoalescingEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](e_.ID, objc.Sel("mouseCoalescingEnabled"))
	return rv
}


// A Boolean value that indicates whether the system coalesces mouse movement events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/isMouseCoalescingEnabled
func (e_ Event) SetMouseCoalescingEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setMouseCoalescingEnabled:"), value)
}


// A Boolean value that indicates whether to track fluid swipe gestures using scroll events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/isSwipeTrackingFromScrollEventsEnabled
func (e_ Event) SwipeTrackingFromScrollEventsEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](e_.ID, objc.Sel("swipeTrackingFromScrollEventsEnabled"))
	return rv
}


// The virtual code for the key associated with the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/keyCode
func (e_ Event) KeyCode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("keyCode"))
	return rv
}


// The number of seconds someone must hold down a key before the first key repeat event occurs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/keyRepeatDelay
func (e_ Event) KeyRepeatDelay() TimeInterval /* not a class type */ {
	rv := objc.Send[TimeInterval](e_.ID, objc.Sel("keyRepeatDelay"))
	return rv
}


// The number of seconds someone must hold down a key to generate key-repeat events after the initial delay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/keyRepeatInterval
func (e_ Event) KeyRepeatInterval() TimeInterval /* not a class type */ {
	rv := objc.Send[TimeInterval](e_.ID, objc.Sel("keyRepeatInterval"))
	return rv
}


// The event location in the base coordinate system of the associated window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/locationInWindow
func (e_ Event) LocationInWindow() objc.IObject /* cross-framework: Point */ {
	rv := objc.Send[Point](e_.ID, objc.Sel("locationInWindow"))
	return rv
}


// The amount of change to add to a magnification gesture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/magnification
func (e_ Event) Magnification() float64 /* primitive/slice/pointer. */ {
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
func (e_ Event) MouseLocation() objc.IObject /* cross-framework: Point */ {
	rv := objc.Send[Point](e_.ID, objc.Sel("mouseLocation"))
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


// The index of the pointing device currently in proximity with the tablet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/pointingDeviceID
func (e_ Event) PointingDeviceID() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](e_.ID, objc.Sel("pointingDeviceID"))
	return rv
}


// The vendor-assigned serial number of a pointing device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/pointingDeviceSerialNumber
func (e_ Event) PointingDeviceSerialNumber() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](e_.ID, objc.Sel("pointingDeviceSerialNumber"))
	return rv
}


// The kind of pointing device associated with this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/pointingDeviceType-swift.property
func (e_ Event) PointingDeviceType() PointingDeviceType {
	rv := objc.Send[PointingDeviceType](e_.ID, objc.Sel("pointingDeviceType"))
	return rv
}


// The indices of the currently pressed mouse buttons.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/pressedMouseButtons
func (e_ Event) PressedMouseButtons() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](e_.ID, objc.Sel("pressedMouseButtons"))
	return rv
}


// A normalized value that indicates the degree of pressure applied to an appropriate input device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/pressure
func (e_ Event) Pressure() float32 /* primitive/slice/pointer. */ {
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
func (e_ Event) Rotation() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](e_.ID, objc.Sel("rotation"))
	return rv
}


// The scroll wheel’s horizontal delta.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/scrollingDeltaX
func (e_ Event) ScrollingDeltaX() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](e_.ID, objc.Sel("scrollingDeltaX"))
	return rv
}


// The scroll wheel’s vertical delta.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/scrollingDeltaY
func (e_ Event) ScrollingDeltaY() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](e_.ID, objc.Sel("scrollingDeltaY"))
	return rv
}


// A value that indicates the stage of a pressure gesture event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/stage
func (e_ Event) Stage() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](e_.ID, objc.Sel("stage"))
	return rv
}


// The transition value for the stage of a pressure gesture event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/stageTransition
func (e_ Event) StageTransition() float64 /* primitive/slice/pointer. */ {
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


// The index of the tablet device connected to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/systemTabletID
func (e_ Event) SystemTabletID() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](e_.ID, objc.Sel("systemTabletID"))
	return rv
}


// The USB model identifier of the tablet device associated with this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/tabletID
func (e_ Event) TabletID() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](e_.ID, objc.Sel("tabletID"))
	return rv
}


// The tangential pressure on the device that generated this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/tangentialPressure
func (e_ Event) TangentialPressure() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](e_.ID, objc.Sel("tangentialPressure"))
	return rv
}


// The scaled tilt values of the pointing device that generated this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/tilt
func (e_ Event) Tilt() objc.IObject /* cross-framework: Point */ {
	rv := objc.Send[Point](e_.ID, objc.Sel("tilt"))
	return rv
}


// The time when the event occurred in seconds since system startup.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/timestamp
func (e_ Event) Timestamp() TimeInterval /* not a class type */ {
	rv := objc.Send[TimeInterval](e_.ID, objc.Sel("timestamp"))
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
func (e_ Event) TrackingNumber() int /* primitive/slice/pointer. */ {
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


// The unique identifier of the pointing device that generated this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/uniqueID
func (e_ Event) UniqueID() uint64 /* primitive/slice/pointer. */ {
	rv := objc.Send[uint64](e_.ID, objc.Sel("uniqueID"))
	return rv
}


// The data associated with a mouse-tracking event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/userData
func (e_ Event) UserData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("userData"))
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


// The vendor identifier of the tablet associated with the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/vendorID
func (e_ Event) VendorID() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](e_.ID, objc.Sel("vendorID"))
	return rv
}


// A coded bit field whose set bits indicate the type of pointing device (within a vendor selection) associated with the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/vendorPointingDeviceType
func (e_ Event) VendorPointingDeviceType() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](e_.ID, objc.Sel("vendorPointingDeviceType"))
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
func (e_ Event) WindowNumber() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](e_.ID, objc.Sel("windowNumber"))
	return rv
}


// A Boolean value that indicates whether the key event is a repeat.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/isarepeat
func (e_ Event) IsARepeat() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](e_.ID, objc.Sel("isARepeat"))
	return rv
}


// A Boolean value that indicates whether the key event is a repeat.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/isarepeat
func (e_ Event) SetIsARepeat(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setIsARepeat:"), value)
}


// A Boolean value that indicates whether the user has changed the device inversion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/isdirectioninvertedfromdevice
func (e_ Event) IsDirectionInvertedFromDevice() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](e_.ID, objc.Sel("isDirectionInvertedFromDevice"))
	return rv
}


// A Boolean value that indicates whether the user has changed the device inversion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/isdirectioninvertedfromdevice
func (e_ Event) SetIsDirectionInvertedFromDevice(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setIsDirectionInvertedFromDevice:"), value)
}


// A Boolean value that indicates whether a pointing device is entering or leaving the proximity of its tablet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/isenteringproximity
func (e_ Event) IsEnteringProximity() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](e_.ID, objc.Sel("isEnteringProximity"))
	return rv
}


// A Boolean value that indicates whether a pointing device is entering or leaving the proximity of its tablet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/isenteringproximity
func (e_ Event) SetIsEnteringProximity(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setIsEnteringProximity:"), value)
}


// The code associated with a function key or other special key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/specialkey-swift.property
func (e_ Event) SpecialKey() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("specialKey"))
	return rv
}


// The code associated with a function key or other special key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/specialkey-swift.property
func (e_ Event) SetSpecialKey(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setSpecialKey:"), value)
}


