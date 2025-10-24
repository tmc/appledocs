// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class PDFOutline */


/* debug [class_header]: Header for PDFOutline */
// The class instance for the [PDFOutline] class.
var (
	PDFOutlineClass     _PDFOutlineClass
	PDFOutlineClassOnce sync.Once
)

func getPDFOutlineClass() _PDFOutlineClass {
	PDFOutlineClassOnce.Do(func() {
		PDFOutlineClass = _PDFOutlineClass{objc.GetClass("PDFOutline")}
	})
	return PDFOutlineClass
}

type _PDFOutlineClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PDFOutline */
// An interface definition for the [PDFOutline] class.
type IPDFOutline interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PDFOutline */
	// properties:
	Action() IPDFAction
	SetAction(value IPDFAction)
	Destination() IPDFDestination
	SetDestination(value IPDFDestination)
	Document() IPDFDocument
	Index() uint
	IsOpen() bool
	SetIsOpen(value bool)
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	NumberOfChildren() uint
	Parent() IPDFOutline
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PDFOutline */
	// methods:
	ChildAtIndex(index uint) IPDFOutline
	InsertChildAtIndex(child IPDFOutline, index uint)
	RemoveFromParent()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PDFOutline */
// Alloc allocates a new instance without initialization.
func (pc _PDFOutlineClass) Alloc() PDFOutline {
	rv := objc.Send[PDFOutline](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PDFOutlineClass) New() PDFOutline {
	rv := objc.Send[PDFOutline](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFOutline) Init() PDFOutline {
	rv := objc.Send[PDFOutline](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFOutline) Autorelease() PDFOutline {
	rv := objc.Send[PDFOutline](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFOutline creates a new PDFOutline instance.
func NewPDFOutline() PDFOutline {
	return getPDFOutlineClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PDFOutline */
// A object is an element in a tree-structured hierarchy that can represent the structure of a PDF document.
//
// An outline is an optional component of a PDF document, useful for viewing the structure of the document and for navigating within it. Outlines are created by the document’s author. If you represent a PDF document outline using outline objects, the root of the hierarchy is obtained from the PDF document itself. This root outline is not visible and serves merely as a container for the visible outlines.


// A object is an element in a tree-structured hierarchy that can represent the structure of a PDF document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFOutline
type PDFOutline struct {
	objectivec.Object
}

// PDFOutlineFrom constructs a [PDFOutline] from an unsafe.Pointer.
//
// A object is an element in a tree-structured hierarchy that can represent the structure of a PDF document.
func PDFOutlineFrom(ptr unsafe.Pointer) PDFOutline {
	return PDFOutline{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PDFOutline */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PDFOutline */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PDFOutline */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PDFOutline */

// Returns the child outline object at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFOutline/child(at:)
func (p_ PDFOutline) ChildAtIndex(index uint) PDFOutline {
	rv := objc.Send[PDFOutline](p_.ID, objc.Sel("childAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: ChildAtIndex */


// Inserts the specified outline object at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFOutline/insertChild(_:at:)
func (p_ PDFOutline) InsertChildAtIndex(child IPDFOutline, index uint) {
	objc.Send[objc.ID](p_.ID, objc.Sel("insertChild:atIndex:"), child, index)
}/* debug [instance_methods/method]: InsertChildAtIndex */


// Removes the outline object from its parent (does nothing if outline object is the root outline object).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFOutline/removeFromParent()
func (p_ PDFOutline) RemoveFromParent() {
	objc.Send[objc.ID](p_.ID, objc.Sel("removeFromParent"))
}/* debug [instance_methods/method]: RemoveFromParent */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PDFOutline */

// Returns the action performed when users click the outline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFOutline/action
func (p_ PDFOutline) Action() IPDFAction {
	rv := objc.Send[PDFAction](p_.ID, objc.Sel("action"))
	return rv
}/* debug [instance_properties/getter]: action */


// Returns the action performed when users click the outline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFOutline/action
func (p_ PDFOutline) SetAction(value IPDFAction) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAction:"), value)
}/* debug [instance_properties/setter]: action */


// Returns the destination associated with the outline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFOutline/destination
func (p_ PDFOutline) Destination() IPDFDestination {
	rv := objc.Send[PDFDestination](p_.ID, objc.Sel("destination"))
	return rv
}/* debug [instance_properties/getter]: destination */


// Returns the destination associated with the outline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFOutline/destination
func (p_ PDFOutline) SetDestination(value IPDFDestination) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDestination:"), value)
}/* debug [instance_properties/setter]: destination */


// Returns the document with which the outline is associated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFOutline/document
func (p_ PDFOutline) Document() IPDFDocument {
	rv := objc.Send[PDFDocument](p_.ID, objc.Sel("document"))
	return rv
}/* debug [instance_properties/getter]: document */


// Returns the index of the outline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFOutline/index
func (p_ PDFOutline) Index() uint {
	rv := objc.Send[uint](p_.ID, objc.Sel("index"))
	return rv
}/* debug [instance_properties/getter]: index */


// Returns a Boolean value that indicates whether the outline object is initially disclosed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFOutline/isOpen
func (p_ PDFOutline) IsOpen() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isOpen"))
	return rv
}/* debug [instance_properties/getter]: isOpen */


// Returns a Boolean value that indicates whether the outline object is initially disclosed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFOutline/isOpen
func (p_ PDFOutline) SetIsOpen(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsOpen:"), value)
}/* debug [instance_properties/setter]: isOpen */


// Returns the label for the outline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFOutline/label
func (p_ PDFOutline) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// Returns the label for the outline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFOutline/label
func (p_ PDFOutline) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// Returns the number of child outline objects in the outline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFOutline/numberOfChildren
func (p_ PDFOutline) NumberOfChildren() uint {
	rv := objc.Send[uint](p_.ID, objc.Sel("numberOfChildren"))
	return rv
}/* debug [instance_properties/getter]: numberOfChildren */


// Returns the parent outline object of the outline (returns if called on the root outline object).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFOutline/parent
func (p_ PDFOutline) Parent() IPDFOutline {
	rv := objc.Send[PDFOutline](p_.ID, objc.Sel("parent"))
	return rv
}/* debug [instance_properties/getter]: parent */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PDFOutline */


