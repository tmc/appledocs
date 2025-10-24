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

// The class instance for the [PDFPage] class.
var (
	PDFPageClass     _PDFPageClass
	PDFPageClassOnce sync.Once
)

func getPDFPageClass() _PDFPageClass {
	PDFPageClassOnce.Do(func() {
		PDFPageClass = _PDFPageClass{objc.GetClass("PDFPage")}
	})
	return PDFPageClass
}

type _PDFPageClass struct {
	class objc.Class
}

// An interface definition for the [PDFPage] class.
type IPDFPage interface {
	objectivec.IObject
	// properties:
	Annotations() []IPDFAnnotation
	AttributedString() objc.IObject /* cross-framework: AttributedString */
	DataRepresentation() objc.IObject /* cross-framework: NSData */
	DisplaysAnnotations() bool
	SetDisplaysAnnotations(value bool)
	Document() IPDFDocument
	Label() objc.IObject /* cross-framework: NSString */
	NumberOfCharacters() uint
	PageRef() PDFPageRef /* not a class type */
	Rotation() int
	SetRotation(value int)
	String() objc.IObject /* cross-framework: NSString */
	// methods:
	AddAnnotation(annotation IPDFAnnotation)
	AnnotationAtPoint(point objc.IObject /* cross-framework: Point */) IPDFAnnotation
	BoundsForBox(box PDFDisplayBox) objc.IObject /* cross-framework: Rect */
	CharacterBoundsAtIndex(index int) objc.IObject /* cross-framework: Rect */
	CharacterIndexAtPoint(point objc.IObject /* cross-framework: Point */) int
	DrawWithBoxToContext(box PDFDisplayBox, context ContextRef /* not a class type */)
	RemoveAnnotation(annotation IPDFAnnotation)
	SelectionForRange(range_ objc.IObject /* cross-framework: Range */) IPDFSelection
	SelectionForRect(rect objc.IObject /* cross-framework: Rect */) IPDFSelection
	SelectionFromPointToPoint(startPoint objc.IObject /* cross-framework: Point */, endPoint objc.IObject /* cross-framework: Point */) IPDFSelection
	SelectionForLineAtPoint(point objc.IObject /* cross-framework: Point */) IPDFSelection
	SelectionForWordAtPoint(point objc.IObject /* cross-framework: Point */) IPDFSelection
	SetBoundsForBox(bounds objc.IObject /* cross-framework: Rect */, box PDFDisplayBox)
	ThumbnailOfSizeForBox(size objc.IObject /* cross-framework: Size */, box PDFDisplayBox) objc.IObject /* cross-framework: Image */
	TransformContextForBox(context ContextRef /* not a class type */, box PDFDisplayBox)
	TransformForBox(box PDFDisplayBox) objc.IObject /* cross-framework: AffineTransform */
}

// , a subclass of , defines methods used to render PDF pages and work with annotations, text, and selections.
//
// objects are flexible and powerful. With them you can render PDF content onscreen or to a printer, add annotations, count characters, define selections, and get the textual content of a page as an object. Your application instantiates a object by asking for one from a object. For simple display and navigation of PDF documents within your application, you don’t need to use . You need only use .


// , a subclass of , defines methods used to render PDF pages and work with annotations, text, and selections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage
type PDFPage struct {
	objectivec.Object
}

// PDFPageFrom constructs a [PDFPage] from an unsafe.Pointer.
//
// , a subclass of , defines methods used to render PDF pages and work with annotations, text, and selections.
func PDFPageFrom(ptr unsafe.Pointer) PDFPage {
	return PDFPage{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PDFPageClass) Alloc() PDFPage {
	rv := objc.Send[PDFPage](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PDFPageClass) New() PDFPage {
	rv := objc.Send[PDFPage](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFPage) Init() PDFPage {
	rv := objc.Send[PDFPage](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFPage) Autorelease() PDFPage {
	rv := objc.Send[PDFPage](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFPage creates a new PDFPage instance.
func NewPDFPage() PDFPage {
	return getPDFPageClass().New()
}



// Creates a new object and initializes it with the specified object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/init(image:)
func NewPDFPageWithImage(image objc.IObject /* cross-framework: Image */) PDFPage {
	instance := getPDFPageClass().Alloc()
	rv := objc.Send[PDFPage](instance.ID, objc.Sel("initWithImage:"), image)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/init(image:options:)
func NewPDFPageWithImageOptions(image objc.IObject /* cross-framework: Image */, options foundation.IDictionary) PDFPage {
	instance := getPDFPageClass().Alloc()
	rv := objc.Send[PDFPage](instance.ID, objc.Sel("initWithImage:options:"), image, options)
	rv.Autorelease()
	return rv
}



// Adds the specified annotation object to the page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/addAnnotation(_:)
func (p_ PDFPage) AddAnnotation(annotation IPDFAnnotation) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addAnnotation:"), annotation)
}


// Returns the annotation, if there is one, at the specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/annotation(at:)
func (p_ PDFPage) AnnotationAtPoint(point objc.IObject /* cross-framework: Point */) IPDFAnnotation {
	rv := objc.Send[PDFAnnotation](p_.ID, objc.Sel("annotationAtPoint:"), point)
	return rv
}


// Returns the bounds for the specified PDF display box.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/bounds(for:)
func (p_ PDFPage) BoundsForBox(box PDFDisplayBox) objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](p_.ID, objc.Sel("boundsForBox:"), box)
	return rv
}


// Returns the bounds, in page space, of the character at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/characterBounds(at:)
func (p_ PDFPage) CharacterBoundsAtIndex(index int) objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](p_.ID, objc.Sel("characterBoundsAtIndex:"), index)
	return rv
}


// Returns the character index value for the specified point in page space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/characterIndex(at:)
func (p_ PDFPage) CharacterIndexAtPoint(point objc.IObject /* cross-framework: Point */) int {
	rv := objc.Send[int](p_.ID, objc.Sel("characterIndexAtPoint:"), point)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/draw(with:to:)
func (p_ PDFPage) DrawWithBoxToContext(box PDFDisplayBox, context ContextRef /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("drawWithBox:toContext:"), box, context)
}


// Removes the specified annotation from the page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/removeAnnotation(_:)
func (p_ PDFPage) RemoveAnnotation(annotation IPDFAnnotation) {
	objc.Send[objc.ID](p_.ID, objc.Sel("removeAnnotation:"), annotation)
}


// Returns the text contained within the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/selection(for:)-20y9d
func (p_ PDFPage) SelectionForRange(range_ objc.IObject /* cross-framework: Range */) IPDFSelection {
	rv := objc.Send[PDFSelection](p_.ID, objc.Sel("selectionForRange:"), range_)
	return rv
}


// Returns the text enclosed within the specified rectangle, expressed in page (user) coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/selection(for:)-2ckpi
func (p_ PDFPage) SelectionForRect(rect objc.IObject /* cross-framework: Rect */) IPDFSelection {
	rv := objc.Send[PDFSelection](p_.ID, objc.Sel("selectionForRect:"), rect)
	return rv
}


// Returns the text between the two specified points in page space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/selection(from:to:)
func (p_ PDFPage) SelectionFromPointToPoint(startPoint objc.IObject /* cross-framework: Point */, endPoint objc.IObject /* cross-framework: Point */) IPDFSelection {
	rv := objc.Send[PDFSelection](p_.ID, objc.Sel("selectionFromPoint:toPoint:"), startPoint, endPoint)
	return rv
}


// Returns the whole line of text that includes the specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/selectionForLine(at:)
func (p_ PDFPage) SelectionForLineAtPoint(point objc.IObject /* cross-framework: Point */) IPDFSelection {
	rv := objc.Send[PDFSelection](p_.ID, objc.Sel("selectionForLineAtPoint:"), point)
	return rv
}


// Returns the whole word that includes the specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/selectionForWord(at:)
func (p_ PDFPage) SelectionForWordAtPoint(point objc.IObject /* cross-framework: Point */) IPDFSelection {
	rv := objc.Send[PDFSelection](p_.ID, objc.Sel("selectionForWordAtPoint:"), point)
	return rv
}


// Sets the bounds for the specified box.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/setBounds(_:for:)
func (p_ PDFPage) SetBoundsForBox(bounds objc.IObject /* cross-framework: Rect */, box PDFDisplayBox) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBounds:forBox:"), bounds, box)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/thumbnail(of:for:)
func (p_ PDFPage) ThumbnailOfSizeForBox(size objc.IObject /* cross-framework: Size */, box PDFDisplayBox) objc.IObject /* cross-framework: Image */ {
	rv := objc.Send[appkit.Image](p_.ID, objc.Sel("thumbnailOfSize:forBox:"), size, box)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/transform(_:for:)
func (p_ PDFPage) TransformContextForBox(context ContextRef /* not a class type */, box PDFDisplayBox) {
	objc.Send[objc.ID](p_.ID, objc.Sel("transformContext:forBox:"), context, box)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/transform(for:)
func (p_ PDFPage) TransformForBox(box PDFDisplayBox) objc.IObject /* cross-framework: AffineTransform */ {
	rv := objc.Send[corefoundation.AffineTransform](p_.ID, objc.Sel("transformForBox:"), box)
	return rv
}


// Returns an array containing the page’s annotations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/annotations
func (p_ PDFPage) Annotations() []IPDFAnnotation {
	rv := objc.Send[[]PDFAnnotation](p_.ID, objc.Sel("annotations"))
	return rv
}


// Returns an object representing the text on the page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/attributedString
func (p_ PDFPage) AttributedString() objc.IObject /* cross-framework: AttributedString */ {
	rv := objc.Send[foundation.AttributedString](p_.ID, objc.Sel("attributedString"))
	return rv
}


// Returns the PDF data (that is, a PDF document) representing this page. This method does not preserve external page links.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/dataRepresentation
func (p_ PDFPage) DataRepresentation() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](p_.ID, objc.Sel("dataRepresentation"))
	return rv
}


// Returns a Boolean value indicating whether annotations are displayed for the page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/displaysAnnotations
func (p_ PDFPage) DisplaysAnnotations() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("displaysAnnotations"))
	return rv
}


// Returns a Boolean value indicating whether annotations are displayed for the page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/displaysAnnotations
func (p_ PDFPage) SetDisplaysAnnotations(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDisplaysAnnotations:"), value)
}


// Returns the object with which the page is associated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/document
func (p_ PDFPage) Document() IPDFDocument {
	rv := objc.Send[PDFDocument](p_.ID, objc.Sel("document"))
	return rv
}


// Returns the label for the page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/label
func (p_ PDFPage) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("label"))
	return rv
}


// Returns the number of characters on the page, including whitespace characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/numberOfCharacters
func (p_ PDFPage) NumberOfCharacters() uint {
	rv := objc.Send[uint](p_.ID, objc.Sel("numberOfCharacters"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/pageRef
func (p_ PDFPage) PageRef() PDFPageRef /* not a class type */ {
	rv := objc.Send[PDFPageRef](p_.ID, objc.Sel("pageRef"))
	return rv
}


// Sets the rotation angle for the page in degrees.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/rotation
func (p_ PDFPage) Rotation() int {
	rv := objc.Send[int](p_.ID, objc.Sel("rotation"))
	return rv
}


// Sets the rotation angle for the page in degrees.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/rotation
func (p_ PDFPage) SetRotation(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRotation:"), value)
}


// Returns an object representing the text on the page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/string
func (p_ PDFPage) String() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("string"))
	return rv
}


