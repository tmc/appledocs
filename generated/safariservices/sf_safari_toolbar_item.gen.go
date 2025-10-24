// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SFSafariToolbarItem */


/* debug [class_header]: Header for SFSafariToolbarItem */
// The class instance for the [SFSafariToolbarItem] class.
var (
	SFSafariToolbarItemClass     _SFSafariToolbarItemClass
	SFSafariToolbarItemClassOnce sync.Once
)

func getSFSafariToolbarItemClass() _SFSafariToolbarItemClass {
	SFSafariToolbarItemClassOnce.Do(func() {
		SFSafariToolbarItemClass = _SFSafariToolbarItemClass{objc.GetClass("SFSafariToolbarItem")}
	})
	return SFSafariToolbarItemClass
}

type _SFSafariToolbarItemClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SFSafariToolbarItem */
// An interface definition for the [SFSafariToolbarItem] class.
type ISFSafariToolbarItem interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SFSafariToolbarItem */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SFSafariToolbarItem */
	// methods:
	SetBadgeText(badgeText objc.IObject /* cross-framework: NSString */)
	SetEnabled(enabled bool)
	SetImage(image appkit.Image)
	SetLabel(label objc.IObject /* cross-framework: NSString */)
	ShowPopover()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SFSafariToolbarItem */
// Alloc allocates a new instance without initialization.
func (sc _SFSafariToolbarItemClass) Alloc() SFSafariToolbarItem {
	rv := objc.Send[SFSafariToolbarItem](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SFSafariToolbarItemClass) New() SFSafariToolbarItem {
	rv := objc.Send[SFSafariToolbarItem](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFSafariToolbarItem) Init() SFSafariToolbarItem {
	rv := objc.Send[SFSafariToolbarItem](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFSafariToolbarItem) Autorelease() SFSafariToolbarItem {
	rv := objc.Send[SFSafariToolbarItem](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSafariToolbarItem creates a new SFSafariToolbarItem instance.
func NewSFSafariToolbarItem() SFSafariToolbarItem {
	return getSFSafariToolbarItemClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SFSafariToolbarItem */
// A proxy for a Safari app extension toolbar item in a Safari window.
//
// Your app extension only uses this object when it wants to explicitly set the toolbar item state. Typically, other state changes occur automatically. Safari calls on your app extension handler when changes, such as navigation to a webpage, could affect the state of the toolbar item.


// A proxy for a Safari app extension toolbar item in a Safari window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariToolbarItem
type SFSafariToolbarItem struct {
	objectivec.Object
}

// SFSafariToolbarItemFrom constructs a [SFSafariToolbarItem] from an unsafe.Pointer.
//
// A proxy for a Safari app extension toolbar item in a Safari window.
func SFSafariToolbarItemFrom(ptr unsafe.Pointer) SFSafariToolbarItem {
	return SFSafariToolbarItem{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SFSafariToolbarItem *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SFSafariToolbarItem */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SFSafariToolbarItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SFSafariToolbarItem */

// Sets the badge text for the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariToolbarItem/setBadgeText(_:)
func (s_ SFSafariToolbarItem) SetBadgeText(badgeText objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBadgeText:"), badgeText)
}/* debug [instance_methods/method]: SetBadgeText */


// Sets whether the toolbar item is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariToolbarItem/setEnabled(_:)
func (s_ SFSafariToolbarItem) SetEnabled(enabled bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setEnabled:"), enabled)
}/* debug [instance_methods/method]: SetEnabled */


// Sets the image displayed in the toolbar button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariToolbarItem/setImage(_:)
func (s_ SFSafariToolbarItem) SetImage(image appkit.Image) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setImage:"), image)
}/* debug [instance_methods/method]: SetImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariToolbarItem/setLabel(_:)
func (s_ SFSafariToolbarItem) SetLabel(label objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLabel:"), label)
}/* debug [instance_methods/method]: SetLabel */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariToolbarItem/showPopover()
func (s_ SFSafariToolbarItem) ShowPopover() {
	objc.Send[objc.ID](s_.ID, objc.Sel("showPopover"))
}/* debug [instance_methods/method]: ShowPopover */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SFSafariToolbarItem */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SFSafariToolbarItem */





