// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Application] class.
var (
	applicationClass     _ApplicationClass
	applicationClassOnce sync.Once
)

func getApplicationClass() _ApplicationClass {
	applicationClassOnce.Do(func() {
		applicationClass = _ApplicationClass{objc.GetClass("NSApplication")}
	})
	return applicationClass
}

type _ApplicationClass struct {
	class objc.Class
}

// An interface definition for the [Application] class.
type IApplication interface {
	IResponder
	BeginModalSessionForWindow(window unsafe.Pointer) unsafe.Pointer
	DiscardEventsMatchingMaskBeforeEvent(mask unsafe.Pointer, lastEvent unsafe.Pointer)
	NextEventMatchingMaskUntilDateInModeDequeue(mask unsafe.Pointer, expiration unsafe.Pointer, mode unsafe.Pointer, deqFlag bool) unsafe.Pointer
	OrderFrontStandardAboutPanel(sender objc.ID)
	PostEventAtStart(event unsafe.Pointer, atStart bool)
	RegisterForRemoteNotifications()
	ReportException(exception unsafe.Pointer)
	RunModalForWindow(window unsafe.Pointer) unsafe.Pointer
	RunModalSession(session unsafe.Pointer) unsafe.Pointer
	SendActionToFrom(action objc.SEL, target objc.ID, sender objc.ID) bool
	SendEvent(event unsafe.Pointer)
	SetWindowsNeedUpdate(needUpdate bool)
	ToggleTouchBarCustomizationPalette(sender objc.ID)
	UpdateWindows()
	ValidRequestorForSendTypeReturnType(sendType unsafe.Pointer, returnType unsafe.Pointer) objc.ID
}

// An object that manages an app’s main event loop and resources used by all of that app’s objects.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication
type Application struct {
	Responder
}

// ApplicationFrom constructs a [Application] from an unsafe.Pointer.
//
// An object that manages an app’s main event loop and resources used by all of that app’s objects.
func ApplicationFrom(ptr unsafe.Pointer) Application {
	return Application{
		Responder: ResponderFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _ApplicationClass) Alloc() Application {
	rv := objc.Send[Application](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _ApplicationClass) New() Application {
	rv := objc.Send[Application](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ Application) Init() Application {
	rv := objc.Send[Application](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ Application) Autorelease() Application {
	rv := objc.Send[Application](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewApplication creates a new Application instance.
func NewApplication() Application {
	return getApplicationClass().New()
}


// Sets up a modal session with the given window and returns a pointer to the structure representing the session.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/beginModalSession(for:)
func (a_ Application) BeginModalSessionForWindow(window unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("beginModalSessionForWindow:"), window)
	return rv
}

// Removes all events matching the given mask and generated before the specified event.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/discardEvents(matching:before:)
func (a_ Application) DiscardEventsMatchingMaskBeforeEvent(mask unsafe.Pointer, lastEvent unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("discardEventsMatchingMask:beforeEvent:"), mask, lastEvent)
}

// Returns the next event matching a given mask, or if no such event is found before a specified expiration date.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/nextEvent(matching:until:inMode:dequeue:)
func (a_ Application) NextEventMatchingMaskUntilDateInModeDequeue(mask unsafe.Pointer, expiration unsafe.Pointer, mode unsafe.Pointer, deqFlag bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("nextEventMatchingMask:untilDate:inMode:dequeue:"), mask, expiration, mode, deqFlag)
	return rv
}

// Displays a standard About window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/orderFrontStandardAboutPanel(_:)
func (a_ Application) OrderFrontStandardAboutPanel(sender objc.ID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("orderFrontStandardAboutPanel:"), sender)
}

// Adds a given event to the receiver’s event queue.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/postEvent(_:atStart:)
func (a_ Application) PostEventAtStart(event unsafe.Pointer, atStart bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("postEvent:atStart:"), event, atStart)
}

// Register for notifications sent by Apple Push Notification service (APNs).
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/registerForRemoteNotifications()
func (a_ Application) RegisterForRemoteNotifications() {
	objc.Send[objc.ID](a_.ID, objc.Sel("registerForRemoteNotifications"))
}

// Logs a given exception by calling .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/reportException(_:)
func (a_ Application) ReportException(exception unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("reportException:"), exception)
}

// Starts a modal event loop for the specified window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/runModal(for:)
func (a_ Application) RunModalForWindow(window unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("runModalForWindow:"), window)
	return rv
}

// Runs a given modal session, as defined in a previous invocation of .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/runModalSession(_:)
func (a_ Application) RunModalSession(session unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("runModalSession:"), session)
	return rv
}

// Sends the given action message to the given target.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/sendAction(_:to:from:)
func (a_ Application) SendActionToFrom(action objc.SEL, target objc.ID, sender objc.ID) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("sendAction:to:from:"), action, target, sender)
	return rv
}

// Dispatches an event to other objects.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/sendEvent(_:)
func (a_ Application) SendEvent(event unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("sendEvent:"), event)
}

// Sets whether the receiver’s windows need updating when the receiver has finished processing the current event.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/setWindowsNeedUpdate(_:)
func (a_ Application) SetWindowsNeedUpdate(needUpdate bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setWindowsNeedUpdate:"), needUpdate)
}

// Show or hides the interface for customizing the Touch Bar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/toggleTouchBarCustomizationPalette(_:)
func (a_ Application) ToggleTouchBarCustomizationPalette(sender objc.ID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("toggleTouchBarCustomizationPalette:"), sender)
}

// Sends an message to each onscreen window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/updateWindows()
func (a_ Application) UpdateWindows() {
	objc.Send[objc.ID](a_.ID, objc.Sel("updateWindows"))
}

// Indicates whether the receiver can send and receive the specified pasteboard types.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/validRequestor(forSendType:returnType:)
func (a_ Application) ValidRequestorForSendTypeReturnType(sendType unsafe.Pointer, returnType unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("validRequestorForSendType:returnType:"), sendType, returnType)
	return rv
}



