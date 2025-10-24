// Code generated from Apple documentation for Accounts. DO NOT EDIT.

package accounts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ACAccountCredential] class.
var (
	ACAccountCredentialClass     _ACAccountCredentialClass
	ACAccountCredentialClassOnce sync.Once
)

func getACAccountCredentialClass() _ACAccountCredentialClass {
	ACAccountCredentialClassOnce.Do(func() {
		ACAccountCredentialClass = _ACAccountCredentialClass{objc.GetClass("ACAccountCredential")}
	})
	return ACAccountCredentialClass
}

type _ACAccountCredentialClass struct {
	class objc.Class
}

// An interface definition for the [ACAccountCredential] class.
type IACAccountCredential interface {
	objectivec.IObject
	// properties:
	OauthToken() objc.IObject /* cross-framework: NSString */
	SetOauthToken(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

// A credential object that encapsulates the information needed to authenticate a user.
//
// To create an account credential that uses the OAuth open authentication standard, use the method.


// A credential object that encapsulates the information needed to authenticate a user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccountCredential
type ACAccountCredential struct {
	objectivec.Object
}

// ACAccountCredentialFrom constructs a [ACAccountCredential] from an unsafe.Pointer.
//
// A credential object that encapsulates the information needed to authenticate a user.
func ACAccountCredentialFrom(ptr unsafe.Pointer) ACAccountCredential {
	return ACAccountCredential{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _ACAccountCredentialClass) Alloc() ACAccountCredential {
	rv := objc.Send[ACAccountCredential](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _ACAccountCredentialClass) New() ACAccountCredential {
	rv := objc.Send[ACAccountCredential](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ ACAccountCredential) Init() ACAccountCredential {
	rv := objc.Send[ACAccountCredential](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ ACAccountCredential) Autorelease() ACAccountCredential {
	rv := objc.Send[ACAccountCredential](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewACAccountCredential creates a new ACAccountCredential instance.
func NewACAccountCredential() ACAccountCredential {
	return getACAccountCredentialClass().New()
}



// Initializes an account credential using OAuth 2.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccountCredential/init(oAuth2Token:refreshToken:expiryDate:)
func NewACAccountCredentialWithOAuth2TokenRefreshTokenExpiryDate(token objc.IObject /* cross-framework: NSString */, refreshToken objc.IObject /* cross-framework: NSString */, expiryDate objc.IObject /* cross-framework: NSDate */) ACAccountCredential {
	instance := getACAccountCredentialClass().Alloc()
	rv := objc.Send[ACAccountCredential](instance.ID, objc.Sel("initWithOAuth2Token:refreshToken:expiryDate:"), token, refreshToken, expiryDate)
	rv.Autorelease()
	return rv
}


// Initializes an account credential using OAuth.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccountCredential/init(oAuthToken:tokenSecret:)
func NewACAccountCredentialWithOAuthTokenTokenSecret(token objc.IObject /* cross-framework: NSString */, secret objc.IObject /* cross-framework: NSString */) ACAccountCredential {
	instance := getACAccountCredentialClass().Alloc()
	rv := objc.Send[ACAccountCredential](instance.ID, objc.Sel("initWithOAuthToken:tokenSecret:"), token, secret)
	rv.Autorelease()
	return rv
}



// The token used for the credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccountCredential/oauthToken
func (a_ ACAccountCredential) OauthToken() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("oauthToken"))
	return rv
}


// The token used for the credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccountCredential/oauthToken
func (a_ ACAccountCredential) SetOauthToken(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOauthToken:"), value)
}


