// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NEAppProxyProvider] class.
var (
	nEAppProxyProviderClass     _NEAppProxyProviderClass
	nEAppProxyProviderClassOnce sync.Once
)

func getNEAppProxyProviderClass() _NEAppProxyProviderClass {
	nEAppProxyProviderClassOnce.Do(func() {
		nEAppProxyProviderClass = _NEAppProxyProviderClass{objc.GetClass("NEAppProxyProvider")}
	})
	return nEAppProxyProviderClass
}

type _NEAppProxyProviderClass struct {
	class objc.Class
}

// An interface definition for the [NEAppProxyProvider] class.
type INEAppProxyProvider interface {
	objectivec.IObject
}

// A parent class referenced by other NetworkExtension classes. [Full Topic]
type NEAppProxyProvider struct {
	objectivec.Object
}

// NEAppProxyProviderFrom constructs a [NEAppProxyProvider] from an unsafe.Pointer.
//
// A parent class referenced by other NetworkExtension classes.
func NEAppProxyProviderFrom(ptr unsafe.Pointer) NEAppProxyProvider {
	return NEAppProxyProvider{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NEAppProxyProviderClass) Alloc() NEAppProxyProvider {
	rv := objc.Send[NEAppProxyProvider](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEAppProxyProviderClass) New() NEAppProxyProvider {
	rv := objc.Send[NEAppProxyProvider](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEAppProxyProvider) Init() NEAppProxyProvider {
	rv := objc.Send[NEAppProxyProvider](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEAppProxyProvider) Autorelease() NEAppProxyProvider {
	rv := objc.Send[NEAppProxyProvider](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEAppProxyProvider creates a new NEAppProxyProvider instance.
func NewNEAppProxyProvider() NEAppProxyProvider {
	return getNEAppProxyProviderClass().New()
}




