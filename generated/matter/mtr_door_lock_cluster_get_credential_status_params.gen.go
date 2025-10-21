// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDoorLockClusterGetCredentialStatusParams] class.
var (
	MTRDoorLockClusterGetCredentialStatusParamsClass     _MTRDoorLockClusterGetCredentialStatusParamsClass
	MTRDoorLockClusterGetCredentialStatusParamsClassOnce sync.Once
)

func getMTRDoorLockClusterGetCredentialStatusParamsClass() _MTRDoorLockClusterGetCredentialStatusParamsClass {
	MTRDoorLockClusterGetCredentialStatusParamsClassOnce.Do(func() {
		MTRDoorLockClusterGetCredentialStatusParamsClass = _MTRDoorLockClusterGetCredentialStatusParamsClass{objc.GetClass("MTRDoorLockClusterGetCredentialStatusParams")}
	})
	return MTRDoorLockClusterGetCredentialStatusParamsClass
}

type _MTRDoorLockClusterGetCredentialStatusParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterGetCredentialStatusParams] class.
type IMTRDoorLockClusterGetCredentialStatusParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetCredentialStatusParams
type MTRDoorLockClusterGetCredentialStatusParams struct {
	objectivec.Object
}

// MTRDoorLockClusterGetCredentialStatusParamsFrom constructs a [MTRDoorLockClusterGetCredentialStatusParams] from an unsafe.Pointer.
func MTRDoorLockClusterGetCredentialStatusParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterGetCredentialStatusParams {
	return MTRDoorLockClusterGetCredentialStatusParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterGetCredentialStatusParamsClass) Alloc() MTRDoorLockClusterGetCredentialStatusParams {
	rv := objc.Send[MTRDoorLockClusterGetCredentialStatusParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterGetCredentialStatusParamsClass) New() MTRDoorLockClusterGetCredentialStatusParams {
	rv := objc.Send[MTRDoorLockClusterGetCredentialStatusParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterGetCredentialStatusParams) Init() MTRDoorLockClusterGetCredentialStatusParams {
	rv := objc.Send[MTRDoorLockClusterGetCredentialStatusParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterGetCredentialStatusParams) Autorelease() MTRDoorLockClusterGetCredentialStatusParams {
	rv := objc.Send[MTRDoorLockClusterGetCredentialStatusParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterGetCredentialStatusParams creates a new MTRDoorLockClusterGetCredentialStatusParams instance.
func NewMTRDoorLockClusterGetCredentialStatusParams() MTRDoorLockClusterGetCredentialStatusParams {
	return getMTRDoorLockClusterGetCredentialStatusParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetcredentialstatusparams/credential
func (m_ MTRDoorLockClusterGetCredentialStatusParams) Credential() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("credential"))
	return rv
}


// SetCredential sets the value of the credential property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetcredentialstatusparams/credential
func (m_ MTRDoorLockClusterGetCredentialStatusParams) SetCredential(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCredential:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetcredentialstatusparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterGetCredentialStatusParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetcredentialstatusparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterGetCredentialStatusParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetcredentialstatusparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterGetCredentialStatusParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetcredentialstatusparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterGetCredentialStatusParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



