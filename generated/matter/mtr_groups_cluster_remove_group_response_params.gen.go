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
	// properties:
	GroupID() objc.IObject /* cross-framework: NSNumber */
	SetGroupID(value objc.IObject /* cross-framework: NSNumber */)
	GroupId() objc.IObject /* cross-framework: NSNumber */
	SetGroupId(value objc.IObject /* cross-framework: NSNumber */)
	Status() objc.IObject /* cross-framework: NSNumber */
	SetStatus(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterremovegroupresponseparams/groupid-5taii
func (m_ MTRGroupsClusterRemoveGroupResponseParams) GroupID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("groupID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterremovegroupresponseparams/groupid-5taii
func (m_ MTRGroupsClusterRemoveGroupResponseParams) SetGroupID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterremovegroupresponseparams/groupid-5tahm
func (m_ MTRGroupsClusterRemoveGroupResponseParams) GroupId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("groupId"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterremovegroupresponseparams/groupid-5tahm
func (m_ MTRGroupsClusterRemoveGroupResponseParams) SetGroupId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupId:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterremovegroupresponseparams/status
func (m_ MTRGroupsClusterRemoveGroupResponseParams) Status() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("status"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterremovegroupresponseparams/status
func (m_ MTRGroupsClusterRemoveGroupResponseParams) SetStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterremovegroupresponseparams/timedinvoketimeoutms
func (m_ MTRGroupsClusterRemoveGroupResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterremovegroupresponseparams/timedinvoketimeoutms
func (m_ MTRGroupsClusterRemoveGroupResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



