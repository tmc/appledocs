// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVMediaSelection */


/* debug [class_header]: Header for AVMediaSelection */
// The class instance for the [MediaSelection] class.
var (
	MediaSelectionClass     _MediaSelectionClass
	MediaSelectionClassOnce sync.Once
)

func getMediaSelectionClass() _MediaSelectionClass {
	MediaSelectionClassOnce.Do(func() {
		MediaSelectionClass = _MediaSelectionClass{objc.GetClass("AVMediaSelection")}
	})
	return MediaSelectionClass
}

type _MediaSelectionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MediaSelection */
// An interface definition for the [MediaSelection] class.
type IMediaSelection interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MediaSelection */
	// properties:
	Asset() IAVAsset
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MediaSelection */
	// methods:
	MediaSelectionCriteriaCanBeAppliedAutomaticallyToMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) bool
	SelectedMediaOptionInMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) IMediaSelectionOption
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MediaSelection */
// Alloc allocates a new instance without initialization.
func (mc _MediaSelectionClass) Alloc() MediaSelection {
	rv := objc.Send[MediaSelection](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MediaSelectionClass) New() MediaSelection {
	rv := objc.Send[MediaSelection](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MediaSelection) Init() MediaSelection {
	rv := objc.Send[MediaSelection](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MediaSelection) Autorelease() MediaSelection {
	rv := objc.Send[MediaSelection](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMediaSelection creates a new MediaSelection instance.
func NewMediaSelection() MediaSelection {
	return getMediaSelectionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MediaSelection */
// An object that represents a complete rendition of media selection options on an asset.


// An object that represents a complete rendition of media selection options on an asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaSelection
type MediaSelection struct {
	objectivec.Object
}

// MediaSelectionFrom constructs a [MediaSelection] from an unsafe.Pointer.
//
// An object that represents a complete rendition of media selection options on an asset.
func MediaSelectionFrom(ptr unsafe.Pointer) MediaSelection {
	return MediaSelection{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MediaSelection *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MediaSelection */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MediaSelection */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MediaSelection */

// Indicates whether the specified media selection group is subject to automatic media selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaSelection/mediaSelectionCriteriaCanBeAppliedAutomatically(to:)
func (m_ MediaSelection) MediaSelectionCriteriaCanBeAppliedAutomaticallyToMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("mediaSelectionCriteriaCanBeAppliedAutomaticallyToMediaSelectionGroup:"), mediaSelectionGroup)
	return rv
}/* debug [instance_methods/method]: MediaSelectionCriteriaCanBeAppliedAutomaticallyToMediaSelectionGroup */


// Returns the media selection option that’s currently selected in the specified group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaSelection/selectedMediaOption(in:)
func (m_ MediaSelection) SelectedMediaOptionInMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) IMediaSelectionOption {
	rv := objc.Send[MediaSelectionOption](m_.ID, objc.Sel("selectedMediaOptionInMediaSelectionGroup:"), mediaSelectionGroup)
	return rv
}/* debug [instance_methods/method]: SelectedMediaOptionInMediaSelectionGroup */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MediaSelection */

// The asset associated with the media selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaSelection/asset
func (m_ MediaSelection) Asset() IAVAsset {
	rv := objc.Send[Asset](m_.ID, objc.Sel("asset"))
	return rv
}/* debug [instance_properties/getter]: asset */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMediaSelection */



