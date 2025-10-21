// Code generated from Apple documentation for StoreKitTest. DO NOT EDIT.

package storekittest

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AdTestPostbackResponse] class.
var (
	AdTestPostbackResponseClass     _AdTestPostbackResponseClass
	AdTestPostbackResponseClassOnce sync.Once
)

func getAdTestPostbackResponseClass() _AdTestPostbackResponseClass {
	AdTestPostbackResponseClassOnce.Do(func() {
		AdTestPostbackResponseClass = _AdTestPostbackResponseClass{objc.GetClass("SKAdTestPostbackResponse")}
	})
	return AdTestPostbackResponseClass
}

type _AdTestPostbackResponseClass struct {
	class objc.Class
}

// An interface definition for the [AdTestPostbackResponse] class.
type IAdTestPostbackResponse interface {
	objectivec.IObject
}

// The status and error information for a postback that the system sends in the testing environment.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostbackResponse
type AdTestPostbackResponse struct {
	objectivec.Object
}

// AdTestPostbackResponseFrom constructs a [AdTestPostbackResponse] from an unsafe.Pointer.
//
// The status and error information for a postback that the system sends in the testing environment.
func AdTestPostbackResponseFrom(ptr unsafe.Pointer) AdTestPostbackResponse {
	return AdTestPostbackResponse{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AdTestPostbackResponseClass) Alloc() AdTestPostbackResponse {
	rv := objc.Send[AdTestPostbackResponse](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AdTestPostbackResponseClass) New() AdTestPostbackResponse {
	rv := objc.Send[AdTestPostbackResponse](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AdTestPostbackResponse) Init() AdTestPostbackResponse {
	rv := objc.Send[AdTestPostbackResponse](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AdTestPostbackResponse) Autorelease() AdTestPostbackResponse {
	rv := objc.Send[AdTestPostbackResponse](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAdTestPostbackResponse creates a new AdTestPostbackResponse instance.
func NewAdTestPostbackResponse() AdTestPostbackResponse {
	return getAdTestPostbackResponseClass().New()
}


// A Boolean value that indicates whether the system successfully delivered the test postback.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostbackResponse/didSucceed
func (a_ AdTestPostbackResponse) DidSucceed() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("didSucceed"))
	return rv
}


// SetDidSucceed sets the value of the didSucceed property.
// A Boolean value that indicates whether the system successfully delivered the test postback.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostbackResponse/didSucceed
func (a_ AdTestPostbackResponse) SetDidSucceed(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDidSucceed:"), value)
}

// An error the test session reports if sending a test postbacks fails.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostbackResponse/error
func (a_ AdTestPostbackResponse) Error() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("error"))
	return rv
}


// SetError sets the value of the error property.
// An error the test session reports if sending a test postbacks fails.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostbackResponse/error
func (a_ AdTestPostbackResponse) SetError(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setError:"), value)
}

// The HTTP response from the server receiving the test postback.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostbackResponse/httpResponse
func (a_ AdTestPostbackResponse) HttpResponse() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("httpResponse"))
	return rv
}


// SetHttpResponse sets the value of the httpResponse property.
// The HTTP response from the server receiving the test postback.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostbackResponse/httpResponse
func (a_ AdTestPostbackResponse) SetHttpResponse(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setHttpResponse:"), value)
}



