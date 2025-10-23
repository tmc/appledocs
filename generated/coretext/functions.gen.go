// Code generated from Apple documentation for CoreText. DO NOT EDIT.

package coretext

import (
	"unsafe"

	"github.com/ebitengine/purego"
	coregraphics "github.com/tmc/appledocs/generated/coregraphics"
)


// CoreText Functions (203 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_CTFontCollectionCopyExclusionDescriptors func(unsafe.Pointer) unsafe.Pointer
	_CTFontCollectionCopyFontAttribute func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTFontCollectionCopyFontAttributes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTFontCollectionCopyQueryDescriptors func(unsafe.Pointer) unsafe.Pointer
	_CTFontCollectionCreateCopyWithFontDescriptors func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTFontCollectionCreateFromAvailableFonts func(unsafe.Pointer) unsafe.Pointer
	_CTFontCollectionCreateMatchingFontDescriptors func(unsafe.Pointer) unsafe.Pointer
	_CTFontCollectionCreateMatchingFontDescriptorsForFamily func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTFontCollectionCreateMatchingFontDescriptorsSortedWithCallback func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTFontCollectionCreateMatchingFontDescriptorsWithOptions func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTFontCollectionCreateMutableCopy func(unsafe.Pointer) unsafe.Pointer
	_CTFontCollectionCreateWithFontDescriptors func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTFontCollectionGetTypeID func() unsafe.Pointer
	_CTFontCollectionSetExclusionDescriptors func(unsafe.Pointer, unsafe.Pointer)
	_CTFontCollectionSetQueryDescriptors func(unsafe.Pointer, unsafe.Pointer)
	_CTFontCopyAttribute func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTFontCopyAvailableTables func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTFontCopyCharacterSet func(unsafe.Pointer) unsafe.Pointer
	_CTFontCopyDefaultCascadeListForLanguages func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTFontCopyDisplayName func(unsafe.Pointer) unsafe.Pointer
	_CTFontCopyFamilyName func(unsafe.Pointer) unsafe.Pointer
	_CTFontCopyFeatureSettings func(unsafe.Pointer) unsafe.Pointer
	_CTFontCopyFeatures func(unsafe.Pointer) unsafe.Pointer
	_CTFontCopyFontDescriptor func(unsafe.Pointer) unsafe.Pointer
	_CTFontCopyFullName func(unsafe.Pointer) unsafe.Pointer
	_CTFontCopyGraphicsFont func(unsafe.Pointer, unsafe.Pointer) coregraphics.FontRef
	_CTFontCopyLocalizedName func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTFontCopyName func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTFontCopyNameForGlyph func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTFontCopyPostScriptName func(unsafe.Pointer) unsafe.Pointer
	_CTFontCopySupportedLanguages func(unsafe.Pointer) unsafe.Pointer
	_CTFontCopyTable func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTFontCopyTraits func(unsafe.Pointer) unsafe.Pointer
	_CTFontCopyVariation func(unsafe.Pointer) unsafe.Pointer
	_CTFontCopyVariationAxes func(unsafe.Pointer) unsafe.Pointer
	_CTFontCreateCopyWithAttributes func(unsafe.Pointer, float64, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTFontCreateCopyWithFamily func(unsafe.Pointer, float64, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTFontCreateCopyWithSymbolicTraits func(unsafe.Pointer, float64, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTFontCreateForString func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTFontCreateForStringWithLanguage func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTFontCreatePathForGlyph func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) coregraphics.PathRef
	_CTFontCreateUIFontForLanguage func(unsafe.Pointer, float64, unsafe.Pointer) unsafe.Pointer
	_CTFontCreateWithFontDescriptor func(unsafe.Pointer, float64, unsafe.Pointer) unsafe.Pointer
	_CTFontCreateWithFontDescriptorAndOptions func(unsafe.Pointer, float64, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTFontCreateWithGraphicsFont func(coregraphics.FontRef, float64, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTFontCreateWithName func(unsafe.Pointer, float64, unsafe.Pointer) unsafe.Pointer
	_CTFontCreateWithNameAndOptions func(unsafe.Pointer, float64, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTFontCreateWithPlatformFont func(unsafe.Pointer, float64, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTFontCreateWithQuickdrawInstance func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, float64) unsafe.Pointer
	_CTFontDescriptorCopyAttribute func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTFontDescriptorCopyAttributes func(unsafe.Pointer) unsafe.Pointer
	_CTFontDescriptorCopyLocalizedAttribute func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTFontDescriptorCreateCopyWithAttributes func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTFontDescriptorCreateCopyWithFamily func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTFontDescriptorCreateCopyWithFeature func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTFontDescriptorCreateCopyWithSymbolicTraits func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTFontDescriptorCreateCopyWithVariation func(unsafe.Pointer, unsafe.Pointer, float64) unsafe.Pointer
	_CTFontDescriptorCreateMatchingFontDescriptor func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTFontDescriptorCreateMatchingFontDescriptors func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTFontDescriptorCreateWithAttributes func(unsafe.Pointer) unsafe.Pointer
	_CTFontDescriptorCreateWithNameAndSize func(unsafe.Pointer, float64) unsafe.Pointer
	_CTFontDescriptorGetTypeID func() unsafe.Pointer
	_CTFontDescriptorMatchFontDescriptorsWithProgressHandler func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_CTFontDrawGlyphs func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uintptr, coregraphics.ContextRef)
	_CTFontDrawImageFromAdaptiveImageProviderAtPoint func(unsafe.Pointer, unsafe.Pointer, coregraphics.CGPoint, coregraphics.ContextRef)
	_CTFontGetAdvancesForGlyphs func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, coregraphics.CGSize, unsafe.Pointer) float64
	_CTFontGetAscent func(unsafe.Pointer) float64
	_CTFontGetBoundingBox func(unsafe.Pointer) coregraphics.CGRect
	_CTFontGetBoundingRectsForGlyphs func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, coregraphics.CGRect, unsafe.Pointer) coregraphics.CGRect
	_CTFontGetCapHeight func(unsafe.Pointer) float64
	_CTFontGetDescent func(unsafe.Pointer) float64
	_CTFontGetGlyphCount func(unsafe.Pointer) unsafe.Pointer
	_CTFontGetGlyphWithName func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTFontGetGlyphsForCharacters func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_CTFontGetLeading func(unsafe.Pointer) float64
	_CTFontGetLigatureCaretPositions func(unsafe.Pointer, unsafe.Pointer, float64, unsafe.Pointer) unsafe.Pointer
	_CTFontGetMatrix func(unsafe.Pointer) coregraphics.CGAffineTransform
	_CTFontGetOpticalBoundsForGlyphs func(unsafe.Pointer, unsafe.Pointer, coregraphics.CGRect, unsafe.Pointer, unsafe.Pointer) coregraphics.CGRect
	_CTFontGetPlatformFont func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTFontGetSize func(unsafe.Pointer) float64
	_CTFontGetSlantAngle func(unsafe.Pointer) float64
	_CTFontGetStringEncoding func(unsafe.Pointer) unsafe.Pointer
	_CTFontGetSymbolicTraits func(unsafe.Pointer) unsafe.Pointer
	_CTFontGetTypeID func() unsafe.Pointer
	_CTFontGetTypographicBoundsForAdaptiveImageProvider func(unsafe.Pointer, unsafe.Pointer) coregraphics.CGRect
	_CTFontGetUnderlinePosition func(unsafe.Pointer) float64
	_CTFontGetUnderlineThickness func(unsafe.Pointer) float64
	_CTFontGetUnitsPerEm func(unsafe.Pointer) unsafe.Pointer
	_CTFontGetVerticalTranslationsForGlyphs func(unsafe.Pointer, unsafe.Pointer, coregraphics.CGSize, unsafe.Pointer)
	_CTFontGetXHeight func(unsafe.Pointer) float64
	_CTFontHasTable func(unsafe.Pointer, unsafe.Pointer) bool
	_CTFontManagerCompareFontFamilyNames func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTFontManagerCopyAvailableFontFamilyNames func() unsafe.Pointer
	_CTFontManagerCopyAvailableFontURLs func() unsafe.Pointer
	_CTFontManagerCopyAvailablePostScriptNames func() unsafe.Pointer
	_CTFontManagerCopyRegisteredFontDescriptors func(unsafe.Pointer, bool) unsafe.Pointer
	_CTFontManagerCreateFontDescriptorFromData func(unsafe.Pointer) unsafe.Pointer
	_CTFontManagerCreateFontDescriptorsFromData func(unsafe.Pointer) unsafe.Pointer
	_CTFontManagerCreateFontDescriptorsFromURL func(unsafe.Pointer) unsafe.Pointer
	_CTFontManagerCreateFontRequestRunLoopSource func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTFontManagerEnableFontDescriptors func(unsafe.Pointer, bool)
	_CTFontManagerGetAutoActivationSetting func(unsafe.Pointer) unsafe.Pointer
	_CTFontManagerGetScopeForURL func(unsafe.Pointer) unsafe.Pointer
	_CTFontManagerIsSupportedFont func(unsafe.Pointer) bool
	_CTFontManagerRegisterFontDescriptors func(unsafe.Pointer, unsafe.Pointer, bool, bool)
	_CTFontManagerRegisterFontURLs func(unsafe.Pointer, unsafe.Pointer, bool, bool)
	_CTFontManagerRegisterFontsForURL func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_CTFontManagerRegisterFontsForURLs func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_CTFontManagerRegisterFontsWithAssetNames func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, bool, bool)
	_CTFontManagerRegisterGraphicsFont func(coregraphics.FontRef, unsafe.Pointer) bool
	_CTFontManagerRequestFonts func(unsafe.Pointer)
	_CTFontManagerSetAutoActivationSetting func(unsafe.Pointer, unsafe.Pointer)
	_CTFontManagerUnregisterFontDescriptors func(unsafe.Pointer, unsafe.Pointer, bool)
	_CTFontManagerUnregisterFontURLs func(unsafe.Pointer, unsafe.Pointer, bool)
	_CTFontManagerUnregisterFontsForURL func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_CTFontManagerUnregisterFontsForURLs func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_CTFontManagerUnregisterGraphicsFont func(coregraphics.FontRef, unsafe.Pointer) bool
	_CTFrameDraw func(unsafe.Pointer, coregraphics.ContextRef)
	_CTFrameGetFrameAttributes func(unsafe.Pointer) unsafe.Pointer
	_CTFrameGetLineOrigins func(unsafe.Pointer, unsafe.Pointer, coregraphics.CGPoint)
	_CTFrameGetLines func(unsafe.Pointer) unsafe.Pointer
	_CTFrameGetPath func(unsafe.Pointer) coregraphics.PathRef
	_CTFrameGetStringRange func(unsafe.Pointer) unsafe.Pointer
	_CTFrameGetTypeID func() unsafe.Pointer
	_CTFrameGetVisibleStringRange func(unsafe.Pointer) unsafe.Pointer
	_CTFramesetterCreateFrame func(unsafe.Pointer, unsafe.Pointer, coregraphics.PathRef, unsafe.Pointer) unsafe.Pointer
	_CTFramesetterCreateWithAttributedString func(unsafe.Pointer) unsafe.Pointer
	_CTFramesetterCreateWithTypesetter func(unsafe.Pointer) unsafe.Pointer
	_CTFramesetterGetTypeID func() unsafe.Pointer
	_CTFramesetterGetTypesetter func(unsafe.Pointer) unsafe.Pointer
	_CTFramesetterSuggestFrameSizeWithConstraints func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, coregraphics.CGSize, unsafe.Pointer) coregraphics.CGSize
	_CTGetCoreTextVersion func() uint32
	_CTGlyphInfoCreateWithCharacterIdentifier func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTGlyphInfoCreateWithGlyph func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTGlyphInfoCreateWithGlyphName func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTGlyphInfoGetCharacterCollection func(unsafe.Pointer) unsafe.Pointer
	_CTGlyphInfoGetCharacterIdentifier func(unsafe.Pointer) unsafe.Pointer
	_CTGlyphInfoGetGlyph func(unsafe.Pointer) unsafe.Pointer
	_CTGlyphInfoGetGlyphName func(unsafe.Pointer) unsafe.Pointer
	_CTGlyphInfoGetTypeID func() unsafe.Pointer
	_CTLineCreateJustifiedLine func(unsafe.Pointer, float64, float64) unsafe.Pointer
	_CTLineCreateTruncatedLine func(unsafe.Pointer, float64, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTLineCreateWithAttributedString func(unsafe.Pointer) unsafe.Pointer
	_CTLineDraw func(unsafe.Pointer, coregraphics.ContextRef)
	_CTLineEnumerateCaretOffsets func(unsafe.Pointer)
	_CTLineGetBoundsWithOptions func(unsafe.Pointer, unsafe.Pointer) coregraphics.CGRect
	_CTLineGetGlyphCount func(unsafe.Pointer) unsafe.Pointer
	_CTLineGetGlyphRuns func(unsafe.Pointer) unsafe.Pointer
	_CTLineGetImageBounds func(unsafe.Pointer, coregraphics.ContextRef) coregraphics.CGRect
	_CTLineGetOffsetForStringIndex func(unsafe.Pointer, unsafe.Pointer, []float64) float64
	_CTLineGetPenOffsetForFlush func(unsafe.Pointer, float64, float64) float64
	_CTLineGetStringIndexForPosition func(unsafe.Pointer, coregraphics.CGPoint) unsafe.Pointer
	_CTLineGetStringRange func(unsafe.Pointer) unsafe.Pointer
	_CTLineGetTrailingWhitespaceWidth func(unsafe.Pointer) float64
	_CTLineGetTypeID func() unsafe.Pointer
	_CTLineGetTypographicBounds func(unsafe.Pointer, []float64, []float64, []float64) float64
	_CTParagraphStyleCreate func(unsafe.Pointer, uintptr) unsafe.Pointer
	_CTParagraphStyleCreateCopy func(unsafe.Pointer) unsafe.Pointer
	_CTParagraphStyleGetTypeID func() unsafe.Pointer
	_CTParagraphStyleGetValueForSpecifier func(unsafe.Pointer, unsafe.Pointer, uintptr, unsafe.Pointer) bool
	_CTRubyAnnotationCreate func(unsafe.Pointer, unsafe.Pointer, float64, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTRubyAnnotationCreateCopy func(unsafe.Pointer) unsafe.Pointer
	_CTRubyAnnotationCreateWithAttributes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTRubyAnnotationGetAlignment func(unsafe.Pointer) unsafe.Pointer
	_CTRubyAnnotationGetOverhang func(unsafe.Pointer) unsafe.Pointer
	_CTRubyAnnotationGetSizeFactor func(unsafe.Pointer) float64
	_CTRubyAnnotationGetTextForPosition func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTRubyAnnotationGetTypeID func() unsafe.Pointer
	_CTRunDelegateCreate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTRunDelegateGetRefCon func(unsafe.Pointer) unsafe.Pointer
	_CTRunDelegateGetTypeID func() unsafe.Pointer
	_CTRunDraw func(unsafe.Pointer, coregraphics.ContextRef, unsafe.Pointer)
	_CTRunGetAdvances func(unsafe.Pointer, unsafe.Pointer, coregraphics.CGSize)
	_CTRunGetAdvancesPtr func(unsafe.Pointer) unsafe.Pointer
	_CTRunGetAttributes func(unsafe.Pointer) unsafe.Pointer
	_CTRunGetBaseAdvancesAndOrigins func(unsafe.Pointer, unsafe.Pointer, coregraphics.CGSize, coregraphics.CGPoint)
	_CTRunGetGlyphCount func(unsafe.Pointer) unsafe.Pointer
	_CTRunGetGlyphs func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_CTRunGetGlyphsPtr func(unsafe.Pointer) unsafe.Pointer
	_CTRunGetImageBounds func(unsafe.Pointer, coregraphics.ContextRef, unsafe.Pointer) coregraphics.CGRect
	_CTRunGetPositions func(unsafe.Pointer, unsafe.Pointer, coregraphics.CGPoint)
	_CTRunGetPositionsPtr func(unsafe.Pointer) unsafe.Pointer
	_CTRunGetStatus func(unsafe.Pointer) unsafe.Pointer
	_CTRunGetStringIndices func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_CTRunGetStringIndicesPtr func(unsafe.Pointer) unsafe.Pointer
	_CTRunGetStringRange func(unsafe.Pointer) unsafe.Pointer
	_CTRunGetTextMatrix func(unsafe.Pointer) coregraphics.CGAffineTransform
	_CTRunGetTypeID func() unsafe.Pointer
	_CTRunGetTypographicBounds func(unsafe.Pointer, unsafe.Pointer, []float64, []float64, []float64) float64
	_CTTextTabCreate func(unsafe.Pointer, float64, unsafe.Pointer) unsafe.Pointer
	_CTTextTabGetAlignment func(unsafe.Pointer) unsafe.Pointer
	_CTTextTabGetLocation func(unsafe.Pointer) float64
	_CTTextTabGetOptions func(unsafe.Pointer) unsafe.Pointer
	_CTTextTabGetTypeID func() unsafe.Pointer
	_CTTypesetterCreateLine func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTTypesetterCreateLineWithOffset func(unsafe.Pointer, unsafe.Pointer, float64) unsafe.Pointer
	_CTTypesetterCreateWithAttributedString func(unsafe.Pointer) unsafe.Pointer
	_CTTypesetterCreateWithAttributedStringAndOptions func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CTTypesetterGetTypeID func() unsafe.Pointer
	_CTTypesetterSuggestClusterBreak func(unsafe.Pointer, unsafe.Pointer, float64) unsafe.Pointer
	_CTTypesetterSuggestClusterBreakWithOffset func(unsafe.Pointer, unsafe.Pointer, float64, float64) unsafe.Pointer
	_CTTypesetterSuggestLineBreak func(unsafe.Pointer, unsafe.Pointer, float64) unsafe.Pointer
	_CTTypesetterSuggestLineBreakWithOffset func(unsafe.Pointer, unsafe.Pointer, float64, float64) unsafe.Pointer
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
	tryRegister(&_CTFontCopyFeatureSettings, lib, "CTFontCopyFeatureSettings")
	tryRegister(&_CTFontCopyFeatures, lib, "CTFontCopyFeatures")
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
	tryRegister(&_CTFontGetGlyphWithName, lib, "CTFontGetGlyphWithName")
	tryRegister(&_CTFontGetGlyphsForCharacters, lib, "CTFontGetGlyphsForCharacters")
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
	tryRegister(&_CTFontManagerRegisterFontURLs, lib, "CTFontManagerRegisterFontURLs")
	tryRegister(&_CTFontManagerRegisterFontsForURL, lib, "CTFontManagerRegisterFontsForURL")
	tryRegister(&_CTFontManagerRegisterFontsForURLs, lib, "CTFontManagerRegisterFontsForURLs")
	tryRegister(&_CTFontManagerRegisterFontsWithAssetNames, lib, "CTFontManagerRegisterFontsWithAssetNames")
	tryRegister(&_CTFontManagerRegisterGraphicsFont, lib, "CTFontManagerRegisterGraphicsFont")
	tryRegister(&_CTFontManagerRequestFonts, lib, "CTFontManagerRequestFonts")
	tryRegister(&_CTFontManagerSetAutoActivationSetting, lib, "CTFontManagerSetAutoActivationSetting")
	tryRegister(&_CTFontManagerUnregisterFontDescriptors, lib, "CTFontManagerUnregisterFontDescriptors")
	tryRegister(&_CTFontManagerUnregisterFontURLs, lib, "CTFontManagerUnregisterFontURLs")
	tryRegister(&_CTFontManagerUnregisterFontsForURL, lib, "CTFontManagerUnregisterFontsForURL")
	tryRegister(&_CTFontManagerUnregisterFontsForURLs, lib, "CTFontManagerUnregisterFontsForURLs")
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
func CTFontCollectionCopyExclusionDescriptors(collection unsafe.Pointer) unsafe.Pointer {
	return _CTFontCollectionCopyExclusionDescriptors(collection)
}

// Retrieves an array of font descriptor attribute values.
//
// Added in macOS 10.7.
// Retrieves an array of font descriptor attribute values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCollectionCopyFontAttribute(_:_:_:)
func CTFontCollectionCopyFontAttribute(collection unsafe.Pointer, attributeName unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	return _CTFontCollectionCopyFontAttribute(collection, attributeName, options)
}

// Retrieves an array of dictionaries containing font descriptor attribute values.
//
// Added in macOS 10.7.
// Retrieves an array of dictionaries containing font descriptor attribute values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCollectionCopyFontAttributes(_:_:_:)
func CTFontCollectionCopyFontAttributes(collection unsafe.Pointer, attributeNames unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	return _CTFontCollectionCopyFontAttributes(collection, attributeNames, options)
}

// Retrieves the array of descriptors for font matching.
//
// Added in macOS 10.7.
// Retrieves the array of descriptors for font matching.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCollectionCopyQueryDescriptors(_:)
func CTFontCollectionCopyQueryDescriptors(collection unsafe.Pointer) unsafe.Pointer {
	return _CTFontCollectionCopyQueryDescriptors(collection)
}

// Returns a copy of the original collection augmented with the given new font descriptors.
//
// Added in macOS 10.5.
// Returns a copy of the original collection augmented with the given new font descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCollectionCreateCopyWithFontDescriptors(_:_:_:)
func CTFontCollectionCreateCopyWithFontDescriptors(original unsafe.Pointer, queryDescriptors unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	return _CTFontCollectionCreateCopyWithFontDescriptors(original, queryDescriptors, options)
}

// Returns a new font collection containing all available fonts.
//
// Added in macOS 10.5.
// Returns a new font collection containing all available fonts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCollectionCreateFromAvailableFonts(_:)
func CTFontCollectionCreateFromAvailableFonts(options unsafe.Pointer) unsafe.Pointer {
	return _CTFontCollectionCreateFromAvailableFonts(options)
}

// Returns an array of font descriptors matching the collection.
//
// Added in macOS 10.5.
// Returns an array of font descriptors matching the collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCollectionCreateMatchingFontDescriptors(_:)
func CTFontCollectionCreateMatchingFontDescriptors(collection unsafe.Pointer) unsafe.Pointer {
	return _CTFontCollectionCreateMatchingFontDescriptors(collection)
}

// Retrieves an array of font descriptors that match the specified family, one descriptor for each style in the collection.
//
// Added in macOS 10.7.
// Retrieves an array of font descriptors that match the specified family, one descriptor for each style in the collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCollectionCreateMatchingFontDescriptorsForFamily(_:_:_:)
func CTFontCollectionCreateMatchingFontDescriptorsForFamily(collection unsafe.Pointer, familyName unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	return _CTFontCollectionCreateMatchingFontDescriptorsForFamily(collection, familyName, options)
}

// Returns the array of matching font descriptors sorted with the callback function.
//
// Added in macOS 10.5.
// Returns the array of matching font descriptors sorted with the callback function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCollectionCreateMatchingFontDescriptorsSortedWithCallback(_:_:_:)
func CTFontCollectionCreateMatchingFontDescriptorsSortedWithCallback(collection unsafe.Pointer, sortCallback unsafe.Pointer, refCon unsafe.Pointer) unsafe.Pointer {
	return _CTFontCollectionCreateMatchingFontDescriptorsSortedWithCallback(collection, sortCallback, refCon)
}

// Creates an array of font descriptors that match the specified collection.
//
// Added in macOS 10.7.
// Creates an array of font descriptors that match the specified collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCollectionCreateMatchingFontDescriptorsWithOptions(_:_:)
func CTFontCollectionCreateMatchingFontDescriptorsWithOptions(collection unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	return _CTFontCollectionCreateMatchingFontDescriptorsWithOptions(collection, options)
}

// Creates a mutable copy of the original collection.
//
// Added in macOS 10.7.
// Creates a mutable copy of the original collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCollectionCreateMutableCopy(_:)
func CTFontCollectionCreateMutableCopy(original unsafe.Pointer) unsafe.Pointer {
	return _CTFontCollectionCreateMutableCopy(original)
}

// Returns a new font collection based on the given array of font descriptors.
//
// Added in macOS 10.5.
// Returns a new font collection based on the given array of font descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCollectionCreateWithFontDescriptors(_:_:)
func CTFontCollectionCreateWithFontDescriptors(queryDescriptors unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	return _CTFontCollectionCreateWithFontDescriptors(queryDescriptors, options)
}

// Returns the type identifier for Core Text font collection references.
//
// Added in macOS 10.5.
// Returns the type identifier for Core Text font collection references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCollectionGetTypeID()
func CTFontCollectionGetTypeID() unsafe.Pointer {
	return _CTFontCollectionGetTypeID()
}

// Replaces the array of descriptors to exclude from the match.
//
// Added in macOS 10.7.
// Replaces the array of descriptors to exclude from the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCollectionSetExclusionDescriptors(_:_:)
func CTFontCollectionSetExclusionDescriptors(collection unsafe.Pointer, descriptors unsafe.Pointer) {
	_CTFontCollectionSetExclusionDescriptors(collection, descriptors)
}

// Replaces the array of descriptors for font matching.
//
// Added in macOS 10.7.
// Replaces the array of descriptors for font matching.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCollectionSetQueryDescriptors(_:_:)
func CTFontCollectionSetQueryDescriptors(collection unsafe.Pointer, descriptors unsafe.Pointer) {
	_CTFontCollectionSetQueryDescriptors(collection, descriptors)
}

// Returns the value associated with an arbitrary attribute of the given font.
//
// Added in macOS 10.5.
// Returns the value associated with an arbitrary attribute of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCopyAttribute(_:_:)
func CTFontCopyAttribute(font unsafe.Pointer, attribute unsafe.Pointer) unsafe.Pointer {
	return _CTFontCopyAttribute(font, attribute)
}

// Returns an array of font table tags.
//
// Added in macOS 10.5.
// Returns an array of font table tags.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCopyAvailableTables(_:_:)
func CTFontCopyAvailableTables(font unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	return _CTFontCopyAvailableTables(font, options)
}

// Returns the Unicode character set of the font.
//
// Added in macOS 10.5.
// Returns the Unicode character set of the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCopyCharacterSet(_:)
func CTFontCopyCharacterSet(font unsafe.Pointer) unsafe.Pointer {
	return _CTFontCopyCharacterSet(font)
}

// Retrieves an ordered list of font substitution preferences.
//
// Added in macOS 10.8.
// Retrieves an ordered list of font substitution preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCopyDefaultCascadeListForLanguages(_:_:)
func CTFontCopyDefaultCascadeListForLanguages(font unsafe.Pointer, languagePrefList unsafe.Pointer) unsafe.Pointer {
	return _CTFontCopyDefaultCascadeListForLanguages(font, languagePrefList)
}

// Returns the display name of the given font.
//
// Added in macOS 10.5.
// Returns the display name of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCopyDisplayName(_:)
func CTFontCopyDisplayName(font unsafe.Pointer) unsafe.Pointer {
	return _CTFontCopyDisplayName(font)
}

// Returns the family name of the given font.
//
// Added in macOS 10.5.
// Returns the family name of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCopyFamilyName(_:)
func CTFontCopyFamilyName(font unsafe.Pointer) unsafe.Pointer {
	return _CTFontCopyFamilyName(font)
}

// Returns an array of font feature-setting tuples.
//
// Added in macOS 10.5.
// Returns an array of font feature-setting tuples.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCopyFeatureSettings(_:)
func CTFontCopyFeatureSettings(font unsafe.Pointer) unsafe.Pointer {
	return _CTFontCopyFeatureSettings(font)
}

// Returns an array of font features.
//
// Added in macOS 10.5.
// Returns an array of font features.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCopyFeatures(_:)
func CTFontCopyFeatures(font unsafe.Pointer) unsafe.Pointer {
	return _CTFontCopyFeatures(font)
}

// Returns the normalized font descriptor for the given font reference.
//
// Added in macOS 10.5.
// Returns the normalized font descriptor for the given font reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCopyFontDescriptor(_:)
func CTFontCopyFontDescriptor(font unsafe.Pointer) unsafe.Pointer {
	return _CTFontCopyFontDescriptor(font)
}

// Returns the full name of the given font.
//
// Added in macOS 10.5.
// Returns the full name of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCopyFullName(_:)
func CTFontCopyFullName(font unsafe.Pointer) unsafe.Pointer {
	return _CTFontCopyFullName(font)
}

// Returns a Core Graphics font reference and attributes.
//
// Added in macOS 10.5.
// Returns a Core Graphics font reference and attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCopyGraphicsFont(_:_:)
func CTFontCopyGraphicsFont(font unsafe.Pointer, attributes unsafe.Pointer) coregraphics.FontRef {
	return _CTFontCopyGraphicsFont(font, attributes)
}

// Returns a reference to a localized name for the given font.
//
// Added in macOS 10.5.
// Returns a reference to a localized name for the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCopyLocalizedName(_:_:_:)
func CTFontCopyLocalizedName(font unsafe.Pointer, nameKey unsafe.Pointer, actualLanguage unsafe.Pointer) unsafe.Pointer {
	return _CTFontCopyLocalizedName(font, nameKey, actualLanguage)
}

// Returns a reference to the requested name of the given font.
//
// Added in macOS 10.5.
// Returns a reference to the requested name of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCopyName(_:_:)
func CTFontCopyName(font unsafe.Pointer, nameKey unsafe.Pointer) unsafe.Pointer {
	return _CTFontCopyName(font, nameKey)
}

// Retrieves the name for the specified glyph.
//
// Added in macOS 10.8.
// Retrieves the name for the specified glyph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCopyNameForGlyph(_:_:)
func CTFontCopyNameForGlyph(font unsafe.Pointer, glyph unsafe.Pointer) unsafe.Pointer {
	return _CTFontCopyNameForGlyph(font, glyph)
}

// Returns the PostScript name of the given font.
//
// Added in macOS 10.5.
// Returns the PostScript name of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCopyPostScriptName(_:)
func CTFontCopyPostScriptName(font unsafe.Pointer) unsafe.Pointer {
	return _CTFontCopyPostScriptName(font)
}

// Returns an array of languages supported by the font.
//
// Added in macOS 10.5.
// Returns an array of languages supported by the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCopySupportedLanguages(_:)
func CTFontCopySupportedLanguages(font unsafe.Pointer) unsafe.Pointer {
	return _CTFontCopySupportedLanguages(font)
}

// Returns a reference to the font table data.
//
// Added in macOS 10.5.
// Returns a reference to the font table data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCopyTable(_:_:_:)
func CTFontCopyTable(font unsafe.Pointer, table unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	return _CTFontCopyTable(font, table, options)
}

// Returns the traits dictionary of the given font.
//
// Added in macOS 10.5.
// Returns the traits dictionary of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCopyTraits(_:)
func CTFontCopyTraits(font unsafe.Pointer) unsafe.Pointer {
	return _CTFontCopyTraits(font)
}

// Returns a variation dictionary from the font reference.
//
// Added in macOS 10.5.
// Returns a variation dictionary from the font reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCopyVariation(_:)
func CTFontCopyVariation(font unsafe.Pointer) unsafe.Pointer {
	return _CTFontCopyVariation(font)
}

// Returns an array of variation axes.
//
// Added in macOS 10.5.
// Returns an array of variation axes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCopyVariationAxes(_:)
func CTFontCopyVariationAxes(font unsafe.Pointer) unsafe.Pointer {
	return _CTFontCopyVariationAxes(font)
}

// Returns a new font with additional attributes based on the original font.
//
// Added in macOS 10.5.
// Returns a new font with additional attributes based on the original font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCreateCopyWithAttributes(_:_:_:_:)
func CTFontCreateCopyWithAttributes(font unsafe.Pointer, size float64, matrix unsafe.Pointer, attributes unsafe.Pointer) unsafe.Pointer {
	return _CTFontCreateCopyWithAttributes(font, size, matrix, attributes)
}

// Returns a new font in the specified family based on the traits of the original font.
//
// Added in macOS 10.5.
// Returns a new font in the specified family based on the traits of the original font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCreateCopyWithFamily(_:_:_:_:)
func CTFontCreateCopyWithFamily(font unsafe.Pointer, size float64, matrix unsafe.Pointer, family unsafe.Pointer) unsafe.Pointer {
	return _CTFontCreateCopyWithFamily(font, size, matrix, family)
}

// Returns a new font in the same font family as the original with the specified symbolic traits.
//
// Added in macOS 10.5.
// Returns a new font in the same font family as the original with the specified symbolic traits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCreateCopyWithSymbolicTraits(_:_:_:_:_:)
func CTFontCreateCopyWithSymbolicTraits(font unsafe.Pointer, size float64, matrix unsafe.Pointer, symTraitValue unsafe.Pointer, symTraitMask unsafe.Pointer) unsafe.Pointer {
	return _CTFontCreateCopyWithSymbolicTraits(font, size, matrix, symTraitValue, symTraitMask)
}

// Returns a font reference that most accurately maps the string range based on the current font.
//
// Added in macOS 10.5.
// Returns a font reference that most accurately maps the string range based on the current font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCreateForString(_:_:_:)
func CTFontCreateForString(currentFont unsafe.Pointer, string_ unsafe.Pointer, range_ unsafe.Pointer) unsafe.Pointer {
	return _CTFontCreateForString(currentFont, string_, range_)
}

// Returns a font reference that most accurately maps the string range based on the current font and language.
//
// Added in macOS 10.9.
// Returns a font reference that most accurately maps the string range based on the current font and language.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCreateForStringWithLanguage(_:_:_:_:)
func CTFontCreateForStringWithLanguage(currentFont unsafe.Pointer, string_ unsafe.Pointer, range_ unsafe.Pointer, language unsafe.Pointer) unsafe.Pointer {
	return _CTFontCreateForStringWithLanguage(currentFont, string_, range_, language)
}

// Creates a path for the specified glyph.
//
// Added in macOS 10.5.
// Creates a path for the specified glyph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCreatePathForGlyph(_:_:_:)
func CTFontCreatePathForGlyph(font unsafe.Pointer, glyph unsafe.Pointer, matrix unsafe.Pointer) coregraphics.PathRef {
	return _CTFontCreatePathForGlyph(font, glyph, matrix)
}

// Returns the special user-interface font for the given language and user-interface type.
//
// Added in macOS 10.5.
// Returns the special user-interface font for the given language and user-interface type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCreateUIFontForLanguage(_:_:_:)
func CTFontCreateUIFontForLanguage(uiType unsafe.Pointer, size float64, language unsafe.Pointer) unsafe.Pointer {
	return _CTFontCreateUIFontForLanguage(uiType, size, language)
}

// Returns a new font reference that best matches the given font descriptor.
//
// Added in macOS 10.5.
// Returns a new font reference that best matches the given font descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCreateWithFontDescriptor(_:_:_:)
func CTFontCreateWithFontDescriptor(descriptor unsafe.Pointer, size float64, matrix unsafe.Pointer) unsafe.Pointer {
	return _CTFontCreateWithFontDescriptor(descriptor, size, matrix)
}

// Returns a new font reference that best matches the given font descriptor.
//
// Added in macOS 10.6.
// Returns a new font reference that best matches the given font descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCreateWithFontDescriptorAndOptions(_:_:_:_:)
func CTFontCreateWithFontDescriptorAndOptions(descriptor unsafe.Pointer, size float64, matrix unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	return _CTFontCreateWithFontDescriptorAndOptions(descriptor, size, matrix, options)
}

// Creates a new font reference from an existing Core Graphics font reference.
//
// Added in macOS 10.5.
// Creates a new font reference from an existing Core Graphics font reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCreateWithGraphicsFont(_:_:_:_:)
func CTFontCreateWithGraphicsFont(graphicsFont coregraphics.FontRef, size float64, matrix unsafe.Pointer, attributes unsafe.Pointer) unsafe.Pointer {
	return _CTFontCreateWithGraphicsFont(graphicsFont, size, matrix, attributes)
}

// Returns a new font reference for the given name.
//
// Added in macOS 10.5.
// Returns a new font reference for the given name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCreateWithName(_:_:_:)
func CTFontCreateWithName(name unsafe.Pointer, size float64, matrix unsafe.Pointer) unsafe.Pointer {
	return _CTFontCreateWithName(name, size, matrix)
}

// Returns a new font reference for the given name.
//
// Added in macOS 10.6.
// Returns a new font reference for the given name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCreateWithNameAndOptions(_:_:_:_:)
func CTFontCreateWithNameAndOptions(name unsafe.Pointer, size float64, matrix unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	return _CTFontCreateWithNameAndOptions(name, size, matrix, options)
}

// Creates a new font reference from an ATS font reference.
//
// Deprecated: This function was deprecated in macOS 11.0.
//
// Added in macOS 10.5.
// Creates a new font reference from an ATS font reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCreateWithPlatformFont(_:_:_:_:)
func CTFontCreateWithPlatformFont(platformFont unsafe.Pointer, size float64, matrix unsafe.Pointer, attributes unsafe.Pointer) unsafe.Pointer {
	return _CTFontCreateWithPlatformFont(platformFont, size, matrix, attributes)
}

// Returns a font reference for the given QuickDraw instance.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.5.
// Returns a font reference for the given QuickDraw instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCreateWithQuickdrawInstance(_:_:_:_:)
func CTFontCreateWithQuickdrawInstance(name unsafe.Pointer, identifier unsafe.Pointer, style unsafe.Pointer, size float64) unsafe.Pointer {
	return _CTFontCreateWithQuickdrawInstance(name, identifier, style, size)
}

// Returns the value associated with an arbitrary attribute.
//
// Added in macOS 10.5.
// Returns the value associated with an arbitrary attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorCopyAttribute(_:_:)
func CTFontDescriptorCopyAttribute(descriptor unsafe.Pointer, attribute unsafe.Pointer) unsafe.Pointer {
	return _CTFontDescriptorCopyAttribute(descriptor, attribute)
}

// Returns the attributes dictionary of the font descriptor.
//
// Added in macOS 10.5.
// Returns the attributes dictionary of the font descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorCopyAttributes(_:)
func CTFontDescriptorCopyAttributes(descriptor unsafe.Pointer) unsafe.Pointer {
	return _CTFontDescriptorCopyAttributes(descriptor)
}

// Returns a localized value for the requested attribute, if available.
//
// Added in macOS 10.5.
// Returns a localized value for the requested attribute, if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorCopyLocalizedAttribute(_:_:_:)
func CTFontDescriptorCopyLocalizedAttribute(descriptor unsafe.Pointer, attribute unsafe.Pointer, language unsafe.Pointer) unsafe.Pointer {
	return _CTFontDescriptorCopyLocalizedAttribute(descriptor, attribute, language)
}

// Creates a copy of the original font descriptor with new attributes.
//
// Added in macOS 10.5.
// Creates a copy of the original font descriptor with new attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorCreateCopyWithAttributes(_:_:)
func CTFontDescriptorCreateCopyWithAttributes(original unsafe.Pointer, attributes unsafe.Pointer) unsafe.Pointer {
	return _CTFontDescriptorCreateCopyWithAttributes(original, attributes)
}

// Creates a copy of the font descriptor in the specified family based on the traits of the original.
//
// Added in macOS 10.9.
// Creates a copy of the font descriptor in the specified family based on the traits of the original.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorCreateCopyWithFamily(_:_:)
func CTFontDescriptorCreateCopyWithFamily(original unsafe.Pointer, family unsafe.Pointer) unsafe.Pointer {
	return _CTFontDescriptorCreateCopyWithFamily(original, family)
}

// Copies a font descriptor with new feature settings.
//
// Added in macOS 10.5.
// Copies a font descriptor with new feature settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorCreateCopyWithFeature(_:_:_:)
func CTFontDescriptorCreateCopyWithFeature(original unsafe.Pointer, featureTypeIdentifier unsafe.Pointer, featureSelectorIdentifier unsafe.Pointer) unsafe.Pointer {
	return _CTFontDescriptorCreateCopyWithFeature(original, featureTypeIdentifier, featureSelectorIdentifier)
}

// Creates a copy of the font descriptor with the specified symbolic traits as the original.
//
// Added in macOS 10.9.
// Creates a copy of the font descriptor with the specified symbolic traits as the original.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorCreateCopyWithSymbolicTraits(_:_:_:)
func CTFontDescriptorCreateCopyWithSymbolicTraits(original unsafe.Pointer, symTraitValue unsafe.Pointer, symTraitMask unsafe.Pointer) unsafe.Pointer {
	return _CTFontDescriptorCreateCopyWithSymbolicTraits(original, symTraitValue, symTraitMask)
}

// Creates a copy of the original font descriptor with a new variation instance.
//
// Added in macOS 10.5.
// Creates a copy of the original font descriptor with a new variation instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorCreateCopyWithVariation(_:_:_:)
func CTFontDescriptorCreateCopyWithVariation(original unsafe.Pointer, variationIdentifier unsafe.Pointer, variationValue float64) unsafe.Pointer {
	return _CTFontDescriptorCreateCopyWithVariation(original, variationIdentifier, variationValue)
}

// Returns the single preferred matching font descriptor based on the original descriptor and system precedence.
//
// Added in macOS 10.5.
// Returns the single preferred matching font descriptor based on the original descriptor and system precedence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorCreateMatchingFontDescriptor(_:_:)
func CTFontDescriptorCreateMatchingFontDescriptor(descriptor unsafe.Pointer, mandatoryAttributes unsafe.Pointer) unsafe.Pointer {
	return _CTFontDescriptorCreateMatchingFontDescriptor(descriptor, mandatoryAttributes)
}

// Returns an array of normalized font descriptors matching the provided descriptor.
//
// Added in macOS 10.5.
// Returns an array of normalized font descriptors matching the provided descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorCreateMatchingFontDescriptors(_:_:)
func CTFontDescriptorCreateMatchingFontDescriptors(descriptor unsafe.Pointer, mandatoryAttributes unsafe.Pointer) unsafe.Pointer {
	return _CTFontDescriptorCreateMatchingFontDescriptors(descriptor, mandatoryAttributes)
}

// Creates a new font descriptor reference from a dictionary of attributes.
//
// Added in macOS 10.5.
// Creates a new font descriptor reference from a dictionary of attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorCreateWithAttributes(_:)
func CTFontDescriptorCreateWithAttributes(attributes unsafe.Pointer) unsafe.Pointer {
	return _CTFontDescriptorCreateWithAttributes(attributes)
}

// Creates a new font descriptor with the provided PostScript name and size.
//
// Added in macOS 10.5.
// Creates a new font descriptor with the provided PostScript name and size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorCreateWithNameAndSize(_:_:)
func CTFontDescriptorCreateWithNameAndSize(name unsafe.Pointer, size float64) unsafe.Pointer {
	return _CTFontDescriptorCreateWithNameAndSize(name, size)
}

// Returns the type identifier for Core Text font descriptor references.
//
// Added in macOS 10.5.
// Returns the type identifier for Core Text font descriptor references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorGetTypeID()
func CTFontDescriptorGetTypeID() unsafe.Pointer {
	return _CTFontDescriptorGetTypeID()
}

// Matches font descriptors and tracks progress with a progress handler.
//
// Added in macOS 10.9.
// Matches font descriptors and tracks progress with a progress handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptorMatchFontDescriptorsWithProgressHandler(_:_:_:)
func CTFontDescriptorMatchFontDescriptorsWithProgressHandler(descriptors unsafe.Pointer, mandatoryAttributes unsafe.Pointer, progressBlock unsafe.Pointer) bool {
	return _CTFontDescriptorMatchFontDescriptorsWithProgressHandler(descriptors, mandatoryAttributes, progressBlock)
}

// Renders the given glyphs of a font at the specified positions in the supplied graphics context.
//
// Added in macOS 10.7.
// Renders the given glyphs of a font at the specified positions in the supplied graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDrawGlyphs(_:_:_:_:_:)
func CTFontDrawGlyphs(font unsafe.Pointer, glyphs unsafe.Pointer, positions unsafe.Pointer, count uintptr, context coregraphics.ContextRef) {
	_CTFontDrawGlyphs(font, glyphs, positions, count, context)
}

// CTFontDrawImageFromAdaptiveImageProviderAtPoint is a CoreText function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDrawImageFromAdaptiveImageProviderAtPoint(_:_:_:_:)
func CTFontDrawImageFromAdaptiveImageProviderAtPoint(font unsafe.Pointer, provider unsafe.Pointer, point coregraphics.CGPoint, context coregraphics.ContextRef) {
	_CTFontDrawImageFromAdaptiveImageProviderAtPoint(font, provider, point, context)
}

// Calculates the advances for an array of glyphs and returns the summed advance.
//
// Added in macOS 10.5.
// Calculates the advances for an array of glyphs and returns the summed advance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetAdvancesForGlyphs(_:_:_:_:_:)
func CTFontGetAdvancesForGlyphs(font unsafe.Pointer, orientation unsafe.Pointer, glyphs unsafe.Pointer, advances coregraphics.CGSize, count unsafe.Pointer) float64 {
	return _CTFontGetAdvancesForGlyphs(font, orientation, glyphs, advances, count)
}

// Returns the scaled font-ascent metric of the given font.
//
// Added in macOS 10.5.
// Returns the scaled font-ascent metric of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetAscent(_:)
func CTFontGetAscent(font unsafe.Pointer) float64 {
	return _CTFontGetAscent(font)
}

// Returns the scaled bounding box of the given font.
//
// Added in macOS 10.5.
// Returns the scaled bounding box of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetBoundingBox(_:)
func CTFontGetBoundingBox(font unsafe.Pointer) coregraphics.CGRect {
	return _CTFontGetBoundingBox(font)
}

// Calculates the bounding rects for an array of glyphs and returns the overall bounding rectangle for the glyph run.
//
// Added in macOS 10.5.
// Calculates the bounding rects for an array of glyphs and returns the overall bounding rectangle for the glyph run.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetBoundingRectsForGlyphs(_:_:_:_:_:)
func CTFontGetBoundingRectsForGlyphs(font unsafe.Pointer, orientation unsafe.Pointer, glyphs unsafe.Pointer, boundingRects coregraphics.CGRect, count unsafe.Pointer) coregraphics.CGRect {
	return _CTFontGetBoundingRectsForGlyphs(font, orientation, glyphs, boundingRects, count)
}

// Returns the cap-height metric of the given font.
//
// Added in macOS 10.5.
// Returns the cap-height metric of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetCapHeight(_:)
func CTFontGetCapHeight(font unsafe.Pointer) float64 {
	return _CTFontGetCapHeight(font)
}

// Returns the scaled font-descent metric of the given font.
//
// Added in macOS 10.5.
// Returns the scaled font-descent metric of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetDescent(_:)
func CTFontGetDescent(font unsafe.Pointer) float64 {
	return _CTFontGetDescent(font)
}

// Returns the number of glyphs of the given font.
//
// Added in macOS 10.5.
// Returns the number of glyphs of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetGlyphCount(_:)
func CTFontGetGlyphCount(font unsafe.Pointer) unsafe.Pointer {
	return _CTFontGetGlyphCount(font)
}

// Returns the glyph for the specified name.
//
// Added in macOS 10.5.
// Returns the glyph for the specified name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetGlyphWithName(_:_:)
func CTFontGetGlyphWithName(font unsafe.Pointer, glyphName unsafe.Pointer) unsafe.Pointer {
	return _CTFontGetGlyphWithName(font, glyphName)
}

// Performs basic character-to-glyph mapping.
//
// Added in macOS 10.5.
// Performs basic character-to-glyph mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetGlyphsForCharacters(_:_:_:_:)
func CTFontGetGlyphsForCharacters(font unsafe.Pointer, characters unsafe.Pointer, glyphs unsafe.Pointer, count unsafe.Pointer) bool {
	return _CTFontGetGlyphsForCharacters(font, characters, glyphs, count)
}

// Returns the scaled font-leading metric of the given font.
//
// Added in macOS 10.5.
// Returns the scaled font-leading metric of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetLeading(_:)
func CTFontGetLeading(font unsafe.Pointer) float64 {
	return _CTFontGetLeading(font)
}

// Returns caret positions within a glyph.
//
// Added in macOS 10.5.
// Returns caret positions within a glyph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetLigatureCaretPositions(_:_:_:_:)
func CTFontGetLigatureCaretPositions(font unsafe.Pointer, glyph unsafe.Pointer, positions float64, maxPositions unsafe.Pointer) unsafe.Pointer {
	return _CTFontGetLigatureCaretPositions(font, glyph, positions, maxPositions)
}

// Returns the transformation matrix of the given font.
//
// Added in macOS 10.5.
// Returns the transformation matrix of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetMatrix(_:)
func CTFontGetMatrix(font unsafe.Pointer) coregraphics.CGAffineTransform {
	return _CTFontGetMatrix(font)
}

// Calculates the optical bounds for an array of glyphs and returns the overall optical bounds for the run.
//
// Added in macOS 10.8.
// Calculates the optical bounds for an array of glyphs and returns the overall optical bounds for the run.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetOpticalBoundsForGlyphs(_:_:_:_:_:)
func CTFontGetOpticalBoundsForGlyphs(font unsafe.Pointer, glyphs unsafe.Pointer, boundingRects coregraphics.CGRect, count unsafe.Pointer, options unsafe.Pointer) coregraphics.CGRect {
	return _CTFontGetOpticalBoundsForGlyphs(font, glyphs, boundingRects, count, options)
}

// Returns an ATS font reference and attributes.
//
// Deprecated: This function was deprecated in macOS 11.0.
//
// Added in macOS 10.5.
// Returns an ATS font reference and attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetPlatformFont(_:_:)
func CTFontGetPlatformFont(font unsafe.Pointer, attributes unsafe.Pointer) unsafe.Pointer {
	return _CTFontGetPlatformFont(font, attributes)
}

// Returns the point size of the given font.
//
// Added in macOS 10.5.
// Returns the point size of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetSize(_:)
func CTFontGetSize(font unsafe.Pointer) float64 {
	return _CTFontGetSize(font)
}

// Returns the slant angle of the given font.
//
// Added in macOS 10.5.
// Returns the slant angle of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetSlantAngle(_:)
func CTFontGetSlantAngle(font unsafe.Pointer) float64 {
	return _CTFontGetSlantAngle(font)
}

// Returns the best string encoding for legacy format support.
//
// Added in macOS 10.5.
// Returns the best string encoding for legacy format support.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetStringEncoding(_:)
func CTFontGetStringEncoding(font unsafe.Pointer) unsafe.Pointer {
	return _CTFontGetStringEncoding(font)
}

// Returns the symbolic traits of the given font.
//
// Added in macOS 10.5.
// Returns the symbolic traits of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetSymbolicTraits(_:)
func CTFontGetSymbolicTraits(font unsafe.Pointer) unsafe.Pointer {
	return _CTFontGetSymbolicTraits(font)
}

// Returns the type identifier for Core Text font references.
//
// Added in macOS 10.5.
// Returns the type identifier for Core Text font references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetTypeID()
func CTFontGetTypeID() unsafe.Pointer {
	return _CTFontGetTypeID()
}

// CTFontGetTypographicBoundsForAdaptiveImageProvider is a CoreText function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetTypographicBoundsForAdaptiveImageProvider(_:_:)
func CTFontGetTypographicBoundsForAdaptiveImageProvider(font unsafe.Pointer, provider unsafe.Pointer) coregraphics.CGRect {
	return _CTFontGetTypographicBoundsForAdaptiveImageProvider(font, provider)
}

// Returns the scaled underline position of the given font.
//
// Added in macOS 10.5.
// Returns the scaled underline position of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetUnderlinePosition(_:)
func CTFontGetUnderlinePosition(font unsafe.Pointer) float64 {
	return _CTFontGetUnderlinePosition(font)
}

// Returns the scaled underline-thickness metric of the given font.
//
// Added in macOS 10.5.
// Returns the scaled underline-thickness metric of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetUnderlineThickness(_:)
func CTFontGetUnderlineThickness(font unsafe.Pointer) float64 {
	return _CTFontGetUnderlineThickness(font)
}

// Returns the units-per-em metric of the given font.
//
// Added in macOS 10.5.
// Returns the units-per-em metric of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetUnitsPerEm(_:)
func CTFontGetUnitsPerEm(font unsafe.Pointer) unsafe.Pointer {
	return _CTFontGetUnitsPerEm(font)
}

// Calculates the offset from the default (horizontal) origin to the vertical origin for an array of glyphs.
//
// Added in macOS 10.5.
// Calculates the offset from the default (horizontal) origin to the vertical origin for an array of glyphs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetVerticalTranslationsForGlyphs(_:_:_:_:)
func CTFontGetVerticalTranslationsForGlyphs(font unsafe.Pointer, glyphs unsafe.Pointer, translations coregraphics.CGSize, count unsafe.Pointer) {
	_CTFontGetVerticalTranslationsForGlyphs(font, glyphs, translations, count)
}

// Returns the x-height metric of the given font.
//
// Added in macOS 10.5.
// Returns the x-height metric of the given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontGetXHeight(_:)
func CTFontGetXHeight(font unsafe.Pointer) float64 {
	return _CTFontGetXHeight(font)
}

// CTFontHasTable is a CoreText function.
//
// Added in macOS 10.15.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontHasTable(_:_:)
func CTFontHasTable(font unsafe.Pointer, tag unsafe.Pointer) bool {
	return _CTFontHasTable(font, tag)
}

// A comparator function to compare font family names and sort them according to Apple guidelines.
//
// Added in macOS 10.6.
// A comparator function to compare font family names and sort them according to Apple guidelines.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerCompareFontFamilyNames(_:_:_:)
func CTFontManagerCompareFontFamilyNames(family1 unsafe.Pointer, family2 unsafe.Pointer, context unsafe.Pointer) unsafe.Pointer {
	return _CTFontManagerCompareFontFamilyNames(family1, family2, context)
}

// Returns an array of visible font family names sorted for user interface display.
//
// Added in macOS 10.6.
// Returns an array of visible font family names sorted for user interface display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerCopyAvailableFontFamilyNames()
func CTFontManagerCopyAvailableFontFamilyNames() unsafe.Pointer {
	return _CTFontManagerCopyAvailableFontFamilyNames()
}

// Returns an array of font URLs.
//
// Added in macOS 10.6.
// Returns an array of font URLs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerCopyAvailableFontURLs()
func CTFontManagerCopyAvailableFontURLs() unsafe.Pointer {
	return _CTFontManagerCopyAvailableFontURLs()
}

// Returns an array of unique PostScript font names for the fonts.
//
// Added in macOS 10.6.
// Returns an array of unique PostScript font names for the fonts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerCopyAvailablePostScriptNames()
func CTFontManagerCopyAvailablePostScriptNames() unsafe.Pointer {
	return _CTFontManagerCopyAvailablePostScriptNames()
}

// Retrieves the font descriptors that were registered with the font manager.

// Retrieves the font descriptors that were registered with the font manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerCopyRegisteredFontDescriptors(_:_:)
func CTFontManagerCopyRegisteredFontDescriptors(scope unsafe.Pointer, enabled bool) unsafe.Pointer {
	return _CTFontManagerCopyRegisteredFontDescriptors(scope, enabled)
}

// Creates a font descriptor representing the font in the supplied data.
//
// Added in macOS 10.7.
// Creates a font descriptor representing the font in the supplied data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerCreateFontDescriptorFromData(_:)
func CTFontManagerCreateFontDescriptorFromData(data unsafe.Pointer) unsafe.Pointer {
	return _CTFontManagerCreateFontDescriptorFromData(data)
}

// Creates an array of font descriptors for the fonts in the supplied data.
//
// Added in macOS 10.13.
// Creates an array of font descriptors for the fonts in the supplied data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerCreateFontDescriptorsFromData(_:)
func CTFontManagerCreateFontDescriptorsFromData(data unsafe.Pointer) unsafe.Pointer {
	return _CTFontManagerCreateFontDescriptorsFromData(data)
}

// Returns an array of font descriptors representing each of the fonts in the specified URL.
//
// Added in macOS 10.6.
// Returns an array of font descriptors representing each of the fonts in the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerCreateFontDescriptorsFromURL(_:)
func CTFontManagerCreateFontDescriptorsFromURL(fileURL unsafe.Pointer) unsafe.Pointer {
	return _CTFontManagerCreateFontDescriptorsFromURL(fileURL)
}

// Creates a reference to a run loop source used to convey font requests from the Font Manager.
//
// Deprecated: This function was deprecated in macOS 11.0.
//
// Added in macOS 10.6.
// Creates a reference to a run loop source used to convey font requests from the Font Manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerCreateFontRequestRunLoopSource(_:_:)
func CTFontManagerCreateFontRequestRunLoopSource(sourceOrder unsafe.Pointer, createMatchesCallback unsafe.Pointer) unsafe.Pointer {
	return _CTFontManagerCreateFontRequestRunLoopSource(sourceOrder, createMatchesCallback)
}

// Enables or disables the matching font descriptors for font descriptor matching.
//
// Added in macOS 10.6.
// Enables or disables the matching font descriptors for font descriptor matching.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerEnableFontDescriptors(_:_:)
func CTFontManagerEnableFontDescriptors(descriptors unsafe.Pointer, enable bool) {
	_CTFontManagerEnableFontDescriptors(descriptors, enable)
}

// Gets the auto-activation setting for the specified bundle identifier.
//
// Added in macOS 10.6.
// Gets the auto-activation setting for the specified bundle identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerGetAutoActivationSetting(_:)
func CTFontManagerGetAutoActivationSetting(bundleIdentifier unsafe.Pointer) unsafe.Pointer {
	return _CTFontManagerGetAutoActivationSetting(bundleIdentifier)
}

// Returns the registration scope of the specified URL.
//
// Added in macOS 10.6.
// Returns the registration scope of the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerGetScopeForURL(_:)
func CTFontManagerGetScopeForURL(fontURL unsafe.Pointer) unsafe.Pointer {
	return _CTFontManagerGetScopeForURL(fontURL)
}

// Determines whether a file is in a supported font format.
//
// Added in macOS 10.6.
// Determines whether a file is in a supported font format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerIsSupportedFont(_:)
func CTFontManagerIsSupportedFont(fontURL unsafe.Pointer) bool {
	return _CTFontManagerIsSupportedFont(fontURL)
}

// Registers font descriptors with the font manager.
//
// Added in macOS 10.15.
// Registers font descriptors with the font manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerRegisterFontDescriptors(_:_:_:_:)
func CTFontManagerRegisterFontDescriptors(fontDescriptors unsafe.Pointer, scope unsafe.Pointer, enabled bool, registrationHandler bool) {
	_CTFontManagerRegisterFontDescriptors(fontDescriptors, scope, enabled, registrationHandler)
}

// Registers fonts from the specified font URLs with the font manager.
//
// Added in macOS 10.15.
// Registers fonts from the specified font URLs with the font manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerRegisterFontURLs(_:_:_:_:)
func CTFontManagerRegisterFontURLs(fontURLs unsafe.Pointer, scope unsafe.Pointer, enabled bool, registrationHandler bool) {
	_CTFontManagerRegisterFontURLs(fontURLs, scope, enabled, registrationHandler)
}

// Registers fonts from the specified font URL with the Font Manager. Registered fonts are discoverable through font descriptor matching.
//
// Added in macOS 10.6.
// Registers fonts from the specified font URL with the Font Manager. Registered fonts are discoverable through font descriptor matching.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerRegisterFontsForURL(_:_:_:)
func CTFontManagerRegisterFontsForURL(fontURL unsafe.Pointer, scope unsafe.Pointer, error_ unsafe.Pointer) bool {
	return _CTFontManagerRegisterFontsForURL(fontURL, scope, error_)
}

// Registers fonts from the specified array of font URLs with the Font Manager. Registered fonts are discoverable through font descriptor matching.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.6.
// Registers fonts from the specified array of font URLs with the Font Manager. Registered fonts are discoverable through font descriptor matching.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerRegisterFontsForURLs(_:_:_:)
func CTFontManagerRegisterFontsForURLs(fontURLs unsafe.Pointer, scope unsafe.Pointer, errors unsafe.Pointer) bool {
	return _CTFontManagerRegisterFontsForURLs(fontURLs, scope, errors)
}

// Registers named font assets in the specified bundle with the font manager.

// Registers named font assets in the specified bundle with the font manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerRegisterFontsWithAssetNames(_:_:_:_:_:)
func CTFontManagerRegisterFontsWithAssetNames(fontAssetNames unsafe.Pointer, bundle unsafe.Pointer, scope unsafe.Pointer, enabled bool, registrationHandler bool) {
	_CTFontManagerRegisterFontsWithAssetNames(fontAssetNames, bundle, scope, enabled, registrationHandler)
}

// Registers the specified graphics font with the font manager.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.8.
// Registers the specified graphics font with the font manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerRegisterGraphicsFont(_:_:)
func CTFontManagerRegisterGraphicsFont(font coregraphics.FontRef, error_ unsafe.Pointer) bool {
	return _CTFontManagerRegisterGraphicsFont(font, error_)
}

// Resolves font descriptors specified on input.

// Resolves font descriptors specified on input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerRequestFonts(_:_:)
func CTFontManagerRequestFonts(fontDescriptors unsafe.Pointer) {
	_CTFontManagerRequestFonts(fontDescriptors)
}

// Sets the auto-activation setting for the specified bundle identifier.
//
// Added in macOS 10.6.
// Sets the auto-activation setting for the specified bundle identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerSetAutoActivationSetting(_:_:)
func CTFontManagerSetAutoActivationSetting(bundleIdentifier unsafe.Pointer, setting unsafe.Pointer) {
	_CTFontManagerSetAutoActivationSetting(bundleIdentifier, setting)
}

// Unregisters font descriptors with the font manager.
//
// Added in macOS 10.15.
// Unregisters font descriptors with the font manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerUnregisterFontDescriptors(_:_:_:)
func CTFontManagerUnregisterFontDescriptors(fontDescriptors unsafe.Pointer, scope unsafe.Pointer, registrationHandler bool) {
	_CTFontManagerUnregisterFontDescriptors(fontDescriptors, scope, registrationHandler)
}

// Unregisters fonts from the specified font URLs with the font manager.
//
// Added in macOS 10.15.
// Unregisters fonts from the specified font URLs with the font manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerUnregisterFontURLs(_:_:_:)
func CTFontManagerUnregisterFontURLs(fontURLs unsafe.Pointer, scope unsafe.Pointer, registrationHandler bool) {
	_CTFontManagerUnregisterFontURLs(fontURLs, scope, registrationHandler)
}

// Unregisters fonts from the specified font URL with the Font Manager. Unregistered fonts are no longer discoverable through font descriptor matching.
//
// Added in macOS 10.6.
// Unregisters fonts from the specified font URL with the Font Manager. Unregistered fonts are no longer discoverable through font descriptor matching.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerUnregisterFontsForURL(_:_:_:)
func CTFontManagerUnregisterFontsForURL(fontURL unsafe.Pointer, scope unsafe.Pointer, error_ unsafe.Pointer) bool {
	return _CTFontManagerUnregisterFontsForURL(fontURL, scope, error_)
}

// Unregisters fonts from the specified array of font URLs with the Font Manager. Unregistered fonts are no longer discoverable through font descriptor matching.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.6.
// Unregisters fonts from the specified array of font URLs with the Font Manager. Unregistered fonts are no longer discoverable through font descriptor matching.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerUnregisterFontsForURLs(_:_:_:)
func CTFontManagerUnregisterFontsForURLs(fontURLs unsafe.Pointer, scope unsafe.Pointer, errors unsafe.Pointer) bool {
	return _CTFontManagerUnregisterFontsForURLs(fontURLs, scope, errors)
}

// Unregisters the specified graphics font with the font manager.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.8.
// Unregisters the specified graphics font with the font manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontManagerUnregisterGraphicsFont(_:_:)
func CTFontManagerUnregisterGraphicsFont(font coregraphics.FontRef, error_ unsafe.Pointer) bool {
	return _CTFontManagerUnregisterGraphicsFont(font, error_)
}

// Draws an entire frame into a context.
//
// Added in macOS 10.5.
// Draws an entire frame into a context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFrameDraw(_:_:)
func CTFrameDraw(frame unsafe.Pointer, context coregraphics.ContextRef) {
	_CTFrameDraw(frame, context)
}

// Returns the frame attributes used to create the frame.
//
// Added in macOS 10.5.
// Returns the frame attributes used to create the frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFrameGetFrameAttributes(_:)
func CTFrameGetFrameAttributes(frame unsafe.Pointer) unsafe.Pointer {
	return _CTFrameGetFrameAttributes(frame)
}

// Copies a range of line origins for a frame.
//
// Added in macOS 10.5.
// Copies a range of line origins for a frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFrameGetLineOrigins(_:_:_:)
func CTFrameGetLineOrigins(frame unsafe.Pointer, range_ unsafe.Pointer, origins coregraphics.CGPoint) {
	_CTFrameGetLineOrigins(frame, range_, origins)
}

// Returns an array of lines stored in the frame.
//
// Added in macOS 10.5.
// Returns an array of lines stored in the frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFrameGetLines(_:)
func CTFrameGetLines(frame unsafe.Pointer) unsafe.Pointer {
	return _CTFrameGetLines(frame)
}

// Returns the path used to create the frame.
//
// Added in macOS 10.5.
// Returns the path used to create the frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFrameGetPath(_:)
func CTFrameGetPath(frame unsafe.Pointer) coregraphics.PathRef {
	return _CTFrameGetPath(frame)
}

// Returns the range of characters originally requested to fill the frame.
//
// Added in macOS 10.5.
// Returns the range of characters originally requested to fill the frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFrameGetStringRange(_:)
func CTFrameGetStringRange(frame unsafe.Pointer) unsafe.Pointer {
	return _CTFrameGetStringRange(frame)
}

// Returns the type identifier for the CTFrame opaque type.
//
// Added in macOS 10.5.
// Returns the type identifier for the CTFrame opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFrameGetTypeID()
func CTFrameGetTypeID() unsafe.Pointer {
	return _CTFrameGetTypeID()
}

// Returns the range of characters that actually fit in the frame.
//
// Added in macOS 10.5.
// Returns the range of characters that actually fit in the frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFrameGetVisibleStringRange(_:)
func CTFrameGetVisibleStringRange(frame unsafe.Pointer) unsafe.Pointer {
	return _CTFrameGetVisibleStringRange(frame)
}

// Creates an immutable frame using a framesetter.
//
// Added in macOS 10.5.
// Creates an immutable frame using a framesetter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFramesetterCreateFrame(_:_:_:_:)
func CTFramesetterCreateFrame(framesetter unsafe.Pointer, stringRange unsafe.Pointer, path coregraphics.PathRef, frameAttributes unsafe.Pointer) unsafe.Pointer {
	return _CTFramesetterCreateFrame(framesetter, stringRange, path, frameAttributes)
}

// Creates an immutable framesetter object from an attributed string.
//
// Added in macOS 10.5.
// Creates an immutable framesetter object from an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFramesetterCreateWithAttributedString(_:)
func CTFramesetterCreateWithAttributedString(attrString unsafe.Pointer) unsafe.Pointer {
	return _CTFramesetterCreateWithAttributedString(attrString)
}

// Creates a framesetter directly from a typesetter.
//
// Added in macOS 10.14.
// Creates a framesetter directly from a typesetter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFramesetterCreateWithTypesetter(_:)
func CTFramesetterCreateWithTypesetter(typesetter unsafe.Pointer) unsafe.Pointer {
	return _CTFramesetterCreateWithTypesetter(typesetter)
}

// Returns the Core Foundation type identifier of the framesetter object.
//
// Added in macOS 10.5.
// Returns the Core Foundation type identifier of the framesetter object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFramesetterGetTypeID()
func CTFramesetterGetTypeID() unsafe.Pointer {
	return _CTFramesetterGetTypeID()
}

// Returns the typesetter object being used by the framesetter.
//
// Added in macOS 10.5.
// Returns the typesetter object being used by the framesetter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFramesetterGetTypesetter(_:)
func CTFramesetterGetTypesetter(framesetter unsafe.Pointer) unsafe.Pointer {
	return _CTFramesetterGetTypesetter(framesetter)
}

// Determines the frame size needed for a string range.
//
// Added in macOS 10.5.
// Determines the frame size needed for a string range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFramesetterSuggestFrameSizeWithConstraints(_:_:_:_:_:)
func CTFramesetterSuggestFrameSizeWithConstraints(framesetter unsafe.Pointer, stringRange unsafe.Pointer, frameAttributes unsafe.Pointer, constraints coregraphics.CGSize, fitRange unsafe.Pointer) coregraphics.CGSize {
	return _CTFramesetterSuggestFrameSizeWithConstraints(framesetter, stringRange, frameAttributes, constraints, fitRange)
}

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
}

// Creates an immutable glyph info object with a character identifier.
//
// Added in macOS 10.5.
// Creates an immutable glyph info object with a character identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTGlyphInfoCreateWithCharacterIdentifier(_:_:_:)
func CTGlyphInfoCreateWithCharacterIdentifier(cid unsafe.Pointer, collection unsafe.Pointer, baseString unsafe.Pointer) unsafe.Pointer {
	return _CTGlyphInfoCreateWithCharacterIdentifier(cid, collection, baseString)
}

// Creates an immutable glyph info object with a glyph index.
//
// Added in macOS 10.5.
// Creates an immutable glyph info object with a glyph index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTGlyphInfoCreateWithGlyph(_:_:_:)
func CTGlyphInfoCreateWithGlyph(glyph unsafe.Pointer, font unsafe.Pointer, baseString unsafe.Pointer) unsafe.Pointer {
	return _CTGlyphInfoCreateWithGlyph(glyph, font, baseString)
}

// Creates an immutable glyph info object with a glyph name.
//
// Added in macOS 10.5.
// Creates an immutable glyph info object with a glyph name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTGlyphInfoCreateWithGlyphName(_:_:_:)
func CTGlyphInfoCreateWithGlyphName(glyphName unsafe.Pointer, font unsafe.Pointer, baseString unsafe.Pointer) unsafe.Pointer {
	return _CTGlyphInfoCreateWithGlyphName(glyphName, font, baseString)
}

// Gets the character collection for a glyph info object.
//
// Added in macOS 10.5.
// Gets the character collection for a glyph info object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTGlyphInfoGetCharacterCollection(_:)
func CTGlyphInfoGetCharacterCollection(glyphInfo unsafe.Pointer) unsafe.Pointer {
	return _CTGlyphInfoGetCharacterCollection(glyphInfo)
}

// Gets the character identifier for a glyph info object.
//
// Added in macOS 10.5.
// Gets the character identifier for a glyph info object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTGlyphInfoGetCharacterIdentifier(_:)
func CTGlyphInfoGetCharacterIdentifier(glyphInfo unsafe.Pointer) unsafe.Pointer {
	return _CTGlyphInfoGetCharacterIdentifier(glyphInfo)
}

// Retrieves the glyph for a glyph info, if that object exists.
//
// Added in macOS 10.15.
// Retrieves the glyph for a glyph info, if that object exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTGlyphInfoGetGlyph(_:)
func CTGlyphInfoGetGlyph(glyphInfo unsafe.Pointer) unsafe.Pointer {
	return _CTGlyphInfoGetGlyph(glyphInfo)
}

// Retrieves the glyph name for a glyph info object, if that object exists.
//
// Added in macOS 10.5.
// Retrieves the glyph name for a glyph info object, if that object exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTGlyphInfoGetGlyphName(_:)
func CTGlyphInfoGetGlyphName(glyphInfo unsafe.Pointer) unsafe.Pointer {
	return _CTGlyphInfoGetGlyphName(glyphInfo)
}

// Returns the Core Foundation type identifier of the glyph info object
//
// Added in macOS 10.5.
// Returns the Core Foundation type identifier of the glyph info object
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTGlyphInfoGetTypeID()
func CTGlyphInfoGetTypeID() unsafe.Pointer {
	return _CTGlyphInfoGetTypeID()
}

// Creates a justified line from an existing line.
//
// Added in macOS 10.5.
// Creates a justified line from an existing line.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineCreateJustifiedLine(_:_:_:)
func CTLineCreateJustifiedLine(line unsafe.Pointer, justificationFactor float64, justificationWidth float64) unsafe.Pointer {
	return _CTLineCreateJustifiedLine(line, justificationFactor, justificationWidth)
}

// Creates a truncated line from an existing line.
//
// Added in macOS 10.5.
// Creates a truncated line from an existing line.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineCreateTruncatedLine(_:_:_:_:)
func CTLineCreateTruncatedLine(line unsafe.Pointer, width float64, truncationType unsafe.Pointer, truncationToken unsafe.Pointer) unsafe.Pointer {
	return _CTLineCreateTruncatedLine(line, width, truncationType, truncationToken)
}

// Creates a single immutable line object from an attributed string.
//
// Added in macOS 10.5.
// Creates a single immutable line object from an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineCreateWithAttributedString(_:)
func CTLineCreateWithAttributedString(attrString unsafe.Pointer) unsafe.Pointer {
	return _CTLineCreateWithAttributedString(attrString)
}

// Draws a complete line.
//
// Added in macOS 10.5.
// Draws a complete line.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineDraw(_:_:)
func CTLineDraw(line unsafe.Pointer, context coregraphics.ContextRef) {
	_CTLineDraw(line, context)
}

// Enumerates caret offsets for characters in a line.
//
// Added in macOS 10.11.
// Enumerates caret offsets for characters in a line.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineEnumerateCaretOffsets(_:_:)
func CTLineEnumerateCaretOffsets(line unsafe.Pointer) {
	_CTLineEnumerateCaretOffsets(line)
}

// Calculates the bounds for a line.
//
// Added in macOS 10.8.
// Calculates the bounds for a line.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineGetBoundsWithOptions(_:_:)
func CTLineGetBoundsWithOptions(line unsafe.Pointer, options unsafe.Pointer) coregraphics.CGRect {
	return _CTLineGetBoundsWithOptions(line, options)
}

// Returns the total glyph count for the line object.
//
// Added in macOS 10.5.
// Returns the total glyph count for the line object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineGetGlyphCount(_:)
func CTLineGetGlyphCount(line unsafe.Pointer) unsafe.Pointer {
	return _CTLineGetGlyphCount(line)
}

// Returns the array of glyph runs that make up the line object.
//
// Added in macOS 10.5.
// Returns the array of glyph runs that make up the line object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineGetGlyphRuns(_:)
func CTLineGetGlyphRuns(line unsafe.Pointer) unsafe.Pointer {
	return _CTLineGetGlyphRuns(line)
}

// Calculates the image bounds for a line.
//
// Added in macOS 10.5.
// Calculates the image bounds for a line.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineGetImageBounds(_:_:)
func CTLineGetImageBounds(line unsafe.Pointer, context coregraphics.ContextRef) coregraphics.CGRect {
	return _CTLineGetImageBounds(line, context)
}

// Determines the graphical offset or offsets for a string index.
//
// Added in macOS 10.5.
// Determines the graphical offset or offsets for a string index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineGetOffsetForStringIndex(_:_:_:)
func CTLineGetOffsetForStringIndex(line unsafe.Pointer, charIndex unsafe.Pointer, secondaryOffset []float64) float64 {
	return _CTLineGetOffsetForStringIndex(line, charIndex, secondaryOffset)
}

// Gets the pen offset required to draw flush text.
//
// Added in macOS 10.5.
// Gets the pen offset required to draw flush text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineGetPenOffsetForFlush(_:_:_:)
func CTLineGetPenOffsetForFlush(line unsafe.Pointer, flushFactor float64, flushWidth float64) float64 {
	return _CTLineGetPenOffsetForFlush(line, flushFactor, flushWidth)
}

// Performs hit testing.
//
// Added in macOS 10.5.
// Performs hit testing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineGetStringIndexForPosition(_:_:)
func CTLineGetStringIndexForPosition(line unsafe.Pointer, position coregraphics.CGPoint) unsafe.Pointer {
	return _CTLineGetStringIndexForPosition(line, position)
}

// Gets the range of characters that originally spawned the glyphs in the line.
//
// Added in macOS 10.5.
// Gets the range of characters that originally spawned the glyphs in the line.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineGetStringRange(_:)
func CTLineGetStringRange(line unsafe.Pointer) unsafe.Pointer {
	return _CTLineGetStringRange(line)
}

// Returns the trailing whitespace width for a line.
//
// Added in macOS 10.5.
// Returns the trailing whitespace width for a line.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineGetTrailingWhitespaceWidth(_:)
func CTLineGetTrailingWhitespaceWidth(line unsafe.Pointer) float64 {
	return _CTLineGetTrailingWhitespaceWidth(line)
}

// Returns the Core Foundation type identifier of the line object.
//
// Added in macOS 10.5.
// Returns the Core Foundation type identifier of the line object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineGetTypeID()
func CTLineGetTypeID() unsafe.Pointer {
	return _CTLineGetTypeID()
}

// Calculates the typographic bounds of a line.
//
// Added in macOS 10.5.
// Calculates the typographic bounds of a line.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLineGetTypographicBounds(_:_:_:_:)
func CTLineGetTypographicBounds(line unsafe.Pointer, ascent []float64, descent []float64, leading []float64) float64 {
	return _CTLineGetTypographicBounds(line, ascent, descent, leading)
}

// Creates an immutable paragraph style.
//
// Added in macOS 10.5.
// Creates an immutable paragraph style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleCreate(_:_:)
func CTParagraphStyleCreate(settings unsafe.Pointer, settingCount uintptr) unsafe.Pointer {
	return _CTParagraphStyleCreate(settings, settingCount)
}

// Creates an immutable copy of a paragraph style.
//
// Added in macOS 10.5.
// Creates an immutable copy of a paragraph style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleCreateCopy(_:)
func CTParagraphStyleCreateCopy(paragraphStyle unsafe.Pointer) unsafe.Pointer {
	return _CTParagraphStyleCreateCopy(paragraphStyle)
}

// Returns the Core Foundation type identifier of the paragraph style object.
//
// Added in macOS 10.5.
// Returns the Core Foundation type identifier of the paragraph style object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleGetTypeID()
func CTParagraphStyleGetTypeID() unsafe.Pointer {
	return _CTParagraphStyleGetTypeID()
}

// Obtains the current value for a single setting specifier.
//
// Added in macOS 10.5.
// Obtains the current value for a single setting specifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleGetValueForSpecifier(_:_:_:_:)
func CTParagraphStyleGetValueForSpecifier(paragraphStyle unsafe.Pointer, spec unsafe.Pointer, valueBufferSize uintptr, valueBuffer unsafe.Pointer) bool {
	return _CTParagraphStyleGetValueForSpecifier(paragraphStyle, spec, valueBufferSize, valueBuffer)
}

// Creates an immutable ruby annotation object.
//
// Added in macOS 10.10.
// Creates an immutable ruby annotation object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyAnnotationCreate(_:_:_:_:)
func CTRubyAnnotationCreate(alignment unsafe.Pointer, overhang unsafe.Pointer, sizeFactor float64, text unsafe.Pointer, p4 unsafe.Pointer) unsafe.Pointer {
	return _CTRubyAnnotationCreate(alignment, overhang, sizeFactor, text, p4)
}

// Creates an immutable copy of a ruby annotation object.
//
// Added in macOS 10.10.
// Creates an immutable copy of a ruby annotation object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyAnnotationCreateCopy(_:)
func CTRubyAnnotationCreateCopy(rubyAnnotation unsafe.Pointer) unsafe.Pointer {
	return _CTRubyAnnotationCreateCopy(rubyAnnotation)
}

// Creates an immutable ruby annotation object with the specified attributes.
//
// Added in macOS 10.12.
// Creates an immutable ruby annotation object with the specified attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyAnnotationCreateWithAttributes(_:_:_:_:_:)
func CTRubyAnnotationCreateWithAttributes(alignment unsafe.Pointer, overhang unsafe.Pointer, position unsafe.Pointer, string_ unsafe.Pointer, attributes unsafe.Pointer) unsafe.Pointer {
	return _CTRubyAnnotationCreateWithAttributes(alignment, overhang, position, string_, attributes)
}

// Retrieves the alignment value of a ruby annotation object.
//
// Added in macOS 10.10.
// Retrieves the alignment value of a ruby annotation object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyAnnotationGetAlignment(_:)
func CTRubyAnnotationGetAlignment(rubyAnnotation unsafe.Pointer) unsafe.Pointer {
	return _CTRubyAnnotationGetAlignment(rubyAnnotation)
}

// Retrieves the overhang value of a ruby annotation object.
//
// Added in macOS 10.10.
// Retrieves the overhang value of a ruby annotation object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyAnnotationGetOverhang(_:)
func CTRubyAnnotationGetOverhang(rubyAnnotation unsafe.Pointer) unsafe.Pointer {
	return _CTRubyAnnotationGetOverhang(rubyAnnotation)
}

// Retrieves the size factor of a ruby annotation object.
//
// Added in macOS 10.10.
// Retrieves the size factor of a ruby annotation object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyAnnotationGetSizeFactor(_:)
func CTRubyAnnotationGetSizeFactor(rubyAnnotation unsafe.Pointer) float64 {
	return _CTRubyAnnotationGetSizeFactor(rubyAnnotation)
}

// Retrieves the ruby text for a particular position in a ruby annotation.
//
// Added in macOS 10.10.
// Retrieves the ruby text for a particular position in a ruby annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyAnnotationGetTextForPosition(_:_:)
func CTRubyAnnotationGetTextForPosition(rubyAnnotation unsafe.Pointer, position unsafe.Pointer) unsafe.Pointer {
	return _CTRubyAnnotationGetTextForPosition(rubyAnnotation, position)
}

// Retrieves the type of the ruby annotation object.
//
// Added in macOS 10.10.
// Retrieves the type of the ruby annotation object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyAnnotationGetTypeID()
func CTRubyAnnotationGetTypeID() unsafe.Pointer {
	return _CTRubyAnnotationGetTypeID()
}

// Creates an immutable instance of a run delegate.
//
// Added in macOS 10.5.
// Creates an immutable instance of a run delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunDelegateCreate(_:_:)
func CTRunDelegateCreate(callbacks unsafe.Pointer, refCon unsafe.Pointer) unsafe.Pointer {
	return _CTRunDelegateCreate(callbacks, refCon)
}

// Returns a run delegate’s “refCon” value.
//
// Added in macOS 10.5.
// Returns a run delegate’s “refCon” value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunDelegateGetRefCon(_:)
func CTRunDelegateGetRefCon(runDelegate unsafe.Pointer) unsafe.Pointer {
	return _CTRunDelegateGetRefCon(runDelegate)
}

// Returns the type of CTRunDelegate objects.
//
// Added in macOS 10.5.
// Returns the type of CTRunDelegate objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunDelegateGetTypeID()
func CTRunDelegateGetTypeID() unsafe.Pointer {
	return _CTRunDelegateGetTypeID()
}

// Draws a complete run or part of one.
//
// Added in macOS 10.5.
// Draws a complete run or part of one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunDraw(_:_:_:)
func CTRunDraw(run unsafe.Pointer, context coregraphics.ContextRef, range_ unsafe.Pointer) {
	_CTRunDraw(run, context, range_)
}

// Copies a range of glyph advances into a user-provided buffer.
//
// Added in macOS 10.5.
// Copies a range of glyph advances into a user-provided buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunGetAdvances(_:_:_:)
func CTRunGetAdvances(run unsafe.Pointer, range_ unsafe.Pointer, buffer coregraphics.CGSize) {
	_CTRunGetAdvances(run, range_, buffer)
}

// Returns a direct pointer for the glyph advance array stored in the run.
//
// Added in macOS 10.5.
// Returns a direct pointer for the glyph advance array stored in the run.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunGetAdvancesPtr(_:)
func CTRunGetAdvancesPtr(run unsafe.Pointer) unsafe.Pointer {
	return _CTRunGetAdvancesPtr(run)
}

// Returns the attribute dictionary that was used to create the glyph run.
//
// Added in macOS 10.5.
// Returns the attribute dictionary that was used to create the glyph run.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunGetAttributes(_:)
func CTRunGetAttributes(run unsafe.Pointer) unsafe.Pointer {
	return _CTRunGetAttributes(run)
}

// Copies a range of base advances and origins into user-provided buffers.
//
// Added in macOS 10.11.
// Copies a range of base advances and origins into user-provided buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunGetBaseAdvancesAndOrigins(_:_:_:_:)
func CTRunGetBaseAdvancesAndOrigins(runRef unsafe.Pointer, range_ unsafe.Pointer, advancesBuffer coregraphics.CGSize, originsBuffer coregraphics.CGPoint) {
	_CTRunGetBaseAdvancesAndOrigins(runRef, range_, advancesBuffer, originsBuffer)
}

// Gets the glyph count for the run.
//
// Added in macOS 10.5.
// Gets the glyph count for the run.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunGetGlyphCount(_:)
func CTRunGetGlyphCount(run unsafe.Pointer) unsafe.Pointer {
	return _CTRunGetGlyphCount(run)
}

// Copies a range of glyphs into a user-provided buffer.
//
// Added in macOS 10.5.
// Copies a range of glyphs into a user-provided buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunGetGlyphs(_:_:_:)
func CTRunGetGlyphs(run unsafe.Pointer, range_ unsafe.Pointer, buffer unsafe.Pointer) {
	_CTRunGetGlyphs(run, range_, buffer)
}

// Returns a direct pointer for the glyph array stored in the run.
//
// Added in macOS 10.5.
// Returns a direct pointer for the glyph array stored in the run.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunGetGlyphsPtr(_:)
func CTRunGetGlyphsPtr(run unsafe.Pointer) unsafe.Pointer {
	return _CTRunGetGlyphsPtr(run)
}

// Calculates the image bounds for a glyph range.
//
// Added in macOS 10.5.
// Calculates the image bounds for a glyph range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunGetImageBounds(_:_:_:)
func CTRunGetImageBounds(run unsafe.Pointer, context coregraphics.ContextRef, range_ unsafe.Pointer) coregraphics.CGRect {
	return _CTRunGetImageBounds(run, context, range_)
}

// Copies a range of glyph positions into a user-provided buffer.
//
// Added in macOS 10.5.
// Copies a range of glyph positions into a user-provided buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunGetPositions(_:_:_:)
func CTRunGetPositions(run unsafe.Pointer, range_ unsafe.Pointer, buffer coregraphics.CGPoint) {
	_CTRunGetPositions(run, range_, buffer)
}

// Returns a direct pointer for the glyph position array stored in the run.
//
// Added in macOS 10.5.
// Returns a direct pointer for the glyph position array stored in the run.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunGetPositionsPtr(_:)
func CTRunGetPositionsPtr(run unsafe.Pointer) unsafe.Pointer {
	return _CTRunGetPositionsPtr(run)
}

// Returns the run’s status.
//
// Added in macOS 10.5.
// Returns the run’s status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunGetStatus(_:)
func CTRunGetStatus(run unsafe.Pointer) unsafe.Pointer {
	return _CTRunGetStatus(run)
}

// Copies a range of string indices into a user-provided buffer.
//
// Added in macOS 10.5.
// Copies a range of string indices into a user-provided buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunGetStringIndices(_:_:_:)
func CTRunGetStringIndices(run unsafe.Pointer, range_ unsafe.Pointer, buffer unsafe.Pointer) {
	_CTRunGetStringIndices(run, range_, buffer)
}

// Returns a direct pointer for the string indices stored in the run.
//
// Added in macOS 10.5.
// Returns a direct pointer for the string indices stored in the run.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunGetStringIndicesPtr(_:)
func CTRunGetStringIndicesPtr(run unsafe.Pointer) unsafe.Pointer {
	return _CTRunGetStringIndicesPtr(run)
}

// Gets the range of characters that originally spawned the glyphs in the run.
//
// Added in macOS 10.5.
// Gets the range of characters that originally spawned the glyphs in the run.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunGetStringRange(_:)
func CTRunGetStringRange(run unsafe.Pointer) unsafe.Pointer {
	return _CTRunGetStringRange(run)
}

// Returns the text matrix needed to draw this run.
//
// Added in macOS 10.5.
// Returns the text matrix needed to draw this run.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunGetTextMatrix(_:)
func CTRunGetTextMatrix(run unsafe.Pointer) coregraphics.CGAffineTransform {
	return _CTRunGetTextMatrix(run)
}

// Returns the Core Foundation type identifier of the run object.
//
// Added in macOS 10.5.
// Returns the Core Foundation type identifier of the run object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunGetTypeID()
func CTRunGetTypeID() unsafe.Pointer {
	return _CTRunGetTypeID()
}

// Gets the typographic bounds of the run.
//
// Added in macOS 10.5.
// Gets the typographic bounds of the run.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunGetTypographicBounds(_:_:_:_:_:)
func CTRunGetTypographicBounds(run unsafe.Pointer, range_ unsafe.Pointer, ascent []float64, descent []float64, leading []float64) float64 {
	return _CTRunGetTypographicBounds(run, range_, ascent, descent, leading)
}

// Creates and initializes a new text tab object.
//
// Added in macOS 10.5.
// Creates and initializes a new text tab object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTextTabCreate(_:_:_:)
func CTTextTabCreate(alignment unsafe.Pointer, location float64, options unsafe.Pointer) unsafe.Pointer {
	return _CTTextTabCreate(alignment, location, options)
}

// Returns the text alignment of the tab.
//
// Added in macOS 10.5.
// Returns the text alignment of the tab.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTextTabGetAlignment(_:)
func CTTextTabGetAlignment(tab unsafe.Pointer) unsafe.Pointer {
	return _CTTextTabGetAlignment(tab)
}

// Returns the tab’s ruler location.
//
// Added in macOS 10.5.
// Returns the tab’s ruler location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTextTabGetLocation(_:)
func CTTextTabGetLocation(tab unsafe.Pointer) float64 {
	return _CTTextTabGetLocation(tab)
}

// Returns the dictionary of attributes associated with the tab.
//
// Added in macOS 10.5.
// Returns the dictionary of attributes associated with the tab.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTextTabGetOptions(_:)
func CTTextTabGetOptions(tab unsafe.Pointer) unsafe.Pointer {
	return _CTTextTabGetOptions(tab)
}

// Returns the Core Foundation type identifier of the text tab object.
//
// Added in macOS 10.5.
// Returns the Core Foundation type identifier of the text tab object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTextTabGetTypeID()
func CTTextTabGetTypeID() unsafe.Pointer {
	return _CTTextTabGetTypeID()
}

// Creates an immutable line from the typesetter.
//
// Added in macOS 10.5.
// Creates an immutable line from the typesetter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTypesetterCreateLine(_:_:)
func CTTypesetterCreateLine(typesetter unsafe.Pointer, stringRange unsafe.Pointer) unsafe.Pointer {
	return _CTTypesetterCreateLine(typesetter, stringRange)
}

// Creates an immutable line from the typesetter at a specified line offset.
//
// Added in macOS 10.6.
// Creates an immutable line from the typesetter at a specified line offset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTypesetterCreateLineWithOffset(_:_:_:)
func CTTypesetterCreateLineWithOffset(typesetter unsafe.Pointer, stringRange unsafe.Pointer, offset float64) unsafe.Pointer {
	return _CTTypesetterCreateLineWithOffset(typesetter, stringRange, offset)
}

// Creates an immutable typesetter object using an attributed string.
//
// Added in macOS 10.5.
// Creates an immutable typesetter object using an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTypesetterCreateWithAttributedString(_:)
func CTTypesetterCreateWithAttributedString(string_ unsafe.Pointer) unsafe.Pointer {
	return _CTTypesetterCreateWithAttributedString(string_)
}

// Creates an immutable typesetter object using an attributed string and a dictionary of options.
//
// Added in macOS 10.5.
// Creates an immutable typesetter object using an attributed string and a dictionary of options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTypesetterCreateWithAttributedStringAndOptions(_:_:)
func CTTypesetterCreateWithAttributedStringAndOptions(string_ unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	return _CTTypesetterCreateWithAttributedStringAndOptions(string_, options)
}

// Returns the Core Foundation type identifier of the typesetter object.
//
// Added in macOS 10.5.
// Returns the Core Foundation type identifier of the typesetter object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTypesetterGetTypeID()
func CTTypesetterGetTypeID() unsafe.Pointer {
	return _CTTypesetterGetTypeID()
}

// Suggests a cluster line breakpoint based on the width provided.
//
// Added in macOS 10.5.
// Suggests a cluster line breakpoint based on the width provided.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTypesetterSuggestClusterBreak(_:_:_:)
func CTTypesetterSuggestClusterBreak(typesetter unsafe.Pointer, startIndex unsafe.Pointer, width float64) unsafe.Pointer {
	return _CTTypesetterSuggestClusterBreak(typesetter, startIndex, width)
}

// Suggests a cluster line breakpoint based on the specified width and line offset.
//
// Added in macOS 10.6.
// Suggests a cluster line breakpoint based on the specified width and line offset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTypesetterSuggestClusterBreakWithOffset(_:_:_:_:)
func CTTypesetterSuggestClusterBreakWithOffset(typesetter unsafe.Pointer, startIndex unsafe.Pointer, width float64, offset float64) unsafe.Pointer {
	return _CTTypesetterSuggestClusterBreakWithOffset(typesetter, startIndex, width, offset)
}

// Suggests a contextual line breakpoint based on the width provided.
//
// Added in macOS 10.5.
// Suggests a contextual line breakpoint based on the width provided.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTypesetterSuggestLineBreak(_:_:_:)
func CTTypesetterSuggestLineBreak(typesetter unsafe.Pointer, startIndex unsafe.Pointer, width float64) unsafe.Pointer {
	return _CTTypesetterSuggestLineBreak(typesetter, startIndex, width)
}

// Suggests a contextual line breakpoint based on the width provided and the specified offset.
//
// Added in macOS 10.6.
// Suggests a contextual line breakpoint based on the width provided and the specified offset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTypesetterSuggestLineBreakWithOffset(_:_:_:_:)
func CTTypesetterSuggestLineBreakWithOffset(typesetter unsafe.Pointer, startIndex unsafe.Pointer, width float64, offset float64) unsafe.Pointer {
	return _CTTypesetterSuggestLineBreakWithOffset(typesetter, startIndex, width, offset)
}



