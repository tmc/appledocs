// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INPlayMediaIntent] class.
var (
	INPlayMediaIntentClass     _INPlayMediaIntentClass
	INPlayMediaIntentClassOnce sync.Once
)

func getINPlayMediaIntentClass() _INPlayMediaIntentClass {
	INPlayMediaIntentClassOnce.Do(func() {
		INPlayMediaIntentClass = _INPlayMediaIntentClass{objc.GetClass("INPlayMediaIntent")}
	})
	return INPlayMediaIntentClass
}

type _INPlayMediaIntentClass struct {
	class objc.Class
}

// An interface definition for the [INPlayMediaIntent] class.
type IINPlayMediaIntent interface {
	IINIntent
	// properties:
	MediaContainer() INMediaItem
	SetMediaContainer(value INMediaItem)
	MediaItems() INMediaItem
	SetMediaItems(value INMediaItem)
	MediaSearch() INMediaSearch
	SetMediaSearch(value INMediaSearch)
	PlayShuffled() bool
	SetPlayShuffled(value bool)
	PlaybackQueueLocation() unsafe.Pointer
	SetPlaybackQueueLocation(value unsafe.Pointer)
	PlaybackRepeatMode() unsafe.Pointer
	SetPlaybackRepeatMode(value unsafe.Pointer)
	PlaybackSpeed() float64
	SetPlaybackSpeed(value float64)
	ResumePlayback() bool
	SetResumePlayback(value bool)
	// methods:
}

// An intent that contains information about media playable from your app.
//
// Use to donate songs, movies, and other media the user plays in your app, or to recommend upcoming media, such as new TV shows or podcast episodes. The system uses these donations to offer the user relevant search results and suggestions. supports playing audio in the background. See for more information about supporting background audio in your app.


// An intent that contains information about media playable from your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INPlayMediaIntent
type INPlayMediaIntent struct {
	INIntent
}

// INPlayMediaIntentFrom constructs a [INPlayMediaIntent] from an unsafe.Pointer.
//
// An intent that contains information about media playable from your app.
func INPlayMediaIntentFrom(ptr unsafe.Pointer) INPlayMediaIntent {
	return INPlayMediaIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INPlayMediaIntentClass) Alloc() INPlayMediaIntent {
	rv := objc.Send[INPlayMediaIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INPlayMediaIntentClass) New() INPlayMediaIntent {
	rv := objc.Send[INPlayMediaIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INPlayMediaIntent) Init() INPlayMediaIntent {
	rv := objc.Send[INPlayMediaIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INPlayMediaIntent) Autorelease() INPlayMediaIntent {
	rv := objc.Send[INPlayMediaIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINPlayMediaIntent creates a new INPlayMediaIntent instance.
func NewINPlayMediaIntent() INPlayMediaIntent {
	return getINPlayMediaIntentClass().New()
}



// The media item container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inplaymediaintent/mediacontainer
func (i_ INPlayMediaIntent) MediaContainer() INMediaItem {
	rv := objc.Send[INMediaItem](i_.ID, objc.Sel("mediaContainer"))
	return rv
}


// The media item container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inplaymediaintent/mediacontainer
func (i_ INPlayMediaIntent) SetMediaContainer(value INMediaItem) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMediaContainer:"), value)
}


// The media content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inplaymediaintent/mediaitems
func (i_ INPlayMediaIntent) MediaItems() INMediaItem {
	rv := objc.Send[INMediaItem](i_.ID, objc.Sel("mediaItems"))
	return rv
}


// The media content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inplaymediaintent/mediaitems
func (i_ INPlayMediaIntent) SetMediaItems(value INMediaItem) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMediaItems:"), value)
}


// An object that contains the search parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inplaymediaintent/mediasearch
func (i_ INPlayMediaIntent) MediaSearch() INMediaSearch {
	rv := objc.Send[INMediaSearch](i_.ID, objc.Sel("mediaSearch"))
	return rv
}


// An object that contains the search parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inplaymediaintent/mediasearch
func (i_ INPlayMediaIntent) SetMediaSearch(value INMediaSearch) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMediaSearch:"), value)
}


// A Boolean value that indicates whether the media plays in a shuffled order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inplaymediaintent/playshuffled-2btex
func (i_ INPlayMediaIntent) PlayShuffled() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("playShuffled"))
	return rv
}


// A Boolean value that indicates whether the media plays in a shuffled order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inplaymediaintent/playshuffled-2btex
func (i_ INPlayMediaIntent) SetPlayShuffled(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPlayShuffled:"), value)
}


// The queue location for a media item during playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inplaymediaintent/playbackqueuelocation
func (i_ INPlayMediaIntent) PlaybackQueueLocation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("playbackQueueLocation"))
	return rv
}


// The queue location for a media item during playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inplaymediaintent/playbackqueuelocation
func (i_ INPlayMediaIntent) SetPlaybackQueueLocation(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPlaybackQueueLocation:"), value)
}


// The repeat mode setting at the time the user plays the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inplaymediaintent/playbackrepeatmode
func (i_ INPlayMediaIntent) PlaybackRepeatMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("playbackRepeatMode"))
	return rv
}


// The repeat mode setting at the time the user plays the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inplaymediaintent/playbackrepeatmode
func (i_ INPlayMediaIntent) SetPlaybackRepeatMode(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPlaybackRepeatMode:"), value)
}


// The playback speed for a media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inplaymediaintent/playbackspeed-17r2t
func (i_ INPlayMediaIntent) PlaybackSpeed() float64 {
	rv := objc.Send[float64](i_.ID, objc.Sel("playbackSpeed"))
	return rv
}


// The playback speed for a media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inplaymediaintent/playbackspeed-17r2t
func (i_ INPlayMediaIntent) SetPlaybackSpeed(value float64) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPlaybackSpeed:"), value)
}


// The resume playback setting at the time the user plays the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inplaymediaintent/resumeplayback-1xw6r
func (i_ INPlayMediaIntent) ResumePlayback() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("resumePlayback"))
	return rv
}


// The resume playback setting at the time the user plays the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inplaymediaintent/resumeplayback-1xw6r
func (i_ INPlayMediaIntent) SetResumePlayback(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setResumePlayback:"), value)
}



