// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [trustUsesSystemKeychain] class.
var (
	TrustUsesSystemKeychainClass     _trustUsesSystemKeychainClass
	TrustUsesSystemKeychainClassOnce sync.Once
)

func gettrustUsesSystemKeychainClass() _trustUsesSystemKeychainClass {
	TrustUsesSystemKeychainClassOnce.Do(func() {
		TrustUsesSystemKeychainClass = _trustUsesSystemKeychainClass{objc.GetClass("trustUsesSystemKeychain")}
	})
	return TrustUsesSystemKeychainClass
}

type _trustUsesSystemKeychainClass struct {
	class objc.Class
}

// An interface definition for the [trustUsesSystemKeychain] class.
type ItrustUsesSystemKeychain interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/trustUsesSystemKeychain-c.ivar
type trustUsesSystemKeychain struct {
	objectivec.Object
}

// trustUsesSystemKeychainFrom constructs a [trustUsesSystemKeychain] from an unsafe.Pointer.
func trustUsesSystemKeychainFrom(ptr unsafe.Pointer) trustUsesSystemKeychain {
	return trustUsesSystemKeychain{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _trustUsesSystemKeychainClass) Alloc() trustUsesSystemKeychain {
	rv := objc.Send[trustUsesSystemKeychain](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _trustUsesSystemKeychainClass) New() trustUsesSystemKeychain {
	rv := objc.Send[trustUsesSystemKeychain](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ trustUsesSystemKeychain) Init() trustUsesSystemKeychain {
	rv := objc.Send[trustUsesSystemKeychain](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ trustUsesSystemKeychain) Autorelease() trustUsesSystemKeychain {
	rv := objc.Send[trustUsesSystemKeychain](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewtrustUsesSystemKeychain creates a new trustUsesSystemKeychain instance.
func NewtrustUsesSystemKeychain() trustUsesSystemKeychain {
	return gettrustUsesSystemKeychainClass().New()
}




