// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MTRDoorLockClusterLockOperationErrorEvent] class.
type IMTRDoorLockClusterLockOperationErrorEvent interface {
	objectivec.IObject
	// properties:
	Credentials() unsafe.Pointer
	SetCredentials(value unsafe.Pointer)
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
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockOperationErrorEvent
type MTRDoorLockClusterLockOperationErrorEvent struct {
	objectivec.Object
}

// MTRDoorLockClusterLockOperationErrorEventFrom constructs a [MTRDoorLockClusterLockOperationErrorEvent] from an unsafe.Pointer.
func MTRDoorLockClusterLockOperationErrorEventFrom(ptr unsafe.Pointer) MTRDoorLockClusterLockOperationErrorEvent {
	return MTRDoorLockClusterLockOperationErrorEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterLockOperationErrorEventClass) Alloc() MTRDoorLockClusterLockOperationErrorEvent {
	rv := objc.Send[MTRDoorLockClusterLockOperationErrorEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockoperationerrorevent/credentials
func (m_ MTRDoorLockClusterLockOperationErrorEvent) Credentials() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("credentials"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockoperationerrorevent/credentials
func (m_ MTRDoorLockClusterLockOperationErrorEvent) SetCredentials(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCredentials:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockoperationerrorevent/fabricindex
func (m_ MTRDoorLockClusterLockOperationErrorEvent) FabricIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockoperationerrorevent/fabricindex
func (m_ MTRDoorLockClusterLockOperationErrorEvent) SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockoperationerrorevent/lockoperationtype
func (m_ MTRDoorLockClusterLockOperationErrorEvent) LockOperationType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("lockOperationType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockoperationerrorevent/lockoperationtype
func (m_ MTRDoorLockClusterLockOperationErrorEvent) SetLockOperationType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLockOperationType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockoperationerrorevent/operationerror
func (m_ MTRDoorLockClusterLockOperationErrorEvent) OperationError() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("operationError"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockoperationerrorevent/operationerror
func (m_ MTRDoorLockClusterLockOperationErrorEvent) SetOperationError(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationError:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockoperationerrorevent/operationsource
func (m_ MTRDoorLockClusterLockOperationErrorEvent) OperationSource() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("operationSource"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockoperationerrorevent/operationsource
func (m_ MTRDoorLockClusterLockOperationErrorEvent) SetOperationSource(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationSource:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockoperationerrorevent/sourcenode
func (m_ MTRDoorLockClusterLockOperationErrorEvent) SourceNode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("sourceNode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockoperationerrorevent/sourcenode
func (m_ MTRDoorLockClusterLockOperationErrorEvent) SetSourceNode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceNode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockoperationerrorevent/userindex
func (m_ MTRDoorLockClusterLockOperationErrorEvent) UserIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterlockoperationerrorevent/userindex
func (m_ MTRDoorLockClusterLockOperationErrorEvent) SetUserIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}



