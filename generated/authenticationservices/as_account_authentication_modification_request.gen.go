// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AccountAuthenticationModificationRequest] class.
var (
	AccountAuthenticationModificationRequestClass     _AccountAuthenticationModificationRequestClass
	AccountAuthenticationModificationRequestClassOnce sync.Once
)

func getAccountAuthenticationModificationRequestClass() _AccountAuthenticationModificationRequestClass {
	AccountAuthenticationModificationRequestClassOnce.Do(func() {
		AccountAuthenticationModificationRequestClass = _AccountAuthenticationModificationRequestClass{objc.GetClass("ASAccountAuthenticationModificationRequest")}
	})
	return AccountAuthenticationModificationRequestClass
}

type _AccountAuthenticationModificationRequestClass struct {
	class objc.Class
}

// An interface definition for the [AccountAuthenticationModificationRequest] class.
type IAccountAuthenticationModificationRequest interface {
	objectivec.IObject
}

// A request to modify an account’s authentication properties.
//
// To initiate an account authentication modification request from your app, use either or .
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAccountAuthenticationModificationRequest
type AccountAuthenticationModificationRequest struct {
	objectivec.Object
}

// AccountAuthenticationModificationRequestFrom constructs a [AccountAuthenticationModificationRequest] from an unsafe.Pointer.
//
// A request to modify an account’s authentication properties.
func AccountAuthenticationModificationRequestFrom(ptr unsafe.Pointer) AccountAuthenticationModificationRequest {
	return AccountAuthenticationModificationRequest{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AccountAuthenticationModificationRequestClass) Alloc() AccountAuthenticationModificationRequest {
	rv := objc.Send[AccountAuthenticationModificationRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AccountAuthenticationModificationRequestClass) New() AccountAuthenticationModificationRequest {
	rv := objc.Send[AccountAuthenticationModificationRequest](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccountAuthenticationModificationRequest) Init() AccountAuthenticationModificationRequest {
	rv := objc.Send[AccountAuthenticationModificationRequest](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccountAuthenticationModificationRequest) Autorelease() AccountAuthenticationModificationRequest {
	rv := objc.Send[AccountAuthenticationModificationRequest](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccountAuthenticationModificationRequest creates a new AccountAuthenticationModificationRequest instance.
func NewAccountAuthenticationModificationRequest() AccountAuthenticationModificationRequest {
	return getAccountAuthenticationModificationRequestClass().New()
}




