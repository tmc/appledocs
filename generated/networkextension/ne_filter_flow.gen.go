// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NEFilterFlow */


/* debug [class_header]: Header for NEFilterFlow */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEFilterFlow */
// An interface definition for the [NEFilterFlow] class.
type INEFilterFlow interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NEFilterFlow */
	// properties:
	Direction() NETrafficDirection
	Identifier() foundation.UUID
	SourceAppAuditToken() objc.IObject /* cross-framework: NSData */
	SourceProcessAuditToken() objc.IObject /* cross-framework: NSData */
	URL() objc.IObject /* cross-framework: NSURL */
	NEFilterFlowBytesMax() uint64
	SetNEFilterFlowBytesMax(value uint64)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEFilterFlow */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEFilterFlow */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEFilterFlow */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEFilterFlow *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEFilterFlow */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEFilterFlow */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEFilterFlow */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEFilterFlow */

// The initial direction of the flow: incoming or outgoing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterFlow/direction
func (n_ NEFilterFlow) Direction() NETrafficDirection {
	rv := objc.Send[NETrafficDirection](n_.ID, objc.Sel("direction"))
	return rv
}/* debug [instance_properties/getter]: direction */


// The unique identifier of the flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterFlow/identifier
func (n_ NEFilterFlow) Identifier() foundation.UUID {
	rv := objc.Send[foundation.UUID](n_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// The audit token of the source application of the flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterFlow/sourceAppAuditToken
func (n_ NEFilterFlow) SourceAppAuditToken() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](n_.ID, objc.Sel("sourceAppAuditToken"))
	return rv
}/* debug [instance_properties/getter]: sourceAppAuditToken */


// The audit token of the process that created the flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterFlow/sourceProcessAuditToken
func (n_ NEFilterFlow) SourceProcessAuditToken() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](n_.ID, objc.Sel("sourceProcessAuditToken"))
	return rv
}/* debug [instance_properties/getter]: sourceProcessAuditToken */


// The flow’s HTTP URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterFlow/url
func (n_ NEFilterFlow) URL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](n_.ID, objc.Sel("URL"))
	return rv
}/* debug [instance_properties/getter]: URL */


// The maximum number of bytes to pass or peek for a flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterflowbytesmax
func (n_ NEFilterFlow) NEFilterFlowBytesMax() uint64 {
	rv := objc.Send[uint64](n_.ID, objc.Sel("NEFilterFlowBytesMax"))
	return rv
}/* debug [instance_properties/getter]: NEFilterFlowBytesMax */


// The maximum number of bytes to pass or peek for a flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterflowbytesmax
func (n_ NEFilterFlow) SetNEFilterFlowBytesMax(value uint64) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNEFilterFlowBytesMax:"), value)
}/* debug [instance_properties/setter]: NEFilterFlowBytesMax */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEFilterFlow */


