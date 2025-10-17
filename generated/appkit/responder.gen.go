// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Responder] class.
var ResponderClass objc.Class

func init() {
	ResponderClass = objc.GetClass("NSResponder")
}

type Responder struct {
	objc.ID
}

func ResponderFrom(ptr unsafe.Pointer) Responder {
	return Responder{
		ID: objc.ID(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc Responder) Alloc() Responder {
	ret := objc.ID(ResponderClass).Send(objc.RegisterName("alloc"))
	return Responder{ret}
}

// Init initializes the instance.
func (r_ Responder) Init() Responder {
	ret := r_.ID.Send(objc.RegisterName("init"))
	return Responder{ret}
}
// Creates a new responder object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/init()
func NewResponder() Responder {
	instance := Responder{}.Alloc()
	instance = instance.Init()
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Creates a new responder object with data in an unarchiver. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/init(coder:)
func NewResponderWithCoder(coder unsafe.Pointer) Responder {
	instance := Responder{}.Alloc()
	sel := objc.RegisterName("initWithCoder:")
	ret := instance.ID.Send(sel, coder)
	instance = Responder{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}


// Returns the classes that support secure coding. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/allowedClasses(forRestorableStateKeyPath:)
func (rc Responder) AllowedClassesForRestorableStateKeyPath(keyPath unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("allowedClassesForRestorableStateKeyPath:")
	ret := objc.ID(ResponderClass).Send(sel, keyPath)
	return unsafe.Pointer(ret)
}
// Notifies the receiver that it’s about to become first responder in its  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/becomeFirstResponder()
func (r_ Responder) BecomeFirstResponder() bool {
	sel := objc.RegisterName("becomeFirstResponder")
	ret := r_.ID.Send(sel)
	return ret != 0
}
// Informs the receiver that the user has begun a touch gesture. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/beginGesture(with:)
func (r_ Responder) BeginGestureWithEvent(event unsafe.Pointer) {
	sel := objc.RegisterName("beginGestureWithEvent:")
	r_.ID.Send(sel, event)
}
// Informs the responder that performed a double-tap on the side of an Apple Pencil. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/changeMode(with:)
func (r_ Responder) ChangeModeWithEvent(event unsafe.Pointer) {
	sel := objc.RegisterName("changeModeWithEvent:")
	r_.ID.Send(sel, event)
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/contextMenuKeyDown(_:)
func (r_ Responder) ContextMenuKeyDown(event unsafe.Pointer) {
	sel := objc.RegisterName("contextMenuKeyDown:")
	r_.ID.Send(sel, event)
}
// Informs the receiver that the mouse cursor has moved into a cursor rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/cursorUpdate(with:)
func (r_ Responder) CursorUpdate(event unsafe.Pointer) {
	sel := objc.RegisterName("cursorUpdate:")
	r_.ID.Send(sel, event)
}
// Saves the interface-related state of the responder. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/encodeRestorableState(with:)
func (r_ Responder) EncodeRestorableStateWithCoder(coder unsafe.Pointer) {
	sel := objc.RegisterName("encodeRestorableStateWithCoder:")
	r_.ID.Send(sel, coder)
}
// Saves the interface-related state of the responder to a keyed archiver either synchronously or asynchronously on the given operation queue. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/encodeRestorableState(with:backgroundQueue:)
func (r_ Responder) EncodeRestorableStateWithCoderBackgroundQueue(coder unsafe.Pointer, queue unsafe.Pointer) {
	sel := objc.RegisterName("encodeRestorableStateWithCoder:backgroundQueue:")
	r_.ID.Send(sel, coder, queue)
}
// Informs the receiver that the user has ended a touch gesture. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/endGesture(with:)
func (r_ Responder) EndGestureWithEvent(event unsafe.Pointer) {
	sel := objc.RegisterName("endGestureWithEvent:")
	r_.ID.Send(sel, event)
}
// Informs the receiver that the user has pressed or released a modifier key (Shift, Control, and so on). [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/flagsChanged(with:)
func (r_ Responder) FlagsChanged(event unsafe.Pointer) {
	sel := objc.RegisterName("flagsChanged:")
	r_.ID.Send(sel, event)
}
// Clears any unprocessed key events when overridden by subclasses. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/flushBufferedKeyEvents()
func (r_ Responder) FlushBufferedKeyEvents() {
	sel := objc.RegisterName("flushBufferedKeyEvents")
	r_.ID.Send(sel)
}
// Displays context-sensitive help for the receiver if help has been registered. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/helpRequested(_:)
func (r_ Responder) HelpRequested(eventPtr unsafe.Pointer) {
	sel := objc.RegisterName("helpRequested:")
	r_.ID.Send(sel, eventPtr)
}
// Returns the receiver’s interface style. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/interfaceStyle
func (r_ Responder) InterfaceStyle() unsafe.Pointer {
	sel := objc.RegisterName("interfaceStyle")
	ret := r_.ID.Send(sel)
	return unsafe.Pointer(ret)
}
// Handles a series of key events. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/interpretKeyEvents(_:)
func (r_ Responder) InterpretKeyEvents(eventArray unsafe.Pointer) {
	sel := objc.RegisterName("interpretKeyEvents:")
	r_.ID.Send(sel, eventArray)
}
// Marks the responder’s interface-related state as dirty. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/invalidateRestorableState()
func (r_ Responder) InvalidateRestorableState() {
	sel := objc.RegisterName("invalidateRestorableState")
	r_.ID.Send(sel)
}
// Informs the receiver that the user has pressed a key. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/keyDown(with:)
func (r_ Responder) KeyDown(event unsafe.Pointer) {
	sel := objc.RegisterName("keyDown:")
	r_.ID.Send(sel, event)
}
// Informs the receiver that the user has released a key. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/keyUp(with:)
func (r_ Responder) KeyUp(event unsafe.Pointer) {
	sel := objc.RegisterName("keyUp:")
	r_.ID.Send(sel, event)
}
// Informs the receiver that the user has begun a pinch gesture. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/magnify(with:)
func (r_ Responder) MagnifyWithEvent(event unsafe.Pointer) {
	sel := objc.RegisterName("magnifyWithEvent:")
	r_.ID.Send(sel, event)
}
// Your custom subclass of the   class should override this method to create and configure your subclass’s default   object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/makeTouchBar()
func (r_ Responder) MakeTouchBar() unsafe.Pointer {
	sel := objc.RegisterName("makeTouchBar")
	ret := r_.ID.Send(sel)
	return unsafe.Pointer(ret)
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/mouseCancelled(with:)
func (r_ Responder) MouseCancelled(event unsafe.Pointer) {
	sel := objc.RegisterName("mouseCancelled:")
	r_.ID.Send(sel, event)
}
// Informs the receiver that the user has pressed the left mouse button. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/mouseDown(with:)
func (r_ Responder) MouseDown(event unsafe.Pointer) {
	sel := objc.RegisterName("mouseDown:")
	r_.ID.Send(sel, event)
}
// Informs the receiver that the user has moved the mouse with the left button pressed. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/mouseDragged(with:)
func (r_ Responder) MouseDragged(event unsafe.Pointer) {
	sel := objc.RegisterName("mouseDragged:")
	r_.ID.Send(sel, event)
}
// Informs the receiver that the cursor has entered a tracking rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/mouseEntered(with:)
func (r_ Responder) MouseEntered(event unsafe.Pointer) {
	sel := objc.RegisterName("mouseEntered:")
	r_.ID.Send(sel, event)
}
// Informs the receiver that the cursor has exited a tracking rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/mouseExited(with:)
func (r_ Responder) MouseExited(event unsafe.Pointer) {
	sel := objc.RegisterName("mouseExited:")
	r_.ID.Send(sel, event)
}
// Informs the receiver that the mouse has moved. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/mouseMoved(with:)
func (r_ Responder) MouseMoved(event unsafe.Pointer) {
	sel := objc.RegisterName("mouseMoved:")
	r_.ID.Send(sel, event)
}
// Informs the receiver that the user has released the left mouse button. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/mouseUp(with:)
func (r_ Responder) MouseUp(event unsafe.Pointer) {
	sel := objc.RegisterName("mouseUp:")
	r_.ID.Send(sel, event)
}
// Creates a new window to show as a tab in a tabbed window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/newWindowForTab(_:)
func (r_ Responder) NewWindowForTab(sender objc.ID) {
	sel := objc.RegisterName("newWindowForTab:")
	r_.ID.Send(sel, sender)
}
// Handles the case where an event or action message falls off the end of the responder chain. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/noResponder(for:)
func (r_ Responder) NoResponderFor(eventSelector objc.SEL) {
	sel := objc.RegisterName("noResponderFor:")
	r_.ID.Send(sel, eventSelector)
}
// Informs the receiver that the user has pressed a mouse button other than the left or right one. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/otherMouseDown(with:)
func (r_ Responder) OtherMouseDown(event unsafe.Pointer) {
	sel := objc.RegisterName("otherMouseDown:")
	r_.ID.Send(sel, event)
}
// Informs the receiver that the user has moved the mouse with a button other than the left or right button pressed. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/otherMouseDragged(with:)
func (r_ Responder) OtherMouseDragged(event unsafe.Pointer) {
	sel := objc.RegisterName("otherMouseDragged:")
	r_.ID.Send(sel, event)
}
// Informs the receiver that the user has released a mouse button other than the left or right button. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/otherMouseUp(with:)
func (r_ Responder) OtherMouseUp(event unsafe.Pointer) {
	sel := objc.RegisterName("otherMouseUp:")
	r_.ID.Send(sel, event)
}
// Handle a key equivalent. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/performKeyEquivalent(with:)
func (r_ Responder) PerformKeyEquivalent(event unsafe.Pointer) bool {
	sel := objc.RegisterName("performKeyEquivalent:")
	ret := r_.ID.Send(sel, event)
	return ret != 0
}
// Handle a mnemonic. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/performMnemonic:
func (r_ Responder) PerformMnemonic(string unsafe.Pointer) bool {
	sel := objc.RegisterName("performMnemonic:")
	ret := r_.ID.Send(sel, string)
	return ret != 0
}
// Performs all find oriented actions. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/performTextFinderAction(_:)
func (r_ Responder) PerformTextFinderAction(sender objc.ID) {
	sel := objc.RegisterName("performTextFinderAction:")
	r_.ID.Send(sel, sender)
}
// Presents an error alert to the user as an application-modal dialog. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/presentError(_:)
func (r_ Responder) PresentError(error unsafe.Pointer) bool {
	sel := objc.RegisterName("presentError:")
	ret := r_.ID.Send(sel, error)
	return ret != 0
}
// Presents an error alert to the user as a document-modal sheet attached to document window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/presentError(_:modalFor:delegate:didPresent:contextInfo:)
func (r_ Responder) PresentErrorModalForWindowDelegateDidPresentSelectorContextInfo(error unsafe.Pointer, window unsafe.Pointer, delegate objc.ID, didPresentSelector objc.SEL, contextInfo unsafe.Pointer) {
	sel := objc.RegisterName("presentError:modalForWindow:delegate:didPresentSelector:contextInfo:")
	r_.ID.Send(sel, error, window, delegate, didPresentSelector, contextInfo)
}
// Indicates a pressure change as the result of a user input event on a system that supports pressure sensitivity. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/pressureChange(with:)
func (r_ Responder) PressureChangeWithEvent(event unsafe.Pointer) {
	sel := objc.RegisterName("pressureChangeWithEvent:")
	r_.ID.Send(sel, event)
}
// Performs a Quick Look on the content at the location specified by the supplied event. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/quickLook(with:)
func (r_ Responder) QuickLookWithEvent(event unsafe.Pointer) {
	sel := objc.RegisterName("quickLookWithEvent:")
	r_.ID.Send(sel, event)
}
// Notifies the receiver that it’s been asked to relinquish its status as first responder in its window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/resignFirstResponder()
func (r_ Responder) ResignFirstResponder() bool {
	sel := objc.RegisterName("resignFirstResponder")
	ret := r_.ID.Send(sel)
	return ret != 0
}
// Restores the interface-related state of the responder. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/restoreState(with:)
func (r_ Responder) RestoreStateWithCoder(coder unsafe.Pointer) {
	sel := objc.RegisterName("restoreStateWithCoder:")
	r_.ID.Send(sel, coder)
}
// Informs the receiver that the user has pressed the right mouse button. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/rightMouseDown(with:)
func (r_ Responder) RightMouseDown(event unsafe.Pointer) {
	sel := objc.RegisterName("rightMouseDown:")
	r_.ID.Send(sel, event)
}
// Informs the receiver that the user has moved the mouse with the right button pressed. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/rightMouseDragged(with:)
func (r_ Responder) RightMouseDragged(event unsafe.Pointer) {
	sel := objc.RegisterName("rightMouseDragged:")
	r_.ID.Send(sel, event)
}
// Informs the receiver that the user has released the right mouse button. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/rightMouseUp(with:)
func (r_ Responder) RightMouseUp(event unsafe.Pointer) {
	sel := objc.RegisterName("rightMouseUp:")
	r_.ID.Send(sel, event)
}
// Informs the receiver that the user has begun a rotation gesture. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/rotate(with:)
func (r_ Responder) RotateWithEvent(event unsafe.Pointer) {
	sel := objc.RegisterName("rotateWithEvent:")
	r_.ID.Send(sel, event)
}
// Informs the receiver that the mouse’s scroll wheel has moved. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/scrollWheel(with:)
func (r_ Responder) ScrollWheel(event unsafe.Pointer) {
	sel := objc.RegisterName("scrollWheel:")
	r_.ID.Send(sel, event)
}
// Sets the receiver’s style to the style specified by  , such as   or  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/setInterfaceStyle:
func (r_ Responder) SetInterfaceStyle(interfaceStyle unsafe.Pointer) {
	sel := objc.RegisterName("setInterfaceStyle:")
	r_.ID.Send(sel, interfaceStyle)
}
// Indicates whether a pen-down event should be treated as an ink event. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/shouldBeTreatedAsInkEvent(_:)
func (r_ Responder) ShouldBeTreatedAsInkEvent(event unsafe.Pointer) bool {
	sel := objc.RegisterName("shouldBeTreatedAsInkEvent:")
	ret := r_.ID.Send(sel, event)
	return ret != 0
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/showWritingTools(_:)
func (r_ Responder) ShowWritingTools(sender objc.ID) {
	sel := objc.RegisterName("showWritingTools:")
	r_.ID.Send(sel, sender)
}
// Informs the receiver that the user performed a smart zoom gesture. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/smartMagnify(with:)
func (r_ Responder) SmartMagnifyWithEvent(event unsafe.Pointer) {
	sel := objc.RegisterName("smartMagnifyWithEvent:")
	r_.ID.Send(sel, event)
}
// Finds a target for an action method. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/supplementalTarget(forAction:sender:)
func (r_ Responder) SupplementalTargetForActionSender(action objc.SEL, sender objc.ID) objc.ID {
	sel := objc.RegisterName("supplementalTargetForAction:sender:")
	ret := r_.ID.Send(sel, action, sender)
	return ret
}
// Informs the receiver that the user has begun a swipe gesture. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/swipe(with:)
func (r_ Responder) SwipeWithEvent(event unsafe.Pointer) {
	sel := objc.RegisterName("swipeWithEvent:")
	r_.ID.Send(sel, event)
}
// Informs the receiver that a tablet-point event has occurred. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/tabletPoint(with:)
func (r_ Responder) TabletPoint(event unsafe.Pointer) {
	sel := objc.RegisterName("tabletPoint:")
	r_.ID.Send(sel, event)
}
// Informs the receiver that a tablet-proximity event has occurred. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/tabletProximity(with:)
func (r_ Responder) TabletProximity(event unsafe.Pointer) {
	sel := objc.RegisterName("tabletProximity:")
	r_.ID.Send(sel, event)
}
// Informs the receiver that new set of touches has been recognized. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/touchesBegan(with:)
func (r_ Responder) TouchesBeganWithEvent(event unsafe.Pointer) {
	sel := objc.RegisterName("touchesBeganWithEvent:")
	r_.ID.Send(sel, event)
}
// Informs the receiver that tracking of touches has been cancelled for any reason. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/touchesCancelled(with:)
func (r_ Responder) TouchesCancelledWithEvent(event unsafe.Pointer) {
	sel := objc.RegisterName("touchesCancelledWithEvent:")
	r_.ID.Send(sel, event)
}
// Returns that a set of touches have been removed. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/touchesEnded(with:)
func (r_ Responder) TouchesEndedWithEvent(event unsafe.Pointer) {
	sel := objc.RegisterName("touchesEndedWithEvent:")
	r_.ID.Send(sel, event)
}
// Informs the receiver that one or more touches has moved. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/touchesMoved(with:)
func (r_ Responder) TouchesMovedWithEvent(event unsafe.Pointer) {
	sel := objc.RegisterName("touchesMovedWithEvent:")
	r_.ID.Send(sel, event)
}
// Attempts to perform the method indicated by an action with a specified argument. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/tryToPerform(_:with:)
func (r_ Responder) TryToPerformWith(action objc.SEL, object objc.ID) bool {
	sel := objc.RegisterName("tryToPerform:with:")
	ret := r_.ID.Send(sel, action, object)
	return ret != 0
}
// Updates the state of the given user activity. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/updateUserActivityState(_:)
func (r_ Responder) UpdateUserActivityState(userActivity unsafe.Pointer) {
	sel := objc.RegisterName("updateUserActivityState:")
	r_.ID.Send(sel, userActivity)
}
// Overridden by subclasses to determine what services are available. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/validRequestor(forSendType:returnType:)
func (r_ Responder) ValidRequestorForSendTypeReturnType(sendType unsafe.Pointer, returnType unsafe.Pointer) objc.ID {
	sel := objc.RegisterName("validRequestorForSendType:returnType:")
	ret := r_.ID.Send(sel, sendType, returnType)
	return ret
}
// Allows controls to determine when they should become first responder. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/validateProposedFirstResponder(_:for:)
func (r_ Responder) ValidateProposedFirstResponderForEvent(responder unsafe.Pointer, event unsafe.Pointer) bool {
	sel := objc.RegisterName("validateProposedFirstResponder:forEvent:")
	ret := r_.ID.Send(sel, responder, event)
	return ret != 0
}
// Returns whether to forward elastic scrolling gesture events up the responder. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/wantsForwardedScrollEvents(for:)
func (r_ Responder) WantsForwardedScrollEventsForAxis(axis unsafe.Pointer) bool {
	sel := objc.RegisterName("wantsForwardedScrollEventsForAxis:")
	ret := r_.ID.Send(sel, axis)
	return ret != 0
}
// Implement this method to track gesture scroll events such as a swipe. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/wantsScrollEventsForSwipeTracking(on:)
func (r_ Responder) WantsScrollEventsForSwipeTrackingOnAxis(axis unsafe.Pointer) bool {
	sel := objc.RegisterName("wantsScrollEventsForSwipeTrackingOnAxis:")
	ret := r_.ID.Send(sel, axis)
	return ret != 0
}
// Returns a custom version of the supplied error object that’s more suitable for presentation in alert sheets and dialogs. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSResponder/willPresentError(_:)
func (r_ Responder) WillPresentError(error unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("willPresentError:")
	ret := r_.ID.Send(sel, error)
	return unsafe.Pointer(ret)
}

