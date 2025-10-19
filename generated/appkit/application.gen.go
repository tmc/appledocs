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
// Every app uses a single instance of to control the main event loop, keep track of the app’s windows and menus, distribute events to the appropriate objects (that’s, itself or one of its windows), set up autorelease pools, and receive notification of app-level events. An object has a delegate (an object that you assign) that’s notified when the app starts or terminates, is hidden or activated, should open a file selected by the user, and so forth. By setting the delegate and implementing the delegate methods, you customize the behavior of your app without having to subclass . In your app’s function, create the instance by calling the class method. After creating the application object, the function should load your app’s main nib file and then start the event loop by sending the application object a message. If you create an Application project in Xcode, this function is created for you. The function Xcode creates begins by calling a function named , which is functionally similar to the following: The class method initializes the display environment and connects your program to the window server and the display server. The object maintains a list of all the objects the app uses, so it can retrieve any of the app’s objects. The method also initializes the global variable , which you use to retrieve the instance. only performs the initialization once. If you invoke it more than once, it returns the application object it created previously. The shared object performs the important task of receiving events from the window server and distributing them to the proper objects. translates an event into an object, then forwards the event object to the affected object. All keyboard and mouse events go directly to the object associated with the event. The only exception to this rule is if the Command key is pressed when a key-down event occurs; in this case, every object has an opportunity to respond to the event. When a window object receives an object from , it distributes it to the objects in its view hierarchy. is also responsible for dispatching certain Apple events received by the app. For example, macOS sends Apple events to your app at various times, such as when the app is launched or reopened. installs Apple event handlers to handle these events by sending a message to the appropriate object. You can also use the class to register your own Apple event handlers. The method is generally the best place to do so. For more information on how events are handled and how you can modify the default behavior, including information on working with Apple events in scriptable apps, see in . The class sets up block during initialization and inside the event loop—specifically, within its initialization (or ) and methods. Similarly, the methods AppKit adds to employ blocks during the loading of nib files. These blocks aren’t accessible outside the scope of the respective and methods. Typically, an app creates objects either while the event loop is running or by loading objects from nib files, so this lack of access usually isn’t a problem. However, if you do need to use Cocoa classes within the function itself (other than to load nib files or to instantiate ), you should create an block to contain the code using the classes.
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

// The appearance associated with the app’s windows.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/appearance
func (a_ Application) Appearance() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("appearance"))
	return rv
}

// SetAppearance sets the value of the appearance property.
// The appearance associated with the app’s windows.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/appearance
func (a_ Application) SetAppearance(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAppearance:"), value)
}
// The last event object that the app retrieved from the event queue.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/currentEvent
func (a_ Application) CurrentEvent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("currentEvent"))
	return rv
}
// The appearance that AppKit uses to draw the app’s interface.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/effectiveAppearance
func (a_ Application) EffectiveAppearance() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("effectiveAppearance"))
	return rv
}
// A Boolean value indicating whether the main menu contains an item for customizing the contents of the Touch Bar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/isAutomaticCustomizeTouchBarMenuItemEnabled
func (a_ Application) AutomaticCustomizeTouchBarMenuItemEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("automaticCustomizeTouchBarMenuItemEnabled"))
	return rv
}

// SetAutomaticCustomizeTouchBarMenuItemEnabled sets the value of the automaticCustomizeTouchBarMenuItemEnabled property.
// A Boolean value indicating whether the main menu contains an item for customizing the contents of the Touch Bar.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/isAutomaticCustomizeTouchBarMenuItemEnabled
func (a_ Application) SetAutomaticCustomizeTouchBarMenuItemEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAutomaticCustomizeTouchBarMenuItemEnabled:"), value)
}


