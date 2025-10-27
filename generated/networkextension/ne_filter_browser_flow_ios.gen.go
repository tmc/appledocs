//go:build darwin && ios

// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// iOS-only methods for NEFilterBrowserFlow


// iOS-only properties

// A URL of the web page that’s responsible for the flow’s creation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterBrowserFlow/parentURL
func (n_ NEFilterBrowserFlow) ParentURL() foundation.foundation.INSURL {
	rv := objc.Send[foundation.NSURL](n_.ID, objc.Sel("parentURL"))
	return rv
}

// An HTTP request of the flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterBrowserFlow/request
func (n_ NEFilterBrowserFlow) Request() foundation.URLRequest {
	rv := objc.Send[foundation.URLRequest](n_.ID, objc.Sel("request"))
	return rv
}

// An HTTP response of the flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterBrowserFlow/response
func (n_ NEFilterBrowserFlow) Response() foundation.URLResponse {
	rv := objc.Send[foundation.URLResponse](n_.ID, objc.Sel("response"))
	return rv
}





