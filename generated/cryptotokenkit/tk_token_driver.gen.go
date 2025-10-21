// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TKTokenDriver] class.
var (
	TKTokenDriverClass     _TKTokenDriverClass
	TKTokenDriverClassOnce sync.Once
)

func getTKTokenDriverClass() _TKTokenDriverClass {
	TKTokenDriverClassOnce.Do(func() {
		TKTokenDriverClass = _TKTokenDriverClass{objc.GetClass("TKTokenDriver")}
	})
	return TKTokenDriverClass
}

type _TKTokenDriverClass struct {
	class objc.Class
}

// An interface definition for the [TKTokenDriver] class.
type ITKTokenDriver interface {
	objectivec.IObject
}

// A base class for building token drivers.
//
// When using the class, implement the protocol with the method, which the system invokes when it requests the creation of a token instance. After you create the token driver, it can examine and to implement your desired functionality. An implementation can also access its associated token configuration using the property.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenDriver
type TKTokenDriver struct {
	objectivec.Object
}

// TKTokenDriverFrom constructs a [TKTokenDriver] from an unsafe.Pointer.
//
// A base class for building token drivers.
func TKTokenDriverFrom(ptr unsafe.Pointer) TKTokenDriver {
	return TKTokenDriver{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TKTokenDriverClass) Alloc() TKTokenDriver {
	rv := objc.Send[TKTokenDriver](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TKTokenDriverClass) New() TKTokenDriver {
	rv := objc.Send[TKTokenDriver](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKTokenDriver) Init() TKTokenDriver {
	rv := objc.Send[TKTokenDriver](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKTokenDriver) Autorelease() TKTokenDriver {
	rv := objc.Send[TKTokenDriver](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKTokenDriver creates a new TKTokenDriver instance.
func NewTKTokenDriver() TKTokenDriver {
	return getTKTokenDriverClass().New()
}


// The token driver delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenDriver/delegate
func (t_ TKTokenDriver) Delegate() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The token driver delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenDriver/delegate
func (t_ TKTokenDriver) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDelegate:"), value)
}



