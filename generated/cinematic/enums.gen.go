// Code generated from Apple documentation for Cinematic. DO NOT EDIT.

package cinematic

// Enum types and constants
// CNDetectionType - The type of object detected, such as face, torso, cat, dog and so on.
//
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetectionType
type CNDetectionType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetectionType/catBody
	CNDetectionTypeCatBody CNDetectionType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetectionType/unknown
	CNDetectionTypeUnknown CNDetectionType = 0
)

// CNRenderingQuality - The rendering quality, such as thumbnail, preview, export and so on.
//
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNRenderingQuality
type CNRenderingQuality uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNRenderingQuality/export
	CNRenderingQualityExport CNRenderingQuality = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNRenderingQuality/exportHigh
	CNRenderingQualityExportHigh CNRenderingQuality = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNRenderingQuality/preview
	CNRenderingQualityPreview CNRenderingQuality = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNRenderingQuality/thumbnail
	CNRenderingQualityThumbnail CNRenderingQuality = 0
)

// CNSpatialAudioContentType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNSpatialAudioContentType
type CNSpatialAudioContentType uint

const (
	// CNSpatialAudioContentTypeSpatial - Export settings to generate an asset with spatial audio and effect burned in
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNSpatialAudioContentType/spatial
	CNSpatialAudioContentTypeSpatial CNSpatialAudioContentType = 0
	// CNSpatialAudioContentTypeStereo - Export settings to generate an asset with stereo audio and effect burned in
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNSpatialAudioContentType/stereo
	CNSpatialAudioContentTypeStereo CNSpatialAudioContentType = 0
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
	// CNSpatialAudioRenderingStyleStudioBackgroundStem - Isolates the ambience when foreground is studio Audio Mix and place it in a spatial stem. There is no voice stem.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNSpatialAudioRenderingStyle/studioBackgroundStem
	CNSpatialAudioRenderingStyleStudioBackgroundStem CNSpatialAudioRenderingStyle = 0
)


