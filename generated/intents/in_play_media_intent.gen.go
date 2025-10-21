// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
}

// An intent that contains information about media playable from your app.
//
// Use to donate songs, movies, and other media the user plays in your app, or to recommend upcoming media, such as new TV shows or podcast episodes. The system uses these donations to offer the user relevant search results and suggestions. supports playing audio in the background. See for more information about supporting background audio in your app.
//
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




// Initialize an intent that describes media content such as a song, podcast episode, or movie.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INPlayMediaIntent/initWithMediaItems:mediaContainer:playShuffled:playbackRepeatMode:resumePlayback:playbackQueueLocation:playbackSpeed:mediaSearch:
func NewINPlayMediaIntentWithMediaItemsMediaContainerPlayShuffledPlaybackRepeatModeResumePlaybackPlaybackQueueLocationPlaybackSpeedMediaSearch(mediaItems unsafe.Pointer, mediaContainer unsafe.Pointer, playShuffled foundation.Number, playbackRepeatMode unsafe.Pointer, resumePlayback foundation.Number, playbackQueueLocation unsafe.Pointer, playbackSpeed foundation.Number, mediaSearch unsafe.Pointer) INPlayMediaIntent {
	instance := getINPlayMediaIntentClass().Alloc()
	rv := objc.Send[INPlayMediaIntent](instance.ID, objc.Sel("initWithMediaItems:mediaContainer:playShuffled:playbackRepeatMode:resumePlayback:playbackQueueLocation:playbackSpeed:mediaSearch:"), mediaItems, mediaContainer, playShuffled, playbackRepeatMode, resumePlayback, playbackQueueLocation, playbackSpeed, mediaSearch)
	rv.Autorelease()
	return rv
}


// The media content.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INPlayMediaIntent/mediaItems
func (i_ INPlayMediaIntent) MediaItems() []INMediaItem {
	rv := objc.Send[[]INMediaItem](i_.ID, objc.Sel("mediaItems"))
	return rv
}

// An object that contains the search parameters.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INPlayMediaIntent/mediaSearch
func (i_ INPlayMediaIntent) MediaSearch() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("mediaSearch"))
	return rv
}

// The playback speed for a media item.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INPlayMediaIntent/playbackSpeed-6ngbq
func (i_ INPlayMediaIntent) PlaybackSpeed() foundation.Number {
	rv := objc.Send[foundation.Number](i_.ID, objc.Sel("playbackSpeed"))
	return rv
}

// The resume playback setting at the time the user plays the media item.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INPlayMediaIntent/resumePlayback-9zfyp
func (i_ INPlayMediaIntent) ResumePlayback() foundation.Number {
	rv := objc.Send[foundation.Number](i_.ID, objc.Sel("resumePlayback"))
	return rv
}

// The media item container.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inplaymediaintent/mediacontainer
func (i_ INPlayMediaIntent) MediaContainer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("mediaContainer"))
	return rv
}


// SetMediaContainer sets the value of the mediaContainer property.
// The media item container.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inplaymediaintent/mediacontainer
func (i_ INPlayMediaIntent) SetMediaContainer(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMediaContainer:"), value)
}

// A Boolean value that indicates whether the media plays in a shuffled order.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inplaymediaintent/playshuffled-2btex
func (i_ INPlayMediaIntent) PlayShuffled() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("playShuffled"))
	return rv
}


// SetPlayShuffled sets the value of the playShuffled property.
// A Boolean value that indicates whether the media plays in a shuffled order.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inplaymediaintent/playshuffled-2btex
func (i_ INPlayMediaIntent) SetPlayShuffled(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPlayShuffled:"), value)
}

// The queue location for a media item during playback.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inplaymediaintent/playbackqueuelocation
func (i_ INPlayMediaIntent) PlaybackQueueLocation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("playbackQueueLocation"))
	return rv
}


// SetPlaybackQueueLocation sets the value of the playbackQueueLocation property.
// The queue location for a media item during playback.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inplaymediaintent/playbackqueuelocation
func (i_ INPlayMediaIntent) SetPlaybackQueueLocation(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPlaybackQueueLocation:"), value)
}

// The repeat mode setting at the time the user plays the media item.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inplaymediaintent/playbackrepeatmode
func (i_ INPlayMediaIntent) PlaybackRepeatMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("playbackRepeatMode"))
	return rv
}


// SetPlaybackRepeatMode sets the value of the playbackRepeatMode property.
// The repeat mode setting at the time the user plays the media item.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inplaymediaintent/playbackrepeatmode
func (i_ INPlayMediaIntent) SetPlaybackRepeatMode(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPlaybackRepeatMode:"), value)
}


