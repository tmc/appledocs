// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Responder] class.
var responderClass = _ResponderClass{objc.GetClass("NSResponder")}

type _ResponderClass struct {
	class objc.Class
}

// An abstract class that forms the basis of event and command processing in AppKit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder

type Responder struct {
	objectivec.Object
}

// ResponderFrom constructs a [Responder] from an unsafe.Pointer.
//
// An abstract class that forms the basis of event and command processing in AppKit.
func ResponderFrom(ptr unsafe.Pointer) Responder {
	return Responder{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (rc _ResponderClass) Alloc() Responder {
	rv := objc.Send[Responder](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (rc _ResponderClass) New() Responder {
	rv := objc.Send[Responder](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ Responder) Init() Responder {
	rv := objc.Send[Responder](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ Responder) Autorelease() Responder {
	rv := objc.Send[Responder](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewResponder creates a new Responder instance.
func NewResponder() Responder {
	return responderClass.New()
}
// Creates a new responder object with data in an unarchiver. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/init(coder:)
func NewResponderWithCoder(coder unsafe.Pointer) Responder {
	instance := responderClass.Alloc()
	rv := objc.Send[Responder](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}


// Returns the classes that support secure coding. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/allowedClasses(forRestorableStateKeyPath:)
func (rc _ResponderClass) AllowedClassesForRestorableStateKeyPath(keyPath string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(rc.class), objc.Sel("allowedClassesForRestorableStateKeyPath:"), keyPath)
	return rv
}
// Notifies the receiver that it’s about to become first responder in its . [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/becomeFirstResponder()
func (r_ Responder) BecomeFirstResponder() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("becomeFirstResponder"))
	return rv
}
// Informs the receiver that the user has begun a touch gesture. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/beginGesture(with:)
func (r_ Responder) BeginGestureWithEvent(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("beginGestureWithEvent:"), event)
}
// Informs the responder that performed a double-tap on the side of an Apple Pencil. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/changeMode(with:)
func (r_ Responder) ChangeModeWithEvent(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("changeModeWithEvent:"), event)
}
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/contextMenuKeyDown(_:)
func (r_ Responder) ContextMenuKeyDown(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("contextMenuKeyDown:"), event)
}
// Informs the receiver that the mouse cursor has moved into a cursor rectangle. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/cursorUpdate(with:)
func (r_ Responder) CursorUpdate(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("cursorUpdate:"), event)
}
// Saves the interface-related state of the responder. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/encodeRestorableState(with:)
func (r_ Responder) EncodeRestorableStateWithCoder(coder unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("encodeRestorableStateWithCoder:"), coder)
}
// Saves the interface-related state of the responder to a keyed archiver either synchronously or asynchronously on the given operation queue. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/encodeRestorableState(with:backgroundQueue:)
func (r_ Responder) EncodeRestorableStateWithCoderBackgroundQueue(coder unsafe.Pointer, queue unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("encodeRestorableStateWithCoder:backgroundQueue:"), coder, queue)
}
// Informs the receiver that the user has ended a touch gesture. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/endGesture(with:)
func (r_ Responder) EndGestureWithEvent(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("endGestureWithEvent:"), event)
}
// Informs the receiver that the user has pressed or released a modifier key (Shift, Control, and so on). [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/flagsChanged(with:)
func (r_ Responder) FlagsChanged(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("flagsChanged:"), event)
}
// Clears any unprocessed key events when overridden by subclasses. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/flushBufferedKeyEvents()
func (r_ Responder) FlushBufferedKeyEvents() {
	objc.Send[objc.ID](r_.ID, objc.Sel("flushBufferedKeyEvents"))
}
// Displays context-sensitive help for the receiver if help has been registered. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/helpRequested(_:)
func (r_ Responder) HelpRequested(eventPtr unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("helpRequested:"), eventPtr)
}
// Returns the receiver’s interface style. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/interfaceStyle
func (r_ Responder) InterfaceStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("interfaceStyle"))
	return rv
}
// Handles a series of key events. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/interpretKeyEvents(_:)
func (r_ Responder) InterpretKeyEvents(eventArray unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("interpretKeyEvents:"), eventArray)
}
// Marks the responder’s interface-related state as dirty. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/invalidateRestorableState()
func (r_ Responder) InvalidateRestorableState() {
	objc.Send[objc.ID](r_.ID, objc.Sel("invalidateRestorableState"))
}
// Informs the receiver that the user has pressed a key. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/keyDown(with:)
func (r_ Responder) KeyDown(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("keyDown:"), event)
}
// Informs the receiver that the user has released a key. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/keyUp(with:)
func (r_ Responder) KeyUp(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("keyUp:"), event)
}
// Informs the receiver that the user has begun a pinch gesture. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/magnify(with:)
func (r_ Responder) MagnifyWithEvent(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("magnifyWithEvent:"), event)
}
// Your custom subclass of the class should override this method to create and configure your subclass’s default object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/makeTouchBar()
func (r_ Responder) MakeTouchBar() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("makeTouchBar"))
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/mouseCancelled(with:)
func (r_ Responder) MouseCancelled(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("mouseCancelled:"), event)
}
// Informs the receiver that the user has pressed the left mouse button. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/mouseDown(with:)
func (r_ Responder) MouseDown(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("mouseDown:"), event)
}
// Informs the receiver that the user has moved the mouse with the left button pressed. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/mouseDragged(with:)
func (r_ Responder) MouseDragged(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("mouseDragged:"), event)
}
// Informs the receiver that the cursor has entered a tracking rectangle. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/mouseEntered(with:)
func (r_ Responder) MouseEntered(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("mouseEntered:"), event)
}
// Informs the receiver that the cursor has exited a tracking rectangle. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/mouseExited(with:)
func (r_ Responder) MouseExited(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("mouseExited:"), event)
}
// Informs the receiver that the mouse has moved. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/mouseMoved(with:)
func (r_ Responder) MouseMoved(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("mouseMoved:"), event)
}
// Informs the receiver that the user has released the left mouse button. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/mouseUp(with:)
func (r_ Responder) MouseUp(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("mouseUp:"), event)
}
// Creates a new window to show as a tab in a tabbed window. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/newWindowForTab(_:)
func (r_ Responder) NewWindowForTab(sender objc.ID) {
	objc.Send[objc.ID](r_.ID, objc.Sel("newWindowForTab:"), sender)
}
// Handles the case where an event or action message falls off the end of the responder chain. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/noResponder(for:)
func (r_ Responder) NoResponderFor(eventSelector objc.SEL) {
	objc.Send[objc.ID](r_.ID, objc.Sel("noResponderFor:"), eventSelector)
}
// Informs the receiver that the user has pressed a mouse button other than the left or right one. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/otherMouseDown(with:)
func (r_ Responder) OtherMouseDown(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("otherMouseDown:"), event)
}
// Informs the receiver that the user has moved the mouse with a button other than the left or right button pressed. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/otherMouseDragged(with:)
func (r_ Responder) OtherMouseDragged(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("otherMouseDragged:"), event)
}
// Informs the receiver that the user has released a mouse button other than the left or right button. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/otherMouseUp(with:)
func (r_ Responder) OtherMouseUp(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("otherMouseUp:"), event)
}
// Handle a key equivalent. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/performKeyEquivalent(with:)
func (r_ Responder) PerformKeyEquivalent(event unsafe.Pointer) bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("performKeyEquivalent:"), event)
	return rv
}
// Handle a mnemonic. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/performMnemonic:
func (r_ Responder) PerformMnemonic(string string) bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("performMnemonic:"), string)
	return rv
}
// Performs all find oriented actions. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/performTextFinderAction(_:)
func (r_ Responder) PerformTextFinderAction(sender objc.ID) {
	objc.Send[objc.ID](r_.ID, objc.Sel("performTextFinderAction:"), sender)
}
// Presents an error alert to the user as an application-modal dialog. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/presentError(_:)
func (r_ Responder) PresentError(error unsafe.Pointer) bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("presentError:"), error)
	return rv
}
// Presents an error alert to the user as a document-modal sheet attached to document window. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/presentError(_:modalFor:delegate:didPresent:contextInfo:)
func (r_ Responder) PresentErrorModalForWindowDelegateDidPresentSelectorContextInfo(error unsafe.Pointer, window unsafe.Pointer, delegate objc.ID, didPresentSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("presentError:modalForWindow:delegate:didPresentSelector:contextInfo:"), error, window, delegate, didPresentSelector, contextInfo)
}
// Indicates a pressure change as the result of a user input event on a system that supports pressure sensitivity. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/pressureChange(with:)
func (r_ Responder) PressureChangeWithEvent(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("pressureChangeWithEvent:"), event)
}
// Performs a Quick Look on the content at the location specified by the supplied event. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/quickLook(with:)
func (r_ Responder) QuickLookWithEvent(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("quickLookWithEvent:"), event)
}
// Notifies the receiver that it’s been asked to relinquish its status as first responder in its window. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/resignFirstResponder()
func (r_ Responder) ResignFirstResponder() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("resignFirstResponder"))
	return rv
}
// Restores the interface-related state of the responder. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/restoreState(with:)
func (r_ Responder) RestoreStateWithCoder(coder unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("restoreStateWithCoder:"), coder)
}
// Informs the receiver that the user has pressed the right mouse button. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/rightMouseDown(with:)
func (r_ Responder) RightMouseDown(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("rightMouseDown:"), event)
}
// Informs the receiver that the user has moved the mouse with the right button pressed. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/rightMouseDragged(with:)
func (r_ Responder) RightMouseDragged(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("rightMouseDragged:"), event)
}
// Informs the receiver that the user has released the right mouse button. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/rightMouseUp(with:)
func (r_ Responder) RightMouseUp(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("rightMouseUp:"), event)
}
// Informs the receiver that the user has begun a rotation gesture. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/rotate(with:)
func (r_ Responder) RotateWithEvent(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("rotateWithEvent:"), event)
}
// Informs the receiver that the mouse’s scroll wheel has moved. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/scrollWheel(with:)
func (r_ Responder) ScrollWheel(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("scrollWheel:"), event)
}
// Sets the receiver’s style to the style specified by , such as or . [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/setInterfaceStyle:
func (r_ Responder) SetInterfaceStyle(interfaceStyle unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setInterfaceStyle:"), interfaceStyle)
}
// Indicates whether a pen-down event should be treated as an ink event. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/shouldBeTreatedAsInkEvent(_:)
func (r_ Responder) ShouldBeTreatedAsInkEvent(event unsafe.Pointer) bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("shouldBeTreatedAsInkEvent:"), event)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/showWritingTools(_:)
func (r_ Responder) ShowWritingTools(sender objc.ID) {
	objc.Send[objc.ID](r_.ID, objc.Sel("showWritingTools:"), sender)
}
// Informs the receiver that the user performed a smart zoom gesture. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/smartMagnify(with:)
func (r_ Responder) SmartMagnifyWithEvent(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("smartMagnifyWithEvent:"), event)
}
// Finds a target for an action method. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/supplementalTarget(forAction:sender:)
func (r_ Responder) SupplementalTargetForActionSender(action objc.SEL, sender objc.ID) objc.ID {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("supplementalTargetForAction:sender:"), action, sender)
	return rv
}
// Informs the receiver that the user has begun a swipe gesture. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/swipe(with:)
func (r_ Responder) SwipeWithEvent(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("swipeWithEvent:"), event)
}
// Informs the receiver that a tablet-point event has occurred. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/tabletPoint(with:)
func (r_ Responder) TabletPoint(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("tabletPoint:"), event)
}
// Informs the receiver that a tablet-proximity event has occurred. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/tabletProximity(with:)
func (r_ Responder) TabletProximity(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("tabletProximity:"), event)
}
// Informs the receiver that new set of touches has been recognized. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/touchesBegan(with:)
func (r_ Responder) TouchesBeganWithEvent(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("touchesBeganWithEvent:"), event)
}
// Informs the receiver that tracking of touches has been cancelled for any reason. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/touchesCancelled(with:)
func (r_ Responder) TouchesCancelledWithEvent(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("touchesCancelledWithEvent:"), event)
}
// Returns that a set of touches have been removed. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/touchesEnded(with:)
func (r_ Responder) TouchesEndedWithEvent(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("touchesEndedWithEvent:"), event)
}
// Informs the receiver that one or more touches has moved. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/touchesMoved(with:)
func (r_ Responder) TouchesMovedWithEvent(event unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("touchesMovedWithEvent:"), event)
}
// Attempts to perform the method indicated by an action with a specified argument. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/tryToPerform(_:with:)
func (r_ Responder) TryToPerformWith(action objc.SEL, object objc.ID) bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("tryToPerform:with:"), action, object)
	return rv
}
// Updates the state of the given user activity. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/updateUserActivityState(_:)
func (r_ Responder) UpdateUserActivityState(userActivity unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("updateUserActivityState:"), userActivity)
}
// Overridden by subclasses to determine what services are available. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/validRequestor(forSendType:returnType:)
func (r_ Responder) ValidRequestorForSendTypeReturnType(sendType unsafe.Pointer, returnType unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("validRequestorForSendType:returnType:"), sendType, returnType)
	return rv
}
// Allows controls to determine when they should become first responder. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/validateProposedFirstResponder(_:for:)
func (r_ Responder) ValidateProposedFirstResponderForEvent(responder unsafe.Pointer, event unsafe.Pointer) bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("validateProposedFirstResponder:forEvent:"), responder, event)
	return rv
}
// Returns whether to forward elastic scrolling gesture events up the responder. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/wantsForwardedScrollEvents(for:)
func (r_ Responder) WantsForwardedScrollEventsForAxis(axis unsafe.Pointer) bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("wantsForwardedScrollEventsForAxis:"), axis)
	return rv
}
// Implement this method to track gesture scroll events such as a swipe. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/wantsScrollEventsForSwipeTracking(on:)
func (r_ Responder) WantsScrollEventsForSwipeTrackingOnAxis(axis unsafe.Pointer) bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("wantsScrollEventsForSwipeTrackingOnAxis:"), axis)
	return rv
}
// Returns a custom version of the supplied error object that’s more suitable for presentation in alert sheets and dialogs. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/willPresentError(_:)
func (r_ Responder) WillPresentError(error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("willPresentError:"), error)
	return rv
}

