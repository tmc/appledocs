// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRLaundryWasherModeClusterChangeToModeParams] class.
var (
	MTRLaundryWasherModeClusterChangeToModeParamsClass     _MTRLaundryWasherModeClusterChangeToModeParamsClass
	MTRLaundryWasherModeClusterChangeToModeParamsClassOnce sync.Once
)

func getMTRLaundryWasherModeClusterChangeToModeParamsClass() _MTRLaundryWasherModeClusterChangeToModeParamsClass {
	MTRLaundryWasherModeClusterChangeToModeParamsClassOnce.Do(func() {
		MTRLaundryWasherModeClusterChangeToModeParamsClass = _MTRLaundryWasherModeClusterChangeToModeParamsClass{objc.GetClass("MTRLaundryWasherModeClusterChangeToModeParams")}
	})
	return MTRLaundryWasherModeClusterChangeToModeParamsClass
}

type _MTRLaundryWasherModeClusterChangeToModeParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRLaundryWasherModeClusterChangeToModeParams] class.
type IMTRLaundryWasherModeClusterChangeToModeParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLaundryWasherModeClusterChangeToModeParams
type MTRLaundryWasherModeClusterChangeToModeParams struct {
	objectivec.Object
}

// MTRLaundryWasherModeClusterChangeToModeParamsFrom constructs a [MTRLaundryWasherModeClusterChangeToModeParams] from an unsafe.Pointer.
func MTRLaundryWasherModeClusterChangeToModeParamsFrom(ptr unsafe.Pointer) MTRLaundryWasherModeClusterChangeToModeParams {
	return MTRLaundryWasherModeClusterChangeToModeParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRLaundryWasherModeClusterChangeToModeParamsClass) Alloc() MTRLaundryWasherModeClusterChangeToModeParams {
	rv := objc.Send[MTRLaundryWasherModeClusterChangeToModeParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRLaundryWasherModeClusterChangeToModeParamsClass) New() MTRLaundryWasherModeClusterChangeToModeParams {
	rv := objc.Send[MTRLaundryWasherModeClusterChangeToModeParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRLaundryWasherModeClusterChangeToModeParams) Init() MTRLaundryWasherModeClusterChangeToModeParams {
	rv := objc.Send[MTRLaundryWasherModeClusterChangeToModeParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRLaundryWasherModeClusterChangeToModeParams) Autorelease() MTRLaundryWasherModeClusterChangeToModeParams {
	rv := objc.Send[MTRLaundryWasherModeClusterChangeToModeParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRLaundryWasherModeClusterChangeToModeParams creates a new MTRLaundryWasherModeClusterChangeToModeParams instance.
func NewMTRLaundryWasherModeClusterChangeToModeParams() MTRLaundryWasherModeClusterChangeToModeParams {
	return getMTRLaundryWasherModeClusterChangeToModeParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLaundryWasherModeClusterChangeToModeParams/newMode
func (m_ MTRLaundryWasherModeClusterChangeToModeParams) NewMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("newMode"))
	return rv
}


// SetNewMode sets the value of the newMode property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLaundryWasherModeClusterChangeToModeParams/newMode
func (m_ MTRLaundryWasherModeClusterChangeToModeParams) SetNewMode(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNewMode:"), value)
}
// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLaundryWasherModeClusterChangeToModeParams/serverSideProcessingTimeout
func (m_ MTRLaundryWasherModeClusterChangeToModeParams) ServerSideProcessingTimeout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLaundryWasherModeClusterChangeToModeParams/serverSideProcessingTimeout
func (m_ MTRLaundryWasherModeClusterChangeToModeParams) SetServerSideProcessingTimeout(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}
// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLaundryWasherModeClusterChangeToModeParams/timedInvokeTimeoutMs
func (m_ MTRLaundryWasherModeClusterChangeToModeParams) TimedInvokeTimeoutMs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLaundryWasherModeClusterChangeToModeParams/timedInvokeTimeoutMs
func (m_ MTRLaundryWasherModeClusterChangeToModeParams) SetTimedInvokeTimeoutMs(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


