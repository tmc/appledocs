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
	"github.com/tmc/appledocs/generated/vision"
)

/* debug [class.gen.go]: Generating class PDFPage */


/* debug [class_header]: Header for PDFPage */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PDFPage */
// An interface definition for the [PDFPage] class.
type IPDFPage interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PDFPage */
	// properties:
	Annotations() []PDFAnnotation
	AttributedString() foundation.AttributedString
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PDFPage */
	// methods:
	AddAnnotation(annotation IPDFAnnotation)
	AnnotationAtPoint(point vision.Point) IPDFAnnotation
	BoundsForBox(box PDFDisplayBox) Rect /* not a class type */
	CharacterBoundsAtIndex(index int) corefoundation.CGRect
	CharacterIndexAtPoint(point vision.Point) int
	DrawWithBoxToContext(box PDFDisplayBox, context ContextRef /* not a class type */)
	RemoveAnnotation(annotation IPDFAnnotation)
	SelectionForRange(range_ corefoundation.Range) IPDFSelection
	SelectionForRect(rect Rect /* not a class type */) IPDFSelection
	SelectionFromPointToPoint(startPoint vision.Point, endPoint vision.Point) IPDFSelection
	SelectionForLineAtPoint(point vision.Point) IPDFSelection
	SelectionForWordAtPoint(point vision.Point) IPDFSelection
	SetBoundsForBox(bounds Rect /* not a class type */, box PDFDisplayBox)
	ThumbnailOfSizeForBox(size corefoundation.CGSize, box PDFDisplayBox) appkit.Image
	TransformContextForBox(context ContextRef /* not a class type */, box PDFDisplayBox)
	TransformForBox(box PDFDisplayBox) corefoundation.CGAffineTransform
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PDFPage */
// Alloc allocates a new instance without initialization.
func (pc _PDFPageClass) Alloc() PDFPage {
	rv := objc.Send[PDFPage](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PDFPage */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PDFPage */

// Creates a new object and initializes it with the specified object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/init(image:)
func NewPDFPageWithImage(image appkit.Image) PDFPage {
	instance := getPDFPageClass().Alloc()
	rv := objc.Send[PDFPage](instance.ID, objc.Sel("initWithImage:"), image)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPDFPageWithImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/init(image:options:)
func NewPDFPageWithImageOptions(image appkit.Image, options foundation.IDictionary) PDFPage {
	instance := getPDFPageClass().Alloc()
	rv := objc.Send[PDFPage](instance.ID, objc.Sel("initWithImage:options:"), image, options)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPDFPageWithImageOptions */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PDFPage */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PDFPage */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PDFPage */

// Adds the specified annotation object to the page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/addAnnotation(_:)
func (p_ PDFPage) AddAnnotation(annotation IPDFAnnotation) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addAnnotation:"), annotation)
}/* debug [instance_methods/method]: AddAnnotation */


// Returns the annotation, if there is one, at the specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/annotation(at:)
func (p_ PDFPage) AnnotationAtPoint(point vision.Point) IPDFAnnotation {
	rv := objc.Send[PDFAnnotation](p_.ID, objc.Sel("annotationAtPoint:"), point)
	return rv
}/* debug [instance_methods/method]: AnnotationAtPoint */


// Returns the bounds for the specified PDF display box.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/bounds(for:)
func (p_ PDFPage) BoundsForBox(box PDFDisplayBox) Rect /* not a class type */ {
	rv := objc.Send[Rect](p_.ID, objc.Sel("boundsForBox:"), box)
	return rv
}/* debug [instance_methods/method]: BoundsForBox */


// Returns the bounds, in page space, of the character at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/characterBounds(at:)
func (p_ PDFPage) CharacterBoundsAtIndex(index int) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](p_.ID, objc.Sel("characterBoundsAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: CharacterBoundsAtIndex */


// Returns the character index value for the specified point in page space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/characterIndex(at:)
func (p_ PDFPage) CharacterIndexAtPoint(point vision.Point) int {
	rv := objc.Send[int](p_.ID, objc.Sel("characterIndexAtPoint:"), point)
	return rv
}/* debug [instance_methods/method]: CharacterIndexAtPoint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/draw(with:to:)
func (p_ PDFPage) DrawWithBoxToContext(box PDFDisplayBox, context ContextRef /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("drawWithBox:toContext:"), box, context)
}/* debug [instance_methods/method]: DrawWithBoxToContext */


// Removes the specified annotation from the page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/removeAnnotation(_:)
func (p_ PDFPage) RemoveAnnotation(annotation IPDFAnnotation) {
	objc.Send[objc.ID](p_.ID, objc.Sel("removeAnnotation:"), annotation)
}/* debug [instance_methods/method]: RemoveAnnotation */


// Returns the text contained within the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/selection(for:)-20y9d
func (p_ PDFPage) SelectionForRange(range_ corefoundation.Range) IPDFSelection {
	rv := objc.Send[PDFSelection](p_.ID, objc.Sel("selectionForRange:"), range_)
	return rv
}/* debug [instance_methods/method]: SelectionForRange */


// Returns the text enclosed within the specified rectangle, expressed in page (user) coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/selection(for:)-2ckpi
func (p_ PDFPage) SelectionForRect(rect Rect /* not a class type */) IPDFSelection {
	rv := objc.Send[PDFSelection](p_.ID, objc.Sel("selectionForRect:"), rect)
	return rv
}/* debug [instance_methods/method]: SelectionForRect */


// Returns the text between the two specified points in page space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/selection(from:to:)
func (p_ PDFPage) SelectionFromPointToPoint(startPoint vision.Point, endPoint vision.Point) IPDFSelection {
	rv := objc.Send[PDFSelection](p_.ID, objc.Sel("selectionFromPoint:toPoint:"), startPoint, endPoint)
	return rv
}/* debug [instance_methods/method]: SelectionFromPointToPoint */


// Returns the whole line of text that includes the specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/selectionForLine(at:)
func (p_ PDFPage) SelectionForLineAtPoint(point vision.Point) IPDFSelection {
	rv := objc.Send[PDFSelection](p_.ID, objc.Sel("selectionForLineAtPoint:"), point)
	return rv
}/* debug [instance_methods/method]: SelectionForLineAtPoint */


// Returns the whole word that includes the specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/selectionForWord(at:)
func (p_ PDFPage) SelectionForWordAtPoint(point vision.Point) IPDFSelection {
	rv := objc.Send[PDFSelection](p_.ID, objc.Sel("selectionForWordAtPoint:"), point)
	return rv
}/* debug [instance_methods/method]: SelectionForWordAtPoint */


// Sets the bounds for the specified box.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/setBounds(_:for:)
func (p_ PDFPage) SetBoundsForBox(bounds Rect /* not a class type */, box PDFDisplayBox) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBounds:forBox:"), bounds, box)
}/* debug [instance_methods/method]: SetBoundsForBox */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/thumbnail(of:for:)
func (p_ PDFPage) ThumbnailOfSizeForBox(size corefoundation.CGSize, box PDFDisplayBox) appkit.Image {
	rv := objc.Send[appkit.Image](p_.ID, objc.Sel("thumbnailOfSize:forBox:"), size, box)
	return rv
}/* debug [instance_methods/method]: ThumbnailOfSizeForBox */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/transform(_:for:)
func (p_ PDFPage) TransformContextForBox(context ContextRef /* not a class type */, box PDFDisplayBox) {
	objc.Send[objc.ID](p_.ID, objc.Sel("transformContext:forBox:"), context, box)
}/* debug [instance_methods/method]: TransformContextForBox */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/transform(for:)
func (p_ PDFPage) TransformForBox(box PDFDisplayBox) corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](p_.ID, objc.Sel("transformForBox:"), box)
	return rv
}/* debug [instance_methods/method]: TransformForBox */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PDFPage */

// Returns an array containing the page’s annotations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/annotations
func (p_ PDFPage) Annotations() []PDFAnnotation {
	rv := objc.Send[[]PDFAnnotation](p_.ID, objc.Sel("annotations"))
	return rv
}/* debug [instance_properties/getter]: annotations */


// Returns an object representing the text on the page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/attributedString
func (p_ PDFPage) AttributedString() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](p_.ID, objc.Sel("attributedString"))
	return rv
}/* debug [instance_properties/getter]: attributedString */


// Returns the PDF data (that is, a PDF document) representing this page. This method does not preserve external page links.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/dataRepresentation
func (p_ PDFPage) DataRepresentation() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](p_.ID, objc.Sel("dataRepresentation"))
	return rv
}/* debug [instance_properties/getter]: dataRepresentation */


// Returns a Boolean value indicating whether annotations are displayed for the page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/displaysAnnotations
func (p_ PDFPage) DisplaysAnnotations() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("displaysAnnotations"))
	return rv
}/* debug [instance_properties/getter]: displaysAnnotations */


// Returns a Boolean value indicating whether annotations are displayed for the page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/displaysAnnotations
func (p_ PDFPage) SetDisplaysAnnotations(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDisplaysAnnotations:"), value)
}/* debug [instance_properties/setter]: displaysAnnotations */


// Returns the object with which the page is associated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/document
func (p_ PDFPage) Document() IPDFDocument {
	rv := objc.Send[PDFDocument](p_.ID, objc.Sel("document"))
	return rv
}/* debug [instance_properties/getter]: document */


// Returns the label for the page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/label
func (p_ PDFPage) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// Returns the number of characters on the page, including whitespace characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/numberOfCharacters
func (p_ PDFPage) NumberOfCharacters() uint {
	rv := objc.Send[uint](p_.ID, objc.Sel("numberOfCharacters"))
	return rv
}/* debug [instance_properties/getter]: numberOfCharacters */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/pageRef
func (p_ PDFPage) PageRef() PDFPageRef /* not a class type */ {
	rv := objc.Send[PDFPageRef](p_.ID, objc.Sel("pageRef"))
	return rv
}/* debug [instance_properties/getter]: pageRef */


// Sets the rotation angle for the page in degrees.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/rotation
func (p_ PDFPage) Rotation() int {
	rv := objc.Send[int](p_.ID, objc.Sel("rotation"))
	return rv
}/* debug [instance_properties/getter]: rotation */


// Sets the rotation angle for the page in degrees.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/rotation
func (p_ PDFPage) SetRotation(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRotation:"), value)
}/* debug [instance_properties/setter]: rotation */


// Returns an object representing the text on the page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFPage/string
func (p_ PDFPage) String() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("string"))
	return rv
}/* debug [instance_properties/getter]: string */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PDFPage */


