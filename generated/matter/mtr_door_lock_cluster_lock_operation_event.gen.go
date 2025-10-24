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
	// properties:
	Credentials() unsafe.Pointer
	SetCredentials(value unsafe.Pointer)
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
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockoperationevent/credentials
func (m_ MTRDoorLockClusterLockOperationEvent) Credentials() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("credentials"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockoperationevent/credentials
func (m_ MTRDoorLockClusterLockOperationEvent) SetCredentials(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCredentials:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockoperationevent/fabricindex
func (m_ MTRDoorLockClusterLockOperationEvent) FabricIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockoperationevent/fabricindex
func (m_ MTRDoorLockClusterLockOperationEvent) SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockoperationevent/lockoperationtype
func (m_ MTRDoorLockClusterLockOperationEvent) LockOperationType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("lockOperationType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockoperationevent/lockoperationtype
func (m_ MTRDoorLockClusterLockOperationEvent) SetLockOperationType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLockOperationType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockoperationevent/operationsource
func (m_ MTRDoorLockClusterLockOperationEvent) OperationSource() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("operationSource"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockoperationevent/operationsource
func (m_ MTRDoorLockClusterLockOperationEvent) SetOperationSource(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationSource:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockoperationevent/sourcenode
func (m_ MTRDoorLockClusterLockOperationEvent) SourceNode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("sourceNode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockoperationevent/sourcenode
func (m_ MTRDoorLockClusterLockOperationEvent) SetSourceNode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceNode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockoperationevent/userindex
func (m_ MTRDoorLockClusterLockOperationEvent) UserIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockoperationevent/userindex
func (m_ MTRDoorLockClusterLockOperationEvent) SetUserIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}



