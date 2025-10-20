// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Application] class.
var (
	ApplicationClass     _ApplicationClass
	ApplicationClassOnce sync.Once
)

func getApplicationClass() _ApplicationClass {
	ApplicationClassOnce.Do(func() {
		ApplicationClass = _ApplicationClass{objc.GetClass("NSApplication")}
	})
	return ApplicationClass
}

type _ApplicationClass struct {
	class objc.Class
}

// An interface definition for the [Application] class.
type IApplication interface {
	IResponder
	ActivateIgnoringOtherApps(ignoreOtherApps bool)
	ActivationPolicy() unsafe.Pointer
	ApplicationPrintFiles(sender unsafe.Pointer, filenames unsafe.Pointer)
	BeginModalSessionForWindow(window unsafe.Pointer) unsafe.Pointer
	BeginSheetModalForWindowModalDelegateDidEndSelectorContextInfo(sheet unsafe.Pointer, docWindow unsafe.Pointer, modalDelegate objc.ID, didEndSelector objc.SEL, contextInfo unsafe.Pointer)
	DiscardEventsMatchingMaskBeforeEvent(mask unsafe.Pointer, lastEvent unsafe.Pointer)
	EndModalSession(session unsafe.Pointer)
	EndSheet(sheet unsafe.Pointer)
	EndSheetReturnCode(sheet unsafe.Pointer, returnCode int)
	FinishLaunching()
	Hide(sender objc.ID)
	NextEventMatchingMaskUntilDateInModeDequeue(mask unsafe.Pointer, expiration unsafe.Pointer, mode unsafe.Pointer, deqFlag bool) unsafe.Pointer
	OrderFrontCharacterPalette(sender objc.ID)
	OrderFrontColorPanel(sender objc.ID)
	OrderFrontStandardAboutPanel(sender objc.ID)
	OrderFrontStandardAboutPanelWithOptions(optionsDictionary unsafe.Pointer)
	PostEventAtStart(event unsafe.Pointer, atStart bool)
	RegisterForRemoteNotifications()
	RegisterForRemoteNotificationTypes(types unsafe.Pointer)
	RegisterServicesMenuSendTypesReturnTypes(sendTypes unsafe.Pointer, returnTypes unsafe.Pointer)
	RegisterUserInterfaceItemSearchHandler(handler objc.ID)
	ReplyToOpenOrPrint(reply unsafe.Pointer)
	ReportException(exception unsafe.Pointer)
	Run()
	RunModalForWindow(window unsafe.Pointer) unsafe.Pointer
	RunModalForWindowRelativeToWindow(window unsafe.Pointer, docWindow unsafe.Pointer) int
	RunModalSession(session unsafe.Pointer) unsafe.Pointer
	RunPageLayout(sender objc.ID)
	SendActionToFrom(action objc.SEL, target objc.ID, sender objc.ID) bool
	SendEvent(event unsafe.Pointer)
	SetActivationPolicy(activationPolicy unsafe.Pointer) bool
	SetWindowsNeedUpdate(needUpdate bool)
	StopModal()
	Terminate(sender objc.ID)
	ToggleTouchBarCustomizationPalette(sender objc.ID)
	UnhideWithoutActivation()
	UnregisterForRemoteNotifications()
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

// Makes the receiver the active app.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/activate(ignoringOtherApps:)
func (a_ Application) ActivateIgnoringOtherApps(ignoreOtherApps bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("activateIgnoringOtherApps:"), ignoreOtherApps)
}

// Returns the app’s activation policy.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/activationPolicy()
func (a_ Application) ActivationPolicy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("activationPolicy"))
	return rv
}

// Use the delegate method instead.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/application:printFiles:
func (a_ Application) ApplicationPrintFiles(sender unsafe.Pointer, filenames unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("application:printFiles:"), sender, filenames)
}

// Sets up a modal session with the given window and returns a pointer to the structure representing the session.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/beginModalSession(for:)
func (a_ Application) BeginModalSessionForWindow(window unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("beginModalSessionForWindow:"), window)
	return rv
}

// Starts a document modal session.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/beginSheet(_:modalFor:modalDelegate:didEnd:contextInfo:)
func (a_ Application) BeginSheetModalForWindowModalDelegateDidEndSelectorContextInfo(sheet unsafe.Pointer, docWindow unsafe.Pointer, modalDelegate objc.ID, didEndSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("beginSheet:modalForWindow:modalDelegate:didEndSelector:contextInfo:"), sheet, docWindow, modalDelegate, didEndSelector, contextInfo)
}

// Removes all events matching the given mask and generated before the specified event.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/discardEvents(matching:before:)
func (a_ Application) DiscardEventsMatchingMaskBeforeEvent(mask unsafe.Pointer, lastEvent unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("discardEventsMatchingMask:beforeEvent:"), mask, lastEvent)
}

// Finishes a modal session.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/endModalSession(_:)
func (a_ Application) EndModalSession(session unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("endModalSession:"), session)
}

// Ends a document modal session by specifying the sheet window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/endSheet(_:)
func (a_ Application) EndSheet(sheet unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("endSheet:"), sheet)
}

// Ends a document modal session by specifying the sheet window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/endSheet(_:returnCode:)
func (a_ Application) EndSheetReturnCode(sheet unsafe.Pointer, returnCode int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("endSheet:returnCode:"), sheet, returnCode)
}

// Activates the app, opens any files specified by the user default, and unhighlights the app’s icon.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/finishLaunching()
func (a_ Application) FinishLaunching() {
	objc.Send[objc.ID](a_.ID, objc.Sel("finishLaunching"))
}

// Hides all the receiver’s windows, and the next app in line is activated.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/hide(_:)
func (a_ Application) Hide(sender objc.ID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("hide:"), sender)
}

// Returns the next event matching a given mask, or if no such event is found before a specified expiration date.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/nextEvent(matching:until:inMode:dequeue:)
func (a_ Application) NextEventMatchingMaskUntilDateInModeDequeue(mask unsafe.Pointer, expiration unsafe.Pointer, mode unsafe.Pointer, deqFlag bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("nextEventMatchingMask:untilDate:inMode:dequeue:"), mask, expiration, mode, deqFlag)
	return rv
}

// Opens the character palette.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/orderFrontCharacterPalette(_:)
func (a_ Application) OrderFrontCharacterPalette(sender objc.ID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("orderFrontCharacterPalette:"), sender)
}

// Brings up the color panel, an instance of .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/orderFrontColorPanel(_:)
func (a_ Application) OrderFrontColorPanel(sender objc.ID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("orderFrontColorPanel:"), sender)
}

// Displays a standard About window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/orderFrontStandardAboutPanel(_:)
func (a_ Application) OrderFrontStandardAboutPanel(sender objc.ID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("orderFrontStandardAboutPanel:"), sender)
}

// Displays a standard About window with information from a given options dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/orderFrontStandardAboutPanel(options:)
func (a_ Application) OrderFrontStandardAboutPanelWithOptions(optionsDictionary unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("orderFrontStandardAboutPanelWithOptions:"), optionsDictionary)
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

// Register to receive notifications of the specified types from a provider through the Apple Push Notification service.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/registerForRemoteNotifications(matching:)
func (a_ Application) RegisterForRemoteNotificationTypes(types unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("registerForRemoteNotificationTypes:"), types)
}

// Registers the pasteboard types the receiver can send and receive in response to service requests.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/registerServicesMenuSendTypes(_:returnTypes:)
func (a_ Application) RegisterServicesMenuSendTypesReturnTypes(sendTypes unsafe.Pointer, returnTypes unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("registerServicesMenuSendTypes:returnTypes:"), sendTypes, returnTypes)
}

// Register an object that provides help data to your app.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/registerUserInterfaceItemSearchHandler(_:)
func (a_ Application) RegisterUserInterfaceItemSearchHandler(handler objc.ID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("registerUserInterfaceItemSearchHandler:"), handler)
}

// Handles errors that might occur when the user attempts to open or print files.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/reply(toOpenOrPrint:)
func (a_ Application) ReplyToOpenOrPrint(reply unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("replyToOpenOrPrint:"), reply)
}

// Logs a given exception by calling .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/reportException(_:)
func (a_ Application) ReportException(exception unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("reportException:"), exception)
}

// Starts the main event loop.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/run()
func (a_ Application) Run() {
	objc.Send[objc.ID](a_.ID, objc.Sel("run"))
}

// Starts a modal event loop for the specified window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/runModal(for:)
func (a_ Application) RunModalForWindow(window unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("runModalForWindow:"), window)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/runModalForWindow:relativeToWindow:
func (a_ Application) RunModalForWindowRelativeToWindow(window unsafe.Pointer, docWindow unsafe.Pointer) int {
	rv := objc.Send[int](a_.ID, objc.Sel("runModalForWindow:relativeToWindow:"), window, docWindow)
	return rv
}

// Runs a given modal session, as defined in a previous invocation of .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/runModalSession(_:)
func (a_ Application) RunModalSession(session unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("runModalSession:"), session)
	return rv
}

// Displays the receiver’s page layout panel, an instance of .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/runPageLayout(_:)
func (a_ Application) RunPageLayout(sender objc.ID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("runPageLayout:"), sender)
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

// Attempts to modify the app’s activation policy.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/setActivationPolicy(_:)
func (a_ Application) SetActivationPolicy(activationPolicy unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setActivationPolicy:"), activationPolicy)
	return rv
}

// Sets whether the receiver’s windows need updating when the receiver has finished processing the current event.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/setWindowsNeedUpdate(_:)
func (a_ Application) SetWindowsNeedUpdate(needUpdate bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setWindowsNeedUpdate:"), needUpdate)
}

// Stops a modal event loop.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/stopModal()
func (a_ Application) StopModal() {
	objc.Send[objc.ID](a_.ID, objc.Sel("stopModal"))
}

// Terminates the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/terminate(_:)
func (a_ Application) Terminate(sender objc.ID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("terminate:"), sender)
}

// Show or hides the interface for customizing the Touch Bar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/toggleTouchBarCustomizationPalette(_:)
func (a_ Application) ToggleTouchBarCustomizationPalette(sender objc.ID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("toggleTouchBarCustomizationPalette:"), sender)
}

// Restores hidden windows without activating their owner (the receiver).
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/unhideWithoutActivation()
func (a_ Application) UnhideWithoutActivation() {
	objc.Send[objc.ID](a_.ID, objc.Sel("unhideWithoutActivation"))
}

// Unregister for notifications received from Apple Push Notification service.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/unregisterForRemoteNotifications()
func (a_ Application) UnregisterForRemoteNotifications() {
	objc.Send[objc.ID](a_.ID, objc.Sel("unregisterForRemoteNotifications"))
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

// The types of push notifications that the app accepts.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/enabledRemoteNotificationTypes
func (a_ Application) EnabledRemoteNotificationTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("enabledRemoteNotificationTypes"))
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

// A Boolean value indicating whether the app is registered with Apple Push Notification service (APNs).
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/isRegisteredForRemoteNotifications
func (a_ Application) RegisteredForRemoteNotifications() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("registeredForRemoteNotifications"))
	return rv
}

// The app’s Services menu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/servicesMenu
func (a_ Application) ServicesMenu() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("servicesMenu"))
	return rv
}

// SetServicesMenu sets the value of the servicesMenu property.
// The app’s Services menu.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/servicesMenu
func (a_ Application) SetServicesMenu(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setServicesMenu:"), value)
}

// The object that provides the services the current app advertises in the Services menu of other apps.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/servicesProvider
func (a_ Application) ServicesProvider() objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("servicesProvider"))
	return rv
}

// SetServicesProvider sets the value of the servicesProvider property.
// The object that provides the services the current app advertises in the Services menu of other apps.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/servicesProvider
func (a_ Application) SetServicesProvider(value objc.ID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setServicesProvider:"), value)
}
