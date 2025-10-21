// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TKTokenSession] class.
var (
	TKTokenSessionClass     _TKTokenSessionClass
	TKTokenSessionClassOnce sync.Once
)

func getTKTokenSessionClass() _TKTokenSessionClass {
	TKTokenSessionClassOnce.Do(func() {
		TKTokenSessionClass = _TKTokenSessionClass{objc.GetClass("TKTokenSession")}
	})
	return TKTokenSessionClass
}

type _TKTokenSessionClass struct {
	class objc.Class
}

// An interface definition for the [TKTokenSession] class.
type ITKTokenSession interface {
	objectivec.IObject
}

// A token session that manages the authentication state of a token.
//
// A token session communicates with its delegate to perform operations with its token that are bound to the authentication state. A session is always instantiated by a instance through the token’s delegate when the framework detects access to the token from a new authentication session.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenSession
type TKTokenSession struct {
	objectivec.Object
}

// TKTokenSessionFrom constructs a [TKTokenSession] from an unsafe.Pointer.
//
// A token session that manages the authentication state of a token.
func TKTokenSessionFrom(ptr unsafe.Pointer) TKTokenSession {
	return TKTokenSession{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TKTokenSessionClass) Alloc() TKTokenSession {
	rv := objc.Send[TKTokenSession](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TKTokenSessionClass) New() TKTokenSession {
	rv := objc.Send[TKTokenSession](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKTokenSession) Init() TKTokenSession {
	rv := objc.Send[TKTokenSession](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKTokenSession) Autorelease() TKTokenSession {
	rv := objc.Send[TKTokenSession](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKTokenSession creates a new TKTokenSession instance.
func NewTKTokenSession() TKTokenSession {
	return getTKTokenSessionClass().New()
}


// Initializes a token session with the specified token.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenSession/init(token:)
func NewTKTokenSessionWithToken(token unsafe.Pointer) TKTokenSession {
	instance := getTKTokenSessionClass().Alloc()
	rv := objc.Send[TKTokenSession](instance.ID, objc.Sel("initWithToken:"), token)
	rv.Autorelease()
	return rv
}


// The token session delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenSession/delegate
func (t_ TKTokenSession) Delegate() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The token session delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenSession/delegate
func (t_ TKTokenSession) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDelegate:"), value)
}
// The token to which the session is bound.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenSession/token
func (t_ TKTokenSession) Token() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("token"))
	return rv
}


