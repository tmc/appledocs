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
	GlyphIndexOffset unsafe.Pointer
	HorizontalAdvance unsafe.Pointer
	VerticalAdvance unsafe.Pointer
	XOffsetToHOrigin unsafe.Pointer
	YOffsetToVOrigin unsafe.Pointer
}/* debug [types.gen.go/struct]: ALMXGlyphEntry */

// ALMXHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/ALMXHeader
type ALMXHeader struct {
	FirstGlyph unsafe.Pointer
	Flags unsafe.Pointer
	LastGlyph unsafe.Pointer
	Lookup SFNTLookupTable
	NMasters unsafe.Pointer
	Version unsafe.Pointer
}/* debug [types.gen.go/struct]: ALMXHeader */

// AnchorPoint
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/AnchorPoint
type AnchorPoint struct {
	X unsafe.Pointer
	Y unsafe.Pointer
}/* debug [types.gen.go/struct]: AnchorPoint */

// AnchorPointTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/AnchorPointTable
type AnchorPointTable struct {
	NPoints unsafe.Pointer
	Points AnchorPoint
}/* debug [types.gen.go/struct]: AnchorPointTable */

// AnkrTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/AnkrTable
type AnkrTable struct {
	AnchorPointTableOffset unsafe.Pointer
	Flags unsafe.Pointer
	LookupTableOffset unsafe.Pointer
	Version unsafe.Pointer
}/* debug [types.gen.go/struct]: AnkrTable */

// BslnFormat0Part
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/BslnFormat0Part
type BslnFormat0Part struct {
	Deltas unsafe.Pointer
}/* debug [types.gen.go/struct]: BslnFormat0Part */

// BslnFormat1Part
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/BslnFormat1Part
type BslnFormat1Part struct {
	Deltas unsafe.Pointer
	MappingData SFNTLookupTable
}/* debug [types.gen.go/struct]: BslnFormat1Part */

// BslnFormat2Part
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/BslnFormat2Part
type BslnFormat2Part struct {
	CtlPoints unsafe.Pointer
	StdGlyph unsafe.Pointer
}/* debug [types.gen.go/struct]: BslnFormat2Part */

// BslnFormat3Part
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/BslnFormat3Part
type BslnFormat3Part struct {
	CtlPoints unsafe.Pointer
	MappingData SFNTLookupTable
	StdGlyph unsafe.Pointer
}/* debug [types.gen.go/struct]: BslnFormat3Part */

// BslnTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/BslnTable
type BslnTable struct {
	DefaultBaseline unsafe.Pointer
	Format BslnTableFormat
	Parts unsafe.Pointer
	Version unsafe.Pointer
}/* debug [types.gen.go/struct]: BslnTable */

// CTParagraphStyleSetting - This structure is used to alter the paragraph style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSetting
type CTParagraphStyleSetting struct {
	Spec ParagraphStyleSpecifier // The specifier of the setting. See   for possible values.
	Value unsafe.Pointer // A reference to the value of the setting specified by the   field. The value must be in the proper range for the   value and at least as large as the size specified in  .
	ValueSize uintptr // The size of the value pointed to by the   field. This value must match the size of the value required by the   set in the   field.
}/* debug [types.gen.go/struct]: CTParagraphStyleSetting */

// CTRunDelegateCallbacks - A structure holding pointers to callbacks implemented by the run delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunDelegateCallbacks
type CTRunDelegateCallbacks struct {
	Dealloc RunDelegateDeallocateCallback // The callback invoked when the retain count of a CTRunDelegate reaches 0 and the CTRunDelegate is deallocated. This callback may be  .
	GetAscent RunDelegateGetAscentCallback // The callback invoked to request the run delegate to determine and return the typographic ascent of glyphs in the run. This callback may be  , which is equivalent to a   callback that always returns 0.
	GetDescent RunDelegateGetDescentCallback // The callback invoked to request the run delegate to determine and return the typographic descent of glyphs in the run. This callback may be  , which is equivalent to a   callback that always returns 0.
	GetWidth RunDelegateGetWidthCallback // The callback invoked to request the run delegate to determine and return the typographic width of glyphs in the run. This callback may be  , which is equivalent to a   callback that always returns 0.
	Version Index // The version number of the callbacks being passed in as a parameter to  . The initial version is  .
}/* debug [types.gen.go/struct]: CTRunDelegateCallbacks */

// FontVariation
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/FontVariation
type FontVariation struct {
	Name unsafe.Pointer
	Value unsafe.Pointer
}/* debug [types.gen.go/struct]: FontVariation */

// JustDirectionTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/JustDirectionTable
type JustDirectionTable struct {
	JustClass unsafe.Pointer
	Lookup SFNTLookupTable
	Postcomp unsafe.Pointer
	WidthDeltaClusters unsafe.Pointer
}/* debug [types.gen.go/struct]: JustDirectionTable */

// JustPCAction
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/JustPCAction
type JustPCAction struct {
	ActionCount unsafe.Pointer
	Actions JustPCActionSubrecord
}/* debug [types.gen.go/struct]: JustPCAction */

// JustPCActionSubrecord
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/JustPCActionSubrecord
type JustPCActionSubrecord struct {
	Data unsafe.Pointer
	Length unsafe.Pointer
	TheClass unsafe.Pointer
	TheType JustPCActionType
}/* debug [types.gen.go/struct]: JustPCActionSubrecord */

// JustPCConditionalAddAction
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/JustPCConditionalAddAction
type JustPCConditionalAddAction struct {
	AddGlyph unsafe.Pointer
	SubstGlyph unsafe.Pointer
	SubstThreshold unsafe.Pointer
}/* debug [types.gen.go/struct]: JustPCConditionalAddAction */

// JustPCDecompositionAction
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/JustPCDecompositionAction
type JustPCDecompositionAction struct {
	Count unsafe.Pointer
	Glyphs unsafe.Pointer
	LowerLimit unsafe.Pointer
	Order unsafe.Pointer
	UpperLimit unsafe.Pointer
}/* debug [types.gen.go/struct]: JustPCDecompositionAction */

// JustPCDuctilityAction
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/JustPCDuctilityAction
type JustPCDuctilityAction struct {
	DuctilityAxis unsafe.Pointer
	MaximumLimit unsafe.Pointer
	MinimumLimit unsafe.Pointer
	NoStretchValue unsafe.Pointer
}/* debug [types.gen.go/struct]: JustPCDuctilityAction */

// JustPCGlyphRepeatAddAction
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/JustPCGlyphRepeatAddAction
type JustPCGlyphRepeatAddAction struct {
	Flags unsafe.Pointer
	Glyph unsafe.Pointer
}/* debug [types.gen.go/struct]: JustPCGlyphRepeatAddAction */

// JustPostcompTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/JustPostcompTable
type JustPostcompTable struct {
	LookupTable SFNTLookupTable
}/* debug [types.gen.go/struct]: JustPostcompTable */

// JustTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/JustTable
type JustTable struct {
	Format unsafe.Pointer
	HorizHeaderOffset unsafe.Pointer
	Version unsafe.Pointer
	VertHeaderOffset unsafe.Pointer
}/* debug [types.gen.go/struct]: JustTable */

// JustWidthDeltaEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/JustWidthDeltaEntry
type JustWidthDeltaEntry struct {
	AfterGrowLimit unsafe.Pointer
	AfterShrinkLimit unsafe.Pointer
	BeforeGrowLimit unsafe.Pointer
	BeforeShrinkLimit unsafe.Pointer
	GrowFlags JustificationFlags
	JustClass unsafe.Pointer
	ShrinkFlags JustificationFlags
}/* debug [types.gen.go/struct]: JustWidthDeltaEntry */

// JustWidthDeltaGroup
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/JustWidthDeltaGroup
type JustWidthDeltaGroup struct {
	Count unsafe.Pointer
	Entries JustWidthDeltaEntry
}/* debug [types.gen.go/struct]: JustWidthDeltaGroup */

// KernIndexArrayHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernIndexArrayHeader
type KernIndexArrayHeader struct {
	Flags unsafe.Pointer
	GlyphCount unsafe.Pointer
	KernIndex unsafe.Pointer
	KernValue unsafe.Pointer
	KernValueCount unsafe.Pointer
	LeftClass unsafe.Pointer
	LeftClassCount unsafe.Pointer
	RightClass unsafe.Pointer
	RightClassCount unsafe.Pointer
}/* debug [types.gen.go/struct]: KernIndexArrayHeader */

// KernKerningPair
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernKerningPair
type KernKerningPair struct {
	Left unsafe.Pointer
	Right unsafe.Pointer
}/* debug [types.gen.go/struct]: KernKerningPair */

// KernOffsetTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernOffsetTable
type KernOffsetTable struct {
	FirstGlyph unsafe.Pointer
	NGlyphs unsafe.Pointer
	OffsetTable KernArrayOffset
}/* debug [types.gen.go/struct]: KernOffsetTable */

// KernOrderedListEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernOrderedListEntry
type KernOrderedListEntry struct {
	Pair KernKerningPair
	Value KernKerningValue
}/* debug [types.gen.go/struct]: KernOrderedListEntry */

// KernOrderedListHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernOrderedListHeader
type KernOrderedListHeader struct {
	EntrySelector unsafe.Pointer
	NPairs unsafe.Pointer
	RangeShift unsafe.Pointer
	SearchRange unsafe.Pointer
	Table unsafe.Pointer
}/* debug [types.gen.go/struct]: KernOrderedListHeader */

// KernSimpleArrayHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernSimpleArrayHeader
type KernSimpleArrayHeader struct {
	FirstTable unsafe.Pointer
	LeftOffsetTable unsafe.Pointer
	RightOffsetTable unsafe.Pointer
	RowWidth unsafe.Pointer
	TheArray KernArrayOffset
}/* debug [types.gen.go/struct]: KernSimpleArrayHeader */

// KernStateEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernStateEntry
type KernStateEntry struct {
	Flags unsafe.Pointer
	NewState unsafe.Pointer
}/* debug [types.gen.go/struct]: KernStateEntry */

// KernStateHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernStateHeader
type KernStateHeader struct {
	FirstTable unsafe.Pointer
	Header STHeader
	ValueTable unsafe.Pointer
}/* debug [types.gen.go/struct]: KernStateHeader */

// KernSubtableHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernSubtableHeader
type KernSubtableHeader struct {
	FsHeader unsafe.Pointer
	Length unsafe.Pointer
	StInfo KernSubtableInfo
	TupleIndex unsafe.Pointer
}/* debug [types.gen.go/struct]: KernSubtableHeader */

// KernTableHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernTableHeader
type KernTableHeader struct {
	FirstSubtable unsafe.Pointer
	NTables unsafe.Pointer
	Version unsafe.Pointer
}/* debug [types.gen.go/struct]: KernTableHeader */

// KernVersion0Header
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernVersion0Header
type KernVersion0Header struct {
	FirstSubtable unsafe.Pointer
	NTables unsafe.Pointer
	Version unsafe.Pointer
}/* debug [types.gen.go/struct]: KernVersion0Header */

// KernVersion0SubtableHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernVersion0SubtableHeader
type KernVersion0SubtableHeader struct {
	FsHeader unsafe.Pointer
	Length unsafe.Pointer
	StInfo KernSubtableInfo
	Version unsafe.Pointer
}/* debug [types.gen.go/struct]: KernVersion0SubtableHeader */

// KerxAnchorPointAction
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxAnchorPointAction
type KerxAnchorPointAction struct {
	CurrAnchorPoint unsafe.Pointer
	MarkAnchorPoint unsafe.Pointer
}/* debug [types.gen.go/struct]: KerxAnchorPointAction */

// KerxControlPointAction
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxControlPointAction
type KerxControlPointAction struct {
	CurrControlPoint unsafe.Pointer
	MarkControlPoint unsafe.Pointer
}/* debug [types.gen.go/struct]: KerxControlPointAction */

// KerxControlPointEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxControlPointEntry
type KerxControlPointEntry struct {
	ActionIndex unsafe.Pointer
	Flags unsafe.Pointer
	NewState unsafe.Pointer
}/* debug [types.gen.go/struct]: KerxControlPointEntry */

// KerxControlPointHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxControlPointHeader
type KerxControlPointHeader struct {
	FirstTable unsafe.Pointer
	Flags unsafe.Pointer
	Header STXHeader
}/* debug [types.gen.go/struct]: KerxControlPointHeader */

// KerxCoordinateAction
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxCoordinateAction
type KerxCoordinateAction struct {
	CurrX unsafe.Pointer
	CurrY unsafe.Pointer
	MarkX unsafe.Pointer
	MarkY unsafe.Pointer
}/* debug [types.gen.go/struct]: KerxCoordinateAction */

// KerxIndexArrayHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxIndexArrayHeader
type KerxIndexArrayHeader struct {
	ColumnCount unsafe.Pointer
	ColumnIndexTableOffset unsafe.Pointer
	Flags unsafe.Pointer
	KerningArrayOffset unsafe.Pointer
	KerningVectorOffset unsafe.Pointer
	RowCount unsafe.Pointer
	RowIndexTableOffset unsafe.Pointer
}/* debug [types.gen.go/struct]: KerxIndexArrayHeader */

// KerxKerningPair
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxKerningPair
type KerxKerningPair struct {
	Left unsafe.Pointer
	Right unsafe.Pointer
}/* debug [types.gen.go/struct]: KerxKerningPair */

// KerxOrderedListEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxOrderedListEntry
type KerxOrderedListEntry struct {
	Pair KerxKerningPair
	Value KernKerningValue
}/* debug [types.gen.go/struct]: KerxOrderedListEntry */

// KerxOrderedListHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxOrderedListHeader
type KerxOrderedListHeader struct {
	EntrySelector unsafe.Pointer
	NPairs unsafe.Pointer
	RangeShift unsafe.Pointer
	SearchRange unsafe.Pointer
	Table unsafe.Pointer
}/* debug [types.gen.go/struct]: KerxOrderedListHeader */

// KerxSimpleArrayHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxSimpleArrayHeader
type KerxSimpleArrayHeader struct {
	FirstTable unsafe.Pointer
	LeftOffsetTable unsafe.Pointer
	RightOffsetTable unsafe.Pointer
	RowWidth unsafe.Pointer
	TheArray KerxArrayOffset
}/* debug [types.gen.go/struct]: KerxSimpleArrayHeader */

// KerxStateEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxStateEntry
type KerxStateEntry struct {
	Flags unsafe.Pointer
	NewState unsafe.Pointer
	ValueIndex unsafe.Pointer
}/* debug [types.gen.go/struct]: KerxStateEntry */

// KerxStateHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxStateHeader
type KerxStateHeader struct {
	FirstTable unsafe.Pointer
	Header STXHeader
	ValueTable unsafe.Pointer
}/* debug [types.gen.go/struct]: KerxStateHeader */

// KerxSubtableHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxSubtableHeader
type KerxSubtableHeader struct {
	FsHeader unsafe.Pointer
	Length unsafe.Pointer
	StInfo KerxSubtableCoverage
	TupleCount unsafe.Pointer
}/* debug [types.gen.go/struct]: KerxSubtableHeader */

// KerxTableHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxTableHeader
type KerxTableHeader struct {
	FirstSubtable unsafe.Pointer
	NTables unsafe.Pointer
	Version unsafe.Pointer
}/* debug [types.gen.go/struct]: KerxTableHeader */

// LcarCaretClassEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/LcarCaretClassEntry
type LcarCaretClassEntry struct {
	Count unsafe.Pointer
	Partials unsafe.Pointer
}/* debug [types.gen.go/struct]: LcarCaretClassEntry */

// LcarCaretTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/LcarCaretTable
type LcarCaretTable struct {
	Format unsafe.Pointer
	Lookup SFNTLookupTable
	Version unsafe.Pointer
}/* debug [types.gen.go/struct]: LcarCaretTable */

// LtagStringRange
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/LtagStringRange
type LtagStringRange struct {
	Length unsafe.Pointer
	Offset unsafe.Pointer
}/* debug [types.gen.go/struct]: LtagStringRange */

// LtagTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/LtagTable
type LtagTable struct {
	Flags unsafe.Pointer
	NumTags unsafe.Pointer
	TagRange LtagStringRange
	Version unsafe.Pointer
}/* debug [types.gen.go/struct]: LtagTable */

// MortChain
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MortChain
type MortChain struct {
	DefaultFlags MortSubtableMaskFlags
	FeatureEntries MortFeatureEntry
	Length unsafe.Pointer
	NFeatures unsafe.Pointer
	NSubtables unsafe.Pointer
}/* debug [types.gen.go/struct]: MortChain */

// MortContextualSubtable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MortContextualSubtable
type MortContextualSubtable struct {
	Header STHeader
	SubstitutionTableOffset unsafe.Pointer
}/* debug [types.gen.go/struct]: MortContextualSubtable */

// MortFeatureEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MortFeatureEntry
type MortFeatureEntry struct {
	DisableFlags MortSubtableMaskFlags
	EnableFlags MortSubtableMaskFlags
	FeatureSelector unsafe.Pointer
	FeatureType unsafe.Pointer
}/* debug [types.gen.go/struct]: MortFeatureEntry */

// MortInsertionSubtable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MortInsertionSubtable
type MortInsertionSubtable struct {
	Header STHeader
}/* debug [types.gen.go/struct]: MortInsertionSubtable */

// MortLigatureSubtable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MortLigatureSubtable
type MortLigatureSubtable struct {
	ComponentTableOffset unsafe.Pointer
	Header STHeader
	LigatureActionTableOffset unsafe.Pointer
	LigatureTableOffset unsafe.Pointer
}/* debug [types.gen.go/struct]: MortLigatureSubtable */

// MortRearrangementSubtable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MortRearrangementSubtable
type MortRearrangementSubtable struct {
	Header STHeader
}/* debug [types.gen.go/struct]: MortRearrangementSubtable */

// MortSubtable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MortSubtable
type MortSubtable struct {
	Coverage unsafe.Pointer
	Flags MortSubtableMaskFlags
	Length unsafe.Pointer
	U unsafe.Pointer
}/* debug [types.gen.go/struct]: MortSubtable */

// MortSwashSubtable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MortSwashSubtable
type MortSwashSubtable struct {
	Lookup SFNTLookupTable
}/* debug [types.gen.go/struct]: MortSwashSubtable */

// MortTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MortTable
type MortTable struct {
	Chains MortChain
	NChains unsafe.Pointer
	Version unsafe.Pointer
}/* debug [types.gen.go/struct]: MortTable */

// MorxChain
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MorxChain
type MorxChain struct {
	DefaultFlags MortSubtableMaskFlags
	FeatureEntries MortFeatureEntry
	Length unsafe.Pointer
	NFeatures unsafe.Pointer
	NSubtables unsafe.Pointer
}/* debug [types.gen.go/struct]: MorxChain */

// MorxContextualSubtable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MorxContextualSubtable
type MorxContextualSubtable struct {
	Header STXHeader
	SubstitutionTableOffset unsafe.Pointer
}/* debug [types.gen.go/struct]: MorxContextualSubtable */

// MorxInsertionSubtable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MorxInsertionSubtable
type MorxInsertionSubtable struct {
	Header STXHeader
	InsertionGlyphTableOffset unsafe.Pointer
}/* debug [types.gen.go/struct]: MorxInsertionSubtable */

// MorxLigatureSubtable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MorxLigatureSubtable
type MorxLigatureSubtable struct {
	ComponentTableOffset unsafe.Pointer
	Header STXHeader
	LigatureActionTableOffset unsafe.Pointer
	LigatureTableOffset unsafe.Pointer
}/* debug [types.gen.go/struct]: MorxLigatureSubtable */

// MorxRearrangementSubtable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MorxRearrangementSubtable
type MorxRearrangementSubtable struct {
	Header STXHeader
}/* debug [types.gen.go/struct]: MorxRearrangementSubtable */

// MorxSubtable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MorxSubtable
type MorxSubtable struct {
	Coverage unsafe.Pointer
	Flags MortSubtableMaskFlags
	Length unsafe.Pointer
	U unsafe.Pointer
}/* debug [types.gen.go/struct]: MorxSubtable */

// MorxTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MorxTable
type MorxTable struct {
	Chains MorxChain
	NChains unsafe.Pointer
	Version unsafe.Pointer
}/* debug [types.gen.go/struct]: MorxTable */

// OpbdSideValues
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/OpbdSideValues
type OpbdSideValues struct {
	BottomSideShift unsafe.Pointer
	LeftSideShift unsafe.Pointer
	RightSideShift unsafe.Pointer
	TopSideShift unsafe.Pointer
}/* debug [types.gen.go/struct]: OpbdSideValues */

// OpbdTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/OpbdTable
type OpbdTable struct {
	Format OpbdTableFormat
	LookupTable SFNTLookupTable
	Version unsafe.Pointer
}/* debug [types.gen.go/struct]: OpbdTable */

// PropLookupSegment
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/PropLookupSegment
type PropLookupSegment struct {
	FirstGlyph unsafe.Pointer
	LastGlyph unsafe.Pointer
	Value unsafe.Pointer
}/* debug [types.gen.go/struct]: PropLookupSegment */

// PropLookupSingle
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/PropLookupSingle
type PropLookupSingle struct {
	Glyph unsafe.Pointer
	Props PropCharProperties
}/* debug [types.gen.go/struct]: PropLookupSingle */

// PropTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/PropTable
type PropTable struct {
	DefaultProps PropCharProperties
	Format unsafe.Pointer
	Lookup SFNTLookupTable
	Version unsafe.Pointer
}/* debug [types.gen.go/struct]: PropTable */

// ROTAGlyphEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/ROTAGlyphEntry
type ROTAGlyphEntry struct {
	GlyphIndexOffset unsafe.Pointer
	HBaselineOffset unsafe.Pointer
	VBaselineOffset unsafe.Pointer
}/* debug [types.gen.go/struct]: ROTAGlyphEntry */

// ROTAHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/ROTAHeader
type ROTAHeader struct {
	FirstGlyph unsafe.Pointer
	Flags unsafe.Pointer
	LastGlyph unsafe.Pointer
	Lookup SFNTLookupTable
	NMasters unsafe.Pointer
	Version unsafe.Pointer
}/* debug [types.gen.go/struct]: ROTAHeader */

// sfntCMapEncoding
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntCMapEncoding
type sfntCMapEncoding struct {
	Offset unsafe.Pointer
	PlatformID unsafe.Pointer
	ScriptID unsafe.Pointer
}/* debug [types.gen.go/struct]: sfntCMapEncoding */

// sfntCMapExtendedSubHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntCMapExtendedSubHeader
type sfntCMapExtendedSubHeader struct {
	Format unsafe.Pointer
	Language unsafe.Pointer
	Length unsafe.Pointer
	Reserved unsafe.Pointer
}/* debug [types.gen.go/struct]: sfntCMapExtendedSubHeader */

// sfntCMapHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntCMapHeader
type sfntCMapHeader struct {
	Encoding SfntCMapEncoding
	NumTables unsafe.Pointer
	Version unsafe.Pointer
}/* debug [types.gen.go/struct]: sfntCMapHeader */

// sfntCMapSubHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntCMapSubHeader
type sfntCMapSubHeader struct {
	Format unsafe.Pointer
	LanguageID unsafe.Pointer
	Length unsafe.Pointer
}/* debug [types.gen.go/struct]: sfntCMapSubHeader */

// sfntDescriptorHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntDescriptorHeader
type sfntDescriptorHeader struct {
	Descriptor SfntFontDescriptor
	DescriptorCount unsafe.Pointer
	Version unsafe.Pointer
}/* debug [types.gen.go/struct]: sfntDescriptorHeader */

// sfntDirectory
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntDirectory
type sfntDirectory struct {
	EntrySelector unsafe.Pointer
	Format unsafe.Pointer
	NumOffsets unsafe.Pointer
	RangeShift unsafe.Pointer
	SearchRange unsafe.Pointer
	Table SfntDirectoryEntry
}/* debug [types.gen.go/struct]: sfntDirectory */

// sfntDirectoryEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntDirectoryEntry
type sfntDirectoryEntry struct {
	CheckSum unsafe.Pointer
	Length unsafe.Pointer
	Offset unsafe.Pointer
	TableTag unsafe.Pointer
}/* debug [types.gen.go/struct]: sfntDirectoryEntry */

// sfntFeatureHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntFeatureHeader
type sfntFeatureHeader struct {
	FeatureNameCount unsafe.Pointer
	FeatureSetCount unsafe.Pointer
	Names SfntFeatureName
	Reserved unsafe.Pointer
	Runs SfntFontRunFeature
	Settings SfntFontFeatureSetting
	Version unsafe.Pointer
}/* debug [types.gen.go/struct]: sfntFeatureHeader */

// sfntFeatureName
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntFeatureName
type sfntFeatureName struct {
	FeatureFlags unsafe.Pointer
	FeatureType unsafe.Pointer
	NameID unsafe.Pointer
	OffsetToSettings unsafe.Pointer
	SettingCount unsafe.Pointer
}/* debug [types.gen.go/struct]: sfntFeatureName */

// sfntFontDescriptor
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntFontDescriptor
type sfntFontDescriptor struct {
	Name unsafe.Pointer
	Value unsafe.Pointer
}/* debug [types.gen.go/struct]: sfntFontDescriptor */

// sfntFontFeatureSetting
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntFontFeatureSetting
type sfntFontFeatureSetting struct {
	NameID unsafe.Pointer
	Setting unsafe.Pointer
}/* debug [types.gen.go/struct]: sfntFontFeatureSetting */

// sfntFontRunFeature
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntFontRunFeature
type sfntFontRunFeature struct {
	FeatureType unsafe.Pointer
	Setting unsafe.Pointer
}/* debug [types.gen.go/struct]: sfntFontRunFeature */

// sfntInstance
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntInstance
type sfntInstance struct {
	Coord unsafe.Pointer
	Flags unsafe.Pointer
	NameID unsafe.Pointer
}/* debug [types.gen.go/struct]: sfntInstance */

// SFNTLookupArrayHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/SFNTLookupArrayHeader
type SFNTLookupArrayHeader struct {
	LookupValues SFNTLookupValue
}/* debug [types.gen.go/struct]: SFNTLookupArrayHeader */

// SFNTLookupBinarySearchHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/SFNTLookupBinarySearchHeader
type SFNTLookupBinarySearchHeader struct {
	EntrySelector unsafe.Pointer
	NUnits unsafe.Pointer
	RangeShift unsafe.Pointer
	SearchRange unsafe.Pointer
	UnitSize unsafe.Pointer
}/* debug [types.gen.go/struct]: SFNTLookupBinarySearchHeader */

// SFNTLookupSegment
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/SFNTLookupSegment
type SFNTLookupSegment struct {
	FirstGlyph unsafe.Pointer
	LastGlyph unsafe.Pointer
	Value unsafe.Pointer
}/* debug [types.gen.go/struct]: SFNTLookupSegment */

// SFNTLookupSegmentHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/SFNTLookupSegmentHeader
type SFNTLookupSegmentHeader struct {
	BinSearch SFNTLookupBinarySearchHeader
	Segments SFNTLookupSegment
}/* debug [types.gen.go/struct]: SFNTLookupSegmentHeader */

// SFNTLookupSingle
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/SFNTLookupSingle
type SFNTLookupSingle struct {
	Glyph unsafe.Pointer
	Value unsafe.Pointer
}/* debug [types.gen.go/struct]: SFNTLookupSingle */

// SFNTLookupSingleHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/SFNTLookupSingleHeader
type SFNTLookupSingleHeader struct {
	BinSearch SFNTLookupBinarySearchHeader
	Entries SFNTLookupSingle
}/* debug [types.gen.go/struct]: SFNTLookupSingleHeader */

// SFNTLookupTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/SFNTLookupTable
type SFNTLookupTable struct {
	Format SFNTLookupTableFormat
	FsHeader unsafe.Pointer
}/* debug [types.gen.go/struct]: SFNTLookupTable */

// SFNTLookupTrimmedArrayHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/SFNTLookupTrimmedArrayHeader
type SFNTLookupTrimmedArrayHeader struct {
	Count unsafe.Pointer
	FirstGlyph unsafe.Pointer
	ValueArray SFNTLookupValue
}/* debug [types.gen.go/struct]: SFNTLookupTrimmedArrayHeader */

// SFNTLookupVectorHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/SFNTLookupVectorHeader
type SFNTLookupVectorHeader struct {
	Count unsafe.Pointer
	FirstGlyph unsafe.Pointer
	Values unsafe.Pointer
	ValueSize unsafe.Pointer
}/* debug [types.gen.go/struct]: SFNTLookupVectorHeader */

// sfntNameHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntNameHeader
type sfntNameHeader struct {
	Count unsafe.Pointer
	Format unsafe.Pointer
	Rec SfntNameRecord
	StringOffset unsafe.Pointer
}/* debug [types.gen.go/struct]: sfntNameHeader */

// sfntNameRecord
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntNameRecord
type sfntNameRecord struct {
	LanguageID unsafe.Pointer
	Length unsafe.Pointer
	NameID unsafe.Pointer
	Offset unsafe.Pointer
	PlatformID unsafe.Pointer
	ScriptID unsafe.Pointer
}/* debug [types.gen.go/struct]: sfntNameRecord */

// sfntVariationAxis
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntVariationAxis
type sfntVariationAxis struct {
	AxisTag unsafe.Pointer
	DefaultValue unsafe.Pointer
	Flags unsafe.Pointer
	MaxValue unsafe.Pointer
	MinValue unsafe.Pointer
	NameID unsafe.Pointer
}/* debug [types.gen.go/struct]: sfntVariationAxis */

// sfntVariationHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntVariationHeader
type sfntVariationHeader struct {
	Axis SfntVariationAxis
	AxisCount unsafe.Pointer
	AxisSize unsafe.Pointer
	CountSizePairs unsafe.Pointer
	Instance SfntInstance
	InstanceCount unsafe.Pointer
	InstanceSize unsafe.Pointer
	OffsetToData unsafe.Pointer
	Version unsafe.Pointer
}/* debug [types.gen.go/struct]: sfntVariationHeader */

// STClassTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/STClassTable
type STClassTable struct {
	Classes STClass
	FirstGlyph unsafe.Pointer
	NGlyphs unsafe.Pointer
}/* debug [types.gen.go/struct]: STClassTable */

// STEntryOne
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/STEntryOne
type STEntryOne struct {
	Flags unsafe.Pointer
	NewState unsafe.Pointer
	Offset1 unsafe.Pointer
}/* debug [types.gen.go/struct]: STEntryOne */

// STEntryTwo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/STEntryTwo
type STEntryTwo struct {
	Flags unsafe.Pointer
	NewState unsafe.Pointer
	Offset1 unsafe.Pointer
	Offset2 unsafe.Pointer
}/* debug [types.gen.go/struct]: STEntryTwo */

// STEntryZero
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/STEntryZero
type STEntryZero struct {
	Flags unsafe.Pointer
	NewState unsafe.Pointer
}/* debug [types.gen.go/struct]: STEntryZero */

// STHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/STHeader
type STHeader struct {
	ClassTableOffset unsafe.Pointer
	EntryTableOffset unsafe.Pointer
	Filler unsafe.Pointer
	NClasses STClass
	StateArrayOffset unsafe.Pointer
}/* debug [types.gen.go/struct]: STHeader */

// STXEntryOne
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/STXEntryOne
type STXEntryOne struct {
	Flags unsafe.Pointer
	Index1 unsafe.Pointer
	NewState STXStateIndex
}/* debug [types.gen.go/struct]: STXEntryOne */

// STXEntryTwo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/STXEntryTwo
type STXEntryTwo struct {
	Flags unsafe.Pointer
	Index1 unsafe.Pointer
	Index2 unsafe.Pointer
	NewState STXStateIndex
}/* debug [types.gen.go/struct]: STXEntryTwo */

// STXEntryZero
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/STXEntryZero
type STXEntryZero struct {
	Flags unsafe.Pointer
	NewState STXStateIndex
}/* debug [types.gen.go/struct]: STXEntryZero */

// STXHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/STXHeader
type STXHeader struct {
	ClassTableOffset unsafe.Pointer
	EntryTableOffset unsafe.Pointer
	NClasses unsafe.Pointer
	StateArrayOffset unsafe.Pointer
}/* debug [types.gen.go/struct]: STXHeader */

// TrakTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/TrakTable
type TrakTable struct {
	Format unsafe.Pointer
	HorizOffset unsafe.Pointer
	Version unsafe.Pointer
	VertOffset unsafe.Pointer
}/* debug [types.gen.go/struct]: TrakTable */

// TrakTableData
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/TrakTableData
type TrakTableData struct {
	NSizes unsafe.Pointer
	NTracks unsafe.Pointer
	SizeTableOffset unsafe.Pointer
	TrakTable TrakTableEntry
}/* debug [types.gen.go/struct]: TrakTableData */

// TrakTableEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/TrakTableEntry
type TrakTableEntry struct {
	NameTableIndex unsafe.Pointer
	SizesOffset unsafe.Pointer
	Track unsafe.Pointer
}/* debug [types.gen.go/struct]: TrakTableEntry */





