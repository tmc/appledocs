// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterThreadNetworkDirectory] class.
var (
	MTRClusterThreadNetworkDirectoryClass     _MTRClusterThreadNetworkDirectoryClass
	MTRClusterThreadNetworkDirectoryClassOnce sync.Once
)

func getMTRClusterThreadNetworkDirectoryClass() _MTRClusterThreadNetworkDirectoryClass {
	MTRClusterThreadNetworkDirectoryClassOnce.Do(func() {
		MTRClusterThreadNetworkDirectoryClass = _MTRClusterThreadNetworkDirectoryClass{objc.GetClass("MTRClusterThreadNetworkDirectory")}
	})
	return MTRClusterThreadNetworkDirectoryClass
}

type _MTRClusterThreadNetworkDirectoryClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterThreadNetworkDirectory] class.
type IMTRClusterThreadNetworkDirectory interface {
	IMTRGenericCluster
	AddNetworkWithParamsExpectedValuesExpectedValueIntervalCompletion(params unsafe.Pointer, expectedDataValueDictionaries unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer, completion unsafe.Pointer)
	GetOperationalDatasetWithParamsExpectedValuesExpectedValueIntervalCompletion(params unsafe.Pointer, expectedDataValueDictionaries unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer, completion unsafe.Pointer)
	ReadAttributeAcceptedCommandListWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeAttributeListWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeClusterRevisionWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeFeatureMapWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeGeneratedCommandListWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributePreferredExtendedPanIDWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeThreadNetworkTableSizeWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeThreadNetworksWithParams(params unsafe.Pointer) unsafe.Pointer
	RemoveNetworkWithParamsExpectedValuesExpectedValueIntervalCompletion(params unsafe.Pointer, expectedDataValueDictionaries unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer, completion unsafe.Pointer)
	WriteAttributePreferredExtendedPanIDWithValueExpectedValueInterval(dataValueDictionary unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer)
	WriteAttributePreferredExtendedPanIDWithValueExpectedValueIntervalParams(dataValueDictionary unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer, params unsafe.Pointer)
}

// Cluster Thread Network Directory Manages the names and credentials of Thread networks visible to the user.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterThreadNetworkDirectory
type MTRClusterThreadNetworkDirectory struct {
	MTRGenericCluster
}

// MTRClusterThreadNetworkDirectoryFrom constructs a [MTRClusterThreadNetworkDirectory] from an unsafe.Pointer.
//
// Cluster Thread Network Directory Manages the names and credentials of Thread networks visible to the user.
func MTRClusterThreadNetworkDirectoryFrom(ptr unsafe.Pointer) MTRClusterThreadNetworkDirectory {
	return MTRClusterThreadNetworkDirectory{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterThreadNetworkDirectoryClass) Alloc() MTRClusterThreadNetworkDirectory {
	rv := objc.Send[MTRClusterThreadNetworkDirectory](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterThreadNetworkDirectoryClass) New() MTRClusterThreadNetworkDirectory {
	rv := objc.Send[MTRClusterThreadNetworkDirectory](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterThreadNetworkDirectory) Init() MTRClusterThreadNetworkDirectory {
	rv := objc.Send[MTRClusterThreadNetworkDirectory](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterThreadNetworkDirectory) Autorelease() MTRClusterThreadNetworkDirectory {
	rv := objc.Send[MTRClusterThreadNetworkDirectory](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterThreadNetworkDirectory creates a new MTRClusterThreadNetworkDirectory instance.
func NewMTRClusterThreadNetworkDirectory() MTRClusterThreadNetworkDirectory {
	return getMTRClusterThreadNetworkDirectoryClass().New()
}




// For all instance methods that take a completion (i.e. command invocations), the completion will be called on the provided queue.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterThreadNetworkDirectory/init(device:endpointID:queue:)
func NewMTRClusterThreadNetworkDirectoryWithDeviceEndpointIDQueue(device unsafe.Pointer, endpointID unsafe.Pointer, queue unsafe.Pointer) MTRClusterThreadNetworkDirectory {
	instance := getMTRClusterThreadNetworkDirectoryClass().Alloc()
	rv := objc.Send[MTRClusterThreadNetworkDirectory](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterThreadNetworkDirectory/addNetwork(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterThreadNetworkDirectory) AddNetworkWithParamsExpectedValuesExpectedValueIntervalCompletion(params unsafe.Pointer, expectedDataValueDictionaries unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addNetworkWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterThreadNetworkDirectory/getOperationalDataset(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterThreadNetworkDirectory) GetOperationalDatasetWithParamsExpectedValuesExpectedValueIntervalCompletion(params unsafe.Pointer, expectedDataValueDictionaries unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("getOperationalDatasetWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterThreadNetworkDirectory/readAttributeAcceptedCommandList(with:)
func (m_ MTRClusterThreadNetworkDirectory) ReadAttributeAcceptedCommandListWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterThreadNetworkDirectory/readAttributeAttributeList(with:)
func (m_ MTRClusterThreadNetworkDirectory) ReadAttributeAttributeListWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeAttributeListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterThreadNetworkDirectory/readAttributeClusterRevision(with:)
func (m_ MTRClusterThreadNetworkDirectory) ReadAttributeClusterRevisionWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeClusterRevisionWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterThreadNetworkDirectory/readAttributeFeatureMap(with:)
func (m_ MTRClusterThreadNetworkDirectory) ReadAttributeFeatureMapWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeFeatureMapWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterThreadNetworkDirectory/readAttributeGeneratedCommandList(with:)
func (m_ MTRClusterThreadNetworkDirectory) ReadAttributeGeneratedCommandListWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterThreadNetworkDirectory/readAttributePreferredExtendedPanID(with:)
func (m_ MTRClusterThreadNetworkDirectory) ReadAttributePreferredExtendedPanIDWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributePreferredExtendedPanIDWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterThreadNetworkDirectory/readAttributeThreadNetworkTableSize(with:)
func (m_ MTRClusterThreadNetworkDirectory) ReadAttributeThreadNetworkTableSizeWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeThreadNetworkTableSizeWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterThreadNetworkDirectory/readAttributeThreadNetworks(with:)
func (m_ MTRClusterThreadNetworkDirectory) ReadAttributeThreadNetworksWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeThreadNetworksWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterThreadNetworkDirectory/removeNetwork(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterThreadNetworkDirectory) RemoveNetworkWithParamsExpectedValuesExpectedValueIntervalCompletion(params unsafe.Pointer, expectedDataValueDictionaries unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeNetworkWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterThreadNetworkDirectory/writeAttributePreferredExtendedPanID(withValue:expectedValueInterval:)
func (m_ MTRClusterThreadNetworkDirectory) WriteAttributePreferredExtendedPanIDWithValueExpectedValueInterval(dataValueDictionary unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributePreferredExtendedPanIDWithValue:expectedValueInterval:"), dataValueDictionary, expectedValueIntervalMs)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterThreadNetworkDirectory/writeAttributePreferredExtendedPanID(withValue:expectedValueInterval:params:)
func (m_ MTRClusterThreadNetworkDirectory) WriteAttributePreferredExtendedPanIDWithValueExpectedValueIntervalParams(dataValueDictionary unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer, params unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributePreferredExtendedPanIDWithValue:expectedValueInterval:params:"), dataValueDictionary, expectedValueIntervalMs, params)
}


