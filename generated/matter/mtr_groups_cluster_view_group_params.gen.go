// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
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

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterviewgroupparams/groupid-4nbva
func (m_ MTRGroupsClusterViewGroupParams) GroupID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("groupID"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterviewgroupparams/groupid-4nbva
func (m_ MTRGroupsClusterViewGroupParams) SetGroupID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupID:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterviewgroupparams/groupid-4nbw6
func (m_ MTRGroupsClusterViewGroupParams) GroupId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("groupId"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterviewgroupparams/groupid-4nbw6
func (m_ MTRGroupsClusterViewGroupParams) SetGroupId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupId:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterviewgroupparams/serversideprocessingtimeout
func (m_ MTRGroupsClusterViewGroupParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterviewgroupparams/serversideprocessingtimeout
func (m_ MTRGroupsClusterViewGroupParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterviewgroupparams/timedinvoketimeoutms
func (m_ MTRGroupsClusterViewGroupParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterviewgroupparams/timedinvoketimeoutms
func (m_ MTRGroupsClusterViewGroupParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}
