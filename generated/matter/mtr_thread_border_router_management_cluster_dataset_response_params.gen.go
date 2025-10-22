// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRThreadBorderRouterManagementClusterDatasetResponseParams] class.
var (
	MTRThreadBorderRouterManagementClusterDatasetResponseParamsClass     _MTRThreadBorderRouterManagementClusterDatasetResponseParamsClass
	MTRThreadBorderRouterManagementClusterDatasetResponseParamsClassOnce sync.Once
)

func getMTRThreadBorderRouterManagementClusterDatasetResponseParamsClass() _MTRThreadBorderRouterManagementClusterDatasetResponseParamsClass {
	MTRThreadBorderRouterManagementClusterDatasetResponseParamsClassOnce.Do(func() {
		MTRThreadBorderRouterManagementClusterDatasetResponseParamsClass = _MTRThreadBorderRouterManagementClusterDatasetResponseParamsClass{objc.GetClass("MTRThreadBorderRouterManagementClusterDatasetResponseParams")}
	})
	return MTRThreadBorderRouterManagementClusterDatasetResponseParamsClass
}

type _MTRThreadBorderRouterManagementClusterDatasetResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRThreadBorderRouterManagementClusterDatasetResponseParams] class.
type IMTRThreadBorderRouterManagementClusterDatasetResponseParams interface {
	objectivec.IObject
	Dataset() foundation.NSData
	SetDataset(value foundation.IData)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterDatasetResponseParams
type MTRThreadBorderRouterManagementClusterDatasetResponseParams struct {
	objectivec.Object
}

// MTRThreadBorderRouterManagementClusterDatasetResponseParamsFrom constructs a [MTRThreadBorderRouterManagementClusterDatasetResponseParams] from an unsafe.Pointer.
func MTRThreadBorderRouterManagementClusterDatasetResponseParamsFrom(ptr unsafe.Pointer) MTRThreadBorderRouterManagementClusterDatasetResponseParams {
	return MTRThreadBorderRouterManagementClusterDatasetResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThreadBorderRouterManagementClusterDatasetResponseParamsClass) Alloc() MTRThreadBorderRouterManagementClusterDatasetResponseParams {
	rv := objc.Send[MTRThreadBorderRouterManagementClusterDatasetResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThreadBorderRouterManagementClusterDatasetResponseParamsClass) New() MTRThreadBorderRouterManagementClusterDatasetResponseParams {
	rv := objc.Send[MTRThreadBorderRouterManagementClusterDatasetResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThreadBorderRouterManagementClusterDatasetResponseParams) Init() MTRThreadBorderRouterManagementClusterDatasetResponseParams {
	rv := objc.Send[MTRThreadBorderRouterManagementClusterDatasetResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThreadBorderRouterManagementClusterDatasetResponseParams) Autorelease() MTRThreadBorderRouterManagementClusterDatasetResponseParams {
	rv := objc.Send[MTRThreadBorderRouterManagementClusterDatasetResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThreadBorderRouterManagementClusterDatasetResponseParams creates a new MTRThreadBorderRouterManagementClusterDatasetResponseParams instance.
func NewMTRThreadBorderRouterManagementClusterDatasetResponseParams() MTRThreadBorderRouterManagementClusterDatasetResponseParams {
	return getMTRThreadBorderRouterManagementClusterDatasetResponseParamsClass().New()
}




// Initialize an MTRThreadBorderRouterManagementClusterDatasetResponseParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterDatasetResponseParams/init(responseValue:)
func NewMTRThreadBorderRouterManagementClusterDatasetResponseParamsWithResponseValueError(responseValue unsafe.Pointer, error_ unsafe.Pointer) MTRThreadBorderRouterManagementClusterDatasetResponseParams {
	instance := getMTRThreadBorderRouterManagementClusterDatasetResponseParamsClass().Alloc()
	rv := objc.Send[MTRThreadBorderRouterManagementClusterDatasetResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterDatasetResponseParams/dataset
func (m_ MTRThreadBorderRouterManagementClusterDatasetResponseParams) Dataset() foundation.NSData {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("dataset"))
	return rv
}


// SetDataset sets the value of the dataset property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadBorderRouterManagementClusterDatasetResponseParams/dataset
func (m_ MTRThreadBorderRouterManagementClusterDatasetResponseParams) SetDataset(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDataset:"), value)
}


