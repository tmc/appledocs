// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRServiceAreaClusterSkipAreaParams] class.
var (
	MTRServiceAreaClusterSkipAreaParamsClass     _MTRServiceAreaClusterSkipAreaParamsClass
	MTRServiceAreaClusterSkipAreaParamsClassOnce sync.Once
)

func getMTRServiceAreaClusterSkipAreaParamsClass() _MTRServiceAreaClusterSkipAreaParamsClass {
	MTRServiceAreaClusterSkipAreaParamsClassOnce.Do(func() {
		MTRServiceAreaClusterSkipAreaParamsClass = _MTRServiceAreaClusterSkipAreaParamsClass{objc.GetClass("MTRServiceAreaClusterSkipAreaParams")}
	})
	return MTRServiceAreaClusterSkipAreaParamsClass
}

type _MTRServiceAreaClusterSkipAreaParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRServiceAreaClusterSkipAreaParams] class.
type IMTRServiceAreaClusterSkipAreaParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSkipAreaParams
type MTRServiceAreaClusterSkipAreaParams struct {
	objectivec.Object
}

// MTRServiceAreaClusterSkipAreaParamsFrom constructs a [MTRServiceAreaClusterSkipAreaParams] from an unsafe.Pointer.
func MTRServiceAreaClusterSkipAreaParamsFrom(ptr unsafe.Pointer) MTRServiceAreaClusterSkipAreaParams {
	return MTRServiceAreaClusterSkipAreaParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRServiceAreaClusterSkipAreaParamsClass) Alloc() MTRServiceAreaClusterSkipAreaParams {
	rv := objc.Send[MTRServiceAreaClusterSkipAreaParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRServiceAreaClusterSkipAreaParamsClass) New() MTRServiceAreaClusterSkipAreaParams {
	rv := objc.Send[MTRServiceAreaClusterSkipAreaParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRServiceAreaClusterSkipAreaParams) Init() MTRServiceAreaClusterSkipAreaParams {
	rv := objc.Send[MTRServiceAreaClusterSkipAreaParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRServiceAreaClusterSkipAreaParams) Autorelease() MTRServiceAreaClusterSkipAreaParams {
	rv := objc.Send[MTRServiceAreaClusterSkipAreaParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRServiceAreaClusterSkipAreaParams creates a new MTRServiceAreaClusterSkipAreaParams instance.
func NewMTRServiceAreaClusterSkipAreaParams() MTRServiceAreaClusterSkipAreaParams {
	return getMTRServiceAreaClusterSkipAreaParamsClass().New()
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSkipAreaParams/serverSideProcessingTimeout
func (m_ MTRServiceAreaClusterSkipAreaParams) ServerSideProcessingTimeout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSkipAreaParams/serverSideProcessingTimeout
func (m_ MTRServiceAreaClusterSkipAreaParams) SetServerSideProcessingTimeout(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSkipAreaParams/skippedArea
func (m_ MTRServiceAreaClusterSkipAreaParams) SkippedArea() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("skippedArea"))
	return rv
}


// SetSkippedArea sets the value of the skippedArea property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSkipAreaParams/skippedArea
func (m_ MTRServiceAreaClusterSkipAreaParams) SetSkippedArea(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSkippedArea:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSkipAreaParams/timedInvokeTimeoutMs
func (m_ MTRServiceAreaClusterSkipAreaParams) TimedInvokeTimeoutMs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSkipAreaParams/timedInvokeTimeoutMs
func (m_ MTRServiceAreaClusterSkipAreaParams) SetTimedInvokeTimeoutMs(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



