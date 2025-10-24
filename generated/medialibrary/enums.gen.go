// Code generated from Apple documentation for MediaLibrary. DO NOT EDIT.

package medialibrary

/* debug [enums.gen.go]: Generating 2 enums for MediaLibrary */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum MLMediaSourceType (3 cases) */
// MLMediaSourceType - Specifies the source type associated with a particular media source. Source type reflects the primary type of media within the source. These constants are used to specify values for 
//
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaSourceType
type MLMediaSourceType uint

const (
	// MLMediaSourceTypeAudio - Audio source type. Includes iTunes, GarageBand, and Logic.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaSourceType/audio
	MLMediaSourceTypeAudio MLMediaSourceType = 0
	// MLMediaSourceTypeImage - Image source type. Includes iPhoto, Aperture, and Photo Booth.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaSourceType/image
	MLMediaSourceTypeImage MLMediaSourceType = 0
	// MLMediaSourceTypeMovie - Movie source type. Includes iMovie and Final Cut Pro.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaSourceType/movie
	MLMediaSourceTypeMovie MLMediaSourceType = 0
)

/* debug [enums.gen.go]: Processing enum MLMediaType (3 cases) */
// MLMediaType - Specifies the media type associated with a particular media object. These constants are used to specify a media object’s 
//
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaType
type MLMediaType uint

const (
	// MLMediaTypeAudio - Audio media type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaType/audio
	MLMediaTypeAudio MLMediaType = 0
	// MLMediaTypeImage - Image media type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaType/image
	MLMediaTypeImage MLMediaType = 0
	// MLMediaTypeMovie - Video media type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaType/movie
	MLMediaTypeMovie MLMediaType = 0
)


