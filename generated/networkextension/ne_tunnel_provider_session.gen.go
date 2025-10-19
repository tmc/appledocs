// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NETunnelProviderSession] class.
var (
	nETunnelProviderSessionClass     _NETunnelProviderSessionClass
	nETunnelProviderSessionClassOnce sync.Once
)

func getNETunnelProviderSessionClass() _NETunnelProviderSessionClass {
	nETunnelProviderSessionClassOnce.Do(func() {
		nETunnelProviderSessionClass = _NETunnelProviderSessionClass{objc.GetClass("NETunnelProviderSession")}
	})
	return nETunnelProviderSessionClass
}

type _NETunnelProviderSessionClass struct {
	class objc.Class
}

// An interface definition for the [NETunnelProviderSession] class.
type INETunnelProviderSession interface {
	INEVPNConnection
}

// An object to start and stop a tunnel connection and get its status. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderSession
type NETunnelProviderSession struct {
	NEVPNConnection
}

// NETunnelProviderSessionFrom constructs a [NETunnelProviderSession] from an unsafe.Pointer.
//
// An object to start and stop a tunnel connection and get its status.
func NETunnelProviderSessionFrom(ptr unsafe.Pointer) NETunnelProviderSession {
	return NETunnelProviderSession{
		NEVPNConnection: NEVPNConnectionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NETunnelProviderSessionClass) Alloc() NETunnelProviderSession {
	rv := objc.Send[NETunnelProviderSession](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NETunnelProviderSessionClass) New() NETunnelProviderSession {
	rv := objc.Send[NETunnelProviderSession](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NETunnelProviderSession) Init() NETunnelProviderSession {
	rv := objc.Send[NETunnelProviderSession](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NETunnelProviderSession) Autorelease() NETunnelProviderSession {
	rv := objc.Send[NETunnelProviderSession](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNETunnelProviderSession creates a new NETunnelProviderSession instance.
func NewNETunnelProviderSession() NETunnelProviderSession {
	return getNETunnelProviderSessionClass().New()
}




