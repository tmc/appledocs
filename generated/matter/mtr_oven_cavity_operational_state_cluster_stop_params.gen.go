// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROvenCavityOperationalStateClusterStopParams] class.
var (
	MTROvenCavityOperationalStateClusterStopParamsClass     _MTROvenCavityOperationalStateClusterStopParamsClass
	MTROvenCavityOperationalStateClusterStopParamsClassOnce sync.Once
)

func getMTROvenCavityOperationalStateClusterStopParamsClass() _MTROvenCavityOperationalStateClusterStopParamsClass {
	MTROvenCavityOperationalStateClusterStopParamsClassOnce.Do(func() {
		MTROvenCavityOperationalStateClusterStopParamsClass = _MTROvenCavityOperationalStateClusterStopParamsClass{objc.GetClass("MTROvenCavityOperationalStateClusterStopParams")}
	})
	return MTROvenCavityOperationalStateClusterStopParamsClass
}

type _MTROvenCavityOperationalStateClusterStopParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROvenCavityOperationalStateClusterStopParams] class.
type IMTROvenCavityOperationalStateClusterStopParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterStopParams
type MTROvenCavityOperationalStateClusterStopParams struct {
	objectivec.Object
}

// MTROvenCavityOperationalStateClusterStopParamsFrom constructs a [MTROvenCavityOperationalStateClusterStopParams] from an unsafe.Pointer.
func MTROvenCavityOperationalStateClusterStopParamsFrom(ptr unsafe.Pointer) MTROvenCavityOperationalStateClusterStopParams {
	return MTROvenCavityOperationalStateClusterStopParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROvenCavityOperationalStateClusterStopParamsClass) Alloc() MTROvenCavityOperationalStateClusterStopParams {
	rv := objc.Send[MTROvenCavityOperationalStateClusterStopParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROvenCavityOperationalStateClusterStopParamsClass) New() MTROvenCavityOperationalStateClusterStopParams {
	rv := objc.Send[MTROvenCavityOperationalStateClusterStopParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROvenCavityOperationalStateClusterStopParams) Init() MTROvenCavityOperationalStateClusterStopParams {
	rv := objc.Send[MTROvenCavityOperationalStateClusterStopParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROvenCavityOperationalStateClusterStopParams) Autorelease() MTROvenCavityOperationalStateClusterStopParams {
	rv := objc.Send[MTROvenCavityOperationalStateClusterStopParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROvenCavityOperationalStateClusterStopParams creates a new MTROvenCavityOperationalStateClusterStopParams instance.
func NewMTROvenCavityOperationalStateClusterStopParams() MTROvenCavityOperationalStateClusterStopParams {
	return getMTROvenCavityOperationalStateClusterStopParamsClass().New()
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterStopParams/serverSideProcessingTimeout
func (m_ MTROvenCavityOperationalStateClusterStopParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterStopParams/serverSideProcessingTimeout
func (m_ MTROvenCavityOperationalStateClusterStopParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterStopParams/timedInvokeTimeoutMs
func (m_ MTROvenCavityOperationalStateClusterStopParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterStopParams/timedInvokeTimeoutMs
func (m_ MTROvenCavityOperationalStateClusterStopParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



