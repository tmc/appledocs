
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Responder] class.
var ResponderClass _ResponderClass

func init() {
	ResponderClass = _ResponderClass{objc.GetClass("NSResponder")}
}

type _ResponderClass struct {
	objc.Class
}

// An interface definition for the [Responder] class.
type IResponder interface {
	ID() objc.ID
	BecomeFirstResponder() bool
	BeginGestureWithEvent(event unsafe.Pointer)
	ChangeModeWithEvent(event unsafe.Pointer)
	ContextMenuKeyDown(event unsafe.Pointer)
	CursorUpdate(event unsafe.Pointer)
	EncodeRestorableStateWithCoder(coder unsafe.Pointer)
	EncodeRestorableStateWithCoderBackgroundQueue(coder unsafe.Pointer, queue unsafe.Pointer)
	EndGestureWithEvent(event unsafe.Pointer)
	FlagsChanged(event unsafe.Pointer)
	FlushBufferedKeyEvents()
	HelpRequested(eventPtr unsafe.Pointer)
	InterfaceStyle() unsafe.Pointer
	InterpretKeyEvents(eventArray unsafe.Pointer)
	InvalidateRestorableState()
	KeyDown(event unsafe.Pointer)
	KeyUp(event unsafe.Pointer)
	MagnifyWithEvent(event unsafe.Pointer)
	MakeTouchBar() unsafe.Pointer
	MouseCancelled(event unsafe.Pointer)
	MouseDown(event unsafe.Pointer)
	MouseDragged(event unsafe.Pointer)
	MouseEntered(event unsafe.Pointer)
	MouseExited(event unsafe.Pointer)
	MouseMoved(event unsafe.Pointer)
	MouseUp(event unsafe.Pointer)
	NewWindowForTab(sender objc.ID)
	NoResponderFor(eventSelector objc.SEL)
	OtherMouseDown(event unsafe.Pointer)
	OtherMouseDragged(event unsafe.Pointer)
	OtherMouseUp(event unsafe.Pointer)
	PerformKeyEquivalent(event unsafe.Pointer) bool
	PerformMnemonic(string string) bool
	PerformTextFinderAction(sender objc.ID)
	PresentError(error unsafe.Pointer) bool
	PresentErrorModalForWindowDelegateDidPresentSelectorContextInfo(error unsafe.Pointer, window unsafe.Pointer, delegate objc.ID, didPresentSelector objc.SEL, contextInfo unsafe.Pointer)
	PressureChangeWithEvent(event unsafe.Pointer)
	QuickLookWithEvent(event unsafe.Pointer)
	ResignFirstResponder() bool
	RestoreStateWithCoder(coder unsafe.Pointer)
	RightMouseDown(event unsafe.Pointer)
	RightMouseDragged(event unsafe.Pointer)
	RightMouseUp(event unsafe.Pointer)
	RotateWithEvent(event unsafe.Pointer)
	ScrollWheel(event unsafe.Pointer)
	SetInterfaceStyle(interfaceStyle unsafe.Pointer)
	ShouldBeTreatedAsInkEvent(event unsafe.Pointer) bool
	ShowWritingTools(sender objc.ID)
	SmartMagnifyWithEvent(event unsafe.Pointer)
	SupplementalTargetForActionSender(action objc.SEL, sender objc.ID) objc.ID
	SwipeWithEvent(event unsafe.Pointer)
	TabletPoint(event unsafe.Pointer)
	TabletProximity(event unsafe.Pointer)
	TouchesBeganWithEvent(event unsafe.Pointer)
	TouchesCancelledWithEvent(event unsafe.Pointer)
	TouchesEndedWithEvent(event unsafe.Pointer)
	TouchesMovedWithEvent(event unsafe.Pointer)
	TryToPerformWith(action objc.SEL, object objc.ID) bool
	UpdateUserActivityState(userActivity unsafe.Pointer)
	ValidRequestorForSendTypeReturnType(sendType unsafe.Pointer, returnType unsafe.Pointer) objc.ID
	ValidateProposedFirstResponderForEvent(responder unsafe.Pointer, event unsafe.Pointer) bool
	WantsForwardedScrollEventsForAxis(axis unsafe.Pointer) bool
	WantsScrollEventsForSwipeTrackingOnAxis(axis unsafe.Pointer) bool
	WillPresentError(error unsafe.Pointer) unsafe.Pointer
}

type Responder struct {
	id objc.ID
}

func ResponderFrom(ptr unsafe.Pointer) Responder {
	return Responder{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ Responder) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _ResponderClass) Alloc() Responder {
	rv := objc.Send[Responder](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _ResponderClass) New() Responder {
	rv := objc.Send[Responder](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewResponder creates and returns a new initialized instance.
func NewResponder() Responder {
	return ResponderClass.New()
}

// Init initializes the instance.
func (r_ Responder) Init() Responder {
	rv := objc.Send[Responder](r_.ID(), selInit)
	return rv
}
// Returns the classes that support secure coding. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/allowedClasses(forRestorableStateKeyPath:)
func (rc _ResponderClass) AllowedClassesForRestorableStateKeyPath(keyPath string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(rc.Class), objc.RegisterName("allowedClassesForRestorableStateKeyPath:"), keyPath)
	return rv
}
// Notifies the receiver that it’s about to become first responder in its  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/becomeFirstResponder()
func (r_ Responder) BecomeFirstResponder() bool {
	rv := objc.Send[bool](r_.ID(), objc.RegisterName("becomeFirstResponder"))
	return rv
}
// Informs the receiver that the user has begun a touch gesture. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/beginGesture(with:)
func (r_ Responder) BeginGestureWithEvent(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("beginGestureWithEvent:"), event)
}
// Informs the responder that performed a double-tap on the side of an Apple Pencil. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/changeMode(with:)
func (r_ Responder) ChangeModeWithEvent(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("changeModeWithEvent:"), event)
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/contextMenuKeyDown(_:)
func (r_ Responder) ContextMenuKeyDown(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("contextMenuKeyDown:"), event)
}
// Informs the receiver that the mouse cursor has moved into a cursor rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/cursorUpdate(with:)
func (r_ Responder) CursorUpdate(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("cursorUpdate:"), event)
}
// Saves the interface-related state of the responder. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/encodeRestorableState(with:)
func (r_ Responder) EncodeRestorableStateWithCoder(coder unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("encodeRestorableStateWithCoder:"), coder)
}
// Saves the interface-related state of the responder to a keyed archiver either synchronously or asynchronously on the given operation queue. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/encodeRestorableState(with:backgroundQueue:)
func (r_ Responder) EncodeRestorableStateWithCoderBackgroundQueue(coder unsafe.Pointer, queue unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("encodeRestorableStateWithCoder:backgroundQueue:"), coder, queue)
}
// Informs the receiver that the user has ended a touch gesture. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/endGesture(with:)
func (r_ Responder) EndGestureWithEvent(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("endGestureWithEvent:"), event)
}
// Informs the receiver that the user has pressed or released a modifier key (Shift, Control, and so on). [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/flagsChanged(with:)
func (r_ Responder) FlagsChanged(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("flagsChanged:"), event)
}
// Clears any unprocessed key events when overridden by subclasses. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/flushBufferedKeyEvents()
func (r_ Responder) FlushBufferedKeyEvents() {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("flushBufferedKeyEvents"))
}
// Displays context-sensitive help for the receiver if help has been registered. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/helpRequested(_:)
func (r_ Responder) HelpRequested(eventPtr unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("helpRequested:"), eventPtr)
}
// Returns the receiver’s interface style. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/interfaceStyle
func (r_ Responder) InterfaceStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID(), objc.RegisterName("interfaceStyle"))
	return rv
}
// Handles a series of key events. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/interpretKeyEvents(_:)
func (r_ Responder) InterpretKeyEvents(eventArray unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("interpretKeyEvents:"), eventArray)
}
// Marks the responder’s interface-related state as dirty. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/invalidateRestorableState()
func (r_ Responder) InvalidateRestorableState() {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("invalidateRestorableState"))
}
// Informs the receiver that the user has pressed a key. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/keyDown(with:)
func (r_ Responder) KeyDown(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("keyDown:"), event)
}
// Informs the receiver that the user has released a key. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/keyUp(with:)
func (r_ Responder) KeyUp(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("keyUp:"), event)
}
// Informs the receiver that the user has begun a pinch gesture. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/magnify(with:)
func (r_ Responder) MagnifyWithEvent(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("magnifyWithEvent:"), event)
}
// Your custom subclass of the   class should override this method to create and configure your subclass’s default   object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/makeTouchBar()
func (r_ Responder) MakeTouchBar() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID(), objc.RegisterName("makeTouchBar"))
	return rv
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/mouseCancelled(with:)
func (r_ Responder) MouseCancelled(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("mouseCancelled:"), event)
}
// Informs the receiver that the user has pressed the left mouse button. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/mouseDown(with:)
func (r_ Responder) MouseDown(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("mouseDown:"), event)
}
// Informs the receiver that the user has moved the mouse with the left button pressed. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/mouseDragged(with:)
func (r_ Responder) MouseDragged(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("mouseDragged:"), event)
}
// Informs the receiver that the cursor has entered a tracking rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/mouseEntered(with:)
func (r_ Responder) MouseEntered(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("mouseEntered:"), event)
}
// Informs the receiver that the cursor has exited a tracking rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/mouseExited(with:)
func (r_ Responder) MouseExited(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("mouseExited:"), event)
}
// Informs the receiver that the mouse has moved. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/mouseMoved(with:)
func (r_ Responder) MouseMoved(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("mouseMoved:"), event)
}
// Informs the receiver that the user has released the left mouse button. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/mouseUp(with:)
func (r_ Responder) MouseUp(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("mouseUp:"), event)
}
// Creates a new window to show as a tab in a tabbed window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/newWindowForTab(_:)
func (r_ Responder) NewWindowForTab(sender objc.ID) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("newWindowForTab:"), sender)
}
// Handles the case where an event or action message falls off the end of the responder chain. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/noResponder(for:)
func (r_ Responder) NoResponderFor(eventSelector objc.SEL) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("noResponderFor:"), eventSelector)
}
// Informs the receiver that the user has pressed a mouse button other than the left or right one. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/otherMouseDown(with:)
func (r_ Responder) OtherMouseDown(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("otherMouseDown:"), event)
}
// Informs the receiver that the user has moved the mouse with a button other than the left or right button pressed. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/otherMouseDragged(with:)
func (r_ Responder) OtherMouseDragged(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("otherMouseDragged:"), event)
}
// Informs the receiver that the user has released a mouse button other than the left or right button. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/otherMouseUp(with:)
func (r_ Responder) OtherMouseUp(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("otherMouseUp:"), event)
}
// Handle a key equivalent. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/performKeyEquivalent(with:)
func (r_ Responder) PerformKeyEquivalent(event unsafe.Pointer) bool {
	rv := objc.Send[bool](r_.ID(), objc.RegisterName("performKeyEquivalent:"), event)
	return rv
}
// Handle a mnemonic. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/performMnemonic:
func (r_ Responder) PerformMnemonic(string string) bool {
	rv := objc.Send[bool](r_.ID(), objc.RegisterName("performMnemonic:"), string)
	return rv
}
// Performs all find oriented actions. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/performTextFinderAction(_:)
func (r_ Responder) PerformTextFinderAction(sender objc.ID) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("performTextFinderAction:"), sender)
}
// Presents an error alert to the user as an application-modal dialog. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/presentError(_:)
func (r_ Responder) PresentError(error unsafe.Pointer) bool {
	rv := objc.Send[bool](r_.ID(), objc.RegisterName("presentError:"), error)
	return rv
}
// Presents an error alert to the user as a document-modal sheet attached to document window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/presentError(_:modalFor:delegate:didPresent:contextInfo:)
func (r_ Responder) PresentErrorModalForWindowDelegateDidPresentSelectorContextInfo(error unsafe.Pointer, window unsafe.Pointer, delegate objc.ID, didPresentSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("presentError:modalForWindow:delegate:didPresentSelector:contextInfo:"), error, window, delegate, didPresentSelector, contextInfo)
}
// Indicates a pressure change as the result of a user input event on a system that supports pressure sensitivity. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/pressureChange(with:)
func (r_ Responder) PressureChangeWithEvent(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("pressureChangeWithEvent:"), event)
}
// Performs a Quick Look on the content at the location specified by the supplied event. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/quickLook(with:)
func (r_ Responder) QuickLookWithEvent(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("quickLookWithEvent:"), event)
}
// Notifies the receiver that it’s been asked to relinquish its status as first responder in its window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/resignFirstResponder()
func (r_ Responder) ResignFirstResponder() bool {
	rv := objc.Send[bool](r_.ID(), objc.RegisterName("resignFirstResponder"))
	return rv
}
// Restores the interface-related state of the responder. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/restoreState(with:)
func (r_ Responder) RestoreStateWithCoder(coder unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("restoreStateWithCoder:"), coder)
}
// Informs the receiver that the user has pressed the right mouse button. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/rightMouseDown(with:)
func (r_ Responder) RightMouseDown(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("rightMouseDown:"), event)
}
// Informs the receiver that the user has moved the mouse with the right button pressed. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/rightMouseDragged(with:)
func (r_ Responder) RightMouseDragged(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("rightMouseDragged:"), event)
}
// Informs the receiver that the user has released the right mouse button. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/rightMouseUp(with:)
func (r_ Responder) RightMouseUp(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("rightMouseUp:"), event)
}
// Informs the receiver that the user has begun a rotation gesture. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/rotate(with:)
func (r_ Responder) RotateWithEvent(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("rotateWithEvent:"), event)
}
// Informs the receiver that the mouse’s scroll wheel has moved. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/scrollWheel(with:)
func (r_ Responder) ScrollWheel(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("scrollWheel:"), event)
}
// Sets the receiver’s style to the style specified by  , such as   or  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/setInterfaceStyle:
func (r_ Responder) SetInterfaceStyle(interfaceStyle unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("setInterfaceStyle:"), interfaceStyle)
}
// Indicates whether a pen-down event should be treated as an ink event. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/shouldBeTreatedAsInkEvent(_:)
func (r_ Responder) ShouldBeTreatedAsInkEvent(event unsafe.Pointer) bool {
	rv := objc.Send[bool](r_.ID(), objc.RegisterName("shouldBeTreatedAsInkEvent:"), event)
	return rv
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/showWritingTools(_:)
func (r_ Responder) ShowWritingTools(sender objc.ID) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("showWritingTools:"), sender)
}
// Informs the receiver that the user performed a smart zoom gesture. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/smartMagnify(with:)
func (r_ Responder) SmartMagnifyWithEvent(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("smartMagnifyWithEvent:"), event)
}
// Finds a target for an action method. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/supplementalTarget(forAction:sender:)
func (r_ Responder) SupplementalTargetForActionSender(action objc.SEL, sender objc.ID) objc.ID {
	rv := objc.Send[objc.ID](r_.ID(), objc.RegisterName("supplementalTargetForAction:sender:"), action, sender)
	return rv
}
// Informs the receiver that the user has begun a swipe gesture. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/swipe(with:)
func (r_ Responder) SwipeWithEvent(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("swipeWithEvent:"), event)
}
// Informs the receiver that a tablet-point event has occurred. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/tabletPoint(with:)
func (r_ Responder) TabletPoint(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("tabletPoint:"), event)
}
// Informs the receiver that a tablet-proximity event has occurred. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/tabletProximity(with:)
func (r_ Responder) TabletProximity(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("tabletProximity:"), event)
}
// Informs the receiver that new set of touches has been recognized. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/touchesBegan(with:)
func (r_ Responder) TouchesBeganWithEvent(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("touchesBeganWithEvent:"), event)
}
// Informs the receiver that tracking of touches has been cancelled for any reason. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/touchesCancelled(with:)
func (r_ Responder) TouchesCancelledWithEvent(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("touchesCancelledWithEvent:"), event)
}
// Returns that a set of touches have been removed. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/touchesEnded(with:)
func (r_ Responder) TouchesEndedWithEvent(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("touchesEndedWithEvent:"), event)
}
// Informs the receiver that one or more touches has moved. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/touchesMoved(with:)
func (r_ Responder) TouchesMovedWithEvent(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("touchesMovedWithEvent:"), event)
}
// Attempts to perform the method indicated by an action with a specified argument. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/tryToPerform(_:with:)
func (r_ Responder) TryToPerformWith(action objc.SEL, object objc.ID) bool {
	rv := objc.Send[bool](r_.ID(), objc.RegisterName("tryToPerform:with:"), action, object)
	return rv
}
// Updates the state of the given user activity. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/updateUserActivityState(_:)
func (r_ Responder) UpdateUserActivityState(userActivity unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("updateUserActivityState:"), userActivity)
}
// Overridden by subclasses to determine what services are available. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/validRequestor(forSendType:returnType:)
func (r_ Responder) ValidRequestorForSendTypeReturnType(sendType unsafe.Pointer, returnType unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](r_.ID(), objc.RegisterName("validRequestorForSendType:returnType:"), sendType, returnType)
	return rv
}
// Allows controls to determine when they should become first responder. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/validateProposedFirstResponder(_:for:)
func (r_ Responder) ValidateProposedFirstResponderForEvent(responder unsafe.Pointer, event unsafe.Pointer) bool {
	rv := objc.Send[bool](r_.ID(), objc.RegisterName("validateProposedFirstResponder:forEvent:"), responder, event)
	return rv
}
// Returns whether to forward elastic scrolling gesture events up the responder. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/wantsForwardedScrollEvents(for:)
func (r_ Responder) WantsForwardedScrollEventsForAxis(axis unsafe.Pointer) bool {
	rv := objc.Send[bool](r_.ID(), objc.RegisterName("wantsForwardedScrollEventsForAxis:"), axis)
	return rv
}
// Implement this method to track gesture scroll events such as a swipe. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/wantsScrollEventsForSwipeTracking(on:)
func (r_ Responder) WantsScrollEventsForSwipeTrackingOnAxis(axis unsafe.Pointer) bool {
	rv := objc.Send[bool](r_.ID(), objc.RegisterName("wantsScrollEventsForSwipeTrackingOnAxis:"), axis)
	return rv
}
// Returns a custom version of the supplied error object that’s more suitable for presentation in alert sheets and dialogs. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/willPresentError(_:)
func (r_ Responder) WillPresentError(error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID(), objc.RegisterName("willPresentError:"), error)
	return rv
}
// A Boolean value that indicates whether the responder accepts first responder status. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/acceptsFirstResponder
func (r_ Responder) AcceptsFirstResponder() bool {
	rv := objc.Send[bool](r_.ID(), objc.RegisterName("acceptsFirstResponder"))
	return rv
}
// Returns the responder’s menu. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/menu
func (r_ Responder) Menu() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID(), objc.RegisterName("menu"))
	return rv
}
// SetMenu sets the value of the menu property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/menu
func (r_ Responder) SetMenu(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("setMenu:"), value)
}
// The next responder after this one, or   if it has none. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/nextResponder
func (r_ Responder) NextResponder() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID(), objc.RegisterName("nextResponder"))
	return rv
}
// SetNextResponder sets the value of the nextResponder property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/nextResponder
func (r_ Responder) SetNextResponder(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("setNextResponder:"), value)
}
// The   object associated with the responder. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/touchBar
func (r_ Responder) TouchBar() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID(), objc.RegisterName("touchBar"))
	return rv
}
// SetTouchBar sets the value of the touchBar property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/touchBar
func (r_ Responder) SetTouchBar(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("setTouchBar:"), value)
}
// The undo manager for this responder. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/undoManager
func (r_ Responder) UndoManager() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID(), objc.RegisterName("undoManager"))
	return rv
}
// An object encapsulating a user activity supported by this responder. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/userActivity
func (r_ Responder) UserActivity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID(), objc.RegisterName("userActivity"))
	return rv
}
// SetUserActivity sets the value of the userActivity property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/userActivity
func (r_ Responder) SetUserActivity(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("setUserActivity:"), value)
}
