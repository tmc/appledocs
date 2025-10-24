// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDoorLockClusterLockUserChangeEvent */


/* debug [class_header]: Header for MTRDoorLockClusterLockUserChangeEvent */
// The class instance for the [MTRDoorLockClusterLockUserChangeEvent] class.
var (
	MTRDoorLockClusterLockUserChangeEventClass     _MTRDoorLockClusterLockUserChangeEventClass
	MTRDoorLockClusterLockUserChangeEventClassOnce sync.Once
)

func getMTRDoorLockClusterLockUserChangeEventClass() _MTRDoorLockClusterLockUserChangeEventClass {
	MTRDoorLockClusterLockUserChangeEventClassOnce.Do(func() {
		MTRDoorLockClusterLockUserChangeEventClass = _MTRDoorLockClusterLockUserChangeEventClass{objc.GetClass("MTRDoorLockClusterLockUserChangeEvent")}
	})
	return MTRDoorLockClusterLockUserChangeEventClass
}

type _MTRDoorLockClusterLockUserChangeEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDoorLockClusterLockUserChangeEvent */
// An interface definition for the [MTRDoorLockClusterLockUserChangeEvent] class.
type IMTRDoorLockClusterLockUserChangeEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDoorLockClusterLockUserChangeEvent */
	// properties:
	DataIndex() objc.IObject /* cross-framework: NSNumber */
	SetDataIndex(value objc.IObject /* cross-framework: NSNumber */)
	DataOperationType() objc.IObject /* cross-framework: NSNumber */
	SetDataOperationType(value objc.IObject /* cross-framework: NSNumber */)
	FabricIndex() objc.IObject /* cross-framework: NSNumber */
	SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */)
	LockDataType() objc.IObject /* cross-framework: NSNumber */
	SetLockDataType(value objc.IObject /* cross-framework: NSNumber */)
	OperationSource() objc.IObject /* cross-framework: NSNumber */
	SetOperationSource(value objc.IObject /* cross-framework: NSNumber */)
	SourceNode() objc.IObject /* cross-framework: NSNumber */
	SetSourceNode(value objc.IObject /* cross-framework: NSNumber */)
	UserIndex() objc.IObject /* cross-framework: NSNumber */
	SetUserIndex(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDoorLockClusterLockUserChangeEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDoorLockClusterLockUserChangeEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterLockUserChangeEventClass) Alloc() MTRDoorLockClusterLockUserChangeEvent {
	rv := objc.Send[MTRDoorLockClusterLockUserChangeEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDoorLockClusterLockUserChangeEventClass) New() MTRDoorLockClusterLockUserChangeEvent {
	rv := objc.Send[MTRDoorLockClusterLockUserChangeEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterLockUserChangeEvent) Init() MTRDoorLockClusterLockUserChangeEvent {
	rv := objc.Send[MTRDoorLockClusterLockUserChangeEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterLockUserChangeEvent) Autorelease() MTRDoorLockClusterLockUserChangeEvent {
	rv := objc.Send[MTRDoorLockClusterLockUserChangeEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterLockUserChangeEvent creates a new MTRDoorLockClusterLockUserChangeEvent instance.
func NewMTRDoorLockClusterLockUserChangeEvent() MTRDoorLockClusterLockUserChangeEvent {
	return getMTRDoorLockClusterLockUserChangeEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDoorLockClusterLockUserChangeEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockUserChangeEvent
type MTRDoorLockClusterLockUserChangeEvent struct {
	objectivec.Object
}

// MTRDoorLockClusterLockUserChangeEventFrom constructs a [MTRDoorLockClusterLockUserChangeEvent] from an unsafe.Pointer.
func MTRDoorLockClusterLockUserChangeEventFrom(ptr unsafe.Pointer) MTRDoorLockClusterLockUserChangeEvent {
	return MTRDoorLockClusterLockUserChangeEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDoorLockClusterLockUserChangeEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDoorLockClusterLockUserChangeEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDoorLockClusterLockUserChangeEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDoorLockClusterLockUserChangeEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDoorLockClusterLockUserChangeEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockUserChangeEvent/dataIndex
func (m_ MTRDoorLockClusterLockUserChangeEvent) DataIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("dataIndex"))
	return rv
}/* debug [instance_properties/getter]: dataIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockUserChangeEvent/dataIndex
func (m_ MTRDoorLockClusterLockUserChangeEvent) SetDataIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDataIndex:"), value)
}/* debug [instance_properties/setter]: dataIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockUserChangeEvent/dataOperationType
func (m_ MTRDoorLockClusterLockUserChangeEvent) DataOperationType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("dataOperationType"))
	return rv
}/* debug [instance_properties/getter]: dataOperationType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockUserChangeEvent/dataOperationType
func (m_ MTRDoorLockClusterLockUserChangeEvent) SetDataOperationType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDataOperationType:"), value)
}/* debug [instance_properties/setter]: dataOperationType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockUserChangeEvent/fabricIndex
func (m_ MTRDoorLockClusterLockUserChangeEvent) FabricIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricIndex"))
	return rv
}/* debug [instance_properties/getter]: fabricIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockUserChangeEvent/fabricIndex
func (m_ MTRDoorLockClusterLockUserChangeEvent) SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}/* debug [instance_properties/setter]: fabricIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockUserChangeEvent/lockDataType
func (m_ MTRDoorLockClusterLockUserChangeEvent) LockDataType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("lockDataType"))
	return rv
}/* debug [instance_properties/getter]: lockDataType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockUserChangeEvent/lockDataType
func (m_ MTRDoorLockClusterLockUserChangeEvent) SetLockDataType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLockDataType:"), value)
}/* debug [instance_properties/setter]: lockDataType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockUserChangeEvent/operationSource
func (m_ MTRDoorLockClusterLockUserChangeEvent) OperationSource() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("operationSource"))
	return rv
}/* debug [instance_properties/getter]: operationSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockUserChangeEvent/operationSource
func (m_ MTRDoorLockClusterLockUserChangeEvent) SetOperationSource(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationSource:"), value)
}/* debug [instance_properties/setter]: operationSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockUserChangeEvent/sourceNode
func (m_ MTRDoorLockClusterLockUserChangeEvent) SourceNode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("sourceNode"))
	return rv
}/* debug [instance_properties/getter]: sourceNode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockUserChangeEvent/sourceNode
func (m_ MTRDoorLockClusterLockUserChangeEvent) SetSourceNode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceNode:"), value)
}/* debug [instance_properties/setter]: sourceNode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockUserChangeEvent/userIndex
func (m_ MTRDoorLockClusterLockUserChangeEvent) UserIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userIndex"))
	return rv
}/* debug [instance_properties/getter]: userIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockUserChangeEvent/userIndex
func (m_ MTRDoorLockClusterLockUserChangeEvent) SetUserIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}/* debug [instance_properties/setter]: userIndex */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDoorLockClusterLockUserChangeEvent */



