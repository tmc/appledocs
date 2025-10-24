//go:build darwin && ios

// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for MKLookAroundSnapshotOptions


// iOS-only properties

// A collection of traits that describes orientation and other characteristics of the snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundSnapshotter/Options/traitCollection
func (m_ MKLookAroundSnapshotOptions) TraitCollection() TraitCollection /* not a class type */ {
	rv := objc.Send[TraitCollection](m_.ID, objc.Sel("traitCollection"))
	return rv
}
func (m_ MKLookAroundSnapshotOptions) SetTraitCollection(value TraitCollection /* not a class type */) {
	m_.ID.Send(objc.RegisterName("setTraitCollection:"), value)
}





