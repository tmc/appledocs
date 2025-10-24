//go:build darwin && ios

// Code generated from Apple documentation for StoreKitTest. DO NOT EDIT.

package storekittest

import (
	"github.com/tmc/appledocs/generated/coretelephony"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

// iOS-only methods for AdTestPostbackResponse

// iOS-only properties

// A Boolean value that indicates whether the system successfully delivered the test postback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostbackResponse/didSucceed
func (a_ AdTestPostbackResponse) DidSucceed() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("didSucceed"))
	return rv
}
func (a_ AdTestPostbackResponse) SetDidSucceed(value bool) {
	a_.ID.Send(objc.RegisterName("setDidSucceed:"), value)
}

// An error the test session reports if sending a test postbacks fails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostbackResponse/error
func (a_ AdTestPostbackResponse) Error() objc.IObject /* cross-framework: Error */ {
	rv := objc.Send[coretelephony.Error](a_.ID, objc.Sel("error"))
	return rv
}
func (a_ AdTestPostbackResponse) SetError(value objc.IObject /* cross-framework: Error */) {
	a_.ID.Send(objc.RegisterName("setError:"), value)
}

// The HTTP response from the server receiving the test postback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostbackResponse/httpResponse
func (a_ AdTestPostbackResponse) HttpResponse() foundation.HTTPURLResponse {
	rv := objc.Send[foundation.HTTPURLResponse](a_.ID, objc.Sel("httpResponse"))
	return rv
}
func (a_ AdTestPostbackResponse) SetHttpResponse(value foundation.HTTPURLResponse) {
	a_.ID.Send(objc.RegisterName("setHttpResponse:"), value)
}
