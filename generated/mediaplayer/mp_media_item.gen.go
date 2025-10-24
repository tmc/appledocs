// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MediaItem] class.
var (
	MediaItemClass     _MediaItemClass
	MediaItemClassOnce sync.Once
)

func getMediaItemClass() _MediaItemClass {
	MediaItemClassOnce.Do(func() {
		MediaItemClass = _MediaItemClass{objc.GetClass("MPMediaItem")}
	})
	return MediaItemClass
}

type _MediaItemClass struct {
	class objc.Class
}

// An interface definition for the [MediaItem] class.
type IMediaItem interface {
	IMediaEntity
	// properties:
	MPMediaEntityPropertyPersistentID() objc.IObject /* cross-framework: NSString */
	Composer() objc.IObject /* cross-framework: NSString */
	SetComposer(value objc.IObject /* cross-framework: NSString */)
	HasProtectedAsset() bool
	SetHasProtectedAsset(value bool)
	IsCloudItem() bool
	SetIsCloudItem(value bool)
	IsCompilation() bool
	SetIsCompilation(value bool)
	IsExplicitItem() bool
	SetIsExplicitItem(value bool)
	IsPreorder() bool
	SetIsPreorder(value bool)
	MPMediaItemPropertyPersistentID() objc.IObject /* cross-framework: NSString */
	// methods:
}

// A collection of properties that represents a single item in the media library.
//
// A media item has an overall unique identifier, accessed using the property key, as well as specific identifiers for its metadata. These identifiers persists across application launches. A media item can have a wide range of metadata associated with it. You access this metadata using the method along with the property keys described in this document. You can also access metadata in a batch fashion using the method. Anytime the app accesses more than one property, enumerating over a set of property keys is more efficient than fetching each individual property. defines both of these methods, the abstract superclass of , and described in . You use attributes of media items to build media queries for searching the Media library. , , and describe these attributes. In addition, describes the property, and describes media queries.


// A collection of properties that represents a single item in the media library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem
type MediaItem struct {
	MediaEntity
}

// MediaItemFrom constructs a [MediaItem] from an unsafe.Pointer.
//
// A collection of properties that represents a single item in the media library.
func MediaItemFrom(ptr unsafe.Pointer) MediaItem {
	return MediaItem{
		MediaEntity: MediaEntityFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MediaItemClass) Alloc() MediaItem {
	rv := objc.Send[MediaItem](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MediaItemClass) New() MediaItem {
	rv := objc.Send[MediaItem](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MediaItem) Init() MediaItem {
	rv := objc.Send[MediaItem](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MediaItem) Autorelease() MediaItem {
	rv := objc.Send[MediaItem](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMediaItem creates a new MediaItem instance.
func NewMediaItem() MediaItem {
	return getMediaItemClass().New()
}



// Obtains the persistent identifier key for a specified grouping type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/persistentIDProperty(forGroupingType:)
func (mc _MediaItemClass) PersistentIDPropertyForGroupingType(groupingType MediaGrouping) objc.IObject /* cross-framework: String */ {
	rv := objc.Send[foundation.String](objc.ID(mc.class), objc.Sel("persistentIDPropertyForGroupingType:"), groupingType)
	return rv
}


// Obtains the title key for a specified grouping type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItem/titleProperty(forGroupingType:)
func (mc _MediaItemClass) TitlePropertyForGroupingType(groupingType MediaGrouping) objc.IObject /* cross-framework: String */ {
	rv := objc.Send[foundation.String](objc.ID(mc.class), objc.Sel("titlePropertyForGroupingType:"), groupingType)
	return rv
}


// The persistent identifier for a media entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaentitypropertypersistentid
func (m_ MediaItem) MPMediaEntityPropertyPersistentID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("MPMediaEntityPropertyPersistentID"))
	return rv
}


// The musical composer for the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitem/composer
func (m_ MediaItem) Composer() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("composer"))
	return rv
}


// The musical composer for the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitem/composer
func (m_ MediaItem) SetComposer(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setComposer:"), value)
}


// A Boolean value that indicates whether the media item has a protected asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitem/hasprotectedasset
func (m_ MediaItem) HasProtectedAsset() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasProtectedAsset"))
	return rv
}


// A Boolean value that indicates whether the media item has a protected asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitem/hasprotectedasset
func (m_ MediaItem) SetHasProtectedAsset(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHasProtectedAsset:"), value)
}


// A Boolean value that indicates whether the media item is an iCloud Music Library item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitem/isclouditem
func (m_ MediaItem) IsCloudItem() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isCloudItem"))
	return rv
}


// A Boolean value that indicates whether the media item is an iCloud Music Library item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitem/isclouditem
func (m_ MediaItem) SetIsCloudItem(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsCloudItem:"), value)
}


// A Boolean value that indicates whether the media item is part of a compilation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitem/iscompilation
func (m_ MediaItem) IsCompilation() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isCompilation"))
	return rv
}


// A Boolean value that indicates whether the media item is part of a compilation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitem/iscompilation
func (m_ MediaItem) SetIsCompilation(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsCompilation:"), value)
}


// A Boolean value that indicates whether the media item has explicit (adult) lyrics or language.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitem/isexplicititem
func (m_ MediaItem) IsExplicitItem() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isExplicitItem"))
	return rv
}


// A Boolean value that indicates whether the media item has explicit (adult) lyrics or language.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitem/isexplicititem
func (m_ MediaItem) SetIsExplicitItem(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsExplicitItem:"), value)
}


// A Boolean value that indicates whether the media item is a preorder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitem/ispreorder
func (m_ MediaItem) IsPreorder() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isPreorder"))
	return rv
}


// A Boolean value that indicates whether the media item is a preorder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitem/ispreorder
func (m_ MediaItem) SetIsPreorder(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsPreorder:"), value)
}


// The key for the persistent identifier for the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitempropertypersistentid
func (m_ MediaItem) MPMediaItemPropertyPersistentID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("MPMediaItemPropertyPersistentID"))
	return rv
}


