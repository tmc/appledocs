// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

/* debug [enums.gen.go]: Generating 26 enums for MediaPlayer */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum MPErrorCode (8 cases) */
// MPErrorCode - An enumeration that represents error codes for framework operations.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPError/Code
type MPErrorCode uint

const (
	// MPErrorCancelled - An error that indicates the system canceled the requested operation before it completed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPError/Code/cancelled
	MPErrorCancelled MPErrorCode = 0
	// MPErrorCloudServiceCapabilityMissing - An error that indicates the operation can’t complete because iCloud services aren’t enabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPError/Code/cloudServiceCapabilityMissing
	MPErrorCloudServiceCapabilityMissing MPErrorCode = 0
	// MPErrorNetworkConnectionFailed - An error that indicates the operation failed because the device can’t connect to the network.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPError/Code/networkConnectionFailed
	MPErrorNetworkConnectionFailed MPErrorCode = 0
	// MPErrorNotFound - An error that indicates the operation failed because the system can’t find the requested identifier in the current storefront.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPError/Code/notFound
	MPErrorNotFound MPErrorCode = 0
	// MPErrorNotSupported - An error that indicates the requested operation failed because the system doesn’t support it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPError/Code/notSupported
	MPErrorNotSupported MPErrorCode = 0
	// MPErrorPermissionDenied - An error that indicates the operation can’t complete because the user doesn’t have permission to execute the operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPError/Code/permissionDenied
	MPErrorPermissionDenied MPErrorCode = 0
	// MPErrorRequestTimedOut - An error that indicates the requested operation timed out.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPError/Code/requestTimedOut
	MPErrorRequestTimedOut MPErrorCode = 0
	// MPErrorUnknown - An error that indicates the requested operation can’t complete due to an unknown error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPError/Code/unknown
	MPErrorUnknown MPErrorCode = 0
)

/* debug [enums.gen.go]: Processing enum MPChangeLanguageOptionSetting (3 cases) */
// MPChangeLanguageOptionSetting - The states that determine when language option changes take effect.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPChangeLanguageOptionSetting
type MPChangeLanguageOptionSetting uint

const (
	// MPChangeLanguageOptionSettingNone - No language option change is to be made.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPChangeLanguageOptionSetting/none
	MPChangeLanguageOptionSettingNone MPChangeLanguageOptionSetting = 0
	// MPChangeLanguageOptionSettingNowPlayingItemOnly - The language option change is applied to the now playing item only.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPChangeLanguageOptionSetting/nowPlayingItemOnly
	MPChangeLanguageOptionSettingNowPlayingItemOnly MPChangeLanguageOptionSetting = 0
	// MPChangeLanguageOptionSettingPermanent - The language option change is applied to all future playback items.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPChangeLanguageOptionSetting/permanent
	MPChangeLanguageOptionSettingPermanent MPChangeLanguageOptionSetting = 0
)

/* debug [enums.gen.go]: Processing enum MPMediaGrouping (7 cases) */
// MPMediaGrouping - Keys used to configure a media query.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaGrouping
type MPMediaGrouping uint

const (
	// MPMediaGroupingAlbum - Groups and sorts media item collections by album, and sorts songs within an album by track order.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaGrouping/album
	MPMediaGroupingAlbum MPMediaGrouping = 0
	// MPMediaGroupingAlbumArtist - Groups and sorts media item collections by album artist (the primary performing artist for an album as a whole).
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaGrouping/albumArtist
	MPMediaGroupingAlbumArtist MPMediaGrouping = 0
	// MPMediaGroupingArtist - Groups and sorts media item collections by performing artist.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaGrouping/artist
	MPMediaGroupingArtist MPMediaGrouping = 0
	// MPMediaGroupingGenre - Groups and sorts media item collections by musical or film genre.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaGrouping/genre
	MPMediaGroupingGenre MPMediaGrouping = 0
	// MPMediaGroupingPlaylist - Groups and sorts media item collections by playlist.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaGrouping/playlist
	MPMediaGroupingPlaylist MPMediaGrouping = 0
	// MPMediaGroupingPodcastTitle - Groups and sorts media item collections by podcast title.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaGrouping/podcastTitle
	MPMediaGroupingPodcastTitle MPMediaGrouping = 0
	// MPMediaGroupingTitle - Groups and sorts media item collections by title. For songs, for example, the title is the song name. This is the default grouping key.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaGrouping/title
	MPMediaGroupingTitle MPMediaGrouping = 0
)

/* debug [enums.gen.go]: Processing enum MPMediaLibraryAuthorizationStatus (4 cases) */
// MPMediaLibraryAuthorizationStatus - The list of possible states for authorization to access to the user’s media library.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaLibraryAuthorizationStatus
type MPMediaLibraryAuthorizationStatus uint

const (
	// MPMediaLibraryAuthorizationStatusAuthorized - Your app may access items in the user’s media library.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaLibraryAuthorizationStatus/authorized
	MPMediaLibraryAuthorizationStatusAuthorized MPMediaLibraryAuthorizationStatus = 0
	// MPMediaLibraryAuthorizationStatusDenied - The app may not access the items in the user’s media library.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaLibraryAuthorizationStatus/denied
	MPMediaLibraryAuthorizationStatusDenied MPMediaLibraryAuthorizationStatus = 0
	// MPMediaLibraryAuthorizationStatusNotDetermined - The user hasn’t determined whether to authorize the use of their media library.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaLibraryAuthorizationStatus/notDetermined
	MPMediaLibraryAuthorizationStatusNotDetermined MPMediaLibraryAuthorizationStatus = 0
	// MPMediaLibraryAuthorizationStatusRestricted - The app may access some of the content in the user’s media library.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaLibraryAuthorizationStatus/restricted
	MPMediaLibraryAuthorizationStatusRestricted MPMediaLibraryAuthorizationStatus = 0
)

/* debug [enums.gen.go]: Processing enum MPMediaPlaylistAttribute (4 cases) */
// MPMediaPlaylistAttribute - Attributes define the type of playlist.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPlaylistAttribute
type MPMediaPlaylistAttribute uint

const (
	// MPMediaPlaylistAttributeGenius - A Genius playlist includes items related to other items in your Music library.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPlaylistAttribute/genius
	MPMediaPlaylistAttributeGenius MPMediaPlaylistAttribute = 0
	// MPMediaPlaylistAttributeNone - A playlist with no attributes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPlaylistAttribute/MPMediaPlaylistAttributeNone
	MPMediaPlaylistAttributeNone MPMediaPlaylistAttribute = 0
	// MPMediaPlaylistAttributeOnTheGo - A playlist created on a device rather than synced from the Music app.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPlaylistAttribute/onTheGo
	MPMediaPlaylistAttributeOnTheGo MPMediaPlaylistAttribute = 0
	// MPMediaPlaylistAttributeSmart - A smart playlist includes items that match one or more user-specified rules.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPlaylistAttribute/smart
	MPMediaPlaylistAttributeSmart MPMediaPlaylistAttribute = 0
)

/* debug [enums.gen.go]: Processing enum MPMediaPredicateComparison (2 cases) */
// MPMediaPredicateComparison - Logical comparison types for media queries.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPredicateComparison
type MPMediaPredicateComparison uint

const (
	// MPMediaPredicateComparisonContains - Matches when a media item’s value for a given property is contained in the value of the media property predicate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPredicateComparison/contains
	MPMediaPredicateComparisonContains MPMediaPredicateComparison = 0
	// MPMediaPredicateComparisonEqualTo - Matches when a media item’s value for a given property is equal to the value in the media property predicate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPredicateComparison/equalTo
	MPMediaPredicateComparisonEqualTo MPMediaPredicateComparison = 0
)

/* debug [enums.gen.go]: Processing enum MPMediaType (13 cases) */
// MPMediaType - The properties for defining the type for a media item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaType
type MPMediaType uint

const (
	// MPMediaTypeAny - The media item contains an unspecified type of media content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaType/any
	MPMediaTypeAny MPMediaType = 0
	// MPMediaTypeAnyAudio - The media item contains an unspecified type of audio content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaType/anyAudio
	MPMediaTypeAnyAudio MPMediaType = 0
	// MPMediaTypeAnyVideo - The media item contains an unspecified type of video content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaType/anyVideo
	MPMediaTypeAnyVideo MPMediaType = 0
	// MPMediaTypeAudioBook - The media item contains an audio book.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaType/audioBook
	MPMediaTypeAudioBook MPMediaType = 0
	// MPMediaTypeAudioITunesU - The media item contains an iTunes U audio lesson.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaType/audioITunesU
	MPMediaTypeAudioITunesU MPMediaType = 0
	// MPMediaTypeHomeVideo - The media item contains a home video.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaType/homeVideo
	MPMediaTypeHomeVideo MPMediaType = 0
	// MPMediaTypeMovie - The media item contains a movie.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaType/movie
	MPMediaTypeMovie MPMediaType = 0
	// MPMediaTypeMusic - The media item contains music.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaType/music
	MPMediaTypeMusic MPMediaType = 0
	// MPMediaTypeMusicVideo - The media item contains a music video.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaType/musicVideo
	MPMediaTypeMusicVideo MPMediaType = 0
	// MPMediaTypePodcast - The media item contains a podcast.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaType/podcast
	MPMediaTypePodcast MPMediaType = 0
	// MPMediaTypeTVShow - The media item contains a TV show.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaType/tvShow
	MPMediaTypeTVShow MPMediaType = 0
	// MPMediaTypeVideoITunesU - The media item contains an iTunes U video.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaType/videoITunesU
	MPMediaTypeVideoITunesU MPMediaType = 0
	// MPMediaTypeVideoPodcast - The media item contains a video podcast.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaType/videoPodcast
	MPMediaTypeVideoPodcast MPMediaType = 0
)

/* debug [enums.gen.go]: Processing enum MPMovieControlStyle (4 cases) */
// MPMovieControlStyle - Constants describing the style of the playback controls.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieControlStyle
type MPMovieControlStyle uint

const (
	// MPMovieControlStyleDefault - Fullscreen controls are displayed by default.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieControlStyle/default
	MPMovieControlStyleDefault MPMovieControlStyle = 0
	// MPMovieControlStyleEmbedded - Controls for an embedded view are displayed. The controls include a start/pause button, a scrubber bar, and a button for toggling between fullscreen and embedded display modes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieControlStyle/embedded
	MPMovieControlStyleEmbedded MPMovieControlStyle = 0
	// MPMovieControlStyleFullscreen - Controls for fullscreen playback are displayed. The controls include a start/pause button, a scrubber bar, forward and reverse seeking buttons, a button for toggling between fullscreen and embedded display modes, a button for toggling the aspect fill mode, and a Done button. Tapping the done button pauses the video and exits fullscreen mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieControlStyle/fullscreen
	MPMovieControlStyleFullscreen MPMovieControlStyle = 0
	// MPMovieControlStyleNone - No controls are displayed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieControlStyle/none
	MPMovieControlStyleNone MPMovieControlStyle = 0
)

/* debug [enums.gen.go]: Processing enum MPMovieFinishReason (3 cases) */
// MPMovieFinishReason - Constants describing the reason that playback ended.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieFinishReason
type MPMovieFinishReason uint

const (
	// MPMovieFinishReasonPlaybackEnded - The end of the movie was reached.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieFinishReason/playbackEnded
	MPMovieFinishReasonPlaybackEnded MPMovieFinishReason = 0
	// MPMovieFinishReasonPlaybackError - There was an error during playback.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieFinishReason/playbackError
	MPMovieFinishReasonPlaybackError MPMovieFinishReason = 0
	// MPMovieFinishReasonUserExited - The user stopped playback.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieFinishReason/userExited
	MPMovieFinishReasonUserExited MPMovieFinishReason = 0
)

/* debug [enums.gen.go]: Processing enum MPMovieLoadState (4 cases) */
// MPMovieLoadState - Constants describing the network load state of the movie player.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieLoadState
type MPMovieLoadState uint

const (
	// MPMovieLoadStateUnknown - The load state is not known.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieLoadState/MPMovieLoadStateUnknown
	MPMovieLoadStateUnknown MPMovieLoadState = 0
	// MPMovieLoadStatePlayable - The buffer has enough data that playback can begin, but it may run out of data before playback finishes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieLoadState/playable
	MPMovieLoadStatePlayable MPMovieLoadState = 0
	// MPMovieLoadStatePlaythroughOK - Enough data has been buffered for playback to continue uninterrupted.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieLoadState/playthroughOK
	MPMovieLoadStatePlaythroughOK MPMovieLoadState = 0
	// MPMovieLoadStateStalled - The buffering of data has stalled. If started now, playback may pause automatically if the player runs out of buffered data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieLoadState/stalled
	MPMovieLoadStateStalled MPMovieLoadState = 0
)

/* debug [enums.gen.go]: Processing enum MPMovieMediaTypeMask (3 cases) */
// MPMovieMediaTypeMask - The types of content available in the movie file.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieMediaTypeMask
type MPMovieMediaTypeMask uint

const (
	// MPMovieMediaTypeMaskAudio - The movie file contains audio media.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieMediaTypeMask/audio
	MPMovieMediaTypeMaskAudio MPMovieMediaTypeMask = 0
	// MPMovieMediaTypeMaskNone - The types of media available in the media are not yet known.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieMediaTypeMask/MPMovieMediaTypeMaskNone
	MPMovieMediaTypeMaskNone MPMovieMediaTypeMask = 0
	// MPMovieMediaTypeMaskVideo - The movie file contains video media.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieMediaTypeMask/video
	MPMovieMediaTypeMaskVideo MPMovieMediaTypeMask = 0
)

/* debug [enums.gen.go]: Processing enum MPMoviePlaybackState (6 cases) */
// MPMoviePlaybackState - Constants describing the current playback state of the movie player.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlaybackState
type MPMoviePlaybackState uint

const (
	// MPMoviePlaybackStateInterrupted - Playback is temporarily interrupted, perhaps because the buffer ran out of content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlaybackState/interrupted
	MPMoviePlaybackStateInterrupted MPMoviePlaybackState = 0
	// MPMoviePlaybackStatePaused - Playback is currently paused. Playback will resume from the point where it was paused.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlaybackState/paused
	MPMoviePlaybackStatePaused MPMoviePlaybackState = 0
	// MPMoviePlaybackStatePlaying - Playback is currently under way.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlaybackState/playing
	MPMoviePlaybackStatePlaying MPMoviePlaybackState = 0
	// MPMoviePlaybackStateSeekingBackward - The movie player is currently seeking towards the beginning of the movie.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlaybackState/seekingBackward
	MPMoviePlaybackStateSeekingBackward MPMoviePlaybackState = 0
	// MPMoviePlaybackStateSeekingForward - The movie player is currently seeking towards the end of the movie.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlaybackState/seekingForward
	MPMoviePlaybackStateSeekingForward MPMoviePlaybackState = 0
	// MPMoviePlaybackStateStopped - Playback is currently stopped. Playback will commence from the beginning of the movie.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlaybackState/stopped
	MPMoviePlaybackStateStopped MPMoviePlaybackState = 0
)

/* debug [enums.gen.go]: Processing enum MPMovieRepeatMode (2 cases) */
// MPMovieRepeatMode - Constants describing how the movie player repeats content at the end of playback.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieRepeatMode
type MPMovieRepeatMode uint

const (
	// MPMovieRepeatModeNone - Content is not repeated when playback finishes
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieRepeatMode/none
	MPMovieRepeatModeNone MPMovieRepeatMode = 0
	// MPMovieRepeatModeOne - The current movie is repeated when it finishes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieRepeatMode/one
	MPMovieRepeatModeOne MPMovieRepeatMode = 0
)

/* debug [enums.gen.go]: Processing enum MPMovieScalingMode (4 cases) */
// MPMovieScalingMode - Constants describing how the movie content is scaled to fit the frame of its view.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieScalingMode
type MPMovieScalingMode uint

const (
	// MPMovieScalingModeAspectFill - Scale the movie uniformly until the movie fills the visible bounds of the view. Content at the edges of the larger of the two dimensions is clipped so that the other dimension fits the view exactly. The aspect ratio of the movie is preserved.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieScalingMode/aspectFill
	MPMovieScalingModeAspectFill MPMovieScalingMode = 0
	// MPMovieScalingModeAspectFit - Scale the movie uniformly until one dimension fits the visible bounds of the view exactly. In the other dimension, the region between the edge of the movie and the edge of the view is filled with a black bar. The aspect ratio of the movie is preserved.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieScalingMode/aspectFit
	MPMovieScalingModeAspectFit MPMovieScalingMode = 0
	// MPMovieScalingModeFill - Scale the movie until both dimensions fit the visible bounds of the view exactly. The aspect ratio of the movie is not preserved.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieScalingMode/fill
	MPMovieScalingModeFill MPMovieScalingMode = 0
	// MPMovieScalingModeNone - Do not scale the movie.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieScalingMode/none
	MPMovieScalingModeNone MPMovieScalingMode = 0
)

/* debug [enums.gen.go]: Processing enum MPMovieSourceType (3 cases) */
// MPMovieSourceType - Specifies the type of the movie file.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieSourceType
type MPMovieSourceType uint

const (
	// MPMovieSourceTypeFile - The movie is a local file or is a file that can be downloaded from the network.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieSourceType/file
	MPMovieSourceTypeFile MPMovieSourceType = 0
	// MPMovieSourceTypeStreaming - The movie is a live or on-demand stream.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieSourceType/streaming
	MPMovieSourceTypeStreaming MPMovieSourceType = 0
	// MPMovieSourceTypeUnknown - The movie type is not yet known.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieSourceType/unknown
	MPMovieSourceTypeUnknown MPMovieSourceType = 0
)

/* debug [enums.gen.go]: Processing enum MPMovieTimeOption (2 cases) */
// MPMovieTimeOption - Constants describing which frame to use when generating thumbnail images.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieTimeOption
type MPMovieTimeOption uint

const (
	// MPMovieTimeOptionExact - Use the exact current frame.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieTimeOption/exact
	MPMovieTimeOptionExact MPMovieTimeOption = 0
	// MPMovieTimeOptionNearestKeyFrame - Generate a thumbnail image using the nearest key frame. This frame could be several frames away from the current frame. This option generally offers better performance than trying to find the exact frame.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieTimeOption/nearestKeyFrame
	MPMovieTimeOptionNearestKeyFrame MPMovieTimeOption = 0
)

/* debug [enums.gen.go]: Processing enum MPMusicPlaybackState (6 cases) */
// MPMusicPlaybackState - The music player playback state modes.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlaybackState
type MPMusicPlaybackState uint

const (
	// MPMusicPlaybackStateInterrupted - The music player has been interrupted, such as by an incoming phone call.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlaybackState/interrupted
	MPMusicPlaybackStateInterrupted MPMusicPlaybackState = 0
	// MPMusicPlaybackStatePaused - The music player is paused.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlaybackState/paused
	MPMusicPlaybackStatePaused MPMusicPlaybackState = 0
	// MPMusicPlaybackStatePlaying - The music player is playing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlaybackState/playing
	MPMusicPlaybackStatePlaying MPMusicPlaybackState = 0
	// MPMusicPlaybackStateSeekingBackward - The music player is seeking backward.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlaybackState/seekingBackward
	MPMusicPlaybackStateSeekingBackward MPMusicPlaybackState = 0
	// MPMusicPlaybackStateSeekingForward - The music player is seeking forward.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlaybackState/seekingForward
	MPMusicPlaybackStateSeekingForward MPMusicPlaybackState = 0
	// MPMusicPlaybackStateStopped - The music player is stopped.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlaybackState/stopped
	MPMusicPlaybackStateStopped MPMusicPlaybackState = 0
)

/* debug [enums.gen.go]: Processing enum MPMusicRepeatMode (4 cases) */
// MPMusicRepeatMode - The repeat modes for the media player.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicRepeatMode
type MPMusicRepeatMode uint

const (
	// MPMusicRepeatModeAll - The music player will repeat the current playlist.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicRepeatMode/all
	MPMusicRepeatModeAll MPMusicRepeatMode = 0
	// MPMusicRepeatModeDefault - The user’s preferred repeat mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicRepeatMode/default
	MPMusicRepeatModeDefault MPMusicRepeatMode = 0
	// MPMusicRepeatModeNone - The music player will not repeat the current song or playlist.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicRepeatMode/none
	MPMusicRepeatModeNone MPMusicRepeatMode = 0
	// MPMusicRepeatModeOne - The music player will repeat the current song.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicRepeatMode/one
	MPMusicRepeatModeOne MPMusicRepeatMode = 0
)

/* debug [enums.gen.go]: Processing enum MPMusicShuffleMode (4 cases) */
// MPMusicShuffleMode - The shuffle modes for the media player.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicShuffleMode
type MPMusicShuffleMode uint

const (
	// MPMusicShuffleModeAlbums - The playlist is shuffled by album.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicShuffleMode/albums
	MPMusicShuffleModeAlbums MPMusicShuffleMode = 0
	// MPMusicShuffleModeDefault - The user’s preferred shuffle mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicShuffleMode/default
	MPMusicShuffleModeDefault MPMusicShuffleMode = 0
	// MPMusicShuffleModeOff - The playlist is not shuffled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicShuffleMode/off
	MPMusicShuffleModeOff MPMusicShuffleMode = 0
	// MPMusicShuffleModeSongs - The playlist is shuffled by song.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicShuffleMode/songs
	MPMusicShuffleModeSongs MPMusicShuffleMode = 0
)

/* debug [enums.gen.go]: Processing enum MPNowPlayingInfoLanguageOptionType (2 cases) */
// MPNowPlayingInfoLanguageOptionType - The language option type to use for the Now Playing item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoLanguageOptionType
type MPNowPlayingInfoLanguageOptionType uint

const (
	// MPNowPlayingInfoLanguageOptionTypeAudible - Indicates an audible language option is used.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoLanguageOptionType/audible
	MPNowPlayingInfoLanguageOptionTypeAudible MPNowPlayingInfoLanguageOptionType = 0
	// MPNowPlayingInfoLanguageOptionTypeLegible - Indicates a written language option is used.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoLanguageOptionType/legible
	MPNowPlayingInfoLanguageOptionTypeLegible MPNowPlayingInfoLanguageOptionType = 0
)

/* debug [enums.gen.go]: Processing enum MPNowPlayingInfoMediaType (3 cases) */
// MPNowPlayingInfoMediaType - The type of media currently playing.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoMediaType
type MPNowPlayingInfoMediaType uint

const (
	// MPNowPlayingInfoMediaTypeAudio - The now playing media item is an audio item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoMediaType/audio
	MPNowPlayingInfoMediaTypeAudio MPNowPlayingInfoMediaType = 0
	// MPNowPlayingInfoMediaTypeNone - There is no now playing media item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoMediaType/none
	MPNowPlayingInfoMediaTypeNone MPNowPlayingInfoMediaType = 0
	// MPNowPlayingInfoMediaTypeVideo - The now playing media item is a video item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoMediaType/video
	MPNowPlayingInfoMediaTypeVideo MPNowPlayingInfoMediaType = 0
)

/* debug [enums.gen.go]: Processing enum MPNowPlayingPlaybackState (5 cases) */
// MPNowPlayingPlaybackState - The playback state of the app.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingPlaybackState
type MPNowPlayingPlaybackState uint

const (
	// MPNowPlayingPlaybackStateInterrupted - The app has been interrupted during playback.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingPlaybackState/interrupted
	MPNowPlayingPlaybackStateInterrupted MPNowPlayingPlaybackState = 0
	// MPNowPlayingPlaybackStatePaused - The app is currently paused.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingPlaybackState/paused
	MPNowPlayingPlaybackStatePaused MPNowPlayingPlaybackState = 0
	// MPNowPlayingPlaybackStatePlaying - The app is currently playing a media item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingPlaybackState/playing
	MPNowPlayingPlaybackStatePlaying MPNowPlayingPlaybackState = 0
	// MPNowPlayingPlaybackStateStopped - The app has stopped playing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingPlaybackState/stopped
	MPNowPlayingPlaybackStateStopped MPNowPlayingPlaybackState = 0
	// MPNowPlayingPlaybackStateUnknown - The current state of the app is unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingPlaybackState/unknown
	MPNowPlayingPlaybackStateUnknown MPNowPlayingPlaybackState = 0
)

/* debug [enums.gen.go]: Processing enum MPRemoteCommandHandlerStatus (5 cases) */
// MPRemoteCommandHandlerStatus - Constants indicating the status of a command.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandHandlerStatus
type MPRemoteCommandHandlerStatus uint

const (
	// MPRemoteCommandHandlerStatusCommandFailed - The requested command failed to execute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandHandlerStatus/commandFailed
	MPRemoteCommandHandlerStatusCommandFailed MPRemoteCommandHandlerStatus = 0
	// MPRemoteCommandHandlerStatusDeviceNotFound - The requested command couldn’t execute because a required device isn’t available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandHandlerStatus/deviceNotFound
	MPRemoteCommandHandlerStatusDeviceNotFound MPRemoteCommandHandlerStatus = 0
	// MPRemoteCommandHandlerStatusNoActionableNowPlayingItem - The requested command couldn’t execute because no Now Playing item is available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandHandlerStatus/noActionableNowPlayingItem
	MPRemoteCommandHandlerStatusNoActionableNowPlayingItem MPRemoteCommandHandlerStatus = 0
	// MPRemoteCommandHandlerStatusNoSuchContent - The requested command couldn’t execute because its required content isn’t available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandHandlerStatus/noSuchContent
	MPRemoteCommandHandlerStatusNoSuchContent MPRemoteCommandHandlerStatus = 0
	// MPRemoteCommandHandlerStatusSuccess - The requested command executed successfully.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandHandlerStatus/success
	MPRemoteCommandHandlerStatusSuccess MPRemoteCommandHandlerStatus = 0
)

/* debug [enums.gen.go]: Processing enum MPRepeatType (3 cases) */
// MPRepeatType - Indicates which items to play repeatedly.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRepeatType
type MPRepeatType uint

const (
	// MPRepeatTypeAll - The current container or playlist is repeated indefinitely.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRepeatType/all
	MPRepeatTypeAll MPRepeatType = 0
	// MPRepeatTypeOff - Nothing is repeated during playback.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRepeatType/off
	MPRepeatTypeOff MPRepeatType = 0
	// MPRepeatTypeOne - A single item is repeated indefinitely.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRepeatType/one
	MPRepeatTypeOne MPRepeatType = 0
)

/* debug [enums.gen.go]: Processing enum MPSeekCommandEventType (2 cases) */
// MPSeekCommandEventType - Defines the beginning and ending of seek events.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPSeekCommandEventType
type MPSeekCommandEventType uint

const (
	// MPSeekCommandEventTypeBeginSeeking - Indicates the external media player began seeking.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPSeekCommandEventType/beginSeeking
	MPSeekCommandEventTypeBeginSeeking MPSeekCommandEventType = 0
	// MPSeekCommandEventTypeEndSeeking - Indicates the external media player stopped seeking.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPSeekCommandEventType/endSeeking
	MPSeekCommandEventTypeEndSeeking MPSeekCommandEventType = 0
)

/* debug [enums.gen.go]: Processing enum MPShuffleType (3 cases) */
// MPShuffleType - Indicates which item types to shuffle.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPShuffleType
type MPShuffleType uint

const (
	// MPShuffleTypeCollections - Collections of items are shuffled during playback.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPShuffleType/collections
	MPShuffleTypeCollections MPShuffleType = 0
	// MPShuffleTypeItems - Individual items are shuffled during playback.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPShuffleType/items
	MPShuffleTypeItems MPShuffleType = 0
	// MPShuffleTypeOff - Nothing is shuffled during playback.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPShuffleType/off
	MPShuffleTypeOff MPShuffleType = 0
)


