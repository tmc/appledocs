// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterOvenMode] class.
var (
	MTRClusterOvenModeClass     _MTRClusterOvenModeClass
	MTRClusterOvenModeClassOnce sync.Once
)

func getMTRClusterOvenModeClass() _MTRClusterOvenModeClass {
	MTRClusterOvenModeClassOnce.Do(func() {
		MTRClusterOvenModeClass = _MTRClusterOvenModeClass{objc.GetClass("MTRClusterOvenMode")}
	})
	return MTRClusterOvenModeClass
}

type _MTRClusterOvenModeClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterOvenMode] class.
type IMTRClusterOvenMode interface {
	IMTRGenericCluster
	ChangeToModeWithParamsExpectedValuesExpectedValueIntervalCompletion(params unsafe.Pointer, expectedDataValueDictionaries unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer, completion unsafe.Pointer)
	ReadAttributeAcceptedCommandListWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeAttributeListWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeClusterRevisionWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeCurrentModeWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeFeatureMapWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeGeneratedCommandListWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeSupportedModesWithParams(params unsafe.Pointer) unsafe.Pointer
}

// Cluster Oven Mode Attributes and commands for selecting a mode from a list of supported options.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterOvenMode
type MTRClusterOvenMode struct {
	MTRGenericCluster
}

// MTRClusterOvenModeFrom constructs a [MTRClusterOvenMode] from an unsafe.Pointer.
//
// Cluster Oven Mode Attributes and commands for selecting a mode from a list of supported options.
func MTRClusterOvenModeFrom(ptr unsafe.Pointer) MTRClusterOvenMode {
	return MTRClusterOvenMode{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterOvenModeClass) Alloc() MTRClusterOvenMode {
	rv := objc.Send[MTRClusterOvenMode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterOvenModeClass) New() MTRClusterOvenMode {
	rv := objc.Send[MTRClusterOvenMode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterOvenMode) Init() MTRClusterOvenMode {
	rv := objc.Send[MTRClusterOvenMode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterOvenMode) Autorelease() MTRClusterOvenMode {
	rv := objc.Send[MTRClusterOvenMode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterOvenMode creates a new MTRClusterOvenMode instance.
func NewMTRClusterOvenMode() MTRClusterOvenMode {
	return getMTRClusterOvenModeClass().New()
}




// For all instance methods that take a completion (i.e. command invocations), the completion will be called on the provided queue.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterOvenMode/init(device:endpointID:queue:)
func NewMTRClusterOvenModeWithDeviceEndpointIDQueue(device unsafe.Pointer, endpointID unsafe.Pointer, queue unsafe.Pointer) MTRClusterOvenMode {
	instance := getMTRClusterOvenModeClass().Alloc()
	rv := objc.Send[MTRClusterOvenMode](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterOvenMode/changeToMode(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterOvenMode) ChangeToModeWithParamsExpectedValuesExpectedValueIntervalCompletion(params unsafe.Pointer, expectedDataValueDictionaries unsafe.Pointer, expectedValueIntervalMs unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("changeToModeWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterOvenMode/readAttributeAcceptedCommandList(with:)
func (m_ MTRClusterOvenMode) ReadAttributeAcceptedCommandListWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterOvenMode/readAttributeAttributeList(with:)
func (m_ MTRClusterOvenMode) ReadAttributeAttributeListWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeAttributeListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterOvenMode/readAttributeClusterRevision(with:)
func (m_ MTRClusterOvenMode) ReadAttributeClusterRevisionWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeClusterRevisionWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterOvenMode/readAttributeCurrentMode(with:)
func (m_ MTRClusterOvenMode) ReadAttributeCurrentModeWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeCurrentModeWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterOvenMode/readAttributeFeatureMap(with:)
func (m_ MTRClusterOvenMode) ReadAttributeFeatureMapWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeFeatureMapWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterOvenMode/readAttributeGeneratedCommandList(with:)
func (m_ MTRClusterOvenMode) ReadAttributeGeneratedCommandListWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterOvenMode/readAttributeSupportedModes(with:)
func (m_ MTRClusterOvenMode) ReadAttributeSupportedModesWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeSupportedModesWithParams:"), params)
	return rv
}


