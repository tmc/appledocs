// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPMediaPredicate */


/* debug [class_header]: Header for MPMediaPredicate */
// The class instance for the [MediaPredicate] class.
var (
	MediaPredicateClass     _MediaPredicateClass
	MediaPredicateClassOnce sync.Once
)

func getMediaPredicateClass() _MediaPredicateClass {
	MediaPredicateClassOnce.Do(func() {
		MediaPredicateClass = _MediaPredicateClass{objc.GetClass("MPMediaPredicate")}
	})
	return MediaPredicateClass
}

type _MediaPredicateClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MediaPredicate */
// An interface definition for the [MediaPredicate] class.
type IMediaPredicate interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MediaPredicate */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MediaPredicate */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MediaPredicate */
// Alloc allocates a new instance without initialization.
func (mc _MediaPredicateClass) Alloc() MediaPredicate {
	rv := objc.Send[MediaPredicate](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MediaPredicateClass) New() MediaPredicate {
	rv := objc.Send[MediaPredicate](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MediaPredicate) Init() MediaPredicate {
	rv := objc.Send[MediaPredicate](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MediaPredicate) Autorelease() MediaPredicate {
	rv := objc.Send[MediaPredicate](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMediaPredicate creates a new MediaPredicate instance.
func NewMediaPredicate() MediaPredicate {
	return getMediaPredicateClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MediaPredicate */
// An abstract class that defines classes for filtering media in a media query.
//
// In media queries, a is a statement of a logical condition that you want to test each media item against. The system returns the media items that satisfy the condition in the query result. Use this class’s concrete subclass, described in , to define the filter in a media query to retrieve a subset of media items from the library. For more information about media queries, see .


// An abstract class that defines classes for filtering media in a media query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPredicate
type MediaPredicate struct {
	objectivec.Object
}

// MediaPredicateFrom constructs a [MediaPredicate] from an unsafe.Pointer.
//
// An abstract class that defines classes for filtering media in a media query.
func MediaPredicateFrom(ptr unsafe.Pointer) MediaPredicate {
	return MediaPredicate{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MediaPredicate *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MediaPredicate */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MediaPredicate */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MediaPredicate */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MediaPredicate */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPMediaPredicate */



