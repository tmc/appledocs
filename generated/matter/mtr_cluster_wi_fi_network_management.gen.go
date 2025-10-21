// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterWiFiNetworkManagement] class.
var (
	MTRClusterWiFiNetworkManagementClass     _MTRClusterWiFiNetworkManagementClass
	MTRClusterWiFiNetworkManagementClassOnce sync.Once
)

func getMTRClusterWiFiNetworkManagementClass() _MTRClusterWiFiNetworkManagementClass {
	MTRClusterWiFiNetworkManagementClassOnce.Do(func() {
		MTRClusterWiFiNetworkManagementClass = _MTRClusterWiFiNetworkManagementClass{objc.GetClass("MTRClusterWiFiNetworkManagement")}
	})
	return MTRClusterWiFiNetworkManagementClass
}

type _MTRClusterWiFiNetworkManagementClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterWiFiNetworkManagement] class.
type IMTRClusterWiFiNetworkManagement interface {
	IMTRGenericCluster
	NetworkPassphraseRequestWithParamsExpectedValuesExpectedValueIntervalCompletion(params unsafe.Pointer, expectedDataValueDictionaries unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer, completion unsafe.Pointer)
	NetworkPassphraseRequestWithExpectedValuesExpectedValueIntervalCompletion(expectedValues unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer, completion unsafe.Pointer)
	ReadAttributeAcceptedCommandListWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeAttributeListWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeClusterRevisionWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeFeatureMapWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeGeneratedCommandListWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributePassphraseSurrogateWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeSSIDWithParams(params unsafe.Pointer) unsafe.Pointer
}

// Cluster Wi-Fi Network Management Functionality to retrieve operational information about a managed Wi-Fi network.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWiFiNetworkManagement
type MTRClusterWiFiNetworkManagement struct {
	MTRGenericCluster
}

// MTRClusterWiFiNetworkManagementFrom constructs a [MTRClusterWiFiNetworkManagement] from an unsafe.Pointer.
//
// Cluster Wi-Fi Network Management Functionality to retrieve operational information about a managed Wi-Fi network.
func MTRClusterWiFiNetworkManagementFrom(ptr unsafe.Pointer) MTRClusterWiFiNetworkManagement {
	return MTRClusterWiFiNetworkManagement{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterWiFiNetworkManagementClass) Alloc() MTRClusterWiFiNetworkManagement {
	rv := objc.Send[MTRClusterWiFiNetworkManagement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterWiFiNetworkManagementClass) New() MTRClusterWiFiNetworkManagement {
	rv := objc.Send[MTRClusterWiFiNetworkManagement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterWiFiNetworkManagement) Init() MTRClusterWiFiNetworkManagement {
	rv := objc.Send[MTRClusterWiFiNetworkManagement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterWiFiNetworkManagement) Autorelease() MTRClusterWiFiNetworkManagement {
	rv := objc.Send[MTRClusterWiFiNetworkManagement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterWiFiNetworkManagement creates a new MTRClusterWiFiNetworkManagement instance.
func NewMTRClusterWiFiNetworkManagement() MTRClusterWiFiNetworkManagement {
	return getMTRClusterWiFiNetworkManagementClass().New()
}


// For all instance methods that take a completion (i.e. command invocations), the completion will be called on the provided queue.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWiFiNetworkManagement/init(device:endpointID:queue:)
func NewMTRClusterWiFiNetworkManagementWithDeviceEndpointIDQueue(device unsafe.Pointer, endpointID unsafe.Pointer, queue unsafe.Pointer) MTRClusterWiFiNetworkManagement {
	instance := getMTRClusterWiFiNetworkManagementClass().Alloc()
	rv := objc.Send[MTRClusterWiFiNetworkManagement](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWiFiNetworkManagement/networkPassphraseRequest(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterWiFiNetworkManagement) NetworkPassphraseRequestWithParamsExpectedValuesExpectedValueIntervalCompletion(params unsafe.Pointer, expectedDataValueDictionaries unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("networkPassphraseRequestWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWiFiNetworkManagement/networkPassphraseRequest(withExpectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterWiFiNetworkManagement) NetworkPassphraseRequestWithExpectedValuesExpectedValueIntervalCompletion(expectedValues unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("networkPassphraseRequestWithExpectedValues:expectedValueInterval:completion:"), expectedValues, expectedValueIntervalMs, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWiFiNetworkManagement/readAttributeAcceptedCommandList(with:)
func (m_ MTRClusterWiFiNetworkManagement) ReadAttributeAcceptedCommandListWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWiFiNetworkManagement/readAttributeAttributeList(with:)
func (m_ MTRClusterWiFiNetworkManagement) ReadAttributeAttributeListWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeAttributeListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWiFiNetworkManagement/readAttributeClusterRevision(with:)
func (m_ MTRClusterWiFiNetworkManagement) ReadAttributeClusterRevisionWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeClusterRevisionWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWiFiNetworkManagement/readAttributeFeatureMap(with:)
func (m_ MTRClusterWiFiNetworkManagement) ReadAttributeFeatureMapWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeFeatureMapWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWiFiNetworkManagement/readAttributeGeneratedCommandList(with:)
func (m_ MTRClusterWiFiNetworkManagement) ReadAttributeGeneratedCommandListWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWiFiNetworkManagement/readAttributePassphraseSurrogate(with:)
func (m_ MTRClusterWiFiNetworkManagement) ReadAttributePassphraseSurrogateWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributePassphraseSurrogateWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWiFiNetworkManagement/readAttributeSSID(with:)
func (m_ MTRClusterWiFiNetworkManagement) ReadAttributeSSIDWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeSSIDWithParams:"), params)
	return rv
}


