// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRClusterLaundryDryerControls] class.
var (
	MTRClusterLaundryDryerControlsClass     _MTRClusterLaundryDryerControlsClass
	MTRClusterLaundryDryerControlsClassOnce sync.Once
)

func getMTRClusterLaundryDryerControlsClass() _MTRClusterLaundryDryerControlsClass {
	MTRClusterLaundryDryerControlsClassOnce.Do(func() {
		MTRClusterLaundryDryerControlsClass = _MTRClusterLaundryDryerControlsClass{objc.GetClass("MTRClusterLaundryDryerControls")}
	})
	return MTRClusterLaundryDryerControlsClass
}

type _MTRClusterLaundryDryerControlsClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterLaundryDryerControls] class.
type IMTRClusterLaundryDryerControls interface {
	IMTRGenericCluster
	ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeAttributeListWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeClusterRevisionWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeFeatureMapWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeGeneratedCommandListWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeSelectedDrynessLevelWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeSupportedDrynessLevelsWithParams(params IMTRReadParams) unsafe.Pointer
	WriteAttributeSelectedDrynessLevelWithValueExpectedValueInterval(dataValueDictionary unsafe.Pointer, expectedValueIntervalMs foundation.INumber)
	WriteAttributeSelectedDrynessLevelWithValueExpectedValueIntervalParams(dataValueDictionary unsafe.Pointer, expectedValueIntervalMs foundation.INumber, params IMTRWriteParams)
}

// Cluster Laundry Dryer Controls This cluster provides a way to access options associated with the operation of a laundry dryer device type.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLaundryDryerControls
type MTRClusterLaundryDryerControls struct {
	MTRGenericCluster
}

// MTRClusterLaundryDryerControlsFrom constructs a [MTRClusterLaundryDryerControls] from an unsafe.Pointer.
//
// Cluster Laundry Dryer Controls This cluster provides a way to access options associated with the operation of a laundry dryer device type.
func MTRClusterLaundryDryerControlsFrom(ptr unsafe.Pointer) MTRClusterLaundryDryerControls {
	return MTRClusterLaundryDryerControls{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterLaundryDryerControlsClass) Alloc() MTRClusterLaundryDryerControls {
	rv := objc.Send[MTRClusterLaundryDryerControls](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterLaundryDryerControlsClass) New() MTRClusterLaundryDryerControls {
	rv := objc.Send[MTRClusterLaundryDryerControls](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterLaundryDryerControls) Init() MTRClusterLaundryDryerControls {
	rv := objc.Send[MTRClusterLaundryDryerControls](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterLaundryDryerControls) Autorelease() MTRClusterLaundryDryerControls {
	rv := objc.Send[MTRClusterLaundryDryerControls](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterLaundryDryerControls creates a new MTRClusterLaundryDryerControls instance.
func NewMTRClusterLaundryDryerControls() MTRClusterLaundryDryerControls {
	return getMTRClusterLaundryDryerControlsClass().New()
}




// The queue is currently unused, but may be used in the future for calling completions for command invocations if commands are added to this cluster.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLaundryDryerControls/init(device:endpointID:queue:)
func NewMTRClusterLaundryDryerControlsWithDeviceEndpointIDQueue(device IMTRDevice, endpointID foundation.INumber, queue unsafe.Pointer) MTRClusterLaundryDryerControls {
	instance := getMTRClusterLaundryDryerControlsClass().Alloc()
	rv := objc.Send[MTRClusterLaundryDryerControls](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLaundryDryerControls/readAttributeAcceptedCommandList(with:)
func (m_ MTRClusterLaundryDryerControls) ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLaundryDryerControls/readAttributeAttributeList(with:)
func (m_ MTRClusterLaundryDryerControls) ReadAttributeAttributeListWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeAttributeListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLaundryDryerControls/readAttributeClusterRevision(with:)
func (m_ MTRClusterLaundryDryerControls) ReadAttributeClusterRevisionWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeClusterRevisionWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLaundryDryerControls/readAttributeFeatureMap(with:)
func (m_ MTRClusterLaundryDryerControls) ReadAttributeFeatureMapWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeFeatureMapWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLaundryDryerControls/readAttributeGeneratedCommandList(with:)
func (m_ MTRClusterLaundryDryerControls) ReadAttributeGeneratedCommandListWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLaundryDryerControls/readAttributeSelectedDrynessLevel(with:)
func (m_ MTRClusterLaundryDryerControls) ReadAttributeSelectedDrynessLevelWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeSelectedDrynessLevelWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLaundryDryerControls/readAttributeSupportedDrynessLevels(with:)
func (m_ MTRClusterLaundryDryerControls) ReadAttributeSupportedDrynessLevelsWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeSupportedDrynessLevelsWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLaundryDryerControls/writeAttributeSelectedDrynessLevel(withValue:expectedValueInterval:)
func (m_ MTRClusterLaundryDryerControls) WriteAttributeSelectedDrynessLevelWithValueExpectedValueInterval(dataValueDictionary unsafe.Pointer, expectedValueIntervalMs foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeSelectedDrynessLevelWithValue:expectedValueInterval:"), dataValueDictionary, expectedValueIntervalMs)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLaundryDryerControls/writeAttributeSelectedDrynessLevel(withValue:expectedValueInterval:params:)
func (m_ MTRClusterLaundryDryerControls) WriteAttributeSelectedDrynessLevelWithValueExpectedValueIntervalParams(dataValueDictionary unsafe.Pointer, expectedValueIntervalMs foundation.INumber, params IMTRWriteParams) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeSelectedDrynessLevelWithValue:expectedValueInterval:params:"), dataValueDictionary, expectedValueIntervalMs, params)
}


