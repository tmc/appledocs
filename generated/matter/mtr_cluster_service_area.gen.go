// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRClusterServiceArea] class.
var (
	MTRClusterServiceAreaClass     _MTRClusterServiceAreaClass
	MTRClusterServiceAreaClassOnce sync.Once
)

func getMTRClusterServiceAreaClass() _MTRClusterServiceAreaClass {
	MTRClusterServiceAreaClassOnce.Do(func() {
		MTRClusterServiceAreaClass = _MTRClusterServiceAreaClass{objc.GetClass("MTRClusterServiceArea")}
	})
	return MTRClusterServiceAreaClass
}

type _MTRClusterServiceAreaClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterServiceArea] class.
type IMTRClusterServiceArea interface {
	IMTRGenericCluster
	// properties:
	// methods:
	ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeAttributeListWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeClusterRevisionWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeCurrentAreaWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeEstimatedEndTimeWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeFeatureMapWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeGeneratedCommandListWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeProgressWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeSelectedAreasWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeSupportedAreasWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeSupportedMapsWithParams(params IMTRReadParams) foundation.IDictionary
	SelectAreasWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRServiceAreaClusterSelectAreasParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	SkipAreaWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRServiceAreaClusterSkipAreaParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
}

// Cluster Service Area The Service Area cluster provides an interface for controlling the areas where a device should operate, and for querying the current area being serviced.


// Cluster Service Area The Service Area cluster provides an interface for controlling the areas where a device should operate, and for querying the current area being serviced.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterServiceArea
type MTRClusterServiceArea struct {
	MTRGenericCluster
}

// MTRClusterServiceAreaFrom constructs a [MTRClusterServiceArea] from an unsafe.Pointer.
//
// Cluster Service Area The Service Area cluster provides an interface for controlling the areas where a device should operate, and for querying the current area being serviced.
func MTRClusterServiceAreaFrom(ptr unsafe.Pointer) MTRClusterServiceArea {
	return MTRClusterServiceArea{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterServiceAreaClass) Alloc() MTRClusterServiceArea {
	rv := objc.Send[MTRClusterServiceArea](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterServiceAreaClass) New() MTRClusterServiceArea {
	rv := objc.Send[MTRClusterServiceArea](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterServiceArea) Init() MTRClusterServiceArea {
	rv := objc.Send[MTRClusterServiceArea](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterServiceArea) Autorelease() MTRClusterServiceArea {
	rv := objc.Send[MTRClusterServiceArea](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterServiceArea creates a new MTRClusterServiceArea instance.
func NewMTRClusterServiceArea() MTRClusterServiceArea {
	return getMTRClusterServiceAreaClass().New()
}



// For all instance methods that take a completion (i.e. command invocations), the completion will be called on the provided queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterServiceArea/init(device:endpointID:queue:)
func NewMTRClusterServiceAreaWithDeviceEndpointIDQueue(device IMTRDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRClusterServiceArea {
	instance := getMTRClusterServiceAreaClass().Alloc()
	rv := objc.Send[MTRClusterServiceArea](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterServiceArea/readAttributeAcceptedCommandList(with:)
func (m_ MTRClusterServiceArea) ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterServiceArea/readAttributeAttributeList(with:)
func (m_ MTRClusterServiceArea) ReadAttributeAttributeListWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeAttributeListWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterServiceArea/readAttributeClusterRevision(with:)
func (m_ MTRClusterServiceArea) ReadAttributeClusterRevisionWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeClusterRevisionWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterServiceArea/readAttributeCurrentArea(with:)
func (m_ MTRClusterServiceArea) ReadAttributeCurrentAreaWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeCurrentAreaWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterServiceArea/readAttributeEstimatedEndTime(with:)
func (m_ MTRClusterServiceArea) ReadAttributeEstimatedEndTimeWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeEstimatedEndTimeWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterServiceArea/readAttributeFeatureMap(with:)
func (m_ MTRClusterServiceArea) ReadAttributeFeatureMapWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeFeatureMapWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterServiceArea/readAttributeGeneratedCommandList(with:)
func (m_ MTRClusterServiceArea) ReadAttributeGeneratedCommandListWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterServiceArea/readAttributeProgress(with:)
func (m_ MTRClusterServiceArea) ReadAttributeProgressWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeProgressWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterServiceArea/readAttributeSelectedAreas(with:)
func (m_ MTRClusterServiceArea) ReadAttributeSelectedAreasWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeSelectedAreasWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterServiceArea/readAttributeSupportedAreas(with:)
func (m_ MTRClusterServiceArea) ReadAttributeSupportedAreasWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeSupportedAreasWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterServiceArea/readAttributeSupportedMaps(with:)
func (m_ MTRClusterServiceArea) ReadAttributeSupportedMapsWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeSupportedMapsWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterServiceArea/selectAreas(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterServiceArea) SelectAreasWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRServiceAreaClusterSelectAreasParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("selectAreasWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterServiceArea/skip(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterServiceArea) SkipAreaWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRServiceAreaClusterSkipAreaParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("skipAreaWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}


