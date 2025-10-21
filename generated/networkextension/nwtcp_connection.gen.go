// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [NWTCPConnection] class.
var (
	NWTCPConnectionClass     _NWTCPConnectionClass
	NWTCPConnectionClassOnce sync.Once
)

func getNWTCPConnectionClass() _NWTCPConnectionClass {
	NWTCPConnectionClassOnce.Do(func() {
		NWTCPConnectionClass = _NWTCPConnectionClass{objc.GetClass("NWTCPConnection")}
	})
	return NWTCPConnectionClass
}

type _NWTCPConnectionClass struct {
	class objc.Class
}

// An interface definition for the [NWTCPConnection] class.
type INWTCPConnection interface {
	objectivec.IObject
}

// An object to manage a TCP connection, with or without TLS.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWTCPConnection
type NWTCPConnection struct {
	objectivec.Object
}

// NWTCPConnectionFrom constructs a [NWTCPConnection] from an unsafe.Pointer.
//
// An object to manage a TCP connection, with or without TLS.
func NWTCPConnectionFrom(ptr unsafe.Pointer) NWTCPConnection {
	return NWTCPConnection{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NWTCPConnectionClass) Alloc() NWTCPConnection {
	rv := objc.Send[NWTCPConnection](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NWTCPConnectionClass) New() NWTCPConnection {
	rv := objc.Send[NWTCPConnection](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NWTCPConnection) Init() NWTCPConnection {
	rv := objc.Send[NWTCPConnection](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NWTCPConnection) Autorelease() NWTCPConnection {
	rv := objc.Send[NWTCPConnection](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNWTCPConnection creates a new NWTCPConnection instance.
func NewNWTCPConnection() NWTCPConnection {
	return getNWTCPConnectionClass().New()
}




