// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams] class.
var (
	MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParamsClass     _MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParamsClass
	MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParamsClassOnce sync.Once
)

func getMTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParamsClass() _MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParamsClass {
	MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParamsClassOnce.Do(func() {
		MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParamsClass = _MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParamsClass{objc.GetClass("MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams")}
	})
	return MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParamsClass
}

type _MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams] class.
type IMTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams
type MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams struct {
	objectivec.Object
}

// MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParamsFrom constructs a [MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams] from an unsafe.Pointer.
func MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParamsFrom(ptr unsafe.Pointer) MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams {
	return MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParamsClass) Alloc() MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams {
	rv := objc.Send[MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParamsClass) New() MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams {
	rv := objc.Send[MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams) Init() MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams {
	rv := objc.Send[MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams) Autorelease() MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams {
	rv := objc.Send[MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams creates a new MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams instance.
func NewMTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams() MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams {
	return getMTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParamsClass().New()
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams/serverSideProcessingTimeout
func (m_ MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams) ServerSideProcessingTimeout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams/serverSideProcessingTimeout
func (m_ MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams) SetServerSideProcessingTimeout(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}
// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams/timedInvokeTimeoutMs
func (m_ MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams) TimedInvokeTimeoutMs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams/timedInvokeTimeoutMs
func (m_ MTRThreadBorderRouterManagementClusterGetActiveDatasetRequestParams) SetTimedInvokeTimeoutMs(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


