// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRGeneralCommissioningClusterCommissioningCompleteParams] class.
var (
	MTRGeneralCommissioningClusterCommissioningCompleteParamsClass     _MTRGeneralCommissioningClusterCommissioningCompleteParamsClass
	MTRGeneralCommissioningClusterCommissioningCompleteParamsClassOnce sync.Once
)

func getMTRGeneralCommissioningClusterCommissioningCompleteParamsClass() _MTRGeneralCommissioningClusterCommissioningCompleteParamsClass {
	MTRGeneralCommissioningClusterCommissioningCompleteParamsClassOnce.Do(func() {
		MTRGeneralCommissioningClusterCommissioningCompleteParamsClass = _MTRGeneralCommissioningClusterCommissioningCompleteParamsClass{objc.GetClass("MTRGeneralCommissioningClusterCommissioningCompleteParams")}
	})
	return MTRGeneralCommissioningClusterCommissioningCompleteParamsClass
}

type _MTRGeneralCommissioningClusterCommissioningCompleteParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRGeneralCommissioningClusterCommissioningCompleteParams] class.
type IMTRGeneralCommissioningClusterCommissioningCompleteParams interface {
	objectivec.IObject
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralCommissioningClusterCommissioningCompleteParams
type MTRGeneralCommissioningClusterCommissioningCompleteParams struct {
	objectivec.Object
}

// MTRGeneralCommissioningClusterCommissioningCompleteParamsFrom constructs a [MTRGeneralCommissioningClusterCommissioningCompleteParams] from an unsafe.Pointer.
func MTRGeneralCommissioningClusterCommissioningCompleteParamsFrom(ptr unsafe.Pointer) MTRGeneralCommissioningClusterCommissioningCompleteParams {
	return MTRGeneralCommissioningClusterCommissioningCompleteParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGeneralCommissioningClusterCommissioningCompleteParamsClass) Alloc() MTRGeneralCommissioningClusterCommissioningCompleteParams {
	rv := objc.Send[MTRGeneralCommissioningClusterCommissioningCompleteParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGeneralCommissioningClusterCommissioningCompleteParamsClass) New() MTRGeneralCommissioningClusterCommissioningCompleteParams {
	rv := objc.Send[MTRGeneralCommissioningClusterCommissioningCompleteParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGeneralCommissioningClusterCommissioningCompleteParams) Init() MTRGeneralCommissioningClusterCommissioningCompleteParams {
	rv := objc.Send[MTRGeneralCommissioningClusterCommissioningCompleteParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGeneralCommissioningClusterCommissioningCompleteParams) Autorelease() MTRGeneralCommissioningClusterCommissioningCompleteParams {
	rv := objc.Send[MTRGeneralCommissioningClusterCommissioningCompleteParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGeneralCommissioningClusterCommissioningCompleteParams creates a new MTRGeneralCommissioningClusterCommissioningCompleteParams instance.
func NewMTRGeneralCommissioningClusterCommissioningCompleteParams() MTRGeneralCommissioningClusterCommissioningCompleteParams {
	return getMTRGeneralCommissioningClusterCommissioningCompleteParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclustercommissioningcompleteparams/serversideprocessingtimeout
func (m_ MTRGeneralCommissioningClusterCommissioningCompleteParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclustercommissioningcompleteparams/serversideprocessingtimeout
func (m_ MTRGeneralCommissioningClusterCommissioningCompleteParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclustercommissioningcompleteparams/timedinvoketimeoutms
func (m_ MTRGeneralCommissioningClusterCommissioningCompleteParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclustercommissioningcompleteparams/timedinvoketimeoutms
func (m_ MTRGeneralCommissioningClusterCommissioningCompleteParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



