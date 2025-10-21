// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRICDManagementClusterUnregisterClientParams] class.
var (
	MTRICDManagementClusterUnregisterClientParamsClass     _MTRICDManagementClusterUnregisterClientParamsClass
	MTRICDManagementClusterUnregisterClientParamsClassOnce sync.Once
)

func getMTRICDManagementClusterUnregisterClientParamsClass() _MTRICDManagementClusterUnregisterClientParamsClass {
	MTRICDManagementClusterUnregisterClientParamsClassOnce.Do(func() {
		MTRICDManagementClusterUnregisterClientParamsClass = _MTRICDManagementClusterUnregisterClientParamsClass{objc.GetClass("MTRICDManagementClusterUnregisterClientParams")}
	})
	return MTRICDManagementClusterUnregisterClientParamsClass
}

type _MTRICDManagementClusterUnregisterClientParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRICDManagementClusterUnregisterClientParams] class.
type IMTRICDManagementClusterUnregisterClientParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterUnregisterClientParams
type MTRICDManagementClusterUnregisterClientParams struct {
	objectivec.Object
}

// MTRICDManagementClusterUnregisterClientParamsFrom constructs a [MTRICDManagementClusterUnregisterClientParams] from an unsafe.Pointer.
func MTRICDManagementClusterUnregisterClientParamsFrom(ptr unsafe.Pointer) MTRICDManagementClusterUnregisterClientParams {
	return MTRICDManagementClusterUnregisterClientParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRICDManagementClusterUnregisterClientParamsClass) Alloc() MTRICDManagementClusterUnregisterClientParams {
	rv := objc.Send[MTRICDManagementClusterUnregisterClientParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRICDManagementClusterUnregisterClientParamsClass) New() MTRICDManagementClusterUnregisterClientParams {
	rv := objc.Send[MTRICDManagementClusterUnregisterClientParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRICDManagementClusterUnregisterClientParams) Init() MTRICDManagementClusterUnregisterClientParams {
	rv := objc.Send[MTRICDManagementClusterUnregisterClientParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRICDManagementClusterUnregisterClientParams) Autorelease() MTRICDManagementClusterUnregisterClientParams {
	rv := objc.Send[MTRICDManagementClusterUnregisterClientParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRICDManagementClusterUnregisterClientParams creates a new MTRICDManagementClusterUnregisterClientParams instance.
func NewMTRICDManagementClusterUnregisterClientParams() MTRICDManagementClusterUnregisterClientParams {
	return getMTRICDManagementClusterUnregisterClientParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterUnregisterClientParams/checkInNodeID
func (m_ MTRICDManagementClusterUnregisterClientParams) CheckInNodeID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("checkInNodeID"))
	return rv
}


// SetCheckInNodeID sets the value of the checkInNodeID property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterUnregisterClientParams/checkInNodeID
func (m_ MTRICDManagementClusterUnregisterClientParams) SetCheckInNodeID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCheckInNodeID:"), value)
}

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterUnregisterClientParams/serverSideProcessingTimeout
func (m_ MTRICDManagementClusterUnregisterClientParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterUnregisterClientParams/serverSideProcessingTimeout
func (m_ MTRICDManagementClusterUnregisterClientParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterUnregisterClientParams/timedInvokeTimeoutMs
func (m_ MTRICDManagementClusterUnregisterClientParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterUnregisterClientParams/timedInvokeTimeoutMs
func (m_ MTRICDManagementClusterUnregisterClientParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterUnregisterClientParams/verificationKey
func (m_ MTRICDManagementClusterUnregisterClientParams) VerificationKey() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("verificationKey"))
	return rv
}


// SetVerificationKey sets the value of the verificationKey property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterUnregisterClientParams/verificationKey
func (m_ MTRICDManagementClusterUnregisterClientParams) SetVerificationKey(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVerificationKey:"), value)
}



