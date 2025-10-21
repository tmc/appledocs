// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	ApplicationPrintFiles(sender IApplication, filenames []string)
	BeginModalSessionForWindow(window IWindow) unsafe.Pointer
	BeginSheetModalForWindowModalDelegateDidEndSelectorContextInfo(sheet IWindow, docWindow IWindow, modalDelegate objectivec.IObject, didEndSelector objc.SEL, contextInfo unsafe.Pointer)
	DiscardEventsMatchingMaskBeforeEvent(mask EventMask, lastEvent IEvent)
	EndModalSession(session unsafe.Pointer)
	EndSheet(sheet IWindow)
	EndSheetReturnCode(sheet IWindow, returnCode int)
	FinishLaunching()
	Hide(sender objectivec.IObject)
	NextEventMatchingMaskUntilDateInModeDequeue(mask EventMask, expiration foundation.IDate, mode unsafe.Pointer, deqFlag bool) Event
	OrderFrontCharacterPalette(sender objectivec.IObject)
	OrderFrontColorPanel(sender objectivec.IObject)
	OrderFrontStandardAboutPanel(sender objectivec.IObject)
	OrderFrontStandardAboutPanelWithOptions(optionsDictionary unsafe.Pointer)
	PostEventAtStart(event IEvent, atStart bool)
	RegisterForRemoteNotifications()
	RegisterForRemoteNotificationTypes(types RemoteNotificationType)
	RegisterServicesMenuSendTypesReturnTypes(sendTypes []string, returnTypes []string)
	RegisterUserInterfaceItemSearchHandler(handler objectivec.IObject)
	ReplyToOpenOrPrint(reply unsafe.Pointer)
	ReportException(exception foundation.IException)
	Run()
	RunModalForWindow(window IWindow) ModalResponse
	RunModalForWindowRelativeToWindow(window IWindow, docWindow IWindow) int
	RunModalSession(session unsafe.Pointer) ModalResponse
	RunPageLayout(sender objectivec.IObject)
	SendActionToFrom(action objc.SEL, target objectivec.IObject, sender objectivec.IObject) bool
	SendEvent(event IEvent)
	SetWindowsNeedUpdate(needUpdate bool)
	StopModal()
	Terminate(sender objectivec.IObject)
	ToggleTouchBarCustomizationPalette(sender objectivec.IObject)
	UnhideWithoutActivation()
	UnregisterForRemoteNotifications()
	UpdateWindows()
	ValidRequestorForSendTypeReturnType(sendType PasteboardType, returnType PasteboardType) objc.ID
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


// Returns the application instance, creating it if it doesn’t exist yet.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/shared
func (ac _ApplicationClass) SharedApplication() Application {
	rv := objc.Send[NSApplication](objc.ID(ac.class), objc.Sel("sharedApplication"))
	return rv
}
// Makes the receiver the active app.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/activate(ignoringOtherApps:)
func (a_ Application) ActivateIgnoringOtherApps(ignoreOtherApps bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("activateIgnoringOtherApps:"), ignoreOtherApps)
}

// Use the delegate method instead.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/application:printFiles:
func (a_ Application) ApplicationPrintFiles(sender IApplication, filenames []string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("application:printFiles:"), sender, filenames)
}

// Sets up a modal session with the given window and returns a pointer to the structure representing the session.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/beginModalSession(for:)
func (a_ Application) BeginModalSessionForWindow(window IWindow) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("beginModalSessionForWindow:"), window)
	return rv
}

// Starts a document modal session.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/beginSheet(_:modalFor:modalDelegate:didEnd:contextInfo:)
func (a_ Application) BeginSheetModalForWindowModalDelegateDidEndSelectorContextInfo(sheet IWindow, docWindow IWindow, modalDelegate objectivec.IObject, didEndSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("beginSheet:modalForWindow:modalDelegate:didEndSelector:contextInfo:"), sheet, docWindow, modalDelegate, didEndSelector, contextInfo)
}

// Removes all events matching the given mask and generated before the specified event.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/discardEvents(matching:before:)
func (a_ Application) DiscardEventsMatchingMaskBeforeEvent(mask EventMask, lastEvent IEvent) {
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
func (a_ Application) EndSheet(sheet IWindow) {
	objc.Send[objc.ID](a_.ID, objc.Sel("endSheet:"), sheet)
}

// Ends a document modal session by specifying the sheet window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/endSheet(_:returnCode:)
func (a_ Application) EndSheetReturnCode(sheet IWindow, returnCode int) {
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
func (a_ Application) Hide(sender objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("hide:"), sender)
}

// Returns the next event matching a given mask, or if no such event is found before a specified expiration date.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/nextEvent(matching:until:inMode:dequeue:)
func (a_ Application) NextEventMatchingMaskUntilDateInModeDequeue(mask EventMask, expiration foundation.IDate, mode unsafe.Pointer, deqFlag bool) Event {
	rv := objc.Send[Event](a_.ID, objc.Sel("nextEventMatchingMask:untilDate:inMode:dequeue:"), mask, expiration, mode, deqFlag)
	return rv
}

// Opens the character palette.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/orderFrontCharacterPalette(_:)
func (a_ Application) OrderFrontCharacterPalette(sender objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("orderFrontCharacterPalette:"), sender)
}

// Brings up the color panel, an instance of .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/orderFrontColorPanel(_:)
func (a_ Application) OrderFrontColorPanel(sender objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("orderFrontColorPanel:"), sender)
}

// Displays a standard About window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/orderFrontStandardAboutPanel(_:)
func (a_ Application) OrderFrontStandardAboutPanel(sender objectivec.IObject) {
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
func (a_ Application) PostEventAtStart(event IEvent, atStart bool) {
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
func (a_ Application) RegisterForRemoteNotificationTypes(types RemoteNotificationType) {
	objc.Send[objc.ID](a_.ID, objc.Sel("registerForRemoteNotificationTypes:"), types)
}

// Registers the pasteboard types the receiver can send and receive in response to service requests.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/registerServicesMenuSendTypes(_:returnTypes:)
func (a_ Application) RegisterServicesMenuSendTypesReturnTypes(sendTypes []string, returnTypes []string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("registerServicesMenuSendTypes:returnTypes:"), sendTypes, returnTypes)
}

// Register an object that provides help data to your app.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/registerUserInterfaceItemSearchHandler(_:)
func (a_ Application) RegisterUserInterfaceItemSearchHandler(handler objectivec.IObject) {
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
func (a_ Application) ReportException(exception foundation.IException) {
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
func (a_ Application) RunModalForWindow(window IWindow) ModalResponse {
	rv := objc.Send[ModalResponse](a_.ID, objc.Sel("runModalForWindow:"), window)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/runModalForWindow:relativeToWindow:
func (a_ Application) RunModalForWindowRelativeToWindow(window IWindow, docWindow IWindow) int {
	rv := objc.Send[int](a_.ID, objc.Sel("runModalForWindow:relativeToWindow:"), window, docWindow)
	return rv
}

// Runs a given modal session, as defined in a previous invocation of .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/runModalSession(_:)
func (a_ Application) RunModalSession(session unsafe.Pointer) ModalResponse {
	rv := objc.Send[ModalResponse](a_.ID, objc.Sel("runModalSession:"), session)
	return rv
}

// Displays the receiver’s page layout panel, an instance of .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/runPageLayout(_:)
func (a_ Application) RunPageLayout(sender objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("runPageLayout:"), sender)
}

// Sends the given action message to the given target.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/sendAction(_:to:from:)
func (a_ Application) SendActionToFrom(action objc.SEL, target objectivec.IObject, sender objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("sendAction:to:from:"), action, target, sender)
	return rv
}

// Dispatches an event to other objects.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/sendEvent(_:)
func (a_ Application) SendEvent(event IEvent) {
	objc.Send[objc.ID](a_.ID, objc.Sel("sendEvent:"), event)
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
func (a_ Application) Terminate(sender objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("terminate:"), sender)
}

// Show or hides the interface for customizing the Touch Bar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/toggleTouchBarCustomizationPalette(_:)
func (a_ Application) ToggleTouchBarCustomizationPalette(sender objectivec.IObject) {
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
func (a_ Application) ValidRequestorForSendTypeReturnType(sendType PasteboardType, returnType PasteboardType) objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("validRequestorForSendType:returnType:"), sendType, returnType)
	return rv
}

// The appearance associated with the app’s windows.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/appearance
func (a_ Application) Appearance() NSAppearance {
	rv := objc.Send[NSAppearance](a_.ID, objc.Sel("appearance"))
	return rv
}


// SetAppearance sets the value of the appearance property.
// The appearance associated with the app’s windows.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/appearance
func (a_ Application) SetAppearance(value IAppearance) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAppearance:"), value)
}

// The image used for the app’s icon.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/applicationIconImage
func (a_ Application) ApplicationIconImage() Image {
	rv := objc.Send[Image](a_.ID, objc.Sel("applicationIconImage"))
	return rv
}


// SetApplicationIconImage sets the value of the applicationIconImage property.
// The image used for the app’s icon.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/applicationIconImage
func (a_ Application) SetApplicationIconImage(value IImage) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setApplicationIconImage:"), value)
}

// The last event object that the app retrieved from the event queue.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/currentEvent
func (a_ Application) CurrentEvent() NSEvent {
	rv := objc.Send[NSEvent](a_.ID, objc.Sel("currentEvent"))
	return rv
}

// The appearance that AppKit uses to draw the app’s interface.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/effectiveAppearance
func (a_ Application) EffectiveAppearance() NSAppearance {
	rv := objc.Send[NSAppearance](a_.ID, objc.Sel("effectiveAppearance"))
	return rv
}

// The types of push notifications that the app accepts.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/enabledRemoteNotificationTypes
func (a_ Application) EnabledRemoteNotificationTypes() RemoteNotificationType {
	rv := objc.Send[RemoteNotificationType](a_.ID, objc.Sel("enabledRemoteNotificationTypes"))
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
func (a_ Application) ServicesMenu() NSMenu {
	rv := objc.Send[NSMenu](a_.ID, objc.Sel("servicesMenu"))
	return rv
}


// SetServicesMenu sets the value of the servicesMenu property.
// The app’s Services menu.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/servicesMenu
func (a_ Application) SetServicesMenu(value IMenu) {
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

// Returns the application instance, creating it if it doesn’t exist yet.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/shared
func (a_ Application) SharedApplication() NSApplication {
	rv := objc.Send[NSApplication](a_.ID, objc.Sel("sharedApplication"))
	return rv
}

// The global variable for the shared app instance.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapp
func (a_ Application) NSApp() NSApplication {
	rv := objc.Send[NSApplication](a_.ID, objc.Sel("NSApp"))
	return rv
}


// SetNSApp sets the value of the NSApp property.
// The global variable for the shared app instance.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapp
func (a_ Application) SetNSApp(value IApplication) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setNSApp:"), value)
}

// A boolean value indicating whether your application should suppress HDR content based on established policy.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/applicationshouldsuppresshighdynamicrangecontent
func (a_ Application) ApplicationShouldSuppressHighDynamicRangeContent() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("applicationShouldSuppressHighDynamicRangeContent"))
	return rv
}


// SetApplicationShouldSuppressHighDynamicRangeContent sets the value of the applicationShouldSuppressHighDynamicRangeContent property.
// A boolean value indicating whether your application should suppress HDR content based on established policy.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/applicationshouldsuppresshighdynamicrangecontent
func (a_ Application) SetApplicationShouldSuppressHighDynamicRangeContent(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setApplicationShouldSuppressHighDynamicRangeContent:"), value)
}

// The set of app presentation options that are currently in effect for the system.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/currentsystempresentationoptions
func (a_ Application) CurrentSystemPresentationOptions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("currentSystemPresentationOptions"))
	return rv
}


// SetCurrentSystemPresentationOptions sets the value of the currentSystemPresentationOptions property.
// The set of app presentation options that are currently in effect for the system.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/currentsystempresentationoptions
func (a_ Application) SetCurrentSystemPresentationOptions(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurrentSystemPresentationOptions:"), value)
}

// The app delegate object.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/delegate
func (a_ Application) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The app delegate object.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/delegate
func (a_ Application) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDelegate:"), value)
}

// The app’s Dock tile.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/docktile
func (a_ Application) DockTile() NSDockTile {
	rv := objc.Send[NSDockTile](a_.ID, objc.Sel("dockTile"))
	return rv
}


// SetDockTile sets the value of the dockTile property.
// The app’s Dock tile.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/docktile
func (a_ Application) SetDockTile(value IDockTile) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDockTile:"), value)
}

// The help menu used by the app.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/helpmenu
func (a_ Application) HelpMenu() NSMenu {
	rv := objc.Send[NSMenu](a_.ID, objc.Sel("helpMenu"))
	return rv
}


// SetHelpMenu sets the value of the helpMenu property.
// The help menu used by the app.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/helpmenu
func (a_ Application) SetHelpMenu(value IMenu) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setHelpMenu:"), value)
}

// A Boolean value indicating whether this is the active app.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/isactive
func (a_ Application) IsActive() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isActive"))
	return rv
}


// SetIsActive sets the value of the isActive property.
// A Boolean value indicating whether this is the active app.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/isactive
func (a_ Application) SetIsActive(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsActive:"), value)
}

// A Boolean value indicating whether Full Keyboard Access is enabled in the Keyboard preference pane.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/isfullkeyboardaccessenabled
func (a_ Application) IsFullKeyboardAccessEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isFullKeyboardAccessEnabled"))
	return rv
}


// SetIsFullKeyboardAccessEnabled sets the value of the isFullKeyboardAccessEnabled property.
// A Boolean value indicating whether Full Keyboard Access is enabled in the Keyboard preference pane.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/isfullkeyboardaccessenabled
func (a_ Application) SetIsFullKeyboardAccessEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsFullKeyboardAccessEnabled:"), value)
}

// A Boolean value indicating whether the app is registered with Apple Push Notification service (APNs).
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/isregisteredforremotenotifications
func (a_ Application) IsRegisteredForRemoteNotifications() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isRegisteredForRemoteNotifications"))
	return rv
}


// SetIsRegisteredForRemoteNotifications sets the value of the isRegisteredForRemoteNotifications property.
// A Boolean value indicating whether the app is registered with Apple Push Notification service (APNs).

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/isregisteredforremotenotifications
func (a_ Application) SetIsRegisteredForRemoteNotifications(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsRegisteredForRemoteNotifications:"), value)
}

// A Boolean value indicating whether the main event loop is running.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/isrunning
func (a_ Application) IsRunning() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isRunning"))
	return rv
}


// SetIsRunning sets the value of the isRunning property.
// A Boolean value indicating whether the main event loop is running.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/isrunning
func (a_ Application) SetIsRunning(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsRunning:"), value)
}

// An array of document objects arranged according to the front-to-back ordering of their associated windows.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/ordereddocuments
func (a_ Application) OrderedDocuments() NSDocument {
	rv := objc.Send[NSDocument](a_.ID, objc.Sel("orderedDocuments"))
	return rv
}


// SetOrderedDocuments sets the value of the orderedDocuments property.
// An array of document objects arranged according to the front-to-back ordering of their associated windows.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/ordereddocuments
func (a_ Application) SetOrderedDocuments(value IDocument) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOrderedDocuments:"), value)
}

// An array of window objects arranged according to their front-to-back ordering on the screen.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/orderedwindows
func (a_ Application) OrderedWindows() NSWindow {
	rv := objc.Send[NSWindow](a_.ID, objc.Sel("orderedWindows"))
	return rv
}


// SetOrderedWindows sets the value of the orderedWindows property.
// An array of window objects arranged according to their front-to-back ordering on the screen.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/orderedwindows
func (a_ Application) SetOrderedWindows(value IWindow) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOrderedWindows:"), value)
}

// The presentation options that should be in effect for the system when this app is active.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/presentationoptions-swift.property
func (a_ Application) PresentationOptions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("presentationOptions"))
	return rv
}


// SetPresentationOptions sets the value of the presentationOptions property.
// The presentation options that should be in effect for the system when this app is active.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/presentationoptions-swift.property
func (a_ Application) SetPresentationOptions(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPresentationOptions:"), value)
}

// The layout direction of the user interface.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/userinterfacelayoutdirection
func (a_ Application) UserInterfaceLayoutDirection() UserInterfaceLayoutDirection {
	rv := objc.Send[UserInterfaceLayoutDirection](a_.ID, objc.Sel("userInterfaceLayoutDirection"))
	return rv
}


// SetUserInterfaceLayoutDirection sets the value of the userInterfaceLayoutDirection property.
// The layout direction of the user interface.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/userinterfacelayoutdirection
func (a_ Application) SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setUserInterfaceLayoutDirection:"), value)
}

// Indicates the activation policy of the application.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/activationpolicy
func (a_ Application) ActivationPolicy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("activationPolicy"))
	return rv
}


// SetActivationPolicy sets the value of the activationPolicy property.
// Indicates the activation policy of the application.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/activationpolicy
func (a_ Application) SetActivationPolicy(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setActivationPolicy:"), value)
}



