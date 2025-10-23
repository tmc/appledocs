// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AuthorizationProviderExtensionAuthorizationResult] class.
var (
	AuthorizationProviderExtensionAuthorizationResultClass     _AuthorizationProviderExtensionAuthorizationResultClass
	AuthorizationProviderExtensionAuthorizationResultClassOnce sync.Once
)

func getAuthorizationProviderExtensionAuthorizationResultClass() _AuthorizationProviderExtensionAuthorizationResultClass {
	AuthorizationProviderExtensionAuthorizationResultClassOnce.Do(func() {
		AuthorizationProviderExtensionAuthorizationResultClass = _AuthorizationProviderExtensionAuthorizationResultClass{objc.GetClass("ASAuthorizationProviderExtensionAuthorizationResult")}
	})
	return AuthorizationProviderExtensionAuthorizationResultClass
}

type _AuthorizationProviderExtensionAuthorizationResultClass struct {
	class objc.Class
}

// An interface definition for the [AuthorizationProviderExtensionAuthorizationResult] class.
type IAuthorizationProviderExtensionAuthorizationResult interface {
	objectivec.IObject
	HttpAuthorizationHeaders() string
	SetHttpAuthorizationHeaders(value string)
	HttpBody() foundation.Data
	SetHttpBody(value foundation.IData)
	HttpResponse() foundation.HTTPURLResponse
	SetHttpResponse(value foundation.IHTTPURLResponse)
}

// The result of an authorization request.


// The result of an authorization request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionAuthorizationResult
type AuthorizationProviderExtensionAuthorizationResult struct {
	objectivec.Object
}

// AuthorizationProviderExtensionAuthorizationResultFrom constructs a [AuthorizationProviderExtensionAuthorizationResult] from an unsafe.Pointer.
//
// The result of an authorization request.
func AuthorizationProviderExtensionAuthorizationResultFrom(ptr unsafe.Pointer) AuthorizationProviderExtensionAuthorizationResult {
	return AuthorizationProviderExtensionAuthorizationResult{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AuthorizationProviderExtensionAuthorizationResultClass) Alloc() AuthorizationProviderExtensionAuthorizationResult {
	rv := objc.Send[AuthorizationProviderExtensionAuthorizationResult](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AuthorizationProviderExtensionAuthorizationResultClass) New() AuthorizationProviderExtensionAuthorizationResult {
	rv := objc.Send[AuthorizationProviderExtensionAuthorizationResult](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationProviderExtensionAuthorizationResult) Init() AuthorizationProviderExtensionAuthorizationResult {
	rv := objc.Send[AuthorizationProviderExtensionAuthorizationResult](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationProviderExtensionAuthorizationResult) Autorelease() AuthorizationProviderExtensionAuthorizationResult {
	rv := objc.Send[AuthorizationProviderExtensionAuthorizationResult](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationProviderExtensionAuthorizationResult creates a new AuthorizationProviderExtensionAuthorizationResult instance.
func NewAuthorizationProviderExtensionAuthorizationResult() AuthorizationProviderExtensionAuthorizationResult {
	return getAuthorizationProviderExtensionAuthorizationResultClass().New()
}



// A dictionary of authorization HTTP headers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionauthorizationresult/httpauthorizationheaders
func (a_ AuthorizationProviderExtensionAuthorizationResult) HttpAuthorizationHeaders() string {
	rv := objc.Send[string](a_.ID, objc.Sel("httpAuthorizationHeaders"))
	return rv
}


// A dictionary of authorization HTTP headers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionauthorizationresult/httpauthorizationheaders
func (a_ AuthorizationProviderExtensionAuthorizationResult) SetHttpAuthorizationHeaders(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setHttpAuthorizationHeaders:"), objc.String(value))
}


// The HTTP response body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionauthorizationresult/httpbody
func (a_ AuthorizationProviderExtensionAuthorizationResult) HttpBody() foundation.Data {
	rv := objc.Send[foundation.Data](a_.ID, objc.Sel("httpBody"))
	return rv
}


// The HTTP response body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionauthorizationresult/httpbody
func (a_ AuthorizationProviderExtensionAuthorizationResult) SetHttpBody(value foundation.IData) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setHttpBody:"), value)
}


// The HTTP response for authentications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionauthorizationresult/httpresponse
func (a_ AuthorizationProviderExtensionAuthorizationResult) HttpResponse() foundation.HTTPURLResponse {
	rv := objc.Send[foundation.HTTPURLResponse](a_.ID, objc.Sel("httpResponse"))
	return rv
}


// The HTTP response for authentications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionauthorizationresult/httpresponse
func (a_ AuthorizationProviderExtensionAuthorizationResult) SetHttpResponse(value foundation.IHTTPURLResponse) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setHttpResponse:"), value)
}



