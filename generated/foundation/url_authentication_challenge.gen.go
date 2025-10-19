// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLAuthenticationChallenge] class.
var (
	uRLAuthenticationChallengeClass     _URLAuthenticationChallengeClass
	uRLAuthenticationChallengeClassOnce sync.Once
)

func getURLAuthenticationChallengeClass() _URLAuthenticationChallengeClass {
	uRLAuthenticationChallengeClassOnce.Do(func() {
		uRLAuthenticationChallengeClass = _URLAuthenticationChallengeClass{objc.GetClass("NSURLAuthenticationChallenge")}
	})
	return uRLAuthenticationChallengeClass
}

type _URLAuthenticationChallengeClass struct {
	class objc.Class
}

// An interface definition for the [URLAuthenticationChallenge] class.
type IURLAuthenticationChallenge interface {
	objectivec.IObject
}

// A challenge from a server requiring authentication from the client.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLAuthenticationChallenge
type URLAuthenticationChallenge struct {
	objectivec.Object
}

// URLAuthenticationChallengeFrom constructs a [URLAuthenticationChallenge] from an unsafe.Pointer.
//
// A challenge from a server requiring authentication from the client.
func URLAuthenticationChallengeFrom(ptr unsafe.Pointer) URLAuthenticationChallenge {
	return URLAuthenticationChallenge{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _URLAuthenticationChallengeClass) Alloc() URLAuthenticationChallenge {
	rv := objc.Send[URLAuthenticationChallenge](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _URLAuthenticationChallengeClass) New() URLAuthenticationChallenge {
	rv := objc.Send[URLAuthenticationChallenge](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLAuthenticationChallenge) Init() URLAuthenticationChallenge {
	rv := objc.Send[URLAuthenticationChallenge](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLAuthenticationChallenge) Autorelease() URLAuthenticationChallenge {
	rv := objc.Send[URLAuthenticationChallenge](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLAuthenticationChallenge creates a new URLAuthenticationChallenge instance.
func NewURLAuthenticationChallenge() URLAuthenticationChallenge {
	return getURLAuthenticationChallengeClass().New()
}




