// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRBindingClusterTargetStruct] class.
var (
	MTRBindingClusterTargetStructClass     _MTRBindingClusterTargetStructClass
	MTRBindingClusterTargetStructClassOnce sync.Once
)

func getMTRBindingClusterTargetStructClass() _MTRBindingClusterTargetStructClass {
	MTRBindingClusterTargetStructClassOnce.Do(func() {
		MTRBindingClusterTargetStructClass = _MTRBindingClusterTargetStructClass{objc.GetClass("MTRBindingClusterTargetStruct")}
	})
	return MTRBindingClusterTargetStructClass
}

type _MTRBindingClusterTargetStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRBindingClusterTargetStruct] class.
type IMTRBindingClusterTargetStruct interface {
	objectivec.IObject
	// properties:
	Cluster() objc.IObject /* cross-framework: NSNumber */
	SetCluster(value objc.IObject /* cross-framework: NSNumber */)
	Endpoint() objc.IObject /* cross-framework: NSNumber */
	SetEndpoint(value objc.IObject /* cross-framework: NSNumber */)
	FabricIndex() objc.IObject /* cross-framework: NSNumber */
	SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */)
	Group() objc.IObject /* cross-framework: NSNumber */
	SetGroup(value objc.IObject /* cross-framework: NSNumber */)
	Node() objc.IObject /* cross-framework: NSNumber */
	SetNode(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBindingClusterTargetStruct
type MTRBindingClusterTargetStruct struct {
	objectivec.Object
}

// MTRBindingClusterTargetStructFrom constructs a [MTRBindingClusterTargetStruct] from an unsafe.Pointer.
func MTRBindingClusterTargetStructFrom(ptr unsafe.Pointer) MTRBindingClusterTargetStruct {
	return MTRBindingClusterTargetStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBindingClusterTargetStructClass) Alloc() MTRBindingClusterTargetStruct {
	rv := objc.Send[MTRBindingClusterTargetStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBindingClusterTargetStructClass) New() MTRBindingClusterTargetStruct {
	rv := objc.Send[MTRBindingClusterTargetStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBindingClusterTargetStruct) Init() MTRBindingClusterTargetStruct {
	rv := objc.Send[MTRBindingClusterTargetStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBindingClusterTargetStruct) Autorelease() MTRBindingClusterTargetStruct {
	rv := objc.Send[MTRBindingClusterTargetStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBindingClusterTargetStruct creates a new MTRBindingClusterTargetStruct instance.
func NewMTRBindingClusterTargetStruct() MTRBindingClusterTargetStruct {
	return getMTRBindingClusterTargetStructClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbindingclustertargetstruct/cluster
func (m_ MTRBindingClusterTargetStruct) Cluster() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("cluster"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbindingclustertargetstruct/cluster
func (m_ MTRBindingClusterTargetStruct) SetCluster(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCluster:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbindingclustertargetstruct/endpoint
func (m_ MTRBindingClusterTargetStruct) Endpoint() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("endpoint"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbindingclustertargetstruct/endpoint
func (m_ MTRBindingClusterTargetStruct) SetEndpoint(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndpoint:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbindingclustertargetstruct/fabricindex
func (m_ MTRBindingClusterTargetStruct) FabricIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbindingclustertargetstruct/fabricindex
func (m_ MTRBindingClusterTargetStruct) SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbindingclustertargetstruct/group
func (m_ MTRBindingClusterTargetStruct) Group() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("group"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbindingclustertargetstruct/group
func (m_ MTRBindingClusterTargetStruct) SetGroup(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroup:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbindingclustertargetstruct/node
func (m_ MTRBindingClusterTargetStruct) Node() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("node"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbindingclustertargetstruct/node
func (m_ MTRBindingClusterTargetStruct) SetNode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNode:"), value)
}



