// Code generated from Apple documentation for PhotosUI. DO NOT EDIT.

package photosui

// Enum types and constants
// PHLivePhotoBadgeOptions - Options for the semantic use and display style of icons for badging Live Photo assets, used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHLivePhotoBadgeOptions
type PHLivePhotoBadgeOptions uint

const (
	// PHLivePhotoBadgeOptionsLiveOff - Return an icon for identifying assets whose additional Live Photo content is disabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHLivePhotoBadgeOptions/liveOff
	PHLivePhotoBadgeOptionsLiveOff PHLivePhotoBadgeOptions = 0
	// PHLivePhotoBadgeOptionsOverContent - Return a variant icon for use on a variable background such as an animating Live Photo view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHLivePhotoBadgeOptions/overContent
	PHLivePhotoBadgeOptionsOverContent PHLivePhotoBadgeOptions = 0
)

// PHLivePhotoViewContentMode - The enumerated Live Photo content modes.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHLivePhotoViewContentMode
type PHLivePhotoViewContentMode uint

const (
	// PHLivePhotoViewContentModeAspectFill - A mode that resizes the content to fill its horizontal or vertical dimension.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHLivePhotoViewContentMode/aspectFill
	PHLivePhotoViewContentModeAspectFill PHLivePhotoViewContentMode = 0
	// PHLivePhotoViewContentModeAspectFit - A mode that resizes the content to fit the view’s bounds, while preserving the content’s aspect ratio.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHLivePhotoViewContentMode/aspectFit
	PHLivePhotoViewContentModeAspectFit PHLivePhotoViewContentMode = 0
)

// PHLivePhotoViewPlaybackStyle - Options for how much of the motion and sound content of a Live Photo to play, used in the 
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHLivePhotoViewPlaybackStyle
type PHLivePhotoViewPlaybackStyle uint

const (
	// PHLivePhotoViewPlaybackStyleFull - Plays back the entire motion and sound content of the Live Photo, including transition effects at the start and end.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHLivePhotoViewPlaybackStyle/full
	PHLivePhotoViewPlaybackStyleFull PHLivePhotoViewPlaybackStyle = 0
	// PHLivePhotoViewPlaybackStyleHint - Plays back only a brief section of the motion content of the Live Photo, without sound.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHLivePhotoViewPlaybackStyle/hint
	PHLivePhotoViewPlaybackStyleHint PHLivePhotoViewPlaybackStyle = 0
	// PHLivePhotoViewPlaybackStyleUndefined - This value is invalid for use.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHLivePhotoViewPlaybackStyle/undefined
	PHLivePhotoViewPlaybackStyleUndefined PHLivePhotoViewPlaybackStyle = 0
)

// PHPickerCapabilities - Options that customize the look and behavior of the photos picker.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerCapabilities
type PHPickerCapabilities uint

const (
	// PHPickerCapabilitiesNone - An option that represents no capabilities.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerCapabilities/PHPickerCapabilitiesNone
	PHPickerCapabilitiesNone PHPickerCapabilities = 0
	// PHPickerCapabilitiesCollectionNavigation - A capability that corresponds to a sidebar or the Albums tab.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerCapabilities/collectionNavigation
	PHPickerCapabilitiesCollectionNavigation PHPickerCapabilities = 0
	// PHPickerCapabilitiesSearch - A capability that corresponds to the search bar.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerCapabilities/search
	PHPickerCapabilitiesSearch PHPickerCapabilities = 0
	// PHPickerCapabilitiesSelectionActions - A cabability that represents the Cancel and Add buttons.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerCapabilities/selectionActions
	PHPickerCapabilitiesSelectionActions PHPickerCapabilities = 0
	// PHPickerCapabilitiesSensitivityAnalysisIntervention - A capability that prompts for confirmation if a person selects a photo that contains nudity.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerCapabilities/sensitivityAnalysisIntervention
	PHPickerCapabilitiesSensitivityAnalysisIntervention PHPickerCapabilities = 0
	// PHPickerCapabilitiesStagingArea - A capability that corresponds to an area in which the selected photos display.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerCapabilities/stagingArea
	PHPickerCapabilitiesStagingArea PHPickerCapabilities = 0
)

// PHPickerConfigurationAssetRepresentationMode - Constants identifying the mode the system uses when many representations exist for an asset.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerConfigurationAssetRepresentationMode
type PHPickerConfigurationAssetRepresentationMode uint

const (
	// PHPickerConfigurationAssetRepresentationModeAutomatic - A mode that indicates that the system chooses the appropriate asset representation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerConfigurationAssetRepresentationMode/automatic
	PHPickerConfigurationAssetRepresentationModeAutomatic PHPickerConfigurationAssetRepresentationMode = 0
	// PHPickerConfigurationAssetRepresentationModeCompatible - A mode that uses the most compatible asset representation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerConfigurationAssetRepresentationMode/compatible
	PHPickerConfigurationAssetRepresentationModeCompatible PHPickerConfigurationAssetRepresentationMode = 0
	// PHPickerConfigurationAssetRepresentationModeCurrent - A mode that uses the current representation to avoid transcoding, if possible.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerConfigurationAssetRepresentationMode/current
	PHPickerConfigurationAssetRepresentationModeCurrent PHPickerConfigurationAssetRepresentationMode = 0
)

// PHPickerConfigurationSelection - Options that represent differing selection behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerConfigurationSelection
type PHPickerConfigurationSelection uint

const (
	// PHPickerConfigurationSelectionContinuous - An option that provides the app a person’s selection immediately.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerConfigurationSelection/continuous
	PHPickerConfigurationSelectionContinuous PHPickerConfigurationSelection = 0
	// PHPickerConfigurationSelectionContinuousAndOrdered - An option that provides the app a person’s selection immediately and displays selected photos with a numbered badge.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerConfigurationSelection/continuousAndOrdered
	PHPickerConfigurationSelectionContinuousAndOrdered PHPickerConfigurationSelection = 0
	// PHPickerConfigurationSelectionDefault - An option that provides selected photos to the app in the default order after the user confirms the selection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerConfigurationSelection/default
	PHPickerConfigurationSelectionDefault PHPickerConfigurationSelection = 0
	// PHPickerConfigurationSelectionOrdered - An option that provides selected photos to the app in the chosen order after the user confirms the selection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerConfigurationSelection/ordered
	PHPickerConfigurationSelectionOrdered PHPickerConfigurationSelection = 0
)

// PHPickerMode - Layout options that determine how the picker orders photos visually.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerMode-c.enum
type PHPickerMode uint

const (
	// PHPickerModeCompact - A linear layout that’s conducive to a smaller area onscreen.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerMode-c.enum/PHPickerModeCompact
	PHPickerModeCompact PHPickerMode = 0
	// PHPickerModeDefault - A grid-based layout that’s conducive to a larger area onscreen.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHPickerMode-c.enum/PHPickerModeDefault
	PHPickerModeDefault PHPickerMode = 0
)

// PHProjectSectionType - The intended usage of the section: cover, content, or auxiliary.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectSection/SectionType-swift.enum
type PHProjectSectionType uint

const (
	// PHProjectSectionTypeAuxiliary - An auxiliary section.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectSection/SectionType-swift.enum/auxiliary
	PHProjectSectionTypeAuxiliary PHProjectSectionType = 0
	// PHProjectSectionTypeContent - A content section.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectSection/SectionType-swift.enum/content
	PHProjectSectionTypeContent PHProjectSectionType = 0
	// PHProjectSectionTypeCover - A cover section.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectSection/SectionType-swift.enum/cover
	PHProjectSectionTypeCover PHProjectSectionType = 0
	// PHProjectSectionTypeUndefined - A blank or undefined section.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectSection/SectionType-swift.enum/undefined
	PHProjectSectionTypeUndefined PHProjectSectionType = 0
)

// PHProjectTextElementType - An enumeration of the type of text element.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectTextElement/ElementType
type PHProjectTextElementType uint


