// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MKMapSnapshot] class.
var (
	MKMapSnapshotClass     _MKMapSnapshotClass
	MKMapSnapshotClassOnce sync.Once
)

func getMKMapSnapshotClass() _MKMapSnapshotClass {
	MKMapSnapshotClassOnce.Do(func() {
		MKMapSnapshotClass = _MKMapSnapshotClass{objc.GetClass("MKMapSnapshot")}
	})
	return MKMapSnapshotClass
}

type _MKMapSnapshotClass struct {
	class objc.Class
}

// An interface definition for the [MKMapSnapshot] class.
type IMKMapSnapshot interface {
	objectivec.IObject
	PointForCoordinate(coordinate unsafe.Pointer) coregraphics.CGPoint
}

// An image that a snapshotter object generates.
//
// You don’t create instances of this class directly. Instead, you use an object to capture the map contents asynchronously. An object contains the image that the snapshotter generates from the map contents. Snapshot images don’t include any custom overlays or annotations that your app adds to the map view. If you want your annotations and overlays to appear on the final image, you need to draw them yourself. To position those items correctly on the image, use the method of this class to translate the overlay or annotation coordinate value to an appropriate location inside the image’s coordinate space.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Snapshot
type MKMapSnapshot struct {
	objectivec.Object
}

// MKMapSnapshotFrom constructs a [MKMapSnapshot] from an unsafe.Pointer.
//
// An image that a snapshotter object generates.
func MKMapSnapshotFrom(ptr unsafe.Pointer) MKMapSnapshot {
	return MKMapSnapshot{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MKMapSnapshotClass) Alloc() MKMapSnapshot {
	rv := objc.Send[MKMapSnapshot](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKMapSnapshotClass) New() MKMapSnapshot {
	rv := objc.Send[MKMapSnapshot](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKMapSnapshot) Init() MKMapSnapshot {
	rv := objc.Send[MKMapSnapshot](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKMapSnapshot) Autorelease() MKMapSnapshot {
	rv := objc.Send[MKMapSnapshot](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKMapSnapshot creates a new MKMapSnapshot instance.
func NewMKMapSnapshot() MKMapSnapshot {
	return getMKMapSnapshotClass().New()
}


// Converts the specified map coordinate to a point in the coordinate space of the image.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Snapshot/point(for:)
func (m_ MKMapSnapshot) PointForCoordinate(coordinate unsafe.Pointer) coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](m_.ID, objc.Sel("pointForCoordinate:"), coordinate)
	return rv
}

// The visual style that MapKit uses when rendering the snapshot.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Snapshot/appearance
func (m_ MKMapSnapshot) Appearance() appkit.Appearance {
	rv := objc.Send[appkit.Appearance](m_.ID, objc.Sel("appearance"))
	return rv
}

// The image of the map’s content.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapsnapshotter/snapshot/image
func (m_ MKMapSnapshot) Image() appkit.Image {
	rv := objc.Send[appkit.Image](m_.ID, objc.Sel("image"))
	return rv
}


// SetImage sets the value of the image property.
// The image of the map’s content.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapsnapshotter/snapshot/image
func (m_ MKMapSnapshot) SetImage(value appkit.IImage) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setImage:"), value)
}

// Traits to use when creating the snapshot.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapsnapshotter/snapshot/traitcollection
func (m_ MKMapSnapshot) TraitCollection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("traitCollection"))
	return rv
}


// SetTraitCollection sets the value of the traitCollection property.
// Traits to use when creating the snapshot.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapsnapshotter/snapshot/traitcollection
func (m_ MKMapSnapshot) SetTraitCollection(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTraitCollection:"), value)
}



