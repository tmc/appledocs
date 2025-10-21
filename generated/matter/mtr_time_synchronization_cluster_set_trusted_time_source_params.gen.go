// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRTimeSynchronizationClusterSetTrustedTimeSourceParams] class.
var (
	MTRTimeSynchronizationClusterSetTrustedTimeSourceParamsClass     _MTRTimeSynchronizationClusterSetTrustedTimeSourceParamsClass
	MTRTimeSynchronizationClusterSetTrustedTimeSourceParamsClassOnce sync.Once
)

func getMTRTimeSynchronizationClusterSetTrustedTimeSourceParamsClass() _MTRTimeSynchronizationClusterSetTrustedTimeSourceParamsClass {
	MTRTimeSynchronizationClusterSetTrustedTimeSourceParamsClassOnce.Do(func() {
		MTRTimeSynchronizationClusterSetTrustedTimeSourceParamsClass = _MTRTimeSynchronizationClusterSetTrustedTimeSourceParamsClass{objc.GetClass("MTRTimeSynchronizationClusterSetTrustedTimeSourceParams")}
	})
	return MTRTimeSynchronizationClusterSetTrustedTimeSourceParamsClass
}

type _MTRTimeSynchronizationClusterSetTrustedTimeSourceParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTimeSynchronizationClusterSetTrustedTimeSourceParams] class.
type IMTRTimeSynchronizationClusterSetTrustedTimeSourceParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetTrustedTimeSourceParams
type MTRTimeSynchronizationClusterSetTrustedTimeSourceParams struct {
	objectivec.Object
}

// MTRTimeSynchronizationClusterSetTrustedTimeSourceParamsFrom constructs a [MTRTimeSynchronizationClusterSetTrustedTimeSourceParams] from an unsafe.Pointer.
func MTRTimeSynchronizationClusterSetTrustedTimeSourceParamsFrom(ptr unsafe.Pointer) MTRTimeSynchronizationClusterSetTrustedTimeSourceParams {
	return MTRTimeSynchronizationClusterSetTrustedTimeSourceParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTimeSynchronizationClusterSetTrustedTimeSourceParamsClass) Alloc() MTRTimeSynchronizationClusterSetTrustedTimeSourceParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetTrustedTimeSourceParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTimeSynchronizationClusterSetTrustedTimeSourceParamsClass) New() MTRTimeSynchronizationClusterSetTrustedTimeSourceParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetTrustedTimeSourceParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTimeSynchronizationClusterSetTrustedTimeSourceParams) Init() MTRTimeSynchronizationClusterSetTrustedTimeSourceParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetTrustedTimeSourceParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTimeSynchronizationClusterSetTrustedTimeSourceParams) Autorelease() MTRTimeSynchronizationClusterSetTrustedTimeSourceParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetTrustedTimeSourceParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTimeSynchronizationClusterSetTrustedTimeSourceParams creates a new MTRTimeSynchronizationClusterSetTrustedTimeSourceParams instance.
func NewMTRTimeSynchronizationClusterSetTrustedTimeSourceParams() MTRTimeSynchronizationClusterSetTrustedTimeSourceParams {
	return getMTRTimeSynchronizationClusterSetTrustedTimeSourceParamsClass().New()
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetTrustedTimeSourceParams/serverSideProcessingTimeout
func (m_ MTRTimeSynchronizationClusterSetTrustedTimeSourceParams) ServerSideProcessingTimeout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetTrustedTimeSourceParams/serverSideProcessingTimeout
func (m_ MTRTimeSynchronizationClusterSetTrustedTimeSourceParams) SetServerSideProcessingTimeout(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetTrustedTimeSourceParams/timedInvokeTimeoutMs
func (m_ MTRTimeSynchronizationClusterSetTrustedTimeSourceParams) TimedInvokeTimeoutMs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetTrustedTimeSourceParams/timedInvokeTimeoutMs
func (m_ MTRTimeSynchronizationClusterSetTrustedTimeSourceParams) SetTimedInvokeTimeoutMs(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetTrustedTimeSourceParams/trustedTimeSource
func (m_ MTRTimeSynchronizationClusterSetTrustedTimeSourceParams) TrustedTimeSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("trustedTimeSource"))
	return rv
}


// SetTrustedTimeSource sets the value of the trustedTimeSource property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetTrustedTimeSourceParams/trustedTimeSource
func (m_ MTRTimeSynchronizationClusterSetTrustedTimeSourceParams) SetTrustedTimeSource(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTrustedTimeSource:"), value)
}



