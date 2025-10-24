// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRRVCRunModeClusterChangeToModeParams] class.
var (
	MTRRVCRunModeClusterChangeToModeParamsClass     _MTRRVCRunModeClusterChangeToModeParamsClass
	MTRRVCRunModeClusterChangeToModeParamsClassOnce sync.Once
)

func getMTRRVCRunModeClusterChangeToModeParamsClass() _MTRRVCRunModeClusterChangeToModeParamsClass {
	MTRRVCRunModeClusterChangeToModeParamsClassOnce.Do(func() {
		MTRRVCRunModeClusterChangeToModeParamsClass = _MTRRVCRunModeClusterChangeToModeParamsClass{objc.GetClass("MTRRVCRunModeClusterChangeToModeParams")}
	})
	return MTRRVCRunModeClusterChangeToModeParamsClass
}

type _MTRRVCRunModeClusterChangeToModeParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRRVCRunModeClusterChangeToModeParams] class.
type IMTRRVCRunModeClusterChangeToModeParams interface {
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
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRVCRunModeClusterChangeToModeParams
type MTRRVCRunModeClusterChangeToModeParams struct {
	objectivec.Object
}

// MTRRVCRunModeClusterChangeToModeParamsFrom constructs a [MTRRVCRunModeClusterChangeToModeParams] from an unsafe.Pointer.
func MTRRVCRunModeClusterChangeToModeParamsFrom(ptr unsafe.Pointer) MTRRVCRunModeClusterChangeToModeParams {
	return MTRRVCRunModeClusterChangeToModeParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRRVCRunModeClusterChangeToModeParamsClass) Alloc() MTRRVCRunModeClusterChangeToModeParams {
	rv := objc.Send[MTRRVCRunModeClusterChangeToModeParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRRVCRunModeClusterChangeToModeParamsClass) New() MTRRVCRunModeClusterChangeToModeParams {
	rv := objc.Send[MTRRVCRunModeClusterChangeToModeParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRRVCRunModeClusterChangeToModeParams) Init() MTRRVCRunModeClusterChangeToModeParams {
	rv := objc.Send[MTRRVCRunModeClusterChangeToModeParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRRVCRunModeClusterChangeToModeParams) Autorelease() MTRRVCRunModeClusterChangeToModeParams {
	rv := objc.Send[MTRRVCRunModeClusterChangeToModeParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRRVCRunModeClusterChangeToModeParams creates a new MTRRVCRunModeClusterChangeToModeParams instance.
func NewMTRRVCRunModeClusterChangeToModeParams() MTRRVCRunModeClusterChangeToModeParams {
	return getMTRRVCRunModeClusterChangeToModeParamsClass().New()
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvcrunmodeclusterchangetomodeparams/newmode
func (m_ MTRRVCRunModeClusterChangeToModeParams) NewMode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("newMode"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvcrunmodeclusterchangetomodeparams/newmode
func (m_ MTRRVCRunModeClusterChangeToModeParams) SetNewMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNewMode:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvcrunmodeclusterchangetomodeparams/serversideprocessingtimeout
func (m_ MTRRVCRunModeClusterChangeToModeParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvcrunmodeclusterchangetomodeparams/serversideprocessingtimeout
func (m_ MTRRVCRunModeClusterChangeToModeParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvcrunmodeclusterchangetomodeparams/timedinvoketimeoutms
func (m_ MTRRVCRunModeClusterChangeToModeParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvcrunmodeclusterchangetomodeparams/timedinvoketimeoutms
func (m_ MTRRVCRunModeClusterChangeToModeParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}
