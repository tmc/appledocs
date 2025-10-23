// Code generated from Apple documentation for Cinematic. DO NOT EDIT.

package cinematic

// Enum types and constants
// CNCinematicErrorCode enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNCinematicError/Code
type CNCinematicErrorCode uint

// CNDetectionType - The type of object detected, such as face, torso, cat, dog and so on.
//
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetectionType
type CNDetectionType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetectionType/catBody
	CNDetectionTypeCatBody CNDetectionType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetectionType/humanTorso
	CNDetectionTypeHumanTorso CNDetectionType = 0
)

// CNRenderingQuality - The rendering quality, such as thumbnail, preview, export and so on.
//
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNRenderingQuality
type CNRenderingQuality uint

// CNSpatialAudioContentType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNSpatialAudioContentType
type CNSpatialAudioContentType uint

const (
	// CNSpatialAudioContentTypeSpatial - Export settings to generate an asset with spatial audio and effect burned in
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNSpatialAudioContentType/spatial
	CNSpatialAudioContentTypeSpatial CNSpatialAudioContentType = 0
)

// CNSpatialAudioRenderingStyle enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNSpatialAudioRenderingStyle
type CNSpatialAudioRenderingStyle uint

const (
	// CNSpatialAudioRenderingStyleStandard - This produces a spatial stem of the original recording that is unprocessed. This is the default rendering style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNSpatialAudioRenderingStyle/standard
	CNSpatialAudioRenderingStyleStandard CNSpatialAudioRenderingStyle = 0
	// CNSpatialAudioRenderingStyleStudioForegroundStem - Isolates all voices, add a studio/proximity effect in the voice track and place them in a mono stem. There is no ambience stem.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNSpatialAudioRenderingStyle/studioForegroundStem
	CNSpatialAudioRenderingStyleStudioForegroundStem CNSpatialAudioRenderingStyle = 0
)


