// Code generated from Apple documentation for CoreText. DO NOT EDIT.

package coretext
import (
"unsafe"
)

// Type aliases and typedefs
// FontRef - A font object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFont
// CTFontRef has base type: const struct __CTFont *
type FontRef uintptr
// FontCollectionRef - A font collection.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCollection
// CTFontCollectionRef has base type: const struct __CTFontCollection *
type FontCollectionRef uintptr
// FontDescriptorRef - A font descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontDescriptor
// CTFontDescriptorRef has base type: const struct __CTFontDescriptor *
type FontDescriptorRef uintptr
// FrameRef - A frame.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFrame
// CTFrameRef has base type: const struct __CTFrame *
type FrameRef uintptr
// FramesetterRef - Generate text frames.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFramesetter
// CTFramesetterRef has base type: const struct __CTFramesetter *
type FramesetterRef uintptr
// GlyphInfoRef - Override a font’s specified mapping from Unicode to the glyph ID.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTGlyphInfo
// CTGlyphInfoRef has base type: const struct __CTGlyphInfo *
type GlyphInfoRef uintptr
// LineRef - A line of text.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTLine
// CTLineRef has base type: const struct __CTLine *
type LineRef uintptr
// ParagraphStyleRef - Paragraph or ruler attributes in an attributed string.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyle
// CTParagraphStyleRef has base type: const struct __CTParagraphStyle *
type ParagraphStyleRef uintptr
// RubyAnnotationRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRubyAnnotation
// CTRubyAnnotationRef has base type: const struct __CTRubyAnnotation *
type RubyAnnotationRef uintptr
// RunRef - A glyph run.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRun
// CTRunRef has base type: const struct __CTRun *
type RunRef uintptr
// RunDelegateRef - A run delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunDelegate
// CTRunDelegateRef has base type: const struct __CTRunDelegate *
type RunDelegateRef uintptr
// TextTabRef - A tab in a paragraph style, storing an alignment type and location.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTextTab
// CTTextTabRef has base type: const struct __CTTextTab *
type TextTabRef uintptr
// TypesetterRef - A typesetter which performs line layout.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTTypesetter
// CTTypesetterRef has base type: const struct __CTTypesetter *
type TypesetterRef uintptr
// ATSFontRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/ATSFontRef
// ATSFontRef has base type: UInt32
type ATSFontRef uintptr
// BslnBaselineClass type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/BslnBaselineClass
// BslnBaselineClass has base type: UInt32
type BslnBaselineClass uintptr
// BslnBaselineRecord type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/BslnBaselineRecord
// BslnBaselineRecord has base type: Fixed[32]
type BslnBaselineRecord uintptr
// BslnTableFormat type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/BslnTableFormat
// BslnTableFormat has base type: UInt16
type BslnTableFormat uintptr
// BslnTablePtr type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/BslnTablePtr
// BslnTablePtr has base type: BslnTable *
type BslnTablePtr uintptr
// FontCollectionSortDescriptorsCallback - The collection sorting callback type.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontCollectionSortDescriptorsCallback
// CTFontCollectionSortDescriptorsCallback is a callback function
// C type: enum CFComparisonResult (*)(const struct __CTFontDescriptor *, const struct __CTFontDescriptor *, void *)
type FontCollectionSortDescriptorsCallback = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) ComparisonResult
// FontPriority - The priority of font descriptors when resolving duplicates and sorting match results.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontPriority
// CTFontPriority has base type: uint32_t
type FontPriority uintptr
// FontTableTag - Font table tags provide access to font table data.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFontTableTag
// CTFontTableTag has base type: FourCharCode
type FontTableTag uintptr
// MutableFontCollectionRef - A reference to a mutable font collection.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTMutableFontCollection
// CTMutableFontCollectionRef has base type: struct __CTFontCollection *
type MutableFontCollectionRef uintptr
// RunDelegateDeallocateCallback - Defines a pointer to a function that is invoked when a CTRunDelegate object is deallocated.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunDelegateDeallocateCallback
// CTRunDelegateDeallocateCallback is a callback function
// C type: void (*)(void *)
type RunDelegateDeallocateCallback = func(unsafe.Pointer)
// RunDelegateGetAscentCallback - Defines a pointer to a function that determines typographic ascent of glyphs in the run.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunDelegateGetAscentCallback
// CTRunDelegateGetAscentCallback is a callback function
// C type: double (*)(void *)
type RunDelegateGetAscentCallback = func(unsafe.Pointer) float64
// RunDelegateGetDescentCallback - Defines a pointer to a function that determines typographic descent of glyphs in the run.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunDelegateGetDescentCallback
// CTRunDelegateGetDescentCallback is a callback function
// C type: double (*)(void *)
type RunDelegateGetDescentCallback = func(unsafe.Pointer) float64
// RunDelegateGetWidthCallback - Defines a pointer to a function that determines the typographic width of glyphs in the run.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunDelegateGetWidthCallback
// CTRunDelegateGetWidthCallback is a callback function
// C type: double (*)(void *)
type RunDelegateGetWidthCallback = func(unsafe.Pointer) float64
// FontLanguageCode type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/FontLanguageCode
// FontLanguageCode has base type: UInt32
type FontLanguageCode uintptr
// FontNameCode type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/FontNameCode
// FontNameCode has base type: UInt32
type FontNameCode uintptr
// FontPlatformCode type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/FontPlatformCode
// FontPlatformCode has base type: UInt32
type FontPlatformCode uintptr
// FontScriptCode type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/FontScriptCode
// FontScriptCode has base type: UInt32
type FontScriptCode uintptr
// JustificationFlags type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/JustificationFlags
// JustificationFlags has base type: UInt16
type JustificationFlags uintptr
// JustPCActionType type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/JustPCActionType
// JustPCActionType has base type: UInt16
type JustPCActionType uintptr
// JustPCUnconditionalAddAction type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/JustPCUnconditionalAddAction
// JustPCUnconditionalAddAction has base type: UInt16
type JustPCUnconditionalAddAction uintptr
// KernArrayOffset type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernArrayOffset
// KernArrayOffset has base type: UInt16
type KernArrayOffset uintptr
// KernKerningValue type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernKerningValue
// KernKerningValue has base type: SInt16
type KernKerningValue uintptr
// KernOffsetTablePtr type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernOffsetTablePtr
// KernOffsetTablePtr has base type: KernOffsetTable *
type KernOffsetTablePtr uintptr
// KernOrderedListEntryPtr type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernOrderedListEntryPtr
// KernOrderedListEntryPtr has base type: KernOrderedListEntry *
type KernOrderedListEntryPtr uintptr
// KernSubtableHeaderPtr type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernSubtableHeaderPtr
// KernSubtableHeaderPtr has base type: KernSubtableHeader *
type KernSubtableHeaderPtr uintptr
// KernSubtableInfo type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernSubtableInfo
// KernSubtableInfo has base type: UInt16
type KernSubtableInfo uintptr
// KernTableFormat type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernTableFormat
// KernTableFormat has base type: UInt8
type KernTableFormat uintptr
// KernTableHeaderHandle type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernTableHeaderHandle
// KernTableHeaderHandle has base type: KernTableHeaderPtr *
type KernTableHeaderHandle uintptr
// KernTableHeaderPtr type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernTableHeaderPtr
// KernTableHeaderPtr has base type: KernTableHeader *
type KernTableHeaderPtr uintptr
// KerxArrayOffset type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxArrayOffset
// KerxArrayOffset has base type: UInt32
type KerxArrayOffset uintptr
// KerxOrderedListEntryPtr type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxOrderedListEntryPtr
// KerxOrderedListEntryPtr has base type: KerxOrderedListEntry *
type KerxOrderedListEntryPtr uintptr
// KerxSubtableCoverage type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxSubtableCoverage
// KerxSubtableCoverage has base type: UInt32
type KerxSubtableCoverage uintptr
// KerxSubtableHeaderPtr type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxSubtableHeaderPtr
// KerxSubtableHeaderPtr has base type: KerxSubtableHeader *
type KerxSubtableHeaderPtr uintptr
// KerxTableHeaderHandle type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxTableHeaderHandle
// KerxTableHeaderHandle has base type: KerxTableHeaderPtr *
type KerxTableHeaderHandle uintptr
// KerxTableHeaderPtr type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxTableHeaderPtr
// KerxTableHeaderPtr has base type: KerxTableHeader *
type KerxTableHeaderPtr uintptr
// LcarCaretTablePtr type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/LcarCaretTablePtr
// LcarCaretTablePtr has base type: LcarCaretTable *
type LcarCaretTablePtr uintptr
// MortLigatureActionEntry type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MortLigatureActionEntry
// MortLigatureActionEntry has base type: UInt32
type MortLigatureActionEntry uintptr
// MortSubtableMaskFlags type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MortSubtableMaskFlags
// MortSubtableMaskFlags has base type: UInt32
type MortSubtableMaskFlags uintptr
// OpbdTableFormat type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/OpbdTableFormat
// OpbdTableFormat has base type: UInt16
type OpbdTableFormat uintptr
// PropCharProperties type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/PropCharProperties
// PropCharProperties has base type: UInt16
type PropCharProperties uintptr
// SFNTLookupKind type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/SFNTLookupKind
// SFNTLookupKind has base type: UInt32
type SFNTLookupKind uintptr
// SFNTLookupOffset type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/SFNTLookupOffset
// SFNTLookupOffset has base type: UInt16
type SFNTLookupOffset uintptr
// SFNTLookupTableFormat type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/SFNTLookupTableFormat
// SFNTLookupTableFormat has base type: UInt16
type SFNTLookupTableFormat uintptr
// SFNTLookupTableHandle type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/SFNTLookupTableHandle
// SFNTLookupTableHandle has base type: SFNTLookupTablePtr *
type SFNTLookupTableHandle uintptr
// SFNTLookupTablePtr type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/SFNTLookupTablePtr
// SFNTLookupTablePtr has base type: SFNTLookupTable *
type SFNTLookupTablePtr uintptr
// SFNTLookupValue type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/SFNTLookupValue
// SFNTLookupValue has base type: UInt16
type SFNTLookupValue uintptr
// STClass type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/STClass
// STClass has base type: UInt8
type STClass uintptr
// STEntryIndex type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/STEntryIndex
// STEntryIndex has base type: UInt8
type STEntryIndex uintptr
// STXClass type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/STXClass
// STXClass has base type: UInt16
type STXClass uintptr
// STXClassTable type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/STXClassTable
// STXClassTable has base type: SFNTLookupTable
type STXClassTable uintptr
// STXEntryIndex type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/STXEntryIndex
// STXEntryIndex has base type: UInt16
type STXEntryIndex uintptr
// STXStateIndex type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/STXStateIndex
// STXStateIndex has base type: UInt16
type STXStateIndex uintptr
// TrakValue type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/TrakValue
// TrakValue has base type: SInt16
type TrakValue uintptr

