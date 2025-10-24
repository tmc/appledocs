// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRClusterTemperatureControl] class.
var (
	MTRClusterTemperatureControlClass     _MTRClusterTemperatureControlClass
	MTRClusterTemperatureControlClassOnce sync.Once
)

func getMTRClusterTemperatureControlClass() _MTRClusterTemperatureControlClass {
	MTRClusterTemperatureControlClassOnce.Do(func() {
		MTRClusterTemperatureControlClass = _MTRClusterTemperatureControlClass{objc.GetClass("MTRClusterTemperatureControl")}
	})
	return MTRClusterTemperatureControlClass
}

type _MTRClusterTemperatureControlClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterTemperatureControl] class.
type IMTRClusterTemperatureControl interface {
	IMTRGenericCluster
	// properties:
	// methods:
	ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeAttributeListWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeClusterRevisionWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeFeatureMapWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeGeneratedCommandListWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeMaxTemperatureWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeMinTemperatureWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeSelectedTemperatureLevelWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeStepWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeSupportedTemperatureLevelsWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeTemperatureSetpointWithParams(params IMTRReadParams) foundation.IDictionary
	SetTemperatureWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRTemperatureControlClusterSetTemperatureParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	SetTemperatureWithExpectedValuesExpectedValueIntervalCompletion(expectedValues foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
}

// Cluster Temperature Control Attributes and commands for configuring the temperature control, and reporting temperature.


// Cluster Temperature Control Attributes and commands for configuring the temperature control, and reporting temperature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTemperatureControl
type MTRClusterTemperatureControl struct {
	MTRGenericCluster
}

// MTRClusterTemperatureControlFrom constructs a [MTRClusterTemperatureControl] from an unsafe.Pointer.
//
// Cluster Temperature Control Attributes and commands for configuring the temperature control, and reporting temperature.
func MTRClusterTemperatureControlFrom(ptr unsafe.Pointer) MTRClusterTemperatureControl {
	return MTRClusterTemperatureControl{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterTemperatureControlClass) Alloc() MTRClusterTemperatureControl {
	rv := objc.Send[MTRClusterTemperatureControl](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterTemperatureControlClass) New() MTRClusterTemperatureControl {
	rv := objc.Send[MTRClusterTemperatureControl](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterTemperatureControl) Init() MTRClusterTemperatureControl {
	rv := objc.Send[MTRClusterTemperatureControl](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterTemperatureControl) Autorelease() MTRClusterTemperatureControl {
	rv := objc.Send[MTRClusterTemperatureControl](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterTemperatureControl creates a new MTRClusterTemperatureControl instance.
func NewMTRClusterTemperatureControl() MTRClusterTemperatureControl {
	return getMTRClusterTemperatureControlClass().New()
}



// For all instance methods that take a completion (i.e. command invocations), the completion will be called on the provided queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTemperatureControl/init(device:endpointID:queue:)
func NewMTRClusterTemperatureControlWithDeviceEndpointIDQueue(device IMTRDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRClusterTemperatureControl {
	instance := getMTRClusterTemperatureControlClass().Alloc()
	rv := objc.Send[MTRClusterTemperatureControl](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTemperatureControl/readAttributeAcceptedCommandList(with:)
func (m_ MTRClusterTemperatureControl) ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTemperatureControl/readAttributeAttributeList(with:)
func (m_ MTRClusterTemperatureControl) ReadAttributeAttributeListWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeAttributeListWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTemperatureControl/readAttributeClusterRevision(with:)
func (m_ MTRClusterTemperatureControl) ReadAttributeClusterRevisionWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeClusterRevisionWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTemperatureControl/readAttributeFeatureMap(with:)
func (m_ MTRClusterTemperatureControl) ReadAttributeFeatureMapWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeFeatureMapWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTemperatureControl/readAttributeGeneratedCommandList(with:)
func (m_ MTRClusterTemperatureControl) ReadAttributeGeneratedCommandListWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTemperatureControl/readAttributeMaxTemperature(with:)
func (m_ MTRClusterTemperatureControl) ReadAttributeMaxTemperatureWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeMaxTemperatureWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTemperatureControl/readAttributeMinTemperature(with:)
func (m_ MTRClusterTemperatureControl) ReadAttributeMinTemperatureWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeMinTemperatureWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTemperatureControl/readAttributeSelectedTemperatureLevel(with:)
func (m_ MTRClusterTemperatureControl) ReadAttributeSelectedTemperatureLevelWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeSelectedTemperatureLevelWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTemperatureControl/readAttributeStep(with:)
func (m_ MTRClusterTemperatureControl) ReadAttributeStepWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeStepWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTemperatureControl/readAttributeSupportedTemperatureLevels(with:)
func (m_ MTRClusterTemperatureControl) ReadAttributeSupportedTemperatureLevelsWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeSupportedTemperatureLevelsWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTemperatureControl/readAttributeTemperatureSetpoint(with:)
func (m_ MTRClusterTemperatureControl) ReadAttributeTemperatureSetpointWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeTemperatureSetpointWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTemperatureControl/setTemperatureWith(_:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterTemperatureControl) SetTemperatureWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRTemperatureControlClusterSetTemperatureParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTemperatureWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTemperatureControl/setTemperatureWithExpectedValues(_:expectedValueInterval:completion:)
func (m_ MTRClusterTemperatureControl) SetTemperatureWithExpectedValuesExpectedValueIntervalCompletion(expectedValues foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTemperatureWithExpectedValues:expectedValueInterval:completion:"), expectedValues, expectedValueIntervalMs, completion)
}


