// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRGroupsClusterRemoveAllGroupsParams] class.
var (
	MTRGroupsClusterRemoveAllGroupsParamsClass     _MTRGroupsClusterRemoveAllGroupsParamsClass
	MTRGroupsClusterRemoveAllGroupsParamsClassOnce sync.Once
)

func getMTRGroupsClusterRemoveAllGroupsParamsClass() _MTRGroupsClusterRemoveAllGroupsParamsClass {
	MTRGroupsClusterRemoveAllGroupsParamsClassOnce.Do(func() {
		MTRGroupsClusterRemoveAllGroupsParamsClass = _MTRGroupsClusterRemoveAllGroupsParamsClass{objc.GetClass("MTRGroupsClusterRemoveAllGroupsParams")}
	})
	return MTRGroupsClusterRemoveAllGroupsParamsClass
}

type _MTRGroupsClusterRemoveAllGroupsParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRGroupsClusterRemoveAllGroupsParams] class.
type IMTRGroupsClusterRemoveAllGroupsParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGroupsClusterRemoveAllGroupsParams
type MTRGroupsClusterRemoveAllGroupsParams struct {
	objectivec.Object
}

// MTRGroupsClusterRemoveAllGroupsParamsFrom constructs a [MTRGroupsClusterRemoveAllGroupsParams] from an unsafe.Pointer.
func MTRGroupsClusterRemoveAllGroupsParamsFrom(ptr unsafe.Pointer) MTRGroupsClusterRemoveAllGroupsParams {
	return MTRGroupsClusterRemoveAllGroupsParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGroupsClusterRemoveAllGroupsParamsClass) Alloc() MTRGroupsClusterRemoveAllGroupsParams {
	rv := objc.Send[MTRGroupsClusterRemoveAllGroupsParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGroupsClusterRemoveAllGroupsParamsClass) New() MTRGroupsClusterRemoveAllGroupsParams {
	rv := objc.Send[MTRGroupsClusterRemoveAllGroupsParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGroupsClusterRemoveAllGroupsParams) Init() MTRGroupsClusterRemoveAllGroupsParams {
	rv := objc.Send[MTRGroupsClusterRemoveAllGroupsParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGroupsClusterRemoveAllGroupsParams) Autorelease() MTRGroupsClusterRemoveAllGroupsParams {
	rv := objc.Send[MTRGroupsClusterRemoveAllGroupsParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGroupsClusterRemoveAllGroupsParams creates a new MTRGroupsClusterRemoveAllGroupsParams instance.
func NewMTRGroupsClusterRemoveAllGroupsParams() MTRGroupsClusterRemoveAllGroupsParams {
	return getMTRGroupsClusterRemoveAllGroupsParamsClass().New()
}




