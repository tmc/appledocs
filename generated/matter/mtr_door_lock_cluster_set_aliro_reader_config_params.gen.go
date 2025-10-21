// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDoorLockClusterSetAliroReaderConfigParams] class.
var (
	MTRDoorLockClusterSetAliroReaderConfigParamsClass     _MTRDoorLockClusterSetAliroReaderConfigParamsClass
	MTRDoorLockClusterSetAliroReaderConfigParamsClassOnce sync.Once
)

func getMTRDoorLockClusterSetAliroReaderConfigParamsClass() _MTRDoorLockClusterSetAliroReaderConfigParamsClass {
	MTRDoorLockClusterSetAliroReaderConfigParamsClassOnce.Do(func() {
		MTRDoorLockClusterSetAliroReaderConfigParamsClass = _MTRDoorLockClusterSetAliroReaderConfigParamsClass{objc.GetClass("MTRDoorLockClusterSetAliroReaderConfigParams")}
	})
	return MTRDoorLockClusterSetAliroReaderConfigParamsClass
}

type _MTRDoorLockClusterSetAliroReaderConfigParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterSetAliroReaderConfigParams] class.
type IMTRDoorLockClusterSetAliroReaderConfigParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetAliroReaderConfigParams
type MTRDoorLockClusterSetAliroReaderConfigParams struct {
	objectivec.Object
}

// MTRDoorLockClusterSetAliroReaderConfigParamsFrom constructs a [MTRDoorLockClusterSetAliroReaderConfigParams] from an unsafe.Pointer.
func MTRDoorLockClusterSetAliroReaderConfigParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterSetAliroReaderConfigParams {
	return MTRDoorLockClusterSetAliroReaderConfigParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterSetAliroReaderConfigParamsClass) Alloc() MTRDoorLockClusterSetAliroReaderConfigParams {
	rv := objc.Send[MTRDoorLockClusterSetAliroReaderConfigParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterSetAliroReaderConfigParamsClass) New() MTRDoorLockClusterSetAliroReaderConfigParams {
	rv := objc.Send[MTRDoorLockClusterSetAliroReaderConfigParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterSetAliroReaderConfigParams) Init() MTRDoorLockClusterSetAliroReaderConfigParams {
	rv := objc.Send[MTRDoorLockClusterSetAliroReaderConfigParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterSetAliroReaderConfigParams) Autorelease() MTRDoorLockClusterSetAliroReaderConfigParams {
	rv := objc.Send[MTRDoorLockClusterSetAliroReaderConfigParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterSetAliroReaderConfigParams creates a new MTRDoorLockClusterSetAliroReaderConfigParams instance.
func NewMTRDoorLockClusterSetAliroReaderConfigParams() MTRDoorLockClusterSetAliroReaderConfigParams {
	return getMTRDoorLockClusterSetAliroReaderConfigParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetAliroReaderConfigParams/groupIdentifier
func (m_ MTRDoorLockClusterSetAliroReaderConfigParams) GroupIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("groupIdentifier"))
	return rv
}


// SetGroupIdentifier sets the value of the groupIdentifier property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetAliroReaderConfigParams/groupIdentifier
func (m_ MTRDoorLockClusterSetAliroReaderConfigParams) SetGroupIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupIdentifier:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetAliroReaderConfigParams/groupResolvingKey
func (m_ MTRDoorLockClusterSetAliroReaderConfigParams) GroupResolvingKey() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("groupResolvingKey"))
	return rv
}


// SetGroupResolvingKey sets the value of the groupResolvingKey property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetAliroReaderConfigParams/groupResolvingKey
func (m_ MTRDoorLockClusterSetAliroReaderConfigParams) SetGroupResolvingKey(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupResolvingKey:"), value)
}

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetAliroReaderConfigParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterSetAliroReaderConfigParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetAliroReaderConfigParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterSetAliroReaderConfigParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetAliroReaderConfigParams/signingKey
func (m_ MTRDoorLockClusterSetAliroReaderConfigParams) SigningKey() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("signingKey"))
	return rv
}


// SetSigningKey sets the value of the signingKey property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetAliroReaderConfigParams/signingKey
func (m_ MTRDoorLockClusterSetAliroReaderConfigParams) SetSigningKey(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSigningKey:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetAliroReaderConfigParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterSetAliroReaderConfigParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetAliroReaderConfigParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterSetAliroReaderConfigParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetAliroReaderConfigParams/verificationKey
func (m_ MTRDoorLockClusterSetAliroReaderConfigParams) VerificationKey() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("verificationKey"))
	return rv
}


// SetVerificationKey sets the value of the verificationKey property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetAliroReaderConfigParams/verificationKey
func (m_ MTRDoorLockClusterSetAliroReaderConfigParams) SetVerificationKey(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVerificationKey:"), value)
}



