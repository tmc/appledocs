//go:build darwin && ios

// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for MediaPlaylist


// Adds an array of media items to the end of the playlist.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPlaylist/add(_:completionHandler:)
func (m_ MediaPlaylist) AddMediaItemsCompletionHandler(mediaItems []IMediaItem, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addMediaItems:completionHandler:"), mediaItems, completionHandler)
}

// Adds the item associated with the product identifier to the end of the playlist.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPlaylist/addItem(withProductID:completionHandler:)
func (m_ MediaPlaylist) AddItemWithProductIDCompletionHandler(productID objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addItemWithProductID:completionHandler:"), productID, completionHandler)
}

// iOS-only properties

// The display name for the playlist defined in the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPlaylist/authorDisplayName
func (m_ MediaPlaylist) AuthorDisplayName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("authorDisplayName"))
	return rv
}

// The cloud identifier for the playlist.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPlaylist/cloudGlobalID
func (m_ MediaPlaylist) CloudGlobalID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("cloudGlobalID"))
	return rv
}

// User supplied text that describes the playlist.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPlaylist/descriptionText
func (m_ MediaPlaylist) DescriptionText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("descriptionText"))
	return rv
}

// The name of the playlist.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPlaylist/name
func (m_ MediaPlaylist) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}

// The persistent identifier for the playlist.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPlaylist/persistentID
func (m_ MediaPlaylist) PersistentID() objc.IObject /* cross-framework: MediaEntityPersistentID */ {
	rv := objc.Send[MediaEntityPersistentID](m_.ID, objc.Sel("persistentID"))
	return rv
}

// The attributes associated with the playlist.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPlaylist/playlistAttributes
func (m_ MediaPlaylist) PlaylistAttributes() MediaPlaylistAttribute {
	rv := objc.Send[MediaPlaylistAttribute](m_.ID, objc.Sel("playlistAttributes"))
	return rv
}

// The items seeded to generate the playlist; applies only to Genius playlists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPlaylist/seedItems
func (m_ MediaPlaylist) SeedItems() []IMediaItem {
	rv := objc.Send[[]MediaItem](m_.ID, objc.Sel("seedItems"))
	return rv
}





