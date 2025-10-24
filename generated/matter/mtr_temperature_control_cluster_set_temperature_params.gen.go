// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRTemperatureControlClusterSetTemperatureParams] class.
var (
	MTRTemperatureControlClusterSetTemperatureParamsClass     _MTRTemperatureControlClusterSetTemperatureParamsClass
	MTRTemperatureControlClusterSetTemperatureParamsClassOnce sync.Once
)

func getMTRTemperatureControlClusterSetTemperatureParamsClass() _MTRTemperatureControlClusterSetTemperatureParamsClass {
	MTRTemperatureControlClusterSetTemperatureParamsClassOnce.Do(func() {
		MTRTemperatureControlClusterSetTemperatureParamsClass = _MTRTemperatureControlClusterSetTemperatureParamsClass{objc.GetClass("MTRTemperatureControlClusterSetTemperatureParams")}
	})
	return MTRTemperatureControlClusterSetTemperatureParamsClass
}

type _MTRTemperatureControlClusterSetTemperatureParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTemperatureControlClusterSetTemperatureParams] class.
type IMTRTemperatureControlClusterSetTemperatureParams interface {
	objectivec.IObject
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TargetTemperature() objc.IObject /* cross-framework: NSNumber */
	SetTargetTemperature(value objc.IObject /* cross-framework: NSNumber */)
	TargetTemperatureLevel() objc.IObject /* cross-framework: NSNumber */
	SetTargetTemperatureLevel(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTemperatureControlClusterSetTemperatureParams
type MTRTemperatureControlClusterSetTemperatureParams struct {
	objectivec.Object
}

// MTRTemperatureControlClusterSetTemperatureParamsFrom constructs a [MTRTemperatureControlClusterSetTemperatureParams] from an unsafe.Pointer.
func MTRTemperatureControlClusterSetTemperatureParamsFrom(ptr unsafe.Pointer) MTRTemperatureControlClusterSetTemperatureParams {
	return MTRTemperatureControlClusterSetTemperatureParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTemperatureControlClusterSetTemperatureParamsClass) Alloc() MTRTemperatureControlClusterSetTemperatureParams {
	rv := objc.Send[MTRTemperatureControlClusterSetTemperatureParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTemperatureControlClusterSetTemperatureParamsClass) New() MTRTemperatureControlClusterSetTemperatureParams {
	rv := objc.Send[MTRTemperatureControlClusterSetTemperatureParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTemperatureControlClusterSetTemperatureParams) Init() MTRTemperatureControlClusterSetTemperatureParams {
	rv := objc.Send[MTRTemperatureControlClusterSetTemperatureParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTemperatureControlClusterSetTemperatureParams) Autorelease() MTRTemperatureControlClusterSetTemperatureParams {
	rv := objc.Send[MTRTemperatureControlClusterSetTemperatureParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTemperatureControlClusterSetTemperatureParams creates a new MTRTemperatureControlClusterSetTemperatureParams instance.
func NewMTRTemperatureControlClusterSetTemperatureParams() MTRTemperatureControlClusterSetTemperatureParams {
	return getMTRTemperatureControlClusterSetTemperatureParamsClass().New()
}



// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTemperatureControlClusterSetTemperatureParams/serverSideProcessingTimeout
func (m_ MTRTemperatureControlClusterSetTemperatureParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTemperatureControlClusterSetTemperatureParams/serverSideProcessingTimeout
func (m_ MTRTemperatureControlClusterSetTemperatureParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTemperatureControlClusterSetTemperatureParams/targetTemperature
func (m_ MTRTemperatureControlClusterSetTemperatureParams) TargetTemperature() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("targetTemperature"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTemperatureControlClusterSetTemperatureParams/targetTemperature
func (m_ MTRTemperatureControlClusterSetTemperatureParams) SetTargetTemperature(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTargetTemperature:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTemperatureControlClusterSetTemperatureParams/targetTemperatureLevel
func (m_ MTRTemperatureControlClusterSetTemperatureParams) TargetTemperatureLevel() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("targetTemperatureLevel"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTemperatureControlClusterSetTemperatureParams/targetTemperatureLevel
func (m_ MTRTemperatureControlClusterSetTemperatureParams) SetTargetTemperatureLevel(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTargetTemperatureLevel:"), value)
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTemperatureControlClusterSetTemperatureParams/timedInvokeTimeoutMs
func (m_ MTRTemperatureControlClusterSetTemperatureParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTemperatureControlClusterSetTemperatureParams/timedInvokeTimeoutMs
func (m_ MTRTemperatureControlClusterSetTemperatureParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



