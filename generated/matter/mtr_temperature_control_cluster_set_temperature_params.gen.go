// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
}

//
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
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTemperatureControlClusterSetTemperatureParams/serverSideProcessingTimeout
func (m_ MTRTemperatureControlClusterSetTemperatureParams) ServerSideProcessingTimeout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTemperatureControlClusterSetTemperatureParams/serverSideProcessingTimeout
func (m_ MTRTemperatureControlClusterSetTemperatureParams) SetServerSideProcessingTimeout(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTemperatureControlClusterSetTemperatureParams/targetTemperature
func (m_ MTRTemperatureControlClusterSetTemperatureParams) TargetTemperature() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("targetTemperature"))
	return rv
}


// SetTargetTemperature sets the value of the targetTemperature property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTemperatureControlClusterSetTemperatureParams/targetTemperature
func (m_ MTRTemperatureControlClusterSetTemperatureParams) SetTargetTemperature(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTargetTemperature:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTemperatureControlClusterSetTemperatureParams/targetTemperatureLevel
func (m_ MTRTemperatureControlClusterSetTemperatureParams) TargetTemperatureLevel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("targetTemperatureLevel"))
	return rv
}


// SetTargetTemperatureLevel sets the value of the targetTemperatureLevel property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTemperatureControlClusterSetTemperatureParams/targetTemperatureLevel
func (m_ MTRTemperatureControlClusterSetTemperatureParams) SetTargetTemperatureLevel(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTargetTemperatureLevel:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTemperatureControlClusterSetTemperatureParams/timedInvokeTimeoutMs
func (m_ MTRTemperatureControlClusterSetTemperatureParams) TimedInvokeTimeoutMs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTemperatureControlClusterSetTemperatureParams/timedInvokeTimeoutMs
func (m_ MTRTemperatureControlClusterSetTemperatureParams) SetTimedInvokeTimeoutMs(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



