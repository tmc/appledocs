// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AuthorizationPublicKeyCredentialLargeBlobAssertionOutput] class.
var (
	AuthorizationPublicKeyCredentialLargeBlobAssertionOutputClass     _AuthorizationPublicKeyCredentialLargeBlobAssertionOutputClass
	AuthorizationPublicKeyCredentialLargeBlobAssertionOutputClassOnce sync.Once
)

func getAuthorizationPublicKeyCredentialLargeBlobAssertionOutputClass() _AuthorizationPublicKeyCredentialLargeBlobAssertionOutputClass {
	AuthorizationPublicKeyCredentialLargeBlobAssertionOutputClassOnce.Do(func() {
		AuthorizationPublicKeyCredentialLargeBlobAssertionOutputClass = _AuthorizationPublicKeyCredentialLargeBlobAssertionOutputClass{objc.GetClass("ASAuthorizationPublicKeyCredentialLargeBlobAssertionOutput")}
	})
	return AuthorizationPublicKeyCredentialLargeBlobAssertionOutputClass
}

type _AuthorizationPublicKeyCredentialLargeBlobAssertionOutputClass struct {
	class objc.Class
}

// An interface definition for the [AuthorizationPublicKeyCredentialLargeBlobAssertionOutput] class.
type IAuthorizationPublicKeyCredentialLargeBlobAssertionOutput interface {
	objectivec.IObject
	// properties:
	DidWrite() bool
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialLargeBlobAssertionOutput-c.class
type AuthorizationPublicKeyCredentialLargeBlobAssertionOutput struct {
	objectivec.Object
}

// AuthorizationPublicKeyCredentialLargeBlobAssertionOutputFrom constructs a [AuthorizationPublicKeyCredentialLargeBlobAssertionOutput] from an unsafe.Pointer.
func AuthorizationPublicKeyCredentialLargeBlobAssertionOutputFrom(ptr unsafe.Pointer) AuthorizationPublicKeyCredentialLargeBlobAssertionOutput {
	return AuthorizationPublicKeyCredentialLargeBlobAssertionOutput{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AuthorizationPublicKeyCredentialLargeBlobAssertionOutputClass) Alloc() AuthorizationPublicKeyCredentialLargeBlobAssertionOutput {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobAssertionOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AuthorizationPublicKeyCredentialLargeBlobAssertionOutputClass) New() AuthorizationPublicKeyCredentialLargeBlobAssertionOutput {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobAssertionOutput](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationPublicKeyCredentialLargeBlobAssertionOutput) Init() AuthorizationPublicKeyCredentialLargeBlobAssertionOutput {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobAssertionOutput](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationPublicKeyCredentialLargeBlobAssertionOutput) Autorelease() AuthorizationPublicKeyCredentialLargeBlobAssertionOutput {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobAssertionOutput](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationPublicKeyCredentialLargeBlobAssertionOutput creates a new AuthorizationPublicKeyCredentialLargeBlobAssertionOutput instance.
func NewAuthorizationPublicKeyCredentialLargeBlobAssertionOutput() AuthorizationPublicKeyCredentialLargeBlobAssertionOutput {
	return getAuthorizationPublicKeyCredentialLargeBlobAssertionOutputClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialLargeBlobAssertionOutput-c.class/didWrite
func (a_ AuthorizationPublicKeyCredentialLargeBlobAssertionOutput) DidWrite() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("didWrite"))
	return rv
}



