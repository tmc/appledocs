// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAssetTrackGroup */


/* debug [class_header]: Header for AVAssetTrackGroup */
// The class instance for the [AssetTrackGroup] class.
var (
	AssetTrackGroupClass     _AssetTrackGroupClass
	AssetTrackGroupClassOnce sync.Once
)

func getAssetTrackGroupClass() _AssetTrackGroupClass {
	AssetTrackGroupClassOnce.Do(func() {
		AssetTrackGroupClass = _AssetTrackGroupClass{objc.GetClass("AVAssetTrackGroup")}
	})
	return AssetTrackGroupClass
}

type _AssetTrackGroupClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AssetTrackGroup */
// An interface definition for the [AssetTrackGroup] class.
type IAssetTrackGroup interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AssetTrackGroup */
	// properties:
	TrackIDs() []foundation.Number
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AssetTrackGroup */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AssetTrackGroup */
// Alloc allocates a new instance without initialization.
func (ac _AssetTrackGroupClass) Alloc() AssetTrackGroup {
	rv := objc.Send[AssetTrackGroup](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AssetTrackGroupClass) New() AssetTrackGroup {
	rv := objc.Send[AssetTrackGroup](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetTrackGroup) Init() AssetTrackGroup {
	rv := objc.Send[AssetTrackGroup](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetTrackGroup) Autorelease() AssetTrackGroup {
	rv := objc.Send[AssetTrackGroup](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetTrackGroup creates a new AssetTrackGroup instance.
func NewAssetTrackGroup() AssetTrackGroup {
	return getAssetTrackGroupClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AssetTrackGroup */
// A group of related tracks in an asset.
//
// A track group describes a group of related alternative tracks, only one of which should play at a time. Groups of alternative tracks typically contain variations of the same content, like subtitles in multiple translations. You can inspect an asset’s track groups by loading the value of its property.


// A group of related tracks in an asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrackGroup
type AssetTrackGroup struct {
	objectivec.Object
}

// AssetTrackGroupFrom constructs a [AssetTrackGroup] from an unsafe.Pointer.
//
// A group of related tracks in an asset.
func AssetTrackGroupFrom(ptr unsafe.Pointer) AssetTrackGroup {
	return AssetTrackGroup{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AssetTrackGroup *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AssetTrackGroup */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AssetTrackGroup */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AssetTrackGroup */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AssetTrackGroup */

// The IDs of the tracks in the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrackGroup/trackIDs
func (a_ AssetTrackGroup) TrackIDs() []foundation.Number {
	rv := objc.Send[[]foundation.Number](a_.ID, objc.Sel("trackIDs"))
	return rv
}/* debug [instance_properties/getter]: trackIDs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAssetTrackGroup */



