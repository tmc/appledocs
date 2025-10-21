// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NowPlayingInfoCenter] class.
var (
	NowPlayingInfoCenterClass     _NowPlayingInfoCenterClass
	NowPlayingInfoCenterClassOnce sync.Once
)

func getNowPlayingInfoCenterClass() _NowPlayingInfoCenterClass {
	NowPlayingInfoCenterClassOnce.Do(func() {
		NowPlayingInfoCenterClass = _NowPlayingInfoCenterClass{objc.GetClass("MPNowPlayingInfoCenter")}
	})
	return NowPlayingInfoCenterClass
}

type _NowPlayingInfoCenterClass struct {
	class objc.Class
}

// An interface definition for the [NowPlayingInfoCenter] class.
type INowPlayingInfoCenter interface {
	objectivec.IObject
}

// An object for setting the Now Playing information for media that your app plays.
//
// If your app also provides Now Playing information containing information about the current track, use this object to update that information at appropriate times. This object contains a dictionary describing the playing item. The system displays Now Playing information on the device’s Lock Screen and in the media controls in Control Center. If the user directs playback of your media to Apple TV using AirPlay, the Now Playing information appears on the television screen. If the user connects a device to an iPod accessory, such as in a car, the accessory may display Now Playing information. The information you can specify includes all of the Now Playing metadata properties (see the Accessing Now Playing metadata properties topic group below), and the following subset of properties: You don’t have direct control over what information the system displays, or its formatting. You set the values in the dictionary and the system or the connected accessory handles displaying the information in a consistent manner for all apps. You can ensure that your app interacts well with other apps providing Now Playing information by following the best practices in the sample code project.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoCenter
type NowPlayingInfoCenter struct {
	objectivec.Object
}

// NowPlayingInfoCenterFrom constructs a [NowPlayingInfoCenter] from an unsafe.Pointer.
//
// An object for setting the Now Playing information for media that your app plays.
func NowPlayingInfoCenterFrom(ptr unsafe.Pointer) NowPlayingInfoCenter {
	return NowPlayingInfoCenter{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NowPlayingInfoCenterClass) Alloc() NowPlayingInfoCenter {
	rv := objc.Send[NowPlayingInfoCenter](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NowPlayingInfoCenterClass) New() NowPlayingInfoCenter {
	rv := objc.Send[NowPlayingInfoCenter](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NowPlayingInfoCenter) Init() NowPlayingInfoCenter {
	rv := objc.Send[NowPlayingInfoCenter](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NowPlayingInfoCenter) Autorelease() NowPlayingInfoCenter {
	rv := objc.Send[NowPlayingInfoCenter](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNowPlayingInfoCenter creates a new NowPlayingInfoCenter instance.
func NewNowPlayingInfoCenter() NowPlayingInfoCenter {
	return getNowPlayingInfoCenterClass().New()
}


// Returns the singleton Now Playing info center.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoCenter/default()
func (nc _NowPlayingInfoCenterClass) DefaultCenter() NowPlayingInfoCenter {
	rv := objc.Send[NowPlayingInfoCenter](objc.ID(nc.class), objc.Sel("defaultCenter"))
	return rv
}

// Keys related to animated artwork that are supported by the current platform.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoCenter/supportedAnimatedArtworkKeys
func (nc _NowPlayingInfoCenterClass) SupportedAnimatedArtworkKeys() []string {
	rv := objc.Send[[]string](objc.ID(nc.class), objc.Sel("supportedAnimatedArtworkKeys"))
	return rv
}
// The current Now Playing information for the default Now Playing info center.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoCenter/nowPlayingInfo
func (n_ NowPlayingInfoCenter) NowPlayingInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("nowPlayingInfo"))
	return rv
}


// SetNowPlayingInfo sets the value of the nowPlayingInfo property.
// The current Now Playing information for the default Now Playing info center.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoCenter/nowPlayingInfo
func (n_ NowPlayingInfoCenter) SetNowPlayingInfo(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNowPlayingInfo:"), value)
}

// The current playback state of the app.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoCenter/playbackState
func (n_ NowPlayingInfoCenter) PlaybackState() NowPlayingPlaybackState {
	rv := objc.Send[NowPlayingPlaybackState](n_.ID, objc.Sel("playbackState"))
	return rv
}


// SetPlaybackState sets the value of the playbackState property.
// The current playback state of the app.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoCenter/playbackState
func (n_ NowPlayingInfoCenter) SetPlaybackState(value NowPlayingPlaybackState) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPlaybackState:"), value)
}

// Keys related to animated artwork that are supported by the current platform.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoCenter/supportedAnimatedArtworkKeys
func (n_ NowPlayingInfoCenter) SupportedAnimatedArtworkKeys() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("supportedAnimatedArtworkKeys"))
	return rv
}

// The title of an album.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitempropertyalbumtitle
func (n_ NowPlayingInfoCenter) MPMediaItemPropertyAlbumTitle() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("MPMediaItemPropertyAlbumTitle"))
	return rv
}

// The number of tracks for the album that contains the media item.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitempropertyalbumtrackcount
func (n_ NowPlayingInfoCenter) MPMediaItemPropertyAlbumTrackCount() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("MPMediaItemPropertyAlbumTrackCount"))
	return rv
}

// The track number of the media item, for a media item that is part of an album.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitempropertyalbumtracknumber
func (n_ NowPlayingInfoCenter) MPMediaItemPropertyAlbumTrackNumber() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("MPMediaItemPropertyAlbumTrackNumber"))
	return rv
}

// The performing artists for a media item — which may vary from the primary artist for the album that a media item belongs to.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitempropertyartist
func (n_ NowPlayingInfoCenter) MPMediaItemPropertyArtist() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("MPMediaItemPropertyArtist"))
	return rv
}

// The artwork image for the media item.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitempropertyartwork
func (n_ NowPlayingInfoCenter) MPMediaItemPropertyArtwork() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("MPMediaItemPropertyArtwork"))
	return rv
}

// The musical composer for the media item.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitempropertycomposer
func (n_ NowPlayingInfoCenter) MPMediaItemPropertyComposer() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("MPMediaItemPropertyComposer"))
	return rv
}

// The number of discs for the album that contains the media item.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitempropertydisccount
func (n_ NowPlayingInfoCenter) MPMediaItemPropertyDiscCount() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("MPMediaItemPropertyDiscCount"))
	return rv
}

// The disc number of the media item, for a media item that is part of a multidisc album.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitempropertydiscnumber
func (n_ NowPlayingInfoCenter) MPMediaItemPropertyDiscNumber() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("MPMediaItemPropertyDiscNumber"))
	return rv
}

// The music or film genre of the media item.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitempropertygenre
func (n_ NowPlayingInfoCenter) MPMediaItemPropertyGenre() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("MPMediaItemPropertyGenre"))
	return rv
}

// The media type of the media item.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitempropertymediatype
func (n_ NowPlayingInfoCenter) MPMediaItemPropertyMediaType() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("MPMediaItemPropertyMediaType"))
	return rv
}

// The key for the persistent identifier for the media item.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitempropertypersistentid
func (n_ NowPlayingInfoCenter) MPMediaItemPropertyPersistentID() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("MPMediaItemPropertyPersistentID"))
	return rv
}

// The playback duration of the media item.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitempropertyplaybackduration
func (n_ NowPlayingInfoCenter) MPMediaItemPropertyPlaybackDuration() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("MPMediaItemPropertyPlaybackDuration"))
	return rv
}

// The title or name of the media item.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitempropertytitle
func (n_ NowPlayingInfoCenter) MPMediaItemPropertyTitle() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("MPMediaItemPropertyTitle"))
	return rv
}

// The identifier of the collection the Now Playing item belongs to.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfocollectionidentifier
func (n_ NowPlayingInfoCenter) MPNowPlayingInfoCollectionIdentifier() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("MPNowPlayingInfoCollectionIdentifier"))
	return rv
}

// 1:1 (square) animated artwork for the current media item.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfoproperty1x1animatedartwork
func (n_ NowPlayingInfoCenter) MPNowPlayingInfoProperty1x1AnimatedArtwork() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("MPNowPlayingInfoProperty1x1AnimatedArtwork"))
	return rv
}

// 3:4 (tall) animated artwork for the current media item.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfoproperty3x4animatedartwork
func (n_ NowPlayingInfoCenter) MPNowPlayingInfoProperty3x4AnimatedArtwork() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("MPNowPlayingInfoProperty3x4AnimatedArtwork"))
	return rv
}

// A list of ad breaks in the Now Playing item.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfopropertyadtimeranges
func (n_ NowPlayingInfoCenter) MPNowPlayingInfoPropertyAdTimeRanges() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("MPNowPlayingInfoPropertyAdTimeRanges"))
	return rv
}

// The URL pointing to the Now Playing item’s underlying asset.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfopropertyasseturl
func (n_ NowPlayingInfoCenter) MPNowPlayingInfoPropertyAssetURL() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("MPNowPlayingInfoPropertyAssetURL"))
	return rv
}

// The available language option groups for the Now Playing item.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfopropertyavailablelanguageoptions
func (n_ NowPlayingInfoCenter) MPNowPlayingInfoPropertyAvailableLanguageOptions() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("MPNowPlayingInfoPropertyAvailableLanguageOptions"))
	return rv
}

// The total number of chapters in the Now Playing item.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfopropertychaptercount
func (n_ NowPlayingInfoCenter) MPNowPlayingInfoPropertyChapterCount() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("MPNowPlayingInfoPropertyChapterCount"))
	return rv
}

// The number corresponding to the currently playing chapter.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfopropertychapternumber
func (n_ NowPlayingInfoCenter) MPNowPlayingInfoPropertyChapterNumber() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("MPNowPlayingInfoPropertyChapterNumber"))
	return rv
}

// The start time for the credits, in seconds, without ads, for the Now Playing item.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfopropertycreditsstarttime
func (n_ NowPlayingInfoCenter) MPNowPlayingInfoPropertyCreditsStartTime() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("MPNowPlayingInfoPropertyCreditsStartTime"))
	return rv
}

// The currently active language options for the Now Playing item.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfopropertycurrentlanguageoptions
func (n_ NowPlayingInfoCenter) MPNowPlayingInfoPropertyCurrentLanguageOptions() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("MPNowPlayingInfoPropertyCurrentLanguageOptions"))
	return rv
}

// The date associated with the current elapsed playback time.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfopropertycurrentplaybackdate
func (n_ NowPlayingInfoCenter) MPNowPlayingInfoPropertyCurrentPlaybackDate() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("MPNowPlayingInfoPropertyCurrentPlaybackDate"))
	return rv
}

// The default playback rate for the Now Playing item.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfopropertydefaultplaybackrate
func (n_ NowPlayingInfoCenter) MPNowPlayingInfoPropertyDefaultPlaybackRate() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("MPNowPlayingInfoPropertyDefaultPlaybackRate"))
	return rv
}

// The elapsed time of the Now Playing item, in seconds.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfopropertyelapsedplaybacktime
func (n_ NowPlayingInfoCenter) MPNowPlayingInfoPropertyElapsedPlaybackTime() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("MPNowPlayingInfoPropertyElapsedPlaybackTime"))
	return rv
}

// A number that denotes whether to exclude the Now Playing item from content suggestions.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfopropertyexcludefromsuggestions
func (n_ NowPlayingInfoCenter) MPNowPlayingInfoPropertyExcludeFromSuggestions() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("MPNowPlayingInfoPropertyExcludeFromSuggestions"))
	return rv
}

// The opaque identifier that uniquely identifies the Now Playing item, even through app relaunches.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfopropertyexternalcontentidentifier
func (n_ NowPlayingInfoCenter) MPNowPlayingInfoPropertyExternalContentIdentifier() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("MPNowPlayingInfoPropertyExternalContentIdentifier"))
	return rv
}

// The opaque identifier that uniquely identifies the profile the Now Playing item plays from, even through app relaunches.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfopropertyexternaluserprofileidentifier
func (n_ NowPlayingInfoCenter) MPNowPlayingInfoPropertyExternalUserProfileIdentifier() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("MPNowPlayingInfoPropertyExternalUserProfileIdentifier"))
	return rv
}

// The International Standard Recording Code (ISRC) of the Now Playing item.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfopropertyinternationalstandardrecordingcode
func (n_ NowPlayingInfoCenter) MPNowPlayingInfoPropertyInternationalStandardRecordingCode() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("MPNowPlayingInfoPropertyInternationalStandardRecordingCode"))
	return rv
}

// A number that denotes whether the Now Playing item is a live stream.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfopropertyislivestream
func (n_ NowPlayingInfoCenter) MPNowPlayingInfoPropertyIsLiveStream() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("MPNowPlayingInfoPropertyIsLiveStream"))
	return rv
}

// The media type of the Now Playing item.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfopropertymediatype
func (n_ NowPlayingInfoCenter) MPNowPlayingInfoPropertyMediaType() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("MPNowPlayingInfoPropertyMediaType"))
	return rv
}

// The current progress of the Now Playing item.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfopropertyplaybackprogress
func (n_ NowPlayingInfoCenter) MPNowPlayingInfoPropertyPlaybackProgress() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("MPNowPlayingInfoPropertyPlaybackProgress"))
	return rv
}

// The total number of items in the app’s playback queue.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfopropertyplaybackqueuecount
func (n_ NowPlayingInfoCenter) MPNowPlayingInfoPropertyPlaybackQueueCount() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("MPNowPlayingInfoPropertyPlaybackQueueCount"))
	return rv
}

// The index of the Now Playing item in the app’s playback queue.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfopropertyplaybackqueueindex
func (n_ NowPlayingInfoCenter) MPNowPlayingInfoPropertyPlaybackQueueIndex() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("MPNowPlayingInfoPropertyPlaybackQueueIndex"))
	return rv
}

// The playback rate of the Now Playing item.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfopropertyplaybackrate
func (n_ NowPlayingInfoCenter) MPNowPlayingInfoPropertyPlaybackRate() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("MPNowPlayingInfoPropertyPlaybackRate"))
	return rv
}

// The service provider associated with the Now Playing item.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfopropertyserviceidentifier
func (n_ NowPlayingInfoCenter) MPNowPlayingInfoPropertyServiceIdentifier() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("MPNowPlayingInfoPropertyServiceIdentifier"))
	return rv
}



