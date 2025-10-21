// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRGroupsClusterAddGroupParams] class.
var (
	MTRGroupsClusterAddGroupParamsClass     _MTRGroupsClusterAddGroupParamsClass
	MTRGroupsClusterAddGroupParamsClassOnce sync.Once
)

func getMTRGroupsClusterAddGroupParamsClass() _MTRGroupsClusterAddGroupParamsClass {
	MTRGroupsClusterAddGroupParamsClassOnce.Do(func() {
		MTRGroupsClusterAddGroupParamsClass = _MTRGroupsClusterAddGroupParamsClass{objc.GetClass("MTRGroupsClusterAddGroupParams")}
	})
	return MTRGroupsClusterAddGroupParamsClass
}

type _MTRGroupsClusterAddGroupParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRGroupsClusterAddGroupParams] class.
type IMTRGroupsClusterAddGroupParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGroupsClusterAddGroupParams
type MTRGroupsClusterAddGroupParams struct {
	objectivec.Object
}

// MTRGroupsClusterAddGroupParamsFrom constructs a [MTRGroupsClusterAddGroupParams] from an unsafe.Pointer.
func MTRGroupsClusterAddGroupParamsFrom(ptr unsafe.Pointer) MTRGroupsClusterAddGroupParams {
	return MTRGroupsClusterAddGroupParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGroupsClusterAddGroupParamsClass) Alloc() MTRGroupsClusterAddGroupParams {
	rv := objc.Send[MTRGroupsClusterAddGroupParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGroupsClusterAddGroupParamsClass) New() MTRGroupsClusterAddGroupParams {
	rv := objc.Send[MTRGroupsClusterAddGroupParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGroupsClusterAddGroupParams) Init() MTRGroupsClusterAddGroupParams {
	rv := objc.Send[MTRGroupsClusterAddGroupParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGroupsClusterAddGroupParams) Autorelease() MTRGroupsClusterAddGroupParams {
	rv := objc.Send[MTRGroupsClusterAddGroupParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGroupsClusterAddGroupParams creates a new MTRGroupsClusterAddGroupParams instance.
func NewMTRGroupsClusterAddGroupParams() MTRGroupsClusterAddGroupParams {
	return getMTRGroupsClusterAddGroupParamsClass().New()
}




