// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRClusterDeviceEnergyManagementMode] class.
var (
	MTRClusterDeviceEnergyManagementModeClass     _MTRClusterDeviceEnergyManagementModeClass
	MTRClusterDeviceEnergyManagementModeClassOnce sync.Once
)

func getMTRClusterDeviceEnergyManagementModeClass() _MTRClusterDeviceEnergyManagementModeClass {
	MTRClusterDeviceEnergyManagementModeClassOnce.Do(func() {
		MTRClusterDeviceEnergyManagementModeClass = _MTRClusterDeviceEnergyManagementModeClass{objc.GetClass("MTRClusterDeviceEnergyManagementMode")}
	})
	return MTRClusterDeviceEnergyManagementModeClass
}

type _MTRClusterDeviceEnergyManagementModeClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterDeviceEnergyManagementMode] class.
type IMTRClusterDeviceEnergyManagementMode interface {
	IMTRGenericCluster
	// properties:
	// methods:
	ChangeToModeWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRDeviceEnergyManagementModeClusterChangeToModeParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeAttributeListWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeClusterRevisionWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeCurrentModeWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeFeatureMapWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeGeneratedCommandListWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeSupportedModesWithParams(params IMTRReadParams) foundation.IDictionary
}

// Cluster Device Energy Management Mode Attributes and commands for selecting a mode from a list of supported options.


// Cluster Device Energy Management Mode Attributes and commands for selecting a mode from a list of supported options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDeviceEnergyManagementMode
type MTRClusterDeviceEnergyManagementMode struct {
	MTRGenericCluster
}

// MTRClusterDeviceEnergyManagementModeFrom constructs a [MTRClusterDeviceEnergyManagementMode] from an unsafe.Pointer.
//
// Cluster Device Energy Management Mode Attributes and commands for selecting a mode from a list of supported options.
func MTRClusterDeviceEnergyManagementModeFrom(ptr unsafe.Pointer) MTRClusterDeviceEnergyManagementMode {
	return MTRClusterDeviceEnergyManagementMode{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterDeviceEnergyManagementModeClass) Alloc() MTRClusterDeviceEnergyManagementMode {
	rv := objc.Send[MTRClusterDeviceEnergyManagementMode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterDeviceEnergyManagementModeClass) New() MTRClusterDeviceEnergyManagementMode {
	rv := objc.Send[MTRClusterDeviceEnergyManagementMode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterDeviceEnergyManagementMode) Init() MTRClusterDeviceEnergyManagementMode {
	rv := objc.Send[MTRClusterDeviceEnergyManagementMode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterDeviceEnergyManagementMode) Autorelease() MTRClusterDeviceEnergyManagementMode {
	rv := objc.Send[MTRClusterDeviceEnergyManagementMode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterDeviceEnergyManagementMode creates a new MTRClusterDeviceEnergyManagementMode instance.
func NewMTRClusterDeviceEnergyManagementMode() MTRClusterDeviceEnergyManagementMode {
	return getMTRClusterDeviceEnergyManagementModeClass().New()
}



// For all instance methods that take a completion (i.e. command invocations), the completion will be called on the provided queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDeviceEnergyManagementMode/init(device:endpointID:queue:)
func NewMTRClusterDeviceEnergyManagementModeWithDeviceEndpointIDQueue(device IMTRDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRClusterDeviceEnergyManagementMode {
	instance := getMTRClusterDeviceEnergyManagementModeClass().Alloc()
	rv := objc.Send[MTRClusterDeviceEnergyManagementMode](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDeviceEnergyManagementMode/changeToMode(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterDeviceEnergyManagementMode) ChangeToModeWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRDeviceEnergyManagementModeClusterChangeToModeParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("changeToModeWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDeviceEnergyManagementMode/readAttributeAcceptedCommandList(with:)
func (m_ MTRClusterDeviceEnergyManagementMode) ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDeviceEnergyManagementMode/readAttributeAttributeList(with:)
func (m_ MTRClusterDeviceEnergyManagementMode) ReadAttributeAttributeListWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeAttributeListWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDeviceEnergyManagementMode/readAttributeClusterRevision(with:)
func (m_ MTRClusterDeviceEnergyManagementMode) ReadAttributeClusterRevisionWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeClusterRevisionWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDeviceEnergyManagementMode/readAttributeCurrentMode(with:)
func (m_ MTRClusterDeviceEnergyManagementMode) ReadAttributeCurrentModeWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeCurrentModeWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDeviceEnergyManagementMode/readAttributeFeatureMap(with:)
func (m_ MTRClusterDeviceEnergyManagementMode) ReadAttributeFeatureMapWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeFeatureMapWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDeviceEnergyManagementMode/readAttributeGeneratedCommandList(with:)
func (m_ MTRClusterDeviceEnergyManagementMode) ReadAttributeGeneratedCommandListWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDeviceEnergyManagementMode/readAttributeSupportedModes(with:)
func (m_ MTRClusterDeviceEnergyManagementMode) ReadAttributeSupportedModesWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeSupportedModesWithParams:"), params)
	return rv
}


