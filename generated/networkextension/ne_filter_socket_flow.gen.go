// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NEFilterSocketFlow] class.
var (
	NEFilterSocketFlowClass     _NEFilterSocketFlowClass
	NEFilterSocketFlowClassOnce sync.Once
)

func getNEFilterSocketFlowClass() _NEFilterSocketFlowClass {
	NEFilterSocketFlowClassOnce.Do(func() {
		NEFilterSocketFlowClass = _NEFilterSocketFlowClass{objc.GetClass("NEFilterSocketFlow")}
	})
	return NEFilterSocketFlowClass
}

type _NEFilterSocketFlowClass struct {
	class objc.Class
}

// An interface definition for the [NEFilterSocketFlow] class.
type INEFilterSocketFlow interface {
	INEFilterFlow
}

// A flow of network data that the filter examines.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterSocketFlow
type NEFilterSocketFlow struct {
	NEFilterFlow
}

// NEFilterSocketFlowFrom constructs a [NEFilterSocketFlow] from an unsafe.Pointer.
//
// A flow of network data that the filter examines.
func NEFilterSocketFlowFrom(ptr unsafe.Pointer) NEFilterSocketFlow {
	return NEFilterSocketFlow{
		NEFilterFlow: NEFilterFlowFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NEFilterSocketFlowClass) Alloc() NEFilterSocketFlow {
	rv := objc.Send[NEFilterSocketFlow](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEFilterSocketFlowClass) New() NEFilterSocketFlow {
	rv := objc.Send[NEFilterSocketFlow](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEFilterSocketFlow) Init() NEFilterSocketFlow {
	rv := objc.Send[NEFilterSocketFlow](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEFilterSocketFlow) Autorelease() NEFilterSocketFlow {
	rv := objc.Send[NEFilterSocketFlow](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterSocketFlow creates a new NEFilterSocketFlow instance.
func NewNEFilterSocketFlow() NEFilterSocketFlow {
	return getNEFilterSocketFlowClass().New()
}




