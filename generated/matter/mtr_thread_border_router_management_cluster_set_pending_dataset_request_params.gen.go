// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams] class.
var (
	MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParamsClass     _MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParamsClass
	MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParamsClassOnce sync.Once
)

func getMTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParamsClass() _MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParamsClass {
	MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParamsClassOnce.Do(func() {
		MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParamsClass = _MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParamsClass{objc.GetClass("MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams")}
	})
	return MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParamsClass
}

type _MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams] class.
type IMTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams
type MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams struct {
	objectivec.Object
}

// MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParamsFrom constructs a [MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams] from an unsafe.Pointer.
func MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParamsFrom(ptr unsafe.Pointer) MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams {
	return MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParamsClass) Alloc() MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams {
	rv := objc.Send[MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParamsClass) New() MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams {
	rv := objc.Send[MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams) Init() MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams {
	rv := objc.Send[MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams) Autorelease() MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams {
	rv := objc.Send[MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams creates a new MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams instance.
func NewMTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams() MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams {
	return getMTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams/pendingDataset
func (m_ MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams) PendingDataset() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("pendingDataset"))
	return rv
}


// SetPendingDataset sets the value of the pendingDataset property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams/pendingDataset
func (m_ MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams) SetPendingDataset(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPendingDataset:"), value)
}

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams/serverSideProcessingTimeout
func (m_ MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams/serverSideProcessingTimeout
func (m_ MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams/timedInvokeTimeoutMs
func (m_ MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams/timedInvokeTimeoutMs
func (m_ MTRThreadBorderRouterManagementClusterSetPendingDatasetRequestParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



