// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class DOMStyleSheet */


/* debug [class_header]: Header for DOMStyleSheet */
// The class instance for the [DOMStyleSheet] class.
var (
	DOMStyleSheetClass     _DOMStyleSheetClass
	DOMStyleSheetClassOnce sync.Once
)

func getDOMStyleSheetClass() _DOMStyleSheetClass {
	DOMStyleSheetClassOnce.Do(func() {
		DOMStyleSheetClass = _DOMStyleSheetClass{objc.GetClass("DOMStyleSheet")}
	})
	return DOMStyleSheetClass
}

type _DOMStyleSheetClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMStyleSheet */
// An interface definition for the [DOMStyleSheet] class.
type IDOMStyleSheet interface {
	IDOMObject
	
/* debug [class_interface_properties]: Properties for DOMStyleSheet */
	// properties:
	Disabled() bool
	SetDisabled(value bool)
	Href() objc.IObject /* cross-framework: NSString */
	Media() IDOMMediaList
	OwnerNode() IDOMNode
	ParentStyleSheet() IDOMStyleSheet
	Title() objc.IObject /* cross-framework: NSString */
	Type() objc.IObject /* cross-framework: NSString */
	Parent() IDOMStyleSheet
	SetParent(value IDOMStyleSheet)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMStyleSheet */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMStyleSheet */
// Alloc allocates a new instance without initialization.
func (dc _DOMStyleSheetClass) Alloc() DOMStyleSheet {
	rv := objc.Send[DOMStyleSheet](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMStyleSheetClass) New() DOMStyleSheet {
	rv := objc.Send[DOMStyleSheet](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMStyleSheet) Init() DOMStyleSheet {
	rv := objc.Send[DOMStyleSheet](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMStyleSheet) Autorelease() DOMStyleSheet {
	rv := objc.Send[DOMStyleSheet](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMStyleSheet creates a new DOMStyleSheet instance.
func NewDOMStyleSheet() DOMStyleSheet {
	return getDOMStyleSheetClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMStyleSheet */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMStyleSheet
type DOMStyleSheet struct {
	DOMObject
}

// DOMStyleSheetFrom constructs a [DOMStyleSheet] from an unsafe.Pointer.
func DOMStyleSheetFrom(ptr unsafe.Pointer) DOMStyleSheet {
	return DOMStyleSheet{
		DOMObject: DOMObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMStyleSheet *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMStyleSheet */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMStyleSheet */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMStyleSheet */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMStyleSheet */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMStyleSheet/disabled
func (d_ DOMStyleSheet) Disabled() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("disabled"))
	return rv
}/* debug [instance_properties/getter]: disabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMStyleSheet/disabled
func (d_ DOMStyleSheet) SetDisabled(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDisabled:"), value)
}/* debug [instance_properties/setter]: disabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMStyleSheet/href
func (d_ DOMStyleSheet) Href() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("href"))
	return rv
}/* debug [instance_properties/getter]: href */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMStyleSheet/media
func (d_ DOMStyleSheet) Media() IDOMMediaList {
	rv := objc.Send[DOMMediaList](d_.ID, objc.Sel("media"))
	return rv
}/* debug [instance_properties/getter]: media */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMStyleSheet/ownerNode
func (d_ DOMStyleSheet) OwnerNode() IDOMNode {
	rv := objc.Send[DOMNode](d_.ID, objc.Sel("ownerNode"))
	return rv
}/* debug [instance_properties/getter]: ownerNode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMStyleSheet/parent
func (d_ DOMStyleSheet) ParentStyleSheet() IDOMStyleSheet {
	rv := objc.Send[DOMStyleSheet](d_.ID, objc.Sel("parentStyleSheet"))
	return rv
}/* debug [instance_properties/getter]: parentStyleSheet */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMStyleSheet/title
func (d_ DOMStyleSheet) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMStyleSheet/type
func (d_ DOMStyleSheet) Type() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/domstylesheet/parent
func (d_ DOMStyleSheet) Parent() IDOMStyleSheet {
	rv := objc.Send[DOMStyleSheet](d_.ID, objc.Sel("parent"))
	return rv
}/* debug [instance_properties/getter]: parent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/domstylesheet/parent
func (d_ DOMStyleSheet) SetParent(value IDOMStyleSheet) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setParent:"), value)
}/* debug [instance_properties/setter]: parent */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMStyleSheet */



