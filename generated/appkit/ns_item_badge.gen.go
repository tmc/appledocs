// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSItemBadge */


/* debug [class_header]: Header for NSItemBadge */
// The class instance for the [ItemBadge] class.
var (
	ItemBadgeClass     _ItemBadgeClass
	ItemBadgeClassOnce sync.Once
)

func getItemBadgeClass() _ItemBadgeClass {
	ItemBadgeClassOnce.Do(func() {
		ItemBadgeClass = _ItemBadgeClass{objc.GetClass("NSItemBadge")}
	})
	return ItemBadgeClass
}

type _ItemBadgeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ItemBadge */
// An interface definition for the [ItemBadge] class.
type IItemBadge interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ItemBadge */
	// properties:
	Text() objc.IObject /* cross-framework: NSString */
	IsBordered() bool
	SetIsBordered(value bool)
	IsEnabled() bool
	SetIsEnabled(value bool)
	IsHidden() bool
	SetIsHidden(value bool)
	IsNavigational() bool
	SetIsNavigational(value bool)
	IsVisible() bool
	SetIsVisible(value bool)
	Style() objectivec.IObject
	SetStyle(value objectivec.IObject)
	Tag() int
	SetTag(value int)
	VisibilityPriority() objectivec.IObject
	SetVisibilityPriority(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ItemBadge */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ItemBadge */
// Alloc allocates a new instance without initialization.
func (ic _ItemBadgeClass) Alloc() ItemBadge {
	rv := objc.Send[ItemBadge](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ItemBadgeClass) New() ItemBadge {
	rv := objc.Send[ItemBadge](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ItemBadge) Init() ItemBadge {
	rv := objc.Send[ItemBadge](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ItemBadge) Autorelease() ItemBadge {
	rv := objc.Send[ItemBadge](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewItemBadge creates a new ItemBadge instance.
func NewItemBadge() ItemBadge {
	return getItemBadgeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ItemBadge */
// represents a badge that can be attached to an .
//
// This badge provides a way to display small visual indicators, such as counts and text labels, within a toolbar item. Badges can be used to highlight important information, such as unread notifications or status indicators.


// represents a badge that can be attached to an .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSItemBadge-c.class
type ItemBadge struct {
	objectivec.Object
}

// ItemBadgeFrom constructs a [ItemBadge] from an unsafe.Pointer.
//
// represents a badge that can be attached to an .
func ItemBadgeFrom(ptr unsafe.Pointer) ItemBadge {
	return ItemBadge{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ItemBadge *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ItemBadge */

// Creates a badge displaying a localized numerical count.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSItemBadge-c.class/badgeWithCount:
func (ic _ItemBadgeClass) BadgeWithCount(count int) IItemBadge {
	rv := objc.Send[ItemBadge](objc.ID(ic.class), objc.Sel("badgeWithCount:"), count)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=BadgeWithCount) */


// Creates a badge displaying a text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSItemBadge-c.class/badgeWithText:
func (ic _ItemBadgeClass) BadgeWithText(text objc.IObject /* cross-framework: NSString */) IItemBadge {
	rv := objc.Send[ItemBadge](objc.ID(ic.class), objc.Sel("badgeWithText:"), text)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=BadgeWithText) */


// Creates a badge styled as an indicator. In this context, an indicator is simply a badge without any text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSItemBadge-c.class/indicatorBadge
func (ic _ItemBadgeClass) IndicatorBadge() IItemBadge {
	rv := objc.Send[ItemBadge](objc.ID(ic.class), objc.Sel("indicatorBadge"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=IndicatorBadge) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ItemBadge */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ItemBadge */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ItemBadge */

// The text to be displayed within the badge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSItemBadge-c.class/text
func (i_ ItemBadge) Text() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](i_.ID, objc.Sel("text"))
	return rv
}/* debug [instance_properties/getter]: text */


// A Boolean value that indicates whether the toolbar item has a bordered style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/isbordered
func (i_ ItemBadge) IsBordered() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isBordered"))
	return rv
}/* debug [instance_properties/getter]: isBordered */


// A Boolean value that indicates whether the toolbar item has a bordered style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/isbordered
func (i_ ItemBadge) SetIsBordered(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsBordered:"), value)
}/* debug [instance_properties/setter]: isBordered */


// A Boolean value that indicates whether the item is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/isenabled
func (i_ ItemBadge) IsEnabled() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isEnabled"))
	return rv
}/* debug [instance_properties/getter]: isEnabled */


// A Boolean value that indicates whether the item is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/isenabled
func (i_ ItemBadge) SetIsEnabled(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsEnabled:"), value)
}/* debug [instance_properties/setter]: isEnabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/ishidden
func (i_ ItemBadge) IsHidden() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isHidden"))
	return rv
}/* debug [instance_properties/getter]: isHidden */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/ishidden
func (i_ ItemBadge) SetIsHidden(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsHidden:"), value)
}/* debug [instance_properties/setter]: isHidden */


// A Boolean value that indicates whether the item behaves as a navigation item in the toolbar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/isnavigational
func (i_ ItemBadge) IsNavigational() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isNavigational"))
	return rv
}/* debug [instance_properties/getter]: isNavigational */


// A Boolean value that indicates whether the item behaves as a navigation item in the toolbar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/isnavigational
func (i_ ItemBadge) SetIsNavigational(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsNavigational:"), value)
}/* debug [instance_properties/setter]: isNavigational */


// A Boolean value that indicates whether the item is currently visible in the toolbar, and not in the overflow menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/isvisible
func (i_ ItemBadge) IsVisible() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isVisible"))
	return rv
}/* debug [instance_properties/getter]: isVisible */


// A Boolean value that indicates whether the item is currently visible in the toolbar, and not in the overflow menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/isvisible
func (i_ ItemBadge) SetIsVisible(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsVisible:"), value)
}/* debug [instance_properties/setter]: isVisible */


// Defines the toolbar item’s appearance. The default style is plain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/style-swift.property
func (i_ ItemBadge) Style() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("style"))
	return rv
}/* debug [instance_properties/getter]: style */


// Defines the toolbar item’s appearance. The default style is plain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/style-swift.property
func (i_ ItemBadge) SetStyle(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setStyle:"), value)
}/* debug [instance_properties/setter]: style */


// An integer tag you can use to identify the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/tag
func (i_ ItemBadge) Tag() int {
	rv := objc.Send[int](i_.ID, objc.Sel("tag"))
	return rv
}/* debug [instance_properties/getter]: tag */


// An integer tag you can use to identify the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/tag
func (i_ ItemBadge) SetTag(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTag:"), value)
}/* debug [instance_properties/setter]: tag */


// The display priority associated with the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/visibilitypriority-swift.property
func (i_ ItemBadge) VisibilityPriority() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("visibilityPriority"))
	return rv
}/* debug [instance_properties/getter]: visibilityPriority */


// The display priority associated with the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/visibilitypriority-swift.property
func (i_ ItemBadge) SetVisibilityPriority(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setVisibilityPriority:"), value)
}/* debug [instance_properties/setter]: visibilityPriority */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSItemBadge */



