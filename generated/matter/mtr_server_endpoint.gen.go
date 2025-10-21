// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRServerEndpoint] class.
var (
	MTRServerEndpointClass     _MTRServerEndpointClass
	MTRServerEndpointClassOnce sync.Once
)

func getMTRServerEndpointClass() _MTRServerEndpointClass {
	MTRServerEndpointClassOnce.Do(func() {
		MTRServerEndpointClass = _MTRServerEndpointClass{objc.GetClass("MTRServerEndpoint")}
	})
	return MTRServerEndpointClass
}

type _MTRServerEndpointClass struct {
	class objc.Class
}

// An interface definition for the [MTRServerEndpoint] class.
type IMTRServerEndpoint interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServerEndpoint
type MTRServerEndpoint struct {
	objectivec.Object
}

// MTRServerEndpointFrom constructs a [MTRServerEndpoint] from an unsafe.Pointer.
func MTRServerEndpointFrom(ptr unsafe.Pointer) MTRServerEndpoint {
	return MTRServerEndpoint{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRServerEndpointClass) Alloc() MTRServerEndpoint {
	rv := objc.Send[MTRServerEndpoint](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRServerEndpointClass) New() MTRServerEndpoint {
	rv := objc.Send[MTRServerEndpoint](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRServerEndpoint) Init() MTRServerEndpoint {
	rv := objc.Send[MTRServerEndpoint](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRServerEndpoint) Autorelease() MTRServerEndpoint {
	rv := objc.Send[MTRServerEndpoint](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRServerEndpoint creates a new MTRServerEndpoint instance.
func NewMTRServerEndpoint() MTRServerEndpoint {
	return getMTRServerEndpointClass().New()
}




