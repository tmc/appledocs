// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROvenModeClusterChangeToModeParams] class.
var (
	MTROvenModeClusterChangeToModeParamsClass     _MTROvenModeClusterChangeToModeParamsClass
	MTROvenModeClusterChangeToModeParamsClassOnce sync.Once
)

func getMTROvenModeClusterChangeToModeParamsClass() _MTROvenModeClusterChangeToModeParamsClass {
	MTROvenModeClusterChangeToModeParamsClassOnce.Do(func() {
		MTROvenModeClusterChangeToModeParamsClass = _MTROvenModeClusterChangeToModeParamsClass{objc.GetClass("MTROvenModeClusterChangeToModeParams")}
	})
	return MTROvenModeClusterChangeToModeParamsClass
}

type _MTROvenModeClusterChangeToModeParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROvenModeClusterChangeToModeParams] class.
type IMTROvenModeClusterChangeToModeParams interface {
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
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterChangeToModeParams
type MTROvenModeClusterChangeToModeParams struct {
	objectivec.Object
}

// MTROvenModeClusterChangeToModeParamsFrom constructs a [MTROvenModeClusterChangeToModeParams] from an unsafe.Pointer.
func MTROvenModeClusterChangeToModeParamsFrom(ptr unsafe.Pointer) MTROvenModeClusterChangeToModeParams {
	return MTROvenModeClusterChangeToModeParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROvenModeClusterChangeToModeParamsClass) Alloc() MTROvenModeClusterChangeToModeParams {
	rv := objc.Send[MTROvenModeClusterChangeToModeParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROvenModeClusterChangeToModeParamsClass) New() MTROvenModeClusterChangeToModeParams {
	rv := objc.Send[MTROvenModeClusterChangeToModeParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROvenModeClusterChangeToModeParams) Init() MTROvenModeClusterChangeToModeParams {
	rv := objc.Send[MTROvenModeClusterChangeToModeParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROvenModeClusterChangeToModeParams) Autorelease() MTROvenModeClusterChangeToModeParams {
	rv := objc.Send[MTROvenModeClusterChangeToModeParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROvenModeClusterChangeToModeParams creates a new MTROvenModeClusterChangeToModeParams instance.
func NewMTROvenModeClusterChangeToModeParams() MTROvenModeClusterChangeToModeParams {
	return getMTROvenModeClusterChangeToModeParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterChangeToModeParams/newMode
func (m_ MTROvenModeClusterChangeToModeParams) NewMode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("newMode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterChangeToModeParams/newMode
func (m_ MTROvenModeClusterChangeToModeParams) SetNewMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNewMode:"), value)
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterChangeToModeParams/serverSideProcessingTimeout
func (m_ MTROvenModeClusterChangeToModeParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterChangeToModeParams/serverSideProcessingTimeout
func (m_ MTROvenModeClusterChangeToModeParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterChangeToModeParams/timedInvokeTimeoutMs
func (m_ MTROvenModeClusterChangeToModeParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterChangeToModeParams/timedInvokeTimeoutMs
func (m_ MTROvenModeClusterChangeToModeParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



