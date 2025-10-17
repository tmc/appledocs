
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Application] class.
var ApplicationClass _ApplicationClass

func init() {
	ApplicationClass = _ApplicationClass{objc.GetClass("NSApplication")}
}

type _ApplicationClass struct {
	objc.Class
}

// An interface definition for the [Application] class.
type IApplication interface {
	ID() objc.ID
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

type Application struct {
	id objc.ID
}

func ApplicationFrom(ptr unsafe.Pointer) Application {
	return Application{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ Application) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _ApplicationClass) Alloc() Application {
	rv := objc.Send[Application](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _ApplicationClass) New() Application {
	rv := objc.Send[Application](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewApplication creates and returns a new initialized instance.
func NewApplication() Application {
	return ApplicationClass.New()
}

// Init initializes the instance.
func (a_ Application) Init() Application {
	rv := objc.Send[Application](a_.ID(), selInit)
	return rv
}
// Sets up a modal session with the given window and returns a pointer to the   structure representing the session. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSApplication/beginModalSession(for:)
func (a_ Application) BeginModalSessionForWindow(window unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID(), objc.RegisterName("beginModalSessionForWindow:"), window)
	return rv
}
// Removes all events matching the given mask and generated before the specified event. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSApplication/discardEvents(matching:before:)
func (a_ Application) DiscardEventsMatchingMaskBeforeEvent(mask unsafe.Pointer, lastEvent unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID(), objc.RegisterName("discardEventsMatchingMask:beforeEvent:"), mask, lastEvent)
}
// Returns the next event matching a given mask, or   if no such event is found before a specified expiration date. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSApplication/nextEvent(matching:until:inMode:dequeue:)
func (a_ Application) NextEventMatchingMaskUntilDateInModeDequeue(mask unsafe.Pointer, expiration unsafe.Pointer, mode unsafe.Pointer, deqFlag bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID(), objc.RegisterName("nextEventMatchingMask:untilDate:inMode:dequeue:"), mask, expiration, mode, deqFlag)
	return rv
}
// Displays a standard About window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSApplication/orderFrontStandardAboutPanel(_:)
func (a_ Application) OrderFrontStandardAboutPanel(sender objc.ID) {
	objc.Send[objc.ID](a_.ID(), objc.RegisterName("orderFrontStandardAboutPanel:"), sender)
}
// Adds a given event to the receiver’s event queue. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSApplication/postEvent(_:atStart:)
func (a_ Application) PostEventAtStart(event unsafe.Pointer, atStart bool) {
	objc.Send[objc.ID](a_.ID(), objc.RegisterName("postEvent:atStart:"), event, atStart)
}
// Register for notifications sent by Apple Push Notification service (APNs). [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSApplication/registerForRemoteNotifications()
func (a_ Application) RegisterForRemoteNotifications() {
	objc.Send[objc.ID](a_.ID(), objc.RegisterName("registerForRemoteNotifications"))
}
// Logs a given exception by calling  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSApplication/reportException(_:)
func (a_ Application) ReportException(exception unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID(), objc.RegisterName("reportException:"), exception)
}
// Starts a modal event loop for the specified window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSApplication/runModal(for:)
func (a_ Application) RunModalForWindow(window unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID(), objc.RegisterName("runModalForWindow:"), window)
	return rv
}
// Runs a given modal session, as defined in a previous invocation of  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSApplication/runModalSession(_:)
func (a_ Application) RunModalSession(session unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID(), objc.RegisterName("runModalSession:"), session)
	return rv
}
// Sends the given action message to the given target. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSApplication/sendAction(_:to:from:)
func (a_ Application) SendActionToFrom(action objc.SEL, target objc.ID, sender objc.ID) bool {
	rv := objc.Send[bool](a_.ID(), objc.RegisterName("sendAction:to:from:"), action, target, sender)
	return rv
}
// Dispatches an event to other objects. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSApplication/sendEvent(_:)
func (a_ Application) SendEvent(event unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID(), objc.RegisterName("sendEvent:"), event)
}
// Sets whether the receiver’s windows need updating when the receiver has finished processing the current event. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSApplication/setWindowsNeedUpdate(_:)
func (a_ Application) SetWindowsNeedUpdate(needUpdate bool) {
	objc.Send[objc.ID](a_.ID(), objc.RegisterName("setWindowsNeedUpdate:"), needUpdate)
}
// Show or hides the interface for customizing the Touch Bar. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSApplication/toggleTouchBarCustomizationPalette(_:)
func (a_ Application) ToggleTouchBarCustomizationPalette(sender objc.ID) {
	objc.Send[objc.ID](a_.ID(), objc.RegisterName("toggleTouchBarCustomizationPalette:"), sender)
}
// Sends an   message to each onscreen window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSApplication/updateWindows()
func (a_ Application) UpdateWindows() {
	objc.Send[objc.ID](a_.ID(), objc.RegisterName("updateWindows"))
}
// Indicates whether the receiver can send and receive the specified pasteboard types. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSApplication/validRequestor(forSendType:returnType:)
func (a_ Application) ValidRequestorForSendTypeReturnType(sendType unsafe.Pointer, returnType unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](a_.ID(), objc.RegisterName("validRequestorForSendType:returnType:"), sendType, returnType)
	return rv
}
// The appearance associated with the app’s windows. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSApplication/appearance
func (a_ Application) Appearance() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID(), objc.RegisterName("appearance"))
	return rv
}
// SetAppearance sets the value of the appearance property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSApplication/appearance
func (a_ Application) SetAppearance(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID(), objc.RegisterName("setAppearance:"), value)
}
// The last event object that the app retrieved from the event queue. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSApplication/currentEvent
func (a_ Application) CurrentEvent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID(), objc.RegisterName("currentEvent"))
	return rv
}
// The appearance that AppKit uses to draw the app’s interface. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSApplication/effectiveAppearance
func (a_ Application) EffectiveAppearance() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID(), objc.RegisterName("effectiveAppearance"))
	return rv
}
// A Boolean value indicating whether the main menu contains an item for customizing the contents of the Touch Bar. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSApplication/isAutomaticCustomizeTouchBarMenuItemEnabled
func (a_ Application) AutomaticCustomizeTouchBarMenuItemEnabled() bool {
	rv := objc.Send[bool](a_.ID(), objc.RegisterName("automaticCustomizeTouchBarMenuItemEnabled"))
	return rv
}
// SetAutomaticCustomizeTouchBarMenuItemEnabled sets the value of the automaticCustomizeTouchBarMenuItemEnabled property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSApplication/isAutomaticCustomizeTouchBarMenuItemEnabled
func (a_ Application) SetAutomaticCustomizeTouchBarMenuItemEnabled(value bool) {
	objc.Send[objc.ID](a_.ID(), objc.RegisterName("setAutomaticCustomizeTouchBarMenuItemEnabled:"), value)
}
