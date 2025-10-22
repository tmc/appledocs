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
	DataIndex() foundation.Number
	SetDataIndex(value foundation.INumber)
	DataOperationType() foundation.Number
	SetDataOperationType(value foundation.INumber)
	FabricIndex() foundation.Number
	SetFabricIndex(value foundation.INumber)
	LockDataType() foundation.Number
	SetLockDataType(value foundation.INumber)
	OperationSource() foundation.Number
	SetOperationSource(value foundation.INumber)
	SourceNode() foundation.Number
	SetSourceNode(value foundation.INumber)
	UserIndex() foundation.Number
	SetUserIndex(value foundation.INumber)
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockuserchangeevent/dataindex
func (m_ MTRDoorLockClusterLockUserChangeEvent) DataIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("dataIndex"))
	return rv
}


// SetDataIndex sets the value of the dataIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockuserchangeevent/dataindex
func (m_ MTRDoorLockClusterLockUserChangeEvent) SetDataIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDataIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockuserchangeevent/dataoperationtype
func (m_ MTRDoorLockClusterLockUserChangeEvent) DataOperationType() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("dataOperationType"))
	return rv
}


// SetDataOperationType sets the value of the dataOperationType property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockuserchangeevent/dataoperationtype
func (m_ MTRDoorLockClusterLockUserChangeEvent) SetDataOperationType(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDataOperationType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockuserchangeevent/fabricindex
func (m_ MTRDoorLockClusterLockUserChangeEvent) FabricIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("fabricIndex"))
	return rv
}


// SetFabricIndex sets the value of the fabricIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockuserchangeevent/fabricindex
func (m_ MTRDoorLockClusterLockUserChangeEvent) SetFabricIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockuserchangeevent/lockdatatype
func (m_ MTRDoorLockClusterLockUserChangeEvent) LockDataType() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("lockDataType"))
	return rv
}


// SetLockDataType sets the value of the lockDataType property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockuserchangeevent/lockdatatype
func (m_ MTRDoorLockClusterLockUserChangeEvent) SetLockDataType(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLockDataType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockuserchangeevent/operationsource
func (m_ MTRDoorLockClusterLockUserChangeEvent) OperationSource() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("operationSource"))
	return rv
}


// SetOperationSource sets the value of the operationSource property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockuserchangeevent/operationsource
func (m_ MTRDoorLockClusterLockUserChangeEvent) SetOperationSource(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationSource:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockuserchangeevent/sourcenode
func (m_ MTRDoorLockClusterLockUserChangeEvent) SourceNode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("sourceNode"))
	return rv
}


// SetSourceNode sets the value of the sourceNode property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockuserchangeevent/sourcenode
func (m_ MTRDoorLockClusterLockUserChangeEvent) SetSourceNode(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceNode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockuserchangeevent/userindex
func (m_ MTRDoorLockClusterLockUserChangeEvent) UserIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("userIndex"))
	return rv
}


// SetUserIndex sets the value of the userIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockuserchangeevent/userindex
func (m_ MTRDoorLockClusterLockUserChangeEvent) SetUserIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}



