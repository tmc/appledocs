// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AuthorizationPublicKeyCredentialPRFRegistrationInput] class.
var (
	AuthorizationPublicKeyCredentialPRFRegistrationInputClass     _AuthorizationPublicKeyCredentialPRFRegistrationInputClass
	AuthorizationPublicKeyCredentialPRFRegistrationInputClassOnce sync.Once
)

func getAuthorizationPublicKeyCredentialPRFRegistrationInputClass() _AuthorizationPublicKeyCredentialPRFRegistrationInputClass {
	AuthorizationPublicKeyCredentialPRFRegistrationInputClassOnce.Do(func() {
		AuthorizationPublicKeyCredentialPRFRegistrationInputClass = _AuthorizationPublicKeyCredentialPRFRegistrationInputClass{objc.GetClass("ASAuthorizationPublicKeyCredentialPRFRegistrationInput")}
	})
	return AuthorizationPublicKeyCredentialPRFRegistrationInputClass
}

type _AuthorizationPublicKeyCredentialPRFRegistrationInputClass struct {
	class objc.Class
}

// An interface definition for the [AuthorizationPublicKeyCredentialPRFRegistrationInput] class.
type IAuthorizationPublicKeyCredentialPRFRegistrationInput interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialPRFRegistrationInput-c.class
type AuthorizationPublicKeyCredentialPRFRegistrationInput struct {
	objectivec.Object
}

// AuthorizationPublicKeyCredentialPRFRegistrationInputFrom constructs a [AuthorizationPublicKeyCredentialPRFRegistrationInput] from an unsafe.Pointer.
func AuthorizationPublicKeyCredentialPRFRegistrationInputFrom(ptr unsafe.Pointer) AuthorizationPublicKeyCredentialPRFRegistrationInput {
	return AuthorizationPublicKeyCredentialPRFRegistrationInput{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AuthorizationPublicKeyCredentialPRFRegistrationInputClass) Alloc() AuthorizationPublicKeyCredentialPRFRegistrationInput {
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFRegistrationInput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AuthorizationPublicKeyCredentialPRFRegistrationInputClass) New() AuthorizationPublicKeyCredentialPRFRegistrationInput {
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFRegistrationInput](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationPublicKeyCredentialPRFRegistrationInput) Init() AuthorizationPublicKeyCredentialPRFRegistrationInput {
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFRegistrationInput](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationPublicKeyCredentialPRFRegistrationInput) Autorelease() AuthorizationPublicKeyCredentialPRFRegistrationInput {
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFRegistrationInput](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationPublicKeyCredentialPRFRegistrationInput creates a new AuthorizationPublicKeyCredentialPRFRegistrationInput instance.
func NewAuthorizationPublicKeyCredentialPRFRegistrationInput() AuthorizationPublicKeyCredentialPRFRegistrationInput {
	return getAuthorizationPublicKeyCredentialPRFRegistrationInputClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialPRFRegistrationInput-c.class/initWithInputValues:
func NewAuthorizationPublicKeyCredentialPRFRegistrationInputWithInputValues(inputValues unsafe.Pointer) AuthorizationPublicKeyCredentialPRFRegistrationInput {
	instance := getAuthorizationPublicKeyCredentialPRFRegistrationInputClass().Alloc()
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFRegistrationInput](instance.ID, objc.Sel("initWithInputValues:"), inputValues)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialPRFRegistrationInput-c.class/checkForSupport
func (ac _AuthorizationPublicKeyCredentialPRFRegistrationInputClass) CheckForSupport() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("checkForSupport"))
	return rv
}


