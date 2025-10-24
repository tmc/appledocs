//go:build darwin && ios

// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objc"
)

// iOS-only methods for WebView

// iOS-only properties

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/findInteraction
func (w_ WebView) FindInteraction() FindInteraction /* not a class type */ {
	rv := objc.Send[FindInteraction](w_.ID, objc.Sel("findInteraction"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/isFindInteractionEnabled
func (w_ WebView) FindInteractionEnabled() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("findInteractionEnabled"))
	return rv
}
func (w_ WebView) SetFindInteractionEnabled(value bool) {
	w_.ID.Send(objc.RegisterName("setFindInteractionEnabled:"), value)
}

// The scroll view associated with the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebView/scrollView
func (w_ WebView) ScrollView() appkit.ScrollView {
	rv := objc.Send[appkit.ScrollView](w_.ID, objc.Sel("scrollView"))
	return rv
}
