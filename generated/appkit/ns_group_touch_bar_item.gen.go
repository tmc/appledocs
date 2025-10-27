// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [GroupTouchBarItem] class.
type IGroupTouchBarItem interface {
	ITouchBarItem
	

	// properties:
	CustomizationLabel() foundation.foundation.INSString
	SetCustomizationLabel(value foundation.foundation.INSString)
	EffectiveCompressionOptions() IUserInterfaceCompressionOptions
	GroupTouchBar() TouchBar /* not a class type */
	SetGroupTouchBar(value TouchBar /* not a class type */)
	GroupUserInterfaceLayoutDirection() UserInterfaceLayoutDirection
	SetGroupUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection)
	PreferredItemWidth() float64
	SetPreferredItemWidth(value float64)
	PrefersEqualWidths() bool
	SetPrefersEqualWidths(value bool)
	PrioritizedCompressionOptions() []UserInterfaceCompressionOptions
	SetPrioritizedCompressionOptions(value []UserInterfaceCompressionOptions)


	

	// methods:


}





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






// Initializes and returns a group item configured to match system alerts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/init(alertStyleWithIdentifier:)
func NewGroupTouchBarItemAlertStyleGroupItemWithIdentifier(identifier TouchBarItemIdentifier) GroupTouchBarItem {
	rv := objc.Send[GroupTouchBarItem](objc.ID(getGroupTouchBarItemClass().class), objc.Sel("alertStyleGroupItemWithIdentifier:"), identifier)
	return rv
}


// Initializes and returns a group item whose bar is constructed from the supplied items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/init(identifier:items:)
func NewGroupTouchBarItemGroupItemWithIdentifierItems(identifier TouchBarItemIdentifier, items []TouchBarItem) GroupTouchBarItem {
	rv := objc.Send[GroupTouchBarItem](objc.ID(getGroupTouchBarItemClass().class), objc.Sel("groupItemWithIdentifier:items:"), identifier, items)
	return rv
}


// Initializes and returns a group item whose bar is constructed from the supplied items, and with the specified compression options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/init(identifier:items:allowedCompressionOptions:)
func NewGroupTouchBarItemGroupItemWithIdentifierItemsAllowedCompressionOptions(identifier TouchBarItemIdentifier, items []TouchBarItem, allowedCompressionOptions IUserInterfaceCompressionOptions) GroupTouchBarItem {
	rv := objc.Send[GroupTouchBarItem](objc.ID(getGroupTouchBarItemClass().class), objc.Sel("groupItemWithIdentifier:items:allowedCompressionOptions:"), identifier, items, allowedCompressionOptions)
	return rv
}







// Initializes and returns a group item configured to match system alerts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/init(alertStyleWithIdentifier:)
func (gc _GroupTouchBarItemClass) AlertStyleGroupItemWithIdentifier(identifier TouchBarItemIdentifier) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(gc.class), objc.Sel("alertStyleGroupItemWithIdentifier:"), identifier)
	return rv
}


// Initializes and returns a group item whose bar is constructed from the supplied items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/init(identifier:items:)
func (gc _GroupTouchBarItemClass) GroupItemWithIdentifierItems(identifier TouchBarItemIdentifier, items []TouchBarItem) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(gc.class), objc.Sel("groupItemWithIdentifier:items:"), identifier, items)
	return rv
}


// Initializes and returns a group item whose bar is constructed from the supplied items, and with the specified compression options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/init(identifier:items:allowedCompressionOptions:)
func (gc _GroupTouchBarItemClass) GroupItemWithIdentifierItemsAllowedCompressionOptions(identifier TouchBarItemIdentifier, items []TouchBarItem, allowedCompressionOptions IUserInterfaceCompressionOptions) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(gc.class), objc.Sel("groupItemWithIdentifier:items:allowedCompressionOptions:"), identifier, items, allowedCompressionOptions)
	return rv
}

















// The user-visible string identifying this item during bar customization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/customizationLabel
func (g_ GroupTouchBarItem) CustomizationLabel() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("customizationLabel"))
	return rv
}


// The user-visible string identifying this item during bar customization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/customizationLabel
func (g_ GroupTouchBarItem) SetCustomizationLabel(value foundation.foundation.INSString) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCustomizationLabel:"), value)
}


// The compression options that are currently active on the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/effectiveCompressionOptions
func (g_ GroupTouchBarItem) EffectiveCompressionOptions() IUserInterfaceCompressionOptions {
	rv := objc.Send[UserInterfaceCompressionOptions](g_.ID, objc.Sel("effectiveCompressionOptions"))
	return rv
}


// A bar that holds this group’s items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/groupTouchBar
func (g_ GroupTouchBarItem) GroupTouchBar() TouchBar /* not a class type */ {
	rv := objc.Send[TouchBar](g_.ID, objc.Sel("groupTouchBar"))
	return rv
}


// A bar that holds this group’s items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/groupTouchBar
func (g_ GroupTouchBarItem) SetGroupTouchBar(value TouchBar /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGroupTouchBar:"), value)
}


// The user interface direction that controls the layout order of the items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/groupUserInterfaceLayoutDirection
func (g_ GroupTouchBarItem) GroupUserInterfaceLayoutDirection() UserInterfaceLayoutDirection {
	rv := objc.Send[UserInterfaceLayoutDirection](g_.ID, objc.Sel("groupUserInterfaceLayoutDirection"))
	return rv
}


// The user interface direction that controls the layout order of the items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/groupUserInterfaceLayoutDirection
func (g_ GroupTouchBarItem) SetGroupUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGroupUserInterfaceLayoutDirection:"), value)
}


// The preferred width for items in the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/preferredItemWidth
func (g_ GroupTouchBarItem) PreferredItemWidth() float64 {
	rv := objc.Send[float64](g_.ID, objc.Sel("preferredItemWidth"))
	return rv
}


// The preferred width for items in the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/preferredItemWidth
func (g_ GroupTouchBarItem) SetPreferredItemWidth(value float64) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPreferredItemWidth:"), value)
}


// A Boolean value that specifies that items should have equal widths when possible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/prefersEqualWidths
func (g_ GroupTouchBarItem) PrefersEqualWidths() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("prefersEqualWidths"))
	return rv
}


// A Boolean value that specifies that items should have equal widths when possible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/prefersEqualWidths
func (g_ GroupTouchBarItem) SetPrefersEqualWidths(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPrefersEqualWidths:"), value)
}


// The allowed compression options, in the order they should be applied.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/prioritizedCompressionOptions
func (g_ GroupTouchBarItem) PrioritizedCompressionOptions() []UserInterfaceCompressionOptions {
	rv := objc.Send[[]UserInterfaceCompressionOptions](g_.ID, objc.Sel("prioritizedCompressionOptions"))
	return rv
}


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
}







