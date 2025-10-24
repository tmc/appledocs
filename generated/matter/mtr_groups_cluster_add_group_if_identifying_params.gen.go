// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRGroupsClusterAddGroupIfIdentifyingParams] class.
var (
	MTRGroupsClusterAddGroupIfIdentifyingParamsClass     _MTRGroupsClusterAddGroupIfIdentifyingParamsClass
	MTRGroupsClusterAddGroupIfIdentifyingParamsClassOnce sync.Once
)

func getMTRGroupsClusterAddGroupIfIdentifyingParamsClass() _MTRGroupsClusterAddGroupIfIdentifyingParamsClass {
	MTRGroupsClusterAddGroupIfIdentifyingParamsClassOnce.Do(func() {
		MTRGroupsClusterAddGroupIfIdentifyingParamsClass = _MTRGroupsClusterAddGroupIfIdentifyingParamsClass{objc.GetClass("MTRGroupsClusterAddGroupIfIdentifyingParams")}
	})
	return MTRGroupsClusterAddGroupIfIdentifyingParamsClass
}

type _MTRGroupsClusterAddGroupIfIdentifyingParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRGroupsClusterAddGroupIfIdentifyingParams] class.
type IMTRGroupsClusterAddGroupIfIdentifyingParams interface {
	objectivec.IObject
	// properties:
	GroupID() objc.IObject /* cross-framework: NSNumber */
	SetGroupID(value objc.IObject /* cross-framework: NSNumber */)
	GroupId() objc.IObject /* cross-framework: NSNumber */
	SetGroupId(value objc.IObject /* cross-framework: NSNumber */)
	GroupName() objc.IObject /* cross-framework: NSString */
	SetGroupName(value objc.IObject /* cross-framework: NSString */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGroupsClusterAddGroupIfIdentifyingParams
type MTRGroupsClusterAddGroupIfIdentifyingParams struct {
	objectivec.Object
}

// MTRGroupsClusterAddGroupIfIdentifyingParamsFrom constructs a [MTRGroupsClusterAddGroupIfIdentifyingParams] from an unsafe.Pointer.
func MTRGroupsClusterAddGroupIfIdentifyingParamsFrom(ptr unsafe.Pointer) MTRGroupsClusterAddGroupIfIdentifyingParams {
	return MTRGroupsClusterAddGroupIfIdentifyingParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGroupsClusterAddGroupIfIdentifyingParamsClass) Alloc() MTRGroupsClusterAddGroupIfIdentifyingParams {
	rv := objc.Send[MTRGroupsClusterAddGroupIfIdentifyingParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGroupsClusterAddGroupIfIdentifyingParamsClass) New() MTRGroupsClusterAddGroupIfIdentifyingParams {
	rv := objc.Send[MTRGroupsClusterAddGroupIfIdentifyingParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGroupsClusterAddGroupIfIdentifyingParams) Init() MTRGroupsClusterAddGroupIfIdentifyingParams {
	rv := objc.Send[MTRGroupsClusterAddGroupIfIdentifyingParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGroupsClusterAddGroupIfIdentifyingParams) Autorelease() MTRGroupsClusterAddGroupIfIdentifyingParams {
	rv := objc.Send[MTRGroupsClusterAddGroupIfIdentifyingParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGroupsClusterAddGroupIfIdentifyingParams creates a new MTRGroupsClusterAddGroupIfIdentifyingParams instance.
func NewMTRGroupsClusterAddGroupIfIdentifyingParams() MTRGroupsClusterAddGroupIfIdentifyingParams {
	return getMTRGroupsClusterAddGroupIfIdentifyingParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusteraddgroupifidentifyingparams/groupid-1fuio
func (m_ MTRGroupsClusterAddGroupIfIdentifyingParams) GroupID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("groupID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusteraddgroupifidentifyingparams/groupid-1fuio
func (m_ MTRGroupsClusterAddGroupIfIdentifyingParams) SetGroupID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusteraddgroupifidentifyingparams/groupid-1fujk
func (m_ MTRGroupsClusterAddGroupIfIdentifyingParams) GroupId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("groupId"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusteraddgroupifidentifyingparams/groupid-1fujk
func (m_ MTRGroupsClusterAddGroupIfIdentifyingParams) SetGroupId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupId:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusteraddgroupifidentifyingparams/groupname
func (m_ MTRGroupsClusterAddGroupIfIdentifyingParams) GroupName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("groupName"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusteraddgroupifidentifyingparams/groupname
func (m_ MTRGroupsClusterAddGroupIfIdentifyingParams) SetGroupName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupName:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusteraddgroupifidentifyingparams/serversideprocessingtimeout
func (m_ MTRGroupsClusterAddGroupIfIdentifyingParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusteraddgroupifidentifyingparams/serversideprocessingtimeout
func (m_ MTRGroupsClusterAddGroupIfIdentifyingParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusteraddgroupifidentifyingparams/timedinvoketimeoutms
func (m_ MTRGroupsClusterAddGroupIfIdentifyingParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusteraddgroupifidentifyingparams/timedinvoketimeoutms
func (m_ MTRGroupsClusterAddGroupIfIdentifyingParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



