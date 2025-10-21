// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTREnergyEVSEModeClusterChangeToModeParams] class.
var (
	MTREnergyEVSEModeClusterChangeToModeParamsClass     _MTREnergyEVSEModeClusterChangeToModeParamsClass
	MTREnergyEVSEModeClusterChangeToModeParamsClassOnce sync.Once
)

func getMTREnergyEVSEModeClusterChangeToModeParamsClass() _MTREnergyEVSEModeClusterChangeToModeParamsClass {
	MTREnergyEVSEModeClusterChangeToModeParamsClassOnce.Do(func() {
		MTREnergyEVSEModeClusterChangeToModeParamsClass = _MTREnergyEVSEModeClusterChangeToModeParamsClass{objc.GetClass("MTREnergyEVSEModeClusterChangeToModeParams")}
	})
	return MTREnergyEVSEModeClusterChangeToModeParamsClass
}

type _MTREnergyEVSEModeClusterChangeToModeParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTREnergyEVSEModeClusterChangeToModeParams] class.
type IMTREnergyEVSEModeClusterChangeToModeParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterChangeToModeParams
type MTREnergyEVSEModeClusterChangeToModeParams struct {
	objectivec.Object
}

// MTREnergyEVSEModeClusterChangeToModeParamsFrom constructs a [MTREnergyEVSEModeClusterChangeToModeParams] from an unsafe.Pointer.
func MTREnergyEVSEModeClusterChangeToModeParamsFrom(ptr unsafe.Pointer) MTREnergyEVSEModeClusterChangeToModeParams {
	return MTREnergyEVSEModeClusterChangeToModeParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTREnergyEVSEModeClusterChangeToModeParamsClass) Alloc() MTREnergyEVSEModeClusterChangeToModeParams {
	rv := objc.Send[MTREnergyEVSEModeClusterChangeToModeParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTREnergyEVSEModeClusterChangeToModeParamsClass) New() MTREnergyEVSEModeClusterChangeToModeParams {
	rv := objc.Send[MTREnergyEVSEModeClusterChangeToModeParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREnergyEVSEModeClusterChangeToModeParams) Init() MTREnergyEVSEModeClusterChangeToModeParams {
	rv := objc.Send[MTREnergyEVSEModeClusterChangeToModeParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREnergyEVSEModeClusterChangeToModeParams) Autorelease() MTREnergyEVSEModeClusterChangeToModeParams {
	rv := objc.Send[MTREnergyEVSEModeClusterChangeToModeParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREnergyEVSEModeClusterChangeToModeParams creates a new MTREnergyEVSEModeClusterChangeToModeParams instance.
func NewMTREnergyEVSEModeClusterChangeToModeParams() MTREnergyEVSEModeClusterChangeToModeParams {
	return getMTREnergyEVSEModeClusterChangeToModeParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterChangeToModeParams/newMode
func (m_ MTREnergyEVSEModeClusterChangeToModeParams) NewMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("newMode"))
	return rv
}


// SetNewMode sets the value of the newMode property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterChangeToModeParams/newMode
func (m_ MTREnergyEVSEModeClusterChangeToModeParams) SetNewMode(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNewMode:"), value)
}

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterChangeToModeParams/serverSideProcessingTimeout
func (m_ MTREnergyEVSEModeClusterChangeToModeParams) ServerSideProcessingTimeout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterChangeToModeParams/serverSideProcessingTimeout
func (m_ MTREnergyEVSEModeClusterChangeToModeParams) SetServerSideProcessingTimeout(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterChangeToModeParams/timedInvokeTimeoutMs
func (m_ MTREnergyEVSEModeClusterChangeToModeParams) TimedInvokeTimeoutMs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterChangeToModeParams/timedInvokeTimeoutMs
func (m_ MTREnergyEVSEModeClusterChangeToModeParams) SetTimedInvokeTimeoutMs(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



