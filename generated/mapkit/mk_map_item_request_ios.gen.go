//go:build darwin && ios

// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for MKMapItemRequest


// iOS-only properties

// The feature annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItemRequest/featureAnnotation
func (m_ MKMapItemRequest) FeatureAnnotation() IMKMapFeatureAnnotation {
	rv := objc.Send[MKMapFeatureAnnotation](m_.ID, objc.Sel("featureAnnotation"))
	return rv
}

// The feature annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItemRequest/mapFeatureAnnotation
func (m_ MKMapItemRequest) MapFeatureAnnotation() IMKMapFeatureAnnotation {
	rv := objc.Send[MKMapFeatureAnnotation](m_.ID, objc.Sel("mapFeatureAnnotation"))
	return rv
}




