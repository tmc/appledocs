// Code generated from Apple documentation for Accounts. DO NOT EDIT.

package accounts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ACAccountCredential */


/* debug [class_header]: Header for ACAccountCredential */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ACAccountCredential */
// An interface definition for the [ACAccountCredential] class.
type IACAccountCredential interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ACAccountCredential */
	// properties:
	OauthToken() objc.IObject /* cross-framework: NSString */
	SetOauthToken(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ACAccountCredential */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ACAccountCredential */
// Alloc allocates a new instance without initialization.
func (ac _ACAccountCredentialClass) Alloc() ACAccountCredential {
	rv := objc.Send[ACAccountCredential](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ACAccountCredential */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ACAccountCredential */

// Initializes an account credential using OAuth 2.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccountCredential/init(oAuth2Token:refreshToken:expiryDate:)
func NewACAccountCredentialWithOAuth2TokenRefreshTokenExpiryDate(token objc.IObject /* cross-framework: NSString */, refreshToken objc.IObject /* cross-framework: NSString */, expiryDate objc.IObject /* cross-framework: NSDate */) ACAccountCredential {
	instance := getACAccountCredentialClass().Alloc()
	rv := objc.Send[ACAccountCredential](instance.ID, objc.Sel("initWithOAuth2Token:refreshToken:expiryDate:"), token, refreshToken, expiryDate)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewACAccountCredentialWithOAuth2TokenRefreshTokenExpiryDate */


// Initializes an account credential using OAuth.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccountCredential/init(oAuthToken:tokenSecret:)
func NewACAccountCredentialWithOAuthTokenTokenSecret(token objc.IObject /* cross-framework: NSString */, secret objc.IObject /* cross-framework: NSString */) ACAccountCredential {
	instance := getACAccountCredentialClass().Alloc()
	rv := objc.Send[ACAccountCredential](instance.ID, objc.Sel("initWithOAuthToken:tokenSecret:"), token, secret)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewACAccountCredentialWithOAuthTokenTokenSecret */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ACAccountCredential */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ACAccountCredential */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ACAccountCredential */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ACAccountCredential */

// The token used for the credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccountCredential/oauthToken
func (a_ ACAccountCredential) OauthToken() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("oauthToken"))
	return rv
}/* debug [instance_properties/getter]: oauthToken */


// The token used for the credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccountCredential/oauthToken
func (a_ ACAccountCredential) SetOauthToken(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOauthToken:"), value)
}/* debug [instance_properties/setter]: oauthToken */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ACAccountCredential */


