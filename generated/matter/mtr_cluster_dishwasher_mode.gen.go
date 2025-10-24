// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRClusterDishwasherMode] class.
var (
	MTRClusterDishwasherModeClass     _MTRClusterDishwasherModeClass
	MTRClusterDishwasherModeClassOnce sync.Once
)

func getMTRClusterDishwasherModeClass() _MTRClusterDishwasherModeClass {
	MTRClusterDishwasherModeClassOnce.Do(func() {
		MTRClusterDishwasherModeClass = _MTRClusterDishwasherModeClass{objc.GetClass("MTRClusterDishwasherMode")}
	})
	return MTRClusterDishwasherModeClass
}

type _MTRClusterDishwasherModeClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterDishwasherMode] class.
type IMTRClusterDishwasherMode interface {
	IMTRGenericCluster
	// properties:
	// methods:
	ChangeToModeWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRDishwasherModeClusterChangeToModeParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeAttributeListWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeClusterRevisionWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeCurrentModeWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeFeatureMapWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeGeneratedCommandListWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeSupportedModesWithParams(params IMTRReadParams) foundation.IDictionary
}

// Cluster Dishwasher Mode Attributes and commands for selecting a mode from a list of supported options.


// Cluster Dishwasher Mode Attributes and commands for selecting a mode from a list of supported options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDishwasherMode
type MTRClusterDishwasherMode struct {
	MTRGenericCluster
}

// MTRClusterDishwasherModeFrom constructs a [MTRClusterDishwasherMode] from an unsafe.Pointer.
//
// Cluster Dishwasher Mode Attributes and commands for selecting a mode from a list of supported options.
func MTRClusterDishwasherModeFrom(ptr unsafe.Pointer) MTRClusterDishwasherMode {
	return MTRClusterDishwasherMode{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterDishwasherModeClass) Alloc() MTRClusterDishwasherMode {
	rv := objc.Send[MTRClusterDishwasherMode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterDishwasherModeClass) New() MTRClusterDishwasherMode {
	rv := objc.Send[MTRClusterDishwasherMode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterDishwasherMode) Init() MTRClusterDishwasherMode {
	rv := objc.Send[MTRClusterDishwasherMode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterDishwasherMode) Autorelease() MTRClusterDishwasherMode {
	rv := objc.Send[MTRClusterDishwasherMode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterDishwasherMode creates a new MTRClusterDishwasherMode instance.
func NewMTRClusterDishwasherMode() MTRClusterDishwasherMode {
	return getMTRClusterDishwasherModeClass().New()
}



// For all instance methods that take a completion (i.e. command invocations), the completion will be called on the provided queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDishwasherMode/init(device:endpointID:queue:)
func NewMTRClusterDishwasherModeWithDeviceEndpointIDQueue(device IMTRDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRClusterDishwasherMode {
	instance := getMTRClusterDishwasherModeClass().Alloc()
	rv := objc.Send[MTRClusterDishwasherMode](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDishwasherMode/changeToMode(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterDishwasherMode) ChangeToModeWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRDishwasherModeClusterChangeToModeParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("changeToModeWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDishwasherMode/readAttributeAcceptedCommandList(with:)
func (m_ MTRClusterDishwasherMode) ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDishwasherMode/readAttributeAttributeList(with:)
func (m_ MTRClusterDishwasherMode) ReadAttributeAttributeListWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeAttributeListWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDishwasherMode/readAttributeClusterRevision(with:)
func (m_ MTRClusterDishwasherMode) ReadAttributeClusterRevisionWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeClusterRevisionWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDishwasherMode/readAttributeCurrentMode(with:)
func (m_ MTRClusterDishwasherMode) ReadAttributeCurrentModeWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeCurrentModeWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDishwasherMode/readAttributeFeatureMap(with:)
func (m_ MTRClusterDishwasherMode) ReadAttributeFeatureMapWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeFeatureMapWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDishwasherMode/readAttributeGeneratedCommandList(with:)
func (m_ MTRClusterDishwasherMode) ReadAttributeGeneratedCommandListWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDishwasherMode/readAttributeSupportedModes(with:)
func (m_ MTRClusterDishwasherMode) ReadAttributeSupportedModesWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeSupportedModesWithParams:"), params)
	return rv
}


