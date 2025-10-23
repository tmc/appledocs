// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
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
	LocationInNode(node unsafe.Pointer) coregraphics.CGPoint
	CGEvent() coregraphics.CGEventRef
	EventRef() unsafe.Pointer
	LocationInWindow() coregraphics.CGPoint
	ModifierFlags() EventModifierFlags
	Timestamp() float64
	Window() NSWindow
	WindowNumber() int
	AbsoluteX() int
	SetAbsoluteX(value int)
	AbsoluteY() int
	SetAbsoluteY(value int)
	AbsoluteZ() int
	SetAbsoluteZ(value int)
	AssociatedEventsMask() unsafe.Pointer
	SetAssociatedEventsMask(value unsafe.Pointer)
	ButtonMask() unsafe.Pointer
	SetButtonMask(value unsafe.Pointer)
	ButtonNumber() int
	SetButtonNumber(value int)
	CapabilityMask() int
	SetCapabilityMask(value int)
	Characters() string
	SetCharacters(value string)
	CharactersIgnoringModifiers() string
	SetCharactersIgnoringModifiers(value string)
	ClickCount() int
	SetClickCount(value int)
	Context() NSGraphicsContext
	SetContext(value IGraphicsContext)
	Data1() int
	SetData1(value int)
	Data2() int
	SetData2(value int)
	DeltaX() float64
	SetDeltaX(value float64)
	DeltaY() float64
	SetDeltaY(value float64)
	DeltaZ() float64
	SetDeltaZ(value float64)
	DeviceID() int
	SetDeviceID(value int)
	EventNumber() int
	SetEventNumber(value int)
	HasPreciseScrollingDeltas() bool
	SetHasPreciseScrollingDeltas(value bool)
	IsARepeat() bool
	SetIsARepeat(value bool)
	IsDirectionInvertedFromDevice() bool
	SetIsDirectionInvertedFromDevice(value bool)
	IsEnteringProximity() bool
	SetIsEnteringProximity(value bool)
	KeyCode() unsafe.Pointer
	SetKeyCode(value unsafe.Pointer)
	Magnification() float64
	SetMagnification(value float64)
	MomentumPhase() unsafe.Pointer
	SetMomentumPhase(value unsafe.Pointer)
	Phase() unsafe.Pointer
	SetPhase(value unsafe.Pointer)
	PointingDeviceID() int
	SetPointingDeviceID(value int)
	PointingDeviceSerialNumber() int
	SetPointingDeviceSerialNumber(value int)
	PointingDeviceType() unsafe.Pointer
	SetPointingDeviceType(value unsafe.Pointer)
	Pressure() float32
	SetPressure(value float32)
	PressureBehavior() unsafe.Pointer
	SetPressureBehavior(value unsafe.Pointer)
	Rotation() float32
	SetRotation(value float32)
	ScrollingDeltaX() float64
	SetScrollingDeltaX(value float64)
	ScrollingDeltaY() float64
	SetScrollingDeltaY(value float64)
	SpecialKey() unsafe.Pointer
	SetSpecialKey(value unsafe.Pointer)
	Stage() int
	SetStage(value int)
	StageTransition() float64
	SetStageTransition(value float64)
	Subtype() unsafe.Pointer
	SetSubtype(value unsafe.Pointer)
	SystemTabletID() int
	SetSystemTabletID(value int)
	TabletID() int
	SetTabletID(value int)
	TangentialPressure() float32
	SetTangentialPressure(value float32)
	Tilt() coregraphics.CGPoint
	SetTilt(value coregraphics.CGPoint)
	TrackingArea() NSTrackingArea
	SetTrackingArea(value ITrackingArea)
	TrackingNumber() int
	SetTrackingNumber(value int)
	Type() EventType
	SetType(value EventType)
	UniqueID() uint64
	SetUniqueID(value uint64)
	UserData() unsafe.Pointer
	SetUserData(value unsafe.Pointer)
	VendorDefined() unsafe.Pointer
	SetVendorDefined(value unsafe.Pointer)
	VendorID() int
	SetVendorID(value int)
	VendorPointingDeviceType() int
	SetVendorPointingDeviceType(value int)
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



// Installs an event monitor that receives copies of events the system posts to this app prior to their dispatch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/addLocalMonitorForEvents(matching:handler:)
func (ec _EventClass) AddLocalMonitorForEventsMatchingMaskHandler(mask EventMask, block unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(ec.class), objc.Sel("addLocalMonitorForEventsMatchingMask:handler:"), mask, block)
	return rv
}


// Reports the current mouse position in screen coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/mouseLocation
func (ec _EventClass) MouseLocation() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](objc.ID(ec.class), objc.Sel("mouseLocation"))
	return rv
}

// Returns the location of the receiver in the coordinate system of the given node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/location(in:)
func (e_ Event) LocationInNode(node unsafe.Pointer) coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](e_.ID, objc.Sel("locationInNode:"), node)
	return rv
}


// The Core Graphics event object corresponding to this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/cgEvent
func (e_ Event) CGEvent() coregraphics.CGEventRef {
	rv := objc.Send[coregraphics.CGEventRef](e_.ID, objc.Sel("CGEvent"))
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


// The event location in the base coordinate system of the associated window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/locationInWindow
func (e_ Event) LocationInWindow() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](e_.ID, objc.Sel("locationInWindow"))
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


// Reports the current mouse position in screen coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/mouseLocation
func (e_ Event) MouseLocation() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](e_.ID, objc.Sel("mouseLocation"))
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


// The window object associated with the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/window
func (e_ Event) Window() NSWindow {
	rv := objc.Send[NSWindow](e_.ID, objc.Sel("window"))
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


// The absolute x coordinate of a pointing device on its tablet at full tablet resolution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/absolutex
func (e_ Event) AbsoluteX() int {
	rv := objc.Send[int](e_.ID, objc.Sel("absoluteX"))
	return rv
}


// The absolute x coordinate of a pointing device on its tablet at full tablet resolution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/absolutex
func (e_ Event) SetAbsoluteX(value int) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setAbsoluteX:"), value)
}


// The absolute y coordinate of a pointing device on its tablet at full tablet resolution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/absolutey
func (e_ Event) AbsoluteY() int {
	rv := objc.Send[int](e_.ID, objc.Sel("absoluteY"))
	return rv
}


// The absolute y coordinate of a pointing device on its tablet at full tablet resolution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/absolutey
func (e_ Event) SetAbsoluteY(value int) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setAbsoluteY:"), value)
}


// The absolute z coordinate of pointing device on its tablet at full tablet resolution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/absolutez
func (e_ Event) AbsoluteZ() int {
	rv := objc.Send[int](e_.ID, objc.Sel("absoluteZ"))
	return rv
}


// The absolute z coordinate of pointing device on its tablet at full tablet resolution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/absolutez
func (e_ Event) SetAbsoluteZ(value int) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setAbsoluteZ:"), value)
}


// The associated events mask of a mouse event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/associatedeventsmask
func (e_ Event) AssociatedEventsMask() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("associatedEventsMask"))
	return rv
}


// The associated events mask of a mouse event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/associatedeventsmask
func (e_ Event) SetAssociatedEventsMask(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setAssociatedEventsMask:"), value)
}


// A bit mask identifying the buttons pressed for a tablet event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/buttonmask-swift.property
func (e_ Event) ButtonMask() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("buttonMask"))
	return rv
}


// A bit mask identifying the buttons pressed for a tablet event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/buttonmask-swift.property
func (e_ Event) SetButtonMask(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setButtonMask:"), value)
}


// The button number for a mouse event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/buttonnumber
func (e_ Event) ButtonNumber() int {
	rv := objc.Send[int](e_.ID, objc.Sel("buttonNumber"))
	return rv
}


// The button number for a mouse event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/buttonnumber
func (e_ Event) SetButtonNumber(value int) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setButtonNumber:"), value)
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


// The characters associated with a key-up or key-down event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/characters
func (e_ Event) Characters() string {
	rv := objc.Send[string](e_.ID, objc.Sel("characters"))
	return rv
}


// The characters associated with a key-up or key-down event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/characters
func (e_ Event) SetCharacters(value string) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setCharacters:"), objc.String(value))
}


// The characters generated by a key event as if no modifier key (except for Shift) applies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/charactersignoringmodifiers
func (e_ Event) CharactersIgnoringModifiers() string {
	rv := objc.Send[string](e_.ID, objc.Sel("charactersIgnoringModifiers"))
	return rv
}


// The characters generated by a key event as if no modifier key (except for Shift) applies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/charactersignoringmodifiers
func (e_ Event) SetCharactersIgnoringModifiers(value string) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setCharactersIgnoringModifiers:"), objc.String(value))
}


// The number of mouse clicks associated with a mouse-down or mouse-up event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/clickcount
func (e_ Event) ClickCount() int {
	rv := objc.Send[int](e_.ID, objc.Sel("clickCount"))
	return rv
}


// The number of mouse clicks associated with a mouse-down or mouse-up event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/clickcount
func (e_ Event) SetClickCount(value int) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setClickCount:"), value)
}


// The display graphics context for this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/context
func (e_ Event) Context() NSGraphicsContext {
	rv := objc.Send[NSGraphicsContext](e_.ID, objc.Sel("context"))
	return rv
}


// The display graphics context for this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/context
func (e_ Event) SetContext(value IGraphicsContext) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setContext:"), value)
}


// Additional data associated with this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/data1
func (e_ Event) Data1() int {
	rv := objc.Send[int](e_.ID, objc.Sel("data1"))
	return rv
}


// Additional data associated with this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/data1
func (e_ Event) SetData1(value int) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setData1:"), value)
}


// Additional data associated with this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/data2
func (e_ Event) Data2() int {
	rv := objc.Send[int](e_.ID, objc.Sel("data2"))
	return rv
}


// Additional data associated with this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/data2
func (e_ Event) SetData2(value int) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setData2:"), value)
}


// The x-coordinate change for scroll wheel, mouse-move, mouse-drag, and swipe events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/deltax
func (e_ Event) DeltaX() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("deltaX"))
	return rv
}


// The x-coordinate change for scroll wheel, mouse-move, mouse-drag, and swipe events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/deltax
func (e_ Event) SetDeltaX(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setDeltaX:"), value)
}


// The y-coordinate change for scroll wheel, mouse-move, mouse-drag, and swipe events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/deltay
func (e_ Event) DeltaY() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("deltaY"))
	return rv
}


// The y-coordinate change for scroll wheel, mouse-move, mouse-drag, and swipe events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/deltay
func (e_ Event) SetDeltaY(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setDeltaY:"), value)
}


// The z-coordinate change for a scroll wheel, mouse-move, or mouse-drag event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/deltaz
func (e_ Event) DeltaZ() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("deltaZ"))
	return rv
}


// The z-coordinate change for a scroll wheel, mouse-move, or mouse-drag event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/deltaz
func (e_ Event) SetDeltaZ(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setDeltaZ:"), value)
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


// The counter value of the latest mouse or tracking-rectangle event object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/eventnumber
func (e_ Event) EventNumber() int {
	rv := objc.Send[int](e_.ID, objc.Sel("eventNumber"))
	return rv
}


// The counter value of the latest mouse or tracking-rectangle event object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/eventnumber
func (e_ Event) SetEventNumber(value int) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setEventNumber:"), value)
}


// A Boolean value that indicates whether precise scrolling deltas are available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/hasprecisescrollingdeltas
func (e_ Event) HasPreciseScrollingDeltas() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("hasPreciseScrollingDeltas"))
	return rv
}


// A Boolean value that indicates whether precise scrolling deltas are available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/hasprecisescrollingdeltas
func (e_ Event) SetHasPreciseScrollingDeltas(value bool) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setHasPreciseScrollingDeltas:"), value)
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


// The virtual code for the key associated with the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/keycode
func (e_ Event) KeyCode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("keyCode"))
	return rv
}


// The virtual code for the key associated with the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/keycode
func (e_ Event) SetKeyCode(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setKeyCode:"), value)
}


// The amount of change to add to a magnification gesture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/magnification
func (e_ Event) Magnification() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("magnification"))
	return rv
}


// The amount of change to add to a magnification gesture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/magnification
func (e_ Event) SetMagnification(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setMagnification:"), value)
}


// The momentum phase for a scroll or flick gesture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/momentumphase
func (e_ Event) MomentumPhase() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("momentumPhase"))
	return rv
}


// The momentum phase for a scroll or flick gesture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/momentumphase
func (e_ Event) SetMomentumPhase(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setMomentumPhase:"), value)
}


// The phase of a gesture event, such as a magnify, scroll, or pressure change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/phase-swift.property
func (e_ Event) Phase() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("phase"))
	return rv
}


// The phase of a gesture event, such as a magnify, scroll, or pressure change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/phase-swift.property
func (e_ Event) SetPhase(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setPhase:"), value)
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
func (e_ Event) PointingDeviceType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("pointingDeviceType"))
	return rv
}


// The kind of pointing device associated with this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/pointingdevicetype-swift.property
func (e_ Event) SetPointingDeviceType(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setPointingDeviceType:"), value)
}


// A normalized value that indicates the degree of pressure applied to an appropriate input device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/pressure
func (e_ Event) Pressure() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("pressure"))
	return rv
}


// A normalized value that indicates the degree of pressure applied to an appropriate input device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/pressure
func (e_ Event) SetPressure(value float32) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setPressure:"), value)
}


// The behavior and progression for a pressure event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/pressurebehavior-swift.property
func (e_ Event) PressureBehavior() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("pressureBehavior"))
	return rv
}


// The behavior and progression for a pressure event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/pressurebehavior-swift.property
func (e_ Event) SetPressureBehavior(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setPressureBehavior:"), value)
}


// The rotation in degrees of the tablet pointing device associated with this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/rotation
func (e_ Event) Rotation() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("rotation"))
	return rv
}


// The rotation in degrees of the tablet pointing device associated with this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/rotation
func (e_ Event) SetRotation(value float32) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setRotation:"), value)
}


// The scroll wheel’s horizontal delta.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/scrollingdeltax
func (e_ Event) ScrollingDeltaX() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("scrollingDeltaX"))
	return rv
}


// The scroll wheel’s horizontal delta.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/scrollingdeltax
func (e_ Event) SetScrollingDeltaX(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setScrollingDeltaX:"), value)
}


// The scroll wheel’s vertical delta.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/scrollingdeltay
func (e_ Event) ScrollingDeltaY() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("scrollingDeltaY"))
	return rv
}


// The scroll wheel’s vertical delta.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/scrollingdeltay
func (e_ Event) SetScrollingDeltaY(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setScrollingDeltaY:"), value)
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


// A value that indicates the stage of a pressure gesture event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/stage
func (e_ Event) Stage() int {
	rv := objc.Send[int](e_.ID, objc.Sel("stage"))
	return rv
}


// A value that indicates the stage of a pressure gesture event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/stage
func (e_ Event) SetStage(value int) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setStage:"), value)
}


// The transition value for the stage of a pressure gesture event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/stagetransition
func (e_ Event) StageTransition() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("stageTransition"))
	return rv
}


// The transition value for the stage of a pressure gesture event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/stagetransition
func (e_ Event) SetStageTransition(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setStageTransition:"), value)
}


// The event’s subtype.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/subtype
func (e_ Event) Subtype() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("subtype"))
	return rv
}


// The event’s subtype.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/subtype
func (e_ Event) SetSubtype(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setSubtype:"), value)
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


// The tangential pressure on the device that generated this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/tangentialpressure
func (e_ Event) TangentialPressure() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("tangentialPressure"))
	return rv
}


// The tangential pressure on the device that generated this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/tangentialpressure
func (e_ Event) SetTangentialPressure(value float32) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setTangentialPressure:"), value)
}


// The scaled tilt values of the pointing device that generated this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/tilt
func (e_ Event) Tilt() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](e_.ID, objc.Sel("tilt"))
	return rv
}


// The scaled tilt values of the pointing device that generated this event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/tilt
func (e_ Event) SetTilt(value coregraphics.CGPoint) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setTilt:"), value)
}


// The tracking area for the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/trackingarea
func (e_ Event) TrackingArea() NSTrackingArea {
	rv := objc.Send[NSTrackingArea](e_.ID, objc.Sel("trackingArea"))
	return rv
}


// The tracking area for the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/trackingarea
func (e_ Event) SetTrackingArea(value ITrackingArea) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setTrackingArea:"), value)
}


// The identifier of a mouse-tracking event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/trackingnumber
func (e_ Event) TrackingNumber() int {
	rv := objc.Send[int](e_.ID, objc.Sel("trackingNumber"))
	return rv
}


// The identifier of a mouse-tracking event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/trackingnumber
func (e_ Event) SetTrackingNumber(value int) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setTrackingNumber:"), value)
}


// The event’s type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/type
func (e_ Event) Type() EventType {
	rv := objc.Send[EventType](e_.ID, objc.Sel("type"))
	return rv
}


// The event’s type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/type
func (e_ Event) SetType(value EventType) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setType:"), value)
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


// The data associated with a mouse-tracking event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/userdata
func (e_ Event) UserData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("userData"))
	return rv
}


// The data associated with a mouse-tracking event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/userdata
func (e_ Event) SetUserData(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setUserData:"), value)
}


// An array of three vendor-defined number objects associated with a pointing-type event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/vendordefined
func (e_ Event) VendorDefined() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("vendorDefined"))
	return rv
}


// An array of three vendor-defined number objects associated with a pointing-type event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/vendordefined
func (e_ Event) SetVendorDefined(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setVendorDefined:"), value)
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



