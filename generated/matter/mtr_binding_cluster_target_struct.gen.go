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
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbindingclustertargetstruct/cluster
func (m_ MTRBindingClusterTargetStruct) Cluster() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("cluster"))
	return rv
}


// SetCluster sets the value of the cluster property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbindingclustertargetstruct/cluster
func (m_ MTRBindingClusterTargetStruct) SetCluster(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCluster:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbindingclustertargetstruct/endpoint
func (m_ MTRBindingClusterTargetStruct) Endpoint() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("endpoint"))
	return rv
}


// SetEndpoint sets the value of the endpoint property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbindingclustertargetstruct/endpoint
func (m_ MTRBindingClusterTargetStruct) SetEndpoint(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndpoint:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbindingclustertargetstruct/fabricindex
func (m_ MTRBindingClusterTargetStruct) FabricIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("fabricIndex"))
	return rv
}


// SetFabricIndex sets the value of the fabricIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbindingclustertargetstruct/fabricindex
func (m_ MTRBindingClusterTargetStruct) SetFabricIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbindingclustertargetstruct/group
func (m_ MTRBindingClusterTargetStruct) Group() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("group"))
	return rv
}


// SetGroup sets the value of the group property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbindingclustertargetstruct/group
func (m_ MTRBindingClusterTargetStruct) SetGroup(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroup:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbindingclustertargetstruct/node
func (m_ MTRBindingClusterTargetStruct) Node() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("node"))
	return rv
}


// SetNode sets the value of the node property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbindingclustertargetstruct/node
func (m_ MTRBindingClusterTargetStruct) SetNode(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNode:"), value)
}



