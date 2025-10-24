// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AuthorizationPublicKeyCredentialLargeBlobAssertionInput] class.
var (
	AuthorizationPublicKeyCredentialLargeBlobAssertionInputClass     _AuthorizationPublicKeyCredentialLargeBlobAssertionInputClass
	AuthorizationPublicKeyCredentialLargeBlobAssertionInputClassOnce sync.Once
)

func getAuthorizationPublicKeyCredentialLargeBlobAssertionInputClass() _AuthorizationPublicKeyCredentialLargeBlobAssertionInputClass {
	AuthorizationPublicKeyCredentialLargeBlobAssertionInputClassOnce.Do(func() {
		AuthorizationPublicKeyCredentialLargeBlobAssertionInputClass = _AuthorizationPublicKeyCredentialLargeBlobAssertionInputClass{objc.GetClass("ASAuthorizationPublicKeyCredentialLargeBlobAssertionInput")}
	})
	return AuthorizationPublicKeyCredentialLargeBlobAssertionInputClass
}

type _AuthorizationPublicKeyCredentialLargeBlobAssertionInputClass struct {
	class objc.Class
}

// An interface definition for the [AuthorizationPublicKeyCredentialLargeBlobAssertionInput] class.
type IAuthorizationPublicKeyCredentialLargeBlobAssertionInput interface {
	objectivec.IObject
	// properties:
	DataToWrite() objc.IObject /* cross-framework: NSData */
	SetDataToWrite(value objc.IObject /* cross-framework: NSData */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialLargeBlobAssertionInput-c.class
type AuthorizationPublicKeyCredentialLargeBlobAssertionInput struct {
	objectivec.Object
}

// AuthorizationPublicKeyCredentialLargeBlobAssertionInputFrom constructs a [AuthorizationPublicKeyCredentialLargeBlobAssertionInput] from an unsafe.Pointer.
func AuthorizationPublicKeyCredentialLargeBlobAssertionInputFrom(ptr unsafe.Pointer) AuthorizationPublicKeyCredentialLargeBlobAssertionInput {
	return AuthorizationPublicKeyCredentialLargeBlobAssertionInput{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AuthorizationPublicKeyCredentialLargeBlobAssertionInputClass) Alloc() AuthorizationPublicKeyCredentialLargeBlobAssertionInput {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobAssertionInput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AuthorizationPublicKeyCredentialLargeBlobAssertionInputClass) New() AuthorizationPublicKeyCredentialLargeBlobAssertionInput {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobAssertionInput](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationPublicKeyCredentialLargeBlobAssertionInput) Init() AuthorizationPublicKeyCredentialLargeBlobAssertionInput {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobAssertionInput](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationPublicKeyCredentialLargeBlobAssertionInput) Autorelease() AuthorizationPublicKeyCredentialLargeBlobAssertionInput {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobAssertionInput](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationPublicKeyCredentialLargeBlobAssertionInput creates a new AuthorizationPublicKeyCredentialLargeBlobAssertionInput instance.
func NewAuthorizationPublicKeyCredentialLargeBlobAssertionInput() AuthorizationPublicKeyCredentialLargeBlobAssertionInput {
	return getAuthorizationPublicKeyCredentialLargeBlobAssertionInputClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialLargeBlobAssertionInput-c.class/initWithOperation:
func NewAuthorizationPublicKeyCredentialLargeBlobAssertionInputWithOperation(operation AuthorizationPublicKeyCredentialLargeBlobAssertionOperation /* not a class type */) AuthorizationPublicKeyCredentialLargeBlobAssertionInput {
	instance := getAuthorizationPublicKeyCredentialLargeBlobAssertionInputClass().Alloc()
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobAssertionInput](instance.ID, objc.Sel("initWithOperation:"), operation)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialLargeBlobAssertionInput-c.class/dataToWrite
func (a_ AuthorizationPublicKeyCredentialLargeBlobAssertionInput) DataToWrite() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("dataToWrite"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialLargeBlobAssertionInput-c.class/dataToWrite
func (a_ AuthorizationPublicKeyCredentialLargeBlobAssertionInput) SetDataToWrite(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDataToWrite:"), value)
}


