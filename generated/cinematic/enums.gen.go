// Code generated from Apple documentation for Cinematic. DO NOT EDIT.

package cinematic

/* debug [enums.gen.go]: Generating 5 enums for Cinematic */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum CNCinematicErrorCode (7 cases) */
// CNCinematicErrorCode enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNCinematicError/Code
type CNCinematicErrorCode uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNCinematicError/Code/cancelled
	CNCinematicErrorCodeCancelled CNCinematicErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNCinematicError/Code/incompatible
	CNCinematicErrorCodeIncompatible CNCinematicErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNCinematicError/Code/incomplete
	CNCinematicErrorCodeIncomplete CNCinematicErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNCinematicError/Code/malformed
	CNCinematicErrorCodeMalformed CNCinematicErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNCinematicError/Code/unknown
	CNCinematicErrorCodeUnknown CNCinematicErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNCinematicError/Code/unreadable
	CNCinematicErrorCodeUnreadable CNCinematicErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNCinematicError/Code/unsupported
	CNCinematicErrorCodeUnsupported CNCinematicErrorCode = 0
)

/* debug [enums.gen.go]: Processing enum CNDetectionType (12 cases) */
// CNDetectionType - The type of object detected, such as face, torso, cat, dog and so on.
//
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetectionType
type CNDetectionType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetectionType/autoFocus
	CNDetectionTypeAutoFocus CNDetectionType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetectionType/catBody
	CNDetectionTypeCatBody CNDetectionType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetectionType/catHead
	CNDetectionTypeCatHead CNDetectionType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetectionType/custom
	CNDetectionTypeCustom CNDetectionType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetectionType/dogBody
	CNDetectionTypeDogBody CNDetectionType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetectionType/dogHead
	CNDetectionTypeDogHead CNDetectionType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetectionType/fixedFocus
	CNDetectionTypeFixedFocus CNDetectionType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetectionType/humanFace
	CNDetectionTypeHumanFace CNDetectionType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetectionType/humanHead
	CNDetectionTypeHumanHead CNDetectionType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetectionType/humanTorso
	CNDetectionTypeHumanTorso CNDetectionType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetectionType/sportsBall
	CNDetectionTypeSportsBall CNDetectionType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetectionType/unknown
	CNDetectionTypeUnknown CNDetectionType = 0
)

/* debug [enums.gen.go]: Processing enum CNSpatialAudioContentType (2 cases) */
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

/* debug [enums.gen.go]: Processing enum CNSpatialAudioRenderingStyle (10 cases) */
// CNSpatialAudioRenderingStyle enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNSpatialAudioRenderingStyle
type CNSpatialAudioRenderingStyle uint

const (
	// CNSpatialAudioRenderingStyleCinematic - Isolates the ambience and place it in a spatial stem. Isolates all voices and place them in a mono stem.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNSpatialAudioRenderingStyle/cinematic
	CNSpatialAudioRenderingStyleCinematic CNSpatialAudioRenderingStyle = 0
	// CNSpatialAudioRenderingStyleCinematicBackgroundStem - Isolates the ambience when foreground is cinematic Audio Mix and place it in a spatial stem. There is no voice stem.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNSpatialAudioRenderingStyle/cinematicBackgroundStem
	CNSpatialAudioRenderingStyleCinematicBackgroundStem CNSpatialAudioRenderingStyle = 0
	// CNSpatialAudioRenderingStyleCinematicForegroundStem - Isolates all voices and places them in a mono stem. There is no ambience stem.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNSpatialAudioRenderingStyle/cinematicForegroundStem
	CNSpatialAudioRenderingStyleCinematicForegroundStem CNSpatialAudioRenderingStyle = 0
	// CNSpatialAudioRenderingStyleInFrame - Isolates the ambience and place it in a spatial stem. Isolates only voices from the camera field of view and place them in a mono stem.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNSpatialAudioRenderingStyle/inFrame
	CNSpatialAudioRenderingStyleInFrame CNSpatialAudioRenderingStyle = 0
	// CNSpatialAudioRenderingStyleInFrameBackgroundStem - Isolates the ambience and foreground that is out of frame and place it in a spatial stem. There is no voice stem.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNSpatialAudioRenderingStyle/inFrameBackgroundStem
	CNSpatialAudioRenderingStyleInFrameBackgroundStem CNSpatialAudioRenderingStyle = 0
	// CNSpatialAudioRenderingStyleInFrameForegroundStem - Isolates only voices from the camera field of view and place them in a mono stem. There is no ambience stem.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNSpatialAudioRenderingStyle/inFrameForegroundStem
	CNSpatialAudioRenderingStyleInFrameForegroundStem CNSpatialAudioRenderingStyle = 0
	// CNSpatialAudioRenderingStyleStandard - This produces a spatial stem of the original recording that is unprocessed. This is the default rendering style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNSpatialAudioRenderingStyle/standard
	CNSpatialAudioRenderingStyleStandard CNSpatialAudioRenderingStyle = 0
	// CNSpatialAudioRenderingStyleStudio - Isolates the ambience and place it in a spatial stem. Isolates all voices, add a studio/proximity effect in the voice track and place them in a mono stem.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNSpatialAudioRenderingStyle/studio
	CNSpatialAudioRenderingStyleStudio CNSpatialAudioRenderingStyle = 0
	// CNSpatialAudioRenderingStyleStudioBackgroundStem - Isolates the ambience when foreground is studio Audio Mix and place it in a spatial stem. There is no voice stem.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNSpatialAudioRenderingStyle/studioBackgroundStem
	CNSpatialAudioRenderingStyleStudioBackgroundStem CNSpatialAudioRenderingStyle = 0
	// CNSpatialAudioRenderingStyleStudioForegroundStem - Isolates all voices, add a studio/proximity effect in the voice track and place them in a mono stem. There is no ambience stem.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNSpatialAudioRenderingStyle/studioForegroundStem
	CNSpatialAudioRenderingStyleStudioForegroundStem CNSpatialAudioRenderingStyle = 0
)

/* debug [enums.gen.go]: Processing enum CNRenderingQuality (4 cases) */
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


