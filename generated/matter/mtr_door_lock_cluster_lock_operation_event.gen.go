// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDoorLockClusterLockOperationEvent */


/* debug [class_header]: Header for MTRDoorLockClusterLockOperationEvent */
// The class instance for the [MTRDoorLockClusterLockOperationEvent] class.
var (
	MTRDoorLockClusterLockOperationEventClass     _MTRDoorLockClusterLockOperationEventClass
	MTRDoorLockClusterLockOperationEventClassOnce sync.Once
)

func getMTRDoorLockClusterLockOperationEventClass() _MTRDoorLockClusterLockOperationEventClass {
	MTRDoorLockClusterLockOperationEventClassOnce.Do(func() {
		MTRDoorLockClusterLockOperationEventClass = _MTRDoorLockClusterLockOperationEventClass{objc.GetClass("MTRDoorLockClusterLockOperationEvent")}
	})
	return MTRDoorLockClusterLockOperationEventClass
}

type _MTRDoorLockClusterLockOperationEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDoorLockClusterLockOperationEvent */
// An interface definition for the [MTRDoorLockClusterLockOperationEvent] class.
type IMTRDoorLockClusterLockOperationEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDoorLockClusterLockOperationEvent */
	// properties:
	Credentials() objc.IObject /* cross-framework: NSArray */
	SetCredentials(value objc.IObject /* cross-framework: NSArray */)
	FabricIndex() objc.IObject /* cross-framework: NSNumber */
	SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */)
	LockOperationType() objc.IObject /* cross-framework: NSNumber */
	SetLockOperationType(value objc.IObject /* cross-framework: NSNumber */)
	OperationSource() objc.IObject /* cross-framework: NSNumber */
	SetOperationSource(value objc.IObject /* cross-framework: NSNumber */)
	SourceNode() objc.IObject /* cross-framework: NSNumber */
	SetSourceNode(value objc.IObject /* cross-framework: NSNumber */)
	UserIndex() objc.IObject /* cross-framework: NSNumber */
	SetUserIndex(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDoorLockClusterLockOperationEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDoorLockClusterLockOperationEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterLockOperationEventClass) Alloc() MTRDoorLockClusterLockOperationEvent {
	rv := objc.Send[MTRDoorLockClusterLockOperationEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDoorLockClusterLockOperationEventClass) New() MTRDoorLockClusterLockOperationEvent {
	rv := objc.Send[MTRDoorLockClusterLockOperationEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterLockOperationEvent) Init() MTRDoorLockClusterLockOperationEvent {
	rv := objc.Send[MTRDoorLockClusterLockOperationEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterLockOperationEvent) Autorelease() MTRDoorLockClusterLockOperationEvent {
	rv := objc.Send[MTRDoorLockClusterLockOperationEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterLockOperationEvent creates a new MTRDoorLockClusterLockOperationEvent instance.
func NewMTRDoorLockClusterLockOperationEvent() MTRDoorLockClusterLockOperationEvent {
	return getMTRDoorLockClusterLockOperationEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDoorLockClusterLockOperationEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockOperationEvent
type MTRDoorLockClusterLockOperationEvent struct {
	objectivec.Object
}

// MTRDoorLockClusterLockOperationEventFrom constructs a [MTRDoorLockClusterLockOperationEvent] from an unsafe.Pointer.
func MTRDoorLockClusterLockOperationEventFrom(ptr unsafe.Pointer) MTRDoorLockClusterLockOperationEvent {
	return MTRDoorLockClusterLockOperationEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDoorLockClusterLockOperationEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDoorLockClusterLockOperationEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDoorLockClusterLockOperationEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDoorLockClusterLockOperationEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDoorLockClusterLockOperationEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockOperationEvent/credentials
func (m_ MTRDoorLockClusterLockOperationEvent) Credentials() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("credentials"))
	return rv
}/* debug [instance_properties/getter]: credentials */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockOperationEvent/credentials
func (m_ MTRDoorLockClusterLockOperationEvent) SetCredentials(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCredentials:"), value)
}/* debug [instance_properties/setter]: credentials */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockOperationEvent/fabricIndex
func (m_ MTRDoorLockClusterLockOperationEvent) FabricIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricIndex"))
	return rv
}/* debug [instance_properties/getter]: fabricIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockOperationEvent/fabricIndex
func (m_ MTRDoorLockClusterLockOperationEvent) SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}/* debug [instance_properties/setter]: fabricIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockOperationEvent/lockOperationType
func (m_ MTRDoorLockClusterLockOperationEvent) LockOperationType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("lockOperationType"))
	return rv
}/* debug [instance_properties/getter]: lockOperationType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockOperationEvent/lockOperationType
func (m_ MTRDoorLockClusterLockOperationEvent) SetLockOperationType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLockOperationType:"), value)
}/* debug [instance_properties/setter]: lockOperationType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockOperationEvent/operationSource
func (m_ MTRDoorLockClusterLockOperationEvent) OperationSource() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("operationSource"))
	return rv
}/* debug [instance_properties/getter]: operationSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockOperationEvent/operationSource
func (m_ MTRDoorLockClusterLockOperationEvent) SetOperationSource(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationSource:"), value)
}/* debug [instance_properties/setter]: operationSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockOperationEvent/sourceNode
func (m_ MTRDoorLockClusterLockOperationEvent) SourceNode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("sourceNode"))
	return rv
}/* debug [instance_properties/getter]: sourceNode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockOperationEvent/sourceNode
func (m_ MTRDoorLockClusterLockOperationEvent) SetSourceNode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceNode:"), value)
}/* debug [instance_properties/setter]: sourceNode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockOperationEvent/userIndex
func (m_ MTRDoorLockClusterLockOperationEvent) UserIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userIndex"))
	return rv
}/* debug [instance_properties/getter]: userIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockOperationEvent/userIndex
func (m_ MTRDoorLockClusterLockOperationEvent) SetUserIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}/* debug [instance_properties/setter]: userIndex */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDoorLockClusterLockOperationEvent */



