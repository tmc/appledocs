// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRGroupsClusterRemoveGroupParams] class.
var (
	MTRGroupsClusterRemoveGroupParamsClass     _MTRGroupsClusterRemoveGroupParamsClass
	MTRGroupsClusterRemoveGroupParamsClassOnce sync.Once
)

func getMTRGroupsClusterRemoveGroupParamsClass() _MTRGroupsClusterRemoveGroupParamsClass {
	MTRGroupsClusterRemoveGroupParamsClassOnce.Do(func() {
		MTRGroupsClusterRemoveGroupParamsClass = _MTRGroupsClusterRemoveGroupParamsClass{objc.GetClass("MTRGroupsClusterRemoveGroupParams")}
	})
	return MTRGroupsClusterRemoveGroupParamsClass
}

type _MTRGroupsClusterRemoveGroupParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRGroupsClusterRemoveGroupParams] class.
type IMTRGroupsClusterRemoveGroupParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGroupsClusterRemoveGroupParams
type MTRGroupsClusterRemoveGroupParams struct {
	objectivec.Object
}

// MTRGroupsClusterRemoveGroupParamsFrom constructs a [MTRGroupsClusterRemoveGroupParams] from an unsafe.Pointer.
func MTRGroupsClusterRemoveGroupParamsFrom(ptr unsafe.Pointer) MTRGroupsClusterRemoveGroupParams {
	return MTRGroupsClusterRemoveGroupParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGroupsClusterRemoveGroupParamsClass) Alloc() MTRGroupsClusterRemoveGroupParams {
	rv := objc.Send[MTRGroupsClusterRemoveGroupParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGroupsClusterRemoveGroupParamsClass) New() MTRGroupsClusterRemoveGroupParams {
	rv := objc.Send[MTRGroupsClusterRemoveGroupParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGroupsClusterRemoveGroupParams) Init() MTRGroupsClusterRemoveGroupParams {
	rv := objc.Send[MTRGroupsClusterRemoveGroupParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGroupsClusterRemoveGroupParams) Autorelease() MTRGroupsClusterRemoveGroupParams {
	rv := objc.Send[MTRGroupsClusterRemoveGroupParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGroupsClusterRemoveGroupParams creates a new MTRGroupsClusterRemoveGroupParams instance.
func NewMTRGroupsClusterRemoveGroupParams() MTRGroupsClusterRemoveGroupParams {
	return getMTRGroupsClusterRemoveGroupParamsClass().New()
}




