// Code generated from Apple documentation for CoreText. DO NOT EDIT.

package coretext

// Enum types and constants
// CTCharacterCollection - Constants that specify character collections.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTCharacterCollection
type CTCharacterCollection uint

// CTFontCollectionCopyOptions - Option bits for use with CTFontCollectionCopyFontAttribute(s).
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCollectionCopyOptions
type CTFontCollectionCopyOptions uint

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

// CTFontFormat - The recognized format of the font.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontFormat
type CTFontFormat uint

// CTFontManagerAutoActivationSetting - Sets the auto-activation for the specified bundle identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerAutoActivationSetting
type CTFontManagerAutoActivationSetting uint

// CTFontManagerError - Errors that prevent unregistration of fonts for a specified font file URL.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerError
type CTFontManagerError uint

// CTFontManagerScope - Constants that define the scope for font registration.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerScope
type CTFontManagerScope uint

// CTFontOptions - Options for font creation and descriptor matching.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontOptions
type CTFontOptions uint

// CTFontOrientation - The intended rendering orientation of the font for obtaining glyph metrics.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontOrientation
type CTFontOrientation uint

// CTFontStylisticClass - The stylistic class values of the font.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontStylisticClass
type CTFontStylisticClass uint

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

// CTFontTableOptions - Constants that describe font table options.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontTableOptions
type CTFontTableOptions uint

// CTFontUIFontType - Constants that represent the specific user-interface purpose to specify for font creation.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontUIFontType
type CTFontUIFontType uint

// CTFramePathFillRule - These constants specify the fill rule used by a frame
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFramePathFillRule
type CTFramePathFillRule uint

// CTFrameProgression - Constants that specify frame progression types.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFrameProgression
type CTFrameProgression uint

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

// CTLineBreakMode - These constants specify what happens when a line is too long for its frame.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineBreakMode
type CTLineBreakMode uint

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

// CTParagraphStyleSpecifier - Constants used to query and modify a paragraph style object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier
type CTParagraphStyleSpecifier uint

// CTRubyAlignment - Constants that specify how to align the ruby text and the base text relative to each other when they have different lengths.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyAlignment
type CTRubyAlignment uint

// CTRubyOverhang - Constants that specify whether, and on which side, ruby text can overhang adjacent text if it’s wider than the base text.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyOverhang
type CTRubyOverhang uint

// CTRubyPosition - Constants that specify the position of the ruby text relative to to the base text.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyPosition
type CTRubyPosition uint

// CTRunStatus - A bitfield that represents the disposition of the run.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunStatus
type CTRunStatus uint

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

// CTUnderlineStyleModifiers - Underline style modifiers.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTUnderlineStyleModifiers
type CTUnderlineStyleModifiers uint

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


