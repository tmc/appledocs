// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDoorLockClusterClearYearDayScheduleParams] class.
var (
	MTRDoorLockClusterClearYearDayScheduleParamsClass     _MTRDoorLockClusterClearYearDayScheduleParamsClass
	MTRDoorLockClusterClearYearDayScheduleParamsClassOnce sync.Once
)

func getMTRDoorLockClusterClearYearDayScheduleParamsClass() _MTRDoorLockClusterClearYearDayScheduleParamsClass {
	MTRDoorLockClusterClearYearDayScheduleParamsClassOnce.Do(func() {
		MTRDoorLockClusterClearYearDayScheduleParamsClass = _MTRDoorLockClusterClearYearDayScheduleParamsClass{objc.GetClass("MTRDoorLockClusterClearYearDayScheduleParams")}
	})
	return MTRDoorLockClusterClearYearDayScheduleParamsClass
}

type _MTRDoorLockClusterClearYearDayScheduleParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterClearYearDayScheduleParams] class.
type IMTRDoorLockClusterClearYearDayScheduleParams interface {
	objectivec.IObject
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	UserIndex() objc.IObject /* cross-framework: NSNumber */
	SetUserIndex(value objc.IObject /* cross-framework: NSNumber */)
	YearDayIndex() objc.IObject /* cross-framework: NSNumber */
	SetYearDayIndex(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearYearDayScheduleParams
type MTRDoorLockClusterClearYearDayScheduleParams struct {
	objectivec.Object
}

// MTRDoorLockClusterClearYearDayScheduleParamsFrom constructs a [MTRDoorLockClusterClearYearDayScheduleParams] from an unsafe.Pointer.
func MTRDoorLockClusterClearYearDayScheduleParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterClearYearDayScheduleParams {
	return MTRDoorLockClusterClearYearDayScheduleParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterClearYearDayScheduleParamsClass) Alloc() MTRDoorLockClusterClearYearDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterClearYearDayScheduleParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterClearYearDayScheduleParamsClass) New() MTRDoorLockClusterClearYearDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterClearYearDayScheduleParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterClearYearDayScheduleParams) Init() MTRDoorLockClusterClearYearDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterClearYearDayScheduleParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterClearYearDayScheduleParams) Autorelease() MTRDoorLockClusterClearYearDayScheduleParams {
	rv := objc.Send[MTRDoorLockClusterClearYearDayScheduleParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterClearYearDayScheduleParams creates a new MTRDoorLockClusterClearYearDayScheduleParams instance.
func NewMTRDoorLockClusterClearYearDayScheduleParams() MTRDoorLockClusterClearYearDayScheduleParams {
	return getMTRDoorLockClusterClearYearDayScheduleParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearyeardayscheduleparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterClearYearDayScheduleParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearyeardayscheduleparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterClearYearDayScheduleParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearyeardayscheduleparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterClearYearDayScheduleParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearyeardayscheduleparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterClearYearDayScheduleParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearyeardayscheduleparams/userindex
func (m_ MTRDoorLockClusterClearYearDayScheduleParams) UserIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearyeardayscheduleparams/userindex
func (m_ MTRDoorLockClusterClearYearDayScheduleParams) SetUserIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearyeardayscheduleparams/yeardayindex
func (m_ MTRDoorLockClusterClearYearDayScheduleParams) YearDayIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("yearDayIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearyeardayscheduleparams/yeardayindex
func (m_ MTRDoorLockClusterClearYearDayScheduleParams) SetYearDayIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setYearDayIndex:"), value)
}



