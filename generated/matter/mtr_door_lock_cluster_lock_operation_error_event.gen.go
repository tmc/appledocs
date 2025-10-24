// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDoorLockClusterLockOperationErrorEvent */


/* debug [class_header]: Header for MTRDoorLockClusterLockOperationErrorEvent */
// The class instance for the [MTRDoorLockClusterLockOperationErrorEvent] class.
var (
	MTRDoorLockClusterLockOperationErrorEventClass     _MTRDoorLockClusterLockOperationErrorEventClass
	MTRDoorLockClusterLockOperationErrorEventClassOnce sync.Once
)

func getMTRDoorLockClusterLockOperationErrorEventClass() _MTRDoorLockClusterLockOperationErrorEventClass {
	MTRDoorLockClusterLockOperationErrorEventClassOnce.Do(func() {
		MTRDoorLockClusterLockOperationErrorEventClass = _MTRDoorLockClusterLockOperationErrorEventClass{objc.GetClass("MTRDoorLockClusterLockOperationErrorEvent")}
	})
	return MTRDoorLockClusterLockOperationErrorEventClass
}

type _MTRDoorLockClusterLockOperationErrorEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDoorLockClusterLockOperationErrorEvent */
// An interface definition for the [MTRDoorLockClusterLockOperationErrorEvent] class.
type IMTRDoorLockClusterLockOperationErrorEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDoorLockClusterLockOperationErrorEvent */
	// properties:
	Credentials() objc.IObject /* cross-framework: NSArray */
	SetCredentials(value objc.IObject /* cross-framework: NSArray */)
	FabricIndex() objc.IObject /* cross-framework: NSNumber */
	SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */)
	LockOperationType() objc.IObject /* cross-framework: NSNumber */
	SetLockOperationType(value objc.IObject /* cross-framework: NSNumber */)
	OperationError() objc.IObject /* cross-framework: NSNumber */
	SetOperationError(value objc.IObject /* cross-framework: NSNumber */)
	OperationSource() objc.IObject /* cross-framework: NSNumber */
	SetOperationSource(value objc.IObject /* cross-framework: NSNumber */)
	SourceNode() objc.IObject /* cross-framework: NSNumber */
	SetSourceNode(value objc.IObject /* cross-framework: NSNumber */)
	UserIndex() objc.IObject /* cross-framework: NSNumber */
	SetUserIndex(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDoorLockClusterLockOperationErrorEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDoorLockClusterLockOperationErrorEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterLockOperationErrorEventClass) Alloc() MTRDoorLockClusterLockOperationErrorEvent {
	rv := objc.Send[MTRDoorLockClusterLockOperationErrorEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDoorLockClusterLockOperationErrorEventClass) New() MTRDoorLockClusterLockOperationErrorEvent {
	rv := objc.Send[MTRDoorLockClusterLockOperationErrorEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterLockOperationErrorEvent) Init() MTRDoorLockClusterLockOperationErrorEvent {
	rv := objc.Send[MTRDoorLockClusterLockOperationErrorEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterLockOperationErrorEvent) Autorelease() MTRDoorLockClusterLockOperationErrorEvent {
	rv := objc.Send[MTRDoorLockClusterLockOperationErrorEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterLockOperationErrorEvent creates a new MTRDoorLockClusterLockOperationErrorEvent instance.
func NewMTRDoorLockClusterLockOperationErrorEvent() MTRDoorLockClusterLockOperationErrorEvent {
	return getMTRDoorLockClusterLockOperationErrorEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDoorLockClusterLockOperationErrorEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockOperationErrorEvent
type MTRDoorLockClusterLockOperationErrorEvent struct {
	objectivec.Object
}

// MTRDoorLockClusterLockOperationErrorEventFrom constructs a [MTRDoorLockClusterLockOperationErrorEvent] from an unsafe.Pointer.
func MTRDoorLockClusterLockOperationErrorEventFrom(ptr unsafe.Pointer) MTRDoorLockClusterLockOperationErrorEvent {
	return MTRDoorLockClusterLockOperationErrorEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDoorLockClusterLockOperationErrorEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDoorLockClusterLockOperationErrorEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDoorLockClusterLockOperationErrorEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDoorLockClusterLockOperationErrorEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDoorLockClusterLockOperationErrorEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockOperationErrorEvent/credentials
func (m_ MTRDoorLockClusterLockOperationErrorEvent) Credentials() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("credentials"))
	return rv
}/* debug [instance_properties/getter]: credentials */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockOperationErrorEvent/credentials
func (m_ MTRDoorLockClusterLockOperationErrorEvent) SetCredentials(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCredentials:"), value)
}/* debug [instance_properties/setter]: credentials */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockOperationErrorEvent/fabricIndex
func (m_ MTRDoorLockClusterLockOperationErrorEvent) FabricIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricIndex"))
	return rv
}/* debug [instance_properties/getter]: fabricIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockOperationErrorEvent/fabricIndex
func (m_ MTRDoorLockClusterLockOperationErrorEvent) SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}/* debug [instance_properties/setter]: fabricIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockOperationErrorEvent/lockOperationType
func (m_ MTRDoorLockClusterLockOperationErrorEvent) LockOperationType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("lockOperationType"))
	return rv
}/* debug [instance_properties/getter]: lockOperationType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockOperationErrorEvent/lockOperationType
func (m_ MTRDoorLockClusterLockOperationErrorEvent) SetLockOperationType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLockOperationType:"), value)
}/* debug [instance_properties/setter]: lockOperationType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockOperationErrorEvent/operationError
func (m_ MTRDoorLockClusterLockOperationErrorEvent) OperationError() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("operationError"))
	return rv
}/* debug [instance_properties/getter]: operationError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockOperationErrorEvent/operationError
func (m_ MTRDoorLockClusterLockOperationErrorEvent) SetOperationError(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationError:"), value)
}/* debug [instance_properties/setter]: operationError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockOperationErrorEvent/operationSource
func (m_ MTRDoorLockClusterLockOperationErrorEvent) OperationSource() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("operationSource"))
	return rv
}/* debug [instance_properties/getter]: operationSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockOperationErrorEvent/operationSource
func (m_ MTRDoorLockClusterLockOperationErrorEvent) SetOperationSource(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationSource:"), value)
}/* debug [instance_properties/setter]: operationSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockOperationErrorEvent/sourceNode
func (m_ MTRDoorLockClusterLockOperationErrorEvent) SourceNode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("sourceNode"))
	return rv
}/* debug [instance_properties/getter]: sourceNode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockOperationErrorEvent/sourceNode
func (m_ MTRDoorLockClusterLockOperationErrorEvent) SetSourceNode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceNode:"), value)
}/* debug [instance_properties/setter]: sourceNode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockOperationErrorEvent/userIndex
func (m_ MTRDoorLockClusterLockOperationErrorEvent) UserIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userIndex"))
	return rv
}/* debug [instance_properties/getter]: userIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockOperationErrorEvent/userIndex
func (m_ MTRDoorLockClusterLockOperationErrorEvent) SetUserIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}/* debug [instance_properties/setter]: userIndex */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDoorLockClusterLockOperationErrorEvent */



