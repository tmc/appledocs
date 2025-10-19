// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NEFilterFlow] class.
var (
	nEFilterFlowClass     _NEFilterFlowClass
	nEFilterFlowClassOnce sync.Once
)

func getNEFilterFlowClass() _NEFilterFlowClass {
	nEFilterFlowClassOnce.Do(func() {
		nEFilterFlowClass = _NEFilterFlowClass{objc.GetClass("NEFilterFlow")}
	})
	return nEFilterFlowClass
}

type _NEFilterFlowClass struct {
	class objc.Class
}

// An interface definition for the [NEFilterFlow] class.
type INEFilterFlow interface {
	objectivec.IObject
}

// The abstract base class for types that represent flows of network data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterFlow
type NEFilterFlow struct {
	objectivec.Object
}

// NEFilterFlowFrom constructs a [NEFilterFlow] from an unsafe.Pointer.
//
// The abstract base class for types that represent flows of network data.
func NEFilterFlowFrom(ptr unsafe.Pointer) NEFilterFlow {
	return NEFilterFlow{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NEFilterFlowClass) Alloc() NEFilterFlow {
	rv := objc.Send[NEFilterFlow](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEFilterFlowClass) New() NEFilterFlow {
	rv := objc.Send[NEFilterFlow](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEFilterFlow) Init() NEFilterFlow {
	rv := objc.Send[NEFilterFlow](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEFilterFlow) Autorelease() NEFilterFlow {
	rv := objc.Send[NEFilterFlow](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterFlow creates a new NEFilterFlow instance.
func NewNEFilterFlow() NEFilterFlow {
	return getNEFilterFlowClass().New()
}




