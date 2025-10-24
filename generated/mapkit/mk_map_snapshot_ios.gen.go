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

// iOS-only methods for MKMapSnapshot


// iOS-only properties

// Traits to use when creating the snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Snapshot/traitCollection
func (m_ MKMapSnapshot) TraitCollection() TraitCollection /* not a class type */ {
	rv := objc.Send[TraitCollection](m_.ID, objc.Sel("traitCollection"))
	return rv
}





