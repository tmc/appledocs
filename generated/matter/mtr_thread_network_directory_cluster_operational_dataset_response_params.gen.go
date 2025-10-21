// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams] class.
var (
	MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParamsClass     _MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParamsClass
	MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParamsClassOnce sync.Once
)

func getMTRThreadNetworkDirectoryClusterOperationalDatasetResponseParamsClass() _MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParamsClass {
	MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParamsClassOnce.Do(func() {
		MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParamsClass = _MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParamsClass{objc.GetClass("MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams")}
	})
	return MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParamsClass
}

type _MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams] class.
type IMTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams
type MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams struct {
	objectivec.Object
}

// MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParamsFrom constructs a [MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams] from an unsafe.Pointer.
func MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParamsFrom(ptr unsafe.Pointer) MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams {
	return MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParamsClass) Alloc() MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams {
	rv := objc.Send[MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParamsClass) New() MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams {
	rv := objc.Send[MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams) Init() MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams {
	rv := objc.Send[MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams) Autorelease() MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams {
	rv := objc.Send[MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams creates a new MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams instance.
func NewMTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams() MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams {
	return getMTRThreadNetworkDirectoryClusterOperationalDatasetResponseParamsClass().New()
}




// Initialize an MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams/init(responseValue:)
func NewMTRThreadNetworkDirectoryClusterOperationalDatasetResponseParamsWithResponseValueError(responseValue unsafe.Pointer, error_ unsafe.Pointer) MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams {
	instance := getMTRThreadNetworkDirectoryClusterOperationalDatasetResponseParamsClass().Alloc()
	rv := objc.Send[MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams/operationalDataset
func (m_ MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams) OperationalDataset() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("operationalDataset"))
	return rv
}


// SetOperationalDataset sets the value of the operationalDataset property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams/operationalDataset
func (m_ MTRThreadNetworkDirectoryClusterOperationalDatasetResponseParams) SetOperationalDataset(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationalDataset:"), value)
}


