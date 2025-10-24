// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVMutableMediaSelection */


/* debug [class_header]: Header for AVMutableMediaSelection */
// The class instance for the [MutableMediaSelection] class.
var (
	MutableMediaSelectionClass     _MutableMediaSelectionClass
	MutableMediaSelectionClassOnce sync.Once
)

func getMutableMediaSelectionClass() _MutableMediaSelectionClass {
	MutableMediaSelectionClassOnce.Do(func() {
		MutableMediaSelectionClass = _MutableMediaSelectionClass{objc.GetClass("AVMutableMediaSelection")}
	})
	return MutableMediaSelectionClass
}

type _MutableMediaSelectionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MutableMediaSelection */
// An interface definition for the [MutableMediaSelection] class.
type IMutableMediaSelection interface {
	IMediaSelection
	
/* debug [class_interface_properties]: Properties for MutableMediaSelection */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MutableMediaSelection */
	// methods:
	SelectMediaOptionInMediaSelectionGroup(mediaSelectionOption IAVMediaSelectionOption, mediaSelectionGroup IAVMediaSelectionGroup)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MutableMediaSelection */
// Alloc allocates a new instance without initialization.
func (mc _MutableMediaSelectionClass) Alloc() MutableMediaSelection {
	rv := objc.Send[MutableMediaSelection](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MutableMediaSelectionClass) New() MutableMediaSelection {
	rv := objc.Send[MutableMediaSelection](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutableMediaSelection) Init() MutableMediaSelection {
	rv := objc.Send[MutableMediaSelection](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutableMediaSelection) Autorelease() MutableMediaSelection {
	rv := objc.Send[MutableMediaSelection](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutableMediaSelection creates a new MutableMediaSelection instance.
func NewMutableMediaSelection() MutableMediaSelection {
	return getMutableMediaSelectionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MutableMediaSelection */
// A mutable object that represents a complete rendition of media selection options on an asset.


// A mutable object that represents a complete rendition of media selection options on an asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMediaSelection
type MutableMediaSelection struct {
	MediaSelection
}

// MutableMediaSelectionFrom constructs a [MutableMediaSelection] from an unsafe.Pointer.
//
// A mutable object that represents a complete rendition of media selection options on an asset.
func MutableMediaSelectionFrom(ptr unsafe.Pointer) MutableMediaSelection {
	return MutableMediaSelection{
		MediaSelection: MediaSelectionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MutableMediaSelection *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MutableMediaSelection */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MutableMediaSelection */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MutableMediaSelection */

// Selects the media option in the specified media selection group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMediaSelection/select(_:in:)
func (m_ MutableMediaSelection) SelectMediaOptionInMediaSelectionGroup(mediaSelectionOption IAVMediaSelectionOption, mediaSelectionGroup IAVMediaSelectionGroup) {
	objc.Send[objc.ID](m_.ID, objc.Sel("selectMediaOption:inMediaSelectionGroup:"), mediaSelectionOption, mediaSelectionGroup)
}/* debug [instance_methods/method]: SelectMediaOptionInMediaSelectionGroup */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MutableMediaSelection */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMutableMediaSelection */



