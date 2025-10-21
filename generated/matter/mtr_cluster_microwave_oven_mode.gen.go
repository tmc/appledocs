// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRClusterMicrowaveOvenMode] class.
var (
	MTRClusterMicrowaveOvenModeClass     _MTRClusterMicrowaveOvenModeClass
	MTRClusterMicrowaveOvenModeClassOnce sync.Once
)

func getMTRClusterMicrowaveOvenModeClass() _MTRClusterMicrowaveOvenModeClass {
	MTRClusterMicrowaveOvenModeClassOnce.Do(func() {
		MTRClusterMicrowaveOvenModeClass = _MTRClusterMicrowaveOvenModeClass{objc.GetClass("MTRClusterMicrowaveOvenMode")}
	})
	return MTRClusterMicrowaveOvenModeClass
}

type _MTRClusterMicrowaveOvenModeClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterMicrowaveOvenMode] class.
type IMTRClusterMicrowaveOvenMode interface {
	IMTRGenericCluster
	ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeAttributeListWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeClusterRevisionWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeCurrentModeWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeFeatureMapWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeGeneratedCommandListWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeSupportedModesWithParams(params IMTRReadParams) unsafe.Pointer
}

// Cluster Microwave Oven Mode Attributes and commands for selecting a mode from a list of supported options.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMicrowaveOvenMode
type MTRClusterMicrowaveOvenMode struct {
	MTRGenericCluster
}

// MTRClusterMicrowaveOvenModeFrom constructs a [MTRClusterMicrowaveOvenMode] from an unsafe.Pointer.
//
// Cluster Microwave Oven Mode Attributes and commands for selecting a mode from a list of supported options.
func MTRClusterMicrowaveOvenModeFrom(ptr unsafe.Pointer) MTRClusterMicrowaveOvenMode {
	return MTRClusterMicrowaveOvenMode{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterMicrowaveOvenModeClass) Alloc() MTRClusterMicrowaveOvenMode {
	rv := objc.Send[MTRClusterMicrowaveOvenMode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterMicrowaveOvenModeClass) New() MTRClusterMicrowaveOvenMode {
	rv := objc.Send[MTRClusterMicrowaveOvenMode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterMicrowaveOvenMode) Init() MTRClusterMicrowaveOvenMode {
	rv := objc.Send[MTRClusterMicrowaveOvenMode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterMicrowaveOvenMode) Autorelease() MTRClusterMicrowaveOvenMode {
	rv := objc.Send[MTRClusterMicrowaveOvenMode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterMicrowaveOvenMode creates a new MTRClusterMicrowaveOvenMode instance.
func NewMTRClusterMicrowaveOvenMode() MTRClusterMicrowaveOvenMode {
	return getMTRClusterMicrowaveOvenModeClass().New()
}




// The queue is currently unused, but may be used in the future for calling completions for command invocations if commands are added to this cluster.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMicrowaveOvenMode/init(device:endpointID:queue:)
func NewMTRClusterMicrowaveOvenModeWithDeviceEndpointIDQueue(device IMTRDevice, endpointID foundation.INumber, queue unsafe.Pointer) MTRClusterMicrowaveOvenMode {
	instance := getMTRClusterMicrowaveOvenModeClass().Alloc()
	rv := objc.Send[MTRClusterMicrowaveOvenMode](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMicrowaveOvenMode/readAttributeAcceptedCommandList(with:)
func (m_ MTRClusterMicrowaveOvenMode) ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMicrowaveOvenMode/readAttributeAttributeList(with:)
func (m_ MTRClusterMicrowaveOvenMode) ReadAttributeAttributeListWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeAttributeListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMicrowaveOvenMode/readAttributeClusterRevision(with:)
func (m_ MTRClusterMicrowaveOvenMode) ReadAttributeClusterRevisionWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeClusterRevisionWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMicrowaveOvenMode/readAttributeCurrentMode(with:)
func (m_ MTRClusterMicrowaveOvenMode) ReadAttributeCurrentModeWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeCurrentModeWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMicrowaveOvenMode/readAttributeFeatureMap(with:)
func (m_ MTRClusterMicrowaveOvenMode) ReadAttributeFeatureMapWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeFeatureMapWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMicrowaveOvenMode/readAttributeGeneratedCommandList(with:)
func (m_ MTRClusterMicrowaveOvenMode) ReadAttributeGeneratedCommandListWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMicrowaveOvenMode/readAttributeSupportedModes(with:)
func (m_ MTRClusterMicrowaveOvenMode) ReadAttributeSupportedModesWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeSupportedModesWithParams:"), params)
	return rv
}


