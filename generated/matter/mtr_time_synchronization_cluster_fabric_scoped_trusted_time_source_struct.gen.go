// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct] class.
var (
	MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStructClass     _MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStructClass
	MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStructClassOnce sync.Once
)

func getMTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStructClass() _MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStructClass {
	MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStructClassOnce.Do(func() {
		MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStructClass = _MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStructClass{objc.GetClass("MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct")}
	})
	return MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStructClass
}

type _MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct] class.
type IMTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct
type MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct struct {
	objectivec.Object
}

// MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStructFrom constructs a [MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct] from an unsafe.Pointer.
func MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStructFrom(ptr unsafe.Pointer) MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct {
	return MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStructClass) Alloc() MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct {
	rv := objc.Send[MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStructClass) New() MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct {
	rv := objc.Send[MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct) Init() MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct {
	rv := objc.Send[MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct) Autorelease() MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct {
	rv := objc.Send[MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct creates a new MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct instance.
func NewMTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct() MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct {
	return getMTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct/endpoint
func (m_ MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct) Endpoint() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("endpoint"))
	return rv
}


// SetEndpoint sets the value of the endpoint property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct/endpoint
func (m_ MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct) SetEndpoint(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndpoint:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct/nodeID
func (m_ MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct) NodeID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nodeID"))
	return rv
}


// SetNodeID sets the value of the nodeID property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct/nodeID
func (m_ MTRTimeSynchronizationClusterFabricScopedTrustedTimeSourceStruct) SetNodeID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNodeID:"), value)
}



