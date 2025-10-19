// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NERelay] class.
var (
	nERelayClass     _NERelayClass
	nERelayClassOnce sync.Once
)

func getNERelayClass() _NERelayClass {
	nERelayClassOnce.Do(func() {
		nERelayClass = _NERelayClass{objc.GetClass("NERelay")}
	})
	return nERelayClass
}

type _NERelayClass struct {
	class objc.Class
}

// An interface definition for the [NERelay] class.
type INERelay interface {
	objectivec.IObject
}

// A single relay server configuration that you can chain together with other relays. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelay
type NERelay struct {
	objectivec.Object
}

// NERelayFrom constructs a [NERelay] from an unsafe.Pointer.
//
// A single relay server configuration that you can chain together with other relays.
func NERelayFrom(ptr unsafe.Pointer) NERelay {
	return NERelay{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NERelayClass) Alloc() NERelay {
	rv := objc.Send[NERelay](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NERelayClass) New() NERelay {
	rv := objc.Send[NERelay](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NERelay) Init() NERelay {
	rv := objc.Send[NERelay](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NERelay) Autorelease() NERelay {
	rv := objc.Send[NERelay](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNERelay creates a new NERelay instance.
func NewNERelay() NERelay {
	return getNERelayClass().New()
}




