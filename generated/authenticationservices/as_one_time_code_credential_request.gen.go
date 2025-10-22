// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [OneTimeCodeCredentialRequest] class.
var (
	OneTimeCodeCredentialRequestClass     _OneTimeCodeCredentialRequestClass
	OneTimeCodeCredentialRequestClassOnce sync.Once
)

func getOneTimeCodeCredentialRequestClass() _OneTimeCodeCredentialRequestClass {
	OneTimeCodeCredentialRequestClassOnce.Do(func() {
		OneTimeCodeCredentialRequestClass = _OneTimeCodeCredentialRequestClass{objc.GetClass("ASOneTimeCodeCredentialRequest")}
	})
	return OneTimeCodeCredentialRequestClass
}

type _OneTimeCodeCredentialRequestClass struct {
	class objc.Class
}

// An interface definition for the [OneTimeCodeCredentialRequest] class.
type IOneTimeCodeCredentialRequest interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASOneTimeCodeCredentialRequest

type OneTimeCodeCredentialRequest struct {
	objectivec.Object
}

// OneTimeCodeCredentialRequestFrom constructs a [OneTimeCodeCredentialRequest] from an unsafe.Pointer.
func OneTimeCodeCredentialRequestFrom(ptr unsafe.Pointer) OneTimeCodeCredentialRequest {
	return OneTimeCodeCredentialRequest{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _OneTimeCodeCredentialRequestClass) Alloc() OneTimeCodeCredentialRequest {
	rv := objc.Send[OneTimeCodeCredentialRequest](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _OneTimeCodeCredentialRequestClass) New() OneTimeCodeCredentialRequest {
	rv := objc.Send[OneTimeCodeCredentialRequest](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OneTimeCodeCredentialRequest) Init() OneTimeCodeCredentialRequest {
	rv := objc.Send[OneTimeCodeCredentialRequest](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OneTimeCodeCredentialRequest) Autorelease() OneTimeCodeCredentialRequest {
	rv := objc.Send[OneTimeCodeCredentialRequest](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOneTimeCodeCredentialRequest creates a new OneTimeCodeCredentialRequest instance.
func NewOneTimeCodeCredentialRequest() OneTimeCodeCredentialRequest {
	return getOneTimeCodeCredentialRequestClass().New()
}




