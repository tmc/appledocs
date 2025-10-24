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
	// properties:
	GroupIdentifier() objc.IObject /* cross-framework: NSData */
	SetGroupIdentifier(value objc.IObject /* cross-framework: NSData */)
	GroupResolvingKey() objc.IObject /* cross-framework: NSData */
	SetGroupResolvingKey(value objc.IObject /* cross-framework: NSData */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	SigningKey() objc.IObject /* cross-framework: NSData */
	SetSigningKey(value objc.IObject /* cross-framework: NSData */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	VerificationKey() objc.IObject /* cross-framework: NSData */
	SetVerificationKey(value objc.IObject /* cross-framework: NSData */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetAliroReaderConfigParams/groupIdentifier
func (m_ MTRDoorLockClusterSetAliroReaderConfigParams) GroupIdentifier() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("groupIdentifier"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetAliroReaderConfigParams/groupIdentifier
func (m_ MTRDoorLockClusterSetAliroReaderConfigParams) SetGroupIdentifier(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupIdentifier:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetAliroReaderConfigParams/groupResolvingKey
func (m_ MTRDoorLockClusterSetAliroReaderConfigParams) GroupResolvingKey() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("groupResolvingKey"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetAliroReaderConfigParams/groupResolvingKey
func (m_ MTRDoorLockClusterSetAliroReaderConfigParams) SetGroupResolvingKey(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupResolvingKey:"), value)
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetAliroReaderConfigParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterSetAliroReaderConfigParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetAliroReaderConfigParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterSetAliroReaderConfigParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetAliroReaderConfigParams/signingKey
func (m_ MTRDoorLockClusterSetAliroReaderConfigParams) SigningKey() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("signingKey"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetAliroReaderConfigParams/signingKey
func (m_ MTRDoorLockClusterSetAliroReaderConfigParams) SetSigningKey(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSigningKey:"), value)
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetAliroReaderConfigParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterSetAliroReaderConfigParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetAliroReaderConfigParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterSetAliroReaderConfigParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetAliroReaderConfigParams/verificationKey
func (m_ MTRDoorLockClusterSetAliroReaderConfigParams) VerificationKey() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("verificationKey"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetAliroReaderConfigParams/verificationKey
func (m_ MTRDoorLockClusterSetAliroReaderConfigParams) SetVerificationKey(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVerificationKey:"), value)
}



