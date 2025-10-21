// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRTimeSynchronizationClusterSetDefaultNTPParams] class.
var (
	MTRTimeSynchronizationClusterSetDefaultNTPParamsClass     _MTRTimeSynchronizationClusterSetDefaultNTPParamsClass
	MTRTimeSynchronizationClusterSetDefaultNTPParamsClassOnce sync.Once
)

func getMTRTimeSynchronizationClusterSetDefaultNTPParamsClass() _MTRTimeSynchronizationClusterSetDefaultNTPParamsClass {
	MTRTimeSynchronizationClusterSetDefaultNTPParamsClassOnce.Do(func() {
		MTRTimeSynchronizationClusterSetDefaultNTPParamsClass = _MTRTimeSynchronizationClusterSetDefaultNTPParamsClass{objc.GetClass("MTRTimeSynchronizationClusterSetDefaultNTPParams")}
	})
	return MTRTimeSynchronizationClusterSetDefaultNTPParamsClass
}

type _MTRTimeSynchronizationClusterSetDefaultNTPParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTimeSynchronizationClusterSetDefaultNTPParams] class.
type IMTRTimeSynchronizationClusterSetDefaultNTPParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetDefaultNTPParams
type MTRTimeSynchronizationClusterSetDefaultNTPParams struct {
	objectivec.Object
}

// MTRTimeSynchronizationClusterSetDefaultNTPParamsFrom constructs a [MTRTimeSynchronizationClusterSetDefaultNTPParams] from an unsafe.Pointer.
func MTRTimeSynchronizationClusterSetDefaultNTPParamsFrom(ptr unsafe.Pointer) MTRTimeSynchronizationClusterSetDefaultNTPParams {
	return MTRTimeSynchronizationClusterSetDefaultNTPParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTimeSynchronizationClusterSetDefaultNTPParamsClass) Alloc() MTRTimeSynchronizationClusterSetDefaultNTPParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetDefaultNTPParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTimeSynchronizationClusterSetDefaultNTPParamsClass) New() MTRTimeSynchronizationClusterSetDefaultNTPParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetDefaultNTPParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTimeSynchronizationClusterSetDefaultNTPParams) Init() MTRTimeSynchronizationClusterSetDefaultNTPParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetDefaultNTPParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTimeSynchronizationClusterSetDefaultNTPParams) Autorelease() MTRTimeSynchronizationClusterSetDefaultNTPParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetDefaultNTPParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTimeSynchronizationClusterSetDefaultNTPParams creates a new MTRTimeSynchronizationClusterSetDefaultNTPParams instance.
func NewMTRTimeSynchronizationClusterSetDefaultNTPParams() MTRTimeSynchronizationClusterSetDefaultNTPParams {
	return getMTRTimeSynchronizationClusterSetDefaultNTPParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetDefaultNTPParams/defaultNTP
func (m_ MTRTimeSynchronizationClusterSetDefaultNTPParams) DefaultNTP() string {
	rv := objc.Send[string](m_.ID, objc.Sel("defaultNTP"))
	return rv
}


// SetDefaultNTP sets the value of the defaultNTP property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetDefaultNTPParams/defaultNTP
func (m_ MTRTimeSynchronizationClusterSetDefaultNTPParams) SetDefaultNTP(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDefaultNTP:"), objc.String(value))
}

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetDefaultNTPParams/serverSideProcessingTimeout
func (m_ MTRTimeSynchronizationClusterSetDefaultNTPParams) ServerSideProcessingTimeout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetDefaultNTPParams/serverSideProcessingTimeout
func (m_ MTRTimeSynchronizationClusterSetDefaultNTPParams) SetServerSideProcessingTimeout(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetDefaultNTPParams/timedInvokeTimeoutMs
func (m_ MTRTimeSynchronizationClusterSetDefaultNTPParams) TimedInvokeTimeoutMs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetDefaultNTPParams/timedInvokeTimeoutMs
func (m_ MTRTimeSynchronizationClusterSetDefaultNTPParams) SetTimedInvokeTimeoutMs(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



