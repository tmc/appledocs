// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTREnergyEVSEClusterClearTargetsParams] class.
var (
	MTREnergyEVSEClusterClearTargetsParamsClass     _MTREnergyEVSEClusterClearTargetsParamsClass
	MTREnergyEVSEClusterClearTargetsParamsClassOnce sync.Once
)

func getMTREnergyEVSEClusterClearTargetsParamsClass() _MTREnergyEVSEClusterClearTargetsParamsClass {
	MTREnergyEVSEClusterClearTargetsParamsClassOnce.Do(func() {
		MTREnergyEVSEClusterClearTargetsParamsClass = _MTREnergyEVSEClusterClearTargetsParamsClass{objc.GetClass("MTREnergyEVSEClusterClearTargetsParams")}
	})
	return MTREnergyEVSEClusterClearTargetsParamsClass
}

type _MTREnergyEVSEClusterClearTargetsParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTREnergyEVSEClusterClearTargetsParams] class.
type IMTREnergyEVSEClusterClearTargetsParams interface {
	objectivec.IObject
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterClearTargetsParams
type MTREnergyEVSEClusterClearTargetsParams struct {
	objectivec.Object
}

// MTREnergyEVSEClusterClearTargetsParamsFrom constructs a [MTREnergyEVSEClusterClearTargetsParams] from an unsafe.Pointer.
func MTREnergyEVSEClusterClearTargetsParamsFrom(ptr unsafe.Pointer) MTREnergyEVSEClusterClearTargetsParams {
	return MTREnergyEVSEClusterClearTargetsParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTREnergyEVSEClusterClearTargetsParamsClass) Alloc() MTREnergyEVSEClusterClearTargetsParams {
	rv := objc.Send[MTREnergyEVSEClusterClearTargetsParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTREnergyEVSEClusterClearTargetsParamsClass) New() MTREnergyEVSEClusterClearTargetsParams {
	rv := objc.Send[MTREnergyEVSEClusterClearTargetsParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREnergyEVSEClusterClearTargetsParams) Init() MTREnergyEVSEClusterClearTargetsParams {
	rv := objc.Send[MTREnergyEVSEClusterClearTargetsParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREnergyEVSEClusterClearTargetsParams) Autorelease() MTREnergyEVSEClusterClearTargetsParams {
	rv := objc.Send[MTREnergyEVSEClusterClearTargetsParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREnergyEVSEClusterClearTargetsParams creates a new MTREnergyEVSEClusterClearTargetsParams instance.
func NewMTREnergyEVSEClusterClearTargetsParams() MTREnergyEVSEClusterClearTargetsParams {
	return getMTREnergyEVSEClusterClearTargetsParamsClass().New()
}



// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterClearTargetsParams/serverSideProcessingTimeout
func (m_ MTREnergyEVSEClusterClearTargetsParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterClearTargetsParams/serverSideProcessingTimeout
func (m_ MTREnergyEVSEClusterClearTargetsParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterClearTargetsParams/timedInvokeTimeoutMs
func (m_ MTREnergyEVSEClusterClearTargetsParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterClearTargetsParams/timedInvokeTimeoutMs
func (m_ MTREnergyEVSEClusterClearTargetsParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



