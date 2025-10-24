//go:build darwin && ios

// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for MKMapFeatureAnnotation


// iOS-only properties

// The type of map feature this annotation represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapFeatureAnnotation/featureType-swift.property
func (m_ MKMapFeatureAnnotation) FeatureType() MKMapFeatureType {
	rv := objc.Send[MKMapFeatureType](m_.ID, objc.Sel("featureType"))
	return rv
}

// The icon style of a feature annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapFeatureAnnotation/iconStyle
func (m_ MKMapFeatureAnnotation) IconStyle() IMKIconStyle {
	rv := objc.Send[MKIconStyle](m_.ID, objc.Sel("iconStyle"))
	return rv
}

// The feature annotation’s point of interest category.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapFeatureAnnotation/pointOfInterestCategory
func (m_ MKMapFeatureAnnotation) PointOfInterestCategory() MKPointOfInterestCategory /* typedef */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("pointOfInterestCategory"))
	return rv
}





