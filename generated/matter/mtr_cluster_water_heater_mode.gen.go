// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRClusterWaterHeaterMode] class.
var (
	MTRClusterWaterHeaterModeClass     _MTRClusterWaterHeaterModeClass
	MTRClusterWaterHeaterModeClassOnce sync.Once
)

func getMTRClusterWaterHeaterModeClass() _MTRClusterWaterHeaterModeClass {
	MTRClusterWaterHeaterModeClassOnce.Do(func() {
		MTRClusterWaterHeaterModeClass = _MTRClusterWaterHeaterModeClass{objc.GetClass("MTRClusterWaterHeaterMode")}
	})
	return MTRClusterWaterHeaterModeClass
}

type _MTRClusterWaterHeaterModeClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterWaterHeaterMode] class.
type IMTRClusterWaterHeaterMode interface {
	IMTRGenericCluster
	// properties:
	// methods:
	ChangeToModeWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRWaterHeaterModeClusterChangeToModeParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeAttributeListWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeClusterRevisionWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeCurrentModeWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeFeatureMapWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeGeneratedCommandListWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeSupportedModesWithParams(params IMTRReadParams) foundation.IDictionary
}

// Cluster Water Heater Mode Attributes and commands for selecting a mode from a list of supported options.


// Cluster Water Heater Mode Attributes and commands for selecting a mode from a list of supported options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWaterHeaterMode
type MTRClusterWaterHeaterMode struct {
	MTRGenericCluster
}

// MTRClusterWaterHeaterModeFrom constructs a [MTRClusterWaterHeaterMode] from an unsafe.Pointer.
//
// Cluster Water Heater Mode Attributes and commands for selecting a mode from a list of supported options.
func MTRClusterWaterHeaterModeFrom(ptr unsafe.Pointer) MTRClusterWaterHeaterMode {
	return MTRClusterWaterHeaterMode{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterWaterHeaterModeClass) Alloc() MTRClusterWaterHeaterMode {
	rv := objc.Send[MTRClusterWaterHeaterMode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterWaterHeaterModeClass) New() MTRClusterWaterHeaterMode {
	rv := objc.Send[MTRClusterWaterHeaterMode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterWaterHeaterMode) Init() MTRClusterWaterHeaterMode {
	rv := objc.Send[MTRClusterWaterHeaterMode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterWaterHeaterMode) Autorelease() MTRClusterWaterHeaterMode {
	rv := objc.Send[MTRClusterWaterHeaterMode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterWaterHeaterMode creates a new MTRClusterWaterHeaterMode instance.
func NewMTRClusterWaterHeaterMode() MTRClusterWaterHeaterMode {
	return getMTRClusterWaterHeaterModeClass().New()
}



// For all instance methods that take a completion (i.e. command invocations), the completion will be called on the provided queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWaterHeaterMode/init(device:endpointID:queue:)
func NewMTRClusterWaterHeaterModeWithDeviceEndpointIDQueue(device IMTRDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRClusterWaterHeaterMode {
	instance := getMTRClusterWaterHeaterModeClass().Alloc()
	rv := objc.Send[MTRClusterWaterHeaterMode](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWaterHeaterMode/changeToMode(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterWaterHeaterMode) ChangeToModeWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRWaterHeaterModeClusterChangeToModeParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("changeToModeWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWaterHeaterMode/readAttributeAcceptedCommandList(with:)
func (m_ MTRClusterWaterHeaterMode) ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWaterHeaterMode/readAttributeAttributeList(with:)
func (m_ MTRClusterWaterHeaterMode) ReadAttributeAttributeListWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeAttributeListWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWaterHeaterMode/readAttributeClusterRevision(with:)
func (m_ MTRClusterWaterHeaterMode) ReadAttributeClusterRevisionWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeClusterRevisionWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWaterHeaterMode/readAttributeCurrentMode(with:)
func (m_ MTRClusterWaterHeaterMode) ReadAttributeCurrentModeWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeCurrentModeWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWaterHeaterMode/readAttributeFeatureMap(with:)
func (m_ MTRClusterWaterHeaterMode) ReadAttributeFeatureMapWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeFeatureMapWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWaterHeaterMode/readAttributeGeneratedCommandList(with:)
func (m_ MTRClusterWaterHeaterMode) ReadAttributeGeneratedCommandListWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWaterHeaterMode/readAttributeSupportedModes(with:)
func (m_ MTRClusterWaterHeaterMode) ReadAttributeSupportedModesWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeSupportedModesWithParams:"), params)
	return rv
}


