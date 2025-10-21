// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterviewgroupparams/groupid-4nbva
func (m_ MTRGroupsClusterViewGroupParams) GroupID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("groupID"))
	return rv
}


// SetGroupID sets the value of the groupID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterviewgroupparams/groupid-4nbva
func (m_ MTRGroupsClusterViewGroupParams) SetGroupID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterviewgroupparams/groupid-4nbw6
func (m_ MTRGroupsClusterViewGroupParams) GroupId() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("groupId"))
	return rv
}


// SetGroupId sets the value of the groupId property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterviewgroupparams/groupid-4nbw6
func (m_ MTRGroupsClusterViewGroupParams) SetGroupId(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupId:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterviewgroupparams/serversideprocessingtimeout
func (m_ MTRGroupsClusterViewGroupParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterviewgroupparams/serversideprocessingtimeout
func (m_ MTRGroupsClusterViewGroupParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterviewgroupparams/timedinvoketimeoutms
func (m_ MTRGroupsClusterViewGroupParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterviewgroupparams/timedinvoketimeoutms
func (m_ MTRGroupsClusterViewGroupParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



