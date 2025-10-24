// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSGroupTouchBarItem */


/* debug [class_header]: Header for NSGroupTouchBarItem */
// The class instance for the [GroupTouchBarItem] class.
var (
	GroupTouchBarItemClass     _GroupTouchBarItemClass
	GroupTouchBarItemClassOnce sync.Once
)

func getGroupTouchBarItemClass() _GroupTouchBarItemClass {
	GroupTouchBarItemClassOnce.Do(func() {
		GroupTouchBarItemClass = _GroupTouchBarItemClass{objc.GetClass("NSGroupTouchBarItem")}
	})
	return GroupTouchBarItemClass
}

type _GroupTouchBarItemClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GroupTouchBarItem */
// An interface definition for the [GroupTouchBarItem] class.
type IGroupTouchBarItem interface {
	ITouchBarItem
	
/* debug [class_interface_properties]: Properties for GroupTouchBarItem */
	// properties:
	CustomizationLabel() objc.IObject /* cross-framework: NSString */
	SetCustomizationLabel(value objc.IObject /* cross-framework: NSString */)
	EffectiveCompressionOptions() IUserInterfaceCompressionOptions
	GroupTouchBar() objc.IObject /* cross-framework: TouchBar */
	SetGroupTouchBar(value objc.IObject /* cross-framework: TouchBar */)
	GroupUserInterfaceLayoutDirection() UserInterfaceLayoutDirection
	SetGroupUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection)
	PreferredItemWidth() float64
	SetPreferredItemWidth(value float64)
	PrefersEqualWidths() bool
	SetPrefersEqualWidths(value bool)
	PrioritizedCompressionOptions() []UserInterfaceCompressionOptions
	SetPrioritizedCompressionOptions(value []UserInterfaceCompressionOptions)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GroupTouchBarItem */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GroupTouchBarItem */
// Alloc allocates a new instance without initialization.
func (gc _GroupTouchBarItemClass) Alloc() GroupTouchBarItem {
	rv := objc.Send[GroupTouchBarItem](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GroupTouchBarItemClass) New() GroupTouchBarItem {
	rv := objc.Send[GroupTouchBarItem](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GroupTouchBarItem) Init() GroupTouchBarItem {
	rv := objc.Send[GroupTouchBarItem](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GroupTouchBarItem) Autorelease() GroupTouchBarItem {
	rv := objc.Send[GroupTouchBarItem](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGroupTouchBarItem creates a new GroupTouchBarItem instance.
func NewGroupTouchBarItem() GroupTouchBarItem {
	return getGroupTouchBarItemClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GroupTouchBarItem */
// A bar item that provides a bar to contain other items.


// A bar item that provides a bar to contain other items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem
type GroupTouchBarItem struct {
	TouchBarItem
}

// GroupTouchBarItemFrom constructs a [GroupTouchBarItem] from an unsafe.Pointer.
//
// A bar item that provides a bar to contain other items.
func GroupTouchBarItemFrom(ptr unsafe.Pointer) GroupTouchBarItem {
	return GroupTouchBarItem{
		TouchBarItem: TouchBarItemFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GroupTouchBarItem */

// Initializes and returns a group item configured to match system alerts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/init(alertStyleWithIdentifier:)
func NewGroupTouchBarItemAlertStyleGroupItemWithIdentifier(identifier TouchBarItemIdentifier /* typedef */) GroupTouchBarItem {
	rv := objc.Send[GroupTouchBarItem](objc.ID(getGroupTouchBarItemClass().class), objc.Sel("alertStyleGroupItemWithIdentifier:"), identifier)
	return rv
}/* debug [class_init_methods/constructor]: NewGroupTouchBarItemAlertStyleGroupItemWithIdentifier */


// Initializes and returns a group item whose bar is constructed from the supplied items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/init(identifier:items:)
func NewGroupTouchBarItemGroupItemWithIdentifierItems(identifier TouchBarItemIdentifier /* typedef */, items []TouchBarItem) GroupTouchBarItem {
	rv := objc.Send[GroupTouchBarItem](objc.ID(getGroupTouchBarItemClass().class), objc.Sel("groupItemWithIdentifier:items:"), identifier, items)
	return rv
}/* debug [class_init_methods/constructor]: NewGroupTouchBarItemGroupItemWithIdentifierItems */


// Initializes and returns a group item whose bar is constructed from the supplied items, and with the specified compression options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/init(identifier:items:allowedCompressionOptions:)
func NewGroupTouchBarItemGroupItemWithIdentifierItemsAllowedCompressionOptions(identifier TouchBarItemIdentifier /* typedef */, items []TouchBarItem, allowedCompressionOptions IUserInterfaceCompressionOptions) GroupTouchBarItem {
	rv := objc.Send[GroupTouchBarItem](objc.ID(getGroupTouchBarItemClass().class), objc.Sel("groupItemWithIdentifier:items:allowedCompressionOptions:"), identifier, items, allowedCompressionOptions)
	return rv
}/* debug [class_init_methods/constructor]: NewGroupTouchBarItemGroupItemWithIdentifierItemsAllowedCompressionOptions */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GroupTouchBarItem */

// Initializes and returns a group item configured to match system alerts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/init(alertStyleWithIdentifier:)
func (gc _GroupTouchBarItemClass) AlertStyleGroupItemWithIdentifier(identifier TouchBarItemIdentifier /* typedef */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(gc.class), objc.Sel("alertStyleGroupItemWithIdentifier:"), identifier)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AlertStyleGroupItemWithIdentifier) */


// Initializes and returns a group item whose bar is constructed from the supplied items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/init(identifier:items:)
func (gc _GroupTouchBarItemClass) GroupItemWithIdentifierItems(identifier TouchBarItemIdentifier /* typedef */, items []TouchBarItem) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(gc.class), objc.Sel("groupItemWithIdentifier:items:"), identifier, items)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GroupItemWithIdentifierItems) */


// Initializes and returns a group item whose bar is constructed from the supplied items, and with the specified compression options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/init(identifier:items:allowedCompressionOptions:)
func (gc _GroupTouchBarItemClass) GroupItemWithIdentifierItemsAllowedCompressionOptions(identifier TouchBarItemIdentifier /* typedef */, items []TouchBarItem, allowedCompressionOptions IUserInterfaceCompressionOptions) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(gc.class), objc.Sel("groupItemWithIdentifier:items:allowedCompressionOptions:"), identifier, items, allowedCompressionOptions)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GroupItemWithIdentifierItemsAllowedCompressionOptions) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GroupTouchBarItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GroupTouchBarItem */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GroupTouchBarItem */

// The user-visible string identifying this item during bar customization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/customizationLabel
func (g_ GroupTouchBarItem) CustomizationLabel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("customizationLabel"))
	return rv
}/* debug [instance_properties/getter]: customizationLabel */


// The user-visible string identifying this item during bar customization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/customizationLabel
func (g_ GroupTouchBarItem) SetCustomizationLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCustomizationLabel:"), value)
}/* debug [instance_properties/setter]: customizationLabel */


// The compression options that are currently active on the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/effectiveCompressionOptions
func (g_ GroupTouchBarItem) EffectiveCompressionOptions() IUserInterfaceCompressionOptions {
	rv := objc.Send[UserInterfaceCompressionOptions](g_.ID, objc.Sel("effectiveCompressionOptions"))
	return rv
}/* debug [instance_properties/getter]: effectiveCompressionOptions */


// A bar that holds this group’s items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/groupTouchBar
func (g_ GroupTouchBarItem) GroupTouchBar() objc.IObject /* cross-framework: TouchBar */ {
	rv := objc.Send[objc.ID](g_.ID, objc.Sel("groupTouchBar"))
	return rv
}/* debug [instance_properties/getter]: groupTouchBar */


// A bar that holds this group’s items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/groupTouchBar
func (g_ GroupTouchBarItem) SetGroupTouchBar(value objc.IObject /* cross-framework: TouchBar */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGroupTouchBar:"), value)
}/* debug [instance_properties/setter]: groupTouchBar */


// The user interface direction that controls the layout order of the items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/groupUserInterfaceLayoutDirection
func (g_ GroupTouchBarItem) GroupUserInterfaceLayoutDirection() UserInterfaceLayoutDirection {
	rv := objc.Send[UserInterfaceLayoutDirection](g_.ID, objc.Sel("groupUserInterfaceLayoutDirection"))
	return rv
}/* debug [instance_properties/getter]: groupUserInterfaceLayoutDirection */


// The user interface direction that controls the layout order of the items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/groupUserInterfaceLayoutDirection
func (g_ GroupTouchBarItem) SetGroupUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGroupUserInterfaceLayoutDirection:"), value)
}/* debug [instance_properties/setter]: groupUserInterfaceLayoutDirection */


// The preferred width for items in the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/preferredItemWidth
func (g_ GroupTouchBarItem) PreferredItemWidth() float64 {
	rv := objc.Send[float64](g_.ID, objc.Sel("preferredItemWidth"))
	return rv
}/* debug [instance_properties/getter]: preferredItemWidth */


// The preferred width for items in the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/preferredItemWidth
func (g_ GroupTouchBarItem) SetPreferredItemWidth(value float64) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPreferredItemWidth:"), value)
}/* debug [instance_properties/setter]: preferredItemWidth */


// A Boolean value that specifies that items should have equal widths when possible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/prefersEqualWidths
func (g_ GroupTouchBarItem) PrefersEqualWidths() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("prefersEqualWidths"))
	return rv
}/* debug [instance_properties/getter]: prefersEqualWidths */


// A Boolean value that specifies that items should have equal widths when possible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/prefersEqualWidths
func (g_ GroupTouchBarItem) SetPrefersEqualWidths(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPrefersEqualWidths:"), value)
}/* debug [instance_properties/setter]: prefersEqualWidths */


// The allowed compression options, in the order they should be applied.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/prioritizedCompressionOptions
func (g_ GroupTouchBarItem) PrioritizedCompressionOptions() []UserInterfaceCompressionOptions {
	rv := objc.Send[[]UserInterfaceCompressionOptions](g_.ID, objc.Sel("prioritizedCompressionOptions"))
	return rv
}/* debug [instance_properties/getter]: prioritizedCompressionOptions */


// The allowed compression options, in the order they should be applied.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/prioritizedCompressionOptions
func (g_ GroupTouchBarItem) SetPrioritizedCompressionOptions(value []UserInterfaceCompressionOptions) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](g_.ID, objc.Sel("setPrioritizedCompressionOptions:"), nsArray)
}/* debug [instance_properties/setter]: prioritizedCompressionOptions */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSGroupTouchBarItem */


