//go:build darwin && ios

// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for WebViewConfiguration


// iOS-only properties

// A Boolean value that indicates whether HTML5 videos play inline or use the native full-screen controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/allowsInlineMediaPlayback
func (w_ WebViewConfiguration) AllowsInlineMediaPlayback() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("allowsInlineMediaPlayback"))
	return rv
}
func (w_ WebViewConfiguration) SetAllowsInlineMediaPlayback(value bool) {
	w_.ID.Send(objc.RegisterName("setAllowsInlineMediaPlayback:"), value)
}

// A Boolean value that indicates whether HTML5 videos can play Picture in Picture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/allowsPictureInPictureMediaPlayback
func (w_ WebViewConfiguration) AllowsPictureInPictureMediaPlayback() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("allowsPictureInPictureMediaPlayback"))
	return rv
}
func (w_ WebViewConfiguration) SetAllowsPictureInPictureMediaPlayback(value bool) {
	w_.ID.Send(objc.RegisterName("setAllowsPictureInPictureMediaPlayback:"), value)
}

// The types of data detectors to apply to the web view’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/dataDetectorTypes
func (w_ WebViewConfiguration) DataDetectorTypes() DataDetectorTypes {
	rv := objc.Send[DataDetectorTypes](w_.ID, objc.Sel("dataDetectorTypes"))
	return rv
}
func (w_ WebViewConfiguration) SetDataDetectorTypes(value DataDetectorTypes) {
	w_.ID.Send(objc.RegisterName("setDataDetectorTypes:"), value)
}

// A Boolean value that determines whether a web view allows scaling of the webpage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/ignoresViewportScaleLimits
func (w_ WebViewConfiguration) IgnoresViewportScaleLimits() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("ignoresViewportScaleLimits"))
	return rv
}
func (w_ WebViewConfiguration) SetIgnoresViewportScaleLimits(value bool) {
	w_.ID.Send(objc.RegisterName("setIgnoresViewportScaleLimits:"), value)
}

// Deprecated property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/mediaPlaybackAllowsAirPlay
func (w_ WebViewConfiguration) MediaPlaybackAllowsAirPlay() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("mediaPlaybackAllowsAirPlay"))
	return rv
}
func (w_ WebViewConfiguration) SetMediaPlaybackAllowsAirPlay(value bool) {
	w_.ID.Send(objc.RegisterName("setMediaPlaybackAllowsAirPlay:"), value)
}

// Deprecated property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/mediaPlaybackRequiresUserAction
func (w_ WebViewConfiguration) MediaPlaybackRequiresUserAction() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("mediaPlaybackRequiresUserAction"))
	return rv
}
func (w_ WebViewConfiguration) SetMediaPlaybackRequiresUserAction(value bool) {
	w_.ID.Send(objc.RegisterName("setMediaPlaybackRequiresUserAction:"), value)
}

// A Boolean value that indicates whether HTML5 videos require the user to start playing them ( ) or whether the videos play automatically ( ).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/requiresUserActionForMediaPlayback
func (w_ WebViewConfiguration) RequiresUserActionForMediaPlayback() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("requiresUserActionForMediaPlayback"))
	return rv
}
func (w_ WebViewConfiguration) SetRequiresUserActionForMediaPlayback(value bool) {
	w_.ID.Send(objc.RegisterName("setRequiresUserActionForMediaPlayback:"), value)
}

// The level of granularity with which the user can interactively select web view content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/selectionGranularity
func (w_ WebViewConfiguration) SelectionGranularity() SelectionGranularity {
	rv := objc.Send[SelectionGranularity](w_.ID, objc.Sel("selectionGranularity"))
	return rv
}
func (w_ WebViewConfiguration) SetSelectionGranularity(value SelectionGranularity) {
	w_.ID.Send(objc.RegisterName("setSelectionGranularity:"), value)
}





