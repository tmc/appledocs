// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRClusterDishwasherAlarm] class.
var (
	MTRClusterDishwasherAlarmClass     _MTRClusterDishwasherAlarmClass
	MTRClusterDishwasherAlarmClassOnce sync.Once
)

func getMTRClusterDishwasherAlarmClass() _MTRClusterDishwasherAlarmClass {
	MTRClusterDishwasherAlarmClassOnce.Do(func() {
		MTRClusterDishwasherAlarmClass = _MTRClusterDishwasherAlarmClass{objc.GetClass("MTRClusterDishwasherAlarm")}
	})
	return MTRClusterDishwasherAlarmClass
}

type _MTRClusterDishwasherAlarmClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterDishwasherAlarm] class.
type IMTRClusterDishwasherAlarm interface {
	IMTRGenericCluster
	// properties:
	// methods:
	ModifyEnabledAlarmsWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRDishwasherAlarmClusterModifyEnabledAlarmsParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeAttributeListWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeClusterRevisionWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeFeatureMapWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeGeneratedCommandListWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeLatchWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeMaskWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeStateWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeSupportedWithParams(params IMTRReadParams) foundation.IDictionary
	ResetWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRDishwasherAlarmClusterResetParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
}

// Cluster Dishwasher Alarm Attributes and commands for configuring the Dishwasher alarm.


// Cluster Dishwasher Alarm Attributes and commands for configuring the Dishwasher alarm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDishwasherAlarm
type MTRClusterDishwasherAlarm struct {
	MTRGenericCluster
}

// MTRClusterDishwasherAlarmFrom constructs a [MTRClusterDishwasherAlarm] from an unsafe.Pointer.
//
// Cluster Dishwasher Alarm Attributes and commands for configuring the Dishwasher alarm.
func MTRClusterDishwasherAlarmFrom(ptr unsafe.Pointer) MTRClusterDishwasherAlarm {
	return MTRClusterDishwasherAlarm{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterDishwasherAlarmClass) Alloc() MTRClusterDishwasherAlarm {
	rv := objc.Send[MTRClusterDishwasherAlarm](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterDishwasherAlarmClass) New() MTRClusterDishwasherAlarm {
	rv := objc.Send[MTRClusterDishwasherAlarm](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterDishwasherAlarm) Init() MTRClusterDishwasherAlarm {
	rv := objc.Send[MTRClusterDishwasherAlarm](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterDishwasherAlarm) Autorelease() MTRClusterDishwasherAlarm {
	rv := objc.Send[MTRClusterDishwasherAlarm](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterDishwasherAlarm creates a new MTRClusterDishwasherAlarm instance.
func NewMTRClusterDishwasherAlarm() MTRClusterDishwasherAlarm {
	return getMTRClusterDishwasherAlarmClass().New()
}



// For all instance methods that take a completion (i.e. command invocations), the completion will be called on the provided queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDishwasherAlarm/init(device:endpointID:queue:)
func NewMTRClusterDishwasherAlarmWithDeviceEndpointIDQueue(device IMTRDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRClusterDishwasherAlarm {
	instance := getMTRClusterDishwasherAlarmClass().Alloc()
	rv := objc.Send[MTRClusterDishwasherAlarm](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDishwasherAlarm/modifyEnabledAlarms(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterDishwasherAlarm) ModifyEnabledAlarmsWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRDishwasherAlarmClusterModifyEnabledAlarmsParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("modifyEnabledAlarmsWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDishwasherAlarm/readAttributeAcceptedCommandList(with:)
func (m_ MTRClusterDishwasherAlarm) ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDishwasherAlarm/readAttributeAttributeList(with:)
func (m_ MTRClusterDishwasherAlarm) ReadAttributeAttributeListWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeAttributeListWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDishwasherAlarm/readAttributeClusterRevision(with:)
func (m_ MTRClusterDishwasherAlarm) ReadAttributeClusterRevisionWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeClusterRevisionWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDishwasherAlarm/readAttributeFeatureMap(with:)
func (m_ MTRClusterDishwasherAlarm) ReadAttributeFeatureMapWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeFeatureMapWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDishwasherAlarm/readAttributeGeneratedCommandList(with:)
func (m_ MTRClusterDishwasherAlarm) ReadAttributeGeneratedCommandListWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDishwasherAlarm/readAttributeLatch(with:)
func (m_ MTRClusterDishwasherAlarm) ReadAttributeLatchWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeLatchWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDishwasherAlarm/readAttributeMask(with:)
func (m_ MTRClusterDishwasherAlarm) ReadAttributeMaskWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeMaskWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDishwasherAlarm/readAttributeState(with:)
func (m_ MTRClusterDishwasherAlarm) ReadAttributeStateWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeStateWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDishwasherAlarm/readAttributeSupported(with:)
func (m_ MTRClusterDishwasherAlarm) ReadAttributeSupportedWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeSupportedWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDishwasherAlarm/reset(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterDishwasherAlarm) ResetWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRDishwasherAlarmClusterResetParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("resetWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}


