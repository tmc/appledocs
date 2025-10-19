// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NERelayManager] class.
var (
	nERelayManagerClass     _NERelayManagerClass
	nERelayManagerClassOnce sync.Once
)

func getNERelayManagerClass() _NERelayManagerClass {
	nERelayManagerClassOnce.Do(func() {
		nERelayManagerClass = _NERelayManagerClass{objc.GetClass("NERelayManager")}
	})
	return nERelayManagerClass
}

type _NERelayManagerClass struct {
	class objc.Class
}

// An interface definition for the [NERelayManager] class.
type INERelayManager interface {
	objectivec.IObject
}

// An object you use to create and manage a network relay configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManager
type NERelayManager struct {
	objectivec.Object
}

// NERelayManagerFrom constructs a [NERelayManager] from an unsafe.Pointer.
//
// An object you use to create and manage a network relay configuration.
func NERelayManagerFrom(ptr unsafe.Pointer) NERelayManager {
	return NERelayManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NERelayManagerClass) Alloc() NERelayManager {
	rv := objc.Send[NERelayManager](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NERelayManagerClass) New() NERelayManager {
	rv := objc.Send[NERelayManager](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NERelayManager) Init() NERelayManager {
	rv := objc.Send[NERelayManager](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NERelayManager) Autorelease() NERelayManager {
	rv := objc.Send[NERelayManager](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNERelayManager creates a new NERelayManager instance.
func NewNERelayManager() NERelayManager {
	return getNERelayManagerClass().New()
}




