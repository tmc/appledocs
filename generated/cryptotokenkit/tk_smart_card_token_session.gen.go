// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TKSmartCardTokenSession] class.
var (
	TKSmartCardTokenSessionClass     _TKSmartCardTokenSessionClass
	TKSmartCardTokenSessionClassOnce sync.Once
)

func getTKSmartCardTokenSessionClass() _TKSmartCardTokenSessionClass {
	TKSmartCardTokenSessionClassOnce.Do(func() {
		TKSmartCardTokenSessionClass = _TKSmartCardTokenSessionClass{objc.GetClass("TKSmartCardTokenSession")}
	})
	return TKSmartCardTokenSessionClass
}

type _TKSmartCardTokenSessionClass struct {
	class objc.Class
}

// An interface definition for the [TKSmartCardTokenSession] class.
type ITKSmartCardTokenSession interface {
	ITKTokenSession
	GetSmartCardWithError(error_ unsafe.Pointer) TKSmartCard
}

// A token session that is based on a smart card token.
//
// You can use the property to access and send APDUs to the underlying smart card.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardTokenSession
type TKSmartCardTokenSession struct {
	TKTokenSession
}

// TKSmartCardTokenSessionFrom constructs a [TKSmartCardTokenSession] from an unsafe.Pointer.
//
// A token session that is based on a smart card token.
func TKSmartCardTokenSessionFrom(ptr unsafe.Pointer) TKSmartCardTokenSession {
	return TKSmartCardTokenSession{
		TKTokenSession: TKTokenSessionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc _TKSmartCardTokenSessionClass) Alloc() TKSmartCardTokenSession {
	rv := objc.Send[TKSmartCardTokenSession](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TKSmartCardTokenSessionClass) New() TKSmartCardTokenSession {
	rv := objc.Send[TKSmartCardTokenSession](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKSmartCardTokenSession) Init() TKSmartCardTokenSession {
	rv := objc.Send[TKSmartCardTokenSession](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKSmartCardTokenSession) Autorelease() TKSmartCardTokenSession {
	rv := objc.Send[TKSmartCardTokenSession](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKSmartCardTokenSession creates a new TKSmartCardTokenSession instance.
func NewTKSmartCardTokenSession() TKSmartCardTokenSession {
	return getTKSmartCardTokenSessionClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardTokenSession/getSmartCard()
func (t_ TKSmartCardTokenSession) GetSmartCardWithError(error_ unsafe.Pointer) TKSmartCard {
	rv := objc.Send[TKSmartCard](t_.ID, objc.Sel("getSmartCardWithError:"), error_)
	return rv
}

// The smart card for the active exclusive session and selected application.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardTokenSession/smartCard
func (t_ TKSmartCardTokenSession) SmartCard() TKSmartCard {
	rv := objc.Send[TKSmartCard](t_.ID, objc.Sel("smartCard"))
	return rv
}



