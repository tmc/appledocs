// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRGroupsClusterViewGroupResponseParams] class.
var (
	MTRGroupsClusterViewGroupResponseParamsClass     _MTRGroupsClusterViewGroupResponseParamsClass
	MTRGroupsClusterViewGroupResponseParamsClassOnce sync.Once
)

func getMTRGroupsClusterViewGroupResponseParamsClass() _MTRGroupsClusterViewGroupResponseParamsClass {
	MTRGroupsClusterViewGroupResponseParamsClassOnce.Do(func() {
		MTRGroupsClusterViewGroupResponseParamsClass = _MTRGroupsClusterViewGroupResponseParamsClass{objc.GetClass("MTRGroupsClusterViewGroupResponseParams")}
	})
	return MTRGroupsClusterViewGroupResponseParamsClass
}

type _MTRGroupsClusterViewGroupResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRGroupsClusterViewGroupResponseParams] class.
type IMTRGroupsClusterViewGroupResponseParams interface {
	objectivec.IObject
	// properties:
	GroupID() objc.IObject /* cross-framework: NSNumber */
	SetGroupID(value objc.IObject /* cross-framework: NSNumber */)
	GroupId() objc.IObject /* cross-framework: NSNumber */
	SetGroupId(value objc.IObject /* cross-framework: NSNumber */)
	GroupName() objc.IObject /* cross-framework: NSString */
	SetGroupName(value objc.IObject /* cross-framework: NSString */)
	Status() objc.IObject /* cross-framework: NSNumber */
	SetStatus(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGroupsClusterViewGroupResponseParams
type MTRGroupsClusterViewGroupResponseParams struct {
	objectivec.Object
}

// MTRGroupsClusterViewGroupResponseParamsFrom constructs a [MTRGroupsClusterViewGroupResponseParams] from an unsafe.Pointer.
func MTRGroupsClusterViewGroupResponseParamsFrom(ptr unsafe.Pointer) MTRGroupsClusterViewGroupResponseParams {
	return MTRGroupsClusterViewGroupResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGroupsClusterViewGroupResponseParamsClass) Alloc() MTRGroupsClusterViewGroupResponseParams {
	rv := objc.Send[MTRGroupsClusterViewGroupResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGroupsClusterViewGroupResponseParamsClass) New() MTRGroupsClusterViewGroupResponseParams {
	rv := objc.Send[MTRGroupsClusterViewGroupResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGroupsClusterViewGroupResponseParams) Init() MTRGroupsClusterViewGroupResponseParams {
	rv := objc.Send[MTRGroupsClusterViewGroupResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGroupsClusterViewGroupResponseParams) Autorelease() MTRGroupsClusterViewGroupResponseParams {
	rv := objc.Send[MTRGroupsClusterViewGroupResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGroupsClusterViewGroupResponseParams creates a new MTRGroupsClusterViewGroupResponseParams instance.
func NewMTRGroupsClusterViewGroupResponseParams() MTRGroupsClusterViewGroupResponseParams {
	return getMTRGroupsClusterViewGroupResponseParamsClass().New()
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterviewgroupresponseparams/groupid-7jt17
func (m_ MTRGroupsClusterViewGroupResponseParams) GroupID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("groupID"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterviewgroupresponseparams/groupid-7jt17
func (m_ MTRGroupsClusterViewGroupResponseParams) SetGroupID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupID:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterviewgroupresponseparams/groupid-7jt0b
func (m_ MTRGroupsClusterViewGroupResponseParams) GroupId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("groupId"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterviewgroupresponseparams/groupid-7jt0b
func (m_ MTRGroupsClusterViewGroupResponseParams) SetGroupId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupId:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterviewgroupresponseparams/groupname
func (m_ MTRGroupsClusterViewGroupResponseParams) GroupName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("groupName"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterviewgroupresponseparams/groupname
func (m_ MTRGroupsClusterViewGroupResponseParams) SetGroupName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupName:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterviewgroupresponseparams/status
func (m_ MTRGroupsClusterViewGroupResponseParams) Status() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("status"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterviewgroupresponseparams/status
func (m_ MTRGroupsClusterViewGroupResponseParams) SetStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterviewgroupresponseparams/timedinvoketimeoutms
func (m_ MTRGroupsClusterViewGroupResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterviewgroupresponseparams/timedinvoketimeoutms
func (m_ MTRGroupsClusterViewGroupResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}
