// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AuthorizationPublicKeyCredentialPRFAssertionInputValues] class.
var (
	AuthorizationPublicKeyCredentialPRFAssertionInputValuesClass     _AuthorizationPublicKeyCredentialPRFAssertionInputValuesClass
	AuthorizationPublicKeyCredentialPRFAssertionInputValuesClassOnce sync.Once
)

func getAuthorizationPublicKeyCredentialPRFAssertionInputValuesClass() _AuthorizationPublicKeyCredentialPRFAssertionInputValuesClass {
	AuthorizationPublicKeyCredentialPRFAssertionInputValuesClassOnce.Do(func() {
		AuthorizationPublicKeyCredentialPRFAssertionInputValuesClass = _AuthorizationPublicKeyCredentialPRFAssertionInputValuesClass{objc.GetClass("ASAuthorizationPublicKeyCredentialPRFAssertionInputValues")}
	})
	return AuthorizationPublicKeyCredentialPRFAssertionInputValuesClass
}

type _AuthorizationPublicKeyCredentialPRFAssertionInputValuesClass struct {
	class objc.Class
}

// An interface definition for the [AuthorizationPublicKeyCredentialPRFAssertionInputValues] class.
type IAuthorizationPublicKeyCredentialPRFAssertionInputValues interface {
	objectivec.IObject
	SaltInput2() foundation.NSData
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialPRFAssertionInputValues

type AuthorizationPublicKeyCredentialPRFAssertionInputValues struct {
	objectivec.Object
}

// AuthorizationPublicKeyCredentialPRFAssertionInputValuesFrom constructs a [AuthorizationPublicKeyCredentialPRFAssertionInputValues] from an unsafe.Pointer.
func AuthorizationPublicKeyCredentialPRFAssertionInputValuesFrom(ptr unsafe.Pointer) AuthorizationPublicKeyCredentialPRFAssertionInputValues {
	return AuthorizationPublicKeyCredentialPRFAssertionInputValues{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AuthorizationPublicKeyCredentialPRFAssertionInputValuesClass) Alloc() AuthorizationPublicKeyCredentialPRFAssertionInputValues {
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFAssertionInputValues](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AuthorizationPublicKeyCredentialPRFAssertionInputValuesClass) New() AuthorizationPublicKeyCredentialPRFAssertionInputValues {
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFAssertionInputValues](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationPublicKeyCredentialPRFAssertionInputValues) Init() AuthorizationPublicKeyCredentialPRFAssertionInputValues {
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFAssertionInputValues](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationPublicKeyCredentialPRFAssertionInputValues) Autorelease() AuthorizationPublicKeyCredentialPRFAssertionInputValues {
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFAssertionInputValues](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationPublicKeyCredentialPRFAssertionInputValues creates a new AuthorizationPublicKeyCredentialPRFAssertionInputValues instance.
func NewAuthorizationPublicKeyCredentialPRFAssertionInputValues() AuthorizationPublicKeyCredentialPRFAssertionInputValues {
	return getAuthorizationPublicKeyCredentialPRFAssertionInputValuesClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialPRFAssertionInputValues/saltInput2

func (a_ AuthorizationPublicKeyCredentialPRFAssertionInputValues) SaltInput2() foundation.NSData {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("saltInput2"))
	return rv
}



