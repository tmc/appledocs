// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Application] class.
var ApplicationClass objc.Class

func init() {
	ApplicationClass = objc.GetClass("NSApplication")
}

type Application struct {
	objc.ID
}

func ApplicationFrom(ptr unsafe.Pointer) Application {
	return Application{
		ID: objc.ID(ptr),
	}
}


// Sets up a modal session with the given window and returns a pointer to the   structure representing the session. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSApplication/beginModalSession(for:)
func (a_ Application) BeginModalSessionForWindow(window unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("beginModalSessionForWindow:")
	ret := a_.ID.Send(sel, window)
	return unsafe.Pointer(ret)
}
// Removes all events matching the given mask and generated before the specified event. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSApplication/discardEvents(matching:before:)
func (a_ Application) DiscardEventsMatchingMaskBeforeEvent(mask unsafe.Pointer, lastEvent unsafe.Pointer) {
	sel := objc.RegisterName("discardEventsMatchingMask:beforeEvent:")
	a_.ID.Send(sel, mask, lastEvent)
}
// Returns the next event matching a given mask, or   if no such event is found before a specified expiration date. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSApplication/nextEvent(matching:until:inMode:dequeue:)
func (a_ Application) NextEventMatchingMaskUntilDateInModeDequeue(mask unsafe.Pointer, expiration unsafe.Pointer, mode unsafe.Pointer, deqFlag bool) unsafe.Pointer {
	sel := objc.RegisterName("nextEventMatchingMask:untilDate:inMode:dequeue:")
	ret := a_.ID.Send(sel, mask, expiration, mode, deqFlag)
	return unsafe.Pointer(ret)
}
// Displays a standard About window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSApplication/orderFrontStandardAboutPanel(_:)
func (a_ Application) OrderFrontStandardAboutPanel(sender objc.ID) {
	sel := objc.RegisterName("orderFrontStandardAboutPanel:")
	a_.ID.Send(sel, sender)
}
// Adds a given event to the receiver’s event queue. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSApplication/postEvent(_:atStart:)
func (a_ Application) PostEventAtStart(event unsafe.Pointer, atStart bool) {
	sel := objc.RegisterName("postEvent:atStart:")
	a_.ID.Send(sel, event, atStart)
}
// Register for notifications sent by Apple Push Notification service (APNs). [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSApplication/registerForRemoteNotifications()
func (a_ Application) RegisterForRemoteNotifications() {
	sel := objc.RegisterName("registerForRemoteNotifications")
	a_.ID.Send(sel)
}
// Logs a given exception by calling  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSApplication/reportException(_:)
func (a_ Application) ReportException(exception unsafe.Pointer) {
	sel := objc.RegisterName("reportException:")
	a_.ID.Send(sel, exception)
}
// Starts a modal event loop for the specified window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSApplication/runModal(for:)
func (a_ Application) RunModalForWindow(window unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("runModalForWindow:")
	ret := a_.ID.Send(sel, window)
	return unsafe.Pointer(ret)
}
// Runs a given modal session, as defined in a previous invocation of  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSApplication/runModalSession(_:)
func (a_ Application) RunModalSession(session unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("runModalSession:")
	ret := a_.ID.Send(sel, session)
	return unsafe.Pointer(ret)
}
// Sends the given action message to the given target. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSApplication/sendAction(_:to:from:)
func (a_ Application) SendActionToFrom(action objc.SEL, target objc.ID, sender objc.ID) bool {
	sel := objc.RegisterName("sendAction:to:from:")
	ret := a_.ID.Send(sel, action, target, sender)
	return ret != 0
}
// Dispatches an event to other objects. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSApplication/sendEvent(_:)
func (a_ Application) SendEvent(event unsafe.Pointer) {
	sel := objc.RegisterName("sendEvent:")
	a_.ID.Send(sel, event)
}
// Sets whether the receiver’s windows need updating when the receiver has finished processing the current event. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSApplication/setWindowsNeedUpdate(_:)
func (a_ Application) SetWindowsNeedUpdate(needUpdate bool) {
	sel := objc.RegisterName("setWindowsNeedUpdate:")
	a_.ID.Send(sel, needUpdate)
}
// Show or hides the interface for customizing the Touch Bar. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSApplication/toggleTouchBarCustomizationPalette(_:)
func (a_ Application) ToggleTouchBarCustomizationPalette(sender objc.ID) {
	sel := objc.RegisterName("toggleTouchBarCustomizationPalette:")
	a_.ID.Send(sel, sender)
}
// Sends an   message to each onscreen window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSApplication/updateWindows()
func (a_ Application) UpdateWindows() {
	sel := objc.RegisterName("updateWindows")
	a_.ID.Send(sel)
}
// Indicates whether the receiver can send and receive the specified pasteboard types. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSApplication/validRequestor(forSendType:returnType:)
func (a_ Application) ValidRequestorForSendTypeReturnType(sendType unsafe.Pointer, returnType unsafe.Pointer) objc.ID {
	sel := objc.RegisterName("validRequestorForSendType:returnType:")
	ret := a_.ID.Send(sel, sendType, returnType)
	return ret
}

