// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRBasicClusterMfgSpecificPingParams] class.
var (
	MTRBasicClusterMfgSpecificPingParamsClass     _MTRBasicClusterMfgSpecificPingParamsClass
	MTRBasicClusterMfgSpecificPingParamsClassOnce sync.Once
)

func getMTRBasicClusterMfgSpecificPingParamsClass() _MTRBasicClusterMfgSpecificPingParamsClass {
	MTRBasicClusterMfgSpecificPingParamsClassOnce.Do(func() {
		MTRBasicClusterMfgSpecificPingParamsClass = _MTRBasicClusterMfgSpecificPingParamsClass{objc.GetClass("MTRBasicClusterMfgSpecificPingParams")}
	})
	return MTRBasicClusterMfgSpecificPingParamsClass
}

type _MTRBasicClusterMfgSpecificPingParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRBasicClusterMfgSpecificPingParams] class.
type IMTRBasicClusterMfgSpecificPingParams interface {
	objectivec.IObject
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBasicClusterMfgSpecificPingParams
type MTRBasicClusterMfgSpecificPingParams struct {
	objectivec.Object
}

// MTRBasicClusterMfgSpecificPingParamsFrom constructs a [MTRBasicClusterMfgSpecificPingParams] from an unsafe.Pointer.
func MTRBasicClusterMfgSpecificPingParamsFrom(ptr unsafe.Pointer) MTRBasicClusterMfgSpecificPingParams {
	return MTRBasicClusterMfgSpecificPingParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBasicClusterMfgSpecificPingParamsClass) Alloc() MTRBasicClusterMfgSpecificPingParams {
	rv := objc.Send[MTRBasicClusterMfgSpecificPingParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBasicClusterMfgSpecificPingParamsClass) New() MTRBasicClusterMfgSpecificPingParams {
	rv := objc.Send[MTRBasicClusterMfgSpecificPingParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBasicClusterMfgSpecificPingParams) Init() MTRBasicClusterMfgSpecificPingParams {
	rv := objc.Send[MTRBasicClusterMfgSpecificPingParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBasicClusterMfgSpecificPingParams) Autorelease() MTRBasicClusterMfgSpecificPingParams {
	rv := objc.Send[MTRBasicClusterMfgSpecificPingParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBasicClusterMfgSpecificPingParams creates a new MTRBasicClusterMfgSpecificPingParams instance.
func NewMTRBasicClusterMfgSpecificPingParams() MTRBasicClusterMfgSpecificPingParams {
	return getMTRBasicClusterMfgSpecificPingParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbasicclustermfgspecificpingparams/serversideprocessingtimeout
func (m_ MTRBasicClusterMfgSpecificPingParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbasicclustermfgspecificpingparams/serversideprocessingtimeout
func (m_ MTRBasicClusterMfgSpecificPingParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbasicclustermfgspecificpingparams/timedinvoketimeoutms
func (m_ MTRBasicClusterMfgSpecificPingParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbasicclustermfgspecificpingparams/timedinvoketimeoutms
func (m_ MTRBasicClusterMfgSpecificPingParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



