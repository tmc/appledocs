// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRCommissionerControlClusterCommissioningRequestResultEvent */


/* debug [class_header]: Header for MTRCommissionerControlClusterCommissioningRequestResultEvent */
// The class instance for the [MTRCommissionerControlClusterCommissioningRequestResultEvent] class.
var (
	MTRCommissionerControlClusterCommissioningRequestResultEventClass     _MTRCommissionerControlClusterCommissioningRequestResultEventClass
	MTRCommissionerControlClusterCommissioningRequestResultEventClassOnce sync.Once
)

func getMTRCommissionerControlClusterCommissioningRequestResultEventClass() _MTRCommissionerControlClusterCommissioningRequestResultEventClass {
	MTRCommissionerControlClusterCommissioningRequestResultEventClassOnce.Do(func() {
		MTRCommissionerControlClusterCommissioningRequestResultEventClass = _MTRCommissionerControlClusterCommissioningRequestResultEventClass{objc.GetClass("MTRCommissionerControlClusterCommissioningRequestResultEvent")}
	})
	return MTRCommissionerControlClusterCommissioningRequestResultEventClass
}

type _MTRCommissionerControlClusterCommissioningRequestResultEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRCommissionerControlClusterCommissioningRequestResultEvent */
// An interface definition for the [MTRCommissionerControlClusterCommissioningRequestResultEvent] class.
type IMTRCommissionerControlClusterCommissioningRequestResultEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRCommissionerControlClusterCommissioningRequestResultEvent */
	// properties:
	ClientNodeID() objc.IObject /* cross-framework: NSNumber */
	SetClientNodeID(value objc.IObject /* cross-framework: NSNumber */)
	FabricIndex() objc.IObject /* cross-framework: NSNumber */
	SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */)
	RequestID() objc.IObject /* cross-framework: NSNumber */
	SetRequestID(value objc.IObject /* cross-framework: NSNumber */)
	StatusCode() objc.IObject /* cross-framework: NSNumber */
	SetStatusCode(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRCommissionerControlClusterCommissioningRequestResultEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRCommissionerControlClusterCommissioningRequestResultEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRCommissionerControlClusterCommissioningRequestResultEventClass) Alloc() MTRCommissionerControlClusterCommissioningRequestResultEvent {
	rv := objc.Send[MTRCommissionerControlClusterCommissioningRequestResultEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRCommissionerControlClusterCommissioningRequestResultEventClass) New() MTRCommissionerControlClusterCommissioningRequestResultEvent {
	rv := objc.Send[MTRCommissionerControlClusterCommissioningRequestResultEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRCommissionerControlClusterCommissioningRequestResultEvent) Init() MTRCommissionerControlClusterCommissioningRequestResultEvent {
	rv := objc.Send[MTRCommissionerControlClusterCommissioningRequestResultEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRCommissionerControlClusterCommissioningRequestResultEvent) Autorelease() MTRCommissionerControlClusterCommissioningRequestResultEvent {
	rv := objc.Send[MTRCommissionerControlClusterCommissioningRequestResultEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRCommissionerControlClusterCommissioningRequestResultEvent creates a new MTRCommissionerControlClusterCommissioningRequestResultEvent instance.
func NewMTRCommissionerControlClusterCommissioningRequestResultEvent() MTRCommissionerControlClusterCommissioningRequestResultEvent {
	return getMTRCommissionerControlClusterCommissioningRequestResultEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRCommissionerControlClusterCommissioningRequestResultEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterCommissioningRequestResultEvent
type MTRCommissionerControlClusterCommissioningRequestResultEvent struct {
	objectivec.Object
}

// MTRCommissionerControlClusterCommissioningRequestResultEventFrom constructs a [MTRCommissionerControlClusterCommissioningRequestResultEvent] from an unsafe.Pointer.
func MTRCommissionerControlClusterCommissioningRequestResultEventFrom(ptr unsafe.Pointer) MTRCommissionerControlClusterCommissioningRequestResultEvent {
	return MTRCommissionerControlClusterCommissioningRequestResultEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRCommissionerControlClusterCommissioningRequestResultEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRCommissionerControlClusterCommissioningRequestResultEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRCommissionerControlClusterCommissioningRequestResultEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRCommissionerControlClusterCommissioningRequestResultEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRCommissionerControlClusterCommissioningRequestResultEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterCommissioningRequestResultEvent/clientNodeID
func (m_ MTRCommissionerControlClusterCommissioningRequestResultEvent) ClientNodeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("clientNodeID"))
	return rv
}/* debug [instance_properties/getter]: clientNodeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterCommissioningRequestResultEvent/clientNodeID
func (m_ MTRCommissionerControlClusterCommissioningRequestResultEvent) SetClientNodeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setClientNodeID:"), value)
}/* debug [instance_properties/setter]: clientNodeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissionercontrolclustercommissioningrequestresultevent/fabricindex
func (m_ MTRCommissionerControlClusterCommissioningRequestResultEvent) FabricIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricIndex"))
	return rv
}/* debug [instance_properties/getter]: fabricIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissionercontrolclustercommissioningrequestresultevent/fabricindex
func (m_ MTRCommissionerControlClusterCommissioningRequestResultEvent) SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}/* debug [instance_properties/setter]: fabricIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissionercontrolclustercommissioningrequestresultevent/requestid
func (m_ MTRCommissionerControlClusterCommissioningRequestResultEvent) RequestID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("requestID"))
	return rv
}/* debug [instance_properties/getter]: requestID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissionercontrolclustercommissioningrequestresultevent/requestid
func (m_ MTRCommissionerControlClusterCommissioningRequestResultEvent) SetRequestID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequestID:"), value)
}/* debug [instance_properties/setter]: requestID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissionercontrolclustercommissioningrequestresultevent/statuscode
func (m_ MTRCommissionerControlClusterCommissioningRequestResultEvent) StatusCode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("statusCode"))
	return rv
}/* debug [instance_properties/getter]: statusCode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissionercontrolclustercommissioningrequestresultevent/statuscode
func (m_ MTRCommissionerControlClusterCommissioningRequestResultEvent) SetStatusCode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatusCode:"), value)
}/* debug [instance_properties/setter]: statusCode */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRCommissionerControlClusterCommissioningRequestResultEvent */



