//go:build darwin && ios

// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for MKMapSnapshotOptions


// iOS-only properties

// The scale factor to use when creating the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/scale
func (m_ MKMapSnapshotOptions) Scale() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("scale"))
	return rv
}
func (m_ MKMapSnapshotOptions) SetScale(value float64) {
	m_.ID.Send(objc.RegisterName("setScale:"), value)
}

// Traits that determine the appearance of the map snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/traitCollection
func (m_ MKMapSnapshotOptions) TraitCollection() TraitCollection /* not a class type */ {
	rv := objc.Send[TraitCollection](m_.ID, objc.Sel("traitCollection"))
	return rv
}
func (m_ MKMapSnapshotOptions) SetTraitCollection(value TraitCollection /* not a class type */) {
	m_.ID.Send(objc.RegisterName("setTraitCollection:"), value)
}





