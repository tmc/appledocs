// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	// methods:
	NetworkPassphraseRequestWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	NetworkPassphraseRequestWithExpectedValuesExpectedValueIntervalCompletion(expectedValues foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeAttributeListWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeClusterRevisionWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeFeatureMapWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeGeneratedCommandListWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributePassphraseSurrogateWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeSSIDWithParams(params IMTRReadParams) foundation.IDictionary
}

// Cluster Wi-Fi Network Management Functionality to retrieve operational information about a managed Wi-Fi network.


// Cluster Wi-Fi Network Management Functionality to retrieve operational information about a managed Wi-Fi network.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWiFiNetworkManagement/init(device:endpointID:queue:)
func NewMTRClusterWiFiNetworkManagementWithDeviceEndpointIDQueue(device IMTRDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRClusterWiFiNetworkManagement {
	instance := getMTRClusterWiFiNetworkManagementClass().Alloc()
	rv := objc.Send[MTRClusterWiFiNetworkManagement](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWiFiNetworkManagement/networkPassphraseRequest(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterWiFiNetworkManagement) NetworkPassphraseRequestWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("networkPassphraseRequestWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWiFiNetworkManagement/networkPassphraseRequest(withExpectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterWiFiNetworkManagement) NetworkPassphraseRequestWithExpectedValuesExpectedValueIntervalCompletion(expectedValues foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("networkPassphraseRequestWithExpectedValues:expectedValueInterval:completion:"), expectedValues, expectedValueIntervalMs, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWiFiNetworkManagement/readAttributeAcceptedCommandList(with:)
func (m_ MTRClusterWiFiNetworkManagement) ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWiFiNetworkManagement/readAttributeAttributeList(with:)
func (m_ MTRClusterWiFiNetworkManagement) ReadAttributeAttributeListWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeAttributeListWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWiFiNetworkManagement/readAttributeClusterRevision(with:)
func (m_ MTRClusterWiFiNetworkManagement) ReadAttributeClusterRevisionWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeClusterRevisionWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWiFiNetworkManagement/readAttributeFeatureMap(with:)
func (m_ MTRClusterWiFiNetworkManagement) ReadAttributeFeatureMapWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeFeatureMapWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWiFiNetworkManagement/readAttributeGeneratedCommandList(with:)
func (m_ MTRClusterWiFiNetworkManagement) ReadAttributeGeneratedCommandListWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWiFiNetworkManagement/readAttributePassphraseSurrogate(with:)
func (m_ MTRClusterWiFiNetworkManagement) ReadAttributePassphraseSurrogateWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributePassphraseSurrogateWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWiFiNetworkManagement/readAttributeSSID(with:)
func (m_ MTRClusterWiFiNetworkManagement) ReadAttributeSSIDWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeSSIDWithParams:"), params)
	return rv
}


