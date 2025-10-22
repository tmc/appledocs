// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MTRDoorLockClusterLockOperationEvent] class.
type IMTRDoorLockClusterLockOperationEvent interface {
	objectivec.IObject
	Credentials() unsafe.Pointer
	SetCredentials(value unsafe.Pointer)
	FabricIndex() foundation.Number
	SetFabricIndex(value foundation.INumber)
	LockOperationType() foundation.Number
	SetLockOperationType(value foundation.INumber)
	OperationSource() foundation.Number
	SetOperationSource(value foundation.INumber)
	SourceNode() foundation.Number
	SetSourceNode(value foundation.INumber)
	UserIndex() foundation.Number
	SetUserIndex(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockOperationEvent
type MTRDoorLockClusterLockOperationEvent struct {
	objectivec.Object
}

// MTRDoorLockClusterLockOperationEventFrom constructs a [MTRDoorLockClusterLockOperationEvent] from an unsafe.Pointer.
func MTRDoorLockClusterLockOperationEventFrom(ptr unsafe.Pointer) MTRDoorLockClusterLockOperationEvent {
	return MTRDoorLockClusterLockOperationEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterLockOperationEventClass) Alloc() MTRDoorLockClusterLockOperationEvent {
	rv := objc.Send[MTRDoorLockClusterLockOperationEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockoperationevent/credentials
func (m_ MTRDoorLockClusterLockOperationEvent) Credentials() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("credentials"))
	return rv
}


// SetCredentials sets the value of the credentials property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockoperationevent/credentials
func (m_ MTRDoorLockClusterLockOperationEvent) SetCredentials(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCredentials:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockoperationevent/fabricindex
func (m_ MTRDoorLockClusterLockOperationEvent) FabricIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("fabricIndex"))
	return rv
}


// SetFabricIndex sets the value of the fabricIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockoperationevent/fabricindex
func (m_ MTRDoorLockClusterLockOperationEvent) SetFabricIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockoperationevent/lockoperationtype
func (m_ MTRDoorLockClusterLockOperationEvent) LockOperationType() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("lockOperationType"))
	return rv
}


// SetLockOperationType sets the value of the lockOperationType property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockoperationevent/lockoperationtype
func (m_ MTRDoorLockClusterLockOperationEvent) SetLockOperationType(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLockOperationType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockoperationevent/operationsource
func (m_ MTRDoorLockClusterLockOperationEvent) OperationSource() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("operationSource"))
	return rv
}


// SetOperationSource sets the value of the operationSource property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockoperationevent/operationsource
func (m_ MTRDoorLockClusterLockOperationEvent) SetOperationSource(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationSource:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockoperationevent/sourcenode
func (m_ MTRDoorLockClusterLockOperationEvent) SourceNode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("sourceNode"))
	return rv
}


// SetSourceNode sets the value of the sourceNode property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockoperationevent/sourcenode
func (m_ MTRDoorLockClusterLockOperationEvent) SetSourceNode(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceNode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockoperationevent/userindex
func (m_ MTRDoorLockClusterLockOperationEvent) UserIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("userIndex"))
	return rv
}


// SetUserIndex sets the value of the userIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockoperationevent/userindex
func (m_ MTRDoorLockClusterLockOperationEvent) SetUserIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}



