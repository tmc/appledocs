// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ContentItem] class.
var (
	ContentItemClass     _ContentItemClass
	ContentItemClassOnce sync.Once
)

func getContentItemClass() _ContentItemClass {
	ContentItemClassOnce.Do(func() {
		ContentItemClass = _ContentItemClass{objc.GetClass("MPContentItem")}
	})
	return ContentItemClass
}

type _ContentItemClass struct {
	class objc.Class
}

// An interface definition for the [ContentItem] class.
type IContentItem interface {
	objectivec.IObject
}

// An object that contains the information for a displayed media item.
//
// This object represents a media item such as a song, movie, radio station, or podcast episode. The media player displays the information stored in it. Update this object by changing its properties during runtime or by creating a new object with new property values, but with the same identifier as the object to change. Use the and methods found in to update several objects at once.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPContentItem
type ContentItem struct {
	objectivec.Object
}

// ContentItemFrom constructs a [ContentItem] from an unsafe.Pointer.
//
// An object that contains the information for a displayed media item.
func ContentItemFrom(ptr unsafe.Pointer) ContentItem {
	return ContentItem{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _ContentItemClass) Alloc() ContentItem {
	rv := objc.Send[ContentItem](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ContentItemClass) New() ContentItem {
	rv := objc.Send[ContentItem](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ContentItem) Init() ContentItem {
	rv := objc.Send[ContentItem](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ContentItem) Autorelease() ContentItem {
	rv := objc.Send[ContentItem](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewContentItem creates a new ContentItem instance.
func NewContentItem() ContentItem {
	return getContentItemClass().New()
}




// Sets the identifier for a media item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPContentItem/init(identifier:)
func NewContentItemWithIdentifier(identifier appkit.string) ContentItem {
	instance := getContentItemClass().Alloc()
	rv := objc.Send[ContentItem](instance.ID, objc.Sel("initWithIdentifier:"), identifier)
	rv.Autorelease()
	return rv
}


// A single image that’s associated with the media item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPContentItem/artwork
func (c_ ContentItem) Artwork() MPMediaItemArtwork {
	rv := objc.Send[MPMediaItemArtwork](c_.ID, objc.Sel("artwork"))
	return rv
}


// SetArtwork sets the value of the artwork property.
// A single image that’s associated with the media item.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPContentItem/artwork
func (c_ ContentItem) SetArtwork(value IMPMediaItemArtwork) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setArtwork:"), value)
}

// The unique identifier for the media item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPContentItem/identifier
func (c_ ContentItem) Identifier() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("identifier"))
	return rv
}

// A Boolean value that indicates whether a media item is container of other items.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPContentItem/isContainer
func (c_ ContentItem) Container() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("container"))
	return rv
}


// SetContainer sets the value of the container property.
// A Boolean value that indicates whether a media item is container of other items.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPContentItem/isContainer
func (c_ ContentItem) SetContainer(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContainer:"), value)
}

// A Boolean value that indicates whether the media item contains explicit content.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPContentItem/isExplicitContent
func (c_ ContentItem) ExplicitContent() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("explicitContent"))
	return rv
}


// SetExplicitContent sets the value of the explicitContent property.
// A Boolean value that indicates whether the media item contains explicit content.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPContentItem/isExplicitContent
func (c_ ContentItem) SetExplicitContent(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExplicitContent:"), value)
}

// A Boolean value that indicates whether a media item is able to be played.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPContentItem/isPlayable
func (c_ ContentItem) Playable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("playable"))
	return rv
}


// SetPlayable sets the value of the playable property.
// A Boolean value that indicates whether a media item is able to be played.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPContentItem/isPlayable
func (c_ ContentItem) SetPlayable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPlayable:"), value)
}

// A Boolean value that indicates whether the content item is streaming content.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPContentItem/isStreamingContent
func (c_ ContentItem) StreamingContent() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("streamingContent"))
	return rv
}


// SetStreamingContent sets the value of the streamingContent property.
// A Boolean value that indicates whether the content item is streaming content.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPContentItem/isStreamingContent
func (c_ ContentItem) SetStreamingContent(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStreamingContent:"), value)
}

// The amount of content played for the media item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPContentItem/playbackProgress
func (c_ ContentItem) PlaybackProgress() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("playbackProgress"))
	return rv
}


// SetPlaybackProgress sets the value of the playbackProgress property.
// The amount of content played for the media item.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPContentItem/playbackProgress
func (c_ ContentItem) SetPlaybackProgress(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPlaybackProgress:"), value)
}

// A secondary designator for the media item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPContentItem/subtitle
func (c_ ContentItem) Subtitle() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("subtitle"))
	return rv
}


// SetSubtitle sets the value of the subtitle property.
// A secondary designator for the media item.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPContentItem/subtitle
func (c_ ContentItem) SetSubtitle(value appkit.string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubtitle:"), value)
}

// The public name of the media item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPContentItem/title
func (c_ ContentItem) Title() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("title"))
	return rv
}


// SetTitle sets the value of the title property.
// The public name of the media item.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPContentItem/title
func (c_ ContentItem) SetTitle(value appkit.string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTitle:"), value)
}

// A Boolean value that indicates whether a media item is container of other items.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpcontentitem/iscontainer
func (c_ ContentItem) IsContainer() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isContainer"))
	return rv
}


// SetIsContainer sets the value of the isContainer property.
// A Boolean value that indicates whether a media item is container of other items.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpcontentitem/iscontainer
func (c_ ContentItem) SetIsContainer(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsContainer:"), value)
}

// A Boolean value that indicates whether the media item contains explicit content.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpcontentitem/isexplicitcontent
func (c_ ContentItem) IsExplicitContent() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isExplicitContent"))
	return rv
}


// SetIsExplicitContent sets the value of the isExplicitContent property.
// A Boolean value that indicates whether the media item contains explicit content.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpcontentitem/isexplicitcontent
func (c_ ContentItem) SetIsExplicitContent(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsExplicitContent:"), value)
}

// A Boolean value that indicates whether a media item is able to be played.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpcontentitem/isplayable
func (c_ ContentItem) IsPlayable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isPlayable"))
	return rv
}


// SetIsPlayable sets the value of the isPlayable property.
// A Boolean value that indicates whether a media item is able to be played.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpcontentitem/isplayable
func (c_ ContentItem) SetIsPlayable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsPlayable:"), value)
}

// A Boolean value that indicates whether the content item is streaming content.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpcontentitem/isstreamingcontent
func (c_ ContentItem) IsStreamingContent() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isStreamingContent"))
	return rv
}


// SetIsStreamingContent sets the value of the isStreamingContent property.
// A Boolean value that indicates whether the content item is streaming content.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpcontentitem/isstreamingcontent
func (c_ ContentItem) SetIsStreamingContent(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsStreamingContent:"), value)
}


