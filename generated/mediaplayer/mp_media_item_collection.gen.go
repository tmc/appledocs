// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPMediaItemCollection */


/* debug [class_header]: Header for MPMediaItemCollection */
// The class instance for the [MediaItemCollection] class.
var (
	MediaItemCollectionClass     _MediaItemCollectionClass
	MediaItemCollectionClassOnce sync.Once
)

func getMediaItemCollectionClass() _MediaItemCollectionClass {
	MediaItemCollectionClassOnce.Do(func() {
		MediaItemCollectionClass = _MediaItemCollectionClass{objc.GetClass("MPMediaItemCollection")}
	})
	return MediaItemCollectionClass
}

type _MediaItemCollectionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MediaItemCollection */
// An interface definition for the [MediaItemCollection] class.
type IMediaItemCollection interface {
	IMediaEntity
	
/* debug [class_interface_properties]: Properties for MediaItemCollection */
	// properties:
	Collections() IMPMediaItemCollection
	SetCollections(value IMPMediaItemCollection)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MediaItemCollection */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MediaItemCollection */
// Alloc allocates a new instance without initialization.
func (mc _MediaItemCollectionClass) Alloc() MediaItemCollection {
	rv := objc.Send[MediaItemCollection](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MediaItemCollectionClass) New() MediaItemCollection {
	rv := objc.Send[MediaItemCollection](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MediaItemCollection) Init() MediaItemCollection {
	rv := objc.Send[MediaItemCollection](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MediaItemCollection) Autorelease() MediaItemCollection {
	rv := objc.Send[MediaItemCollection](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMediaItemCollection creates a new MediaItemCollection instance.
func NewMediaItemCollection() MediaItemCollection {
	return getMediaItemCollectionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MediaItemCollection */
// A sorted set of media items from the media library.
//
// Typically, you use this class by requesting an array of from a media query by way of its collections property. describes media queries. The grouping type for the media query determines the arrangement of the media items you obtain. You also use the media query property to obtain synced playlists, as described in . A media item collection can have a wide range of metadata associated with it. You access this metadata using the method along with the property keys described in this document. You can also access metadata in a batch fashion using the method. In some cases, this is more efficient. defines and describes both of these methods.


// A sorted set of media items from the media library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItemCollection
type MediaItemCollection struct {
	MediaEntity
}

// MediaItemCollectionFrom constructs a [MediaItemCollection] from an unsafe.Pointer.
//
// A sorted set of media items from the media library.
func MediaItemCollectionFrom(ptr unsafe.Pointer) MediaItemCollection {
	return MediaItemCollection{
		MediaEntity: MediaEntityFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MediaItemCollection */

// Initializes a media item collection with an array of media items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItemCollection/init(items:)
func NewMediaItemCollectionWithItems(items []MediaItem) MediaItemCollection {
	instance := getMediaItemCollectionClass().Alloc()
	rv := objc.Send[MediaItemCollection](instance.ID, objc.Sel("initWithItems:"), items)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMediaItemCollectionWithItems */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MediaItemCollection */

// Creates a media item collection by copying an array of media items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItemCollection/collectionWithItems:
func (mc _MediaItemCollectionClass) CollectionWithItems(items []MediaItem) IMediaItemCollection {
	rv := objc.Send[MediaItemCollection](objc.ID(mc.class), objc.Sel("collectionWithItems:"), items)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CollectionWithItems) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MediaItemCollection */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MediaItemCollection */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MediaItemCollection */

// An array of media item collections whose contained items match the query’s media property predicate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaquery/collections
func (m_ MediaItemCollection) Collections() IMPMediaItemCollection {
	rv := objc.Send[MediaItemCollection](m_.ID, objc.Sel("collections"))
	return rv
}/* debug [instance_properties/getter]: collections */


// An array of media item collections whose contained items match the query’s media property predicate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaquery/collections
func (m_ MediaItemCollection) SetCollections(value IMPMediaItemCollection) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCollections:"), value)
}/* debug [instance_properties/setter]: collections */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPMediaItemCollection */


