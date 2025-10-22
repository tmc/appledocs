// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AuthorizationSingleSignOnCredential] class.
var (
	AuthorizationSingleSignOnCredentialClass     _AuthorizationSingleSignOnCredentialClass
	AuthorizationSingleSignOnCredentialClassOnce sync.Once
)

func getAuthorizationSingleSignOnCredentialClass() _AuthorizationSingleSignOnCredentialClass {
	AuthorizationSingleSignOnCredentialClassOnce.Do(func() {
		AuthorizationSingleSignOnCredentialClass = _AuthorizationSingleSignOnCredentialClass{objc.GetClass("ASAuthorizationSingleSignOnCredential")}
	})
	return AuthorizationSingleSignOnCredentialClass
}

type _AuthorizationSingleSignOnCredentialClass struct {
	class objc.Class
}

// An interface definition for the [AuthorizationSingleSignOnCredential] class.
type IAuthorizationSingleSignOnCredential interface {
	objectivec.IObject
	AuthorizedScopes() []string
	AccessToken() foundation.Data
	SetAccessToken(value foundation.IData)
	AuthenticatedResponse() foundation.HTTPURLResponse
	SetAuthenticatedResponse(value foundation.IHTTPURLResponse)
	IdentityToken() foundation.Data
	SetIdentityToken(value foundation.IData)
	State() string
	SetState(value string)
}

// A credential that results from a successful single sign-on (SSO) authentication.


// A credential that results from a successful single sign-on (SSO) authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSingleSignOnCredential

type AuthorizationSingleSignOnCredential struct {
	objectivec.Object
}

// AuthorizationSingleSignOnCredentialFrom constructs a [AuthorizationSingleSignOnCredential] from an unsafe.Pointer.
//
// A credential that results from a successful single sign-on (SSO) authentication.
func AuthorizationSingleSignOnCredentialFrom(ptr unsafe.Pointer) AuthorizationSingleSignOnCredential {
	return AuthorizationSingleSignOnCredential{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AuthorizationSingleSignOnCredentialClass) Alloc() AuthorizationSingleSignOnCredential {
	rv := objc.Send[AuthorizationSingleSignOnCredential](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AuthorizationSingleSignOnCredentialClass) New() AuthorizationSingleSignOnCredential {
	rv := objc.Send[AuthorizationSingleSignOnCredential](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationSingleSignOnCredential) Init() AuthorizationSingleSignOnCredential {
	rv := objc.Send[AuthorizationSingleSignOnCredential](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationSingleSignOnCredential) Autorelease() AuthorizationSingleSignOnCredential {
	rv := objc.Send[AuthorizationSingleSignOnCredential](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationSingleSignOnCredential creates a new AuthorizationSingleSignOnCredential instance.
func NewAuthorizationSingleSignOnCredential() AuthorizationSingleSignOnCredential {
	return getAuthorizationSingleSignOnCredentialClass().New()
}



// The contact information the user authorized your app to access.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSingleSignOnCredential/authorizedScopes

func (a_ AuthorizationSingleSignOnCredential) AuthorizedScopes() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("authorizedScopes"))
	return rv
}


// An access token used to get an identity token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationsinglesignoncredential/accesstoken

func (a_ AuthorizationSingleSignOnCredential) AccessToken() foundation.Data {
	rv := objc.Send[foundation.Data](a_.ID, objc.Sel("accessToken"))
	return rv
}


// An access token used to get an identity token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationsinglesignoncredential/accesstoken

func (a_ AuthorizationSingleSignOnCredential) SetAccessToken(value foundation.IData) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAccessToken:"), value)
}


// The complete response authentication, including technology-specific values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationsinglesignoncredential/authenticatedresponse

func (a_ AuthorizationSingleSignOnCredential) AuthenticatedResponse() foundation.HTTPURLResponse {
	rv := objc.Send[foundation.HTTPURLResponse](a_.ID, objc.Sel("authenticatedResponse"))
	return rv
}


// The complete response authentication, including technology-specific values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationsinglesignoncredential/authenticatedresponse

func (a_ AuthorizationSingleSignOnCredential) SetAuthenticatedResponse(value foundation.IHTTPURLResponse) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAuthenticatedResponse:"), value)
}


// A JSON Web Token (JWT) that securely communicates information about the user to your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationsinglesignoncredential/identitytoken

func (a_ AuthorizationSingleSignOnCredential) IdentityToken() foundation.Data {
	rv := objc.Send[foundation.Data](a_.ID, objc.Sel("identityToken"))
	return rv
}


// A JSON Web Token (JWT) that securely communicates information about the user to your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationsinglesignoncredential/identitytoken

func (a_ AuthorizationSingleSignOnCredential) SetIdentityToken(value foundation.IData) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIdentityToken:"), value)
}


// An arbitrary string that your app provided to the request that generated this credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationsinglesignoncredential/state

func (a_ AuthorizationSingleSignOnCredential) State() string {
	rv := objc.Send[string](a_.ID, objc.Sel("state"))
	return rv
}


// An arbitrary string that your app provided to the request that generated this credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationsinglesignoncredential/state

func (a_ AuthorizationSingleSignOnCredential) SetState(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setState:"), objc.String(value))
}



