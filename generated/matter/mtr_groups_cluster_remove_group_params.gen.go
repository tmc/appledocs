// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	GroupID() objc.IObject /* cross-framework: NSNumber */
	SetGroupID(value objc.IObject /* cross-framework: NSNumber */)
	GroupId() objc.IObject /* cross-framework: NSNumber */
	SetGroupId(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}

// [Full Topic]
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

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterremovegroupparams/groupid-44m5l
func (m_ MTRGroupsClusterRemoveGroupParams) GroupID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("groupID"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterremovegroupparams/groupid-44m5l
func (m_ MTRGroupsClusterRemoveGroupParams) SetGroupID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupID:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterremovegroupparams/groupid-44m4p
func (m_ MTRGroupsClusterRemoveGroupParams) GroupId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("groupId"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterremovegroupparams/groupid-44m4p
func (m_ MTRGroupsClusterRemoveGroupParams) SetGroupId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupId:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterremovegroupparams/serversideprocessingtimeout
func (m_ MTRGroupsClusterRemoveGroupParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterremovegroupparams/serversideprocessingtimeout
func (m_ MTRGroupsClusterRemoveGroupParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterremovegroupparams/timedinvoketimeoutms
func (m_ MTRGroupsClusterRemoveGroupParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterremovegroupparams/timedinvoketimeoutms
func (m_ MTRGroupsClusterRemoveGroupParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}
