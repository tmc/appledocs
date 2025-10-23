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

// CTLineBreakMode - These constants specify what happens when a line is too long for its frame.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineBreakMode
type CTLineBreakMode uint

// CTLineTruncationType - Truncation types required by the 
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineTruncationType
type CTLineTruncationType uint

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

// CTUnderlineStyleModifiers - Underline style modifiers.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTUnderlineStyleModifiers
type CTUnderlineStyleModifiers uint

// CTWritingDirection - These constants specify the writing direction.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTWritingDirection
type CTWritingDirection uint


