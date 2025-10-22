// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ExtensionContext] class.
var (
	ExtensionContextClass     _ExtensionContextClass
	ExtensionContextClassOnce sync.Once
)

func getExtensionContextClass() _ExtensionContextClass {
	ExtensionContextClassOnce.Do(func() {
		ExtensionContextClass = _ExtensionContextClass{objc.GetClass("NSExtensionContext")}
	})
	return ExtensionContextClass
}

type _ExtensionContextClass struct {
	class objc.Class
}

// An interface definition for the [ExtensionContext] class.
type IExtensionContext interface {
	objectivec.IObject
	CancelRequestWithError(error_ IError)
	CompleteRequestReturningItemsCompletionHandler(items objectivec.IObject, completionHandler unsafe.Pointer)
	CompleteRequestWithBroadcastURLBroadcastConfigurationSetupInfo(broadcastURL IURL, broadcastConfiguration objectivec.IObject, setupInfo IDictionary)
	CompleteRequestWithBroadcastURLSetupInfo(broadcastURL IURL, setupInfo IDictionary)
	DismissNotificationContentExtension()
	InterfaceParametersDescription() String
	LoadBroadcastingApplicationInfoWithCompletion(handler unsafe.Pointer)
	MediaPlayingPaused()
	MediaPlayingStarted()
	OpenURLCompletionHandler(URL IURL, completionHandler unsafe.Pointer)
	PerformNotificationDefaultAction()
	WidgetMaximumSizeForDisplayMode(displayMode unsafe.Pointer) coregraphics.CGSize
	HostedViewMaximumAllowedSize() coregraphics.CGSize
	HostedViewMinimumAllowedSize() coregraphics.CGSize
	InputItems() objc.ID
	NotificationActions() []objc.ID
	SetNotificationActions(value []objc.ID)
	WidgetActiveDisplayMode() unsafe.Pointer
	WidgetLargestAvailableDisplayMode() unsafe.Pointer
	SetWidgetLargestAvailableDisplayMode(value unsafe.Pointer)
	NSExtensionItemsAndErrorsKey() string
}

// The host app context from which an app extension is invoked.
//
// When a host app sends a request to an app extension, it provides an extension context. For many app extensions, the most important part of the context is the data the user wants to work with, which is contained in the property.


// The host app context from which an app extension is invoked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionContext
type ExtensionContext struct {
	objectivec.Object
}

// ExtensionContextFrom constructs a [ExtensionContext] from an unsafe.Pointer.
//
// The host app context from which an app extension is invoked.
func ExtensionContextFrom(ptr unsafe.Pointer) ExtensionContext {
	return ExtensionContext{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ec _ExtensionContextClass) Alloc() ExtensionContext {
	rv := objc.Send[ExtensionContext](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _ExtensionContextClass) New() ExtensionContext {
	rv := objc.Send[ExtensionContext](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ ExtensionContext) Init() ExtensionContext {
	rv := objc.Send[ExtensionContext](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ ExtensionContext) Autorelease() ExtensionContext {
	rv := objc.Send[ExtensionContext](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewExtensionContext creates a new ExtensionContext instance.
func NewExtensionContext() ExtensionContext {
	return getExtensionContextClass().New()
}



// Tells the host app to cancel the app extension request, with a supplied error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionContext/cancelRequest(withError:)
func (e_ ExtensionContext) CancelRequestWithError(error_ IError) {
	objc.Send[objc.ID](e_.ID, objc.Sel("cancelRequestWithError:"), error_)
}


// Tells the host app to complete the app extension request with an array of result items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionContext/completeRequest(returningItems:completionHandler:)
func (e_ ExtensionContext) CompleteRequestReturningItemsCompletionHandler(items objectivec.IObject, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("completeRequestReturningItems:completionHandler:"), items, completionHandler)
}


// Tells the host app to complete the app extension request with the specified broadcast information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionContext/completeRequest(withBroadcast:broadcastConfiguration:setupInfo:)
func (e_ ExtensionContext) CompleteRequestWithBroadcastURLBroadcastConfigurationSetupInfo(broadcastURL IURL, broadcastConfiguration objectivec.IObject, setupInfo IDictionary) {
	objc.Send[objc.ID](e_.ID, objc.Sel("completeRequestWithBroadcastURL:broadcastConfiguration:setupInfo:"), broadcastURL, broadcastConfiguration, setupInfo)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionContext/completeRequest(withBroadcast:setupInfo:)
func (e_ ExtensionContext) CompleteRequestWithBroadcastURLSetupInfo(broadcastURL IURL, setupInfo IDictionary) {
	objc.Send[objc.ID](e_.ID, objc.Sel("completeRequestWithBroadcastURL:setupInfo:"), broadcastURL, setupInfo)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionContext/dismissNotificationContentExtension()
func (e_ ExtensionContext) DismissNotificationContentExtension() {
	objc.Send[objc.ID](e_.ID, objc.Sel("dismissNotificationContentExtension"))
}


// Returns a human-readable string describing the data that SiriKit displays to the user when you handle an intent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionContext/interfaceParametersDescription()
func (e_ ExtensionContext) InterfaceParametersDescription() String {
	rv := objc.Send[String](e_.ID, objc.Sel("interfaceParametersDescription"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionContext/loadBroadcastingApplicationInfo(completion:)
func (e_ ExtensionContext) LoadBroadcastingApplicationInfoWithCompletion(handler unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("loadBroadcastingApplicationInfoWithCompletion:"), handler)
}


// Tells the system that the Notification Content app extension stopped playing a media file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionContext/mediaPlayingPaused()
func (e_ ExtensionContext) MediaPlayingPaused() {
	objc.Send[objc.ID](e_.ID, objc.Sel("mediaPlayingPaused"))
}


// Tells the system that the Notification Content app extension began playing a media file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionContext/mediaPlayingStarted()
func (e_ ExtensionContext) MediaPlayingStarted() {
	objc.Send[objc.ID](e_.ID, objc.Sel("mediaPlayingStarted"))
}


// Asks the system to open a URL on behalf of the currently running app extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionContext/open(_:completionHandler:)
func (e_ ExtensionContext) OpenURLCompletionHandler(URL IURL, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("openURL:completionHandler:"), URL, completionHandler)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionContext/performNotificationDefaultAction()
func (e_ ExtensionContext) PerformNotificationDefaultAction() {
	objc.Send[objc.ID](e_.ID, objc.Sel("performNotificationDefaultAction"))
}


// Returns the maximum size for the specified widget display mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionContext/widgetMaximumSize(for:)
func (e_ ExtensionContext) WidgetMaximumSizeForDisplayMode(displayMode unsafe.Pointer) coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](e_.ID, objc.Sel("widgetMaximumSizeForDisplayMode:"), displayMode)
	return rv
}


// The maximum size for a Siri hosted view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionContext/hostedViewMaximumAllowedSize
func (e_ ExtensionContext) HostedViewMaximumAllowedSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](e_.ID, objc.Sel("hostedViewMaximumAllowedSize"))
	return rv
}


// The minimum size for a Siri hosted view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionContext/hostedViewMinimumAllowedSize
func (e_ ExtensionContext) HostedViewMinimumAllowedSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](e_.ID, objc.Sel("hostedViewMinimumAllowedSize"))
	return rv
}


// The list of input objects associated with the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionContext/inputItems
func (e_ ExtensionContext) InputItems() objc.ID {
	rv := objc.Send[objc.ID](e_.ID, objc.Sel("inputItems"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionContext/notificationActions
func (e_ ExtensionContext) NotificationActions() []objc.ID {
	rv := objc.Send[[]objc.ID](e_.ID, objc.Sel("notificationActions"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionContext/notificationActions
func (e_ ExtensionContext) SetNotificationActions(value []objc.ID) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](e_.ID, objc.Sel("setNotificationActions:"), nsArray)
}


// The active display mode of the widget.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionContext/widgetActiveDisplayMode
func (e_ ExtensionContext) WidgetActiveDisplayMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("widgetActiveDisplayMode"))
	return rv
}


// The largest display mode the widget supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionContext/widgetLargestAvailableDisplayMode
func (e_ ExtensionContext) WidgetLargestAvailableDisplayMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("widgetLargestAvailableDisplayMode"))
	return rv
}


// The largest display mode the widget supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionContext/widgetLargestAvailableDisplayMode
func (e_ ExtensionContext) SetWidgetLargestAvailableDisplayMode(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setWidgetLargestAvailableDisplayMode:"), value)
}


// The extension items and errors key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensionitemsanderrorskey
func (e_ ExtensionContext) NSExtensionItemsAndErrorsKey() string {
	rv := objc.Send[string](e_.ID, objc.Sel("NSExtensionItemsAndErrorsKey"))
	return rv
}



