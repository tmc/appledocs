// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRICDManagementClusterRegisterClientParams] class.
var (
	MTRICDManagementClusterRegisterClientParamsClass     _MTRICDManagementClusterRegisterClientParamsClass
	MTRICDManagementClusterRegisterClientParamsClassOnce sync.Once
)

func getMTRICDManagementClusterRegisterClientParamsClass() _MTRICDManagementClusterRegisterClientParamsClass {
	MTRICDManagementClusterRegisterClientParamsClassOnce.Do(func() {
		MTRICDManagementClusterRegisterClientParamsClass = _MTRICDManagementClusterRegisterClientParamsClass{objc.GetClass("MTRICDManagementClusterRegisterClientParams")}
	})
	return MTRICDManagementClusterRegisterClientParamsClass
}

type _MTRICDManagementClusterRegisterClientParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRICDManagementClusterRegisterClientParams] class.
type IMTRICDManagementClusterRegisterClientParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterRegisterClientParams
type MTRICDManagementClusterRegisterClientParams struct {
	objectivec.Object
}

// MTRICDManagementClusterRegisterClientParamsFrom constructs a [MTRICDManagementClusterRegisterClientParams] from an unsafe.Pointer.
func MTRICDManagementClusterRegisterClientParamsFrom(ptr unsafe.Pointer) MTRICDManagementClusterRegisterClientParams {
	return MTRICDManagementClusterRegisterClientParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRICDManagementClusterRegisterClientParamsClass) Alloc() MTRICDManagementClusterRegisterClientParams {
	rv := objc.Send[MTRICDManagementClusterRegisterClientParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRICDManagementClusterRegisterClientParamsClass) New() MTRICDManagementClusterRegisterClientParams {
	rv := objc.Send[MTRICDManagementClusterRegisterClientParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRICDManagementClusterRegisterClientParams) Init() MTRICDManagementClusterRegisterClientParams {
	rv := objc.Send[MTRICDManagementClusterRegisterClientParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRICDManagementClusterRegisterClientParams) Autorelease() MTRICDManagementClusterRegisterClientParams {
	rv := objc.Send[MTRICDManagementClusterRegisterClientParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRICDManagementClusterRegisterClientParams creates a new MTRICDManagementClusterRegisterClientParams instance.
func NewMTRICDManagementClusterRegisterClientParams() MTRICDManagementClusterRegisterClientParams {
	return getMTRICDManagementClusterRegisterClientParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterRegisterClientParams/checkInNodeID
func (m_ MTRICDManagementClusterRegisterClientParams) CheckInNodeID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("checkInNodeID"))
	return rv
}


// SetCheckInNodeID sets the value of the checkInNodeID property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterRegisterClientParams/checkInNodeID
func (m_ MTRICDManagementClusterRegisterClientParams) SetCheckInNodeID(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCheckInNodeID:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterRegisterClientParams/clientType
func (m_ MTRICDManagementClusterRegisterClientParams) ClientType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("clientType"))
	return rv
}


// SetClientType sets the value of the clientType property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterRegisterClientParams/clientType
func (m_ MTRICDManagementClusterRegisterClientParams) SetClientType(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setClientType:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterRegisterClientParams/key
func (m_ MTRICDManagementClusterRegisterClientParams) Key() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("key"))
	return rv
}


// SetKey sets the value of the key property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterRegisterClientParams/key
func (m_ MTRICDManagementClusterRegisterClientParams) SetKey(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setKey:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterRegisterClientParams/monitoredSubject
func (m_ MTRICDManagementClusterRegisterClientParams) MonitoredSubject() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("monitoredSubject"))
	return rv
}


// SetMonitoredSubject sets the value of the monitoredSubject property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterRegisterClientParams/monitoredSubject
func (m_ MTRICDManagementClusterRegisterClientParams) SetMonitoredSubject(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMonitoredSubject:"), value)
}
// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterRegisterClientParams/serverSideProcessingTimeout
func (m_ MTRICDManagementClusterRegisterClientParams) ServerSideProcessingTimeout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterRegisterClientParams/serverSideProcessingTimeout
func (m_ MTRICDManagementClusterRegisterClientParams) SetServerSideProcessingTimeout(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}
// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterRegisterClientParams/timedInvokeTimeoutMs
func (m_ MTRICDManagementClusterRegisterClientParams) TimedInvokeTimeoutMs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterRegisterClientParams/timedInvokeTimeoutMs
func (m_ MTRICDManagementClusterRegisterClientParams) SetTimedInvokeTimeoutMs(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterRegisterClientParams/verificationKey
func (m_ MTRICDManagementClusterRegisterClientParams) VerificationKey() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("verificationKey"))
	return rv
}


// SetVerificationKey sets the value of the verificationKey property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterRegisterClientParams/verificationKey
func (m_ MTRICDManagementClusterRegisterClientParams) SetVerificationKey(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVerificationKey:"), value)
}


