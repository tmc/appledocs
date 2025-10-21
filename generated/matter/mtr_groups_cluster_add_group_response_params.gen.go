// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRGroupsClusterAddGroupResponseParams] class.
var (
	MTRGroupsClusterAddGroupResponseParamsClass     _MTRGroupsClusterAddGroupResponseParamsClass
	MTRGroupsClusterAddGroupResponseParamsClassOnce sync.Once
)

func getMTRGroupsClusterAddGroupResponseParamsClass() _MTRGroupsClusterAddGroupResponseParamsClass {
	MTRGroupsClusterAddGroupResponseParamsClassOnce.Do(func() {
		MTRGroupsClusterAddGroupResponseParamsClass = _MTRGroupsClusterAddGroupResponseParamsClass{objc.GetClass("MTRGroupsClusterAddGroupResponseParams")}
	})
	return MTRGroupsClusterAddGroupResponseParamsClass
}

type _MTRGroupsClusterAddGroupResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRGroupsClusterAddGroupResponseParams] class.
type IMTRGroupsClusterAddGroupResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGroupsClusterAddGroupResponseParams
type MTRGroupsClusterAddGroupResponseParams struct {
	objectivec.Object
}

// MTRGroupsClusterAddGroupResponseParamsFrom constructs a [MTRGroupsClusterAddGroupResponseParams] from an unsafe.Pointer.
func MTRGroupsClusterAddGroupResponseParamsFrom(ptr unsafe.Pointer) MTRGroupsClusterAddGroupResponseParams {
	return MTRGroupsClusterAddGroupResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGroupsClusterAddGroupResponseParamsClass) Alloc() MTRGroupsClusterAddGroupResponseParams {
	rv := objc.Send[MTRGroupsClusterAddGroupResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGroupsClusterAddGroupResponseParamsClass) New() MTRGroupsClusterAddGroupResponseParams {
	rv := objc.Send[MTRGroupsClusterAddGroupResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGroupsClusterAddGroupResponseParams) Init() MTRGroupsClusterAddGroupResponseParams {
	rv := objc.Send[MTRGroupsClusterAddGroupResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGroupsClusterAddGroupResponseParams) Autorelease() MTRGroupsClusterAddGroupResponseParams {
	rv := objc.Send[MTRGroupsClusterAddGroupResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGroupsClusterAddGroupResponseParams creates a new MTRGroupsClusterAddGroupResponseParams instance.
func NewMTRGroupsClusterAddGroupResponseParams() MTRGroupsClusterAddGroupResponseParams {
	return getMTRGroupsClusterAddGroupResponseParamsClass().New()
}




