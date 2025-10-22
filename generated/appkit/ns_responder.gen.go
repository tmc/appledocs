// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Responder] class.
var (
	ResponderClass     _ResponderClass
	ResponderClassOnce sync.Once
)

func getResponderClass() _ResponderClass {
	ResponderClassOnce.Do(func() {
		ResponderClass = _ResponderClass{objc.GetClass("NSResponder")}
	})
	return ResponderClass
}

type _ResponderClass struct {
	class objc.Class
}

// An interface definition for the [Responder] class.
type IResponder interface {
	objectivec.IObject
	BecomeFirstResponder() bool
	BeginGestureWithEvent(event IEvent)
	ChangeModeWithEvent(event IEvent)
	ContextMenuKeyDown(event IEvent)
	CursorUpdate(event IEvent)
	EncodeRestorableStateWithCoder(coder foundation.ICoder)
	EncodeRestorableStateWithCoderBackgroundQueue(coder foundation.ICoder, queue foundation.IOperationQueue)
	EndGestureWithEvent(event IEvent)
	FlagsChanged(event IEvent)
	FlushBufferedKeyEvents()
	HelpRequested(eventPtr IEvent)
	InterfaceStyle() unsafe.Pointer
	InterpretKeyEvents(eventArray []Event)
	InvalidateRestorableState()
	KeyDown(event IEvent)
	KeyUp(event IEvent)
	MagnifyWithEvent(event IEvent)
	MakeTouchBar() TouchBar
	MouseCancelled(event IEvent)
	MouseDown(event IEvent)
	MouseDragged(event IEvent)
	MouseEntered(event IEvent)
	MouseExited(event IEvent)
	MouseMoved(event IEvent)
	MouseUp(event IEvent)
	NewWindowForTab(sender objectivec.IObject)
	NoResponderFor(eventSelector objc.SEL)
	OtherMouseDown(event IEvent)
	OtherMouseDragged(event IEvent)
	OtherMouseUp(event IEvent)
	PerformKeyEquivalent(event IEvent) bool
	PerformMnemonic(string_ string) bool
	PerformTextFinderAction(sender objectivec.IObject)
	PresentError(error_ foundation.IError) bool
	PresentErrorModalForWindowDelegateDidPresentSelectorContextInfo(error_ foundation.IError, window IWindow, delegate objectivec.IObject, didPresentSelector objc.SEL, contextInfo unsafe.Pointer)
	PressureChangeWithEvent(event IEvent)
	QuickLookWithEvent(event IEvent)
	ResignFirstResponder() bool
	RestoreStateWithCoder(coder foundation.ICoder)
	RightMouseDown(event IEvent)
	RightMouseDragged(event IEvent)
	RightMouseUp(event IEvent)
	RotateWithEvent(event IEvent)
	ScrollWheel(event IEvent)
	SetInterfaceStyle(interfaceStyle unsafe.Pointer)
	ShouldBeTreatedAsInkEvent(event IEvent) bool
	ShowWritingTools(sender objectivec.IObject)
	SmartMagnifyWithEvent(event IEvent)
	SupplementalTargetForActionSender(action objc.SEL, sender objectivec.IObject) objc.ID
	SwipeWithEvent(event IEvent)
	TabletPoint(event IEvent)
	TabletProximity(event IEvent)
	TouchesBeganWithEvent(event IEvent)
	TouchesCancelledWithEvent(event IEvent)
	TouchesEndedWithEvent(event IEvent)
	TouchesMovedWithEvent(event IEvent)
	TryToPerformWith(action objc.SEL, object objectivec.IObject) bool
	UpdateUserActivityState(userActivity foundation.IUserActivity)
	ValidRequestorForSendTypeReturnType(sendType PasteboardType, returnType PasteboardType) objc.ID
	ValidateProposedFirstResponderForEvent(responder IResponder, event IEvent) bool
	WantsForwardedScrollEventsForAxis(axis IEventGestureAxis) bool
	WantsScrollEventsForSwipeTrackingOnAxis(axis IEventGestureAxis) bool
	WillPresentError(error_ foundation.IError) foundation.Error
	AcceptsFirstResponder() bool
	Menu() NSMenu
	SetMenu(value IMenu)
	NextResponder() NSResponder
	SetNextResponder(value IResponder)
	TouchBar() NSTouchBar
	SetTouchBar(value ITouchBar)
	UndoManager() foundation.UndoManager
	UserActivity() foundation.UserActivity
	SetUserActivity(value foundation.IUserActivity)
}

// An abstract class that forms the basis of event and command processing in AppKit.
//
// The core classes— , , and —inherit from , as must any class that handles events. The responder model uses three components: event messages, action messages, and the responder chain. also plays an important role in the presentation of error information. The default implementations of the and methods send to , thereby giving subclasses the opportunity to customize the localized information presented in error alerts. then forwards the message to the next responder, passing it the customized object. The exact path up the modified responder chain depends on the type of application window: Window that the document owns: view > superviews > window > window controller > document object > document controller > the application object Window with window controller but no documents: view > superviews > window > window controller > the application object Window with no window controllers: view > superviews > window > the application object displays a document-modal error alert and, if the error object has a recovery attempter, gives it a chance to recover from the error. A recovery attempter is an object that conforms to the informal protocol.


// An abstract class that forms the basis of event and command processing in AppKit.
//
// [Full Topic]
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getResponderClass().New()
}




// Creates a new responder object with data in an unarchiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/init(coder:)

func NewResponderWithCoder(coder foundation.ICoder) Responder {
	instance := getResponderClass().Alloc()
	rv := objc.Send[Responder](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}



// Returns the classes that support secure coding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/allowedClasses(forRestorableStateKeyPath:)

func (rc _ResponderClass) AllowedClassesForRestorableStateKeyPath(keyPath string) []objc.Class {
	rv := objc.Send[[]objc.Class](objc.ID(rc.class), objc.Sel("allowedClassesForRestorableStateKeyPath:"), objc.String(keyPath))
	return rv
}


// Returns an array of key paths representing the restorable attributes of the responder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/restorableStateKeyPaths

func (rc _ResponderClass) RestorableStateKeyPaths() []string {
	rv := objc.Send[[]string](objc.ID(rc.class), objc.Sel("restorableStateKeyPaths"))
	return rv
}


// Notifies the receiver that it’s about to become first responder in its .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/becomeFirstResponder()

func (r_ Responder) BecomeFirstResponder() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("becomeFirstResponder"))
	return rv
}



// Informs the receiver that the user has begun a touch gesture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/beginGesture(with:)

func (r_ Responder) BeginGestureWithEvent(event IEvent) {
	objc.Send[objc.ID](r_.ID, objc.Sel("beginGestureWithEvent:"), event)
}



// Informs the responder that performed a double-tap on the side of an Apple Pencil.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/changeMode(with:)

func (r_ Responder) ChangeModeWithEvent(event IEvent) {
	objc.Send[objc.ID](r_.ID, objc.Sel("changeModeWithEvent:"), event)
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/contextMenuKeyDown(_:)

func (r_ Responder) ContextMenuKeyDown(event IEvent) {
	objc.Send[objc.ID](r_.ID, objc.Sel("contextMenuKeyDown:"), event)
}



// Informs the receiver that the mouse cursor has moved into a cursor rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/cursorUpdate(with:)

func (r_ Responder) CursorUpdate(event IEvent) {
	objc.Send[objc.ID](r_.ID, objc.Sel("cursorUpdate:"), event)
}



// Saves the interface-related state of the responder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/encodeRestorableState(with:)

func (r_ Responder) EncodeRestorableStateWithCoder(coder foundation.ICoder) {
	objc.Send[objc.ID](r_.ID, objc.Sel("encodeRestorableStateWithCoder:"), coder)
}



// Saves the interface-related state of the responder to a keyed archiver either synchronously or asynchronously on the given operation queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/encodeRestorableState(with:backgroundQueue:)

func (r_ Responder) EncodeRestorableStateWithCoderBackgroundQueue(coder foundation.ICoder, queue foundation.IOperationQueue) {
	objc.Send[objc.ID](r_.ID, objc.Sel("encodeRestorableStateWithCoder:backgroundQueue:"), coder, queue)
}



// Informs the receiver that the user has ended a touch gesture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/endGesture(with:)

func (r_ Responder) EndGestureWithEvent(event IEvent) {
	objc.Send[objc.ID](r_.ID, objc.Sel("endGestureWithEvent:"), event)
}



// Informs the receiver that the user has pressed or released a modifier key (Shift, Control, and so on).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/flagsChanged(with:)

func (r_ Responder) FlagsChanged(event IEvent) {
	objc.Send[objc.ID](r_.ID, objc.Sel("flagsChanged:"), event)
}



// Clears any unprocessed key events when overridden by subclasses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/flushBufferedKeyEvents()

func (r_ Responder) FlushBufferedKeyEvents() {
	objc.Send[objc.ID](r_.ID, objc.Sel("flushBufferedKeyEvents"))
}



// Displays context-sensitive help for the receiver if help has been registered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/helpRequested(_:)

func (r_ Responder) HelpRequested(eventPtr IEvent) {
	objc.Send[objc.ID](r_.ID, objc.Sel("helpRequested:"), eventPtr)
}



// Returns the receiver’s interface style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/interfaceStyle

func (r_ Responder) InterfaceStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("interfaceStyle"))
	return rv
}



// Handles a series of key events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/interpretKeyEvents(_:)

func (r_ Responder) InterpretKeyEvents(eventArray []Event) {
	objc.Send[objc.ID](r_.ID, objc.Sel("interpretKeyEvents:"), eventArray)
}



// Marks the responder’s interface-related state as dirty.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/invalidateRestorableState()

func (r_ Responder) InvalidateRestorableState() {
	objc.Send[objc.ID](r_.ID, objc.Sel("invalidateRestorableState"))
}



// Informs the receiver that the user has pressed a key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/keyDown(with:)

func (r_ Responder) KeyDown(event IEvent) {
	objc.Send[objc.ID](r_.ID, objc.Sel("keyDown:"), event)
}



// Informs the receiver that the user has released a key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/keyUp(with:)

func (r_ Responder) KeyUp(event IEvent) {
	objc.Send[objc.ID](r_.ID, objc.Sel("keyUp:"), event)
}



// Informs the receiver that the user has begun a pinch gesture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/magnify(with:)

func (r_ Responder) MagnifyWithEvent(event IEvent) {
	objc.Send[objc.ID](r_.ID, objc.Sel("magnifyWithEvent:"), event)
}



// Your custom subclass of the class should override this method to create and configure your subclass’s default object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/makeTouchBar()

func (r_ Responder) MakeTouchBar() TouchBar {
	rv := objc.Send[TouchBar](r_.ID, objc.Sel("makeTouchBar"))
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/mouseCancelled(with:)

func (r_ Responder) MouseCancelled(event IEvent) {
	objc.Send[objc.ID](r_.ID, objc.Sel("mouseCancelled:"), event)
}



// Informs the receiver that the user has pressed the left mouse button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/mouseDown(with:)

func (r_ Responder) MouseDown(event IEvent) {
	objc.Send[objc.ID](r_.ID, objc.Sel("mouseDown:"), event)
}



// Informs the receiver that the user has moved the mouse with the left button pressed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/mouseDragged(with:)

func (r_ Responder) MouseDragged(event IEvent) {
	objc.Send[objc.ID](r_.ID, objc.Sel("mouseDragged:"), event)
}



// Informs the receiver that the cursor has entered a tracking rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/mouseEntered(with:)

func (r_ Responder) MouseEntered(event IEvent) {
	objc.Send[objc.ID](r_.ID, objc.Sel("mouseEntered:"), event)
}



// Informs the receiver that the cursor has exited a tracking rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/mouseExited(with:)

func (r_ Responder) MouseExited(event IEvent) {
	objc.Send[objc.ID](r_.ID, objc.Sel("mouseExited:"), event)
}



// Informs the receiver that the mouse has moved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/mouseMoved(with:)

func (r_ Responder) MouseMoved(event IEvent) {
	objc.Send[objc.ID](r_.ID, objc.Sel("mouseMoved:"), event)
}



// Informs the receiver that the user has released the left mouse button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/mouseUp(with:)

func (r_ Responder) MouseUp(event IEvent) {
	objc.Send[objc.ID](r_.ID, objc.Sel("mouseUp:"), event)
}



// Creates a new window to show as a tab in a tabbed window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/newWindowForTab(_:)

func (r_ Responder) NewWindowForTab(sender objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("newWindowForTab:"), sender)
}



// Handles the case where an event or action message falls off the end of the responder chain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/noResponder(for:)

func (r_ Responder) NoResponderFor(eventSelector objc.SEL) {
	objc.Send[objc.ID](r_.ID, objc.Sel("noResponderFor:"), eventSelector)
}



// Informs the receiver that the user has pressed a mouse button other than the left or right one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/otherMouseDown(with:)

func (r_ Responder) OtherMouseDown(event IEvent) {
	objc.Send[objc.ID](r_.ID, objc.Sel("otherMouseDown:"), event)
}



// Informs the receiver that the user has moved the mouse with a button other than the left or right button pressed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/otherMouseDragged(with:)

func (r_ Responder) OtherMouseDragged(event IEvent) {
	objc.Send[objc.ID](r_.ID, objc.Sel("otherMouseDragged:"), event)
}



// Informs the receiver that the user has released a mouse button other than the left or right button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/otherMouseUp(with:)

func (r_ Responder) OtherMouseUp(event IEvent) {
	objc.Send[objc.ID](r_.ID, objc.Sel("otherMouseUp:"), event)
}



// Handle a key equivalent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/performKeyEquivalent(with:)

func (r_ Responder) PerformKeyEquivalent(event IEvent) bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("performKeyEquivalent:"), event)
	return rv
}



// Handle a mnemonic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/performMnemonic:

func (r_ Responder) PerformMnemonic(string_ string) bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("performMnemonic:"), objc.String(string_))
	return rv
}



// Performs all find oriented actions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/performTextFinderAction(_:)

func (r_ Responder) PerformTextFinderAction(sender objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("performTextFinderAction:"), sender)
}



// Presents an error alert to the user as an application-modal dialog.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/presentError(_:)

func (r_ Responder) PresentError(error_ foundation.IError) bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("presentError:"), error_)
	return rv
}



// Presents an error alert to the user as a document-modal sheet attached to document window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/presentError(_:modalFor:delegate:didPresent:contextInfo:)

func (r_ Responder) PresentErrorModalForWindowDelegateDidPresentSelectorContextInfo(error_ foundation.IError, window IWindow, delegate objectivec.IObject, didPresentSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("presentError:modalForWindow:delegate:didPresentSelector:contextInfo:"), error_, window, delegate, didPresentSelector, contextInfo)
}



// Indicates a pressure change as the result of a user input event on a system that supports pressure sensitivity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/pressureChange(with:)

func (r_ Responder) PressureChangeWithEvent(event IEvent) {
	objc.Send[objc.ID](r_.ID, objc.Sel("pressureChangeWithEvent:"), event)
}



// Performs a Quick Look on the content at the location specified by the supplied event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/quickLook(with:)

func (r_ Responder) QuickLookWithEvent(event IEvent) {
	objc.Send[objc.ID](r_.ID, objc.Sel("quickLookWithEvent:"), event)
}



// Notifies the receiver that it’s been asked to relinquish its status as first responder in its window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/resignFirstResponder()

func (r_ Responder) ResignFirstResponder() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("resignFirstResponder"))
	return rv
}



// Restores the interface-related state of the responder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/restoreState(with:)

func (r_ Responder) RestoreStateWithCoder(coder foundation.ICoder) {
	objc.Send[objc.ID](r_.ID, objc.Sel("restoreStateWithCoder:"), coder)
}



// Informs the receiver that the user has pressed the right mouse button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/rightMouseDown(with:)

func (r_ Responder) RightMouseDown(event IEvent) {
	objc.Send[objc.ID](r_.ID, objc.Sel("rightMouseDown:"), event)
}



// Informs the receiver that the user has moved the mouse with the right button pressed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/rightMouseDragged(with:)

func (r_ Responder) RightMouseDragged(event IEvent) {
	objc.Send[objc.ID](r_.ID, objc.Sel("rightMouseDragged:"), event)
}



// Informs the receiver that the user has released the right mouse button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/rightMouseUp(with:)

func (r_ Responder) RightMouseUp(event IEvent) {
	objc.Send[objc.ID](r_.ID, objc.Sel("rightMouseUp:"), event)
}



// Informs the receiver that the user has begun a rotation gesture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/rotate(with:)

func (r_ Responder) RotateWithEvent(event IEvent) {
	objc.Send[objc.ID](r_.ID, objc.Sel("rotateWithEvent:"), event)
}



// Informs the receiver that the mouse’s scroll wheel has moved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/scrollWheel(with:)

func (r_ Responder) ScrollWheel(event IEvent) {
	objc.Send[objc.ID](r_.ID, objc.Sel("scrollWheel:"), event)
}



// Sets the receiver’s style to the style specified by , such as or .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/setInterfaceStyle:

func (r_ Responder) SetInterfaceStyle(interfaceStyle unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setInterfaceStyle:"), interfaceStyle)
}



// Indicates whether a pen-down event should be treated as an ink event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/shouldBeTreatedAsInkEvent(_:)

func (r_ Responder) ShouldBeTreatedAsInkEvent(event IEvent) bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("shouldBeTreatedAsInkEvent:"), event)
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/showWritingTools(_:)

func (r_ Responder) ShowWritingTools(sender objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("showWritingTools:"), sender)
}



// Informs the receiver that the user performed a smart zoom gesture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/smartMagnify(with:)

func (r_ Responder) SmartMagnifyWithEvent(event IEvent) {
	objc.Send[objc.ID](r_.ID, objc.Sel("smartMagnifyWithEvent:"), event)
}



// Finds a target for an action method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/supplementalTarget(forAction:sender:)

func (r_ Responder) SupplementalTargetForActionSender(action objc.SEL, sender objectivec.IObject) objc.ID {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("supplementalTargetForAction:sender:"), action, sender)
	return rv
}



// Informs the receiver that the user has begun a swipe gesture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/swipe(with:)

func (r_ Responder) SwipeWithEvent(event IEvent) {
	objc.Send[objc.ID](r_.ID, objc.Sel("swipeWithEvent:"), event)
}



// Informs the receiver that a tablet-point event has occurred.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/tabletPoint(with:)

func (r_ Responder) TabletPoint(event IEvent) {
	objc.Send[objc.ID](r_.ID, objc.Sel("tabletPoint:"), event)
}



// Informs the receiver that a tablet-proximity event has occurred.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/tabletProximity(with:)

func (r_ Responder) TabletProximity(event IEvent) {
	objc.Send[objc.ID](r_.ID, objc.Sel("tabletProximity:"), event)
}



// Informs the receiver that new set of touches has been recognized.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/touchesBegan(with:)

func (r_ Responder) TouchesBeganWithEvent(event IEvent) {
	objc.Send[objc.ID](r_.ID, objc.Sel("touchesBeganWithEvent:"), event)
}



// Informs the receiver that tracking of touches has been cancelled for any reason.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/touchesCancelled(with:)

func (r_ Responder) TouchesCancelledWithEvent(event IEvent) {
	objc.Send[objc.ID](r_.ID, objc.Sel("touchesCancelledWithEvent:"), event)
}



// Returns that a set of touches have been removed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/touchesEnded(with:)

func (r_ Responder) TouchesEndedWithEvent(event IEvent) {
	objc.Send[objc.ID](r_.ID, objc.Sel("touchesEndedWithEvent:"), event)
}



// Informs the receiver that one or more touches has moved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/touchesMoved(with:)

func (r_ Responder) TouchesMovedWithEvent(event IEvent) {
	objc.Send[objc.ID](r_.ID, objc.Sel("touchesMovedWithEvent:"), event)
}



// Attempts to perform the method indicated by an action with a specified argument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/tryToPerform(_:with:)

func (r_ Responder) TryToPerformWith(action objc.SEL, object objectivec.IObject) bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("tryToPerform:with:"), action, object)
	return rv
}



// Updates the state of the given user activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/updateUserActivityState(_:)

func (r_ Responder) UpdateUserActivityState(userActivity foundation.IUserActivity) {
	objc.Send[objc.ID](r_.ID, objc.Sel("updateUserActivityState:"), userActivity)
}



// Overridden by subclasses to determine what services are available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/validRequestor(forSendType:returnType:)

func (r_ Responder) ValidRequestorForSendTypeReturnType(sendType PasteboardType, returnType PasteboardType) objc.ID {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("validRequestorForSendType:returnType:"), sendType, returnType)
	return rv
}



// Allows controls to determine when they should become first responder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/validateProposedFirstResponder(_:for:)

func (r_ Responder) ValidateProposedFirstResponderForEvent(responder IResponder, event IEvent) bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("validateProposedFirstResponder:forEvent:"), responder, event)
	return rv
}



// Returns whether to forward elastic scrolling gesture events up the responder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/wantsForwardedScrollEvents(for:)

func (r_ Responder) WantsForwardedScrollEventsForAxis(axis IEventGestureAxis) bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("wantsForwardedScrollEventsForAxis:"), axis)
	return rv
}



// Implement this method to track gesture scroll events such as a swipe.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/wantsScrollEventsForSwipeTracking(on:)

func (r_ Responder) WantsScrollEventsForSwipeTrackingOnAxis(axis IEventGestureAxis) bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("wantsScrollEventsForSwipeTrackingOnAxis:"), axis)
	return rv
}



// Returns a custom version of the supplied error object that’s more suitable for presentation in alert sheets and dialogs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/willPresentError(_:)

func (r_ Responder) WillPresentError(error_ foundation.IError) foundation.Error {
	rv := objc.Send[foundation.Error](r_.ID, objc.Sel("willPresentError:"), error_)
	return rv
}


// A Boolean value that indicates whether the responder accepts first responder status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/acceptsFirstResponder

func (r_ Responder) AcceptsFirstResponder() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("acceptsFirstResponder"))
	return rv
}


// Returns the responder’s menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/menu

func (r_ Responder) Menu() NSMenu {
	rv := objc.Send[NSMenu](r_.ID, objc.Sel("menu"))
	return rv
}


// Returns the responder’s menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/menu

func (r_ Responder) SetMenu(value IMenu) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setMenu:"), value)
}


// The next responder after this one, or if it has none.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/nextResponder

func (r_ Responder) NextResponder() NSResponder {
	rv := objc.Send[NSResponder](r_.ID, objc.Sel("nextResponder"))
	return rv
}


// The next responder after this one, or if it has none.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/nextResponder

func (r_ Responder) SetNextResponder(value IResponder) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setNextResponder:"), value)
}


// Returns an array of key paths representing the restorable attributes of the responder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/restorableStateKeyPaths

func (r_ Responder) RestorableStateKeyPaths() []string {
	rv := objc.Send[[]string](r_.ID, objc.Sel("restorableStateKeyPaths"))
	return rv
}


// The object associated with the responder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/touchBar

func (r_ Responder) TouchBar() NSTouchBar {
	rv := objc.Send[NSTouchBar](r_.ID, objc.Sel("touchBar"))
	return rv
}


// The object associated with the responder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/touchBar

func (r_ Responder) SetTouchBar(value ITouchBar) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setTouchBar:"), value)
}


// The undo manager for this responder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/undoManager

func (r_ Responder) UndoManager() foundation.UndoManager {
	rv := objc.Send[foundation.UndoManager](r_.ID, objc.Sel("undoManager"))
	return rv
}


// An object encapsulating a user activity supported by this responder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/userActivity

func (r_ Responder) UserActivity() foundation.UserActivity {
	rv := objc.Send[foundation.UserActivity](r_.ID, objc.Sel("userActivity"))
	return rv
}


// An object encapsulating a user activity supported by this responder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/userActivity

func (r_ Responder) SetUserActivity(value foundation.IUserActivity) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setUserActivity:"), value)
}


