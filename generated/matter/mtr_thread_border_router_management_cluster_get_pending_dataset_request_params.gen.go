// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams] class.
var (
	MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParamsClass     _MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParamsClass
	MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParamsClassOnce sync.Once
)

func getMTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParamsClass() _MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParamsClass {
	MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParamsClassOnce.Do(func() {
		MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParamsClass = _MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParamsClass{objc.GetClass("MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams")}
	})
	return MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParamsClass
}

type _MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams] class.
type IMTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams
type MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams struct {
	objectivec.Object
}

// MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParamsFrom constructs a [MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams] from an unsafe.Pointer.
func MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParamsFrom(ptr unsafe.Pointer) MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams {
	return MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParamsClass) Alloc() MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams {
	rv := objc.Send[MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParamsClass) New() MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams {
	rv := objc.Send[MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams) Init() MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams {
	rv := objc.Send[MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams) Autorelease() MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams {
	rv := objc.Send[MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams creates a new MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams instance.
func NewMTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams() MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams {
	return getMTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParamsClass().New()
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams/serverSideProcessingTimeout
func (m_ MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams/serverSideProcessingTimeout
func (m_ MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams/timedInvokeTimeoutMs
func (m_ MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams/timedInvokeTimeoutMs
func (m_ MTRThreadBorderRouterManagementClusterGetPendingDatasetRequestParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



