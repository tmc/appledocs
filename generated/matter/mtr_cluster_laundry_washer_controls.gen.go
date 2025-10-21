// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterLaundryWasherControls] class.
var (
	MTRClusterLaundryWasherControlsClass     _MTRClusterLaundryWasherControlsClass
	MTRClusterLaundryWasherControlsClassOnce sync.Once
)

func getMTRClusterLaundryWasherControlsClass() _MTRClusterLaundryWasherControlsClass {
	MTRClusterLaundryWasherControlsClassOnce.Do(func() {
		MTRClusterLaundryWasherControlsClass = _MTRClusterLaundryWasherControlsClass{objc.GetClass("MTRClusterLaundryWasherControls")}
	})
	return MTRClusterLaundryWasherControlsClass
}

type _MTRClusterLaundryWasherControlsClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterLaundryWasherControls] class.
type IMTRClusterLaundryWasherControls interface {
	IMTRGenericCluster
	ReadAttributeAcceptedCommandListWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeAttributeListWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeClusterRevisionWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeFeatureMapWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeGeneratedCommandListWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeNumberOfRinsesWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeSpinSpeedCurrentWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeSpinSpeedsWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeSupportedRinsesWithParams(params unsafe.Pointer) unsafe.Pointer
	WriteAttributeNumberOfRinsesWithValueExpectedValueInterval(dataValueDictionary unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer)
	WriteAttributeNumberOfRinsesWithValueExpectedValueIntervalParams(dataValueDictionary unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer, params unsafe.Pointer)
	WriteAttributeSpinSpeedCurrentWithValueExpectedValueInterval(dataValueDictionary unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer)
	WriteAttributeSpinSpeedCurrentWithValueExpectedValueIntervalParams(dataValueDictionary unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer, params unsafe.Pointer)
}

// Cluster Laundry Washer Controls This cluster supports remotely monitoring and controlling the different types of functionality available to a washing device, such as a washing machine.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLaundryWasherControls
type MTRClusterLaundryWasherControls struct {
	MTRGenericCluster
}

// MTRClusterLaundryWasherControlsFrom constructs a [MTRClusterLaundryWasherControls] from an unsafe.Pointer.
//
// Cluster Laundry Washer Controls This cluster supports remotely monitoring and controlling the different types of functionality available to a washing device, such as a washing machine.
func MTRClusterLaundryWasherControlsFrom(ptr unsafe.Pointer) MTRClusterLaundryWasherControls {
	return MTRClusterLaundryWasherControls{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterLaundryWasherControlsClass) Alloc() MTRClusterLaundryWasherControls {
	rv := objc.Send[MTRClusterLaundryWasherControls](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterLaundryWasherControlsClass) New() MTRClusterLaundryWasherControls {
	rv := objc.Send[MTRClusterLaundryWasherControls](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterLaundryWasherControls) Init() MTRClusterLaundryWasherControls {
	rv := objc.Send[MTRClusterLaundryWasherControls](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterLaundryWasherControls) Autorelease() MTRClusterLaundryWasherControls {
	rv := objc.Send[MTRClusterLaundryWasherControls](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterLaundryWasherControls creates a new MTRClusterLaundryWasherControls instance.
func NewMTRClusterLaundryWasherControls() MTRClusterLaundryWasherControls {
	return getMTRClusterLaundryWasherControlsClass().New()
}


// The queue is currently unused, but may be used in the future for calling completions for command invocations if commands are added to this cluster.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLaundryWasherControls/init(device:endpointID:queue:)
func NewMTRClusterLaundryWasherControlsWithDeviceEndpointIDQueue(device unsafe.Pointer, endpointID unsafe.Pointer, queue unsafe.Pointer) MTRClusterLaundryWasherControls {
	instance := getMTRClusterLaundryWasherControlsClass().Alloc()
	rv := objc.Send[MTRClusterLaundryWasherControls](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLaundryWasherControls/readAttributeAcceptedCommandList(with:)
func (m_ MTRClusterLaundryWasherControls) ReadAttributeAcceptedCommandListWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLaundryWasherControls/readAttributeAttributeList(with:)
func (m_ MTRClusterLaundryWasherControls) ReadAttributeAttributeListWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeAttributeListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLaundryWasherControls/readAttributeClusterRevision(with:)
func (m_ MTRClusterLaundryWasherControls) ReadAttributeClusterRevisionWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeClusterRevisionWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLaundryWasherControls/readAttributeFeatureMap(with:)
func (m_ MTRClusterLaundryWasherControls) ReadAttributeFeatureMapWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeFeatureMapWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLaundryWasherControls/readAttributeGeneratedCommandList(with:)
func (m_ MTRClusterLaundryWasherControls) ReadAttributeGeneratedCommandListWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLaundryWasherControls/readAttributeNumberOfRinses(with:)
func (m_ MTRClusterLaundryWasherControls) ReadAttributeNumberOfRinsesWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeNumberOfRinsesWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLaundryWasherControls/readAttributeSpinSpeedCurrent(with:)
func (m_ MTRClusterLaundryWasherControls) ReadAttributeSpinSpeedCurrentWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeSpinSpeedCurrentWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLaundryWasherControls/readAttributeSpinSpeeds(with:)
func (m_ MTRClusterLaundryWasherControls) ReadAttributeSpinSpeedsWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeSpinSpeedsWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLaundryWasherControls/readAttributeSupportedRinses(with:)
func (m_ MTRClusterLaundryWasherControls) ReadAttributeSupportedRinsesWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeSupportedRinsesWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLaundryWasherControls/writeAttributeNumberOfRinses(withValue:expectedValueInterval:)
func (m_ MTRClusterLaundryWasherControls) WriteAttributeNumberOfRinsesWithValueExpectedValueInterval(dataValueDictionary unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeNumberOfRinsesWithValue:expectedValueInterval:"), dataValueDictionary, expectedValueIntervalMs)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLaundryWasherControls/writeAttributeNumberOfRinses(withValue:expectedValueInterval:params:)
func (m_ MTRClusterLaundryWasherControls) WriteAttributeNumberOfRinsesWithValueExpectedValueIntervalParams(dataValueDictionary unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer, params unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeNumberOfRinsesWithValue:expectedValueInterval:params:"), dataValueDictionary, expectedValueIntervalMs, params)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLaundryWasherControls/writeAttributeSpinSpeedCurrent(withValue:expectedValueInterval:)
func (m_ MTRClusterLaundryWasherControls) WriteAttributeSpinSpeedCurrentWithValueExpectedValueInterval(dataValueDictionary unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeSpinSpeedCurrentWithValue:expectedValueInterval:"), dataValueDictionary, expectedValueIntervalMs)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLaundryWasherControls/writeAttributeSpinSpeedCurrent(withValue:expectedValueInterval:params:)
func (m_ MTRClusterLaundryWasherControls) WriteAttributeSpinSpeedCurrentWithValueExpectedValueIntervalParams(dataValueDictionary unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer, params unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeSpinSpeedCurrentWithValue:expectedValueInterval:params:"), dataValueDictionary, expectedValueIntervalMs, params)
}


