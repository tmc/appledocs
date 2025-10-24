// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRTimeSynchronizationClusterTrustedTimeSourceStruct] class.
var (
	MTRTimeSynchronizationClusterTrustedTimeSourceStructClass     _MTRTimeSynchronizationClusterTrustedTimeSourceStructClass
	MTRTimeSynchronizationClusterTrustedTimeSourceStructClassOnce sync.Once
)

func getMTRTimeSynchronizationClusterTrustedTimeSourceStructClass() _MTRTimeSynchronizationClusterTrustedTimeSourceStructClass {
	MTRTimeSynchronizationClusterTrustedTimeSourceStructClassOnce.Do(func() {
		MTRTimeSynchronizationClusterTrustedTimeSourceStructClass = _MTRTimeSynchronizationClusterTrustedTimeSourceStructClass{objc.GetClass("MTRTimeSynchronizationClusterTrustedTimeSourceStruct")}
	})
	return MTRTimeSynchronizationClusterTrustedTimeSourceStructClass
}

type _MTRTimeSynchronizationClusterTrustedTimeSourceStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRTimeSynchronizationClusterTrustedTimeSourceStruct] class.
type IMTRTimeSynchronizationClusterTrustedTimeSourceStruct interface {
	objectivec.IObject
	// properties:
	Endpoint() objc.IObject /* cross-framework: NSNumber */
	SetEndpoint(value objc.IObject /* cross-framework: NSNumber */)
	FabricIndex() objc.IObject /* cross-framework: NSNumber */
	SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */)
	NodeID() objc.IObject /* cross-framework: NSNumber */
	SetNodeID(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterTrustedTimeSourceStruct
type MTRTimeSynchronizationClusterTrustedTimeSourceStruct struct {
	objectivec.Object
}

// MTRTimeSynchronizationClusterTrustedTimeSourceStructFrom constructs a [MTRTimeSynchronizationClusterTrustedTimeSourceStruct] from an unsafe.Pointer.
func MTRTimeSynchronizationClusterTrustedTimeSourceStructFrom(ptr unsafe.Pointer) MTRTimeSynchronizationClusterTrustedTimeSourceStruct {
	return MTRTimeSynchronizationClusterTrustedTimeSourceStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTimeSynchronizationClusterTrustedTimeSourceStructClass) Alloc() MTRTimeSynchronizationClusterTrustedTimeSourceStruct {
	rv := objc.Send[MTRTimeSynchronizationClusterTrustedTimeSourceStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTimeSynchronizationClusterTrustedTimeSourceStructClass) New() MTRTimeSynchronizationClusterTrustedTimeSourceStruct {
	rv := objc.Send[MTRTimeSynchronizationClusterTrustedTimeSourceStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTimeSynchronizationClusterTrustedTimeSourceStruct) Init() MTRTimeSynchronizationClusterTrustedTimeSourceStruct {
	rv := objc.Send[MTRTimeSynchronizationClusterTrustedTimeSourceStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTimeSynchronizationClusterTrustedTimeSourceStruct) Autorelease() MTRTimeSynchronizationClusterTrustedTimeSourceStruct {
	rv := objc.Send[MTRTimeSynchronizationClusterTrustedTimeSourceStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTimeSynchronizationClusterTrustedTimeSourceStruct creates a new MTRTimeSynchronizationClusterTrustedTimeSourceStruct instance.
func NewMTRTimeSynchronizationClusterTrustedTimeSourceStruct() MTRTimeSynchronizationClusterTrustedTimeSourceStruct {
	return getMTRTimeSynchronizationClusterTrustedTimeSourceStructClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterTrustedTimeSourceStruct/endpoint
func (m_ MTRTimeSynchronizationClusterTrustedTimeSourceStruct) Endpoint() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("endpoint"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterTrustedTimeSourceStruct/endpoint
func (m_ MTRTimeSynchronizationClusterTrustedTimeSourceStruct) SetEndpoint(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndpoint:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterTrustedTimeSourceStruct/fabricIndex
func (m_ MTRTimeSynchronizationClusterTrustedTimeSourceStruct) FabricIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterTrustedTimeSourceStruct/fabricIndex
func (m_ MTRTimeSynchronizationClusterTrustedTimeSourceStruct) SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterTrustedTimeSourceStruct/nodeID
func (m_ MTRTimeSynchronizationClusterTrustedTimeSourceStruct) NodeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nodeID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterTrustedTimeSourceStruct/nodeID
func (m_ MTRTimeSynchronizationClusterTrustedTimeSourceStruct) SetNodeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNodeID:"), value)
}



