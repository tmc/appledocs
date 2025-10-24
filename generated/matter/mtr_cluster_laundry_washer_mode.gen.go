// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRClusterLaundryWasherMode] class.
var (
	MTRClusterLaundryWasherModeClass     _MTRClusterLaundryWasherModeClass
	MTRClusterLaundryWasherModeClassOnce sync.Once
)

func getMTRClusterLaundryWasherModeClass() _MTRClusterLaundryWasherModeClass {
	MTRClusterLaundryWasherModeClassOnce.Do(func() {
		MTRClusterLaundryWasherModeClass = _MTRClusterLaundryWasherModeClass{objc.GetClass("MTRClusterLaundryWasherMode")}
	})
	return MTRClusterLaundryWasherModeClass
}

type _MTRClusterLaundryWasherModeClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterLaundryWasherMode] class.
type IMTRClusterLaundryWasherMode interface {
	IMTRGenericCluster
	// properties:
	// methods:
	ChangeToModeWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRLaundryWasherModeClusterChangeToModeParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeAttributeListWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeClusterRevisionWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeCurrentModeWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeFeatureMapWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeGeneratedCommandListWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeSupportedModesWithParams(params IMTRReadParams) foundation.IDictionary
}

// Cluster Laundry Washer Mode Attributes and commands for selecting a mode from a list of supported options.


// Cluster Laundry Washer Mode Attributes and commands for selecting a mode from a list of supported options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLaundryWasherMode
type MTRClusterLaundryWasherMode struct {
	MTRGenericCluster
}

// MTRClusterLaundryWasherModeFrom constructs a [MTRClusterLaundryWasherMode] from an unsafe.Pointer.
//
// Cluster Laundry Washer Mode Attributes and commands for selecting a mode from a list of supported options.
func MTRClusterLaundryWasherModeFrom(ptr unsafe.Pointer) MTRClusterLaundryWasherMode {
	return MTRClusterLaundryWasherMode{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterLaundryWasherModeClass) Alloc() MTRClusterLaundryWasherMode {
	rv := objc.Send[MTRClusterLaundryWasherMode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterLaundryWasherModeClass) New() MTRClusterLaundryWasherMode {
	rv := objc.Send[MTRClusterLaundryWasherMode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterLaundryWasherMode) Init() MTRClusterLaundryWasherMode {
	rv := objc.Send[MTRClusterLaundryWasherMode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterLaundryWasherMode) Autorelease() MTRClusterLaundryWasherMode {
	rv := objc.Send[MTRClusterLaundryWasherMode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterLaundryWasherMode creates a new MTRClusterLaundryWasherMode instance.
func NewMTRClusterLaundryWasherMode() MTRClusterLaundryWasherMode {
	return getMTRClusterLaundryWasherModeClass().New()
}



// For all instance methods that take a completion (i.e. command invocations), the completion will be called on the provided queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLaundryWasherMode/init(device:endpointID:queue:)
func NewMTRClusterLaundryWasherModeWithDeviceEndpointIDQueue(device IMTRDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRClusterLaundryWasherMode {
	instance := getMTRClusterLaundryWasherModeClass().Alloc()
	rv := objc.Send[MTRClusterLaundryWasherMode](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLaundryWasherMode/changeToMode(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterLaundryWasherMode) ChangeToModeWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRLaundryWasherModeClusterChangeToModeParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("changeToModeWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLaundryWasherMode/readAttributeAcceptedCommandList(with:)
func (m_ MTRClusterLaundryWasherMode) ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLaundryWasherMode/readAttributeAttributeList(with:)
func (m_ MTRClusterLaundryWasherMode) ReadAttributeAttributeListWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeAttributeListWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLaundryWasherMode/readAttributeClusterRevision(with:)
func (m_ MTRClusterLaundryWasherMode) ReadAttributeClusterRevisionWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeClusterRevisionWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLaundryWasherMode/readAttributeCurrentMode(with:)
func (m_ MTRClusterLaundryWasherMode) ReadAttributeCurrentModeWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeCurrentModeWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLaundryWasherMode/readAttributeFeatureMap(with:)
func (m_ MTRClusterLaundryWasherMode) ReadAttributeFeatureMapWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeFeatureMapWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLaundryWasherMode/readAttributeGeneratedCommandList(with:)
func (m_ MTRClusterLaundryWasherMode) ReadAttributeGeneratedCommandListWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLaundryWasherMode/readAttributeSupportedModes(with:)
func (m_ MTRClusterLaundryWasherMode) ReadAttributeSupportedModesWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeSupportedModesWithParams:"), params)
	return rv
}


