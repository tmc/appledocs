// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

// Enum types and constants
// MPErrorCode - An enumeration that represents error codes for framework operations.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPError/Code
type MPErrorCode uint

// MPMediaGrouping - Keys used to configure a media query.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaGrouping
type MPMediaGrouping uint

const (
	// MPMediaGroupingAlbum - Groups and sorts media item collections by album, and sorts songs within an album by track order.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaGrouping/album
	MPMediaGroupingAlbum MPMediaGrouping = 0
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

// MPMediaLibraryAuthorizationStatus - The list of possible states for authorization to access to the user’s media library.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaLibraryAuthorizationStatus
type MPMediaLibraryAuthorizationStatus uint

// MPMediaPlaylistAttribute - Attributes define the type of playlist.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPlaylistAttribute
type MPMediaPlaylistAttribute uint

// MPMediaPredicateComparison - Logical comparison types for media queries.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPredicateComparison
type MPMediaPredicateComparison uint

// MPMediaType - The properties for defining the type for a media item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaType
type MPMediaType uint

const (
	// MPMediaTypeAny - The media item contains an unspecified type of media content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaType/any
	MPMediaTypeAny MPMediaType = 0
	// MPMediaTypeAudioBook - The media item contains an audio book.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaType/audioBook
	MPMediaTypeAudioBook MPMediaType = 0
	// MPMediaTypeMusic - The media item contains music.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaType/music
	MPMediaTypeMusic MPMediaType = 0
	// MPMediaTypePodcast - The media item contains a podcast.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaType/podcast
	MPMediaTypePodcast MPMediaType = 0
)

// MPMovieLoadState - Constants describing the network load state of the movie player.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieLoadState
type MPMovieLoadState uint

// MPMovieMediaTypeMask - The types of content available in the movie file.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieMediaTypeMask
type MPMovieMediaTypeMask uint

// MPMusicPlaybackState - The music player playback state modes.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlaybackState
type MPMusicPlaybackState uint

// MPMusicRepeatMode - The repeat modes for the media player.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicRepeatMode
type MPMusicRepeatMode uint

// MPMusicShuffleMode - The shuffle modes for the media player.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicShuffleMode
type MPMusicShuffleMode uint

// MPNowPlayingInfoLanguageOptionType - The language option type to use for the Now Playing item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoLanguageOptionType
type MPNowPlayingInfoLanguageOptionType uint

// MPNowPlayingInfoMediaType - The type of media currently playing.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoMediaType
type MPNowPlayingInfoMediaType uint

// MPNowPlayingPlaybackState - The playback state of the app.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingPlaybackState
type MPNowPlayingPlaybackState uint


