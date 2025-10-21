// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRGroupsClusterViewGroupParams] class.
var (
	MTRGroupsClusterViewGroupParamsClass     _MTRGroupsClusterViewGroupParamsClass
	MTRGroupsClusterViewGroupParamsClassOnce sync.Once
)

func getMTRGroupsClusterViewGroupParamsClass() _MTRGroupsClusterViewGroupParamsClass {
	MTRGroupsClusterViewGroupParamsClassOnce.Do(func() {
		MTRGroupsClusterViewGroupParamsClass = _MTRGroupsClusterViewGroupParamsClass{objc.GetClass("MTRGroupsClusterViewGroupParams")}
	})
	return MTRGroupsClusterViewGroupParamsClass
}

type _MTRGroupsClusterViewGroupParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRGroupsClusterViewGroupParams] class.
type IMTRGroupsClusterViewGroupParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGroupsClusterViewGroupParams
type MTRGroupsClusterViewGroupParams struct {
	objectivec.Object
}

// MTRGroupsClusterViewGroupParamsFrom constructs a [MTRGroupsClusterViewGroupParams] from an unsafe.Pointer.
func MTRGroupsClusterViewGroupParamsFrom(ptr unsafe.Pointer) MTRGroupsClusterViewGroupParams {
	return MTRGroupsClusterViewGroupParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGroupsClusterViewGroupParamsClass) Alloc() MTRGroupsClusterViewGroupParams {
	rv := objc.Send[MTRGroupsClusterViewGroupParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGroupsClusterViewGroupParamsClass) New() MTRGroupsClusterViewGroupParams {
	rv := objc.Send[MTRGroupsClusterViewGroupParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGroupsClusterViewGroupParams) Init() MTRGroupsClusterViewGroupParams {
	rv := objc.Send[MTRGroupsClusterViewGroupParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGroupsClusterViewGroupParams) Autorelease() MTRGroupsClusterViewGroupParams {
	rv := objc.Send[MTRGroupsClusterViewGroupParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGroupsClusterViewGroupParams creates a new MTRGroupsClusterViewGroupParams instance.
func NewMTRGroupsClusterViewGroupParams() MTRGroupsClusterViewGroupParams {
	return getMTRGroupsClusterViewGroupParamsClass().New()
}




