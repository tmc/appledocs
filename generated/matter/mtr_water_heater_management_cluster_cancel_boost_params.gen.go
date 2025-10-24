// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterCancelBoostParams/serverSideProcessingTimeout
func (m_ MTRWaterHeaterManagementClusterCancelBoostParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterCancelBoostParams/serverSideProcessingTimeout
func (m_ MTRWaterHeaterManagementClusterCancelBoostParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterCancelBoostParams/timedInvokeTimeoutMs
func (m_ MTRWaterHeaterManagementClusterCancelBoostParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterCancelBoostParams/timedInvokeTimeoutMs
func (m_ MTRWaterHeaterManagementClusterCancelBoostParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



