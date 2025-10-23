// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MKMapFeatureAnnotation] class.
var (
	MKMapFeatureAnnotationClass     _MKMapFeatureAnnotationClass
	MKMapFeatureAnnotationClassOnce sync.Once
)

func getMKMapFeatureAnnotationClass() _MKMapFeatureAnnotationClass {
	MKMapFeatureAnnotationClassOnce.Do(func() {
		MKMapFeatureAnnotationClass = _MKMapFeatureAnnotationClass{objc.GetClass("MKMapFeatureAnnotation")}
	})
	return MKMapFeatureAnnotationClass
}

type _MKMapFeatureAnnotationClass struct {
	class objc.Class
}

// An interface definition for the [MKMapFeatureAnnotation] class.
type IMKMapFeatureAnnotation interface {
	objectivec.IObject
	// properties:
	FeatureType() unsafe.Pointer
	SetFeatureType(value unsafe.Pointer)
	IconStyle() IMKIconStyle
	SetIconStyle(value IMKIconStyle)
	PointOfInterestCategory() MKPointOfInterestCategory /* typedef */
	SetPointOfInterestCategory(value MKPointOfInterestCategory /* typedef */)
	// methods:
}

// A class that describes an annotation element on the map’s display such as a point of interest, territorial boundary, or physical feature.


// A class that describes an annotation element on the map’s display such as a point of interest, territorial boundary, or physical feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapFeatureAnnotation
type MKMapFeatureAnnotation struct {
	objectivec.Object
}

// MKMapFeatureAnnotationFrom constructs a [MKMapFeatureAnnotation] from an unsafe.Pointer.
//
// A class that describes an annotation element on the map’s display such as a point of interest, territorial boundary, or physical feature.
func MKMapFeatureAnnotationFrom(ptr unsafe.Pointer) MKMapFeatureAnnotation {
	return MKMapFeatureAnnotation{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MKMapFeatureAnnotationClass) Alloc() MKMapFeatureAnnotation {
	rv := objc.Send[MKMapFeatureAnnotation](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKMapFeatureAnnotationClass) New() MKMapFeatureAnnotation {
	rv := objc.Send[MKMapFeatureAnnotation](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKMapFeatureAnnotation) Init() MKMapFeatureAnnotation {
	rv := objc.Send[MKMapFeatureAnnotation](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKMapFeatureAnnotation) Autorelease() MKMapFeatureAnnotation {
	rv := objc.Send[MKMapFeatureAnnotation](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKMapFeatureAnnotation creates a new MKMapFeatureAnnotation instance.
func NewMKMapFeatureAnnotation() MKMapFeatureAnnotation {
	return getMKMapFeatureAnnotationClass().New()
}



// The type of map feature this annotation represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapfeatureannotation/featuretype-swift.property
func (m_ MKMapFeatureAnnotation) FeatureType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("featureType"))
	return rv
}


// The type of map feature this annotation represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapfeatureannotation/featuretype-swift.property
func (m_ MKMapFeatureAnnotation) SetFeatureType(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFeatureType:"), value)
}


// The icon style of a feature annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapfeatureannotation/iconstyle
func (m_ MKMapFeatureAnnotation) IconStyle() IMKIconStyle {
	rv := objc.Send[MKIconStyle](m_.ID, objc.Sel("iconStyle"))
	return rv
}


// The icon style of a feature annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapfeatureannotation/iconstyle
func (m_ MKMapFeatureAnnotation) SetIconStyle(value IMKIconStyle) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIconStyle:"), value)
}


// The feature annotation’s point of interest category.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapfeatureannotation/pointofinterestcategory
func (m_ MKMapFeatureAnnotation) PointOfInterestCategory() MKPointOfInterestCategory /* typedef */ {
	rv := objc.Send[MKPointOfInterestCategory](m_.ID, objc.Sel("pointOfInterestCategory"))
	return rv
}


// The feature annotation’s point of interest category.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapfeatureannotation/pointofinterestcategory
func (m_ MKMapFeatureAnnotation) SetPointOfInterestCategory(value MKPointOfInterestCategory /* typedef */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPointOfInterestCategory:"), value)
}




