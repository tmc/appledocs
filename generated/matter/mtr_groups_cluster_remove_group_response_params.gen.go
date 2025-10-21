// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRGroupsClusterRemoveGroupResponseParams] class.
var (
	MTRGroupsClusterRemoveGroupResponseParamsClass     _MTRGroupsClusterRemoveGroupResponseParamsClass
	MTRGroupsClusterRemoveGroupResponseParamsClassOnce sync.Once
)

func getMTRGroupsClusterRemoveGroupResponseParamsClass() _MTRGroupsClusterRemoveGroupResponseParamsClass {
	MTRGroupsClusterRemoveGroupResponseParamsClassOnce.Do(func() {
		MTRGroupsClusterRemoveGroupResponseParamsClass = _MTRGroupsClusterRemoveGroupResponseParamsClass{objc.GetClass("MTRGroupsClusterRemoveGroupResponseParams")}
	})
	return MTRGroupsClusterRemoveGroupResponseParamsClass
}

type _MTRGroupsClusterRemoveGroupResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRGroupsClusterRemoveGroupResponseParams] class.
type IMTRGroupsClusterRemoveGroupResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGroupsClusterRemoveGroupResponseParams
type MTRGroupsClusterRemoveGroupResponseParams struct {
	objectivec.Object
}

// MTRGroupsClusterRemoveGroupResponseParamsFrom constructs a [MTRGroupsClusterRemoveGroupResponseParams] from an unsafe.Pointer.
func MTRGroupsClusterRemoveGroupResponseParamsFrom(ptr unsafe.Pointer) MTRGroupsClusterRemoveGroupResponseParams {
	return MTRGroupsClusterRemoveGroupResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGroupsClusterRemoveGroupResponseParamsClass) Alloc() MTRGroupsClusterRemoveGroupResponseParams {
	rv := objc.Send[MTRGroupsClusterRemoveGroupResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGroupsClusterRemoveGroupResponseParamsClass) New() MTRGroupsClusterRemoveGroupResponseParams {
	rv := objc.Send[MTRGroupsClusterRemoveGroupResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGroupsClusterRemoveGroupResponseParams) Init() MTRGroupsClusterRemoveGroupResponseParams {
	rv := objc.Send[MTRGroupsClusterRemoveGroupResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGroupsClusterRemoveGroupResponseParams) Autorelease() MTRGroupsClusterRemoveGroupResponseParams {
	rv := objc.Send[MTRGroupsClusterRemoveGroupResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGroupsClusterRemoveGroupResponseParams creates a new MTRGroupsClusterRemoveGroupResponseParams instance.
func NewMTRGroupsClusterRemoveGroupResponseParams() MTRGroupsClusterRemoveGroupResponseParams {
	return getMTRGroupsClusterRemoveGroupResponseParamsClass().New()
}




