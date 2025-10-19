// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NETransparentProxyProvider] class.
var (
	nETransparentProxyProviderClass     _NETransparentProxyProviderClass
	nETransparentProxyProviderClassOnce sync.Once
)

func getNETransparentProxyProviderClass() _NETransparentProxyProviderClass {
	nETransparentProxyProviderClassOnce.Do(func() {
		nETransparentProxyProviderClass = _NETransparentProxyProviderClass{objc.GetClass("NETransparentProxyProvider")}
	})
	return nETransparentProxyProviderClass
}

type _NETransparentProxyProviderClass struct {
	class objc.Class
}

// An interface definition for the [NETransparentProxyProvider] class.
type INETransparentProxyProvider interface {
	INEAppProxyProvider
}

// An object that implements the client side of a custom transparent network proxy solution.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETransparentProxyProvider
type NETransparentProxyProvider struct {
	NEAppProxyProvider
}

// NETransparentProxyProviderFrom constructs a [NETransparentProxyProvider] from an unsafe.Pointer.
//
// An object that implements the client side of a custom transparent network proxy solution.
func NETransparentProxyProviderFrom(ptr unsafe.Pointer) NETransparentProxyProvider {
	return NETransparentProxyProvider{
		NEAppProxyProvider: NEAppProxyProviderFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NETransparentProxyProviderClass) Alloc() NETransparentProxyProvider {
	rv := objc.Send[NETransparentProxyProvider](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NETransparentProxyProviderClass) New() NETransparentProxyProvider {
	rv := objc.Send[NETransparentProxyProvider](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NETransparentProxyProvider) Init() NETransparentProxyProvider {
	rv := objc.Send[NETransparentProxyProvider](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NETransparentProxyProvider) Autorelease() NETransparentProxyProvider {
	rv := objc.Send[NETransparentProxyProvider](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNETransparentProxyProvider creates a new NETransparentProxyProvider instance.
func NewNETransparentProxyProvider() NETransparentProxyProvider {
	return getNETransparentProxyProviderClass().New()
}




