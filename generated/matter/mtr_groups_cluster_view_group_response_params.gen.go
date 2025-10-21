// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRGroupsClusterViewGroupResponseParams] class.
var (
	MTRGroupsClusterViewGroupResponseParamsClass     _MTRGroupsClusterViewGroupResponseParamsClass
	MTRGroupsClusterViewGroupResponseParamsClassOnce sync.Once
)

func getMTRGroupsClusterViewGroupResponseParamsClass() _MTRGroupsClusterViewGroupResponseParamsClass {
	MTRGroupsClusterViewGroupResponseParamsClassOnce.Do(func() {
		MTRGroupsClusterViewGroupResponseParamsClass = _MTRGroupsClusterViewGroupResponseParamsClass{objc.GetClass("MTRGroupsClusterViewGroupResponseParams")}
	})
	return MTRGroupsClusterViewGroupResponseParamsClass
}

type _MTRGroupsClusterViewGroupResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRGroupsClusterViewGroupResponseParams] class.
type IMTRGroupsClusterViewGroupResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGroupsClusterViewGroupResponseParams
type MTRGroupsClusterViewGroupResponseParams struct {
	objectivec.Object
}

// MTRGroupsClusterViewGroupResponseParamsFrom constructs a [MTRGroupsClusterViewGroupResponseParams] from an unsafe.Pointer.
func MTRGroupsClusterViewGroupResponseParamsFrom(ptr unsafe.Pointer) MTRGroupsClusterViewGroupResponseParams {
	return MTRGroupsClusterViewGroupResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGroupsClusterViewGroupResponseParamsClass) Alloc() MTRGroupsClusterViewGroupResponseParams {
	rv := objc.Send[MTRGroupsClusterViewGroupResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGroupsClusterViewGroupResponseParamsClass) New() MTRGroupsClusterViewGroupResponseParams {
	rv := objc.Send[MTRGroupsClusterViewGroupResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGroupsClusterViewGroupResponseParams) Init() MTRGroupsClusterViewGroupResponseParams {
	rv := objc.Send[MTRGroupsClusterViewGroupResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGroupsClusterViewGroupResponseParams) Autorelease() MTRGroupsClusterViewGroupResponseParams {
	rv := objc.Send[MTRGroupsClusterViewGroupResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGroupsClusterViewGroupResponseParams creates a new MTRGroupsClusterViewGroupResponseParams instance.
func NewMTRGroupsClusterViewGroupResponseParams() MTRGroupsClusterViewGroupResponseParams {
	return getMTRGroupsClusterViewGroupResponseParamsClass().New()
}




