// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSPathControlItem */


/* debug [class_header]: Header for NSPathControlItem */
// The class instance for the [PathControlItem] class.
var (
	PathControlItemClass     _PathControlItemClass
	PathControlItemClassOnce sync.Once
)

func getPathControlItemClass() _PathControlItemClass {
	PathControlItemClassOnce.Do(func() {
		PathControlItemClass = _PathControlItemClass{objc.GetClass("NSPathControlItem")}
	})
	return PathControlItemClass
}

type _PathControlItemClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PathControlItem */
// An interface definition for the [PathControlItem] class.
type IPathControlItem interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PathControlItem */
	// properties:
	AttributedTitle() foundation.AttributedString
	SetAttributedTitle(value foundation.AttributedString)
	Image() IImage
	SetImage(value IImage)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	URL() objc.IObject /* cross-framework: NSURL */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PathControlItem */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PathControlItem */
// Alloc allocates a new instance without initialization.
func (pc _PathControlItemClass) Alloc() PathControlItem {
	rv := objc.Send[PathControlItem](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PathControlItemClass) New() PathControlItem {
	rv := objc.Send[PathControlItem](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PathControlItem) Init() PathControlItem {
	rv := objc.Send[PathControlItem](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PathControlItem) Autorelease() PathControlItem {
	rv := objc.Send[PathControlItem](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPathControlItem creates a new PathControlItem instance.
func NewPathControlItem() PathControlItem {
	return getPathControlItemClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PathControlItem */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControlItem
type PathControlItem struct {
	objectivec.Object
}

// PathControlItemFrom constructs a [PathControlItem] from an unsafe.Pointer.
func PathControlItemFrom(ptr unsafe.Pointer) PathControlItem {
	return PathControlItem{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PathControlItem *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PathControlItem */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PathControlItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PathControlItem */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PathControlItem */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControlItem/attributedTitle
func (p_ PathControlItem) AttributedTitle() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](p_.ID, objc.Sel("attributedTitle"))
	return rv
}/* debug [instance_properties/getter]: attributedTitle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControlItem/attributedTitle
func (p_ PathControlItem) SetAttributedTitle(value foundation.AttributedString) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAttributedTitle:"), value)
}/* debug [instance_properties/setter]: attributedTitle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControlItem/image
func (p_ PathControlItem) Image() IImage {
	rv := objc.Send[Image](p_.ID, objc.Sel("image"))
	return rv
}/* debug [instance_properties/getter]: image */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControlItem/image
func (p_ PathControlItem) SetImage(value IImage) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setImage:"), value)
}/* debug [instance_properties/setter]: image */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControlItem/title
func (p_ PathControlItem) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControlItem/title
func (p_ PathControlItem) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTitle:"), value)
}/* debug [instance_properties/setter]: title */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControlItem/url
func (p_ PathControlItem) URL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](p_.ID, objc.Sel("URL"))
	return rv
}/* debug [instance_properties/getter]: URL */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSPathControlItem */



