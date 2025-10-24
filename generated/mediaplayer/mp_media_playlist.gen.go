// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MediaPlaylist] class.
var (
	MediaPlaylistClass     _MediaPlaylistClass
	MediaPlaylistClassOnce sync.Once
)

func getMediaPlaylistClass() _MediaPlaylistClass {
	MediaPlaylistClassOnce.Do(func() {
		MediaPlaylistClass = _MediaPlaylistClass{objc.GetClass("MPMediaPlaylist")}
	})
	return MediaPlaylistClass
}

type _MediaPlaylistClass struct {
	class objc.Class
}

// An interface definition for the [MediaPlaylist] class.
type IMediaPlaylist interface {
	IMediaItemCollection
	// properties:
	// methods:
}

// A playable collection of related media items.
//
// Each playlist has a name, a set of attributes, and a unique identifier that persists across application launches. Users configure playlists using iTunes or by creating a playlist on the device. Playlists are read-only to your iOS app. To obtain playlists, configure a media query that’s grouped by playlist. Each returned media item collection is a media playlist. The following code snippet illustrates this by logging playlist and song names to the Xcode debugger console: and describe the API for building a media query. describes the methods for querying media playlist property values.


// A playable collection of related media items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPlaylist
type MediaPlaylist struct {
	MediaItemCollection
}

// MediaPlaylistFrom constructs a [MediaPlaylist] from an unsafe.Pointer.
//
// A playable collection of related media items.
func MediaPlaylistFrom(ptr unsafe.Pointer) MediaPlaylist {
	return MediaPlaylist{
		MediaItemCollection: MediaItemCollectionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MediaPlaylistClass) Alloc() MediaPlaylist {
	rv := objc.Send[MediaPlaylist](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MediaPlaylistClass) New() MediaPlaylist {
	rv := objc.Send[MediaPlaylist](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MediaPlaylist) Init() MediaPlaylist {
	rv := objc.Send[MediaPlaylist](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MediaPlaylist) Autorelease() MediaPlaylist {
	rv := objc.Send[MediaPlaylist](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMediaPlaylist creates a new MediaPlaylist instance.
func NewMediaPlaylist() MediaPlaylist {
	return getMediaPlaylistClass().New()
}



