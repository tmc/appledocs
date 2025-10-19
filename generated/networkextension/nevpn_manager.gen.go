// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NEVPNManager] class.
var (
	nEVPNManagerClass     _NEVPNManagerClass
	nEVPNManagerClassOnce sync.Once
)

func getNEVPNManagerClass() _NEVPNManagerClass {
	nEVPNManagerClassOnce.Do(func() {
		nEVPNManagerClass = _NEVPNManagerClass{objc.GetClass("NEVPNManager")}
	})
	return nEVPNManagerClass
}

type _NEVPNManagerClass struct {
	class objc.Class
}

// An interface definition for the [NEVPNManager] class.
type INEVPNManager interface {
	objectivec.IObject
}

// An object to create and manage a Personal VPN configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNManager
type NEVPNManager struct {
	objectivec.Object
}

// NEVPNManagerFrom constructs a [NEVPNManager] from an unsafe.Pointer.
//
// An object to create and manage a Personal VPN configuration.
func NEVPNManagerFrom(ptr unsafe.Pointer) NEVPNManager {
	return NEVPNManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NEVPNManagerClass) Alloc() NEVPNManager {
	rv := objc.Send[NEVPNManager](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEVPNManagerClass) New() NEVPNManager {
	rv := objc.Send[NEVPNManager](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEVPNManager) Init() NEVPNManager {
	rv := objc.Send[NEVPNManager](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEVPNManager) Autorelease() NEVPNManager {
	rv := objc.Send[NEVPNManager](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEVPNManager creates a new NEVPNManager instance.
func NewNEVPNManager() NEVPNManager {
	return getNEVPNManagerClass().New()
}


// Access the single instance of .
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNManager/shared()
func (nc _NEVPNManagerClass) SharedManager() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(nc.class), objc.Sel("sharedManager"))
	return rv
}


