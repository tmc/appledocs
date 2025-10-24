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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusteraddgroupparams/groupid-9do9s
func (m_ MTRGroupsClusterAddGroupParams) GroupID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("groupID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusteraddgroupparams/groupid-9do9s
func (m_ MTRGroupsClusterAddGroupParams) SetGroupID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusteraddgroupparams/groupid-9do8w
func (m_ MTRGroupsClusterAddGroupParams) GroupId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("groupId"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusteraddgroupparams/groupid-9do8w
func (m_ MTRGroupsClusterAddGroupParams) SetGroupId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupId:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusteraddgroupparams/groupname
func (m_ MTRGroupsClusterAddGroupParams) GroupName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("groupName"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusteraddgroupparams/groupname
func (m_ MTRGroupsClusterAddGroupParams) SetGroupName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupName:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusteraddgroupparams/serversideprocessingtimeout
func (m_ MTRGroupsClusterAddGroupParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusteraddgroupparams/serversideprocessingtimeout
func (m_ MTRGroupsClusterAddGroupParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusteraddgroupparams/timedinvoketimeoutms
func (m_ MTRGroupsClusterAddGroupParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusteraddgroupparams/timedinvoketimeoutms
func (m_ MTRGroupsClusterAddGroupParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



