// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	CheckInNodeID() objc.IObject /* cross-framework: NSNumber */
	SetCheckInNodeID(value objc.IObject /* cross-framework: NSNumber */)
	ClientType() objc.IObject /* cross-framework: NSNumber */
	SetClientType(value objc.IObject /* cross-framework: NSNumber */)
	Key() objc.IObject /* cross-framework: NSData */
	SetKey(value objc.IObject /* cross-framework: NSData */)
	MonitoredSubject() objc.IObject /* cross-framework: NSNumber */
	SetMonitoredSubject(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	VerificationKey() objc.IObject /* cross-framework: NSData */
	SetVerificationKey(value objc.IObject /* cross-framework: NSData */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterRegisterClientParams/checkInNodeID
func (m_ MTRICDManagementClusterRegisterClientParams) CheckInNodeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("checkInNodeID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterRegisterClientParams/checkInNodeID
func (m_ MTRICDManagementClusterRegisterClientParams) SetCheckInNodeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCheckInNodeID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterRegisterClientParams/clientType
func (m_ MTRICDManagementClusterRegisterClientParams) ClientType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("clientType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterRegisterClientParams/clientType
func (m_ MTRICDManagementClusterRegisterClientParams) SetClientType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setClientType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterRegisterClientParams/key
func (m_ MTRICDManagementClusterRegisterClientParams) Key() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("key"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterRegisterClientParams/key
func (m_ MTRICDManagementClusterRegisterClientParams) SetKey(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setKey:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterRegisterClientParams/monitoredSubject
func (m_ MTRICDManagementClusterRegisterClientParams) MonitoredSubject() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("monitoredSubject"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterRegisterClientParams/monitoredSubject
func (m_ MTRICDManagementClusterRegisterClientParams) SetMonitoredSubject(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMonitoredSubject:"), value)
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterRegisterClientParams/serverSideProcessingTimeout
func (m_ MTRICDManagementClusterRegisterClientParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterRegisterClientParams/serverSideProcessingTimeout
func (m_ MTRICDManagementClusterRegisterClientParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterRegisterClientParams/timedInvokeTimeoutMs
func (m_ MTRICDManagementClusterRegisterClientParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterRegisterClientParams/timedInvokeTimeoutMs
func (m_ MTRICDManagementClusterRegisterClientParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterRegisterClientParams/verificationKey
func (m_ MTRICDManagementClusterRegisterClientParams) VerificationKey() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("verificationKey"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterRegisterClientParams/verificationKey
func (m_ MTRICDManagementClusterRegisterClientParams) SetVerificationKey(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVerificationKey:"), value)
}



