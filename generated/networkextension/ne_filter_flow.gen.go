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
	NEFilterFlowClass     _NEFilterFlowClass
	NEFilterFlowClassOnce sync.Once
)

func getNEFilterFlowClass() _NEFilterFlowClass {
	NEFilterFlowClassOnce.Do(func() {
		NEFilterFlowClass = _NEFilterFlowClass{objc.GetClass("NEFilterFlow")}
	})
	return NEFilterFlowClass
}

type _NEFilterFlowClass struct {
	class objc.Class
}





// An interface definition for the [NEFilterFlow] class.
type INEFilterFlow interface {
	objectivec.IObject
	

	// properties:
	Direction() NETrafficDirection
	Identifier() foundation.UUID
	SourceAppAuditToken() foundation.foundation.INSData
	SourceProcessAuditToken() foundation.foundation.INSData
	URL() foundation.foundation.INSURL
	NEFilterFlowBytesMax() uint64
	SetNEFilterFlowBytesMax(value uint64)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (nc _NEFilterFlowClass) Alloc() NEFilterFlow {
	rv := objc.Send[NEFilterFlow](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// The abstract base class for types that represent flows of network data.


// The abstract base class for types that represent flows of network data.
//
// [Full Topic]
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

























// The initial direction of the flow: incoming or outgoing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterFlow/direction
func (n_ NEFilterFlow) Direction() NETrafficDirection {
	rv := objc.Send[NETrafficDirection](n_.ID, objc.Sel("direction"))
	return rv
}


// The unique identifier of the flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterFlow/identifier
func (n_ NEFilterFlow) Identifier() foundation.UUID {
	rv := objc.Send[foundation.UUID](n_.ID, objc.Sel("identifier"))
	return rv
}


// The audit token of the source application of the flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterFlow/sourceAppAuditToken
func (n_ NEFilterFlow) SourceAppAuditToken() foundation.foundation.INSData {
	rv := objc.Send[foundation.NSData](n_.ID, objc.Sel("sourceAppAuditToken"))
	return rv
}


// The audit token of the process that created the flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterFlow/sourceProcessAuditToken
func (n_ NEFilterFlow) SourceProcessAuditToken() foundation.foundation.INSData {
	rv := objc.Send[foundation.NSData](n_.ID, objc.Sel("sourceProcessAuditToken"))
	return rv
}


// The flow’s HTTP URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterFlow/url
func (n_ NEFilterFlow) URL() foundation.foundation.INSURL {
	rv := objc.Send[foundation.NSURL](n_.ID, objc.Sel("URL"))
	return rv
}


// The maximum number of bytes to pass or peek for a flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterflowbytesmax
func (n_ NEFilterFlow) NEFilterFlowBytesMax() uint64 {
	rv := objc.Send[uint64](n_.ID, objc.Sel("NEFilterFlowBytesMax"))
	return rv
}


// The maximum number of bytes to pass or peek for a flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterflowbytesmax
func (n_ NEFilterFlow) SetNEFilterFlowBytesMax(value uint64) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNEFilterFlowBytesMax:"), value)
}







