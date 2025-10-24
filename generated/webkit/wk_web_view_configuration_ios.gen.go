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





