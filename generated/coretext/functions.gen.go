// Code generated from Apple documentation for CoreText. DO NOT EDIT.

package coretext

/* debug [functions.gen.go]: Generating 203 functions for CoreText */
import (
	"unsafe"

	"github.com/ebitengine/purego"
	corefoundation "github.com/tmc/appledocs/generated/corefoundation"
)


// CoreText Functions (203 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_CTFontCollectionCopyExclusionDescriptors func(FontCollectionRef) ArrayRef
	_CTFontCollectionCopyFontAttribute func(FontCollectionRef, StringRef, FontCollectionCopyOptions) ArrayRef
	_CTFontCollectionCopyFontAttributes func(FontCollectionRef, SetRef, FontCollectionCopyOptions) ArrayRef
	_CTFontCollectionCopyQueryDescriptors func(FontCollectionRef) ArrayRef
	_CTFontCollectionCreateCopyWithFontDescriptors func(FontCollectionRef, ArrayRef, DictionaryRef) FontCollectionRef
	_CTFontCollectionCreateFromAvailableFonts func(DictionaryRef) FontCollectionRef
	_CTFontCollectionCreateMatchingFontDescriptors func(FontCollectionRef) ArrayRef
	_CTFontCollectionCreateMatchingFontDescriptorsForFamily func(FontCollectionRef, StringRef, DictionaryRef) ArrayRef
	_CTFontCollectionCreateMatchingFontDescriptorsSortedWithCallback func(FontCollectionRef, FontCollectionSortDescriptorsCallback, unsafe.Pointer) ArrayRef
	_CTFontCollectionCreateMatchingFontDescriptorsWithOptions func(FontCollectionRef, DictionaryRef) ArrayRef
	_CTFontCollectionCreateMutableCopy func(FontCollectionRef) MutableFontCollectionRef
	_CTFontCollectionCreateWithFontDescriptors func(ArrayRef, DictionaryRef) FontCollectionRef
	_CTFontCollectionGetTypeID func() TypeID
	_CTFontCollectionSetExclusionDescriptors func(MutableFontCollectionRef, ArrayRef)
	_CTFontCollectionSetQueryDescriptors func(MutableFontCollectionRef, ArrayRef)
	_CTFontCopyAttribute func(FontRef, StringRef) TypeRef
	_CTFontCopyAvailableTables func(FontRef, FontTableOptions) ArrayRef
	_CTFontCopyCharacterSet func(FontRef) CharacterSetRef
	_CTFontCopyDefaultCascadeListForLanguages func(FontRef, ArrayRef) ArrayRef
	_CTFontCopyDisplayName func(FontRef) StringRef
	_CTFontCopyFamilyName func(FontRef) StringRef
	_CTFontCopyFeatures func(FontRef) ArrayRef
	_CTFontCopyFeatureSettings func(FontRef) ArrayRef
	_CTFontCopyFontDescriptor func(FontRef) FontDescriptorRef
	_CTFontCopyFullName func(FontRef) StringRef
	_CTFontCopyGraphicsFont func(FontRef, unsafe.Pointer) FontRef
	_CTFontCopyLocalizedName func(FontRef, StringRef, unsafe.Pointer) StringRef
	_CTFontCopyName func(FontRef, StringRef) StringRef
	_CTFontCopyNameForGlyph func(FontRef, Glyph) StringRef
	_CTFontCopyPostScriptName func(FontRef) StringRef
	_CTFontCopySupportedLanguages func(FontRef) ArrayRef
	_CTFontCopyTable func(FontRef, FontTableTag, FontTableOptions) DataRef
	_CTFontCopyTraits func(FontRef) DictionaryRef
	_CTFontCopyVariation func(FontRef) DictionaryRef
	_CTFontCopyVariationAxes func(FontRef) ArrayRef
	_CTFontCreateCopyWithAttributes func(FontRef, float64, unsafe.Pointer, FontDescriptorRef) FontRef
	_CTFontCreateCopyWithFamily func(FontRef, float64, unsafe.Pointer, StringRef) FontRef
	_CTFontCreateCopyWithSymbolicTraits func(FontRef, float64, unsafe.Pointer, FontSymbolicTraits, FontSymbolicTraits) FontRef
	_CTFontCreateForString func(FontRef, StringRef, corefoundation.Range) FontRef
	_CTFontCreateForStringWithLanguage func(FontRef, StringRef, corefoundation.Range, StringRef) FontRef
	_CTFontCreatePathForGlyph func(FontRef, Glyph, unsafe.Pointer) PathRef
	_CTFontCreateUIFontForLanguage func(FontUIFontType, float64, StringRef) FontRef
	_CTFontCreateWithFontDescriptor func(FontDescriptorRef, float64, unsafe.Pointer) FontRef
	_CTFontCreateWithFontDescriptorAndOptions func(FontDescriptorRef, float64, unsafe.Pointer, FontOptions) FontRef
	_CTFontCreateWithGraphicsFont func(FontRef, float64, unsafe.Pointer, FontDescriptorRef) FontRef
	_CTFontCreateWithName func(StringRef, float64, unsafe.Pointer) FontRef
	_CTFontCreateWithNameAndOptions func(StringRef, float64, unsafe.Pointer, FontOptions) FontRef
	_CTFontCreateWithPlatformFont func(ATSFontRef, float64, unsafe.Pointer, FontDescriptorRef) FontRef
	_CTFontCreateWithQuickdrawInstance func(unsafe.Pointer, int16, uint8, float64) FontRef
	_CTFontDescriptorCopyAttribute func(FontDescriptorRef, StringRef) TypeRef
	_CTFontDescriptorCopyAttributes func(FontDescriptorRef) DictionaryRef
	_CTFontDescriptorCopyLocalizedAttribute func(FontDescriptorRef, StringRef, unsafe.Pointer) TypeRef
	_CTFontDescriptorCreateCopyWithAttributes func(FontDescriptorRef, DictionaryRef) FontDescriptorRef
	_CTFontDescriptorCreateCopyWithFamily func(FontDescriptorRef, StringRef) FontDescriptorRef
	_CTFontDescriptorCreateCopyWithFeature func(FontDescriptorRef, NumberRef, NumberRef) FontDescriptorRef
	_CTFontDescriptorCreateCopyWithSymbolicTraits func(FontDescriptorRef, FontSymbolicTraits, FontSymbolicTraits) FontDescriptorRef
	_CTFontDescriptorCreateCopyWithVariation func(FontDescriptorRef, NumberRef, float64) FontDescriptorRef
	_CTFontDescriptorCreateMatchingFontDescriptor func(FontDescriptorRef, SetRef) FontDescriptorRef
	_CTFontDescriptorCreateMatchingFontDescriptors func(FontDescriptorRef, SetRef) ArrayRef
	_CTFontDescriptorCreateWithAttributes func(DictionaryRef) FontDescriptorRef
	_CTFontDescriptorCreateWithNameAndSize func(StringRef, float64) FontDescriptorRef
	_CTFontDescriptorGetTypeID func() TypeID
	_CTFontDescriptorMatchFontDescriptorsWithProgressHandler func(ArrayRef, SetRef, FontDescriptorProgressHandler) bool
	_CTFontDrawGlyphs func(FontRef, unsafe.Pointer, unsafe.Pointer, uintptr, ContextRef)
	_CTFontDrawImageFromAdaptiveImageProviderAtPoint func(FontRef, unsafe.Pointer, corefoundation.CGPoint, ContextRef)
	_CTFontGetAdvancesForGlyphs func(FontRef, FontOrientation, unsafe.Pointer, corefoundation.CGSize, Index) float64
	_CTFontGetAscent func(FontRef) float64
	_CTFontGetBoundingBox func(FontRef) corefoundation.CGRect
	_CTFontGetBoundingRectsForGlyphs func(FontRef, FontOrientation, unsafe.Pointer, corefoundation.CGRect, Index) corefoundation.CGRect
	_CTFontGetCapHeight func(FontRef) float64
	_CTFontGetDescent func(FontRef) float64
	_CTFontGetGlyphCount func(FontRef) Index
	_CTFontGetGlyphsForCharacters func(FontRef, unsafe.Pointer, Glyph, Index) bool
	_CTFontGetGlyphWithName func(FontRef, StringRef) Glyph
	_CTFontGetLeading func(FontRef) float64
	_CTFontGetLigatureCaretPositions func(FontRef, Glyph, float64, Index) Index
	_CTFontGetMatrix func(FontRef) corefoundation.CGAffineTransform
	_CTFontGetOpticalBoundsForGlyphs func(FontRef, unsafe.Pointer, corefoundation.CGRect, Index, OptionFlags) corefoundation.CGRect
	_CTFontGetPlatformFont func(FontRef, unsafe.Pointer) ATSFontRef
	_CTFontGetSize func(FontRef) float64
	_CTFontGetSlantAngle func(FontRef) float64
	_CTFontGetStringEncoding func(FontRef) StringEncoding
	_CTFontGetSymbolicTraits func(FontRef) FontSymbolicTraits
	_CTFontGetTypeID func() TypeID
	_CTFontGetTypographicBoundsForAdaptiveImageProvider func(FontRef, unsafe.Pointer) corefoundation.CGRect
	_CTFontGetUnderlinePosition func(FontRef) float64
	_CTFontGetUnderlineThickness func(FontRef) float64
	_CTFontGetUnitsPerEm func(FontRef) unsafe.Pointer
	_CTFontGetVerticalTranslationsForGlyphs func(FontRef, unsafe.Pointer, corefoundation.CGSize, Index)
	_CTFontGetXHeight func(FontRef) float64
	_CTFontHasTable func(FontRef, FontTableTag) bool
	_CTFontManagerCompareFontFamilyNames func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) ComparisonResult
	_CTFontManagerCopyAvailableFontFamilyNames func() ArrayRef
	_CTFontManagerCopyAvailableFontURLs func() ArrayRef
	_CTFontManagerCopyAvailablePostScriptNames func() ArrayRef
	_CTFontManagerCopyRegisteredFontDescriptors func(FontManagerScope, bool) ArrayRef
	_CTFontManagerCreateFontDescriptorFromData func(DataRef) FontDescriptorRef
	_CTFontManagerCreateFontDescriptorsFromData func(DataRef) ArrayRef
	_CTFontManagerCreateFontDescriptorsFromURL func(URLRef) ArrayRef
	_CTFontManagerCreateFontRequestRunLoopSource func(Index, ArrayRef) RunLoopSourceRef
	_CTFontManagerEnableFontDescriptors func(ArrayRef, bool)
	_CTFontManagerGetAutoActivationSetting func(StringRef) FontManagerAutoActivationSetting
	_CTFontManagerGetScopeForURL func(URLRef) FontManagerScope
	_CTFontManagerIsSupportedFont func(URLRef) bool
	_CTFontManagerRegisterFontDescriptors func(ArrayRef, FontManagerScope, bool, bool)
	_CTFontManagerRegisterFontsForURL func(URLRef, FontManagerScope, unsafe.Pointer) bool
	_CTFontManagerRegisterFontsForURLs func(ArrayRef, FontManagerScope, unsafe.Pointer) bool
	_CTFontManagerRegisterFontsWithAssetNames func(ArrayRef, BundleRef, FontManagerScope, bool, bool)
	_CTFontManagerRegisterFontURLs func(ArrayRef, FontManagerScope, bool, bool)
	_CTFontManagerRegisterGraphicsFont func(FontRef, unsafe.Pointer) bool
	_CTFontManagerRequestFonts func(ArrayRef)
	_CTFontManagerSetAutoActivationSetting func(StringRef, FontManagerAutoActivationSetting)
	_CTFontManagerUnregisterFontDescriptors func(ArrayRef, FontManagerScope, bool)
	_CTFontManagerUnregisterFontsForURL func(URLRef, FontManagerScope, unsafe.Pointer) bool
	_CTFontManagerUnregisterFontsForURLs func(ArrayRef, FontManagerScope, unsafe.Pointer) bool
	_CTFontManagerUnregisterFontURLs func(ArrayRef, FontManagerScope, bool)
	_CTFontManagerUnregisterGraphicsFont func(FontRef, unsafe.Pointer) bool
	_CTFrameDraw func(FrameRef, ContextRef)
	_CTFrameGetFrameAttributes func(FrameRef) DictionaryRef
	_CTFrameGetLineOrigins func(FrameRef, corefoundation.Range, corefoundation.CGPoint)
	_CTFrameGetLines func(FrameRef) ArrayRef
	_CTFrameGetPath func(FrameRef) PathRef
	_CTFrameGetStringRange func(FrameRef) corefoundation.Range
	_CTFrameGetTypeID func() TypeID
	_CTFrameGetVisibleStringRange func(FrameRef) corefoundation.Range
	_CTFramesetterCreateFrame func(FramesetterRef, corefoundation.Range, PathRef, DictionaryRef) FrameRef
	_CTFramesetterCreateWithAttributedString func(AttributedStringRef) FramesetterRef
	_CTFramesetterCreateWithTypesetter func(TypesetterRef) FramesetterRef
	_CTFramesetterGetTypeID func() TypeID
	_CTFramesetterGetTypesetter func(FramesetterRef) TypesetterRef
	_CTFramesetterSuggestFrameSizeWithConstraints func(FramesetterRef, corefoundation.Range, DictionaryRef, corefoundation.CGSize, unsafe.Pointer) corefoundation.CGSize
	_CTGetCoreTextVersion func() uint32
	_CTGlyphInfoCreateWithCharacterIdentifier func(FontIndex, CharacterCollection, StringRef) GlyphInfoRef
	_CTGlyphInfoCreateWithGlyph func(Glyph, FontRef, StringRef) GlyphInfoRef
	_CTGlyphInfoCreateWithGlyphName func(StringRef, FontRef, StringRef) GlyphInfoRef
	_CTGlyphInfoGetCharacterCollection func(GlyphInfoRef) CharacterCollection
	_CTGlyphInfoGetCharacterIdentifier func(GlyphInfoRef) FontIndex
	_CTGlyphInfoGetGlyph func(GlyphInfoRef) Glyph
	_CTGlyphInfoGetGlyphName func(GlyphInfoRef) StringRef
	_CTGlyphInfoGetTypeID func() TypeID
	_CTLineCreateJustifiedLine func(LineRef, float64, float64) LineRef
	_CTLineCreateTruncatedLine func(LineRef, float64, LineTruncationType, LineRef) LineRef
	_CTLineCreateWithAttributedString func(AttributedStringRef) LineRef
	_CTLineDraw func(LineRef, ContextRef)
	_CTLineEnumerateCaretOffsets func(LineRef)
	_CTLineGetBoundsWithOptions func(LineRef, LineBoundsOptions) corefoundation.CGRect
	_CTLineGetGlyphCount func(LineRef) Index
	_CTLineGetGlyphRuns func(LineRef) ArrayRef
	_CTLineGetImageBounds func(LineRef, ContextRef) corefoundation.CGRect
	_CTLineGetOffsetForStringIndex func(LineRef, Index, []float64) float64
	_CTLineGetPenOffsetForFlush func(LineRef, float64, float64) float64
	_CTLineGetStringIndexForPosition func(LineRef, corefoundation.CGPoint) Index
	_CTLineGetStringRange func(LineRef) corefoundation.Range
	_CTLineGetTrailingWhitespaceWidth func(LineRef) float64
	_CTLineGetTypeID func() TypeID
	_CTLineGetTypographicBounds func(LineRef, []float64, []float64, []float64) float64
	_CTParagraphStyleCreate func(unsafe.Pointer, uintptr) ParagraphStyleRef
	_CTParagraphStyleCreateCopy func(ParagraphStyleRef) ParagraphStyleRef
	_CTParagraphStyleGetTypeID func() TypeID
	_CTParagraphStyleGetValueForSpecifier func(ParagraphStyleRef, ParagraphStyleSpecifier, uintptr, unsafe.Pointer) bool
	_CTRubyAnnotationCreate func(RubyAlignment, RubyOverhang, float64, StringRef, unsafe.Pointer) RubyAnnotationRef
	_CTRubyAnnotationCreateCopy func(RubyAnnotationRef) RubyAnnotationRef
	_CTRubyAnnotationCreateWithAttributes func(RubyAlignment, RubyOverhang, RubyPosition, StringRef, DictionaryRef) RubyAnnotationRef
	_CTRubyAnnotationGetAlignment func(RubyAnnotationRef) RubyAlignment
	_CTRubyAnnotationGetOverhang func(RubyAnnotationRef) RubyOverhang
	_CTRubyAnnotationGetSizeFactor func(RubyAnnotationRef) float64
	_CTRubyAnnotationGetTextForPosition func(RubyAnnotationRef, RubyPosition) StringRef
	_CTRubyAnnotationGetTypeID func() TypeID
	_CTRunDelegateCreate func(unsafe.Pointer, unsafe.Pointer) RunDelegateRef
	_CTRunDelegateGetRefCon func(RunDelegateRef) unsafe.Pointer
	_CTRunDelegateGetTypeID func() TypeID
	_CTRunDraw func(RunRef, ContextRef, corefoundation.Range)
	_CTRunGetAdvances func(RunRef, corefoundation.Range, corefoundation.CGSize)
	_CTRunGetAdvancesPtr func(RunRef) unsafe.Pointer
	_CTRunGetAttributes func(RunRef) DictionaryRef
	_CTRunGetBaseAdvancesAndOrigins func(RunRef, corefoundation.Range, corefoundation.CGSize, corefoundation.CGPoint)
	_CTRunGetGlyphCount func(RunRef) Index
	_CTRunGetGlyphs func(RunRef, corefoundation.Range, Glyph)
	_CTRunGetGlyphsPtr func(RunRef) unsafe.Pointer
	_CTRunGetImageBounds func(RunRef, ContextRef, corefoundation.Range) corefoundation.CGRect
	_CTRunGetPositions func(RunRef, corefoundation.Range, corefoundation.CGPoint)
	_CTRunGetPositionsPtr func(RunRef) unsafe.Pointer
	_CTRunGetStatus func(RunRef) RunStatus
	_CTRunGetStringIndices func(RunRef, corefoundation.Range, Index)
	_CTRunGetStringIndicesPtr func(RunRef) unsafe.Pointer
	_CTRunGetStringRange func(RunRef) corefoundation.Range
	_CTRunGetTextMatrix func(RunRef) corefoundation.CGAffineTransform
	_CTRunGetTypeID func() TypeID
	_CTRunGetTypographicBounds func(RunRef, corefoundation.Range, []float64, []float64, []float64) float64
	_CTTextTabCreate func(TextAlignment, float64, DictionaryRef) TextTabRef
	_CTTextTabGetAlignment func(TextTabRef) TextAlignment
	_CTTextTabGetLocation func(TextTabRef) float64
	_CTTextTabGetOptions func(TextTabRef) DictionaryRef
	_CTTextTabGetTypeID func() TypeID
	_CTTypesetterCreateLine func(TypesetterRef, corefoundation.Range) LineRef
	_CTTypesetterCreateLineWithOffset func(TypesetterRef, corefoundation.Range, float64) LineRef
	_CTTypesetterCreateWithAttributedString func(AttributedStringRef) TypesetterRef
	_CTTypesetterCreateWithAttributedStringAndOptions func(AttributedStringRef, DictionaryRef) TypesetterRef
	_CTTypesetterGetTypeID func() TypeID
	_CTTypesetterSuggestClusterBreak func(TypesetterRef, Index, float64) Index
	_CTTypesetterSuggestClusterBreakWithOffset func(TypesetterRef, Index, float64, float64) Index
	_CTTypesetterSuggestLineBreak func(TypesetterRef, Index, float64) Index
	_CTTypesetterSuggestLineBreakWithOffset func(TypesetterRef, Index, float64, float64) Index
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_CTFontCollectionCopyExclusionDescriptors, lib, "CTFontCollectionCopyExclusionDescriptors")
	tryRegister(&_CTFontCollectionCopyFontAttribute, lib, "CTFontCollectionCopyFontAttribute")
	tryRegister(&_CTFontCollectionCopyFontAttributes, lib, "CTFontCollectionCopyFontAttributes")
	tryRegister(&_CTFontCollectionCopyQueryDescriptors, lib, "CTFontCollectionCopyQueryDescriptors")
	tryRegister(&_CTFontCollectionCreateCopyWithFontDescriptors, lib, "CTFontCollectionCreateCopyWithFontDescriptors")
	tryRegister(&_CTFontCollectionCreateFromAvailableFonts, lib, "CTFontCollectionCreateFromAvailableFonts")
	tryRegister(&_CTFontCollectionCreateMatchingFontDescriptors, lib, "CTFontCollectionCreateMatchingFontDescriptors")
	tryRegister(&_CTFontCollectionCreateMatchingFontDescriptorsForFamily, lib, "CTFontCollectionCreateMatchingFontDescriptorsForFamily")
	tryRegister(&_CTFontCollectionCreateMatchingFontDescriptorsSortedWithCallback, lib, "CTFontCollectionCreateMatchingFontDescriptorsSortedWithCallback")
	tryRegister(&_CTFontCollectionCreateMatchingFontDescriptorsWithOptions, lib, "CTFontCollectionCreateMatchingFontDescriptorsWithOptions")
	tryRegister(&_CTFontCollectionCreateMutableCopy, lib, "CTFontCollectionCreateMutableCopy")
	tryRegister(&_CTFontCollectionCreateWithFontDescriptors, lib, "CTFontCollectionCreateWithFontDescriptors")
	tryRegister(&_CTFontCollectionGetTypeID, lib, "CTFontCollectionGetTypeID")
	tryRegister(&_CTFontCollectionSetExclusionDescriptors, lib, "CTFontCollectionSetExclusionDescriptors")
	tryRegister(&_CTFontCollectionSetQueryDescriptors, lib, "CTFontCollectionSetQueryDescriptors")
	tryRegister(&_CTFontCopyAttribute, lib, "CTFontCopyAttribute")
	tryRegister(&_CTFontCopyAvailableTables, lib, "CTFontCopyAvailableTables")
	tryRegister(&_CTFontCopyCharacterSet, lib, "CTFontCopyCharacterSet")
	tryRegister(&_CTFontCopyDefaultCascadeListForLanguages, lib, "CTFontCopyDefaultCascadeListForLanguages")
	tryRegister(&_CTFontCopyDisplayName, lib, "CTFontCopyDisplayName")
	tryRegister(&_CTFontCopyFamilyName, lib, "CTFontCopyFamilyName")
	tryRegister(&_CTFontCopyFeatures, lib, "CTFontCopyFeatures")
	tryRegister(&_CTFontCopyFeatureSettings, lib, "CTFontCopyFeatureSettings")
	tryRegister(&_CTFontCopyFontDescriptor, lib, "CTFontCopyFontDescriptor")
	tryRegister(&_CTFontCopyFullName, lib, "CTFontCopyFullName")
	tryRegister(&_CTFontCopyGraphicsFont, lib, "CTFontCopyGraphicsFont")
	tryRegister(&_CTFontCopyLocalizedName, lib, "CTFontCopyLocalizedName")
	tryRegister(&_CTFontCopyName, lib, "CTFontCopyName")
	tryRegister(&_CTFontCopyNameForGlyph, lib, "CTFontCopyNameForGlyph")
	tryRegister(&_CTFontCopyPostScriptName, lib, "CTFontCopyPostScriptName")
	tryRegister(&_CTFontCopySupportedLanguages, lib, "CTFontCopySupportedLanguages")
	tryRegister(&_CTFontCopyTable, lib, "CTFontCopyTable")
	tryRegister(&_CTFontCopyTraits, lib, "CTFontCopyTraits")
	tryRegister(&_CTFontCopyVariation, lib, "CTFontCopyVariation")
	tryRegister(&_CTFontCopyVariationAxes, lib, "CTFontCopyVariationAxes")
	tryRegister(&_CTFontCreateCopyWithAttributes, lib, "CTFontCreateCopyWithAttributes")
	tryRegister(&_CTFontCreateCopyWithFamily, lib, "CTFontCreateCopyWithFamily")
	tryRegister(&_CTFontCreateCopyWithSymbolicTraits, lib, "CTFontCreateCopyWithSymbolicTraits")
	tryRegister(&_CTFontCreateForString, lib, "CTFontCreateForString")
	tryRegister(&_CTFontCreateForStringWithLanguage, lib, "CTFontCreateForStringWithLanguage")
	tryRegister(&_CTFontCreatePathForGlyph, lib, "CTFontCreatePathForGlyph")
	tryRegister(&_CTFontCreateUIFontForLanguage, lib, "CTFontCreateUIFontForLanguage")
	tryRegister(&_CTFontCreateWithFontDescriptor, lib, "CTFontCreateWithFontDescriptor")
	tryRegister(&_CTFontCreateWithFontDescriptorAndOptions, lib, "CTFontCreateWithFontDescriptorAndOptions")
	tryRegister(&_CTFontCreateWithGraphicsFont, lib, "CTFontCreateWithGraphicsFont")
	tryRegister(&_CTFontCreateWithName, lib, "CTFontCreateWithName")
	tryRegister(&_CTFontCreateWithNameAndOptions, lib, "CTFontCreateWithNameAndOptions")
	tryRegister(&_CTFontCreateWithPlatformFont, lib, "CTFontCreateWithPlatformFont")
	tryRegister(&_CTFontCreateWithQuickdrawInstance, lib, "CTFontCreateWithQuickdrawInstance")
	tryRegister(&_CTFontDescriptorCopyAttribute, lib, "CTFontDescriptorCopyAttribute")
	tryRegister(&_CTFontDescriptorCopyAttributes, lib, "CTFontDescriptorCopyAttributes")
	tryRegister(&_CTFontDescriptorCopyLocalizedAttribute, lib, "CTFontDescriptorCopyLocalizedAttribute")
	tryRegister(&_CTFontDescriptorCreateCopyWithAttributes, lib, "CTFontDescriptorCreateCopyWithAttributes")
	tryRegister(&_CTFontDescriptorCreateCopyWithFamily, lib, "CTFontDescriptorCreateCopyWithFamily")
	tryRegister(&_CTFontDescriptorCreateCopyWithFeature, lib, "CTFontDescriptorCreateCopyWithFeature")
	tryRegister(&_CTFontDescriptorCreateCopyWithSymbolicTraits, lib, "CTFontDescriptorCreateCopyWithSymbolicTraits")
	tryRegister(&_CTFontDescriptorCreateCopyWithVariation, lib, "CTFontDescriptorCreateCopyWithVariation")
	tryRegister(&_CTFontDescriptorCreateMatchingFontDescriptor, lib, "CTFontDescriptorCreateMatchingFontDescriptor")
	tryRegister(&_CTFontDescriptorCreateMatchingFontDescriptors, lib, "CTFontDescriptorCreateMatchingFontDescriptors")
	tryRegister(&_CTFontDescriptorCreateWithAttributes, lib, "CTFontDescriptorCreateWithAttributes")
	tryRegister(&_CTFontDescriptorCreateWithNameAndSize, lib, "CTFontDescriptorCreateWithNameAndSize")
	tryRegister(&_CTFontDescriptorGetTypeID, lib, "CTFontDescriptorGetTypeID")
	tryRegister(&_CTFontDescriptorMatchFontDescriptorsWithProgressHandler, lib, "CTFontDescriptorMatchFontDescriptorsWithProgressHandler")
	tryRegister(&_CTFontDrawGlyphs, lib, "CTFontDrawGlyphs")
	tryRegister(&_CTFontDrawImageFromAdaptiveImageProviderAtPoint, lib, "CTFontDrawImageFromAdaptiveImageProviderAtPoint")
	tryRegister(&_CTFontGetAdvancesForGlyphs, lib, "CTFontGetAdvancesForGlyphs")
	tryRegister(&_CTFontGetAscent, lib, "CTFontGetAscent")
	tryRegister(&_CTFontGetBoundingBox, lib, "CTFontGetBoundingBox")
	tryRegister(&_CTFontGetBoundingRectsForGlyphs, lib, "CTFontGetBoundingRectsForGlyphs")
	tryRegister(&_CTFontGetCapHeight, lib, "CTFontGetCapHeight")
	tryRegister(&_CTFontGetDescent, lib, "CTFontGetDescent")
	tryRegister(&_CTFontGetGlyphCount, lib, "CTFontGetGlyphCount")
	tryRegister(&_CTFontGetGlyphsForCharacters, lib, "CTFontGetGlyphsForCharacters")
	tryRegister(&_CTFontGetGlyphWithName, lib, "CTFontGetGlyphWithName")
	tryRegister(&_CTFontGetLeading, lib, "CTFontGetLeading")
	tryRegister(&_CTFontGetLigatureCaretPositions, lib, "CTFontGetLigatureCaretPositions")
	tryRegister(&_CTFontGetMatrix, lib, "CTFontGetMatrix")
	tryRegister(&_CTFontGetOpticalBoundsForGlyphs, lib, "CTFontGetOpticalBoundsForGlyphs")
	tryRegister(&_CTFontGetPlatformFont, lib, "CTFontGetPlatformFont")
	tryRegister(&_CTFontGetSize, lib, "CTFontGetSize")
	tryRegister(&_CTFontGetSlantAngle, lib, "CTFontGetSlantAngle")
	tryRegister(&_CTFontGetStringEncoding, lib, "CTFontGetStringEncoding")
	tryRegister(&_CTFontGetSymbolicTraits, lib, "CTFontGetSymbolicTraits")
	tryRegister(&_CTFontGetTypeID, lib, "CTFontGetTypeID")
	tryRegister(&_CTFontGetTypographicBoundsForAdaptiveImageProvider, lib, "CTFontGetTypographicBoundsForAdaptiveImageProvider")
	tryRegister(&_CTFontGetUnderlinePosition, lib, "CTFontGetUnderlinePosition")
	tryRegister(&_CTFontGetUnderlineThickness, lib, "CTFontGetUnderlineThickness")
	tryRegister(&_CTFontGetUnitsPerEm, lib, "CTFontGetUnitsPerEm")
	tryRegister(&_CTFontGetVerticalTranslationsForGlyphs, lib, "CTFontGetVerticalTranslationsForGlyphs")
	tryRegister(&_CTFontGetXHeight, lib, "CTFontGetXHeight")
	tryRegister(&_CTFontHasTable, lib, "CTFontHasTable")
	tryRegister(&_CTFontManagerCompareFontFamilyNames, lib, "CTFontManagerCompareFontFamilyNames")
	tryRegister(&_CTFontManagerCopyAvailableFontFamilyNames, lib, "CTFontManagerCopyAvailableFontFamilyNames")
	tryRegister(&_CTFontManagerCopyAvailableFontURLs, lib, "CTFontManagerCopyAvailableFontURLs")
	tryRegister(&_CTFontManagerCopyAvailablePostScriptNames, lib, "CTFontManagerCopyAvailablePostScriptNames")
	tryRegister(&_CTFontManagerCopyRegisteredFontDescriptors, lib, "CTFontManagerCopyRegisteredFontDescriptors")
	tryRegister(&_CTFontManagerCreateFontDescriptorFromData, lib, "CTFontManagerCreateFontDescriptorFromData")
	tryRegister(&_CTFontManagerCreateFontDescriptorsFromData, lib, "CTFontManagerCreateFontDescriptorsFromData")
	tryRegister(&_CTFontManagerCreateFontDescriptorsFromURL, lib, "CTFontManagerCreateFontDescriptorsFromURL")
	tryRegister(&_CTFontManagerCreateFontRequestRunLoopSource, lib, "CTFontManagerCreateFontRequestRunLoopSource")
	tryRegister(&_CTFontManagerEnableFontDescriptors, lib, "CTFontManagerEnableFontDescriptors")
	tryRegister(&_CTFontManagerGetAutoActivationSetting, lib, "CTFontManagerGetAutoActivationSetting")
	tryRegister(&_CTFontManagerGetScopeForURL, lib, "CTFontManagerGetScopeForURL")
	tryRegister(&_CTFontManagerIsSupportedFont, lib, "CTFontManagerIsSupportedFont")
	tryRegister(&_CTFontManagerRegisterFontDescriptors, lib, "CTFontManagerRegisterFontDescriptors")
	tryRegister(&_CTFontManagerRegisterFontsForURL, lib, "CTFontManagerRegisterFontsForURL")
	tryRegister(&_CTFontManagerRegisterFontsForURLs, lib, "CTFontManagerRegisterFontsForURLs")
	tryRegister(&_CTFontManagerRegisterFontsWithAssetNames, lib, "CTFontManagerRegisterFontsWithAssetNames")
	tryRegister(&_CTFontManagerRegisterFontURLs, lib, "CTFontManagerRegisterFontURLs")
	tryRegister(&_CTFontManagerRegisterGraphicsFont, lib, "CTFontManagerRegisterGraphicsFont")
	tryRegister(&_CTFontManagerRequestFonts, lib, "CTFontManagerRequestFonts")
	tryRegister(&_CTFontManagerSetAutoActivationSetting, lib, "CTFontManagerSetAutoActivationSetting")
	tryRegister(&_CTFontManagerUnregisterFontDescriptors, lib, "CTFontManagerUnregisterFontDescriptors")
	tryRegister(&_CTFontManagerUnregisterFontsForURL, lib, "CTFontManagerUnregisterFontsForURL")
	tryRegister(&_CTFontManagerUnregisterFontsForURLs, lib, "CTFontManagerUnregisterFontsForURLs")
	tryRegister(&_CTFontManagerUnregisterFontURLs, lib, "CTFontManagerUnregisterFontURLs")
	tryRegister(&_CTFontManagerUnregisterGraphicsFont, lib, "CTFontManagerUnregisterGraphicsFont")
	tryRegister(&_CTFrameDraw, lib, "CTFrameDraw")
	tryRegister(&_CTFrameGetFrameAttributes, lib, "CTFrameGetFrameAttributes")
	tryRegister(&_CTFrameGetLineOrigins, lib, "CTFrameGetLineOrigins")
	tryRegister(&_CTFrameGetLines, lib, "CTFrameGetLines")
	tryRegister(&_CTFrameGetPath, lib, "CTFrameGetPath")
	tryRegister(&_CTFrameGetStringRange, lib, "CTFrameGetStringRange")
	tryRegister(&_CTFrameGetTypeID, lib, "CTFrameGetTypeID")
	tryRegister(&_CTFrameGetVisibleStringRange, lib, "CTFrameGetVisibleStringRange")
	tryRegister(&_CTFramesetterCreateFrame, lib, "CTFramesetterCreateFrame")
	tryRegister(&_CTFramesetterCreateWithAttributedString, lib, "CTFramesetterCreateWithAttributedString")
	tryRegister(&_CTFramesetterCreateWithTypesetter, lib, "CTFramesetterCreateWithTypesetter")
	tryRegister(&_CTFramesetterGetTypeID, lib, "CTFramesetterGetTypeID")
	tryRegister(&_CTFramesetterGetTypesetter, lib, "CTFramesetterGetTypesetter")
	tryRegister(&_CTFramesetterSuggestFrameSizeWithConstraints, lib, "CTFramesetterSuggestFrameSizeWithConstraints")
	tryRegister(&_CTGetCoreTextVersion, lib, "CTGetCoreTextVersion")
	tryRegister(&_CTGlyphInfoCreateWithCharacterIdentifier, lib, "CTGlyphInfoCreateWithCharacterIdentifier")
	tryRegister(&_CTGlyphInfoCreateWithGlyph, lib, "CTGlyphInfoCreateWithGlyph")
	tryRegister(&_CTGlyphInfoCreateWithGlyphName, lib, "CTGlyphInfoCreateWithGlyphName")
	tryRegister(&_CTGlyphInfoGetCharacterCollection, lib, "CTGlyphInfoGetCharacterCollection")
	tryRegister(&_CTGlyphInfoGetCharacterIdentifier, lib, "CTGlyphInfoGetCharacterIdentifier")
	tryRegister(&_CTGlyphInfoGetGlyph, lib, "CTGlyphInfoGetGlyph")
	tryRegister(&_CTGlyphInfoGetGlyphName, lib, "CTGlyphInfoGetGlyphName")
	tryRegister(&_CTGlyphInfoGetTypeID, lib, "CTGlyphInfoGetTypeID")
	tryRegister(&_CTLineCreateJustifiedLine, lib, "CTLineCreateJustifiedLine")
	tryRegister(&_CTLineCreateTruncatedLine, lib, "CTLineCreateTruncatedLine")
	tryRegister(&_CTLineCreateWithAttributedString, lib, "CTLineCreateWithAttributedString")
	tryRegister(&_CTLineDraw, lib, "CTLineDraw")
	tryRegister(&_CTLineEnumerateCaretOffsets, lib, "CTLineEnumerateCaretOffsets")
	tryRegister(&_CTLineGetBoundsWithOptions, lib, "CTLineGetBoundsWithOptions")
	tryRegister(&_CTLineGetGlyphCount, lib, "CTLineGetGlyphCount")
	tryRegister(&_CTLineGetGlyphRuns, lib, "CTLineGetGlyphRuns")
	tryRegister(&_CTLineGetImageBounds, lib, "CTLineGetImageBounds")
	tryRegister(&_CTLineGetOffsetForStringIndex, lib, "CTLineGetOffsetForStringIndex")
	tryRegister(&_CTLineGetPenOffsetForFlush, lib, "CTLineGetPenOffsetForFlush")
	tryRegister(&_CTLineGetStringIndexForPosition, lib, "CTLineGetStringIndexForPosition")
	tryRegister(&_CTLineGetStringRange, lib, "CTLineGetStringRange")
	tryRegister(&_CTLineGetTrailingWhitespaceWidth, lib, "CTLineGetTrailingWhitespaceWidth")
	tryRegister(&_CTLineGetTypeID, lib, "CTLineGetTypeID")
	tryRegister(&_CTLineGetTypographicBounds, lib, "CTLineGetTypographicBounds")
	tryRegister(&_CTParagraphStyleCreate, lib, "CTParagraphStyleCreate")
	tryRegister(&_CTParagraphStyleCreateCopy, lib, "CTParagraphStyleCreateCopy")
	tryRegister(&_CTParagraphStyleGetTypeID, lib, "CTParagraphStyleGetTypeID")
	tryRegister(&_CTParagraphStyleGetValueForSpecifier, lib, "CTParagraphStyleGetValueForSpecifier")
	tryRegister(&_CTRubyAnnotationCreate, lib, "CTRubyAnnotationCreate")
	tryRegister(&_CTRubyAnnotationCreateCopy, lib, "CTRubyAnnotationCreateCopy")
	tryRegister(&_CTRubyAnnotationCreateWithAttributes, lib, "CTRubyAnnotationCreateWithAttributes")
	tryRegister(&_CTRubyAnnotationGetAlignment, lib, "CTRubyAnnotationGetAlignment")
	tryRegister(&_CTRubyAnnotationGetOverhang, lib, "CTRubyAnnotationGetOverhang")
	tryRegister(&_CTRubyAnnotationGetSizeFactor, lib, "CTRubyAnnotationGetSizeFactor")
	tryRegister(&_CTRubyAnnotationGetTextForPosition, lib, "CTRubyAnnotationGetTextForPosition")
	tryRegister(&_CTRubyAnnotationGetTypeID, lib, "CTRubyAnnotationGetTypeID")
	tryRegister(&_CTRunDelegateCreate, lib, "CTRunDelegateCreate")
	tryRegister(&_CTRunDelegateGetRefCon, lib, "CTRunDelegateGetRefCon")
	tryRegister(&_CTRunDelegateGetTypeID, lib, "CTRunDelegateGetTypeID")
	tryRegister(&_CTRunDraw, lib, "CTRunDraw")
	tryRegister(&_CTRunGetAdvances, lib, "CTRunGetAdvances")
	tryRegister(&_CTRunGetAdvancesPtr, lib, "CTRunGetAdvancesPtr")
	tryRegister(&_CTRunGetAttributes, lib, "CTRunGetAttributes")
	tryRegister(&_CTRunGetBaseAdvancesAndOrigins, lib, "CTRunGetBaseAdvancesAndOrigins")
	tryRegister(&_CTRunGetGlyphCount, lib, "CTRunGetGlyphCount")
	tryRegister(&_CTRunGetGlyphs, lib, "CTRunGetGlyphs")
	tryRegister(&_CTRunGetGlyphsPtr, lib, "CTRunGetGlyphsPtr")
	tryRegister(&_CTRunGetImageBounds, lib, "CTRunGetImageBounds")
	tryRegister(&_CTRunGetPositions, lib, "CTRunGetPositions")
	tryRegister(&_CTRunGetPositionsPtr, lib, "CTRunGetPositionsPtr")
	tryRegister(&_CTRunGetStatus, lib, "CTRunGetStatus")
	tryRegister(&_CTRunGetStringIndices, lib, "CTRunGetStringIndices")
	tryRegister(&_CTRunGetStringIndicesPtr, lib, "CTRunGetStringIndicesPtr")
	tryRegister(&_CTRunGetStringRange, lib, "CTRunGetStringRange")
	tryRegister(&_CTRunGetTextMatrix, lib, "CTRunGetTextMatrix")
	tryRegister(&_CTRunGetTypeID, lib, "CTRunGetTypeID")
	tryRegister(&_CTRunGetTypographicBounds, lib, "CTRunGetTypographicBounds")
	tryRegister(&_CTTextTabCreate, lib, "CTTextTabCreate")
	tryRegister(&_CTTextTabGetAlignment, lib, "CTTextTabGetAlignment")
	tryRegister(&_CTTextTabGetLocation, lib, "CTTextTabGetLocation")
	tryRegister(&_CTTextTabGetOptions, lib, "CTTextTabGetOptions")
	tryRegister(&_CTTextTabGetTypeID, lib, "CTTextTabGetTypeID")
	tryRegister(&_CTTypesetterCreateLine, lib, "CTTypesetterCreateLine")
	tryRegister(&_CTTypesetterCreateLineWithOffset, lib, "CTTypesetterCreateLineWithOffset")
	tryRegister(&_CTTypesetterCreateWithAttributedString, lib, "CTTypesetterCreateWithAttributedString")
	tryRegister(&_CTTypesetterCreateWithAttributedStringAndOptions, lib, "CTTypesetterCreateWithAttributedStringAndOptions")
	tryRegister(&_CTTypesetterGetTypeID, lib, "CTTypesetterGetTypeID")
	tryRegister(&_CTTypesetterSuggestClusterBreak, lib, "CTTypesetterSuggestClusterBreak")
	tryRegister(&_CTTypesetterSuggestClusterBreakWithOffset, lib, "CTTypesetterSuggestClusterBreakWithOffset")
	tryRegister(&_CTTypesetterSuggestLineBreak, lib, "CTTypesetterSuggestLineBreak")
	tryRegister(&_CTTypesetterSuggestLineBreakWithOffset, lib, "CTTypesetterSuggestLineBreakWithOffset")
}

// tryRegister attempts to register a function, silently ignoring failures.
// This allows the library to load even if some symbols are missing.
func tryRegister(fn interface{}, lib uintptr, name string) {
	defer func() {
		if r := recover(); r != nil {
			// Symbol not found - function will remain nil and panic when called
			// This is expected for inline functions, macros, or version-specific APIs
		}
	}()
	purego.RegisterLibFunc(fn, lib, name)
}



// Retrieves the array of descriptors to exclude from the match.
//
// Added in macOS 10.7.
// Retrieves the array of descriptors to exclude from the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCollectionCopyExclusionDescriptors(_:)
func CTFontCollectionCopyExclusionDescriptors(collection FontCollectionRef) ArrayRef {
	return _CTFontCollectionCopyExclusionDescriptors(collection)
}/* debug [functions.gen.go/function]: CTFontCollectionCopyExclusionDescriptors */

// Retrieves an array of font descriptor attribute values.
//
// Added in macOS 10.7.
// Retrieves an array of font descriptor attribute values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCollectionCopyFontAttribute(_:_:_:)
func CTFontCollectionCopyFontAttribute(collection FontCollectionRef, attributeName StringRef, options FontCollectionCopyOptions) ArrayRef {
	return _CTFontCollectionCopyFontAttribute(collection, attributeName, options)
}/* debug [functions.gen.go/function]: CTFontCollectionCopyFontAttribute */

// Retrieves an array of dictionaries containing font descriptor attribute values.
//
// Added in macOS 10.7.
// Retrieves an array of dictionaries containing font descriptor attribute values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCollectionCopyFontAttributes(_:_:_:)
func CTFontCollectionCopyFontAttributes(collection FontCollectionRef, attributeNames SetRef, options FontCollectionCopyOptions) ArrayRef {
	return _CTFontCollectionCopyFontAttributes(collection, attributeNames, options)
}/* debug [functions.gen.go/function]: CTFontCollectionCopyFontAttributes */

// Retrieves the array of descriptors for font matching.
//
// Added in macOS 10.7.
// Retrieves the array of descriptors for font matching.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCollectionCopyQueryDescriptors(_:)
func CTFontCollectionCopyQueryDescriptors(collection FontCollectionRef) ArrayRef {
	return _CTFontCollectionCopyQueryDescriptors(collection)
}/* debug [functions.gen.go/function]: CTFontCollectionCopyQueryDescriptors */

// Returns a copy of the original collection augmented with the given new font descriptors.
//
// Added in macOS 10.5.
// Returns a copy of the original collection augmented with the given new font descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCollectionCreateCopyWithFontDescriptors(_:_:_:)
func CTFontCollectionCreateCopyWithFontDescriptors(original FontCollectionRef, queryDescriptors ArrayRef, options DictionaryRef) FontCollectionRef {
	return _CTFontCollectionCreateCopyWithFontDescriptors(original, queryDescriptors, options)
}/* debug [functions.gen.go/function]: CTFontCollectionCreateCopyWithFontDescriptors */

// Returns a new font collection containing all available fonts.
//
// Added in macOS 10.5.
// Returns a new font collection containing all available fonts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCollectionCreateFromAvailableFonts(_:)
func CTFontCollectionCreateFromAvailableFonts(options DictionaryRef) FontCollectionRef {
	return _CTFontCollectionCreateFromAvailableFonts(options)
}/* debug [functions.gen.go/function]: CTFontCollectionCreateFromAvailableFonts */

// Returns an array of font descriptors matching the collection.
//
// Added in macOS 10.5.
// Returns an array of font descriptors matching the collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCollectionCreateMatchingFontDescriptors(_:)
func CTFontCollectionCreateMatchingFontDescriptors(collection FontCollectionRef) ArrayRef {
	return _CTFontCollectionCreateMatchingFontDescriptors(collection)
}/* debug [functions.gen.go/function]: CTFontCollectionCreateMatchingFontDescriptors */

// Retrieves an array of font descriptors that match the specified family, one descriptor for each style in the collection.
//
// Added in macOS 10.7.
// Retrieves an array of font descriptors that match the specified family, one descriptor for each style in the collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCollectionCreateMatchingFontDescriptorsForFamily(_:_:_:)
func CTFontCollectionCreateMatchingFontDescriptorsForFamily(collection FontCollectionRef, familyName StringRef, options DictionaryRef) ArrayRef {
	return _CTFontCollectionCreateMatchingFontDescriptorsForFamily(collection, familyName, options)
}/* debug [functions.gen.go/function]: CTFontCollectionCreateMatchingFontDescriptorsForFamily */

// Returns the array of matching font descriptors sorted with the callback function.
//
// Added in macOS 10.5.
// Returns the array of matching font descriptors sorted with the callback function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCollectionCreateMatchingFontDescriptorsSortedWithCallback(_:_:_:)
func CTFontCollectionCreateMatchingFontDescriptorsSortedWithCallback(collection FontCollectionRef, sortCallback FontCollectionSortDescriptorsCallback, refCon unsafe.Pointer) ArrayRef {
	return _CTFontCollectionCreateMatchingFontDescriptorsSortedWithCallback(collection, sortCallback, refCon)
}/* debug [functions.gen.go/function]: CTFontCollectionCreateMatchingFontDescriptorsSortedWithCallback */

// Creates an array of font descriptors that match the specified collection.
//
// Added in macOS 10.7.
// Creates an array of font descriptors that match the specified collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCollectionCreateMatchingFontDescriptorsWithOptions(_:_:)
func CTFontCollectionCreateMatchingFontDescriptorsWithOptions(collection FontCollectionRef, options DictionaryRef) ArrayRef {
	return _CTFontCollectionCreateMatchingFontDescriptorsWithOptions(collection, options)
}/* debug [functions.gen.go/function]: CTFontCollectionCreateMatchingFontDescriptorsWithOptions */

// Creates a mutable copy of the original collection.
//
// Added in macOS 10.7.
// Creates a mutable copy of the original collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCollectionCreateMutableCopy(_:)
func CTFontCollectionCreateMutableCopy(original FontCollectionRef) MutableFontCollectionRef {
	return _CTFontCollectionCreateMutableCopy(original)
}/* debug [functions.gen.go/function]: CTFontCollectionCreateMutableCopy */

// Returns a new font collection based on the given array of font descriptors.
//
// Added in macOS 10.5.
// Returns a new font collection based on the given array of font descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCollectionCreateWithFontDescriptors(_:_:)
func CTFontCollectionCreateWithFontDescriptors(queryDescriptors ArrayRef, options DictionaryRef) FontCollectionRef {
	return _CTFontCollectionCreateWithFontDescriptors(queryDescriptors, options)
}/* debug [functions.gen.go/function]: CTFontCollectionCreateWithFontDescriptors */

// Returns the type identifier for Core Text font collection references.
//
// Added in macOS 10.5.
// Returns the type identifier for Core Text font collection references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCollectionGetTypeID()
func CTFontCollectionGetTypeID() TypeID {
	return _CTFontCollectionGetTypeID()
}/* debug [functions.gen.go/function]: CTFontCollectionGetTypeID */

// Replaces the array of descriptors to exclude from the match.
//
// Added in macOS 10.7.
// Replaces the array of descriptors to exclude from the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCollectionSetExclusionDescriptors(_:_:)
func CTFontCollectionSetExclusionDescriptors(collection MutableFontCollectionRef, descriptors ArrayRef) {
	_CTFontCollectionSetExclusionDescriptors(collection, descriptors)
}/* debug [functions.gen.go/function]: CTFontCollectionSetExclusionDescriptors */

// Replaces the array of descriptors for font matching.
//
// Added in macOS 10.7.
// Replaces the array of descriptors for font matching.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCollectionSetQueryDescriptors(_:_:)
func CTFontCollectionSetQueryDescriptors(collection MutableFontCollectionRef, descriptors ArrayRef) {
	_CTFontCollectionSetQueryDescriptors(collection, descriptors)
}/* debug [functions.gen.go/function]: CTFontCollectionSetQueryDescriptors */

// Returns the value associated with an arbitrary attribute of the given font.
//
// Added in macOS 10.5.
// Returns the value associated with an arbitrary attribute of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCopyAttribute(_:_:)
func CTFontCopyAttribute(font FontRef, attribute StringRef) TypeRef {
	return _CTFontCopyAttribute(font, attribute)
}/* debug [functions.gen.go/function]: CTFontCopyAttribute */

// Returns an array of font table tags.
//
// Added in macOS 10.5.
// Returns an array of font table tags.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCopyAvailableTables(_:_:)
func CTFontCopyAvailableTables(font FontRef, options FontTableOptions) ArrayRef {
	return _CTFontCopyAvailableTables(font, options)
}/* debug [functions.gen.go/function]: CTFontCopyAvailableTables */

// Returns the Unicode character set of the font.
//
// Added in macOS 10.5.
// Returns the Unicode character set of the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCopyCharacterSet(_:)
func CTFontCopyCharacterSet(font FontRef) CharacterSetRef {
	return _CTFontCopyCharacterSet(font)
}/* debug [functions.gen.go/function]: CTFontCopyCharacterSet */

// Retrieves an ordered list of font substitution preferences.
//
// Added in macOS 10.8.
// Retrieves an ordered list of font substitution preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCopyDefaultCascadeListForLanguages(_:_:)
func CTFontCopyDefaultCascadeListForLanguages(font FontRef, languagePrefList ArrayRef) ArrayRef {
	return _CTFontCopyDefaultCascadeListForLanguages(font, languagePrefList)
}/* debug [functions.gen.go/function]: CTFontCopyDefaultCascadeListForLanguages */

// Returns the display name of the given font.
//
// Added in macOS 10.5.
// Returns the display name of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCopyDisplayName(_:)
func CTFontCopyDisplayName(font FontRef) StringRef {
	return _CTFontCopyDisplayName(font)
}/* debug [functions.gen.go/function]: CTFontCopyDisplayName */

// Returns the family name of the given font.
//
// Added in macOS 10.5.
// Returns the family name of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCopyFamilyName(_:)
func CTFontCopyFamilyName(font FontRef) StringRef {
	return _CTFontCopyFamilyName(font)
}/* debug [functions.gen.go/function]: CTFontCopyFamilyName */

// Returns an array of font features.
//
// Added in macOS 10.5.
// Returns an array of font features.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCopyFeatures(_:)
func CTFontCopyFeatures(font FontRef) ArrayRef {
	return _CTFontCopyFeatures(font)
}/* debug [functions.gen.go/function]: CTFontCopyFeatures */

// Returns an array of font feature-setting tuples.
//
// Added in macOS 10.5.
// Returns an array of font feature-setting tuples.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCopyFeatureSettings(_:)
func CTFontCopyFeatureSettings(font FontRef) ArrayRef {
	return _CTFontCopyFeatureSettings(font)
}/* debug [functions.gen.go/function]: CTFontCopyFeatureSettings */

// Returns the normalized font descriptor for the given font reference.
//
// Added in macOS 10.5.
// Returns the normalized font descriptor for the given font reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCopyFontDescriptor(_:)
func CTFontCopyFontDescriptor(font FontRef) FontDescriptorRef {
	return _CTFontCopyFontDescriptor(font)
}/* debug [functions.gen.go/function]: CTFontCopyFontDescriptor */

// Returns the full name of the given font.
//
// Added in macOS 10.5.
// Returns the full name of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCopyFullName(_:)
func CTFontCopyFullName(font FontRef) StringRef {
	return _CTFontCopyFullName(font)
}/* debug [functions.gen.go/function]: CTFontCopyFullName */

// Returns a Core Graphics font reference and attributes.
//
// Added in macOS 10.5.
// Returns a Core Graphics font reference and attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCopyGraphicsFont(_:_:)
func CTFontCopyGraphicsFont(font FontRef, attributes unsafe.Pointer) FontRef {
	return _CTFontCopyGraphicsFont(font, attributes)
}/* debug [functions.gen.go/function]: CTFontCopyGraphicsFont */

// Returns a reference to a localized name for the given font.
//
// Added in macOS 10.5.
// Returns a reference to a localized name for the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCopyLocalizedName(_:_:_:)
func CTFontCopyLocalizedName(font FontRef, nameKey StringRef, actualLanguage unsafe.Pointer) StringRef {
	return _CTFontCopyLocalizedName(font, nameKey, actualLanguage)
}/* debug [functions.gen.go/function]: CTFontCopyLocalizedName */

// Returns a reference to the requested name of the given font.
//
// Added in macOS 10.5.
// Returns a reference to the requested name of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCopyName(_:_:)
func CTFontCopyName(font FontRef, nameKey StringRef) StringRef {
	return _CTFontCopyName(font, nameKey)
}/* debug [functions.gen.go/function]: CTFontCopyName */

// Retrieves the name for the specified glyph.
//
// Added in macOS 10.8.
// Retrieves the name for the specified glyph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCopyNameForGlyph(_:_:)
func CTFontCopyNameForGlyph(font FontRef, glyph Glyph) StringRef {
	return _CTFontCopyNameForGlyph(font, glyph)
}/* debug [functions.gen.go/function]: CTFontCopyNameForGlyph */

// Returns the PostScript name of the given font.
//
// Added in macOS 10.5.
// Returns the PostScript name of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCopyPostScriptName(_:)
func CTFontCopyPostScriptName(font FontRef) StringRef {
	return _CTFontCopyPostScriptName(font)
}/* debug [functions.gen.go/function]: CTFontCopyPostScriptName */

// Returns an array of languages supported by the font.
//
// Added in macOS 10.5.
// Returns an array of languages supported by the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCopySupportedLanguages(_:)
func CTFontCopySupportedLanguages(font FontRef) ArrayRef {
	return _CTFontCopySupportedLanguages(font)
}/* debug [functions.gen.go/function]: CTFontCopySupportedLanguages */

// Returns a reference to the font table data.
//
// Added in macOS 10.5.
// Returns a reference to the font table data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCopyTable(_:_:_:)
func CTFontCopyTable(font FontRef, table FontTableTag, options FontTableOptions) DataRef {
	return _CTFontCopyTable(font, table, options)
}/* debug [functions.gen.go/function]: CTFontCopyTable */

// Returns the traits dictionary of the given font.
//
// Added in macOS 10.5.
// Returns the traits dictionary of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCopyTraits(_:)
func CTFontCopyTraits(font FontRef) DictionaryRef {
	return _CTFontCopyTraits(font)
}/* debug [functions.gen.go/function]: CTFontCopyTraits */

// Returns a variation dictionary from the font reference.
//
// Added in macOS 10.5.
// Returns a variation dictionary from the font reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCopyVariation(_:)
func CTFontCopyVariation(font FontRef) DictionaryRef {
	return _CTFontCopyVariation(font)
}/* debug [functions.gen.go/function]: CTFontCopyVariation */

// Returns an array of variation axes.
//
// Added in macOS 10.5.
// Returns an array of variation axes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCopyVariationAxes(_:)
func CTFontCopyVariationAxes(font FontRef) ArrayRef {
	return _CTFontCopyVariationAxes(font)
}/* debug [functions.gen.go/function]: CTFontCopyVariationAxes */

// Returns a new font with additional attributes based on the original font.
//
// Added in macOS 10.5.
// Returns a new font with additional attributes based on the original font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCreateCopyWithAttributes(_:_:_:_:)
func CTFontCreateCopyWithAttributes(font FontRef, size float64, matrix unsafe.Pointer, attributes FontDescriptorRef) FontRef {
	return _CTFontCreateCopyWithAttributes(font, size, matrix, attributes)
}/* debug [functions.gen.go/function]: CTFontCreateCopyWithAttributes */

// Returns a new font in the specified family based on the traits of the original font.
//
// Added in macOS 10.5.
// Returns a new font in the specified family based on the traits of the original font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCreateCopyWithFamily(_:_:_:_:)
func CTFontCreateCopyWithFamily(font FontRef, size float64, matrix unsafe.Pointer, family StringRef) FontRef {
	return _CTFontCreateCopyWithFamily(font, size, matrix, family)
}/* debug [functions.gen.go/function]: CTFontCreateCopyWithFamily */

// Returns a new font in the same font family as the original with the specified symbolic traits.
//
// Added in macOS 10.5.
// Returns a new font in the same font family as the original with the specified symbolic traits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCreateCopyWithSymbolicTraits(_:_:_:_:_:)
func CTFontCreateCopyWithSymbolicTraits(font FontRef, size float64, matrix unsafe.Pointer, symTraitValue FontSymbolicTraits, symTraitMask FontSymbolicTraits) FontRef {
	return _CTFontCreateCopyWithSymbolicTraits(font, size, matrix, symTraitValue, symTraitMask)
}/* debug [functions.gen.go/function]: CTFontCreateCopyWithSymbolicTraits */

// Returns a font reference that most accurately maps the string range based on the current font.
//
// Added in macOS 10.5.
// Returns a font reference that most accurately maps the string range based on the current font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCreateForString(_:_:_:)
func CTFontCreateForString(currentFont FontRef, string_ StringRef, range_ corefoundation.Range) FontRef {
	return _CTFontCreateForString(currentFont, string_, range_)
}/* debug [functions.gen.go/function]: CTFontCreateForString */

// Returns a font reference that most accurately maps the string range based on the current font and language.
//
// Added in macOS 10.9.
// Returns a font reference that most accurately maps the string range based on the current font and language.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCreateForStringWithLanguage(_:_:_:_:)
func CTFontCreateForStringWithLanguage(currentFont FontRef, string_ StringRef, range_ corefoundation.Range, language StringRef) FontRef {
	return _CTFontCreateForStringWithLanguage(currentFont, string_, range_, language)
}/* debug [functions.gen.go/function]: CTFontCreateForStringWithLanguage */

// Creates a path for the specified glyph.
//
// Added in macOS 10.5.
// Creates a path for the specified glyph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCreatePathForGlyph(_:_:_:)
func CTFontCreatePathForGlyph(font FontRef, glyph Glyph, matrix unsafe.Pointer) PathRef {
	return _CTFontCreatePathForGlyph(font, glyph, matrix)
}/* debug [functions.gen.go/function]: CTFontCreatePathForGlyph */

// Returns the special user-interface font for the given language and user-interface type.
//
// Added in macOS 10.5.
// Returns the special user-interface font for the given language and user-interface type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCreateUIFontForLanguage(_:_:_:)
func CTFontCreateUIFontForLanguage(uiType FontUIFontType, size float64, language StringRef) FontRef {
	return _CTFontCreateUIFontForLanguage(uiType, size, language)
}/* debug [functions.gen.go/function]: CTFontCreateUIFontForLanguage */

// Returns a new font reference that best matches the given font descriptor.
//
// Added in macOS 10.5.
// Returns a new font reference that best matches the given font descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCreateWithFontDescriptor(_:_:_:)
func CTFontCreateWithFontDescriptor(descriptor FontDescriptorRef, size float64, matrix unsafe.Pointer) FontRef {
	return _CTFontCreateWithFontDescriptor(descriptor, size, matrix)
}/* debug [functions.gen.go/function]: CTFontCreateWithFontDescriptor */

// Returns a new font reference that best matches the given font descriptor.
//
// Added in macOS 10.6.
// Returns a new font reference that best matches the given font descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCreateWithFontDescriptorAndOptions(_:_:_:_:)
func CTFontCreateWithFontDescriptorAndOptions(descriptor FontDescriptorRef, size float64, matrix unsafe.Pointer, options FontOptions) FontRef {
	return _CTFontCreateWithFontDescriptorAndOptions(descriptor, size, matrix, options)
}/* debug [functions.gen.go/function]: CTFontCreateWithFontDescriptorAndOptions */

// Creates a new font reference from an existing Core Graphics font reference.
//
// Added in macOS 10.5.
// Creates a new font reference from an existing Core Graphics font reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCreateWithGraphicsFont(_:_:_:_:)
func CTFontCreateWithGraphicsFont(graphicsFont FontRef, size float64, matrix unsafe.Pointer, attributes FontDescriptorRef) FontRef {
	return _CTFontCreateWithGraphicsFont(graphicsFont, size, matrix, attributes)
}/* debug [functions.gen.go/function]: CTFontCreateWithGraphicsFont */

// Returns a new font reference for the given name.
//
// Added in macOS 10.5.
// Returns a new font reference for the given name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCreateWithName(_:_:_:)
func CTFontCreateWithName(name StringRef, size float64, matrix unsafe.Pointer) FontRef {
	return _CTFontCreateWithName(name, size, matrix)
}/* debug [functions.gen.go/function]: CTFontCreateWithName */

// Returns a new font reference for the given name.
//
// Added in macOS 10.6.
// Returns a new font reference for the given name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCreateWithNameAndOptions(_:_:_:_:)
func CTFontCreateWithNameAndOptions(name StringRef, size float64, matrix unsafe.Pointer, options FontOptions) FontRef {
	return _CTFontCreateWithNameAndOptions(name, size, matrix, options)
}/* debug [functions.gen.go/function]: CTFontCreateWithNameAndOptions */

// Creates a new font reference from an ATS font reference.
//
// Deprecated: This function was deprecated in macOS 11.0.
//
// Added in macOS 10.5.
// Creates a new font reference from an ATS font reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCreateWithPlatformFont(_:_:_:_:)
func CTFontCreateWithPlatformFont(platformFont ATSFontRef, size float64, matrix unsafe.Pointer, attributes FontDescriptorRef) FontRef {
	return _CTFontCreateWithPlatformFont(platformFont, size, matrix, attributes)
}/* debug [functions.gen.go/function]: CTFontCreateWithPlatformFont */

// Returns a font reference for the given QuickDraw instance.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.5.
// Returns a font reference for the given QuickDraw instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCreateWithQuickdrawInstance(_:_:_:_:)
func CTFontCreateWithQuickdrawInstance(name unsafe.Pointer, identifier int16, style uint8, size float64) FontRef {
	return _CTFontCreateWithQuickdrawInstance(name, identifier, style, size)
}/* debug [functions.gen.go/function]: CTFontCreateWithQuickdrawInstance */

// Returns the value associated with an arbitrary attribute.
//
// Added in macOS 10.5.
// Returns the value associated with an arbitrary attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorCopyAttribute(_:_:)
func CTFontDescriptorCopyAttribute(descriptor FontDescriptorRef, attribute StringRef) TypeRef {
	return _CTFontDescriptorCopyAttribute(descriptor, attribute)
}/* debug [functions.gen.go/function]: CTFontDescriptorCopyAttribute */

// Returns the attributes dictionary of the font descriptor.
//
// Added in macOS 10.5.
// Returns the attributes dictionary of the font descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorCopyAttributes(_:)
func CTFontDescriptorCopyAttributes(descriptor FontDescriptorRef) DictionaryRef {
	return _CTFontDescriptorCopyAttributes(descriptor)
}/* debug [functions.gen.go/function]: CTFontDescriptorCopyAttributes */

// Returns a localized value for the requested attribute, if available.
//
// Added in macOS 10.5.
// Returns a localized value for the requested attribute, if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorCopyLocalizedAttribute(_:_:_:)
func CTFontDescriptorCopyLocalizedAttribute(descriptor FontDescriptorRef, attribute StringRef, language unsafe.Pointer) TypeRef {
	return _CTFontDescriptorCopyLocalizedAttribute(descriptor, attribute, language)
}/* debug [functions.gen.go/function]: CTFontDescriptorCopyLocalizedAttribute */

// Creates a copy of the original font descriptor with new attributes.
//
// Added in macOS 10.5.
// Creates a copy of the original font descriptor with new attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorCreateCopyWithAttributes(_:_:)
func CTFontDescriptorCreateCopyWithAttributes(original FontDescriptorRef, attributes DictionaryRef) FontDescriptorRef {
	return _CTFontDescriptorCreateCopyWithAttributes(original, attributes)
}/* debug [functions.gen.go/function]: CTFontDescriptorCreateCopyWithAttributes */

// Creates a copy of the font descriptor in the specified family based on the traits of the original.
//
// Added in macOS 10.9.
// Creates a copy of the font descriptor in the specified family based on the traits of the original.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorCreateCopyWithFamily(_:_:)
func CTFontDescriptorCreateCopyWithFamily(original FontDescriptorRef, family StringRef) FontDescriptorRef {
	return _CTFontDescriptorCreateCopyWithFamily(original, family)
}/* debug [functions.gen.go/function]: CTFontDescriptorCreateCopyWithFamily */

// Copies a font descriptor with new feature settings.
//
// Added in macOS 10.5.
// Copies a font descriptor with new feature settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorCreateCopyWithFeature(_:_:_:)
func CTFontDescriptorCreateCopyWithFeature(original FontDescriptorRef, featureTypeIdentifier NumberRef, featureSelectorIdentifier NumberRef) FontDescriptorRef {
	return _CTFontDescriptorCreateCopyWithFeature(original, featureTypeIdentifier, featureSelectorIdentifier)
}/* debug [functions.gen.go/function]: CTFontDescriptorCreateCopyWithFeature */

// Creates a copy of the font descriptor with the specified symbolic traits as the original.
//
// Added in macOS 10.9.
// Creates a copy of the font descriptor with the specified symbolic traits as the original.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorCreateCopyWithSymbolicTraits(_:_:_:)
func CTFontDescriptorCreateCopyWithSymbolicTraits(original FontDescriptorRef, symTraitValue FontSymbolicTraits, symTraitMask FontSymbolicTraits) FontDescriptorRef {
	return _CTFontDescriptorCreateCopyWithSymbolicTraits(original, symTraitValue, symTraitMask)
}/* debug [functions.gen.go/function]: CTFontDescriptorCreateCopyWithSymbolicTraits */

// Creates a copy of the original font descriptor with a new variation instance.
//
// Added in macOS 10.5.
// Creates a copy of the original font descriptor with a new variation instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorCreateCopyWithVariation(_:_:_:)
func CTFontDescriptorCreateCopyWithVariation(original FontDescriptorRef, variationIdentifier NumberRef, variationValue float64) FontDescriptorRef {
	return _CTFontDescriptorCreateCopyWithVariation(original, variationIdentifier, variationValue)
}/* debug [functions.gen.go/function]: CTFontDescriptorCreateCopyWithVariation */

// Returns the single preferred matching font descriptor based on the original descriptor and system precedence.
//
// Added in macOS 10.5.
// Returns the single preferred matching font descriptor based on the original descriptor and system precedence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorCreateMatchingFontDescriptor(_:_:)
func CTFontDescriptorCreateMatchingFontDescriptor(descriptor FontDescriptorRef, mandatoryAttributes SetRef) FontDescriptorRef {
	return _CTFontDescriptorCreateMatchingFontDescriptor(descriptor, mandatoryAttributes)
}/* debug [functions.gen.go/function]: CTFontDescriptorCreateMatchingFontDescriptor */

// Returns an array of normalized font descriptors matching the provided descriptor.
//
// Added in macOS 10.5.
// Returns an array of normalized font descriptors matching the provided descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorCreateMatchingFontDescriptors(_:_:)
func CTFontDescriptorCreateMatchingFontDescriptors(descriptor FontDescriptorRef, mandatoryAttributes SetRef) ArrayRef {
	return _CTFontDescriptorCreateMatchingFontDescriptors(descriptor, mandatoryAttributes)
}/* debug [functions.gen.go/function]: CTFontDescriptorCreateMatchingFontDescriptors */

// Creates a new font descriptor reference from a dictionary of attributes.
//
// Added in macOS 10.5.
// Creates a new font descriptor reference from a dictionary of attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorCreateWithAttributes(_:)
func CTFontDescriptorCreateWithAttributes(attributes DictionaryRef) FontDescriptorRef {
	return _CTFontDescriptorCreateWithAttributes(attributes)
}/* debug [functions.gen.go/function]: CTFontDescriptorCreateWithAttributes */

// Creates a new font descriptor with the provided PostScript name and size.
//
// Added in macOS 10.5.
// Creates a new font descriptor with the provided PostScript name and size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorCreateWithNameAndSize(_:_:)
func CTFontDescriptorCreateWithNameAndSize(name StringRef, size float64) FontDescriptorRef {
	return _CTFontDescriptorCreateWithNameAndSize(name, size)
}/* debug [functions.gen.go/function]: CTFontDescriptorCreateWithNameAndSize */

// Returns the type identifier for Core Text font descriptor references.
//
// Added in macOS 10.5.
// Returns the type identifier for Core Text font descriptor references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorGetTypeID()
func CTFontDescriptorGetTypeID() TypeID {
	return _CTFontDescriptorGetTypeID()
}/* debug [functions.gen.go/function]: CTFontDescriptorGetTypeID */

// Matches font descriptors and tracks progress with a progress handler.
//
// Added in macOS 10.9.
// Matches font descriptors and tracks progress with a progress handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorMatchFontDescriptorsWithProgressHandler(_:_:_:)
func CTFontDescriptorMatchFontDescriptorsWithProgressHandler(descriptors ArrayRef, mandatoryAttributes SetRef, progressBlock FontDescriptorProgressHandler) bool {
	return _CTFontDescriptorMatchFontDescriptorsWithProgressHandler(descriptors, mandatoryAttributes, progressBlock)
}/* debug [functions.gen.go/function]: CTFontDescriptorMatchFontDescriptorsWithProgressHandler */

// Renders the given glyphs of a font at the specified positions in the supplied graphics context.
//
// Added in macOS 10.7.
// Renders the given glyphs of a font at the specified positions in the supplied graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDrawGlyphs(_:_:_:_:_:)
func CTFontDrawGlyphs(font FontRef, glyphs unsafe.Pointer, positions unsafe.Pointer, count uintptr, context ContextRef) {
	_CTFontDrawGlyphs(font, glyphs, positions, count, context)
}/* debug [functions.gen.go/function]: CTFontDrawGlyphs */

// CTFontDrawImageFromAdaptiveImageProviderAtPoint is a CoreText function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDrawImageFromAdaptiveImageProviderAtPoint(_:_:_:_:)
func CTFontDrawImageFromAdaptiveImageProviderAtPoint(font FontRef, provider unsafe.Pointer, point corefoundation.CGPoint, context ContextRef) {
	_CTFontDrawImageFromAdaptiveImageProviderAtPoint(font, provider, point, context)
}/* debug [functions.gen.go/function]: CTFontDrawImageFromAdaptiveImageProviderAtPoint */

// Calculates the advances for an array of glyphs and returns the summed advance.
//
// Added in macOS 10.5.
// Calculates the advances for an array of glyphs and returns the summed advance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetAdvancesForGlyphs(_:_:_:_:_:)
func CTFontGetAdvancesForGlyphs(font FontRef, orientation FontOrientation, glyphs unsafe.Pointer, advances corefoundation.CGSize, count Index) float64 {
	return _CTFontGetAdvancesForGlyphs(font, orientation, glyphs, advances, count)
}/* debug [functions.gen.go/function]: CTFontGetAdvancesForGlyphs */

// Returns the scaled font-ascent metric of the given font.
//
// Added in macOS 10.5.
// Returns the scaled font-ascent metric of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetAscent(_:)
func CTFontGetAscent(font FontRef) float64 {
	return _CTFontGetAscent(font)
}/* debug [functions.gen.go/function]: CTFontGetAscent */

// Returns the scaled bounding box of the given font.
//
// Added in macOS 10.5.
// Returns the scaled bounding box of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetBoundingBox(_:)
func CTFontGetBoundingBox(font FontRef) corefoundation.CGRect {
	return _CTFontGetBoundingBox(font)
}/* debug [functions.gen.go/function]: CTFontGetBoundingBox */

// Calculates the bounding rects for an array of glyphs and returns the overall bounding rectangle for the glyph run.
//
// Added in macOS 10.5.
// Calculates the bounding rects for an array of glyphs and returns the overall bounding rectangle for the glyph run.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetBoundingRectsForGlyphs(_:_:_:_:_:)
func CTFontGetBoundingRectsForGlyphs(font FontRef, orientation FontOrientation, glyphs unsafe.Pointer, boundingRects corefoundation.CGRect, count Index) corefoundation.CGRect {
	return _CTFontGetBoundingRectsForGlyphs(font, orientation, glyphs, boundingRects, count)
}/* debug [functions.gen.go/function]: CTFontGetBoundingRectsForGlyphs */

// Returns the cap-height metric of the given font.
//
// Added in macOS 10.5.
// Returns the cap-height metric of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetCapHeight(_:)
func CTFontGetCapHeight(font FontRef) float64 {
	return _CTFontGetCapHeight(font)
}/* debug [functions.gen.go/function]: CTFontGetCapHeight */

// Returns the scaled font-descent metric of the given font.
//
// Added in macOS 10.5.
// Returns the scaled font-descent metric of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetDescent(_:)
func CTFontGetDescent(font FontRef) float64 {
	return _CTFontGetDescent(font)
}/* debug [functions.gen.go/function]: CTFontGetDescent */

// Returns the number of glyphs of the given font.
//
// Added in macOS 10.5.
// Returns the number of glyphs of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetGlyphCount(_:)
func CTFontGetGlyphCount(font FontRef) Index {
	return _CTFontGetGlyphCount(font)
}/* debug [functions.gen.go/function]: CTFontGetGlyphCount */

// Performs basic character-to-glyph mapping.
//
// Added in macOS 10.5.
// Performs basic character-to-glyph mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetGlyphsForCharacters(_:_:_:_:)
func CTFontGetGlyphsForCharacters(font FontRef, characters unsafe.Pointer, glyphs Glyph, count Index) bool {
	return _CTFontGetGlyphsForCharacters(font, characters, glyphs, count)
}/* debug [functions.gen.go/function]: CTFontGetGlyphsForCharacters */

// Returns the glyph for the specified name.
//
// Added in macOS 10.5.
// Returns the glyph for the specified name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetGlyphWithName(_:_:)
func CTFontGetGlyphWithName(font FontRef, glyphName StringRef) Glyph {
	return _CTFontGetGlyphWithName(font, glyphName)
}/* debug [functions.gen.go/function]: CTFontGetGlyphWithName */

// Returns the scaled font-leading metric of the given font.
//
// Added in macOS 10.5.
// Returns the scaled font-leading metric of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetLeading(_:)
func CTFontGetLeading(font FontRef) float64 {
	return _CTFontGetLeading(font)
}/* debug [functions.gen.go/function]: CTFontGetLeading */

// Returns caret positions within a glyph.
//
// Added in macOS 10.5.
// Returns caret positions within a glyph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetLigatureCaretPositions(_:_:_:_:)
func CTFontGetLigatureCaretPositions(font FontRef, glyph Glyph, positions float64, maxPositions Index) Index {
	return _CTFontGetLigatureCaretPositions(font, glyph, positions, maxPositions)
}/* debug [functions.gen.go/function]: CTFontGetLigatureCaretPositions */

// Returns the transformation matrix of the given font.
//
// Added in macOS 10.5.
// Returns the transformation matrix of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetMatrix(_:)
func CTFontGetMatrix(font FontRef) corefoundation.CGAffineTransform {
	return _CTFontGetMatrix(font)
}/* debug [functions.gen.go/function]: CTFontGetMatrix */

// Calculates the optical bounds for an array of glyphs and returns the overall optical bounds for the run.
//
// Added in macOS 10.8.
// Calculates the optical bounds for an array of glyphs and returns the overall optical bounds for the run.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetOpticalBoundsForGlyphs(_:_:_:_:_:)
func CTFontGetOpticalBoundsForGlyphs(font FontRef, glyphs unsafe.Pointer, boundingRects corefoundation.CGRect, count Index, options OptionFlags) corefoundation.CGRect {
	return _CTFontGetOpticalBoundsForGlyphs(font, glyphs, boundingRects, count, options)
}/* debug [functions.gen.go/function]: CTFontGetOpticalBoundsForGlyphs */

// Returns an ATS font reference and attributes.
//
// Deprecated: This function was deprecated in macOS 11.0.
//
// Added in macOS 10.5.
// Returns an ATS font reference and attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetPlatformFont(_:_:)
func CTFontGetPlatformFont(font FontRef, attributes unsafe.Pointer) ATSFontRef {
	return _CTFontGetPlatformFont(font, attributes)
}/* debug [functions.gen.go/function]: CTFontGetPlatformFont */

// Returns the point size of the given font.
//
// Added in macOS 10.5.
// Returns the point size of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetSize(_:)
func CTFontGetSize(font FontRef) float64 {
	return _CTFontGetSize(font)
}/* debug [functions.gen.go/function]: CTFontGetSize */

// Returns the slant angle of the given font.
//
// Added in macOS 10.5.
// Returns the slant angle of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetSlantAngle(_:)
func CTFontGetSlantAngle(font FontRef) float64 {
	return _CTFontGetSlantAngle(font)
}/* debug [functions.gen.go/function]: CTFontGetSlantAngle */

// Returns the best string encoding for legacy format support.
//
// Added in macOS 10.5.
// Returns the best string encoding for legacy format support.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetStringEncoding(_:)
func CTFontGetStringEncoding(font FontRef) StringEncoding {
	return _CTFontGetStringEncoding(font)
}/* debug [functions.gen.go/function]: CTFontGetStringEncoding */

// Returns the symbolic traits of the given font.
//
// Added in macOS 10.5.
// Returns the symbolic traits of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetSymbolicTraits(_:)
func CTFontGetSymbolicTraits(font FontRef) FontSymbolicTraits {
	return _CTFontGetSymbolicTraits(font)
}/* debug [functions.gen.go/function]: CTFontGetSymbolicTraits */

// Returns the type identifier for Core Text font references.
//
// Added in macOS 10.5.
// Returns the type identifier for Core Text font references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetTypeID()
func CTFontGetTypeID() TypeID {
	return _CTFontGetTypeID()
}/* debug [functions.gen.go/function]: CTFontGetTypeID */

// CTFontGetTypographicBoundsForAdaptiveImageProvider is a CoreText function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetTypographicBoundsForAdaptiveImageProvider(_:_:)
func CTFontGetTypographicBoundsForAdaptiveImageProvider(font FontRef, provider unsafe.Pointer) corefoundation.CGRect {
	return _CTFontGetTypographicBoundsForAdaptiveImageProvider(font, provider)
}/* debug [functions.gen.go/function]: CTFontGetTypographicBoundsForAdaptiveImageProvider */

// Returns the scaled underline position of the given font.
//
// Added in macOS 10.5.
// Returns the scaled underline position of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetUnderlinePosition(_:)
func CTFontGetUnderlinePosition(font FontRef) float64 {
	return _CTFontGetUnderlinePosition(font)
}/* debug [functions.gen.go/function]: CTFontGetUnderlinePosition */

// Returns the scaled underline-thickness metric of the given font.
//
// Added in macOS 10.5.
// Returns the scaled underline-thickness metric of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetUnderlineThickness(_:)
func CTFontGetUnderlineThickness(font FontRef) float64 {
	return _CTFontGetUnderlineThickness(font)
}/* debug [functions.gen.go/function]: CTFontGetUnderlineThickness */

// Returns the units-per-em metric of the given font.
//
// Added in macOS 10.5.
// Returns the units-per-em metric of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetUnitsPerEm(_:)
func CTFontGetUnitsPerEm(font FontRef) unsafe.Pointer {
	return _CTFontGetUnitsPerEm(font)
}/* debug [functions.gen.go/function]: CTFontGetUnitsPerEm */

// Calculates the offset from the default (horizontal) origin to the vertical origin for an array of glyphs.
//
// Added in macOS 10.5.
// Calculates the offset from the default (horizontal) origin to the vertical origin for an array of glyphs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetVerticalTranslationsForGlyphs(_:_:_:_:)
func CTFontGetVerticalTranslationsForGlyphs(font FontRef, glyphs unsafe.Pointer, translations corefoundation.CGSize, count Index) {
	_CTFontGetVerticalTranslationsForGlyphs(font, glyphs, translations, count)
}/* debug [functions.gen.go/function]: CTFontGetVerticalTranslationsForGlyphs */

// Returns the x-height metric of the given font.
//
// Added in macOS 10.5.
// Returns the x-height metric of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetXHeight(_:)
func CTFontGetXHeight(font FontRef) float64 {
	return _CTFontGetXHeight(font)
}/* debug [functions.gen.go/function]: CTFontGetXHeight */

// CTFontHasTable is a CoreText function.
//
// Added in macOS 10.15.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontHasTable(_:_:)
func CTFontHasTable(font FontRef, tag FontTableTag) bool {
	return _CTFontHasTable(font, tag)
}/* debug [functions.gen.go/function]: CTFontHasTable */

// A comparator function to compare font family names and sort them according to Apple guidelines.
//
// Added in macOS 10.6.
// A comparator function to compare font family names and sort them according to Apple guidelines.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerCompareFontFamilyNames(_:_:_:)
func CTFontManagerCompareFontFamilyNames(family1 unsafe.Pointer, family2 unsafe.Pointer, context unsafe.Pointer) ComparisonResult {
	return _CTFontManagerCompareFontFamilyNames(family1, family2, context)
}/* debug [functions.gen.go/function]: CTFontManagerCompareFontFamilyNames */

// Returns an array of visible font family names sorted for user interface display.
//
// Added in macOS 10.6.
// Returns an array of visible font family names sorted for user interface display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerCopyAvailableFontFamilyNames()
func CTFontManagerCopyAvailableFontFamilyNames() ArrayRef {
	return _CTFontManagerCopyAvailableFontFamilyNames()
}/* debug [functions.gen.go/function]: CTFontManagerCopyAvailableFontFamilyNames */

// Returns an array of font URLs.
//
// Added in macOS 10.6.
// Returns an array of font URLs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerCopyAvailableFontURLs()
func CTFontManagerCopyAvailableFontURLs() ArrayRef {
	return _CTFontManagerCopyAvailableFontURLs()
}/* debug [functions.gen.go/function]: CTFontManagerCopyAvailableFontURLs */

// Returns an array of unique PostScript font names for the fonts.
//
// Added in macOS 10.6.
// Returns an array of unique PostScript font names for the fonts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerCopyAvailablePostScriptNames()
func CTFontManagerCopyAvailablePostScriptNames() ArrayRef {
	return _CTFontManagerCopyAvailablePostScriptNames()
}/* debug [functions.gen.go/function]: CTFontManagerCopyAvailablePostScriptNames */

// Retrieves the font descriptors that were registered with the font manager.

// Retrieves the font descriptors that were registered with the font manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerCopyRegisteredFontDescriptors(_:_:)
func CTFontManagerCopyRegisteredFontDescriptors(scope FontManagerScope, enabled bool) ArrayRef {
	return _CTFontManagerCopyRegisteredFontDescriptors(scope, enabled)
}/* debug [functions.gen.go/function]: CTFontManagerCopyRegisteredFontDescriptors */

// Creates a font descriptor representing the font in the supplied data.
//
// Added in macOS 10.7.
// Creates a font descriptor representing the font in the supplied data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerCreateFontDescriptorFromData(_:)
func CTFontManagerCreateFontDescriptorFromData(data DataRef) FontDescriptorRef {
	return _CTFontManagerCreateFontDescriptorFromData(data)
}/* debug [functions.gen.go/function]: CTFontManagerCreateFontDescriptorFromData */

// Creates an array of font descriptors for the fonts in the supplied data.
//
// Added in macOS 10.13.
// Creates an array of font descriptors for the fonts in the supplied data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerCreateFontDescriptorsFromData(_:)
func CTFontManagerCreateFontDescriptorsFromData(data DataRef) ArrayRef {
	return _CTFontManagerCreateFontDescriptorsFromData(data)
}/* debug [functions.gen.go/function]: CTFontManagerCreateFontDescriptorsFromData */

// Returns an array of font descriptors representing each of the fonts in the specified URL.
//
// Added in macOS 10.6.
// Returns an array of font descriptors representing each of the fonts in the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerCreateFontDescriptorsFromURL(_:)
func CTFontManagerCreateFontDescriptorsFromURL(fileURL URLRef) ArrayRef {
	return _CTFontManagerCreateFontDescriptorsFromURL(fileURL)
}/* debug [functions.gen.go/function]: CTFontManagerCreateFontDescriptorsFromURL */

// Creates a reference to a run loop source used to convey font requests from the Font Manager.
//
// Deprecated: This function was deprecated in macOS 11.0.
//
// Added in macOS 10.6.
// Creates a reference to a run loop source used to convey font requests from the Font Manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerCreateFontRequestRunLoopSource(_:_:)
func CTFontManagerCreateFontRequestRunLoopSource(sourceOrder Index, createMatchesCallback ArrayRef) RunLoopSourceRef {
	return _CTFontManagerCreateFontRequestRunLoopSource(sourceOrder, createMatchesCallback)
}/* debug [functions.gen.go/function]: CTFontManagerCreateFontRequestRunLoopSource */

// Enables or disables the matching font descriptors for font descriptor matching.
//
// Added in macOS 10.6.
// Enables or disables the matching font descriptors for font descriptor matching.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerEnableFontDescriptors(_:_:)
func CTFontManagerEnableFontDescriptors(descriptors ArrayRef, enable bool) {
	_CTFontManagerEnableFontDescriptors(descriptors, enable)
}/* debug [functions.gen.go/function]: CTFontManagerEnableFontDescriptors */

// Gets the auto-activation setting for the specified bundle identifier.
//
// Added in macOS 10.6.
// Gets the auto-activation setting for the specified bundle identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerGetAutoActivationSetting(_:)
func CTFontManagerGetAutoActivationSetting(bundleIdentifier StringRef) FontManagerAutoActivationSetting {
	return _CTFontManagerGetAutoActivationSetting(bundleIdentifier)
}/* debug [functions.gen.go/function]: CTFontManagerGetAutoActivationSetting */

// Returns the registration scope of the specified URL.
//
// Added in macOS 10.6.
// Returns the registration scope of the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerGetScopeForURL(_:)
func CTFontManagerGetScopeForURL(fontURL URLRef) FontManagerScope {
	return _CTFontManagerGetScopeForURL(fontURL)
}/* debug [functions.gen.go/function]: CTFontManagerGetScopeForURL */

// Determines whether a file is in a supported font format.
//
// Added in macOS 10.6.
// Determines whether a file is in a supported font format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerIsSupportedFont(_:)
func CTFontManagerIsSupportedFont(fontURL URLRef) bool {
	return _CTFontManagerIsSupportedFont(fontURL)
}/* debug [functions.gen.go/function]: CTFontManagerIsSupportedFont */

// Registers font descriptors with the font manager.
//
// Added in macOS 10.15.
// Registers font descriptors with the font manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerRegisterFontDescriptors(_:_:_:_:)
func CTFontManagerRegisterFontDescriptors(fontDescriptors ArrayRef, scope FontManagerScope, enabled bool, registrationHandler bool) {
	_CTFontManagerRegisterFontDescriptors(fontDescriptors, scope, enabled, registrationHandler)
}/* debug [functions.gen.go/function]: CTFontManagerRegisterFontDescriptors */

// Registers fonts from the specified font URL with the Font Manager. Registered fonts are discoverable through font descriptor matching.
//
// Added in macOS 10.6.
// Registers fonts from the specified font URL with the Font Manager. Registered fonts are discoverable through font descriptor matching.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerRegisterFontsForURL(_:_:_:)
func CTFontManagerRegisterFontsForURL(fontURL URLRef, scope FontManagerScope, error_ unsafe.Pointer) bool {
	return _CTFontManagerRegisterFontsForURL(fontURL, scope, error_)
}/* debug [functions.gen.go/function]: CTFontManagerRegisterFontsForURL */

// Registers fonts from the specified array of font URLs with the Font Manager. Registered fonts are discoverable through font descriptor matching.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.6.
// Registers fonts from the specified array of font URLs with the Font Manager. Registered fonts are discoverable through font descriptor matching.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerRegisterFontsForURLs(_:_:_:)
func CTFontManagerRegisterFontsForURLs(fontURLs ArrayRef, scope FontManagerScope, errors unsafe.Pointer) bool {
	return _CTFontManagerRegisterFontsForURLs(fontURLs, scope, errors)
}/* debug [functions.gen.go/function]: CTFontManagerRegisterFontsForURLs */

// Registers named font assets in the specified bundle with the font manager.

// Registers named font assets in the specified bundle with the font manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerRegisterFontsWithAssetNames(_:_:_:_:_:)
func CTFontManagerRegisterFontsWithAssetNames(fontAssetNames ArrayRef, bundle BundleRef, scope FontManagerScope, enabled bool, registrationHandler bool) {
	_CTFontManagerRegisterFontsWithAssetNames(fontAssetNames, bundle, scope, enabled, registrationHandler)
}/* debug [functions.gen.go/function]: CTFontManagerRegisterFontsWithAssetNames */

// Registers fonts from the specified font URLs with the font manager.
//
// Added in macOS 10.15.
// Registers fonts from the specified font URLs with the font manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerRegisterFontURLs(_:_:_:_:)
func CTFontManagerRegisterFontURLs(fontURLs ArrayRef, scope FontManagerScope, enabled bool, registrationHandler bool) {
	_CTFontManagerRegisterFontURLs(fontURLs, scope, enabled, registrationHandler)
}/* debug [functions.gen.go/function]: CTFontManagerRegisterFontURLs */

// Registers the specified graphics font with the font manager.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.8.
// Registers the specified graphics font with the font manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerRegisterGraphicsFont(_:_:)
func CTFontManagerRegisterGraphicsFont(font FontRef, error_ unsafe.Pointer) bool {
	return _CTFontManagerRegisterGraphicsFont(font, error_)
}/* debug [functions.gen.go/function]: CTFontManagerRegisterGraphicsFont */

// Resolves font descriptors specified on input.

// Resolves font descriptors specified on input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerRequestFonts(_:_:)
func CTFontManagerRequestFonts(fontDescriptors ArrayRef) {
	_CTFontManagerRequestFonts(fontDescriptors)
}/* debug [functions.gen.go/function]: CTFontManagerRequestFonts */

// Sets the auto-activation setting for the specified bundle identifier.
//
// Added in macOS 10.6.
// Sets the auto-activation setting for the specified bundle identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerSetAutoActivationSetting(_:_:)
func CTFontManagerSetAutoActivationSetting(bundleIdentifier StringRef, setting FontManagerAutoActivationSetting) {
	_CTFontManagerSetAutoActivationSetting(bundleIdentifier, setting)
}/* debug [functions.gen.go/function]: CTFontManagerSetAutoActivationSetting */

// Unregisters font descriptors with the font manager.
//
// Added in macOS 10.15.
// Unregisters font descriptors with the font manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerUnregisterFontDescriptors(_:_:_:)
func CTFontManagerUnregisterFontDescriptors(fontDescriptors ArrayRef, scope FontManagerScope, registrationHandler bool) {
	_CTFontManagerUnregisterFontDescriptors(fontDescriptors, scope, registrationHandler)
}/* debug [functions.gen.go/function]: CTFontManagerUnregisterFontDescriptors */

// Unregisters fonts from the specified font URL with the Font Manager. Unregistered fonts are no longer discoverable through font descriptor matching.
//
// Added in macOS 10.6.
// Unregisters fonts from the specified font URL with the Font Manager. Unregistered fonts are no longer discoverable through font descriptor matching.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerUnregisterFontsForURL(_:_:_:)
func CTFontManagerUnregisterFontsForURL(fontURL URLRef, scope FontManagerScope, error_ unsafe.Pointer) bool {
	return _CTFontManagerUnregisterFontsForURL(fontURL, scope, error_)
}/* debug [functions.gen.go/function]: CTFontManagerUnregisterFontsForURL */

// Unregisters fonts from the specified array of font URLs with the Font Manager. Unregistered fonts are no longer discoverable through font descriptor matching.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.6.
// Unregisters fonts from the specified array of font URLs with the Font Manager. Unregistered fonts are no longer discoverable through font descriptor matching.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerUnregisterFontsForURLs(_:_:_:)
func CTFontManagerUnregisterFontsForURLs(fontURLs ArrayRef, scope FontManagerScope, errors unsafe.Pointer) bool {
	return _CTFontManagerUnregisterFontsForURLs(fontURLs, scope, errors)
}/* debug [functions.gen.go/function]: CTFontManagerUnregisterFontsForURLs */

// Unregisters fonts from the specified font URLs with the font manager.
//
// Added in macOS 10.15.
// Unregisters fonts from the specified font URLs with the font manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerUnregisterFontURLs(_:_:_:)
func CTFontManagerUnregisterFontURLs(fontURLs ArrayRef, scope FontManagerScope, registrationHandler bool) {
	_CTFontManagerUnregisterFontURLs(fontURLs, scope, registrationHandler)
}/* debug [functions.gen.go/function]: CTFontManagerUnregisterFontURLs */

// Unregisters the specified graphics font with the font manager.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.8.
// Unregisters the specified graphics font with the font manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerUnregisterGraphicsFont(_:_:)
func CTFontManagerUnregisterGraphicsFont(font FontRef, error_ unsafe.Pointer) bool {
	return _CTFontManagerUnregisterGraphicsFont(font, error_)
}/* debug [functions.gen.go/function]: CTFontManagerUnregisterGraphicsFont */

// Draws an entire frame into a context.
//
// Added in macOS 10.5.
// Draws an entire frame into a context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFrameDraw(_:_:)
func CTFrameDraw(frame FrameRef, context ContextRef) {
	_CTFrameDraw(frame, context)
}/* debug [functions.gen.go/function]: CTFrameDraw */

// Returns the frame attributes used to create the frame.
//
// Added in macOS 10.5.
// Returns the frame attributes used to create the frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFrameGetFrameAttributes(_:)
func CTFrameGetFrameAttributes(frame FrameRef) DictionaryRef {
	return _CTFrameGetFrameAttributes(frame)
}/* debug [functions.gen.go/function]: CTFrameGetFrameAttributes */

// Copies a range of line origins for a frame.
//
// Added in macOS 10.5.
// Copies a range of line origins for a frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFrameGetLineOrigins(_:_:_:)
func CTFrameGetLineOrigins(frame FrameRef, range_ corefoundation.Range, origins corefoundation.CGPoint) {
	_CTFrameGetLineOrigins(frame, range_, origins)
}/* debug [functions.gen.go/function]: CTFrameGetLineOrigins */

// Returns an array of lines stored in the frame.
//
// Added in macOS 10.5.
// Returns an array of lines stored in the frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFrameGetLines(_:)
func CTFrameGetLines(frame FrameRef) ArrayRef {
	return _CTFrameGetLines(frame)
}/* debug [functions.gen.go/function]: CTFrameGetLines */

// Returns the path used to create the frame.
//
// Added in macOS 10.5.
// Returns the path used to create the frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFrameGetPath(_:)
func CTFrameGetPath(frame FrameRef) PathRef {
	return _CTFrameGetPath(frame)
}/* debug [functions.gen.go/function]: CTFrameGetPath */

// Returns the range of characters originally requested to fill the frame.
//
// Added in macOS 10.5.
// Returns the range of characters originally requested to fill the frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFrameGetStringRange(_:)
func CTFrameGetStringRange(frame FrameRef) corefoundation.Range {
	return _CTFrameGetStringRange(frame)
}/* debug [functions.gen.go/function]: CTFrameGetStringRange */

// Returns the type identifier for the CTFrame opaque type.
//
// Added in macOS 10.5.
// Returns the type identifier for the CTFrame opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFrameGetTypeID()
func CTFrameGetTypeID() TypeID {
	return _CTFrameGetTypeID()
}/* debug [functions.gen.go/function]: CTFrameGetTypeID */

// Returns the range of characters that actually fit in the frame.
//
// Added in macOS 10.5.
// Returns the range of characters that actually fit in the frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFrameGetVisibleStringRange(_:)
func CTFrameGetVisibleStringRange(frame FrameRef) corefoundation.Range {
	return _CTFrameGetVisibleStringRange(frame)
}/* debug [functions.gen.go/function]: CTFrameGetVisibleStringRange */

// Creates an immutable frame using a framesetter.
//
// Added in macOS 10.5.
// Creates an immutable frame using a framesetter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFramesetterCreateFrame(_:_:_:_:)
func CTFramesetterCreateFrame(framesetter FramesetterRef, stringRange corefoundation.Range, path PathRef, frameAttributes DictionaryRef) FrameRef {
	return _CTFramesetterCreateFrame(framesetter, stringRange, path, frameAttributes)
}/* debug [functions.gen.go/function]: CTFramesetterCreateFrame */

// Creates an immutable framesetter object from an attributed string.
//
// Added in macOS 10.5.
// Creates an immutable framesetter object from an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFramesetterCreateWithAttributedString(_:)
func CTFramesetterCreateWithAttributedString(attrString AttributedStringRef) FramesetterRef {
	return _CTFramesetterCreateWithAttributedString(attrString)
}/* debug [functions.gen.go/function]: CTFramesetterCreateWithAttributedString */

// Creates a framesetter directly from a typesetter.
//
// Added in macOS 10.14.
// Creates a framesetter directly from a typesetter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFramesetterCreateWithTypesetter(_:)
func CTFramesetterCreateWithTypesetter(typesetter TypesetterRef) FramesetterRef {
	return _CTFramesetterCreateWithTypesetter(typesetter)
}/* debug [functions.gen.go/function]: CTFramesetterCreateWithTypesetter */

// Returns the Core Foundation type identifier of the framesetter object.
//
// Added in macOS 10.5.
// Returns the Core Foundation type identifier of the framesetter object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFramesetterGetTypeID()
func CTFramesetterGetTypeID() TypeID {
	return _CTFramesetterGetTypeID()
}/* debug [functions.gen.go/function]: CTFramesetterGetTypeID */

// Returns the typesetter object being used by the framesetter.
//
// Added in macOS 10.5.
// Returns the typesetter object being used by the framesetter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFramesetterGetTypesetter(_:)
func CTFramesetterGetTypesetter(framesetter FramesetterRef) TypesetterRef {
	return _CTFramesetterGetTypesetter(framesetter)
}/* debug [functions.gen.go/function]: CTFramesetterGetTypesetter */

// Determines the frame size needed for a string range.
//
// Added in macOS 10.5.
// Determines the frame size needed for a string range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFramesetterSuggestFrameSizeWithConstraints(_:_:_:_:_:)
func CTFramesetterSuggestFrameSizeWithConstraints(framesetter FramesetterRef, stringRange corefoundation.Range, frameAttributes DictionaryRef, constraints corefoundation.CGSize, fitRange unsafe.Pointer) corefoundation.CGSize {
	return _CTFramesetterSuggestFrameSizeWithConstraints(framesetter, stringRange, frameAttributes, constraints, fitRange)
}/* debug [functions.gen.go/function]: CTFramesetterSuggestFrameSizeWithConstraints */

// Returns the version of the Core Text framework.
//
// Deprecated: This function was deprecated in macOS 11.0.
//
// Added in macOS 10.5.
// Returns the version of the Core Text framework.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTGetCoreTextVersion()
func CTGetCoreTextVersion() uint32 {
	return _CTGetCoreTextVersion()
}/* debug [functions.gen.go/function]: CTGetCoreTextVersion */

// Creates an immutable glyph info object with a character identifier.
//
// Added in macOS 10.5.
// Creates an immutable glyph info object with a character identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTGlyphInfoCreateWithCharacterIdentifier(_:_:_:)
func CTGlyphInfoCreateWithCharacterIdentifier(cid FontIndex, collection CharacterCollection, baseString StringRef) GlyphInfoRef {
	return _CTGlyphInfoCreateWithCharacterIdentifier(cid, collection, baseString)
}/* debug [functions.gen.go/function]: CTGlyphInfoCreateWithCharacterIdentifier */

// Creates an immutable glyph info object with a glyph index.
//
// Added in macOS 10.5.
// Creates an immutable glyph info object with a glyph index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTGlyphInfoCreateWithGlyph(_:_:_:)
func CTGlyphInfoCreateWithGlyph(glyph Glyph, font FontRef, baseString StringRef) GlyphInfoRef {
	return _CTGlyphInfoCreateWithGlyph(glyph, font, baseString)
}/* debug [functions.gen.go/function]: CTGlyphInfoCreateWithGlyph */

// Creates an immutable glyph info object with a glyph name.
//
// Added in macOS 10.5.
// Creates an immutable glyph info object with a glyph name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTGlyphInfoCreateWithGlyphName(_:_:_:)
func CTGlyphInfoCreateWithGlyphName(glyphName StringRef, font FontRef, baseString StringRef) GlyphInfoRef {
	return _CTGlyphInfoCreateWithGlyphName(glyphName, font, baseString)
}/* debug [functions.gen.go/function]: CTGlyphInfoCreateWithGlyphName */

// Gets the character collection for a glyph info object.
//
// Added in macOS 10.5.
// Gets the character collection for a glyph info object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTGlyphInfoGetCharacterCollection(_:)
func CTGlyphInfoGetCharacterCollection(glyphInfo GlyphInfoRef) CharacterCollection {
	return _CTGlyphInfoGetCharacterCollection(glyphInfo)
}/* debug [functions.gen.go/function]: CTGlyphInfoGetCharacterCollection */

// Gets the character identifier for a glyph info object.
//
// Added in macOS 10.5.
// Gets the character identifier for a glyph info object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTGlyphInfoGetCharacterIdentifier(_:)
func CTGlyphInfoGetCharacterIdentifier(glyphInfo GlyphInfoRef) FontIndex {
	return _CTGlyphInfoGetCharacterIdentifier(glyphInfo)
}/* debug [functions.gen.go/function]: CTGlyphInfoGetCharacterIdentifier */

// Retrieves the glyph for a glyph info, if that object exists.
//
// Added in macOS 10.15.
// Retrieves the glyph for a glyph info, if that object exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTGlyphInfoGetGlyph(_:)
func CTGlyphInfoGetGlyph(glyphInfo GlyphInfoRef) Glyph {
	return _CTGlyphInfoGetGlyph(glyphInfo)
}/* debug [functions.gen.go/function]: CTGlyphInfoGetGlyph */

// Retrieves the glyph name for a glyph info object, if that object exists.
//
// Added in macOS 10.5.
// Retrieves the glyph name for a glyph info object, if that object exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTGlyphInfoGetGlyphName(_:)
func CTGlyphInfoGetGlyphName(glyphInfo GlyphInfoRef) StringRef {
	return _CTGlyphInfoGetGlyphName(glyphInfo)
}/* debug [functions.gen.go/function]: CTGlyphInfoGetGlyphName */

// Returns the Core Foundation type identifier of the glyph info object
//
// Added in macOS 10.5.
// Returns the Core Foundation type identifier of the glyph info object
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTGlyphInfoGetTypeID()
func CTGlyphInfoGetTypeID() TypeID {
	return _CTGlyphInfoGetTypeID()
}/* debug [functions.gen.go/function]: CTGlyphInfoGetTypeID */

// Creates a justified line from an existing line.
//
// Added in macOS 10.5.
// Creates a justified line from an existing line.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineCreateJustifiedLine(_:_:_:)
func CTLineCreateJustifiedLine(line LineRef, justificationFactor float64, justificationWidth float64) LineRef {
	return _CTLineCreateJustifiedLine(line, justificationFactor, justificationWidth)
}/* debug [functions.gen.go/function]: CTLineCreateJustifiedLine */

// Creates a truncated line from an existing line.
//
// Added in macOS 10.5.
// Creates a truncated line from an existing line.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineCreateTruncatedLine(_:_:_:_:)
func CTLineCreateTruncatedLine(line LineRef, width float64, truncationType LineTruncationType, truncationToken LineRef) LineRef {
	return _CTLineCreateTruncatedLine(line, width, truncationType, truncationToken)
}/* debug [functions.gen.go/function]: CTLineCreateTruncatedLine */

// Creates a single immutable line object from an attributed string.
//
// Added in macOS 10.5.
// Creates a single immutable line object from an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineCreateWithAttributedString(_:)
func CTLineCreateWithAttributedString(attrString AttributedStringRef) LineRef {
	return _CTLineCreateWithAttributedString(attrString)
}/* debug [functions.gen.go/function]: CTLineCreateWithAttributedString */

// Draws a complete line.
//
// Added in macOS 10.5.
// Draws a complete line.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineDraw(_:_:)
func CTLineDraw(line LineRef, context ContextRef) {
	_CTLineDraw(line, context)
}/* debug [functions.gen.go/function]: CTLineDraw */

// Enumerates caret offsets for characters in a line.
//
// Added in macOS 10.11.
// Enumerates caret offsets for characters in a line.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineEnumerateCaretOffsets(_:_:)
func CTLineEnumerateCaretOffsets(line LineRef) {
	_CTLineEnumerateCaretOffsets(line)
}/* debug [functions.gen.go/function]: CTLineEnumerateCaretOffsets */

// Calculates the bounds for a line.
//
// Added in macOS 10.8.
// Calculates the bounds for a line.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineGetBoundsWithOptions(_:_:)
func CTLineGetBoundsWithOptions(line LineRef, options LineBoundsOptions) corefoundation.CGRect {
	return _CTLineGetBoundsWithOptions(line, options)
}/* debug [functions.gen.go/function]: CTLineGetBoundsWithOptions */

// Returns the total glyph count for the line object.
//
// Added in macOS 10.5.
// Returns the total glyph count for the line object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineGetGlyphCount(_:)
func CTLineGetGlyphCount(line LineRef) Index {
	return _CTLineGetGlyphCount(line)
}/* debug [functions.gen.go/function]: CTLineGetGlyphCount */

// Returns the array of glyph runs that make up the line object.
//
// Added in macOS 10.5.
// Returns the array of glyph runs that make up the line object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineGetGlyphRuns(_:)
func CTLineGetGlyphRuns(line LineRef) ArrayRef {
	return _CTLineGetGlyphRuns(line)
}/* debug [functions.gen.go/function]: CTLineGetGlyphRuns */

// Calculates the image bounds for a line.
//
// Added in macOS 10.5.
// Calculates the image bounds for a line.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineGetImageBounds(_:_:)
func CTLineGetImageBounds(line LineRef, context ContextRef) corefoundation.CGRect {
	return _CTLineGetImageBounds(line, context)
}/* debug [functions.gen.go/function]: CTLineGetImageBounds */

// Determines the graphical offset or offsets for a string index.
//
// Added in macOS 10.5.
// Determines the graphical offset or offsets for a string index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineGetOffsetForStringIndex(_:_:_:)
func CTLineGetOffsetForStringIndex(line LineRef, charIndex Index, secondaryOffset []float64) float64 {
	return _CTLineGetOffsetForStringIndex(line, charIndex, secondaryOffset)
}/* debug [functions.gen.go/function]: CTLineGetOffsetForStringIndex */

// Gets the pen offset required to draw flush text.
//
// Added in macOS 10.5.
// Gets the pen offset required to draw flush text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineGetPenOffsetForFlush(_:_:_:)
func CTLineGetPenOffsetForFlush(line LineRef, flushFactor float64, flushWidth float64) float64 {
	return _CTLineGetPenOffsetForFlush(line, flushFactor, flushWidth)
}/* debug [functions.gen.go/function]: CTLineGetPenOffsetForFlush */

// Performs hit testing.
//
// Added in macOS 10.5.
// Performs hit testing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineGetStringIndexForPosition(_:_:)
func CTLineGetStringIndexForPosition(line LineRef, position corefoundation.CGPoint) Index {
	return _CTLineGetStringIndexForPosition(line, position)
}/* debug [functions.gen.go/function]: CTLineGetStringIndexForPosition */

// Gets the range of characters that originally spawned the glyphs in the line.
//
// Added in macOS 10.5.
// Gets the range of characters that originally spawned the glyphs in the line.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineGetStringRange(_:)
func CTLineGetStringRange(line LineRef) corefoundation.Range {
	return _CTLineGetStringRange(line)
}/* debug [functions.gen.go/function]: CTLineGetStringRange */

// Returns the trailing whitespace width for a line.
//
// Added in macOS 10.5.
// Returns the trailing whitespace width for a line.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineGetTrailingWhitespaceWidth(_:)
func CTLineGetTrailingWhitespaceWidth(line LineRef) float64 {
	return _CTLineGetTrailingWhitespaceWidth(line)
}/* debug [functions.gen.go/function]: CTLineGetTrailingWhitespaceWidth */

// Returns the Core Foundation type identifier of the line object.
//
// Added in macOS 10.5.
// Returns the Core Foundation type identifier of the line object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineGetTypeID()
func CTLineGetTypeID() TypeID {
	return _CTLineGetTypeID()
}/* debug [functions.gen.go/function]: CTLineGetTypeID */

// Calculates the typographic bounds of a line.
//
// Added in macOS 10.5.
// Calculates the typographic bounds of a line.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineGetTypographicBounds(_:_:_:_:)
func CTLineGetTypographicBounds(line LineRef, ascent []float64, descent []float64, leading []float64) float64 {
	return _CTLineGetTypographicBounds(line, ascent, descent, leading)
}/* debug [functions.gen.go/function]: CTLineGetTypographicBounds */

// Creates an immutable paragraph style.
//
// Added in macOS 10.5.
// Creates an immutable paragraph style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleCreate(_:_:)
func CTParagraphStyleCreate(settings unsafe.Pointer, settingCount uintptr) ParagraphStyleRef {
	return _CTParagraphStyleCreate(settings, settingCount)
}/* debug [functions.gen.go/function]: CTParagraphStyleCreate */

// Creates an immutable copy of a paragraph style.
//
// Added in macOS 10.5.
// Creates an immutable copy of a paragraph style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleCreateCopy(_:)
func CTParagraphStyleCreateCopy(paragraphStyle ParagraphStyleRef) ParagraphStyleRef {
	return _CTParagraphStyleCreateCopy(paragraphStyle)
}/* debug [functions.gen.go/function]: CTParagraphStyleCreateCopy */

// Returns the Core Foundation type identifier of the paragraph style object.
//
// Added in macOS 10.5.
// Returns the Core Foundation type identifier of the paragraph style object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleGetTypeID()
func CTParagraphStyleGetTypeID() TypeID {
	return _CTParagraphStyleGetTypeID()
}/* debug [functions.gen.go/function]: CTParagraphStyleGetTypeID */

// Obtains the current value for a single setting specifier.
//
// Added in macOS 10.5.
// Obtains the current value for a single setting specifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleGetValueForSpecifier(_:_:_:_:)
func CTParagraphStyleGetValueForSpecifier(paragraphStyle ParagraphStyleRef, spec ParagraphStyleSpecifier, valueBufferSize uintptr, valueBuffer unsafe.Pointer) bool {
	return _CTParagraphStyleGetValueForSpecifier(paragraphStyle, spec, valueBufferSize, valueBuffer)
}/* debug [functions.gen.go/function]: CTParagraphStyleGetValueForSpecifier */

// Creates an immutable ruby annotation object.
//
// Added in macOS 10.10.
// Creates an immutable ruby annotation object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyAnnotationCreate(_:_:_:_:)
func CTRubyAnnotationCreate(alignment RubyAlignment, overhang RubyOverhang, sizeFactor float64, text StringRef, p4 unsafe.Pointer) RubyAnnotationRef {
	return _CTRubyAnnotationCreate(alignment, overhang, sizeFactor, text, p4)
}/* debug [functions.gen.go/function]: CTRubyAnnotationCreate */

// Creates an immutable copy of a ruby annotation object.
//
// Added in macOS 10.10.
// Creates an immutable copy of a ruby annotation object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyAnnotationCreateCopy(_:)
func CTRubyAnnotationCreateCopy(rubyAnnotation RubyAnnotationRef) RubyAnnotationRef {
	return _CTRubyAnnotationCreateCopy(rubyAnnotation)
}/* debug [functions.gen.go/function]: CTRubyAnnotationCreateCopy */

// Creates an immutable ruby annotation object with the specified attributes.
//
// Added in macOS 10.12.
// Creates an immutable ruby annotation object with the specified attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyAnnotationCreateWithAttributes(_:_:_:_:_:)
func CTRubyAnnotationCreateWithAttributes(alignment RubyAlignment, overhang RubyOverhang, position RubyPosition, string_ StringRef, attributes DictionaryRef) RubyAnnotationRef {
	return _CTRubyAnnotationCreateWithAttributes(alignment, overhang, position, string_, attributes)
}/* debug [functions.gen.go/function]: CTRubyAnnotationCreateWithAttributes */

// Retrieves the alignment value of a ruby annotation object.
//
// Added in macOS 10.10.
// Retrieves the alignment value of a ruby annotation object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyAnnotationGetAlignment(_:)
func CTRubyAnnotationGetAlignment(rubyAnnotation RubyAnnotationRef) RubyAlignment {
	return _CTRubyAnnotationGetAlignment(rubyAnnotation)
}/* debug [functions.gen.go/function]: CTRubyAnnotationGetAlignment */

// Retrieves the overhang value of a ruby annotation object.
//
// Added in macOS 10.10.
// Retrieves the overhang value of a ruby annotation object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyAnnotationGetOverhang(_:)
func CTRubyAnnotationGetOverhang(rubyAnnotation RubyAnnotationRef) RubyOverhang {
	return _CTRubyAnnotationGetOverhang(rubyAnnotation)
}/* debug [functions.gen.go/function]: CTRubyAnnotationGetOverhang */

// Retrieves the size factor of a ruby annotation object.
//
// Added in macOS 10.10.
// Retrieves the size factor of a ruby annotation object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyAnnotationGetSizeFactor(_:)
func CTRubyAnnotationGetSizeFactor(rubyAnnotation RubyAnnotationRef) float64 {
	return _CTRubyAnnotationGetSizeFactor(rubyAnnotation)
}/* debug [functions.gen.go/function]: CTRubyAnnotationGetSizeFactor */

// Retrieves the ruby text for a particular position in a ruby annotation.
//
// Added in macOS 10.10.
// Retrieves the ruby text for a particular position in a ruby annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyAnnotationGetTextForPosition(_:_:)
func CTRubyAnnotationGetTextForPosition(rubyAnnotation RubyAnnotationRef, position RubyPosition) StringRef {
	return _CTRubyAnnotationGetTextForPosition(rubyAnnotation, position)
}/* debug [functions.gen.go/function]: CTRubyAnnotationGetTextForPosition */

// Retrieves the type of the ruby annotation object.
//
// Added in macOS 10.10.
// Retrieves the type of the ruby annotation object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyAnnotationGetTypeID()
func CTRubyAnnotationGetTypeID() TypeID {
	return _CTRubyAnnotationGetTypeID()
}/* debug [functions.gen.go/function]: CTRubyAnnotationGetTypeID */

// Creates an immutable instance of a run delegate.
//
// Added in macOS 10.5.
// Creates an immutable instance of a run delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunDelegateCreate(_:_:)
func CTRunDelegateCreate(callbacks unsafe.Pointer, refCon unsafe.Pointer) RunDelegateRef {
	return _CTRunDelegateCreate(callbacks, refCon)
}/* debug [functions.gen.go/function]: CTRunDelegateCreate */

// Returns a run delegate’s “refCon” value.
//
// Added in macOS 10.5.
// Returns a run delegate’s “refCon” value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunDelegateGetRefCon(_:)
func CTRunDelegateGetRefCon(runDelegate RunDelegateRef) unsafe.Pointer {
	return _CTRunDelegateGetRefCon(runDelegate)
}/* debug [functions.gen.go/function]: CTRunDelegateGetRefCon */

// Returns the type of CTRunDelegate objects.
//
// Added in macOS 10.5.
// Returns the type of CTRunDelegate objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunDelegateGetTypeID()
func CTRunDelegateGetTypeID() TypeID {
	return _CTRunDelegateGetTypeID()
}/* debug [functions.gen.go/function]: CTRunDelegateGetTypeID */

// Draws a complete run or part of one.
//
// Added in macOS 10.5.
// Draws a complete run or part of one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunDraw(_:_:_:)
func CTRunDraw(run RunRef, context ContextRef, range_ corefoundation.Range) {
	_CTRunDraw(run, context, range_)
}/* debug [functions.gen.go/function]: CTRunDraw */

// Copies a range of glyph advances into a user-provided buffer.
//
// Added in macOS 10.5.
// Copies a range of glyph advances into a user-provided buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunGetAdvances(_:_:_:)
func CTRunGetAdvances(run RunRef, range_ corefoundation.Range, buffer corefoundation.CGSize) {
	_CTRunGetAdvances(run, range_, buffer)
}/* debug [functions.gen.go/function]: CTRunGetAdvances */

// Returns a direct pointer for the glyph advance array stored in the run.
//
// Added in macOS 10.5.
// Returns a direct pointer for the glyph advance array stored in the run.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunGetAdvancesPtr(_:)
func CTRunGetAdvancesPtr(run RunRef) unsafe.Pointer {
	return _CTRunGetAdvancesPtr(run)
}/* debug [functions.gen.go/function]: CTRunGetAdvancesPtr */

// Returns the attribute dictionary that was used to create the glyph run.
//
// Added in macOS 10.5.
// Returns the attribute dictionary that was used to create the glyph run.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunGetAttributes(_:)
func CTRunGetAttributes(run RunRef) DictionaryRef {
	return _CTRunGetAttributes(run)
}/* debug [functions.gen.go/function]: CTRunGetAttributes */

// Copies a range of base advances and origins into user-provided buffers.
//
// Added in macOS 10.11.
// Copies a range of base advances and origins into user-provided buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunGetBaseAdvancesAndOrigins(_:_:_:_:)
func CTRunGetBaseAdvancesAndOrigins(runRef RunRef, range_ corefoundation.Range, advancesBuffer corefoundation.CGSize, originsBuffer corefoundation.CGPoint) {
	_CTRunGetBaseAdvancesAndOrigins(runRef, range_, advancesBuffer, originsBuffer)
}/* debug [functions.gen.go/function]: CTRunGetBaseAdvancesAndOrigins */

// Gets the glyph count for the run.
//
// Added in macOS 10.5.
// Gets the glyph count for the run.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunGetGlyphCount(_:)
func CTRunGetGlyphCount(run RunRef) Index {
	return _CTRunGetGlyphCount(run)
}/* debug [functions.gen.go/function]: CTRunGetGlyphCount */

// Copies a range of glyphs into a user-provided buffer.
//
// Added in macOS 10.5.
// Copies a range of glyphs into a user-provided buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunGetGlyphs(_:_:_:)
func CTRunGetGlyphs(run RunRef, range_ corefoundation.Range, buffer Glyph) {
	_CTRunGetGlyphs(run, range_, buffer)
}/* debug [functions.gen.go/function]: CTRunGetGlyphs */

// Returns a direct pointer for the glyph array stored in the run.
//
// Added in macOS 10.5.
// Returns a direct pointer for the glyph array stored in the run.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunGetGlyphsPtr(_:)
func CTRunGetGlyphsPtr(run RunRef) unsafe.Pointer {
	return _CTRunGetGlyphsPtr(run)
}/* debug [functions.gen.go/function]: CTRunGetGlyphsPtr */

// Calculates the image bounds for a glyph range.
//
// Added in macOS 10.5.
// Calculates the image bounds for a glyph range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunGetImageBounds(_:_:_:)
func CTRunGetImageBounds(run RunRef, context ContextRef, range_ corefoundation.Range) corefoundation.CGRect {
	return _CTRunGetImageBounds(run, context, range_)
}/* debug [functions.gen.go/function]: CTRunGetImageBounds */

// Copies a range of glyph positions into a user-provided buffer.
//
// Added in macOS 10.5.
// Copies a range of glyph positions into a user-provided buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunGetPositions(_:_:_:)
func CTRunGetPositions(run RunRef, range_ corefoundation.Range, buffer corefoundation.CGPoint) {
	_CTRunGetPositions(run, range_, buffer)
}/* debug [functions.gen.go/function]: CTRunGetPositions */

// Returns a direct pointer for the glyph position array stored in the run.
//
// Added in macOS 10.5.
// Returns a direct pointer for the glyph position array stored in the run.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunGetPositionsPtr(_:)
func CTRunGetPositionsPtr(run RunRef) unsafe.Pointer {
	return _CTRunGetPositionsPtr(run)
}/* debug [functions.gen.go/function]: CTRunGetPositionsPtr */

// Returns the run’s status.
//
// Added in macOS 10.5.
// Returns the run’s status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunGetStatus(_:)
func CTRunGetStatus(run RunRef) RunStatus {
	return _CTRunGetStatus(run)
}/* debug [functions.gen.go/function]: CTRunGetStatus */

// Copies a range of string indices into a user-provided buffer.
//
// Added in macOS 10.5.
// Copies a range of string indices into a user-provided buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunGetStringIndices(_:_:_:)
func CTRunGetStringIndices(run RunRef, range_ corefoundation.Range, buffer Index) {
	_CTRunGetStringIndices(run, range_, buffer)
}/* debug [functions.gen.go/function]: CTRunGetStringIndices */

// Returns a direct pointer for the string indices stored in the run.
//
// Added in macOS 10.5.
// Returns a direct pointer for the string indices stored in the run.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunGetStringIndicesPtr(_:)
func CTRunGetStringIndicesPtr(run RunRef) unsafe.Pointer {
	return _CTRunGetStringIndicesPtr(run)
}/* debug [functions.gen.go/function]: CTRunGetStringIndicesPtr */

// Gets the range of characters that originally spawned the glyphs in the run.
//
// Added in macOS 10.5.
// Gets the range of characters that originally spawned the glyphs in the run.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunGetStringRange(_:)
func CTRunGetStringRange(run RunRef) corefoundation.Range {
	return _CTRunGetStringRange(run)
}/* debug [functions.gen.go/function]: CTRunGetStringRange */

// Returns the text matrix needed to draw this run.
//
// Added in macOS 10.5.
// Returns the text matrix needed to draw this run.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunGetTextMatrix(_:)
func CTRunGetTextMatrix(run RunRef) corefoundation.CGAffineTransform {
	return _CTRunGetTextMatrix(run)
}/* debug [functions.gen.go/function]: CTRunGetTextMatrix */

// Returns the Core Foundation type identifier of the run object.
//
// Added in macOS 10.5.
// Returns the Core Foundation type identifier of the run object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunGetTypeID()
func CTRunGetTypeID() TypeID {
	return _CTRunGetTypeID()
}/* debug [functions.gen.go/function]: CTRunGetTypeID */

// Gets the typographic bounds of the run.
//
// Added in macOS 10.5.
// Gets the typographic bounds of the run.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunGetTypographicBounds(_:_:_:_:_:)
func CTRunGetTypographicBounds(run RunRef, range_ corefoundation.Range, ascent []float64, descent []float64, leading []float64) float64 {
	return _CTRunGetTypographicBounds(run, range_, ascent, descent, leading)
}/* debug [functions.gen.go/function]: CTRunGetTypographicBounds */

// Creates and initializes a new text tab object.
//
// Added in macOS 10.5.
// Creates and initializes a new text tab object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTextTabCreate(_:_:_:)
func CTTextTabCreate(alignment TextAlignment, location float64, options DictionaryRef) TextTabRef {
	return _CTTextTabCreate(alignment, location, options)
}/* debug [functions.gen.go/function]: CTTextTabCreate */

// Returns the text alignment of the tab.
//
// Added in macOS 10.5.
// Returns the text alignment of the tab.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTextTabGetAlignment(_:)
func CTTextTabGetAlignment(tab TextTabRef) TextAlignment {
	return _CTTextTabGetAlignment(tab)
}/* debug [functions.gen.go/function]: CTTextTabGetAlignment */

// Returns the tab’s ruler location.
//
// Added in macOS 10.5.
// Returns the tab’s ruler location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTextTabGetLocation(_:)
func CTTextTabGetLocation(tab TextTabRef) float64 {
	return _CTTextTabGetLocation(tab)
}/* debug [functions.gen.go/function]: CTTextTabGetLocation */

// Returns the dictionary of attributes associated with the tab.
//
// Added in macOS 10.5.
// Returns the dictionary of attributes associated with the tab.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTextTabGetOptions(_:)
func CTTextTabGetOptions(tab TextTabRef) DictionaryRef {
	return _CTTextTabGetOptions(tab)
}/* debug [functions.gen.go/function]: CTTextTabGetOptions */

// Returns the Core Foundation type identifier of the text tab object.
//
// Added in macOS 10.5.
// Returns the Core Foundation type identifier of the text tab object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTextTabGetTypeID()
func CTTextTabGetTypeID() TypeID {
	return _CTTextTabGetTypeID()
}/* debug [functions.gen.go/function]: CTTextTabGetTypeID */

// Creates an immutable line from the typesetter.
//
// Added in macOS 10.5.
// Creates an immutable line from the typesetter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTypesetterCreateLine(_:_:)
func CTTypesetterCreateLine(typesetter TypesetterRef, stringRange corefoundation.Range) LineRef {
	return _CTTypesetterCreateLine(typesetter, stringRange)
}/* debug [functions.gen.go/function]: CTTypesetterCreateLine */

// Creates an immutable line from the typesetter at a specified line offset.
//
// Added in macOS 10.6.
// Creates an immutable line from the typesetter at a specified line offset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTypesetterCreateLineWithOffset(_:_:_:)
func CTTypesetterCreateLineWithOffset(typesetter TypesetterRef, stringRange corefoundation.Range, offset float64) LineRef {
	return _CTTypesetterCreateLineWithOffset(typesetter, stringRange, offset)
}/* debug [functions.gen.go/function]: CTTypesetterCreateLineWithOffset */

// Creates an immutable typesetter object using an attributed string.
//
// Added in macOS 10.5.
// Creates an immutable typesetter object using an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTypesetterCreateWithAttributedString(_:)
func CTTypesetterCreateWithAttributedString(string_ AttributedStringRef) TypesetterRef {
	return _CTTypesetterCreateWithAttributedString(string_)
}/* debug [functions.gen.go/function]: CTTypesetterCreateWithAttributedString */

// Creates an immutable typesetter object using an attributed string and a dictionary of options.
//
// Added in macOS 10.5.
// Creates an immutable typesetter object using an attributed string and a dictionary of options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTypesetterCreateWithAttributedStringAndOptions(_:_:)
func CTTypesetterCreateWithAttributedStringAndOptions(string_ AttributedStringRef, options DictionaryRef) TypesetterRef {
	return _CTTypesetterCreateWithAttributedStringAndOptions(string_, options)
}/* debug [functions.gen.go/function]: CTTypesetterCreateWithAttributedStringAndOptions */

// Returns the Core Foundation type identifier of the typesetter object.
//
// Added in macOS 10.5.
// Returns the Core Foundation type identifier of the typesetter object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTypesetterGetTypeID()
func CTTypesetterGetTypeID() TypeID {
	return _CTTypesetterGetTypeID()
}/* debug [functions.gen.go/function]: CTTypesetterGetTypeID */

// Suggests a cluster line breakpoint based on the width provided.
//
// Added in macOS 10.5.
// Suggests a cluster line breakpoint based on the width provided.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTypesetterSuggestClusterBreak(_:_:_:)
func CTTypesetterSuggestClusterBreak(typesetter TypesetterRef, startIndex Index, width float64) Index {
	return _CTTypesetterSuggestClusterBreak(typesetter, startIndex, width)
}/* debug [functions.gen.go/function]: CTTypesetterSuggestClusterBreak */

// Suggests a cluster line breakpoint based on the specified width and line offset.
//
// Added in macOS 10.6.
// Suggests a cluster line breakpoint based on the specified width and line offset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTypesetterSuggestClusterBreakWithOffset(_:_:_:_:)
func CTTypesetterSuggestClusterBreakWithOffset(typesetter TypesetterRef, startIndex Index, width float64, offset float64) Index {
	return _CTTypesetterSuggestClusterBreakWithOffset(typesetter, startIndex, width, offset)
}/* debug [functions.gen.go/function]: CTTypesetterSuggestClusterBreakWithOffset */

// Suggests a contextual line breakpoint based on the width provided.
//
// Added in macOS 10.5.
// Suggests a contextual line breakpoint based on the width provided.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTypesetterSuggestLineBreak(_:_:_:)
func CTTypesetterSuggestLineBreak(typesetter TypesetterRef, startIndex Index, width float64) Index {
	return _CTTypesetterSuggestLineBreak(typesetter, startIndex, width)
}/* debug [functions.gen.go/function]: CTTypesetterSuggestLineBreak */

// Suggests a contextual line breakpoint based on the width provided and the specified offset.
//
// Added in macOS 10.6.
// Suggests a contextual line breakpoint based on the width provided and the specified offset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTypesetterSuggestLineBreakWithOffset(_:_:_:_:)
func CTTypesetterSuggestLineBreakWithOffset(typesetter TypesetterRef, startIndex Index, width float64, offset float64) Index {
	return _CTTypesetterSuggestLineBreakWithOffset(typesetter, startIndex, width, offset)
}/* debug [functions.gen.go/function]: CTTypesetterSuggestLineBreakWithOffset */




