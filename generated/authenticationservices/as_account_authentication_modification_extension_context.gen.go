// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [AccountAuthenticationModificationExtensionContext] class.
var (
	AccountAuthenticationModificationExtensionContextClass     _AccountAuthenticationModificationExtensionContextClass
	AccountAuthenticationModificationExtensionContextClassOnce sync.Once
)

func getAccountAuthenticationModificationExtensionContextClass() _AccountAuthenticationModificationExtensionContextClass {
	AccountAuthenticationModificationExtensionContextClassOnce.Do(func() {
		AccountAuthenticationModificationExtensionContextClass = _AccountAuthenticationModificationExtensionContextClass{objc.GetClass("ASAccountAuthenticationModificationExtensionContext")}
	})
	return AccountAuthenticationModificationExtensionContextClass
}

type _AccountAuthenticationModificationExtensionContextClass struct {
	class objc.Class
}

// An interface definition for the [AccountAuthenticationModificationExtensionContext] class.
type IAccountAuthenticationModificationExtensionContext interface {
	foundation.IExtensionContext
	// properties:
	ASExtensionLocalizedFailureReasonErrorKey() string /* primitive/slice/pointer. */
	// methods:
}

// An object that you interact with to change an account’s password or to upgrade to Sign in with Apple.


// An object that you interact with to change an account’s password or to upgrade to Sign in with Apple.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAccountAuthenticationModificationExtensionContext
type AccountAuthenticationModificationExtensionContext struct {
	foundation.ExtensionContext
}

// AccountAuthenticationModificationExtensionContextFrom constructs a [AccountAuthenticationModificationExtensionContext] from an unsafe.Pointer.
//
// An object that you interact with to change an account’s password or to upgrade to Sign in with Apple.
func AccountAuthenticationModificationExtensionContextFrom(ptr unsafe.Pointer) AccountAuthenticationModificationExtensionContext {
	return AccountAuthenticationModificationExtensionContext{
		ExtensionContext: foundation.ExtensionContextFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AccountAuthenticationModificationExtensionContextClass) Alloc() AccountAuthenticationModificationExtensionContext {
	rv := objc.Send[AccountAuthenticationModificationExtensionContext](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AccountAuthenticationModificationExtensionContextClass) New() AccountAuthenticationModificationExtensionContext {
	rv := objc.Send[AccountAuthenticationModificationExtensionContext](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccountAuthenticationModificationExtensionContext) Init() AccountAuthenticationModificationExtensionContext {
	rv := objc.Send[AccountAuthenticationModificationExtensionContext](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccountAuthenticationModificationExtensionContext) Autorelease() AccountAuthenticationModificationExtensionContext {
	rv := objc.Send[AccountAuthenticationModificationExtensionContext](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccountAuthenticationModificationExtensionContext creates a new AccountAuthenticationModificationExtensionContext instance.
func NewAccountAuthenticationModificationExtensionContext() AccountAuthenticationModificationExtensionContext {
	return getAccountAuthenticationModificationExtensionContextClass().New()
}



// A key that specifies a string value to show to the user when a request fails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asextensionlocalizedfailurereasonerrorkey
func (a_ AccountAuthenticationModificationExtensionContext) ASExtensionLocalizedFailureReasonErrorKey() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](a_.ID, objc.Sel("ASExtensionLocalizedFailureReasonErrorKey"))
	return rv
}



