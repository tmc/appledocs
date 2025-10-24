// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AuthorizationPublicKeyCredentialPRFAssertionOutput] class.
var (
	AuthorizationPublicKeyCredentialPRFAssertionOutputClass     _AuthorizationPublicKeyCredentialPRFAssertionOutputClass
	AuthorizationPublicKeyCredentialPRFAssertionOutputClassOnce sync.Once
)

func getAuthorizationPublicKeyCredentialPRFAssertionOutputClass() _AuthorizationPublicKeyCredentialPRFAssertionOutputClass {
	AuthorizationPublicKeyCredentialPRFAssertionOutputClassOnce.Do(func() {
		AuthorizationPublicKeyCredentialPRFAssertionOutputClass = _AuthorizationPublicKeyCredentialPRFAssertionOutputClass{objc.GetClass("ASAuthorizationPublicKeyCredentialPRFAssertionOutput")}
	})
	return AuthorizationPublicKeyCredentialPRFAssertionOutputClass
}

type _AuthorizationPublicKeyCredentialPRFAssertionOutputClass struct {
	class objc.Class
}

// An interface definition for the [AuthorizationPublicKeyCredentialPRFAssertionOutput] class.
type IAuthorizationPublicKeyCredentialPRFAssertionOutput interface {
	objectivec.IObject
	// properties:
	First() objc.IObject /* cross-framework: NSData */
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialPRFAssertionOutput-c.class
type AuthorizationPublicKeyCredentialPRFAssertionOutput struct {
	objectivec.Object
}

// AuthorizationPublicKeyCredentialPRFAssertionOutputFrom constructs a [AuthorizationPublicKeyCredentialPRFAssertionOutput] from an unsafe.Pointer.
func AuthorizationPublicKeyCredentialPRFAssertionOutputFrom(ptr unsafe.Pointer) AuthorizationPublicKeyCredentialPRFAssertionOutput {
	return AuthorizationPublicKeyCredentialPRFAssertionOutput{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AuthorizationPublicKeyCredentialPRFAssertionOutputClass) Alloc() AuthorizationPublicKeyCredentialPRFAssertionOutput {
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFAssertionOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AuthorizationPublicKeyCredentialPRFAssertionOutputClass) New() AuthorizationPublicKeyCredentialPRFAssertionOutput {
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFAssertionOutput](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationPublicKeyCredentialPRFAssertionOutput) Init() AuthorizationPublicKeyCredentialPRFAssertionOutput {
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFAssertionOutput](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationPublicKeyCredentialPRFAssertionOutput) Autorelease() AuthorizationPublicKeyCredentialPRFAssertionOutput {
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFAssertionOutput](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationPublicKeyCredentialPRFAssertionOutput creates a new AuthorizationPublicKeyCredentialPRFAssertionOutput instance.
func NewAuthorizationPublicKeyCredentialPRFAssertionOutput() AuthorizationPublicKeyCredentialPRFAssertionOutput {
	return getAuthorizationPublicKeyCredentialPRFAssertionOutputClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialPRFAssertionOutput-c.class/first
func (a_ AuthorizationPublicKeyCredentialPRFAssertionOutput) First() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("first"))
	return rv
}



