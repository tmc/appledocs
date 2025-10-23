// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MKMapItemRequest] class.
var (
	MKMapItemRequestClass     _MKMapItemRequestClass
	MKMapItemRequestClassOnce sync.Once
)

func getMKMapItemRequestClass() _MKMapItemRequestClass {
	MKMapItemRequestClassOnce.Do(func() {
		MKMapItemRequestClass = _MKMapItemRequestClass{objc.GetClass("MKMapItemRequest")}
	})
	return MKMapItemRequestClass
}

type _MKMapItemRequestClass struct {
	class objc.Class
}

// An interface definition for the [MKMapItemRequest] class.
type IMKMapItemRequest interface {
	objectivec.IObject
	// properties:
	Feature() unsafe.Pointer
	SetFeature(value unsafe.Pointer)
	FeatureAnnotation() IMKMapFeatureAnnotation
	SetFeatureAnnotation(value IMKMapFeatureAnnotation)
	IsCancelled() bool
	SetIsCancelled(value bool)
	IsLoading() bool
	SetIsLoading(value bool)
	MapFeature() unsafe.Pointer
	SetMapFeature(value unsafe.Pointer)
	MapFeatureAnnotation() IMKMapFeatureAnnotation
	SetMapFeatureAnnotation(value IMKMapFeatureAnnotation)
	MapItemIdentifier() MKMapItemIdentifier
	SetMapItemIdentifier(value MKMapItemIdentifier)
	PlaceDescriptor() unsafe.Pointer
	SetPlaceDescriptor(value unsafe.Pointer)
	// methods:
}

// A utility class you use to request additional information about a map feature.


// A utility class you use to request additional information about a map feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItemRequest
type MKMapItemRequest struct {
	objectivec.Object
}

// MKMapItemRequestFrom constructs a [MKMapItemRequest] from an unsafe.Pointer.
//
// A utility class you use to request additional information about a map feature.
func MKMapItemRequestFrom(ptr unsafe.Pointer) MKMapItemRequest {
	return MKMapItemRequest{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MKMapItemRequestClass) Alloc() MKMapItemRequest {
	rv := objc.Send[MKMapItemRequest](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKMapItemRequestClass) New() MKMapItemRequest {
	rv := objc.Send[MKMapItemRequest](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKMapItemRequest) Init() MKMapItemRequest {
	rv := objc.Send[MKMapItemRequest](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKMapItemRequest) Autorelease() MKMapItemRequest {
	rv := objc.Send[MKMapItemRequest](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKMapItemRequest creates a new MKMapItemRequest instance.
func NewMKMapItemRequest() MKMapItemRequest {
	return getMKMapItemRequestClass().New()
}



// The map feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitemrequest/feature
func (m_ MKMapItemRequest) Feature() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("feature"))
	return rv
}


// The map feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitemrequest/feature
func (m_ MKMapItemRequest) SetFeature(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFeature:"), value)
}


// The feature annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitemrequest/featureannotation
func (m_ MKMapItemRequest) FeatureAnnotation() IMKMapFeatureAnnotation {
	rv := objc.Send[MKMapFeatureAnnotation](m_.ID, objc.Sel("featureAnnotation"))
	return rv
}


// The feature annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitemrequest/featureannotation
func (m_ MKMapItemRequest) SetFeatureAnnotation(value IMKMapFeatureAnnotation) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFeatureAnnotation:"), value)
}


// A Boolean value that indicates if the cancellation of the request was successful.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitemrequest/iscancelled
func (m_ MKMapItemRequest) IsCancelled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isCancelled"))
	return rv
}


// A Boolean value that indicates if the cancellation of the request was successful.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitemrequest/iscancelled
func (m_ MKMapItemRequest) SetIsCancelled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsCancelled:"), value)
}


// A Boolean value that indicates if the request is loading.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitemrequest/isloading
func (m_ MKMapItemRequest) IsLoading() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isLoading"))
	return rv
}


// A Boolean value that indicates if the request is loading.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitemrequest/isloading
func (m_ MKMapItemRequest) SetIsLoading(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsLoading:"), value)
}


// The map feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitemrequest/mapfeature
func (m_ MKMapItemRequest) MapFeature() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("mapFeature"))
	return rv
}


// The map feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitemrequest/mapfeature
func (m_ MKMapItemRequest) SetMapFeature(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMapFeature:"), value)
}


// The feature annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitemrequest/mapfeatureannotation
func (m_ MKMapItemRequest) MapFeatureAnnotation() IMKMapFeatureAnnotation {
	rv := objc.Send[MKMapFeatureAnnotation](m_.ID, objc.Sel("mapFeatureAnnotation"))
	return rv
}


// The feature annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitemrequest/mapfeatureannotation
func (m_ MKMapItemRequest) SetMapFeatureAnnotation(value IMKMapFeatureAnnotation) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMapFeatureAnnotation:"), value)
}


// The map item identifer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitemrequest/mapitemidentifier
func (m_ MKMapItemRequest) MapItemIdentifier() MKMapItemIdentifier {
	rv := objc.Send[MKMapItemIdentifier](m_.ID, objc.Sel("mapItemIdentifier"))
	return rv
}


// The map item identifer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitemrequest/mapitemidentifier
func (m_ MKMapItemRequest) SetMapItemIdentifier(value MKMapItemIdentifier) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMapItemIdentifier:"), value)
}


// The place descriptor that contains information that’s helpful in uniquely identifying this place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitemrequest/placedescriptor
func (m_ MKMapItemRequest) PlaceDescriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("placeDescriptor"))
	return rv
}


// The place descriptor that contains information that’s helpful in uniquely identifying this place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitemrequest/placedescriptor
func (m_ MKMapItemRequest) SetPlaceDescriptor(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPlaceDescriptor:"), value)
}



