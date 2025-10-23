// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TKToken] class.
var (
	TKTokenClass     _TKTokenClass
	TKTokenClassOnce sync.Once
)

func getTKTokenClass() _TKTokenClass {
	TKTokenClassOnce.Do(func() {
		TKTokenClass = _TKTokenClass{objc.GetClass("TKToken")}
	})
	return TKTokenClass
}

type _TKTokenClass struct {
	class objc.Class
}

// An interface definition for the [TKToken] class.
type ITKToken interface {
	objectivec.IObject
	Configuration() unsafe.Pointer
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	KeychainContents() TKTokenKeychainContents
	SetKeychainContents(value ITKTokenKeychainContents)
	TokenDriver() TKTokenDriver
	SetTokenDriver(value ITKTokenDriver)
}

// A representation of a hardware-based cryptographic token.


// A representation of a hardware-based cryptographic token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKToken
type TKToken struct {
	objectivec.Object
}

// TKTokenFrom constructs a [TKToken] from an unsafe.Pointer.
//
// A representation of a hardware-based cryptographic token.
func TKTokenFrom(ptr unsafe.Pointer) TKToken {
	return TKToken{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TKTokenClass) Alloc() TKToken {
	rv := objc.Send[TKToken](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TKTokenClass) New() TKToken {
	rv := objc.Send[TKToken](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKToken) Init() TKToken {
	rv := objc.Send[TKToken](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKToken) Autorelease() TKToken {
	rv := objc.Send[TKToken](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKToken creates a new TKToken instance.
func NewTKToken() TKToken {
	return getTKTokenClass().New()
}



// Initializes a token with the driver you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKToken/init(tokenDriver:instanceID:)
func NewTKTokenWithTokenDriverInstanceID(tokenDriver ITKTokenDriver, instanceID ITKTokenInstanceID) TKToken {
	instance := getTKTokenClass().Alloc()
	rv := objc.Send[TKToken](instance.ID, objc.Sel("initWithTokenDriver:instanceID:"), tokenDriver, instanceID)
	rv.Autorelease()
	return rv
}



// The current configuration for a token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKToken/configuration-swift.property
func (t_ TKToken) Configuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("configuration"))
	return rv
}


// The token delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tktoken/delegate
func (t_ TKToken) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("delegate"))
	return rv
}


// The token delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tktoken/delegate
func (t_ TKToken) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDelegate:"), value)
}


// The contents of the keychain for this token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tktoken/keychaincontents
func (t_ TKToken) KeychainContents() TKTokenKeychainContents {
	rv := objc.Send[TKTokenKeychainContents](t_.ID, objc.Sel("keychainContents"))
	return rv
}


// The contents of the keychain for this token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tktoken/keychaincontents
func (t_ TKToken) SetKeychainContents(value ITKTokenKeychainContents) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setKeychainContents:"), value)
}


// The token driver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tktoken/tokendriver
func (t_ TKToken) TokenDriver() TKTokenDriver {
	rv := objc.Send[TKTokenDriver](t_.ID, objc.Sel("tokenDriver"))
	return rv
}


// The token driver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tktoken/tokendriver
func (t_ TKToken) SetTokenDriver(value ITKTokenDriver) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTokenDriver:"), value)
}


