// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MTRDoorLockClusterLockUserChangeEvent] class.
type IMTRDoorLockClusterLockUserChangeEvent interface {
	objectivec.IObject
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
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockUserChangeEvent
type MTRDoorLockClusterLockUserChangeEvent struct {
	objectivec.Object
}

// MTRDoorLockClusterLockUserChangeEventFrom constructs a [MTRDoorLockClusterLockUserChangeEvent] from an unsafe.Pointer.
func MTRDoorLockClusterLockUserChangeEventFrom(ptr unsafe.Pointer) MTRDoorLockClusterLockUserChangeEvent {
	return MTRDoorLockClusterLockUserChangeEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterLockUserChangeEventClass) Alloc() MTRDoorLockClusterLockUserChangeEvent {
	rv := objc.Send[MTRDoorLockClusterLockUserChangeEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockuserchangeevent/dataindex
func (m_ MTRDoorLockClusterLockUserChangeEvent) DataIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("dataIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockuserchangeevent/dataindex
func (m_ MTRDoorLockClusterLockUserChangeEvent) SetDataIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDataIndex:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockuserchangeevent/dataoperationtype
func (m_ MTRDoorLockClusterLockUserChangeEvent) DataOperationType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("dataOperationType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockuserchangeevent/dataoperationtype
func (m_ MTRDoorLockClusterLockUserChangeEvent) SetDataOperationType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDataOperationType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockuserchangeevent/fabricindex
func (m_ MTRDoorLockClusterLockUserChangeEvent) FabricIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockuserchangeevent/fabricindex
func (m_ MTRDoorLockClusterLockUserChangeEvent) SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockuserchangeevent/lockdatatype
func (m_ MTRDoorLockClusterLockUserChangeEvent) LockDataType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("lockDataType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockuserchangeevent/lockdatatype
func (m_ MTRDoorLockClusterLockUserChangeEvent) SetLockDataType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLockDataType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockuserchangeevent/operationsource
func (m_ MTRDoorLockClusterLockUserChangeEvent) OperationSource() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("operationSource"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockuserchangeevent/operationsource
func (m_ MTRDoorLockClusterLockUserChangeEvent) SetOperationSource(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationSource:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockuserchangeevent/sourcenode
func (m_ MTRDoorLockClusterLockUserChangeEvent) SourceNode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("sourceNode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockuserchangeevent/sourcenode
func (m_ MTRDoorLockClusterLockUserChangeEvent) SetSourceNode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceNode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockuserchangeevent/userindex
func (m_ MTRDoorLockClusterLockUserChangeEvent) UserIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockuserchangeevent/userindex
func (m_ MTRDoorLockClusterLockUserChangeEvent) SetUserIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}



