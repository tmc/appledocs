// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [PDFOutline] class.
type IPDFOutline interface {
	objectivec.IObject
	ChildAtIndex(index uint) unsafe.Pointer
	InsertChildAtIndex(child unsafe.Pointer, index uint)
	RemoveFromParent()
}

// A object is an element in a tree-structured hierarchy that can represent the structure of a PDF document.
//
// An outline is an optional component of a PDF document, useful for viewing the structure of the document and for navigating within it. Outlines are created by the document’s author. If you represent a PDF document outline using outline objects, the root of the hierarchy is obtained from the PDF document itself. This root outline is not visible and serves merely as a container for the visible outlines.
//
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

// Alloc allocates a new instance without initialization.
func (pc _PDFOutlineClass) Alloc() PDFOutline {
	rv := objc.Send[PDFOutline](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Returns the child outline object at the specified index.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFOutline/child(at:)
func (p_ PDFOutline) ChildAtIndex(index uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("childAtIndex:"), index)
	return rv
}

// Inserts the specified outline object at the specified index.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFOutline/insertChild(_:at:)
func (p_ PDFOutline) InsertChildAtIndex(child unsafe.Pointer, index uint) {
	objc.Send[objc.ID](p_.ID, objc.Sel("insertChild:atIndex:"), child, index)
}

// Removes the outline object from its parent (does nothing if outline object is the root outline object).
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFOutline/removeFromParent()
func (p_ PDFOutline) RemoveFromParent() {
	objc.Send[objc.ID](p_.ID, objc.Sel("removeFromParent"))
}

// Returns the action performed when users click the outline.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFOutline/action
func (p_ PDFOutline) Action() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("action"))
	return rv
}


// SetAction sets the value of the action property.
// Returns the action performed when users click the outline.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFOutline/action
func (p_ PDFOutline) SetAction(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAction:"), value)
}

// Returns the destination associated with the outline.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFOutline/destination
func (p_ PDFOutline) Destination() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("destination"))
	return rv
}


// SetDestination sets the value of the destination property.
// Returns the destination associated with the outline.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFOutline/destination
func (p_ PDFOutline) SetDestination(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDestination:"), value)
}

// Returns the document with which the outline is associated.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFOutline/document
func (p_ PDFOutline) Document() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("document"))
	return rv
}

// Returns the index of the outline.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFOutline/index
func (p_ PDFOutline) Index() uint {
	rv := objc.Send[uint](p_.ID, objc.Sel("index"))
	return rv
}

// Returns a Boolean value that indicates whether the outline object is initially disclosed.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFOutline/isOpen
func (p_ PDFOutline) IsOpen() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isOpen"))
	return rv
}


// SetIsOpen sets the value of the isOpen property.
// Returns a Boolean value that indicates whether the outline object is initially disclosed.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFOutline/isOpen
func (p_ PDFOutline) SetIsOpen(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsOpen:"), value)
}

// Returns the label for the outline.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFOutline/label
func (p_ PDFOutline) Label() string {
	rv := objc.Send[string](p_.ID, objc.Sel("label"))
	return rv
}


// SetLabel sets the value of the label property.
// Returns the label for the outline.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFOutline/label
func (p_ PDFOutline) SetLabel(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLabel:"), objc.String(value))
}

// Returns the number of child outline objects in the outline.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFOutline/numberOfChildren
func (p_ PDFOutline) NumberOfChildren() uint {
	rv := objc.Send[uint](p_.ID, objc.Sel("numberOfChildren"))
	return rv
}

// Returns the parent outline object of the outline (returns if called on the root outline object).
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFOutline/parent
func (p_ PDFOutline) Parent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("parent"))
	return rv
}


