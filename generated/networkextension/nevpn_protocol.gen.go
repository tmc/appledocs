// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NEVPNProtocol] class.
var (
	nEVPNProtocolClass     _NEVPNProtocolClass
	nEVPNProtocolClassOnce sync.Once
)

func getNEVPNProtocolClass() _NEVPNProtocolClass {
	nEVPNProtocolClassOnce.Do(func() {
		nEVPNProtocolClass = _NEVPNProtocolClass{objc.GetClass("NEVPNProtocol")}
	})
	return nEVPNProtocolClass
}

type _NEVPNProtocolClass struct {
	class objc.Class
}

// An interface definition for the [NEVPNProtocol] class.
type INEVPNProtocol interface {
	objectivec.IObject
}

// A parent class referenced by other NetworkExtension classes. [Full Topic]
type NEVPNProtocol struct {
	objectivec.Object
}

// NEVPNProtocolFrom constructs a [NEVPNProtocol] from an unsafe.Pointer.
//
// A parent class referenced by other NetworkExtension classes.
func NEVPNProtocolFrom(ptr unsafe.Pointer) NEVPNProtocol {
	return NEVPNProtocol{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NEVPNProtocolClass) Alloc() NEVPNProtocol {
	rv := objc.Send[NEVPNProtocol](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEVPNProtocolClass) New() NEVPNProtocol {
	rv := objc.Send[NEVPNProtocol](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEVPNProtocol) Init() NEVPNProtocol {
	rv := objc.Send[NEVPNProtocol](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEVPNProtocol) Autorelease() NEVPNProtocol {
	rv := objc.Send[NEVPNProtocol](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEVPNProtocol creates a new NEVPNProtocol instance.
func NewNEVPNProtocol() NEVPNProtocol {
	return getNEVPNProtocolClass().New()
}




