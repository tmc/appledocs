// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRMicrowaveOvenControlClusterAddMoreTimeParams] class.
var (
	MTRMicrowaveOvenControlClusterAddMoreTimeParamsClass     _MTRMicrowaveOvenControlClusterAddMoreTimeParamsClass
	MTRMicrowaveOvenControlClusterAddMoreTimeParamsClassOnce sync.Once
)

func getMTRMicrowaveOvenControlClusterAddMoreTimeParamsClass() _MTRMicrowaveOvenControlClusterAddMoreTimeParamsClass {
	MTRMicrowaveOvenControlClusterAddMoreTimeParamsClassOnce.Do(func() {
		MTRMicrowaveOvenControlClusterAddMoreTimeParamsClass = _MTRMicrowaveOvenControlClusterAddMoreTimeParamsClass{objc.GetClass("MTRMicrowaveOvenControlClusterAddMoreTimeParams")}
	})
	return MTRMicrowaveOvenControlClusterAddMoreTimeParamsClass
}

type _MTRMicrowaveOvenControlClusterAddMoreTimeParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRMicrowaveOvenControlClusterAddMoreTimeParams] class.
type IMTRMicrowaveOvenControlClusterAddMoreTimeParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenControlClusterAddMoreTimeParams
type MTRMicrowaveOvenControlClusterAddMoreTimeParams struct {
	objectivec.Object
}

// MTRMicrowaveOvenControlClusterAddMoreTimeParamsFrom constructs a [MTRMicrowaveOvenControlClusterAddMoreTimeParams] from an unsafe.Pointer.
func MTRMicrowaveOvenControlClusterAddMoreTimeParamsFrom(ptr unsafe.Pointer) MTRMicrowaveOvenControlClusterAddMoreTimeParams {
	return MTRMicrowaveOvenControlClusterAddMoreTimeParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMicrowaveOvenControlClusterAddMoreTimeParamsClass) Alloc() MTRMicrowaveOvenControlClusterAddMoreTimeParams {
	rv := objc.Send[MTRMicrowaveOvenControlClusterAddMoreTimeParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMicrowaveOvenControlClusterAddMoreTimeParamsClass) New() MTRMicrowaveOvenControlClusterAddMoreTimeParams {
	rv := objc.Send[MTRMicrowaveOvenControlClusterAddMoreTimeParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMicrowaveOvenControlClusterAddMoreTimeParams) Init() MTRMicrowaveOvenControlClusterAddMoreTimeParams {
	rv := objc.Send[MTRMicrowaveOvenControlClusterAddMoreTimeParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMicrowaveOvenControlClusterAddMoreTimeParams) Autorelease() MTRMicrowaveOvenControlClusterAddMoreTimeParams {
	rv := objc.Send[MTRMicrowaveOvenControlClusterAddMoreTimeParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMicrowaveOvenControlClusterAddMoreTimeParams creates a new MTRMicrowaveOvenControlClusterAddMoreTimeParams instance.
func NewMTRMicrowaveOvenControlClusterAddMoreTimeParams() MTRMicrowaveOvenControlClusterAddMoreTimeParams {
	return getMTRMicrowaveOvenControlClusterAddMoreTimeParamsClass().New()
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenControlClusterAddMoreTimeParams/serverSideProcessingTimeout
func (m_ MTRMicrowaveOvenControlClusterAddMoreTimeParams) ServerSideProcessingTimeout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenControlClusterAddMoreTimeParams/serverSideProcessingTimeout
func (m_ MTRMicrowaveOvenControlClusterAddMoreTimeParams) SetServerSideProcessingTimeout(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenControlClusterAddMoreTimeParams/timeToAdd
func (m_ MTRMicrowaveOvenControlClusterAddMoreTimeParams) TimeToAdd() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timeToAdd"))
	return rv
}


// SetTimeToAdd sets the value of the timeToAdd property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenControlClusterAddMoreTimeParams/timeToAdd
func (m_ MTRMicrowaveOvenControlClusterAddMoreTimeParams) SetTimeToAdd(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimeToAdd:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenControlClusterAddMoreTimeParams/timedInvokeTimeoutMs
func (m_ MTRMicrowaveOvenControlClusterAddMoreTimeParams) TimedInvokeTimeoutMs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenControlClusterAddMoreTimeParams/timedInvokeTimeoutMs
func (m_ MTRMicrowaveOvenControlClusterAddMoreTimeParams) SetTimedInvokeTimeoutMs(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



