// Code generated from Apple documentation for CoreText. DO NOT EDIT.

package coretext
import (
	"unsafe"
)


// C struct types
// ALMXGlyphEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/ALMXGlyphEntry
type ALMXGlyphEntry struct {
}// ALMXHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/ALMXHeader
type ALMXHeader struct {
}// AnchorPoint
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/AnchorPoint
type AnchorPoint struct {
}// AnchorPointTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/AnchorPointTable
type AnchorPointTable struct {
}// AnkrTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/AnkrTable
type AnkrTable struct {
}// BslnFormat0Part
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/BslnFormat0Part
type BslnFormat0Part struct {
}// BslnFormat1Part
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/BslnFormat1Part
type BslnFormat1Part struct {
}// BslnFormat2Part
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/BslnFormat2Part
type BslnFormat2Part struct {
}// BslnFormat3Part
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/BslnFormat3Part
type BslnFormat3Part struct {
}// BslnTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/BslnTable
type BslnTable struct {
}// CTParagraphStyleSetting - This structure is used to alter the paragraph style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSetting
type CTParagraphStyleSetting struct {
}// CTRunDelegateCallbacks - A structure holding pointers to callbacks implemented by the run delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunDelegateCallbacks
type CTRunDelegateCallbacks struct {
	Dealloc RunDelegateDeallocateCallback // The callback invoked when the retain count of a CTRunDelegate reaches 0 and the CTRunDelegate is deallocated. This callback may be  .
	GetAscent RunDelegateGetAscentCallback // The callback invoked to request the run delegate to determine and return the typographic ascent of glyphs in the run. This callback may be  , which is equivalent to a   callback that always returns 0.
	GetDescent RunDelegateGetDescentCallback // The callback invoked to request the run delegate to determine and return the typographic descent of glyphs in the run. This callback may be  , which is equivalent to a   callback that always returns 0.
	GetWidth RunDelegateGetWidthCallback // The callback invoked to request the run delegate to determine and return the typographic width of glyphs in the run. This callback may be  , which is equivalent to a   callback that always returns 0.
	Version unsafe.Pointer // The version number of the callbacks being passed in as a parameter to  . The initial version is  .
}// FontVariation
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/FontVariation
type FontVariation struct {
}// JustDirectionTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/JustDirectionTable
type JustDirectionTable struct {
}// JustPCAction
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/JustPCAction
type JustPCAction struct {
}// JustPCActionSubrecord
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/JustPCActionSubrecord
type JustPCActionSubrecord struct {
}// JustPCConditionalAddAction
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/JustPCConditionalAddAction
type JustPCConditionalAddAction struct {
}// JustPCDecompositionAction
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/JustPCDecompositionAction
type JustPCDecompositionAction struct {
}// JustPCDuctilityAction
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/JustPCDuctilityAction
type JustPCDuctilityAction struct {
}// JustPCGlyphRepeatAddAction
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/JustPCGlyphRepeatAddAction
type JustPCGlyphRepeatAddAction struct {
}// JustPostcompTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/JustPostcompTable
type JustPostcompTable struct {
}// JustTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/JustTable
type JustTable struct {
}// JustWidthDeltaEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/JustWidthDeltaEntry
type JustWidthDeltaEntry struct {
}// JustWidthDeltaGroup
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/JustWidthDeltaGroup
type JustWidthDeltaGroup struct {
}// KernIndexArrayHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernIndexArrayHeader
type KernIndexArrayHeader struct {
}// KernKerningPair
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernKerningPair
type KernKerningPair struct {
}// KernOffsetTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernOffsetTable
type KernOffsetTable struct {
}// KernOrderedListEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernOrderedListEntry
type KernOrderedListEntry struct {
}// KernOrderedListHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernOrderedListHeader
type KernOrderedListHeader struct {
}// KernSimpleArrayHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernSimpleArrayHeader
type KernSimpleArrayHeader struct {
}// KernStateEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernStateEntry
type KernStateEntry struct {
}// KernStateHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernStateHeader
type KernStateHeader struct {
}// KernSubtableHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernSubtableHeader
type KernSubtableHeader struct {
}// KernTableHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernTableHeader
type KernTableHeader struct {
}// KernVersion0Header
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernVersion0Header
type KernVersion0Header struct {
}// KernVersion0SubtableHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernVersion0SubtableHeader
type KernVersion0SubtableHeader struct {
}// KerxAnchorPointAction
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxAnchorPointAction
type KerxAnchorPointAction struct {
}// KerxControlPointAction
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxControlPointAction
type KerxControlPointAction struct {
}// KerxControlPointEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxControlPointEntry
type KerxControlPointEntry struct {
}// KerxControlPointHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxControlPointHeader
type KerxControlPointHeader struct {
}// KerxCoordinateAction
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxCoordinateAction
type KerxCoordinateAction struct {
}// KerxIndexArrayHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxIndexArrayHeader
type KerxIndexArrayHeader struct {
}// KerxKerningPair
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxKerningPair
type KerxKerningPair struct {
}// KerxOrderedListEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxOrderedListEntry
type KerxOrderedListEntry struct {
}// KerxOrderedListHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxOrderedListHeader
type KerxOrderedListHeader struct {
}// KerxSimpleArrayHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxSimpleArrayHeader
type KerxSimpleArrayHeader struct {
}// KerxStateEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxStateEntry
type KerxStateEntry struct {
}// KerxStateHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxStateHeader
type KerxStateHeader struct {
}// KerxSubtableHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxSubtableHeader
type KerxSubtableHeader struct {
}// KerxTableHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxTableHeader
type KerxTableHeader struct {
}// LcarCaretClassEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/LcarCaretClassEntry
type LcarCaretClassEntry struct {
}// LcarCaretTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/LcarCaretTable
type LcarCaretTable struct {
}// LtagStringRange
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/LtagStringRange
type LtagStringRange struct {
}// LtagTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/LtagTable
type LtagTable struct {
}// MortChain
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MortChain
type MortChain struct {
}// MortContextualSubtable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MortContextualSubtable
type MortContextualSubtable struct {
}// MortFeatureEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MortFeatureEntry
type MortFeatureEntry struct {
}// MortInsertionSubtable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MortInsertionSubtable
type MortInsertionSubtable struct {
}// MortLigatureSubtable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MortLigatureSubtable
type MortLigatureSubtable struct {
}// MortRearrangementSubtable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MortRearrangementSubtable
type MortRearrangementSubtable struct {
}// MortSubtable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MortSubtable
type MortSubtable struct {
}// MortSwashSubtable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MortSwashSubtable
type MortSwashSubtable struct {
}// MortTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MortTable
type MortTable struct {
}// MorxChain
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MorxChain
type MorxChain struct {
}// MorxContextualSubtable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MorxContextualSubtable
type MorxContextualSubtable struct {
}// MorxInsertionSubtable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MorxInsertionSubtable
type MorxInsertionSubtable struct {
}// MorxLigatureSubtable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MorxLigatureSubtable
type MorxLigatureSubtable struct {
}// MorxRearrangementSubtable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MorxRearrangementSubtable
type MorxRearrangementSubtable struct {
}// MorxSubtable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MorxSubtable
type MorxSubtable struct {
}// MorxTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MorxTable
type MorxTable struct {
}// OpbdSideValues
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/OpbdSideValues
type OpbdSideValues struct {
}// OpbdTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/OpbdTable
type OpbdTable struct {
}// PropLookupSegment
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/PropLookupSegment
type PropLookupSegment struct {
}// PropLookupSingle
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/PropLookupSingle
type PropLookupSingle struct {
}// PropTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/PropTable
type PropTable struct {
}// ROTAGlyphEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/ROTAGlyphEntry
type ROTAGlyphEntry struct {
}// ROTAHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/ROTAHeader
type ROTAHeader struct {
}// SFNTLookupArrayHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/SFNTLookupArrayHeader
type SFNTLookupArrayHeader struct {
}// SFNTLookupBinarySearchHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/SFNTLookupBinarySearchHeader
type SFNTLookupBinarySearchHeader struct {
}// SFNTLookupSegment
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/SFNTLookupSegment
type SFNTLookupSegment struct {
}// SFNTLookupSegmentHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/SFNTLookupSegmentHeader
type SFNTLookupSegmentHeader struct {
}// SFNTLookupSingle
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/SFNTLookupSingle
type SFNTLookupSingle struct {
}// SFNTLookupSingleHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/SFNTLookupSingleHeader
type SFNTLookupSingleHeader struct {
}// SFNTLookupTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/SFNTLookupTable
type SFNTLookupTable struct {
}// SFNTLookupTrimmedArrayHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/SFNTLookupTrimmedArrayHeader
type SFNTLookupTrimmedArrayHeader struct {
}// SFNTLookupVectorHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/SFNTLookupVectorHeader
type SFNTLookupVectorHeader struct {
}// STClassTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/STClassTable
type STClassTable struct {
}// STEntryOne
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/STEntryOne
type STEntryOne struct {
}// STEntryTwo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/STEntryTwo
type STEntryTwo struct {
}// STEntryZero
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/STEntryZero
type STEntryZero struct {
}// STHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/STHeader
type STHeader struct {
}// STXEntryOne
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/STXEntryOne
type STXEntryOne struct {
}// STXEntryTwo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/STXEntryTwo
type STXEntryTwo struct {
}// STXEntryZero
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/STXEntryZero
type STXEntryZero struct {
}// STXHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/STXHeader
type STXHeader struct {
}// TrakTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/TrakTable
type TrakTable struct {
}// TrakTableData
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/TrakTableData
type TrakTableData struct {
}// TrakTableEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/TrakTableEntry
type TrakTableEntry struct {
}// sfntCMapEncoding
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntCMapEncoding
type sfntCMapEncoding struct {
}// sfntCMapExtendedSubHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntCMapExtendedSubHeader
type sfntCMapExtendedSubHeader struct {
}// sfntCMapHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntCMapHeader
type sfntCMapHeader struct {
}// sfntCMapSubHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntCMapSubHeader
type sfntCMapSubHeader struct {
}// sfntDescriptorHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntDescriptorHeader
type sfntDescriptorHeader struct {
}// sfntDirectory
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntDirectory
type sfntDirectory struct {
}// sfntDirectoryEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntDirectoryEntry
type sfntDirectoryEntry struct {
}// sfntFeatureHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntFeatureHeader
type sfntFeatureHeader struct {
}// sfntFeatureName
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntFeatureName
type sfntFeatureName struct {
}// sfntFontDescriptor
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntFontDescriptor
type sfntFontDescriptor struct {
}// sfntFontFeatureSetting
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntFontFeatureSetting
type sfntFontFeatureSetting struct {
}// sfntFontRunFeature
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntFontRunFeature
type sfntFontRunFeature struct {
}// sfntInstance
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntInstance
type sfntInstance struct {
}// sfntNameHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntNameHeader
type sfntNameHeader struct {
}// sfntNameRecord
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntNameRecord
type sfntNameRecord struct {
}// sfntVariationAxis
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntVariationAxis
type sfntVariationAxis struct {
}// sfntVariationHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntVariationHeader
type sfntVariationHeader struct {
}



