// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NEAppProxyFlow] class.
var (
	nEAppProxyFlowClass     _NEAppProxyFlowClass
	nEAppProxyFlowClassOnce sync.Once
)

func getNEAppProxyFlowClass() _NEAppProxyFlowClass {
	nEAppProxyFlowClassOnce.Do(func() {
		nEAppProxyFlowClass = _NEAppProxyFlowClass{objc.GetClass("NEAppProxyFlow")}
	})
	return nEAppProxyFlowClass
}

type _NEAppProxyFlowClass struct {
	class objc.Class
}

// An interface definition for the [NEAppProxyFlow] class.
type INEAppProxyFlow interface {
	objectivec.IObject
}

// A parent class referenced by other NetworkExtension classes. [Full Topic]
type NEAppProxyFlow struct {
	objectivec.Object
}

// NEAppProxyFlowFrom constructs a [NEAppProxyFlow] from an unsafe.Pointer.
//
// A parent class referenced by other NetworkExtension classes.
func NEAppProxyFlowFrom(ptr unsafe.Pointer) NEAppProxyFlow {
	return NEAppProxyFlow{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NEAppProxyFlowClass) Alloc() NEAppProxyFlow {
	rv := objc.Send[NEAppProxyFlow](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEAppProxyFlowClass) New() NEAppProxyFlow {
	rv := objc.Send[NEAppProxyFlow](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEAppProxyFlow) Init() NEAppProxyFlow {
	rv := objc.Send[NEAppProxyFlow](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEAppProxyFlow) Autorelease() NEAppProxyFlow {
	rv := objc.Send[NEAppProxyFlow](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEAppProxyFlow creates a new NEAppProxyFlow instance.
func NewNEAppProxyFlow() NEAppProxyFlow {
	return getNEAppProxyFlowClass().New()
}




