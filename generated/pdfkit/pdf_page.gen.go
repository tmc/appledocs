// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/coregraphics"
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
	AddAnnotation(annotation unsafe.Pointer)
	AnnotationAtPoint(point Point) unsafe.Pointer
	BoundsForBox(box unsafe.Pointer) coregraphics.CGRect
	CharacterBoundsAtIndex(index int) coregraphics.CGRect
	CharacterIndexAtPoint(point coregraphics.CGPoint) int
	DrawWithBox(box unsafe.Pointer)
	DrawWithBoxToContext(box unsafe.Pointer, context CGContextRef)
	RemoveAnnotation(annotation unsafe.Pointer)
	SelectionForRange(range_ Range) unsafe.Pointer
	SelectionForRect(rect Rect) unsafe.Pointer
	SelectionFromPointToPoint(startPoint coregraphics.CGPoint, endPoint coregraphics.CGPoint) unsafe.Pointer
	SelectionForLineAtPoint(point Point) unsafe.Pointer
	SelectionForWordAtPoint(point coregraphics.CGPoint) unsafe.Pointer
	SetBoundsForBox(bounds coregraphics.CGRect, box unsafe.Pointer)
	ThumbnailOfSizeForBox(size coregraphics.CGSize, box unsafe.Pointer) unsafe.Pointer
	TransformContextForBox(context CGContextRef, box unsafe.Pointer)
	TransformForBox(box unsafe.Pointer) coregraphics.CGAffineTransform
}

// , a subclass of , defines methods used to render PDF pages and work with annotations, text, and selections.
//
// objects are flexible and powerful. With them you can render PDF content onscreen or to a printer, add annotations, count characters, define selections, and get the textual content of a page as an object. Your application instantiates a object by asking for one from a object. For simple display and navigation of PDF documents within your application, you don’t need to use . You need only use .
//
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
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/init(image:)
func NewPDFPageWithImage(image unsafe.Pointer) PDFPage {
	instance := getPDFPageClass().Alloc()
	rv := objc.Send[PDFPage](instance.ID, objc.Sel("initWithImage:"), image)
	rv.Autorelease()
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/init(image:options:)
func NewPDFPageWithImageOptions(image unsafe.Pointer, options unsafe.Pointer) PDFPage {
	instance := getPDFPageClass().Alloc()
	rv := objc.Send[PDFPage](instance.ID, objc.Sel("initWithImage:options:"), image, options)
	rv.Autorelease()
	return rv
}


// Adds the specified annotation object to the page.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/addAnnotation(_:)
func (p_ PDFPage) AddAnnotation(annotation unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addAnnotation:"), annotation)
}

// Returns the annotation, if there is one, at the specified point.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/annotation(at:)
func (p_ PDFPage) AnnotationAtPoint(point Point) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("annotationAtPoint:"), point)
	return rv
}

// Returns the bounds for the specified PDF display box.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/bounds(for:)
func (p_ PDFPage) BoundsForBox(box unsafe.Pointer) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](p_.ID, objc.Sel("boundsForBox:"), box)
	return rv
}

// Returns the bounds, in page space, of the character at the specified index.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/characterBounds(at:)
func (p_ PDFPage) CharacterBoundsAtIndex(index int) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](p_.ID, objc.Sel("characterBoundsAtIndex:"), index)
	return rv
}

// Returns the character index value for the specified point in page space.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/characterIndex(at:)
func (p_ PDFPage) CharacterIndexAtPoint(point coregraphics.CGPoint) int {
	rv := objc.Send[int](p_.ID, objc.Sel("characterIndexAtPoint:"), point)
	return rv
}

// Draws the page within the specified box.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/draw(with:)
func (p_ PDFPage) DrawWithBox(box unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("drawWithBox:"), box)
}

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/draw(with:to:)
func (p_ PDFPage) DrawWithBoxToContext(box unsafe.Pointer, context CGContextRef) {
	objc.Send[objc.ID](p_.ID, objc.Sel("drawWithBox:toContext:"), box, context)
}

// Removes the specified annotation from the page.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/removeAnnotation(_:)
func (p_ PDFPage) RemoveAnnotation(annotation unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("removeAnnotation:"), annotation)
}

// Returns the text contained within the specified range.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/selection(for:)-20y9d
func (p_ PDFPage) SelectionForRange(range_ Range) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("selectionForRange:"), range_)
	return rv
}

// Returns the text enclosed within the specified rectangle, expressed in page (user) coordinates.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/selection(for:)-2ckpi
func (p_ PDFPage) SelectionForRect(rect Rect) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("selectionForRect:"), rect)
	return rv
}

// Returns the text between the two specified points in page space.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/selection(from:to:)
func (p_ PDFPage) SelectionFromPointToPoint(startPoint coregraphics.CGPoint, endPoint coregraphics.CGPoint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("selectionFromPoint:toPoint:"), startPoint, endPoint)
	return rv
}

// Returns the whole line of text that includes the specified point.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/selectionForLine(at:)
func (p_ PDFPage) SelectionForLineAtPoint(point Point) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("selectionForLineAtPoint:"), point)
	return rv
}

// Returns the whole word that includes the specified point.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/selectionForWord(at:)
func (p_ PDFPage) SelectionForWordAtPoint(point coregraphics.CGPoint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("selectionForWordAtPoint:"), point)
	return rv
}

// Sets the bounds for the specified box.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/setBounds(_:for:)
func (p_ PDFPage) SetBoundsForBox(bounds coregraphics.CGRect, box unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBounds:forBox:"), bounds, box)
}

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/thumbnail(of:for:)
func (p_ PDFPage) ThumbnailOfSizeForBox(size coregraphics.CGSize, box unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("thumbnailOfSize:forBox:"), size, box)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/transform(_:for:)
func (p_ PDFPage) TransformContextForBox(context CGContextRef, box unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("transformContext:forBox:"), context, box)
}

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/transform(for:)
func (p_ PDFPage) TransformForBox(box unsafe.Pointer) coregraphics.CGAffineTransform {
	rv := objc.Send[coregraphics.CGAffineTransform](p_.ID, objc.Sel("transformForBox:"), box)
	return rv
}

// Returns an array containing the page’s annotations.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/annotations
func (p_ PDFPage) Annotations() []PDFAnnotation {
	rv := objc.Send[[]PDFAnnotation](p_.ID, objc.Sel("annotations"))
	return rv
}

// Returns an object representing the text on the page.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/attributedString
func (p_ PDFPage) AttributedString() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("attributedString"))
	return rv
}

// Returns the PDF data (that is, a PDF document) representing this page. This method does not preserve external page links.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/dataRepresentation
func (p_ PDFPage) DataRepresentation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("dataRepresentation"))
	return rv
}

// Returns a Boolean value indicating whether annotations are displayed for the page.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/displaysAnnotations
func (p_ PDFPage) DisplaysAnnotations() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("displaysAnnotations"))
	return rv
}


// SetDisplaysAnnotations sets the value of the displaysAnnotations property.
// Returns a Boolean value indicating whether annotations are displayed for the page.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/displaysAnnotations
func (p_ PDFPage) SetDisplaysAnnotations(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDisplaysAnnotations:"), value)
}
// Returns the object with which the page is associated.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/document
func (p_ PDFPage) Document() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("document"))
	return rv
}

// Returns the label for the page.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/label
func (p_ PDFPage) Label() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("label"))
	return rv
}

// Returns the number of characters on the page, including whitespace characters.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/numberOfCharacters
func (p_ PDFPage) NumberOfCharacters() uint {
	rv := objc.Send[uint](p_.ID, objc.Sel("numberOfCharacters"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/pageRef
func (p_ PDFPage) PageRef() CGPDFPageRef {
	rv := objc.Send[CGPDFPageRef](p_.ID, objc.Sel("pageRef"))
	return rv
}

// Sets the rotation angle for the page in degrees.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/rotation
func (p_ PDFPage) Rotation() int {
	rv := objc.Send[int](p_.ID, objc.Sel("rotation"))
	return rv
}


// SetRotation sets the value of the rotation property.
// Sets the rotation angle for the page in degrees.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/rotation
func (p_ PDFPage) SetRotation(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRotation:"), value)
}
// Returns an object representing the text on the page.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/string
func (p_ PDFPage) String() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("string"))
	return rv
}


