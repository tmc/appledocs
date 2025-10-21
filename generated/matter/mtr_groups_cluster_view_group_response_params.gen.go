// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterviewgroupresponseparams/status
func (m_ MTRGroupsClusterViewGroupResponseParams) Status() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("status"))
	return rv
}


// SetStatus sets the value of the status property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterviewgroupresponseparams/status
func (m_ MTRGroupsClusterViewGroupResponseParams) SetStatus(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterviewgroupresponseparams/groupid-7jt0b
func (m_ MTRGroupsClusterViewGroupResponseParams) GroupId() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("groupId"))
	return rv
}


// SetGroupId sets the value of the groupId property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterviewgroupresponseparams/groupid-7jt0b
func (m_ MTRGroupsClusterViewGroupResponseParams) SetGroupId(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupId:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterviewgroupresponseparams/timedinvoketimeoutms
func (m_ MTRGroupsClusterViewGroupResponseParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterviewgroupresponseparams/timedinvoketimeoutms
func (m_ MTRGroupsClusterViewGroupResponseParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterviewgroupresponseparams/groupid-7jt17
func (m_ MTRGroupsClusterViewGroupResponseParams) GroupID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("groupID"))
	return rv
}


// SetGroupID sets the value of the groupID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterviewgroupresponseparams/groupid-7jt17
func (m_ MTRGroupsClusterViewGroupResponseParams) SetGroupID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterviewgroupresponseparams/groupname
func (m_ MTRGroupsClusterViewGroupResponseParams) GroupName() string {
	rv := objc.Send[string](m_.ID, objc.Sel("groupName"))
	return rv
}


// SetGroupName sets the value of the groupName property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclusterviewgroupresponseparams/groupname
func (m_ MTRGroupsClusterViewGroupResponseParams) SetGroupName(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupName:"), objc.String(value))
}



