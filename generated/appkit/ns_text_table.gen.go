// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
)





// The class instance for the [TextTable] class.
var (
	TextTableClass     _TextTableClass
	TextTableClassOnce sync.Once
)

func getTextTableClass() _TextTableClass {
	TextTableClassOnce.Do(func() {
		TextTableClass = _TextTableClass{objc.GetClass("NSTextTable")}
	})
	return TextTableClass
}

type _TextTableClass struct {
	class objc.Class
}





// An interface definition for the [TextTable] class.
type ITextTable interface {
	ITextBlock
	

	// properties:
	CollapsesBorders() bool
	SetCollapsesBorders(value bool)
	HidesEmptyCells() bool
	SetHidesEmptyCells(value bool)
	LayoutAlgorithm() TextTableLayoutAlgorithm
	SetLayoutAlgorithm(value TextTableLayoutAlgorithm)
	NumberOfColumns() uint
	SetNumberOfColumns(value uint)


	

	// methods:
	BoundsRectForBlockContentRectInRectTextContainerCharacterRange(block ITextTableBlock, contentRect corefoundation.CGRect, rect corefoundation.CGRect, textContainer ITextContainer, charRange foundation.Range) corefoundation.CGRect
	DrawBackgroundForBlockWithFrameInViewCharacterRangeLayoutManager(block ITextTableBlock, frameRect corefoundation.CGRect, controlView IView, charRange foundation.Range, layoutManager ILayoutManager)
	RectForBlockLayoutAtPointInRectTextContainerCharacterRange(block ITextTableBlock, startingPoint corefoundation.CGPoint, rect corefoundation.CGRect, textContainer ITextContainer, charRange foundation.Range) corefoundation.CGRect


}





// Alloc allocates a new instance without initialization.
func (tc _TextTableClass) Alloc() TextTable {
	rv := objc.Send[TextTable](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TextTableClass) New() TextTable {
	rv := objc.Send[TextTable](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextTable) Init() TextTable {
	rv := objc.Send[TextTable](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextTable) Autorelease() TextTable {
	rv := objc.Send[TextTable](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextTable creates a new TextTable instance.
func NewTextTable() TextTable {
	return getTextTableClass().New()
}





// An object that represents a text table as a whole.
//
// A text table is responsible for laying out and drawing the text table blocks it contains, and it maintains the basic parameters of the table.


// An object that represents a text table as a whole.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextTable
type TextTable struct {
	TextBlock
}

// TextTableFrom constructs a [TextTable] from an unsafe.Pointer.
//
// An object that represents a text table as a whole.
func TextTableFrom(ptr unsafe.Pointer) TextTable {
	return TextTable{
		TextBlock: TextBlockFrom(ptr),
	}
}




















// Returns the rectangle the text table block actually occupies, including padding, borders, and margins.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextTable/boundsRect(for:contentRect:in:textContainer:characterRange:)
func (t_ TextTable) BoundsRectForBlockContentRectInRectTextContainerCharacterRange(block ITextTableBlock, contentRect corefoundation.CGRect, rect corefoundation.CGRect, textContainer ITextContainer, charRange foundation.Range) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t_.ID, objc.Sel("boundsRectForBlock:contentRect:inRect:textContainer:characterRange:"), block, contentRect, rect, textContainer, charRange)
	return rv
}


// Draws any colors and other decorations for a text table block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextTable/drawBackground(for:withFrame:in:characterRange:layoutManager:)
func (t_ TextTable) DrawBackgroundForBlockWithFrameInViewCharacterRangeLayoutManager(block ITextTableBlock, frameRect corefoundation.CGRect, controlView IView, charRange foundation.Range, layoutManager ILayoutManager) {
	objc.Send[objc.ID](t_.ID, objc.Sel("drawBackgroundForBlock:withFrame:inView:characterRange:layoutManager:"), block, frameRect, controlView, charRange, layoutManager)
}


// Returns the rectangle within which glyphs should be laid out for a text table block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextTable/rect(for:layoutAt:in:textContainer:characterRange:)
func (t_ TextTable) RectForBlockLayoutAtPointInRectTextContainerCharacterRange(block ITextTableBlock, startingPoint corefoundation.CGPoint, rect corefoundation.CGRect, textContainer ITextContainer, charRange foundation.Range) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t_.ID, objc.Sel("rectForBlock:layoutAtPoint:inRect:textContainer:characterRange:"), block, startingPoint, rect, textContainer, charRange)
	return rv
}







// A Boolean value indicating whether the text table borders are collapsible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextTable/collapsesBorders
func (t_ TextTable) CollapsesBorders() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("collapsesBorders"))
	return rv
}


// A Boolean value indicating whether the text table borders are collapsible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextTable/collapsesBorders
func (t_ TextTable) SetCollapsesBorders(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCollapsesBorders:"), value)
}


// A Boolean value indicating whether the text table hides empty cells.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextTable/hidesEmptyCells
func (t_ TextTable) HidesEmptyCells() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("hidesEmptyCells"))
	return rv
}


// A Boolean value indicating whether the text table hides empty cells.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextTable/hidesEmptyCells
func (t_ TextTable) SetHidesEmptyCells(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setHidesEmptyCells:"), value)
}


// The text table layout algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextTable/layoutAlgorithm-swift.property
func (t_ TextTable) LayoutAlgorithm() TextTableLayoutAlgorithm {
	rv := objc.Send[TextTableLayoutAlgorithm](t_.ID, objc.Sel("layoutAlgorithm"))
	return rv
}


// The text table layout algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextTable/layoutAlgorithm-swift.property
func (t_ TextTable) SetLayoutAlgorithm(value TextTableLayoutAlgorithm) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLayoutAlgorithm:"), value)
}


// The number of columns in the text table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextTable/numberOfColumns
func (t_ TextTable) NumberOfColumns() uint {
	rv := objc.Send[uint](t_.ID, objc.Sel("numberOfColumns"))
	return rv
}


// The number of columns in the text table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextTable/numberOfColumns
func (t_ TextTable) SetNumberOfColumns(value uint) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setNumberOfColumns:"), value)
}








