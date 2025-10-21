// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterMicrowaveOvenControl] class.
var (
	MTRClusterMicrowaveOvenControlClass     _MTRClusterMicrowaveOvenControlClass
	MTRClusterMicrowaveOvenControlClassOnce sync.Once
)

func getMTRClusterMicrowaveOvenControlClass() _MTRClusterMicrowaveOvenControlClass {
	MTRClusterMicrowaveOvenControlClassOnce.Do(func() {
		MTRClusterMicrowaveOvenControlClass = _MTRClusterMicrowaveOvenControlClass{objc.GetClass("MTRClusterMicrowaveOvenControl")}
	})
	return MTRClusterMicrowaveOvenControlClass
}

type _MTRClusterMicrowaveOvenControlClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterMicrowaveOvenControl] class.
type IMTRClusterMicrowaveOvenControl interface {
	IMTRGenericCluster
	AddMoreTimeWithParamsExpectedValuesExpectedValueIntervalCompletion(params unsafe.Pointer, expectedDataValueDictionaries unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer, completion unsafe.Pointer)
	ReadAttributeAcceptedCommandListWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeAttributeListWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeClusterRevisionWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeCookTimeWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeFeatureMapWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeGeneratedCommandListWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeMaxCookTimeWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeMaxPowerWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeMinPowerWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributePowerSettingWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributePowerStepWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeWattRatingWithParams(params unsafe.Pointer) unsafe.Pointer
	SetCookingParametersWithParamsExpectedValuesExpectedValueIntervalCompletion(params unsafe.Pointer, expectedDataValueDictionaries unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer, completion unsafe.Pointer)
	SetCookingParametersWithExpectedValuesExpectedValueIntervalCompletion(expectedValues unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer, completion unsafe.Pointer)
}

// Cluster Microwave Oven Control Attributes and commands for configuring the microwave oven control, and reporting cooking stats.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMicrowaveOvenControl
type MTRClusterMicrowaveOvenControl struct {
	MTRGenericCluster
}

// MTRClusterMicrowaveOvenControlFrom constructs a [MTRClusterMicrowaveOvenControl] from an unsafe.Pointer.
//
// Cluster Microwave Oven Control Attributes and commands for configuring the microwave oven control, and reporting cooking stats.
func MTRClusterMicrowaveOvenControlFrom(ptr unsafe.Pointer) MTRClusterMicrowaveOvenControl {
	return MTRClusterMicrowaveOvenControl{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterMicrowaveOvenControlClass) Alloc() MTRClusterMicrowaveOvenControl {
	rv := objc.Send[MTRClusterMicrowaveOvenControl](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterMicrowaveOvenControlClass) New() MTRClusterMicrowaveOvenControl {
	rv := objc.Send[MTRClusterMicrowaveOvenControl](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterMicrowaveOvenControl) Init() MTRClusterMicrowaveOvenControl {
	rv := objc.Send[MTRClusterMicrowaveOvenControl](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterMicrowaveOvenControl) Autorelease() MTRClusterMicrowaveOvenControl {
	rv := objc.Send[MTRClusterMicrowaveOvenControl](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterMicrowaveOvenControl creates a new MTRClusterMicrowaveOvenControl instance.
func NewMTRClusterMicrowaveOvenControl() MTRClusterMicrowaveOvenControl {
	return getMTRClusterMicrowaveOvenControlClass().New()
}


// For all instance methods that take a completion (i.e. command invocations), the completion will be called on the provided queue.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMicrowaveOvenControl/init(device:endpointID:queue:)
func NewMTRClusterMicrowaveOvenControlWithDeviceEndpointIDQueue(device unsafe.Pointer, endpointID unsafe.Pointer, queue unsafe.Pointer) MTRClusterMicrowaveOvenControl {
	instance := getMTRClusterMicrowaveOvenControlClass().Alloc()
	rv := objc.Send[MTRClusterMicrowaveOvenControl](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMicrowaveOvenControl/addMoreTime(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterMicrowaveOvenControl) AddMoreTimeWithParamsExpectedValuesExpectedValueIntervalCompletion(params unsafe.Pointer, expectedDataValueDictionaries unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addMoreTimeWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMicrowaveOvenControl/readAttributeAcceptedCommandList(with:)
func (m_ MTRClusterMicrowaveOvenControl) ReadAttributeAcceptedCommandListWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMicrowaveOvenControl/readAttributeAttributeList(with:)
func (m_ MTRClusterMicrowaveOvenControl) ReadAttributeAttributeListWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeAttributeListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMicrowaveOvenControl/readAttributeClusterRevision(with:)
func (m_ MTRClusterMicrowaveOvenControl) ReadAttributeClusterRevisionWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeClusterRevisionWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMicrowaveOvenControl/readAttributeCookTime(with:)
func (m_ MTRClusterMicrowaveOvenControl) ReadAttributeCookTimeWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeCookTimeWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMicrowaveOvenControl/readAttributeFeatureMap(with:)
func (m_ MTRClusterMicrowaveOvenControl) ReadAttributeFeatureMapWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeFeatureMapWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMicrowaveOvenControl/readAttributeGeneratedCommandList(with:)
func (m_ MTRClusterMicrowaveOvenControl) ReadAttributeGeneratedCommandListWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMicrowaveOvenControl/readAttributeMaxCookTime(with:)
func (m_ MTRClusterMicrowaveOvenControl) ReadAttributeMaxCookTimeWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeMaxCookTimeWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMicrowaveOvenControl/readAttributeMaxPower(with:)
func (m_ MTRClusterMicrowaveOvenControl) ReadAttributeMaxPowerWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeMaxPowerWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMicrowaveOvenControl/readAttributeMinPower(with:)
func (m_ MTRClusterMicrowaveOvenControl) ReadAttributeMinPowerWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeMinPowerWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMicrowaveOvenControl/readAttributePowerSetting(with:)
func (m_ MTRClusterMicrowaveOvenControl) ReadAttributePowerSettingWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributePowerSettingWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMicrowaveOvenControl/readAttributePowerStep(with:)
func (m_ MTRClusterMicrowaveOvenControl) ReadAttributePowerStepWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributePowerStepWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMicrowaveOvenControl/readAttributeWattRating(with:)
func (m_ MTRClusterMicrowaveOvenControl) ReadAttributeWattRatingWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeWattRatingWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMicrowaveOvenControl/setCookingParametersWith(_:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterMicrowaveOvenControl) SetCookingParametersWithParamsExpectedValuesExpectedValueIntervalCompletion(params unsafe.Pointer, expectedDataValueDictionaries unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCookingParametersWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMicrowaveOvenControl/setCookingParametersWithExpectedValues(_:expectedValueInterval:completion:)
func (m_ MTRClusterMicrowaveOvenControl) SetCookingParametersWithExpectedValuesExpectedValueIntervalCompletion(expectedValues unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCookingParametersWithExpectedValues:expectedValueInterval:completion:"), expectedValues, expectedValueIntervalMs, completion)
}


