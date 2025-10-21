// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRClusterTimeSynchronization] class.
var (
	MTRClusterTimeSynchronizationClass     _MTRClusterTimeSynchronizationClass
	MTRClusterTimeSynchronizationClassOnce sync.Once
)

func getMTRClusterTimeSynchronizationClass() _MTRClusterTimeSynchronizationClass {
	MTRClusterTimeSynchronizationClassOnce.Do(func() {
		MTRClusterTimeSynchronizationClass = _MTRClusterTimeSynchronizationClass{objc.GetClass("MTRClusterTimeSynchronization")}
	})
	return MTRClusterTimeSynchronizationClass
}

type _MTRClusterTimeSynchronizationClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterTimeSynchronization] class.
type IMTRClusterTimeSynchronization interface {
	IMTRGenericCluster
	ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeAttributeListWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeClusterRevisionWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeDSTOffsetWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeDSTOffsetListMaxSizeWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeDefaultNTPWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeFeatureMapWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeGeneratedCommandListWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeGranularityWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeLocalTimeWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeNTPServerAvailableWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeSupportsDNSResolveWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeTimeSourceWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeTimeZoneWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeTimeZoneDatabaseWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeTimeZoneListMaxSizeWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeTrustedTimeSourceWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeUTCTimeWithParams(params IMTRReadParams) unsafe.Pointer
	SetDSTOffsetWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRTimeSynchronizationClusterSetDSTOffsetParams, expectedDataValueDictionaries []foundation.IDictionary, expectedValueIntervalMs foundation.INumber, completion unsafe.Pointer)
	SetDefaultNTPWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRTimeSynchronizationClusterSetDefaultNTPParams, expectedDataValueDictionaries []foundation.IDictionary, expectedValueIntervalMs foundation.INumber, completion unsafe.Pointer)
	SetTimeZoneWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRTimeSynchronizationClusterSetTimeZoneParams, expectedDataValueDictionaries []foundation.IDictionary, expectedValueIntervalMs foundation.INumber, completion unsafe.Pointer)
	SetTrustedTimeSourceWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRTimeSynchronizationClusterSetTrustedTimeSourceParams, expectedDataValueDictionaries []foundation.IDictionary, expectedValueIntervalMs foundation.INumber, completion unsafe.Pointer)
	SetUTCTimeWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRTimeSynchronizationClusterSetUTCTimeParams, expectedDataValueDictionaries []foundation.IDictionary, expectedValueIntervalMs foundation.INumber, completion unsafe.Pointer)
}

// Cluster Time Synchronization Accurate time is required for a number of reasons, including scheduling, display and validating security materials.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTimeSynchronization
type MTRClusterTimeSynchronization struct {
	MTRGenericCluster
}

// MTRClusterTimeSynchronizationFrom constructs a [MTRClusterTimeSynchronization] from an unsafe.Pointer.
//
// Cluster Time Synchronization Accurate time is required for a number of reasons, including scheduling, display and validating security materials.
func MTRClusterTimeSynchronizationFrom(ptr unsafe.Pointer) MTRClusterTimeSynchronization {
	return MTRClusterTimeSynchronization{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterTimeSynchronizationClass) Alloc() MTRClusterTimeSynchronization {
	rv := objc.Send[MTRClusterTimeSynchronization](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterTimeSynchronizationClass) New() MTRClusterTimeSynchronization {
	rv := objc.Send[MTRClusterTimeSynchronization](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterTimeSynchronization) Init() MTRClusterTimeSynchronization {
	rv := objc.Send[MTRClusterTimeSynchronization](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterTimeSynchronization) Autorelease() MTRClusterTimeSynchronization {
	rv := objc.Send[MTRClusterTimeSynchronization](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterTimeSynchronization creates a new MTRClusterTimeSynchronization instance.
func NewMTRClusterTimeSynchronization() MTRClusterTimeSynchronization {
	return getMTRClusterTimeSynchronizationClass().New()
}




// For all instance methods that take a completion (i.e. command invocations), the completion will be called on the provided queue.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTimeSynchronization/init(device:endpointID:queue:)
func NewMTRClusterTimeSynchronizationWithDeviceEndpointIDQueue(device IMTRDevice, endpointID foundation.INumber, queue unsafe.Pointer) MTRClusterTimeSynchronization {
	instance := getMTRClusterTimeSynchronizationClass().Alloc()
	rv := objc.Send[MTRClusterTimeSynchronization](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTimeSynchronization/readAttributeAcceptedCommandList(with:)
func (m_ MTRClusterTimeSynchronization) ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTimeSynchronization/readAttributeAttributeList(with:)
func (m_ MTRClusterTimeSynchronization) ReadAttributeAttributeListWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeAttributeListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTimeSynchronization/readAttributeClusterRevision(with:)
func (m_ MTRClusterTimeSynchronization) ReadAttributeClusterRevisionWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeClusterRevisionWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTimeSynchronization/readAttributeDSTOffset(with:)
func (m_ MTRClusterTimeSynchronization) ReadAttributeDSTOffsetWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeDSTOffsetWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTimeSynchronization/readAttributeDSTOffsetListMaxSize(with:)
func (m_ MTRClusterTimeSynchronization) ReadAttributeDSTOffsetListMaxSizeWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeDSTOffsetListMaxSizeWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTimeSynchronization/readAttributeDefaultNTP(with:)
func (m_ MTRClusterTimeSynchronization) ReadAttributeDefaultNTPWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeDefaultNTPWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTimeSynchronization/readAttributeFeatureMap(with:)
func (m_ MTRClusterTimeSynchronization) ReadAttributeFeatureMapWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeFeatureMapWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTimeSynchronization/readAttributeGeneratedCommandList(with:)
func (m_ MTRClusterTimeSynchronization) ReadAttributeGeneratedCommandListWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTimeSynchronization/readAttributeGranularity(with:)
func (m_ MTRClusterTimeSynchronization) ReadAttributeGranularityWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeGranularityWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTimeSynchronization/readAttributeLocalTime(with:)
func (m_ MTRClusterTimeSynchronization) ReadAttributeLocalTimeWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeLocalTimeWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTimeSynchronization/readAttributeNTPServerAvailable(with:)
func (m_ MTRClusterTimeSynchronization) ReadAttributeNTPServerAvailableWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeNTPServerAvailableWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTimeSynchronization/readAttributeSupportsDNSResolve(with:)
func (m_ MTRClusterTimeSynchronization) ReadAttributeSupportsDNSResolveWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeSupportsDNSResolveWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTimeSynchronization/readAttributeTimeSource(with:)
func (m_ MTRClusterTimeSynchronization) ReadAttributeTimeSourceWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeTimeSourceWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTimeSynchronization/readAttributeTimeZone(with:)
func (m_ MTRClusterTimeSynchronization) ReadAttributeTimeZoneWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeTimeZoneWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTimeSynchronization/readAttributeTimeZoneDatabase(with:)
func (m_ MTRClusterTimeSynchronization) ReadAttributeTimeZoneDatabaseWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeTimeZoneDatabaseWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTimeSynchronization/readAttributeTimeZoneListMaxSize(with:)
func (m_ MTRClusterTimeSynchronization) ReadAttributeTimeZoneListMaxSizeWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeTimeZoneListMaxSizeWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTimeSynchronization/readAttributeTrustedTimeSource(with:)
func (m_ MTRClusterTimeSynchronization) ReadAttributeTrustedTimeSourceWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeTrustedTimeSourceWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTimeSynchronization/readAttributeUTCTime(with:)
func (m_ MTRClusterTimeSynchronization) ReadAttributeUTCTimeWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeUTCTimeWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTimeSynchronization/setDSTOffsetWith(_:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterTimeSynchronization) SetDSTOffsetWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRTimeSynchronizationClusterSetDSTOffsetParams, expectedDataValueDictionaries []foundation.IDictionary, expectedValueIntervalMs foundation.INumber, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDSTOffsetWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTimeSynchronization/setDefaultNTPWith(_:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterTimeSynchronization) SetDefaultNTPWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRTimeSynchronizationClusterSetDefaultNTPParams, expectedDataValueDictionaries []foundation.IDictionary, expectedValueIntervalMs foundation.INumber, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDefaultNTPWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTimeSynchronization/setTimeZoneWith(_:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterTimeSynchronization) SetTimeZoneWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRTimeSynchronizationClusterSetTimeZoneParams, expectedDataValueDictionaries []foundation.IDictionary, expectedValueIntervalMs foundation.INumber, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimeZoneWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTimeSynchronization/setTrustedTimeSourceWith(_:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterTimeSynchronization) SetTrustedTimeSourceWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRTimeSynchronizationClusterSetTrustedTimeSourceParams, expectedDataValueDictionaries []foundation.IDictionary, expectedValueIntervalMs foundation.INumber, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTrustedTimeSourceWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTimeSynchronization/setUTCTimeWith(_:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterTimeSynchronization) SetUTCTimeWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRTimeSynchronizationClusterSetUTCTimeParams, expectedDataValueDictionaries []foundation.IDictionary, expectedValueIntervalMs foundation.INumber, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUTCTimeWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}


