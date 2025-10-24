// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKLookAroundSnapshot */


/* debug [class_header]: Header for MKLookAroundSnapshot */
// The class instance for the [MKLookAroundSnapshot] class.
var (
	MKLookAroundSnapshotClass     _MKLookAroundSnapshotClass
	MKLookAroundSnapshotClassOnce sync.Once
)

func getMKLookAroundSnapshotClass() _MKLookAroundSnapshotClass {
	MKLookAroundSnapshotClassOnce.Do(func() {
		MKLookAroundSnapshotClass = _MKLookAroundSnapshotClass{objc.GetClass("MKLookAroundSnapshot")}
	})
	return MKLookAroundSnapshotClass
}

type _MKLookAroundSnapshotClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKLookAroundSnapshot */
// An interface definition for the [MKLookAroundSnapshot] class.
type IMKLookAroundSnapshot interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKLookAroundSnapshot */
	// properties:
	Image() appkit.Image
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKLookAroundSnapshot */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKLookAroundSnapshot */
// Alloc allocates a new instance without initialization.
func (mc _MKLookAroundSnapshotClass) Alloc() MKLookAroundSnapshot {
	rv := objc.Send[MKLookAroundSnapshot](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKLookAroundSnapshotClass) New() MKLookAroundSnapshot {
	rv := objc.Send[MKLookAroundSnapshot](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKLookAroundSnapshot) Init() MKLookAroundSnapshot {
	rv := objc.Send[MKLookAroundSnapshot](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKLookAroundSnapshot) Autorelease() MKLookAroundSnapshot {
	rv := objc.Send[MKLookAroundSnapshot](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKLookAroundSnapshot creates a new MKLookAroundSnapshot instance.
func NewMKLookAroundSnapshot() MKLookAroundSnapshot {
	return getMKLookAroundSnapshotClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKLookAroundSnapshot */
// An object that contains a snapshot image.


// An object that contains a snapshot image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundSnapshotter/Snapshot
type MKLookAroundSnapshot struct {
	objectivec.Object
}

// MKLookAroundSnapshotFrom constructs a [MKLookAroundSnapshot] from an unsafe.Pointer.
//
// An object that contains a snapshot image.
func MKLookAroundSnapshotFrom(ptr unsafe.Pointer) MKLookAroundSnapshot {
	return MKLookAroundSnapshot{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKLookAroundSnapshot *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKLookAroundSnapshot */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKLookAroundSnapshot */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKLookAroundSnapshot */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKLookAroundSnapshot */

// The image returned by the snapshot request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundSnapshotter/Snapshot/image
func (m_ MKLookAroundSnapshot) Image() appkit.Image {
	rv := objc.Send[appkit.Image](m_.ID, objc.Sel("image"))
	return rv
}/* debug [instance_properties/getter]: image */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKLookAroundSnapshot */



