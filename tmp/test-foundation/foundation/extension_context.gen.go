// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var ExtensionContextClass _ExtensionContextClass

func init() {
	ExtensionContextClass = _ExtensionContextClass{objc.GetClass("NSExtensionContext")}
}

type _ExtensionContextClass struct {
	class objc.Class
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
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionContext/completeRequest(withBroadcast:broadcastConfiguration:setupInfo:)
func (e_ ExtensionContext) CompleteRequestWithBroadcastURLBroadcastConfigurationSetupInfo(broadcastURL unsafe.Pointer, broadcastConfiguration unsafe.Pointer, setupInfo unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("completeRequestWithBroadcastURL:broadcastConfiguration:setupInfo:"), broadcastURL, broadcastConfiguration, setupInfo)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionContext/completeRequest(withBroadcast:setupInfo:)
func (e_ ExtensionContext) CompleteRequestWithBroadcastURLSetupInfo(broadcastURL unsafe.Pointer, setupInfo unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("completeRequestWithBroadcastURL:setupInfo:"), broadcastURL, setupInfo)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionContext/dismissNotificationContentExtension()
func (e_ ExtensionContext) DismissNotificationContentExtension() {
	objc.Send[objc.ID](e_.ID, objc.Sel("dismissNotificationContentExtension"))
}
// Returns a human-readable string describing the data that SiriKit displays to the user when you handle an intent. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionContext/interfaceParametersDescription()
func (e_ ExtensionContext) InterfaceParametersDescription() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("interfaceParametersDescription"))
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionContext/loadBroadcastingApplicationInfo(completion:)
func (e_ ExtensionContext) LoadBroadcastingApplicationInfoWithCompletion(handler unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("loadBroadcastingApplicationInfoWithCompletion:"), handler)
}
// Tells the system that the Notification Content app extension stopped playing a media file. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionContext/mediaPlayingPaused()
func (e_ ExtensionContext) MediaPlayingPaused() {
	objc.Send[objc.ID](e_.ID, objc.Sel("mediaPlayingPaused"))
}
// Tells the system that the Notification Content app extension began playing a media file. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionContext/mediaPlayingStarted()
func (e_ ExtensionContext) MediaPlayingStarted() {
	objc.Send[objc.ID](e_.ID, objc.Sel("mediaPlayingStarted"))
}
// Asks the system to open a URL on behalf of the currently running app extension. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionContext/open(_:completionHandler:)
func (e_ ExtensionContext) OpenURLCompletionHandler(URL unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("openURL:completionHandler:"), URL, completionHandler)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionContext/performNotificationDefaultAction()
func (e_ ExtensionContext) PerformNotificationDefaultAction() {
	objc.Send[objc.ID](e_.ID, objc.Sel("performNotificationDefaultAction"))
}
// Returns the maximum size for the specified widget display mode. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionContext/widgetMaximumSize(for:)
func (e_ ExtensionContext) WidgetMaximumSizeForDisplayMode(displayMode unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("widgetMaximumSizeForDisplayMode:"), displayMode)
	return rv
}


