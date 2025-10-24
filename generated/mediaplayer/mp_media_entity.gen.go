// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPMediaEntity */


/* debug [class_header]: Header for MPMediaEntity */
// The class instance for the [MediaEntity] class.
var (
	MediaEntityClass     _MediaEntityClass
	MediaEntityClassOnce sync.Once
)

func getMediaEntityClass() _MediaEntityClass {
	MediaEntityClassOnce.Do(func() {
		MediaEntityClass = _MediaEntityClass{objc.GetClass("MPMediaEntity")}
	})
	return MediaEntityClass
}

type _MediaEntityClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MediaEntity */
// An interface definition for the [MediaEntity] class.
type IMediaEntity interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MediaEntity */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MediaEntity */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MediaEntity */
// Alloc allocates a new instance without initialization.
func (mc _MediaEntityClass) Alloc() MediaEntity {
	rv := objc.Send[MediaEntity](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MediaEntityClass) New() MediaEntity {
	rv := objc.Send[MediaEntity](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MediaEntity) Init() MediaEntity {
	rv := objc.Send[MediaEntity](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MediaEntity) Autorelease() MediaEntity {
	rv := objc.Send[MediaEntity](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMediaEntity creates a new MediaEntity instance.
func NewMediaEntity() MediaEntity {
	return getMediaEntityClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MediaEntity */
// The abstract superclass for media items, media item collections, and media playlist instances.
//
// This is the superclass for and instances, and in turn for instances.


// The abstract superclass for media items, media item collections, and media playlist instances.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaEntity
type MediaEntity struct {
	objectivec.Object
}

// MediaEntityFrom constructs a [MediaEntity] from an unsafe.Pointer.
//
// The abstract superclass for media items, media item collections, and media playlist instances.
func MediaEntityFrom(ptr unsafe.Pointer) MediaEntity {
	return MediaEntity{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MediaEntity *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MediaEntity */

// Indicates whether you can use the media property key that you specify to construct a media property predicate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaEntity/canFilter(byProperty:)
func (mc _MediaEntityClass) CanFilterByProperty(property objc.IObject /* cross-framework: NSString */) bool {
	rv := objc.Send[bool](objc.ID(mc.class), objc.Sel("canFilterByProperty:"), property)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CanFilterByProperty) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MediaEntity */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MediaEntity */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MediaEntity */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPMediaEntity */


