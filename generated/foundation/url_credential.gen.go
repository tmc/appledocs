// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLCredential] class.
var (
	uRLCredentialClass     _URLCredentialClass
	uRLCredentialClassOnce sync.Once
)

func getURLCredentialClass() _URLCredentialClass {
	uRLCredentialClassOnce.Do(func() {
		uRLCredentialClass = _URLCredentialClass{objc.GetClass("NSURLCredential")}
	})
	return uRLCredentialClass
}

type _URLCredentialClass struct {
	class objc.Class
}

// An interface definition for the [URLCredential] class.
type IURLCredential interface {
	objectivec.IObject
}

// n authentication credential consisting of information specific to the type of credential and the type of persistent storage to use, if any.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredential
type URLCredential struct {
	objectivec.Object
}

// URLCredentialFrom constructs a [URLCredential] from an unsafe.Pointer.
//
// n authentication credential consisting of information specific to the type of credential and the type of persistent storage to use, if any.
func URLCredentialFrom(ptr unsafe.Pointer) URLCredential {
	return URLCredential{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _URLCredentialClass) Alloc() URLCredential {
	rv := objc.Send[URLCredential](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _URLCredentialClass) New() URLCredential {
	rv := objc.Send[URLCredential](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLCredential) Init() URLCredential {
	rv := objc.Send[URLCredential](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLCredential) Autorelease() URLCredential {
	rv := objc.Send[URLCredential](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLCredential creates a new URLCredential instance.
func NewURLCredential() URLCredential {
	return getURLCredentialClass().New()
}




