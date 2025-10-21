// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRWaterHeaterManagementClusterCancelBoostParams] class.
var (
	MTRWaterHeaterManagementClusterCancelBoostParamsClass     _MTRWaterHeaterManagementClusterCancelBoostParamsClass
	MTRWaterHeaterManagementClusterCancelBoostParamsClassOnce sync.Once
)

func getMTRWaterHeaterManagementClusterCancelBoostParamsClass() _MTRWaterHeaterManagementClusterCancelBoostParamsClass {
	MTRWaterHeaterManagementClusterCancelBoostParamsClassOnce.Do(func() {
		MTRWaterHeaterManagementClusterCancelBoostParamsClass = _MTRWaterHeaterManagementClusterCancelBoostParamsClass{objc.GetClass("MTRWaterHeaterManagementClusterCancelBoostParams")}
	})
	return MTRWaterHeaterManagementClusterCancelBoostParamsClass
}

type _MTRWaterHeaterManagementClusterCancelBoostParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRWaterHeaterManagementClusterCancelBoostParams] class.
type IMTRWaterHeaterManagementClusterCancelBoostParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterCancelBoostParams
type MTRWaterHeaterManagementClusterCancelBoostParams struct {
	objectivec.Object
}

// MTRWaterHeaterManagementClusterCancelBoostParamsFrom constructs a [MTRWaterHeaterManagementClusterCancelBoostParams] from an unsafe.Pointer.
func MTRWaterHeaterManagementClusterCancelBoostParamsFrom(ptr unsafe.Pointer) MTRWaterHeaterManagementClusterCancelBoostParams {
	return MTRWaterHeaterManagementClusterCancelBoostParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRWaterHeaterManagementClusterCancelBoostParamsClass) Alloc() MTRWaterHeaterManagementClusterCancelBoostParams {
	rv := objc.Send[MTRWaterHeaterManagementClusterCancelBoostParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRWaterHeaterManagementClusterCancelBoostParamsClass) New() MTRWaterHeaterManagementClusterCancelBoostParams {
	rv := objc.Send[MTRWaterHeaterManagementClusterCancelBoostParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWaterHeaterManagementClusterCancelBoostParams) Init() MTRWaterHeaterManagementClusterCancelBoostParams {
	rv := objc.Send[MTRWaterHeaterManagementClusterCancelBoostParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWaterHeaterManagementClusterCancelBoostParams) Autorelease() MTRWaterHeaterManagementClusterCancelBoostParams {
	rv := objc.Send[MTRWaterHeaterManagementClusterCancelBoostParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWaterHeaterManagementClusterCancelBoostParams creates a new MTRWaterHeaterManagementClusterCancelBoostParams instance.
func NewMTRWaterHeaterManagementClusterCancelBoostParams() MTRWaterHeaterManagementClusterCancelBoostParams {
	return getMTRWaterHeaterManagementClusterCancelBoostParamsClass().New()
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterCancelBoostParams/serverSideProcessingTimeout
func (m_ MTRWaterHeaterManagementClusterCancelBoostParams) ServerSideProcessingTimeout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterCancelBoostParams/serverSideProcessingTimeout
func (m_ MTRWaterHeaterManagementClusterCancelBoostParams) SetServerSideProcessingTimeout(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}
// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterCancelBoostParams/timedInvokeTimeoutMs
func (m_ MTRWaterHeaterManagementClusterCancelBoostParams) TimedInvokeTimeoutMs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterCancelBoostParams/timedInvokeTimeoutMs
func (m_ MTRWaterHeaterManagementClusterCancelBoostParams) SetTimedInvokeTimeoutMs(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


