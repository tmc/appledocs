// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [OneTimeCodeCredentialIdentity] class.
var (
	OneTimeCodeCredentialIdentityClass     _OneTimeCodeCredentialIdentityClass
	OneTimeCodeCredentialIdentityClassOnce sync.Once
)

func getOneTimeCodeCredentialIdentityClass() _OneTimeCodeCredentialIdentityClass {
	OneTimeCodeCredentialIdentityClassOnce.Do(func() {
		OneTimeCodeCredentialIdentityClass = _OneTimeCodeCredentialIdentityClass{objc.GetClass("ASOneTimeCodeCredentialIdentity")}
	})
	return OneTimeCodeCredentialIdentityClass
}

type _OneTimeCodeCredentialIdentityClass struct {
	class objc.Class
}

// An interface definition for the [OneTimeCodeCredentialIdentity] class.
type IOneTimeCodeCredentialIdentity interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASOneTimeCodeCredentialIdentity
type OneTimeCodeCredentialIdentity struct {
	objectivec.Object
}

// OneTimeCodeCredentialIdentityFrom constructs a [OneTimeCodeCredentialIdentity] from an unsafe.Pointer.
func OneTimeCodeCredentialIdentityFrom(ptr unsafe.Pointer) OneTimeCodeCredentialIdentity {
	return OneTimeCodeCredentialIdentity{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _OneTimeCodeCredentialIdentityClass) Alloc() OneTimeCodeCredentialIdentity {
	rv := objc.Send[OneTimeCodeCredentialIdentity](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _OneTimeCodeCredentialIdentityClass) New() OneTimeCodeCredentialIdentity {
	rv := objc.Send[OneTimeCodeCredentialIdentity](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OneTimeCodeCredentialIdentity) Init() OneTimeCodeCredentialIdentity {
	rv := objc.Send[OneTimeCodeCredentialIdentity](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OneTimeCodeCredentialIdentity) Autorelease() OneTimeCodeCredentialIdentity {
	rv := objc.Send[OneTimeCodeCredentialIdentity](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOneTimeCodeCredentialIdentity creates a new OneTimeCodeCredentialIdentity instance.
func NewOneTimeCodeCredentialIdentity() OneTimeCodeCredentialIdentity {
	return getOneTimeCodeCredentialIdentityClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASOneTimeCodeCredentialIdentity/label
func (o_ OneTimeCodeCredentialIdentity) Label() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("label"))
	return rv
}



