// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROvenCavityOperationalStateClusterStartParams] class.
var (
	MTROvenCavityOperationalStateClusterStartParamsClass     _MTROvenCavityOperationalStateClusterStartParamsClass
	MTROvenCavityOperationalStateClusterStartParamsClassOnce sync.Once
)

func getMTROvenCavityOperationalStateClusterStartParamsClass() _MTROvenCavityOperationalStateClusterStartParamsClass {
	MTROvenCavityOperationalStateClusterStartParamsClassOnce.Do(func() {
		MTROvenCavityOperationalStateClusterStartParamsClass = _MTROvenCavityOperationalStateClusterStartParamsClass{objc.GetClass("MTROvenCavityOperationalStateClusterStartParams")}
	})
	return MTROvenCavityOperationalStateClusterStartParamsClass
}

type _MTROvenCavityOperationalStateClusterStartParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROvenCavityOperationalStateClusterStartParams] class.
type IMTROvenCavityOperationalStateClusterStartParams interface {
	objectivec.IObject
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterStartParams
type MTROvenCavityOperationalStateClusterStartParams struct {
	objectivec.Object
}

// MTROvenCavityOperationalStateClusterStartParamsFrom constructs a [MTROvenCavityOperationalStateClusterStartParams] from an unsafe.Pointer.
func MTROvenCavityOperationalStateClusterStartParamsFrom(ptr unsafe.Pointer) MTROvenCavityOperationalStateClusterStartParams {
	return MTROvenCavityOperationalStateClusterStartParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROvenCavityOperationalStateClusterStartParamsClass) Alloc() MTROvenCavityOperationalStateClusterStartParams {
	rv := objc.Send[MTROvenCavityOperationalStateClusterStartParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROvenCavityOperationalStateClusterStartParamsClass) New() MTROvenCavityOperationalStateClusterStartParams {
	rv := objc.Send[MTROvenCavityOperationalStateClusterStartParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROvenCavityOperationalStateClusterStartParams) Init() MTROvenCavityOperationalStateClusterStartParams {
	rv := objc.Send[MTROvenCavityOperationalStateClusterStartParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROvenCavityOperationalStateClusterStartParams) Autorelease() MTROvenCavityOperationalStateClusterStartParams {
	rv := objc.Send[MTROvenCavityOperationalStateClusterStartParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROvenCavityOperationalStateClusterStartParams creates a new MTROvenCavityOperationalStateClusterStartParams instance.
func NewMTROvenCavityOperationalStateClusterStartParams() MTROvenCavityOperationalStateClusterStartParams {
	return getMTROvenCavityOperationalStateClusterStartParamsClass().New()
}



// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterStartParams/serverSideProcessingTimeout
func (m_ MTROvenCavityOperationalStateClusterStartParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterStartParams/serverSideProcessingTimeout
func (m_ MTROvenCavityOperationalStateClusterStartParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterStartParams/timedInvokeTimeoutMs
func (m_ MTROvenCavityOperationalStateClusterStartParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterStartParams/timedInvokeTimeoutMs
func (m_ MTROvenCavityOperationalStateClusterStartParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



