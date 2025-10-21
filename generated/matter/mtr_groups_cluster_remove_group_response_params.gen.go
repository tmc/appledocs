// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterremovegroupresponseparams/groupid-5taii
func (m_ MTRGroupsClusterRemoveGroupResponseParams) GroupID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("groupID"))
	return rv
}


// SetGroupID sets the value of the groupID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterremovegroupresponseparams/groupid-5taii
func (m_ MTRGroupsClusterRemoveGroupResponseParams) SetGroupID(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterremovegroupresponseparams/groupid-5tahm
func (m_ MTRGroupsClusterRemoveGroupResponseParams) GroupId() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("groupId"))
	return rv
}


// SetGroupId sets the value of the groupId property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterremovegroupresponseparams/groupid-5tahm
func (m_ MTRGroupsClusterRemoveGroupResponseParams) SetGroupId(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupId:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterremovegroupresponseparams/status
func (m_ MTRGroupsClusterRemoveGroupResponseParams) Status() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("status"))
	return rv
}


// SetStatus sets the value of the status property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterremovegroupresponseparams/status
func (m_ MTRGroupsClusterRemoveGroupResponseParams) SetStatus(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterremovegroupresponseparams/timedinvoketimeoutms
func (m_ MTRGroupsClusterRemoveGroupResponseParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterremovegroupresponseparams/timedinvoketimeoutms
func (m_ MTRGroupsClusterRemoveGroupResponseParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



