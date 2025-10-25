// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVFragmentedAssetTrack */


/* debug [class_header]: Header for AVFragmentedAssetTrack */
// The class instance for the [FragmentedAssetTrack] class.
var (
	FragmentedAssetTrackClass     _FragmentedAssetTrackClass
	FragmentedAssetTrackClassOnce sync.Once
)

func getFragmentedAssetTrackClass() _FragmentedAssetTrackClass {
	FragmentedAssetTrackClassOnce.Do(func() {
		FragmentedAssetTrackClass = _FragmentedAssetTrackClass{objc.GetClass("AVFragmentedAssetTrack")}
	})
	return FragmentedAssetTrackClass
}

type _FragmentedAssetTrackClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FragmentedAssetTrack */
// An interface definition for the [FragmentedAssetTrack] class.
type IFragmentedAssetTrack interface {
	IAssetTrack
	
/* debug [class_interface_properties]: Properties for FragmentedAssetTrack */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FragmentedAssetTrack */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FragmentedAssetTrack */
// Alloc allocates a new instance without initialization.
func (fc _FragmentedAssetTrackClass) Alloc() FragmentedAssetTrack {
	rv := objc.Send[FragmentedAssetTrack](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FragmentedAssetTrackClass) New() FragmentedAssetTrack {
	rv := objc.Send[FragmentedAssetTrack](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FragmentedAssetTrack) Init() FragmentedAssetTrack {
	rv := objc.Send[FragmentedAssetTrack](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FragmentedAssetTrack) Autorelease() FragmentedAssetTrack {
	rv := objc.Send[FragmentedAssetTrack](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFragmentedAssetTrack creates a new FragmentedAssetTrack instance.
func NewFragmentedAssetTrack() FragmentedAssetTrack {
	return getFragmentedAssetTrackClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FragmentedAssetTrack */
// An object that provides the track-level interface to inspect a fragmented asset’s media tracks.
//
// This class subclasses . It has no methods or properties of its own.


// An object that provides the track-level interface to inspect a fragmented asset’s media tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVFragmentedAssetTrack
type FragmentedAssetTrack struct {
	AssetTrack
}

// FragmentedAssetTrackFrom constructs a [FragmentedAssetTrack] from an unsafe.Pointer.
//
// An object that provides the track-level interface to inspect a fragmented asset’s media tracks.
func FragmentedAssetTrackFrom(ptr unsafe.Pointer) FragmentedAssetTrack {
	return FragmentedAssetTrack{
		AssetTrack: AssetTrackFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FragmentedAssetTrack *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FragmentedAssetTrack */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FragmentedAssetTrack */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FragmentedAssetTrack */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FragmentedAssetTrack */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVFragmentedAssetTrack */



