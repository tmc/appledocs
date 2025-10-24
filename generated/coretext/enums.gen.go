// Code generated from Apple documentation for CoreText. DO NOT EDIT.

package coretext

/* debug [enums.gen.go]: Generating 27 enums for CoreText */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum CTCharacterCollection (12 cases) */
// CTCharacterCollection - Constants that specify character collections.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTCharacterCollection
type CTCharacterCollection uint

const (
	// kCTCharacterCollectionAdobeCNS1 - The Adobe-CNS1 mapping.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTCharacterCollection/adobeCNS1
	kCTCharacterCollectionAdobeCNS1 CTCharacterCollection = 0
	// kCTCharacterCollectionAdobeGB1 - The Adobe-GB1 mapping.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTCharacterCollection/adobeGB1
	kCTCharacterCollectionAdobeGB1 CTCharacterCollection = 0
	// kCTCharacterCollectionAdobeJapan1 - The Adobe-Japan1 mapping.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTCharacterCollection/adobeJapan1
	kCTCharacterCollectionAdobeJapan1 CTCharacterCollection = 0
	// kCTCharacterCollectionAdobeJapan2 - The Adobe-Japan2 mapping.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTCharacterCollection/adobeJapan2
	kCTCharacterCollectionAdobeJapan2 CTCharacterCollection = 0
	// kCTCharacterCollectionAdobeKorea1 - The Adobe-Korea1 mapping.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTCharacterCollection/adobeKorea1
	kCTCharacterCollectionAdobeKorea1 CTCharacterCollection = 0
	// kCTCharacterCollectionIdentityMapping - The character identifier is equal to the glyph index.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTCharacterCollection/identityMapping
	kCTCharacterCollectionIdentityMapping CTCharacterCollection = 0
	// kCTAdobeCNS1CharacterCollection - The Adobe-CNS1 mapping.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTCharacterCollection/kCTAdobeCNS1CharacterCollection
	kCTAdobeCNS1CharacterCollection CTCharacterCollection = 0
	// kCTAdobeGB1CharacterCollection - The Adobe-GB1 mapping.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTCharacterCollection/kCTAdobeGB1CharacterCollection
	kCTAdobeGB1CharacterCollection CTCharacterCollection = 0
	// kCTAdobeJapan1CharacterCollection - The Adobe-Japan1 mapping.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTCharacterCollection/kCTAdobeJapan1CharacterCollection
	kCTAdobeJapan1CharacterCollection CTCharacterCollection = 0
	// kCTAdobeJapan2CharacterCollection - The Adobe-Japan2 mapping.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTCharacterCollection/kCTAdobeJapan2CharacterCollection
	kCTAdobeJapan2CharacterCollection CTCharacterCollection = 0
	// kCTAdobeKorea1CharacterCollection - The Adobe-Korea1 mapping.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTCharacterCollection/kCTAdobeKorea1CharacterCollection
	kCTAdobeKorea1CharacterCollection CTCharacterCollection = 0
	// kCTIdentityMappingCharacterCollection - The character identifier is equal to the glyph index.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTCharacterCollection/kCTIdentityMappingCharacterCollection
	kCTIdentityMappingCharacterCollection CTCharacterCollection = 0
)

/* debug [enums.gen.go]: Processing enum CTFontCollectionCopyOptions (3 cases) */
// CTFontCollectionCopyOptions - Option bits for use with CTFontCollectionCopyFontAttribute(s).
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCollectionCopyOptions
type CTFontCollectionCopyOptions uint

const (
	// kCTFontCollectionCopyDefaultOptions - Passing this option indicates that defaults are to be used.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCollectionCopyOptions/kCTFontCollectionCopyDefaultOptions
	kCTFontCollectionCopyDefaultOptions CTFontCollectionCopyOptions = 0
	// kCTFontCollectionCopyStandardSort - Passing this option indicates that the return values should be sorted in standard UI order, suitable for display to the user. This is the same sorting behavior used by   and Font Book.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCollectionCopyOptions/standardSort
	kCTFontCollectionCopyStandardSort CTFontCollectionCopyOptions = 0
	// kCTFontCollectionCopyUnique - Passing this option indicates that duplicate values should be removed from the results.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCollectionCopyOptions/unique
	kCTFontCollectionCopyUnique CTFontCollectionCopyOptions = 0
)

/* debug [enums.gen.go]: Processing enum CTFontDescriptorMatchingState (9 cases) */
// CTFontDescriptorMatchingState - Constants that track the progress of font descriptor matching.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorMatchingState
type CTFontDescriptorMatchingState uint

const (
	// kCTFontDescriptorMatchingDidBegin - A state that indicates matching is about to begin.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorMatchingState/didBegin
	kCTFontDescriptorMatchingDidBegin CTFontDescriptorMatchingState = 0
	// kCTFontDescriptorMatchingDidFailWithError - A state that indicates an error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorMatchingState/didFailWithError
	kCTFontDescriptorMatchingDidFailWithError CTFontDescriptorMatchingState = 0
	// kCTFontDescriptorMatchingDidFinish - A state that indicates matching is done.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorMatchingState/didFinish
	kCTFontDescriptorMatchingDidFinish CTFontDescriptorMatchingState = 0
	// kCTFontDescriptorMatchingDidFinishDownloading - A state that indicates downloading is done.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorMatchingState/didFinishDownloading
	kCTFontDescriptorMatchingDidFinishDownloading CTFontDescriptorMatchingState = 0
	// kCTFontDescriptorMatchingDidMatch - A state that indicates the font descriptor match is successful.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorMatchingState/didMatch
	kCTFontDescriptorMatchingDidMatch CTFontDescriptorMatchingState = 0
	// kCTFontDescriptorMatchingDownloading - A state that indicates downloading is in progress.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorMatchingState/downloading
	kCTFontDescriptorMatchingDownloading CTFontDescriptorMatchingState = 0
	// kCTFontDescriptorMatchingStalled - A state that indicates that matching is stalled, such as while waiting for a server response.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorMatchingState/stalled
	kCTFontDescriptorMatchingStalled CTFontDescriptorMatchingState = 0
	// kCTFontDescriptorMatchingWillBeginDownloading - A state that indicates downloading is about to begin.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorMatchingState/willBeginDownloading
	kCTFontDescriptorMatchingWillBeginDownloading CTFontDescriptorMatchingState = 0
	// kCTFontDescriptorMatchingWillBeginQuerying - A state that indicates communication with the server is about to begin.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorMatchingState/willBeginQuerying
	kCTFontDescriptorMatchingWillBeginQuerying CTFontDescriptorMatchingState = 0
)

/* debug [enums.gen.go]: Processing enum CTFontFormat (6 cases) */
// CTFontFormat - The recognized format of the font.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontFormat
type CTFontFormat uint

const (
	// kCTFontFormatBitmap - The font is a bitmap-only format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontFormat/bitmap
	kCTFontFormatBitmap CTFontFormat = 0
	// kCTFontFormatOpenTypePostScript - The font is an OpenType format containing PostScript data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontFormat/openTypePostScript
	kCTFontFormatOpenTypePostScript CTFontFormat = 0
	// kCTFontFormatOpenTypeTrueType - The font is an OpenType format containing TrueType data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontFormat/openTypeTrueType
	kCTFontFormatOpenTypeTrueType CTFontFormat = 0
	// kCTFontFormatPostScript - The font is a recognized PostScript format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontFormat/postScript
	kCTFontFormatPostScript CTFontFormat = 0
	// kCTFontFormatTrueType - The font is a recognized TrueType format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontFormat/trueType
	kCTFontFormatTrueType CTFontFormat = 0
	// kCTFontFormatUnrecognized - The font is not a recognized format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontFormat/unrecognized
	kCTFontFormatUnrecognized CTFontFormat = 0
)

/* debug [enums.gen.go]: Processing enum CTFontManagerAutoActivationSetting (4 cases) */
// CTFontManagerAutoActivationSetting - Sets the auto-activation for the specified bundle identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerAutoActivationSetting
type CTFontManagerAutoActivationSetting uint

const (
	// kCTFontManagerAutoActivationDefault - Default auto-activation setting. When specified, the application uses the global setting.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerAutoActivationSetting/default
	kCTFontManagerAutoActivationDefault CTFontManagerAutoActivationSetting = 0
	// kCTFontManagerAutoActivationDisabled - Disables auto-activation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerAutoActivationSetting/disabled
	kCTFontManagerAutoActivationDisabled CTFontManagerAutoActivationSetting = 0
	// kCTFontManagerAutoActivationEnabled - Enables auto-activation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerAutoActivationSetting/enabled
	kCTFontManagerAutoActivationEnabled CTFontManagerAutoActivationSetting = 0
	// kCTFontManagerAutoActivationPromptUser - Requires user input for auto-activation. A dialog is presented to the user to confirm auto-activation of the font.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerAutoActivationSetting/promptUser
	kCTFontManagerAutoActivationPromptUser CTFontManagerAutoActivationSetting = 0
)

/* debug [enums.gen.go]: Processing enum CTFontManagerError (17 cases) */
// CTFontManagerError - Errors that prevent unregistration of fonts for a specified font file URL.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerError
type CTFontManagerError uint

const (
	// kCTFontManagerErrorAlreadyRegistered - An error that indicates the file is already registered in the specified scope.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerError/alreadyRegistered
	kCTFontManagerErrorAlreadyRegistered CTFontManagerError = 0
	// kCTFontManagerErrorAssetNotFound - An error that indicates the asset isn’t found.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerError/assetNotFound
	kCTFontManagerErrorAssetNotFound CTFontManagerError = 0
	// kCTFontManagerErrorCancelledByUser - An error that indicates the user cancelled the operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerError/cancelledByUser
	kCTFontManagerErrorCancelledByUser CTFontManagerError = 0
	// kCTFontManagerErrorDuplicatedName - An error that indicates the file can’t register because of a duplicate font name.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerError/duplicatedName
	kCTFontManagerErrorDuplicatedName CTFontManagerError = 0
	// kCTFontManagerErrorExceededResourceLimit - An error that indicates an operation failure due to a system limitation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerError/exceededResourceLimit
	kCTFontManagerErrorExceededResourceLimit CTFontManagerError = 0
	// kCTFontManagerErrorFileNotFound - An error that indicates the file doesn’t exist at the specified URL.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerError/fileNotFound
	kCTFontManagerErrorFileNotFound CTFontManagerError = 0
	// kCTFontManagerErrorInsufficientInfo - An error that indicates the font descriptor doesn’t have the necessary information to specify a font file.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerError/insufficientInfo
	kCTFontManagerErrorInsufficientInfo CTFontManagerError = 0
	// kCTFontManagerErrorInsufficientPermissions - An error that indicates insufficient permissions to access the file.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerError/insufficientPermissions
	kCTFontManagerErrorInsufficientPermissions CTFontManagerError = 0
	// kCTFontManagerErrorInUse - An error that indicates the font file is actively in use and can’t be unregistered.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerError/inUse
	kCTFontManagerErrorInUse CTFontManagerError = 0
	// kCTFontManagerErrorInvalidFilePath - An error that indicates the file isn’t in an allowed location, which must be either in the app’s bundle or an on-demand resource.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerError/invalidFilePath
	kCTFontManagerErrorInvalidFilePath CTFontManagerError = 0
	// kCTFontManagerErrorInvalidFontData - An error that indicates the file contains invalid font data that could cause system problems.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerError/invalidFontData
	kCTFontManagerErrorInvalidFontData CTFontManagerError = 0
	// kCTFontManagerErrorMissingEntitlement - An error that indicates the file can’t be processed because the provider doesn’t have a necessary entitlement.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerError/missingEntitlement
	kCTFontManagerErrorMissingEntitlement CTFontManagerError = 0
	// kCTFontManagerErrorNotRegistered - An error that indicates the file isn’t registered in the specified scope.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerError/notRegistered
	kCTFontManagerErrorNotRegistered CTFontManagerError = 0
	// kCTFontManagerErrorRegistrationFailed - An error that indicates the file can’t be processed due to an unexpected FontProvider error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerError/registrationFailed
	kCTFontManagerErrorRegistrationFailed CTFontManagerError = 0
	// kCTFontManagerErrorSystemRequired - An error that indicates the file is required by the system and can’t be unregistered.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerError/systemRequired
	kCTFontManagerErrorSystemRequired CTFontManagerError = 0
	// kCTFontManagerErrorUnrecognizedFormat - An error that indicates the file’s format is unrecognized or unsupported.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerError/unrecognizedFormat
	kCTFontManagerErrorUnrecognizedFormat CTFontManagerError = 0
	// kCTFontManagerErrorUnsupportedScope - An error that indicates the specified scope isn’t supported.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerError/unsupportedScope
	kCTFontManagerErrorUnsupportedScope CTFontManagerError = 0
)

/* debug [enums.gen.go]: Processing enum CTFontManagerScope (5 cases) */
// CTFontManagerScope - Constants that define the scope for font registration.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerScope
type CTFontManagerScope uint

const (
	// kCTFontManagerScopeNone - No scope is defined.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerScope/none
	kCTFontManagerScopeNone CTFontManagerScope = 0
	// kCTFontManagerScopePersistent - The font is available to all processes for the current user session and will be available in subsequent sessions unless unregistered.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerScope/persistent
	kCTFontManagerScopePersistent CTFontManagerScope = 0
	// kCTFontManagerScopeProcess - The font is available to the current process for the duration of the process unless directly unregistered.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerScope/process
	kCTFontManagerScopeProcess CTFontManagerScope = 0
	// kCTFontManagerScopeSession - The font is available to the current user session but won’t be available in subsequent sessions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerScope/session
	kCTFontManagerScopeSession CTFontManagerScope = 0
	// kCTFontManagerScopeUser - The font is available to all processes for the current user session and will be available in subsequent sessions unless unregistered.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerScope/user
	kCTFontManagerScopeUser CTFontManagerScope = 0
)

/* debug [enums.gen.go]: Processing enum CTFontOptions (4 cases) */
// CTFontOptions - Options for font creation and descriptor matching.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontOptions
type CTFontOptions uint

const (
	// kCTFontOptionsDefault - Default options are used.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontOptions/kCTFontOptionsDefault
	kCTFontOptionsDefault CTFontOptions = 0
	// kCTFontOptionsPreferSystemFont - Font matching prefers to match Apple system fonts.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontOptions/preferSystemFont
	kCTFontOptionsPreferSystemFont CTFontOptions = 0
	// kCTFontOptionsPreventAutoActivation - Prevents automatic font activation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontOptions/preventAutoActivation
	kCTFontOptionsPreventAutoActivation CTFontOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontOptions/preventAutoDownload
	kCTFontOptionsPreventAutoDownload CTFontOptions = 0
)

/* debug [enums.gen.go]: Processing enum CTFontOrientation (6 cases) */
// CTFontOrientation - The intended rendering orientation of the font for obtaining glyph metrics.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontOrientation
type CTFontOrientation uint

const (
	// kCTFontOrientationDefault - The native orientation of the font.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontOrientation/default
	kCTFontOrientationDefault CTFontOrientation = 0
	// kCTFontOrientationHorizontal - The horizontal orientation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontOrientation/horizontal
	kCTFontOrientationHorizontal CTFontOrientation = 0
	// kCTFontDefaultOrientation - The native orientation of the font.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontOrientation/kCTFontDefaultOrientation
	kCTFontDefaultOrientation CTFontOrientation = 0
	// kCTFontHorizontalOrientation - The horizontal orientation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontOrientation/kCTFontHorizontalOrientation
	kCTFontHorizontalOrientation CTFontOrientation = 0
	// kCTFontVerticalOrientation - The vertical orientation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontOrientation/kCTFontVerticalOrientation
	kCTFontVerticalOrientation CTFontOrientation = 0
	// kCTFontOrientationVertical - The vertical orientation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontOrientation/vertical
	kCTFontOrientationVertical CTFontOrientation = 0
)

/* debug [enums.gen.go]: Processing enum CTFontStylisticClass (22 cases) */
// CTFontStylisticClass - The stylistic class values of the font.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontStylisticClass
type CTFontStylisticClass uint

const (
	// kCTFontClarendonSerifsClass - The font’s style is a variation of the Oldstyle Serifs and the Transitional Serifs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontStylisticClass/clarendonSerifsClass
	kCTFontClarendonSerifsClass CTFontStylisticClass = 0
	// kCTFontClassClarendonSerifs - A font style variation of the Oldstyle Serifs and the Transitional Serifs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontStylisticClass/classClarendonSerifs
	kCTFontClassClarendonSerifs CTFontStylisticClass = 0
	// kCTFontClassFreeformSerifs - A font style that includes serifs but expresses a design freedom that doesn’t generally fit within the other serif design classifications.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontStylisticClass/classFreeformSerifs
	kCTFontClassFreeformSerifs CTFontStylisticClass = 0
	// kCTFontClassModernSerifs - A font style based on the Latin printing style of the 20th century.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontStylisticClass/classModernSerifs
	kCTFontClassModernSerifs CTFontStylisticClass = 0
	// kCTFontClassOldStyleSerifs - A font style based on the Latin printing style of the 15th to 17th century.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontStylisticClass/classOldStyleSerifs
	kCTFontClassOldStyleSerifs CTFontStylisticClass = 0
	// kCTFontClassOrnamentals - A font style that includes highly decorated or stylized character shapes such as those typically used in headlines.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontStylisticClass/classOrnamentals
	kCTFontClassOrnamentals CTFontStylisticClass = 0
	// kCTFontClassSansSerif - A font style that includes most basic letter forms (excluding Scripts and Ornamentals) that do not have serifs on the strokes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontStylisticClass/classSansSerif
	kCTFontClassSansSerif CTFontStylisticClass = 0
	// kCTFontClassScripts - A font style among those typefaces designed to simulate handwriting.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontStylisticClass/classScripts
	kCTFontClassScripts CTFontStylisticClass = 0
	// kCTFontClassSlabSerifs - A font style characterized by serifs with a square transition between the strokes and the serifs (no brackets).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontStylisticClass/classSlabSerifs
	kCTFontClassSlabSerifs CTFontStylisticClass = 0
	// kCTFontClassSymbolic - A generally design-independent font style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontStylisticClass/classSymbolic
	kCTFontClassSymbolic CTFontStylisticClass = 0
	// kCTFontClassTransitionalSerifs - A font style based on the Latin printing style of the 18th to 19th century.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontStylisticClass/classTransitionalSerifs
	kCTFontClassTransitionalSerifs CTFontStylisticClass = 0
	// kCTFontFreeformSerifsClass - The font’s style includes serifs but expresses a design freedom that doesn’t generally fit within the other serif design classifications.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontStylisticClass/freeformSerifsClass
	kCTFontFreeformSerifsClass CTFontStylisticClass = 0
	// kCTFontClassUnknown - The font has no design classification.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontStylisticClass/kCTFontClassUnknown
	kCTFontClassUnknown CTFontStylisticClass = 0
	// kCTFontUnknownClass - The font has no design classification.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontStylisticClass/kCTFontUnknownClass
	kCTFontUnknownClass CTFontStylisticClass = 0
	// kCTFontModernSerifsClass - The font’s style is based on the Latin printing style of the 20th century.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontStylisticClass/modernSerifsClass
	kCTFontModernSerifsClass CTFontStylisticClass = 0
	// kCTFontOldStyleSerifsClass - The font’s style is based on the Latin printing style of the 15th to 17th century.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontStylisticClass/oldStyleSerifsClass
	kCTFontOldStyleSerifsClass CTFontStylisticClass = 0
	// kCTFontOrnamentalsClass - The font’s style includes highly decorated or stylized character shapes such as those typically used in headlines.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontStylisticClass/ornamentalsClass
	kCTFontOrnamentalsClass CTFontStylisticClass = 0
	// kCTFontSansSerifClass - The font’s style includes most basic letter forms (excluding Scripts and Ornamentals) that do not have serifs on the strokes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontStylisticClass/sansSerifClass
	kCTFontSansSerifClass CTFontStylisticClass = 0
	// kCTFontScriptsClass - The font’s style is among those typefaces designed to simulate handwriting.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontStylisticClass/scriptsClass
	kCTFontScriptsClass CTFontStylisticClass = 0
	// kCTFontSlabSerifsClass - The font’s style is characterized by serifs with a square transition between the strokes and the serifs (no brackets).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontStylisticClass/slabSerifsClass
	kCTFontSlabSerifsClass CTFontStylisticClass = 0
	// kCTFontSymbolicClass - The font’s style is generally design independent.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontStylisticClass/symbolicClass
	kCTFontSymbolicClass CTFontStylisticClass = 0
	// kCTFontTransitionalSerifsClass - The font’s style is based on the Latin printing style of the 18th to 19th century.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontStylisticClass/transitionalSerifsClass
	kCTFontTransitionalSerifsClass CTFontStylisticClass = 0
)

/* debug [enums.gen.go]: Processing enum CTFontSymbolicTraits (20 cases) */
// CTFontSymbolicTraits - The symbolic representation of stylistic font attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontSymbolicTraits
type CTFontSymbolicTraits uint

const (
	// kCTFontBoldTrait - The font typestyle is boldface.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontSymbolicTraits/boldTrait
	kCTFontBoldTrait CTFontSymbolicTraits = 0
	// kCTFontClassMaskTrait - Mask for the font class.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontSymbolicTraits/classMaskTrait
	kCTFontClassMaskTrait CTFontSymbolicTraits = 0
	// kCTFontColorGlyphsTrait - The font contains color glyphs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontSymbolicTraits/colorGlyphsTrait
	kCTFontColorGlyphsTrait CTFontSymbolicTraits = 0
	// kCTFontCompositeTrait - The font is in Composite Font Reference format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontSymbolicTraits/compositeTrait
	kCTFontCompositeTrait CTFontSymbolicTraits = 0
	// kCTFontCondensedTrait - The font typestyle is condensed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontSymbolicTraits/condensedTrait
	kCTFontCondensedTrait CTFontSymbolicTraits = 0
	// kCTFontExpandedTrait - The font typestyle is expanded.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontSymbolicTraits/expandedTrait
	kCTFontExpandedTrait CTFontSymbolicTraits = 0
	// kCTFontItalicTrait - The font typestyle is italic.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontSymbolicTraits/italicTrait
	kCTFontItalicTrait CTFontSymbolicTraits = 0
	// kCTFontMonoSpaceTrait - The font uses fixed-pitch glyphs if available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontSymbolicTraits/monoSpaceTrait
	kCTFontMonoSpaceTrait CTFontSymbolicTraits = 0
	// kCTFontTraitBold - The font typestyle is boldface.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontSymbolicTraits/traitBold
	kCTFontTraitBold CTFontSymbolicTraits = 0
	// kCTFontTraitClassMask - Mask for the font class.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontSymbolicTraits/traitClassMask
	kCTFontTraitClassMask CTFontSymbolicTraits = 0
	// kCTFontTraitColorGlyphs - The font contains color glyphs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontSymbolicTraits/traitColorGlyphs
	kCTFontTraitColorGlyphs CTFontSymbolicTraits = 0
	// kCTFontTraitComposite - The font is in Composite Font Reference format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontSymbolicTraits/traitComposite
	kCTFontTraitComposite CTFontSymbolicTraits = 0
	// kCTFontTraitCondensed - The font typestyle is condensed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontSymbolicTraits/traitCondensed
	kCTFontTraitCondensed CTFontSymbolicTraits = 0
	// kCTFontTraitExpanded - The font typestyle is expanded.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontSymbolicTraits/traitExpanded
	kCTFontTraitExpanded CTFontSymbolicTraits = 0
	// kCTFontTraitItalic - The font typestyle is italic.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontSymbolicTraits/traitItalic
	kCTFontTraitItalic CTFontSymbolicTraits = 0
	// kCTFontTraitMonoSpace - The font uses fixed-pitch glyphs if available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontSymbolicTraits/traitMonoSpace
	kCTFontTraitMonoSpace CTFontSymbolicTraits = 0
	// kCTFontTraitUIOptimized - The font synthesizes appropriate attributes for user interface rendering, such as control titles, if necessary.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontSymbolicTraits/traitUIOptimized
	kCTFontTraitUIOptimized CTFontSymbolicTraits = 0
	// kCTFontTraitVertical - The font uses vertical glyph variants and metrics.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontSymbolicTraits/traitVertical
	kCTFontTraitVertical CTFontSymbolicTraits = 0
	// kCTFontUIOptimizedTrait - The font synthesizes appropriate attributes for user interface rendering, such as control titles, if necessary.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontSymbolicTraits/uiOptimizedTrait
	kCTFontUIOptimizedTrait CTFontSymbolicTraits = 0
	// kCTFontVerticalTrait - The font uses vertical glyph variants and metrics.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontSymbolicTraits/verticalTrait
	kCTFontVerticalTrait CTFontSymbolicTraits = 0
)

/* debug [enums.gen.go]: Processing enum CTFontTableOptions (2 cases) */
// CTFontTableOptions - Constants that describe font table options.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontTableOptions
type CTFontTableOptions uint

const (
	// kCTFontTableOptionExcludeSynthetic - The font table excludes synthetic font data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontTableOptions/excludeSynthetic
	kCTFontTableOptionExcludeSynthetic CTFontTableOptions = 0
	// kCTFontTableOptionNoOptions - No font table options are specified.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontTableOptions/kCTFontTableOptionNoOptions
	kCTFontTableOptionNoOptions CTFontTableOptions = 0
)

/* debug [enums.gen.go]: Processing enum CTFontUIFontType (56 cases) */
// CTFontUIFontType - Constants that represent the specific user-interface purpose to specify for font creation.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType
type CTFontUIFontType uint

const (
	// kCTFontUIFontAlertHeader - The font for alert headers.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/alertHeader
	kCTFontUIFontAlertHeader CTFontUIFontType = 0
	// kCTFontUIFontApplication - The default font for text documents.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/application
	kCTFontUIFontApplication CTFontUIFontType = 0
	// kCTFontUIFontControlContent - The font for contents of user-interface controls.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/controlContent
	kCTFontUIFontControlContent CTFontUIFontType = 0
	// kCTFontUIFontEmphasizedSystem - The system font for emphasis in alerts.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/emphasizedSystem
	kCTFontUIFontEmphasizedSystem CTFontUIFontType = 0
	// kCTFontUIFontEmphasizedSystemDetail - The system font for emphasis in details.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/emphasizedSystemDetail
	kCTFontUIFontEmphasizedSystemDetail CTFontUIFontType = 0
	// kCTFontAlertHeaderFontType - The font used for alert headers.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/kCTFontAlertHeaderFontType
	kCTFontAlertHeaderFontType CTFontUIFontType = 0
	// kCTFontApplicationFontType - The default font for text documents.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/kCTFontApplicationFontType
	kCTFontApplicationFontType CTFontUIFontType = 0
	// kCTFontControlContentFontType - The font used for contents of user-interface controls.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/kCTFontControlContentFontType
	kCTFontControlContentFontType CTFontUIFontType = 0
	// kCTFontEmphasizedSystemDetailFontType - The system font used for emphasis in details.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/kCTFontEmphasizedSystemDetailFontType
	kCTFontEmphasizedSystemDetailFontType CTFontUIFontType = 0
	// kCTFontEmphasizedSystemFontType - The system font used for emphasis in alerts.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/kCTFontEmphasizedSystemFontType
	kCTFontEmphasizedSystemFontType CTFontUIFontType = 0
	// kCTFontLabelFontType - The font used for labels and tick marks on full-size sliders.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/kCTFontLabelFontType
	kCTFontLabelFontType CTFontUIFontType = 0
	// kCTFontMenuItemCmdKeyFontType - The font used for menu-item command-key equivalents.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/kCTFontMenuItemCmdKeyFontType
	kCTFontMenuItemCmdKeyFontType CTFontUIFontType = 0
	// kCTFontMenuItemFontType - The font used for menu items.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/kCTFontMenuItemFontType
	kCTFontMenuItemFontType CTFontUIFontType = 0
	// kCTFontMenuItemMarkFontType - The font used to draw menu-item marks.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/kCTFontMenuItemMarkFontType
	kCTFontMenuItemMarkFontType CTFontUIFontType = 0
	// kCTFontMenuTitleFontType - The font used for menu titles.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/kCTFontMenuTitleFontType
	kCTFontMenuTitleFontType CTFontUIFontType = 0
	// kCTFontMessageFontType - The font used for standard interface items, such as button labels and menu items.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/kCTFontMessageFontType
	kCTFontMessageFontType CTFontUIFontType = 0
	// kCTFontMiniEmphasizedSystemFontType - The miniature system font used for emphasis.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/kCTFontMiniEmphasizedSystemFontType
	kCTFontMiniEmphasizedSystemFontType CTFontUIFontType = 0
	// kCTFontMiniSystemFontType - The standard miniature system font used for mini controls and utility window labels and text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/kCTFontMiniSystemFontType
	kCTFontMiniSystemFontType CTFontUIFontType = 0
	// kCTFontNoFontType - The user-interface font type isn’t specified.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/kCTFontNoFontType
	kCTFontNoFontType CTFontUIFontType = 0
	// kCTFontPaletteFontType - The font used in tool palettes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/kCTFontPaletteFontType
	kCTFontPaletteFontType CTFontUIFontType = 0
	// kCTFontPushButtonFontType - The font used for a push button, a rounded rectangular button with a text label on it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/kCTFontPushButtonFontType
	kCTFontPushButtonFontType CTFontUIFontType = 0
	// kCTFontSmallEmphasizedSystemFontType - The small system font used for emphasis.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/kCTFontSmallEmphasizedSystemFontType
	kCTFontSmallEmphasizedSystemFontType CTFontUIFontType = 0
	// kCTFontSmallSystemFontType - The standard small system font used for informative text in alerts, column headings in lists, help tags, and small controls.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/kCTFontSmallSystemFontType
	kCTFontSmallSystemFontType CTFontUIFontType = 0
	// kCTFontSmallToolbarFontType - The small font used for labels of toolbar items.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/kCTFontSmallToolbarFontType
	kCTFontSmallToolbarFontType CTFontUIFontType = 0
	// kCTFontSystemDetailFontType - The standard system font used for details.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/kCTFontSystemDetailFontType
	kCTFontSystemDetailFontType CTFontUIFontType = 0
	// kCTFontSystemFontType - The system font used for standard user-interface items, such as button labels and menu items.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/kCTFontSystemFontType
	kCTFontSystemFontType CTFontUIFontType = 0
	// kCTFontToolbarFontType - The font used for labels of toolbar items.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/kCTFontToolbarFontType
	kCTFontToolbarFontType CTFontUIFontType = 0
	// kCTFontToolTipFontType - The font used for tool tips.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/kCTFontToolTipFontType
	kCTFontToolTipFontType CTFontUIFontType = 0
	// kCTFontUserFixedPitchFontType - The font used by default for documents and other text under the user’s control when that font is fixed-pitch.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/kCTFontUserFixedPitchFontType
	kCTFontUserFixedPitchFontType CTFontUIFontType = 0
	// kCTFontUserFontType - The font used by default for documents and other text under the user’s control.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/kCTFontUserFontType
	kCTFontUserFontType CTFontUIFontType = 0
	// kCTFontUtilityWindowTitleFontType - The font used for utility window titles.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/kCTFontUtilityWindowTitleFontType
	kCTFontUtilityWindowTitleFontType CTFontUIFontType = 0
	// kCTFontViewsFontType - The view font used as the default font of text in lists and tables.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/kCTFontViewsFontType
	kCTFontViewsFontType CTFontUIFontType = 0
	// kCTFontWindowTitleFontType - The font used for window titles.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/kCTFontWindowTitleFontType
	kCTFontWindowTitleFontType CTFontUIFontType = 0
	// kCTFontUIFontLabel - The font for labels and tick marks on full-size sliders.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/label
	kCTFontUIFontLabel CTFontUIFontType = 0
	// kCTFontUIFontMenuItem - The font for menu items.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/menuItem
	kCTFontUIFontMenuItem CTFontUIFontType = 0
	// kCTFontUIFontMenuItemCmdKey - The font for menu-item command-key equivalents.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/menuItemCmdKey
	kCTFontUIFontMenuItemCmdKey CTFontUIFontType = 0
	// kCTFontUIFontMenuItemMark - The font to draw menu-item marks.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/menuItemMark
	kCTFontUIFontMenuItemMark CTFontUIFontType = 0
	// kCTFontUIFontMenuTitle - The font for menu titles.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/menuTitle
	kCTFontUIFontMenuTitle CTFontUIFontType = 0
	// kCTFontUIFontMessage - The font for standard interface items, such as button labels and menu items.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/message
	kCTFontUIFontMessage CTFontUIFontType = 0
	// kCTFontUIFontMiniEmphasizedSystem - The miniature system font for emphasis.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/miniEmphasizedSystem
	kCTFontUIFontMiniEmphasizedSystem CTFontUIFontType = 0
	// kCTFontUIFontMiniSystem - The standard miniature system font for mini controls and utility window labels and text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/miniSystem
	kCTFontUIFontMiniSystem CTFontUIFontType = 0
	// kCTFontUIFontNone - The user-interface font type isn’t specified.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/none
	kCTFontUIFontNone CTFontUIFontType = 0
	// kCTFontUIFontPalette - The font in tool palettes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/palette
	kCTFontUIFontPalette CTFontUIFontType = 0
	// kCTFontUIFontPushButton - The font for a push button, a rounded rectangular button with a text label on it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/pushButton
	kCTFontUIFontPushButton CTFontUIFontType = 0
	// kCTFontUIFontSmallEmphasizedSystem - The small system font for emphasis.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/smallEmphasizedSystem
	kCTFontUIFontSmallEmphasizedSystem CTFontUIFontType = 0
	// kCTFontUIFontSmallSystem - The standard small system font for informative text in alerts, column headings in lists, help tags, and small controls.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/smallSystem
	kCTFontUIFontSmallSystem CTFontUIFontType = 0
	// kCTFontUIFontSmallToolbar - The small font for labels of toolbar items.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/smallToolbar
	kCTFontUIFontSmallToolbar CTFontUIFontType = 0
	// kCTFontUIFontSystem - The system font for standard user-interface items, such as button labels and menu items.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/system
	kCTFontUIFontSystem CTFontUIFontType = 0
	// kCTFontUIFontSystemDetail - The standard system font for details.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/systemDetail
	kCTFontUIFontSystemDetail CTFontUIFontType = 0
	// kCTFontUIFontToolbar - The font used for labels of toolbar items.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/toolbar
	kCTFontUIFontToolbar CTFontUIFontType = 0
	// kCTFontUIFontToolTip - The font for tool tips.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/toolTip
	kCTFontUIFontToolTip CTFontUIFontType = 0
	// kCTFontUIFontUser - The default font for documents and other text whose font the user can typically change.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/user
	kCTFontUIFontUser CTFontUIFontType = 0
	// kCTFontUIFontUserFixedPitch - The default font for documents and other text under the user’s control when that font is fixed-pitch.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/userFixedPitch
	kCTFontUIFontUserFixedPitch CTFontUIFontType = 0
	// kCTFontUIFontUtilityWindowTitle - The font for utility window titles.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/utilityWindowTitle
	kCTFontUIFontUtilityWindowTitle CTFontUIFontType = 0
	// kCTFontUIFontViews - The default view font for text in lists and tables.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/views
	kCTFontUIFontViews CTFontUIFontType = 0
	// kCTFontUIFontWindowTitle - The font for window titles.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType/windowTitle
	kCTFontUIFontWindowTitle CTFontUIFontType = 0
)

/* debug [enums.gen.go]: Processing enum CTFramePathFillRule (2 cases) */
// CTFramePathFillRule - These constants specify the fill rule used by a frame
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFramePathFillRule
type CTFramePathFillRule uint

const (
	// kCTFramePathFillEvenOdd - Paints the area using the even-odd fill rule.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFramePathFillRule/evenOdd
	kCTFramePathFillEvenOdd CTFramePathFillRule = 0
	// kCTFramePathFillWindingNumber - Paints the area using the nonzero winding number rule.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFramePathFillRule/windingNumber
	kCTFramePathFillWindingNumber CTFramePathFillRule = 0
)

/* debug [enums.gen.go]: Processing enum CTFrameProgression (3 cases) */
// CTFrameProgression - Constants that specify frame progression types.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFrameProgression
type CTFrameProgression uint

const (
	// kCTFrameProgressionLeftToRight - Lines stack left to right for vertical text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFrameProgression/leftToRight
	kCTFrameProgressionLeftToRight CTFrameProgression = 0
	// kCTFrameProgressionRightToLeft - Lines stack right to left for vertical text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFrameProgression/rightToLeft
	kCTFrameProgressionRightToLeft CTFrameProgression = 0
	// kCTFrameProgressionTopToBottom - Lines stack top to bottom for horizontal text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFrameProgression/topToBottom
	kCTFrameProgressionTopToBottom CTFrameProgression = 0
)

/* debug [enums.gen.go]: Processing enum CTLineBoundsOptions (6 cases) */
// CTLineBoundsOptions - Options for getting the bounds of a line of text.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineBoundsOptions
type CTLineBoundsOptions uint

const (
	// kCTLineBoundsExcludeTypographicLeading - An option to exclude typographic leading.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineBoundsOptions/excludeTypographicLeading
	kCTLineBoundsExcludeTypographicLeading CTLineBoundsOptions = 0
	// kCTLineBoundsExcludeTypographicShifts - An option to ignore cross-stream shifts due to positioning, such as kerning or baseline alignment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineBoundsOptions/excludeTypographicShifts
	kCTLineBoundsExcludeTypographicShifts CTLineBoundsOptions = 0
	// kCTLineBoundsIncludeLanguageExtents - An option to include additional space based on common glyph sequences for various languages.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineBoundsOptions/includeLanguageExtents
	kCTLineBoundsIncludeLanguageExtents CTLineBoundsOptions = 0
	// kCTLineBoundsUseGlyphPathBounds - An option to use glyph path bounds rather than the default typographic bounds.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineBoundsOptions/useGlyphPathBounds
	kCTLineBoundsUseGlyphPathBounds CTLineBoundsOptions = 0
	// kCTLineBoundsUseHangingPunctuation - An option to enable hanging punctuation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineBoundsOptions/useHangingPunctuation
	kCTLineBoundsUseHangingPunctuation CTLineBoundsOptions = 0
	// kCTLineBoundsUseOpticalBounds - An option to use optical bounds.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineBoundsOptions/useOpticalBounds
	kCTLineBoundsUseOpticalBounds CTLineBoundsOptions = 0
)

/* debug [enums.gen.go]: Processing enum CTLineBreakMode (6 cases) */
// CTLineBreakMode - These constants specify what happens when a line is too long for its frame.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineBreakMode
type CTLineBreakMode uint

const (
	// kCTLineBreakByCharWrapping - Wrapping occurs before the first character that doesn’t fit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineBreakMode/byCharWrapping
	kCTLineBreakByCharWrapping CTLineBreakMode = 0
	// kCTLineBreakByClipping - Lines are simply not drawn past the edge of the frame.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineBreakMode/byClipping
	kCTLineBreakByClipping CTLineBreakMode = 0
	// kCTLineBreakByTruncatingHead - Each line is displayed so that the end fits in the frame and the missing text is indicated by an ellipsis glyph.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineBreakMode/byTruncatingHead
	kCTLineBreakByTruncatingHead CTLineBreakMode = 0
	// kCTLineBreakByTruncatingMiddle - Each line is displayed so that the beginning and end fit in the container and the missing text is indicated by an ellipsis glyph in the middle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineBreakMode/byTruncatingMiddle
	kCTLineBreakByTruncatingMiddle CTLineBreakMode = 0
	// kCTLineBreakByTruncatingTail - Each line is displayed so that the beginning fits in the container and the missing text is indicated by an ellipsis glyph.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineBreakMode/byTruncatingTail
	kCTLineBreakByTruncatingTail CTLineBreakMode = 0
	// kCTLineBreakByWordWrapping - Wrapping occurs at word boundaries unless the word itself doesn’t fit on a single line.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineBreakMode/byWordWrapping
	kCTLineBreakByWordWrapping CTLineBreakMode = 0
)

/* debug [enums.gen.go]: Processing enum CTLineTruncationType (3 cases) */
// CTLineTruncationType - Truncation types required by the 
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineTruncationType
type CTLineTruncationType uint

const (
	// kCTLineTruncationEnd - Truncate the end of the line, leaving the start portion visible.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineTruncationType/end
	kCTLineTruncationEnd CTLineTruncationType = 0
	// kCTLineTruncationMiddle - Truncate the middle of the line, leaving both the start and the end portions visible.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineTruncationType/middle
	kCTLineTruncationMiddle CTLineTruncationType = 0
	// kCTLineTruncationStart - Truncate the beginning of the line, leaving the end portion visible.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineTruncationType/start
	kCTLineTruncationStart CTLineTruncationType = 0
)

/* debug [enums.gen.go]: Processing enum CTParagraphStyleSpecifier (19 cases) */
// CTParagraphStyleSpecifier - Constants used to query and modify a paragraph style object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier
type CTParagraphStyleSpecifier uint

const (
	// kCTParagraphStyleSpecifierAlignment - The text alignment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier/alignment
	kCTParagraphStyleSpecifierAlignment CTParagraphStyleSpecifier = 0
	// kCTParagraphStyleSpecifierBaseWritingDirection - The base writing direction of the lines.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier/baseWritingDirection
	kCTParagraphStyleSpecifierBaseWritingDirection CTParagraphStyleSpecifier = 0
	// kCTParagraphStyleSpecifierCount - The number of style specifiers.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier/count
	kCTParagraphStyleSpecifierCount CTParagraphStyleSpecifier = 0
	// kCTParagraphStyleSpecifierDefaultTabInterval - The document-wide default tab interval.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier/defaultTabInterval
	kCTParagraphStyleSpecifierDefaultTabInterval CTParagraphStyleSpecifier = 0
	// kCTParagraphStyleSpecifierFirstLineHeadIndent - The distance, in points, from the leading margin of a frame to the beginning of the paragraph’s first line.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier/firstLineHeadIndent
	kCTParagraphStyleSpecifierFirstLineHeadIndent CTParagraphStyleSpecifier = 0
	// kCTParagraphStyleSpecifierHeadIndent - The distance, in points, from the leading margin of a text container to the beginning of lines other than the first.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier/headIndent
	kCTParagraphStyleSpecifierHeadIndent CTParagraphStyleSpecifier = 0
	// kCTParagraphStyleSpecifierLineBoundsOptions - Options that control the alignment of the line edges with the leading and trailing margins.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier/lineBoundsOptions
	kCTParagraphStyleSpecifierLineBoundsOptions CTParagraphStyleSpecifier = 0
	// kCTParagraphStyleSpecifierLineBreakMode - The mode that should be used to break lines when laying out the paragraph’s text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier/lineBreakMode
	kCTParagraphStyleSpecifierLineBreakMode CTParagraphStyleSpecifier = 0
	// kCTParagraphStyleSpecifierLineHeightMultiple - The line height multiple.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier/lineHeightMultiple
	kCTParagraphStyleSpecifierLineHeightMultiple CTParagraphStyleSpecifier = 0
	// kCTParagraphStyleSpecifierLineSpacing - The space in points added between lines within the paragraph (commonly known as leading).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier/lineSpacing
	kCTParagraphStyleSpecifierLineSpacing CTParagraphStyleSpecifier = 0
	// kCTParagraphStyleSpecifierLineSpacingAdjustment - The space in points added between lines within the paragraph (commonly known as leading).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier/lineSpacingAdjustment
	kCTParagraphStyleSpecifierLineSpacingAdjustment CTParagraphStyleSpecifier = 0
	// kCTParagraphStyleSpecifierMaximumLineHeight - The maximum height that any line in the frame will occupy, regardless of the font size or size of any attached graphic.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier/maximumLineHeight
	kCTParagraphStyleSpecifierMaximumLineHeight CTParagraphStyleSpecifier = 0
	// kCTParagraphStyleSpecifierMaximumLineSpacing - The maximum space in points between lines within the paragraph (commonly known as leading).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier/maximumLineSpacing
	kCTParagraphStyleSpecifierMaximumLineSpacing CTParagraphStyleSpecifier = 0
	// kCTParagraphStyleSpecifierMinimumLineHeight - The minimum height that any line in the frame will occupy, regardless of the font size or size of any attached graphic.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier/minimumLineHeight
	kCTParagraphStyleSpecifierMinimumLineHeight CTParagraphStyleSpecifier = 0
	// kCTParagraphStyleSpecifierMinimumLineSpacing - The minimum space in points between lines within the paragraph (commonly known as leading).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier/minimumLineSpacing
	kCTParagraphStyleSpecifierMinimumLineSpacing CTParagraphStyleSpecifier = 0
	// kCTParagraphStyleSpecifierParagraphSpacing - The space added at the end of the paragraph to separate it from the following paragraph.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier/paragraphSpacing
	kCTParagraphStyleSpecifierParagraphSpacing CTParagraphStyleSpecifier = 0
	// kCTParagraphStyleSpecifierParagraphSpacingBefore - The distance between the paragraph’s top and the beginning of its text content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier/paragraphSpacingBefore
	kCTParagraphStyleSpecifierParagraphSpacingBefore CTParagraphStyleSpecifier = 0
	// kCTParagraphStyleSpecifierTabStops - The text tab objects, sorted by location, that define the tab stops for the paragraph style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier/tabStops
	kCTParagraphStyleSpecifierTabStops CTParagraphStyleSpecifier = 0
	// kCTParagraphStyleSpecifierTailIndent - The distance, in points, from the margin of a frame to the end of lines.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier/tailIndent
	kCTParagraphStyleSpecifierTailIndent CTParagraphStyleSpecifier = 0
)

/* debug [enums.gen.go]: Processing enum CTRubyAlignment (8 cases) */
// CTRubyAlignment - Constants that specify how to align the ruby text and the base text relative to each other when they have different lengths.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyAlignment
type CTRubyAlignment uint

const (
	// kCTRubyAlignmentAuto - Core Text automatically determines the alignment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyAlignment/auto
	kCTRubyAlignmentAuto CTRubyAlignment = 0
	// kCTRubyAlignmentCenter - Centers the ruby text within the width of the base text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyAlignment/center
	kCTRubyAlignmentCenter CTRubyAlignment = 0
	// kCTRubyAlignmentDistributeLetter - Distributes the ruby text evenly over the width of the base text, aligning the first and last characters of the ruby text with the first and last characters of the base text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyAlignment/distributeLetter
	kCTRubyAlignmentDistributeLetter CTRubyAlignment = 0
	// kCTRubyAlignmentDistributeSpace - Distributes the ruby text evenly over the width of the base text, adding space before the first and after the last character.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyAlignment/distributeSpace
	kCTRubyAlignmentDistributeSpace CTRubyAlignment = 0
	// kCTRubyAlignmentEnd - Aligns the ruby text with the ending edge of the base text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyAlignment/end
	kCTRubyAlignmentEnd CTRubyAlignment = 0
	// kCTRubyAlignmentInvalid - The alignment is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyAlignment/invalid
	kCTRubyAlignmentInvalid CTRubyAlignment = 0
	// kCTRubyAlignmentLineEdge - Aligns the ruby text to an adjacent line edge.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyAlignment/lineEdge
	kCTRubyAlignmentLineEdge CTRubyAlignment = 0
	// kCTRubyAlignmentStart - Aligns the ruby text with the starting edge of the base text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyAlignment/start
	kCTRubyAlignmentStart CTRubyAlignment = 0
)

/* debug [enums.gen.go]: Processing enum CTRubyOverhang (5 cases) */
// CTRubyOverhang - Constants that specify whether, and on which side, ruby text can overhang adjacent text if it’s wider than the base text.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyOverhang
type CTRubyOverhang uint

const (
	// kCTRubyOverhangAuto - The ruby text can overhang adjacent text on both sides.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyOverhang/auto
	kCTRubyOverhangAuto CTRubyOverhang = 0
	// kCTRubyOverhangEnd - The ruby text can overhang the text that follows it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyOverhang/end
	kCTRubyOverhangEnd CTRubyOverhang = 0
	// kCTRubyOverhangInvalid - The overhang specification is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyOverhang/invalid
	kCTRubyOverhangInvalid CTRubyOverhang = 0
	// kCTRubyOverhangNone - The ruby text can’t overhang the preceding or following text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyOverhang/none
	kCTRubyOverhangNone CTRubyOverhang = 0
	// kCTRubyOverhangStart - The ruby text can overhang the text that precedes it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyOverhang/start
	kCTRubyOverhangStart CTRubyOverhang = 0
)

/* debug [enums.gen.go]: Processing enum CTRubyPosition (5 cases) */
// CTRubyPosition - Constants that specify the position of the ruby text relative to to the base text.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyPosition
type CTRubyPosition uint

const (
	// kCTRubyPositionAfter - The ruby text is positioned after the base text, appearing below horizontal text and to the left of vertical text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyPosition/after
	kCTRubyPositionAfter CTRubyPosition = 0
	// kCTRubyPositionBefore - The ruby text is positioned before the base text, appearing above horizontal text and to the right of vertical text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyPosition/before
	kCTRubyPositionBefore CTRubyPosition = 0
	// kCTRubyPositionCount - A constant that accounts for all ruby positions during ruby annotation creation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyPosition/count
	kCTRubyPositionCount CTRubyPosition = 0
	// kCTRubyPositionInline - The ruby text follows the base text with no special styling.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyPosition/inline
	kCTRubyPositionInline CTRubyPosition = 0
	// kCTRubyPositionInterCharacter - The ruby text is positioned to the right of the base text, regardless of whether it’s horizontal or vertical.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyPosition/interCharacter
	kCTRubyPositionInterCharacter CTRubyPosition = 0
)

/* debug [enums.gen.go]: Processing enum CTRunStatus (4 cases) */
// CTRunStatus - A bitfield that represents the disposition of the run.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunStatus
type CTRunStatus uint

const (
	// kCTRunStatusHasNonIdentityMatrix - The run requires a specific text matrix to be set in the current Core Graphics context for proper drawing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunStatus/hasNonIdentityMatrix
	kCTRunStatusHasNonIdentityMatrix CTRunStatus = 0
	// kCTRunStatusNoStatus - The run has no special attributes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunStatus/kCTRunStatusNoStatus
	kCTRunStatusNoStatus CTRunStatus = 0
	// kCTRunStatusNonMonotonic - The run isn’t in strictly increasing or decreasing order.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunStatus/nonMonotonic
	kCTRunStatusNonMonotonic CTRunStatus = 0
	// kCTRunStatusRightToLeft - The run proceeds from right to left.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunStatus/rightToLeft
	kCTRunStatusRightToLeft CTRunStatus = 0
)

/* debug [enums.gen.go]: Processing enum CTTextAlignment (10 cases) */
// CTTextAlignment - Constants that specify text alignment.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTextAlignment
type CTTextAlignment uint

const (
	// kCTTextAlignmentCenter - Text is visually center-aligned.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTextAlignment/center
	kCTTextAlignmentCenter CTTextAlignment = 0
	// kCTTextAlignmentJustified - Text is fully justified.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTextAlignment/justified
	kCTTextAlignmentJustified CTTextAlignment = 0
	// kCTCenterTextAlignment - Text is visually center-aligned.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTextAlignment/kCTCenterTextAlignment
	kCTCenterTextAlignment CTTextAlignment = 0
	// kCTJustifiedTextAlignment - Text is fully justified.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTextAlignment/kCTJustifiedTextAlignment
	kCTJustifiedTextAlignment CTTextAlignment = 0
	// kCTLeftTextAlignment - Text is visually left-aligned.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTextAlignment/kCTLeftTextAlignment
	kCTLeftTextAlignment CTTextAlignment = 0
	// kCTNaturalTextAlignment - Text uses the natural alignment of the text’s script.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTextAlignment/kCTNaturalTextAlignment
	kCTNaturalTextAlignment CTTextAlignment = 0
	// kCTRightTextAlignment - Text is visually right-aligned.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTextAlignment/kCTRightTextAlignment
	kCTRightTextAlignment CTTextAlignment = 0
	// kCTTextAlignmentLeft - Text is visually left-aligned.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTextAlignment/left
	kCTTextAlignmentLeft CTTextAlignment = 0
	// kCTTextAlignmentNatural - Text uses the natural alignment of the text’s script.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTextAlignment/natural
	kCTTextAlignmentNatural CTTextAlignment = 0
	// kCTTextAlignmentRight - Text is visually right-aligned.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTextAlignment/right
	kCTTextAlignmentRight CTTextAlignment = 0
)

/* debug [enums.gen.go]: Processing enum CTUnderlineStyle (4 cases) */
// CTUnderlineStyle - Underline style specifiers.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTUnderlineStyle
type CTUnderlineStyle uint

const (
	// kCTUnderlineStyleDouble - A specifier that indicates to draw an underline consisting of a double line.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTUnderlineStyle/double
	kCTUnderlineStyleDouble CTUnderlineStyle = 0
	// kCTUnderlineStyleNone - A specifier that indicates not to draw an underline.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTUnderlineStyle/kCTUnderlineStyleNone
	kCTUnderlineStyleNone CTUnderlineStyle = 0
	// kCTUnderlineStyleSingle - A specifier that indicates to draw an underline consisting of a single line.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTUnderlineStyle/single
	kCTUnderlineStyleSingle CTUnderlineStyle = 0
	// kCTUnderlineStyleThick - A specifier that indicates to draw an underline consisting of a thick line.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTUnderlineStyle/thick
	kCTUnderlineStyleThick CTUnderlineStyle = 0
)

/* debug [enums.gen.go]: Processing enum CTUnderlineStyleModifiers (5 cases) */
// CTUnderlineStyleModifiers - Underline style modifiers.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTUnderlineStyleModifiers
type CTUnderlineStyleModifiers uint

const (
	// kCTUnderlinePatternDash - A modifier that indicates to draw an underline using a pattern of dashes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTUnderlineStyleModifiers/patternDash
	kCTUnderlinePatternDash CTUnderlineStyleModifiers = 0
	// kCTUnderlinePatternDashDot - A modifier that indicates to draw an underline using a pattern of alternating dashes and dots.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTUnderlineStyleModifiers/patternDashDot
	kCTUnderlinePatternDashDot CTUnderlineStyleModifiers = 0
	// kCTUnderlinePatternDashDotDot - A modifier that indicates to draw an underline using a pattern of a dash followed by two dots.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTUnderlineStyleModifiers/patternDashDotDot
	kCTUnderlinePatternDashDotDot CTUnderlineStyleModifiers = 0
	// kCTUnderlinePatternDot - A modifier that indicates to draw an underline using a pattern of dots.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTUnderlineStyleModifiers/patternDot
	kCTUnderlinePatternDot CTUnderlineStyleModifiers = 0
	// kCTUnderlinePatternSolid - A modifier that indicates to draw a solid underline.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTUnderlineStyleModifiers/patternSolid
	kCTUnderlinePatternSolid CTUnderlineStyleModifiers = 0
)

/* debug [enums.gen.go]: Processing enum CTWritingDirection (3 cases) */
// CTWritingDirection - These constants specify the writing direction.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTWritingDirection
type CTWritingDirection uint

const (
	// kCTWritingDirectionLeftToRight - The writing direction is left to right.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTWritingDirection/leftToRight
	kCTWritingDirectionLeftToRight CTWritingDirection = 0
	// kCTWritingDirectionNatural - The writing direction is algorithmically determined using the Unicode Bidirectional Algorithm rules P2 and P3.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTWritingDirection/natural
	kCTWritingDirectionNatural CTWritingDirection = 0
	// kCTWritingDirectionRightToLeft - The writing direction is right to left.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTWritingDirection/rightToLeft
	kCTWritingDirectionRightToLeft CTWritingDirection = 0
)


