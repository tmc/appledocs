// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRClusterRefrigeratorAlarm] class.
var (
	MTRClusterRefrigeratorAlarmClass     _MTRClusterRefrigeratorAlarmClass
	MTRClusterRefrigeratorAlarmClassOnce sync.Once
)

func getMTRClusterRefrigeratorAlarmClass() _MTRClusterRefrigeratorAlarmClass {
	MTRClusterRefrigeratorAlarmClassOnce.Do(func() {
		MTRClusterRefrigeratorAlarmClass = _MTRClusterRefrigeratorAlarmClass{objc.GetClass("MTRClusterRefrigeratorAlarm")}
	})
	return MTRClusterRefrigeratorAlarmClass
}

type _MTRClusterRefrigeratorAlarmClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterRefrigeratorAlarm] class.
type IMTRClusterRefrigeratorAlarm interface {
	IMTRGenericCluster
	// properties:
	// methods:
	ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeAttributeListWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeClusterRevisionWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeFeatureMapWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeGeneratedCommandListWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeMaskWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeStateWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeSupportedWithParams(params IMTRReadParams) foundation.IDictionary
}

// Cluster Refrigerator Alarm Attributes and commands for configuring the Refrigerator alarm.


// Cluster Refrigerator Alarm Attributes and commands for configuring the Refrigerator alarm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterRefrigeratorAlarm
type MTRClusterRefrigeratorAlarm struct {
	MTRGenericCluster
}

// MTRClusterRefrigeratorAlarmFrom constructs a [MTRClusterRefrigeratorAlarm] from an unsafe.Pointer.
//
// Cluster Refrigerator Alarm Attributes and commands for configuring the Refrigerator alarm.
func MTRClusterRefrigeratorAlarmFrom(ptr unsafe.Pointer) MTRClusterRefrigeratorAlarm {
	return MTRClusterRefrigeratorAlarm{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterRefrigeratorAlarmClass) Alloc() MTRClusterRefrigeratorAlarm {
	rv := objc.Send[MTRClusterRefrigeratorAlarm](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterRefrigeratorAlarmClass) New() MTRClusterRefrigeratorAlarm {
	rv := objc.Send[MTRClusterRefrigeratorAlarm](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterRefrigeratorAlarm) Init() MTRClusterRefrigeratorAlarm {
	rv := objc.Send[MTRClusterRefrigeratorAlarm](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterRefrigeratorAlarm) Autorelease() MTRClusterRefrigeratorAlarm {
	rv := objc.Send[MTRClusterRefrigeratorAlarm](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterRefrigeratorAlarm creates a new MTRClusterRefrigeratorAlarm instance.
func NewMTRClusterRefrigeratorAlarm() MTRClusterRefrigeratorAlarm {
	return getMTRClusterRefrigeratorAlarmClass().New()
}



// The queue is currently unused, but may be used in the future for calling completions for command invocations if commands are added to this cluster.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterRefrigeratorAlarm/init(device:endpointID:queue:)
func NewMTRClusterRefrigeratorAlarmWithDeviceEndpointIDQueue(device IMTRDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRClusterRefrigeratorAlarm {
	instance := getMTRClusterRefrigeratorAlarmClass().Alloc()
	rv := objc.Send[MTRClusterRefrigeratorAlarm](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterRefrigeratorAlarm/readAttributeAcceptedCommandList(with:)
func (m_ MTRClusterRefrigeratorAlarm) ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterRefrigeratorAlarm/readAttributeAttributeList(with:)
func (m_ MTRClusterRefrigeratorAlarm) ReadAttributeAttributeListWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeAttributeListWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterRefrigeratorAlarm/readAttributeClusterRevision(with:)
func (m_ MTRClusterRefrigeratorAlarm) ReadAttributeClusterRevisionWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeClusterRevisionWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterRefrigeratorAlarm/readAttributeFeatureMap(with:)
func (m_ MTRClusterRefrigeratorAlarm) ReadAttributeFeatureMapWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeFeatureMapWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterRefrigeratorAlarm/readAttributeGeneratedCommandList(with:)
func (m_ MTRClusterRefrigeratorAlarm) ReadAttributeGeneratedCommandListWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterRefrigeratorAlarm/readAttributeMask(with:)
func (m_ MTRClusterRefrigeratorAlarm) ReadAttributeMaskWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeMaskWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterRefrigeratorAlarm/readAttributeState(with:)
func (m_ MTRClusterRefrigeratorAlarm) ReadAttributeStateWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeStateWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterRefrigeratorAlarm/readAttributeSupported(with:)
func (m_ MTRClusterRefrigeratorAlarm) ReadAttributeSupportedWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeSupportedWithParams:"), params)
	return rv
}


