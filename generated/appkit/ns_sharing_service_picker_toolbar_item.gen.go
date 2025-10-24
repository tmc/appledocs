// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSSharingServicePickerToolbarItem */


/* debug [class_header]: Header for NSSharingServicePickerToolbarItem */
// The class instance for the [SharingServicePickerToolbarItem] class.
var (
	SharingServicePickerToolbarItemClass     _SharingServicePickerToolbarItemClass
	SharingServicePickerToolbarItemClassOnce sync.Once
)

func getSharingServicePickerToolbarItemClass() _SharingServicePickerToolbarItemClass {
	SharingServicePickerToolbarItemClassOnce.Do(func() {
		SharingServicePickerToolbarItemClass = _SharingServicePickerToolbarItemClass{objc.GetClass("NSSharingServicePickerToolbarItem")}
	})
	return SharingServicePickerToolbarItemClass
}

type _SharingServicePickerToolbarItemClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SharingServicePickerToolbarItem */
// An interface definition for the [SharingServicePickerToolbarItem] class.
type ISharingServicePickerToolbarItem interface {
	IToolbarItem
	
/* debug [class_interface_properties]: Properties for SharingServicePickerToolbarItem */
	// properties:
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SharingServicePickerToolbarItem */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SharingServicePickerToolbarItem */
// Alloc allocates a new instance without initialization.
func (sc _SharingServicePickerToolbarItemClass) Alloc() SharingServicePickerToolbarItem {
	rv := objc.Send[SharingServicePickerToolbarItem](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SharingServicePickerToolbarItemClass) New() SharingServicePickerToolbarItem {
	rv := objc.Send[SharingServicePickerToolbarItem](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SharingServicePickerToolbarItem) Init() SharingServicePickerToolbarItem {
	rv := objc.Send[SharingServicePickerToolbarItem](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SharingServicePickerToolbarItem) Autorelease() SharingServicePickerToolbarItem {
	rv := objc.Send[SharingServicePickerToolbarItem](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSharingServicePickerToolbarItem creates a new SharingServicePickerToolbarItem instance.
func NewSharingServicePickerToolbarItem() SharingServicePickerToolbarItem {
	return getSharingServicePickerToolbarItemClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SharingServicePickerToolbarItem */
// A toolbar item that displays the macOS share sheet.
//
// An object is a standard item you add to your window’s toolbar. When someone clicks it, the item displays the macOS share sheet. Use this item to share the selected or focal content from the current window. For example, you might share the photo someone is viewing, the currently selected text, or the window’s associated document. Provide the items to share using the associated object. For an app built using Mac Catalyst, provide the items from the object in the property.


// A toolbar item that displays the macOS share sheet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerToolbarItem
type SharingServicePickerToolbarItem struct {
	ToolbarItem
}

// SharingServicePickerToolbarItemFrom constructs a [SharingServicePickerToolbarItem] from an unsafe.Pointer.
//
// A toolbar item that displays the macOS share sheet.
func SharingServicePickerToolbarItemFrom(ptr unsafe.Pointer) SharingServicePickerToolbarItem {
	return SharingServicePickerToolbarItem{
		ToolbarItem: ToolbarItemFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SharingServicePickerToolbarItem *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SharingServicePickerToolbarItem */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SharingServicePickerToolbarItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SharingServicePickerToolbarItem */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SharingServicePickerToolbarItem */

// The custom object from your app that provides the items to share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerToolbarItem/delegate
func (s_ SharingServicePickerToolbarItem) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The custom object from your app that provides the items to share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerToolbarItem/delegate
func (s_ SharingServicePickerToolbarItem) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSSharingServicePickerToolbarItem */


