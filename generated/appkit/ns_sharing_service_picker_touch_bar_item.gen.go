// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class NSSharingServicePickerTouchBarItem */


/* debug [class_header]: Header for NSSharingServicePickerTouchBarItem */
// The class instance for the [SharingServicePickerTouchBarItem] class.
var (
	SharingServicePickerTouchBarItemClass     _SharingServicePickerTouchBarItemClass
	SharingServicePickerTouchBarItemClassOnce sync.Once
)

func getSharingServicePickerTouchBarItemClass() _SharingServicePickerTouchBarItemClass {
	SharingServicePickerTouchBarItemClassOnce.Do(func() {
		SharingServicePickerTouchBarItemClass = _SharingServicePickerTouchBarItemClass{objc.GetClass("NSSharingServicePickerTouchBarItem")}
	})
	return SharingServicePickerTouchBarItemClass
}

type _SharingServicePickerTouchBarItemClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SharingServicePickerTouchBarItem */
// An interface definition for the [SharingServicePickerTouchBarItem] class.
type ISharingServicePickerTouchBarItem interface {
	ITouchBarItem
	
/* debug [class_interface_properties]: Properties for SharingServicePickerTouchBarItem */
	// properties:
	ButtonImage() IImage
	SetButtonImage(value IImage)
	ButtonTitle() objc.IObject /* cross-framework: NSString */
	SetButtonTitle(value objc.IObject /* cross-framework: NSString */)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	Enabled() bool
	SetEnabled(value bool)
	IsEnabled() bool
	SetIsEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SharingServicePickerTouchBarItem */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SharingServicePickerTouchBarItem */
// Alloc allocates a new instance without initialization.
func (sc _SharingServicePickerTouchBarItemClass) Alloc() SharingServicePickerTouchBarItem {
	rv := objc.Send[SharingServicePickerTouchBarItem](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SharingServicePickerTouchBarItemClass) New() SharingServicePickerTouchBarItem {
	rv := objc.Send[SharingServicePickerTouchBarItem](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SharingServicePickerTouchBarItem) Init() SharingServicePickerTouchBarItem {
	rv := objc.Send[SharingServicePickerTouchBarItem](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SharingServicePickerTouchBarItem) Autorelease() SharingServicePickerTouchBarItem {
	rv := objc.Send[SharingServicePickerTouchBarItem](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSharingServicePickerTouchBarItem creates a new SharingServicePickerTouchBarItem instance.
func NewSharingServicePickerTouchBarItem() SharingServicePickerTouchBarItem {
	return getSharingServicePickerTouchBarItemClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SharingServicePickerTouchBarItem */
// A bar item that, along with its delegate, provides a list of objects eligible for sharing.


// A bar item that, along with its delegate, provides a list of objects eligible for sharing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerTouchBarItem
type SharingServicePickerTouchBarItem struct {
	TouchBarItem
}

// SharingServicePickerTouchBarItemFrom constructs a [SharingServicePickerTouchBarItem] from an unsafe.Pointer.
//
// A bar item that, along with its delegate, provides a list of objects eligible for sharing.
func SharingServicePickerTouchBarItemFrom(ptr unsafe.Pointer) SharingServicePickerTouchBarItem {
	return SharingServicePickerTouchBarItem{
		TouchBarItem: TouchBarItemFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SharingServicePickerTouchBarItem *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SharingServicePickerTouchBarItem */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SharingServicePickerTouchBarItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SharingServicePickerTouchBarItem */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SharingServicePickerTouchBarItem */

// The image displayed in the sharing service picker item button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerTouchBarItem/buttonImage
func (s_ SharingServicePickerTouchBarItem) ButtonImage() IImage {
	rv := objc.Send[Image](s_.ID, objc.Sel("buttonImage"))
	return rv
}/* debug [instance_properties/getter]: buttonImage */


// The image displayed in the sharing service picker item button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerTouchBarItem/buttonImage
func (s_ SharingServicePickerTouchBarItem) SetButtonImage(value IImage) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setButtonImage:"), value)
}/* debug [instance_properties/setter]: buttonImage */


// The text displayed in the sharing service picker item button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerTouchBarItem/buttonTitle
func (s_ SharingServicePickerTouchBarItem) ButtonTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("buttonTitle"))
	return rv
}/* debug [instance_properties/getter]: buttonTitle */


// The text displayed in the sharing service picker item button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerTouchBarItem/buttonTitle
func (s_ SharingServicePickerTouchBarItem) SetButtonTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setButtonTitle:"), value)
}/* debug [instance_properties/setter]: buttonTitle */


// The object that acts as the delegate of the sharing service picker bar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerTouchBarItem/delegate
func (s_ SharingServicePickerTouchBarItem) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The object that acts as the delegate of the sharing service picker bar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerTouchBarItem/delegate
func (s_ SharingServicePickerTouchBarItem) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// A Boolean value that specifies whether the sharing service picker item is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerTouchBarItem/isEnabled
func (s_ SharingServicePickerTouchBarItem) Enabled() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("enabled"))
	return rv
}/* debug [instance_properties/getter]: enabled */


// A Boolean value that specifies whether the sharing service picker item is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerTouchBarItem/isEnabled
func (s_ SharingServicePickerTouchBarItem) SetEnabled(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setEnabled:"), value)
}/* debug [instance_properties/setter]: enabled */


// A Boolean value that specifies whether the sharing service picker item is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertouchbaritem/isenabled
func (s_ SharingServicePickerTouchBarItem) IsEnabled() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isEnabled"))
	return rv
}/* debug [instance_properties/getter]: isEnabled */


// A Boolean value that specifies whether the sharing service picker item is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepickertouchbaritem/isenabled
func (s_ SharingServicePickerTouchBarItem) SetIsEnabled(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsEnabled:"), value)
}/* debug [instance_properties/setter]: isEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSSharingServicePickerTouchBarItem */


