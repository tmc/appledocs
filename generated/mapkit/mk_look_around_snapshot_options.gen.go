// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKLookAroundSnapshotOptions */


/* debug [class_header]: Header for MKLookAroundSnapshotOptions */
// The class instance for the [MKLookAroundSnapshotOptions] class.
var (
	MKLookAroundSnapshotOptionsClass     _MKLookAroundSnapshotOptionsClass
	MKLookAroundSnapshotOptionsClassOnce sync.Once
)

func getMKLookAroundSnapshotOptionsClass() _MKLookAroundSnapshotOptionsClass {
	MKLookAroundSnapshotOptionsClassOnce.Do(func() {
		MKLookAroundSnapshotOptionsClass = _MKLookAroundSnapshotOptionsClass{objc.GetClass("MKLookAroundSnapshotOptions")}
	})
	return MKLookAroundSnapshotOptionsClass
}

type _MKLookAroundSnapshotOptionsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKLookAroundSnapshotOptions */
// An interface definition for the [MKLookAroundSnapshotOptions] class.
type IMKLookAroundSnapshotOptions interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKLookAroundSnapshotOptions */
	// properties:
	PointOfInterestFilter() IMKPointOfInterestFilter
	SetPointOfInterestFilter(value IMKPointOfInterestFilter)
	Size() corefoundation.CGSize
	SetSize(value corefoundation.CGSize)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKLookAroundSnapshotOptions */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKLookAroundSnapshotOptions */
// Alloc allocates a new instance without initialization.
func (mc _MKLookAroundSnapshotOptionsClass) Alloc() MKLookAroundSnapshotOptions {
	rv := objc.Send[MKLookAroundSnapshotOptions](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKLookAroundSnapshotOptionsClass) New() MKLookAroundSnapshotOptions {
	rv := objc.Send[MKLookAroundSnapshotOptions](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKLookAroundSnapshotOptions) Init() MKLookAroundSnapshotOptions {
	rv := objc.Send[MKLookAroundSnapshotOptions](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKLookAroundSnapshotOptions) Autorelease() MKLookAroundSnapshotOptions {
	rv := objc.Send[MKLookAroundSnapshotOptions](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKLookAroundSnapshotOptions creates a new MKLookAroundSnapshotOptions instance.
func NewMKLookAroundSnapshotOptions() MKLookAroundSnapshotOptions {
	return getMKLookAroundSnapshotOptionsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKLookAroundSnapshotOptions */
// Values you use to customize LookAround snapshots.


// Values you use to customize LookAround snapshots.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundSnapshotter/Options
type MKLookAroundSnapshotOptions struct {
	objectivec.Object
}

// MKLookAroundSnapshotOptionsFrom constructs a [MKLookAroundSnapshotOptions] from an unsafe.Pointer.
//
// Values you use to customize LookAround snapshots.
func MKLookAroundSnapshotOptionsFrom(ptr unsafe.Pointer) MKLookAroundSnapshotOptions {
	return MKLookAroundSnapshotOptions{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKLookAroundSnapshotOptions *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKLookAroundSnapshotOptions */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKLookAroundSnapshotOptions */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKLookAroundSnapshotOptions */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKLookAroundSnapshotOptions */

// A filter to use to customize what map features are visible in the snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundSnapshotter/Options/pointOfInterestFilter
func (m_ MKLookAroundSnapshotOptions) PointOfInterestFilter() IMKPointOfInterestFilter {
	rv := objc.Send[MKPointOfInterestFilter](m_.ID, objc.Sel("pointOfInterestFilter"))
	return rv
}/* debug [instance_properties/getter]: pointOfInterestFilter */


// A filter to use to customize what map features are visible in the snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundSnapshotter/Options/pointOfInterestFilter
func (m_ MKLookAroundSnapshotOptions) SetPointOfInterestFilter(value IMKPointOfInterestFilter) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPointOfInterestFilter:"), value)
}/* debug [instance_properties/setter]: pointOfInterestFilter */


// The requested size of the snapshot image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundSnapshotter/Options/size
func (m_ MKLookAroundSnapshotOptions) Size() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](m_.ID, objc.Sel("size"))
	return rv
}/* debug [instance_properties/getter]: size */


// The requested size of the snapshot image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundSnapshotter/Options/size
func (m_ MKLookAroundSnapshotOptions) SetSize(value corefoundation.CGSize) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSize:"), value)
}/* debug [instance_properties/setter]: size */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKLookAroundSnapshotOptions */


