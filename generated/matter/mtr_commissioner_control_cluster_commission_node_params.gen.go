// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRCommissionerControlClusterCommissionNodeParams] class.
var (
	MTRCommissionerControlClusterCommissionNodeParamsClass     _MTRCommissionerControlClusterCommissionNodeParamsClass
	MTRCommissionerControlClusterCommissionNodeParamsClassOnce sync.Once
)

func getMTRCommissionerControlClusterCommissionNodeParamsClass() _MTRCommissionerControlClusterCommissionNodeParamsClass {
	MTRCommissionerControlClusterCommissionNodeParamsClassOnce.Do(func() {
		MTRCommissionerControlClusterCommissionNodeParamsClass = _MTRCommissionerControlClusterCommissionNodeParamsClass{objc.GetClass("MTRCommissionerControlClusterCommissionNodeParams")}
	})
	return MTRCommissionerControlClusterCommissionNodeParamsClass
}

type _MTRCommissionerControlClusterCommissionNodeParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRCommissionerControlClusterCommissionNodeParams] class.
type IMTRCommissionerControlClusterCommissionNodeParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterCommissionNodeParams
type MTRCommissionerControlClusterCommissionNodeParams struct {
	objectivec.Object
}

// MTRCommissionerControlClusterCommissionNodeParamsFrom constructs a [MTRCommissionerControlClusterCommissionNodeParams] from an unsafe.Pointer.
func MTRCommissionerControlClusterCommissionNodeParamsFrom(ptr unsafe.Pointer) MTRCommissionerControlClusterCommissionNodeParams {
	return MTRCommissionerControlClusterCommissionNodeParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRCommissionerControlClusterCommissionNodeParamsClass) Alloc() MTRCommissionerControlClusterCommissionNodeParams {
	rv := objc.Send[MTRCommissionerControlClusterCommissionNodeParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRCommissionerControlClusterCommissionNodeParamsClass) New() MTRCommissionerControlClusterCommissionNodeParams {
	rv := objc.Send[MTRCommissionerControlClusterCommissionNodeParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRCommissionerControlClusterCommissionNodeParams) Init() MTRCommissionerControlClusterCommissionNodeParams {
	rv := objc.Send[MTRCommissionerControlClusterCommissionNodeParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRCommissionerControlClusterCommissionNodeParams) Autorelease() MTRCommissionerControlClusterCommissionNodeParams {
	rv := objc.Send[MTRCommissionerControlClusterCommissionNodeParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRCommissionerControlClusterCommissionNodeParams creates a new MTRCommissionerControlClusterCommissionNodeParams instance.
func NewMTRCommissionerControlClusterCommissionNodeParams() MTRCommissionerControlClusterCommissionNodeParams {
	return getMTRCommissionerControlClusterCommissionNodeParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterCommissionNodeParams/requestID
func (m_ MTRCommissionerControlClusterCommissionNodeParams) RequestID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("requestID"))
	return rv
}


// SetRequestID sets the value of the requestID property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterCommissionNodeParams/requestID
func (m_ MTRCommissionerControlClusterCommissionNodeParams) SetRequestID(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequestID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterCommissionNodeParams/responseTimeoutSeconds
func (m_ MTRCommissionerControlClusterCommissionNodeParams) ResponseTimeoutSeconds() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("responseTimeoutSeconds"))
	return rv
}


// SetResponseTimeoutSeconds sets the value of the responseTimeoutSeconds property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterCommissionNodeParams/responseTimeoutSeconds
func (m_ MTRCommissionerControlClusterCommissionNodeParams) SetResponseTimeoutSeconds(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setResponseTimeoutSeconds:"), value)
}

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterCommissionNodeParams/serverSideProcessingTimeout
func (m_ MTRCommissionerControlClusterCommissionNodeParams) ServerSideProcessingTimeout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterCommissionNodeParams/serverSideProcessingTimeout
func (m_ MTRCommissionerControlClusterCommissionNodeParams) SetServerSideProcessingTimeout(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterCommissionNodeParams/timedInvokeTimeoutMs
func (m_ MTRCommissionerControlClusterCommissionNodeParams) TimedInvokeTimeoutMs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterCommissionNodeParams/timedInvokeTimeoutMs
func (m_ MTRCommissionerControlClusterCommissionNodeParams) SetTimedInvokeTimeoutMs(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



