// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class NSPathComponentCell */


/* debug [class_header]: Header for NSPathComponentCell */
// The class instance for the [PathComponentCell] class.
var (
	PathComponentCellClass     _PathComponentCellClass
	PathComponentCellClassOnce sync.Once
)

func getPathComponentCellClass() _PathComponentCellClass {
	PathComponentCellClassOnce.Do(func() {
		PathComponentCellClass = _PathComponentCellClass{objc.GetClass("NSPathComponentCell")}
	})
	return PathComponentCellClass
}

type _PathComponentCellClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PathComponentCell */
// An interface definition for the [PathComponentCell] class.
type IPathComponentCell interface {
	ITextFieldCell
	
/* debug [class_interface_properties]: Properties for PathComponentCell */
	// properties:
	Image() IImage
	SetImage(value IImage)
	URL() objc.IObject /* cross-framework: NSURL */
	SetURL(value objc.IObject /* cross-framework: NSURL */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PathComponentCell */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PathComponentCell */
// Alloc allocates a new instance without initialization.
func (pc _PathComponentCellClass) Alloc() PathComponentCell {
	rv := objc.Send[PathComponentCell](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PathComponentCellClass) New() PathComponentCell {
	rv := objc.Send[PathComponentCell](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PathComponentCell) Init() PathComponentCell {
	rv := objc.Send[PathComponentCell](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PathComponentCell) Autorelease() PathComponentCell {
	rv := objc.Send[PathComponentCell](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPathComponentCell creates a new PathComponentCell instance.
func NewPathComponentCell() PathComponentCell {
	return getPathComponentCellClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PathComponentCell */
// A component of a path.
//
// An object manages a collection of objects, in conjunction with an object, to represent a path.


// A component of a path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathComponentCell
type PathComponentCell struct {
	TextFieldCell
}

// PathComponentCellFrom constructs a [PathComponentCell] from an unsafe.Pointer.
//
// A component of a path.
func PathComponentCellFrom(ptr unsafe.Pointer) PathComponentCell {
	return PathComponentCell{
		TextFieldCell: TextFieldCellFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PathComponentCell *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PathComponentCell */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PathComponentCell */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PathComponentCell */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PathComponentCell */

// The image displayed for this component cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathComponentCell/image
func (p_ PathComponentCell) Image() IImage {
	rv := objc.Send[Image](p_.ID, objc.Sel("image"))
	return rv
}/* debug [instance_properties/getter]: image */


// The image displayed for this component cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathComponentCell/image
func (p_ PathComponentCell) SetImage(value IImage) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setImage:"), value)
}/* debug [instance_properties/setter]: image */


// The portion of the path from the root through the component represented by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathComponentCell/url
func (p_ PathComponentCell) URL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](p_.ID, objc.Sel("URL"))
	return rv
}/* debug [instance_properties/getter]: URL */


// The portion of the path from the root through the component represented by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathComponentCell/url
func (p_ PathComponentCell) SetURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setURL:"), value)
}/* debug [instance_properties/setter]: URL */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSPathComponentCell */



