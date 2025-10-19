// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NEProvider] class.
var (
	nEProviderClass     _NEProviderClass
	nEProviderClassOnce sync.Once
)

func getNEProviderClass() _NEProviderClass {
	nEProviderClassOnce.Do(func() {
		nEProviderClass = _NEProviderClass{objc.GetClass("NEProvider")}
	})
	return nEProviderClass
}

type _NEProviderClass struct {
	class objc.Class
}

// An interface definition for the [NEProvider] class.
type INEProvider interface {
	objectivec.IObject
}

// An abstract base class for all NetworkExtension providers.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProvider
type NEProvider struct {
	objectivec.Object
}

// NEProviderFrom constructs a [NEProvider] from an unsafe.Pointer.
//
// An abstract base class for all NetworkExtension providers.
func NEProviderFrom(ptr unsafe.Pointer) NEProvider {
	return NEProvider{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NEProviderClass) Alloc() NEProvider {
	rv := objc.Send[NEProvider](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEProviderClass) New() NEProvider {
	rv := objc.Send[NEProvider](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEProvider) Init() NEProvider {
	rv := objc.Send[NEProvider](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEProvider) Autorelease() NEProvider {
	rv := objc.Send[NEProvider](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEProvider creates a new NEProvider instance.
func NewNEProvider() NEProvider {
	return getNEProviderClass().New()
}




