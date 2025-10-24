// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSTextAttachmentCell */


/* debug [class_header]: Header for NSTextAttachmentCell */
// The class instance for the [TextAttachmentCell] class.
var (
	TextAttachmentCellClass     _TextAttachmentCellClass
	TextAttachmentCellClassOnce sync.Once
)

func getTextAttachmentCellClass() _TextAttachmentCellClass {
	TextAttachmentCellClassOnce.Do(func() {
		TextAttachmentCellClass = _TextAttachmentCellClass{objc.GetClass("NSTextAttachmentCell")}
	})
	return TextAttachmentCellClass
}

type _TextAttachmentCellClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TextAttachmentCell */
// An interface definition for the [TextAttachmentCell] class.
type ITextAttachmentCell interface {
	ICell
	
/* debug [class_interface_properties]: Properties for TextAttachmentCell */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TextAttachmentCell */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TextAttachmentCell */
// Alloc allocates a new instance without initialization.
func (tc _TextAttachmentCellClass) Alloc() TextAttachmentCell {
	rv := objc.Send[TextAttachmentCell](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TextAttachmentCellClass) New() TextAttachmentCell {
	rv := objc.Send[TextAttachmentCell](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextAttachmentCell) Init() TextAttachmentCell {
	rv := objc.Send[TextAttachmentCell](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextAttachmentCell) Autorelease() TextAttachmentCell {
	rv := objc.Send[TextAttachmentCell](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextAttachmentCell creates a new TextAttachmentCell instance.
func NewTextAttachmentCell() TextAttachmentCell {
	return getTextAttachmentCellClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TextAttachmentCell */
// An object that implements the functionality of the text attachment cell protocol.
//
// This specification describes only those methods whose implementations have features that are particular to this class. For a general discussion of the protocol’s methods, see .


// An object that implements the functionality of the text attachment cell protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachmentCell-swift.class
type TextAttachmentCell struct {
	Cell
}

// TextAttachmentCellFrom constructs a [TextAttachmentCell] from an unsafe.Pointer.
//
// An object that implements the functionality of the text attachment cell protocol.
func TextAttachmentCellFrom(ptr unsafe.Pointer) TextAttachmentCell {
	return TextAttachmentCell{
		Cell: CellFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TextAttachmentCell *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TextAttachmentCell */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TextAttachmentCell */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TextAttachmentCell */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TextAttachmentCell */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTextAttachmentCell */



