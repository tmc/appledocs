// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [trustUsesMutualAuthentication] class.
var (
	TrustUsesMutualAuthenticationClass     _trustUsesMutualAuthenticationClass
	TrustUsesMutualAuthenticationClassOnce sync.Once
)

func gettrustUsesMutualAuthenticationClass() _trustUsesMutualAuthenticationClass {
	TrustUsesMutualAuthenticationClassOnce.Do(func() {
		TrustUsesMutualAuthenticationClass = _trustUsesMutualAuthenticationClass{objc.GetClass("trustUsesMutualAuthentication")}
	})
	return TrustUsesMutualAuthenticationClass
}

type _trustUsesMutualAuthenticationClass struct {
	class objc.Class
}

// An interface definition for the [trustUsesMutualAuthentication] class.
type ItrustUsesMutualAuthentication interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/trustUsesMutualAuthentication-c.ivar
type trustUsesMutualAuthentication struct {
	objectivec.Object
}

// trustUsesMutualAuthenticationFrom constructs a [trustUsesMutualAuthentication] from an unsafe.Pointer.
func trustUsesMutualAuthenticationFrom(ptr unsafe.Pointer) trustUsesMutualAuthentication {
	return trustUsesMutualAuthentication{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _trustUsesMutualAuthenticationClass) Alloc() trustUsesMutualAuthentication {
	rv := objc.Send[trustUsesMutualAuthentication](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _trustUsesMutualAuthenticationClass) New() trustUsesMutualAuthentication {
	rv := objc.Send[trustUsesMutualAuthentication](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ trustUsesMutualAuthentication) Init() trustUsesMutualAuthentication {
	rv := objc.Send[trustUsesMutualAuthentication](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ trustUsesMutualAuthentication) Autorelease() trustUsesMutualAuthentication {
	rv := objc.Send[trustUsesMutualAuthentication](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewtrustUsesMutualAuthentication creates a new trustUsesMutualAuthentication instance.
func NewtrustUsesMutualAuthentication() trustUsesMutualAuthentication {
	return gettrustUsesMutualAuthenticationClass().New()
}




