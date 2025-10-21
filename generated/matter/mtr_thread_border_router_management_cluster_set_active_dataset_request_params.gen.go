// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams] class.
var (
	MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParamsClass     _MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParamsClass
	MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParamsClassOnce sync.Once
)

func getMTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParamsClass() _MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParamsClass {
	MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParamsClassOnce.Do(func() {
		MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParamsClass = _MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParamsClass{objc.GetClass("MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams")}
	})
	return MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParamsClass
}

type _MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams] class.
type IMTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams
type MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams struct {
	objectivec.Object
}

// MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParamsFrom constructs a [MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams] from an unsafe.Pointer.
func MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParamsFrom(ptr unsafe.Pointer) MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams {
	return MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParamsClass) Alloc() MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams {
	rv := objc.Send[MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParamsClass) New() MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams {
	rv := objc.Send[MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams) Init() MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams {
	rv := objc.Send[MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams) Autorelease() MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams {
	rv := objc.Send[MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams creates a new MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams instance.
func NewMTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams() MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams {
	return getMTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams/activeDataset
func (m_ MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams) ActiveDataset() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("activeDataset"))
	return rv
}


// SetActiveDataset sets the value of the activeDataset property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams/activeDataset
func (m_ MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams) SetActiveDataset(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActiveDataset:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams/breadcrumb
func (m_ MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams) Breadcrumb() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("breadcrumb"))
	return rv
}


// SetBreadcrumb sets the value of the breadcrumb property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams/breadcrumb
func (m_ MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams) SetBreadcrumb(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBreadcrumb:"), value)
}

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams/serverSideProcessingTimeout
func (m_ MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams) ServerSideProcessingTimeout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams/serverSideProcessingTimeout
func (m_ MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams) SetServerSideProcessingTimeout(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams/timedInvokeTimeoutMs
func (m_ MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams) TimedInvokeTimeoutMs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams/timedInvokeTimeoutMs
func (m_ MTRThreadBorderRouterManagementClusterSetActiveDatasetRequestParams) SetTimedInvokeTimeoutMs(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



