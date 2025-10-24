// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class PDFSelection */


/* debug [class_header]: Header for PDFSelection */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PDFSelection */
// An interface definition for the [PDFSelection] class.
type IPDFSelection interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PDFSelection */
	// properties:
	AttributedString() foundation.AttributedString
	Color() appkit.Color
	SetColor(value appkit.Color)
	Pages() []PDFPage
	String() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PDFSelection */
	// methods:
	AddSelections(selections []PDFSelection)
	AddSelection(selection IPDFSelection)
	BoundsForPage(page IPDFPage) Rect /* not a class type */
	DrawForPageActive(page IPDFPage, active bool)
	DrawForPageWithBoxActive(page IPDFPage, box PDFDisplayBox, active bool)
	ExtendSelectionAtEnd(succeed int)
	ExtendSelectionAtStart(precede int)
	ExtendSelectionForLineBoundaries()
	NumberOfTextRangesOnPage(page IPDFPage) uint
	RangeAtIndexOnPage(index uint, page IPDFPage) corefoundation.Range
	SelectionsByLine() []PDFSelection
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PDFSelection */
// Alloc allocates a new instance without initialization.
func (pc _PDFSelectionClass) Alloc() PDFSelection {
	rv := objc.Send[PDFSelection](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PDFSelection */
// A object identifies a contiguous or noncontiguous selection of text in a PDF document.


// A object identifies a contiguous or noncontiguous selection of text in a PDF document.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PDFSelection */

// Returns an empty object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFSelection/init(document:)
func NewPDFSelectionWithDocument(document IPDFDocument) PDFSelection {
	instance := getPDFSelectionClass().Alloc()
	rv := objc.Send[PDFSelection](instance.ID, objc.Sel("initWithDocument:"), document)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPDFSelectionWithDocument */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PDFSelection */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PDFSelection */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PDFSelection */

// Adds the specified array of selections to the receiving selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFSelection/add(_:)-3fyld
func (p_ PDFSelection) AddSelections(selections []PDFSelection) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addSelections:"), selections)
}/* debug [instance_methods/method]: AddSelections */


// Adds the specified selection to the receiving selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFSelection/add(_:)-8c2r
func (p_ PDFSelection) AddSelection(selection IPDFSelection) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addSelection:"), selection)
}/* debug [instance_methods/method]: AddSelection */


// Returns the bounds of the selection on the specified page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFSelection/bounds(for:)
func (p_ PDFSelection) BoundsForPage(page IPDFPage) Rect /* not a class type */ {
	rv := objc.Send[Rect](p_.ID, objc.Sel("boundsForPage:"), page)
	return rv
}/* debug [instance_methods/method]: BoundsForPage */


// Calls with a default value for box parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFSelection/draw(for:active:)
func (p_ PDFSelection) DrawForPageActive(page IPDFPage, active bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("drawForPage:active:"), page, active)
}/* debug [instance_methods/method]: DrawForPageActive */


// Draws the selection relative to the origin of the specified box in page space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFSelection/draw(for:with:active:)
func (p_ PDFSelection) DrawForPageWithBoxActive(page IPDFPage, box PDFDisplayBox, active bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("drawForPage:withBox:active:"), page, box, active)
}/* debug [instance_methods/method]: DrawForPageWithBoxActive */


// Extends the selection from its end toward the end of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFSelection/extend(atEnd:)
func (p_ PDFSelection) ExtendSelectionAtEnd(succeed int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("extendSelectionAtEnd:"), succeed)
}/* debug [instance_methods/method]: ExtendSelectionAtEnd */


// Extends the selection from its start toward the beginning of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFSelection/extend(atStart:)
func (p_ PDFSelection) ExtendSelectionAtStart(precede int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("extendSelectionAtStart:"), precede)
}/* debug [instance_methods/method]: ExtendSelectionAtStart */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFSelection/extendForLineBoundaries()
func (p_ PDFSelection) ExtendSelectionForLineBoundaries() {
	objc.Send[objc.ID](p_.ID, objc.Sel("extendSelectionForLineBoundaries"))
}/* debug [instance_methods/method]: ExtendSelectionForLineBoundaries */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFSelection/numberOfTextRanges(on:)
func (p_ PDFSelection) NumberOfTextRangesOnPage(page IPDFPage) uint {
	rv := objc.Send[uint](p_.ID, objc.Sel("numberOfTextRangesOnPage:"), page)
	return rv
}/* debug [instance_methods/method]: NumberOfTextRangesOnPage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFSelection/range(at:on:)
func (p_ PDFSelection) RangeAtIndexOnPage(index uint, page IPDFPage) corefoundation.Range {
	rv := objc.Send[corefoundation.Range](p_.ID, objc.Sel("rangeAtIndex:onPage:"), index, page)
	return rv
}/* debug [instance_methods/method]: RangeAtIndexOnPage */


// Returns an array of selections, one for each line of text covered by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFSelection/selectionsByLine()
func (p_ PDFSelection) SelectionsByLine() []PDFSelection {
	rv := objc.Send[[]PDFSelection](p_.ID, objc.Sel("selectionsByLine"))
	return rv
}/* debug [instance_methods/method]: SelectionsByLine */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PDFSelection */

// Returns an object representing the text contained in the selection (may contain linefeed characters).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFSelection/attributedString
func (p_ PDFSelection) AttributedString() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](p_.ID, objc.Sel("attributedString"))
	return rv
}/* debug [instance_properties/getter]: attributedString */


// Sets the color used for the drawing of a selection in both active and inactive states.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFSelection/color
func (p_ PDFSelection) Color() appkit.Color {
	rv := objc.Send[appkit.Color](p_.ID, objc.Sel("color"))
	return rv
}/* debug [instance_properties/getter]: color */


// Sets the color used for the drawing of a selection in both active and inactive states.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFSelection/color
func (p_ PDFSelection) SetColor(value appkit.Color) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setColor:"), value)
}/* debug [instance_properties/setter]: color */


// Returns the array of pages contained in the selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFSelection/pages
func (p_ PDFSelection) Pages() []PDFPage {
	rv := objc.Send[[]PDFPage](p_.ID, objc.Sel("pages"))
	return rv
}/* debug [instance_properties/getter]: pages */


// Returns an object representing the text contained in the selection (may contain linefeed characters).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFSelection/string
func (p_ PDFSelection) String() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("string"))
	return rv
}/* debug [instance_properties/getter]: string */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PDFSelection */


