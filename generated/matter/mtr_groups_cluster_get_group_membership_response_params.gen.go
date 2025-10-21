// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRGroupsClusterGetGroupMembershipResponseParams] class.
var (
	MTRGroupsClusterGetGroupMembershipResponseParamsClass     _MTRGroupsClusterGetGroupMembershipResponseParamsClass
	MTRGroupsClusterGetGroupMembershipResponseParamsClassOnce sync.Once
)

func getMTRGroupsClusterGetGroupMembershipResponseParamsClass() _MTRGroupsClusterGetGroupMembershipResponseParamsClass {
	MTRGroupsClusterGetGroupMembershipResponseParamsClassOnce.Do(func() {
		MTRGroupsClusterGetGroupMembershipResponseParamsClass = _MTRGroupsClusterGetGroupMembershipResponseParamsClass{objc.GetClass("MTRGroupsClusterGetGroupMembershipResponseParams")}
	})
	return MTRGroupsClusterGetGroupMembershipResponseParamsClass
}

type _MTRGroupsClusterGetGroupMembershipResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRGroupsClusterGetGroupMembershipResponseParams] class.
type IMTRGroupsClusterGetGroupMembershipResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGroupsClusterGetGroupMembershipResponseParams
type MTRGroupsClusterGetGroupMembershipResponseParams struct {
	objectivec.Object
}

// MTRGroupsClusterGetGroupMembershipResponseParamsFrom constructs a [MTRGroupsClusterGetGroupMembershipResponseParams] from an unsafe.Pointer.
func MTRGroupsClusterGetGroupMembershipResponseParamsFrom(ptr unsafe.Pointer) MTRGroupsClusterGetGroupMembershipResponseParams {
	return MTRGroupsClusterGetGroupMembershipResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGroupsClusterGetGroupMembershipResponseParamsClass) Alloc() MTRGroupsClusterGetGroupMembershipResponseParams {
	rv := objc.Send[MTRGroupsClusterGetGroupMembershipResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGroupsClusterGetGroupMembershipResponseParamsClass) New() MTRGroupsClusterGetGroupMembershipResponseParams {
	rv := objc.Send[MTRGroupsClusterGetGroupMembershipResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGroupsClusterGetGroupMembershipResponseParams) Init() MTRGroupsClusterGetGroupMembershipResponseParams {
	rv := objc.Send[MTRGroupsClusterGetGroupMembershipResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGroupsClusterGetGroupMembershipResponseParams) Autorelease() MTRGroupsClusterGetGroupMembershipResponseParams {
	rv := objc.Send[MTRGroupsClusterGetGroupMembershipResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGroupsClusterGetGroupMembershipResponseParams creates a new MTRGroupsClusterGetGroupMembershipResponseParams instance.
func NewMTRGroupsClusterGetGroupMembershipResponseParams() MTRGroupsClusterGetGroupMembershipResponseParams {
	return getMTRGroupsClusterGetGroupMembershipResponseParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclustergetgroupmembershipresponseparams/capacity
func (m_ MTRGroupsClusterGetGroupMembershipResponseParams) Capacity() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("capacity"))
	return rv
}


// SetCapacity sets the value of the capacity property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclustergetgroupmembershipresponseparams/capacity
func (m_ MTRGroupsClusterGetGroupMembershipResponseParams) SetCapacity(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCapacity:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclustergetgroupmembershipresponseparams/grouplist
func (m_ MTRGroupsClusterGetGroupMembershipResponseParams) GroupList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("groupList"))
	return rv
}


// SetGroupList sets the value of the groupList property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclustergetgroupmembershipresponseparams/grouplist
func (m_ MTRGroupsClusterGetGroupMembershipResponseParams) SetGroupList(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupList:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclustergetgroupmembershipresponseparams/timedinvoketimeoutms
func (m_ MTRGroupsClusterGetGroupMembershipResponseParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclustergetgroupmembershipresponseparams/timedinvoketimeoutms
func (m_ MTRGroupsClusterGetGroupMembershipResponseParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



