// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [OneTimeCodeCredential] class.
var (
	OneTimeCodeCredentialClass     _OneTimeCodeCredentialClass
	OneTimeCodeCredentialClassOnce sync.Once
)

func getOneTimeCodeCredentialClass() _OneTimeCodeCredentialClass {
	OneTimeCodeCredentialClassOnce.Do(func() {
		OneTimeCodeCredentialClass = _OneTimeCodeCredentialClass{objc.GetClass("ASOneTimeCodeCredential")}
	})
	return OneTimeCodeCredentialClass
}

type _OneTimeCodeCredentialClass struct {
	class objc.Class
}

// An interface definition for the [OneTimeCodeCredential] class.
type IOneTimeCodeCredential interface {
	objectivec.IObject
}

// A one-time passcode (OTP) credential.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASOneTimeCodeCredential
type OneTimeCodeCredential struct {
	objectivec.Object
}

// OneTimeCodeCredentialFrom constructs a [OneTimeCodeCredential] from an unsafe.Pointer.
//
// A one-time passcode (OTP) credential.
func OneTimeCodeCredentialFrom(ptr unsafe.Pointer) OneTimeCodeCredential {
	return OneTimeCodeCredential{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _OneTimeCodeCredentialClass) Alloc() OneTimeCodeCredential {
	rv := objc.Send[OneTimeCodeCredential](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _OneTimeCodeCredentialClass) New() OneTimeCodeCredential {
	rv := objc.Send[OneTimeCodeCredential](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OneTimeCodeCredential) Init() OneTimeCodeCredential {
	rv := objc.Send[OneTimeCodeCredential](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OneTimeCodeCredential) Autorelease() OneTimeCodeCredential {
	rv := objc.Send[OneTimeCodeCredential](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOneTimeCodeCredential creates a new OneTimeCodeCredential instance.
func NewOneTimeCodeCredential() OneTimeCodeCredential {
	return getOneTimeCodeCredentialClass().New()
}


// The one-time passcode.
//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asonetimecodecredential/code
func (o_ OneTimeCodeCredential) Code() appkit.string {
	rv := objc.Send[appkit.string](o_.ID, objc.Sel("code"))
	return rv
}


// SetCode sets the value of the code property.
// The one-time passcode.

//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asonetimecodecredential/code
func (o_ OneTimeCodeCredential) SetCode(value appkit.string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setCode:"), value)
}



