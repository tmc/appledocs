// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [AuthorizationPublicKeyCredentialLargeBlobRegistrationInput] class.
var (
	AuthorizationPublicKeyCredentialLargeBlobRegistrationInputClass     _AuthorizationPublicKeyCredentialLargeBlobRegistrationInputClass
	AuthorizationPublicKeyCredentialLargeBlobRegistrationInputClassOnce sync.Once
)

func getAuthorizationPublicKeyCredentialLargeBlobRegistrationInputClass() _AuthorizationPublicKeyCredentialLargeBlobRegistrationInputClass {
	AuthorizationPublicKeyCredentialLargeBlobRegistrationInputClassOnce.Do(func() {
		AuthorizationPublicKeyCredentialLargeBlobRegistrationInputClass = _AuthorizationPublicKeyCredentialLargeBlobRegistrationInputClass{objc.GetClass("ASAuthorizationPublicKeyCredentialLargeBlobRegistrationInput")}
	})
	return AuthorizationPublicKeyCredentialLargeBlobRegistrationInputClass
}

type _AuthorizationPublicKeyCredentialLargeBlobRegistrationInputClass struct {
	class objc.Class
}

// An interface definition for the [AuthorizationPublicKeyCredentialLargeBlobRegistrationInput] class.
type IAuthorizationPublicKeyCredentialLargeBlobRegistrationInput interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialLargeBlobRegistrationInput-c.class
type AuthorizationPublicKeyCredentialLargeBlobRegistrationInput struct {
	objectivec.Object
}

// AuthorizationPublicKeyCredentialLargeBlobRegistrationInputFrom constructs a [AuthorizationPublicKeyCredentialLargeBlobRegistrationInput] from an unsafe.Pointer.
func AuthorizationPublicKeyCredentialLargeBlobRegistrationInputFrom(ptr unsafe.Pointer) AuthorizationPublicKeyCredentialLargeBlobRegistrationInput {
	return AuthorizationPublicKeyCredentialLargeBlobRegistrationInput{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AuthorizationPublicKeyCredentialLargeBlobRegistrationInputClass) Alloc() AuthorizationPublicKeyCredentialLargeBlobRegistrationInput {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobRegistrationInput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AuthorizationPublicKeyCredentialLargeBlobRegistrationInputClass) New() AuthorizationPublicKeyCredentialLargeBlobRegistrationInput {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobRegistrationInput](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationPublicKeyCredentialLargeBlobRegistrationInput) Init() AuthorizationPublicKeyCredentialLargeBlobRegistrationInput {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobRegistrationInput](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationPublicKeyCredentialLargeBlobRegistrationInput) Autorelease() AuthorizationPublicKeyCredentialLargeBlobRegistrationInput {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobRegistrationInput](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationPublicKeyCredentialLargeBlobRegistrationInput creates a new AuthorizationPublicKeyCredentialLargeBlobRegistrationInput instance.
func NewAuthorizationPublicKeyCredentialLargeBlobRegistrationInput() AuthorizationPublicKeyCredentialLargeBlobRegistrationInput {
	return getAuthorizationPublicKeyCredentialLargeBlobRegistrationInputClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialLargeBlobRegistrationInput-c.class/initWithSupportRequirement:
func NewAuthorizationPublicKeyCredentialLargeBlobRegistrationInputWithSupportRequirement(requirement unsafe.Pointer) AuthorizationPublicKeyCredentialLargeBlobRegistrationInput {
	instance := getAuthorizationPublicKeyCredentialLargeBlobRegistrationInputClass().Alloc()
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobRegistrationInput](instance.ID, objc.Sel("initWithSupportRequirement:"), requirement)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialLargeBlobRegistrationInput-c.class/supportRequirement
func (a_ AuthorizationPublicKeyCredentialLargeBlobRegistrationInput) SupportRequirement() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("supportRequirement"))
	return rv
}


// SetSupportRequirement sets the value of the supportRequirement property.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialLargeBlobRegistrationInput-c.class/supportRequirement
func (a_ AuthorizationPublicKeyCredentialLargeBlobRegistrationInput) SetSupportRequirement(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSupportRequirement:"), value)
}

