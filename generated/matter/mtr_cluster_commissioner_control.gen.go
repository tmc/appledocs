// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRClusterCommissionerControl] class.
var (
	MTRClusterCommissionerControlClass     _MTRClusterCommissionerControlClass
	MTRClusterCommissionerControlClassOnce sync.Once
)

func getMTRClusterCommissionerControlClass() _MTRClusterCommissionerControlClass {
	MTRClusterCommissionerControlClassOnce.Do(func() {
		MTRClusterCommissionerControlClass = _MTRClusterCommissionerControlClass{objc.GetClass("MTRClusterCommissionerControl")}
	})
	return MTRClusterCommissionerControlClass
}

type _MTRClusterCommissionerControlClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterCommissionerControl] class.
type IMTRClusterCommissionerControl interface {
	IMTRGenericCluster
	// properties:
	// methods:
	CommissionNodeWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRCommissionerControlClusterCommissionNodeParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeAttributeListWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeClusterRevisionWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeFeatureMapWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeGeneratedCommandListWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeSupportedDeviceCategoriesWithParams(params IMTRReadParams) foundation.IDictionary
	RequestCommissioningApprovalWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRCommissionerControlClusterRequestCommissioningApprovalParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
}

// Cluster Commissioner Control Supports the ability for clients to request the commissioning of themselves or other nodes onto a fabric which the cluster server can commission onto.


// Cluster Commissioner Control Supports the ability for clients to request the commissioning of themselves or other nodes onto a fabric which the cluster server can commission onto.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterCommissionerControl
type MTRClusterCommissionerControl struct {
	MTRGenericCluster
}

// MTRClusterCommissionerControlFrom constructs a [MTRClusterCommissionerControl] from an unsafe.Pointer.
//
// Cluster Commissioner Control Supports the ability for clients to request the commissioning of themselves or other nodes onto a fabric which the cluster server can commission onto.
func MTRClusterCommissionerControlFrom(ptr unsafe.Pointer) MTRClusterCommissionerControl {
	return MTRClusterCommissionerControl{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterCommissionerControlClass) Alloc() MTRClusterCommissionerControl {
	rv := objc.Send[MTRClusterCommissionerControl](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterCommissionerControlClass) New() MTRClusterCommissionerControl {
	rv := objc.Send[MTRClusterCommissionerControl](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterCommissionerControl) Init() MTRClusterCommissionerControl {
	rv := objc.Send[MTRClusterCommissionerControl](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterCommissionerControl) Autorelease() MTRClusterCommissionerControl {
	rv := objc.Send[MTRClusterCommissionerControl](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterCommissionerControl creates a new MTRClusterCommissionerControl instance.
func NewMTRClusterCommissionerControl() MTRClusterCommissionerControl {
	return getMTRClusterCommissionerControlClass().New()
}



// For all instance methods that take a completion (i.e. command invocations), the completion will be called on the provided queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterCommissionerControl/init(device:endpointID:queue:)
func NewMTRClusterCommissionerControlWithDeviceEndpointIDQueue(device IMTRDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRClusterCommissionerControl {
	instance := getMTRClusterCommissionerControlClass().Alloc()
	rv := objc.Send[MTRClusterCommissionerControl](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterCommissionerControl/commissionNode(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterCommissionerControl) CommissionNodeWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRCommissionerControlClusterCommissionNodeParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("commissionNodeWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterCommissionerControl/readAttributeAcceptedCommandList(with:)
func (m_ MTRClusterCommissionerControl) ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterCommissionerControl/readAttributeAttributeList(with:)
func (m_ MTRClusterCommissionerControl) ReadAttributeAttributeListWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeAttributeListWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterCommissionerControl/readAttributeClusterRevision(with:)
func (m_ MTRClusterCommissionerControl) ReadAttributeClusterRevisionWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeClusterRevisionWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterCommissionerControl/readAttributeFeatureMap(with:)
func (m_ MTRClusterCommissionerControl) ReadAttributeFeatureMapWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeFeatureMapWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterCommissionerControl/readAttributeGeneratedCommandList(with:)
func (m_ MTRClusterCommissionerControl) ReadAttributeGeneratedCommandListWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterCommissionerControl/readAttributeSupportedDeviceCategories(with:)
func (m_ MTRClusterCommissionerControl) ReadAttributeSupportedDeviceCategoriesWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeSupportedDeviceCategoriesWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterCommissionerControl/requestCommissioningApproval(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterCommissionerControl) RequestCommissioningApprovalWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRCommissionerControlClusterRequestCommissioningApprovalParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("requestCommissioningApprovalWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}


