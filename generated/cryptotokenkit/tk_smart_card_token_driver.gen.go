// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TKSmartCardTokenDriver] class.
var (
	TKSmartCardTokenDriverClass     _TKSmartCardTokenDriverClass
	TKSmartCardTokenDriverClassOnce sync.Once
)

func getTKSmartCardTokenDriverClass() _TKSmartCardTokenDriverClass {
	TKSmartCardTokenDriverClassOnce.Do(func() {
		TKSmartCardTokenDriverClass = _TKSmartCardTokenDriverClass{objc.GetClass("TKSmartCardTokenDriver")}
	})
	return TKSmartCardTokenDriverClass
}

type _TKSmartCardTokenDriverClass struct {
	class objc.Class
}

// An interface definition for the [TKSmartCardTokenDriver] class.
type ITKSmartCardTokenDriver interface {
	ITKTokenDriver
	// properties:
	// methods:
}

// The driver that acts as an entry point for smart card app extensions.


// The driver that acts as an entry point for smart card app extensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardTokenDriver
type TKSmartCardTokenDriver struct {
	TKTokenDriver
}

// TKSmartCardTokenDriverFrom constructs a [TKSmartCardTokenDriver] from an unsafe.Pointer.
//
// The driver that acts as an entry point for smart card app extensions.
func TKSmartCardTokenDriverFrom(ptr unsafe.Pointer) TKSmartCardTokenDriver {
	return TKSmartCardTokenDriver{
		TKTokenDriver: TKTokenDriverFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc _TKSmartCardTokenDriverClass) Alloc() TKSmartCardTokenDriver {
	rv := objc.Send[TKSmartCardTokenDriver](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TKSmartCardTokenDriverClass) New() TKSmartCardTokenDriver {
	rv := objc.Send[TKSmartCardTokenDriver](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKSmartCardTokenDriver) Init() TKSmartCardTokenDriver {
	rv := objc.Send[TKSmartCardTokenDriver](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKSmartCardTokenDriver) Autorelease() TKSmartCardTokenDriver {
	rv := objc.Send[TKSmartCardTokenDriver](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKSmartCardTokenDriver creates a new TKSmartCardTokenDriver instance.
func NewTKSmartCardTokenDriver() TKSmartCardTokenDriver {
	return getTKSmartCardTokenDriverClass().New()
}




