// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

// Enum types and constants
// MPErrorCode - An enumeration that represents error codes for framework operations.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPError/Code
type ErrorCode uint

const (
	// ErrorCancelled - An error that indicates the system canceled the requested operation before it completed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPError/Code/cancelled
	ErrorCancelled ErrorCode = 0
	// ErrorCloudServiceCapabilityMissing - An error that indicates the operation can’t complete because iCloud services aren’t enabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPError/Code/cloudServiceCapabilityMissing
	ErrorCloudServiceCapabilityMissing ErrorCode = 0
	// ErrorNetworkConnectionFailed - An error that indicates the operation failed because the device can’t connect to the network.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPError/Code/networkConnectionFailed
	ErrorNetworkConnectionFailed ErrorCode = 0
	// ErrorNotFound - An error that indicates the operation failed because the system can’t find the requested identifier in the current storefront.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPError/Code/notFound
	ErrorNotFound ErrorCode = 0
	// ErrorNotSupported - An error that indicates the requested operation failed because the system doesn’t support it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPError/Code/notSupported
	ErrorNotSupported ErrorCode = 0
	// ErrorPermissionDenied - An error that indicates the operation can’t complete because the user doesn’t have permission to execute the operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPError/Code/permissionDenied
	ErrorPermissionDenied ErrorCode = 0
	// ErrorRequestTimedOut - An error that indicates the requested operation timed out.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPError/Code/requestTimedOut
	ErrorRequestTimedOut ErrorCode = 0
	// ErrorUnknown - An error that indicates the requested operation can’t complete due to an unknown error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPError/Code/unknown
	ErrorUnknown ErrorCode = 0
)

// MPMediaGrouping - Keys used to configure a media query.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaGrouping
type MediaGrouping uint

const (
	// MediaGroupingAlbum - Groups and sorts media item collections by album, and sorts songs within an album by track order.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaGrouping/album
	MediaGroupingAlbum MediaGrouping = 0
	// MediaGroupingArtist - Groups and sorts media item collections by performing artist.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaGrouping/artist
	MediaGroupingArtist MediaGrouping = 0
	// MediaGroupingGenre - Groups and sorts media item collections by musical or film genre.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaGrouping/genre
	MediaGroupingGenre MediaGrouping = 0
	// MediaGroupingPlaylist - Groups and sorts media item collections by playlist.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaGrouping/playlist
	MediaGroupingPlaylist MediaGrouping = 0
	// MediaGroupingPodcastTitle - Groups and sorts media item collections by podcast title.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaGrouping/podcastTitle
	MediaGroupingPodcastTitle MediaGrouping = 0
	// MediaGroupingTitle - Groups and sorts media item collections by title. For songs, for example, the title is the song name. This is the default grouping key.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaGrouping/title
	MediaGroupingTitle MediaGrouping = 0
)

// MPMediaLibraryAuthorizationStatus - The list of possible states for authorization to access to the user’s media library.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaLibraryAuthorizationStatus
type MediaLibraryAuthorizationStatus uint

const (
	// MediaLibraryAuthorizationStatusAuthorized - Your app may access items in the user’s media library.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaLibraryAuthorizationStatus/authorized
	MediaLibraryAuthorizationStatusAuthorized MediaLibraryAuthorizationStatus = 0
	// MediaLibraryAuthorizationStatusDenied - The app may not access the items in the user’s media library.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaLibraryAuthorizationStatus/denied
	MediaLibraryAuthorizationStatusDenied MediaLibraryAuthorizationStatus = 0
	// MediaLibraryAuthorizationStatusNotDetermined - The user hasn’t determined whether to authorize the use of their media library.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaLibraryAuthorizationStatus/notDetermined
	MediaLibraryAuthorizationStatusNotDetermined MediaLibraryAuthorizationStatus = 0
	// MediaLibraryAuthorizationStatusRestricted - The app may access some of the content in the user’s media library.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaLibraryAuthorizationStatus/restricted
	MediaLibraryAuthorizationStatusRestricted MediaLibraryAuthorizationStatus = 0
)

// MPMediaPlaylistAttribute - Attributes define the type of playlist.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPlaylistAttribute
type MediaPlaylistAttribute uint

// MPMediaPredicateComparison - Logical comparison types for media queries.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPredicateComparison
type MediaPredicateComparison uint

const (
	// MediaPredicateComparisonEqualTo - Matches when a media item’s value for a given property is equal to the value in the media property predicate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPredicateComparison/equalTo
	MediaPredicateComparisonEqualTo MediaPredicateComparison = 0
)

// MPMediaType - The properties for defining the type for a media item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaType
type MediaType uint

const (
	// MediaTypeAny - The media item contains an unspecified type of media content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaType/any
	MediaTypeAny MediaType = 0
	// MediaTypeAudioBook - The media item contains an audio book.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaType/audioBook
	MediaTypeAudioBook MediaType = 0
	// MediaTypeMusic - The media item contains music.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaType/music
	MediaTypeMusic MediaType = 0
	// MediaTypePodcast - The media item contains a podcast.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaType/podcast
	MediaTypePodcast MediaType = 0
)

// MPMovieLoadState - Constants describing the network load state of the movie player.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieLoadState
type MovieLoadState uint

// MPMovieMediaTypeMask - The types of content available in the movie file.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieMediaTypeMask
type MovieMediaTypeMask uint

// MPMusicPlaybackState - The music player playback state modes.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlaybackState
type MusicPlaybackState uint

// MPMusicRepeatMode - The repeat modes for the media player.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicRepeatMode
type MusicRepeatMode uint

// MPMusicShuffleMode - The shuffle modes for the media player.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicShuffleMode
type MusicShuffleMode uint

// MPNowPlayingInfoLanguageOptionType - The language option type to use for the Now Playing item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoLanguageOptionType
type NowPlayingInfoLanguageOptionType uint

// MPNowPlayingInfoMediaType - The type of media currently playing.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoMediaType
type NowPlayingInfoMediaType uint

// MPNowPlayingPlaybackState - The playback state of the app.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingPlaybackState
type NowPlayingPlaybackState uint

// MPRepeatType - Indicates which items to play repeatedly.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRepeatType
type RepeatType uint

// MPShuffleType - Indicates which item types to shuffle.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPShuffleType
type ShuffleType uint


