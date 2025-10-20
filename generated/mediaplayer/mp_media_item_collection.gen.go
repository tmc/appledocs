// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [MediaItemCollection] class.
type IMediaItemCollection interface {
	IMediaEntity
}

// A sorted set of media items from the media library.
//
// Typically, you use this class by requesting an array of from a media query by way of its collections property. describes media queries. The grouping type for the media query determines the arrangement of the media items you obtain. You also use the media query property to obtain synced playlists, as described in . A media item collection can have a wide range of metadata associated with it. You access this metadata using the method along with the property keys described in this document. You can also access metadata in a batch fashion using the method. In some cases, this is more efficient. defines and describes both of these methods.
//
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

// Alloc allocates a new instance without initialization.
func (mc _MediaItemCollectionClass) Alloc() MediaItemCollection {
	rv := objc.Send[MediaItemCollection](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Initializes a media item collection with an array of media items.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItemCollection/init(items:)
func NewMediaItemCollectionWithItems(items unsafe.Pointer) MediaItemCollection {
	instance := getMediaItemCollectionClass().Alloc()
	rv := objc.Send[MediaItemCollection](instance.ID, objc.Sel("initWithItems:"), items)
	rv.Autorelease()
	return rv
}


// Creates a media item collection by copying an array of media items.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItemCollection/collectionWithItems:
func (mc _MediaItemCollectionClass) CollectionWithItems(items unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("collectionWithItems:"), items)
	return rv
}

// The number of media items in a collection.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItemCollection/count
func (m_ MediaItemCollection) Count() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("count"))
	return rv
}

// The media items in a media item collection.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItemCollection/items
func (m_ MediaItemCollection) Items() []MediaItem {
	rv := objc.Send[[]MediaItem](m_.ID, objc.Sel("items"))
	return rv
}

// The types of the media items in a collection.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItemCollection/mediaTypes
func (m_ MediaItemCollection) MediaTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("mediaTypes"))
	return rv
}

// A media item whose properties are representative of the other media items in a collection.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItemCollection/representativeItem
func (m_ MediaItemCollection) RepresentativeItem() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("representativeItem"))
	return rv
}


