// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ExtensionContext] class.
var ExtensionContextClass objc.Class

func init() {
	ExtensionContextClass = objc.GetClass("NSExtensionContext")
}

type ExtensionContext struct {
	objc.ID
}

func ExtensionContextFrom(ptr unsafe.Pointer) ExtensionContext {
	return ExtensionContext{
		ID: objc.ID(ptr),
	}
}


// Tells the host app to complete the app extension request with the specified broadcast information. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSExtensionContext/completeRequest(withBroadcast:broadcastConfiguration:setupInfo:)
func (e_ ExtensionContext) CompleteRequestWithBroadcastURLBroadcastConfigurationSetupInfo(broadcastURL unsafe.Pointer, broadcastConfiguration unsafe.Pointer, setupInfo unsafe.Pointer) {
	sel := objc.RegisterName("completeRequestWithBroadcastURL:broadcastConfiguration:setupInfo:")
	e_.ID.Send(sel, broadcastURL, broadcastConfiguration, setupInfo)
}
//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSExtensionContext/completeRequest(withBroadcast:setupInfo:)
func (e_ ExtensionContext) CompleteRequestWithBroadcastURLSetupInfo(broadcastURL unsafe.Pointer, setupInfo unsafe.Pointer) {
	sel := objc.RegisterName("completeRequestWithBroadcastURL:setupInfo:")
	e_.ID.Send(sel, broadcastURL, setupInfo)
}
//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSExtensionContext/dismissNotificationContentExtension()
func (e_ ExtensionContext) DismissNotificationContentExtension() {
	sel := objc.RegisterName("dismissNotificationContentExtension")
	e_.ID.Send(sel)
}
// Returns a human-readable string describing the data that SiriKit displays to the user when you handle an intent. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSExtensionContext/interfaceParametersDescription()
func (e_ ExtensionContext) InterfaceParametersDescription() unsafe.Pointer {
	sel := objc.RegisterName("interfaceParametersDescription")
	ret := e_.ID.Send(sel)
	return unsafe.Pointer(ret)
}
//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSExtensionContext/loadBroadcastingApplicationInfo(completion:)
func (e_ ExtensionContext) LoadBroadcastingApplicationInfoWithCompletion(handler unsafe.Pointer) {
	sel := objc.RegisterName("loadBroadcastingApplicationInfoWithCompletion:")
	e_.ID.Send(sel, handler)
}
// Tells the system that the Notification Content app extension stopped playing a media file. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSExtensionContext/mediaPlayingPaused()
func (e_ ExtensionContext) MediaPlayingPaused() {
	sel := objc.RegisterName("mediaPlayingPaused")
	e_.ID.Send(sel)
}
// Tells the system that the Notification Content app extension began playing a media file. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSExtensionContext/mediaPlayingStarted()
func (e_ ExtensionContext) MediaPlayingStarted() {
	sel := objc.RegisterName("mediaPlayingStarted")
	e_.ID.Send(sel)
}
// Asks the system to open a URL on behalf of the currently running app extension. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSExtensionContext/open(_:completionHandler:)
func (e_ ExtensionContext) OpenURLCompletionHandler(URL unsafe.Pointer, completionHandler unsafe.Pointer) {
	sel := objc.RegisterName("openURL:completionHandler:")
	e_.ID.Send(sel, URL, completionHandler)
}
//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSExtensionContext/performNotificationDefaultAction()
func (e_ ExtensionContext) PerformNotificationDefaultAction() {
	sel := objc.RegisterName("performNotificationDefaultAction")
	e_.ID.Send(sel)
}
// Returns the maximum size for the specified widget display mode. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSExtensionContext/widgetMaximumSize(for:)
func (e_ ExtensionContext) WidgetMaximumSizeForDisplayMode(displayMode unsafe.Pointer) Size {
	sel := objc.RegisterName("widgetMaximumSizeForDisplayMode:")
	ret := e_.ID.Send(sel, displayMode)
	return Size(ret)
}

