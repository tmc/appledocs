// Code generated from Apple documentation for CoreText. DO NOT EDIT.

package coretext

// Enum types and constants
// CTCharacterCollection - Constants that specify character collections.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTCharacterCollection
type CharacterCollection uint

const (
	// kCTCharacterCollectionAdobeCNS1 - The Adobe-CNS1 mapping.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTCharacterCollection/adobeCNS1
	kCTCharacterCollectionAdobeCNS1 CharacterCollection = 0
	// kCTCharacterCollectionAdobeGB1 - The Adobe-GB1 mapping.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTCharacterCollection/adobeGB1
	kCTCharacterCollectionAdobeGB1 CharacterCollection = 0
	// kCTCharacterCollectionAdobeJapan1 - The Adobe-Japan1 mapping.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTCharacterCollection/adobeJapan1
	kCTCharacterCollectionAdobeJapan1 CharacterCollection = 0
	// kCTCharacterCollectionAdobeJapan2 - The Adobe-Japan2 mapping.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTCharacterCollection/adobeJapan2
	kCTCharacterCollectionAdobeJapan2 CharacterCollection = 0
	// kCTCharacterCollectionAdobeKorea1 - The Adobe-Korea1 mapping.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTCharacterCollection/adobeKorea1
	kCTCharacterCollectionAdobeKorea1 CharacterCollection = 0
	// kCTCharacterCollectionIdentityMapping - The character identifier is equal to the glyph index.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTCharacterCollection/identityMapping
	kCTCharacterCollectionIdentityMapping CharacterCollection = 0
	// kCTAdobeCNS1CharacterCollection - The Adobe-CNS1 mapping.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTCharacterCollection/kCTAdobeCNS1CharacterCollection
	kCTAdobeCNS1CharacterCollection CharacterCollection = 0
	// kCTAdobeGB1CharacterCollection - The Adobe-GB1 mapping.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTCharacterCollection/kCTAdobeGB1CharacterCollection
	kCTAdobeGB1CharacterCollection CharacterCollection = 0
	// kCTAdobeJapan1CharacterCollection - The Adobe-Japan1 mapping.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTCharacterCollection/kCTAdobeJapan1CharacterCollection
	kCTAdobeJapan1CharacterCollection CharacterCollection = 0
	// kCTAdobeJapan2CharacterCollection - The Adobe-Japan2 mapping.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTCharacterCollection/kCTAdobeJapan2CharacterCollection
	kCTAdobeJapan2CharacterCollection CharacterCollection = 0
	// kCTAdobeKorea1CharacterCollection - The Adobe-Korea1 mapping.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTCharacterCollection/kCTAdobeKorea1CharacterCollection
	kCTAdobeKorea1CharacterCollection CharacterCollection = 0
	// kCTIdentityMappingCharacterCollection - The character identifier is equal to the glyph index.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTCharacterCollection/kCTIdentityMappingCharacterCollection
	kCTIdentityMappingCharacterCollection CharacterCollection = 0
)

// CTFontCollectionCopyOptions - Option bits for use with CTFontCollectionCopyFontAttribute(s).
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCollectionCopyOptions
type FontCollectionCopyOptions uint

const (
	// kCTFontCollectionCopyDefaultOptions - Passing this option indicates that defaults are to be used.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCollectionCopyOptions/kCTFontCollectionCopyDefaultOptions
	kCTFontCollectionCopyDefaultOptions FontCollectionCopyOptions = 0
	// kCTFontCollectionCopyStandardSort - Passing this option indicates that the return values should be sorted in standard UI order, suitable for display to the user. This is the same sorting behavior used by   and Font Book.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCollectionCopyOptions/standardSort
	kCTFontCollectionCopyStandardSort FontCollectionCopyOptions = 0
	// kCTFontCollectionCopyUnique - Passing this option indicates that duplicate values should be removed from the results.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCollectionCopyOptions/unique
	kCTFontCollectionCopyUnique FontCollectionCopyOptions = 0
)

// CTFontDescriptorMatchingState - Constants that track the progress of font descriptor matching.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorMatchingState
type FontDescriptorMatchingState uint

const (
	// kCTFontDescriptorMatchingDidBegin - A state that indicates matching is about to begin.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorMatchingState/didBegin
	kCTFontDescriptorMatchingDidBegin FontDescriptorMatchingState = 0
	// kCTFontDescriptorMatchingDidFinishDownloading - A state that indicates downloading is done.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorMatchingState/didFinishDownloading
	kCTFontDescriptorMatchingDidFinishDownloading FontDescriptorMatchingState = 0
	// kCTFontDescriptorMatchingDidMatch - A state that indicates the font descriptor match is successful.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorMatchingState/didMatch
	kCTFontDescriptorMatchingDidMatch FontDescriptorMatchingState = 0
	// kCTFontDescriptorMatchingDownloading - A state that indicates downloading is in progress.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorMatchingState/downloading
	kCTFontDescriptorMatchingDownloading FontDescriptorMatchingState = 0
	// kCTFontDescriptorMatchingStalled - A state that indicates that matching is stalled, such as while waiting for a server response.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorMatchingState/stalled
	kCTFontDescriptorMatchingStalled FontDescriptorMatchingState = 0
	// kCTFontDescriptorMatchingWillBeginQuerying - A state that indicates communication with the server is about to begin.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorMatchingState/willBeginQuerying
	kCTFontDescriptorMatchingWillBeginQuerying FontDescriptorMatchingState = 0
)

// CTFontFormat - The recognized format of the font.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontFormat
type FontFormat uint

const (
	// kCTFontFormatBitmap - The font is a bitmap-only format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontFormat/bitmap
	kCTFontFormatBitmap FontFormat = 0
	// kCTFontFormatOpenTypePostScript - The font is an OpenType format containing PostScript data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontFormat/openTypePostScript
	kCTFontFormatOpenTypePostScript FontFormat = 0
	// kCTFontFormatOpenTypeTrueType - The font is an OpenType format containing TrueType data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontFormat/openTypeTrueType
	kCTFontFormatOpenTypeTrueType FontFormat = 0
	// kCTFontFormatPostScript - The font is a recognized PostScript format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontFormat/postScript
	kCTFontFormatPostScript FontFormat = 0
	// kCTFontFormatTrueType - The font is a recognized TrueType format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontFormat/trueType
	kCTFontFormatTrueType FontFormat = 0
	// kCTFontFormatUnrecognized - The font is not a recognized format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontFormat/unrecognized
	kCTFontFormatUnrecognized FontFormat = 0
)

// CTFontManagerAutoActivationSetting - Sets the auto-activation for the specified bundle identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerAutoActivationSetting
type FontManagerAutoActivationSetting uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerAutoActivationSetting/default
	kCTFontManagerAutoActivationDefault FontManagerAutoActivationSetting = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerAutoActivationSetting/disabled
	kCTFontManagerAutoActivationDisabled FontManagerAutoActivationSetting = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerAutoActivationSetting/enabled
	kCTFontManagerAutoActivationEnabled FontManagerAutoActivationSetting = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerAutoActivationSetting/promptUser
	kCTFontManagerAutoActivationPromptUser FontManagerAutoActivationSetting = 0
)

// CTFontManagerError - Errors that prevent unregistration of fonts for a specified font file URL.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerError
type FontManagerError uint

const (
	// kCTFontManagerErrorAlreadyRegistered - An error that indicates the file is already registered in the specified scope.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerError/alreadyRegistered
	kCTFontManagerErrorAlreadyRegistered FontManagerError = 0
	// kCTFontManagerErrorAssetNotFound - An error that indicates the asset isn’t found.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerError/assetNotFound
	kCTFontManagerErrorAssetNotFound FontManagerError = 0
	// kCTFontManagerErrorCancelledByUser - An error that indicates the user cancelled the operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerError/cancelledByUser
	kCTFontManagerErrorCancelledByUser FontManagerError = 0
	// kCTFontManagerErrorDuplicatedName - An error that indicates the file can’t register because of a duplicate font name.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerError/duplicatedName
	kCTFontManagerErrorDuplicatedName FontManagerError = 0
	// kCTFontManagerErrorExceededResourceLimit - An error that indicates an operation failure due to a system limitation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerError/exceededResourceLimit
	kCTFontManagerErrorExceededResourceLimit FontManagerError = 0
	// kCTFontManagerErrorFileNotFound - An error that indicates the file doesn’t exist at the specified URL.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerError/fileNotFound
	kCTFontManagerErrorFileNotFound FontManagerError = 0
	// kCTFontManagerErrorInUse - An error that indicates the font file is actively in use and can’t be unregistered.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerError/inUse
	kCTFontManagerErrorInUse FontManagerError = 0
	// kCTFontManagerErrorInsufficientInfo - An error that indicates the font descriptor doesn’t have the necessary information to specify a font file.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerError/insufficientInfo
	kCTFontManagerErrorInsufficientInfo FontManagerError = 0
	// kCTFontManagerErrorInsufficientPermissions - An error that indicates insufficient permissions to access the file.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerError/insufficientPermissions
	kCTFontManagerErrorInsufficientPermissions FontManagerError = 0
	// kCTFontManagerErrorInvalidFilePath - An error that indicates the file isn’t in an allowed location, which must be either in the app’s bundle or an on-demand resource.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerError/invalidFilePath
	kCTFontManagerErrorInvalidFilePath FontManagerError = 0
	// kCTFontManagerErrorInvalidFontData - An error that indicates the file contains invalid font data that could cause system problems.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerError/invalidFontData
	kCTFontManagerErrorInvalidFontData FontManagerError = 0
	// kCTFontManagerErrorMissingEntitlement - An error that indicates the file can’t be processed because the provider doesn’t have a necessary entitlement.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerError/missingEntitlement
	kCTFontManagerErrorMissingEntitlement FontManagerError = 0
	// kCTFontManagerErrorNotRegistered - An error that indicates the file isn’t registered in the specified scope.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerError/notRegistered
	kCTFontManagerErrorNotRegistered FontManagerError = 0
	// kCTFontManagerErrorRegistrationFailed - An error that indicates the file can’t be processed due to an unexpected FontProvider error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerError/registrationFailed
	kCTFontManagerErrorRegistrationFailed FontManagerError = 0
	// kCTFontManagerErrorSystemRequired - An error that indicates the file is required by the system and can’t be unregistered.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerError/systemRequired
	kCTFontManagerErrorSystemRequired FontManagerError = 0
	// kCTFontManagerErrorUnrecognizedFormat - An error that indicates the file’s format is unrecognized or unsupported.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerError/unrecognizedFormat
	kCTFontManagerErrorUnrecognizedFormat FontManagerError = 0
	// kCTFontManagerErrorUnsupportedScope - An error that indicates the specified scope isn’t supported.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerError/unsupportedScope
	kCTFontManagerErrorUnsupportedScope FontManagerError = 0
)

// CTFontManagerScope - Constants that define the scope for font registration.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerScope
type FontManagerScope uint

const (
	// kCTFontManagerScopeNone - No scope is defined.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerScope/none
	kCTFontManagerScopeNone FontManagerScope = 0
	// kCTFontManagerScopePersistent - The font is available to all processes for the current user session and will be available in subsequent sessions unless unregistered.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerScope/persistent
	kCTFontManagerScopePersistent FontManagerScope = 0
	// kCTFontManagerScopeProcess - The font is available to the current process for the duration of the process unless directly unregistered.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerScope/process
	kCTFontManagerScopeProcess FontManagerScope = 0
	// kCTFontManagerScopeSession - The font is available to the current user session but won’t be available in subsequent sessions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerScope/session
	kCTFontManagerScopeSession FontManagerScope = 0
	// kCTFontManagerScopeUser - The font is available to all processes for the current user session and will be available in subsequent sessions unless unregistered.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerScope/user
	kCTFontManagerScopeUser FontManagerScope = 0
)

// CTFontOptions - Options for font creation and descriptor matching.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontOptions
type FontOptions uint

// CTFontOrientation - The intended rendering orientation of the font for obtaining glyph metrics.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontOrientation
type FontOrientation uint

const (
	// kCTFontOrientationDefault - The native orientation of the font.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontOrientation/default
	kCTFontOrientationDefault FontOrientation = 0
	// kCTFontOrientationHorizontal - The horizontal orientation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontOrientation/horizontal
	kCTFontOrientationHorizontal FontOrientation = 0
	// kCTFontDefaultOrientation - The native orientation of the font.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontOrientation/kCTFontDefaultOrientation
	kCTFontDefaultOrientation FontOrientation = 0
	// kCTFontHorizontalOrientation - The horizontal orientation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontOrientation/kCTFontHorizontalOrientation
	kCTFontHorizontalOrientation FontOrientation = 0
	// kCTFontVerticalOrientation - The vertical orientation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontOrientation/kCTFontVerticalOrientation
	kCTFontVerticalOrientation FontOrientation = 0
	// kCTFontOrientationVertical - The vertical orientation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontOrientation/vertical
	kCTFontOrientationVertical FontOrientation = 0
)

// CTFontStylisticClass - The stylistic class values of the font.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontStylisticClass
type FontStylisticClass uint

// CTFontSymbolicTraits - The symbolic representation of stylistic font attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontSymbolicTraits
type FontSymbolicTraits uint

// CTFontTableOptions - Constants that describe font table options.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontTableOptions
type FontTableOptions uint

// CTFontUIFontType - Constants that represent the specific user-interface purpose to specify for font creation.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType
type FontUIFontType uint

// CTFramePathFillRule - These constants specify the fill rule used by a frame
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFramePathFillRule
type FramePathFillRule uint

const (
	// kCTFramePathFillEvenOdd - Paints the area using the even-odd fill rule.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFramePathFillRule/evenOdd
	kCTFramePathFillEvenOdd FramePathFillRule = 0
	// kCTFramePathFillWindingNumber - Paints the area using the nonzero winding number rule.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFramePathFillRule/windingNumber
	kCTFramePathFillWindingNumber FramePathFillRule = 0
)

// CTFrameProgression - Constants that specify frame progression types.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFrameProgression
type FrameProgression uint

const (
	// kCTFrameProgressionLeftToRight - Lines stack left to right for vertical text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFrameProgression/leftToRight
	kCTFrameProgressionLeftToRight FrameProgression = 0
	// kCTFrameProgressionTopToBottom - Lines stack top to bottom for horizontal text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFrameProgression/topToBottom
	kCTFrameProgressionTopToBottom FrameProgression = 0
)

// CTLineBoundsOptions - Options for getting the bounds of a line of text.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineBoundsOptions
type LineBoundsOptions uint

const (
	// kCTLineBoundsExcludeTypographicLeading - An option to exclude typographic leading.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineBoundsOptions/excludeTypographicLeading
	kCTLineBoundsExcludeTypographicLeading LineBoundsOptions = 0
	// kCTLineBoundsExcludeTypographicShifts - An option to ignore cross-stream shifts due to positioning, such as kerning or baseline alignment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineBoundsOptions/excludeTypographicShifts
	kCTLineBoundsExcludeTypographicShifts LineBoundsOptions = 0
	// kCTLineBoundsIncludeLanguageExtents - An option to include additional space based on common glyph sequences for various languages.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineBoundsOptions/includeLanguageExtents
	kCTLineBoundsIncludeLanguageExtents LineBoundsOptions = 0
	// kCTLineBoundsUseGlyphPathBounds - An option to use glyph path bounds rather than the default typographic bounds.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineBoundsOptions/useGlyphPathBounds
	kCTLineBoundsUseGlyphPathBounds LineBoundsOptions = 0
	// kCTLineBoundsUseHangingPunctuation - An option to enable hanging punctuation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineBoundsOptions/useHangingPunctuation
	kCTLineBoundsUseHangingPunctuation LineBoundsOptions = 0
	// kCTLineBoundsUseOpticalBounds - An option to use optical bounds.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineBoundsOptions/useOpticalBounds
	kCTLineBoundsUseOpticalBounds LineBoundsOptions = 0
)

// CTLineBreakMode - These constants specify what happens when a line is too long for its frame.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineBreakMode
type LineBreakMode uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineBreakMode/byCharWrapping
	kCTLineBreakByCharWrapping LineBreakMode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineBreakMode/byClipping
	kCTLineBreakByClipping LineBreakMode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineBreakMode/byTruncatingHead
	kCTLineBreakByTruncatingHead LineBreakMode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineBreakMode/byTruncatingMiddle
	kCTLineBreakByTruncatingMiddle LineBreakMode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineBreakMode/byTruncatingTail
	kCTLineBreakByTruncatingTail LineBreakMode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineBreakMode/byWordWrapping
	kCTLineBreakByWordWrapping LineBreakMode = 0
)

// CTLineTruncationType - Truncation types required by the 
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineTruncationType
type LineTruncationType uint

// CTParagraphStyleSpecifier - Constants used to query and modify a paragraph style object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier
type ParagraphStyleSpecifier uint

const (
	// kCTParagraphStyleSpecifierAlignment - The text alignment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier/alignment
	kCTParagraphStyleSpecifierAlignment ParagraphStyleSpecifier = 0
	// kCTParagraphStyleSpecifierBaseWritingDirection - The base writing direction of the lines.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier/baseWritingDirection
	kCTParagraphStyleSpecifierBaseWritingDirection ParagraphStyleSpecifier = 0
	// kCTParagraphStyleSpecifierCount - The number of style specifiers.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier/count
	kCTParagraphStyleSpecifierCount ParagraphStyleSpecifier = 0
	// kCTParagraphStyleSpecifierDefaultTabInterval - The document-wide default tab interval.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier/defaultTabInterval
	kCTParagraphStyleSpecifierDefaultTabInterval ParagraphStyleSpecifier = 0
	// kCTParagraphStyleSpecifierFirstLineHeadIndent - The distance, in points, from the leading margin of a frame to the beginning of the paragraph’s first line.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier/firstLineHeadIndent
	kCTParagraphStyleSpecifierFirstLineHeadIndent ParagraphStyleSpecifier = 0
	// kCTParagraphStyleSpecifierHeadIndent - The distance, in points, from the leading margin of a text container to the beginning of lines other than the first.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier/headIndent
	kCTParagraphStyleSpecifierHeadIndent ParagraphStyleSpecifier = 0
	// kCTParagraphStyleSpecifierLineBoundsOptions - Options that control the alignment of the line edges with the leading and trailing margins.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier/lineBoundsOptions
	kCTParagraphStyleSpecifierLineBoundsOptions ParagraphStyleSpecifier = 0
	// kCTParagraphStyleSpecifierLineBreakMode - The mode that should be used to break lines when laying out the paragraph’s text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier/lineBreakMode
	kCTParagraphStyleSpecifierLineBreakMode ParagraphStyleSpecifier = 0
	// kCTParagraphStyleSpecifierLineHeightMultiple - The line height multiple.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier/lineHeightMultiple
	kCTParagraphStyleSpecifierLineHeightMultiple ParagraphStyleSpecifier = 0
	// kCTParagraphStyleSpecifierLineSpacing - The space in points added between lines within the paragraph (commonly known as leading).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier/lineSpacing
	kCTParagraphStyleSpecifierLineSpacing ParagraphStyleSpecifier = 0
	// kCTParagraphStyleSpecifierLineSpacingAdjustment - The space in points added between lines within the paragraph (commonly known as leading).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier/lineSpacingAdjustment
	kCTParagraphStyleSpecifierLineSpacingAdjustment ParagraphStyleSpecifier = 0
	// kCTParagraphStyleSpecifierMaximumLineHeight - The maximum height that any line in the frame will occupy, regardless of the font size or size of any attached graphic.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier/maximumLineHeight
	kCTParagraphStyleSpecifierMaximumLineHeight ParagraphStyleSpecifier = 0
	// kCTParagraphStyleSpecifierMaximumLineSpacing - The maximum space in points between lines within the paragraph (commonly known as leading).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier/maximumLineSpacing
	kCTParagraphStyleSpecifierMaximumLineSpacing ParagraphStyleSpecifier = 0
	// kCTParagraphStyleSpecifierMinimumLineHeight - The minimum height that any line in the frame will occupy, regardless of the font size or size of any attached graphic.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier/minimumLineHeight
	kCTParagraphStyleSpecifierMinimumLineHeight ParagraphStyleSpecifier = 0
	// kCTParagraphStyleSpecifierMinimumLineSpacing - The minimum space in points between lines within the paragraph (commonly known as leading).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier/minimumLineSpacing
	kCTParagraphStyleSpecifierMinimumLineSpacing ParagraphStyleSpecifier = 0
	// kCTParagraphStyleSpecifierParagraphSpacing - The space added at the end of the paragraph to separate it from the following paragraph.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier/paragraphSpacing
	kCTParagraphStyleSpecifierParagraphSpacing ParagraphStyleSpecifier = 0
	// kCTParagraphStyleSpecifierParagraphSpacingBefore - The distance between the paragraph’s top and the beginning of its text content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier/paragraphSpacingBefore
	kCTParagraphStyleSpecifierParagraphSpacingBefore ParagraphStyleSpecifier = 0
	// kCTParagraphStyleSpecifierTabStops - The text tab objects, sorted by location, that define the tab stops for the paragraph style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier/tabStops
	kCTParagraphStyleSpecifierTabStops ParagraphStyleSpecifier = 0
	// kCTParagraphStyleSpecifierTailIndent - The distance, in points, from the margin of a frame to the end of lines.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier/tailIndent
	kCTParagraphStyleSpecifierTailIndent ParagraphStyleSpecifier = 0
)

// CTRubyAlignment - Constants that specify how to align the ruby text and the base text relative to each other when they have different lengths.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyAlignment
type RubyAlignment uint

// CTRubyOverhang - Constants that specify whether, and on which side, ruby text can overhang adjacent text if it’s wider than the base text.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyOverhang
type RubyOverhang uint

const (
	// kCTRubyOverhangAuto - The ruby text can overhang adjacent text on both sides.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyOverhang/auto
	kCTRubyOverhangAuto RubyOverhang = 0
	// kCTRubyOverhangEnd - The ruby text can overhang the text that follows it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyOverhang/end
	kCTRubyOverhangEnd RubyOverhang = 0
	// kCTRubyOverhangInvalid - The overhang specification is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyOverhang/invalid
	kCTRubyOverhangInvalid RubyOverhang = 0
	// kCTRubyOverhangNone - The ruby text can’t overhang the preceding or following text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyOverhang/none
	kCTRubyOverhangNone RubyOverhang = 0
	// kCTRubyOverhangStart - The ruby text can overhang the text that precedes it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyOverhang/start
	kCTRubyOverhangStart RubyOverhang = 0
)

// CTRubyPosition - Constants that specify the position of the ruby text relative to to the base text.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyPosition
type RubyPosition uint

const (
	// kCTRubyPositionAfter - The ruby text is positioned after the base text, appearing below horizontal text and to the left of vertical text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyPosition/after
	kCTRubyPositionAfter RubyPosition = 0
	// kCTRubyPositionBefore - The ruby text is positioned before the base text, appearing above horizontal text and to the right of vertical text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyPosition/before
	kCTRubyPositionBefore RubyPosition = 0
	// kCTRubyPositionCount - A constant that accounts for all ruby positions during ruby annotation creation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyPosition/count
	kCTRubyPositionCount RubyPosition = 0
	// kCTRubyPositionInline - The ruby text follows the base text with no special styling.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyPosition/inline
	kCTRubyPositionInline RubyPosition = 0
	// kCTRubyPositionInterCharacter - The ruby text is positioned to the right of the base text, regardless of whether it’s horizontal or vertical.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyPosition/interCharacter
	kCTRubyPositionInterCharacter RubyPosition = 0
)

// CTRunStatus - A bitfield that represents the disposition of the run.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunStatus
type RunStatus uint

// CTTextAlignment - Constants that specify text alignment.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTextAlignment
type TextAlignment uint

const (
	// kCTTextAlignmentCenter - Text is visually center-aligned.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTextAlignment/center
	kCTTextAlignmentCenter TextAlignment = 0
	// kCTTextAlignmentJustified - Text is fully justified.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTextAlignment/justified
	kCTTextAlignmentJustified TextAlignment = 0
	// kCTNaturalTextAlignment - Text uses the natural alignment of the text’s script.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTextAlignment/kCTNaturalTextAlignment
	kCTNaturalTextAlignment TextAlignment = 0
	// kCTTextAlignmentLeft - Text is visually left-aligned.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTextAlignment/left
	kCTTextAlignmentLeft TextAlignment = 0
	// kCTTextAlignmentNatural - Text uses the natural alignment of the text’s script.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTextAlignment/natural
	kCTTextAlignmentNatural TextAlignment = 0
	// kCTTextAlignmentRight - Text is visually right-aligned.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTextAlignment/right
	kCTTextAlignmentRight TextAlignment = 0
)

// CTUnderlineStyle - Underline style specifiers.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTUnderlineStyle
type UnderlineStyle uint

const (
	// kCTUnderlineStyleNone - A specifier that indicates not to draw an underline.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTUnderlineStyle/kCTUnderlineStyleNone
	kCTUnderlineStyleNone UnderlineStyle = 0
	// kCTUnderlineStyleThick - A specifier that indicates to draw an underline consisting of a thick line.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTUnderlineStyle/thick
	kCTUnderlineStyleThick UnderlineStyle = 0
)

// CTUnderlineStyleModifiers - Underline style modifiers.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTUnderlineStyleModifiers
type UnderlineStyleModifiers uint

const (
	// kCTUnderlinePatternDash - A modifier that indicates to draw an underline using a pattern of dashes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTUnderlineStyleModifiers/patternDash
	kCTUnderlinePatternDash UnderlineStyleModifiers = 0
	// kCTUnderlinePatternDashDot - A modifier that indicates to draw an underline using a pattern of alternating dashes and dots.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTUnderlineStyleModifiers/patternDashDot
	kCTUnderlinePatternDashDot UnderlineStyleModifiers = 0
	// kCTUnderlinePatternDashDotDot - A modifier that indicates to draw an underline using a pattern of a dash followed by two dots.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTUnderlineStyleModifiers/patternDashDotDot
	kCTUnderlinePatternDashDotDot UnderlineStyleModifiers = 0
	// kCTUnderlinePatternDot - A modifier that indicates to draw an underline using a pattern of dots.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTUnderlineStyleModifiers/patternDot
	kCTUnderlinePatternDot UnderlineStyleModifiers = 0
	// kCTUnderlinePatternSolid - A modifier that indicates to draw a solid underline.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTUnderlineStyleModifiers/patternSolid
	kCTUnderlinePatternSolid UnderlineStyleModifiers = 0
)

// CTWritingDirection - These constants specify the writing direction.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTWritingDirection
type WritingDirection uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTWritingDirection/leftToRight
	kCTWritingDirectionLeftToRight WritingDirection = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTWritingDirection/natural
	kCTWritingDirectionNatural WritingDirection = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTWritingDirection/rightToLeft
	kCTWritingDirectionRightToLeft WritingDirection = 0
)


