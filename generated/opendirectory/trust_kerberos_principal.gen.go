// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [trustKerberosPrincipal] class.
var (
	TrustKerberosPrincipalClass     _trustKerberosPrincipalClass
	TrustKerberosPrincipalClassOnce sync.Once
)

func gettrustKerberosPrincipalClass() _trustKerberosPrincipalClass {
	TrustKerberosPrincipalClassOnce.Do(func() {
		TrustKerberosPrincipalClass = _trustKerberosPrincipalClass{objc.GetClass("trustKerberosPrincipal")}
	})
	return TrustKerberosPrincipalClass
}

type _trustKerberosPrincipalClass struct {
	class objc.Class
}

// An interface definition for the [trustKerberosPrincipal] class.
type ItrustKerberosPrincipal interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/trustKerberosPrincipal-c.ivar
type trustKerberosPrincipal struct {
	objectivec.Object
}

// trustKerberosPrincipalFrom constructs a [trustKerberosPrincipal] from an unsafe.Pointer.
func trustKerberosPrincipalFrom(ptr unsafe.Pointer) trustKerberosPrincipal {
	return trustKerberosPrincipal{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _trustKerberosPrincipalClass) Alloc() trustKerberosPrincipal {
	rv := objc.Send[trustKerberosPrincipal](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _trustKerberosPrincipalClass) New() trustKerberosPrincipal {
	rv := objc.Send[trustKerberosPrincipal](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ trustKerberosPrincipal) Init() trustKerberosPrincipal {
	rv := objc.Send[trustKerberosPrincipal](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ trustKerberosPrincipal) Autorelease() trustKerberosPrincipal {
	rv := objc.Send[trustKerberosPrincipal](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewtrustKerberosPrincipal creates a new trustKerberosPrincipal instance.
func NewtrustKerberosPrincipal() trustKerberosPrincipal {
	return gettrustKerberosPrincipalClass().New()
}




