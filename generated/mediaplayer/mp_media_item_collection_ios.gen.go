//go:build darwin && ios

// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// iOS-only methods for MediaItemCollection


// iOS-only properties

// The number of media items in a collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItemCollection/count
func (m_ MediaItemCollection) Count() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("count"))
	return rv
}

// The media items in a media item collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItemCollection/items
func (m_ MediaItemCollection) Items() []IMediaItem {
	rv := objc.Send[[]MediaItem](m_.ID, objc.Sel("items"))
	return rv
}

// The types of the media items in a collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItemCollection/mediaTypes
func (m_ MediaItemCollection) MediaTypes() MediaType {
	rv := objc.Send[MediaType](m_.ID, objc.Sel("mediaTypes"))
	return rv
}

// A media item whose properties are representative of the other media items in a collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItemCollection/representativeItem
func (m_ MediaItemCollection) RepresentativeItem() IMPMediaItem {
	rv := objc.Send[MediaItem](m_.ID, objc.Sel("representativeItem"))
	return rv
}




