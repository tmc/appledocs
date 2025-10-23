// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AuthorizationPublicKeyCredentialPRFAssertionInput] class.
var (
	AuthorizationPublicKeyCredentialPRFAssertionInputClass     _AuthorizationPublicKeyCredentialPRFAssertionInputClass
	AuthorizationPublicKeyCredentialPRFAssertionInputClassOnce sync.Once
)

func getAuthorizationPublicKeyCredentialPRFAssertionInputClass() _AuthorizationPublicKeyCredentialPRFAssertionInputClass {
	AuthorizationPublicKeyCredentialPRFAssertionInputClassOnce.Do(func() {
		AuthorizationPublicKeyCredentialPRFAssertionInputClass = _AuthorizationPublicKeyCredentialPRFAssertionInputClass{objc.GetClass("ASAuthorizationPublicKeyCredentialPRFAssertionInput")}
	})
	return AuthorizationPublicKeyCredentialPRFAssertionInputClass
}

type _AuthorizationPublicKeyCredentialPRFAssertionInputClass struct {
	class objc.Class
}

// An interface definition for the [AuthorizationPublicKeyCredentialPRFAssertionInput] class.
type IAuthorizationPublicKeyCredentialPRFAssertionInput interface {
	objectivec.IObject
	// properties:
	PerCredentialInputValues() foundation.IDictionary /* already interface */
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialPRFAssertionInput-c.class
type AuthorizationPublicKeyCredentialPRFAssertionInput struct {
	objectivec.Object
}

// AuthorizationPublicKeyCredentialPRFAssertionInputFrom constructs a [AuthorizationPublicKeyCredentialPRFAssertionInput] from an unsafe.Pointer.
func AuthorizationPublicKeyCredentialPRFAssertionInputFrom(ptr unsafe.Pointer) AuthorizationPublicKeyCredentialPRFAssertionInput {
	return AuthorizationPublicKeyCredentialPRFAssertionInput{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AuthorizationPublicKeyCredentialPRFAssertionInputClass) Alloc() AuthorizationPublicKeyCredentialPRFAssertionInput {
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFAssertionInput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AuthorizationPublicKeyCredentialPRFAssertionInputClass) New() AuthorizationPublicKeyCredentialPRFAssertionInput {
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFAssertionInput](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationPublicKeyCredentialPRFAssertionInput) Init() AuthorizationPublicKeyCredentialPRFAssertionInput {
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFAssertionInput](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationPublicKeyCredentialPRFAssertionInput) Autorelease() AuthorizationPublicKeyCredentialPRFAssertionInput {
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFAssertionInput](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationPublicKeyCredentialPRFAssertionInput creates a new AuthorizationPublicKeyCredentialPRFAssertionInput instance.
func NewAuthorizationPublicKeyCredentialPRFAssertionInput() AuthorizationPublicKeyCredentialPRFAssertionInput {
	return getAuthorizationPublicKeyCredentialPRFAssertionInputClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialPRFAssertionInput-c.class/perCredentialInputValues
func (a_ AuthorizationPublicKeyCredentialPRFAssertionInput) PerCredentialInputValues() foundation.IDictionary /* already interface */ {
	rv := objc.Send[foundation.IDictionary](a_.ID, objc.Sel("perCredentialInputValues"))
	return rv
}



