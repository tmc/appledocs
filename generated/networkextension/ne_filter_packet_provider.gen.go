// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NEFilterPacketProvider] class.
var (
	NEFilterPacketProviderClass     _NEFilterPacketProviderClass
	NEFilterPacketProviderClassOnce sync.Once
)

func getNEFilterPacketProviderClass() _NEFilterPacketProviderClass {
	NEFilterPacketProviderClassOnce.Do(func() {
		NEFilterPacketProviderClass = _NEFilterPacketProviderClass{objc.GetClass("NEFilterPacketProvider")}
	})
	return NEFilterPacketProviderClass
}

type _NEFilterPacketProviderClass struct {
	class objc.Class
}

// An interface definition for the [NEFilterPacketProvider] class.
type INEFilterPacketProvider interface {
	INEFilterProvider
}

// A filter provider that evaluates network packets and decides whether to block, allow, or delay the packets.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterPacketProvider
type NEFilterPacketProvider struct {
	NEFilterProvider
}

// NEFilterPacketProviderFrom constructs a [NEFilterPacketProvider] from an unsafe.Pointer.
//
// A filter provider that evaluates network packets and decides whether to block, allow, or delay the packets.
func NEFilterPacketProviderFrom(ptr unsafe.Pointer) NEFilterPacketProvider {
	return NEFilterPacketProvider{
		NEFilterProvider: NEFilterProviderFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NEFilterPacketProviderClass) Alloc() NEFilterPacketProvider {
	rv := objc.Send[NEFilterPacketProvider](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEFilterPacketProviderClass) New() NEFilterPacketProvider {
	rv := objc.Send[NEFilterPacketProvider](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEFilterPacketProvider) Init() NEFilterPacketProvider {
	rv := objc.Send[NEFilterPacketProvider](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEFilterPacketProvider) Autorelease() NEFilterPacketProvider {
	rv := objc.Send[NEFilterPacketProvider](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterPacketProvider creates a new NEFilterPacketProvider instance.
func NewNEFilterPacketProvider() NEFilterPacketProvider {
	return getNEFilterPacketProviderClass().New()
}




