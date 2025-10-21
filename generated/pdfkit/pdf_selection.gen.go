// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [PDFSelection] class.
var (
	PDFSelectionClass     _PDFSelectionClass
	PDFSelectionClassOnce sync.Once
)

func getPDFSelectionClass() _PDFSelectionClass {
	PDFSelectionClassOnce.Do(func() {
		PDFSelectionClass = _PDFSelectionClass{objc.GetClass("PDFSelection")}
	})
	return PDFSelectionClass
}

type _PDFSelectionClass struct {
	class objc.Class
}

// An interface definition for the [PDFSelection] class.
type IPDFSelection interface {
	objectivec.IObject
	AddSelections(selections unsafe.Pointer)
	AddSelection(selection unsafe.Pointer)
	BoundsForPage(page unsafe.Pointer) coregraphics.CGRect
	DrawForPageActive(page unsafe.Pointer, active bool)
	DrawForPageWithBoxActive(page unsafe.Pointer, box unsafe.Pointer, active bool)
	ExtendSelectionAtEnd(succeed int)
	ExtendSelectionAtStart(precede int)
	ExtendSelectionForLineBoundaries()
	NumberOfTextRangesOnPage(page unsafe.Pointer) uint
	RangeAtIndexOnPage(index uint, page unsafe.Pointer) Range
	SelectionsByLine() []PDFSelection
}

// A object identifies a contiguous or noncontiguous selection of text in a PDF document.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFSelection
type PDFSelection struct {
	objectivec.Object
}

// PDFSelectionFrom constructs a [PDFSelection] from an unsafe.Pointer.
//
// A object identifies a contiguous or noncontiguous selection of text in a PDF document.
func PDFSelectionFrom(ptr unsafe.Pointer) PDFSelection {
	return PDFSelection{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PDFSelectionClass) Alloc() PDFSelection {
	rv := objc.Send[PDFSelection](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PDFSelectionClass) New() PDFSelection {
	rv := objc.Send[PDFSelection](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFSelection) Init() PDFSelection {
	rv := objc.Send[PDFSelection](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFSelection) Autorelease() PDFSelection {
	rv := objc.Send[PDFSelection](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFSelection creates a new PDFSelection instance.
func NewPDFSelection() PDFSelection {
	return getPDFSelectionClass().New()
}


// Returns an empty object.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFSelection/init(document:)
func NewPDFSelectionWithDocument(document unsafe.Pointer) PDFSelection {
	instance := getPDFSelectionClass().Alloc()
	rv := objc.Send[PDFSelection](instance.ID, objc.Sel("initWithDocument:"), document)
	rv.Autorelease()
	return rv
}


// Adds the specified array of selections to the receiving selection.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFSelection/add(_:)-3fyld
func (p_ PDFSelection) AddSelections(selections unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addSelections:"), selections)
}

// Adds the specified selection to the receiving selection.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFSelection/add(_:)-8c2r
func (p_ PDFSelection) AddSelection(selection unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addSelection:"), selection)
}

// Returns the bounds of the selection on the specified page.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFSelection/bounds(for:)
func (p_ PDFSelection) BoundsForPage(page unsafe.Pointer) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](p_.ID, objc.Sel("boundsForPage:"), page)
	return rv
}

// Calls with a default value for box parameter.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFSelection/draw(for:active:)
func (p_ PDFSelection) DrawForPageActive(page unsafe.Pointer, active bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("drawForPage:active:"), page, active)
}

// Draws the selection relative to the origin of the specified box in page space.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFSelection/draw(for:with:active:)
func (p_ PDFSelection) DrawForPageWithBoxActive(page unsafe.Pointer, box unsafe.Pointer, active bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("drawForPage:withBox:active:"), page, box, active)
}

// Extends the selection from its end toward the end of the document.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFSelection/extend(atEnd:)
func (p_ PDFSelection) ExtendSelectionAtEnd(succeed int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("extendSelectionAtEnd:"), succeed)
}

// Extends the selection from its start toward the beginning of the document.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFSelection/extend(atStart:)
func (p_ PDFSelection) ExtendSelectionAtStart(precede int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("extendSelectionAtStart:"), precede)
}

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFSelection/extendForLineBoundaries()
func (p_ PDFSelection) ExtendSelectionForLineBoundaries() {
	objc.Send[objc.ID](p_.ID, objc.Sel("extendSelectionForLineBoundaries"))
}

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFSelection/numberOfTextRanges(on:)
func (p_ PDFSelection) NumberOfTextRangesOnPage(page unsafe.Pointer) uint {
	rv := objc.Send[uint](p_.ID, objc.Sel("numberOfTextRangesOnPage:"), page)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFSelection/range(at:on:)
func (p_ PDFSelection) RangeAtIndexOnPage(index uint, page unsafe.Pointer) Range {
	rv := objc.Send[Range](p_.ID, objc.Sel("rangeAtIndex:onPage:"), index, page)
	return rv
}

// Returns an array of selections, one for each line of text covered by the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFSelection/selectionsByLine()
func (p_ PDFSelection) SelectionsByLine() []PDFSelection {
	rv := objc.Send[[]PDFSelection](p_.ID, objc.Sel("selectionsByLine"))
	return rv
}

// Returns an object representing the text contained in the selection (may contain linefeed characters).
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFSelection/attributedString
func (p_ PDFSelection) AttributedString() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("attributedString"))
	return rv
}

// Sets the color used for the drawing of a selection in both active and inactive states.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFSelection/color
func (p_ PDFSelection) Color() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("color"))
	return rv
}


// SetColor sets the value of the color property.
// Sets the color used for the drawing of a selection in both active and inactive states.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFSelection/color
func (p_ PDFSelection) SetColor(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setColor:"), value)
}
// Returns the array of pages contained in the selection.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFSelection/pages
func (p_ PDFSelection) Pages() []PDFPage {
	rv := objc.Send[[]PDFPage](p_.ID, objc.Sel("pages"))
	return rv
}

// Returns an object representing the text contained in the selection (may contain linefeed characters).
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFSelection/string
func (p_ PDFSelection) String() string {
	rv := objc.Send[string](p_.ID, objc.Sel("string"))
	return rv
}


