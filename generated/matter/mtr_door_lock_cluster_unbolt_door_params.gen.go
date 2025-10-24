// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDoorLockClusterUnboltDoorParams] class.
var (
	MTRDoorLockClusterUnboltDoorParamsClass     _MTRDoorLockClusterUnboltDoorParamsClass
	MTRDoorLockClusterUnboltDoorParamsClassOnce sync.Once
)

func getMTRDoorLockClusterUnboltDoorParamsClass() _MTRDoorLockClusterUnboltDoorParamsClass {
	MTRDoorLockClusterUnboltDoorParamsClassOnce.Do(func() {
		MTRDoorLockClusterUnboltDoorParamsClass = _MTRDoorLockClusterUnboltDoorParamsClass{objc.GetClass("MTRDoorLockClusterUnboltDoorParams")}
	})
	return MTRDoorLockClusterUnboltDoorParamsClass
}

type _MTRDoorLockClusterUnboltDoorParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterUnboltDoorParams] class.
type IMTRDoorLockClusterUnboltDoorParams interface {
	objectivec.IObject
	// properties:
	PinCode() objc.IObject /* cross-framework: NSData */
	SetPinCode(value objc.IObject /* cross-framework: NSData */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterUnboltDoorParams
type MTRDoorLockClusterUnboltDoorParams struct {
	objectivec.Object
}

// MTRDoorLockClusterUnboltDoorParamsFrom constructs a [MTRDoorLockClusterUnboltDoorParams] from an unsafe.Pointer.
func MTRDoorLockClusterUnboltDoorParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterUnboltDoorParams {
	return MTRDoorLockClusterUnboltDoorParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterUnboltDoorParamsClass) Alloc() MTRDoorLockClusterUnboltDoorParams {
	rv := objc.Send[MTRDoorLockClusterUnboltDoorParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterUnboltDoorParamsClass) New() MTRDoorLockClusterUnboltDoorParams {
	rv := objc.Send[MTRDoorLockClusterUnboltDoorParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterUnboltDoorParams) Init() MTRDoorLockClusterUnboltDoorParams {
	rv := objc.Send[MTRDoorLockClusterUnboltDoorParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterUnboltDoorParams) Autorelease() MTRDoorLockClusterUnboltDoorParams {
	rv := objc.Send[MTRDoorLockClusterUnboltDoorParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterUnboltDoorParams creates a new MTRDoorLockClusterUnboltDoorParams instance.
func NewMTRDoorLockClusterUnboltDoorParams() MTRDoorLockClusterUnboltDoorParams {
	return getMTRDoorLockClusterUnboltDoorParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterUnboltDoorParams/pinCode
func (m_ MTRDoorLockClusterUnboltDoorParams) PinCode() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("pinCode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterUnboltDoorParams/pinCode
func (m_ MTRDoorLockClusterUnboltDoorParams) SetPinCode(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPinCode:"), value)
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterUnboltDoorParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterUnboltDoorParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterUnboltDoorParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterUnboltDoorParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterUnboltDoorParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterUnboltDoorParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterUnboltDoorParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterUnboltDoorParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



