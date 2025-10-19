// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NEVPNConnection] class.
var (
	nEVPNConnectionClass     _NEVPNConnectionClass
	nEVPNConnectionClassOnce sync.Once
)

func getNEVPNConnectionClass() _NEVPNConnectionClass {
	nEVPNConnectionClassOnce.Do(func() {
		nEVPNConnectionClass = _NEVPNConnectionClass{objc.GetClass("NEVPNConnection")}
	})
	return nEVPNConnectionClass
}

type _NEVPNConnectionClass struct {
	class objc.Class
}

// An interface definition for the [NEVPNConnection] class.
type INEVPNConnection interface {
	objectivec.IObject
}

// An object to start and stop a Personal VPN connection and get its status. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnection
type NEVPNConnection struct {
	objectivec.Object
}

// NEVPNConnectionFrom constructs a [NEVPNConnection] from an unsafe.Pointer.
//
// An object to start and stop a Personal VPN connection and get its status.
func NEVPNConnectionFrom(ptr unsafe.Pointer) NEVPNConnection {
	return NEVPNConnection{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NEVPNConnectionClass) Alloc() NEVPNConnection {
	rv := objc.Send[NEVPNConnection](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEVPNConnectionClass) New() NEVPNConnection {
	rv := objc.Send[NEVPNConnection](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEVPNConnection) Init() NEVPNConnection {
	rv := objc.Send[NEVPNConnection](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEVPNConnection) Autorelease() NEVPNConnection {
	rv := objc.Send[NEVPNConnection](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEVPNConnection creates a new NEVPNConnection instance.
func NewNEVPNConnection() NEVPNConnection {
	return getNEVPNConnectionClass().New()
}




