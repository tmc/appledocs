// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRGroupsClusterGetGroupMembershipParams] class.
var (
	MTRGroupsClusterGetGroupMembershipParamsClass     _MTRGroupsClusterGetGroupMembershipParamsClass
	MTRGroupsClusterGetGroupMembershipParamsClassOnce sync.Once
)

func getMTRGroupsClusterGetGroupMembershipParamsClass() _MTRGroupsClusterGetGroupMembershipParamsClass {
	MTRGroupsClusterGetGroupMembershipParamsClassOnce.Do(func() {
		MTRGroupsClusterGetGroupMembershipParamsClass = _MTRGroupsClusterGetGroupMembershipParamsClass{objc.GetClass("MTRGroupsClusterGetGroupMembershipParams")}
	})
	return MTRGroupsClusterGetGroupMembershipParamsClass
}

type _MTRGroupsClusterGetGroupMembershipParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRGroupsClusterGetGroupMembershipParams] class.
type IMTRGroupsClusterGetGroupMembershipParams interface {
	objectivec.IObject
	// properties:
	GroupList() unsafe.Pointer
	SetGroupList(value unsafe.Pointer)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGroupsClusterGetGroupMembershipParams
type MTRGroupsClusterGetGroupMembershipParams struct {
	objectivec.Object
}

// MTRGroupsClusterGetGroupMembershipParamsFrom constructs a [MTRGroupsClusterGetGroupMembershipParams] from an unsafe.Pointer.
func MTRGroupsClusterGetGroupMembershipParamsFrom(ptr unsafe.Pointer) MTRGroupsClusterGetGroupMembershipParams {
	return MTRGroupsClusterGetGroupMembershipParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGroupsClusterGetGroupMembershipParamsClass) Alloc() MTRGroupsClusterGetGroupMembershipParams {
	rv := objc.Send[MTRGroupsClusterGetGroupMembershipParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGroupsClusterGetGroupMembershipParamsClass) New() MTRGroupsClusterGetGroupMembershipParams {
	rv := objc.Send[MTRGroupsClusterGetGroupMembershipParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGroupsClusterGetGroupMembershipParams) Init() MTRGroupsClusterGetGroupMembershipParams {
	rv := objc.Send[MTRGroupsClusterGetGroupMembershipParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGroupsClusterGetGroupMembershipParams) Autorelease() MTRGroupsClusterGetGroupMembershipParams {
	rv := objc.Send[MTRGroupsClusterGetGroupMembershipParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGroupsClusterGetGroupMembershipParams creates a new MTRGroupsClusterGetGroupMembershipParams instance.
func NewMTRGroupsClusterGetGroupMembershipParams() MTRGroupsClusterGetGroupMembershipParams {
	return getMTRGroupsClusterGetGroupMembershipParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclustergetgroupmembershipparams/grouplist
func (m_ MTRGroupsClusterGetGroupMembershipParams) GroupList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("groupList"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclustergetgroupmembershipparams/grouplist
func (m_ MTRGroupsClusterGetGroupMembershipParams) SetGroupList(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupList:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclustergetgroupmembershipparams/serversideprocessingtimeout
func (m_ MTRGroupsClusterGetGroupMembershipParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclustergetgroupmembershipparams/serversideprocessingtimeout
func (m_ MTRGroupsClusterGetGroupMembershipParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclustergetgroupmembershipparams/timedinvoketimeoutms
func (m_ MTRGroupsClusterGetGroupMembershipParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupsclustergetgroupmembershipparams/timedinvoketimeoutms
func (m_ MTRGroupsClusterGetGroupMembershipParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



