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
	AcceptsFirstResponder() bool
	SetAcceptsFirstResponder(value bool)
	Menu() IMenu
	SetMenu(value IMenu)
	NextResponder() IResponder
	SetNextResponder(value IResponder)
	TouchBar() ITouchBar
	SetTouchBar(value ITouchBar)
	UndoManager() foundation.UndoManager
	SetUndoManager(value foundation.UndoManager)
	UserActivity() foundation.UserActivity
	SetUserActivity(value foundation.UserActivity)
	FlagsChanged(event IEvent)
	KeyDown(event IEvent)
	KeyUp(event IEvent)
	MouseDown(event IEvent)
	MouseDragged(event IEvent)
	MouseEntered(event IEvent)
	MouseExited(event IEvent)
	MouseMoved(event IEvent)
	MouseUp(event IEvent)
	OtherMouseDown(event IEvent)
	OtherMouseDragged(event IEvent)
	OtherMouseUp(event IEvent)
	PresentError(error_ foundation.Error) bool
	PresentErrorModalForWindowDelegateDidPresentSelectorContextInfo(error_ foundation.Error, window IWindow, delegate objectivec.IObject, didPresentSelector objc.SEL, contextInfo unsafe.Pointer)
	RightMouseDown(event IEvent)
	RightMouseDragged(event IEvent)
	RightMouseUp(event IEvent)
	ScrollWheel(event IEvent)
	SupplementalTargetForActionSender(action objc.SEL, sender objectivec.IObject) objc.ID
	TabletPoint(event IEvent)
	TabletProximity(event IEvent)
	UpdateUserActivityState(userActivity foundation.UserActivity)
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



// Informs the receiver that the user has pressed or released a modifier key (Shift, Control, and so on).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/flagsChanged(with:)
func (r_ Responder) FlagsChanged(event IEvent) {
	objc.Send[objc.ID](r_.ID, objc.Sel("flagsChanged:"), event)
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


// Presents an error alert to the user as an application-modal dialog.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/presentError(_:)
func (r_ Responder) PresentError(error_ foundation.Error) bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("presentError:"), error_)
	return rv
}


// Presents an error alert to the user as a document-modal sheet attached to document window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/presentError(_:modalFor:delegate:didPresent:contextInfo:)
func (r_ Responder) PresentErrorModalForWindowDelegateDidPresentSelectorContextInfo(error_ foundation.Error, window IWindow, delegate objectivec.IObject, didPresentSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("presentError:modalForWindow:delegate:didPresentSelector:contextInfo:"), error_, window, delegate, didPresentSelector, contextInfo)
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


// Informs the receiver that the mouse’s scroll wheel has moved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/scrollWheel(with:)
func (r_ Responder) ScrollWheel(event IEvent) {
	objc.Send[objc.ID](r_.ID, objc.Sel("scrollWheel:"), event)
}


// Finds a target for an action method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/supplementalTarget(forAction:sender:)
func (r_ Responder) SupplementalTargetForActionSender(action objc.SEL, sender objectivec.IObject) objc.ID {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("supplementalTargetForAction:sender:"), action, sender)
	return rv
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


// Updates the state of the given user activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSResponder/updateUserActivityState(_:)
func (r_ Responder) UpdateUserActivityState(userActivity foundation.UserActivity) {
	objc.Send[objc.ID](r_.ID, objc.Sel("updateUserActivityState:"), userActivity)
}


// A Boolean value that indicates whether the responder accepts first responder status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/acceptsfirstresponder
func (r_ Responder) AcceptsFirstResponder() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("acceptsFirstResponder"))
	return rv
}


// A Boolean value that indicates whether the responder accepts first responder status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/acceptsfirstresponder
func (r_ Responder) SetAcceptsFirstResponder(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setAcceptsFirstResponder:"), value)
}


// Returns the responder’s menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/menu
func (r_ Responder) Menu() IMenu {
	rv := objc.Send[Menu](r_.ID, objc.Sel("menu"))
	return rv
}


// Returns the responder’s menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/menu
func (r_ Responder) SetMenu(value IMenu) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setMenu:"), value)
}


// The next responder after this one, or
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/nextresponder
func (r_ Responder) NextResponder() IResponder {
	rv := objc.Send[Responder](r_.ID, objc.Sel("nextResponder"))
	return rv
}


// The next responder after this one, or
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/nextresponder
func (r_ Responder) SetNextResponder(value IResponder) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setNextResponder:"), value)
}


// The
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/touchbar
func (r_ Responder) TouchBar() ITouchBar {
	rv := objc.Send[TouchBar](r_.ID, objc.Sel("touchBar"))
	return rv
}


// The
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/touchbar
func (r_ Responder) SetTouchBar(value ITouchBar) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setTouchBar:"), value)
}


// The undo manager for this responder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/undomanager
func (r_ Responder) UndoManager() foundation.UndoManager {
	rv := objc.Send[foundation.UndoManager](r_.ID, objc.Sel("undoManager"))
	return rv
}


// The undo manager for this responder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/undomanager
func (r_ Responder) SetUndoManager(value foundation.UndoManager) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setUndoManager:"), value)
}


// An object encapsulating a user activity supported by this responder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/useractivity
func (r_ Responder) UserActivity() foundation.UserActivity {
	rv := objc.Send[foundation.UserActivity](r_.ID, objc.Sel("userActivity"))
	return rv
}


// An object encapsulating a user activity supported by this responder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/useractivity
func (r_ Responder) SetUserActivity(value foundation.UserActivity) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setUserActivity:"), value)
}



