// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NWEndpoint] class.
var (
	nWEndpointClass     _NWEndpointClass
	nWEndpointClassOnce sync.Once
)

func getNWEndpointClass() _NWEndpointClass {
	nWEndpointClassOnce.Do(func() {
		nWEndpointClass = _NWEndpointClass{objc.GetClass("NWEndpoint")}
	})
	return nWEndpointClass
}

type _NWEndpointClass struct {
	class objc.Class
}

// An interface definition for the [NWEndpoint] class.
type INWEndpoint interface {
	objectivec.IObject
}

// An abstract base class, shared by or , that represents the source or destination of a network connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWEndpoint
type NWEndpoint struct {
	objectivec.Object
}

// NWEndpointFrom constructs a [NWEndpoint] from an unsafe.Pointer.
//
// An abstract base class, shared by or , that represents the source or destination of a network connection.
func NWEndpointFrom(ptr unsafe.Pointer) NWEndpoint {
	return NWEndpoint{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NWEndpointClass) Alloc() NWEndpoint {
	rv := objc.Send[NWEndpoint](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NWEndpointClass) New() NWEndpoint {
	rv := objc.Send[NWEndpoint](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NWEndpoint) Init() NWEndpoint {
	rv := objc.Send[NWEndpoint](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NWEndpoint) Autorelease() NWEndpoint {
	rv := objc.Send[NWEndpoint](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNWEndpoint creates a new NWEndpoint instance.
func NewNWEndpoint() NWEndpoint {
	return getNWEndpointClass().New()
}



