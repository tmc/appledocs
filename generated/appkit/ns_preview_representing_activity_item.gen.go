// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSPreviewRepresentingActivityItem */


/* debug [class_header]: Header for NSPreviewRepresentingActivityItem */
// The class instance for the [PreviewRepresentingActivityItem] class.
var (
	PreviewRepresentingActivityItemClass     _PreviewRepresentingActivityItemClass
	PreviewRepresentingActivityItemClassOnce sync.Once
)

func getPreviewRepresentingActivityItemClass() _PreviewRepresentingActivityItemClass {
	PreviewRepresentingActivityItemClassOnce.Do(func() {
		PreviewRepresentingActivityItemClass = _PreviewRepresentingActivityItemClass{objc.GetClass("NSPreviewRepresentingActivityItem")}
	})
	return PreviewRepresentingActivityItemClass
}

type _PreviewRepresentingActivityItemClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PreviewRepresentingActivityItem */
// An interface definition for the [PreviewRepresentingActivityItem] class.
type IPreviewRepresentingActivityItem interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PreviewRepresentingActivityItem */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PreviewRepresentingActivityItem */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PreviewRepresentingActivityItem */
// Alloc allocates a new instance without initialization.
func (pc _PreviewRepresentingActivityItemClass) Alloc() PreviewRepresentingActivityItem {
	rv := objc.Send[PreviewRepresentingActivityItem](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PreviewRepresentingActivityItemClass) New() PreviewRepresentingActivityItem {
	rv := objc.Send[PreviewRepresentingActivityItem](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PreviewRepresentingActivityItem) Init() PreviewRepresentingActivityItem {
	rv := objc.Send[PreviewRepresentingActivityItem](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PreviewRepresentingActivityItem) Autorelease() PreviewRepresentingActivityItem {
	rv := objc.Send[PreviewRepresentingActivityItem](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPreviewRepresentingActivityItem creates a new PreviewRepresentingActivityItem instance.
func NewPreviewRepresentingActivityItem() PreviewRepresentingActivityItem {
	return getPreviewRepresentingActivityItemClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PreviewRepresentingActivityItem */
// A type that adds metadata to an item you share using the macOS share sheet.
//
// An object provides a concrete implementation of the protocol. Use it to create shareable items for common types such as strings or images, or when you don’t want to adopt the protocol directly in your app’s objects. To share the item from your app, initialize the object with this object.


// A type that adds metadata to an item you share using the macOS share sheet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPreviewRepresentingActivityItem
type PreviewRepresentingActivityItem struct {
	objectivec.Object
}

// PreviewRepresentingActivityItemFrom constructs a [PreviewRepresentingActivityItem] from an unsafe.Pointer.
//
// A type that adds metadata to an item you share using the macOS share sheet.
func PreviewRepresentingActivityItemFrom(ptr unsafe.Pointer) PreviewRepresentingActivityItem {
	return PreviewRepresentingActivityItem{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PreviewRepresentingActivityItem */

// Creates a metadata object with the title, image, and icon for a shareable item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPreviewRepresentingActivityItem/init(item:title:image:icon:)
func NewPreviewRepresentingActivityItemWithItemTitleImageIcon(item objc.IObject, title objc.IObject /* cross-framework: NSString */, image IImage, icon IImage) PreviewRepresentingActivityItem {
	instance := getPreviewRepresentingActivityItemClass().Alloc()
	rv := objc.Send[PreviewRepresentingActivityItem](instance.ID, objc.Sel("initWithItem:title:image:icon:"), item, title, image, icon)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPreviewRepresentingActivityItemWithItemTitleImageIcon */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PreviewRepresentingActivityItem */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PreviewRepresentingActivityItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PreviewRepresentingActivityItem */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PreviewRepresentingActivityItem */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSPreviewRepresentingActivityItem */


