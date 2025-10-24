// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
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
	// properties:
	Appearance() IAppearance
	SetAppearance(value IAppearance)
	ApplicationIconImage() IImage
	SetApplicationIconImage(value IImage)
	CurrentEvent() IEvent
	EffectiveAppearance() IAppearance
	HelpMenu() IMenu
	SetHelpMenu(value IMenu)
	Running() bool
	ModalWindow() IWindow
	OrderedWindows() []IWindow
	ServicesProvider() objc.ID
	SetServicesProvider(value objc.ID)
	UserInterfaceLayoutDirection() UserInterfaceLayoutDirection
	NSApp() IApplication
	SetNSApp(value IApplication)
	ApplicationShouldSuppressHighDynamicRangeContent() bool
	SetApplicationShouldSuppressHighDynamicRangeContent(value bool)
	CurrentSystemPresentationOptions() unsafe.Pointer
	SetCurrentSystemPresentationOptions(value unsafe.Pointer)
	Delegate() ApplicationDelegate /* not a class type */
	SetDelegate(value ApplicationDelegate /* not a class type */)
	DockTile() IDockTile
	SetDockTile(value IDockTile)
	EnabledRemoteNotificationTypes() unsafe.Pointer
	SetEnabledRemoteNotificationTypes(value unsafe.Pointer)
	IsActive() bool
	SetIsActive(value bool)
	IsFullKeyboardAccessEnabled() bool
	SetIsFullKeyboardAccessEnabled(value bool)
	IsRegisteredForRemoteNotifications() bool
	SetIsRegisteredForRemoteNotifications(value bool)
	IsRunning() bool
	SetIsRunning(value bool)
	OrderedDocuments() objc.IObject /* cross-framework: Document */
	SetOrderedDocuments(value objc.IObject /* cross-framework: Document */)
	PresentationOptions() unsafe.Pointer
	SetPresentationOptions(value unsafe.Pointer)
	ActivationPolicy() unsafe.Pointer
	SetActivationPolicy(value unsafe.Pointer)
	// methods:
	AbortModal()
	Activate()
	BeginModalSessionForWindow(window IWindow) objc.IObject /* cross-framework: ModalSession */
	CancelUserAttentionRequest(request int)
	Deactivate()
	DiscardEventsMatchingMaskBeforeEvent(mask EventMask, lastEvent IEvent)
	EnableRelaunchOnLogin()
	EndModalSession(session objc.IObject /* cross-framework: ModalSession */)
	FinishLaunching()
	Hide(sender objectivec.IObject)
	NextEventMatchingMaskUntilDateInModeDequeue(mask EventMask, expiration objc.IObject /* cross-framework: NSDate */, mode RunLoopMode /* not a class type */, deqFlag bool) IEvent
	OrderFrontStandardAboutPanel(sender objectivec.IObject)
	PostEventAtStart(event IEvent, atStart bool)
	PreventWindowOrdering()
	RegisterForRemoteNotifications()
	RegisterServicesMenuSendTypesReturnTypes(sendTypes []string, returnTypes []string)
	RegisterUserInterfaceItemSearchHandler(handler objectivec.IObject)
	ReplyToApplicationShouldTerminate(shouldTerminate bool)
	ReplyToOpenOrPrint(reply ApplicationDelegateReply /* not a class type */)
	RequestUserAttention(requestType RequestUserAttentionType) int
	Run()
	RunModalForWindow(window IWindow) objc.IObject /* cross-framework: ModalResponse */
	RunModalSession(session objc.IObject /* cross-framework: ModalSession */) objc.IObject /* cross-framework: ModalResponse */
	SearchStringInUserInterfaceItemStringSearchRangeFoundRange(searchString objc.IObject /* cross-framework: NSString */, stringToSearch objc.IObject /* cross-framework: NSString */, searchRange objc.IObject /* cross-framework: Range */, foundRange objc.IObject /* cross-framework: Range */) bool
	SendActionToFrom(action objc.SEL, target objectivec.IObject, sender objectivec.IObject) bool
	SendEvent(event IEvent)
	SetWindowsNeedUpdate(needUpdate bool)
	ShowHelp(sender objectivec.IObject)
	Stop(sender objectivec.IObject)
	StopModal()
	StopModalWithCode(returnCode objc.IObject /* cross-framework: ModalResponse */)
	TargetForAction(action objc.SEL) objc.ID
	TargetForActionToFrom(action objc.SEL, target objectivec.IObject, sender objectivec.IObject) objc.ID
	Terminate(sender objectivec.IObject)
	TryToPerformWith(action objc.SEL, object objectivec.IObject) bool
	UnhideAllApplications(sender objectivec.IObject)
	UnhideWithoutActivation()
	UnregisterUserInterfaceItemSearchHandler(handler objectivec.IObject)
	UpdateWindows()
	ValidRequestorForSendTypeReturnType(sendType objc.IObject /* cross-framework: PasteboardType */, returnType objc.IObject /* cross-framework: PasteboardType */) objc.ID
	YieldActivationToApplication(application IRunningApplication)
	YieldActivationToApplicationWithBundleIdentifier(bundleIdentifier objc.IObject /* cross-framework: NSString */)
}

// An object that manages an app’s main event loop and resources used by all of that app’s objects.
//
// Every app uses a single instance of to control the main event loop, keep track of the app’s windows and menus, distribute events to the appropriate objects (that’s, itself or one of its windows), set up autorelease pools, and receive notification of app-level events. An object has a delegate (an object that you assign) that’s notified when the app starts or terminates, is hidden or activated, should open a file selected by the user, and so forth. By setting the delegate and implementing the delegate methods, you customize the behavior of your app without having to subclass . In your app’s function, create the instance by calling the class method. After creating the application object, the function should load your app’s main nib file and then start the event loop by sending the application object a message. If you create an Application project in Xcode, this function is created for you. The function Xcode creates begins by calling a function named , which is functionally similar to the following: The class method initializes the display environment and connects your program to the window server and the display server. The object maintains a list of all the objects the app uses, so it can retrieve any of the app’s objects. The method also initializes the global variable , which you use to retrieve the instance. only performs the initialization once. If you invoke it more than once, it returns the application object it created previously. The shared object performs the important task of receiving events from the window server and distributing them to the proper objects. translates an event into an object, then forwards the event object to the affected object. All keyboard and mouse events go directly to the object associated with the event. The only exception to this rule is if the Command key is pressed when a key-down event occurs; in this case, every object has an opportunity to respond to the event. When a window object receives an object from , it distributes it to the objects in its view hierarchy. is also responsible for dispatching certain Apple events received by the app. For example, macOS sends Apple events to your app at various times, such as when the app is launched or reopened. installs Apple event handlers to handle these events by sending a message to the appropriate object. You can also use the class to register your own Apple event handlers. The method is generally the best place to do so. For more information on how events are handled and how you can modify the default behavior, including information on working with Apple events in scriptable apps, see in . The class sets up block during initialization and inside the event loop—specifically, within its initialization (or ) and methods. Similarly, the methods AppKit adds to employ blocks during the loading of nib files. These blocks aren’t accessible outside the scope of the respective and methods. Typically, an app creates objects either while the event loop is running or by loading objects from nib files, so this lack of access usually isn’t a problem. However, if you do need to use Cocoa classes within the function itself (other than to load nib files or to instantiate ), you should create an block to contain the code using the classes.


// An object that manages an app’s main event loop and resources used by all of that app’s objects.
//
// [Full Topic]
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



// Creates and executes a new thread based on the specified target and selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/detachDrawingThread(_:toTarget:with:)
func (ac _ApplicationClass) DetachDrawingThreadToTargetWithObject(selector objc.SEL, target objectivec.IObject, argument objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("detachDrawingThread:toTarget:withObject:"), selector, target, argument)
}


// Returns the application instance, creating it if it doesn’t exist yet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/shared
func (ac _ApplicationClass) SharedApplication() Application {
	rv := objc.Send[Application](objc.ID(ac.class), objc.Sel("sharedApplication"))
	return rv
}

// Aborts the event loop started by or .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/abortModal()
func (a_ Application) AbortModal() {
	objc.Send[objc.ID](a_.ID, objc.Sel("abortModal"))
}


// Activates the receiver app, if appropriate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/activate()
func (a_ Application) Activate() {
	objc.Send[objc.ID](a_.ID, objc.Sel("activate"))
}


// Sets up a modal session with the given window and returns a pointer to the structure representing the session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/beginModalSession(for:)
func (a_ Application) BeginModalSessionForWindow(window IWindow) objc.IObject /* cross-framework: ModalSession */ {
	rv := objc.Send[ModalSession](a_.ID, objc.Sel("beginModalSessionForWindow:"), window)
	return rv
}


// Cancels a previous user attention request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/cancelUserAttentionRequest(_:)
func (a_ Application) CancelUserAttentionRequest(request int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("cancelUserAttentionRequest:"), request)
}


// Deactivates the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/deactivate()
func (a_ Application) Deactivate() {
	objc.Send[objc.ID](a_.ID, objc.Sel("deactivate"))
}


// Removes all events matching the given mask and generated before the specified event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/discardEvents(matching:before:)
func (a_ Application) DiscardEventsMatchingMaskBeforeEvent(mask EventMask, lastEvent IEvent) {
	objc.Send[objc.ID](a_.ID, objc.Sel("discardEventsMatchingMask:beforeEvent:"), mask, lastEvent)
}


// Enables relaunching the app on login.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/enableRelaunchOnLogin()
func (a_ Application) EnableRelaunchOnLogin() {
	objc.Send[objc.ID](a_.ID, objc.Sel("enableRelaunchOnLogin"))
}


// Finishes a modal session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/endModalSession(_:)
func (a_ Application) EndModalSession(session objc.IObject /* cross-framework: ModalSession */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("endModalSession:"), session)
}


// Activates the app, opens any files specified by the user default, and unhighlights the app’s icon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/finishLaunching()
func (a_ Application) FinishLaunching() {
	objc.Send[objc.ID](a_.ID, objc.Sel("finishLaunching"))
}


// Hides all the receiver’s windows, and the next app in line is activated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/hide(_:)
func (a_ Application) Hide(sender objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("hide:"), sender)
}


// Returns the next event matching a given mask, or if no such event is found before a specified expiration date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/nextEvent(matching:until:inMode:dequeue:)
func (a_ Application) NextEventMatchingMaskUntilDateInModeDequeue(mask EventMask, expiration objc.IObject /* cross-framework: NSDate */, mode RunLoopMode /* not a class type */, deqFlag bool) IEvent {
	rv := objc.Send[Event](a_.ID, objc.Sel("nextEventMatchingMask:untilDate:inMode:dequeue:"), mask, expiration, mode, deqFlag)
	return rv
}


// Displays a standard About window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/orderFrontStandardAboutPanel(_:)
func (a_ Application) OrderFrontStandardAboutPanel(sender objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("orderFrontStandardAboutPanel:"), sender)
}


// Adds a given event to the receiver’s event queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/postEvent(_:atStart:)
func (a_ Application) PostEventAtStart(event IEvent, atStart bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("postEvent:atStart:"), event, atStart)
}


// Suppresses the usual window ordering in handling the most recent mouse-down event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/preventWindowOrdering()
func (a_ Application) PreventWindowOrdering() {
	objc.Send[objc.ID](a_.ID, objc.Sel("preventWindowOrdering"))
}


// Register for notifications sent by Apple Push Notification service (APNs).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/registerForRemoteNotifications()
func (a_ Application) RegisterForRemoteNotifications() {
	objc.Send[objc.ID](a_.ID, objc.Sel("registerForRemoteNotifications"))
}


// Registers the pasteboard types the receiver can send and receive in response to service requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/registerServicesMenuSendTypes(_:returnTypes:)
func (a_ Application) RegisterServicesMenuSendTypesReturnTypes(sendTypes []string, returnTypes []string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("registerServicesMenuSendTypes:returnTypes:"), sendTypes, returnTypes)
}


// Register an object that provides help data to your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/registerUserInterfaceItemSearchHandler(_:)
func (a_ Application) RegisterUserInterfaceItemSearchHandler(handler objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("registerUserInterfaceItemSearchHandler:"), handler)
}


// Responds to once the app knows whether it can terminate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/reply(toApplicationShouldTerminate:)
func (a_ Application) ReplyToApplicationShouldTerminate(shouldTerminate bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("replyToApplicationShouldTerminate:"), shouldTerminate)
}


// Handles errors that might occur when the user attempts to open or print files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/reply(toOpenOrPrint:)
func (a_ Application) ReplyToOpenOrPrint(reply ApplicationDelegateReply /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("replyToOpenOrPrint:"), reply)
}


// Starts a user attention request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/requestUserAttention(_:)
func (a_ Application) RequestUserAttention(requestType RequestUserAttentionType) int {
	rv := objc.Send[int](a_.ID, objc.Sel("requestUserAttention:"), requestType)
	return rv
}


// Starts the main event loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/run()
func (a_ Application) Run() {
	objc.Send[objc.ID](a_.ID, objc.Sel("run"))
}


// Starts a modal event loop for the specified window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/runModal(for:)
func (a_ Application) RunModalForWindow(window IWindow) objc.IObject /* cross-framework: ModalResponse */ {
	rv := objc.Send[ModalResponse](a_.ID, objc.Sel("runModalForWindow:"), window)
	return rv
}


// Runs a given modal session, as defined in a previous invocation of .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/runModalSession(_:)
func (a_ Application) RunModalSession(session objc.IObject /* cross-framework: ModalSession */) objc.IObject /* cross-framework: ModalResponse */ {
	rv := objc.Send[ModalResponse](a_.ID, objc.Sel("runModalSession:"), session)
	return rv
}


// Searches for the string in the user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/searchString(_:inUserInterfaceItemString:range:found:)
func (a_ Application) SearchStringInUserInterfaceItemStringSearchRangeFoundRange(searchString objc.IObject /* cross-framework: NSString */, stringToSearch objc.IObject /* cross-framework: NSString */, searchRange objc.IObject /* cross-framework: Range */, foundRange objc.IObject /* cross-framework: Range */) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("searchString:inUserInterfaceItemString:searchRange:foundRange:"), searchString, stringToSearch, searchRange, foundRange)
	return rv
}


// Sends the given action message to the given target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/sendAction(_:to:from:)
func (a_ Application) SendActionToFrom(action objc.SEL, target objectivec.IObject, sender objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("sendAction:to:from:"), action, target, sender)
	return rv
}


// Dispatches an event to other objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/sendEvent(_:)
func (a_ Application) SendEvent(event IEvent) {
	objc.Send[objc.ID](a_.ID, objc.Sel("sendEvent:"), event)
}


// Sets whether the receiver’s windows need updating when the receiver has finished processing the current event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/setWindowsNeedUpdate(_:)
func (a_ Application) SetWindowsNeedUpdate(needUpdate bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setWindowsNeedUpdate:"), needUpdate)
}


// If your project is properly registered, and the necessary keys have been set in the property list, this method launches Help Viewer and displays the first page of your app’s help book.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/showHelp(_:)
func (a_ Application) ShowHelp(sender objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("showHelp:"), sender)
}


// Stops the main event loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/stop(_:)
func (a_ Application) Stop(sender objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("stop:"), sender)
}


// Stops a modal event loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/stopModal()
func (a_ Application) StopModal() {
	objc.Send[objc.ID](a_.ID, objc.Sel("stopModal"))
}


// Stops a modal event loop, allowing you to return a custom result code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/stopModal(withCode:)
func (a_ Application) StopModalWithCode(returnCode objc.IObject /* cross-framework: ModalResponse */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("stopModalWithCode:"), returnCode)
}


// Returns the object that receives the action message specified by the given selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/target(forAction:)
func (a_ Application) TargetForAction(action objc.SEL) objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("targetForAction:"), action)
	return rv
}


// Searches for an object that can receive the message specified by the given selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/target(forAction:to:from:)
func (a_ Application) TargetForActionToFrom(action objc.SEL, target objectivec.IObject, sender objectivec.IObject) objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("targetForAction:to:from:"), action, target, sender)
	return rv
}


// Terminates the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/terminate(_:)
func (a_ Application) Terminate(sender objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("terminate:"), sender)
}


// Dispatches an action message to the specified target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/tryToPerform(_:with:)
func (a_ Application) TryToPerformWith(action objc.SEL, object objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("tryToPerform:with:"), action, object)
	return rv
}


// Unhides all apps, including the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/unhideAllApplications(_:)
func (a_ Application) UnhideAllApplications(sender objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("unhideAllApplications:"), sender)
}


// Restores hidden windows without activating their owner (the receiver).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/unhideWithoutActivation()
func (a_ Application) UnhideWithoutActivation() {
	objc.Send[objc.ID](a_.ID, objc.Sel("unhideWithoutActivation"))
}


// Unregister an object that provides help data to your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/unregisterUserInterfaceItemSearchHandler(_:)
func (a_ Application) UnregisterUserInterfaceItemSearchHandler(handler objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("unregisterUserInterfaceItemSearchHandler:"), handler)
}


// Sends an message to each onscreen window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/updateWindows()
func (a_ Application) UpdateWindows() {
	objc.Send[objc.ID](a_.ID, objc.Sel("updateWindows"))
}


// Indicates whether the receiver can send and receive the specified pasteboard types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/validRequestor(forSendType:returnType:)
func (a_ Application) ValidRequestorForSendTypeReturnType(sendType objc.IObject /* cross-framework: PasteboardType */, returnType objc.IObject /* cross-framework: PasteboardType */) objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("validRequestorForSendType:returnType:"), sendType, returnType)
	return rv
}


// Explicitly allows another app to make itself active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/yieldActivation(to:)
func (a_ Application) YieldActivationToApplication(application IRunningApplication) {
	objc.Send[objc.ID](a_.ID, objc.Sel("yieldActivationToApplication:"), application)
}


// Explicitly allows another app to make itself active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/yieldActivation(toApplicationWithBundleIdentifier:)
func (a_ Application) YieldActivationToApplicationWithBundleIdentifier(bundleIdentifier objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("yieldActivationToApplicationWithBundleIdentifier:"), bundleIdentifier)
}


// The appearance associated with the app’s windows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/appearance
func (a_ Application) Appearance() IAppearance {
	rv := objc.Send[Appearance](a_.ID, objc.Sel("appearance"))
	return rv
}


// The appearance associated with the app’s windows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/appearance
func (a_ Application) SetAppearance(value IAppearance) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAppearance:"), value)
}


// The image used for the app’s icon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/applicationIconImage
func (a_ Application) ApplicationIconImage() IImage {
	rv := objc.Send[Image](a_.ID, objc.Sel("applicationIconImage"))
	return rv
}


// The image used for the app’s icon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/applicationIconImage
func (a_ Application) SetApplicationIconImage(value IImage) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setApplicationIconImage:"), value)
}


// The last event object that the app retrieved from the event queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/currentEvent
func (a_ Application) CurrentEvent() IEvent {
	rv := objc.Send[Event](a_.ID, objc.Sel("currentEvent"))
	return rv
}


// The appearance that AppKit uses to draw the app’s interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/effectiveAppearance
func (a_ Application) EffectiveAppearance() IAppearance {
	rv := objc.Send[Appearance](a_.ID, objc.Sel("effectiveAppearance"))
	return rv
}


// The help menu used by the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/helpMenu
func (a_ Application) HelpMenu() IMenu {
	rv := objc.Send[Menu](a_.ID, objc.Sel("helpMenu"))
	return rv
}


// The help menu used by the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/helpMenu
func (a_ Application) SetHelpMenu(value IMenu) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setHelpMenu:"), value)
}


// A Boolean value indicating whether the main event loop is running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/isRunning
func (a_ Application) Running() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("running"))
	return rv
}


// The modal window displayed by the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/modalWindow
func (a_ Application) ModalWindow() IWindow {
	rv := objc.Send[Window](a_.ID, objc.Sel("modalWindow"))
	return rv
}


// An array of window objects arranged according to their front-to-back ordering on the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/orderedWindows
func (a_ Application) OrderedWindows() []IWindow {
	rv := objc.Send[[]Window](a_.ID, objc.Sel("orderedWindows"))
	return rv
}


// The object that provides the services the current app advertises in the Services menu of other apps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/servicesProvider
func (a_ Application) ServicesProvider() objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("servicesProvider"))
	return rv
}


// The object that provides the services the current app advertises in the Services menu of other apps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/servicesProvider
func (a_ Application) SetServicesProvider(value objc.ID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setServicesProvider:"), value)
}


// Returns the application instance, creating it if it doesn’t exist yet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/shared
func (a_ Application) SharedApplication() IApplication {
	rv := objc.Send[Application](a_.ID, objc.Sel("sharedApplication"))
	return rv
}


// The layout direction of the user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/userInterfaceLayoutDirection
func (a_ Application) UserInterfaceLayoutDirection() UserInterfaceLayoutDirection {
	rv := objc.Send[UserInterfaceLayoutDirection](a_.ID, objc.Sel("userInterfaceLayoutDirection"))
	return rv
}


// The global variable for the shared app instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapp
func (a_ Application) NSApp() IApplication {
	rv := objc.Send[Application](a_.ID, objc.Sel("NSApp"))
	return rv
}


// The global variable for the shared app instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapp
func (a_ Application) SetNSApp(value IApplication) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setNSApp:"), value)
}


// A boolean value indicating whether your application should suppress HDR content based on established policy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/applicationshouldsuppresshighdynamicrangecontent
func (a_ Application) ApplicationShouldSuppressHighDynamicRangeContent() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("applicationShouldSuppressHighDynamicRangeContent"))
	return rv
}


// A boolean value indicating whether your application should suppress HDR content based on established policy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/applicationshouldsuppresshighdynamicrangecontent
func (a_ Application) SetApplicationShouldSuppressHighDynamicRangeContent(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setApplicationShouldSuppressHighDynamicRangeContent:"), value)
}


// The set of app presentation options that are currently in effect for the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/currentsystempresentationoptions
func (a_ Application) CurrentSystemPresentationOptions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("currentSystemPresentationOptions"))
	return rv
}


// The set of app presentation options that are currently in effect for the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/currentsystempresentationoptions
func (a_ Application) SetCurrentSystemPresentationOptions(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurrentSystemPresentationOptions:"), value)
}


// The app delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/delegate
func (a_ Application) Delegate() ApplicationDelegate /* not a class type */ {
	rv := objc.Send[ApplicationDelegate](a_.ID, objc.Sel("delegate"))
	return rv
}


// The app delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/delegate
func (a_ Application) SetDelegate(value ApplicationDelegate /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDelegate:"), value)
}


// The app’s Dock tile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/docktile
func (a_ Application) DockTile() IDockTile {
	rv := objc.Send[DockTile](a_.ID, objc.Sel("dockTile"))
	return rv
}


// The app’s Dock tile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/docktile
func (a_ Application) SetDockTile(value IDockTile) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDockTile:"), value)
}


// The types of push notifications that the app accepts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/enabledremotenotificationtypes
func (a_ Application) EnabledRemoteNotificationTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("enabledRemoteNotificationTypes"))
	return rv
}


// The types of push notifications that the app accepts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/enabledremotenotificationtypes
func (a_ Application) SetEnabledRemoteNotificationTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setEnabledRemoteNotificationTypes:"), value)
}


// A Boolean value indicating whether this is the active app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/isactive
func (a_ Application) IsActive() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isActive"))
	return rv
}


// A Boolean value indicating whether this is the active app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/isactive
func (a_ Application) SetIsActive(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsActive:"), value)
}


// A Boolean value indicating whether Full Keyboard Access is enabled in the Keyboard preference pane.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/isfullkeyboardaccessenabled
func (a_ Application) IsFullKeyboardAccessEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isFullKeyboardAccessEnabled"))
	return rv
}


// A Boolean value indicating whether Full Keyboard Access is enabled in the Keyboard preference pane.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/isfullkeyboardaccessenabled
func (a_ Application) SetIsFullKeyboardAccessEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsFullKeyboardAccessEnabled:"), value)
}


// A Boolean value indicating whether the app is registered with Apple Push Notification service (APNs).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/isregisteredforremotenotifications
func (a_ Application) IsRegisteredForRemoteNotifications() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isRegisteredForRemoteNotifications"))
	return rv
}


// A Boolean value indicating whether the app is registered with Apple Push Notification service (APNs).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/isregisteredforremotenotifications
func (a_ Application) SetIsRegisteredForRemoteNotifications(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsRegisteredForRemoteNotifications:"), value)
}


// A Boolean value indicating whether the main event loop is running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/isrunning
func (a_ Application) IsRunning() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isRunning"))
	return rv
}


// A Boolean value indicating whether the main event loop is running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/isrunning
func (a_ Application) SetIsRunning(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsRunning:"), value)
}


// An array of document objects arranged according to the front-to-back ordering of their associated windows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/ordereddocuments
func (a_ Application) OrderedDocuments() objc.IObject /* cross-framework: Document */ {
	rv := objc.Send[Document](a_.ID, objc.Sel("orderedDocuments"))
	return rv
}


// An array of document objects arranged according to the front-to-back ordering of their associated windows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/ordereddocuments
func (a_ Application) SetOrderedDocuments(value objc.IObject /* cross-framework: Document */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOrderedDocuments:"), value)
}


// The presentation options that should be in effect for the system when this app is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/presentationoptions-swift.property
func (a_ Application) PresentationOptions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("presentationOptions"))
	return rv
}


// The presentation options that should be in effect for the system when this app is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/presentationoptions-swift.property
func (a_ Application) SetPresentationOptions(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPresentationOptions:"), value)
}


// Indicates the activation policy of the application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/activationpolicy
func (a_ Application) ActivationPolicy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("activationPolicy"))
	return rv
}


// Indicates the activation policy of the application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrunningapplication/activationpolicy
func (a_ Application) SetActivationPolicy(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setActivationPolicy:"), value)
}



