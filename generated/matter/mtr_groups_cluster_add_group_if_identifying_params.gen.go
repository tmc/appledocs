// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRGroupsClusterAddGroupIfIdentifyingParams] class.
var (
	MTRGroupsClusterAddGroupIfIdentifyingParamsClass     _MTRGroupsClusterAddGroupIfIdentifyingParamsClass
	MTRGroupsClusterAddGroupIfIdentifyingParamsClassOnce sync.Once
)

func getMTRGroupsClusterAddGroupIfIdentifyingParamsClass() _MTRGroupsClusterAddGroupIfIdentifyingParamsClass {
	MTRGroupsClusterAddGroupIfIdentifyingParamsClassOnce.Do(func() {
		MTRGroupsClusterAddGroupIfIdentifyingParamsClass = _MTRGroupsClusterAddGroupIfIdentifyingParamsClass{objc.GetClass("MTRGroupsClusterAddGroupIfIdentifyingParams")}
	})
	return MTRGroupsClusterAddGroupIfIdentifyingParamsClass
}

type _MTRGroupsClusterAddGroupIfIdentifyingParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRGroupsClusterAddGroupIfIdentifyingParams] class.
type IMTRGroupsClusterAddGroupIfIdentifyingParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGroupsClusterAddGroupIfIdentifyingParams
type MTRGroupsClusterAddGroupIfIdentifyingParams struct {
	objectivec.Object
}

// MTRGroupsClusterAddGroupIfIdentifyingParamsFrom constructs a [MTRGroupsClusterAddGroupIfIdentifyingParams] from an unsafe.Pointer.
func MTRGroupsClusterAddGroupIfIdentifyingParamsFrom(ptr unsafe.Pointer) MTRGroupsClusterAddGroupIfIdentifyingParams {
	return MTRGroupsClusterAddGroupIfIdentifyingParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGroupsClusterAddGroupIfIdentifyingParamsClass) Alloc() MTRGroupsClusterAddGroupIfIdentifyingParams {
	rv := objc.Send[MTRGroupsClusterAddGroupIfIdentifyingParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGroupsClusterAddGroupIfIdentifyingParamsClass) New() MTRGroupsClusterAddGroupIfIdentifyingParams {
	rv := objc.Send[MTRGroupsClusterAddGroupIfIdentifyingParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGroupsClusterAddGroupIfIdentifyingParams) Init() MTRGroupsClusterAddGroupIfIdentifyingParams {
	rv := objc.Send[MTRGroupsClusterAddGroupIfIdentifyingParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGroupsClusterAddGroupIfIdentifyingParams) Autorelease() MTRGroupsClusterAddGroupIfIdentifyingParams {
	rv := objc.Send[MTRGroupsClusterAddGroupIfIdentifyingParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGroupsClusterAddGroupIfIdentifyingParams creates a new MTRGroupsClusterAddGroupIfIdentifyingParams instance.
func NewMTRGroupsClusterAddGroupIfIdentifyingParams() MTRGroupsClusterAddGroupIfIdentifyingParams {
	return getMTRGroupsClusterAddGroupIfIdentifyingParamsClass().New()
}




