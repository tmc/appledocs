// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRRVCCleanModeClusterChangeToModeParams] class.
var (
	MTRRVCCleanModeClusterChangeToModeParamsClass     _MTRRVCCleanModeClusterChangeToModeParamsClass
	MTRRVCCleanModeClusterChangeToModeParamsClassOnce sync.Once
)

func getMTRRVCCleanModeClusterChangeToModeParamsClass() _MTRRVCCleanModeClusterChangeToModeParamsClass {
	MTRRVCCleanModeClusterChangeToModeParamsClassOnce.Do(func() {
		MTRRVCCleanModeClusterChangeToModeParamsClass = _MTRRVCCleanModeClusterChangeToModeParamsClass{objc.GetClass("MTRRVCCleanModeClusterChangeToModeParams")}
	})
	return MTRRVCCleanModeClusterChangeToModeParamsClass
}

type _MTRRVCCleanModeClusterChangeToModeParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRRVCCleanModeClusterChangeToModeParams] class.
type IMTRRVCCleanModeClusterChangeToModeParams interface {
	objectivec.IObject
	// properties:
	NewMode() objc.IObject /* cross-framework: NSNumber */
	SetNewMode(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRVCCleanModeClusterChangeToModeParams
type MTRRVCCleanModeClusterChangeToModeParams struct {
	objectivec.Object
}

// MTRRVCCleanModeClusterChangeToModeParamsFrom constructs a [MTRRVCCleanModeClusterChangeToModeParams] from an unsafe.Pointer.
func MTRRVCCleanModeClusterChangeToModeParamsFrom(ptr unsafe.Pointer) MTRRVCCleanModeClusterChangeToModeParams {
	return MTRRVCCleanModeClusterChangeToModeParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRRVCCleanModeClusterChangeToModeParamsClass) Alloc() MTRRVCCleanModeClusterChangeToModeParams {
	rv := objc.Send[MTRRVCCleanModeClusterChangeToModeParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRRVCCleanModeClusterChangeToModeParamsClass) New() MTRRVCCleanModeClusterChangeToModeParams {
	rv := objc.Send[MTRRVCCleanModeClusterChangeToModeParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRRVCCleanModeClusterChangeToModeParams) Init() MTRRVCCleanModeClusterChangeToModeParams {
	rv := objc.Send[MTRRVCCleanModeClusterChangeToModeParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRRVCCleanModeClusterChangeToModeParams) Autorelease() MTRRVCCleanModeClusterChangeToModeParams {
	rv := objc.Send[MTRRVCCleanModeClusterChangeToModeParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRRVCCleanModeClusterChangeToModeParams creates a new MTRRVCCleanModeClusterChangeToModeParams instance.
func NewMTRRVCCleanModeClusterChangeToModeParams() MTRRVCCleanModeClusterChangeToModeParams {
	return getMTRRVCCleanModeClusterChangeToModeParamsClass().New()
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvccleanmodeclusterchangetomodeparams/newmode
func (m_ MTRRVCCleanModeClusterChangeToModeParams) NewMode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("newMode"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvccleanmodeclusterchangetomodeparams/newmode
func (m_ MTRRVCCleanModeClusterChangeToModeParams) SetNewMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNewMode:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvccleanmodeclusterchangetomodeparams/serversideprocessingtimeout
func (m_ MTRRVCCleanModeClusterChangeToModeParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvccleanmodeclusterchangetomodeparams/serversideprocessingtimeout
func (m_ MTRRVCCleanModeClusterChangeToModeParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvccleanmodeclusterchangetomodeparams/timedinvoketimeoutms
func (m_ MTRRVCCleanModeClusterChangeToModeParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvccleanmodeclusterchangetomodeparams/timedinvoketimeoutms
func (m_ MTRRVCCleanModeClusterChangeToModeParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}
