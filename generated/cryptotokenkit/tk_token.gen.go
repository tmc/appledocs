// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
}

// A representation of a hardware-based cryptographic token.
//
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
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKToken/init(tokenDriver:instanceID:)
func NewTKTokenWithTokenDriverInstanceID(tokenDriver unsafe.Pointer, instanceID unsafe.Pointer) TKToken {
	instance := getTKTokenClass().Alloc()
	rv := objc.Send[TKToken](instance.ID, objc.Sel("initWithTokenDriver:instanceID:"), tokenDriver, instanceID)
	rv.Autorelease()
	return rv
}


// The current configuration for a token.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKToken/configuration-swift.property
func (t_ TKToken) Configuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("configuration"))
	return rv
}


