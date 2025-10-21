// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TKSmartCardToken] class.
var (
	TKSmartCardTokenClass     _TKSmartCardTokenClass
	TKSmartCardTokenClassOnce sync.Once
)

func getTKSmartCardTokenClass() _TKSmartCardTokenClass {
	TKSmartCardTokenClassOnce.Do(func() {
		TKSmartCardTokenClass = _TKSmartCardTokenClass{objc.GetClass("TKSmartCardToken")}
	})
	return TKSmartCardTokenClass
}

type _TKSmartCardTokenClass struct {
	class objc.Class
}

// An interface definition for the [TKSmartCardToken] class.
type ITKSmartCardToken interface {
	ITKToken
}

// A representation of a smart card based cryptographic token.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardToken
type TKSmartCardToken struct {
	TKToken
}

// TKSmartCardTokenFrom constructs a [TKSmartCardToken] from an unsafe.Pointer.
//
// A representation of a smart card based cryptographic token.
func TKSmartCardTokenFrom(ptr unsafe.Pointer) TKSmartCardToken {
	return TKSmartCardToken{
		TKToken: TKTokenFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc _TKSmartCardTokenClass) Alloc() TKSmartCardToken {
	rv := objc.Send[TKSmartCardToken](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TKSmartCardTokenClass) New() TKSmartCardToken {
	rv := objc.Send[TKSmartCardToken](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKSmartCardToken) Init() TKSmartCardToken {
	rv := objc.Send[TKSmartCardToken](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKSmartCardToken) Autorelease() TKSmartCardToken {
	rv := objc.Send[TKSmartCardToken](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKSmartCardToken creates a new TKSmartCardToken instance.
func NewTKSmartCardToken() TKSmartCardToken {
	return getTKSmartCardTokenClass().New()
}


// Initializes a smart card token with the specified smart card, application identifier, and token driver.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardToken/init(smartCard:aid:instanceID:tokenDriver:)
func NewTKSmartCardTokenWithSmartCardAIDInstanceIDTokenDriver(smartCard unsafe.Pointer, AID unsafe.Pointer, instanceID string, tokenDriver unsafe.Pointer) TKSmartCardToken {
	instance := getTKSmartCardTokenClass().Alloc()
	rv := objc.Send[TKSmartCardToken](instance.ID, objc.Sel("initWithSmartCard:AID:instanceID:tokenDriver:"), smartCard, AID, objc.String(instanceID), tokenDriver)
	rv.Autorelease()
	return rv
}


// The ISO 7816-4 application identifiers of the Smart Card.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardToken/aid
func (t_ TKSmartCardToken) AID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("AID"))
	return rv
}


