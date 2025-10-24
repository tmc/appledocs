//go:build darwin && ios

// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for SFAddToHomeScreenInfo


// iOS-only properties

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFAddToHomeScreenInfo/manifest
func (s_ SFAddToHomeScreenInfo) Manifest() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("manifest"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFAddToHomeScreenInfo/websiteCookies
func (s_ SFAddToHomeScreenInfo) WebsiteCookies() []objc.IObject /* cross-framework: HTTPCookie */ {
	rv := objc.Send[[]foundation.HTTPCookie](s_.ID, objc.Sel("websiteCookies"))
	return rv
}
func (s_ SFAddToHomeScreenInfo) SetWebsiteCookies(value []objc.IObject /* cross-framework: HTTPCookie */) {
	s_.ID.Send(objc.RegisterName("setWebsiteCookies:"), value)
}




