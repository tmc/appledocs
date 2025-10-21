// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterThreadBorderRouterManagement] class.
var (
	MTRClusterThreadBorderRouterManagementClass     _MTRClusterThreadBorderRouterManagementClass
	MTRClusterThreadBorderRouterManagementClassOnce sync.Once
)

func getMTRClusterThreadBorderRouterManagementClass() _MTRClusterThreadBorderRouterManagementClass {
	MTRClusterThreadBorderRouterManagementClassOnce.Do(func() {
		MTRClusterThreadBorderRouterManagementClass = _MTRClusterThreadBorderRouterManagementClass{objc.GetClass("MTRClusterThreadBorderRouterManagement")}
	})
	return MTRClusterThreadBorderRouterManagementClass
}

type _MTRClusterThreadBorderRouterManagementClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterThreadBorderRouterManagement] class.
type IMTRClusterThreadBorderRouterManagement interface {
	IMTRGenericCluster
	GetActiveDatasetRequestWithParamsExpectedValuesExpectedValueIntervalCompletion(params unsafe.Pointer, expectedDataValueDictionaries unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer, completion unsafe.Pointer)
	GetActiveDatasetRequestWithExpectedValuesExpectedValueIntervalCompletion(expectedValues unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer, completion unsafe.Pointer)
	GetPendingDatasetRequestWithParamsExpectedValuesExpectedValueIntervalCompletion(params unsafe.Pointer, expectedDataValueDictionaries unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer, completion unsafe.Pointer)
	GetPendingDatasetRequestWithExpectedValuesExpectedValueIntervalCompletion(expectedValues unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer, completion unsafe.Pointer)
	ReadAttributeAcceptedCommandListWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeActiveDatasetTimestampWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeAttributeListWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeBorderAgentIDWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeBorderRouterNameWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeClusterRevisionWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeFeatureMapWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeGeneratedCommandListWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeInterfaceEnabledWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributePendingDatasetTimestampWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeThreadVersionWithParams(params unsafe.Pointer) unsafe.Pointer
	SetActiveDatasetRequestWithParamsExpectedValuesExpectedValueIntervalCompletion(params unsafe.Pointer, expectedDataValueDictionaries unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer, completion unsafe.Pointer)
	SetPendingDatasetRequestWithParamsExpectedValuesExpectedValueIntervalCompletion(params unsafe.Pointer, expectedDataValueDictionaries unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer, completion unsafe.Pointer)
}

// Cluster Thread Border Router Management Manage the Thread network of Thread Border Router
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterThreadBorderRouterManagement
type MTRClusterThreadBorderRouterManagement struct {
	MTRGenericCluster
}

// MTRClusterThreadBorderRouterManagementFrom constructs a [MTRClusterThreadBorderRouterManagement] from an unsafe.Pointer.
//
// Cluster Thread Border Router Management Manage the Thread network of Thread Border Router
func MTRClusterThreadBorderRouterManagementFrom(ptr unsafe.Pointer) MTRClusterThreadBorderRouterManagement {
	return MTRClusterThreadBorderRouterManagement{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterThreadBorderRouterManagementClass) Alloc() MTRClusterThreadBorderRouterManagement {
	rv := objc.Send[MTRClusterThreadBorderRouterManagement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterThreadBorderRouterManagementClass) New() MTRClusterThreadBorderRouterManagement {
	rv := objc.Send[MTRClusterThreadBorderRouterManagement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterThreadBorderRouterManagement) Init() MTRClusterThreadBorderRouterManagement {
	rv := objc.Send[MTRClusterThreadBorderRouterManagement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterThreadBorderRouterManagement) Autorelease() MTRClusterThreadBorderRouterManagement {
	rv := objc.Send[MTRClusterThreadBorderRouterManagement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterThreadBorderRouterManagement creates a new MTRClusterThreadBorderRouterManagement instance.
func NewMTRClusterThreadBorderRouterManagement() MTRClusterThreadBorderRouterManagement {
	return getMTRClusterThreadBorderRouterManagementClass().New()
}


// For all instance methods that take a completion (i.e. command invocations), the completion will be called on the provided queue.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterThreadBorderRouterManagement/init(device:endpointID:queue:)
func NewMTRClusterThreadBorderRouterManagementWithDeviceEndpointIDQueue(device unsafe.Pointer, endpointID unsafe.Pointer, queue unsafe.Pointer) MTRClusterThreadBorderRouterManagement {
	instance := getMTRClusterThreadBorderRouterManagementClass().Alloc()
	rv := objc.Send[MTRClusterThreadBorderRouterManagement](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterThreadBorderRouterManagement/getActiveDatasetRequest(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterThreadBorderRouterManagement) GetActiveDatasetRequestWithParamsExpectedValuesExpectedValueIntervalCompletion(params unsafe.Pointer, expectedDataValueDictionaries unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("getActiveDatasetRequestWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterThreadBorderRouterManagement/getActiveDatasetRequest(withExpectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterThreadBorderRouterManagement) GetActiveDatasetRequestWithExpectedValuesExpectedValueIntervalCompletion(expectedValues unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("getActiveDatasetRequestWithExpectedValues:expectedValueInterval:completion:"), expectedValues, expectedValueIntervalMs, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterThreadBorderRouterManagement/getPendingDatasetRequest(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterThreadBorderRouterManagement) GetPendingDatasetRequestWithParamsExpectedValuesExpectedValueIntervalCompletion(params unsafe.Pointer, expectedDataValueDictionaries unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("getPendingDatasetRequestWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterThreadBorderRouterManagement/getPendingDatasetRequest(withExpectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterThreadBorderRouterManagement) GetPendingDatasetRequestWithExpectedValuesExpectedValueIntervalCompletion(expectedValues unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("getPendingDatasetRequestWithExpectedValues:expectedValueInterval:completion:"), expectedValues, expectedValueIntervalMs, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterThreadBorderRouterManagement/readAttributeAcceptedCommandList(with:)
func (m_ MTRClusterThreadBorderRouterManagement) ReadAttributeAcceptedCommandListWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterThreadBorderRouterManagement/readAttributeActiveDatasetTimestamp(with:)
func (m_ MTRClusterThreadBorderRouterManagement) ReadAttributeActiveDatasetTimestampWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeActiveDatasetTimestampWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterThreadBorderRouterManagement/readAttributeAttributeList(with:)
func (m_ MTRClusterThreadBorderRouterManagement) ReadAttributeAttributeListWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeAttributeListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterThreadBorderRouterManagement/readAttributeBorderAgentID(with:)
func (m_ MTRClusterThreadBorderRouterManagement) ReadAttributeBorderAgentIDWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeBorderAgentIDWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterThreadBorderRouterManagement/readAttributeBorderRouterName(with:)
func (m_ MTRClusterThreadBorderRouterManagement) ReadAttributeBorderRouterNameWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeBorderRouterNameWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterThreadBorderRouterManagement/readAttributeClusterRevision(with:)
func (m_ MTRClusterThreadBorderRouterManagement) ReadAttributeClusterRevisionWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeClusterRevisionWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterThreadBorderRouterManagement/readAttributeFeatureMap(with:)
func (m_ MTRClusterThreadBorderRouterManagement) ReadAttributeFeatureMapWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeFeatureMapWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterThreadBorderRouterManagement/readAttributeGeneratedCommandList(with:)
func (m_ MTRClusterThreadBorderRouterManagement) ReadAttributeGeneratedCommandListWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterThreadBorderRouterManagement/readAttributeInterfaceEnabled(with:)
func (m_ MTRClusterThreadBorderRouterManagement) ReadAttributeInterfaceEnabledWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeInterfaceEnabledWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterThreadBorderRouterManagement/readAttributePendingDatasetTimestamp(with:)
func (m_ MTRClusterThreadBorderRouterManagement) ReadAttributePendingDatasetTimestampWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributePendingDatasetTimestampWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterThreadBorderRouterManagement/readAttributeThreadVersion(with:)
func (m_ MTRClusterThreadBorderRouterManagement) ReadAttributeThreadVersionWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeThreadVersionWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterThreadBorderRouterManagement/setActiveDatasetRequestWith(_:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterThreadBorderRouterManagement) SetActiveDatasetRequestWithParamsExpectedValuesExpectedValueIntervalCompletion(params unsafe.Pointer, expectedDataValueDictionaries unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActiveDatasetRequestWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterThreadBorderRouterManagement/setPendingDatasetRequestWith(_:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterThreadBorderRouterManagement) SetPendingDatasetRequestWithParamsExpectedValuesExpectedValueIntervalCompletion(params unsafe.Pointer, expectedDataValueDictionaries unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPendingDatasetRequestWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}


