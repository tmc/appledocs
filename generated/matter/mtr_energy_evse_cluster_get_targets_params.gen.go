// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTREnergyEVSEClusterGetTargetsParams] class.
var (
	MTREnergyEVSEClusterGetTargetsParamsClass     _MTREnergyEVSEClusterGetTargetsParamsClass
	MTREnergyEVSEClusterGetTargetsParamsClassOnce sync.Once
)

func getMTREnergyEVSEClusterGetTargetsParamsClass() _MTREnergyEVSEClusterGetTargetsParamsClass {
	MTREnergyEVSEClusterGetTargetsParamsClassOnce.Do(func() {
		MTREnergyEVSEClusterGetTargetsParamsClass = _MTREnergyEVSEClusterGetTargetsParamsClass{objc.GetClass("MTREnergyEVSEClusterGetTargetsParams")}
	})
	return MTREnergyEVSEClusterGetTargetsParamsClass
}

type _MTREnergyEVSEClusterGetTargetsParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTREnergyEVSEClusterGetTargetsParams] class.
type IMTREnergyEVSEClusterGetTargetsParams interface {
	objectivec.IObject
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterGetTargetsParams
type MTREnergyEVSEClusterGetTargetsParams struct {
	objectivec.Object
}

// MTREnergyEVSEClusterGetTargetsParamsFrom constructs a [MTREnergyEVSEClusterGetTargetsParams] from an unsafe.Pointer.
func MTREnergyEVSEClusterGetTargetsParamsFrom(ptr unsafe.Pointer) MTREnergyEVSEClusterGetTargetsParams {
	return MTREnergyEVSEClusterGetTargetsParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTREnergyEVSEClusterGetTargetsParamsClass) Alloc() MTREnergyEVSEClusterGetTargetsParams {
	rv := objc.Send[MTREnergyEVSEClusterGetTargetsParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTREnergyEVSEClusterGetTargetsParamsClass) New() MTREnergyEVSEClusterGetTargetsParams {
	rv := objc.Send[MTREnergyEVSEClusterGetTargetsParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREnergyEVSEClusterGetTargetsParams) Init() MTREnergyEVSEClusterGetTargetsParams {
	rv := objc.Send[MTREnergyEVSEClusterGetTargetsParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREnergyEVSEClusterGetTargetsParams) Autorelease() MTREnergyEVSEClusterGetTargetsParams {
	rv := objc.Send[MTREnergyEVSEClusterGetTargetsParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREnergyEVSEClusterGetTargetsParams creates a new MTREnergyEVSEClusterGetTargetsParams instance.
func NewMTREnergyEVSEClusterGetTargetsParams() MTREnergyEVSEClusterGetTargetsParams {
	return getMTREnergyEVSEClusterGetTargetsParamsClass().New()
}



// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterGetTargetsParams/serverSideProcessingTimeout
func (m_ MTREnergyEVSEClusterGetTargetsParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterGetTargetsParams/serverSideProcessingTimeout
func (m_ MTREnergyEVSEClusterGetTargetsParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterGetTargetsParams/timedInvokeTimeoutMs
func (m_ MTREnergyEVSEClusterGetTargetsParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterGetTargetsParams/timedInvokeTimeoutMs
func (m_ MTREnergyEVSEClusterGetTargetsParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



