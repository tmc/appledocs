// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	Artwork() IMPMediaItemArtwork
	SetArtwork(value IMPMediaItemArtwork)
	Identifier() objc.IObject /* cross-framework: NSString */
	Container() bool
	SetContainer(value bool)
	ExplicitContent() bool
	SetExplicitContent(value bool)
	Playable() bool
	SetPlayable(value bool)
	StreamingContent() bool
	SetStreamingContent(value bool)
	PlaybackProgress() float32
	SetPlaybackProgress(value float32)
	Subtitle() objc.IObject /* cross-framework: NSString */
	SetSubtitle(value objc.IObject /* cross-framework: NSString */)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	IsContainer() bool
	SetIsContainer(value bool)
	IsExplicitContent() bool
	SetIsExplicitContent(value bool)
	IsPlayable() bool
	SetIsPlayable(value bool)
	IsStreamingContent() bool
	SetIsStreamingContent(value bool)
	// methods:
}

// An object that contains the information for a displayed media item.
//
// This object represents a media item such as a song, movie, radio station, or podcast episode. The media player displays the information stored in it. Update this object by changing its properties during runtime or by creating a new object with new property values, but with the same identifier as the object to change. Use the and methods found in to update several objects at once.


// An object that contains the information for a displayed media item.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPContentItem/init(identifier:)
func NewContentItemWithIdentifier(identifier objc.IObject /* cross-framework: NSString */) ContentItem {
	instance := getContentItemClass().Alloc()
	rv := objc.Send[ContentItem](instance.ID, objc.Sel("initWithIdentifier:"), identifier)
	rv.Autorelease()
	return rv
}



// A single image that’s associated with the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPContentItem/artwork
func (c_ ContentItem) Artwork() IMPMediaItemArtwork {
	rv := objc.Send[MediaItemArtwork](c_.ID, objc.Sel("artwork"))
	return rv
}


// A single image that’s associated with the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPContentItem/artwork
func (c_ ContentItem) SetArtwork(value IMPMediaItemArtwork) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setArtwork:"), value)
}


// The unique identifier for the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPContentItem/identifier
func (c_ ContentItem) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("identifier"))
	return rv
}


// A Boolean value that indicates whether a media item is container of other items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPContentItem/isContainer
func (c_ ContentItem) Container() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("container"))
	return rv
}


// A Boolean value that indicates whether a media item is container of other items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPContentItem/isContainer
func (c_ ContentItem) SetContainer(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContainer:"), value)
}


// A Boolean value that indicates whether the media item contains explicit content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPContentItem/isExplicitContent
func (c_ ContentItem) ExplicitContent() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("explicitContent"))
	return rv
}


// A Boolean value that indicates whether the media item contains explicit content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPContentItem/isExplicitContent
func (c_ ContentItem) SetExplicitContent(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExplicitContent:"), value)
}


// A Boolean value that indicates whether a media item is able to be played.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPContentItem/isPlayable
func (c_ ContentItem) Playable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("playable"))
	return rv
}


// A Boolean value that indicates whether a media item is able to be played.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPContentItem/isPlayable
func (c_ ContentItem) SetPlayable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPlayable:"), value)
}


// A Boolean value that indicates whether the content item is streaming content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPContentItem/isStreamingContent
func (c_ ContentItem) StreamingContent() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("streamingContent"))
	return rv
}


// A Boolean value that indicates whether the content item is streaming content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPContentItem/isStreamingContent
func (c_ ContentItem) SetStreamingContent(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStreamingContent:"), value)
}


// The amount of content played for the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPContentItem/playbackProgress
func (c_ ContentItem) PlaybackProgress() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("playbackProgress"))
	return rv
}


// The amount of content played for the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPContentItem/playbackProgress
func (c_ ContentItem) SetPlaybackProgress(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPlaybackProgress:"), value)
}


// A secondary designator for the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPContentItem/subtitle
func (c_ ContentItem) Subtitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("subtitle"))
	return rv
}


// A secondary designator for the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPContentItem/subtitle
func (c_ ContentItem) SetSubtitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubtitle:"), value)
}


// The public name of the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPContentItem/title
func (c_ ContentItem) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("title"))
	return rv
}


// The public name of the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPContentItem/title
func (c_ ContentItem) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTitle:"), value)
}


// A Boolean value that indicates whether a media item is container of other items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpcontentitem/iscontainer
func (c_ ContentItem) IsContainer() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isContainer"))
	return rv
}


// A Boolean value that indicates whether a media item is container of other items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpcontentitem/iscontainer
func (c_ ContentItem) SetIsContainer(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsContainer:"), value)
}


// A Boolean value that indicates whether the media item contains explicit content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpcontentitem/isexplicitcontent
func (c_ ContentItem) IsExplicitContent() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isExplicitContent"))
	return rv
}


// A Boolean value that indicates whether the media item contains explicit content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpcontentitem/isexplicitcontent
func (c_ ContentItem) SetIsExplicitContent(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsExplicitContent:"), value)
}


// A Boolean value that indicates whether a media item is able to be played.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpcontentitem/isplayable
func (c_ ContentItem) IsPlayable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isPlayable"))
	return rv
}


// A Boolean value that indicates whether a media item is able to be played.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpcontentitem/isplayable
func (c_ ContentItem) SetIsPlayable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsPlayable:"), value)
}


// A Boolean value that indicates whether the content item is streaming content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpcontentitem/isstreamingcontent
func (c_ ContentItem) IsStreamingContent() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isStreamingContent"))
	return rv
}


// A Boolean value that indicates whether the content item is streaming content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpcontentitem/isstreamingcontent
func (c_ ContentItem) SetIsStreamingContent(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsStreamingContent:"), value)
}


