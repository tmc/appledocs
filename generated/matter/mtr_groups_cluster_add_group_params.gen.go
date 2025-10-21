// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusteraddgroupparams/groupid-9do9s
func (m_ MTRGroupsClusterAddGroupParams) GroupID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("groupID"))
	return rv
}


// SetGroupID sets the value of the groupID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusteraddgroupparams/groupid-9do9s
func (m_ MTRGroupsClusterAddGroupParams) SetGroupID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusteraddgroupparams/groupid-9do8w
func (m_ MTRGroupsClusterAddGroupParams) GroupId() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("groupId"))
	return rv
}


// SetGroupId sets the value of the groupId property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusteraddgroupparams/groupid-9do8w
func (m_ MTRGroupsClusterAddGroupParams) SetGroupId(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupId:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusteraddgroupparams/groupname
func (m_ MTRGroupsClusterAddGroupParams) GroupName() string {
	rv := objc.Send[string](m_.ID, objc.Sel("groupName"))
	return rv
}


// SetGroupName sets the value of the groupName property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusteraddgroupparams/groupname
func (m_ MTRGroupsClusterAddGroupParams) SetGroupName(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupName:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusteraddgroupparams/serversideprocessingtimeout
func (m_ MTRGroupsClusterAddGroupParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusteraddgroupparams/serversideprocessingtimeout
func (m_ MTRGroupsClusterAddGroupParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusteraddgroupparams/timedinvoketimeoutms
func (m_ MTRGroupsClusterAddGroupParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusteraddgroupparams/timedinvoketimeoutms
func (m_ MTRGroupsClusterAddGroupParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



