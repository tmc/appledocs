// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRDoorLockClusterClearAliroReaderConfigParams] class.
var (
	MTRDoorLockClusterClearAliroReaderConfigParamsClass     _MTRDoorLockClusterClearAliroReaderConfigParamsClass
	MTRDoorLockClusterClearAliroReaderConfigParamsClassOnce sync.Once
)

func getMTRDoorLockClusterClearAliroReaderConfigParamsClass() _MTRDoorLockClusterClearAliroReaderConfigParamsClass {
	MTRDoorLockClusterClearAliroReaderConfigParamsClassOnce.Do(func() {
		MTRDoorLockClusterClearAliroReaderConfigParamsClass = _MTRDoorLockClusterClearAliroReaderConfigParamsClass{objc.GetClass("MTRDoorLockClusterClearAliroReaderConfigParams")}
	})
	return MTRDoorLockClusterClearAliroReaderConfigParamsClass
}

type _MTRDoorLockClusterClearAliroReaderConfigParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterClearAliroReaderConfigParams] class.
type IMTRDoorLockClusterClearAliroReaderConfigParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearAliroReaderConfigParams
type MTRDoorLockClusterClearAliroReaderConfigParams struct {
	objectivec.Object
}

// MTRDoorLockClusterClearAliroReaderConfigParamsFrom constructs a [MTRDoorLockClusterClearAliroReaderConfigParams] from an unsafe.Pointer.
func MTRDoorLockClusterClearAliroReaderConfigParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterClearAliroReaderConfigParams {
	return MTRDoorLockClusterClearAliroReaderConfigParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterClearAliroReaderConfigParamsClass) Alloc() MTRDoorLockClusterClearAliroReaderConfigParams {
	rv := objc.Send[MTRDoorLockClusterClearAliroReaderConfigParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterClearAliroReaderConfigParamsClass) New() MTRDoorLockClusterClearAliroReaderConfigParams {
	rv := objc.Send[MTRDoorLockClusterClearAliroReaderConfigParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterClearAliroReaderConfigParams) Init() MTRDoorLockClusterClearAliroReaderConfigParams {
	rv := objc.Send[MTRDoorLockClusterClearAliroReaderConfigParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterClearAliroReaderConfigParams) Autorelease() MTRDoorLockClusterClearAliroReaderConfigParams {
	rv := objc.Send[MTRDoorLockClusterClearAliroReaderConfigParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterClearAliroReaderConfigParams creates a new MTRDoorLockClusterClearAliroReaderConfigParams instance.
func NewMTRDoorLockClusterClearAliroReaderConfigParams() MTRDoorLockClusterClearAliroReaderConfigParams {
	return getMTRDoorLockClusterClearAliroReaderConfigParamsClass().New()
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearAliroReaderConfigParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterClearAliroReaderConfigParams) ServerSideProcessingTimeout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearAliroReaderConfigParams/serverSideProcessingTimeout
func (m_ MTRDoorLockClusterClearAliroReaderConfigParams) SetServerSideProcessingTimeout(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}
// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearAliroReaderConfigParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterClearAliroReaderConfigParams) TimedInvokeTimeoutMs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearAliroReaderConfigParams/timedInvokeTimeoutMs
func (m_ MTRDoorLockClusterClearAliroReaderConfigParams) SetTimedInvokeTimeoutMs(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


