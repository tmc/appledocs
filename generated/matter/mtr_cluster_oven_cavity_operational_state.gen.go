// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRClusterOvenCavityOperationalState] class.
var (
	MTRClusterOvenCavityOperationalStateClass     _MTRClusterOvenCavityOperationalStateClass
	MTRClusterOvenCavityOperationalStateClassOnce sync.Once
)

func getMTRClusterOvenCavityOperationalStateClass() _MTRClusterOvenCavityOperationalStateClass {
	MTRClusterOvenCavityOperationalStateClassOnce.Do(func() {
		MTRClusterOvenCavityOperationalStateClass = _MTRClusterOvenCavityOperationalStateClass{objc.GetClass("MTRClusterOvenCavityOperationalState")}
	})
	return MTRClusterOvenCavityOperationalStateClass
}

type _MTRClusterOvenCavityOperationalStateClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterOvenCavityOperationalState] class.
type IMTRClusterOvenCavityOperationalState interface {
	IMTRGenericCluster
	ReadAttributeAcceptedCommandListWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeAttributeListWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeClusterRevisionWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeCountdownTimeWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeCurrentPhaseWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeFeatureMapWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeGeneratedCommandListWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeOperationalErrorWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeOperationalStateWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeOperationalStateListWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributePhaseListWithParams(params unsafe.Pointer) unsafe.Pointer
	StartWithParamsExpectedValuesExpectedValueIntervalCompletion(params unsafe.Pointer, expectedDataValueDictionaries unsafe.Pointer, expectedValueIntervalMs foundation.Number, completion unsafe.Pointer)
	StartWithExpectedValuesExpectedValueIntervalCompletion(expectedValues unsafe.Pointer, expectedValueIntervalMs foundation.Number, completion unsafe.Pointer)
	StopWithParamsExpectedValuesExpectedValueIntervalCompletion(params unsafe.Pointer, expectedDataValueDictionaries unsafe.Pointer, expectedValueIntervalMs foundation.Number, completion unsafe.Pointer)
	StopWithExpectedValuesExpectedValueIntervalCompletion(expectedValues unsafe.Pointer, expectedValueIntervalMs foundation.Number, completion unsafe.Pointer)
}

// Cluster Oven Cavity Operational State This cluster supports remotely monitoring and, where supported, changing the operational state of an Oven.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterOvenCavityOperationalState
type MTRClusterOvenCavityOperationalState struct {
	MTRGenericCluster
}

// MTRClusterOvenCavityOperationalStateFrom constructs a [MTRClusterOvenCavityOperationalState] from an unsafe.Pointer.
//
// Cluster Oven Cavity Operational State This cluster supports remotely monitoring and, where supported, changing the operational state of an Oven.
func MTRClusterOvenCavityOperationalStateFrom(ptr unsafe.Pointer) MTRClusterOvenCavityOperationalState {
	return MTRClusterOvenCavityOperationalState{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterOvenCavityOperationalStateClass) Alloc() MTRClusterOvenCavityOperationalState {
	rv := objc.Send[MTRClusterOvenCavityOperationalState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterOvenCavityOperationalStateClass) New() MTRClusterOvenCavityOperationalState {
	rv := objc.Send[MTRClusterOvenCavityOperationalState](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterOvenCavityOperationalState) Init() MTRClusterOvenCavityOperationalState {
	rv := objc.Send[MTRClusterOvenCavityOperationalState](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterOvenCavityOperationalState) Autorelease() MTRClusterOvenCavityOperationalState {
	rv := objc.Send[MTRClusterOvenCavityOperationalState](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterOvenCavityOperationalState creates a new MTRClusterOvenCavityOperationalState instance.
func NewMTRClusterOvenCavityOperationalState() MTRClusterOvenCavityOperationalState {
	return getMTRClusterOvenCavityOperationalStateClass().New()
}




// For all instance methods that take a completion (i.e. command invocations), the completion will be called on the provided queue.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterOvenCavityOperationalState/init(device:endpointID:queue:)
func NewMTRClusterOvenCavityOperationalStateWithDeviceEndpointIDQueue(device unsafe.Pointer, endpointID foundation.Number, queue unsafe.Pointer) MTRClusterOvenCavityOperationalState {
	instance := getMTRClusterOvenCavityOperationalStateClass().Alloc()
	rv := objc.Send[MTRClusterOvenCavityOperationalState](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterOvenCavityOperationalState/readAttributeAcceptedCommandList(with:)
func (m_ MTRClusterOvenCavityOperationalState) ReadAttributeAcceptedCommandListWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterOvenCavityOperationalState/readAttributeAttributeList(with:)
func (m_ MTRClusterOvenCavityOperationalState) ReadAttributeAttributeListWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeAttributeListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterOvenCavityOperationalState/readAttributeClusterRevision(with:)
func (m_ MTRClusterOvenCavityOperationalState) ReadAttributeClusterRevisionWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeClusterRevisionWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterOvenCavityOperationalState/readAttributeCountdownTime(with:)
func (m_ MTRClusterOvenCavityOperationalState) ReadAttributeCountdownTimeWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeCountdownTimeWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterOvenCavityOperationalState/readAttributeCurrentPhase(with:)
func (m_ MTRClusterOvenCavityOperationalState) ReadAttributeCurrentPhaseWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeCurrentPhaseWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterOvenCavityOperationalState/readAttributeFeatureMap(with:)
func (m_ MTRClusterOvenCavityOperationalState) ReadAttributeFeatureMapWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeFeatureMapWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterOvenCavityOperationalState/readAttributeGeneratedCommandList(with:)
func (m_ MTRClusterOvenCavityOperationalState) ReadAttributeGeneratedCommandListWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterOvenCavityOperationalState/readAttributeOperationalError(with:)
func (m_ MTRClusterOvenCavityOperationalState) ReadAttributeOperationalErrorWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeOperationalErrorWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterOvenCavityOperationalState/readAttributeOperationalState(with:)
func (m_ MTRClusterOvenCavityOperationalState) ReadAttributeOperationalStateWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeOperationalStateWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterOvenCavityOperationalState/readAttributeOperationalStateList(with:)
func (m_ MTRClusterOvenCavityOperationalState) ReadAttributeOperationalStateListWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeOperationalStateListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterOvenCavityOperationalState/readAttributePhaseList(with:)
func (m_ MTRClusterOvenCavityOperationalState) ReadAttributePhaseListWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributePhaseListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterOvenCavityOperationalState/start(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterOvenCavityOperationalState) StartWithParamsExpectedValuesExpectedValueIntervalCompletion(params unsafe.Pointer, expectedDataValueDictionaries unsafe.Pointer, expectedValueIntervalMs foundation.Number, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("startWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterOvenCavityOperationalState/start(withExpectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterOvenCavityOperationalState) StartWithExpectedValuesExpectedValueIntervalCompletion(expectedValues unsafe.Pointer, expectedValueIntervalMs foundation.Number, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("startWithExpectedValues:expectedValueInterval:completion:"), expectedValues, expectedValueIntervalMs, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterOvenCavityOperationalState/stop(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterOvenCavityOperationalState) StopWithParamsExpectedValuesExpectedValueIntervalCompletion(params unsafe.Pointer, expectedDataValueDictionaries unsafe.Pointer, expectedValueIntervalMs foundation.Number, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("stopWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterOvenCavityOperationalState/stop(withExpectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterOvenCavityOperationalState) StopWithExpectedValuesExpectedValueIntervalCompletion(expectedValues unsafe.Pointer, expectedValueIntervalMs foundation.Number, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("stopWithExpectedValues:expectedValueInterval:completion:"), expectedValues, expectedValueIntervalMs, completion)
}


