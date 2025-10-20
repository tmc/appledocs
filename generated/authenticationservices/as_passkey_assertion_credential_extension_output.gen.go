// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PasskeyAssertionCredentialExtensionOutput] class.
var (
	PasskeyAssertionCredentialExtensionOutputClass     _PasskeyAssertionCredentialExtensionOutputClass
	PasskeyAssertionCredentialExtensionOutputClassOnce sync.Once
)

func getPasskeyAssertionCredentialExtensionOutputClass() _PasskeyAssertionCredentialExtensionOutputClass {
	PasskeyAssertionCredentialExtensionOutputClassOnce.Do(func() {
		PasskeyAssertionCredentialExtensionOutputClass = _PasskeyAssertionCredentialExtensionOutputClass{objc.GetClass("ASPasskeyAssertionCredentialExtensionOutput")}
	})
	return PasskeyAssertionCredentialExtensionOutputClass
}

type _PasskeyAssertionCredentialExtensionOutputClass struct {
	class objc.Class
}

// An interface definition for the [PasskeyAssertionCredentialExtensionOutput] class.
type IPasskeyAssertionCredentialExtensionOutput interface {
	objectivec.IObject
}

// This class encapsulates output for various WebAuthn extensions used during passkey assertion.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyAssertionCredentialExtensionOutput-c.class
type PasskeyAssertionCredentialExtensionOutput struct {
	objectivec.Object
}

// PasskeyAssertionCredentialExtensionOutputFrom constructs a [PasskeyAssertionCredentialExtensionOutput] from an unsafe.Pointer.
//
// This class encapsulates output for various WebAuthn extensions used during passkey assertion.
func PasskeyAssertionCredentialExtensionOutputFrom(ptr unsafe.Pointer) PasskeyAssertionCredentialExtensionOutput {
	return PasskeyAssertionCredentialExtensionOutput{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PasskeyAssertionCredentialExtensionOutputClass) Alloc() PasskeyAssertionCredentialExtensionOutput {
	rv := objc.Send[PasskeyAssertionCredentialExtensionOutput](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PasskeyAssertionCredentialExtensionOutputClass) New() PasskeyAssertionCredentialExtensionOutput {
	rv := objc.Send[PasskeyAssertionCredentialExtensionOutput](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PasskeyAssertionCredentialExtensionOutput) Init() PasskeyAssertionCredentialExtensionOutput {
	rv := objc.Send[PasskeyAssertionCredentialExtensionOutput](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PasskeyAssertionCredentialExtensionOutput) Autorelease() PasskeyAssertionCredentialExtensionOutput {
	rv := objc.Send[PasskeyAssertionCredentialExtensionOutput](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPasskeyAssertionCredentialExtensionOutput creates a new PasskeyAssertionCredentialExtensionOutput instance.
func NewPasskeyAssertionCredentialExtensionOutput() PasskeyAssertionCredentialExtensionOutput {
	return getPasskeyAssertionCredentialExtensionOutputClass().New()
}


// Output for operation during passkey assertion.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyAssertionCredentialExtensionOutput-c.class/largeBlobAssertionOutput
func (p_ PasskeyAssertionCredentialExtensionOutput) LargeBlobAssertionOutput() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("largeBlobAssertionOutput"))
	return rv
}



